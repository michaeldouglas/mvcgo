
import * as React from 'react';
import { hot } from "react-hot-loader/root";


interface Props {
  message: string;
}

const AppComponent: React.FC<Props> = ({ message }) => {
  return <div>{message}</div>;
};

export default hot(AppComponent);
