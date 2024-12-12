import { useEffect, useState } from 'react';
import './globals.css';

interface Blog {
  id: string;
  content: string;
  caption: string;
  user_name: string;
  picture: string;
}

const DashboardBlog: React.FC = () => {
  const [blogs, setBlogs] = useState<Blog[]>([]);

  useEffect(() => {
    const fetchAllUserBlog = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/blog/blogs', {
          method: 'GET',
        });

        if (!response.ok) {
          throw new Error(`Failed to fetch posts: ${response.status}`);
        }

        const data = await response.json();
        setBlogs(data.data || []); // Gunakan fallback jika data kosong
        console.log(data);
      } catch (error) {
        console.error('Error fetching posts:', error);
      }
    };

    fetchAllUserBlog();
  }, []);

  return (
    <div>
      {blogs.length > 0 ? (
        <div className="post-list">
          {/* Grid responsif */}
          <div className="d-flex justify-center align-item">
            {blogs.map((blog) => (
              <div key={blog.id} className="post-item">
                {blog.content && (
                  <figure className="max-w-full">
                    <div className="max-w-[500px] max-h-[500px] w-auto h-auto aspect-square overflow-hidden rounded-lg mx-auto">
                      <div className="flex items-center py-5">
                        <img
                          src={`http://localhost:8080/api/blog/${blog.picture}`}
                          alt={`${blog.user_name || 'User'}'s profile`}
                          className="user-picture"
                          style={{
                            width: '50px',
                            height: '50px',
                            borderRadius: '50%',
                            objectFit: 'cover',
                          }}
                        />
                        <p className="ml-3 font-medium">{blog.user_name}</p>
                      </div>
                    {/* Gambar responsif */}
                      <img
                        className="w-full h-full object-cover"
                        src={typeof blog.content === 'string' ? `http://localhost:8080/api/blog/${blog.content}` : ''}
                        alt="image description"
                      />
                    </div>
                    <figcaption className="max-w-[500px] max-h-[500px] py-5 mx-auto">
                      <strong>{blog.user_name} </strong> {blog.caption}
                    </figcaption>
                  </figure>
                )}
              </div>
            ))}
          </div>
        </div>
      ) : (
        <p>No posts found</p>
      )}
    </div>
  );
};

export default DashboardBlog;