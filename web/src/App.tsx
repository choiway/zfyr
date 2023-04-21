import { createSignal, onMount } from "solid-js";
import type { Component } from 'solid-js';
import { Routes, Route, Link, A } from "@solidjs/router";

import Shares from "./Shares";
import styles from './App.module.css';

// const fetchUser = async () =>
//     (await fetch(`http://localhost:1323/api`)).json();

const App: Component = () => {
    const [greeting, setGreeting] = createSignal<string>();

    onMount(async () => {
        const res = await fetch(`http://192.168.10.57:1323/api`);
        setGreeting(await res.text());
    });


    return (
        <>
            <p class="text-4xl text-green-600">
                This is the money
            </p>

            <p>{greeting()}</p>
            <a
                class={styles.link}
                href="https://github.com/solidjs/solid"
                target="_blank"
                rel="noopener noreferrer"
            >
                Learn Solid
            </a>

            <div>
                <A href="/shares">Link to Shares</A>
            </div>

            <div>
                <Link href="/about">Link to About</Link>
            </div>

            <Routes>
                <Route
                    path="/"
                    element={<div>Home page</div>}
                />
                <Route path="/shares" component={Shares} />
                <Route
                    path="/about"
                    element={<div>This site was made with Solid</div>}
                />
            </Routes>
        </>
    );
};

export default App;
