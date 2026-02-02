package email

import (
    "fmt"
    "html/template"
    "time"

    "github.com/resend/resend-go/v2"
)

type Service struct {
    client    *resend.Client
    fromEmail string
    fromName  string
    baseURL   string
}

type Config struct {
    APIKey    string
    FromEmail string
    FromName  string
    BaseURL   string
}

func NewService(cfg Config) *Service {
    if cfg.APIKey == "" {
        return &Service{} // disabled mode
    }

    fromName := cfg.FromName
    if fromName == "" {
        fromName = "KULTUR"
    }

    fromEmail := cfg.FromEmail
    if fromEmail == "" {
        fromEmail = "noreply@kulturtt.com"
    }

    return &Service{
        client:    resend.NewClient(cfg.APIKey),
        fromEmail: fromEmail,
        fromName:  fromName,
        baseURL:   cfg.BaseURL,
    }
}

func (s *Service) IsEnabled() bool {
    return s.client != nil
}

func (s *Service) from() string {
    return fmt.Sprintf("%s <%s>", s.fromName, s.fromEmail)
}

func (s *Service) SendConfirmation(toEmail, token string) error {
    if !s.IsEnabled() {
        return nil
    }

    confirmURL := fmt.Sprintf("%s/api/subscribe/confirm/%s", s.baseURL, token)

    body := `<p style="margin: 0 0 16px 0;">
        Thanks for subscribing to <strong>KULTUR</strong>! We're excited to share Trinidad & Tobago's rich cultural heritage with you.
    </p>
    <p style="margin: 0 0 16px 0;">
        Please confirm your email address to start receiving updates about upcoming festivals, first-timer guides, and cultural events.
    </p>
    <p style="margin: 0; padding: 16px; background-color: #fef3c7; border-radius: 8px; border-left: 4px solid ` + ColorGold + `;">
        <strong style="color: #92400e;">Note:</strong> This link expires in 24 hours. If you didn't subscribe, you can safely ignore this email.
    </p>`

    html, err := RenderTemplate(TemplateData{
        PreviewText: "Confirm your email to discover Trinidad's cultural festivals",
        Heading:     "Confirm Your Email",
        Body:        template.HTML(body),
        ButtonText:  "Confirm Email Address",
        ButtonURL:   confirmURL,
        Year:        time.Now().Year(),
    })
    if err != nil {
        return fmt.Errorf("failed to render template: %w", err)
    }

    _, err = s.client.Emails.Send(&resend.SendEmailRequest{
        From:    s.from(),
        To:      []string{toEmail},
        Subject: "Confirm your subscription to KULTUR",
        Html:    html,
    })

    return err
}

