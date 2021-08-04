import { writable } from "svelte/store";

export const messages = writable([])

export const username = writable("")

export const userId = writable("")

export const server = writable()