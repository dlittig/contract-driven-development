import { setupServer } from "msw/node";
import { http, HttpResponse } from "msw";
import createClient from "openapi-fetch";
import { paths } from "../../lib/api/v1";
import { Pet } from "../../lib/api/types";
import { createPet } from "../../lib/api/services";

const server = setupServer();

beforeAll(() => {
  server.listen({
    onUnhandledRequest: (request) => {
      throw new Error(
        `No request handler found for ${request.method} ${request.url}`
      );
    },
  });
});

afterEach(() => server.resetHandlers());
afterAll(() => server.close());

describe("pet api", () => {
  it("creates a new pet successfully", async () => {
    const BASE_URL = "http://127.0.0.1:8080";

    server.use(
      http.post(`${BASE_URL}/pets`, () => {
        return HttpResponse.json(null, {
          status: 201,
        });
      })
    );

    const mockPet: Pet = {
      id: crypto.randomUUID(),
      name: "Bober",
      tag: "Beever",
    };

    const client = createClient<paths>({ baseUrl: BASE_URL });
    const { data, error, response } = await client.POST("/pets", {
      body: mockPet,
    });

    expect(data).toBe(null);
    expect(response.status).toBe(201);
    expect(error).toBeUndefined();
  });

  it("creates a new pet via service successfully", async () => {
    const BASE_URL = "http://127.0.0.1:8080";

    server.use(
      http.post(`${BASE_URL}/pets`, () => {
        return HttpResponse.json(null, {
          status: 201,
        });
      })
    );

    const mockPet: Pet = {
      id: crypto.randomUUID(),
      name: "Bober",
      tag: "Beever",
    };

    const client = createClient<paths>({ baseUrl: BASE_URL });
    const error = await createPet(client, mockPet);

    expect(error).toBeUndefined();
  });

  it("fails to create a new pet when a server error occurs", async () => {
    const BASE_URL = "http://127.0.0.1:8080";

    server.use(
      http.post(`${BASE_URL}/pets`, () => {
        return HttpResponse.json(
          {
            code: 500,
            message: "An unexpected error occurred.",
          },
          {
            status: 500,
          }
        );
      })
    );

    const mockPet: Pet = {
      id: crypto.randomUUID(),
      name: "Bober",
      tag: "Beever",
    };

    const client = createClient<paths>({ baseUrl: BASE_URL });
    const { data, error, response } = await client.POST("/pets", {
      body: mockPet,
    });

    expect(data).toBe(undefined);
    expect(response.status).toBe(500);
    expect(error).toBeDefined();
    expect(error!.code).toBe(500);
  });

  it("fails to create a new pet via service when a server error occurs", async () => {
    const BASE_URL = "http://127.0.0.1:8080";

    server.use(
      http.post(`${BASE_URL}/pets`, () => {
        return HttpResponse.json(
          {
            code: 500,
            message: "An unexpected error occurred.",
          },
          {
            status: 500,
          }
        );
      })
    );

    const mockPet: Pet = {
      id: crypto.randomUUID(),
      name: "Bober",
      tag: "Beever",
    };

    const client = createClient<paths>({ baseUrl: BASE_URL });
    const error = await createPet(client, mockPet);

    expect(error).toBeDefined();
    expect(error!.code).toBe(500);
    expect(error?.message).toBe("An unexpected error occurred.");
  });
});
