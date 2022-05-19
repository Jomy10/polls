import type { Component } from 'solid-js';
import { Switch, Match, createSignal } from 'solid-js';
import Vote from './Vote';
import Home from './Home';

import styles from './App.module.css';

const App: Component = () => {
  // Determines what is being displayed on the screen
  const [show, setShow] = createSignal("home");
  const [pollId, setPollId] = createSignal("");
  
  const showVote = (_pollId: string) => {
    setPollId(_pollId);
    console.log(_pollId);
    setShow("vote");
  }
    
  return (
    <div class={styles.App}>
      <main class={styles.content}>
        <Switch fallback={<div>Not Found</div>}>
          <Match when={show() === "home"}>
            <Home showVote={showVote}/>
          </Match>
          <Match when={show() === "vote"}>
            <Vote pollId={pollId()} />
          </Match>
          <Match when={show() === "create"}>
            <p>Not yet supported</p>
          </Match>
        </Switch>
      </main>
    </div>
  );
};

export default App;
