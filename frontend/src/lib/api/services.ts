import createClient from "openapi-fetch";

import { paths } from "./v1";
import { Pet } from "./types";

export const apiClient = createClient<paths>({
  baseUrl: "http://127.0.0.1:8080",
});

export const createPet = async (client: typeof apiClient, pet: Pet) => {
  try {
    const { error } = await client.POST("/pets", {
      body: pet,
    });

    return error;
  } catch (e) {
    const error = e as Error;
    console.log(error.message);
    return {
      code: 500,
      message: error.name,
    };
  }
};

export const getPets = async (client: typeof apiClient) => {
  try {
    const { data, error } = await client.GET("/pets");

    if (error) {
      return { data: undefined, error };
    }

    return { data, error: undefined };
  } catch (e) {
    const error = e as Error;
    console.log(error.message);
    return {
      data: undefined,
      error: {
        code: 500,
        message: error.name,
      },
    };
  }
};
