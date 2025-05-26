import { api } from "./api";

const region = aws.getRegionOutput().name;

const stage = process.env.SST_STAGE;
export const frontend = new sst.aws.Astro("Frontend", {
  path: "packages/frontend",
  link: [api],
  environment: {
    ASTRO_STAGE: stage,
    ASTRO_REGION: region,
    ASTRO_API_URL: api.url,
  },
  domain: {
    name: process.env.ASTRO_APP_DOMAIN,
    dns: false,
    cert: process.env.FE_ACM_CERT_ARN,
    redirects: [process.env.ASTRO_APP_DOMAIN_REDIRECT],
  },
});
