import React, { useState, ChangeEvent, FormEvent } from 'react';
import { useRouter } from 'next/router';

const CreateBlogForm: React.FC = () => {
  const [caption, setContent] = useState<string>('');
  const [image, setImage] = useState<File | null>(null);
  const router = useRouter();
//   const router = useRouter();

//   const checkAuth = async () => {
//     const response = await fetch('http://localhost:8080/api/user/me', {
//       method: 'GET',
//       credentials: 'include', // Sertakan cookie Authorization
//     });
  
//     if (!response.ok) {
//       alert('You need to log in first!');
//       router.push('/login');
//       return false;
//     }
  
//     return true;
//   };

  const handleBlog = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    // Upload Image (if exists)
    let imageUrl = '';  
    if (image) {
      const formData = new FormData();
      formData.append('image', image);
    
      const uploadResponse = await fetch('http://localhost:8080/api/blog/upload', {
        method: 'POST',
        body: formData,
        credentials: 'include',
      });
    
      // Debug response
      const responseText = await uploadResponse.text(); // Baca respons sebagai teks
      console.log('Server Response:', responseText); // Log respons untuk melihat apa yang dikirim backend
    
      try {
        const uploadData = JSON.parse(responseText); // Parsing manual JSON
        imageUrl = uploadData.content;
      } catch (error) {
        console.error('Error parsing JSON:', error, responseText);
        alert('Failed to upload image.');
        return;
      }
    }

    // Create Post
    const postResponse = await fetch('http://localhost:8080/api/blog/blog', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        caption,
        content: imageUrl,
      }),
      credentials: 'include', // Sertakan cookie Authorization
    });

    if (postResponse.ok) {
      alert('Post created successfully!');
      router.reload();
    } else {
      alert('Failed to create post.');
    } 
  };

  const handleImageChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      setImage(e.target.files[0]);
    }
  };

  return (
    <form className="post-form" onSubmit={handleBlog}>
      <div className="form-group">
        <label className="form-label">Caption</label>
        <textarea
          className="form-textarea"
          value={caption}
          onChange={(e) => setContent(e.target.value)}
          required
        />
      </div>
      <div className="form-group">
        <label className="form-label">Image</label>
        <input
          className="form-input"
          type="file"
          onChange={handleImageChange}
        />
      </div>
      <button className="form-button" type="submit">Create Post</button>
    </form>
  );
};


export default CreateBlogForm;