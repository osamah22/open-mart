# OpenMart HTML Project - Quick Finish Checklist

## 1️⃣ Core Features

- [ ] **Authentication**
  - Signup and login are working.
  - Ensure sessions work correctly (log in/out).
- [ ] **Post Management**
  - Create, read, update, delete posts.
  - Ensure only the owner can edit/delete their post.
- [ ] **Category Management**
  - Create and list categories.
  - Assign posts to categories.

## 2️⃣ Templates & UI

- [ ] **Base Template**
  - Header, footer, nav menu.
- [ ] **Flash/Error Messages**
  - Display validation or business errors.
- [ ] **Forms for Signup/Login/Post**
  - Keep simple; no need for fancy JS.

## 3️⃣ Tests

- [ ] **Service-Level Tests**
  - Keep existing signup/login tests.
  - Add minimal tests for posts/categories.
- [ ] **Optional Handler Tests**
  - Only if time allows; focus on core functionality first.

## 4️⃣ Misc / Finishing Touches

- [ ] **Static Files**
  - CSS, JS, images are served correctly.
- [ ] **Routing**
  - All URLs work (`/`, `/login`, `/signup`, `/posts`, `/categories`).
- [ ] **Session Management**
  - Users stay logged in; logout works.
- [ ] **Error Handling**
  - Keep simple: flash messages in HTML.

## 5️⃣ Optional Polishing (If Time Allows)

- Input trimming/validation in forms.
- Flash messages for success/failure.
- Simple pagination for posts.
- Minor CSS improvements.

---

💡 **Strategy**

- Focus on end-to-end working functionality first.
- Skip optional refactors like `SignupRequest` or fancy error wrapping.
- Once everything works, finish tests and polish the UI.
- After the project is done, move on to REST APIs, caching, or other advanced Go topics.
