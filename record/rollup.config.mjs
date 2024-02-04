import { babel } from "@rollup/plugin-babel";
import { nodeResolve } from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";

const extensions = [".ts", ".js"];

const preventTreeShakingPlugin = () => {
  return {
    name: 'no-treeshaking',
    resolveId(id, importer) {
      if (!importer) {
          // let's not theeshake entry points, as we're not exporting anything in Apps Script files
        return {id, moduleSideEffects: "no-treeshake" }
      }
      return null;
    }
  }
}

export default {
  input: "./src/index.ts",
  output: {
    dir: "build",
    format: "esm",
  },
  plugins: [
    commonjs(),
    preventTreeShakingPlugin(),
    nodeResolve({
      extensions,
    }),
    babel({ extensions, babelHelpers: "runtime" }),
  ],
};