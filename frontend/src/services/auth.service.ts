import axios from "axios";
import { IUser } from "../models/users.model";
import { ApiService } from "./api.service";
import { ErrorService } from "./error.service";

export class AuthService {
    static async register(name: string, firstName: string, email: string, password: string) {
        try {
            const res = await axios.post(`${ApiService.baseURL}/register`, {
                name,
                first_name: firstName,
                email,
                password
            });
            ErrorService.mixinMessage("Registered successfully", "success");
            return res.data;
        } catch(err) {
            ErrorService.mixinMessage("Failed to register", "error");
            throw err;
        }
    }
    
    static async login(email: string, password: string) {
        try {
            const res = await axios.post(`${ApiService.baseURL}/login`, {
                email,
                password
            });
            if(res.data.token) {
                localStorage.setItem("token", res.data.token);
                ErrorService.mixinMessage("Logged in successfully", "success");
                return res.data;
            }
        } catch(err) {
            ErrorService.mixinMessage("Failed to login", "error");
            throw err;
        }
    }

    static async getMe(): Promise<IUser | null> {
        try {
            const token = localStorage.getItem("token");
            if (!token) {
                return null;
            }

            const res = await axios.get(`${ApiService.baseURL}/me`, {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            });
            return res.data;
        } catch(err) {
            ErrorService.mixinMessage("Failed to get user info", "error");
            return null;
        }
    }

    static logout() {
        localStorage.removeItem("token");
    }

    static isAuthenticated(): boolean {
        return !!localStorage.getItem("token");
    }

    static async upgrade(): Promise<void> {
        try {
            const token = localStorage.getItem("token");
            if (!token) {
                throw new Error("Not authenticated");
            }

            await axios.post(`${ApiService.baseURL}/upgrade`, {}, {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            });
            ErrorService.mixinMessage("Successfully upgraded your account", "success");
        } catch(err) {
            ErrorService.mixinMessage("Failed to upgrade account", "error");
            throw err;
        }
    }
}