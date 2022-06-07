@echo off

abigen --bin=_MyERC20_sol_ZecreyRelatedERC20.bin --abi=_MyERC20_sol_ZecreyRelatedERC20.abi --pkg=erc20 --out=ERC20Gen.go && abigen --bin=_Governance_sol_Governance.bin --abi=_Governance_sol_Governance.abi --pkg=governance --out=GovernanceGen.go && abigen --bin=_Governance_sol_Governance.bin --abi=_AssetGovernance_sol_AssetGovernance.abi --pkg=assetGovernance --out=AssetGovernanceGen.go && abigen --bin=_ZecreyVerifier_sol_ZecreyVerifier.bin --abi=_ZecreyVerifier_sol_ZecreyVerifier.abi --pkg=verifier --out=ZecreyVerifierGen.go && abigen --bin=_Zecrey_sol_Zecrey.bin --abi=_Zecrey_sol_Zecrey.abi --pkg=zecrey --out=ZecreyGen.go