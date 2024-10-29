import createClient from "openapi-fetch";

import { paths } from "./v1";
import { Pet } from "./types";

const apiClient = createClient<paths>({
  baseUrl: "http://127.0.0.1:8080/api/v1/",
});

export const createPet = async (pet: Pet) => {
  try {
    const { error } = await apiClient.POST("/pets", {
      body: pet,
    });

    return error;
  } catch (e) {
    return e;
  }
};
