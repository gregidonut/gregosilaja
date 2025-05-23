import React, { useState, useRef, useEffect } from "react";
import { gsap } from "gsap";
import { useGSAP } from "@gsap/react";

import navBurger from "@/assets/nav-burger.svg?url";
import styles from "./reactNavSection.module.css";

import NavList from "./navList/NavList";

gsap.registerPlugin(useGSAP);

export default function (): React.JSX.Element {
    const [navBurgerClicked, setNavBurgerClicked] = useState<boolean>(false);
    const container = useRef(null);
    const tweenRef = useRef<gsap.core.Tween | null>(null);

    useGSAP(
        function () {
            tweenRef.current = gsap.from(".navSection", {
                x: -1500,
                duration: 0.2,
                paused: true,
            });
        },
        { scope: container },
    );

    useEffect(() => {
        if (tweenRef.current) {
            navBurgerClicked
                ? tweenRef.current.play()
                : tweenRef.current.reverse();
        }
    }, [navBurgerClicked]);
    return (
        <div className={styles.navSectionContainer}>
            <div className={styles.navBurger}>
                <button
                    id="nav-burger-btn"
                    onClick={function () {
                        setNavBurgerClicked((prev) => !prev);
                    }}
                >
                    <img src={navBurger} alt="nav-burger" />
                </button>
            </div>
            <div ref={container}>
                <section className={`${styles.section} navSection`}>
                    <NavList />
                </section>
            </div>
        </div>
    );
}
