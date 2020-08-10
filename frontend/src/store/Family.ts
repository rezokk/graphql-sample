import { atom } from "recoil";

export interface Person {
  id: number;
  name: string;
}

const initialFamily: Person[] = [
  { id: 1, name: "cappyzawa" },
  { id: 2, name: "bookun" },
];

export const familyState = atom({
  key: "family",
  default: initialFamily,
});
