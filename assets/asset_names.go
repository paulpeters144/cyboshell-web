package assets

type AssetName string

const (
	// Player assets
	PlayerSeren AssetName = "player_seren"
	PlayerDerek AssetName = "player_derek"

	// Cybo assets
	CyboNuric1 AssetName = "cybo_nuric_1"
	CyboNuric2 AssetName = "cybo_nuric_2"
		CyboNuric3 AssetName = "cybo_nuric_3"
	
		CyboBomake1 AssetName = "cybo_bomake_1"
		CyboBomake2 AssetName = "cybo_bomake_2"
		CyboBomake3 AssetName = "cybo_bomake_3"
	
		CyboGna1 AssetName = "cybo_gna_1"
	
	CyboGna2   AssetName = "cybo_gna_2"
	CyboGna3   AssetName = "cybo_gna_3"

	// Icon assets
	IconStr AssetName = "icon_str"
	IconArm AssetName = "icon_arm"
	IconEvd AssetName = "icon_evd"
	IconAcc AssetName = "icon_acc"
	IconHth AssetName = "icon_hth"

	// Background assets
	BgBattleDaneStreet   AssetName = "bg_battle_dane_street"
	BgBattleFactoryFloor AssetName = "bg_battle_factory_floor"
)

type FontName string

const (
	GameFontTtf       FontName = "game_font_ttf"
	PixelSquare10     FontName = "pixel_square_10"
	PixelSquare10Bold FontName = "pixel_square_10_bold"
	Pixolletta8px     FontName = "pixolletta_8px"
	Pixelmix8px       FontName = "pixelmix_8px"
	PixelmixBold8px   FontName = "pixelmix_bold_8px"
	BigShot12px       FontName = "big_shot_12px"
	Determination12px FontName = "determination_12px"
	Minecraft8px      FontName = "minecraft_8px"
	Monocraft18px     FontName = "monocraft_18px"
)

type ShaderName string

const (
	ShaderFlash   = "shader_flash"
	ShaderSmoke   = "shader_smoke"
	ShaderRadiant = "shader_radiant"
	ShaderSlash     ShaderName = "shader_slash"
	ShaderExplosion ShaderName = "shader_explosion"
	ShaderRefit     ShaderName = "shader_refit"
)
