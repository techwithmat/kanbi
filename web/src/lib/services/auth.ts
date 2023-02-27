import { RegisterURL, LoginURL } from '#constants/urls';

export type RegisterFormData = {
	email: string;
	username: string;
	password: string;
	password_confirmation: string;
};

export async function RegisterUser(formData: RegisterFormData) {
	// TODO: Add Input validation

	try {
		const response = await fetch(RegisterURL, {
			method: 'POST',
			headers: {
				'Content-type': 'application/json'
			},
			body: JSON.stringify(formData)
		});

		if (!response.ok) {
			throw new Error('Registration failed');
		}

		return await response.json();
	} catch (error) {
		console.error('Registration failed:', error);
		throw error;
	}
}

export type LoginFormData = {
	email: string;
	password: string;
};

export async function LoginUser(formData: LoginFormData) {
	// TODO: Add Input validation
  
	try {
		const response = await fetch(LoginURL, {
			method: 'POST',
			headers: {
				'Content-type': 'application/json'
			},
			credentials: 'include',
			body: JSON.stringify(formData)
		});

		if (!response.ok) {
			throw new Error('Login failed');
		}

		return await response.json();
	} catch (error) {
		console.error('Login failed', error);
		throw error;
	}
}
