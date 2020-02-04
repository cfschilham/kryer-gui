module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    "plugin:vue/recommended",
    "@vue/standard",
    "@vue/typescript",
  ],
  rules: {
    "no-console": "off",
    "no-debugger": "off",
    "no-array-constructor": "off",
    quotes: ["error", "double", ],
    "space-before-function-paren": ["error", "never", ],
    "comma-dangle": ["error", "always", ],
    "no-extra-semi": "error",
    semi: ["error", "never", ],
  },
  parserOptions: {
    parser: "@typescript-eslint/parser",
  },
}
