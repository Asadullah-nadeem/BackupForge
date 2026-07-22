'use client';

import dynamic from 'next/dynamic';

const OauthStoragePage = dynamic(
  async () => {
    const { BrowserRouter } = await import('react-router');
    const { OauthStorageComponent } = await import('../../../views/OauthStorageComponent');
    return function WrappedPage() {
      return (
        <BrowserRouter>
          <OauthStorageComponent />
        </BrowserRouter>
      );
    };
  },
  { ssr: false }
);

export default function Page() {
  return <OauthStoragePage />;
}
