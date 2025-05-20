import React from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import DynamicDialog from "./_DynamicDialog";
import type { DialogSet } from "@/utils";

const queryClient = new QueryClient();

type Props = {
    dSet: DialogSet;
};

export default function DynamicDialogWrapper({
    dSet,
}: Props): React.JSX.Element {
    return (
        <QueryClientProvider client={queryClient}>
            <DynamicDialog dSet={dSet} />
        </QueryClientProvider>
    );
}
