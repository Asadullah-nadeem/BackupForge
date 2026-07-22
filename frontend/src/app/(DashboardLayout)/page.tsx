'use client';

import { useEffect, useState } from 'react';
import { userApi } from '../../entity/users';
import { AuthPageComponent } from '../../views/AuthPageComponent';
import { MainScreenComponent } from '../../widgets/main/MainScreenComponent';

export default function Home() {
  const [isAuthorized, setIsAuthorized] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    setIsAuthorized(userApi.isAuthorized());
    setIsLoading(false);

    userApi.addAuthListener(() => {
      setIsAuthorized(userApi.isAuthorized());
    });
  }, []);

  if (isLoading) return null;

  return !isAuthorized ? <AuthPageComponent /> : <MainScreenComponent />;
}
