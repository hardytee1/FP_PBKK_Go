import Link from 'next/link';
import styles from '../styles/Home.module.css';

const Home: React.FC = () => {
  return (
    <div className={styles.container}>
      <h1>Welcome to Our App</h1>
      <div className={styles.buttons}>
        <Link href="/register">
          <button className={styles.button}>Register</button>
        </Link>
        <Link href="/login">
          <button className={styles.button}>Login</button>
        </Link>
      </div>
    </div>
  );
};

export default Home;
