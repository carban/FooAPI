import './globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { GoogleAnalytics } from '@next/third-parties/google'

const inter = Inter({ weight: ["400"], subsets: ['latin'] })

export const metadata: Metadata = {
  title: {
    default: "FooApi",
    template: "FooApi - %s"
  },
  description: 'The All in one Fake Api - Dummy data for your projects, fast and simple. Users, products, posts, comments and more',
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
        <GoogleAnalytics gaId="G-V1GX8Q1LH5" />
      </body>
    </html>
  )
}
