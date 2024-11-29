import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';

const Dashboard: React.FC = () => {
  const router = useRouter();
  const [loading, setLoading] = useState(true);

  const handleLogout = async () => {
    try {
      await fetch('http://localhost:8080/api/user/logout', {
        method: 'POST',
        credentials: 'include', // Ensure cookies are sent
      });
      router.push('/login'); // Redirect to login page
    } catch (error) {
      console.error('Logout failed:', error);
    }
  };

  useEffect(() => {
    const checkAuth = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/user/me', {
          method: 'GET',
          credentials: 'include', // Include cookies in the request
        });

        if (!response.ok) {
          // Redirect to login if not authenticated
          router.push('/login');
        } else {
          setLoading(false); // User is authenticated
        }
      } catch (error) {
        console.error('Auth check failed:', error);
        router.push('/login');
      }
    };

    checkAuth();
  }, [router]);

  if (loading) {
    return <p>Loading...</p>;
  }

  return (
    <div className="dashboard">
      <h1>Welcome to the Dashboard</h1>
      <p>This is your private space where you can manage your account and features.</p>
      <button onClick={handleLogout}>Logout</button>
      <div className="widgets">
        <div className="widget">Widget 1</div>
        <div className="widget">Widget 2</div>
        <div className="widget">Widget 3</div>
      </div>
    </div>
  );
};

export default Dashboard;
