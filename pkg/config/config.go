package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Contracts struct {
	Exchange         common.Address
	FeeModule        common.Address
	NegRiskExchange  common.Address
	NegRiskFeeModule common.Address
	NegRiskAdapter   common.Address
	Collateral       common.Address
	Conditional      common.Address
}

var (
	AMOY_CONTRACTS = &Contracts{
		Exchange:         common.HexToAddress("0xdFE02Eb6733538f8Ea35D585af8DE5958AD99E40"),
		FeeModule:        common.HexToAddress("0x9A9faEf45C671cc57B3e117c5B3053075416490f"),
		NegRiskExchange:  common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
		NegRiskFeeModule: common.HexToAddress("0x78769d50be1763ed1ca0d5e878d93f05aabff29e"),
		NegRiskAdapter:   common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
		Collateral:       common.HexToAddress("0x9c4e1703476e875070ee25b56a58b008cfb8fa78"),
		Conditional:      common.HexToAddress("0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB"),
	}

	MATIC_CONTRACTS = &Contracts{
		Exchange:         common.HexToAddress("0xbb82b482bc4a7e159759f184c724ae76561a680f"),
		FeeModule:        common.HexToAddress("0x56C79347e95530c01A2FC76E732f9566dA16E113"),
		NegRiskExchange:  common.HexToAddress("0x4f8dffd9c1860b3f41f2f51828549371fc243516"),
		NegRiskFeeModule: common.HexToAddress("0x43732487c768ce27160f81d5bd553371b30b7682"),
		NegRiskAdapter:   common.HexToAddress("0x3ebcfdfc53f249ced0a7ec0d74647b1a6cc32471"),
		Collateral:       common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
		Conditional:      common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
	}

	SEPOLIA_CONTRACTS = &Contracts{
		Exchange:         common.HexToAddress("0xD6bcbD99D98395e8B20D820164CDC00B1628573e"),
		FeeModule:        common.HexToAddress("0x388d08A0772617AE049f3C7E098d4418712AD4fB"),
		NegRiskExchange:  common.HexToAddress("0xD6bcbD99D98395e8B20D820164CDC00B1628573e"),
		NegRiskFeeModule: common.HexToAddress("0x388d08A0772617AE049f3C7E098d4418712AD4fB"),
		NegRiskAdapter:   common.HexToAddress("0xf0668D534c383ea6E4143595084726F588866bcB"),
		Collateral:       common.HexToAddress("0x94e80e896a6daE0828B43B7883f100475b3659E7"),
		Conditional:      common.HexToAddress("0x9c20fed2775cdf526ae7d3d0c598ae0551c5023e"),
	}
	BASE_CONTRACTS = &Contracts{
		Exchange:         common.HexToAddress("0xfd989de1beff564f4fc01aa7f26b9b5b596286de"),
		FeeModule:        common.HexToAddress("0x3ce9a3dd61baa17ea7f9159f04c8ad6fa88cddae"),
		NegRiskExchange:  common.HexToAddress("0xfd989de1beff564f4fc01aa7f26b9b5b596286de"),
		NegRiskFeeModule: common.HexToAddress("0x3ce9a3dd61baa17ea7f9159f04c8ad6fa88cddae"),
		NegRiskAdapter:   common.HexToAddress("0x083d44206c5138a60bc2Fc6dA988862798c87790"),
		Collateral:       common.HexToAddress("0xb0262e7fc826b743fe1404ab68e61710d542ae8f"),
		Conditional:      common.HexToAddress("0x2c0c21C3AE9fE1EF87096466308c4F164272bd2b"),
	}
)

func GetContracts(chainId int64) (*Contracts, error) {
	switch chainId {
	case 137:
		return MATIC_CONTRACTS, nil
	case 80002:
		return AMOY_CONTRACTS, nil
	case 11155111:
		return SEPOLIA_CONTRACTS, nil
	case 8453:
		return BASE_CONTRACTS, nil
	default:
		return nil, fmt.Errorf("invalid chain id")
	}
}
