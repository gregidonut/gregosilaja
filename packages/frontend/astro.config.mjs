// @ts-check
import { defineConfig } from "astro/config";
import aws from "astro-sst";

import compressor from "astro-compressor";

// https://astro.build/config
export default defineConfig({
    integrations: [compressor()],
    adapter: aws({
        responseMode: "stream",
    }),
});

