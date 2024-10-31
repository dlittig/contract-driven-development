import { useSetAtom } from "jotai";
import { apiClient, getPets } from "./lib/api/services";
import { petsAtom, PetStore } from "./store";

export const wait = (ms: number) =>
  new Promise((resolve) => {
    setTimeout(() => resolve(true), ms);
  });

export const usePetRetriever = () => {
  const setPets = useSetAtom(petsAtom);

  return async () => {
    const { data: remotePets, error } = await getPets(apiClient);

    if (error) {
      return alert("Failed to fetch pets :(");
    }

    if (remotePets) {
      const allIds = remotePets.map((p) => p.id);
      const byId: PetStore["byId"] = {};

      remotePets.forEach((p) => {
        byId[p.id] = p;
      });

      setPets({
        allIds,
        byId,
      });
    }
  };
};
