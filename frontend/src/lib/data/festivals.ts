import { config } from '$lib/config';
import { festivalsApi } from '$lib/api';
import type { Festival } from '$lib/types/festival';

// Mock data matching backend schema - 10 festivals from PLANNING.md
const mockFestivals: Festival[] = [
    {
        id: '550e8400-e29b-41d4-a716-446655440001',
        slug: 'hosay',
        name: 'Hosay',
        dateType: 'lunar',
        usualMonth: 'July/August',
        date2026Start: '2026-07-25',
        date2026End: '2026-07-28',
        region: 'west',
        heritageType: 'indian',
        festivalType: 'religious',
        summary:
            'A solemn Islamic observance commemorating the martyrdom of Hussein and Hassan at the Battle of Karbala.',
        story: `Hosay was brought to Trinidad by Indian indentured laborers starting in 1845. It commemorates the martyrdom of Imam Hussein and his brother Hassan, grandsons of the Prophet Muhammad, at the Battle of Karbala in 680 AD.

The observance is marked by the building of elaborate tadjahs (ornate tomb replicas) that can reach up to 30 feet tall. These are paraded through the streets accompanied by the hypnotic rhythms of tassa drumming.

**The 1884 Hosay Massacre**: On October 30, 1884, British colonial authorities opened fire on a peaceful Hosay procession in San Fernando. 22 people were killed and over 120 wounded. It remains one of the bloodiest events in Trinidad's colonial history, yet most Trinis have never heard of it.`,
        whatToExpect: `- Massive, colorful tadjahs (tomb replicas) carried through streets
- Hypnotic tassa drumming that you'll feel in your chest
- Processions over multiple nights leading to the final day
- Moon dancers in elaborate costumes
- Solemn atmosphere mixed with communal gathering`,
        howToParticipate: `**Important**: Hosay is NOT a "festival" - it's a solemn religious observance. Approach with respect.

- Dress modestly (cover shoulders and knees)
- Photography is generally welcomed but ask first
- Non-Muslims are welcome to observe respectfully
- Do not touch the tadjahs
- Listen more than you speak - let practitioners share if they wish`,
        practicalInfo: `**Best viewing**: St. James, Port of Spain (largest observance)
**Also in**: Cedros, Couva
**Timing**: Processions happen after sunset
**Getting there**: Parking limited in St. James; consider taxi
**Food**: Street vendors sell doubles and other local food nearby`,
        coverImageUrl:
            'https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg',
        galleryImages: [],
        videoEmbeds: [],
        isPublished: true,
        createdAt: '2026-01-15T00:00:00Z',
    },
    {
        id: '550e8400-e29b-41d4-a716-446655440002',
        slug: 'santa-rosa-festival',
        name: 'Santa Rosa Festival',
        dateType: 'fixed',
        usualMonth: 'August',
        date2026Start: '2026-08-23',
        date2026End: '2026-08-23',
        region: 'north',
        heritageType: 'indigenous',
        festivalType: 'cultural',
        summary:
            'The annual celebration of the Santa Rosa First Peoples Community - the last organized Indigenous group in Trinidad & Tobago.',
        story: `The Santa Rosa First Peoples Community in Arima represents the last organized Indigenous group in Trinidad & Tobago. With approximately 400 registered members, they have been preserving their traditions for over 50 years.

Archaeological evidence shows 8,000 years of continuous Amerindian presence on these islands. The First Peoples' cultural contributions live on in everyday Trinidadian life: cassava bread, the barbecue, parang music, and place names like Arima, Tunapuna, and Chaguanas.

The Santa Rosa Festival honors their patron saint while celebrating and preserving Indigenous culture for future generations.`,
        whatToExpect: `- Traditional Smoke Ceremony and Water Ritual
- Indigenous craft demonstrations
- Traditional foods including cassava bread
- Cultural performances and storytelling
- Exhibition of artifacts and heritage items`,
        howToParticipate: `- This is a community celebration - come with an open heart
- Ask before photographing individuals or ceremonies
- Support local artisans by purchasing crafts
- Listen to elders who share their knowledge
- Children are welcome`,
        practicalInfo: `**Location**: Santa Rosa RC Church grounds, Arima
**Timing**: Usually a Sunday in late August
**Getting there**: Arima is accessible by bus from Port of Spain
**What to bring**: Cash for crafts and food, water, sunscreen`,
        coverImageUrl:
            'https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg',
        galleryImages: [],
        videoEmbeds: [],
        isPublished: true,
        createdAt: '2026-01-15T00:00:00Z',
    },
    {
        id: '550e8400-e29b-41d4-a716-446655440003',
        slug: 'first-peoples-heritage-week',
        name: 'First Peoples Heritage Week',
        dateType: 'fixed',
        usualMonth: 'October',
        date2026Start: '2026-10-12',
        date2026End: '2026-10-18',
        region: 'north',
        heritageType: 'indigenous',
        festivalType: 'cultural',
        summary:
            'A week-long celebration of Indigenous heritage, featuring ceremonies, exhibitions, and educational events.',
        story: `First Peoples Heritage Week was established to bring awareness to Trinidad & Tobago's Indigenous roots. While Columbus arrived in 1498, the islands had been home to Amerindian peoples for millennia.

The week includes the significant Smoke Ceremony, traditionally used for purification and connecting with ancestors. It's a time for the Santa Rosa First Peoples Community to share their living heritage with the broader national community.`,
        whatToExpect: `- Opening Smoke Ceremony
- Heritage exhibitions at NALIS and community spaces
- School visits and educational programs
- Documentary screenings
- Traditional craft workshops
- Closing ceremonies`,
        howToParticipate: `- Attend with genuine curiosity and respect
- Many events are free and open to the public
- Schools can arrange special visits
- Photography policies vary by event - always ask`,
        practicalInfo: `**Primary location**: Arima and Port of Spain
**Check**: Santa Rosa First Peoples Community social media for schedule
**Best for**: Families, students, anyone interested in T&T history`,
        coverImageUrl:
            'https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg',
        galleryImages: [],
        videoEmbeds: [],
        isPublished: true,
        createdAt: '2026-01-15T00:00:00Z',
    },
    {
        id: '550e8400-e29b-41d4-a716-446655440004',
        slug: 'divali',
        name: 'Divali',
        dateType: 'lunar',
        usualMonth: 'October/November',
        date2026Start: '2026-11-08',
        date2026End: '2026-11-08',
        region: 'nationwide',
        heritageType: 'indian',
        festivalType: 'religious',
        summary:
            'The Hindu Festival of Lights, celebrated as a national holiday with the largest observance in the English-speaking Caribbean.',
        story: `Divali celebrates the victory of light over darkness, good over evil, and knowledge over ignorance. It marks Lord Rama's return to Ayodhya after 14 years of exile and his victory over the demon king Ravana.

Trinidad & Tobago's Divali celebration is the largest in the English-speaking Caribbean. Divali Nagar in Chaguanas transforms into a spectacular village of lights, culture, and devotion, attracting hundreds of thousands of visitors over nine nights.`,
        whatToExpect: `- Thousands of deyas (clay lamps) illuminating homes and temples
- Divali Nagar - nine nights of cultural performances, food, and exhibitions
- Lakshmi Puja (prayers to the goddess of prosperity)
- Spectacular fireworks
- Elaborate rangoli (floor decorations) at entrances
- Delicious Indian sweets (mithai)`,
        howToParticipate: `- Visit Divali Nagar any of the nine nights
- Drive through Hindu communities at night to see illuminated homes
- Respect that this is a religious occasion for many
- Try the food - parsad, kurma, and other sweets
- Dress comfortably for walking`,
        practicalInfo: `**Divali Nagar**: NCIC grounds, Chaguanas - runs for 9 nights
**Traffic**: Expect heavy traffic; consider arriving early
**Best night to visit**: The main Divali night (November 8th in 2026)
**Cost**: Entrance to Nagar is free; budget for food and crafts`,
        coverImageUrl:
            'https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg',
        galleryImages: [],
        videoEmbeds: [],
        isPublished: true,
        createdAt: '2026-01-15T00:00:00Z',
    },
    {
        id: '550e8400-e29b-41d4-a716-446655440005',
        slug: 'phagwa',
        name: 'Phagwa (Holi)',
        dateType: 'lunar',
        usualMonth: 'March',
        date2026Start: '2026-03-10',
        date2026End: '2026-03-10',
        region: 'nationwide',
        heritageType: 'indian',
        festivalType: 'religious',
        summary:
            'The vibrant Hindu spring festival of colors, celebrating the victory of good over evil and the arrival of spring.',
        story: `Phagwa, known globally as Holi, celebrates the legend of Prahlad and Holika - a story of devotion triumphing over evil. It marks the arrival of spring and the end of winter.

In Trinidad, Phagwa has become one of the most inclusive celebrations, with people of all backgrounds coming together to throw abeer (colored powder) and celebrate joy and renewal.`,
        whatToExpect: `- Clouds of vibrant colored powder (abeer) everywhere
- Chowtals (traditional folk songs) being sung
- Water play - you WILL get wet
- Joyous atmosphere with music and dancing
- Traditional foods and drinks like mauby`,
        howToParticipate: `- Wear WHITE clothes you don't mind staining permanently
- Bring a change of clothes and towel
- Protect your phone in a waterproof case
- Join in! Throwing colors is encouraged
- Stay hydrated in the sun`,
        practicalInfo: `**Major celebrations**: Tunapuna, Aranguez, Chaguanas
**Timing**: Usually mid-morning to afternoon
**What to bring**: Old clothes, sandals, waterproof phone case
**Warning**: Colors are semi-permanent - oil your hair and skin beforehand`,
        coverImageUrl:
            'https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg',
        galleryImages: [],
        videoEmbeds: [],
        isPublished: true,
        createdAt: '2026-01-15T00:00:00Z',
    },
    {
        id: '550e8400-e29b-41d4-a716-446655440006',
        slug: 'emancipation-day',
        name: 'Emancipation Day',
        dateType: 'fixed',
        usualMonth: 'August',
        date2026Start: '2026-08-01',
        date2026End: '2026-08-01',
        region: 'nationwide',
        heritageType: 'african',
        festivalType: 'national',
        summary:
            'National holiday commemorating the abolition of slavery in the British Empire in 1834.',
        story: `On August 1, 1834, slavery was officially abolished in the British Empire, though full freedom didn't come until 1838 after a period of "apprenticeship." Emancipation Day honors the ancestors who endured and resisted slavery, and celebrates African cultural heritage in Trinidad & Tobago.

The day is marked by cultural celebrations, African-inspired dress, and remembrance of the journey from bondage to freedom.`,
        whatToExpect: `- Cultural celebrations across the country
- African-inspired fashion and headwraps
- Traditional African drumming and dance
- Educational talks and exhibitions
- Dawn ceremonies and libations
- Community gatherings and feasts`,
        howToParticipate: `- Wear African-inspired clothing or colors (red, black, green)
- Attend events at the Lidj Yasu Omowale Emancipation Village
- Learn about African contributions to T&T culture
- Support Black-owned businesses
- Reflect on the journey and ongoing work for justice`,
        practicalInfo: `**Public holiday**: Banks and government offices closed
**Main events**: Port of Spain, San Fernando
**Emancipation Village**: Queen's Park Savannah, Port of Spain
**Dress code**: African attire encouraged but not required`,
        coverImageUrl:
            'https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg',
        galleryImages: [],
        videoEmbeds: [],
        isPublished: true,
        createdAt: '2026-01-15T00:00:00Z',
    },
    {
        id: '550e8400-e29b-41d4-a716-446655440007',
        slug: 'spiritual-baptist-liberation-day',
        name: 'Spiritual Baptist Liberation Day',
        dateType: 'fixed',
        usualMonth: 'March',
        date2026Start: '2026-03-30',
        date2026End: '2026-03-30',
        region: 'nationwide',
        heritageType: 'african',
        festivalType: 'religious',
        summary:
            'Commemorating the 1951 repeal of the Shouter Prohibition Ordinance that had banned the Spiritual Baptist faith for 34 years.',
        story: `The Spiritual Baptist faith blends African spiritual traditions with Christianity, featuring distinctive practices like shouting (Spirit-led vocalizations), bell-ringing, and mourning rituals.

In 1917, the colonial government passed the Shouter Prohibition Ordinance, banning the faith and criminalizing its practitioners. For 34 years, Spiritual Baptists practiced in secret, risking arrest and persecution.

On March 30, 1951, the ordinance was repealed. Trinidad & Tobago became the first country in the world to declare a public holiday honoring a religious minority's liberation from persecution.`,
        whatToExpect: `- Church services featuring distinctive Spiritual Baptist worship
- Bell-ringing and spiritual songs
- Colorful robes and headties
- Public ceremonies and thanksgiving
- Educational programs about the faith's history`,
        howToParticipate: `- Attend a Spiritual Baptist church service (visitors welcome)
- Dress modestly and respectfully
- Observe worship practices without judgment
- Listen to practitioners share their faith journey
- Acknowledge the history of persecution and resilience`,
        practicalInfo: `**Public holiday**: Yes
**Services**: Spiritual Baptist churches throughout T&T
**What to wear**: Modest, respectful attire; women may want to bring a head covering`,
        coverImageUrl:
            'https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg',
        galleryImages: [],
        videoEmbeds: [],
        isPublished: true,
        createdAt: '2026-01-15T00:00:00Z',
    },
    {
        id: '550e8400-e29b-41d4-a716-446655440008',
        slug: 'ramleela',
        name: 'Ramleela',
        dateType: 'lunar',
        usualMonth: 'September/October',
        date2026Start: '2026-10-02',
        date2026End: '2026-10-11',
        region: 'central',
        heritageType: 'indian',
        festivalType: 'religious',
        summary:
            'Open-air theatrical performances of the Hindu epic Ramayana, held at over 35 sites across Trinidad.',
        story: `Ramleela brings the ancient Hindu epic Ramayana to life through elaborate open-air theatrical performances. The epic tells the story of Lord Rama, his wife Sita, and his battle against the demon king Ravana.

Trinidad's Ramleela tradition dates back to the arrival of Indian indentured laborers in the 1800s. Today, over 35 Ramleela sites operate across the country, with performances spanning nine nights culminating in the burning of Ravana's effigy.`,
        whatToExpect: `- Epic theatrical performances under the stars
- Elaborate costumes and makeup
- Live music and devotional singing
- The dramatic burning of massive Ravana effigies
- Traditional Indian food at venue stalls
- Multi-generational family atmosphere`,
        howToParticipate: `- Performances are free and open to all
- Arrive early to get a good viewing spot
- Bring a blanket or low chair to sit on
- Stay for the dramatic climax - Ravana's burning
- Try the food from vendors`,
        practicalInfo: `**Major sites**: Dow Village, Palmiste, Felicity
**Timing**: Nightly performances during Navaratri (nine nights)
**Duration**: Each night's performance runs 2-4 hours
**Best night**: Final night for Ravana's burning (Dussehra)`,
        coverImageUrl:
            'https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg',
        galleryImages: [],
        videoEmbeds: [],
        isPublished: true,
        createdAt: '2026-01-15T00:00:00Z',
    },
    {
        id: '550e8400-e29b-41d4-a716-446655440009',
        slug: 'parang',
        name: 'Parang Season',
        dateType: 'fixed',
        usualMonth: 'October-December',
        date2026Start: '2026-10-01',
        date2026End: '2026-12-31',
        region: 'nationwide',
        heritageType: 'mixed',
        festivalType: 'cultural',
        summary:
            'The traditional Christmas music of Trinidad, blending Spanish, African, and Indigenous influences into joyous celebration.',
        story: `Parang is Trinidad's unique Christmas music tradition, featuring Spanish-influenced songs performed with cuatro, maracas, and box bass. The name comes from "parranda" (spree or fête).

Originally brought by Spanish missionaries and Venezuelan settlers, parang was traditionally performed by roving groups who would go house-to-house during the Christmas season. Today, it's evolved to include "soca parang" - a fusion with calypso rhythms.

The tradition reflects Trinidad's multicultural heritage, blending Spanish lyrics, African rhythms, and Indigenous instruments.`,
        whatToExpect: `- Live parang bands performing at venues across the country
- Traditional instruments: cuatro, maracas, box bass, guitar
- Call-and-response singing
- Dancing and festive atmosphere
- Traditional pastelles, sorrel, and ponche de crème`,
        howToParticipate: `- Attend a parang lime (gathering) at a local venue
- Learn a few common parang songs to sing along
- Try the traditional Christmas foods and drinks
- Support local parang groups
- Host your own parang lime!`,
        practicalInfo: `**Season**: October through December
**Major events**: National Parang Finals, community parang competitions
**Best venues**: Paramin, Lopinot, community centres nationwide
**Food pairings**: Pastelles, black cake, sorrel, ponche de crème`,
        coverImageUrl:
            'https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg',
        galleryImages: [],
        videoEmbeds: [],
        isPublished: true,
        createdAt: '2026-01-15T00:00:00Z',
    },
    {
        id: '550e8400-e29b-41d4-a716-446655440010',
        slug: 'carnival',
        name: 'Trinidad Carnival',
        dateType: 'movable',
        usualMonth: 'February/March',
        date2026Start: '2026-02-16',
        date2026End: '2026-02-17',
        region: 'nationwide',
        heritageType: 'mixed',
        festivalType: 'cultural',
        summary:
            "The Greatest Show on Earth - Trinidad's world-famous celebration of music, mas, and freedom.",
        story: `Trinidad Carnival is the Caribbean's largest and most famous festival, attracting hundreds of thousands of visitors annually. It originated in the late 18th century when French planters brought masquerade traditions to the island.

After Emancipation, formerly enslaved Africans transformed Carnival into a celebration of freedom, incorporating African traditions, stick fighting (kalinda), and eventually the steelpan - the only acoustic musical instrument invented in the 20th century.

Today's Carnival features massive costume bands, soca music, J'ouvert (the pre-dawn paint and mud celebration), and two days of parading through the streets.`,
        whatToExpect: `- J'ouvert (Monday 4am): Paint, mud, and pure abandon
- Panorama: Steelband competition finals
- Mas Monday & Tuesday: Costume bands on the road
- Soca music blasting from massive sound trucks
- Fetes and parties for weeks leading up`,
        howToParticipate: `- Register with a mas band months in advance for costumes
- J'ouvert requires old clothes you can throw away
- Stay hydrated - you'll be in the sun for hours
- Pace yourself - it's a marathon, not a sprint
- Book accommodation 6+ months ahead`,
        practicalInfo: `**Main days**: Monday & Tuesday before Ash Wednesday
**J'ouvert**: Monday 4am start
**Costume price**: $500-$3000+ USD depending on band
**Must-see**: Panorama semis and finals (Savannah)`,
        coverImageUrl:
            'https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg',
        galleryImages: [],
        videoEmbeds: [],
        isPublished: true,
        createdAt: '2026-01-15T00:00:00Z',
    },
];

