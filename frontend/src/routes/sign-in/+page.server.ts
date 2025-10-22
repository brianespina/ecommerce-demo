import type { Actions } from './$types';
import { signIn } from 'supertokens-web-js/recipe/emailpassword';

async function signInClicked(email: string, password: string) {
	try {
		const response = await signIn({
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
			response.formFields.forEach((formField) => {
				if (formField.id === 'email') {
					// Email validation failed (for example incorrect email syntax).
					console.log('email err');
				}
			});
		} else if (response.status === 'WRONG_CREDENTIALS_ERROR') {
			console.log('Email password combination is incorrect.');
		} else if (response.status === 'SIGN_IN_NOT_ALLOWED') {
			// the reason string is a user friendly message
			// about what went wrong. It can also contain a support code which users
			// can tell you so you know why their sign in was not allowed.
			console.log(response.reason);
		} else {
			// sign in successful. The session tokens are automatically handled by
			// the frontend SDK.
			console.log('loged in');
		}
		// eslint-disable-next-line
	} catch (err: any) {
		if (err.isSuperTokensGeneralError === true) {
			// this may be a custom error message sent from the API by you.
			console.log(err.message);
		} else {
			console.log('Oops! Something went wrong.');
		}
	}
}

export const actions = {
	default: async () => {
		signInClicked('espinabrian@gmail.com', '123456');
	}
} satisfies Actions;
