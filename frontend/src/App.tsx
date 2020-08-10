import React from "react";
import "./App.css";
import Family from "./Family";
import { RecoilRoot } from "recoil";

function App() {
  return (
    <div className="App">
      <RecoilRoot>
        <Family />
      </RecoilRoot>
    </div>
  );
}

export default App;
