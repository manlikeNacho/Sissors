(() => {
  "use strict";

  // ---------- Tabs ----------
  const tabButtons = document.querySelectorAll(".tab-btn");
  const panels = {
    shorten: document.getElementById("panel-shorten"),
    lookup: document.getElementById("panel-lookup"),
    account: document.getElementById("panel-account"),
  };

  tabButtons.forEach((btn) => {
    btn.addEventListener("click", () => {
      tabButtons.forEach((b) => {
        b.classList.remove("active");
        b.setAttribute("aria-selected", "false");
      });
      Object.values(panels).forEach((p) => p.classList.remove("active"));

      btn.classList.add("active");
      btn.setAttribute("aria-selected", "true");
      panels[btn.dataset.tab].classList.add("active");
    });
  });

  // ---------- Toasts ----------
  const toastStack = document.getElementById("toast-stack");
  function toast(message, type = "info") {
    const el = document.createElement("div");
    el.className = `toast ${type}`;
    el.textContent = message;
    toastStack.appendChild(el);
    setTimeout(() => el.remove(), 4200);
  }

  // ---------- Helpers ----------
  function setLoading(button, isLoading) {
    const text = button.querySelector(".btn-text");
    const spinner = button.querySelector(".spinner");
    button.disabled = isLoading;
    if (spinner) spinner.hidden = !isLoading;
    if (text) text.style.opacity = isLoading ? "0.6" : "1";
  }

  async function parseJsonSafely(res) {
    try {
      return await res.json();
    } catch {
      return null;
    }
  }

  function extractShortCode(input) {
    const trimmed = input.trim();
    try {
      const url = new URL(trimmed);
      const parts = url.pathname.split("/").filter(Boolean);
      return parts[parts.length - 1] || trimmed;
    } catch {
      // Not a full URL — treat as a bare code, stripping a leading short_url/ segment if pasted.
      return trimmed.replace(/^\/?short_url\//, "");
    }
  }

  // ---------- Shorten ----------
  const shortenForm = document.getElementById("shorten-form");
  const shortenSubmit = document.getElementById("shorten-submit");
  const shortenResult = document.getElementById("shorten-result");
  const resultShort = document.getElementById("result-short");
  const resultMeta = document.getElementById("result-meta");
  const copyBtn = document.getElementById("copy-btn");
  const openBtn = document.getElementById("open-btn");

  let lastShortCode = "";
  let lastOriginalUrl = "";

  shortenForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const longUrl = document.getElementById("long-url").value.trim();
    const customCode = document.getElementById("custom-code").value.trim();

    setLoading(shortenSubmit, true);
    shortenResult.hidden = true;

    try {
      const res = await fetch("/url", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ Url: longUrl, ShortUrl: customCode }),
      });
      const data = await parseJsonSafely(res);

      if (!res.ok) {
        const msg = (data && (data.error?.message || data.error)) || "Could not shorten that URL.";
        toast(typeof msg === "string" ? msg : JSON.stringify(msg), "error");
        return;
      }

      lastShortCode = data.short_url;
      lastOriginalUrl = data.orignal_url || longUrl;

      const shortLink = `${window.location.origin}/short_url/${lastShortCode}`;
      resultShort.value = shortLink;
      resultMeta.textContent = `→ ${lastOriginalUrl}`;
      shortenResult.hidden = false;
      toast("Short URL created", "success");
    } catch (err) {
      toast("Network error — is the server running?", "error");
    } finally {
      setLoading(shortenSubmit, false);
    }
  });

  copyBtn.addEventListener("click", async () => {
    try {
      await navigator.clipboard.writeText(resultShort.value);
      toast("Copied to clipboard", "success");
    } catch {
      resultShort.select();
      document.execCommand("copy");
      toast("Copied to clipboard", "success");
    }
  });

  openBtn.addEventListener("click", () => resolveAndOpen(lastShortCode, lastOriginalUrl));

  async function resolveAndOpen(code, fallbackUrl) {
    try {
      const res = await fetch(`/short_url/${encodeURIComponent(code)}`);
      const data = await parseJsonSafely(res);
      const target = (data && data.Url) || fallbackUrl;
      if (target) {
        window.open(target, "_blank", "noopener");
      } else {
        toast("Could not resolve that short link.", "error");
      }
    } catch {
      toast("Network error while resolving link.", "error");
    }
  }

  // ---------- Lookup ----------
  const lookupForm = document.getElementById("lookup-form");
  const lookupSubmit = document.getElementById("lookup-submit");
  const lookupResult = document.getElementById("lookup-result");
  const lookupOriginal = document.getElementById("lookup-original");
  const lookupOpenBtn = document.getElementById("lookup-open-btn");

  let lookupTarget = "";

  lookupForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const raw = document.getElementById("lookup-code").value;
    const code = extractShortCode(raw);

    setLoading(lookupSubmit, true);
    lookupResult.hidden = true;

    try {
      const res = await fetch(`/short_url/${encodeURIComponent(code)}`);
      const data = await parseJsonSafely(res);

      if (!data || !data.Url) {
        toast("Short URL doesn't exist.", "error");
        return;
      }

      lookupTarget = data.Url;
      lookupOriginal.value = data.Url;
      lookupResult.hidden = false;
    } catch {
      toast("Network error — is the server running?", "error");
    } finally {
      setLoading(lookupSubmit, false);
    }
  });

  lookupOpenBtn.addEventListener("click", () => {
    if (lookupTarget) window.open(lookupTarget, "_blank", "noopener");
  });

  // ---------- Signup ----------
  const signupForm = document.getElementById("signup-form");
  const signupSubmit = document.getElementById("signup-submit");

  signupForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const payload = {
      first_name: document.getElementById("su-first").value.trim(),
      last_name: document.getElementById("su-last").value.trim(),
      email: document.getElementById("su-email").value.trim(),
      phone: Number(document.getElementById("su-phone").value) || 0,
      password: document.getElementById("su-password").value,
      user_type: document.getElementById("su-type").value,
    };

    setLoading(signupSubmit, true);
    try {
      const res = await fetch("/user", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });
      const data = await parseJsonSafely(res);

      if (!res.ok) {
        const msg = (data && (data.error?.message || data.error)) || "Signup failed.";
        toast(typeof msg === "string" ? msg : JSON.stringify(msg), "error");
        return;
      }

      toast("Account created — you can log in now.", "success");
      signupForm.reset();
    } catch {
      toast("Network error — is the server running?", "error");
    } finally {
      setLoading(signupSubmit, false);
    }
  });

  // ---------- Login ----------
  const loginForm = document.getElementById("login-form");
  const loginSubmit = document.getElementById("login-submit");

  loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const payload = {
      email: document.getElementById("li-email").value.trim(),
      password: document.getElementById("li-password").value,
    };

    setLoading(loginSubmit, true);
    try {
      const res = await fetch("user/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });
      const data = await parseJsonSafely(res);
      if (res.ok) {
        toast((data && data.message) || "Logged in", "success");
      } else {
        toast("Login failed.", "error");
      }
    } catch {
      toast("Network error — is the server running?", "error");
    } finally {
      setLoading(loginSubmit, false);
    }
  });
})();
