import './globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'

const inter = Inter({ weight: ["400"], subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'ApiFake',
  description: 'All in one Fake Api: ',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        {children}
      </body>
    </html>
  )
}
