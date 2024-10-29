import Card from "./components/Card";
import Controls from "./components/Controls";
import { Pet } from "./lib/api/types";

const pets: Pet[] = [
  { name: "Snoops", tag: "Dog", id: crypto.randomUUID() },
  { name: "Nemo", tag: "Fish", id: crypto.randomUUID() },
];

const App = () => (
  <div className="flex flex-col gap-4 p-4">
    <h1 className="shrink text-center text-4xl">My Pets</h1>
    <Controls />

    {pets.length === 0 && (
      <p className="text-center text-neutral-500">
        Brighten up your life with a pet :)
      </p>
    )}

    <div className="grid grow grid-cols-3 gap-4">
      {pets.length > 0 &&
        pets.map((pet) => (
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

export default App;
