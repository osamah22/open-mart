# OpenMart Development Plan

## 1️⃣ Core Setup

- [ ] **Project Skeleton**
  - `main.go` bootstraps the app
  - App struct with loggers, templates, services, session manager
- [ ] **Middleware**
  - Logging middleware
  - Panic recovery
  - Request ID
- [ ] **Static Files**
  - Serve `/static/css`, `/static/js`, `/static/img`
- [ ] **Templates**
  - Base template (`base.html`)
  - Home page (`/`)
  - Success/failure notices
  - Error pages (404, 500)

- [ ] **App Helpers**
  - Render helpers for templates
  - Error handling helpers
    ✅ _Test:_ Start server → confirm homepage works, static files load, errors render.

---

## 2️⃣ Session & Flash Infrastructure

- [ ] **Session Manager**
  - Integrate `scs` session manager
- [ ] **Flash Messages**
  - Add helper to set/read flash from session
  - Display flash in base template
    ✅ _Test:_ Set a flash → ensure it appears once, then disappears.

---

## 3️⃣ Database & Service Layer

- [ ] **Migrations**
  - Users table
  - Posts table
  - Categories table
- [ ] **Services**
  - `AuthService`
  - `PostService`
  - `CategoryService`
- [ ] **Testing**
  - Use testcontainers
  - Write tests for signup/login, post create, category create
    ✅ _Test:_ Run integration tests → confirm services behave as expected.

---

## 4️⃣ Authentication UI

- [ ] **Routes**
  - `/signup`, `/login`, `/logout`
- [ ] **Forms & Templates**
  - Signup form
  - Login form
- [ ] **Integration**
  - Hook into `AuthService`
  - Add login-required middleware
    ✅ _Test:_ Signup → login → logout. Confirm flash messages show.

---

## 5️⃣ Posts & Categories UI

- [ ] **Post Routes**
  - `/posts`
  - `/posts/create`
  - `/posts/{id}`
- [ ] **Category Routes**
  - `/categories`
- [ ] **Templates**
  - List posts
  - Show post
  - Create post form
  - List categories
- [ ] **Ownership Check**
  - Only post owner can edit/delete
    ✅ _Test:_ User A creates a post → User B cannot edit/delete.

---

## 6️⃣ Polishing

- [ ] **Validation**
  - Required fields
  - Input trimming
- [ ] **CSS**
  - Minimal styling
    ✅ _Test:_ Try invalid forms → see correct error flash.

---

## 7️⃣ Final Touches

- [ ] **README**
  - Setup instructions
  - Usage instructions
- [ ] **Docker (Optional)**
  - Dockerfile
  - docker-compose for DB
- [ ] **Git Cleanup**
  - Merge feature branches into `main`
  - Delete temp branches
    ✅ _Test:_ Fresh clone → `go run main.go` → app works.
