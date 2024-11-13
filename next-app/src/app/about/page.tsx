import Link from "next/link";
import Foot from "../components/foot";
import Navbar from "../components/navbar";

export default function About() {
    return <div>
        <Navbar />
        <div className="container mx-auto px-5 text-center mb-64">
            <h1 className="mb-10 text-3xl font-semibold tracking-tight text-white">About</h1>
            <p className="text-white text-lg">
                Hi 👋, this is a personal project, I hope it is useful for you.
            </p>
            <p className="text-white text-lg">
                This is an API you can use as placeholder in your app. The idea is to use close realistic data, fast and simple.
            </p>
            <p className="text-white text-lg">
                I am working to finish the features announced and keep adding new features that help people get better dummy data for their projects.
            </p>
            <p className="text-white text-lg">
                If you want to contact me or give me ideas to enhance the project, email to me
                <code className="text-red-500 text-lg hover:underline mb-2 md:mb-0 w-64 flex-1">
                    <Link href="mailto:info@fooapi.com"> info@fooapi.com </Link>
                </code>
            </p>

        </div>

        <Foot />
    </div>
}