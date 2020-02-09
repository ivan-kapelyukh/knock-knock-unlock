export async function login(username, knock) {
  const url = "/login";

  const data = { username, knock };

  const response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  });

  return await response.text();
}
