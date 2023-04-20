import { createSignal, onMount } from "solid-js";
import type { Component } from 'solid-js';

import logo from './logo.svg';
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
        <div class={styles.App}>
            <header class={styles.header}>
                <img src={logo} class={styles.logo} alt="logo" />
                <p>
                    Edit <code>src/App.tsx</code> and save to reload.
                    How about this?
                </p>
                <p class="text-4xl text-red-600">
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
            </header>
        </div>
    );
};

export default App;
