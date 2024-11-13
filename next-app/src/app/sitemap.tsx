import type { MetadataRoute } from 'next'

const paths = [
    "/",
    "/about",
    "/docs",
    "/docs/users",
    "/docs/todos",
    "/docs/products",
    "/docs/posts",
    "/docs/comments",
    "/docs/images",
    "/docs/movies",
    "/docs/songs",
    "/docs/countries",
    "/docs/cities",
    "/docs/mock",
    "/playground"
]

const domain = "https://fooapi.com"

export default function sitemap(): MetadataRoute.Sitemap {
    let sm: MetadataRoute.Sitemap = [];
    for (let p of paths) {
        sm.push({
            url: domain + p,
            lastModified: new Date(),
            changeFrequency: 'yearly',
            priority: p == "/" ? 1.0 : 0.8,
        })
    }
    return sm;
}