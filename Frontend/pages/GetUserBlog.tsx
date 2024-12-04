import { useEffect, useState } from 'react';

interface Blog {
  id: string;
  content: string;
  caption: string;
}

const UserBlog: React.FC = () => {
  const [blogs, setBlogs] = useState<Blog[]>([]);

  useEffect(() => {
    const fetchUserBlog = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/blog/blog', {
          method: 'GET',
          credentials: 'include', // Sertakan cookie Authorization
        });

        if (!response.ok) {
          throw new Error(`Failed to fetch posts: ${response.status}`);
        }

        const data = await response.json();
        setBlogs(data.data || []); // Gunakan fallback jika data kosong
        console.log(data)
      } catch (error) {
        console.error('Error fetching posts:', error);
      }
    };

    fetchUserBlog();
  }, []);

  return (
    <div>
      <h1>Your Posts</h1>
      {blogs.length > 0 ? (
        <div className="post-list">
          {blogs.map((blog) => (
            <div key={blog.id} className="post-item">
                {blog.content && <img src={blog.content} alt="Post image" style={{ width: '200px' }} />}
                <p>{blog.caption}</p>
            </div>
          ))}
        </div>
      ) : (
        <p>No posts found</p>
      )}
    </div>
  );
};

export default UserBlog;
