import { config } from '$lib/config';
import { festivalsApi, memoriesApi } from '$lib/api';
import type { Memory } from '$lib/types/festival';
// import memoriesData from

// Mock memories data for development
const mockMemories: Memory[] = [
    // Carnival memories
    {
        id: 'mem-1',
        festivalId: 'carnival',
        authorName: 'Michelle R.',
        authorEmail: null,
        content: `My first J'Ouvert was life-changing. I was nervous about going alone, but within minutes strangers became family. The mud, the paint, the music — it's impossible to explain until you've felt it. At 4am, covered head to toe in chocolate, dancing with hundreds of people as the sun rose... that's when I understood what it means to be Trini.`,
        yearOfMemory: '2019',
        status: 'approved',
        submittedAt: '2025-11-15T10:30:00Z',
    },
    {
        id: 'mem-2',
        festivalId: 'carnival',
        authorName: 'Jason T.',
        authorEmail: null,
        content: `I'm from the diaspora (grew up in Brooklyn) and finally came home for Carnival in 2023. Playing mas with a band, crossing the stage on Carnival Tuesday — I cried. My grandmother used to tell me stories about Carnival and now I was living it. The energy is unmatched anywhere in the world.`,
        yearOfMemory: '2023',
        status: 'approved',
        submittedAt: '2025-10-22T14:15:00Z',
    },

    // Hosay memories
    {
        id: 'mem-3',
        festivalId: 'hosay',
        authorName: 'Anil M.',
        authorEmail: null,
        content: `My family has been building tadjahs for three generations. People don't understand — this isn't a celebration, it's a commemoration. When you see thousands of people lining the streets of St. James in solemn respect, watching these 30-foot structures pass by to the rhythm of tassa drums, you feel the weight of history.`,
        yearOfMemory: '2022',
        status: 'approved',
        submittedAt: '2025-09-18T08:45:00Z',
    },
    {
        id: 'mem-4',
        festivalId: 'hosay',
        authorName: 'Priya S.',
        authorEmail: null,
        content: `I'm Hindu but I've attended Hosay since childhood. That's Trinidad — we share each other's traditions. Learning about the 1884 massacre as an adult changed how I see this observance. It's not just about Hussein and Hassan; it's about our own history of resistance and survival.`,
        yearOfMemory: '2024',
        status: 'approved',
        submittedAt: '2025-12-01T16:20:00Z',
    },

    // Divali memories
    {
        id: 'mem-5',
        festivalId: 'divali',
        authorName: 'Kamla P.',
        authorEmail: null,
        content: `Divali Nagar is something every Trini should experience at least once. The scale is incredible — hundreds of deyas, cultural performances every night, the food! I take my children every year so they stay connected to their roots. When the whole site lights up at night, it's magical.`,
        yearOfMemory: '2024',
        status: 'approved',
        submittedAt: '2025-11-20T19:00:00Z',
    },

    // Santa Rosa Festival memories
    {
        id: 'mem-6',
        festivalId: 'santa-rosa-festival',
        authorName: 'Ricardo L.',
        authorEmail: null,
        content: `I'm a member of the Santa Rosa First Peoples Community. Our festival is about keeping traditions alive that are 8,000 years old. When visitors come with respect and genuine curiosity, it means the world to us. We're still here, and our culture matters.`,
        yearOfMemory: '2024',
        status: 'approved',
        submittedAt: '2025-08-30T11:00:00Z',
    },

    // First Peoples Heritage Week
    {
        id: 'mem-7',
        festivalId: 'first-peoples-heritage-week',
        authorName: 'Anonymous',
        authorEmail: null,
        content:
            'I had no idea about Amerindian heritage in Trinidad until I attended a Heritage Week event in Arima. The smoke ceremony was powerful — I could feel the connection to something much older than colonial history. Every Trini should learn about this part of who we are.',
        yearOfMemory: '2023',
        status: 'approved',
        submittedAt: '2025-10-15T13:30:00Z',
    },

    // Emancipation Day
    {
        id: 'mem-8',
        festivalId: 'emancipation-day',
        authorName: 'Keisha W.',
        authorEmail: null,
        content: `Walking in the Emancipation Day parade with my grandmother, both of us in African dress, was powerful. She told me stories of her own grandmother who was born just one generation after slavery ended. August 1st isn't just a holiday — it's a reminder of what our ancestors survived and overcame.`,
        yearOfMemory: '2024',
        status: 'approved',
        submittedAt: '2025-08-05T10:15:00Z',
    },

    // Phagwa
    {
        id: 'mem-9',
        festivalId: 'phagwa',
        authorName: 'Davendra R.',
        authorEmail: null,
        content: `Phagwa in Chaguanas is pure joy. Everyone — Hindu, Christian, Muslim, it doesn't matter — comes out to throw colors. By the end you can't tell who is who because everyone is covered in pink, green, purple, yellow. That's the point. That's Trinidad.`,
        yearOfMemory: '2024',
        status: 'approved',
        submittedAt: '2025-03-20T15:45:00Z',
    },

    // Spiritual Baptist Liberation Day
    {
        id: 'mem-10',
        festivalId: 'spiritual-baptist-liberation-day',
        authorName: 'Sister Grace',
        authorEmail: null,
        content: `My church was underground for years. My mother told me about worshipping in secret, about police raids. March 30th means everything to us — it's when we were finally free to worship openly. Every Liberation Day, I wear my white and I thank God for those who fought for our freedom.`,
        yearOfMemory: '2025',
        status: 'approved',
        submittedAt: '2025-04-02T09:00:00Z',
    },
];

