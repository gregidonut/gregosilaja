import React from "react";
import footerFigures from "@/components/footer/footerFigures.json.ts";
import styles from "@/components/footer/footer.module.css";

export default function Footer(): React.JSX.Element {
    return (
        <div className={styles.footerContainer}>
            <footer className={styles.footer}>
                <article className={styles.iconArt}>
                    <h2>socials</h2>
                    <main>
                        {footerFigures.map(function (item, ind) {
                            return (
                                <figure key={ind}>
                                    <a href={item.url}>
                                        <img
                                            src={item.img.src}
                                            alt={item.alt}
                                            loading="lazy"
                                            decoding="async"
                                        />
                                    </a>
                                    <figcaption>
                                        <p>
                                            <a href={item.url}>{item.cap}</a>
                                        </p>
                                    </figcaption>
                                </figure>
                            );
                        })}
                    </main>
                </article>
                <article>
                    <h2>contact</h2>
                    <p>
                        <a href="mailto:gregosilaja@outlook.com">
                            gregosilaja@outlook.com
                        </a>
                    </p>
                    <p className={styles.copyright}>
                        &copy; 2025 gregosilaja.cc. All rights reserved.
                    </p>
                </article>
            </footer>
        </div>
    );
}
