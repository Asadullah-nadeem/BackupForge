export interface RuntimeConfig {
  GITHUB_CLIENT_ID?: string;
  GOOGLE_CLIENT_ID?: string;
  IS_EMAIL_CONFIGURED?: string;
  CLOUDFLARE_TURNSTILE_SITE_KEY?: string;
  CONTAINER_ARCH?: string;
}

declare global {
  interface Window {
    __RUNTIME_CONFIG__?: RuntimeConfig;
  }
}

const runtimeConfig = typeof window !== 'undefined' ? window.__RUNTIME_CONFIG__ : undefined;

export function getApplicationServer() {
  if (typeof window === 'undefined') {
    return 'http://localhost:4005';
  }
  const origin = window.location.origin;
  const url = new URL(origin);

  const isDevelopment = process.env.NODE_ENV === 'development';

  if (isDevelopment) {
    return `${url.protocol}//${url.hostname}:4005`;
  } else {
    return `${url.protocol}//${url.hostname}:${url.port || (url.protocol === 'https:' ? '443' : '80')}`;
  }
}

export const APP_VERSION = (process.env.NEXT_PUBLIC_APP_VERSION as string) || '3.26.0';

export const GITHUB_CLIENT_ID =
  runtimeConfig?.GITHUB_CLIENT_ID || process.env.NEXT_PUBLIC_GITHUB_CLIENT_ID || '';

export const GOOGLE_CLIENT_ID =
  runtimeConfig?.GOOGLE_CLIENT_ID || process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || '';

export const IS_EMAIL_CONFIGURED =
  runtimeConfig?.IS_EMAIL_CONFIGURED === 'true' ||
  process.env.NEXT_PUBLIC_IS_EMAIL_CONFIGURED === 'true';

export const CLOUDFLARE_TURNSTILE_SITE_KEY =
  runtimeConfig?.CLOUDFLARE_TURNSTILE_SITE_KEY ||
  process.env.NEXT_PUBLIC_CLOUDFLARE_TURNSTILE_SITE_KEY ||
  '';

const archMap: Record<string, string> = { amd64: 'x64', arm64: 'arm64' };
const rawArch = runtimeConfig?.CONTAINER_ARCH || 'windows-x64';
export const CONTAINER_ARCH = archMap[rawArch] || rawArch;

export function getOAuthRedirectUri(): string {
  if (typeof window === 'undefined') {
    return 'http://localhost:4005/auth/callback';
  }
  return `${window.location.origin}/auth/callback`;
}
