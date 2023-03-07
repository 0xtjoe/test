import { useState } from 'react'
import Head from 'next/head'
import styles from '@/styles/Home.module.css'

export default function Home() {
  const [count, setCount] = useState(0)
  const [msg, setMsg] = useState("")

  const doClick = async () => {
    // increase current count
    const newCnt = count + 1;

    // send request to server
    const respData = await fetch('/api/fizzbuzz', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        count: newCnt
      }),
    })
    const respResult = await respData.json()

    // update states
    setCount(newCnt)
    setMsg(respResult.message)
  }

  return (
    <>
      <Head>
        <title>Test Frontend</title>
        <meta name="description" content="test frontend" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={styles.main}>
        <p>Your Count</p>
        <p>{count}</p>
        <div>
          <button className={styles.clickBtn} onClick={() => doClick()}>Push me!</button>
        </div>
        {
          msg.length == 0 ? null :
            <p className={styles.msgText}>{msg}</p>
        }
      </main>
    </>
  )
}
