import './globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { GoogleAnalytics } from '@next/third-parties/google'

const inter = Inter({ weight: ["400"], subsets: ['latin'] })

export const metadata: Metadata = {
  title: {
    default: "FooApi | The All in one Fake API",
    template: "%s | FooApi"
  },
  description: 'The All in one Fake API - Dummy data for your projects, fast and simple. Users, products, posts, comments and more',
  keywords: "foo API, fooApi, dummy API, fake data, fake API, fake api, fakeapi, fooapi, foo api, foo api data, foo api graphql, foo api ql, foo API, foo API, dummy fake api, fooAPI, mock data, test data, API services, REST API, GraphQL API, GeoJSON API, FooApi, comments fake api, posts fake api, products fake api, users fake api, todos fake api, songs fake api, movies fake api, countries fake api, cities fake api, images fake api, dummy data API, fake data API, mock data API, test data API, API for developers, REST API services, GraphQL API services, GeoJSON API services, API data, API GraphQL, API GeoJSON, user data API, todo data API, product data API, post data API, comment data API, image data API, movie data API, song data API, country data API, city data API, dummy json, dummyJSON"
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
