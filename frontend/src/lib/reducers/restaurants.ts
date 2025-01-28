import { REQUEST_STATE } from '@/lib/constants';

export interface State {
  fetchState: string;
  restaurantsList: Array<{ id: number; name: string; fee: number; time_required: number }>;
}

export const initialState: State = {
  fetchState: REQUEST_STATE.INITIAL,
  restaurantsList: [],
};

export const restaurantsActionTypes = {
  FETCHING: 'FETCHING',
  FETCH_SUCCESS: 'FETCH_SUCCESS',
} as const;

type Action =
  | { type: 'FETCHING' }
  | { type: 'FETCH_SUCCESS'; payload: { restaurants: Array<any> } };

export const restaurantsReducer = (state: State, action: Action): State => {
  switch (action.type) {
    case restaurantsActionTypes.FETCHING:
      return {
        ...state,
        fetchState: REQUEST_STATE.LOADING,
      };
    case restaurantsActionTypes.FETCH_SUCCESS:
      return {
        fetchState: REQUEST_STATE.OK,
        restaurantsList: action.payload.restaurants,
      };
    default:
      throw new Error('Invalid action type');
  }
};