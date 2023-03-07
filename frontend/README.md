# React/Typescript frontend for test task

### I have used [React/Next.js](https://nextjs.org/blog/next-13) for this test task

You need to fill out the value in .env file.
Please check [.env.development](https://github.com/jamesmern/test/blob/master/frontend/.env.development) file.
```bash
NEXT_PUBLIC_API_BASE=YOUR_BACKEND_API_URL
```

After fill out .env, you can run frontend by running
```bash
yarn install
yarn dev
```

### To avoid cross-origin, I've added api calling code in pages/api/fizzbuzz.ts
```code
const countRequest = await axios.post(
    `${process.env.NEXT_PUBLIC_API_BASE}/fizzbuzz`,
    { count },
    {
        headers: {
            'Content-Type': 'application/json'
        }
    }
)
```

### Currently, frontend is running on 3000 port and we can update the port as we want.
