import type { Component } from 'solid-js';
import { createSignal, Show } from 'solid-js';
import pollInfo from './api/info';
import { ReactiveSet } from '@solid-primitives/set';
import castVote from './api/vote';

// TODO: fix this
export type VoteMap = {
  [key: string]: number;
}

/** poll document */
async function getPoll(pollId: string): Promise<{ title: string, votes: VoteMap }> {
  if (pollId == "") return null;

  const info: { title: string, votes: VoteMap } = await pollInfo(pollId);
  return info;
};

// TODO: support only one option per poll as an option

/**
 * Voting screen
 * @param params {Object} - Parameters of a poll
 */
const Vote: Component = (params: { pollId: string }) => {
  const pollId: string = params.pollId;
  console.log("pollId in Vote", pollId);
  
  let [loaded, setLoaded] = createSignal(false);
  let [title, setTitle] = createSignal("");
  let [options, setOptions] = createSignal([]);
  
  let pollVotes: Object;
  
  getPoll(pollId)
    .then((poll) => {
      if (poll == null) return; // TODO: handle
      setTitle(poll.title);
      for (let key of Object.keys(poll.votes)) {
        setOptions(opts => [...opts, key]);
      }
      
      setLoaded(true);
      
      pollVotes = poll.votes;
    });
    
  const votes = new ReactiveSet<string>([]);
  /** Vote for an option (or unvote) */
  const toggleVote = (opt: string) => {
    votes.has(opt) ? votes.delete(opt) : votes.add(opt);
  };
  
  const [voted, setVoted] = createSignal(false);
  /** Submit votes */
  const submitVote = () => {
    // Update local vote count
    for (let vote of votes) {
      let currentVotes = pollVotes[vote];
      if (currentVotes == null) {
        pollVotes[vote] = 1;
      } else {
        pollVotes[vote] = currentVotes + 1;
      }
      
      // send vote
      castVote(pollId, vote)
        .catch((e) => {
          console.log("Couldn't cast vote", e);
          // TODO: warn user
        });
        
      setVoted(true);
    }
  };
  
  // TODO: styling for votes
    
  return <>
    <Show when={loaded()} fallback={<p>Loading...</p>}>
      <Switch fallback={<p>Error: Unknown state</p>}>
        <Match when={!voted()}>
          <h1>{title()}</h1>
          <ul>
            <For each={options()}>{(opt, i) => 
              <li>
                <button 
                  class={`${ votes.has(opt) ? "selected" : ""}`}
                  onClick={() => { toggleVote(opt) } }
                >
                  {opt}
                </button>
              </li>
            }</For>
          </ul>
          <button onClick={submitVote}>Submit vote</button>
        </Match>
        <Match when={voted()}>
          <h1>{title()}</h1>
          <ul>
            <For each={options()}>{(opt, i) => 
              <li>
                <p>{opt}: {pollVotes[opt]}</p>
              </li>
            }</For>
          </ul>
        </Match>
      </Switch>
   </Show>
  </>;
};

export default Vote;