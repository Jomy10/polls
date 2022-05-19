import type { Component } from 'solid-js';

/**
 * @param params {Object}: an object containing `showVote`, which is a function
 * that taks a number as an argument (the pollId)
 */
const Home: Component = (params: { showVote: (arg: any) => void }) => {
  const onKeyDown = (e: any) => {
    console.log(e);
    if (e.key == "Enter") {
      let pollIdVal = e.target.value;
      params.showVote(pollIdVal);
    }
  }
    
  return <>
    <h1>Polls</h1>
    <input 
      type="text"
      id="pollId"
      placeholder="Poll id"
      onkeydown={onKeyDown}
    />
  </>;
};

export default Home;