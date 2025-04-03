import axios from "axios";
import { ApiService } from "./api.service";
import { ErrorService } from "./error.service";
import { ITradingPair } from "@/models/trading.model";

export class TradingService {
    static async getTradingPairs(): Promise<ITradingPair[]> {
        try {
            const token = localStorage.getItem("token");
            if (!token) {
                throw new Error("Not authenticated");
            }

            const res = await axios.get(`${ApiService.baseURL}/pairs`, {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            });
            return res.data;
        } catch (err) {
            ErrorService.mixinMessage("Failed to fetch trading pairs", "error");
            throw err;
        }
    }
} 