/**
 * Get all festivals - uses API or mock data based on config
 */
export async function getAllFestivals(): Promise<Festival[]> {
    if (config.useApi) {
        try {
            return await festivalsApi.list();
        } catch (error) {
            console.error('Failed to fetch festivals from API, falling back to mock data', error);
            return mockFestivals;
        }
    }
    return mockFestivals;
}

/**
 * Get festival by slug - uses API or mock data based on config
 */
export async function getFestivalBySlug(slug: string): Promise<Festival | undefined> {
    if (config.useApi) {
        try {
            return await festivalsApi.get(slug);
        } catch (error) {
            console.error(
                `Failed to fetch festival ${slug} from API, falling back to mock data`,
                error
            );
            return mockFestivals.find((f) => f.slug === slug);
        }
    }
    return mockFestivals.find((f) => f.slug === slug);
}

/**
 * Synchronous access to festivals (mock data only)
 * Use this for components that need immediate access
 */
export const festivals = mockFestivals;

// Helper function to get festivals for a specific month (0-indexed)
export function getFestivalsByMonth(month: number): Festival[] {
    return mockFestivals.filter((f) => {
        if (!f.date2026Start) return false;
        const startMonth = new Date(f.date2026Start).getMonth();
        const endMonth = f.date2026End ? new Date(f.date2026End).getMonth() : startMonth;
        return month >= startMonth && month <= endMonth;
    });
}

