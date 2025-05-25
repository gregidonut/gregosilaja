import React from "react";
import { useHTMLFooter } from "../../context.ts";

export default function Footer(): React.JSX.Element {
    const htmlFooter = useHTMLFooter();

    return <div dangerouslySetInnerHTML={{ __html: htmlFooter }} />;
}
