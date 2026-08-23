"use client";

import { FormEvent, useEffect, useState } from "react";
import Link from "next/link";
import { demoMovies, Movie } from "./lib/movies";

const API = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";
async function request<T>(path: string, options?: RequestInit): Promise<T> { const r = await fetch(`${API}${path}`, { ...options, headers: { "Content-Type": "application/json", ...(options?.headers ?? {}) } }); if (!r.ok) { const b = await r.json().catch(() => null); throw new Error(b?.error ?? `Request failed (${r.status})`); } return r.json() as Promise<T>; }

export default function Home() {
  const [movies, setMovies] = useState<Movie[]>(demoMovies); const [error, setError] = useState("");
  const loadMovies = async () => { try { setMovies([...(await request<Movie[]>("/movies")), ...demoMovies]); } catch { /* Show demo catalogue while the API starts. */ } };
  useEffect(() => { loadMovies(); }, []);
  async function addMovie(e: FormEvent<HTMLFormElement>) { e.preventDefault(); setError(""); const f = new FormData(e.currentTarget); try { await request("/movies", { method: "POST", body: JSON.stringify({ movieName: f.get("movieName"), uploadedBy: f.get("uploadedBy"), movieLink: f.get("movieLink"), imageUrl: f.get("imageUrl"), trailorUrl: f.get("trailorUrl") }) }); e.currentTarget.reset(); await loadMovies(); } catch (err) { setError(err instanceof Error ? err.message : "Could not add movie."); } }
  return <main><header><p className="eyebrow">MOVIE CHARM · 30 STORIES</p><h1>Find something worth watching.</h1><p>A dark-red catalogue for movies, trailers and honest reactions.</p></header>{error && <p className="error">{error}</p>}<section className="add"><h2>Add a movie</h2><form onSubmit={addMovie}><input name="movieName" placeholder="Movie name" required /><input name="uploadedBy" placeholder="Your name" required /><input name="movieLink" type="url" placeholder="Movie URL" required /><input name="imageUrl" type="url" placeholder="Poster image URL" required /><input name="trailorUrl" type="url" placeholder="YouTube trailer URL" required /><button>Add to shelf</button></form></section><section><h2>Now showing</h2><div className="grid">{movies.map(movie => <Link key={movie.movieId} href={`/movies/${movie.movieId}`} className="card"><img src={movie.imageUrl || "https://placehold.co/600x900/251015/f8d5d5?text=Movie+Charm"} alt={`${movie.movieName} poster`} /><div><p className="eyebrow">TRAILER READY</p><h3>{movie.movieName}</h3><p>Added by {movie.uploadedBy}</p><span>Open movie →</span></div></Link>)}</div></section></main>;
}
