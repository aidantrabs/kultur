// Calendar utility functions

export const MONTH_NAMES = [
    'January',
    'February',
    'March',
    'April',
    'May',
    'June',
    'July',
    'August',
    'September',
    'October',
    'November',
    'December',
] as const;

export const MONTH_NAMES_SHORT = [
    'Jan',
    'Feb',
    'Mar',
    'Apr',
    'May',
    'Jun',
    'Jul',
    'Aug',
    'Sep',
    'Oct',
    'Nov',
    'Dec',
] as const;

export const DAY_NAMES = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'] as const;

export interface CalendarDay {
    date: Date;
    day: number;
    isCurrentMonth: boolean;
    isToday: boolean;
}

/**
 * Get all days to display in a calendar month grid (includes padding days from prev/next month)
 */
export function getCalendarDays(year: number, month: number): CalendarDay[] {
    const today = new Date();
    today.setHours(0, 0, 0, 0);

    const firstDayOfMonth = new Date(year, month, 1);
    const lastDayOfMonth = new Date(year, month + 1, 0);

    const days: CalendarDay[] = [];

    // Add days from previous month to fill the first week
    const firstDayWeekday = firstDayOfMonth.getDay(); // 0 = Sunday
    for (let i = firstDayWeekday - 1; i >= 0; i--) {
        const date = new Date(year, month, -i);
        days.push({
            date,
            day: date.getDate(),
            isCurrentMonth: false,
            isToday: date.getTime() === today.getTime(),
        });
    }

    // Add days of current month
    for (let day = 1; day <= lastDayOfMonth.getDate(); day++) {
        const date = new Date(year, month, day);
        days.push({
            date,
            day,
            isCurrentMonth: true,
            isToday: date.getTime() === today.getTime(),
        });
    }

    // Add days from next month to complete the grid (always show 6 weeks = 42 days)
    const remainingDays = 42 - days.length;
    for (let i = 1; i <= remainingDays; i++) {
        const date = new Date(year, month + 1, i);
        days.push({
            date,
            day: i,
            isCurrentMonth: false,
            isToday: date.getTime() === today.getTime(),
        });
    }

    return days;
}

/**
 * Check if a date falls within a date range
 */
export function isDateInRange(
    date: Date,
    startDate: string | null,
    endDate: string | null
): boolean {
    if (!startDate) return false;

    const checkDate = new Date(date);
    checkDate.setHours(0, 0, 0, 0);

    const start = new Date(startDate);
    start.setHours(0, 0, 0, 0);

    const end = endDate ? new Date(endDate) : start;
    end.setHours(0, 0, 0, 0);

    return checkDate >= start && checkDate <= end;
}

/**
 * Format a date range for display
 */
export function formatDateRange(startDate: string | null, endDate: string | null): string {
    if (!startDate) return 'Date TBD';

    const start = new Date(startDate);
    const options: Intl.DateTimeFormatOptions = { month: 'short', day: 'numeric' };

    if (!endDate || startDate === endDate) {
        return start.toLocaleDateString('en-TT', { ...options, year: 'numeric' });
    }

    const end = new Date(endDate);

    // Same month
    if (start.getMonth() === end.getMonth()) {
        return `${start.toLocaleDateString('en-TT', options)} - ${end.getDate()}, ${end.getFullYear()}`;
    }

    return `${start.toLocaleDateString('en-TT', options)} - ${end.toLocaleDateString('en-TT', { ...options, year: 'numeric' })}`;
}

/**
 * Get relative time description (e.g., "in 3 days", "tomorrow")
 */
export function getRelativeTime(dateString: string): string {
    const date = new Date(dateString);
    const today = new Date();
    today.setHours(0, 0, 0, 0);
    date.setHours(0, 0, 0, 0);

    const diffTime = date.getTime() - today.getTime();
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));

    if (diffDays === 0) return 'Today';
    if (diffDays === 1) return 'Tomorrow';
    if (diffDays < 0) return `${Math.abs(diffDays)} days ago`;
    if (diffDays < 7) return `In ${diffDays} days`;
    if (diffDays < 14) return 'Next week';
    if (diffDays < 30) return `In ${Math.floor(diffDays / 7)} weeks`;

    return `In ${Math.floor(diffDays / 30)} month${Math.floor(diffDays / 30) > 1 ? 's' : ''}`;
}
