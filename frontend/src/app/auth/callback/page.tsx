'use client';

import dynamic from 'next/dynamic';

const AuthCallbackPage = dynamic(
  async () => {
    const { BrowserRouter } = await import('react-router');
    const { OAuthCallbackPage } = await import('../../../views/OAuthCallbackPage');
    return function WrappedPage() {
      return (
        <BrowserRouter>
          <OAuthCallbackPage />
        </BrowserRouter>
      );
    };
  },
  { ssr: false }
);

export default function Page() {
  return <AuthCallbackPage />;
}
