import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import CreateBlogForm from './CreateBlog';
import UserBlog from './GetUserBlog';

const Dashboard: React.FC = () => {
  const router = useRouter();
  const [loading, setLoading] = useState(true);
  const [userName, setUserName] = useState<string | null>(null); // State to store user's name
  const [userPicture, setUserPicture] = useState<string | null>(null); // State to store user's picture URL

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

  const handleManageAccount = () => {
    router.push('/account');
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
          const data = await response.json();
          setUserName(data.data.name); // Set the user's name
          setUserPicture(data.data.picture); // Set the user's picture URL
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
      <div className="welcome-section">
        {userPicture && (
          <img
          src={`http://localhost:8080/api/blog/${userPicture}`}
            alt={`${userName || 'User'}'s profile`}
            className="user-picture"
            style={{
              width: '100px',
              height: '100px',
              borderRadius: '50%',
              objectFit: 'cover',
            }}
          />
        )}
        <h1>Welcome back, {userName || 'User'}!</h1>
      </div>
      <p>This is your private space where you can manage your account and features.</p>
      <button onClick={handleLogout}>Logout</button>
      <button onClick={handleManageAccount}>Manage Account</button>
      <div className="UploadBlog">
        <CreateBlogForm />
      </div>
      <div className="widgets">
        <UserBlog />
      </div>
    </div>
  );
};

export default Dashboard;
