// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import axios from 'axios'

type Data = {
    status: number,
    message: string
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
    const { count } = req.body

    try {
        // send request to /fizzbuzz endpoint
        const countRequest = await axios.post(
            `${process.env.NEXT_PUBLIC_API_BASE}/fizzbuzz`,
            { count },
            {
                headers: {
                    'Content-Type': 'application/json'
                }
            }
        )
        const countResp = countRequest.data;

        // return with returned value
        return res.status(200).json({
            status: countResp.status,
            message: countResp.message
        })
    } catch (err) {
        // console the error
        console.log(`err here: ${err}`)
        return res.json({ message: "Unexpected error", status: 500 })
    }
}
