export function login(username, knock) {
  const url = "/login";
  const data = { username, knock };

  return fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  });
}
