import React from 'react'
import type { Metadata } from 'next'
import { Manrope } from 'next/font/google'
import './css/globals.css'
import { ThemeProvider } from '@/components/theme-provider'
import { Providers } from './providers'

const manrope = Manrope({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: 'BackupForge',
  description: 'A premium native database backup & restore application for Windows',
  icons: {
    icon: '/favicon.svg',
  },
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang='en' suppressHydrationWarning className="dark">
      <head>
        <link rel='icon' href='/favicon.svg' type='image/svg+xml' />
      </head>
      <body className={`${manrope.className}`}>
        <ThemeProvider
          attribute='class'
          defaultTheme='dark'
          enableSystem={false}
          disableTransitionOnChange>
          <Providers>
            {children}
          </Providers>
        </ThemeProvider>
      </body>
    </html>
  )
}
