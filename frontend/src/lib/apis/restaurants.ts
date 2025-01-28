export const fetchRestaurants = async (): Promise<any> => {
  const response = await fetch('http://localhost:8080/api/v1/restaurants');
  if (!response.ok) {
    throw new Error('Failed to fetch restaurants');
  }
  return response.json();
};