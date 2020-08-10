import React, { useEffect } from "react";
import { useRecoilValue } from "recoil";
import { familyState } from "./store/Family";

const Family = () => {
  const family = useRecoilValue(familyState);

  // For debug
  useEffect(() => {
    console.log(family);
  }, [family]);

  return (
    <div>
      <ul>
        {family.map((p) => {
          return <li key={p.id}> {p.name} </li>;
        })}
      </ul>
    </div>
  );
};

export default Family;
