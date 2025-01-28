'use client';

import React, { useReducer, useEffect } from 'react';
import styled from 'styled-components';
import Skeleton from '@mui/material/Skeleton';
import { fetchRestaurants } from '../../lib/apis/restaurants';
import { initialState, restaurantsReducer, restaurantsActionTypes } from '../../lib/reducers/restaurants';
import { REQUEST_STATE } from '../../lib/constants';
import RestaurantImage from '../../../public/restaurant-image.jpg';

const HeaderWrapper = styled.div`
  display: flex;
  justify-content: flex-start;
  padding: 8px 32px;
`;

const MainLogoImage = styled.img`
  height: 90px;
`;

const MainCoverImageWrapper = styled.div`
  text-align: center;
`;

const MainCover = styled.img`
  height: 600px;
`;

const RestaurantsContentsList = styled.div`
  display: flex;
  justify-content: space-around;
  margin-bottom: 150px;
`;

const RestaurantsContentWrapper = styled.div`
  width: 450px;
  height: 300px;
  padding: 48px;
`;

const RestaurantsImageNode = styled.img`
  width: 100%;
`;

const MainText = styled.p`
  color: black;
  font-size: 18px;
`;

const SubText = styled.p`
  color: black;
  font-size: 12px;
`;

export default function RestaurantsPage() {
  const [state, dispatch] = useReducer(restaurantsReducer, initialState);

  useEffect(() => {
    dispatch({ type: restaurantsActionTypes.FETCHING });
    fetchRestaurants().then((data) =>
      dispatch({
        type: restaurantsActionTypes.FETCH_SUCCESS,
        payload: {
          restaurants: data.restaurants,
        },
      })
    );
  }, []);

  return (
    <>
      <HeaderWrapper>
        <MainLogoImage src="/logo.png" alt="main logo" />
      </HeaderWrapper>
      <MainCoverImageWrapper>
        <MainCover src="/main-cover-image.png" alt="main cover" />
      </MainCoverImageWrapper>
      <RestaurantsContentsList>
        {state.fetchState === REQUEST_STATE.LOADING ? (
          <>
            <Skeleton variant="rectangular" width={450} height={300} />
            <Skeleton variant="rectangular" width={450} height={300} />
            <Skeleton variant="rectangular" width={450} height={300} />
          </>
        ) : (
          state.restaurantsList.map((item) => (
            <a
              href={`/restaurants/${item.id}/foods`}
              key={item.id}
              style={{ textDecoration: 'none' }}
            >
              <RestaurantsContentWrapper>
                <RestaurantsImageNode src={RestaurantImage.src} alt="Restaurant" />
                <MainText>{item.name}</MainText>
                <SubText>{`配送料：${item.fee}円 ${item.time_required}分`}</SubText>
              </RestaurantsContentWrapper>
            </a>
          ))
        )}
      </RestaurantsContentsList>
    </>
  );
}