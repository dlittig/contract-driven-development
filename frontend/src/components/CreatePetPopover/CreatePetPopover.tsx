import { useRef, useState } from "react";
import Button from "../Button";
import { apiClient, createPet } from "../../lib/api/services";
import { usePetRetriever } from "../../utils";

const CreatePetPopover = () => {
  const [name, setName] = useState("");
  const [tag, setTag] = useState("");
  const retrievePets = usePetRetriever();
  const popoverRef = useRef<HTMLDivElement>(null);

  const onClick = () => {
    if (name.length === 0 || tag.length === 0) {
      return alert("Please fill all fields.");
    }

    createPet(apiClient, {
      id: crypto.randomUUID(),
      name,
      tag,
    }).then((error) => {
      if (error) {
        return alert("An error hos occured during submit!");
      }

      popoverRef.current?.togglePopover();
      retrievePets();
    });

  };

  return (
    <div
      className="rounded backdrop:backdrop-blur"
      id="create-pet-popover"
      popover="auto"
      ref={popoverRef}
    >
      <div className="flex min-w-[90vw] flex-col gap-6 p-4 md:min-w-[40vw]">
        <h2 className="text-center text-xl">Create new pet</h2>

        <label>
          <p className="pb-2">Name:</p>
          <input
            name="name"
            className="w-full rounded border border-solid border-neutral-200 bg-transparent px-4 py-2 text-white transition-colors focus:border-blue-600"
            placeholder="The new name..."
            defaultValue={name}
            onChange={(event) => setName(event.target.value)}
          />
        </label>
        <label>
          <p className="pb-2">Tag:</p>
          <input
            name="tag"
            className="w-full rounded border border-solid border-neutral-200 bg-transparent px-4 py-2 text-white transition-colors focus:border-blue-600"
            placeholder="The new tag..."
            defaultValue={tag}
            onChange={(event) => setTag(event.target.value)}
          />
        </label>

        <hr className="border-neutral-400" />

        <Button label="Submit" onClick={onClick} />
      </div>
    </div>
  );
};
export default CreatePetPopover;
