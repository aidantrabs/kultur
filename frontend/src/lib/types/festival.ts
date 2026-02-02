// Types matching backend database schema from backend/migrations/001_init.sql

export type DateType = 'fixed' | 'lunar' | 'movable';

export type HeritageType = 'african' | 'indian' | 'indigenous' | 'mixed' | 'christian';

export type Region = 'north' | 'south' | 'central' | 'east' | 'west' | 'tobago' | 'nationwide';

export type FestivalType = 'religious' | 'cultural' | 'national' | 'community';

export interface Festival {
    id: string;
    slug: string;
    name: string;
    dateType: DateType;
    usualMonth: string | null;
    date2026Start: string | null; // ISO date string
    date2026End: string | null; // ISO date string
    region: Region;
    heritageType: HeritageType;
    festivalType: FestivalType;
    summary: string;
    story: string | null;
    whatToExpect: string | null;
    howToParticipate: string | null;
    practicalInfo: string | null;
    coverImageUrl: string | null;
    galleryImages: string[];
    videoEmbeds: string[];
    isPublished: boolean;
    createdAt: string;
}

export interface Memory {
    id: string;
    festivalId: string;
    authorName: string | null;
    authorEmail: string | null;
    content: string;
    yearOfMemory: string | null;
    status: 'pending' | 'approved' | 'rejected';
    submittedAt: string;
}

// Filter state for the calendar
export interface FestivalFilters {
    month: number | null; // 0-11 for Jan-Dec, null for all
    region: Region | null;
    heritageType: HeritageType | null;
}

// Heritage color mapping utility
export const heritageColors: Record<HeritageType, string> = {
    african: 'var(--heritage-african)',
    indian: 'var(--heritage-indian)',
    indigenous: 'var(--heritage-indigenous)',
    mixed: 'var(--heritage-mixed)',
    christian: 'var(--heritage-christian)',
};

export const heritageLabels: Record<HeritageType, string> = {
    african: 'African Heritage',
    indian: 'Indian Heritage',
    indigenous: 'Indigenous/First Peoples',
    mixed: 'Mixed Heritage',
    christian: 'Christian',
};

export const regionLabels: Record<Region, string> = {
    north: 'North Trinidad',
    south: 'South Trinidad',
    central: 'Central Trinidad',
    east: 'East Trinidad',
    west: 'West Trinidad',
    tobago: 'Tobago',
    nationwide: 'Nationwide',
};
