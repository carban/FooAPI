import Link from "next/link";
import { GrGraphQl } from "react-icons/gr";
import { RiBracesFill, RiMapPinLine, RiPaintBrushLine, RiRobot2Line } from "react-icons/ri";
import { useMemo } from "react";
import dynamic from "next/dynamic";

export default async function Home() {

   const Map = useMemo(() => dynamic(
      () => import('@/app/docs/Map'),
      {
         loading: () => <p>A map is loading</p>,
         ssr: false
      }
   ), [])

   return <div>
      <section id="dummy" className="mb-10">
         <div className="inline-flex">
            <RiBracesFill className="w-8 h-8 text-primary" />
            <h1 className="mb-2 text-2xl font-semibold tracking-tight text-white">Dummy Data</h1>
         </div>
         <p className="text-white text-lg">
            On the left side menu you will find how perform the requests on the dummy data. This is a REST service where for every categorie you will find the endpoint, payload and response. You are free to explore the categories and check how they works.      </p>
      </section>
      <hr className="bg-gray-900 border-gray-700"/>
      <section id="ql" className="mb-10 mt-10">

         <div className="inline-flex">
            <GrGraphQl className="w-8 h-8 text-primary" />
            <h1 className="mb-2 text-2xl font-semibold tracking-tight text-white">GraphQL</h1>
         </div>
         <p className="text-white text-lg">
            This is a different way of consume data comparing with REST, we suggest first check your queries in
            <code className="text-red-500 text-base hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/playground"> /playground</Link>
            </code>
            . The GraphQL queries works on the endpoint
            <code className="text-red-500 text-base hover:underline mb-2 md:mb-0 w-64 flex-1">
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
            <code className="text-red-500 text-base hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/playground"> /playground </Link>
            </code>
            docs menu
         </p>
      </section>
      <hr className="bg-gray-900 border-gray-700"/>
      <section id="geo" className="mb-10 mt-10">
         <div className="inline-flex ">
            <RiMapPinLine className="w-8 h-8 text-primary" />
            <h1 className="mb-2 text-2xl font-semibold tracking-tight text-white">GeoJSON</h1>
         </div>
         <p className="text-white text-lg">
            Here is a small map with some data obtained by <code className="text-red-500 text-base hover:underline mb-2 md:mb-0 w-64 flex-1">
               <Link href="/api/cities?limit=3"> /api/cities?limit=3</Link>
            </code>
         </p>
         <div className="mb-5 mt-5">
            <Map position={[12.22, -1.31]} zoom={2} />
         </div>
         <p className="text-white text-lg">
            We are working on the <u className="font-extrabold">Countries</u> data colletion. The plan is that you can get the geometry of some countries :)  
         </p>
      </section>
      <hr className="bg-gray-900 border-gray-700"/>
      <section id="custom" className="mb-10 mt-10">
         <div className="inline-flex ">
            <RiPaintBrushLine className="w-8 h-8 text-primary" />
            <h1 className="mb-2 text-2xl font-semibold tracking-tight text-white">Custom Mock</h1>
         </div>
         <p className="text-white text-lg">
            We are working on it... wait for it soon :)
         </p>
      </section>
      <hr className="bg-gray-900 border-gray-700"/>
      <section id="ai" className="mb-10 mt-10">
         <div className="inline-flex ">
            <RiRobot2Line className="w-8 h-8 text-primary" />
            <h1 className="mb-2 text-2xl font-semibold tracking-tight text-white">AI</h1>
         </div>
         <p className="text-white text-lg">
            We are working on it... wait for it soon :)
         </p>
      </section>
   </div >
}