func (s *Service) SendWelcome(toEmail, unsubscribeToken string) error {
    if !s.IsEnabled() {
        return nil
    }

    unsubURL := fmt.Sprintf("%s/api/unsubscribe/%s", s.baseURL, unsubscribeToken)
    calendarURL := fmt.Sprintf("%s/festivals", s.baseURL)

    body := fmt.Sprintf(`<p style="margin: 0 0 16px 0;">
        Welcome to the KULTUR community! Your subscription is now confirmed.
    </p>

    <p style="margin: 0 0 24px 0;">
        You'll receive updates about:
    </p>

    <table role="presentation" cellspacing="0" cellpadding="0" border="0" style="margin: 0 0 24px 0;">
        <tr>
            <td style="padding: 12px 16px; background-color: #fef2f2; border-radius: 8px; margin-bottom: 8px;">
                <table role="presentation" cellspacing="0" cellpadding="0" border="0">
                    <tr>
                        <td style="padding-right: 12px; vertical-align: top;">
                            <span style="font-size: 20px;">🎭</span>
                        </td>
                        <td>
                            <strong style="color: %s;">Festival Announcements</strong><br>
                            <span style="font-size: 14px; color: #6b7280;">Know when festivals are happening</span>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
        <tr><td style="height: 8px;"></td></tr>
        <tr>
            <td style="padding: 12px 16px; background-color: #fef9c3; border-radius: 8px; margin-bottom: 8px;">
                <table role="presentation" cellspacing="0" cellpadding="0" border="0">
                    <tr>
                        <td style="padding-right: 12px; vertical-align: top;">
                            <span style="font-size: 20px;">📖</span>
                        </td>
                        <td>
                            <strong style="color: %s;">First-Timer Guides</strong><br>
                            <span style="font-size: 14px; color: #6b7280;">What to expect and how to participate</span>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
        <tr><td style="height: 8px;"></td></tr>
        <tr>
            <td style="padding: 12px 16px; background-color: #ecfdf5; border-radius: 8px;">
                <table role="presentation" cellspacing="0" cellpadding="0" border="0">
                    <tr>
                        <td style="padding-right: 12px; vertical-align: top;">
                            <span style="font-size: 20px;">🗓️</span>
                        </td>
                        <td>
                            <strong style="color: %s;">Weekly Digest</strong><br>
                            <span style="font-size: 14px; color: #6b7280;">Cultural events happening this week</span>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>

    <p style="margin: 0;">
        Ready to explore? Check out our <a href="%s" style="color: %s; font-weight: 600;">festival calendar</a> to see what's coming up.
    </p>`, ColorRed, ColorGold, "#059669", calendarURL, ColorRed)

    html, err := RenderTemplate(TemplateData{
        PreviewText:    "Welcome to KULTUR! Start exploring Trinidad's festivals",
        Heading:        "You're All Set! 🎉",
        Body:           template.HTML(body),
        ButtonText:     "Explore Festivals",
        ButtonURL:      calendarURL,
        FooterText:     "Thank you for joining us in celebrating Trinidad & Tobago's cultural heritage.",
        UnsubscribeURL: unsubURL,
        Year:           time.Now().Year(),
    })
    if err != nil {
        return fmt.Errorf("failed to render template: %w", err)
    }

    _, err = s.client.Emails.Send(&resend.SendEmailRequest{
        From:    s.from(),
        To:      []string{toEmail},
        Subject: "Welcome to KULTUR! 🎭",
        Html:    html,
    })

    return err
}

func (s *Service) SendFestivalReminder(toEmail, festivalName, festivalSlug, unsubscribeToken string, daysUntil int) error {
    if !s.IsEnabled() {
        return nil
    }

    unsubURL := fmt.Sprintf("%s/api/unsubscribe/%s", s.baseURL, unsubscribeToken)
    festivalURL := fmt.Sprintf("%s/festivals/%s", s.baseURL, festivalSlug)

    var timeText string
    switch daysUntil {
    case 0:
        timeText = "Today"
    case 1:
        timeText = "Tomorrow"
    case 7:
        timeText = "In 1 week"
    default:
        timeText = fmt.Sprintf("In %d days", daysUntil)
    }

    body := fmt.Sprintf(`<p style="margin: 0 0 16px 0;">
        <strong>%s</strong> is coming up <strong>%s</strong>!
    </p>

    <p style="margin: 0 0 24px 0;">
        Don't miss this opportunity to experience one of Trinidad & Tobago's cultural treasures. Check out our first-timer's guide to know what to expect.
    </p>

    <div style="padding: 20px; background-color: #f9fafb; border-radius: 8px; border: 1px solid %s;">
        <p style="margin: 0; font-size: 14px; color: #6b7280;">
            📍 View the complete guide including what to wear, how to participate, and practical tips.
        </p>
    </div>`, festivalName, timeText, ColorBorder)

    html, err := RenderTemplate(TemplateData{
        PreviewText:    fmt.Sprintf("%s is %s - Don't miss it!", festivalName, timeText),
        Heading:        fmt.Sprintf("%s is %s!", festivalName, timeText),
        Body:           template.HTML(body),
        ButtonText:     "View Festival Guide",
        ButtonURL:      festivalURL,
        UnsubscribeURL: unsubURL,
        Year:           time.Now().Year(),
    })
    if err != nil {
        return fmt.Errorf("failed to render template: %w", err)
    }

    _, err = s.client.Emails.Send(&resend.SendEmailRequest{
        From:    s.from(),
        To:      []string{toEmail},
        Subject: fmt.Sprintf("🎭 %s is %s!", festivalName, timeText),
        Html:    html,
    })

    return err
}
