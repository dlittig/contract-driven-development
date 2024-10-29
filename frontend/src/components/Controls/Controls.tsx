import Button from "../Button";
import CreatePetPopover from "../CreatePetPopover";

const Controls = () => (
  <div className="shrink pb-4">
    <Button label="+ New" popoverTarget="create-pet-popover" />
    <CreatePetPopover />
  </div>
);

export default Controls;
