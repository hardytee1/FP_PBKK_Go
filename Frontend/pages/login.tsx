import { useState } from 'react';
import { login } from '../services/api';
import { useRouter } from 'next/router';

const Login: React.FC = () => {
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await login({ email, password });
      if (response.success) {
        console.log('Redirecting to dashboard...');
        router.push('/profilePage');
      } else {
        setError(response.message);
      }
    } catch (err) {
      setError('Something went wrong.');
    }
  };

  return (
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        minHeight: '100vh',
        background: '#f7f7f7',
        padding: '20px',
      }}
    >
      <div
        style={{
          background: 'white',
          padding: '30px',
          borderRadius: '8px',
          boxShadow: '0 4px 10px rgba(0, 0, 0, 0.1)',
          maxWidth: '400px',
          width: '100%',
        }}
      >
        <h1
          style={{
            textAlign: 'center',
            color: '#333',
            fontFamily: 'Arial, sans-serif',
            marginBottom: '20px',
          }}
        >
          Login
        </h1>
        <form onSubmit={handleSubmit}>
          <div style={{ marginBottom: '15px' }}>
            <input
              type="email"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              style={{
                width: '100%',
                padding: '12px',
                border: '1px solid #ccc',
                borderRadius: '6px',
                fontSize: '16px',
                outline: 'none',
                boxSizing: 'border-box',
                transition: 'border-color 0.3s',
              }}
            />
          </div>
          <div style={{ marginBottom: '25px' }}>
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              autoComplete="current-password"
              required
              style={{
                width: '100%',
                padding: '12px',
                border: '1px solid #ccc',
                borderRadius: '6px',
                fontSize: '16px',
                outline: 'none',
                boxSizing: 'border-box',
                transition: 'border-color 0.3s',
              }}
            />
          </div>
          <button
            type="submit"
            style={{
              width: '100%',
              padding: '12px',
              backgroundColor: '#007bff',
              color: 'white',
              border: 'none',
              borderRadius: '6px',
              fontSize: '16px',
              cursor: 'pointer',
              transition: 'background-color 0.3s',
            }}
          >
            Login
          </button>
        </form>
        {error && (
          <p
            style={{
              color: 'red',
              fontSize: '14px',
              textAlign: 'center',
              marginTop: '15px',
            }}
          >
            {error}
          </p>
        )}
        <p
          style={{
            textAlign: 'center',
            marginTop: '20px',
            fontSize: '14px',
            color: '#555',
          }}
        >
          Don't have an account?{' '}
          <span
            onClick={() => router.push('/register')}
            style={{
              color: '#007bff',
              cursor: 'pointer',
              textDecoration: 'underline',
            }}
          >
            Register here
          </span>
        </p>
      </div>
    </div>
  );
};

export default Login;
