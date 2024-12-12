import { useEffect, useState } from 'react';
import { Button, FileInput, Label, Modal, Textarea } from "flowbite-react";
import './globals.css';
import Router from 'next/router';

interface Blog {
  id: string;
  content: string;
  caption: string;
}

const UserBlog: React.FC = () => {
  const [blogs, setBlogs] = useState<Blog[]>([]);
  const [openModal, setOpenModal] = useState(false);
  const [selectedBlog, setSelectedBlog] = useState<Blog | null>(null);
  const [content, setContent] = useState<File | string>('');
  const [caption, setCaption] = useState('');

  function onCloseModal() {
    setOpenModal(false);
    setSelectedBlog(null);
    setContent('');
    setCaption('');
  }

  useEffect(() => {
    const fetchUserBlog = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/blog/blog', {
          method: 'GET',
          credentials: 'include',
        });

        if (!response.ok) {
          throw new Error(`Failed to fetch posts: ${response.status}`);
        }

        const data = await response.json();
        setBlogs(data.data || []);
      } catch (error) {
        console.error('Error fetching posts:', error);
      }
    };

    fetchUserBlog();
  }, []);

  const handleDeleteBlog = async (blogId: string) => {
    try {
      const response = await fetch(`http://localhost:8080/api/blog/blog/${blogId}`, {
        method: 'DELETE',
        credentials: 'include',
      });

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        console.error(`Failed to delete blog:`, errorData.message || 'Unknown error');
        return;
      }

      setBlogs((prevBlogs) => prevBlogs.filter((blog) => blog.id !== blogId));
      console.log("Blog deleted successfully.");
    } catch (error) {
      console.error('Delete blog failed:', error);
    }
  };

  const handleEditPost = async () => {
    if (!selectedBlog) return;
  
    try {
      let updatedData: { caption: string; content?: string } = { caption };
  
      if (content instanceof File) {
        const formData = new FormData();
        formData.append("content", content);
  
        const uploadResponse = await fetch(
          `http://localhost:8080/api/blog/update/${selectedBlog.id}`,
          {
            method: "POST",
            body: formData,
            credentials: "include",
          }
        );
  
        if (!uploadResponse.ok) {
          throw new Error("Failed to upload file");
        }
  
        const uploadData = await uploadResponse.json();
        updatedData.content = uploadData.content;
      }
  
      const response = await fetch(
        `http://localhost:8080/api/blog/update/${selectedBlog.id}`,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
          body: JSON.stringify(updatedData),
        }
      );
  
      if (!response.ok) {
        const errorData = await response.json();
        console.error(`Failed to edit blog:`, errorData.message || "Unknown error");
        return;
      }
  
      setBlogs((prevBlogs) =>
        prevBlogs.map((blog) =>
          blog.id === selectedBlog.id
            ? { ...blog, ...updatedData }
            : blog
        )
      );
  
      console.log("Blog edited successfully.");
      onCloseModal();
    } catch (error) {
      console.error("Edit blog failed:", error);
    }
    Router.reload();
  };
  

  return (
    <>
      {blogs.length > 0 ? (
        <div className="post-list">
          <div className="py-5 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 mb-4">
            {blogs.map((blog) => (
              <div key={blog.id} className="post-item">
                {blog.content && (
                  <figure className="max-w-full">
                    <div className="w-full aspect-[4/3] overflow-hidden rounded-lg">
                      <img
                        className="w-full h-full object-cover"
                        src={typeof blog.content === 'string' ? `http://localhost:8080/api/blog/${blog.content}` : ''} 
                        alt="Blog Content"
                        />
                    </div>
                    <figcaption className="mt-2 text-sm text-center text-black-500 dark:text-black-500 line-clamp-2">
                      <strong>Description: </strong>{blog.caption}
                    </figcaption>
                  </figure>
                )}

                <div className="button mt-4 flex justify-center space-x-2">
                  <button
                    className="w-24 px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-1 dark:bg-blue-500 dark:hover:bg-blue-600"
                    type="button"
                    onClick={() => {
                      setSelectedBlog(blog);
                      setContent(blog.content);
                      setCaption(blog.caption);
                      setOpenModal(true);
                    }}
                  >
                    Edit
                  </button>

                  <Modal show={openModal} size="md" onClose={onCloseModal} popup>
                    <Modal.Header />
                    <Modal.Body>
                      <div className="space-y-6">
                        <h3 className="text-xl font-medium text-gray-900 dark:text-white">Edit Your Post</h3>
                        <div>
                          <div className="mb-2 block">
                            <Label htmlFor="content" value="Your Photo" />
                          </div>
                          <FileInput
                            id="content"
                            accept=".jpg,.jpeg,.png,.gif"
                            onChange={(e) => setContent(e.target.files ? e.target.files[0] : "")}
                            required
                          />
                        </div>
                        <div>
                          <div className="mb-2 block">
                            <Label htmlFor="description" value="Description Photo" />
                          </div>
                          <Textarea
                            id="description"
                            value={caption}
                            onChange={(event) => setCaption(event.target.value)}
                            required
                          />
                        </div>
                        <div className="w-full">
                          <Button onClick={handleEditPost}>Confirm Edit</Button>
                        </div>
                      </div>
                    </Modal.Body>
                  </Modal>

                  <button
                    type="button"
                    className="w-24 px-4 py-2 text-sm font-medium text-white bg-red-600 rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-1 dark:bg-red-500 dark:hover:bg-red-600"
                    onClick={() => handleDeleteBlog(blog.id)}
                  >
                    Delete
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      ) : (
        <p>No posts found</p>
      )}
    </>
  );
};

export default UserBlog;
