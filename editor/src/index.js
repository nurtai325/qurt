import { createElement } from 'react';
import { createRoot } from 'react-dom/client';

const root = createRoot(document.getElementById("root"));
root.render(createElement("p", undefined, "hello"));
