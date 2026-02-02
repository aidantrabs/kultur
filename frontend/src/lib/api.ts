import { config } from './config';
import type { Festival, Memory } from './types/festival';

/**
 * API client for backend endpoints
 * Base URL configured in .env: PUBLIC_API_URL
 */

const API_URL = config.apiUrl;

interface CreateMemoryRequest {
    festivalId: string;
    authorName: string;
    authorEmail: string;
    content: string;
    yearOfMemory: string;
}

interface SubscribeRequest {
    email: string;
}

/**
 * Generic fetch wrapper with error handling
 */
async function apiFetch<T>(endpoint: string, options?: RequestInit): Promise<T> {
    const url = `${API_URL}${endpoint}`;

    console.log('API fetch:', { url, method: options?.method || 'GET' });

    try {
        const response = await fetch(url, {
            ...options,
            headers: {
                'Content-Type': 'application/json',
                ...options?.headers,
            },
        });

        console.log('API response status:', response.status);

        if (!response.ok) {
            const errorText = await response.text();
            console.error('API error response:', errorText);
            throw new Error(`API error: ${response.status} ${response.statusText}`);
        }

        const data = await response.json();
        console.log('API response data:', data);
        return data;
    } catch (error) {
        console.error(`API request failed: ${endpoint}`, error);
        throw error;
    }
}

/**
 * Festival API endpoints
 */
export const festivalsApi = {
    /**
     * Get all festivals
     * GET /api/festivals
     */
    list: async (): Promise<Festival[]> => {
        return apiFetch<Festival[]>('/api/festivals');
    },

    /**
     * Get upcoming festivals (next 30 days)
     * GET /api/festivals/upcoming
     */
    listUpcoming: async (): Promise<Festival[]> => {
        return apiFetch<Festival[]>('/api/festivals/upcoming');
    },

    /**
     * Get single festival by slug
     * GET /api/festivals/:slug
     */
    get: async (slug: string): Promise<Festival> => {
        return apiFetch<Festival>(`/api/festivals/${slug}`);
    },

    /**
     * Get memories for a festival
     * GET /api/festivals/:slug/memories
     */
    getMemories: async (slug: string): Promise<Memory[]> => {
        return apiFetch<Memory[]>(`/api/festivals/${slug}/memories`);
    },
};

/**
 * Memory API endpoints
 */
export const memoriesApi = {
    /**
     * Create a new memory
     * POST /api/memories
     */
    create: async (memory: CreateMemoryRequest): Promise<Memory> => {
        return apiFetch<Memory>('/api/memories', {
            method: 'POST',
            body: JSON.stringify(memory),
        });
    },
};

/**
 * Subscription API endpoints
 */
export const subscriptionsApi = {
    /**
     * Subscribe to newsletter/reminders
     * POST /api/subscribe
     */
    subscribe: async (data: SubscribeRequest): Promise<{ message: string }> => {
        return apiFetch<{ message: string }>('/api/subscribe', {
            method: 'POST',
            body: JSON.stringify(data),
        });
    },

    /**
     * Confirm subscription
     * GET /api/subscribe/confirm/:token
     */
    confirm: async (token: string): Promise<{ message: string }> => {
        return apiFetch<{ message: string }>(`/api/subscribe/confirm/${token}`);
    },

    /**
     * Unsubscribe
     * GET /api/unsubscribe/:token
     */
    unsubscribe: async (token: string): Promise<{ message: string }> => {
        return apiFetch<{ message: string }>(`/api/unsubscribe/${token}`);
    },
};

/**
 * Combined API object
 */
export const api = {
    festivals: festivalsApi,
    memories: memoriesApi,
    subscriptions: subscriptionsApi,
};
