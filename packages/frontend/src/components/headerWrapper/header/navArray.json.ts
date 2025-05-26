type navItem = {
    name: string;
    url: string;
};
const navArray: readonly navItem[] = [
    { name: "Home", url: "/" },
    // { name: "Portfolio", url: "/portfolio" },
    { name: "Blog", url: "/blog" },
    { name: "About Me", url: "/about" },
    { name: "CV", url: "/cv" },
];
export default navArray;
