import React, { useState, useEffect } from "react";
import { useQuery } from "@tanstack/react-query";
import type { DialogParts, DialogSet } from "@/utils";
import SpeechBubble from "./speechBubble/_SpeechBubble";

import styles from "./_dynamicDialog.module.css";

type Props = {
    dSet: DialogSet;
};

export default function DynamicDialog({ dSet }: Props): React.JSX.Element {
    const [curDSet, setCurDSet] = useState<DialogParts>(dSet[0]);
    const [fetchTriggered, setFetchTriggered] = useState(false);

    useEffect(function () {
        const firstTimer = setTimeout(function () {
            setCurDSet(dSet[1]);

            const secondTimer = setTimeout(function () {
                setCurDSet(dSet[2]);
                setFetchTriggered(true);
            }, 3000);

            return function () {
                clearTimeout(secondTimer);
            };
        }, 1000);

        return function () {
            clearTimeout(firstTimer);
        };
    }, []);
    const { data, isLoading } = useQuery<
        Array<{ q: string; a: string; h: string }>
    >({
        queryKey: ["quotes"],
        queryFn: async function () {
            const resp = await fetch("/api/randomQuote");
            return await resp.json();
        },
        enabled: fetchTriggered,
    });

    useEffect(
        function () {
            if (dSet.indexOf(curDSet) === 2) {
                const thirdTimer = setTimeout(function () {
                    setCurDSet(dSet[3]);
                }, 3000);
                return function () {
                    clearTimeout(thirdTimer);
                };
            }
            if (dSet.indexOf(curDSet) === 3) {
                const fourthTimer = setTimeout(function () {
                    setCurDSet(dSet[0]);
                }, 3000);
                return function () {
                    clearTimeout(fourthTimer);
                };
            }
            return;
        },
        [curDSet],
    );

    if (isLoading) {
        <div className={styles.dynamicDialog}>
            <img
                src={dSet[2].src}
                alt={dSet[2].alt}
                height={250}
                loading="lazy"
                decoding="async"
            />
            <SpeechBubble dialog={dSet[2].dialog} />
        </div>;
    }

    if (data) {
        dSet[0].dialog = `"${data[0]!.q}"\n -${data[0]!.a}`;
    }

    return (
        <div className={styles.dynamicDialog}>
            <img
                src={curDSet.src}
                alt={curDSet.alt}
                height={250}
                loading="lazy"
                decoding="async"
            />
            {curDSet.dialog !== "" ? (
                <SpeechBubble dialog={curDSet.dialog} />
            ) : null}
        </div>
    );
}
