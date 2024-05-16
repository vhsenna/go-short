async function shortenURL() {
  const url = document.getElementById("urlInput").value;
  const response = await fetch("/shorten", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({url}),
  });
  const data = await response.json();
  const shortUrlContainer = document.getElementById("shortUrlContainer");
  if (response.ok) {
    const shortUrl = `${window.location.origin}/r/${data.short_url}`;
    shortUrlContainer.innerHTML = `<a href="${shortUrl}" target="_blank">${shortUrl}</a></p>`;
  } else {
    shortUrlContainer.innerHTML = `<p>Error: ${data.error}</p>`;
  }
}
