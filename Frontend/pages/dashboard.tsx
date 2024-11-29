import { useRouter } from 'next/router';
import { useEffect } from 'react';

const Dashboard: React.FC = () => {
  const router = useRouter();

  const handleLogout = () => {
    // Clear authentication (e.g., token or session)
    localStorage.removeItem('token'); // Adjust as per your auth logic
    router.push('/login'); // Redirect to login page
  };

  useEffect(() => {
    // Example auth check (update with your backend/session logic)
    const token = localStorage.getItem('token');
    if (!token) {
      router.push('/login');
    }
  }, [router]);

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
