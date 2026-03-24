# Prerequisites Workshop Intern Back End Department BCC 2026 - Object Storage & API Documentation

### 1. Project Repository

---

This workshop uses the project repository specifically prepared for learning Object Storage and API Documentation.

Clone the repo using the command below:

```bash
git clone https://github.com/Ablebil/workshop-storage-api-docs.git
```

Navigate into the project directory:

```bash
cd workshop-storage-api-docs
```

Ensure you are on the `main` branch:

```bash
git checkout main
```

Install the dependencies needed by the project:

```bash
go mod tidy
```

Next, copy `.env.example` to `.env` and fill it with the environment variables from the previous workshop:

```env
DB_NAME=
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=

APP_PORT=

JWT_SECRET_KEY=
JWT_EXPIRED_TIME=

GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
GOOGLE_REDIRECT_URL=
```

> **Note:** Make sure your PostgreSQL database is running (you can use `docker compose up -d` if you prefer using the provided Docker configuration).

### 2. Supabase Account & Project Setup

---

To implement Object Storage, we will use Supabase. You need to create an account and set up a new project.

#### Step 1 — Open Supabase

Open the following page:

```text
https://supabase.com/
```

Click **Start your project** or **Sign In** (it is recommended to continue with your GitHub account).

#### Step 2 — Register / Sign In

Complete the registration process until you reach the Supabase Dashboard.

#### Step 3 — Create a New Project

a. Click **New Project** (if prompted, create an Organization first).
b. Fill in the required fields:

- **Organization**: Choose the organization you previously created.
- **Name**: `workshop-4`
- **Database Password**: Click `Generate a password` button (save this just in case, though we won't strictly use the DB feature for this workshop).
- **Region**: Choose the region closest to you (e.g., `Singapore`).
- **Security**: Make sure the `Enable data API` box is checked.

c. Click **Create new project**.

Wait a few minutes for the project to be fully provisioned. Once it's ready, you are good to go! We will set up the storage buckets and API keys together during the workshop.

### 3. Install Postman

---

To test our endpoints and learn how to create API Documentation, we will use Postman.

If you don't have it installed yet, please download and install it from the official website:

```text
https://www.postman.com/downloads/
```

Create a free Postman account and sign in to the app so you can save your collections.
