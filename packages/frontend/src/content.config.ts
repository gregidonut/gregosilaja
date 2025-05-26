import { defineCollection } from "astro:content";
import { Blog } from "@/utils/models";
import axios from "axios";
import axiosRetry from "axios-retry";
axiosRetry(axios, {
    retries: 5,
    retryDelay: axiosRetry.exponentialDelay,
    retryCondition: function (error) {
        return (
            axiosRetry.isNetworkError(error) ||
            axiosRetry.isRetryableError(error) ||
            [502, 503, 504].includes(Number(error.response?.status) || -1)
        );
    },
    onRetry: function (retryCount, error, requestConfig) {
        console.warn(
            `Retrying [${retryCount}]: ${requestConfig.url} (${error.message})`,
        );
    },
});
export const collections = {
    blog: defineCollection({
        loader: async function (): Promise<Blog[]> {
            const ASTRO_API_URL = import.meta.env.ASTRO_API_URL as string;

            const resp = await axios.get(ASTRO_API_URL);
            return resp.data;
        },
    }),
};
