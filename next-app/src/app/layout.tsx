import './globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { GoogleAnalytics } from '@next/third-parties/google'

const inter = Inter({ weight: ["400"], subsets: ['latin'] })

export const metadata: Metadata = {
  title: {
    default: "FooApi | The All in one Fake API",
    template: "FooApi - %s"
  },
  description: 'The All in one Fake API - Dummy data for your projects, fast and simple. Users, products, posts, comments and more',
  keywords: "API, dummy API, fake data, fake API, foo API, mock data, test data, API services, REST API, GraphQL API, GeoJSON API, API documentation, developer tools, FooApi, dummy data API, fake data API, mock data API, test data API, API for developers, REST API services, GraphQL API services, GeoJSON API services, API documentation, developer tools, API data, API GraphQL, API GeoJSON, user data API, todo data API, product data API, post data API, comment data API, image data API, movie data API, song data API, country data API, city data API"
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
