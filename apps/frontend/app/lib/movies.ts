export type Movie = { movieId: number; movieName: string; uploadedBy: string; movieLink: string; imageUrl: string; trailorUrl: string };
const titles = ["Midnight Run", "Red Horizon", "The Last Signal", "Neon City", "Silent River", "After the Storm", "Blackbird", "Glass House", "The Long Way Home", "Echoes", "Fireline", "The Arrival", "Lost & Found", "North Star", "The Night Shift", "Paper Planes", "Wild Hearts", "Second Chance", "Blue Room", "Parallel", "The Crossing", "Golden Hour", "Deep Water", "The Visitor", "Hidden Truth", "Summer Rain", "Broken Road", "The Promise", "City Lights", "Final Chapter"];
const trailers = ["d9MyW72ELq0", "EXeTwQWrcwY", "YoHD9XEInc0", "TcMBFSGVi1c", "zSWdZVtXT7E", "6ZfuNTqbHE8"];
export const demoMovies: Movie[] = titles.map((movieName, index) => ({ movieId: -(index + 1), movieName, uploadedBy: "Movie Charm", movieLink: "https://www.youtube.com", imageUrl: `https://picsum.photos/seed/movie-charm-${index + 1}/700/1000`, trailorUrl: `https://www.youtube.com/watch?v=${trailers[index % trailers.length]}` }));
export function youtubeEmbed(url: string) { const id = url.match(/(?:youtu\.be\/|v=|embed\/)([^?&/]+)/)?.[1]; return id ? `https://www.youtube.com/embed/${id}?autoplay=1&mute=1&rel=0` : ""; }

export type MovieComment = { comment: string; sentiment?: string; message?: string };
const thoughts = ["The pacing pulled me in right away. I would happily watch this again.", "Beautiful atmosphere and a strong ending. This one stayed with me.", "A really enjoyable watch — the soundtrack was the best part for me."];
export function demoComments(movieId: number): MovieComment[] { return thoughts.map((comment, index) => ({ comment: `${comment} — viewer ${Math.abs(movieId) * 3 + index + 1}` })); }
