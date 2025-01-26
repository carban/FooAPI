import Link from "next/link";
import { GrGraphQl } from "react-icons/gr";
import { RiBracesFill, RiMapPinLine, RiPaintBrushLine, RiRobot2Line } from "react-icons/ri";
import { useMemo } from "react";
import dynamic from "next/dynamic";
import { Metadata } from "next";

export const metadata: Metadata = {
   title: "Docs"
}

export default async function Home() {

   const Map = useMemo(() => dynamic(
      () => import('@/app/docs/Map'),
      {
         loading: () => <p>A map is loading</p>,
         ssr: false
      }
   ), [])

   const Map2 = useMemo(() => dynamic(
      () => import('@/app/docs/Map2'),
      {
         loading: () => <p>A map is loading</p>,
         ssr: false
      }
   ), [])

   return <div>
      <h1 className="mb-10 text-3xl font-semibold tracking-tight text-white">Documentation</h1>
      <section id="dummy" className="mb-10">
         <div className="inline-flex">
            <RiBracesFill className="w-8 h-8 text-primary" />
            <h1 className="mb-2 text-2xl font-semibold tracking-tight text-white">Dummy Data</h1>
         </div>
         <p className="text-white text-lg">
            On the left-side menu, you will find instructions on how to perform requests on the dummy data. This is a REST service where, for every category, you will find the endpoint, payload, and response. Feel free to explore the categories and check how they work. For example:
            <code className="text-red-500 text-meidum hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/api/users"> /api/users</Link>
            </code>
         </p>
      </section>
      <hr className="bg-gray-900 border-gray-700" />
      <section id="ql" className="mb-10 mt-10">

         <div className="inline-flex">
            <GrGraphQl className="w-8 h-8 text-primary" />
            <h1 className="mb-2 text-2xl font-semibold tracking-tight text-white">GraphQL</h1>
         </div>
         <p className="text-white text-lg">
            This is a different way to consume data compared to REST. I suggest first checking your queries in
            <code className="text-red-500 text-meidum hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/playground"> /playground</Link>
            </code>
            . The GraphQL queries works on the endpoint
            <code className="text-red-500 text-meidum hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/query"> /query </Link>
            </code>
            . Here is a summary of the Schemas, Queries and Mutations you can perform.
         </p>
         <div className="w-full mt-5 mb-5">
            <table className="w-full border-collapse border border-gray-500">
               <thead>
                  <tr>
                     <th className="border border-gray-500 px-4 py-2">Query</th>
                     <th className="border border-gray-500 px-4 py-2">Mutation</th>
                  </tr>
               </thead>
               <tbody>
                  <tr>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">songs</span>: [<span className="text-pink-400">Song!</span>]!</code>
                        <code className="block"><span className="text-pink-400">song</span>(id: ID!): Song</code>
                     </td>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">createSong</span>(input: CreateSongInput): <span className="text-pink-400">Song!</span></code>
                        <code className="block"><span className="text-pink-400">updateSong</span>(id: ID!, input: UpdateSongInput): <span className="text-pink-400">Song!</span></code>
                        <code className="block"><span className="text-pink-400">deleteSong</span>(id: ID!): <span className="text-pink-400">Song!</span></code>
                     </td>
                  </tr>
                  <tr>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">users</span>: [<span className="text-pink-400">User!</span>]!</code>
                        <code className="block"><span className="text-pink-400">user</span>(id: ID!): User</code>
                     </td>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">createUser</span>(input: CreateUserInput!): <span className="text-pink-400">User!</span></code>
                        <code className="block"><span className="text-pink-400">updateUser</span>(id: ID!, input: UpdateUserInput): <span className="text-pink-400">User!</span></code>
                        <code className="block"><span className="text-pink-400">deleteUser</span>(id: ID!): <span className="text-pink-400">User!</span></code>
                     </td>
                  </tr>
                  <tr>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400"><span className="text-pink-400">posts</span></span>: [<span className="text-pink-400">Post!</span>]!</code>
                        <code className="block"><span className="text-pink-400">post</span>(id: ID!): Post</code>
                     </td>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">createPost</span>(input: CreatePostInput): <span className="text-pink-400">Post!</span></code>
                        <code className="block"><span className="text-pink-400">updatePost</span>(id: ID!, input: UpdatePostInput): <span className="text-pink-400">Post!</span></code>
                        <code className="block"><span className="text-pink-400">deletePost</span>(id: ID!): <span className="text-pink-400">Post!</span></code>
                     </td>
                  </tr>
                  <tr>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">comments</span>: [<span className="text-pink-400">Comment!</span>]!</code>
                        <code className="block"><span className="text-pink-400">comment</span>(id: ID!): Comment</code>
                     </td>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">createComment</span>(input: CreateCommentInput): <span className="text-pink-400">Comment!</span></code>
                        <code className="block"><span className="text-pink-400">updateComment</span>(id: ID!, input: UpdateCommentInput): <span className="text-pink-400">Comment!</span></code>
                        <code className="block"><span className="text-pink-400">deleteComment</span>(id: ID!): <span className="text-pink-400">Comment!</span></code>
                     </td>
                  </tr>
                  <tr>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">products</span>: [<span className="text-pink-400">Product!</span>]!</code>
                        <code className="block"><span className="text-pink-400">product</span>(id: ID!): Product</code>
                     </td>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">createProduct</span>(input: CreateProductInput): <span className="text-pink-400">Product!</span></code>
                        <code className="block"><span className="text-pink-400">updateProduct</span>(id: ID!, input: UpdateProductInput): <span className="text-pink-400">Product!</span></code>
                        <code className="block"><span className="text-pink-400">deleteProduct</span>(id: ID!): <span className="text-pink-400">Product!</span></code>
                     </td>
                  </tr>
                  <tr>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">todos</span>: [<span className="text-pink-400">Todo!</span>]!</code>
                        <code className="block"><span className="text-pink-400">todo</span>(id: ID!): Todo</code>
                     </td>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">createTodo</span>(input: CreateTodoInput): <span className="text-pink-400">Todo!</span></code>
                        <code className="block"><span className="text-pink-400">updateTodo</span>(id: ID!, input: UpdateTodoInput): <span className="text-pink-400">Todo!</span></code>
                        <code className="block"><span className="text-pink-400">deleteTodo</span>(id: ID!): <span className="text-pink-400">Todo!</span></code>
                     </td>
                  </tr>
                  <tr>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">movies</span>: [<span className="text-pink-400">Movie!</span>]!</code>
                        <code className="block"><span className="text-pink-400">movie</span>(id: ID!): Movie</code>

                     </td>
                     <td className="border border-gray-500 px-4 py-2">
                        <code className="block"><span className="text-pink-400">createMovie</span>(input: CreateMovieInput): <span className="text-pink-400">Movie!</span></code>
                        <code className="block"><span className="text-pink-400">updateMovie</span>(id: ID!, input: UpdateMovieInput): <span className="text-pink-400">Movie!</span></code>
                        <code className="block"><span className="text-pink-400">deleteMovie</span>(id: ID!): <span className="text-pink-400">Movie!</span></code>
                     </td>
                  </tr>
               </tbody>
            </table>
         </div>
         <p className="text-lg">
            You can also see this information in the
            <code className="text-red-500 text-meidum hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/playground"> /playground </Link>
            </code>
            docs menu
         </p>
      </section>
      <hr className="bg-gray-900 border-gray-700" />
      <section id="geo" className="mb-10 mt-10">
         <div className="inline-flex ">
            <RiMapPinLine className="w-8 h-8 text-primary" />
            <h1 className="mb-2 text-2xl font-semibold tracking-tight text-white">GeoJSON</h1>
         </div>
         <p className="text-white text-lg">
            It is important highlight that when you use GeoJSON format the coordinates are in the following order <b>[longitude, latitude]</b>. You are free to check the endpoints and data, using a tool like
            <code className="text-red-500 text-meidum hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="https://geojson.io"> https://geojson.io</Link>
            </code> could be useful to see the result in a map quickly.
         </p>
         <p className="text-white text-lg">
            Here is a small map with some data obtained by <code className="text-red-500 text-meidum hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/api/cities?limit=3"> /api/cities?limit=3</Link>
            </code>
         </p>
         <div className="mb-5 mt-5">
            <Map position={[12.22, -1.31]} zoom={2} />
         </div>
         <p className="text-white text-lg">
            Here is another map with some geometry data obtained by <code className="text-red-500 text-meidum hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/api/countries/COL"> /api/countries/COL</Link>
            </code>
            <code className="text-red-500 text-meidum hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/api/countries/EGY"> /api/countries/EGY</Link>
            </code>
            <code className="text-red-500 text-meidum hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/api/countries/IND"> /api/countries/IND</Link>
            </code>
         </p>
         <div className="mb-5 mt-5">
            <Map2 position={[12.22, -1.31]} zoom={2} />
         </div>
      </section>
      <hr className="bg-gray-900 border-gray-700" />
      <section id="custom" className="mb-10 mt-10">
         <div className="inline-flex ">
            <RiPaintBrushLine className="w-8 h-8 text-primary" />
            <h1 className="mb-2 text-2xl font-semibold tracking-tight text-white">Custom Mock</h1>
         </div>
         <p className="text-white text-lg">
            On             <code className="text-red-500 text-meidum hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/api/countries/IND"> /docs/mock</Link>
            </code> you can create custom endpoints in case you need a dummy endpoint to test specific data. The endpoint expires after 1hr so you must to create a new one. In the mean tiem you can edit the last one created. You have the following options available to create your endpoint:
         </p>

         <table className="w-full border-collapse border border-gray-500">
            <thead>
               <tr>
                  <th className="border border-gray-500 px-4 py-2">Method</th>
                  <th className="border border-gray-500 px-4 py-2">Status Code</th>
                  <th className="border border-gray-500 px-4 py-2">Response Body</th>
               </tr>
            </thead>
            <tbody>
               <tr>
                  <td className="border border-gray-500 px-4 py-2">
                     <code className="block"><span className="text-pink-400">GET</span></code>
                     <code className="block"><span className="text-pink-400">POST</span></code>
                     <code className="block"><span className="text-pink-400">PUT</span></code>
                     <code className="block"><span className="text-pink-400">PATCH</span></code>
                     <code className="block"><span className="text-pink-400">DELETE</span></code>
                  </td>
                  <td className="border border-gray-500 px-4 py-2">
                     <code className="block"><span className="text-pink-400">200</span></code>
                     <code className="block"><span className="text-pink-400">400</span></code>
                     <code className="block"><span className="text-pink-400">401</span></code>
                     <code className="block"><span className="text-pink-400">403</span></code>
                     <code className="block"><span className="text-pink-400">404</span></code>
                     <code className="block"><span className="text-pink-400">500</span></code>
                     <code className="block"><span className="text-pink-400">503</span></code>
                  </td>
                  <td>
                     <code className="block"><span className="text-pink-400">{"{...your JSON Object}"}</span></code>
                  </td>
               </tr>
            </tbody>
         </table>
      </section>
      <hr className="bg-gray-900 border-gray-700" />
      <section id="ai" className="mb-10 mt-10">
         <div className="inline-flex ">
            <RiRobot2Line className="w-8 h-8 text-primary" />
            <h1 className="mb-2 text-2xl font-semibold tracking-tight text-white">AI</h1>
         </div>
         <p className="text-white text-lg">
            I am working on it... If you have ideas on this topic contact me :)
         </p>
      </section>
   </div >
}