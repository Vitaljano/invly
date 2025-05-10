import js from "@eslint/js";
import globals from "globals";
import tseslint from "typescript-eslint";
import { defineConfig } from "eslint/config";
import lit from "eslint-plugin-lit";
import wc from "eslint-plugin-wc";



export default defineConfig([
  { files: ["**/*.{js,mjs,cjs,ts}"], plugins: { js }, extends: ["js/recommended"] },
  { files: ["**/*.{js,mjs,cjs,ts}"], languageOptions: { globals: globals.browser } },
  lit.configs['flat/recommended'],
  wc.configs['flat/recommended'],
  tseslint.configs.recommended,
]);
