import { FC } from "react";

import Tag from "../Tag";
import { Pet } from "../../lib/api/types";

const Card: FC<{ title: Pet["name"]; tag: Pet["tag"]; id: Pet["id"] }> = ({
  title,
  tag,
  id,
}) => (
  <div className="flex flex-col gap-4 rounded-xl border border-solid border-neutral-300 p-3 transition-colors hover:bg-white hover:text-neutral-900">
    <p>{title}</p>
    <Tag name={tag} />
    <p className="text-xs text-neutral-500">{id}</p>
  </div>
);

export default Card;
