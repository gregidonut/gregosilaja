import React from "react";
import navArray from "../../navArray.json";
import styles from "./navList.module.css";
import Footer from "./footer/Footer";

export default function NavList(): React.JSX.Element {
    return (
        <ul className={styles.ul}>
            {navArray.map(function (li, i) {
                return (
                    <li key={i}>
                        <a href={li.url}>{li.name}</a>
                    </li>
                );
            })}
            <li>
                <Footer />
            </li>
        </ul>
    );
}
