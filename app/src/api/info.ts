import type { VoteMap } from '../Vote';

export default async function pollInfo(id: string): Promise<{title: string, votes: VoteMap}> {
    const endpoint = `https://vote.jomy.dev/api/info?pollId=${id}`;
    
    let response = await fetch(endpoint, {
        method: 'GET',
        mode: 'cors'
    });
    
    if (!response.ok) {
        throw response.statusText;
    }
    
    return JSON.parse(await response.text());
}
