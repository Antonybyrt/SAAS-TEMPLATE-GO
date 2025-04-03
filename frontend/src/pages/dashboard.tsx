import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { AuthService } from '../services/auth.service';
import { TradingService } from '../services/trading.service';
import { ITradingPair } from '@/models/trading.model';
import { IUser } from '../models/users.model';

const Dashboard = () => {
    const router = useRouter();
    const [user, setUser] = useState<IUser | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [tradingPairs, setTradingPairs] = useState<ITradingPair[]>([]);
    const [isLoadingPairs, setIsLoadingPairs] = useState(true);

    useEffect(() => {
        const checkAuth = async () => {
            try {
                const userData = await AuthService.getMe();
                if (!userData) {
                    router.push('/login');
                    return;
                }
                setUser(userData);
            } catch (error) {
                console.error('Failed to fetch user data:', error);
                router.push('/login');
            } finally {
                setIsLoading(false);
            }
        };
        checkAuth();
    }, [router]);

    useEffect(() => {
        const fetchTradingPairs = async () => {
            try {
                const pairs = await TradingService.getTradingPairs();
                setTradingPairs(pairs);
            } catch (error) {
                console.error('Failed to fetch trading pairs:', error);
            } finally {
                setIsLoadingPairs(false);
            }
        };

        if (!isLoading) {
            fetchTradingPairs();
        }
    }, [isLoading]);

    const handleUpgrade = async () => {
        try {
            await AuthService.upgrade();
            const updatedUser = await AuthService.getMe();
            if (updatedUser) {
                setUser(updatedUser);
            }
        } catch (error) {
            console.error('Failed to upgrade:', error);
        }
    };

    const handleLogout = () => {
        AuthService.logout();
        router.push('/login');
    };

    if (isLoading) {
        return (
            <div className="min-h-screen bg-gray-100 flex items-center justify-center">
                <div className="text-xl">Loading...</div>
            </div>
        );
    }

    if (!user) {
        return null;
    }

    const formatPrice = (price: number) => {
        return new Intl.NumberFormat('en-US', {
            style: 'currency',
            currency: 'USD',
            minimumFractionDigits: 2,
            maximumFractionDigits: 2
        }).format(price);
    };

    const formatChange = (change: number) => {
        return `${change >= 0 ? '+' : ''}${change.toFixed(2)}%`;
    };

    return (
        <div className="min-h-screen bg-gray-100 p-8">
            <div className="max-w-7xl mx-auto">
                {/* User Info Section */}
                <div className="bg-white rounded-lg shadow-md p-6 mb-8">
                    <div className="flex justify-between items-start mb-4">
                        <h1 className="text-2xl font-bold text-gray-900">Welcome, {user.first_name}!</h1>
                        <button
                            onClick={handleLogout}
                            className="bg-red-600 text-white px-4 py-2 rounded-lg hover:bg-red-700 transition-colors"
                        >
                            Logout
                        </button>
                    </div>
                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <p className="text-gray-900 font-bold">Name</p>
                            <p className="font-semibold text-gray-600">{user.name}</p>
                        </div>
                        <div>
                            <p className="text-gray-900 font-bold">Email</p>
                            <p className="font-semibold text-gray-600">{user.email}</p>
                        </div>
                    </div>
                </div>

                {/* Trading Pairs Section */}
                <div className="bg-white rounded-lg shadow-md p-6">
                    <h2 className="text-xl font-bold text-gray-900 mb-4">Trading Pairs</h2>
                    
                    {isLoadingPairs ? (
                        <div className="text-center py-8">
                            <div className="text-gray-600">Loading trading pairs...</div>
                        </div>
                    ) : (
                        <>
                            {/* BTC/USDC Pair (Always visible) */}
                            {tradingPairs.find(pair => pair.symbol === "BTC/USDC") && (
                                <div className="mb-4 p-4 bg-gray-50 rounded-lg">
                                    <div className="flex justify-between items-center">
                                        <div>
                                            <h3 className="font-semibold text-gray-600">BTC/USDC</h3>
                                            <p className="text-gray-600">Bitcoin / USD Coin</p>
                                        </div>
                                        <div className="text-right">
                                            {(() => {
                                                const pair = tradingPairs.find(pair => pair.symbol === "BTC/USDC");
                                                return (
                                                    <>
                                                        <p className="text-2xl font-bold text-gray-600">
                                                            {formatPrice(pair?.price ?? 0)}
                                                        </p>
                                                        <p className={(pair?.change_24h ?? 0) >= 0 ? "text-green-500" : "text-red-500"}>
                                                            {formatChange(pair?.change_24h ?? 0)}
                                                        </p>
                                                    </>
                                                );
                                            })()}
                                        </div>
                                    </div>
                                </div>
                            )}

                            {!user.has_upgraded ? (
                                <div className="text-center py-8">
                                    <p className="text-gray-600 mb-4">Upgrade to see more trading pairs</p>
                                    <button
                                        onClick={handleUpgrade}
                                        className="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors"
                                    >
                                        Upgrade Now
                                    </button>
                                </div>
                            ) : (
                                <div className="space-y-4">
                                    {/* ETH/USDC Pair */}
                                    {tradingPairs.find(pair => pair.symbol === "ETH/USDC") && (
                                        <div className="p-4 bg-gray-50 rounded-lg">
                                            <div className="flex justify-between items-center">
                                                <div>
                                                    <h3 className="font-semibold text-gray-600">ETH/USDC</h3>
                                                    <p className="text-gray-600">Ethereum / USD Coin</p>
                                                </div>
                                                <div className="text-right">
                                                    {(() => {
                                                        const pair = tradingPairs.find(pair => pair.symbol === "ETH/USDC");
                                                        return (
                                                            <>
                                                                <p className="text-2xl font-bold text-gray-600">
                                                                    {formatPrice(pair?.price ?? 0)}
                                                                </p>
                                                                <p className={(pair?.change_24h ?? 0) >= 0 ? "text-green-500" : "text-red-500"}>
                                                                    {formatChange(pair?.change_24h ?? 0)}
                                                                </p>
                                                            </>
                                                        );
                                                    })()}
                                                </div>
                                            </div>
                                        </div>
                                    )}

                                    {/* SOL/USDC Pair */}
                                    {tradingPairs.find(pair => pair.symbol === "SOL/USDC") && (
                                        <div className="p-4 bg-gray-50 rounded-lg">
                                            <div className="flex justify-between items-center">
                                                <div>
                                                    <h3 className="font-semibold text-gray-600">SOL/USDC</h3>
                                                    <p className="text-gray-600">Solana / USD Coin</p>
                                                </div>
                                                <div className="text-right">
                                                    {(() => {
                                                        const pair = tradingPairs.find(pair => pair.symbol === "SOL/USDC");
                                                        return (
                                                            <>
                                                                <p className="text-2xl font-bold text-gray-600">
                                                                    {formatPrice(pair?.price ?? 0)}
                                                                </p>
                                                                <p className={(pair?.change_24h ?? 0) >= 0 ? "text-green-500" : "text-red-500"}>
                                                                    {formatChange(pair?.change_24h ?? 0)}
                                                                </p>
                                                            </>
                                                        );
                                                    })()}
                                                </div>
                                            </div>
                                        </div>
                                    )}

                                    {/* XTZ/USDC Pair */}
                                    {tradingPairs.find(pair => pair.symbol === "XTZ/USDC") && (
                                        <div className="p-4 bg-gray-50 rounded-lg">
                                            <div className="flex justify-between items-center">
                                                <div>
                                                    <h3 className="font-semibold text-gray-600">XTZ/USDC</h3>
                                                    <p className="text-gray-600">Tezos / USD Coin</p>
                                                </div>
                                                <div className="text-right">
                                                    {(() => {
                                                        const pair = tradingPairs.find(pair => pair.symbol === "XTZ/USDC");
                                                        return (
                                                            <>
                                                                <p className="text-2xl font-bold text-gray-600">
                                                                    {formatPrice(pair?.price ?? 0)}
                                                                </p>
                                                                <p className={(pair?.change_24h ?? 0) >= 0 ? "text-green-500" : "text-red-500"}>
                                                                    {formatChange(pair?.change_24h ?? 0)}
                                                                </p>
                                                            </>
                                                        );
                                                    })()}
                                                </div>
                                            </div>
                                        </div>
                                    )}
                                </div>
                            )}
                        </>
                    )}
                </div>
            </div>
        </div>
    );
};

export default Dashboard; 