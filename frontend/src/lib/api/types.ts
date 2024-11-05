import { components, operations } from "./v1";

export type Pet = components["schemas"]["Pet"];

export type CreatePetParams =
  operations["createPets"]["requestBody"]["content"]["application/json"];
