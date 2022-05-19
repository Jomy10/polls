import type { VoteMap } from '../Vote';
import _endpoint from './endpoint';

export default async function pollInfo(id: string): Promise<{title: string, votes: VoteMap}> {
    const endpoint = `${_endpoint}/api/info?pollId=${id}`;
    
    let response = await fetch(endpoint, {
        method: 'GET',
        mode: 'cors'
    });
    
    if (!response.ok) {
        throw response.statusText;
    }
    
    return JSON.parse(await response.text());
}
