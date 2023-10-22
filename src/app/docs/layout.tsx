import VerticalNavbar from '@/app/components/docs/verticalNavbar';
import Foot from "@/app/components/foot";

export default function DocsLayout({ children }: any) {
    return <div>
        <VerticalNavbar>
            {children}
        </VerticalNavbar>
        <Foot />
    </div>
}