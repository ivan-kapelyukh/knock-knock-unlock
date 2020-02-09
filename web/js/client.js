export async function login(user, knocks) {
  const url = "/login";

  const data = {
    user: user,
    knocks: knocks
  };

  const response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  });

  return await response.json();
}
