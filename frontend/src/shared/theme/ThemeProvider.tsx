import { useCallback, useEffect, useMemo, useState } from 'react';
import type { ReactNode } from 'react';

import { ThemeContext } from './themeContext';
import type { ResolvedTheme, ThemeMode } from './themeContext';

const THEME_STORAGE_KEY = 'BackupForge-theme';



function getStoredTheme(): ThemeMode {
  if (typeof window !== 'undefined') {
    const stored = localStorage.getItem(THEME_STORAGE_KEY);

    if (stored === 'light' || stored === 'dark' || stored === 'system') {
      return stored;
    }
  }

  return 'system';
}

interface ThemeProviderProps {
  children: ReactNode;
}

export function ThemeProvider({ children }: ThemeProviderProps) {
  const [theme, setThemeState] = useState<ThemeMode>(getStoredTheme);

  // Apply theme class to document
  const resolvedTheme: ResolvedTheme = 'dark';

  // Apply theme class to document
  useEffect(() => {
    const root = document.documentElement;
    root.classList.add('dark');
  }, []);

  const setTheme = useCallback((newTheme: ThemeMode) => {
    setThemeState(newTheme);
    localStorage.setItem(THEME_STORAGE_KEY, newTheme);
  }, []);

  const value = useMemo(
    () => ({
      theme,
      resolvedTheme,
      setTheme,
    }),
    [theme, resolvedTheme, setTheme],
  );

  return <ThemeContext.Provider value={value}>{children}</ThemeContext.Provider>;
}
