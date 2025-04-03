import React, { useState } from 'react';
import { useRouter } from 'next/router';
import { FormInput } from '../components/FormInput';
import { AuthService } from '../services/auth.service';
import { ErrorService } from '../services/error.service';

export default function Register() {
    const router = useRouter();
    const [formData, setFormData] = useState({
        name: '',
        firstName: '',
        email: '',
        password: ''
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        });
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            await AuthService.register(
                formData.name,
                formData.firstName,
                formData.email,
                formData.password
            );
            router.push('/login');
        } catch (error) {
            ErrorService.mixinMessage("Registration failed", "error");
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-md w-full space-y-8">
                <div>
                    <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
                        Create your account
                    </h2>
                </div>
                <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
                    <div className="rounded-md shadow-sm -space-y-px">
                        <FormInput
                            label="Name"
                            type="text"
                            name="name"
                            value={formData.name}
                            onChange={handleChange}
                            required
                        />
                        <FormInput
                            label="First Name"
                            type="text"
                            name="firstName"
                            value={formData.firstName}
                            onChange={handleChange}
                            required
                        />
                        <FormInput
                            label="Email address"
                            type="email"
                            name="email"
                            value={formData.email}
                            onChange={handleChange}
                            required
                        />
                        <FormInput
                            label="Password"
                            type="password"
                            name="password"
                            value={formData.password}
                            onChange={handleChange}
                            required
                        />
                    </div>

                    <div>
                        <button
                            type="submit"
                            className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                        >
                            Sign up
                        </button>
                    </div>
                </form>
                <div className="text-center">
                    <p className="text-sm text-gray-600">
                        Already have an account?{' '}
                        <a href="/login" className="font-medium text-indigo-600 hover:text-indigo-500">
                            Sign in
                        </a>
                    </p>
                </div>
            </div>
        </div>
    );
} 