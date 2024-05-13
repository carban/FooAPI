import Navbar from '@/app/components/navbar';
import Featured from '@/app/components/home/featured';
import AllCategories from '@/app/components/home/allCategories';
import Foot from './components/foot';

const InitialBanner = () => <center>
   <div className="pt-14 pb-10 hero">
      <h1 className="text-5xl text-white">
         The All in One Fake API
      </h1>
      <h3 className="py-5 text-lg text-white">Dummy data for your projects, fast and simple. Users, products, posts, comments and more!</h3>
   </div>
</center >

const HelloWorld = () => <h1>
   Hello World
</h1>

const MainHome = () => <div>
   <Navbar />
   <InitialBanner />
   <div className="container max-w-full lg:px-20 md:px-10 py-10 bg-gray-900">
      <Featured />
      <AllCategories />
   </div>
   <Foot />
</div >

export default function Home() {
   return <MainHome />
}