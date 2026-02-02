import { env } from '$env/dynamic/public';

/**
 * Application configuration
 * Toggle between mock data and real API by changing PUBLIC_DATA_SOURCE in .env
 */

export const config = {
	// API base URL
	apiUrl: env.PUBLIC_API_URL || 'http://localhost:8080',

	// Data source: 'mock' or 'api'
	// Change this in .env file: PUBLIC_DATA_SOURCE=api
	dataSource: (env.PUBLIC_DATA_SOURCE || 'mock') as 'mock' | 'api',

	// Helper to check if using API
	get useApi() {
		return this.dataSource === 'api';
	},

	// Helper to check if using mock data
	get useMock() {
		return this.dataSource === 'mock';
	}
};