// Helper function to get upcoming festivals within next N days
export function getUpcomingFestivals(days = 30): Festival[] {
    const today = new Date();
    const futureDate = new Date(today);
    futureDate.setDate(today.getDate() + days);

    return mockFestivals
        .filter((f) => {
            if (!f.date2026Start) return false;
            const startDate = new Date(f.date2026Start);
            return startDate >= today && startDate <= futureDate;
        })
        .sort((a, b) => {
            const dateA = new Date(a.date2026Start ?? '');
            const dateB = new Date(b.date2026Start ?? '');
            return dateA.getTime() - dateB.getTime();
        });
}

/**
 * Get upcoming festivals from API
 */
export async function getUpcomingFestivalsApi(): Promise<Festival[]> {
    if (config.useApi) {
        try {
            return await festivalsApi.listUpcoming();
        } catch (error) {
            console.error(
                'Failed to fetch upcoming festivals from API, falling back to mock',
                error
            );
            return getUpcomingFestivals();
        }
    }
    return getUpcomingFestivals();
}

// For demo purposes - get "This Week" festivals relative to a base date
// Using February 2026 as base to show Carnival in "This Week"
export function getThisWeekFestivals(baseDate: Date = new Date()): Festival[] {
    const weekEnd = new Date(baseDate);
    weekEnd.setDate(baseDate.getDate() + 14); // Show next 2 weeks

    return mockFestivals
        .filter((f) => {
            if (!f.date2026Start) return false;
            const startDate = new Date(f.date2026Start);
            return startDate >= baseDate && startDate <= weekEnd;
        })
        .sort((a, b) => {
            const dateA = new Date(a.date2026Start ?? '');
            const dateB = new Date(b.date2026Start ?? '');
            return dateA.getTime() - dateB.getTime();
        });
}
