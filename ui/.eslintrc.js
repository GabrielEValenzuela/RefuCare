module.exports = {
  root: true,
  parser: "@typescript-eslint/parser",
  parserOptions: {
    ecmaVersion: "latest",
    sourceType: "module",
    project: "./tsconfig.json",
  },
  plugins: ["react", "@typescript-eslint", "prettier"],
  extends: [
    "airbnb-typescript/base",
    "next/core-web-vitals",
    "plugin:prettier/recommended",
  ],
  rules: {
    "max-len": ["error", 140],
    quotes: ["error", "double", { avoidEscape: true }],
    "prettier/prettier": ["error", { endOfLine: "auto" }],
  },
};
