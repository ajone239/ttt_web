export async function About() {
  await fetch('/api/board')
    .then(response => response.json())
    .then(data => {
      console.log(data);
    });
}
