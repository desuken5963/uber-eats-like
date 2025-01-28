import React from 'react';
import StyledComponentsRegistry from '@/lib/registry';

export const metadata = {
  title: 'Restaurant List',
  description: 'A list of restaurants',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <StyledComponentsRegistry>{children}</StyledComponentsRegistry>
      </body>
    </html>
  );
}