import adapter from '@sveltejs/adapter-auto';
import { vitePreprocess } from '@sveltejs/kit/vite';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  // Consult https://kit.svelte.dev/docs/integrations#preprocessors
  // for more information about preprocessors
  preprocess: vitePreprocess(),

  kit: {
    adapter: adapter(),
    alias: {
      '#routes': 'src/routes',
      '#components': 'src/lib/components',
      '#icons': 'src/lib/components/Icons',
      '#services': 'src/lib/services',
      '#constants': 'src/lib/constants'
    },
    
  }
};

export default config;
