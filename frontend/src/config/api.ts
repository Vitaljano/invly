const HOST = import.meta.env.VITE_HOST;

export const endpoints = {
	LOGIN: `${HOST}/auth/login`,
	LOGOUT: `${HOST}/logout`,
};
