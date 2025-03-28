// @ts-check
import { defineConfig } from "astro/config";
import tailwindcss from "@tailwindcss/vite";

import react from "@astrojs/react";

import sitemap from "@astrojs/sitemap";

// https://astro.build/config
export default defineConfig({
	integrations: [
		react({
			include: ["**/react/*"],
		}),
		sitemap(),
	],
	vite: {
		ssr: {
			noExternal: ["webcoreui"],
		},
		plugins: [tailwindcss()],
	},
	experimental: {
		contentIntellisense: true,
	},
});