/**
 * Get memories for a specific festival - uses API or mock data based on config
 */
export async function getMemoriesByFestivalSlug(slug: string): Promise<Memory[]> {
  if (config.useApi) {
    try {
      return await festivalsApi.getMemories(slug);
    } catch (error) {
      console.error(`Failed to fetch memories for ${slug} from API, falling back to mock data`, error);
      // For mock fallback, we need to find the festival ID by slug
      const { festivals } = await import('./festivals');
      const festival = festivals.find(f => f.slug === slug);
      if (!festival) return [];
      return mockMemories.filter(
        (m) => m.festivalId === festival.id && m.status === 'approved'
      );
    }
  }
  
  // Mock data mode - find festival by slug
  const { festivals } = await import('./festivals');
  const festival = festivals.find(f => f.slug === slug);
  if (!festival) return [];
  
  return mockMemories.filter(
    (m) => m.festivalId === slug && m.status === 'approved'
  );
}

/**
 * Submit a new memory - uses API or simulates submission based on config
 */
export async function submitMemory(memory: {
    festivalId: string;
    authorName: string;
    authorEmail: string;
    content: string;
    yearOfMemory: string;
}): Promise<Memory> {
    if (config.useApi) {
        return await memoriesApi.create(memory);
    }

    // Simulate API call for mock mode
    await new Promise((resolve) => setTimeout(resolve, 1000));

    return {
        id: `mem-${Date.now()}`,
        festivalId: memory.festivalId,
        authorName: memory.authorName || null,
        authorEmail: memory.authorEmail || null,
        content: memory.content,
        yearOfMemory: memory.yearOfMemory || null,
        status: 'pending',
        submittedAt: new Date().toISOString(),
    };
}

/**
 * Synchronous access to memories (mock data only)
 */
export const memories = mockMemories;

// Get memories for a specific festival (synchronous, mock only)
export function getMemoriesByFestival(festivalId: string): Memory[] {
    return mockMemories.filter((m) => m.festivalId === festivalId && m.status === 'approved');
}

// Get recent memories across all festivals (synchronous, mock only)
export function getRecentMemories(limit = 5): Memory[] {
    return [...mockMemories]
        .filter((m) => m.status === 'approved')
        .sort((a, b) => new Date(b.submittedAt).getTime() - new Date(a.submittedAt).getTime())
        .slice(0, limit);
}
