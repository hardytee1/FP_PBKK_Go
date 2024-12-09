import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';

const Account: React.FC = () => {
  const router = useRouter();
  const [notification, setNotification] = useState<{ type: string; message: string } | null>(null);
  const [user, setUser] = useState<{ name: string; email: string } | null>(null);
  const [newName, setNewName] = useState(''); // State to hold the new name

  useEffect(() => {
    const fetchUserDetails = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/user/me', {
          method: 'GET',
          credentials: 'include',
        });

        if (response.ok) {
          const data = await response.json();
          setUser(data.data);
          setNewName(data.data.name); // Pre-fill the name input with the current name
        } else {
          const errorData = await response.json();
          setNotification({ type: 'error', message: `Failed to fetch user details: ${errorData.message}` });
        }
      } catch (error) {
        console.error('Fetch user details failed:', error);
        setNotification({ type: 'error', message: 'Failed to fetch user details. Please try again later.' });
      }
    };

    fetchUserDetails();
  }, []);

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

  const handleUpdateName = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/user/update', {
        method: 'PUT',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: newName }),
      });

      if (response.ok) {
        setNotification({ type: 'success', message: 'Name updated successfully.' });
        if (user) setUser({ ...user, name: newName }); // Update the name in the state
      } else {
        const errorData = await response.json();
        setNotification({ type: 'error', message: `Failed to update name: ${errorData.message}` });
      }
    } catch (error) {
      console.error('Update name failed:', error);
      setNotification({ type: 'error', message: 'Failed to update name. Please try again later.' });
    }
  };

  return (
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        minHeight: '100vh',
        background: '#f4f6f9',
        padding: '20px',
      }}
    >
      <div
        style={{
          background: 'white',
          padding: '30px',
          borderRadius: '8px',
          boxShadow: '0 4px 10px rgba(0, 0, 0, 0.1)',
          maxWidth: '600px',
          width: '100%',
          textAlign: 'center',
        }}
      >
        <h1
          style={{
            color: '#333',
            fontFamily: 'Arial, sans-serif',
            marginBottom: '20px',
            fontSize: '24px',
          }}
        >
          Manage Your Account
        </h1>
        {notification && (
          <div
            style={{
              color: notification.type === 'error' ? 'red' : 'green',
              marginBottom: '20px',
              fontWeight: 'bold',
            }}
          >
            {notification.message}
          </div>
        )}
        {user ? (
          <div
            style={{
              marginBottom: '20px',
              background: '#f9f9f9',
              padding: '20px',
              borderRadius: '8px',
              boxShadow: '0 2px 5px rgba(0, 0, 0, 0.1)',
            }}
          >
            <p style={{ margin: '10px 0', fontSize: '16px' }}><strong>Name:</strong> {user.name}</p>
            <p style={{ margin: '10px 0', fontSize: '16px' }}><strong>Email:</strong> {user.email}</p>
          </div>
        ) : (
          <p>Loading user details...</p>
        )}
        <div style={{ marginBottom: '30px' }}>
          <label
            style={{
              display: 'block',
              marginBottom: '10px',
              fontSize: '16px',
              color: '#555',
            }}
          >
            Change Name
          </label>
          <input
            type="text"
            value={newName}
            onChange={(e) => setNewName(e.target.value)}
            placeholder="Enter new name"
            style={{
              padding: '10px',
              width: '100%',
              marginBottom: '20px',
              border: '1px solid #ccc',
              borderRadius: '4px',
              fontSize: '16px',
            }}
          />
          <button
            onClick={handleUpdateName}
            style={{
              color: 'white',
              backgroundColor: '#28a745',
              border: 'none',
              padding: '12px 20px',
              cursor: 'pointer',
              fontSize: '16px',
              borderRadius: '5px',
              width: '100%',
            }}
          >
            Update Name
          </button>
        </div>
        <button
          onClick={handleDeleteAccount}
          style={{
            color: 'white',
            backgroundColor: '#dc3545',
            border: 'none',
            padding: '12px 20px',
            cursor: 'pointer',
            fontSize: '16px',
            borderRadius: '5px',
            width: '100%',
            marginBottom: '20px',
          }}
        >
          Delete Account
        </button>
        <button
          onClick={() => router.push('/dashboard')}
          style={{
            color: 'white',
            backgroundColor: '#007bff',
            border: 'none',
            padding: '12px 20px',
            cursor: 'pointer',
            fontSize: '16px',
            borderRadius: '5px',
            width: '100%',
          }}
        >
          Back to Dashboard
        </button>
      </div>
    </div>
  );
};

export default Account;
