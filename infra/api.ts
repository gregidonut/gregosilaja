import { bucket } from "./storage";

export const api = new sst.aws.ApiGatewayV2("BlogApi", {
  link: [bucket],
});

api.route("GET /", {
  handler: "packages/functions/cmd/handlers/blogsGet/main.go",
  runtime: "go",
  ...(process.env.SST_STAGE === "dev" ? { timeout: "15 minutes" } : {}),
});
