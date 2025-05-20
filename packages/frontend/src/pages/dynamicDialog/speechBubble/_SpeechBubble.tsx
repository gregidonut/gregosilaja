import React from "react";
import styles from "./_speechBubble.module.css";

export default function SpeechBubble({
    dialog,
}: {
    dialog: string;
}): React.JSX.Element {
    return (
        <div className={styles.speechBubbleWrapper}>
            <p className={`${styles.bubble} ${styles.speech}`}>{dialog}</p>
        </div>
    );
}
