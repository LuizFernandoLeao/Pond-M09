import React, { useState } from 'react';

const UserForm = ({ onSubmit, initialData }) => {
    const [name, setName] = useState(initialData ? initialData.name : '');
    const [email, setEmail] = useState(initialData ? initialData.email : '');
    const [password, setPassword] = useState('');
    const [profilePicture, setProfilePicture] = useState(null);

    const handleSubmit = (e) => {
        e.preventDefault();
        const userData = {
            name,
            email,
            password,
            profilePicture,
        };
        onSubmit(userData);
    };

    const handleFileChange = (e) => {
        setProfilePicture(e.target.files[0]);
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>Name:</label>
                <input
                    type="text"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    required
                />
            </div>
            <div>
                <label>Email:</label>
                <input
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
            </div>
            <div>
                <label>Password:</label>
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
            </div>
            <div>
                <label>Profile Picture:</label>
                <input
                    type="file"
                    accept="image/*"
                    onChange={handleFileChange}
                />
            </div>
            <button type="submit">Submit</button>
        </form>
    );
};

export default UserForm;