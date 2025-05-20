import type { APIRoute } from "astro";

export const GET: APIRoute = async function () {
    return await fetch("https://zenquotes.io/api/random");
};
