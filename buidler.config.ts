import { BuidlerConfig, usePlugin } from '@nomiclabs/buidler/config';

usePlugin('@nomiclabs/buidler-waffle');
usePlugin('buidler-typechain');

const config: BuidlerConfig = {
  defaultNetwork: 'buidlerevm',
  solc: {
    version: '0.5.15'
  },
  typechain: {
    outDir: 'typechain',
    target: 'ethers'
  }
};

export default config;
