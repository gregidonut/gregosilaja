// @ts-check
import { defineConfig } from "astro/config";
import aws from "astro-sst";

import compressor from "astro-compressor";

import react from "@astrojs/react";

// https://astro.build/config
export default defineConfig({
    site: "https://gregosilaja.cc",
    integrations: [compressor(), react()],
    adapter: aws({
        responseMode: "stream",
    }),
});
