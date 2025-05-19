const region = aws.getRegionOutput().name;

const stage = process.env.SST_STAGE;
export const frontend = new sst.aws.Astro("Frontend", {
  path: "packages/frontend",
  environment: {
    ASTRO_STAGE: stage,
    ASTRO_REGION: region,
  },
  domain: {
    name: process.env.ASTRO_APP_DOMAIN,
    dns: false,
    cert: process.env.FE_ACM_CERT_ARN,
  },
});
