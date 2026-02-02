package middleware

import (
    "net/http"
    "sync"
    "time"

    "github.com/labstack/echo/v4"
)

type RateLimiter struct {
    requests map[string][]time.Time
    mu       sync.RWMutex
    limit    int
    window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    rl := &RateLimiter{
        requests: make(map[string][]time.Time),
        limit:    limit,
        window:   window,
    }

    // cleanup old entries every minute
    go rl.cleanup()

    return rl
}

func (rl *RateLimiter) cleanup() {
    ticker := time.NewTicker(time.Minute)
    for range ticker.C {
        rl.mu.Lock()
        now := time.Now()

        for ip, times := range rl.requests {
            var valid []time.Time
        
						for _, t := range times {
                if now.Sub(t) < rl.window {
                    valid = append(valid, t)
                }
            }
            
						if len(valid) == 0 {
                delete(rl.requests, ip)
            } else {
                rl.requests[ip] = valid
            }
        }

        rl.mu.Unlock()
    }
}

func (rl *RateLimiter) isAllowed(ip string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    now := time.Now()
    windowStart := now.Add(-rl.window)

    // filter requests within window
    var valid []time.Time
    for _, t := range rl.requests[ip] {
        if t.After(windowStart) {
            valid = append(valid, t)
        }
    }

    if len(valid) >= rl.limit {
        rl.requests[ip] = valid
        
				return false
    }

    rl.requests[ip] = append(valid, now)
    
		return true
}

func (rl *RateLimiter) Middleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ip := c.RealIP()

            if !rl.isAllowed(ip) {
                return echo.NewHTTPError(http.StatusTooManyRequests, "rate limit exceeded")
            }

            return next(c)
        }
    }
}
