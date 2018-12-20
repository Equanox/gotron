import * as React from 'react';

import * as styles from 'style.css'

interface RootProps {
}

interface RootState {
}

export class Root extends React.Component<RootProps, RootState> {
  constructor(props: RootProps) {
    super(props);
  }

  render() {
    return (
      <h1 className={styles.topic}>
        Hello Gotron / Typescript-React
      </h1>
    )
  }
}

export default Root;
