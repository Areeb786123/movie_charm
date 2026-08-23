import "./globals.css";

export const metadata = { title: "Movie Charm", description: "Your movie shelf" };

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return <html lang="en"><body>{children}</body></html>;
}
