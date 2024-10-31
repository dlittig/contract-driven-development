import { useEffect, useState } from "react";
import { useAtom, useSetAtom } from "jotai";

import Card from "./components/Card";
import Loading from "./components/Loading";
import Controls from "./components/Controls";
import { apiClient, getPets } from "./lib/api/services";
import { allPetsAtom, petsAtom, PetStore } from "./store";

const App = () => {
  const [allPets] = useAtom(allPetsAtom);
  const setPets = useSetAtom(petsAtom);
  const [isPending, setPending] = useState(false);

  const retrievePets = async () => {
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

  useEffect(() => {
    (async () => {
      setPending(true);
      await retrievePets();
      setPending(false);
    })();
  }, []);

  return (
    <div className="flex flex-col gap-4 p-4">
      <h1 className="shrink text-center text-4xl">My Pets</h1>
      <Controls />

      {allPets.length === 0 && (
        <div className="my-6 flex w-full flex-col items-center">
          <p className="pb-2 text-center text-neutral-500">
            Brighten up your life with a pet :)
          </p>
        </div>
      )}

      {isPending && (
        <div className="my-6 flex w-full flex-col items-center">
          <Loading />
        </div>
      )}

      <div className="grid grow grid-cols-3 gap-4">
        {allPets.length > 0 &&
          allPets.map((pet) => (
            <Card
              key={`pet-${pet.id}`}
              tag={pet.tag}
              title={pet.name}
              id={pet.id}
            />
          ))}
      </div>
    </div>
  );
};

export default App;
