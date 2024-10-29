import { FC } from "react";
import { Pet } from "../../lib/api/types";

const Tag: FC<{ name: Pet["tag"] }> = ({ name }) => (
  <>
    {name && (
      <span className="inline w-fit rounded-xl bg-blue-500 px-3 py-1 text-white">
        {name}
      </span>
    )}
  </>
);

export default Tag;
