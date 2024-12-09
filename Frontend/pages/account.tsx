import { useRouter } from 'next/router';
import { useState } from 'react';

const Account: React.FC = () => {
  const router = useRouter();
  const [notification, setNotification] = useState<{ type: string; message: string } | null>(null);

  const handleDeleteAccount = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/user/delete', {
        method: 'DELETE',
        credentials: 'include',
      });

      if (response.ok) {
        setNotification({ type: 'success', message: 'Account deleted successfully.' });
        setTimeout(() => {
          router.push('/register');
        }, 2000); // Redirect after 2 seconds
      } else {
        const errorData = await response.json();
        setNotification({ type: 'error', message: `Failed to delete account: ${errorData.message}` });
      }
    } catch (error) {
      console.error('Delete account failed:', error);
      setNotification({ type: 'error', message: 'Failed to delete account. Please try again later.' });
    }
  };

  return (
    <div style={{ maxWidth: '600px', margin: '0 auto', textAlign: 'center', padding: '20px' }}>
      <h1>Manage Your Account</h1>
      {notification && (
        <div
          style={{
            color: notification.type === 'error' ? 'red' : 'green',
            margin: '10px 0',
            fontWeight: 'bold',
          }}
        >
          {notification.message}
        </div>
      )}
      <button
        onClick={handleDeleteAccount}
        style={{
          color: 'white',
          backgroundColor: 'red',
          border: 'none',
          padding: '10px 20px',
          cursor: 'pointer',
          marginBottom: '10px',
        }}
      >
        Delete Account
      </button>
      <br />
      <button
        onClick={() => router.push('/dashboard')}
        style={{
          color: 'white',
          backgroundColor: '#007bff',
          border: 'none',
          padding: '10px 20px',
          cursor: 'pointer',
        }}
      >
        Back to Dashboard
      </button>
    </div>
  );
};

export default Account;
