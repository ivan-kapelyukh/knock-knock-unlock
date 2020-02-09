class Client {
  constructor() {}

  async authenticate(user, knocks) {
    const url = "..../login";

    const data = {
      user: user,
      knocks: knocks
    };

    const response = await fetch(url, {
      method: "POST",
      mode: "cors", // TODO: Check this
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    });

    return await response.json();
  }
}
