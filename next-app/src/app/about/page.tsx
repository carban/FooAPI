import Link from "next/link";
import Foot from "../components/foot";
import Navbar from "../components/navbar";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "About"
}

export default function About() {
  return <div>
    <Navbar />
    <div className="container mx-auto px-5 text-center mb-64">
      <h1 className="mb-10 text-3xl font-semibold tracking-tight text-white">About</h1>
      <p className="text-white text-lg">
        Hi 👋, this is a personal project, I hope it is useful for you.
      </p>
      <p className="text-white text-lg">
        This project aims to provide test data and different ways to manipulate it, all in one place, quickly and easily.            </p>
      <p className="text-white text-lg">
        That is why here you can find generic data, geographic data, images and methods like REST and GraphQL.            </p>
      <p className="text-white text-lg">
        I am working to finish the features announced and keep adding new features that help people get better dummy data for their projects.
      </p>
      <p className="text-white text-lg">
        If you want to contact me or give me ideas to enhance the project, email to me
        <code className="text-red-500 text-lg hover:underline mb-2 md:mb-0 w-64 flex-1">
          <Link href="mailto:info@fooapi.com"> info@fooapi.com </Link>
        </code>
      </p>
      <br />
      <br />
      <p className="text-white text-lg">
        If you find this project useful, please consider supporting it:
      </p>
      <br />
      <center>
        <a href="https://www.buymeacoffee.com/carban" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-red.png" alt="Buy Me A Coffee" style={{ height: "60px", width: "217px" }} /></a>
      </center>
    </div>
    <Foot />
  </div>
}