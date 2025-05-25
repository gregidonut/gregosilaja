import { createContext, useContext } from "react";

export const FooterHTMLContext = createContext("");

export function useHTMLFooter(): string {
    return useContext(FooterHTMLContext);
}
