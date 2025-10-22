import type { Actions } from './$types';

import { signUp } from 'supertokens-web-js/recipe/emailpassword';

async function signUpClicked(email: string, password: string) {
	try {
		const response = await signUp({
			formFields: [
				{
					id: 'email',
					value: email
				},
				{
					id: 'password',
					value: password
				}
			]
		});

		if (response.status === 'FIELD_ERROR') {
			// one of the input formFields failed validation
			response.formFields.forEach((formField) => {
				if (formField.id === 'email') {
					// Email validation failed (for example incorrect email syntax),
					// or the email is not unique.
					console.log('email err');
				} else if (formField.id === 'password') {
					// Password validation failed.
					// Maybe it didn't match the password strength
					console.log('pass err');
				}
			});
		} else if (response.status === 'SIGN_UP_NOT_ALLOWED') {
			// the reason string is a user friendly message
			// about what went wrong. It can also contain a support code which users
			// can tell you so you know why their sign up was not allowed.
		} else {
			// sign up successful. The session tokens are automatically handled by
			// the frontend SDK.
			console.log('signed up');
		}
		// eslint-disable-next-line
	} catch (err: any) {
		if (err.isSuperTokensGeneralError === true) {
			// this may be a custom error message sent from the API by you.
			console.log('error');
		} else {
			console.log('oops');
		}
	}
}
export const actions = {
	default: async () => {
		signUpClicked('espinabrian@gmail.com', '123456');
	}
} satisfies Actions;
