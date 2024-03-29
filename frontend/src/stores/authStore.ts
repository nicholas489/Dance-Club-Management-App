import {defineStore} from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        isAuthenticated: false,
        username: '',
        baid: 0,
        initialized: false,
    }),
    actions: {
        async initAuth(): Promise<void> {
            if (this.initialized) return; // Prevent re-initialization
            this.initialized = true; // Mark as initialized

            const response = await fetch('/auth/session');
            switch (response.status) {
                case 200:
                    const res = await response.json();
                    this.username = res.Username;
                    this.baid = res.Baid;
                    this.isAuthenticated = true;
                    break;
                case 401:
                    this.isAuthenticated = false;
                    break;
                default:
                    console.error('Unexpected response status:', response.status);
            }
        },
        async login(formData: {username: string, password: string}): Promise<[boolean, string]> {
            const response = await fetch('/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData)
            });

            if (response.ok) {
                const body = await response.json();
                this.username = body.Username;
                this.baid = body.Baid;
                this.isAuthenticated = true;
                return [true, ''];
            } else {
                return [false, await response.text()];
            }
        },
        async logout(): Promise<void> {
            const response = await fetch('/auth/logout', {method: 'POST'});
            if (response.ok) {
                this.isAuthenticated = false;
                this.username = '';
                this.baid = 0;
            } else {
                console.error('Failed to logout:', response.status);
            }
        },
    },
});