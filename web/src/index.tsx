/* @refresh reload */
import { render } from 'solid-js/web';
import { Router, Route, Routes, Link, A } from "@solidjs/router";

import './index.css';
import App from './App';
import Shares from './Shares';

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
    throw new Error(
        'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got mispelled?',
    );
}

render(
    () =>
        <Router>
            <App />
        </Router>,
    root!
);
