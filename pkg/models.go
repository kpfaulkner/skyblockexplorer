package pkg

type HypixelPlayerResponse struct {
	Success bool `json:"success"`
	Player  struct {
		ID                  string   `json:"_id"`
		UUID                string   `json:"uuid"`
		FirstLogin          int64    `json:"firstLogin"`
		Playername          string   `json:"playername"`
		LastLogin           int64    `json:"lastLogin"`
		Displayname         string   `json:"displayname"`
		KnownAliases        []string `json:"knownAliases"`
		KnownAliasesLower   []string `json:"knownAliasesLower"`
		AchievementsOneTime []string `json:"achievementsOneTime"`
		Stats               struct {
			SkyWars struct {
				Souls                                     int      `json:"souls"`
				ActiveKitSOLO                             string   `json:"activeKit_SOLO"`
				ActiveKitSOLORandom                       bool     `json:"activeKit_SOLO_random"`
				GamesPlayedSkywars                        int      `json:"games_played_skywars"`
				MostKillsGameSolo                         int      `json:"most_kills_game_solo"`
				MostKillsGame                             int      `json:"most_kills_game"`
				WinStreak                                 int      `json:"win_streak"`
				MostKillsGameKitBasicSoloDefault          int      `json:"most_kills_game_kit_basic_solo_default"`
				LastMode                                  string   `json:"lastMode"`
				LongestBowKill                            int      `json:"longest_bow_kill"`
				LongestBowKillKitBasicSoloDefault         int      `json:"longest_bow_kill_kit_basic_solo_default"`
				LongestBowKillSolo                        int      `json:"longest_bow_kill_solo"`
				SurvivedPlayersSolo                       int      `json:"survived_players_solo"`
				MeleeKillsSolo                            int      `json:"melee_kills_solo"`
				BlocksBroken                              int      `json:"blocks_broken"`
				KillsWeeklyA                              int      `json:"kills_weekly_a"`
				Kills                                     int      `json:"kills"`
				RefillChestDestroy                        int      `json:"refill_chest_destroy"`
				KillsKitBasicSoloDefault                  int      `json:"kills_kit_basic_solo_default"`
				SurvivedPlayers                           int      `json:"survived_players"`
				DeathsKitBasicSoloDefault                 int      `json:"deaths_kit_basic_solo_default"`
				MeleeKills                                int      `json:"melee_kills"`
				KillsSolo                                 int      `json:"kills_solo"`
				DeathsSoloInsane                          int      `json:"deaths_solo_insane"`
				Deaths                                    int      `json:"deaths"`
				DeathsSolo                                int      `json:"deaths_solo"`
				BlocksPlaced                              int      `json:"blocks_placed"`
				TimePlayedKitBasicSoloDefault             int      `json:"time_played_kit_basic_solo_default"`
				GamesSolo                                 int      `json:"games_solo"`
				Losses                                    int      `json:"losses"`
				Coins                                     int      `json:"coins"`
				LossesSoloInsane                          int      `json:"losses_solo_insane"`
				MeleeKillsKitBasicSoloDefault             int      `json:"melee_kills_kit_basic_solo_default"`
				TimePlayedSolo                            int      `json:"time_played_solo"`
				LossesSolo                                int      `json:"losses_solo"`
				SoulsGathered                             int      `json:"souls_gathered"`
				TimePlayed                                int      `json:"time_played"`
				SurvivedPlayersKitBasicSoloDefault        int      `json:"survived_players_kit_basic_solo_default"`
				LossesKitBasicSoloDefault                 int      `json:"losses_kit_basic_solo_default"`
				Games                                     int      `json:"games"`
				GamesKitBasicSoloDefault                  int      `json:"games_kit_basic_solo_default"`
				KillsSoloInsane                           int      `json:"kills_solo_insane"`
				KillsMonthlyB                             int      `json:"kills_monthly_b"`
				LongestBowShotKitBasicSoloDefault         int      `json:"longest_bow_shot_kit_basic_solo_default"`
				LongestBowShot                            int      `json:"longest_bow_shot"`
				LongestBowShotSolo                        int      `json:"longest_bow_shot_solo"`
				VoidKillsKitBasicSoloDefault              int      `json:"void_kills_kit_basic_solo_default"`
				KillsWeeklyB                              int      `json:"kills_weekly_b"`
				ArrowsHit                                 int      `json:"arrows_hit"`
				VoidKills                                 int      `json:"void_kills"`
				ChestsOpenedKitBasicSoloDefault           int      `json:"chests_opened_kit_basic_solo_default"`
				ArrowsShotKitBasicSoloDefault             int      `json:"arrows_shot_kit_basic_solo_default"`
				ChestsOpened                              int      `json:"chests_opened"`
				ArrowsHitSolo                             int      `json:"arrows_hit_solo"`
				ArrowsHitKitBasicSoloDefault              int      `json:"arrows_hit_kit_basic_solo_default"`
				ChestsOpenedSolo                          int      `json:"chests_opened_solo"`
				VoidKillsSolo                             int      `json:"void_kills_solo"`
				ArrowsShotSolo                            int      `json:"arrows_shot_solo"`
				ArrowsShot                                int      `json:"arrows_shot"`
				KillsMonthlyA                             int      `json:"kills_monthly_a"`
				MostKillsGameTeam                         int      `json:"most_kills_game_team"`
				MostKillsGameKitAttackingTeamDefault      int      `json:"most_kills_game_kit_attacking_team_default"`
				LongestBowKillKitAttackingTeamDefault     int      `json:"longest_bow_kill_kit_attacking_team_default"`
				LongestBowKillTeam                        int      `json:"longest_bow_kill_team"`
				MeleeKillsTeam                            int      `json:"melee_kills_team"`
				EggThrown                                 int      `json:"egg_thrown"`
				DeathsTeamInsane                          int      `json:"deaths_team_insane"`
				SurvivedPlayersTeam                       int      `json:"survived_players_team"`
				TimePlayedKitAttackingTeamDefault         int      `json:"time_played_kit_attacking_team_default"`
				LossesKitAttackingTeamDefault             int      `json:"losses_kit_attacking_team_default"`
				KillsTeamInsane                           int      `json:"kills_team_insane"`
				DeathsTeam                                int      `json:"deaths_team"`
				KillsTeam                                 int      `json:"kills_team"`
				GamesTeam                                 int      `json:"games_team"`
				DeathsKitAttackingTeamDefault             int      `json:"deaths_kit_attacking_team_default"`
				ChestsOpenedKitAttackingTeamDefault       int      `json:"chests_opened_kit_attacking_team_default"`
				ChestsOpenedTeam                          int      `json:"chests_opened_team"`
				MeleeKillsKitAttackingTeamDefault         int      `json:"melee_kills_kit_attacking_team_default"`
				TimePlayedTeam                            int      `json:"time_played_team"`
				LossesTeam                                int      `json:"losses_team"`
				GamesKitAttackingTeamDefault              int      `json:"games_kit_attacking_team_default"`
				SurvivedPlayersKitAttackingTeamDefault    int      `json:"survived_players_kit_attacking_team_default"`
				KillsKitAttackingTeamDefault              int      `json:"kills_kit_attacking_team_default"`
				LossesTeamInsane                          int      `json:"losses_team_insane"`
				Quits                                     int      `json:"quits"`
				ActiveKitTEAMS                            string   `json:"activeKit_TEAMS"`
				ActiveKitTEAMSRandom                      bool     `json:"activeKit_TEAMS_random"`
				VoidKillsKitAttackingTeamDefault          int      `json:"void_kills_kit_attacking_team_default"`
				VoidKillsTeam                             int      `json:"void_kills_team"`
				AssistsKitAttackingTeamDefault            int      `json:"assists_kit_attacking_team_default"`
				AssistsTeam                               int      `json:"assists_team"`
				Assists                                   int      `json:"assists"`
				LongestBowShotTeam                        int      `json:"longest_bow_shot_team"`
				LongestBowShotKitAttackingTeamDefault     int      `json:"longest_bow_shot_kit_attacking_team_default"`
				ArrowsShotTeam                            int      `json:"arrows_shot_team"`
				BowKillsKitAttackingTeamDefault           int      `json:"bow_kills_kit_attacking_team_default"`
				BowKillsTeam                              int      `json:"bow_kills_team"`
				BowKills                                  int      `json:"bow_kills"`
				ArrowsHitTeam                             int      `json:"arrows_hit_team"`
				ArrowsShotKitAttackingTeamDefault         int      `json:"arrows_shot_kit_attacking_team_default"`
				ArrowsHitKitAttackingTeamDefault          int      `json:"arrows_hit_kit_attacking_team_default"`
				SlimeExplained                            int      `json:"slime_explained"`
				SlimeExplainedLast                        int64    `json:"slime_explained_last"`
				TeamInstantSmelting                       int      `json:"team_instant_smelting"`
				SoulWell                                  int      `json:"soul_well"`
				UsedSoulWell                              bool     `json:"usedSoulWell"`
				ExtraWheels                               int      `json:"extra_wheels"`
				SoloKnowledge                             int      `json:"solo_knowledge"`
				Packages                                  []string `json:"packages"`
				SoulWellRares                             int      `json:"soul_well_rares"`
				TeamEnderMastery                          int      `json:"team_ender_mastery"`
				SoloNourishment                           int      `json:"solo_nourishment"`
				SurvivedPlayersKitBasicSoloScout          int      `json:"survived_players_kit_basic_solo_scout"`
				LossesKitBasicSoloScout                   int      `json:"losses_kit_basic_solo_scout"`
				DeathsKitBasicSoloScout                   int      `json:"deaths_kit_basic_solo_scout"`
				TimePlayedKitBasicSoloScout               int      `json:"time_played_kit_basic_solo_scout"`
				ChestsOpenedKitBasicSoloScout             int      `json:"chests_opened_kit_basic_solo_scout"`
				LongestBowShotKitBasicSoloScout           int      `json:"longest_bow_shot_kit_basic_solo_scout"`
				MostKillsGameKitBasicSoloScout            int      `json:"most_kills_game_kit_basic_solo_scout"`
				LongestBowKillKitBasicSoloScout           int      `json:"longest_bow_kill_kit_basic_solo_scout"`
				KillsKitBasicSoloScout                    int      `json:"kills_kit_basic_solo_scout"`
				MeleeKillsKitBasicSoloScout               int      `json:"melee_kills_kit_basic_solo_scout"`
				ArrowsShotKitBasicSoloScout               int      `json:"arrows_shot_kit_basic_solo_scout"`
				ArrowsHitKitBasicSoloScout                int      `json:"arrows_hit_kit_basic_solo_scout"`
				GamesKitBasicSoloScout                    int      `json:"games_kit_basic_solo_scout"`
				MostKillsGameKitBasicSoloGrenade          int      `json:"most_kills_game_kit_basic_solo_grenade"`
				LongestBowKillKitBasicSoloGrenade         int      `json:"longest_bow_kill_kit_basic_solo_grenade"`
				MeleeKillsKitBasicSoloGrenade             int      `json:"melee_kills_kit_basic_solo_grenade"`
				TimePlayedKitBasicSoloGrenade             int      `json:"time_played_kit_basic_solo_grenade"`
				VoidKillsKitBasicSoloGrenade              int      `json:"void_kills_kit_basic_solo_grenade"`
				KillsKitBasicSoloGrenade                  int      `json:"kills_kit_basic_solo_grenade"`
				SurvivedPlayersKitBasicSoloGrenade        int      `json:"survived_players_kit_basic_solo_grenade"`
				GamesKitBasicSoloGrenade                  int      `json:"games_kit_basic_solo_grenade"`
				DeathsKitBasicSoloGrenade                 int      `json:"deaths_kit_basic_solo_grenade"`
				LossesKitBasicSoloGrenade                 int      `json:"losses_kit_basic_solo_grenade"`
				ChestsOpenedKitBasicSoloGrenade           int      `json:"chests_opened_kit_basic_solo_grenade"`
				KillsSoloNormal                           int      `json:"kills_solo_normal"`
				DeathsSoloNormal                          int      `json:"deaths_solo_normal"`
				LossesSoloNormal                          int      `json:"losses_solo_normal"`
				SkywarsExperience                         int      `json:"skywars_experience"`
				SoloInstantSmelting                       int      `json:"solo_instant_smelting"`
				SoloEnderMastery                          int      `json:"solo_ender_mastery"`
				TeamNourishment                           int      `json:"team_nourishment"`
				TeamKnowledge                             int      `json:"team_knowledge"`
				FastestWinKitBasicSoloScout               int      `json:"fastest_win_kit_basic_solo_scout"`
				FastestWin                                int      `json:"fastest_win"`
				FastestWinSolo                            int      `json:"fastest_win_solo"`
				KillstreakSolo                            int      `json:"killstreak_solo"`
				Killstreak                                int      `json:"killstreak"`
				WinsSoloNormal                            int      `json:"wins_solo_normal"`
				Wins                                      int      `json:"wins"`
				WinstreakSolo                             int      `json:"winstreak_solo"`
				Winstreak                                 int      `json:"winstreak"`
				WinsKitBasicSoloScout                     int      `json:"wins_kit_basic_solo_scout"`
				WinstreakKitBasicSoloScout                int      `json:"winstreak_kit_basic_solo_scout"`
				KillstreakKitBasicSoloScout               int      `json:"killstreak_kit_basic_solo_scout"`
				WinsSolo                                  int      `json:"wins_solo"`
				VoidKillsKitBasicSoloScout                int      `json:"void_kills_kit_basic_solo_scout"`
				SoloBridger                               int      `json:"solo_bridger"`
				SoloResistanceBoost                       int      `json:"solo_resistance_boost"`
				TeamMarksmanship                          int      `json:"team_marksmanship"`
				SoulWellLegendaries                       int      `json:"soul_well_legendaries"`
				LossesKitMiningTeamDefault                int      `json:"losses_kit_mining_team_default"`
				SurvivedPlayersKitMiningTeamDefault       int      `json:"survived_players_kit_mining_team_default"`
				TimePlayedKitMiningTeamDefault            int      `json:"time_played_kit_mining_team_default"`
				DeathsKitMiningTeamDefault                int      `json:"deaths_kit_mining_team_default"`
				SoloLuckyCharm                            int      `json:"solo_lucky_charm"`
				LongestBowShotKitSupportingTeamRookie     int      `json:"longest_bow_shot_kit_supporting_team_rookie"`
				LossesKitSupportingTeamRookie             int      `json:"losses_kit_supporting_team_rookie"`
				ArrowsHitKitSupportingTeamRookie          int      `json:"arrows_hit_kit_supporting_team_rookie"`
				ChestsOpenedKitSupportingTeamRookie       int      `json:"chests_opened_kit_supporting_team_rookie"`
				TimePlayedKitSupportingTeamRookie         int      `json:"time_played_kit_supporting_team_rookie"`
				DeathsKitSupportingTeamRookie             int      `json:"deaths_kit_supporting_team_rookie"`
				SurvivedPlayersKitSupportingTeamRookie    int      `json:"survived_players_kit_supporting_team_rookie"`
				ArrowsShotKitSupportingTeamRookie         int      `json:"arrows_shot_kit_supporting_team_rookie"`
				LongestBowKillKitSupportingTeamRookie     int      `json:"longest_bow_kill_kit_supporting_team_rookie"`
				MostKillsGameKitSupportingTeamRookie      int      `json:"most_kills_game_kit_supporting_team_rookie"`
				GamesKitSupportingTeamRookie              int      `json:"games_kit_supporting_team_rookie"`
				MeleeKillsKitSupportingTeamRookie         int      `json:"melee_kills_kit_supporting_team_rookie"`
				KillsKitSupportingTeamRookie              int      `json:"kills_kit_supporting_team_rookie"`
				MostKillsGameKitBasicSoloRookie           int      `json:"most_kills_game_kit_basic_solo_rookie"`
				SurvivedPlayersKitBasicSoloRookie         int      `json:"survived_players_kit_basic_solo_rookie"`
				KillsKitBasicSoloRookie                   int      `json:"kills_kit_basic_solo_rookie"`
				VoidKillsKitBasicSoloRookie               int      `json:"void_kills_kit_basic_solo_rookie"`
				TimePlayedKitBasicSoloRookie              int      `json:"time_played_kit_basic_solo_rookie"`
				LossesKitBasicSoloRookie                  int      `json:"losses_kit_basic_solo_rookie"`
				ChestsOpenedKitBasicSoloRookie            int      `json:"chests_opened_kit_basic_solo_rookie"`
				DeathsKitBasicSoloRookie                  int      `json:"deaths_kit_basic_solo_rookie"`
				ItemsEnchanted                            int      `json:"items_enchanted"`
				EnderpearlsThrown                         int      `json:"enderpearls_thrown"`
				AssistsKitBasicSoloScout                  int      `json:"assists_kit_basic_solo_scout"`
				AssistsSolo                               int      `json:"assists_solo"`
				LossesKitAttackingTeamMagician            int      `json:"losses_kit_attacking_team_magician"`
				DeathsKitAttackingTeamMagician            int      `json:"deaths_kit_attacking_team_magician"`
				TimePlayedKitAttackingTeamMagician        int      `json:"time_played_kit_attacking_team_magician"`
				ChestsOpenedKitAttackingTeamMagician      int      `json:"chests_opened_kit_attacking_team_magician"`
				SurvivedPlayersKitAttackingTeamMagician   int      `json:"survived_players_kit_attacking_team_magician"`
				SurvivedPlayersKitAttackingTeamScout      int      `json:"survived_players_kit_attacking_team_scout"`
				ChestsOpenedKitAttackingTeamScout         int      `json:"chests_opened_kit_attacking_team_scout"`
				LossesKitAttackingTeamScout               int      `json:"losses_kit_attacking_team_scout"`
				GamesKitAttackingTeamScout                int      `json:"games_kit_attacking_team_scout"`
				DeathsKitAttackingTeamScout               int      `json:"deaths_kit_attacking_team_scout"`
				TimePlayedKitAttackingTeamScout           int      `json:"time_played_kit_attacking_team_scout"`
				MostKillsGameKitAttackingTeamScout        int      `json:"most_kills_game_kit_attacking_team_scout"`
				VoidKillsKitAttackingTeamScout            int      `json:"void_kills_kit_attacking_team_scout"`
				KillsKitAttackingTeamScout                int      `json:"kills_kit_attacking_team_scout"`
				LongestBowKillKitAttackingTeamScout       int      `json:"longest_bow_kill_kit_attacking_team_scout"`
				MeleeKillsKitAttackingTeamScout           int      `json:"melee_kills_kit_attacking_team_scout"`
				VoidKillsKitSupportingTeamRookie          int      `json:"void_kills_kit_supporting_team_rookie"`
				TimePlayedKitBasicSoloPharaoh             int      `json:"time_played_kit_basic_solo_pharaoh"`
				ChestsOpenedKitBasicSoloPharaoh           int      `json:"chests_opened_kit_basic_solo_pharaoh"`
				SurvivedPlayersKitBasicSoloPharaoh        int      `json:"survived_players_kit_basic_solo_pharaoh"`
				LossesKitBasicSoloPharaoh                 int      `json:"losses_kit_basic_solo_pharaoh"`
				DeathsKitBasicSoloPharaoh                 int      `json:"deaths_kit_basic_solo_pharaoh"`
				MegaInstantSmelting                       int      `json:"mega_instant_smelting"`
				LongestBowShotKitAttackingTeamScout       int      `json:"longest_bow_shot_kit_attacking_team_scout"`
				ArrowsShotKitAttackingTeamScout           int      `json:"arrows_shot_kit_attacking_team_scout"`
				ArrowsHitKitAttackingTeamScout            int      `json:"arrows_hit_kit_attacking_team_scout"`
				AssistsKitAttackingTeamScout              int      `json:"assists_kit_attacking_team_scout"`
				FastestWinKitAttackingTeamScout           int      `json:"fastest_win_kit_attacking_team_scout"`
				WinsKitAttackingTeamScout                 int      `json:"wins_kit_attacking_team_scout"`
				WinsSoloInsane                            int      `json:"wins_solo_insane"`
				WinstreakKitAttackingTeamScout            int      `json:"winstreak_kit_attacking_team_scout"`
				KillstreakKitAttackingTeamScout           int      `json:"killstreak_kit_attacking_team_scout"`
				MegaJuggernaut                            int      `json:"mega_juggernaut"`
				LossesKitAttackingTeamEnergix             int      `json:"losses_kit_attacking_team_energix"`
				ChestsOpenedKitAttackingTeamEnergix       int      `json:"chests_opened_kit_attacking_team_energix"`
				TimePlayedKitAttackingTeamEnergix         int      `json:"time_played_kit_attacking_team_energix"`
				ArrowsShotKitAttackingTeamEnergix         int      `json:"arrows_shot_kit_attacking_team_energix"`
				DeathsKitAttackingTeamEnergix             int      `json:"deaths_kit_attacking_team_energix"`
				GamesKitAttackingTeamEnergix              int      `json:"games_kit_attacking_team_energix"`
				SurvivedPlayersKitAttackingTeamEnergix    int      `json:"survived_players_kit_attacking_team_energix"`
				WinstreakTeam                             int      `json:"winstreak_team"`
				WinsTeam                                  int      `json:"wins_team"`
				WinsTeamInsane                            int      `json:"wins_team_insane"`
				BowKillsKitAttackingTeamScout             int      `json:"bow_kills_kit_attacking_team_scout"`
				KillstreakTeam                            int      `json:"killstreak_team"`
				MegaLuckyCharm                            int      `json:"mega_lucky_charm"`
				TeamSavior                                int      `json:"team_savior"`
				MegaNourishment                           int      `json:"mega_nourishment"`
				TeamLuckyCharm                            int      `json:"team_lucky_charm"`
				TeamBlazingArrows                         int      `json:"team_blazing_arrows"`
				LossesKitAdvancedSoloSlime                int      `json:"losses_kit_advanced_solo_slime"`
				ChestsOpenedKitAdvancedSoloSlime          int      `json:"chests_opened_kit_advanced_solo_slime"`
				DeathsKitAdvancedSoloSlime                int      `json:"deaths_kit_advanced_solo_slime"`
				TimePlayedKitAdvancedSoloSlime            int      `json:"time_played_kit_advanced_solo_slime"`
				SurvivedPlayersKitAdvancedSoloSlime       int      `json:"survived_players_kit_advanced_solo_slime"`
				SkywarsChests                             int      `json:"skywars_chests"`
				LongestBowKillKitAdvancedSoloGuardian     int      `json:"longest_bow_kill_kit_advanced_solo_guardian"`
				MostKillsGameKitAdvancedSoloGuardian      int      `json:"most_kills_game_kit_advanced_solo_guardian"`
				LongestBowShotKitAdvancedSoloGuardian     int      `json:"longest_bow_shot_kit_advanced_solo_guardian"`
				FastestWinKitAdvancedSoloGuardian         int      `json:"fastest_win_kit_advanced_solo_guardian"`
				GamesKitAdvancedSoloGuardian              int      `json:"games_kit_advanced_solo_guardian"`
				ArrowsHitKitAdvancedSoloGuardian          int      `json:"arrows_hit_kit_advanced_solo_guardian"`
				ChestsOpenedKitAdvancedSoloGuardian       int      `json:"chests_opened_kit_advanced_solo_guardian"`
				MeleeKillsKitAdvancedSoloGuardian         int      `json:"melee_kills_kit_advanced_solo_guardian"`
				KillstreakKitAdvancedSoloGuardian         int      `json:"killstreak_kit_advanced_solo_guardian"`
				KillsKitAdvancedSoloGuardian              int      `json:"kills_kit_advanced_solo_guardian"`
				WinsKitAdvancedSoloGuardian               int      `json:"wins_kit_advanced_solo_guardian"`
				WinstreakKitAdvancedSoloGuardian          int      `json:"winstreak_kit_advanced_solo_guardian"`
				TimePlayedKitAdvancedSoloGuardian         int      `json:"time_played_kit_advanced_solo_guardian"`
				SurvivedPlayersKitAdvancedSoloGuardian    int      `json:"survived_players_kit_advanced_solo_guardian"`
				ArrowsShotKitAdvancedSoloGuardian         int      `json:"arrows_shot_kit_advanced_solo_guardian"`
				MostKillsGameKitDefendingTeamBatguy       int      `json:"most_kills_game_kit_defending_team_batguy"`
				LongestBowShotKitDefendingTeamBatguy      int      `json:"longest_bow_shot_kit_defending_team_batguy"`
				ArrowsShotKitDefendingTeamBatguy          int      `json:"arrows_shot_kit_defending_team_batguy"`
				ChestsOpenedKitDefendingTeamBatguy        int      `json:"chests_opened_kit_defending_team_batguy"`
				VoidKillsKitDefendingTeamBatguy           int      `json:"void_kills_kit_defending_team_batguy"`
				KillsKitDefendingTeamBatguy               int      `json:"kills_kit_defending_team_batguy"`
				DeathsKitDefendingTeamBatguy              int      `json:"deaths_kit_defending_team_batguy"`
				ArrowsHitKitDefendingTeamBatguy           int      `json:"arrows_hit_kit_defending_team_batguy"`
				TimePlayedKitDefendingTeamBatguy          int      `json:"time_played_kit_defending_team_batguy"`
				SurvivedPlayersKitDefendingTeamBatguy     int      `json:"survived_players_kit_defending_team_batguy"`
				LossesKitDefendingTeamBatguy              int      `json:"losses_kit_defending_team_batguy"`
				LongestBowShotKitAttackingTeamEnergix     int      `json:"longest_bow_shot_kit_attacking_team_energix"`
				AssistsKitAttackingTeamEnergix            int      `json:"assists_kit_attacking_team_energix"`
				ArrowsHitKitAttackingTeamEnergix          int      `json:"arrows_hit_kit_attacking_team_energix"`
				SkyWarsOpenedCommons                      int      `json:"SkyWars_openedCommons"`
				SkywarsChestHistory                       []string `json:"skywars_chest_history"`
				SkyWarsOpenedChests                       int      `json:"SkyWars_openedChests"`
				SkyWarsOpenedRares                        int      `json:"SkyWars_openedRares"`
				ActiveKilleffect                          string   `json:"active_killeffect"`
				ActiveVictorydance                        string   `json:"active_victorydance"`
				MostKillsGameKitAdvancedSoloHunter        int      `json:"most_kills_game_kit_advanced_solo_hunter"`
				LongestBowShotKitAdvancedSoloHunter       int      `json:"longest_bow_shot_kit_advanced_solo_hunter"`
				LongestBowKillKitAdvancedSoloHunter       int      `json:"longest_bow_kill_kit_advanced_solo_hunter"`
				AssistsKitAdvancedSoloHunter              int      `json:"assists_kit_advanced_solo_hunter"`
				VoidKillsKitAdvancedSoloHunter            int      `json:"void_kills_kit_advanced_solo_hunter"`
				ArrowsHitKitAdvancedSoloHunter            int      `json:"arrows_hit_kit_advanced_solo_hunter"`
				DeathsKitAdvancedSoloHunter               int      `json:"deaths_kit_advanced_solo_hunter"`
				KillsKitAdvancedSoloHunter                int      `json:"kills_kit_advanced_solo_hunter"`
				TimePlayedKitAdvancedSoloHunter           int      `json:"time_played_kit_advanced_solo_hunter"`
				ArrowsShotKitAdvancedSoloHunter           int      `json:"arrows_shot_kit_advanced_solo_hunter"`
				GamesKitAdvancedSoloHunter                int      `json:"games_kit_advanced_solo_hunter"`
				SurvivedPlayersKitAdvancedSoloHunter      int      `json:"survived_players_kit_advanced_solo_hunter"`
				ChestsOpenedKitAdvancedSoloHunter         int      `json:"chests_opened_kit_advanced_solo_hunter"`
				LossesKitAdvancedSoloHunter               int      `json:"losses_kit_advanced_solo_hunter"`
				MeleeKillsKitAdvancedSoloHunter           int      `json:"melee_kills_kit_advanced_solo_hunter"`
				BowKillsSolo                              int      `json:"bow_kills_solo"`
				BowKillsKitAdvancedSoloHunter             int      `json:"bow_kills_kit_advanced_solo_hunter"`
				TntMadnessExplainedLast                   int64    `json:"tnt_madness_explained_last"`
				TntMadnessExplained                       int      `json:"tnt_madness_explained"`
				WinStreakLab                              int      `json:"win_streak_lab"`
				ChestsOpenedLabSolo                       int      `json:"chests_opened_lab_solo"`
				DeathsLab                                 int      `json:"deaths_lab"`
				DeathsLabSolo                             int      `json:"deaths_lab_solo"`
				ChestsOpenedLabKitAttackingTeamScout      int      `json:"chests_opened_lab_kit_attacking_team_scout"`
				TimePlayedLabSolo                         int      `json:"time_played_lab_solo"`
				TimePlayedLab                             int      `json:"time_played_lab"`
				DeathsLabKitAttackingTeamScout            int      `json:"deaths_lab_kit_attacking_team_scout"`
				CoinsGainedLab                            int      `json:"coins_gained_lab"`
				LossesLabSolo                             int      `json:"losses_lab_solo"`
				LossesLabKitAttackingTeamScout            int      `json:"losses_lab_kit_attacking_team_scout"`
				ChestsOpenedLab                           int      `json:"chests_opened_lab"`
				TimePlayedLabKitAttackingTeamScout        int      `json:"time_played_lab_kit_attacking_team_scout"`
				LossesLab                                 int      `json:"losses_lab"`
				QuitsLab                                  int      `json:"quits_lab"`
				FastestWinTeam                            int      `json:"fastest_win_team"`
				LongestBowKillKitAttackingTeamEnergix     int      `json:"longest_bow_kill_kit_attacking_team_energix"`
				FastestWinKitAttackingTeamEnergix         int      `json:"fastest_win_kit_attacking_team_energix"`
				MostKillsGameKitAttackingTeamEnergix      int      `json:"most_kills_game_kit_attacking_team_energix"`
				WinstreakKitAttackingTeamEnergix          int      `json:"winstreak_kit_attacking_team_energix"`
				VoidKillsKitAttackingTeamEnergix          int      `json:"void_kills_kit_attacking_team_energix"`
				KillstreakKitAttackingTeamEnergix         int      `json:"killstreak_kit_attacking_team_energix"`
				MeleeKillsKitAttackingTeamEnergix         int      `json:"melee_kills_kit_attacking_team_energix"`
				WinsKitAttackingTeamEnergix               int      `json:"wins_kit_attacking_team_energix"`
				KillsKitAttackingTeamEnergix              int      `json:"kills_kit_attacking_team_energix"`
				FastestWinKitSupportingTeamRookie         int      `json:"fastest_win_kit_supporting_team_rookie"`
				WinstreakKitSupportingTeamRookie          int      `json:"winstreak_kit_supporting_team_rookie"`
				WinsKitSupportingTeamRookie               int      `json:"wins_kit_supporting_team_rookie"`
				MostKillsGameKitBasicSoloPharaoh          int      `json:"most_kills_game_kit_basic_solo_pharaoh"`
				KillsKitBasicSoloPharaoh                  int      `json:"kills_kit_basic_solo_pharaoh"`
				VoidKillsKitBasicSoloPharaoh              int      `json:"void_kills_kit_basic_solo_pharaoh"`
				LongestBowKillKitBasicSoloPharaoh         int      `json:"longest_bow_kill_kit_basic_solo_pharaoh"`
				GamesKitBasicSoloPharaoh                  int      `json:"games_kit_basic_solo_pharaoh"`
				MeleeKillsKitBasicSoloPharaoh             int      `json:"melee_kills_kit_basic_solo_pharaoh"`
				TeamBridger                               int      `json:"team_bridger"`
				MegaBridger                               int      `json:"mega_bridger"`
				MegaEnderMastery                          int      `json:"mega_ender_mastery"`
				SoloBlackMagic                            int      `json:"solo_black_magic"`
				MegaBlazingArrows                         int      `json:"mega_blazing_arrows"`
				LevelFormatted                            string   `json:"levelFormatted"`
				ActiveKitRANKEDRandom                     bool     `json:"activeKit_RANKED_random"`
				ActiveKitRANKED                           string   `json:"activeKit_RANKED"`
				FastestWinKitRankedRankedScout            int      `json:"fastest_win_kit_ranked_ranked_scout"`
				MostKillsGameRanked                       int      `json:"most_kills_game_ranked"`
				MostKillsGameKitRankedRankedScout         int      `json:"most_kills_game_kit_ranked_ranked_scout"`
				LongestBowKillKitRankedRankedScout        int      `json:"longest_bow_kill_kit_ranked_ranked_scout"`
				LongestBowKillRanked                      int      `json:"longest_bow_kill_ranked"`
				FastestWinRanked                          int      `json:"fastest_win_ranked"`
				WinstreakRanked                           int      `json:"winstreak_ranked"`
				ChestsOpenedRanked                        int      `json:"chests_opened_ranked"`
				GamesKitRankedRankedScout                 int      `json:"games_kit_ranked_ranked_scout"`
				TimePlayedRanked                          int      `json:"time_played_ranked"`
				SurvivedPlayersKitRankedRankedScout       int      `json:"survived_players_kit_ranked_ranked_scout"`
				TimePlayedKitRankedRankedScout            int      `json:"time_played_kit_ranked_ranked_scout"`
				GamesRanked                               int      `json:"games_ranked"`
				WinsRankedNormal                          int      `json:"wins_ranked_normal"`
				ChestsOpenedKitRankedRankedScout          int      `json:"chests_opened_kit_ranked_ranked_scout"`
				KillsRankedNormal                         int      `json:"kills_ranked_normal"`
				KillstreakRanked                          int      `json:"killstreak_ranked"`
				MeleeKillsRanked                          int      `json:"melee_kills_ranked"`
				MeleeKillsKitRankedRankedScout            int      `json:"melee_kills_kit_ranked_ranked_scout"`
				SurvivedPlayersRanked                     int      `json:"survived_players_ranked"`
				KillsRanked                               int      `json:"kills_ranked"`
				VoidKillsKitRankedRankedScout             int      `json:"void_kills_kit_ranked_ranked_scout"`
				WinstreakKitRankedRankedScout             int      `json:"winstreak_kit_ranked_ranked_scout"`
				VoidKillsRanked                           int      `json:"void_kills_ranked"`
				KillstreakKitRankedRankedScout            int      `json:"killstreak_kit_ranked_ranked_scout"`
				KillsKitRankedRankedScout                 int      `json:"kills_kit_ranked_ranked_scout"`
				WinsRanked                                int      `json:"wins_ranked"`
				WinsKitRankedRankedScout                  int      `json:"wins_kit_ranked_ranked_scout"`
				MobKills                                  int      `json:"mob_kills"`
				MobKillsKitAttackingTeamScout             int      `json:"mob_kills_kit_attacking_team_scout"`
				MobKillsSolo                              int      `json:"mob_kills_solo"`
				DeathsRanked                              int      `json:"deaths_ranked"`
				LossesRanked                              int      `json:"losses_ranked"`
				DeathsRankedNormal                        int      `json:"deaths_ranked_normal"`
				DeathsKitRankedRankedScout                int      `json:"deaths_kit_ranked_ranked_scout"`
				LossesKitRankedRankedScout                int      `json:"losses_kit_ranked_ranked_scout"`
				LossesRankedNormal                        int      `json:"losses_ranked_normal"`
				AssistsRanked                             int      `json:"assists_ranked"`
				AssistsKitRankedRankedScout               int      `json:"assists_kit_ranked_ranked_scout"`
				LuckyExplained                            int      `json:"lucky_explained"`
				LuckyExplainedLast                        int64    `json:"lucky_explained_last"`
				SurvivedPlayersLabSolo                    int      `json:"survived_players_lab_solo"`
				WinsLab                                   int      `json:"wins_lab"`
				KillstreakLab                             int      `json:"killstreak_lab"`
				VoidKillsLab                              int      `json:"void_kills_lab"`
				VoidKillsLabKitAttackingTeamScout         int      `json:"void_kills_lab_kit_attacking_team_scout"`
				BlocksBrokenLab                           int      `json:"blocks_broken_lab"`
				GamesLab                                  int      `json:"games_lab"`
				SurvivedPlayersLab                        int      `json:"survived_players_lab"`
				KillsLabKitAttackingTeamScout             int      `json:"kills_lab_kit_attacking_team_scout"`
				ArrowsShotLabSolo                         int      `json:"arrows_shot_lab_solo"`
				KillsLab                                  int      `json:"kills_lab"`
				SurvivedPlayersLabKitAttackingTeamScout   int      `json:"survived_players_lab_kit_attacking_team_scout"`
				WinstreakLab                              int      `json:"winstreak_lab"`
				GamesLabSolo                              int      `json:"games_lab_solo"`
				ArrowsHitLabKitAttackingTeamScout         int      `json:"arrows_hit_lab_kit_attacking_team_scout"`
				ArrowsShotLabKitAttackingTeamScout        int      `json:"arrows_shot_lab_kit_attacking_team_scout"`
				MeleeKillsLabSolo                         int      `json:"melee_kills_lab_solo"`
				BlocksPlacedLab                           int      `json:"blocks_placed_lab"`
				GamesLabKitAttackingTeamScout             int      `json:"games_lab_kit_attacking_team_scout"`
				WinstreakLabKitAttackingTeamScout         int      `json:"winstreak_lab_kit_attacking_team_scout"`
				KillsLabSolo                              int      `json:"kills_lab_solo"`
				WinstreakLabSolo                          int      `json:"winstreak_lab_solo"`
				LabWinLuckyBlocksLabSolo                  int      `json:"lab_win_lucky_blocks_lab_solo"`
				LabWinLuckyBlocksLab                      int      `json:"lab_win_lucky_blocks_lab"`
				WinsLabKitAttackingTeamScout              int      `json:"wins_lab_kit_attacking_team_scout"`
				KillstreakLabKitAttackingTeamScout        int      `json:"killstreak_lab_kit_attacking_team_scout"`
				MeleeKillsLabKitAttackingTeamScout        int      `json:"melee_kills_lab_kit_attacking_team_scout"`
				KillstreakLabSolo                         int      `json:"killstreak_lab_solo"`
				LabWinLuckyBlocksLabKitAttackingTeamScout int      `json:"lab_win_lucky_blocks_lab_kit_attacking_team_scout"`
				WinsLabSolo                               int      `json:"wins_lab_solo"`
				ArrowsHitLab                              int      `json:"arrows_hit_lab"`
				VoidKillsLabSolo                          int      `json:"void_kills_lab_solo"`
				MeleeKillsLab                             int      `json:"melee_kills_lab"`
				ArrowsShotLab                             int      `json:"arrows_shot_lab"`
				ArrowsHitLabSolo                          int      `json:"arrows_hit_lab_solo"`
				EnderpearlsThrownLab                      int      `json:"enderpearls_thrown_lab"`
				LongestBowShotRanked                      int      `json:"longest_bow_shot_ranked"`
				LongestBowShotKitRankedRankedScout        int      `json:"longest_bow_shot_kit_ranked_ranked_scout"`
				ArrowsHitRanked                           int      `json:"arrows_hit_ranked"`
				ArrowsShotRanked                          int      `json:"arrows_shot_ranked"`
				ArrowsShotKitRankedRankedScout            int      `json:"arrows_shot_kit_ranked_ranked_scout"`
				ArrowsHitKitRankedRankedScout             int      `json:"arrows_hit_kit_ranked_ranked_scout"`
				DeathsKitRankedRankedArmorer              int      `json:"deaths_kit_ranked_ranked_armorer"`
				LossesKitRankedRankedArmorer              int      `json:"losses_kit_ranked_ranked_armorer"`
				TimePlayedKitRankedRankedArmorer          int      `json:"time_played_kit_ranked_ranked_armorer"`
				SurvivedPlayersKitRankedRankedArmorer     int      `json:"survived_players_kit_ranked_ranked_armorer"`
				Heads                                     int      `json:"heads"`
				HeadsKitAttackingTeamScout                int      `json:"heads_kit_attacking_team_scout"`
				HeadsMeh                                  int      `json:"heads_meh"`
				HeadsMehKitAttackingTeamScout             int      `json:"heads_meh_kit_attacking_team_scout"`
				HeadsMehTeam                              int      `json:"heads_meh_team"`
				HeadsTeam                                 int      `json:"heads_team"`
				HeadCollection                            struct {
					Recent []struct {
						UUID      string `json:"uuid"`
						Timestamp int64  `json:"timestamp"`
						Mode      string `json:"mode"`
						Sacrifice string `json:"sacrifice"`
					} `json:"recent"`
					Prestigious []struct {
						UUID      string `json:"uuid"`
						Timestamp int64  `json:"timestamp"`
						Mode      string `json:"mode"`
						Sacrifice string `json:"sacrifice"`
					} `json:"prestigious"`
				} `json:"head_collection"`
				DeathsKitAttackingTeamSloth               int     `json:"deaths_kit_attacking_team_sloth"`
				LossesKitAttackingTeamSloth               int     `json:"losses_kit_attacking_team_sloth"`
				SurvivedPlayersKitAttackingTeamSloth      int     `json:"survived_players_kit_attacking_team_sloth"`
				TimePlayedKitAttackingTeamSloth           int     `json:"time_played_kit_attacking_team_sloth"`
				ChestsOpenedKitAttackingTeamSloth         int     `json:"chests_opened_kit_attacking_team_sloth"`
				SoloMiningExpertise                       int     `json:"solo_mining_expertise"`
				TeamResistanceBoost                       int     `json:"team_resistance_boost"`
				AssistsKitSupportingTeamEcologist         int     `json:"assists_kit_supporting_team_ecologist"`
				ChestsOpenedKitSupportingTeamEcologist    int     `json:"chests_opened_kit_supporting_team_ecologist"`
				DeathsKitSupportingTeamEcologist          int     `json:"deaths_kit_supporting_team_ecologist"`
				FastestWinKitSupportingTeamEcologist      int     `json:"fastest_win_kit_supporting_team_ecologist"`
				GamesKitSupportingTeamEcologist           int     `json:"games_kit_supporting_team_ecologist"`
				KillsKitSupportingTeamEcologist           int     `json:"kills_kit_supporting_team_ecologist"`
				KillstreakKitSupportingTeamEcologist      int     `json:"killstreak_kit_supporting_team_ecologist"`
				LongestBowKillKitSupportingTeamEcologist  int     `json:"longest_bow_kill_kit_supporting_team_ecologist"`
				MeleeKillsKitSupportingTeamEcologist      int     `json:"melee_kills_kit_supporting_team_ecologist"`
				MostKillsGameKitSupportingTeamEcologist   int     `json:"most_kills_game_kit_supporting_team_ecologist"`
				SurvivedPlayersKitSupportingTeamEcologist int     `json:"survived_players_kit_supporting_team_ecologist"`
				TimePlayedKitSupportingTeamEcologist      int     `json:"time_played_kit_supporting_team_ecologist"`
				WinsKitSupportingTeamEcologist            int     `json:"wins_kit_supporting_team_ecologist"`
				LossesKitSupportingTeamEcologist          int     `json:"losses_kit_supporting_team_ecologist"`
				VoidKillsKitSupportingTeamEcologist       int     `json:"void_kills_kit_supporting_team_ecologist"`
				AssistsKitDefendingTeamArmorer            int     `json:"assists_kit_defending_team_armorer"`
				ChestsOpenedKitDefendingTeamArmorer       int     `json:"chests_opened_kit_defending_team_armorer"`
				DeathsKitDefendingTeamArmorer             int     `json:"deaths_kit_defending_team_armorer"`
				FastestWinKitDefendingTeamArmorer         int     `json:"fastest_win_kit_defending_team_armorer"`
				GamesKitDefendingTeamArmorer              int     `json:"games_kit_defending_team_armorer"`
				KillsKitDefendingTeamArmorer              int     `json:"kills_kit_defending_team_armorer"`
				KillstreakKitDefendingTeamArmorer         int     `json:"killstreak_kit_defending_team_armorer"`
				LongestBowKillKitDefendingTeamArmorer     int     `json:"longest_bow_kill_kit_defending_team_armorer"`
				MeleeKillsKitDefendingTeamArmorer         int     `json:"melee_kills_kit_defending_team_armorer"`
				MostKillsGameKitDefendingTeamArmorer      int     `json:"most_kills_game_kit_defending_team_armorer"`
				SurvivedPlayersKitDefendingTeamArmorer    int     `json:"survived_players_kit_defending_team_armorer"`
				TimePlayedKitDefendingTeamArmorer         int     `json:"time_played_kit_defending_team_armorer"`
				WinsKitDefendingTeamArmorer               int     `json:"wins_kit_defending_team_armorer"`
				ArrowsShotKitDefendingTeamArmorer         int     `json:"arrows_shot_kit_defending_team_armorer"`
				LossesKitDefendingTeamArmorer             int     `json:"losses_kit_defending_team_armorer"`
				VoidKillsKitDefendingTeamArmorer          int     `json:"void_kills_kit_defending_team_armorer"`
				HeadsKitDefendingTeamArmorer              int     `json:"heads_kit_defending_team_armorer"`
				HeadsSalty                                int     `json:"heads_salty"`
				HeadsSaltyKitDefendingTeamArmorer         int     `json:"heads_salty_kit_defending_team_armorer"`
				HeadsSaltyTeam                            int     `json:"heads_salty_team"`
				FreeEventKeySkywarsHalloweenBoxes20192    bool    `json:"free_event_key_skywars_halloween_boxes_2019_2"`
				SkywarsHalloweenBoxes                     int     `json:"skywars_halloween_boxes"`
				TeamBulldozer                             int     `json:"team_bulldozer"`
				TeamFat                                   int     `json:"team_fat"`
				HeadsEww                                  int     `json:"heads_eww"`
				HeadsEwwKitSupportingTeamEcologist        int     `json:"heads_eww_kit_supporting_team_ecologist"`
				HeadsEwwSolo                              int     `json:"heads_eww_solo"`
				HeadsKitSupportingTeamEcologist           int     `json:"heads_kit_supporting_team_ecologist"`
				HeadsMehKitSupportingTeamEcologist        int     `json:"heads_meh_kit_supporting_team_ecologist"`
				HeadsMehSolo                              int     `json:"heads_meh_solo"`
				HeadsSaltyKitSupportingTeamEcologist      int     `json:"heads_salty_kit_supporting_team_ecologist"`
				HeadsSaltySolo                            int     `json:"heads_salty_solo"`
				HeadsSolo                                 int     `json:"heads_solo"`
				HeadsSucculent                            int     `json:"heads_succulent"`
				HeadsSucculentKitSupportingTeamEcologist  int     `json:"heads_succulent_kit_supporting_team_ecologist"`
				HeadsSucculentSolo                        int     `json:"heads_succulent_solo"`
				ArrowsHitKitSupportingTeamEcologist       int     `json:"arrows_hit_kit_supporting_team_ecologist"`
				ArrowsShotKitSupportingTeamEcologist      int     `json:"arrows_shot_kit_supporting_team_ecologist"`
				LongestBowShotKitSupportingTeamEcologist  int     `json:"longest_bow_shot_kit_supporting_team_ecologist"`
				BowKillsKitSupportingTeamEcologist        int     `json:"bow_kills_kit_supporting_team_ecologist"`
				HeadsDivine                               int     `json:"heads_divine"`
				HeadsDivineKitSupportingTeamEcologist     int     `json:"heads_divine_kit_supporting_team_ecologist"`
				HeadsDivineSolo                           int     `json:"heads_divine_solo"`
				HeadsTasty                                int     `json:"heads_tasty"`
				HeadsTastyKitSupportingTeamEcologist      int     `json:"heads_tasty_kit_supporting_team_ecologist"`
				HeadsTastySolo                            int     `json:"heads_tasty_solo"`
				ArrowsHitKitDefendingTeamArmorer          int     `json:"arrows_hit_kit_defending_team_armorer"`
				LongestBowShotKitDefendingTeamArmorer     int     `json:"longest_bow_shot_kit_defending_team_armorer"`
				HeadsHeavenly                             int     `json:"heads_heavenly"`
				HeadsHeavenlyKitSupportingTeamEcologist   int     `json:"heads_heavenly_kit_supporting_team_ecologist"`
				HeadsHeavenlySolo                         int     `json:"heads_heavenly_solo"`
				HeadsSucculentKitDefendingTeamArmorer     int     `json:"heads_succulent_kit_defending_team_armorer"`
				HeadsSucculentTeam                        int     `json:"heads_succulent_team"`
				TeamJuggernaut                            int     `json:"team_juggernaut"`
				SkyWarsSkywarsRating919Position           int     `json:"SkyWars_skywars_rating_9_19_position"`
				SkyWarsSkywarsRating919Rating             float64 `json:"SkyWars_skywars_rating_9_19_rating"`
				MegaTank                                  int     `json:"mega_tank"`
				HeadsTastyKitDefendingTeamArmorer         int     `json:"heads_tasty_kit_defending_team_armorer"`
				HeadsTastyTeam                            int     `json:"heads_tasty_team"`
				MegaEnvironmentalExpert                   int     `json:"mega_environmental_expert"`
				HeadsEwwKitDefendingTeamArmorer           int     `json:"heads_eww_kit_defending_team_armorer"`
				HeadsEwwTeam                              int     `json:"heads_eww_team"`
				HeadsHeavenlyKitDefendingTeamArmorer      int     `json:"heads_heavenly_kit_defending_team_armorer"`
				HeadsHeavenlyTeam                         int     `json:"heads_heavenly_team"`
				HeadsYucky                                int     `json:"heads_yucky"`
				HeadsYuckyKitDefendingTeamArmorer         int     `json:"heads_yucky_kit_defending_team_armorer"`
				HeadsYuckyTeam                            int     `json:"heads_yucky_team"`
				MegaBlackMagic                            int     `json:"mega_black_magic"`
				ActiveKillmessages                        string  `json:"active_killmessages"`
				SkyWarsOpenedEpics                        int     `json:"SkyWars_openedEpics"`
				MobKillsKitDefendingTeamArmorer           int     `json:"mob_kills_kit_defending_team_armorer"`
				MobKillsTeam                              int     `json:"mob_kills_team"`
				TeamSpeedBoost                            int     `json:"team_speed_boost"`
				KitRankedRankedChampionInventory          struct {
					IRONLEGGINGS0   string `json:"IRON_LEGGINGS:0"`
					IRONAXE0        string `json:"IRON_AXE:0"`
					ENCHANTEDBOOK0  string `json:"ENCHANTED_BOOK:0"`
					IRONPICKAXE0    string `json:"IRON_PICKAXE:0"`
					IRONHELMET0     string `json:"IRON_HELMET:0"`
					IRONBOOTS0      string `json:"IRON_BOOTS:0"`
					IRONCHESTPLATE0 string `json:"IRON_CHESTPLATE:0"`
					DIAMONDSWORD0   string `json:"DIAMOND_SWORD:0"`
					ANVIL0          string `json:"ANVIL:0"`
				} `json:"kit_ranked_ranked_champion_inventory"`
				DeathsKitRankedRankedChampion          int `json:"deaths_kit_ranked_ranked_champion"`
				LossesKitRankedRankedChampion          int `json:"losses_kit_ranked_ranked_champion"`
				TimePlayedKitRankedRankedChampion      int `json:"time_played_kit_ranked_ranked_champion"`
				SurvivedPlayersKitRankedRankedChampion int `json:"survived_players_kit_ranked_ranked_champion"`
				ChestsOpenedKitRankedRankedChampion    int `json:"chests_opened_kit_ranked_ranked_champion"`
				FastestWinKitRankedRankedChampion      int `json:"fastest_win_kit_ranked_ranked_champion"`
				GamesKitRankedRankedChampion           int `json:"games_kit_ranked_ranked_champion"`
				KillsKitRankedRankedChampion           int `json:"kills_kit_ranked_ranked_champion"`
				KillstreakKitRankedRankedChampion      int `json:"killstreak_kit_ranked_ranked_champion"`
				LongestBowKillKitRankedRankedChampion  int `json:"longest_bow_kill_kit_ranked_ranked_champion"`
				MeleeKillsKitRankedRankedChampion      int `json:"melee_kills_kit_ranked_ranked_champion"`
				MostKillsGameKitRankedRankedChampion   int `json:"most_kills_game_kit_ranked_ranked_champion"`
				WinsKitRankedRankedChampion            int `json:"wins_kit_ranked_ranked_champion"`
				ArrowsShotKitRankedRankedChampion      int `json:"arrows_shot_kit_ranked_ranked_champion"`
				AssistsKitRankedRankedChampion         int `json:"assists_kit_ranked_ranked_champion"`
				ArrowsHitKitRankedRankedChampion       int `json:"arrows_hit_kit_ranked_ranked_champion"`
				LongestBowShotKitRankedRankedChampion  int `json:"longest_bow_shot_kit_ranked_ranked_champion"`
				VoidKillsKitRankedRankedChampion       int `json:"void_kills_kit_ranked_ranked_champion"`
			} `json:"SkyWars"`
			Arcade struct {
				Coins                     float64 `json:"coins"`
				WeeklyCoinsB              int     `json:"weekly_coins_b"`
				MonthlyCoinsA             int     `json:"monthly_coins_a"`
				StampLevel                int     `json:"stamp_level"`
				TimeStamp                 int64   `json:"time_stamp"`
				MiniwallsActiveKit        string  `json:"miniwalls_activeKit"`
				MonthlyCoinsB             int     `json:"monthly_coins_b"`
				ArrowsHitMiniWalls        int     `json:"arrows_hit_mini_walls"`
				KillsMiniWalls            int     `json:"kills_mini_walls"`
				WinsMiniWalls             int     `json:"wins_mini_walls"`
				DeathsMiniWalls           int     `json:"deaths_mini_walls"`
				ArrowsShotMiniWalls       int     `json:"arrows_shot_mini_walls"`
				WitherDamageMiniWalls     int     `json:"wither_damage_mini_walls"`
				KillsOneinthequiver       int     `json:"kills_oneinthequiver"`
				BountyKillsOneinthequiver int     `json:"bounty_kills_oneinthequiver"`
				DeathsOneinthequiver      int     `json:"deaths_oneinthequiver"`
				RoundsSimonSays           int     `json:"rounds_simon_says"`
				GiftsGrinchSimulatorV2    int     `json:"gifts_grinch_simulator_v2"`
			} `json:"Arcade"`
			Bedwars struct {
				FirstJoin7                                       bool     `json:"first_join_7"`
				Experience                                       int      `json:"Experience"`
				BedwarsBoxes                                     int      `json:"bedwars_boxes"`
				GamesPlayedBedwars1                              int      `json:"games_played_bedwars_1"`
				EightTwoWinstreak                                int      `json:"eight_two_winstreak"`
				FinalDeathsBedwars                               int      `json:"final_deaths_bedwars"`
				GoldResourcesCollectedBedwars                    int      `json:"gold_resources_collected_bedwars"`
				EightTwoBedsBrokenBedwars                        int      `json:"eight_two_beds_broken_bedwars"`
				VoidKillsBedwars                                 int      `json:"void_kills_bedwars"`
				DiamondResourcesCollectedBedwars                 int      `json:"diamond_resources_collected_bedwars"`
				ResourcesCollectedBedwars                        int      `json:"resources_collected_bedwars"`
				EightTwoPermanentItemsPurchasedBedwars           int      `json:"eight_two_permanent _items_purchased_bedwars"`
				EightTwoResourcesCollectedBedwars                int      `json:"eight_two_resources_collected_bedwars"`
				EightTwoDiamondResourcesCollectedBedwars         int      `json:"eight_two_diamond_resources_collected_bedwars"`
				Coins                                            int      `json:"coins"`
				EightTwoKillsBedwars                             int      `json:"eight_two_kills_bedwars"`
				GamesPlayedBedwars                               int      `json:"games_played_bedwars"`
				PermanentItemsPurchasedBedwars                   int      `json:"permanent _items_purchased_bedwars"`
				EightTwoEntityAttackKillsBedwars                 int      `json:"eight_two_entity_attack_kills_bedwars"`
				EightTwoIronResourcesCollectedBedwars            int      `json:"eight_two_iron_resources_collected_bedwars"`
				BedsLostBedwars                                  int      `json:"beds_lost_bedwars"`
				EightTwoVoidKillsBedwars                         int      `json:"eight_two_void_kills_bedwars"`
				KillsBedwars                                     int      `json:"kills_bedwars"`
				EightTwoFinalDeathsBedwars                       int      `json:"eight_two_final_deaths_bedwars"`
				EightTwoVoidFinalDeathsBedwars                   int      `json:"eight_two_void_final_deaths_bedwars"`
				EntityAttackKillsBedwars                         int      `json:"entity_attack_kills_bedwars"`
				EightTwoItemsPurchasedBedwars                    int      `json:"eight_two__items_purchased_bedwars"`
				EightTwoGamesPlayedBedwars                       int      `json:"eight_two_games_played_bedwars"`
				EightTwoItemsPurchasedBedwars2                   int      `json:"eight_two_items_purchased_bedwars"`
				EightTwoBedsLostBedwars                          int      `json:"eight_two_beds_lost_bedwars"`
				VoidFinalDeathsBedwars                           int      `json:"void_final_deaths_bedwars"`
				EightTwoGoldResourcesCollectedBedwars            int      `json:"eight_two_gold_resources_collected_bedwars"`
				LossesBedwars                                    int      `json:"losses_bedwars"`
				ItemsPurchasedBedwars                            int      `json:"items_purchased_bedwars"`
				IronResourcesCollectedBedwars                    int      `json:"iron_resources_collected_bedwars"`
				BedsBrokenBedwars                                int      `json:"beds_broken_bedwars"`
				EightTwoLossesBedwars                            int      `json:"eight_two_losses_bedwars"`
				ItemsPurchasedBedwars2                           int      `json:"_items_purchased_bedwars"`
				EntityAttackFinalDeathsBedwars                   int      `json:"entity_attack_final_deaths_bedwars"`
				EmeraldResourcesCollectedBedwars                 int      `json:"emerald_resources_collected_bedwars"`
				EightTwoEmeraldResourcesCollectedBedwars         int      `json:"eight_two_emerald_resources_collected_bedwars"`
				EightTwoEntityAttackFinalKillsBedwars            int      `json:"eight_two_entity_attack_final_kills_bedwars"`
				EntityAttackFinalKillsBedwars                    int      `json:"entity_attack_final_kills_bedwars"`
				EightTwoEntityAttackFinalDeathsBedwars           int      `json:"eight_two_entity_attack_final_deaths_bedwars"`
				EightTwoVoidFinalKillsBedwars                    int      `json:"eight_two_void_final_kills_bedwars"`
				FinalKillsBedwars                                int      `json:"final_kills_bedwars"`
				VoidFinalKillsBedwars                            int      `json:"void_final_kills_bedwars"`
				EightTwoFinalKillsBedwars                        int      `json:"eight_two_final_kills_bedwars"`
				FourFourWinstreak                                int      `json:"four_four_winstreak"`
				Winstreak                                        int      `json:"winstreak"`
				FourFourBedsBrokenBedwars                        int      `json:"four_four_beds_broken_bedwars"`
				FourFourItemsPurchasedBedwars                    int      `json:"four_four__items_purchased_bedwars"`
				FourFourPermanentItemsPurchasedBedwars           int      `json:"four_four_permanent _items_purchased_bedwars"`
				FourFourVoidKillsBedwars                         int      `json:"four_four_void_kills_bedwars"`
				FourFourKillsBedwars                             int      `json:"four_four_kills_bedwars"`
				FourFourGoldResourcesCollectedBedwars            int      `json:"four_four_gold_resources_collected_bedwars"`
				FourFourIronResourcesCollectedBedwars            int      `json:"four_four_iron_resources_collected_bedwars"`
				FourFourGamesPlayedBedwars                       int      `json:"four_four_games_played_bedwars"`
				FourFourWinsBedwars                              int      `json:"four_four_wins_bedwars"`
				FourFourEntityAttackFinalKillsBedwars            int      `json:"four_four_entity_attack_final_kills_bedwars"`
				WinsBedwars                                      int      `json:"wins_bedwars"`
				FourFourResourcesCollectedBedwars                int      `json:"four_four_resources_collected_bedwars"`
				FourFourItemsPurchasedBedwars2                   int      `json:"four_four_items_purchased_bedwars"`
				FourFourEntityAttackKillsBedwars                 int      `json:"four_four_entity_attack_kills_bedwars"`
				FourFourFinalKillsBedwars                        int      `json:"four_four_final_kills_bedwars"`
				FourFourEmeraldResourcesCollectedBedwars         int      `json:"four_four_emerald_resources_collected_bedwars"`
				FourFourFinalDeathsBedwars                       int      `json:"four_four_final_deaths_bedwars"`
				FallKillsBedwars                                 int      `json:"fall_kills_bedwars"`
				FourFourVoidFinalDeathsBedwars                   int      `json:"four_four_void_final_deaths_bedwars"`
				FourFourBedsLostBedwars                          int      `json:"four_four_beds_lost_bedwars"`
				FourFourLossesBedwars                            int      `json:"four_four_losses_bedwars"`
				FourFourFallKillsBedwars                         int      `json:"four_four_fall_kills_bedwars"`
				VoidDeathsBedwars                                int      `json:"void_deaths_bedwars"`
				DeathsBedwars                                    int      `json:"deaths_bedwars"`
				FallDeathsBedwars                                int      `json:"fall_deaths_bedwars"`
				FourFourFallDeathsBedwars                        int      `json:"four_four_fall_deaths_bedwars"`
				FourFourDeathsBedwars                            int      `json:"four_four_deaths_bedwars"`
				FourFourDiamondResourcesCollectedBedwars         int      `json:"four_four_diamond_resources_collected_bedwars"`
				FourFourEntityAttackDeathsBedwars                int      `json:"four_four_entity_attack_deaths_bedwars"`
				FourFourVoidFinalKillsBedwars                    int      `json:"four_four_void_final_kills_bedwars"`
				EntityAttackDeathsBedwars                        int      `json:"entity_attack_deaths_bedwars"`
				FourFourVoidDeathsBedwars                        int      `json:"four_four_void_deaths_bedwars"`
				SelectedUltimate                                 string   `json:"selected_ultimate"`
				FourFourUltimateWinstreak                        int      `json:"four_four_ultimate_winstreak"`
				FourFourUltimateDeathsBedwars                    int      `json:"four_four_ultimate_deaths_bedwars"`
				FourFourUltimateFinalKillsBedwars                int      `json:"four_four_ultimate_final_kills_bedwars"`
				FourFourUltimatePermanentItemsPurchasedBedwars   int      `json:"four_four_ultimate_permanent _items_purchased_bedwars"`
				FourFourUltimateItemsPurchasedBedwars            int      `json:"four_four_ultimate__items_purchased_bedwars"`
				FourFourUltimateResourcesCollectedBedwars        int      `json:"four_four_ultimate_resources_collected_bedwars"`
				FourFourUltimateWinsBedwars                      int      `json:"four_four_ultimate_wins_bedwars"`
				FourFourUltimateItemsPurchasedBedwars2           int      `json:"four_four_ultimate_items_purchased_bedwars"`
				FourFourUltimateEntityAttackFinalKillsBedwars    int      `json:"four_four_ultimate_entity_attack_final_kills_bedwars"`
				FourFourUltimateDiamondResourcesCollectedBedwars int      `json:"four_four_ultimate_diamond_resources_collected_bedwars"`
				FourFourUltimateIronResourcesCollectedBedwars    int      `json:"four_four_ultimate_iron_resources_collected_bedwars"`
				FourFourUltimateGoldResourcesCollectedBedwars    int      `json:"four_four_ultimate_gold_resources_collected_bedwars"`
				FourFourUltimateGamesPlayedBedwars               int      `json:"four_four_ultimate_games_played_bedwars"`
				FourFourUltimateVoidDeathsBedwars                int      `json:"four_four_ultimate_void_deaths_bedwars"`
				BedwarsChristmasBoxes                            int      `json:"bedwars_christmas_boxes"`
				FourFourEntityAttackFinalDeathsBedwars           int      `json:"four_four_entity_attack_final_deaths_bedwars"`
				FallFinalKillsBedwars                            int      `json:"fall_final_kills_bedwars"`
				FourFourFallFinalKillsBedwars                    int      `json:"four_four_fall_final_kills_bedwars"`
				EightTwoUltimateWinstreak                        int      `json:"eight_two_ultimate_winstreak"`
				EightTwoUltimateIronResourcesCollectedBedwars    int      `json:"eight_two_ultimate_iron_resources_collected_bedwars"`
				EightTwoUltimateVoidFinalDeathsBedwars           int      `json:"eight_two_ultimate_void_final_deaths_bedwars"`
				EightTwoUltimateFinalDeathsBedwars               int      `json:"eight_two_ultimate_final_deaths_bedwars"`
				EightTwoUltimateLossesBedwars                    int      `json:"eight_two_ultimate_losses_bedwars"`
				EightTwoUltimateGoldResourcesCollectedBedwars    int      `json:"eight_two_ultimate_gold_resources_collected_bedwars"`
				EightTwoUltimateResourcesCollectedBedwars        int      `json:"eight_two_ultimate_resources_collected_bedwars"`
				EightTwoUltimateBedsLostBedwars                  int      `json:"eight_two_ultimate_beds_lost_bedwars"`
				EightTwoUltimateGamesPlayedBedwars               int      `json:"eight_two_ultimate_games_played_bedwars"`
				EntityExplosionDeathsBedwars                     int      `json:"entity_explosion_deaths_bedwars"`
				FourFourEntityExplosionDeathsBedwars             int      `json:"four_four_entity_explosion_deaths_bedwars"`
				EightTwoVoidDeathsBedwars                        int      `json:"eight_two_void_deaths_bedwars"`
				EightTwoWinsBedwars                              int      `json:"eight_two_wins_bedwars"`
				EightTwoFallKillsBedwars                         int      `json:"eight_two_fall_kills_bedwars"`
				EightTwoDeathsBedwars                            int      `json:"eight_two_deaths_bedwars"`
				EightTwoEntityAttackDeathsBedwars                int      `json:"eight_two_entity_attack_deaths_bedwars"`
				FreeEventKeyBedwarsChristmasBoxes2018            bool     `json:"free_event_key_bedwars_christmas_boxes_2018"`
				Packages                                         []string `json:"packages"`
				BedwarsOpenedChests                              int      `json:"Bedwars_openedChests"`
				BedwarsOpenedCommons                             int      `json:"Bedwars_openedCommons"`
				BedwarsOpenedRares                               int      `json:"Bedwars_openedRares"`
				ChestHistoryNew                                  []string `json:"chest_history_new"`
				ActiveProjectileTrail                            string   `json:"activeProjectileTrail"`
				ActiveNPCSkin                                    string   `json:"activeNPCSkin"`
				BedwarsOpenedLegendaries                         int      `json:"Bedwars_openedLegendaries"`
				ActiveVictoryDance                               string   `json:"activeVictoryDance"`
				ActiveGlyph                                      string   `json:"activeGlyph"`
				ActiveSprays                                     string   `json:"activeSprays"`
				BedwarsOpenedEpics                               int      `json:"Bedwars_openedEpics"`
				ActiveIslandTopper                               string   `json:"activeIslandTopper"`
				ActiveKillMessages                               string   `json:"activeKillMessages"`
				ActiveDeathCry                                   string   `json:"activeDeathCry"`
				ActiveBedDestroy                                 string   `json:"activeBedDestroy"`
				FallFinalDeathsBedwars                           int      `json:"fall_final_deaths_bedwars"`
				FourFourFallFinalDeathsBedwars                   int      `json:"four_four_fall_final_deaths_bedwars"`
				LastHytaleAd                                     int64    `json:"lastHytaleAd"`
				FourThreeWinstreak                               int      `json:"four_three_winstreak"`
				FourThreeVoidKillsBedwars                        int      `json:"four_three_void_kills_bedwars"`
				FourThreeEntityAttackKillsBedwars                int      `json:"four_three_entity_attack_kills_bedwars"`
				FourThreeFallKillsBedwars                        int      `json:"four_three_fall_kills_bedwars"`
				FourThreeIronResourcesCollectedBedwars           int      `json:"four_three_iron_resources_collected_bedwars"`
				FourThreeEntityAttackFinalDeathsBedwars          int      `json:"four_three_entity_attack_final_deaths_bedwars"`
				FourThreeGamesPlayedBedwars                      int      `json:"four_three_games_played_bedwars"`
				FourThreeItemsPurchasedBedwars                   int      `json:"four_three__items_purchased_bedwars"`
				FourThreeVoidDeathsBedwars                       int      `json:"four_three_void_deaths_bedwars"`
				FourThreeBedsLostBedwars                         int      `json:"four_three_beds_lost_bedwars"`
				FourThreeGoldResourcesCollectedBedwars           int      `json:"four_three_gold_resources_collected_bedwars"`
				FourThreePermanentItemsPurchasedBedwars          int      `json:"four_three_permanent _items_purchased_bedwars"`
				FourThreeDeathsBedwars                           int      `json:"four_three_deaths_bedwars"`
				FourThreeLossesBedwars                           int      `json:"four_three_losses_bedwars"`
				FourThreeItemsPurchasedBedwars2                  int      `json:"four_three_items_purchased_bedwars"`
				FourThreeKillsBedwars                            int      `json:"four_three_kills_bedwars"`
				FourThreeFinalDeathsBedwars                      int      `json:"four_three_final_deaths_bedwars"`
				FourThreeResourcesCollectedBedwars               int      `json:"four_three_resources_collected_bedwars"`
				FourThreeEntityAttackDeathsBedwars               int      `json:"four_three_entity_attack_deaths_bedwars"`
				FourThreeEntityAttackFinalKillsBedwars           int      `json:"four_three_entity_attack_final_kills_bedwars"`
				FourThreeFinalKillsBedwars                       int      `json:"four_three_final_kills_bedwars"`
				FourThreeVoidFinalDeathsBedwars                  int      `json:"four_three_void_final_deaths_bedwars"`
				UnderstandsResourceBank                          bool     `json:"understands_resource_bank"`
				CastleWinstreak                                  int      `json:"castle_winstreak"`
				CastleItemsPurchasedBedwars                      int      `json:"castle_items_purchased_bedwars"`
				CastleDeathsBedwars                              int      `json:"castle_deaths_bedwars"`
				CastleGoldResourcesCollectedBedwars              int      `json:"castle_gold_resources_collected_bedwars"`
				CastleItemsPurchasedBedwars2                     int      `json:"castle__items_purchased_bedwars"`
				CastleResourcesCollectedBedwars                  int      `json:"castle_resources_collected_bedwars"`
				CastleWinsBedwars                                int      `json:"castle_wins_bedwars"`
				CastleVoidDeathsBedwars                          int      `json:"castle_void_deaths_bedwars"`
				CastleIronResourcesCollectedBedwars              int      `json:"castle_iron_resources_collected_bedwars"`
				FourFourFireDeathsBedwars                        int      `json:"four_four_fire_deaths_bedwars"`
				FireDeathsBedwars                                int      `json:"fire_deaths_bedwars"`
				FourFourProjectileDeathsBedwars                  int      `json:"four_four_projectile_deaths_bedwars"`
				ProjectileDeathsBedwars                          int      `json:"projectile_deaths_bedwars"`
				FourFourRushWinstreak                            int      `json:"four_four_rush_winstreak"`
				FourFourRushKillsBedwars                         int      `json:"four_four_rush_kills_bedwars"`
				FourFourRushLossesBedwars                        int      `json:"four_four_rush_losses_bedwars"`
				FourFourRushItemsPurchasedBedwars                int      `json:"four_four_rush_items_purchased_bedwars"`
				FourFourRushGamesPlayedBedwars                   int      `json:"four_four_rush_games_played_bedwars"`
				FourFourRushEntityAttackFinalDeathsBedwars       int      `json:"four_four_rush_entity_attack_final_deaths_bedwars"`
				FourFourRushEmeraldResourcesCollectedBedwars     int      `json:"four_four_rush_emerald_resources_collected_bedwars"`
				FourFourRushBedsLostBedwars                      int      `json:"four_four_rush_beds_lost_bedwars"`
				FourFourRushGoldResourcesCollectedBedwars        int      `json:"four_four_rush_gold_resources_collected_bedwars"`
				FourFourRushResourcesCollectedBedwars            int      `json:"four_four_rush_resources_collected_bedwars"`
				FourFourRushFinalDeathsBedwars                   int      `json:"four_four_rush_final_deaths_bedwars"`
				FourFourRushEntityAttackKillsBedwars             int      `json:"four_four_rush_entity_attack_kills_bedwars"`
				FourFourRushItemsPurchasedBedwars2               int      `json:"four_four_rush__items_purchased_bedwars"`
				FourFourRushPermanentItemsPurchasedBedwars       int      `json:"four_four_rush_permanent _items_purchased_bedwars"`
				FourFourRushIronResourcesCollectedBedwars        int      `json:"four_four_rush_iron_resources_collected_bedwars"`
				FourFourRushDeathsBedwars                        int      `json:"four_four_rush_deaths_bedwars"`
				FourFourRushEntityAttackDeathsBedwars            int      `json:"four_four_rush_entity_attack_deaths_bedwars"`
				FourFourRushVoidKillsBedwars                     int      `json:"four_four_rush_void_kills_bedwars"`
				FourFourRushVoidFinalDeathsBedwars               int      `json:"four_four_rush_void_final_deaths_bedwars"`
				FourFourRushVoidDeathsBedwars                    int      `json:"four_four_rush_void_deaths_bedwars"`
				ActiveKillEffect                                 string   `json:"activeKillEffect"`
				EightOneWinstreak                                int      `json:"eight_one_winstreak"`
				EightOneItemsPurchasedBedwars                    int      `json:"eight_one_items_purchased_bedwars"`
				EightOneVoidDeathsBedwars                        int      `json:"eight_one_void_deaths_bedwars"`
				EightOneDeathsBedwars                            int      `json:"eight_one_deaths_bedwars"`
				EightOneGamesPlayedBedwars                       int      `json:"eight_one_games_played_bedwars"`
				EightOneLossesBedwars                            int      `json:"eight_one_losses_bedwars"`
				EightOneIronResourcesCollectedBedwars            int      `json:"eight_one_iron_resources_collected_bedwars"`
				EightOneResourcesCollectedBedwars                int      `json:"eight_one_resources_collected_bedwars"`
				EightOneItemsPurchasedBedwars2                   int      `json:"eight_one__items_purchased_bedwars"`
				EightOneGoldResourcesCollectedBedwars            int      `json:"eight_one_gold_resources_collected_bedwars"`
				FourThreeVoidFinalKillsBedwars                   int      `json:"four_three_void_final_kills_bedwars"`
				FourThreeBedsBrokenBedwars                       int      `json:"four_three_beds_broken_bedwars"`
				EightOneBedsLostBedwars                          int      `json:"eight_one_beds_lost_bedwars"`
				EightOneFinalDeathsBedwars                       int      `json:"eight_one_final_deaths_bedwars"`
				EightOneEntityAttackFinalDeathsBedwars           int      `json:"eight_one_entity_attack_final_deaths_bedwars"`
				EightOneEntityAttackKillsBedwars                 int      `json:"eight_one_entity_attack_kills_bedwars"`
				EightOneKillsBedwars                             int      `json:"eight_one_kills_bedwars"`
				EightOneEntityAttackFinalKillsBedwars            int      `json:"eight_one_entity_attack_final_kills_bedwars"`
				EightOneEntityAttackDeathsBedwars                int      `json:"eight_one_entity_attack_deaths_bedwars"`
				EightOneFinalKillsBedwars                        int      `json:"eight_one_final_kills_bedwars"`
				EightOneBedsBrokenBedwars                        int      `json:"eight_one_beds_broken_bedwars"`
				EightOneEmeraldResourcesCollectedBedwars         int      `json:"eight_one_emerald_resources_collected_bedwars"`
				EightOneDiamondResourcesCollectedBedwars         int      `json:"eight_one_diamond_resources_collected_bedwars"`
				EightOnePermanentItemsPurchasedBedwars           int      `json:"eight_one_permanent _items_purchased_bedwars"`
				EightOneVoidFinalKillsBedwars                    int      `json:"eight_one_void_final_kills_bedwars"`
				EightOneWinsBedwars                              int      `json:"eight_one_wins_bedwars"`
				BedwarsLunarBoxes                                int      `json:"bedwars_lunar_boxes"`
				FreeEventKeyBedwarsLunarBoxes2019                bool     `json:"free_event_key_bedwars_lunar_boxes_2019"`
				FourThreeWinsBedwars                             int      `json:"four_three_wins_bedwars"`
				FourThreeFallFinalKillsBedwars                   int      `json:"four_three_fall_final_kills_bedwars"`
				FourThreeFallDeathsBedwars                       int      `json:"four_three_fall_deaths_bedwars"`
				EightOneVoidKillsBedwars                         int      `json:"eight_one_void_kills_bedwars"`
				FourFourLuckyWinstreak                           int      `json:"four_four_lucky_winstreak"`
				FourFourLuckyEntityAttackDeathsBedwars           int      `json:"four_four_lucky_entity_attack_deaths_bedwars"`
				FourFourLuckyVoidDeathsBedwars                   int      `json:"four_four_lucky_void_deaths_bedwars"`
				FourFourLuckyItemsPurchasedBedwars               int      `json:"four_four_lucky_items_purchased_bedwars"`
				FourFourLuckyFallKillsBedwars                    int      `json:"four_four_lucky_fall_kills_bedwars"`
				FourFourLuckyLossesBedwars                       int      `json:"four_four_lucky_losses_bedwars"`
				FourFourLuckyGamesPlayedBedwars                  int      `json:"four_four_lucky_games_played_bedwars"`
				FourFourLuckyEntityAttackKillsBedwars            int      `json:"four_four_lucky_entity_attack_kills_bedwars"`
				FourFourLuckyItemsPurchasedBedwars2              int      `json:"four_four_lucky__items_purchased_bedwars"`
				FourFourLuckyKillsBedwars                        int      `json:"four_four_lucky_kills_bedwars"`
				FourFourLuckyResourcesCollectedBedwars           int      `json:"four_four_lucky_resources_collected_bedwars"`
				FourFourLuckyDeathsBedwars                       int      `json:"four_four_lucky_deaths_bedwars"`
				FourFourLuckyIronResourcesCollectedBedwars       int      `json:"four_four_lucky_iron_resources_collected_bedwars"`
				FourFourLuckyVoidKillsBedwars                    int      `json:"four_four_lucky_void_kills_bedwars"`
				FourFourLuckyGoldResourcesCollectedBedwars       int      `json:"four_four_lucky_gold_resources_collected_bedwars"`
				Favourites2                                      string   `json:"favourites_2"`
				EightOneFallKillsBedwars                         int      `json:"eight_one_fall_kills_bedwars"`
				EightOneVoidFinalDeathsBedwars                   int      `json:"eight_one_void_final_deaths_bedwars"`
				EightOneFallDeathsBedwars                        int      `json:"eight_one_fall_deaths_bedwars"`
				FourFourEntityExplosionKillsBedwars              int      `json:"four_four_entity_explosion_kills_bedwars"`
				EntityExplosionKillsBedwars                      int      `json:"entity_explosion_kills_bedwars"`
				FourThreeEmeraldResourcesCollectedBedwars        int      `json:"four_three_emerald_resources_collected_bedwars"`
				EntityExplosionFinalKillsBedwars                 int      `json:"entity_explosion_final_kills_bedwars"`
				FourFourEntityExplosionFinalKillsBedwars         int      `json:"four_four_entity_explosion_final_kills_bedwars"`
				FireTickKillsBedwars                             int      `json:"fire_tick_kills_bedwars"`
				FourFourFireTickKillsBedwars                     int      `json:"four_four_fire_tick_kills_bedwars"`
				FourFourEntityExplosionFinalDeathsBedwars        int      `json:"four_four_entity_explosion_final_deaths_bedwars"`
				EntityExplosionFinalDeathsBedwars                int      `json:"entity_explosion_final_deaths_bedwars"`
				FourFourProjectileFinalDeathsBedwars             int      `json:"four_four_projectile_final_deaths_bedwars"`
				ProjectileFinalDeathsBedwars                     int      `json:"projectile_final_deaths_bedwars"`
				EightOneFallFinalDeathsBedwars                   int      `json:"eight_one_fall_final_deaths_bedwars"`
				FourFourLuckyFinalDeathsBedwars                  int      `json:"four_four_lucky_final_deaths_bedwars"`
				FourFourLuckyBedsLostBedwars                     int      `json:"four_four_lucky_beds_lost_bedwars"`
				FourFourLuckyEntityAttackFinalDeathsBedwars      int      `json:"four_four_lucky_entity_attack_final_deaths_bedwars"`
				FourFourLuckyPermanentItemsPurchasedBedwars      int      `json:"four_four_lucky_permanent _items_purchased_bedwars"`
				FourFourLuckyEntityAttackFinalKillsBedwars       int      `json:"four_four_lucky_entity_attack_final_kills_bedwars"`
				FourFourLuckyBedsBrokenBedwars                   int      `json:"four_four_lucky_beds_broken_bedwars"`
				FourFourLuckyFinalKillsBedwars                   int      `json:"four_four_lucky_final_kills_bedwars"`
				FourFourLuckyVoidFinalKillsBedwars               int      `json:"four_four_lucky_void_final_kills_bedwars"`
				FourFourLuckyProjectileDeathsBedwars             int      `json:"four_four_lucky_projectile_deaths_bedwars"`
				FourFourLuckyDiamondResourcesCollectedBedwars    int      `json:"four_four_lucky_diamond_resources_collected_bedwars"`
				FourFourLuckyEmeraldResourcesCollectedBedwars    int      `json:"four_four_lucky_emerald_resources_collected_bedwars"`
				FourFourLuckyFireTickDeathsBedwars               int      `json:"four_four_lucky_fire_tick_deaths_bedwars"`
				FourFourLuckyWinsBedwars                         int      `json:"four_four_lucky_wins_bedwars"`
				FourFourLuckyProjectileFinalKillsBedwars         int      `json:"four_four_lucky_projectile_final_kills_bedwars"`
				FourFourLuckyFireTickFinalKillsBedwars           int      `json:"four_four_lucky_fire_tick_final_kills_bedwars"`
				EightTwoLuckyWinstreak                           int      `json:"eight_two_lucky_winstreak"`
				EightTwoLuckyDeathsBedwars                       int      `json:"eight_two_lucky_deaths_bedwars"`
				EightTwoLuckyEntityExplosionFinalKillsBedwars    int      `json:"eight_two_lucky_entity_explosion_final_kills_bedwars"`
				EightTwoLuckyVoidKillsBedwars                    int      `json:"eight_two_lucky_void_kills_bedwars"`
				EightTwoLuckyBedsBrokenBedwars                   int      `json:"eight_two_lucky_beds_broken_bedwars"`
				EightTwoLuckyEntityAttackFinalKillsBedwars       int      `json:"eight_two_lucky_entity_attack_final_kills_bedwars"`
				EightTwoLuckyIronResourcesCollectedBedwars       int      `json:"eight_two_lucky_iron_resources_collected_bedwars"`
				EightTwoLuckyDiamondResourcesCollectedBedwars    int      `json:"eight_two_lucky_diamond_resources_collected_bedwars"`
				EightTwoLuckyWinsBedwars                         int      `json:"eight_two_lucky_wins_bedwars"`
				EightTwoLuckyEntityAttackKillsBedwars            int      `json:"eight_two_lucky_entity_attack_kills_bedwars"`
				EightTwoLuckyVoidDeathsBedwars                   int      `json:"eight_two_lucky_void_deaths_bedwars"`
				EightTwoLuckyKillsBedwars                        int      `json:"eight_two_lucky_kills_bedwars"`
				EightTwoLuckyGoldResourcesCollectedBedwars       int      `json:"eight_two_lucky_gold_resources_collected_bedwars"`
				EightTwoLuckyResourcesCollectedBedwars           int      `json:"eight_two_lucky_resources_collected_bedwars"`
				EightTwoLuckyFinalKillsBedwars                   int      `json:"eight_two_lucky_final_kills_bedwars"`
				EightTwoLuckyEmeraldResourcesCollectedBedwars    int      `json:"eight_two_lucky_emerald_resources_collected_bedwars"`
				EightTwoLuckyItemsPurchasedBedwars               int      `json:"eight_two_lucky_items_purchased_bedwars"`
				EightTwoLuckyVoidFinalKillsBedwars               int      `json:"eight_two_lucky_void_final_kills_bedwars"`
				EightTwoLuckyEntityAttackDeathsBedwars           int      `json:"eight_two_lucky_entity_attack_deaths_bedwars"`
				EightTwoLuckyGamesPlayedBedwars                  int      `json:"eight_two_lucky_games_played_bedwars"`
				EightTwoLuckyItemsPurchasedBedwars2              int      `json:"eight_two_lucky__items_purchased_bedwars"`
				EightTwoLuckyEntityAttackFinalDeathsBedwars      int      `json:"eight_two_lucky_entity_attack_final_deaths_bedwars"`
				EightTwoLuckyFinalDeathsBedwars                  int      `json:"eight_two_lucky_final_deaths_bedwars"`
				EightTwoLuckyFireKillsBedwars                    int      `json:"eight_two_lucky_fire_kills_bedwars"`
				EightTwoLuckyLossesBedwars                       int      `json:"eight_two_lucky_losses_bedwars"`
				EightTwoLuckyPermanentItemsPurchasedBedwars      int      `json:"eight_two_lucky_permanent _items_purchased_bedwars"`
				EightTwoLuckyBedsLostBedwars                     int      `json:"eight_two_lucky_beds_lost_bedwars"`
				EightTwoLuckyFireTickKillsBedwars                int      `json:"eight_two_lucky_fire_tick_kills_bedwars"`
				EightTwoLuckyVoidFinalDeathsBedwars              int      `json:"eight_two_lucky_void_final_deaths_bedwars"`
				EightTwoLuckyProjectileKillsBedwars              int      `json:"eight_two_lucky_projectile_kills_bedwars"`
				EightTwoLuckyFireTickFinalKillsBedwars           int      `json:"eight_two_lucky_fire_tick_final_kills_bedwars"`
				EightTwoLuckyFallDeathsBedwars                   int      `json:"eight_two_lucky_fall_deaths_bedwars"`
				EightTwoLuckyFallFinalKillsBedwars               int      `json:"eight_two_lucky_fall_final_kills_bedwars"`
				EightTwoLuckyProjectileDeathsBedwars             int      `json:"eight_two_lucky_projectile_deaths_bedwars"`
				EightTwoLuckyFireTickDeathsBedwars               int      `json:"eight_two_lucky_fire_tick_deaths_bedwars"`
				EightTwoLuckyFireTickFinalDeathsBedwars          int      `json:"eight_two_lucky_fire_tick_final_deaths_bedwars"`
				FourFourLuckyFallDeathsBedwars                   int      `json:"four_four_lucky_fall_deaths_bedwars"`
				FourFourLuckyProjectileKillsBedwars              int      `json:"four_four_lucky_projectile_kills_bedwars"`
				FourFourLuckyVoidFinalDeathsBedwars              int      `json:"four_four_lucky_void_final_deaths_bedwars"`
				EightTwoFallDeathsBedwars                        int      `json:"eight_two_fall_deaths_bedwars"`
				EightTwoEntityExplosionKillsBedwars              int      `json:"eight_two_entity_explosion_kills_bedwars"`
				EightTwoFallFinalKillsBedwars                    int      `json:"eight_two_fall_final_kills_bedwars"`
				EightTwoFallFinalDeathsBedwars                   int      `json:"eight_two_fall_final_deaths_bedwars"`
				EightTwoEntityExplosionDeathsBedwars             int      `json:"eight_two_entity_explosion_deaths_bedwars"`
				FourThreeDiamondResourcesCollectedBedwars        int      `json:"four_three_diamond_resources_collected_bedwars"`
				EightOneEntityExplosionFinalDeathsBedwars        int      `json:"eight_one_entity_explosion_final_deaths_bedwars"`
				FourThreeFallFinalDeathsBedwars                  int      `json:"four_three_fall_final_deaths_bedwars"`
				VotedSnowman                                     bool     `json:"voted_snowman"`
				FourFourLuckyFallFinalKillsBedwars               int      `json:"four_four_lucky_fall_final_kills_bedwars"`
				FourThreeEntityExplosionDeathsBedwars            int      `json:"four_three_entity_explosion_deaths_bedwars"`
				FourThreeEntityExplosionKillsBedwars             int      `json:"four_three_entity_explosion_kills_bedwars"`
				EightTwoLuckyProjectileFinalKillsBedwars         int      `json:"eight_two_lucky_projectile_final_kills_bedwars"`
				EightTwoLuckyEntityExplosionKillsBedwars         int      `json:"eight_two_lucky_entity_explosion_kills_bedwars"`
				EightOneEntityExplosionDeathsBedwars             int      `json:"eight_one_entity_explosion_deaths_bedwars"`
				BedwarsEasterBoxes                               int      `json:"bedwars_easter_boxes"`
				BedwarsHalloweenBoxes                            int      `json:"bedwars_halloween_boxes"`
				FreeEventKeyBedwarsHalloweenBoxes2019            bool     `json:"free_event_key_bedwars_halloween_boxes_2019"`
			} `json:"Bedwars"`
			TNTGames struct {
				Wins                        int      `json:"wins"`
				NewPvprunDoubleJumps        int      `json:"new_pvprun_double_jumps"`
				NewSpleefRepulsor           int      `json:"new_spleef_repulsor"`
				NewTntagSpeedy              int      `json:"new_tntag_speedy"`
				NewTntrunDoubleJumps        int      `json:"new_tntrun_double_jumps"`
				NewBloodwizardRegen         int      `json:"new_bloodwizard_regen"`
				NewFirewizardExplode        int      `json:"new_firewizard_explode"`
				NewWitherwizardExplode      int      `json:"new_witherwizard_explode"`
				NewWitherwizardRegen        int      `json:"new_witherwizard_regen"`
				Packages                    []string `json:"packages"`
				NewSpleefDoubleJumps        int      `json:"new_spleef_double_jumps"`
				NewIcewizardRegen           int      `json:"new_icewizard_regen"`
				NewBloodwizardExplode       int      `json:"new_bloodwizard_explode"`
				NewSpleefTripleshot         int      `json:"new_spleef_tripleshot"`
				NewIcewizardExplode         int      `json:"new_icewizard_explode"`
				NewKineticwizardExplode     int      `json:"new_kineticwizard_explode"`
				NewKineticwizardRegen       int      `json:"new_kineticwizard_regen"`
				NewFirewizardRegen          int      `json:"new_firewizard_regen"`
				Coins                       int      `json:"coins"`
				Winstreak                   int      `json:"winstreak"`
				KillsTntag                  int      `json:"kills_tntag"`
				RunPotionsSplashedOnPlayers int      `json:"run_potions_splashed_on_players"`
				RecordTntrun                int      `json:"record_tntrun"`
				DeathsTntrun                int      `json:"deaths_tntrun"`
				WinsTntag                   int      `json:"wins_tntag"`
				VotesIndustrial             int      `json:"votes_Industrial"`
				TagBlastprotection          int      `json:"tag_blastprotection"`
				TagSpeeditup                int      `json:"tag_speeditup"`
				VotesSciFi                  int      `json:"votes_Sci-Fi"`
				TagSlowitdown               int      `json:"tag_slowitdown"`
				VotesIcicle                 int      `json:"votes_Icicle"`
				WinsTntrun                  int      `json:"wins_tntrun"`
				RecordPvprun                int      `json:"record_pvprun"`
				Flags                       struct {
					GiveDjFeather           bool `json:"give_dj_feather"`
					ShowTipHolograms        bool `json:"show_tip_holograms"`
					ShowTntrunActionbarInfo bool `json:"show_tntrun_actionbar_info"`
					ShowTnttagActionbarInfo bool `json:"show_tnttag_actionbar_info"`
				} `json:"flags"`
			} `json:"TNTGames"`
			Legacy struct {
				NextTokensSeconds int `json:"next_tokens_seconds"`
				ArenaTokens       int `json:"arena_tokens"`
				TotalTokens       int `json:"total_tokens"`
				Tokens            int `json:"tokens"`
			} `json:"Legacy"`
			Arena struct {
				WinStreaks2V2 int `json:"win_streaks_2v2"`
				Deaths2V2     int `json:"deaths_2v2"`
				Coins         int `json:"coins"`
				Losses2V2     int `json:"losses_2v2"`
				Damage2V2     int `json:"damage_2v2"`
				Healed2V2     int `json:"healed_2v2"`
				Games2V2      int `json:"games_2v2"`
				Wins          int `json:"wins"`
				Wins2V2       int `json:"wins_2v2"`
			} `json:"Arena"`
			GingerBread struct {
				ShoesActive   string   `json:"shoes_active"`
				BoosterActive string   `json:"booster_active"`
				Packages      []string `json:"packages"`
				FrameActive   string   `json:"frame_active"`
				PantsActive   string   `json:"pants_active"`
				EngineActive  string   `json:"engine_active"`
				HelmetActive  string   `json:"helmet_active"`
				JacketActive  string   `json:"jacket_active"`
				SkinActive    string   `json:"skin_active"`
				Coins         int      `json:"coins"`
			} `json:"GingerBread"`
			VampireZ struct {
				UpdatedStats bool `json:"updated_stats"`
				Coins        int  `json:"coins"`
			} `json:"VampireZ"`
			Paintball struct {
				Packages []string `json:"packages"`
				Coins    int      `json:"coins"`
			} `json:"Paintball"`
			Quake struct {
				Packages []string `json:"packages"`
				Coins    int      `json:"coins"`
			} `json:"Quake"`
			MurderMystery struct {
				MurdermysteryBooks                                     []string `json:"murdermystery_books"`
				CoinsPickedupMURDERCLASSIC                             int      `json:"coins_pickedup_MURDER_CLASSIC"`
				Wins                                                   int      `json:"wins"`
				WinsWidowSDenMURDERCLASSIC                             int      `json:"wins_widow's_den_MURDER_CLASSIC"`
				Coins                                                  int      `json:"coins"`
				WinsWidowSDen                                          int      `json:"wins_widow's_den"`
				CoinsPickedupWidowSDen                                 int      `json:"coins_pickedup_widow's_den"`
				GamesWidowSDenMURDERCLASSIC                            int      `json:"games_widow's_den_MURDER_CLASSIC"`
				Games                                                  int      `json:"games"`
				GamesMURDERCLASSIC                                     int      `json:"games_MURDER_CLASSIC"`
				CoinsPickedup                                          int      `json:"coins_pickedup"`
				CoinsPickedupWidowSDenMURDERCLASSIC                    int      `json:"coins_pickedup_widow's_den_MURDER_CLASSIC"`
				WinsMURDERCLASSIC                                      int      `json:"wins_MURDER_CLASSIC"`
				GamesWidowSDen                                         int      `json:"games_widow's_den"`
				DetectiveChance                                        int      `json:"detective_chance"`
				MurdererChance                                         int      `json:"murderer_chance"`
				GrantedChests                                          int      `json:"granted_chests"`
				MmChests                                               int      `json:"mm_chests"`
				Packages                                               []string `json:"packages"`
				MurderMysteryOpenedChests                              int      `json:"MurderMystery_openedChests"`
				MurderMysteryOpenedCommons                             int      `json:"MurderMystery_openedCommons"`
				MurderMysteryOpenedRares                               int      `json:"MurderMystery_openedRares"`
				ChestHistoryNew                                        []string `json:"chest_history_new"`
				ActiveKillNote                                         string   `json:"active_kill_note"`
				ActiveAnimatedHat                                      string   `json:"active_animated_hat"`
				MmChristmasChests                                      int      `json:"mm_christmas_chests"`
				QuickestDetectiveWinTimeSecondsMURDERCLASSIC           int      `json:"quickest_detective_win_time_seconds_MURDER_CLASSIC"`
				QuickestDetectiveWinTimeSecondsCruiseShip              int      `json:"quickest_detective_win_time_seconds_cruise_ship"`
				QuickestDetectiveWinTimeSeconds                        int      `json:"quickest_detective_win_time_seconds"`
				QuickestDetectiveWinTimeSecondsCruiseShipMURDERCLASSIC int      `json:"quickest_detective_win_time_seconds_cruise_ship_MURDER_CLASSIC"`
				CoinsPickedupCruiseShipMURDERCLASSIC                   int      `json:"coins_pickedup_cruise_ship_MURDER_CLASSIC"`
				GamesCruiseShipMURDERCLASSIC                           int      `json:"games_cruise_ship_MURDER_CLASSIC"`
				CoinsPickedupCruiseShip                                int      `json:"coins_pickedup_cruise_ship"`
				DetectiveWinsCruiseShip                                int      `json:"detective_wins_cruise_ship"`
				GamesCruiseShip                                        int      `json:"games_cruise_ship"`
				DetectiveWins                                          int      `json:"detective_wins"`
				WinsCruiseShipMURDERCLASSIC                            int      `json:"wins_cruise_ship_MURDER_CLASSIC"`
				DetectiveWinsCruiseShipMURDERCLASSIC                   int      `json:"detective_wins_cruise_ship_MURDER_CLASSIC"`
				DetectiveWinsMURDERCLASSIC                             int      `json:"detective_wins_MURDER_CLASSIC"`
				WinsCruiseShip                                         int      `json:"wins_cruise_ship"`
				CoinsPickedupTransport                                 int      `json:"coins_pickedup_transport"`
				DeathsTransportMURDERCLASSIC                           int      `json:"deaths_transport_MURDER_CLASSIC"`
				GamesTransportMURDERCLASSIC                            int      `json:"games_transport_MURDER_CLASSIC"`
				Deaths                                                 int      `json:"deaths"`
				GamesTransport                                         int      `json:"games_transport"`
				DeathsTransport                                        int      `json:"deaths_transport"`
				CoinsPickedupTransportMURDERCLASSIC                    int      `json:"coins_pickedup_transport_MURDER_CLASSIC"`
				DeathsMURDERCLASSIC                                    int      `json:"deaths_MURDER_CLASSIC"`
				DeathsWidowSDen                                        int      `json:"deaths_widow's_den"`
				DeathsWidowSDenMURDERCLASSIC                           int      `json:"deaths_widow's_den_MURDER_CLASSIC"`
				GamesHypixelWorldMURDERCLASSIC                         int      `json:"games_hypixel_world_MURDER_CLASSIC"`
				WinsHypixelWorldMURDERCLASSIC                          int      `json:"wins_hypixel_world_MURDER_CLASSIC"`
				GamesHypixelWorld                                      int      `json:"games_hypixel_world"`
				CoinsPickedupHypixelWorldMURDERCLASSIC                 int      `json:"coins_pickedup_hypixel_world_MURDER_CLASSIC"`
				CoinsPickedupHypixelWorld                              int      `json:"coins_pickedup_hypixel_world"`
				WinsHypixelWorld                                       int      `json:"wins_hypixel_world"`
				WinsHollywood                                          int      `json:"wins_hollywood"`
				WinsHollywoodMURDERCLASSIC                             int      `json:"wins_hollywood_MURDER_CLASSIC"`
				CoinsPickedupHollywoodMURDERCLASSIC                    int      `json:"coins_pickedup_hollywood_MURDER_CLASSIC"`
				CoinsPickedupHollywood                                 int      `json:"coins_pickedup_hollywood"`
				DeathsHollywoodMURDERCLASSIC                           int      `json:"deaths_hollywood_MURDER_CLASSIC"`
				DeathsHollywood                                        int      `json:"deaths_hollywood"`
				GamesHollywood                                         int      `json:"games_hollywood"`
				GamesHollywoodMURDERCLASSIC                            int      `json:"games_hollywood_MURDER_CLASSIC"`
				GamesAncientTomb                                       int      `json:"games_ancient_tomb"`
				GamesAncientTombMURDERCLASSIC                          int      `json:"games_ancient_tomb_MURDER_CLASSIC"`
				WinsAncientTombMURDERCLASSIC                           int      `json:"wins_ancient_tomb_MURDER_CLASSIC"`
				CoinsPickedupAncientTombMURDERCLASSIC                  int      `json:"coins_pickedup_ancient_tomb_MURDER_CLASSIC"`
				DeathsAncientTombMURDERCLASSIC                         int      `json:"deaths_ancient_tomb_MURDER_CLASSIC"`
				DeathsAncientTomb                                      int      `json:"deaths_ancient_tomb"`
				CoinsPickedupAncientTomb                               int      `json:"coins_pickedup_ancient_tomb"`
				WinsAncientTomb                                        int      `json:"wins_ancient_tomb"`
				WinsSnowfall                                           int      `json:"wins_snowfall"`
				GamesSnowfall                                          int      `json:"games_snowfall"`
				DeathsSnowfall                                         int      `json:"deaths_snowfall"`
				GamesSnowfallMURDERCLASSIC                             int      `json:"games_snowfall_MURDER_CLASSIC"`
				DeathsSnowfallMURDERCLASSIC                            int      `json:"deaths_snowfall_MURDER_CLASSIC"`
				WinsSnowfallMURDERCLASSIC                              int      `json:"wins_snowfall_MURDER_CLASSIC"`
				WinsTransport                                          int      `json:"wins_transport"`
				WinsTransportMURDERCLASSIC                             int      `json:"wins_transport_MURDER_CLASSIC"`
				CoinsPickedupMountainMURDERCLASSIC                     int      `json:"coins_pickedup_mountain_MURDER_CLASSIC"`
				WinsMountain                                           int      `json:"wins_mountain"`
				GamesMountainMURDERCLASSIC                             int      `json:"games_mountain_MURDER_CLASSIC"`
				WinsMountainMURDERCLASSIC                              int      `json:"wins_mountain_MURDER_CLASSIC"`
				GamesMountain                                          int      `json:"games_mountain"`
				DeathsMountain                                         int      `json:"deaths_mountain"`
				CoinsPickedupMountain                                  int      `json:"coins_pickedup_mountain"`
				DeathsMountainMURDERCLASSIC                            int      `json:"deaths_mountain_MURDER_CLASSIC"`
				WinsLibraryMURDERCLASSIC                               int      `json:"wins_library_MURDER_CLASSIC"`
				DeathsLibraryMURDERCLASSIC                             int      `json:"deaths_library_MURDER_CLASSIC"`
				WinsLibrary                                            int      `json:"wins_library"`
				DeathsLibrary                                          int      `json:"deaths_library"`
				GamesLibraryMURDERCLASSIC                              int      `json:"games_library_MURDER_CLASSIC"`
				GamesLibrary                                           int      `json:"games_library"`
				QuickestDetectiveWinTimeSecondsTransport               int      `json:"quickest_detective_win_time_seconds_transport"`
				QuickestDetectiveWinTimeSecondsTransportMURDERCLASSIC  int      `json:"quickest_detective_win_time_seconds_transport_MURDER_CLASSIC"`
				WasHeroTransport                                       int      `json:"was_hero_transport"`
				KillsTransportMURDERCLASSIC                            int      `json:"kills_transport_MURDER_CLASSIC"`
				WasHero                                                int      `json:"was_hero"`
				BowKillsMURDERCLASSIC                                  int      `json:"bow_kills_MURDER_CLASSIC"`
				WasHeroMURDERCLASSIC                                   int      `json:"was_hero_MURDER_CLASSIC"`
				KillsMURDERCLASSIC                                     int      `json:"kills_MURDER_CLASSIC"`
				KillsTransport                                         int      `json:"kills_transport"`
				DetectiveWinsTransport                                 int      `json:"detective_wins_transport"`
				DetectiveWinsTransportMURDERCLASSIC                    int      `json:"detective_wins_transport_MURDER_CLASSIC"`
				BowKillsTransportMURDERCLASSIC                         int      `json:"bow_kills_transport_MURDER_CLASSIC"`
				BowKills                                               int      `json:"bow_kills"`
				BowKillsTransport                                      int      `json:"bow_kills_transport"`
				WasHeroTransportMURDERCLASSIC                          int      `json:"was_hero_transport_MURDER_CLASSIC"`
				Kills                                                  int      `json:"kills"`
				CoinsPickedupAquariumMURDERCLASSIC                     int      `json:"coins_pickedup_aquarium_MURDER_CLASSIC"`
				GamesAquariumMURDERCLASSIC                             int      `json:"games_aquarium_MURDER_CLASSIC"`
				DeathsAquarium                                         int      `json:"deaths_aquarium"`
				CoinsPickedupAquarium                                  int      `json:"coins_pickedup_aquarium"`
				GamesAquarium                                          int      `json:"games_aquarium"`
				WinsAquarium                                           int      `json:"wins_aquarium"`
				WinsAquariumMURDERCLASSIC                              int      `json:"wins_aquarium_MURDER_CLASSIC"`
				DeathsAquariumMURDERCLASSIC                            int      `json:"deaths_aquarium_MURDER_CLASSIC"`
				Showqueuebook                                          bool     `json:"showqueuebook"`
				CoinsPickedupArchives                                  int      `json:"coins_pickedup_archives"`
				WinsArchivesMURDERCLASSIC                              int      `json:"wins_archives_MURDER_CLASSIC"`
				GamesArchivesMURDERCLASSIC                             int      `json:"games_archives_MURDER_CLASSIC"`
				CoinsPickedupArchivesMURDERCLASSIC                     int      `json:"coins_pickedup_archives_MURDER_CLASSIC"`
				GamesArchives                                          int      `json:"games_archives"`
				WinsArchives                                           int      `json:"wins_archives"`
				KnifeKillsMURDERCLASSIC                                int      `json:"knife_kills_MURDER_CLASSIC"`
				KnifeKillsAquariumMURDERCLASSIC                        int      `json:"knife_kills_aquarium_MURDER_CLASSIC"`
				KillsAsMurdererMURDERCLASSIC                           int      `json:"kills_as_murderer_MURDER_CLASSIC"`
				KnifeKillsAquarium                                     int      `json:"knife_kills_aquarium"`
				BowKillsAquariumMURDERCLASSIC                          int      `json:"bow_kills_aquarium_MURDER_CLASSIC"`
				KillsAquarium                                          int      `json:"kills_aquarium"`
				BowKillsAquarium                                       int      `json:"bow_kills_aquarium"`
				KillsAsMurdererAquarium                                int      `json:"kills_as_murderer_aquarium"`
				KillsAsMurderer                                        int      `json:"kills_as_murderer"`
				KillsAquariumMURDERCLASSIC                             int      `json:"kills_aquarium_MURDER_CLASSIC"`
				KnifeKills                                             int      `json:"knife_kills"`
				KillsAsMurdererAquariumMURDERCLASSIC                   int      `json:"kills_as_murderer_aquarium_MURDER_CLASSIC"`
				DeathsHeadquarters                                     int      `json:"deaths_headquarters"`
				GamesHeadquartersMURDERCLASSIC                         int      `json:"games_headquarters_MURDER_CLASSIC"`
				BowKillsHeadquartersMURDERCLASSIC                      int      `json:"bow_kills_headquarters_MURDER_CLASSIC"`
				WinsHeadquartersMURDERCLASSIC                          int      `json:"wins_headquarters_MURDER_CLASSIC"`
				DeathsHeadquartersMURDERCLASSIC                        int      `json:"deaths_headquarters_MURDER_CLASSIC"`
				KillsHeadquarters                                      int      `json:"kills_headquarters"`
				CoinsPickedupHeadquarters                              int      `json:"coins_pickedup_headquarters"`
				WinsHeadquarters                                       int      `json:"wins_headquarters"`
				GamesHeadquarters                                      int      `json:"games_headquarters"`
				BowKillsHeadquarters                                   int      `json:"bow_kills_headquarters"`
				KillsHeadquartersMURDERCLASSIC                         int      `json:"kills_headquarters_MURDER_CLASSIC"`
				CoinsPickedupHeadquartersMURDERCLASSIC                 int      `json:"coins_pickedup_headquarters_MURDER_CLASSIC"`
				WinsGoldRushMURDERCLASSIC                              int      `json:"wins_gold_rush_MURDER_CLASSIC"`
				GamesGoldRushMURDERCLASSIC                             int      `json:"games_gold_rush_MURDER_CLASSIC"`
				WinsGoldRush                                           int      `json:"wins_gold_rush"`
				CoinsPickedupGoldRush                                  int      `json:"coins_pickedup_gold_rush"`
				GamesGoldRush                                          int      `json:"games_gold_rush"`
				CoinsPickedupGoldRushMURDERCLASSIC                     int      `json:"coins_pickedup_gold_rush_MURDER_CLASSIC"`
				CoinsPickedupLibrary                                   int      `json:"coins_pickedup_library"`
				CoinsPickedupLibraryMURDERCLASSIC                      int      `json:"coins_pickedup_library_MURDER_CLASSIC"`
				GamesTowerfallMURDERCLASSIC                            int      `json:"games_towerfall_MURDER_CLASSIC"`
				CoinsPickedupTowerfall                                 int      `json:"coins_pickedup_towerfall"`
				GamesTowerfall                                         int      `json:"games_towerfall"`
				WinsTowerfallMURDERCLASSIC                             int      `json:"wins_towerfall_MURDER_CLASSIC"`
				WinsTowerfall                                          int      `json:"wins_towerfall"`
				CoinsPickedupTowerfallMURDERCLASSIC                    int      `json:"coins_pickedup_towerfall_MURDER_CLASSIC"`
				GamesDarkfallMURDERCLASSIC                             int      `json:"games_darkfall_MURDER_CLASSIC"`
				DeathsDarkfallMURDERCLASSIC                            int      `json:"deaths_darkfall_MURDER_CLASSIC"`
				GamesDarkfall                                          int      `json:"games_darkfall"`
				DeathsDarkfall                                         int      `json:"deaths_darkfall"`
				KnifeKillsDarkfall                                     int      `json:"knife_kills_darkfall"`
				KillsDarkfallMURDERCLASSIC                             int      `json:"kills_darkfall_MURDER_CLASSIC"`
				KillsAsMurdererDarkfall                                int      `json:"kills_as_murderer_darkfall"`
				KillsDarkfall                                          int      `json:"kills_darkfall"`
				KnifeKillsDarkfallMURDERCLASSIC                        int      `json:"knife_kills_darkfall_MURDER_CLASSIC"`
				CoinsPickedupDarkfallMURDERCLASSIC                     int      `json:"coins_pickedup_darkfall_MURDER_CLASSIC"`
				KillsAsMurdererDarkfallMURDERCLASSIC                   int      `json:"kills_as_murderer_darkfall_MURDER_CLASSIC"`
				CoinsPickedupDarkfall                                  int      `json:"coins_pickedup_darkfall"`
				CoinsPickedupSanPeratico                               int      `json:"coins_pickedup_san_peratico"`
				WinsSanPeratico                                        int      `json:"wins_san_peratico"`
				DeathsSanPeraticoMURDERCLASSIC                         int      `json:"deaths_san_peratico_MURDER_CLASSIC"`
				GamesSanPeraticoMURDERCLASSIC                          int      `json:"games_san_peratico_MURDER_CLASSIC"`
				WinsSanPeraticoMURDERCLASSIC                           int      `json:"wins_san_peratico_MURDER_CLASSIC"`
				CoinsPickedupSanPeraticoMURDERCLASSIC                  int      `json:"coins_pickedup_san_peratico_MURDER_CLASSIC"`
				GamesSanPeratico                                       int      `json:"games_san_peratico"`
				DeathsSanPeratico                                      int      `json:"deaths_san_peratico"`
				KillsAsMurdererLibrary                                 int      `json:"kills_as_murderer_library"`
				KillsAsMurdererLibraryMURDERCLASSIC                    int      `json:"kills_as_murderer_library_MURDER_CLASSIC"`
				GamesMURDERASSASSINS                                   int      `json:"games_MURDER_ASSASSINS"`
				DeathsMURDERASSASSINS                                  int      `json:"deaths_MURDER_ASSASSINS"`
				DeathsGoldRushMURDERASSASSINS                          int      `json:"deaths_gold_rush_MURDER_ASSASSINS"`
				DeathsGoldRush                                         int      `json:"deaths_gold_rush"`
				GamesGoldRushMURDERASSASSINS                           int      `json:"games_gold_rush_MURDER_ASSASSINS"`
			} `json:"MurderMystery"`
			HungerGames struct {
				Packages                []string `json:"packages"`
				Coins                   int      `json:"coins"`
				InGamePresentsCap201825 int      `json:"inGamePresentsCap_2018_25"`
				ChestsOpenedKnight      int      `json:"chests_opened_knight"`
				PotionsDrunk            int      `json:"potions_drunk"`
				Deaths                  int      `json:"deaths"`
				Damage                  int      `json:"damage"`
				TimePlayed              int      `json:"time_played"`
				TimePlayedKnight        int      `json:"time_played_knight"`
				DamageKnight            int      `json:"damage_knight"`
				ChestsOpened            int      `json:"chests_opened"`
				DamageTakenKnight       int      `json:"damage_taken_knight"`
				DamageTaken             int      `json:"damage_taken"`
				PotionsDrunkKnight      int      `json:"potions_drunk_knight"`
				GamesPlayedKnight       int      `json:"games_played_knight"`
				GamesPlayed             int      `json:"games_played"`
				DamageTakenScout        int      `json:"damage_taken_scout"`
				TimePlayedScout         int      `json:"time_played_scout"`
				ChestsOpenedScout       int      `json:"chests_opened_scout"`
				GamesPlayedScout        int      `json:"games_played_scout"`
				LastTourneyAd           int64    `json:"lastTourneyAd"`
				DamageScout             int      `json:"damage_scout"`
				ExpScout                int      `json:"exp_scout"`
				Kills                   int      `json:"kills"`
				KillsScout              int      `json:"kills_scout"`
				KillsSoloNormal         int      `json:"kills_solo_normal"`
				PotionsDrunkScout       int      `json:"potions_drunk_scout"`
				WinsSoloNormal          int      `json:"wins_solo_normal"`
				WinsBackup              int      `json:"wins_backup"`
				Autoarmor               bool     `json:"autoarmor"`
				WinsTeamsNormal         int      `json:"wins_teams_normal"`
				Wins                    int      `json:"wins"`
				ExpKnight               int      `json:"exp_knight"`
				KillsKnight             int      `json:"kills_knight"`
			} `json:"HungerGames"`
			Walls3 struct {
				Coins                                    int      `json:"coins"`
				Packages                                 []string `json:"packages"`
				ChosenClass                              string   `json:"chosen_class"`
				HerobrineATotalKillsStandard             int      `json:"herobrine_a_total_kills_standard"`
				HerobrineAssists                         int      `json:"herobrine_assists"`
				HerobrineAActivationsStandard            int      `json:"herobrine_a_activations_standard"`
				EnemiesHitStandard                       int      `json:"enemies_hit_standard"`
				HerobrineIronSwordCrafted                int      `json:"herobrine_iron_sword_crafted"`
				HerobrinePotionsDrunkStandard            int      `json:"herobrine_potions_drunk_standard"`
				HerobrineBlocksBrokenStandard            int      `json:"herobrine_blocks_broken_standard"`
				HerobrineMetersFallenStandard            int      `json:"herobrine_meters_fallen_standard"`
				HerobrineBreadEaten                      int      `json:"herobrine_bread_eaten"`
				HerobrineGamesPlayed                     int      `json:"herobrine_games_played"`
				HerobrineIronOreBroken                   int      `json:"herobrine_iron_ore_broken"`
				TotalDeaths                              int      `json:"total_deaths"`
				HerobrinePotionsDrunk                    int      `json:"herobrine_potions_drunk"`
				HerobrineBlocksBroken                    int      `json:"herobrine_blocks_broken"`
				HerobrineIronSwordCraftedStandard        int      `json:"herobrine_iron_sword_crafted_standard"`
				LossesStandard                           int      `json:"losses_standard"`
				TotalKills                               int      `json:"total_kills"`
				ActivationsStandard                      int      `json:"activations_standard"`
				TreasuresFound                           int      `json:"treasures_found"`
				HerobrineDeaths                          int      `json:"herobrine_deaths"`
				WoodChopped                              int      `json:"wood_chopped"`
				HerobrineDeathsStandard                  int      `json:"herobrine_deaths_standard"`
				HerobrineADamageDealt                    int      `json:"herobrine_a_damage_dealt"`
				HerobrineDamageDealtStandard             int      `json:"herobrine_damage_dealt_standard"`
				BreadEatenStandard                       int      `json:"bread_eaten_standard"`
				HerobrineBlocksPlaced                    int      `json:"herobrine_blocks_placed"`
				HerobrineBreadEatenStandard              int      `json:"herobrine_bread_eaten_standard"`
				TimePlayed                               int      `json:"time_played"`
				HerobrineBlocksPlacedPreparation         int      `json:"herobrine_blocks_placed_preparation"`
				HerobrineFoodEaten                       int      `json:"herobrine_food_eaten"`
				WoodChoppedStandard                      int      `json:"wood_chopped_standard"`
				HerobrineLossesStandard                  int      `json:"herobrine_losses_standard"`
				HerobrineATotalKills                     int      `json:"herobrine_a_total_kills"`
				HerobrineAEnemiesHitStandard             int      `json:"herobrine_a_enemies_hit_standard"`
				Assists                                  int      `json:"assists"`
				FoodEatenStandard                        int      `json:"food_eaten_standard"`
				EnemiesHit                               int      `json:"enemies_hit"`
				GamesPlayed                              int      `json:"games_played"`
				BlocksPlacedStandard                     int      `json:"blocks_placed_standard"`
				HerobrineTotalDeaths                     int      `json:"herobrine_total_deaths"`
				BlocksPlacedPreparation                  int      `json:"blocks_placed_preparation"`
				BlocksPlacedPreparationStandard          int      `json:"blocks_placed_preparation_standard"`
				HerobrineBlocksPlacedStandard            int      `json:"herobrine_blocks_placed_standard"`
				DamageDealt                              int      `json:"damage_dealt"`
				HerobrineActivations                     int      `json:"herobrine_activations"`
				TotalKillsStandard                       int      `json:"total_kills_standard"`
				TimePlayedStandard                       int      `json:"time_played_standard"`
				Deaths                                   int      `json:"deaths"`
				HerobrineAAssists                        int      `json:"herobrine_a_assists"`
				HerobrineWoodChopped                     int      `json:"herobrine_wood_chopped"`
				HerobrineTreasuresFoundStandard          int      `json:"herobrine_treasures_found_standard"`
				HerobrineIronOreBrokenStandard           int      `json:"herobrine_iron_ore_broken_standard"`
				HerobrineEnemiesHitStandard              int      `json:"herobrine_enemies_hit_standard"`
				IronOreBrokenStandard                    int      `json:"iron_ore_broken_standard"`
				TreasuresFoundStandard                   int      `json:"treasures_found_standard"`
				IronSwordCraftedStandard                 int      `json:"iron_sword_crafted_standard"`
				HerobrineEnemiesHit                      int      `json:"herobrine_enemies_hit"`
				FoodEaten                                int      `json:"food_eaten"`
				BlocksBrokenStandard                     int      `json:"blocks_broken_standard"`
				PotionsDrunkStandard                     int      `json:"potions_drunk_standard"`
				HerobrineMetersWalkedSpeed               int      `json:"herobrine_meters_walked_speed"`
				HerobrineSwordCrafted                    int      `json:"herobrine_sword_crafted"`
				Losses                                   int      `json:"losses"`
				DamageDealtStandard                      int      `json:"damage_dealt_standard"`
				MetersWalkedSpeedStandard                int      `json:"meters_walked_speed_standard"`
				HerobrineWoodChoppedStandard             int      `json:"herobrine_wood_chopped_standard"`
				HerobrineAAssistsStandard                int      `json:"herobrine_a_assists_standard"`
				HerobrineMetersWalked                    int      `json:"herobrine_meters_walked"`
				HerobrineTreasuresFound                  int      `json:"herobrine_treasures_found"`
				SwordCrafted                             int      `json:"sword_crafted"`
				HerobrineDamageDealt                     int      `json:"herobrine_damage_dealt"`
				HerobrineBlocksPlacedPreparationStandard int      `json:"herobrine_blocks_placed_preparation_standard"`
				HerobrineTimePlayed                      int      `json:"herobrine_time_played"`
				MetersWalked                             int      `json:"meters_walked"`
				HerobrineActivationsStandard             int      `json:"herobrine_activations_standard"`
				MetersFallenStandard                     int      `json:"meters_fallen_standard"`
				HerobrineMetersWalkedSpeedStandard       int      `json:"herobrine_meters_walked_speed_standard"`
				PotionsDrunk                             int      `json:"potions_drunk"`
				IronSwordCrafted                         int      `json:"iron_sword_crafted"`
				BlocksBroken                             int      `json:"blocks_broken"`
				Activations                              int      `json:"activations"`
				AssistsStandard                          int      `json:"assists_standard"`
				BlocksPlaced                             int      `json:"blocks_placed"`
				HerobrineTotalDeathsStandard             int      `json:"herobrine_total_deaths_standard"`
				MetersWalkedSpeed                        int      `json:"meters_walked_speed"`
				BreadEaten                               int      `json:"bread_eaten"`
				IronOreBroken                            int      `json:"iron_ore_broken"`
				HerobrineFoodEatenStandard               int      `json:"herobrine_food_eaten_standard"`
				HerobrineLosses                          int      `json:"herobrine_losses"`
				DeathsStandard                           int      `json:"deaths_standard"`
				TotalDeathsStandard                      int      `json:"total_deaths_standard"`
				HerobrineTotalKillsStandard              int      `json:"herobrine_total_kills_standard"`
				SwordCraftedStandard                     int      `json:"sword_crafted_standard"`
				HerobrineGamesPlayedStandard             int      `json:"herobrine_games_played_standard"`
				HerobrineSwordCraftedStandard            int      `json:"herobrine_sword_crafted_standard"`
				HerobrineMetersFallen                    int      `json:"herobrine_meters_fallen"`
				HerobrineAEnemiesHit                     int      `json:"herobrine_a_enemies_hit"`
				HerobrineMetersWalkedStandard            int      `json:"herobrine_meters_walked_standard"`
				MetersFallen                             int      `json:"meters_fallen"`
				HerobrineADamageDealtStandard            int      `json:"herobrine_a_damage_dealt_standard"`
				HerobrineAssistsStandard                 int      `json:"herobrine_assists_standard"`
				HerobrineTimePlayedStandard              int      `json:"herobrine_time_played_standard"`
				MetersWalkedStandard                     int      `json:"meters_walked_standard"`
				GamesPlayedStandard                      int      `json:"games_played_standard"`
				HerobrineAActivations                    int      `json:"herobrine_a_activations"`
				HerobrineTotalKills                      int      `json:"herobrine_total_kills"`
			} `json:"Walls3"`
			Battleground struct {
				ChosenClass   string   `json:"chosen_class"`
				ShamanSpec    string   `json:"shaman_spec"`
				MageSpec      string   `json:"mage_spec"`
				WarriorSpec   string   `json:"warrior_spec"`
				Packages      []string `json:"packages"`
				PaladinSpec   string   `json:"paladin_spec"`
				SelectedMount string   `json:"selected_mount"`
				Coins         int      `json:"coins"`
			} `json:"Battleground"`
			Duels struct {
				ShowLbOption                   string   `json:"show_lb_option"`
				GamesPlayedDuels               int      `json:"games_played_duels"`
				ChatEnabled                    string   `json:"chat_enabled"`
				MeleeSwings                    int      `json:"melee_swings"`
				UhcDuelHealthRegenerated       int      `json:"uhc_duel_health_regenerated"`
				MeleeHits                      int      `json:"melee_hits"`
				UhcDuelBowShots                int      `json:"uhc_duel_bow_shots"`
				RoundsPlayed                   int      `json:"rounds_played"`
				UhcDuelMeleeSwings             int      `json:"uhc_duel_melee_swings"`
				BowShots                       int      `json:"bow_shots"`
				GoldenApplesEaten              int      `json:"golden_apples_eaten"`
				UhcDuelRoundsPlayed            int      `json:"uhc_duel_rounds_played"`
				UhcDuelDamageDealt             int      `json:"uhc_duel_damage_dealt"`
				HealthRegenerated              int      `json:"health_regenerated"`
				DamageDealt                    int      `json:"damage_dealt"`
				UhcDuelGoldenApplesEaten       int      `json:"uhc_duel_golden_apples_eaten"`
				UhcDuelMeleeHits               int      `json:"uhc_duel_melee_hits"`
				BlocksPlaced                   int      `json:"blocks_placed"`
				UhcDuelBlocksPlaced            int      `json:"uhc_duel_blocks_placed"`
				Packages                       []string `json:"packages"`
				BridgeDuelHealthRegenerated    int      `json:"bridge_duel_health_regenerated"`
				BridgeDuelDamageDealt          int      `json:"bridge_duel_damage_dealt"`
				BowHits                        int      `json:"bow_hits"`
				BridgeDuelMeleeHits            int      `json:"bridge_duel_melee_hits"`
				BridgeDuelBlocksPlaced         int      `json:"bridge_duel_blocks_placed"`
				BridgeDuelMeleeSwings          int      `json:"bridge_duel_melee_swings"`
				BridgeDuelRoundsPlayed         int      `json:"bridge_duel_rounds_played"`
				BridgeDuelBowHits              int      `json:"bridge_duel_bow_hits"`
				BridgeDuelBowShots             int      `json:"bridge_duel_bow_shots"`
				SkywarsRookieTitlePrestige     int      `json:"skywars_rookie_title_prestige"`
				AllModesRookieTitlePrestige    int      `json:"all_modes_rookie_title_prestige"`
				ComboRookieTitlePrestige       int      `json:"combo_rookie_title_prestige"`
				NoDebuffRookieTitlePrestige    int      `json:"no_debuff_rookie_title_prestige"`
				TournamentRookieTitlePrestige  int      `json:"tournament_rookie_title_prestige"`
				MegaWallsRookieTitlePrestige   int      `json:"mega_walls_rookie_title_prestige"`
				OpRookieTitlePrestige          int      `json:"op_rookie_title_prestige"`
				UhcRookieTitlePrestige         int      `json:"uhc_rookie_title_prestige"`
				TntGamesRookieTitlePrestige    int      `json:"tnt_games_rookie_title_prestige"`
				BlitzRookieTitlePrestige       int      `json:"blitz_rookie_title_prestige"`
				BridgeRookieTitlePrestige      int      `json:"bridge_rookie_title_prestige"`
				BowRookieTitlePrestige         int      `json:"bow_rookie_title_prestige"`
				ClassicRookieTitlePrestige     int      `json:"classic_rookie_title_prestige"`
				SumoRookieTitlePrestige        int      `json:"sumo_rookie_title_prestige"`
				Selected1New                   string   `json:"selected_1_new"`
				Selected2New                   string   `json:"selected_2_new"`
				DuelsRecentlyPlayed            string   `json:"duels_recently_played"`
				BestWinstreakModeComboDuel     int      `json:"best_winstreak_mode_combo_duel"`
				BestComboWinstreak             int      `json:"best_combo_winstreak"`
				MapsWonOn                      []string `json:"maps_won_on"`
				CurrentWinstreak               int      `json:"current_winstreak"`
				CurrentWinstreakModeComboDuel  int      `json:"current_winstreak_mode_combo_duel"`
				CurrentComboWinstreak          int      `json:"current_combo_winstreak"`
				BestOverallWinstreak           int      `json:"best_overall_winstreak"`
				ComboDuelMeleeHits             int      `json:"combo_duel_melee_hits"`
				ComboDuelWins                  int      `json:"combo_duel_wins"`
				Wins                           int      `json:"wins"`
				ComboDuelRoundsPlayed          int      `json:"combo_duel_rounds_played"`
				ComboDuelHealthRegenerated     int      `json:"combo_duel_health_regenerated"`
				Coins                          int      `json:"coins"`
				ComboDuelMeleeSwings           int      `json:"combo_duel_melee_swings"`
				ComboDuelGoldenApplesEaten     int      `json:"combo_duel_golden_apples_eaten"`
				KitMenuOption                  string   `json:"kit_menu_option"`
				UhcTournamentGoldenApplesEaten int      `json:"uhc_tournament_golden_apples_eaten"`
				UhcTournamentLosses            int      `json:"uhc_tournament_losses"`
				UhcTournamentMeleeHits         int      `json:"uhc_tournament_melee_hits"`
				UhcTournamentHealthRegenerated int      `json:"uhc_tournament_health_regenerated"`
				UhcTournamentDeaths            int      `json:"uhc_tournament_deaths"`
				UhcTournamentKills             int      `json:"uhc_tournament_kills"`
				Kills                          int      `json:"kills"`
				UhcTournamentMeleeSwings       int      `json:"uhc_tournament_melee_swings"`
				Losses                         int      `json:"losses"`
				UhcTournamentDamageDealt       int      `json:"uhc_tournament_damage_dealt"`
				UhcTournamentBlocksPlaced      int      `json:"uhc_tournament_blocks_placed"`
				Deaths                         int      `json:"deaths"`
				UhcTournamentRoundsPlayed      int      `json:"uhc_tournament_rounds_played"`
				UhcTournamentBowHits           int      `json:"uhc_tournament_bow_hits"`
				UhcTournamentBowShots          int      `json:"uhc_tournament_bow_shots"`
				ClassicDuelRoundsPlayed        int      `json:"classic_duel_rounds_played"`
				ClassicDuelMeleeHits           int      `json:"classic_duel_melee_hits"`
				ClassicDuelHealthRegenerated   int      `json:"classic_duel_health_regenerated"`
				ClassicDuelDamageDealt         int      `json:"classic_duel_damage_dealt"`
				ClassicDuelBowShots            int      `json:"classic_duel_bow_shots"`
				ClassicDuelMeleeSwings         int      `json:"classic_duel_melee_swings"`
				BridgeDoublesRoundsPlayed      int      `json:"bridge_doubles_rounds_played"`
				BridgeDoublesBlocksPlaced      int      `json:"bridge_doubles_blocks_placed"`
				OpDuelHealthRegenerated        int      `json:"op_duel_health_regenerated"`
				OpDuelDamageDealt              int      `json:"op_duel_damage_dealt"`
				OpDuelRoundsPlayed             int      `json:"op_duel_rounds_played"`
				OpDuelMeleeSwings              int      `json:"op_duel_melee_swings"`
				OpDuelMeleeHits                int      `json:"op_duel_melee_hits"`
				CurrentSumoWinstreak           int      `json:"current_sumo_winstreak"`
				CurrentWinstreakModeSumoDuel   int      `json:"current_winstreak_mode_sumo_duel"`
				BestWinstreakModeSumoDuel      int      `json:"best_winstreak_mode_sumo_duel"`
				BestSumoWinstreak              int      `json:"best_sumo_winstreak"`
				SumoDuelKills                  int      `json:"sumo_duel_kills"`
				SumoDuelMeleeSwings            int      `json:"sumo_duel_melee_swings"`
				SumoDuelWins                   int      `json:"sumo_duel_wins"`
				SumoDuelRoundsPlayed           int      `json:"sumo_duel_rounds_played"`
				SumoDuelMeleeHits              int      `json:"sumo_duel_melee_hits"`
				SumoDuelLosses                 int      `json:"sumo_duel_losses"`
				SumoDuelDeaths                 int      `json:"sumo_duel_deaths"`
				DuelsWinstreakBestSumoDuel     int      `json:"duels_winstreak_best_sumo_duel"`
				LeaderboardPageWinStreak       int      `json:"leaderboardPage_win_streak"`
				SumoIronTitlePrestige          int      `json:"sumo_iron_title_prestige"`
				ComboDuelLosses                int      `json:"combo_duel_losses"`
				ComboDuelDeaths                int      `json:"combo_duel_deaths"`
				UhcDuelBowHits                 int      `json:"uhc_duel_bow_hits"`
				ComboDuelKills                 int      `json:"combo_duel_kills"`
				AllModesIronTitlePrestige      int      `json:"all_modes_iron_title_prestige"`
				SumoGoldTitlePrestige          int      `json:"sumo_gold_title_prestige"`
				DuelsChests                    int      `json:"duels_chests"`
				DuelsOpenedChests              int      `json:"Duels_openedChests"`
				DuelsChestHistory              []string `json:"duels_chest_history"`
				DuelsOpenedRares               int      `json:"Duels_openedRares"`
				DuelsOpenedCommons             int      `json:"Duels_openedCommons"`
				ActiveCosmetictitle            string   `json:"active_cosmetictitle"`
				SwDuelsKitNew3                 string   `json:"sw_duels_kit_new3"`
				CurrentWinstreakModeSwDoubles  int      `json:"current_winstreak_mode_sw_doubles"`
				CurrentSkywarsWinstreak        int      `json:"current_skywars_winstreak"`
				SwDoublesMeleeSwings           int      `json:"sw_doubles_melee_swings"`
				SwDoublesHealthRegenerated     int      `json:"sw_doubles_health_regenerated"`
				SwDoublesDeaths                int      `json:"sw_doubles_deaths"`
				SwDoublesKills                 int      `json:"sw_doubles_kills"`
				SwDoublesMeleeHits             int      `json:"sw_doubles_melee_hits"`
				SwDoublesRoundsPlayed          int      `json:"sw_doubles_rounds_played"`
				SwDoublesLosses                int      `json:"sw_doubles_losses"`
				SwDoublesDamageDealt           int      `json:"sw_doubles_damage_dealt"`
				SwDoublesBlocksPlaced          int      `json:"sw_doubles_blocks_placed"`
				BestWinstreakModeSwDoubles     int      `json:"best_winstreak_mode_sw_doubles"`
				BestSkywarsWinstreak           int      `json:"best_skywars_winstreak"`
				SwDoublesKitWins               int      `json:"sw_doubles_kit_wins"`
				SwDoublesPyromancerKitWins     int      `json:"sw_doubles_pyromancer_kit_wins"`
				KitWins                        int      `json:"kit_wins"`
				PyromancerKitWins              int      `json:"pyromancer_kit_wins"`
				SwDoublesWins                  int      `json:"sw_doubles_wins"`
				DuelsWinstreakBestSwDoubles    int      `json:"duels_winstreak_best_sw_doubles"`
				SwDoublesBowShots              int      `json:"sw_doubles_bow_shots"`
				SwDoublesBowHits               int      `json:"sw_doubles_bow_hits"`
				ActiveAuras                    string   `json:"active_auras"`
				PotionDuelHealPotsUsed         int      `json:"potion_duel_heal_pots_used"`
				PotionDuelMeleeSwings          int      `json:"potion_duel_melee_swings"`
				PotionDuelRoundsPlayed         int      `json:"potion_duel_rounds_played"`
				PotionDuelMeleeHits            int      `json:"potion_duel_melee_hits"`
				PotionDuelHealthRegenerated    int      `json:"potion_duel_health_regenerated"`
				PotionDuelDamageDealt          int      `json:"potion_duel_damage_dealt"`
				HealPotsUsed                   int      `json:"heal_pots_used"`
				SwDuelRoundsPlayed             int      `json:"sw_duel_rounds_played"`
				SwDuelMeleeSwings              int      `json:"sw_duel_melee_swings"`
				SwDuelDamageDealt              int      `json:"sw_duel_damage_dealt"`
				SwDuelHealthRegenerated        int      `json:"sw_duel_health_regenerated"`
				SwDuelMeleeHits                int      `json:"sw_duel_melee_hits"`
				BridgeDoublesBowShots          int      `json:"bridge_doubles_bow_shots"`
				BridgeDoublesDamageDealt       int      `json:"bridge_doubles_damage_dealt"`
				BridgeDoublesHealthRegenerated int      `json:"bridge_doubles_health_regenerated"`
				BridgeDoublesMeleeHits         int      `json:"bridge_doubles_melee_hits"`
				BridgeDoublesMeleeSwings       int      `json:"bridge_doubles_melee_swings"`
				BridgeDoublesBowHits           int      `json:"bridge_doubles_bow_hits"`
				DuelsRecentlyPlayed2           string   `json:"duels_recently_played2"`
				BestWinstreakModeBowDuel       int      `json:"best_winstreak_mode_bow_duel"`
				CurrentWinstreakModeBowDuel    int      `json:"current_winstreak_mode_bow_duel"`
				BestBowWinstreak               int      `json:"best_bow_winstreak"`
				CurrentBowWinstreak            int      `json:"current_bow_winstreak"`
				BowDuelBowHits                 int      `json:"bow_duel_bow_hits"`
				BowDuelBowShots                int      `json:"bow_duel_bow_shots"`
				BowDuelDamageDealt             int      `json:"bow_duel_damage_dealt"`
				BowDuelKills                   int      `json:"bow_duel_kills"`
				BowDuelRoundsPlayed            int      `json:"bow_duel_rounds_played"`
				BowDuelWins                    int      `json:"bow_duel_wins"`
				DuelsWinstreakBestBowDuel      int      `json:"duels_winstreak_best_bow_duel"`
				BowDuelHealthRegenerated       int      `json:"bow_duel_health_regenerated"`
				ActiveWeaponpacks              string   `json:"active_weaponpacks"`
				LayoutBowDuelLayout            struct {
					Num0 string `json:"0"`
					Num3 string `json:"3"`
					Num8 string `json:"8"`
				} `json:"layout_bow_duel_layout"`
				StatusField                       string   `json:"status_field"`
				ActiveProjectileTrail             string   `json:"active_projectile_trail"`
				BowDuelDeaths                     int      `json:"bow_duel_deaths"`
				BowDuelLosses                     int      `json:"bow_duel_losses"`
				CurrentWinstreakModeUhcDuel       int      `json:"current_winstreak_mode_uhc_duel"`
				CurrentUhcWinstreak               int      `json:"current_uhc_winstreak"`
				BestWinstreakModeUhcDuel          int      `json:"best_winstreak_mode_uhc_duel"`
				BestUhcWinstreak                  int      `json:"best_uhc_winstreak"`
				UhcDuelKills                      int      `json:"uhc_duel_kills"`
				UhcDuelWins                       int      `json:"uhc_duel_wins"`
				DuelsWinstreakBestUhcDuel         int      `json:"duels_winstreak_best_uhc_duel"`
				UhcDuelDeaths                     int      `json:"uhc_duel_deaths"`
				UhcDuelLosses                     int      `json:"uhc_duel_losses"`
				BestWinstreakModeUhcDoubles       int      `json:"best_winstreak_mode_uhc_doubles"`
				CurrentWinstreakModeUhcDoubles    int      `json:"current_winstreak_mode_uhc_doubles"`
				UhcDoublesDamageDealt             int      `json:"uhc_doubles_damage_dealt"`
				UhcDoublesGoldenApplesEaten       int      `json:"uhc_doubles_golden_apples_eaten"`
				UhcDoublesHealthRegenerated       int      `json:"uhc_doubles_health_regenerated"`
				UhcDoublesKills                   int      `json:"uhc_doubles_kills"`
				UhcDoublesMeleeHits               int      `json:"uhc_doubles_melee_hits"`
				UhcDoublesMeleeSwings             int      `json:"uhc_doubles_melee_swings"`
				UhcDoublesRoundsPlayed            int      `json:"uhc_doubles_rounds_played"`
				UhcDoublesWins                    int      `json:"uhc_doubles_wins"`
				UhcDoublesBlocksPlaced            int      `json:"uhc_doubles_blocks_placed"`
				UhcDoublesDeaths                  int      `json:"uhc_doubles_deaths"`
				UhcDoublesLosses                  int      `json:"uhc_doubles_losses"`
				UhcDoublesBowHits                 int      `json:"uhc_doubles_bow_hits"`
				UhcDoublesBowShots                int      `json:"uhc_doubles_bow_shots"`
				BestWinstreakModeBridgeDoubles    int      `json:"best_winstreak_mode_bridge_doubles"`
				BridgeMapWins                     []string `json:"bridgeMapWins"`
				CurrentBridgeWinstreak            int      `json:"current_bridge_winstreak"`
				BestBridgeWinstreak               int      `json:"best_bridge_winstreak"`
				CurrentWinstreakModeBridgeDoubles int      `json:"current_winstreak_mode_bridge_doubles"`
				BridgeDeaths                      int      `json:"bridge_deaths"`
				BridgeDoublesBridgeDeaths         int      `json:"bridge_doubles_bridge_deaths"`
				BridgeDoublesBridgeKills          int      `json:"bridge_doubles_bridge_kills"`
				BridgeDoublesGoals                int      `json:"bridge_doubles_goals"`
				BridgeDoublesWins                 int      `json:"bridge_doubles_wins"`
				BridgeKills                       int      `json:"bridge_kills"`
				Goals                             int      `json:"goals"`
				ClassicDuelBowHits                int      `json:"classic_duel_bow_hits"`
				CurrentWinstreakModeClassicDuel   int      `json:"current_winstreak_mode_classic_duel"`
				BestClassicWinstreak              int      `json:"best_classic_winstreak"`
				CurrentClassicWinstreak           int      `json:"current_classic_winstreak"`
				BestWinstreakModeClassicDuel      int      `json:"best_winstreak_mode_classic_duel"`
				ClassicDuelKills                  int      `json:"classic_duel_kills"`
				ClassicDuelWins                   int      `json:"classic_duel_wins"`
			} `json:"Duels"`
			Mcgo struct {
				Coins              int      `json:"coins"`
				PocketChange       int      `json:"pocket_change"`
				GameWinsDeathmatch int      `json:"game_wins_deathmatch"`
				BombsDefused       int      `json:"bombs_defused"`
				GrenadeKills       int      `json:"grenadeKills"`
				GrenadeKills2      int      `json:"grenade_kills"`
				Packages           []string `json:"packages"`
				KillsDeathmatch    int      `json:"kills_deathmatch"`
				GameWins           int      `json:"game_wins"`
				HeadshotKills      int      `json:"headshot_kills"`
				BombsPlanted       int      `json:"bombs_planted"`
				Kills              int      `json:"kills"`
			} `json:"MCGO"`
			SkyClash struct {
				Coins int `json:"coins"`
			} `json:"SkyClash"`
			Uhc struct {
				Coins                 int      `json:"coins"`
				SavedStats            bool     `json:"saved_stats"`
				ClearupAchievement    bool     `json:"clearup_achievement"`
				EquippedKit           string   `json:"equippedKit"`
				Kills                 int      `json:"kills"`
				Score                 int      `json:"score"`
				Deaths                int      `json:"deaths"`
				HeadsEaten            int      `json:"heads_eaten"`
				KitWORKINGTOOLS       int      `json:"kit_WORKING_TOOLS"`
				Packages              []string `json:"packages"`
				PerkToolsmithingLineA int      `json:"perk_toolsmithing_line_a"`
				PerkToolsmithingLineC int      `json:"perk_toolsmithing_line_c"`
				PerkToolsmithingLineB int      `json:"perk_toolsmithing_line_b"`
				DeathsSolo            int      `json:"deaths_solo"`
				UltimatesCraftedSolo  int      `json:"ultimates_crafted_solo"`
				PerkWeaponsmithLineA  int      `json:"perk_weaponsmith_line_a"`
				PerkWeaponsmithLineB  int      `json:"perk_weaponsmith_line_b"`
				PerkApprenticeLineA   int      `json:"perk_apprentice_line_a"`
				PerkApprenticeLineB   int      `json:"perk_apprentice_line_b"`
				PerkApprenticeLineC   int      `json:"perk_apprentice_line_c"`
				PerkEngineeringLineA  int      `json:"perk_engineering_line_a"`
				PerkEngineeringLineC  int      `json:"perk_engineering_line_c"`
				UltimatesCrafted      int      `json:"ultimates_crafted"`
				KillsSolo             int      `json:"kills_solo"`
				HeadsEatenSolo        int      `json:"heads_eaten_solo"`
				PerkCookingLineA      int      `json:"perk_cooking_line_a"`
				PerkCookingLineB      int      `json:"perk_cooking_line_b"`
			} `json:"UHC"`
			SuperSmash struct {
				Coins int `json:"coins"`
			} `json:"SuperSmash"`
			Walls struct {
				Coins int `json:"coins"`
			} `json:"Walls"`
			SpeedUHC struct {
				Coins                                int  `json:"coins"`
				Score                                int  `json:"score"`
				MovedOver                            bool `json:"movedOver"`
				HighestKillstreak                    int  `json:"highestKillstreak"`
				Killstreak                           int  `json:"killstreak"`
				Assists                              int  `json:"assists"`
				AssistsKitBasicNormalDefault         int  `json:"assists_kit_basic_normal_default"`
				AssistsNormal                        int  `json:"assists_normal"`
				AssistsTeam                          int  `json:"assists_team"`
				BlocksBroken                         int  `json:"blocks_broken"`
				BlocksPlaced                         int  `json:"blocks_placed"`
				Deaths                               int  `json:"deaths"`
				DeathsKitBasicNormalDefault          int  `json:"deaths_kit_basic_normal_default"`
				DeathsMasteryWildSpecialist          int  `json:"deaths_mastery_wild_specialist"`
				DeathsNormal                         int  `json:"deaths_normal"`
				DeathsTeam                           int  `json:"deaths_team"`
				DeathsTeamNormal                     int  `json:"deaths_team_normal"`
				Kills                                int  `json:"kills"`
				KillsKitBasicNormalDefault           int  `json:"kills_kit_basic_normal_default"`
				KillsMasteryWildSpecialist           int  `json:"kills_mastery_wild_specialist"`
				KillsMonthlyB                        int  `json:"kills_monthly_b"`
				KillsNormal                          int  `json:"kills_normal"`
				KillsTeam                            int  `json:"kills_team"`
				KillsTeamNormal                      int  `json:"kills_team_normal"`
				KillsWeeklyA                         int  `json:"kills_weekly_a"`
				Losses                               int  `json:"losses"`
				LossesKitBasicNormalDefault          int  `json:"losses_kit_basic_normal_default"`
				LossesMasteryWildSpecialist          int  `json:"losses_mastery_wild_specialist"`
				LossesNormal                         int  `json:"losses_normal"`
				LossesTeam                           int  `json:"losses_team"`
				LossesTeamNormal                     int  `json:"losses_team_normal"`
				Quits                                int  `json:"quits"`
				ScoreKitBasicNormalDefault           int  `json:"score_kit_basic_normal_default"`
				ScoreNormal                          int  `json:"score_normal"`
				ScoreTeam                            int  `json:"score_team"`
				SurvivedPlayers                      int  `json:"survived_players"`
				SurvivedPlayersKitBasicNormalDefault int  `json:"survived_players_kit_basic_normal_default"`
				SurvivedPlayersNormal                int  `json:"survived_players_normal"`
				SurvivedPlayersTeam                  int  `json:"survived_players_team"`
				WinStreak                            int  `json:"win_streak"`
				Games                                int  `json:"games"`
				GamesKitBasicNormalDefault           int  `json:"games_kit_basic_normal_default"`
				GamesNormal                          int  `json:"games_normal"`
				GamesTeam                            int  `json:"games_team"`
				ItemsEnchanted                       int  `json:"items_enchanted"`
				KillsMonthlyA                        int  `json:"kills_monthly_a"`
			} `json:"SpeedUHC"`
			TrueCombat struct {
				Coins    int      `json:"coins"`
				Packages []string `json:"packages"`
			} `json:"TrueCombat"`
			BuildBattle struct {
				GamesPlayed    int `json:"games_played"`
				Score          int `json:"score"`
				Coins          int `json:"coins"`
				MonthlyCoinsA  int `json:"monthly_coins_a"`
				WeeklyCoinsB   int `json:"weekly_coins_b"`
				CorrectGuesses int `json:"correct_guesses"`
				MonthlyCoinsB  int `json:"monthly_coins_b"`
				SoloMostPoints int `json:"solo_most_points"`
				TotalVotes     int `json:"total_votes"`
				WeeklyCoinsA   int `json:"weekly_coins_a"`
			} `json:"BuildBattle"`
			Pit struct {
				Profile struct {
					Renown          int           `json:"renown"`
					OutgoingOffers  []interface{} `json:"outgoing_offers"`
					ContractChoices interface{}   `json:"contract_choices"`
					LastSave        int64         `json:"last_save"`
					KingQuest       struct {
						Kills int `json:"kills"`
					} `json:"king_quest"`
					Prestiges []struct {
						Index        int   `json:"index"`
						XpOnPrestige int   `json:"xp_on_prestige"`
						Timestamp    int64 `json:"timestamp"`
					} `json:"prestiges"`
					ZeroPointThreeGoldTransfer bool `json:"zero_point_three_gold_transfer"`
					Unlocks1                   []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks_1"`
					RenownUnlocks []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"renown_unlocks"`
					InvEnderchest struct {
						Type int   `json:"type"`
						Data []int `json:"data"`
					} `json:"inv_enderchest"`
					DeathRecaps struct {
						Type int   `json:"type"`
						Data []int `json:"data"`
					} `json:"death_recaps"`
					Unlocks2 []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks_2"`
					Cash                   float64 `json:"cash"`
					LastMidfightDisconnect int64   `json:"last_midfight_disconnect"`
					LeaderboardStats       struct {
					} `json:"leaderboard_stats"`
					SelectedPerk3 interface{} `json:"selected_perk_3"`
					SelectedPerk2 interface{} `json:"selected_perk_2"`
					InvArmor      struct {
						Type int   `json:"type"`
						Data []int `json:"data"`
					} `json:"inv_armor"`
					SelectedPerk1       string        `json:"selected_perk_1"`
					SelectedPerk0       string        `json:"selected_perk_0"`
					LastContract        int64         `json:"last_contract"`
					CashDuringPrestige2 float64       `json:"cash_during_prestige_2"`
					LoginMessages       []interface{} `json:"login_messages"`
					HotbarFavorites     []int         `json:"hotbar_favorites"`
					Xp                  int           `json:"xp"`
					InvContents         struct {
						Type int   `json:"type"`
						Data []int `json:"data"`
					} `json:"inv_contents"`
					EndedContracts []struct {
						Difficulty   string `json:"difficulty"`
						GoldReward   int    `json:"gold_reward"`
						Requirements struct {
							Kills int `json:"kills"`
						} `json:"requirements"`
						Progress struct {
							Kills int `json:"kills"`
						} `json:"progress"`
						CompletionDate int64  `json:"completion_date"`
						RemainingTicks int    `json:"remaining_ticks"`
						Key            string `json:"key"`
					} `json:"ended_contracts"`
					Bounties []interface{} `json:"bounties"`
					Unlocks  []struct {
						Tier        int    `json:"tier"`
						AcquireDate int64  `json:"acquireDate"`
						Key         string `json:"key"`
					} `json:"unlocks"`
					CashDuringPrestige1 float64 `json:"cash_during_prestige_1"`
					CashDuringPrestige0 float64 `json:"cash_during_prestige_0"`
				} `json:"profile"`
				PitStatsPtl struct {
					ArrowHits           int `json:"arrow_hits"`
					ArrowsFired         int `json:"arrows_fired"`
					Assists             int `json:"assists"`
					BowDamageDealt      int `json:"bow_damage_dealt"`
					CashEarned          int `json:"cash_earned"`
					DamageDealt         int `json:"damage_dealt"`
					DamageReceived      int `json:"damage_received"`
					Deaths              int `json:"deaths"`
					EnderchestOpened    int `json:"enderchest_opened"`
					FishingRodLaunched  int `json:"fishing_rod_launched"`
					GheadEaten          int `json:"ghead_eaten"`
					Joins               int `json:"joins"`
					JumpedIntoPit       int `json:"jumped_into_pit"`
					Kills               int `json:"kills"`
					LeftClicks          int `json:"left_clicks"`
					MeleeDamageDealt    int `json:"melee_damage_dealt"`
					MeleeDamageReceived int `json:"melee_damage_received"`
					PlaytimeMinutes     int `json:"playtime_minutes"`
					SwordHits           int `json:"sword_hits"`
					BowDamageReceived   int `json:"bow_damage_received"`
				} `json:"pit_stats_ptl"`
			} `json:"Pit"`
		} `json:"stats"`
		Achievements struct {
			SkywarsCages                    int `json:"skywars_cages"`
			BedwarsLevel                    int `json:"bedwars_level"`
			GeneralCoins                    int `json:"general_coins"`
			TntgamesTntBanker               int `json:"tntgames_tnt_banker"`
			TntgamesClinic                  int `json:"tntgames_clinic"`
			GeneralChallenger               int `json:"general_challenger"`
			TntgamesTntTriathlon            int `json:"tntgames_tnt_triathlon"`
			BedwarsBeds                     int `json:"bedwars_beds"`
			BedwarsCollectorsEdition        int `json:"bedwars_collectors_edition"`
			BedwarsBedwarsKiller            int `json:"bedwars_bedwars_killer"`
			ArenaClimbTheRanks              int `json:"arena_climb_the_ranks"`
			SkywarsKillsSolo                int `json:"skywars_kills_solo"`
			ArenaPowerup                    int `json:"arena_powerup"`
			GeneralWins                     int `json:"general_wins"`
			ArenaGladiator                  int `json:"arena_gladiator"`
			MurdermysteryHoarder            int `json:"murdermystery_hoarder"`
			MurdermysteryWinsAsSurvivor     int `json:"murdermystery_wins_as_survivor"`
			TntgamesTntTagWins              int `json:"tntgames_tnt_tag_wins"`
			BedwarsWins                     int `json:"bedwars_wins"`
			SkywarsKillsTeam                int `json:"skywars_kills_team"`
			Christmas2017Advent2018         int `json:"christmas2017_advent_2018"`
			SkywarsKitsSolo                 int `json:"skywars_kits_solo"`
			BedwarsLootBox                  int `json:"bedwars_loot_box"`
			Christmas2017NoChristmas        int `json:"christmas2017_no_christmas"`
			GeneralQuestMaster              int `json:"general_quest_master"`
			SkywarsKitsTeam                 int `json:"skywars_kits_team"`
			SkywarsWinsSolo                 int `json:"skywars_wins_solo"`
			Walls3JackOfAllTrades           int `json:"walls3_jack_of_all_trades"`
			Walls3Coins                     int `json:"walls3_coins"`
			GingerbreadBanker               int `json:"gingerbread_banker"`
			BlitzLooter                     int `json:"blitz_looter"`
			BlitzCoins                      int `json:"blitz_coins"`
			DuelsDuelsWinner                int `json:"duels_duels_winner"`
			DuelsDuelsTraveller             int `json:"duels_duels_traveller"`
			DuelsDuelsWinStreak             int `json:"duels_duels_win_streak"`
			ArcadeArcadeBanker              int `json:"arcade_arcade_banker"`
			ArcadeArcadeWinner              int `json:"arcade_arcade_winner"`
			ArcadeMiniwallsWinner           int `json:"arcade_miniwalls_winner"`
			MurdermysteryKillsAsMurderer    int `json:"murdermystery_kills_as_murderer"`
			SkywarsWinsTeam                 int `json:"skywars_wins_team"`
			ArcadeBountyHunter              int `json:"arcade_bounty_hunter"`
			BuildbattleGuessTheBuildGuesses int `json:"buildbattle_guess_the_build_guesses"`
			BuildbattleBuildBattleScore     int `json:"buildbattle_build_battle_score"`
			MurdermysteryWinsAsMurderer     int `json:"murdermystery_wins_as_murderer"`
			TntgamesTntRunWins              int `json:"tntgames_tnt_run_wins"`
			DuelsDuelsDivision              int `json:"duels_duels_division"`
			EasterThrowEggs                 int `json:"easter_throw_eggs"`
			SkywarsWinsLab                  int `json:"skywars_wins_lab"`
			UhcBounty                       int `json:"uhc_bounty"`
			UhcHunter                       int `json:"uhc_hunter"`
			UhcConsumer                     int `json:"uhc_consumer"`
			UhcMovingUp                     int `json:"uhc_moving_up"`
			SkywarsHeads                    int `json:"skywars_heads"`
			SpeeduhcHunter                  int `json:"speeduhc_hunter"`
			SpeeduhcPromotion               int `json:"speeduhc_promotion"`
			BuildbattleBuildBattleVoter     int `json:"buildbattle_build_battle_voter"`
			BuildbattleBuildBattlePoints    int `json:"buildbattle_build_battle_points"`
			BlitzFightingExpert             int `json:"blitz_fighting_expert"`
			BlitzKills                      int `json:"blitz_kills"`
			BlitzKitExpert                  int `json:"blitz_kit_expert"`
			BlitzKitExperienceCollector     int `json:"blitz_kit_experience_collector"`
			Halloween2017Pumpkinator        int `json:"halloween2017_pumpkinator"`
			UhcUltimatum                    int `json:"uhc_ultimatum"`
			DuelsGoals                      int `json:"duels_goals"`
			DuelsBridgeDoublesWins          int `json:"duels_bridge_doubles_wins"`
			DuelsBridgeWinStreak            int `json:"duels_bridge_win_streak"`
			DuelsBridgeWins                 int `json:"duels_bridge_wins"`
			DuelsUniqueMapWins              int `json:"duels_unique_map_wins"`
		} `json:"achievements"`
		AchievementTracking []interface{} `json:"achievementTracking"`
		AchievementPoints   int           `json:"achievementPoints"`
		FriendRequestsUUID  []interface{} `json:"friendRequestsUuid"`
		Eugene              struct {
			DailyTwoKExp int64 `json:"dailyTwoKExp"`
		} `json:"eugene"`
		NetworkExp float64   `json:"networkExp"`
		LastLogout int64 `json:"lastLogout"`
		Voting     struct {
			Total           int   `json:"total"`
			TotalMcsorg     int   `json:"total_mcsorg"`
			SecondaryMcsorg int   `json:"secondary_mcsorg"`
			LastMcsorg      int64 `json:"last_mcsorg"`
			LastVote        int64 `json:"last_vote"`
			VotesToday      int   `json:"votesToday"`
			SecondaryMcsl   int   `json:"secondary_mcsl"`
			TotalMcsl       int   `json:"total_mcsl"`
			LastMcsl        int64 `json:"last_mcsl"`
		} `json:"voting"`
		LevelingReward0 bool `json:"levelingReward_0"`
		Challenges      struct {
			AllTime struct {
				TNTGAMESTntTagChallenge          int `json:"TNTGAMES__tnt_tag_challenge"`
				ARENACooperationChallenge        int `json:"ARENA__cooperation_challenge"`
				SKYWARSFeedingTheVoidChallenge   int `json:"SKYWARS__feeding_the_void_challenge"`
				BEDWARSSupport                   int `json:"BEDWARS__support"`
				DUELSFeedTheVoidChallenge        int `json:"DUELS__feed_the_void_challenge"`
				BEDWARSOffensive                 int `json:"BEDWARS__offensive"`
				MURDERMYSTERYSherlock            int `json:"MURDER_MYSTERY__sherlock"`
				MURDERMYSTERYHero                int `json:"MURDER_MYSTERY__hero"`
				SKYWARSRushChallenge             int `json:"SKYWARS__rush_challenge"`
				BUILDBATTLEGuesserChallenge      int `json:"BUILD_BATTLE__guesser_challenge"`
				BUILDBATTLETop3Challenge         int `json:"BUILD_BATTLE__top_3_challenge"`
				TNTGAMESTntRunChallenge          int `json:"TNTGAMES__tnt_run_challenge"`
				SKYWARSEndermanChallenge         int `json:"SKYWARS__enderman_challenge"`
				UHCThreatChallenge               int `json:"UHC__threat_challenge"`
				SKYWARSRankedChallenge           int `json:"SKYWARS__ranked_challenge"`
				SURVIVALGAMESResistanceChallenge int `json:"SURVIVAL_GAMES__resistance_challenge"`
				DUELSTargetPracticeChallenge     int `json:"DUELS__target_practice_challenge"`
				SPEEDUHCWizardChallenge          int `json:"SPEED_UHC__wizard_challenge"`
				DUELSTeamsChallenge              int `json:"DUELS__teams_challenge"`
			} `json:"all_time"`
		} `json:"challenges"`
		LevelingReward1 bool `json:"levelingReward_1"`
		PetConsumables  struct {
			CarrotItem   int `json:"CARROT_ITEM"`
			BakedPotato  int `json:"BAKED_POTATO"`
			Feather      int `json:"FEATHER"`
			RottenFlesh  int `json:"ROTTEN_FLESH"`
			SlimeBall    int `json:"SLIME_BALL"`
			Cake         int `json:"CAKE"`
			CookedBeef   int `json:"COOKED_BEEF"`
			WaterBucket  int `json:"WATER_BUCKET"`
			WoodSword    int `json:"WOOD_SWORD"`
			Stick        int `json:"STICK"`
			MilkBucket   int `json:"MILK_BUCKET"`
			GoldRecord   int `json:"GOLD_RECORD"`
			Pork         int `json:"PORK"`
			Leash        int `json:"LEASH"`
			LavaBucket   int `json:"LAVA_BUCKET"`
			MagmaCream   int `json:"MAGMA_CREAM"`
			Apple        int `json:"APPLE"`
			RedRose      int `json:"RED_ROSE"`
			MushroomSoup int `json:"MUSHROOM_SOUP"`
			Cookie       int `json:"COOKIE"`
			Bread        int `json:"BREAD"`
			PumpkinPie   int `json:"PUMPKIN_PIE"`
			Wheat        int `json:"WHEAT"`
			Melon        int `json:"MELON"`
			Bone         int `json:"BONE"`
			HayBlock     int `json:"HAY_BLOCK"`
			RawFish      int `json:"RAW_FISH"`
		} `json:"petConsumables"`
		VanityMeta struct {
			Packages []string `json:"packages"`
		} `json:"vanityMeta"`
		SocialMedia struct {
			Links struct {
			} `json:"links"`
			Prompt bool `json:"prompt"`
		} `json:"socialMedia"`
		LevelingReward2       bool `json:"levelingReward_2"`
		AchievementRewardsNew struct {
			ForPoints200  int64 `json:"for_points_200"`
			ForPoints300  int64 `json:"for_points_300"`
			ForPoints400  int64 `json:"for_points_400"`
			ForPoints500  int64 `json:"for_points_500"`
			ForPoints600  int64 `json:"for_points_600"`
			ForPoints700  int64 `json:"for_points_700"`
			ForPoints800  int64 `json:"for_points_800"`
			ForPoints900  int64 `json:"for_points_900"`
			ForPoints1000 int64 `json:"for_points_1000"`
			ForPoints1100 int64 `json:"for_points_1100"`
			ForPoints1200 int64 `json:"for_points_1200"`
			ForPoints1300 int64 `json:"for_points_1300"`
			ForPoints1400 int64 `json:"for_points_1400"`
			ForPoints1500 int64 `json:"for_points_1500"`
			ForPoints1600 int64 `json:"for_points_1600"`
			ForPoints1700 int64 `json:"for_points_1700"`
			ForPoints1800 int64 `json:"for_points_1800"`
			ForPoints2000 int64 `json:"for_points_2000"`
			ForPoints1900 int64 `json:"for_points_1900"`
			ForPoints2100 int64 `json:"for_points_2100"`
			ForPoints2200 int64 `json:"for_points_2200"`
			ForPoints2300 int64 `json:"for_points_2300"`
		} `json:"achievementRewardsNew"`
		LevelingReward3 bool `json:"levelingReward_3"`
		AchievementSync struct {
			QuakeTiered int `json:"quake_tiered"`
		} `json:"achievementSync"`
		Karma          int    `json:"karma"`
		NewPackageRank string `json:"newPackageRank"`
		LevelUpMVP     int64  `json:"levelUp_MVP"`
		HousingMeta    struct {
			AllowedBlocks      []string `json:"allowedBlocks"`
			Packages           []string `json:"packages"`
			FirstHouseJoinMs   int64    `json:"firstHouseJoinMs"`
			TutorialStep       string   `json:"tutorialStep"`
			GivenCookies104996 []string `json:"given_cookies_104996"`
			GivenCookies105003 []string `json:"given_cookies_105003"`
			GivenCookies105011 []string `json:"given_cookies_105011"`
		} `json:"housingMeta"`
		LevelingReward4         bool  `json:"levelingReward_4"`
		LastAdsenseGenerateTime int64 `json:"lastAdsenseGenerateTime"`
		//LastClaimedReward       int64 `json:"lastClaimedReward"`
		TotalRewards            int   `json:"totalRewards"`
		TotalDailyRewards       int   `json:"totalDailyRewards"`
		RewardStreak            int   `json:"rewardStreak"`
		RewardScore             int   `json:"rewardScore"`
		RewardHighScore         int   `json:"rewardHighScore"`
		LevelingReward5         bool  `json:"levelingReward_5"`
		LevelUpMVPPLUS          int64 `json:"levelUp_MVP_PLUS"`
		LevelingReward6         bool  `json:"levelingReward_6"`
		LevelingReward7         bool  `json:"levelingReward_7"`
		LevelingReward8         bool  `json:"levelingReward_8"`
		AdsenseTokens           int   `json:"adsense_tokens"`
		LevelingReward9         bool  `json:"levelingReward_9"`
		ParkourCheckpointBests  struct {
			Bedwars struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
			} `json:"Bedwars"`
			BuildBattle struct {
				Num0 int `json:"0"`
			} `json:"BuildBattle"`
			Duels struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
			} `json:"Duels"`
			ArcadeGames struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
				Num4 int `json:"4"`
				Num5 int `json:"5"`
				Num6 int `json:"6"`
			} `json:"ArcadeGames"`
			TruePVPParkour struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
			} `json:"TruePVPParkour"`
			CopsnCrims struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
				Num4 int `json:"4"`
			} `json:"CopsnCrims"`
			Legacy struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
			} `json:"Legacy"`
			Warlords struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
			} `json:"Warlords"`
			Prototype struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
			} `json:"Prototype"`
			MainLobby2017 struct {
				Num0 int `json:"0"`
			} `json:"mainLobby2017"`
			Tnt struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
			} `json:"TNT"`
			SkywarsAug2017 struct {
				Num0 int `json:"0"`
			} `json:"SkywarsAug2017"`
		} `json:"parkourCheckpointBests"`
		LevelingReward10 bool `json:"levelingReward_10"`
		LevelingReward11 bool `json:"levelingReward_11"`
		LevelingReward12 bool `json:"levelingReward_12"`
		Quests           struct {
			BedwarsDailyGifts struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"bedwars_daily_gifts"`
			BedwarsDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"bedwars_daily_win"`
			BedwarsDailyOneMore struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"bedwars_daily_one_more"`
			BedwarsWeeklyBedElims struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						BedwarsBedElims int `json:"bedwars_bed_elims"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_weekly_bed_elims"`
			BedwarsWeeklyDreamWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_weekly_dream_win"`
			TntWeeklyPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"tnt_weekly_play"`
			SkywarsArcadeWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_arcade_win"`
			SkywarsTeamWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"skywars_team_win"`
			SkywarsTeamKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"skywars_team_kills"`
			SkywarsSoloWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"skywars_solo_win"`
			SkywarsSoloKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"skywars_solo_kills"`
			SkywarsWeeklyKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"skywars_weekly_kills"`
			PrototypePitDailyKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"prototype_pit_daily_kills"`
			PrototypePitDailyContract struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"prototype_pit_daily_contract"`
			PrototypePitWeeklyGold struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						PrototypePitWeeklyGold int `json:"prototype_pit_weekly_gold"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"prototype_pit_weekly_gold"`
			DuelsPlayer struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"duels_player"`
			DuelsKiller struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"duels_killer"`
			DuelsWinner struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"duels_winner"`
			TntTnttagDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"tnt_tnttag_daily"`
			DuelsWeeklyKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"duels_weekly_kills"`
			DuelsWeeklyWins struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"duels_weekly_wins"`
			SkywarsCorruptWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"skywars_corrupt_win"`
			UhcTeam struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"uhc_team"`
			UhcSolo struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"uhc_solo"`
			UhcDm struct {
				Active struct {
					Objectives struct {
						UhcKills int `json:"uhc_kills"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"uhc_dm"`
			UhcWeekly struct {
				Active struct {
					Objectives struct {
						UhcKills int `json:"uhc_kills"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"uhc_weekly"`
			SoloBrawler struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"solo_brawler"`
			TeamBrawler struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"team_brawler"`
			UhcMadness struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"uhc_madness"`
			BedwarsDailyNightmares struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"bedwars_daily_nightmares"`
			BedwarsWeeklyPumpkinator struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"bedwars_weekly_pumpkinator"`
		} `json:"quests"`
		LevelingReward13   bool `json:"levelingReward_13"`
		LevelingReward14   bool `json:"levelingReward_14"`
		LevelingReward15   bool `json:"levelingReward_15"`
		LevelingReward16   bool `json:"levelingReward_16"`
		LevelingReward17   bool `json:"levelingReward_17"`
		LevelingReward18   bool `json:"levelingReward_18"`
		ParkourCompletions struct {
			Bedwars []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Bedwars"`
			Duels []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Duels"`
			ArcadeGames []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"ArcadeGames"`
			TruePVPParkour []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"TruePVPParkour"`
			CopsnCrims []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"CopsnCrims"`
			Legacy []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Legacy"`
			Warlords []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Warlords"`
			Tnt []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"TNT"`
		} `json:"parkourCompletions"`
		LevelingReward19 bool `json:"levelingReward_19"`
		GiftingMeta      struct {
			GiftsGiven          int      `json:"giftsGiven"`
			BundlesGiven        int      `json:"bundlesGiven"`
			RealBundlesGiven    int      `json:"realBundlesGiven"`
			RealBundlesReceived int      `json:"realBundlesReceived"`
			BundlesReceived     int      `json:"bundlesReceived"`
			Milestones          []string `json:"milestones"`
		} `json:"giftingMeta"`
		LevelingReward20 bool   `json:"levelingReward_20"`
		LevelingReward21 bool   `json:"levelingReward_21"`
		LevelingReward22 bool   `json:"levelingReward_22"`
		LevelingReward23 bool   `json:"levelingReward_23"`
		LevelingReward24 bool   `json:"levelingReward_24"`
		LevelingReward25 bool   `json:"levelingReward_25"`
		LevelingReward26 bool   `json:"levelingReward_26"`
		LevelingReward27 bool   `json:"levelingReward_27"`
		LevelingReward28 bool   `json:"levelingReward_28"`
		LevelingReward29 bool   `json:"levelingReward_29"`
		LevelingReward30 bool   `json:"levelingReward_30"`
		LevelingReward31 bool   `json:"levelingReward_31"`
		LevelingReward32 bool   `json:"levelingReward_32"`
		LevelingReward33 bool   `json:"levelingReward_33"`
		LevelingReward34 bool   `json:"levelingReward_34"`
		RankPlusColor    string `json:"rankPlusColor"`
		LevelingReward35 bool   `json:"levelingReward_35"`
		LevelingReward36 bool   `json:"levelingReward_36"`
		LevelingReward37 bool   `json:"levelingReward_37"`
		LevelingReward38 bool   `json:"levelingReward_38"`
		LevelingReward39 bool   `json:"levelingReward_39"`
		LevelingReward40 bool   `json:"levelingReward_40"`
		LevelingReward41 bool   `json:"levelingReward_41"`
		LevelingReward42 bool   `json:"levelingReward_42"`
		LevelingReward43 bool   `json:"levelingReward_43"`
		LevelingReward44 bool   `json:"levelingReward_44"`
		LevelingReward45 bool   `json:"levelingReward_45"`
		LevelingReward46 bool   `json:"levelingReward_46"`
		LevelingReward47 bool   `json:"levelingReward_47"`
		LevelingReward48 bool   `json:"levelingReward_48"`
		LevelingReward49 bool   `json:"levelingReward_49"`
		LevelingReward50 bool   `json:"levelingReward_50"`
		LevelingReward51 bool   `json:"levelingReward_51"`
		LevelingReward52 bool   `json:"levelingReward_52"`
		LevelingReward53 bool   `json:"levelingReward_53"`
		LevelingReward54 bool   `json:"levelingReward_54"`
		LevelingReward55 bool   `json:"levelingReward_55"`
		LevelingReward56 bool   `json:"levelingReward_56"`
		LevelingReward57 bool   `json:"levelingReward_57"`
		LevelingReward58 bool   `json:"levelingReward_58"`
		CurrentPet       string `json:"currentPet"`
		PetStats         struct {
			Pig struct {
				Name string `json:"name"`
			} `json:"PIG"`
		} `json:"petStats"`
		Channel          string `json:"channel"`
		LevelingReward59 bool   `json:"levelingReward_59"`
		LevelingReward60 bool   `json:"levelingReward_60"`
		LevelingReward61 bool   `json:"levelingReward_61"`
		LevelingReward62 bool   `json:"levelingReward_62"`
		Monthlycrates    struct {
			One2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"1-2019"`
			One12018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
			} `json:"11-2018"`
			One22018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"12-2018"`
			Two2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"2-2019"`
			Three2019 struct {
				MvpPlus bool `json:"MVP_PLUS"`
				Mvp     bool `json:"MVP"`
				VipPlus bool `json:"VIP_PLUS"`
				Vip     bool `json:"VIP"`
				Regular bool `json:"REGULAR"`
			} `json:"3-2019"`
			Four2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				MvpPlus bool `json:"MVP_PLUS"`
				Mvp     bool `json:"MVP"`
			} `json:"4-2019"`
			One12019 struct {
				MvpPlus bool `json:"MVP_PLUS"`
				Mvp     bool `json:"MVP"`
				VipPlus bool `json:"VIP_PLUS"`
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
			} `json:"11-2019"`
		} `json:"monthlycrates"`
		LevelingReward63    bool     `json:"levelingReward_63"`
		LevelingReward64    bool     `json:"levelingReward_64"`
		CurrentCloak        string   `json:"currentCloak"`
		MostRecentGameType  string   `json:"mostRecentGameType"`
		GuildInvites        []string `json:"guildInvites"`
		AdventRewardsV22018 struct {
			Day1  int64 `json:"day1"`
			Day2  int64 `json:"day2"`
			Day3  int64 `json:"day3"`
			Day4  int64 `json:"day4"`
			Day5  int64 `json:"day5"`
			Day6  int64 `json:"day6"`
			Day7  int64 `json:"day7"`
			Day8  int64 `json:"day8"`
			Day9  int64 `json:"day9"`
			Day10 int64 `json:"day10"`
			Day11 int64 `json:"day11"`
			Day12 int64 `json:"day12"`
			Day14 int64 `json:"day14"`
			Day15 int64 `json:"day15"`
			Day16 int64 `json:"day16"`
			Day18 int64 `json:"day18"`
			Day19 int64 `json:"day19"`
			Day20 int64 `json:"day20"`
			Day24 int64 `json:"day24"`
			Day25 int64 `json:"day25"`
		} `json:"adventRewards_v2_2018"`
	} `json:"player"`
}



type HypixelFriendsResponse struct {
	Success bool   `json:"success"`
	UUID    string `json:"uuid"`
	Records []struct {
		ID           string `json:"_id"`
		UUIDSender   string `json:"uuidSender"`
		UUIDReceiver string `json:"uuidReceiver"`
		Started      int64  `json:"started"`
	} `json:"records"`
}

type HypixelStatusResponse struct {
	Success bool   `json:"success"`
	UUID    string `json:"uuid"`
	Session struct {
		Online   bool   `json:"online"`
		GameType string `json:"gameType"`
		Mode     string `json:"mode"`
	} `json:"session"`
}


type HypixelSkyblockResponse struct {
	Success  bool `json:"success"`
	Profiles []struct {
		ProfileID string `json:"profile_id"`
		Members   map[string] struct {
				LastSave       int64 `json:"last_save"`
				CoopInvitation struct {
					Timestamp int64  `json:"timestamp"`
					InvitedBy string `json:"invited_by"`
					Confirmed bool   `json:"confirmed"`
				} `json:"coop_invitation"`
				FirstJoin    int64 `json:"first_join"`
				FirstJoinHub int   `json:"first_join_hub"`
				Stats        struct {
					Deaths                                              float64 `json:"deaths"`
					DeathsVoid                                          float64 `json:"deaths_void"`
					HighestCriticalDamage                               float64 `json:"highest_critical_damage"`
					Kills                                               float64 `json:"kills"`
					KillsPondSquid                                      float64 `json:"kills_pond_squid"`
					KillsZombie                                         float64 `json:"kills_zombie"`
					DeathsEnderman                                      float64 `json:"deaths_enderman"`
					KillsSkeleton                                       float64 `json:"kills_skeleton"`
					KillsSpider                                         float64 `json:"kills_spider"`
					KillsEnderman                                       float64 `json:"kills_enderman"`
					KillsPig                                            float64 `json:"kills_pig"`
					PetMilestoneOresMined                               float64 `json:"pet_milestone_ores_mined"`
					KillsLapisZombie                                    float64 `json:"kills_lapis_zombie"`
					KillsRedstonePigman                                 float64 `json:"kills_redstone_pigman"`
					KillsEmeraldSlime                                   float64 `json:"kills_emerald_slime"`
					KillsDiamondZombie                                  float64 `json:"kills_diamond_zombie"`
					KillsDiamondSkeleton                                float64 `json:"kills_diamond_skeleton"`
					DeathsDiamondZombie                                 float64 `json:"deaths_diamond_zombie"`
					AuctionsBids                                        float64 `json:"auctions_bids"`
					AuctionsHighestBid                                  float64     `json:"auctions_highest_bid"`
					AuctionsWon                                         float64 `json:"auctions_won"`
					AuctionsBoughtUncommon                              float64 `json:"auctions_bought_uncommon"`
					AuctionsGoldSpent                                   float64     `json:"auctions_gold_spent"`
					DeathsSpider                                        float64 `json:"deaths_spider"`
					KillsSheep                                          float64 `json:"kills_sheep"`
					DeathsFire                                          float64 `json:"deaths_fire"`
					DeathsFall                                          float64 `json:"deaths_fall"`
					KillsInvisibleCreeper                               float64 `json:"kills_invisible_creeper"`
					KillsWitherSkeleton                                 float64 `json:"kills_wither_skeleton"`
					KillsMagmaCube                                      float64 `json:"kills_magma_cube"`
					KillsFireballMagmaCube                              float64 `json:"kills_fireball_magma_cube"`
					KillsDasherSpider                                   float64 `json:"kills_dasher_spider"`
					KillsSplitterSpiderSilverfish                       float64 `json:"kills_splitter_spider_silverfish"`
					KillsBlaze                                          float64 `json:"kills_blaze"`
					KillsPigman                                         float64 `json:"kills_pigman"`
					KillsGhast                                          float64 `json:"kills_ghast"`
					DeathsEndermite                                     float64 `json:"deaths_endermite"`
					KillsEndermite                                      float64 `json:"kills_endermite"`
					KillsObsidianWither                                 float64 `json:"kills_obsidian_wither"`
					KillsWatcher                                        float64 `json:"kills_watcher"`
					DeathsWatcher                                       float64 `json:"deaths_watcher"`
					DeathsZealotEnderman                                float64 `json:"deaths_zealot_enderman"`
					KillsZombieVillager                                 float64 `json:"kills_zombie_villager"`
					DeathsZombie                                        float64 `json:"deaths_zombie"`
					DeathsPlayer                                        float64 `json:"deaths_player"`
					AuctionsBoughtEpic                                  float64 `json:"auctions_bought_epic"`
					KillsUnburriedZombie                                float64 `json:"kills_unburried_zombie"`
					DeathsUnburriedZombie                               float64 `json:"deaths_unburried_zombie"`
					AuctionsBoughtCommon                                float64 `json:"auctions_bought_common"`
					KillsSplitterSpider                                 float64 `json:"kills_splitter_spider"`
					KillsSpiderJockey                                   float64 `json:"kills_spider_jockey"`
					KillsJockeySkeleton                                 float64 `json:"kills_jockey_skeleton"`
					KillsVoraciousSpider                                float64 `json:"kills_voracious_spider"`
					KillsWeaverSpider                                   float64 `json:"kills_weaver_spider"`
					KillsJockeyShotSilverfish                           float64 `json:"kills_jockey_shot_silverfish"`
					KillsRuinWolf                                       float64 `json:"kills_ruin_wolf"`
					KillsOldWolf                                        float64 `json:"kills_old_wolf"`
					DeathsOldWolf                                       float64 `json:"deaths_old_wolf"`
					DeathsWolf                                          float64 `json:"deaths_wolf"`
					KillsRandomSlime                                    float64 `json:"kills_random_slime"`
					AuctionsCreated                                     float64 `json:"auctions_created"`
					AuctionsFees                                        float64     `json:"auctions_fees"`
					KillsGeneratorGhast                                 float64 `json:"kills_generator_ghast"`
					AuctionsCompleted                                   float64 `json:"auctions_completed"`
					AuctionsSoldUncommon                                float64 `json:"auctions_sold_uncommon"`
					AuctionsGoldEarned                                  float64     `json:"auctions_gold_earned"`
					AuctionsNoBids                                      float64 `json:"auctions_no_bids"`
					KillsNightRespawningSkeleton                        float64 `json:"kills_night_respawning_skeleton"`
					KillsBatPinata                                      float64 `json:"kills_bat_pinata"`
					KillsHorsemanBat                                    float64 `json:"kills_horseman_bat"`
					DeathsHorsemanHorse                                 float64 `json:"deaths_horseman_horse"`
					KillsHorsemanZombie                                 float64 `json:"kills_horseman_zombie"`
					GiftsGiven                                          float64 `json:"gifts_given"`
					KillsPlayer                                         float64 `json:"kills_player"`
					ItemsFished                                         float64 `json:"items_fished"`
					ItemsFishedNormal                                   float64 `json:"items_fished_normal"`
					PetMilestoneSeaCreaturesKilled                      float64 `json:"pet_milestone_sea_creatures_killed"`
					KillsNightSquid                                     float64 `json:"kills_night_squid"`
					KillsSeaWalker                                      float64 `json:"kills_sea_walker"`
					KillsSeaGuardian                                    float64 `json:"kills_sea_guardian"`
					ItemsFishedTreasure                                 float64 `json:"items_fished_treasure"`
					DeathsSeaGuardian                                   float64 `json:"deaths_sea_guardian"`
					ItemsFishedLargeTreasure                            float64 `json:"items_fished_large_treasure"`
					GiftsReceived                                       float64 `json:"gifts_received"`
					KillsZealotEnderman                                 float64 `json:"kills_zealot_enderman"`
					KillsRespawningSkeleton                             float64 `json:"kills_respawning_skeleton"`
					DeathsYoungDragon                                   float64 `json:"deaths_young_dragon"`
					AuctionsBoughtLegendary                             float64 `json:"auctions_bought_legendary"`
					ForagingRaceBestTime                                float64 `json:"foraging_race_best_time"`
					KillsSeaArcher                                      float64 `json:"kills_sea_archer"`
					DeathsWitherSkeleton                                float64 `json:"deaths_wither_skeleton"`
					DeathsBlaze                                         float64 `json:"deaths_blaze"`
					AuctionsBoughtRare                                  float64 `json:"auctions_bought_rare"`
					DeathsUnknown                                       float64 `json:"deaths_unknown"`
					DeathsUnstableDragon                                float64 `json:"deaths_unstable_dragon"`
					EnderCrystalsDestroyed                              float64 `json:"ender_crystals_destroyed"`
					KillsMagmaCubeBoss                                  float64 `json:"kills_magma_cube_boss"`
					DeathsCorruptedProtector                            float64 `json:"deaths_corrupted_protector"`
					KillsZombieDeep                                     float64 `json:"kills_zombie_deep"`
					KillsGuardianDefender                               float64 `json:"kills_guardian_defender"`
					KillsCreeper                                        float64 `json:"kills_creeper"`
					KillsWitch                                          float64 `json:"kills_witch"`
					DeathsProtectorDragon                               float64 `json:"deaths_protector_dragon"`
					KillsGeneratorSlime                                 float64 `json:"kills_generator_slime"`
					KillsGeneratorMagmaCube                             float64 `json:"kills_generator_magma_cube"`
					DeathsOldDragon                                     float64 `json:"deaths_old_dragon"`
					AuctionsSoldRare                                    float64 `json:"auctions_sold_rare"`
					KillsBroodMotherCaveSpider                          float64 `json:"kills_brood_mother_cave_spider"`
					AuctionsSoldCommon                                  float64 `json:"auctions_sold_common"`
					KillsCatfish                                        float64 `json:"kills_catfish"`
					KillsChickenDeep                                    float64 `json:"kills_chicken_deep"`
					KillsCarrotKing                                     float64 `json:"kills_carrot_king"`
					KillsSeaLeech                                       float64 `json:"kills_sea_leech"`
					KillsWaterHydra                                     float64 `json:"kills_water_hydra"`
					KillsDeepSeaProtector                               float64 `json:"kills_deep_sea_protector"`
					KillsFrozenSteve                                    float64 `json:"kills_frozen_steve"`
					KillsFrostyTheSnowman                               float64 `json:"kills_frosty_the_snowman"`
					KillsChicken                                        float64 `json:"kills_chicken"`
					KillsCow                                            float64 `json:"kills_cow"`
					KillsRabbit                                         float64 `json:"kills_rabbit"`
					KillsCryptLurker                                    float64 `json:"kills_crypt_lurker"`
					KillsCryptTankZombie                                float64 `json:"kills_crypt_tank_zombie"`
					KillsZombieGrunt                                    float64 `json:"kills_zombie_grunt"`
					KillsCryptDreadlord                                 float64 `json:"kills_crypt_dreadlord"`
					DeathsLostAdventurer                                float64 `json:"deaths_lost_adventurer"`
					KillsScaredSkeleton                                 float64 `json:"kills_scared_skeleton"`
					KillsDungeonRespawningSkeleton                      float64 `json:"kills_dungeon_respawning_skeleton"`
					KillsCryptSouleater                                 float64 `json:"kills_crypt_souleater"`
					DeathsDiamondGuy                                    float64 `json:"deaths_diamond_guy"`
					KillsSkeletonGrunt                                  float64 `json:"kills_skeleton_grunt"`
					DeathsSkeletonGrunt                                 float64 `json:"deaths_skeleton_grunt"`
					DeathsWatcherSummonUndead                           float64 `json:"deaths_watcher_summon_undead"`
					KillsWatcherSummonUndead                            float64 `json:"kills_watcher_summon_undead"`
					KillsCellarSpider                                   float64 `json:"kills_cellar_spider"`
					KillsLostAdventurer                                 float64 `json:"kills_lost_adventurer"`
					DeathsTrap                                          float64 `json:"deaths_trap"`
					KillsSkeletonSoldier                                float64 `json:"kills_skeleton_soldier"`
					DeathsScaredSkeleton                                float64 `json:"deaths_scared_skeleton"`
					DeathsCryptSouleater                                float64 `json:"deaths_crypt_souleater"`
					DeathsCryptLurker                                   float64 `json:"deaths_crypt_lurker"`
					DeathsCryptDreadlord                                float64 `json:"deaths_crypt_dreadlord"`
					DeathsSkeletonSoldier                               float64 `json:"deaths_skeleton_soldier"`
					KillsDiamondGuy                                     float64 `json:"kills_diamond_guy"`
					KillsCryptUndead                                    float64 `json:"kills_crypt_undead"`
					DeathsZombieGrunt                                   float64 `json:"deaths_zombie_grunt"`
					KillsSniperSkeleton                                 float64 `json:"kills_sniper_skeleton"`
					KillsCryptUndeadPieter                              float64 `json:"kills_crypt_undead_pieter"`
					DeathsBonzoSummonUndead                             float64 `json:"deaths_bonzo_summon_undead"`
					KillsBonzoSummonUndead                              float64 `json:"kills_bonzo_summon_undead"`
					DeathsBonzo                                         float64 `json:"deaths_bonzo"`
					KillsBlazeHigherOrLower                             float64 `json:"kills_blaze_higher_or_lower"`
					KillsSkeletonMaster                                 float64 `json:"kills_skeleton_master"`
					DeathsScarfWarrior                                  float64 `json:"deaths_scarf_warrior"`
					KillsRevenantZombie                                 float64 `json:"kills_revenant_zombie"`
					KillsDungeonSecretBat                               float64 `json:"kills_dungeon_secret_bat"`
					KillsCryptUndeadChristian                           float64 `json:"kills_crypt_undead_christian"`
					DungeonHubCrystalCoreAnythingNoReturnBestTime       float64 `json:"dungeon_hub_crystal_core_anything_no_return_best_time"`
					DungeonHubCrystalCoreNoPearlsNoReturnBestTime       float64 `json:"dungeon_hub_crystal_core_no_pearls_no_return_best_time"`
					DungeonHubCrystalCoreNoAbilitiesNoReturnBestTime    float64 `json:"dungeon_hub_crystal_core_no_abilities_no_return_best_time"`
					KillsCryptUndeadAlexander                           float64 `json:"kills_crypt_undead_alexander"`
					DeathsScarfArcher                                   float64 `json:"deaths_scarf_archer"`
					DeathsScarfMage                                     float64 `json:"deaths_scarf_mage"`
					DeathsScarf                                         float64 `json:"deaths_scarf"`
					DungeonHubCrystalCoreNothingNoReturnBestTime        float64 `json:"dungeon_hub_crystal_core_nothing_no_return_best_time"`
					DungeonHubGiantMushroomAnythingNoReturnBestTime     float64 `json:"dungeon_hub_giant_mushroom_anything_no_return_best_time"`
					DungeonHubGiantMushroomNoPearlsNoReturnBestTime     float64 `json:"dungeon_hub_giant_mushroom_no_pearls_no_return_best_time"`
					DungeonHubGiantMushroomNoAbilitiesNoReturnBestTime  float64 `json:"dungeon_hub_giant_mushroom_no_abilities_no_return_best_time"`
					DungeonHubGiantMushroomNothingNoReturnBestTime      float64 `json:"dungeon_hub_giant_mushroom_nothing_no_return_best_time"`
					DungeonHubPrecursorRuinsAnythingNoReturnBestTime    float64 `json:"dungeon_hub_precursor_ruins_anything_no_return_best_time"`
					DungeonHubPrecursorRuinsNoPearlsNoReturnBestTime    float64 `json:"dungeon_hub_precursor_ruins_no_pearls_no_return_best_time"`
					DungeonHubPrecursorRuinsNoAbilitiesNoReturnBestTime float64 `json:"dungeon_hub_precursor_ruins_no_abilities_no_return_best_time"`
					DungeonHubPrecursorRuinsNothingNoReturnBestTime     float64 `json:"dungeon_hub_precursor_ruins_nothing_no_return_best_time"`
					KillsTarantulaSpider                                float64 `json:"kills_tarantula_spider"`
					DeathsSniperSkeleton                                float64 `json:"deaths_sniper_skeleton"`
					KillsLonelySpider                                   float64 `json:"kills_lonely_spider"`
					KillsCryptUndeadValentin                            float64 `json:"kills_crypt_undead_valentin"`
					KillsPackSpirit                                     float64 `json:"kills_pack_spirit"`
					KillsSoulOfTheAlpha                                 float64 `json:"kills_soul_of_the_alpha"`
					KillsHowlingSpirit                                  float64 `json:"kills_howling_spirit"`
					KillsDungeonRespawningSkeletonSkull                 float64 `json:"kills_dungeon_respawning_skeleton_skull"`
					KillsCryptUndeadMarius                              float64 `json:"kills_crypt_undead_marius"`
					DeathsObsidianWither                                float64 `json:"deaths_obsidian_wither"`
					KillsCorruptedProtector                             float64 `json:"kills_corrupted_protector"`
					DeathsTarantulaSpider                               float64 `json:"deaths_tarantula_spider"`
					KillsBroodMotherSpider                              float64 `json:"kills_brood_mother_spider"`
					DeathsSkeletor                                      float64 `json:"deaths_skeletor"`
					DeathsSkeletonMaster                                float64 `json:"deaths_skeleton_master"`
					DeathsDungeonRespawningSkeleton                     float64 `json:"deaths_dungeon_respawning_skeleton"`
					DeathsShadowAssassin                                float64 `json:"deaths_shadow_assassin"`
					DeathsCryptTankZombie                               float64 `json:"deaths_crypt_tank_zombie"`
					KillsScarfWarrior                                   float64 `json:"kills_scarf_warrior"`
					KillsParasite                                       float64 `json:"kills_parasite"`
					AuctionsSoldEpic                                    float64 `json:"auctions_sold_epic"`
					MythosKills                                         float64 `json:"mythos_kills"`
					KillsMinosHunter                                    float64 `json:"kills_minos_hunter"`
					MythosBurrowsDugNext                                float64 `json:"mythos_burrows_dug_next"`
					MythosBurrowsDugNextCOMMON                          float64 `json:"mythos_burrows_dug_next_COMMON"`
					MythosBurrowsDugCombat                              float64 `json:"mythos_burrows_dug_combat"`
					MythosBurrowsDugCombatCOMMON                        float64 `json:"mythos_burrows_dug_combat_COMMON"`
					KillsSiameseLynx                                    float64 `json:"kills_siamese_lynx"`
					MythosBurrowsDugTreasure                            float64 `json:"mythos_burrows_dug_treasure"`
					MythosBurrowsDugTreasureCOMMON                      float64 `json:"mythos_burrows_dug_treasure_COMMON"`
					MythosBurrowsChainsComplete                         float64 `json:"mythos_burrows_chains_complete"`
					MythosBurrowsChainsCompleteCOMMON                   float64 `json:"mythos_burrows_chains_complete_COMMON"`
					MythosBurrowsDugNextRARE                            float64 `json:"mythos_burrows_dug_next_RARE"`
					MythosBurrowsDugCombatRARE                          float64 `json:"mythos_burrows_dug_combat_RARE"`
					KillsMinotaur                                       float64 `json:"kills_minotaur"`
					DeathsMinotaur                                      float64 `json:"deaths_minotaur"`
					MythosBurrowsDugTreasureRARE                        float64 `json:"mythos_burrows_dug_treasure_RARE"`
					DeathsGaiaConstruct                                 float64 `json:"deaths_gaia_construct"`
					MythosBurrowsChainsCompleteRARE                     float64 `json:"mythos_burrows_chains_complete_RARE"`
					KillsGaiaConstruct                                  float64 `json:"kills_gaia_construct"`
					KillsLiquidHotMagma                                 float64 `json:"kills_liquid_hot_magma"`
					MostWinterSnowballsHit                              float64 `json:"most_winter_snowballs_hit"`
					MostWinterDamageDealt                               float64 `json:"most_winter_damage_dealt"`
					MostWinterMagmaDamageDealt                          float64 `json:"most_winter_magma_damage_dealt"`
					KillsNurseShark                                     float64 `json:"kills_nurse_shark"`
					KillsBlueShark                                      float64 `json:"kills_blue_shark"`
					KillsSeaWitch                                       float64 `json:"kills_sea_witch"`
					KillsTigerShark                                     float64 `json:"kills_tiger_shark"`
					DeathsWaterHydra                                    float64 `json:"deaths_water_hydra"`
					AuctionsSoldLegendary                               float64 `json:"auctions_sold_legendary"`
					AuctionsBoughtSpecial                               float64 `json:"auctions_bought_special"`
					EndRaceBestTime                                     float64 `json:"end_race_best_time"`
					MythosBurrowsDugNextNull                            float64 `json:"mythos_burrows_dug_next_null"`
					MythosBurrowsDugCombatNull                          float64 `json:"mythos_burrows_dug_combat_null"`
					MythosBurrowsDugTreasureNull                        float64 `json:"mythos_burrows_dug_treasure_null"`
					MythosBurrowsChainsCompleteNull                     float64 `json:"mythos_burrows_chains_complete_null"`
					DeathsSpiritBat                                     float64 `json:"deaths_spirit_bat"`
					DeathsLividClone                                    float64 `json:"deaths_livid_clone"`
					DeathsLivid                                         float64 `json:"deaths_livid"`
					KillsWitherGourd                                    float64 `json:"kills_wither_gourd"`
					KillsPhantomSpirit                                  float64 `json:"kills_phantom_spirit"`
					KillsWraith                                         float64 `json:"kills_wraith"`
					KillsTrickOrTreater                                 float64 `json:"kills_trick_or_treater"`
					KillsBatSpooky                                      float64 `json:"kills_bat_spooky"`
					KillsScaryJerry                                     float64 `json:"kills_scary_jerry"`
					KillsBattyWitch                                     float64 `json:"kills_batty_witch"`
					KillsScarfPriest                                    float64 `json:"kills_scarf_priest"`
					KillsScarfArcher                                    float64 `json:"kills_scarf_archer"`
					KillsZombieSoldier                                  float64 `json:"kills_zombie_soldier"`
					KillsSkeletor                                       float64 `json:"kills_skeletor"`
					KillsZombieKnight                                   float64 `json:"kills_zombie_knight"`
					KillsProfessorGuardianSummon                        float64 `json:"kills_professor_guardian_summon"`
					DeathsZombieSoldier                                 float64 `json:"deaths_zombie_soldier"`
					KillsWitchBat                                       float64 `json:"kills_witch_bat"`
					DeathsDeathmite                                     float64 `json:"deaths_deathmite"`
					DeathsProfessorMageGuardian                         float64 `json:"deaths_professor_mage_guardian"`
					KillsCryptUndeadNicholas                            float64 `json:"kills_crypt_undead_nicholas"`
					KillsShadowAssassin                                 float64 `json:"kills_shadow_assassin"`
					KillsWatcherBonzo                                   float64 `json:"kills_watcher_bonzo"`
					DeathsProfessor                                     float64 `json:"deaths_professor"`
					KillsCryptWitherskeleton                            float64 `json:"kills_crypt_witherskeleton"`
					KillsSuperTankZombie                                float64 `json:"kills_super_tank_zombie"`
					KillsSuperArcher                                    float64 `json:"kills_super_archer"`
					KillsSpiritBull                                     float64 `json:"kills_spirit_bull"`
					KillsSpiritRabbit                                   float64 `json:"kills_spirit_rabbit"`
					KillsSpiritWolf                                     float64 `json:"kills_spirit_wolf"`
					KillsSpiritSheep                                    float64 `json:"kills_spirit_sheep"`
					DeathsSpiritChicken                                 float64 `json:"deaths_spirit_chicken"`
					KillsSpiritBat                                      float64 `json:"kills_spirit_bat"`
					DeathsSpiritWolf                                    float64 `json:"deaths_spirit_wolf"`
					KillsSpiritMiniboss                                 float64 `json:"kills_spirit_miniboss"`
					DeathsSpiritMiniboss                                float64 `json:"deaths_spirit_miniboss"`
					KillsThorn                                          float64 `json:"kills_thorn"`
					KillsSpiritChicken                                  float64 `json:"kills_spirit_chicken"`
					DeathsSpiritBull                                    float64 `json:"deaths_spirit_bull"`
					KillsTentaclees                                     float64 `json:"kills_tentaclees"`
					AuctionsSoldSpecial                                 float64 `json:"auctions_sold_special"`
					DeathsSpiritSheep                                   float64 `json:"deaths_spirit_sheep"`
					DeathsSpiritRabbit                                  float64 `json:"deaths_spirit_rabbit"`
					DeathsWatcherBonzo                                  float64 `json:"deaths_watcher_bonzo"`
					DeathsArmorStand                                    float64 `json:"deaths_armor_stand"`
					DeathsTentaclees                                    float64 `json:"deaths_tentaclees"`
					KillsWatcherScarf                                   float64 `json:"kills_watcher_scarf"`
					DeathsSkeletorPrime                                 float64 `json:"deaths_skeletor_prime"`
					KillsZombieCommander                                float64 `json:"kills_zombie_commander"`
					KillsSkeletorPrime                                  float64 `json:"kills_skeletor_prime"`
					DeathsMimic                                         float64 `json:"deaths_mimic"`
					DeathsWatcherLivid                                  float64 `json:"deaths_watcher_livid"`
					KillsWatcherLivid                                   float64 `json:"kills_watcher_livid"`
					KillsCryptUndeadBernhard                            float64 `json:"kills_crypt_undead_bernhard"`
					DeathsCryptWitherskeleton                           float64 `json:"deaths_crypt_witherskeleton"`
					DeathsSadanStatue                                   float64 `json:"deaths_sadan_statue"`
					KillsKingMidas                                      float64 `json:"kills_king_midas"`
					KillsSadanGolem                                     float64 `json:"kills_sadan_golem"`
					DeathsSadan                                         float64 `json:"deaths_sadan"`
					KillsMimic                                          float64 `json:"kills_mimic"`
					DeathsSadanGolem                                    float64 `json:"deaths_sadan_golem"`
					DeathsSuperTankZombie                               float64 `json:"deaths_super_tank_zombie"`
					KillsSadanStatue                                    float64 `json:"kills_sadan_statue"`
					DeathsWatcherGuardian                               float64 `json:"deaths_watcher_guardian"`
					DeathsWatcherScarf                                  float64 `json:"deaths_watcher_scarf"`
					KillsGrinch                                         float64 `json:"kills_grinch"`
					DeathsSadanGiant                                    float64 `json:"deaths_sadan_giant"`
					KillsSadanGiant                                     float64 `json:"kills_sadan_giant"`
					DeathsKingMidas                                     float64 `json:"deaths_king_midas"`
					KillsGreatWhiteShark                                float64 `json:"kills_great_white_shark"`
					KillsSkeletonLord                                   float64 `json:"kills_skeleton_lord"`
					KillsZombieLord                                     float64 `json:"kills_zombie_lord"`
					DeathsZombieLord                                    float64 `json:"deaths_zombie_lord"`
					DeathsCryptUndeadMinikloon                          float64 `json:"deaths_crypt_undead_minikloon"`
					DeathsLonelySpider                                  float64 `json:"deaths_lonely_spider"`
					DeathsCellarSpider                                  float64 `json:"deaths_cellar_spider"`
					KillsSkeletonEmperor                                float64 `json:"kills_skeleton_emperor"`
					KillsWerewolf                                       float64 `json:"kills_werewolf"`
					KillsNightmare                                      float64 `json:"kills_nightmare"`
					KillsScarecrow                                      float64 `json:"kills_scarecrow"`
					DeathsZombieCommander                               float64 `json:"deaths_zombie_commander"`
					KillsWitherMiner                                    float64 `json:"kills_wither_miner"`
					DeathsMaxor                                         float64 `json:"deaths_maxor"`
					KillsWitherGuard                                    float64 `json:"kills_wither_guard"`
					DeathsWitherGuard                                   float64 `json:"deaths_wither_guard"`
					KillsCryptUndeadFriedrich                           float64 `json:"kills_crypt_undead_friedrich"`
					DeathsWitherMiner                                   float64 `json:"deaths_wither_miner"`
					KillsWatcherGiantBoulder                            float64 `json:"kills_watcher_giant_boulder"`
					DeathsSoulOfTheAlpha                                float64 `json:"deaths_soul_of_the_alpha"`
					KillsWatcherGiantBigfoot                            float64 `json:"kills_watcher_giant_bigfoot"`
					DeathsProfessorArcherGuardian                       float64 `json:"deaths_professor_archer_guardian"`
					KillsWatcherGiantDiamond                            float64 `json:"kills_watcher_giant_diamond"`
					KillsWatcherGiantLaser                              float64 `json:"kills_watcher_giant_laser"`
					DeathsCrushed                                       float64 `json:"deaths_crushed"`
					KillsNecronGuard                                    float64 `json:"kills_necron_guard"`
					DeathsSkeletonLord                                  float64 `json:"deaths_skeleton_lord"`
					DeathsWatcherGiantLaser                             float64 `json:"deaths_watcher_giant_laser"`
					DeathsWatcherGiantBigfoot                           float64 `json:"deaths_watcher_giant_bigfoot"`
					DeathsSuperArcher                                   float64 `json:"deaths_super_archer"`
					DeathsSuffocation                                   float64 `json:"deaths_suffocation"`
					KillsWiseDragon                                     float64 `json:"kills_wise_dragon"`
					KillsMayorJerryBlue                                 float64 `json:"kills_mayor_jerry_blue"`
					KillsMayorJerryGreen                                float64 `json:"kills_mayor_jerry_green"`
					KillsYoungDragon                                    float64 `json:"kills_young_dragon"`
					KillsMayorJerryPurple                               float64 `json:"kills_mayor_jerry_purple"`
					KillsStrongDragon                                   float64 `json:"kills_strong_dragon"`
					KillsMayorJerryGolden                               float64 `json:"kills_mayor_jerry_golden"`
					DeathsCavernsGhost                                  float64 `json:"deaths_caverns_ghost"`
					KillsIceWalker                                      float64 `json:"kills_ice_walker"`
					KillsGoblinWeaklingMelee                            float64 `json:"kills_goblin_weakling_melee"`
					KillsGoblinKnifeThrower                             float64 `json:"kills_goblin_knife_thrower"`
					KillsCavernsGhost                                   float64 `json:"kills_caverns_ghost"`
					KillsGoblin                                         float64 `json:"kills_goblin"`
					KillsTreasureHoarder                                float64 `json:"kills_treasure_hoarder"`
					KillsPowderGhast                                    float64 `json:"kills_powder_ghast"`
					KillsGoblinWeaklingBow                              float64 `json:"kills_goblin_weakling_bow"`
					KillsCrystalSentry                                  float64 `json:"kills_crystal_sentry"`
					KillsGoblinCreepertamer                             float64 `json:"kills_goblin_creepertamer"`
					KillsGoblinCreeper                                  float64 `json:"kills_goblin_creeper"`
					KillsGoblinBattler                                  float64 `json:"kills_goblin_battler"`
					KillsGoblinGolem                                    float64 `json:"kills_goblin_golem"`
					KillsGoblinMurderlover                              float64 `json:"kills_goblin_murderlover"`
					KillsPhantomFisherman                               float64 `json:"kills_phantom_fisherman"`
					KillsYeti                                           float64 `json:"kills_yeti"`
					KillsScarfMage                                      float64 `json:"kills_scarf_mage"`
					MythosBurrowsDugNextEPIC                            float64 `json:"mythos_burrows_dug_next_EPIC"`
					MythosBurrowsDugCombatEPIC                          float64 `json:"mythos_burrows_dug_combat_EPIC"`
					MythosBurrowsDugTreasureEPIC                        float64 `json:"mythos_burrows_dug_treasure_EPIC"`
					MythosBurrowsChainsCompleteEPIC                     float64 `json:"mythos_burrows_chains_complete_EPIC"`
					KillsMinosChampion                                  float64 `json:"kills_minos_champion"`
					MythosBurrowsDugNextLEGENDARY                       float64 `json:"mythos_burrows_dug_next_LEGENDARY"`
					MythosBurrowsDugCombatLEGENDARY                     float64 `json:"mythos_burrows_dug_combat_LEGENDARY"`
					MythosBurrowsDugTreasureLEGENDARY                   float64 `json:"mythos_burrows_dug_treasure_LEGENDARY"`
					MythosBurrowsChainsCompleteLEGENDARY                float64 `json:"mythos_burrows_chains_complete_LEGENDARY"`
					DeathsMinosChampion                                 float64 `json:"deaths_minos_champion"`
					DeathsMinosInquisitor                               float64 `json:"deaths_minos_inquisitor"`
					DeathsSiameseLynx                                   float64 `json:"deaths_siamese_lynx"`
					KillsForestIslandBat                                float64 `json:"kills_forest_island_bat"`
					KillsArachne                                        float64 `json:"kills_arachne"`
					KillsArachneBrood                                   float64 `json:"kills_arachne_brood"`
					DeathsArachneBrood                                  float64 `json:"deaths_arachne_brood"`
					KillsArachneKeeper                                  float64 `json:"kills_arachne_keeper"`
					DeathsArachne                                       float64 `json:"deaths_arachne"`
					DeathsWeaverSpider                                  float64 `json:"deaths_weaver_spider"`
					KillsMasterCryptTankZombie                          float64 `json:"kills_master_crypt_tank_zombie"`
					KillsMasterZombieGrunt                              float64 `json:"kills_master_zombie_grunt"`
					KillsMasterCryptDreadlord                           float64 `json:"kills_master_crypt_dreadlord"`
					KillsMasterCryptLurker                              float64 `json:"kills_master_crypt_lurker"`
					KillsMasterCryptSouleater                           float64 `json:"kills_master_crypt_souleater"`
					KillsMasterScaredSkeleton                           float64 `json:"kills_master_scared_skeleton"`
					KillsMasterSkeletonSoldier                          float64 `json:"kills_master_skeleton_soldier"`
					KillsMasterSkeletonGrunt                            float64 `json:"kills_master_skeleton_grunt"`
					KillsMasterLostAdventurer                           float64 `json:"kills_master_lost_adventurer"`
					KillsMasterDungeonRespawningSkeleton                float64 `json:"kills_master_dungeon_respawning_skeleton"`
					KillsMasterWatcherSummonUndead                      float64 `json:"kills_master_watcher_summon_undead"`
					DeathsMasterWatcherSummonUndead                     float64 `json:"deaths_master_watcher_summon_undead"`
					KillsMasterDiamondGuy                               float64 `json:"kills_master_diamond_guy"`
					KillsMasterCellarSpider                             float64 `json:"kills_master_cellar_spider"`
					KillsMasterBonzoSummonUndead                        float64 `json:"kills_master_bonzo_summon_undead"`
					DeathsMasterBonzoSummonUndead                       float64 `json:"deaths_master_bonzo_summon_undead"`
					DeathsMasterBonzo                                   float64 `json:"deaths_master_bonzo"`
					KillsMushroomCow                                    float64 `json:"kills_mushroom_cow"`
					KillsTrapperRabbit                                  float64 `json:"kills_trapper_rabbit"`
					KillsTrapperChicken                                 float64 `json:"kills_trapper_chicken"`
					KillsTrapperSheep                                   float64 `json:"kills_trapper_sheep"`
					KillsLavaPigman                                     float64 `json:"kills_lava_pigman"`
					KillsLavaBlaze                                      float64 `json:"kills_lava_blaze"`
					KillsSludge                                         float64 `json:"kills_sludge"`
					KillsVoidlingFanatic                                float64 `json:"kills_voidling_fanatic"`
					KillsVoidlingExtremist                              float64 `json:"kills_voidling_extremist"`
					KillsMasterCryptUndeadPieter                        float64 `json:"kills_master_crypt_undead_pieter"`
					KillsMasterCryptUndead                              float64 `json:"kills_master_crypt_undead"`
					KillsMasterSkeletonMaster                           float64 `json:"kills_master_skeleton_master"`
					DeathsMasterScarfWarrior                            float64 `json:"deaths_master_scarf_warrior"`
					DeathsMasterScarfMage                               float64 `json:"deaths_master_scarf_mage"`
					DeathsMasterScarf                                   float64 `json:"deaths_master_scarf"`
					KillsMasterParasite                                 float64 `json:"kills_master_parasite"`
					DeathsMasterSkeletonSoldier                         float64 `json:"deaths_master_skeleton_soldier"`
					KillsMasterCryptUndeadFriedrich                     float64 `json:"kills_master_crypt_undead_friedrich"`
					DeathsMasterLostAdventurer                          float64 `json:"deaths_master_lost_adventurer"`
					DeathsMasterShadowAssassin                          float64 `json:"deaths_master_shadow_assassin"`
					KillsMasterZombieSoldier                            float64 `json:"kills_master_zombie_soldier"`
					DeathsMasterProfessorMageGuardian                   float64 `json:"deaths_master_professor_mage_guardian"`
					KillsMasterSkeletor                                 float64 `json:"kills_master_skeletor"`
					KillsMasterProfessorGuardianSummon                  float64 `json:"kills_master_professor_guardian_summon"`
					KillsMasterZombieKnight                             float64 `json:"kills_master_zombie_knight"`
					KillsMasterSniperSkeleton                           float64 `json:"kills_master_sniper_skeleton"`
					DeathsMasterWatcherBonzo                            float64 `json:"deaths_master_watcher_bonzo"`
					DeathsMasterProfessor                               float64 `json:"deaths_master_professor"`
					DeathsMasterProfessorGuardianSummon                 float64 `json:"deaths_master_professor_guardian_summon"`
					DeathsVoidlingExtremist                             float64 `json:"deaths_voidling_extremist"`
					KillsOldDragon                                      float64 `json:"kills_old_dragon"`
				} `json:"stats"`
				Objectives struct {
					CollectLog struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"collect_log"`
					TalkToGuide struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guide"`
					PublicIsland struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"public_island"`
					ExploreHub struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"explore_hub"`
					ExploreVillage struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"explore_village"`
					TalkToLibrarian struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_librarian"`
					TalkToFarmer struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_farmer"`
					TalkToBlacksmith struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_blacksmith"`
					TalkToLumberjack struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_lumberjack"`
					TalkToEventMaster struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_event_master"`
					TalkToAuctionMaster struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_auction_master"`
					TalkToBanker struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_banker"`
					TalkToFairy struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_fairy"`
					TalkToFisherman1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_fisherman_1"`
					TalkToCarpenter struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_carpenter"`
					TalkToArtist1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_artist_1"`
					PaintCanvas struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"paint_canvas"`
					TalkToPetCollector struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_pet_collector"`
					TalkToPetSitter struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_pet_sitter"`
					ChopTree struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"chop_tree"`
					TalkToLumberjack2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_lumberjack_2"`
					IncreaseForagingSkill struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"increase_foraging_skill"`
					WarpForagingIslands struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"warp_foraging_islands"`
					IncreaseForagingSkill5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"increase_foraging_skill_5"`
					TalkToGustave1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_gustave_1"`
					CollectBirchLogs struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"collect_birch_logs"`
					CollectDarkOakLogs struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"collect_dark_oak_logs"`
					TalkToCharlie struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_charlie"`
					TalkToCharlie2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_charlie_2"`
					DepositCoins struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"deposit_coins"`
					CraftWorkbench struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"craft_workbench"`
					MineCoal struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"mine_coal"`
					TalkToLazyMiner struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_lazy_miner"`
					IncreaseMiningSkill5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"increase_mining_skill_5"`
					TalkToTelekinesisApplier struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_telekinesis_applier"`
					FindPickaxe struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"find_pickaxe"`
					CollectIngots struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						IronIngot   bool   `json:"IRON_INGOT"`
						GoldIngot   bool   `json:"GOLD_INGOT"`
					} `json:"collect_ingots"`
					WarpDeepCaverns struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"warp_deep_caverns"`
					TalkToLapisMiner struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_lapis_miner"`
					TalkToLiftOperator struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_lift_operator"`
					ReachLapisQuarry struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"reach_lapis_quarry"`
					CollectLapis struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						INKSACK4    bool   `json:"INK_SACK:4"`
					} `json:"collect_lapis"`
					ReachPigmensDen struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"reach_pigmens_den"`
					CollectRedstone struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						Redstone    bool   `json:"REDSTONE"`
					} `json:"collect_redstone"`
					TalkToBlacksmith2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_blacksmith_2"`
					TalkToFarmhand1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_farmhand_1"`
					IncreaseFarmingSkill5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"increase_farming_skill_5"`
					WarpMushroomDesert struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"warp_mushroom_desert"`
					KillDangerMobs struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"kill_danger_mobs"`
					TalkToBartender struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_bartender"`
					IncreaseCombatSkill struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"increase_combat_skill"`
					WarpSpidersDen struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"warp_spiders_den"`
					IncreaseCombatSkill5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"increase_combat_skill_5"`
					TalkToRick struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_rick"`
					WarpTheEnd struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"warp_the_end"`
					WarpBlazingFortress struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"warp_blazing_fortress"`
					IncreaseMiningSkill struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"increase_mining_skill"`
					ReforgeItem struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"reforge_item"`
					WarpGoldMine struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"warp_gold_mine"`
					ReachSlimehill struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"reach_slimehill"`
					TalkToGuber1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guber_1"`
					TalkToEndDealer struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_end_dealer"`
					CollectEndStone struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						EnderStone  bool   `json:"ENDER_STONE"`
					} `json:"collect_end_stone"`
					ReachDragonsNest struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"reach_dragons_nest"`
					FightDragon struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fight_dragon"`
					TalkToHaymitch struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_haymitch"`
					CollectEmerald struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						Emerald     bool   `json:"EMERALD"`
					} `json:"collect_emerald"`
					ReachDiamondReserve struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"reach_diamond_reserve"`
					TalkToArtist2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_artist_2"`
					GiveFairySouls struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"give_fairy_souls"`
					CollectDiamond struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						Diamond     bool   `json:"DIAMOND"`
					} `json:"collect_diamond"`
					ReachObsidianSanctuary struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"reach_obsidian_sanctuary"`
					CollectObsidian struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						Obsidian    bool   `json:"OBSIDIAN"`
					} `json:"collect_obsidian"`
					CollectWoolCarpenter struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"collect_wool_carpenter"`
					CompleteTheEndRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_end_race_1"`
					CompleteTheWoodsRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_woods_race_1"`
					TalkToGustave2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_gustave_2"`
					CompleteTheWoodsRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_woods_race_2"`
					TalkToGustave3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_gustave_3"`
					CompleteTheWoodsRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_woods_race_3"`
					CollectWheat struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"collect_wheat"`
					TalkToFarmer2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_farmer_2"`
					IncreaseFarmingSkill struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"increase_farming_skill"`
					WarpBarnIsland struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"warp_barn_island"`
					CraftWheatMinion struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"craft_wheat_minion"`
					TalkToFarmhand2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_farmhand_2"`
					CollectFarmingResources2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						INKSACK3    bool   `json:"INK_SACK:3"`
						Cactus      bool   `json:"CACTUS"`
						SugarCane   bool   `json:"SUGAR_CANE"`
					} `json:"collect_farming_resources_2"`
					CollectFarmAnimalResources2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						Rabbit      bool   `json:"RABBIT"`
						Mutton      bool   `json:"MUTTON"`
					} `json:"collect_farm_animal_resources_2"`
					TalkToElle struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_elle"`
					EnchantItem struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"enchant_item"`
					TalkToMelody struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_melody"`
					TalkToFrosty struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_frosty"`
					TalkToGulliver1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_gulliver_1"`
					CompleteTheChickenRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_chicken_race_1"`
					TalkToGulliver2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_gulliver_2"`
					CompleteTheChickenRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_chicken_race_2"`
					TalkToGuildford1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_1"`
					CompleteTheGiantMushroomAnythingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_with_return_race_1"`
					CompleteTheGiantMushroomNoPearlsWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_with_return_race_1"`
					CompleteTheGiantMushroomNoAbilitiesWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_with_return_race_1"`
					CompleteTheGiantMushroomNothingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_with_return_race_1"`
					CompleteThePrecursorRuinsAnythingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_with_return_race_1"`
					CompleteThePrecursorRuinsNoPearlsWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_with_return_race_1"`
					CompleteThePrecursorRuinsNoAbilitiesWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_with_return_race_1"`
					CompleteThePrecursorRuinsNothingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_with_return_race_1"`
					CompleteTheCrystalCoreAnythingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_with_return_race_1"`
					CompleteTheCrystalCoreNoPearlsWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_with_return_race_1"`
					CompleteTheCrystalCoreNoAbilitiesWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_with_return_race_1"`
					CompleteTheCrystalCoreNothingWithReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_with_return_race_1"`
					CompleteTheGiantMushroomAnythingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_1"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_1"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_1"`
					CompleteTheGiantMushroomNothingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_1"`
					CompleteThePrecursorRuinsAnythingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_1"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_1"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_1"`
					CompleteThePrecursorRuinsNothingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_1"`
					CompleteTheCrystalCoreAnythingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_1"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_1"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_1"`
					CompleteTheCrystalCoreNothingNoReturnRace1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_1"`
					GiveRickIngots struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"give_rick_ingots"`
					CollectFarmingResources struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						PotatoItem  bool   `json:"POTATO_ITEM"`
						CarrotItem  bool   `json:"CARROT_ITEM"`
						Pumpkin     bool   `json:"PUMPKIN"`
						Melon       bool   `json:"MELON"`
					} `json:"collect_farming_resources"`
					CollectFarmAnimalResources struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						RawChicken  bool   `json:"RAW_CHICKEN"`
						Leather     bool   `json:"LEATHER"`
						Pork        bool   `json:"PORK"`
					} `json:"collect_farm_animal_resources"`
					TalkToGuildfordCrystalCoreAnythingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_1"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_1"`
					CompleteTheCrystalCoreAnythingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_2"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_2"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_1"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_2"`
					TalkToGuildfordCrystalCoreAnythingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_2"`
					CompleteTheCrystalCoreAnythingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_3"`
					TalkToGuildfordCrystalCoreAnythingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_3"`
					CompleteTheCrystalCoreAnythingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_anything_no_return_race_4"`
					TalkToGuildfordCrystalCoreAnythingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_anything_no_return_4"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_2"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_3"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_3"`
					CompleteTheCrystalCoreNoPearlsNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_pearls_no_return_race_4"`
					TalkToGuildfordCrystalCoreNoPearlsNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_pearls_no_return_4"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_2"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_3"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_3"`
					CompleteTheCrystalCoreNoAbilitiesNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_no_abilities_no_return_race_4"`
					TalkToGuildfordCrystalCoreNoAbilitiesNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_no_abilities_no_return_4"`
					TalkToGuildfordCrystalCoreNothingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_1"`
					CompleteTheCrystalCoreNothingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_2"`
					TalkToGuildfordCrystalCoreNothingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_2"`
					CompleteTheCrystalCoreNothingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_3"`
					TalkToGuildfordCrystalCoreNothingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_3"`
					CompleteTheCrystalCoreNothingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_crystal_core_nothing_no_return_race_4"`
					TalkToGuildfordCrystalCoreNothingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_crystal_core_nothing_no_return_4"`
					TalkToGuildfordGiantMushroomAnythingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_1"`
					CompleteTheGiantMushroomAnythingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_2"`
					TalkToGuildfordGiantMushroomAnythingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_2"`
					CompleteTheGiantMushroomAnythingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_3"`
					TalkToGuildfordGiantMushroomAnythingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_3"`
					CompleteTheGiantMushroomAnythingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_anything_no_return_race_4"`
					TalkToGuildfordGiantMushroomAnythingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_anything_no_return_4"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_1"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_2"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_2"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_3"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_3"`
					CompleteTheGiantMushroomNoPearlsNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_pearls_no_return_race_4"`
					TalkToGuildfordGiantMushroomNoPearlsNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_pearls_no_return_4"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_1"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_2"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_2"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_3"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_3"`
					CompleteTheGiantMushroomNoAbilitiesNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_no_abilities_no_return_race_4"`
					TalkToGuildfordGiantMushroomNoAbilitiesNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_no_abilities_no_return_4"`
					TalkToGuildfordGiantMushroomNothingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_1"`
					CompleteTheGiantMushroomNothingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_2"`
					TalkToGuildfordGiantMushroomNothingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_2"`
					CompleteTheGiantMushroomNothingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_3"`
					TalkToGuildfordGiantMushroomNothingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_3"`
					CompleteTheGiantMushroomNothingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_giant_mushroom_nothing_no_return_race_4"`
					TalkToGuildfordGiantMushroomNothingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_giant_mushroom_nothing_no_return_4"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_1"`
					CompleteThePrecursorRuinsAnythingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_2"`
					CompleteThePrecursorRuinsAnythingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_3"`
					CompleteThePrecursorRuinsAnythingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_anything_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsAnythingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_anything_no_return_4"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_1"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_2"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_3"`
					CompleteThePrecursorRuinsNoPearlsNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_pearls_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsNoPearlsNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_pearls_no_return_4"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_1"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_2"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_3"`
					CompleteThePrecursorRuinsNoAbilitiesNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_no_abilities_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsNoAbilitiesNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_no_abilities_no_return_4"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn1 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_1"`
					CompleteThePrecursorRuinsNothingNoReturnRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_2"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_2"`
					CompleteThePrecursorRuinsNothingNoReturnRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_3"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_3"`
					CompleteThePrecursorRuinsNothingNoReturnRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_precursor_ruins_nothing_no_return_race_4"`
					TalkToGuildfordPrecursorRuinsNothingNoReturn4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guildford_precursor_ruins_nothing_no_return_4"`
					TalkToGulliver3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_gulliver_3"`
					CompleteTheChickenRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_chicken_race_3"`
					GivePickaxeLapisMiner struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"give_pickaxe_lapis_miner"`
					TalkToGulliver4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_gulliver_4"`
					CompleteTheChickenRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"complete_the_chicken_race_4"`
					TalkToGuber2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guber_2"`
					CompleteTheEndRace2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_end_race_2"`
					TalkToGuber3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guber_3"`
					CompleteTheEndRace3 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_end_race_3"`
					TalkToGuber4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guber_4"`
					CompleteTheEndRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_end_race_4"`
					TalkToGuber5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_guber_5"`
					CraftWoodPickaxe struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"craft_wood_pickaxe"`
					TalkToGustave4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_gustave_4"`
					CompleteTheWoodsRace4 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"complete_the_woods_race_4"`
					TalkToGustave5 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_gustave_5"`
					CollectSpider struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						String      bool   `json:"STRING"`
						SpiderEye   bool   `json:"SPIDER_EYE"`
					} `json:"collect_spider"`
					CollectNetherResources struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
						BlazeRod    bool   `json:"BLAZE_ROD"`
						NetherStalk bool   `json:"NETHER_STALK"`
					} `json:"collect_nether_resources"`
					CollectNetherResources2 struct {
						Status        string `json:"status"`
						Progress      int    `json:"progress"`
						CompletedAt   int64  `json:"completed_at"`
						MagmaCream    bool   `json:"MAGMA_CREAM"`
						GlowstoneDust bool   `json:"GLOWSTONE_DUST"`
						Quartz        bool   `json:"QUARTZ"`
					} `json:"collect_nether_resources_2"`
					TalkToRhys struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_rhys"`
					IncreaseMining12 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"increase_mining_12"`
					HotmGiveMaterials struct {
						Status            string `json:"status"`
						Progress          int    `json:"progress"`
						CompletedAt       int64  `json:"completed_at"`
						Started           bool   `json:"started"`
						EnchantedRedstone int    `json:"ENCHANTED_REDSTONE"`
						EnchantedIron     int    `json:"ENCHANTED_IRON"`
					} `json:"hotm_give_materials"`
					Fetchur240 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-24-0"`
					Fetchur260 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-26-0"`
					Fetchur290 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-29-0"`
					Fetchur300 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"fetchur-30-0"`
					TalkToArchaeologist struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_archaeologist"`
					TalkToShaggy struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_shaggy"`
					FindRelics struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"find_relics"`
					TalkToArchaeologist2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_archaeologist_2"`
					FindUberRelics struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"find_uber_relics"`
					TalkToShaggy2 struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_shaggy_2"`
					TalkToGwendolyn struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_gwendolyn"`
					TalkToBraum struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"talk_to_braum"`
					VisitGreaterMines struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int64  `json:"completed_at"`
					} `json:"visit_greater_mines"`
					TalkToTheGoblinKing struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"talk_to_the_goblin_king"`
					KillAutomatons struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"kill_automatons"`
					EnterDivanMines struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"enter_divan_mines"`
					FindAJungleKey struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"find_a_jungle_key"`
					MineRuby struct {
						Status      string `json:"status"`
						Progress    int    `json:"progress"`
						CompletedAt int    `json:"completed_at"`
					} `json:"mine_ruby"`
				} `json:"objectives"`
				Tutorial []string `json:"tutorial"`
				Quests   struct {
					CollectLog struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"collect_log"`
					ExploreHub struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"explore_hub"`
					ExploreVillage struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"explore_village"`
					TalkToLibrarian struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_librarian"`
					TalkToFarmer struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_farmer"`
					TalkToBlacksmith struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_blacksmith"`
					TalkToLumberjack struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_lumberjack"`
					TalkToAuctionMaster struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_auction_master"`
					TalkToBanker struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_banker"`
					TalkToCarpenter struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_carpenter"`
					TalkToArtist1 struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_artist_1"`
					IncreaseForagingSkill5 struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"increase_foraging_skill_5"`
					TalkToGustave1 struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_gustave_1"`
					TalkToLazyMiner struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_lazy_miner"`
					IncreaseMiningSkill5 struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"increase_mining_skill_5"`
					TalkToLapisMiner struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_lapis_miner"`
					TalkToFarmhand1 struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_farmhand_1"`
					IncreaseFarmingSkill5 struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"increase_farming_skill_5"`
					KillDangerMobs struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"kill_danger_mobs"`
					IncreaseCombatSkill5 struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"increase_combat_skill_5"`
					TalkToRick struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_rick"`
					ReforgeItem struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"reforge_item"`
					TalkToGuber1 struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_guber_1"`
					TalkToEndDealer struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_end_dealer"`
					TalkToGulliver1 struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_gulliver_1"`
					TalkToGuildford1 struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_guildford_1"`
					TalkToRhys struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_rhys"`
					TalkToArchaeologist struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_archaeologist"`
					TalkToGwendolyn struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int64  `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_gwendolyn"`
					TalkToTheGoblinKing struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"talk_to_the_goblin_king"`
					KillAutomatons struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"kill_automatons"`
					EnterDivanMines struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"enter_divan_mines"`
					FindAJungleKey struct {
						Status        string `json:"status"`
						ActivatedAt   int64  `json:"activated_at"`
						ActivatedAtSb int    `json:"activated_at_sb"`
						CompletedAt   int    `json:"completed_at"`
						CompletedAtSb int    `json:"completed_at_sb"`
					} `json:"find_a_jungle_key"`
				} `json:"quests"`
				LastDeath                     int      `json:"last_death"`
				CraftedGenerators             []string `json:"crafted_generators"`
				VisitedZones                  []string `json:"visited_zones"`
				FairySoulsCollected           int      `json:"fairy_souls_collected"`
				FairySouls                    int      `json:"fairy_souls"`
				FairyExchanges                int      `json:"fairy_exchanges"`
				FishingTreasureCaught         int      `json:"fishing_treasure_caught"`
				DeathCount                    int      `json:"death_count"`
				AchievementSpawnedIslandTypes []string `json:"achievement_spawned_island_types"`
				SlayerQuest                   struct {
					Type            string `json:"type"`
					Tier            int    `json:"tier"`
					StartTimestamp  int64  `json:"start_timestamp"`
					CompletionState int    `json:"completion_state"`
					CombatXp        int    `json:"combat_xp"`
					RecentMobKills  []struct {
						Xp        float64 `json:"xp"`
						Timestamp int64   `json:"timestamp"`
					} `json:"recent_mob_kills"`
					LastKilledMobIsland string `json:"last_killed_mob_island"`
				} `json:"slayer_quest"`
				SlayerBosses struct {
					Spider struct {
						ClaimedLevels struct {
							Level1 bool `json:"level_1"`
							Level2 bool `json:"level_2"`
							Level3 bool `json:"level_3"`
							Level4 bool `json:"level_4"`
							Level5 bool `json:"level_5"`
							Level6 bool `json:"level_6"`
							Level7 bool `json:"level_7"`
						} `json:"claimed_levels"`
						BossKillsTier0 int `json:"boss_kills_tier_0"`
						Xp             int `json:"xp"`
						BossKillsTier1 int `json:"boss_kills_tier_1"`
						BossKillsTier2 int `json:"boss_kills_tier_2"`
						BossKillsTier3 int `json:"boss_kills_tier_3"`
					} `json:"spider"`
					Zombie struct {
						ClaimedLevels struct {
							Level1        bool `json:"level_1"`
							Level2        bool `json:"level_2"`
							Level3        bool `json:"level_3"`
							Level4        bool `json:"level_4"`
							Level5        bool `json:"level_5"`
							Level6        bool `json:"level_6"`
							Level7        bool `json:"level_7"`
							Level7Special bool `json:"level_7_special"`
						} `json:"claimed_levels"`
						BossKillsTier0 int `json:"boss_kills_tier_0"`
						Xp             int `json:"xp"`
						BossKillsTier1 int `json:"boss_kills_tier_1"`
						BossKillsTier2 int `json:"boss_kills_tier_2"`
						BossKillsTier3 int `json:"boss_kills_tier_3"`
						BossKillsTier4 int `json:"boss_kills_tier_4"`
					} `json:"zombie"`
					Wolf struct {
						ClaimedLevels struct {
							Level1 bool `json:"level_1"`
							Level2 bool `json:"level_2"`
							Level3 bool `json:"level_3"`
							Level4 bool `json:"level_4"`
							Level5 bool `json:"level_5"`
							Level6 bool `json:"level_6"`
							Level7 bool `json:"level_7"`
						} `json:"claimed_levels"`
						BossKillsTier0 int `json:"boss_kills_tier_0"`
						Xp             int `json:"xp"`
						BossKillsTier1 int `json:"boss_kills_tier_1"`
						BossKillsTier2 int `json:"boss_kills_tier_2"`
						BossKillsTier3 int `json:"boss_kills_tier_3"`
					} `json:"wolf"`
					Enderman struct {
						ClaimedLevels struct {
							Level1 bool `json:"level_1"`
							Level2 bool `json:"level_2"`
							Level3 bool `json:"level_3"`
						} `json:"claimed_levels"`
						BossKillsTier0 int `json:"boss_kills_tier_0"`
						Xp             int `json:"xp"`
						BossKillsTier1 int `json:"boss_kills_tier_1"`
					} `json:"enderman"`
				} `json:"slayer_bosses"`
				Dungeons struct {
					DungeonTypes struct {
						Catacombs struct {
							TimesPlayed struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"times_played"`
							Experience float64 `json:"experience"`
							BestScore  struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"best_score"`
							MobsKilled struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"mobs_killed"`
							MostMobsKilled struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_mobs_killed"`
							MostDamageBerserk struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_damage_berserk"`
							MostHealing struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_healing"`
							TierCompletions struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"tier_completions"`
							FastestTime struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"fastest_time"`
							BestRuns struct {
								Num0 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
									AllyHealing      float64  `json:"ally_healing,omitempty"`
								} `json:"0"`
								Num1 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
									AllyHealing      float64  `json:"ally_healing,omitempty"`
								} `json:"1"`
								Num2 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
									AllyHealing      float64  `json:"ally_healing,omitempty"`
								} `json:"2"`
								Num3 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
									AllyHealing      float64  `json:"ally_healing,omitempty"`
								} `json:"3"`
								Num4 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
									AllyHealing      float64  `json:"ally_healing,omitempty"`
								} `json:"4"`
								Num5 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
									AllyHealing      float64  `json:"ally_healing"`
								} `json:"5"`
								Num6 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
									AllyHealing      float64  `json:"ally_healing"`
								} `json:"6"`
								Num7 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
									AllyHealing      float64  `json:"ally_healing"`
								} `json:"7"`
							} `json:"best_runs"`
							WatcherKills struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"watcher_kills"`
							HighestTierCompleted int `json:"highest_tier_completed"`
							FastestTimeS         struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"fastest_time_s"`
							MostDamageArcher struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_damage_archer"`
							MostDamageHealer struct {
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
							} `json:"most_damage_healer"`
							FastestTimeSPlus struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"fastest_time_s_plus"`
							MostDamageTank struct {
								Num4 float64 `json:"4"`
							} `json:"most_damage_tank"`
							MostDamageMage struct {
								Num0 float64 `json:"0"`
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
								Num4 float64 `json:"4"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"most_damage_mage"`
							MilestoneCompletions struct {
								Num3 float64 `json:"3"`
								Num5 float64 `json:"5"`
								Num6 float64 `json:"6"`
								Num7 float64 `json:"7"`
							} `json:"milestone_completions"`
						} `json:"catacombs"`
						MasterCatacombs struct {
							TierCompletions struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
							} `json:"tier_completions"`
							MilestoneCompletions struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
							} `json:"milestone_completions"`
							HighestTierCompleted int `json:"highest_tier_completed"`
							FastestTime          struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
							} `json:"fastest_time"`
							BestRuns struct {
								Num1 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
									AllyHealing      float64  `json:"ally_healing,omitempty"`
								} `json:"1"`
								Num2 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
								} `json:"2"`
								Num3 []struct {
									Timestamp        int64    `json:"timestamp"`
									ScoreExploration int      `json:"score_exploration"`
									ScoreSpeed       int      `json:"score_speed"`
									ScoreSkill       int      `json:"score_skill"`
									ScoreBonus       int      `json:"score_bonus"`
									DungeonClass     string   `json:"dungeon_class"`
									Teammates        []string `json:"teammates"`
									ElapsedTime      int      `json:"elapsed_time"`
									DamageDealt      float64  `json:"damage_dealt"`
									Deaths           int      `json:"deaths"`
									MobsKilled       int      `json:"mobs_killed"`
									SecretsFound     int      `json:"secrets_found"`
									DamageMitigated  float64  `json:"damage_mitigated"`
									AllyHealing      float64  `json:"ally_healing"`
								} `json:"3"`
							} `json:"best_runs"`
							BestScore struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
							} `json:"best_score"`
							MobsKilled struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
							} `json:"mobs_killed"`
							MostMobsKilled struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
							} `json:"most_mobs_killed"`
							MostDamageBerserk struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
							} `json:"most_damage_berserk"`
							MostHealing struct {
								Num1 float64 `json:"1"`
								Num3 float64 `json:"3"`
							} `json:"most_healing"`
							FastestTimeS struct {
								Num1 float64 `json:"1"`
								Num2 float64 `json:"2"`
								Num3 float64 `json:"3"`
							} `json:"fastest_time_s"`
							MostDamageArcher struct {
								Num1 float64 `json:"1"`
								Num3 float64 `json:"3"`
							} `json:"most_damage_archer"`
							FastestTimeSPlus struct {
								Num1 float64 `json:"1"`
							} `json:"fastest_time_s_plus"`
							MostDamageHealer struct {
								Num3 float64 `json:"3"`
							} `json:"most_damage_healer"`
						} `json:"master_catacombs"`
					} `json:"dungeon_types"`
					PlayerClasses struct {
						Healer struct {
							Experience float64 `json:"experience"`
						} `json:"healer"`
						Mage struct {
							Experience float64 `json:"experience"`
						} `json:"mage"`
						Berserk struct {
							Experience float64 `json:"experience"`
						} `json:"berserk"`
						Archer struct {
							Experience float64 `json:"experience"`
						} `json:"archer"`
						Tank struct {
							Experience float64 `json:"experience"`
						} `json:"tank"`
					} `json:"player_classes"`
					DungeonJournal struct {
						JournalEntries struct {
							KaryllesDiary      []int         `json:"karylles_diary"`
							TheStudy           []int         `json:"the_study"`
							ExpeditionVolume1  []int         `json:"expedition_volume_1"`
							UncannyRemains     []int         `json:"uncanny_remains"`
							ExpeditionVolume2  []int         `json:"expedition_volume_2"`
							GrimAdversity      []int         `json:"grim_adversity"`
							ExpeditionVolume3  []int         `json:"expedition_volume_3"`
							ExpeditionVolume4  []int         `json:"expedition_volume_4"`
							NecronsMagicScroll []interface{} `json:"necrons_magic_scroll"`
							TheWalls           []int         `json:"the_walls"`
						} `json:"journal_entries"`
					} `json:"dungeon_journal"`
					DungeonsBlahBlah     []string `json:"dungeons_blah_blah"`
					SelectedDungeonClass string   `json:"selected_dungeon_class"`
				} `json:"dungeons"`
				Griffin struct {
					Burrows []struct {
						Ts    int64 `json:"ts"`
						X     int   `json:"x"`
						Y     int   `json:"y"`
						Z     int   `json:"z"`
						Type  int   `json:"type"`
						Tier  int   `json:"tier"`
						Chain int   `json:"chain"`
					} `json:"burrows"`
				} `json:"griffin"`
				Jacob2 struct {
					MedalsInv struct {
						Bronze int `json:"bronze"`
						Silver int `json:"silver"`
						Gold   int `json:"gold"`
					} `json:"medals_inv"`
					Perks struct {
						DoubleDrops     int `json:"double_drops"`
						FarmingLevelCap int `json:"farming_level_cap"`
					} `json:"perks"`
					Contests struct {
						One00108SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"100:10_8:SUGAR_CANE"`
						One001014PUMPKIN struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"100:10_14:PUMPKIN"`
						One001020INKSACK3 struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"100:10_20:INK_SACK:3"`
						One001023SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"100:10_23:SUGAR_CANE"`
						One001029POTATOITEM struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"100:10_29:POTATO_ITEM"`
						One001224PUMPKIN struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"100:12_24:PUMPKIN"`
						One001230SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"100:12_30:SUGAR_CANE"`
						One0115PUMPKIN struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:1_5:PUMPKIN"`
						One0118SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:1_8:SUGAR_CANE"`
						One01111POTATOITEM struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:1_11:POTATO_ITEM"`
						One01123SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:1_23:SUGAR_CANE"`
						One01222SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:2_22:SUGAR_CANE"`
						One01327SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:3_27:SUGAR_CANE"`
						One0154SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:5_4:SUGAR_CANE"`
						One0157POTATOITEM struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:5_7:POTATO_ITEM"`
						One01510INKSACK3 struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:5_10:INK_SACK:3"`
						One01516PUMPKIN struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:5_16:PUMPKIN"`
						One01531POTATOITEM struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:5_31:POTATO_ITEM"`
						One01717POTATOITEM struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:7_17:POTATO_ITEM"`
						One011128PUMPKIN struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"101:11_28:PUMPKIN"`
						One02321SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"102:3_21:SUGAR_CANE"`
						One02423INKSACK3 struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"102:4_23:INK_SACK:3"`
						One02426PUMPKIN struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"102:4_26:PUMPKIN"`
						One02516PUMPKIN struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"102:5_16:PUMPKIN"`
						One02522SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"102:5_22:SUGAR_CANE"`
						One02525POTATOITEM struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"102:5_25:POTATO_ITEM"`
						One04417PUMPKIN struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"104:4_17:PUMPKIN"`
						One0454SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"104:5_4:SUGAR_CANE"`
						One04624POTATOITEM struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"104:6_24:POTATO_ITEM"`
						One04711SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"104:7_11:SUGAR_CANE"`
						One04714POTATOITEM struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"104:7_14:POTATO_ITEM"`
						One04114SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"104:11_4:SUGAR_CANE"`
						One0612POTATOITEM struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"106:1_2:POTATO_ITEM"`
						One07717POTATOITEM struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"107:7_17:POTATO_ITEM"`
						One09219NETHERSTALK struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"109:2_19:NETHER_STALK"`
						One09222NETHERSTALK struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"109:2_22:NETHER_STALK"`
						One09231NETHERSTALK struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"109:2_31:NETHER_STALK"`
						One09420SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"109:4_20:SUGAR_CANE"`
						One0954NETHERSTALK struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"109:5_4:NETHER_STALK"`
						One10711NETHERSTALK struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"110:7_11:NETHER_STALK"`
						One10828NETHERSTALK struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"110:8_28:NETHER_STALK"`
						One12126SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"112:1_26:SUGAR_CANE"`
						One12330SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"112:3_30:SUGAR_CANE"`
						One121017NETHERSTALK struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"112:10_17:NETHER_STALK"`
						One121026NETHERSTALK struct {
							Collected int `json:"collected"`
						} `json:"112:10_26:NETHER_STALK"`
						One14711NETHERSTALK struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"114:7_11:NETHER_STALK"`
						One1575SUGARCANE struct {
							Collected int `json:"collected"`
						} `json:"115:7_5:SUGAR_CANE"`
						One151224CARROTITEM struct {
							Collected int `json:"collected"`
						} `json:"115:12_24:CARROT_ITEM"`
						One16129SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"116:1_29:SUGAR_CANE"`
						One1624SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"116:2_4:SUGAR_CANE"`
						One18111NETHERSTALK struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"118:1_11:NETHER_STALK"`
						One18114SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"118:1_14:SUGAR_CANE"`
						One1821PUMPKIN struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"118:2_1:PUMPKIN"`
						One1839SUGARCANE struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"118:3_9:SUGAR_CANE"`
						One18312PUMPKIN struct {
							Collected           int  `json:"collected"`
							ClaimedRewards      bool `json:"claimed_rewards"`
							ClaimedPosition     int  `json:"claimed_position"`
							ClaimedParticipants int  `json:"claimed_participants"`
						} `json:"118:3_12:PUMPKIN"`
						One18810PUMPKIN struct {
							Collected int `json:"collected"`
						} `json:"118:8_10:PUMPKIN"`
						One32129SUGARCANE struct {
							Collected int `json:"collected"`
						} `json:"132:12_9:SUGAR_CANE"`
						One33924NETHERSTALK struct {
							Collected int `json:"collected"`
						} `json:"133:9_24:NETHER_STALK"`
						One52615SUGARCANE struct {
							Collected int `json:"collected"`
						} `json:"152:6_15:SUGAR_CANE"`
					} `json:"contests"`
					Talked       bool     `json:"talked"`
					UniqueGolds2 []string `json:"unique_golds2"`
				} `json:"jacob2"`
				Experimentation struct {
					Simon struct {
						LastAttempt int64 `json:"last_attempt"`
						Attempts0   int   `json:"attempts_0"`
						BonusClicks int   `json:"bonus_clicks"`
						LastClaimed int64 `json:"last_claimed"`
						Claims0     int   `json:"claims_0"`
						BestScore0  int   `json:"best_score_0"`
						Attempts1   int   `json:"attempts_1"`
						Claims1     int   `json:"claims_1"`
						BestScore1  int   `json:"best_score_1"`
						Attempts2   int   `json:"attempts_2"`
						Claims2     int   `json:"claims_2"`
						BestScore2  int   `json:"best_score_2"`
						Attempts3   int   `json:"attempts_3"`
						Claims3     int   `json:"claims_3"`
						BestScore3  int   `json:"best_score_3"`
						Attempts5   int   `json:"attempts_5"`
						Claims5     int   `json:"claims_5"`
						BestScore5  int   `json:"best_score_5"`
					} `json:"simon"`
					Pairings struct {
						LastClaimed int64 `json:"last_claimed"`
						Claims1     int   `json:"claims_1"`
						LastAttempt int   `json:"last_attempt"`
						BestScore1  int   `json:"best_score_1"`
						Claims2     int   `json:"claims_2"`
						BestScore2  int   `json:"best_score_2"`
						Claims3     int   `json:"claims_3"`
						BestScore3  int   `json:"best_score_3"`
						Claims4     int   `json:"claims_4"`
						BestScore4  int   `json:"best_score_4"`
						Claims5     int   `json:"claims_5"`
						BestScore5  int   `json:"best_score_5"`
					} `json:"pairings"`
					Numbers struct {
						LastAttempt int64 `json:"last_attempt"`
						Attempts1   int   `json:"attempts_1"`
						BonusClicks int   `json:"bonus_clicks"`
						LastClaimed int64 `json:"last_claimed"`
						Claims1     int   `json:"claims_1"`
						BestScore1  int   `json:"best_score_1"`
						Attempts2   int   `json:"attempts_2"`
						Claims2     int   `json:"claims_2"`
						BestScore2  int   `json:"best_score_2"`
						Attempts3   int   `json:"attempts_3"`
						Claims3     int   `json:"claims_3"`
						BestScore3  int   `json:"best_score_3"`
					} `json:"numbers"`
					ClaimsResets          int   `json:"claims_resets"`
					ClaimsResetsTimestamp int64 `json:"claims_resets_timestamp"`
				} `json:"experimentation"`
				Perks struct {
					PermanentHealth     int `json:"permanent_health"`
					PermanentStrength   int `json:"permanent_strength"`
					PermanentSpeed      int `json:"permanent_speed"`
					CatacombsBossLuck   int `json:"catacombs_boss_luck"`
					CatacombsCritDamage int `json:"catacombs_crit_damage"`
					CatacombsLooting    int `json:"catacombs_looting"`
					CatacombsStrength   int `json:"catacombs_strength"`
					CatacombsHealth     int `json:"catacombs_health"`
					CatacombsDefense    int `json:"catacombs_defense"`
					PermanentDefense    int `json:"permanent_defense"`
					ForbiddenBlessing   int `json:"forbidden_blessing"`
				} `json:"perks"`
				HarpQuest struct {
					SelectedSong                          string  `json:"selected_song"`
					SelectedSongEpoch                     int64   `json:"selected_song_epoch"`
					SongHymnJoyBestCompletion             float64 `json:"song_hymn_joy_best_completion"`
					SongHymnJoyCompletions                int     `json:"song_hymn_joy_completions"`
					SongHymnJoyPerfectCompletions         int     `json:"song_hymn_joy_perfect_completions"`
					SongFrereJacquesBestCompletion        float64 `json:"song_frere_jacques_best_completion"`
					SongFrereJacquesCompletions           int     `json:"song_frere_jacques_completions"`
					SongFrereJacquesPerfectCompletions    int     `json:"song_frere_jacques_perfect_completions"`
					SongAmazingGraceBestCompletion        float64 `json:"song_amazing_grace_best_completion"`
					SongAmazingGraceCompletions           int     `json:"song_amazing_grace_completions"`
					SongAmazingGracePerfectCompletions    int     `json:"song_amazing_grace_perfect_completions"`
					SongBrahmsBestCompletion              float64 `json:"song_brahms_best_completion"`
					SongBrahmsCompletions                 int     `json:"song_brahms_completions"`
					SongHappyBirthdayBestCompletion       float64 `json:"song_happy_birthday_best_completion"`
					SongHappyBirthdayCompletions          int     `json:"song_happy_birthday_completions"`
					SongHappyBirthdayPerfectCompletions   int     `json:"song_happy_birthday_perfect_completions"`
					SongGreensleevesBestCompletion        float64 `json:"song_greensleeves_best_completion"`
					SongGreensleevesCompletions           int     `json:"song_greensleeves_completions"`
					SongGreensleevesPerfectCompletions    int     `json:"song_greensleeves_perfect_completions"`
					SongJeopardyBestCompletion            float64 `json:"song_jeopardy_best_completion"`
					SongJeopardyCompletions               int     `json:"song_jeopardy_completions"`
					SongJeopardyPerfectCompletions        int     `json:"song_jeopardy_perfect_completions"`
					SongMinuetCompletions                 int     `json:"song_minuet_completions"`
					SongBrahmsPerfectCompletions          int     `json:"song_brahms_perfect_completions"`
					SongMinuetBestCompletion              float64 `json:"song_minuet_best_completion"`
					SongMinuetPerfectCompletions          int     `json:"song_minuet_perfect_completions"`
					SongJoyWorldBestCompletion            float64 `json:"song_joy_world_best_completion"`
					SongJoyWorldCompletions               int     `json:"song_joy_world_completions"`
					SongJoyWorldPerfectCompletions        int     `json:"song_joy_world_perfect_completions"`
					SongPureImaginationBestCompletion     float64 `json:"song_pure_imagination_best_completion"`
					SongPureImaginationCompletions        int     `json:"song_pure_imagination_completions"`
					SongPureImaginationPerfectCompletions int     `json:"song_pure_imagination_perfect_completions"`
					SongVieEnRoseBestCompletion           float64 `json:"song_vie_en_rose_best_completion"`
					SongVieEnRoseCompletions              int     `json:"song_vie_en_rose_completions"`
				} `json:"harp_quest"`
				ActiveEffects []struct {
					Effect    string `json:"effect"`
					Level     int    `json:"level"`
					Modifiers []struct {
						Key string `json:"key"`
						Amp int    `json:"amp"`
					} `json:"modifiers"`
					TicksRemaining int  `json:"ticks_remaining"`
					Infinite       bool `json:"infinite"`
				} `json:"active_effects"`
				PausedEffects         []interface{} `json:"paused_effects"`
				DisabledPotionEffects []interface{} `json:"disabled_potion_effects"`
				VisitedModes          []string      `json:"visited_modes"`
				TempStatBuffs         []struct {
					Stat     int    `json:"stat"`
					Key      string `json:"key"`
					Amount   int    `json:"amount"`
					ExpireAt int64  `json:"expire_at"`
				} `json:"temp_stat_buffs"`
				MiningCore struct {
					Nodes struct {
						MiningSpeed      int `json:"mining_speed"`
						MiningFortune    int `json:"mining_fortune"`
						TitaniumInsanium int `json:"titanium_insanium"`
						MiningSpeedBoost int `json:"mining_speed_boost"`
						DailyPowder      int `json:"daily_powder"`
						EfficientMiner   int `json:"efficient_miner"`
						MiningExperience int `json:"mining_experience"`
					} `json:"nodes"`
					ReceivedFreeTier            bool    `json:"received_free_tier"`
					Tokens                      int     `json:"tokens"`
					PowderMithril               int     `json:"powder_mithril"`
					PowderMithrilTotal          int     `json:"powder_mithril_total"`
					TokensSpent                 int     `json:"tokens_spent"`
					PowderSpentMithril          int     `json:"powder_spent_mithril"`
					Experience                  float64 `json:"experience"`
					LastReset                   int64   `json:"last_reset"`
					RetroactiveTier2Token       bool    `json:"retroactive_tier2_token"`
					SelectedPickaxeAbility      string  `json:"selected_pickaxe_ability"`
					DailyOresMinedDayMithrilOre int     `json:"daily_ores_mined_day_mithril_ore"`
					DailyOresMinedMithrilOre    int     `json:"daily_ores_mined_mithril_ore"`
					GreaterMinesLastAccess      int64   `json:"greater_mines_last_access"`
					Crystals                    struct {
						JadeCrystal struct {
						} `json:"jade_crystal"`
						AmberCrystal struct {
						} `json:"amber_crystal"`
						AmethystCrystal struct {
						} `json:"amethyst_crystal"`
						SapphireCrystal struct {
						} `json:"sapphire_crystal"`
						TopazCrystal struct {
						} `json:"topaz_crystal"`
						JasperCrystal struct {
						} `json:"jasper_crystal"`
						RubyCrystal struct {
						} `json:"ruby_crystal"`
					} `json:"crystals"`
					Biomes struct {
						Dwarven struct {
							StatuesPlaced []interface{} `json:"statues_placed"`
						} `json:"dwarven"`
					} `json:"biomes"`
					PowderGemstone      int `json:"powder_gemstone"`
					PowderGemstoneTotal int `json:"powder_gemstone_total"`
				} `json:"mining_core"`
				Forge struct {
					ForgeProcesses struct {
						Forge1 struct {
						} `json:"forge_1"`
					} `json:"forge_processes"`
				} `json:"forge"`
				ExperienceSkillRunecrafting float64  `json:"experience_skill_runecrafting"`
				ExperienceSkillMining       float64  `json:"experience_skill_mining"`
				UnlockedCollTiers           []string `json:"unlocked_coll_tiers"`
				ExperienceSkillAlchemy      float64  `json:"experience_skill_alchemy"`
				BackpackContents            struct {
					Num0 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"0"`
					Num1 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"1"`
					Num2 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"2"`
					Num3 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"3"`
					Num4 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"4"`
					Num5 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"5"`
					Num6 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"6"`
					Num7 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"7"`
					Num8 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"8"`
					Num9 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"9"`
					Num10 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"10"`
					Num11 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"11"`
					Num12 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"12"`
					Num13 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"13"`
					Num14 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"14"`
					Num15 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"15"`
					Num16 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"16"`
					Num17 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"17"`
				} `json:"backpack_contents"`
				ExperienceSkillTaming float64 `json:"experience_skill_taming"`
				SacksCounts           struct {
					String             int `json:"STRING"`
					Sulphur            int `json:"SULPHUR"`
					BlazeRod           int `json:"BLAZE_ROD"`
					GhastTear          int `json:"GHAST_TEAR"`
					SpiderEye          int `json:"SPIDER_EYE"`
					RottenFlesh        int `json:"ROTTEN_FLESH"`
					Bone               int `json:"BONE"`
					SlimeBall          int `json:"SLIME_BALL"`
					EnderPearl         int `json:"ENDER_PEARL"`
					MagmaCream         int `json:"MAGMA_CREAM"`
					Log                int `json:"LOG"`
					LOG1               int `json:"LOG:1"`
					LOG2               int `json:"LOG:2"`
					LOG21              int `json:"LOG_2:1"`
					Log2               int `json:"LOG_2"`
					LOG3               int `json:"LOG:3"`
					RevenantFlesh      int `json:"REVENANT_FLESH"`
					TarantulaWeb       int `json:"TARANTULA_WEB"`
					RAWFISH1           int `json:"RAW_FISH:1"`
					RawFish            int `json:"RAW_FISH"`
					ClayBall           int `json:"CLAY_BALL"`
					WaterLily          int `json:"WATER_LILY"`
					RAWFISH3           int `json:"RAW_FISH:3"`
					InkSack            int `json:"INK_SACK"`
					PrismarineCrystals int `json:"PRISMARINE_CRYSTALS"`
					PrismarineShard    int `json:"PRISMARINE_SHARD"`
					Sponge             int `json:"SPONGE"`
					RAWFISH2           int `json:"RAW_FISH:2"`
					WolfTooth          int `json:"WOLF_TOOTH"`
					RuneWhiteSpiral1   int `json:"RUNE_WHITE_SPIRAL_1"`
					RuneGem1           int `json:"RUNE_GEM_1"`
					RuneBlood21        int `json:"RUNE_BLOOD_2_1"`
					RuneSnow1          int `json:"RUNE_SNOW_1"`
					RuneSparkling1     int `json:"RUNE_SPARKLING_1"`
					RuneHearts1        int `json:"RUNE_HEARTS_1"`
					RuneGolden1        int `json:"RUNE_GOLDEN_1"`
					RuneLava1          int `json:"RUNE_LAVA_1"`
					RuneRainbow1       int `json:"RUNE_RAINBOW_1"`
					RuneZombieSlayer1  int `json:"RUNE_ZOMBIE_SLAYER_1"`
					RuneZap1           int `json:"RUNE_ZAP_1"`
					RuneClouds1        int `json:"RUNE_CLOUDS_1"`
					RuneHot1           int `json:"RUNE_HOT_1"`
					RuneIce1           int `json:"RUNE_ICE_1"`
					RuneWake1          int `json:"RUNE_WAKE_1"`
					NullSphere         int `json:"NULL_SPHERE"`
					RedGift            int `json:"RED_GIFT"`
					SnowBall           int `json:"SNOW_BALL"`
					IceHunk            int `json:"ICE_HUNK"`
				} `json:"sacks_counts"`
				EssenceUndead int `json:"essence_undead"`
				BackpackIcons struct {
					Num0 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"0"`
					Num1 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"1"`
					Num2 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"2"`
					Num3 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"3"`
					Num4 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"4"`
					Num5 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"5"`
					Num6 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"6"`
					Num7 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"7"`
					Num8 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"8"`
					Num9 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"9"`
					Num10 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"10"`
					Num11 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"11"`
					Num12 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"12"`
					Num13 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"13"`
					Num14 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"14"`
					Num15 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"15"`
					Num16 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"16"`
					Num17 struct {
						Type int    `json:"type"`
						Data string `json:"data"`
					} `json:"17"`
				} `json:"backpack_icons"`
				ExperienceSkillCombat  float64 `json:"experience_skill_combat"`
				EssenceDiamond         int     `json:"essence_diamond"`
				ExperienceSkillFarming float64 `json:"experience_skill_farming"`
				Collection             struct {
					Log                    int `json:"LOG"`
					LOG2                   int `json:"LOG:2"`
					LOG1                   int `json:"LOG:1"`
					LOG21                  int `json:"LOG_2:1"`
					InkSack                int `json:"INK_SACK"`
					WaterLily              int `json:"WATER_LILY"`
					Log2                   int `json:"LOG_2"`
					LOG3                   int `json:"LOG:3"`
					RottenFlesh            int `json:"ROTTEN_FLESH"`
					Cobblestone            int `json:"COBBLESTONE"`
					Seeds                  int `json:"SEEDS"`
					Bone                   int `json:"BONE"`
					String                 int `json:"STRING"`
					EnderPearl             int `json:"ENDER_PEARL"`
					Pork                   int `json:"PORK"`
					Wheat                  int `json:"WHEAT"`
					Coal                   int `json:"COAL"`
					GoldIngot              int `json:"GOLD_INGOT"`
					INKSACK4               int `json:"INK_SACK:4"`
					SpiderEye              int `json:"SPIDER_EYE"`
					SlimeBall              int `json:"SLIME_BALL"`
					Diamond                int `json:"DIAMOND"`
					IronIngot              int `json:"IRON_INGOT"`
					Mutton                 int `json:"MUTTON"`
					CarrotItem             int `json:"CARROT_ITEM"`
					MushroomCollection     int `json:"MUSHROOM_COLLECTION"`
					Sulphur                int `json:"SULPHUR"`
					MagmaCream             int `json:"MAGMA_CREAM"`
					Redstone               int `json:"REDSTONE"`
					BlazeRod               int `json:"BLAZE_ROD"`
					GhastTear              int `json:"GHAST_TEAR"`
					EnderStone             int `json:"ENDER_STONE"`
					Obsidian               int `json:"OBSIDIAN"`
					Emerald                int `json:"EMERALD"`
					RawFish                int `json:"RAW_FISH"`
					RAWFISH1               int `json:"RAW_FISH:1"`
					ClayBall               int `json:"CLAY_BALL"`
					RAWFISH3               int `json:"RAW_FISH:3"`
					RAWFISH2               int `json:"RAW_FISH:2"`
					PrismarineCrystals     int `json:"PRISMARINE_CRYSTALS"`
					PrismarineShard        int `json:"PRISMARINE_SHARD"`
					Melon                  int `json:"MELON"`
					Pumpkin                int `json:"PUMPKIN"`
					PotatoItem             int `json:"POTATO_ITEM"`
					Ice                    int `json:"ICE"`
					GlowstoneDust          int `json:"GLOWSTONE_DUST"`
					Sponge                 int `json:"SPONGE"`
					SugarCane              int `json:"SUGAR_CANE"`
					Gravel                 int `json:"GRAVEL"`
					Feather                int `json:"FEATHER"`
					RawChicken             int `json:"RAW_CHICKEN"`
					Leather                int `json:"LEATHER"`
					Rabbit                 int `json:"RABBIT"`
					Cactus                 int `json:"CACTUS"`
					EnchantedFlint         int `json:"ENCHANTED_FLINT"`
					EnchantedDiamond       int `json:"ENCHANTED_DIAMOND"`
					EnchantedString        int `json:"ENCHANTED_STRING"`
					EnchantedSpiderEye     int `json:"ENCHANTED_SPIDER_EYE"`
					EnchantedRedstone      int `json:"ENCHANTED_REDSTONE"`
					RabbitHide             int `json:"RABBIT_HIDE"`
					RabbitFoot             int `json:"RABBIT_FOOT"`
					EnchantedRabbit        int `json:"ENCHANTED_RABBIT"`
					NetherStalk            int `json:"NETHER_STALK"`
					Netherrack             int `json:"NETHERRACK"`
					Quartz                 int `json:"QUARTZ"`
					Sand                   int `json:"SAND"`
					INKSACK3               int `json:"INK_SACK:3"`
					Egg                    int `json:"EGG"`
					EnchantedGlowstoneDust int `json:"ENCHANTED_GLOWSTONE_DUST"`
					EnchantedCarrot        int `json:"ENCHANTED_CARROT"`
					EnchantedLapisLazuli   int `json:"ENCHANTED_LAPIS_LAZULI"`
					SnowBall               int `json:"SNOW_BALL"`
					EnchantedEmerald       int `json:"ENCHANTED_EMERALD"`
					EnchantedPotato        int `json:"ENCHANTED_POTATO"`
					EnchantedClayBall      int `json:"ENCHANTED_CLAY_BALL"`
					MithrilOre             int `json:"MITHRIL_ORE"`
					EnchantedSnowBlock     int `json:"ENCHANTED_SNOW_BLOCK"`
					HardStone              int `json:"HARD_STONE"`
					GemstoneCollection     int `json:"GEMSTONE_COLLECTION"`
				} `json:"collection"`
				EssenceDragon             int     `json:"essence_dragon"`
				EssenceGold               int     `json:"essence_gold"`
				ExperienceSkillEnchanting float64 `json:"experience_skill_enchanting"`
				ExperienceSkillFishing    float64 `json:"experience_skill_fishing"`
				EssenceIce                int     `json:"essence_ice"`
				EssenceWither             int     `json:"essence_wither"`
				EssenceSpider             int     `json:"essence_spider"`
				ExperienceSkillForaging   float64 `json:"experience_skill_foraging"`
				ExperienceSkillCarpentry  float64 `json:"experience_skill_carpentry"`
			} `json:"members"`
		CommunityUpgrades struct {
			CurrentlyUpgrading interface{} `json:"currently_upgrading"`
			UpgradeStates      []struct {
				Upgrade     string `json:"upgrade"`
				Tier        int    `json:"tier"`
				StartedMs   int64  `json:"started_ms"`
				StartedBy   string `json:"started_by"`
				ClaimedMs   int64  `json:"claimed_ms"`
				ClaimedBy   string `json:"claimed_by"`
				Fasttracked bool   `json:"fasttracked"`
			} `json:"upgrade_states"`
		} `json:"community_upgrades"`
		CuteName string `json:"cute_name"`
		Banking  struct {
			Balance      float64 `json:"balance"`
			Transactions []struct {
				Amount        float64    `json:"amount"`
				Timestamp     int64  `json:"timestamp"`
				Action        string `json:"action"`
				InitiatorName string `json:"initiator_name"`
			} `json:"transactions"`
		} `json:"banking"`
	} `json:"profiles"`
}