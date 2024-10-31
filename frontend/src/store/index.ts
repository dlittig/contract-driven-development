import { atom } from "jotai";
import { Pet } from "../lib/api/types";
import { focusAtom } from "jotai-optics";

export type PetStore = {
  byId: Record<string, Pet>;
  allIds: string[];
};

const pets: PetStore = {
  byId: {},
  allIds: [],
};

export const petsAtom = atom(pets);

export const petsByIdAtom = focusAtom(petsAtom, (optic) => optic.prop("byId"));

export const allPetsAtom = atom((get) => {
  const petStore = get(petsAtom);

  return petStore.allIds
    .map((petId) => {
      if (petStore.byId[petId]) {
        return petStore.byId[petId];
      }
    })
    .filter((item) => !!item);
});
