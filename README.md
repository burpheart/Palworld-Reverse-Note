
# 调试窗口
下面的lua脚本皆为[RE-UE4SS](https://github.com/UE4SS-RE/RE-UE4SS)脚本

开启调试窗口
```lua
RegisterKeyBind(Key.F10, function()
    local PalPlayerController = FindFirstOf("PalPlayerController")
    local PalCheatManager = PalPlayerController.CheatManager
    PalPlayerController:EnableCheats()
    PalCheatManager:DebugWindow()
  end)

```
有一部分命令会触发崩溃

单人离线模式下 也可以使用CRTL+ALT+O开启控制台

控制台在多人模式下 大部分功能不可用


一些离线可用的命令
![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706157223650-fd3fe8d8-14a6-45f7-b35c-168e49d68d06.png#averageHue=%23c8c8c7&clientId=u41804305-9356-4&from=paste&height=1619&id=GQ75f&originHeight=2428&originWidth=2552&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=1397630&status=done&style=none&taskId=u8e87d386-0113-4ff6-9683-6f653fa00eb&title=&width=1701.3333333333333)
![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706157280059-00b51d0f-4979-4bcf-ac40-d5b3b9a35978.png#averageHue=%231c2728&clientId=u41804305-9356-4&from=paste&height=650&id=Yv8yj&originHeight=975&originWidth=2187&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=995308&status=done&style=none&taskId=u9106c625-00f3-4b99-a04f-664d05b6b66&title=&width=1458)
# 游戏配置

游戏大部分可调的配置是在PalGameSetting中储存的

![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706157413191-9d11eab5-240d-45f9-95f1-68a8e17dfe8c.png#averageHue=%23559843&clientId=u41804305-9356-4&from=paste&height=676&id=ue11de5d7&originHeight=1014&originWidth=1360&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=317601&status=done&style=none&taskId=u10a214ed-0c47-44fe-b6a1-3a77aefcf47&title=&width=906.6666666666666)

```csharp

class UPalGameSetting : public UBlueprintFunctionLibrary{
int32 CharacterMaxLevel; 
int32 GuildCharacterMaxLevel; 
int32 OtomoLevelSyncAddMaxLevel; 
enum class EPalPlayerSprintStaminaDecreaseType PlayerSprintStaminaType; 
int32 CharacterMaxRank; 
int32 ConsumStamina_PalThrow; 
float ReturnOtomoPalCoolTime; 
int32 OtomoSlotNum; 
float OtomoWazaCoolDownSpeedRate; 
float PlayerHPRateFromRespawn; 
float PlayerStomachRateFromRespawn; 
float RarePal_AppearanceProbability; 
float RarePal_LevelMultiply; 
int32 BossOrRarePal_TalentMin; 
int32 CharacterRankUpRequiredNumDefault; 
TMap<int32, int32> CharacterRankUpRequiredNumMap; 
float NaturalUpdateSaveParameterInterval; 
float CharacterHUDDisplayRange; 
float CharacterFedEatingTime; 
float CharacterStayingSecondsInSpa; 
int32 PalBoxPageNum; 
int32 PalBoxSlotNumInPage; 
float PlayerBattleJudge_EnemyDistance; 
TMap<int32, float> BodyTemperature_SlipDamage_Percent; 
TMap<int32, float> BodyTemperature_StomachDecreaceRate; 
TSoftObjectPtr<class UNiagaraSystem> SleepFXDefault; 
float LiftupCharacterThrownVelocityScalar; 
float LiftupCharacterClearCollisionDelayTime; 
int32 NickNameMaxLength; 
float IntervalForPalAttackFromBall; 
bool CanShootRiderByFullRide; 
bool HitWazaAttackForMapObject; 
int32 KnockBack_MaxHpPercent; 
float KnockBack_Power; 
float StunTime; 
float StepCooldownTime; 
float Stun_GunDamageRate; 
bool IsEnableAutoReload; 
float DeadShootImpulseRate; 
float DeadShootImpulseMax; 
float OtomoDamageRate_Defense; 
int32 DamageValueMin_MapObject; 
float DamageRate_WealPoint; 
float DamageRate_StrongPoint; 
float DamageRate_SleepHit; 
float FinalDamageRate_Waza; 
float FinalDamageRate_Weapon; 
float FinalDamageRate_Mine; 
float FinalDamageRate_MapObject_Waza; 
float FinalDamageRate_MapObject_Weapon; 
float FinalDamageRate_MapObject_Mine; 
float FoliageDefault_Defense; 
float WazaAttackerLevelDamage_Power; 
float DamageRandomRate_Min; 
float DamageRandomRate_Max; 
float LevelDamageCorrect; 
float LevelStatusAccumulateCorrect; 
float DamageElementMatchRate; 
float MineAttack_DefensePower; 
float StatusCalculate_LevelMultiply_HP; 
int32 StatusCalculate_TribePlus_HP; 
int32 StatusCalculate_ConstPlus_HP; 
float StatusCalculate_LevelMultiply_Attack; 
int32 StatusCalculate_ConstPlus_Attack; 
float StatusCalculate_LevelMultiply_Defense; 
int32 StatusCalculate_ConstPlus_Defense; 
float StatusCalculate_TribeMultiply_CraftSpeed; 
float StatusCalculate_GenkaiToppa_PerAdd; 
float StatusCalculate_Talent_PerAdd; 
float BreakedWeaponDamageRate; 
float BreakedArmorDefenseRate; 
float ArmorDurabilityDamageDivide; 
float ClimbingStamina_Move; 
float ClimbingStamina_Jump; 
float RideWazaStaminaRate; 
bool bIsEnableJumpPreliminary; 
float JumpInterval; 
float FlyMaxHeight; 
float FlyHover_SP; 
float FlyHorizon_SP; 
float FlyHorizon_Dash_SP; 
float FlyVertical_SP; 
float AimingSpeedRateInRide; 
float SlidingEndSpeed; 
int32 JumpSP; 
int32 StepSP; 
int32 MeleeAttackSP; 
float SprintSP; 
float GliderSP; 
float SwimmingFallWaitTimeSec; 
float Swimming_SP_Idle; 
float Swimming_SP_Swim; 
float Swimming_SP_DashSwim; 
float FluidFriction; 
float OverWeightSpeedZero_AddPercent; 
float WalkableFloorAngleForDefault; 
float WalkableFloorAngleForRide; 
bool IsEnableSpeedCollision; 
float CollisionDamageMinSpeed; 
float SpeedCollisionDamagePower; 
float CollisionDamageSpeedMultiplay; 
float CollisionDamageWeightThreshold; 
float AutoHPRegene_Percent_perSecond; 
float AutoHPRegene_Percent_perSecond_Sleeping; 
float PosionHPDecrease_Percent_perSecond; 
float Starvation_DecreaseHP_Percent_perSecond; 
float AutoSANRegene_Percent_perSecond_PalStorage; 
float StomachDecreace_perSecond_Monster; 
float StomachDecreace_perSecond_Player; 
float StomachDecreace_AutoHealing; 
float StomachDecreace_WorkingRate; 
int32 HungerStart_StomachValue; 
int32 FullStomachPalStartEatFood; 
float FullStomachCost_ByWazaUse_Base; 
TMap<int32, float> FullStomachCost_ByWazaUse_RateMap; 
float StomachDecreaceRate_GroundRide_Sprint; 
float StomachDecreaceRate_WaterRide; 
float StomachDecreaceRate_WaterRide_Sprint; 
float StomachDecreaceRate_FlyRide; 
float StomachDecreaceRate_FlyRide_Sprint; 
int32 RemainderOfLife_Second; 
float HpDecreaseRate_Drowning; 
float PlayerShield_RecoverStartTime; 
float PlayerShield_RecoverPercentPerSecond; 
float StaminaRecover_PercentPerSecond; 
float ResuscitationTime; 
int32 PlayerDeath_DropOtomoNum; 
float PlayerDeath_DropOtomoRange; 
int32 PlayerDeath_DropOtomoDisappearHours; 
float PlayerDyingDamagePerTime; 
int32 ElementStatus_ResistanceInitialValue; 
float ElementStatus_AutoDecreasePerSecond; 
int32 ElementStatus_ResetResistanceSecond; 
int32 BuildExp; 
int32 CraftExp; 
int32 PickupItemOnLevelExp; 
int32 MapObjectDestroyProceedExp; 
float MapObjectDistributeExpRange; 
TMap<int32, int32> OtomoExp_LevelDifferenceMap; 
int32 OtomoExp_HigherPlayerLevel; 
int32 CaptureExpBonus_Tier1_TableReferenceNum; 
int32 CaptureExpBonus_Tier2_TableReferenceNum; 
int32 CaptureExpBonus_Tier3_TableReferenceNum; 
TArray<struct FPalDebugOtomoPalInfo> NewGameOtomoPalSet; 
TMap<class FName, int32> NewGameInventoryItemSet; 
TMap<class FName, int32> NewGameLoadoutItemSet; 
struct FVector WorldHUDDisplayOffsetDefault; 
float WorldHUDDisplayRangeDefault; 
float WorldHUDDetailDisplayRange; 
TArray<struct FPalDataTableRowName_ItemData> FarmCropWaterItemIds; 
float FarmCropGrowupSpeedBySec; 
float FarmCropIncreaseRateByWaterFillRate; 
class FString MaxMoney; 
int32 DefaultMoney; 
float SneakAttackBackJudgeAngle_Degree; 
float SneakAttack_PalMeleeWaza_AttackRate; 
float AutoAimCameraMoveRate; 
float AutoAimCharacterMoveRate; 
float AutoAimCameraAdsorptionSpeed; 
float AutoAimLockOnScreenSpaceRate; 
float ForceAutoAimTime; 
float SellItemRate; 
float PalPriceConstantValueA; 
float PalPriceConstantValueB; 
float SellPalRate; 
float SearchRangeOnThrowedCharacterLanded; 
float WorkCompleteReactionRangeFromPlayer; 
int32 WorkerCollectResourceStackMaxNum; 
enum class EPalFacialEyeType FacialTypeHardWork; 
float Timeout_WorkerApproachToTarget; 
float WaitTime_WorkRepairFailedFindPath; 
float WorkerWaitingNotifyInterval; 
float WarpCheckInterval; 
float WarpCheckMoveDistanceThreshold; 
float WarpThreshold; 
float AutoDecreaseHateValue_PercentMaxHP_PerSecond; 
float HateDecreaseDamageRate; 
float Hate_ForceUp_HPRate_OtomoActive; 
float Hate_ForceUp_HPRate_IncidentBattle; 
float CombatEndDistance_BattleStartSelfPos_To_SelfPos; 
float CombatEndDistance_BattleStartSelfPos_To_TargetPos; 
float CombatEndDistance_BattleStartSelfPos_To_TargetPos_AddFirstTargetDistance; 
float NavigationAreaDivideExtents; 
TArray<struct FPalNavigationUpdateFrequencySetting> NavigationUpdateFrequencySettingsFromPlayer; 
float AutoSaveSpan; 
class FName SaveDataName_WorldBaseInfo; 
class FName SaveDataName_World; 
class FName SaveDataName_PlayerDirectory; 
class FName SaveDataName_LocalData; 
class FName SaveDataName_WorldOption; 
TMap<enum class EPalSupportedPlatformType, int32> MaxWorldDataNumMap; 
int32 PalWorldTime_GameStartHour; 
int32 PalWorldMinutes_RealOneMinute; 
int32 NightStartHour; 
int32 NightEndHour; 
int32 PlayerMorningHour; 
int32 PlayerSleepStartHour; 
int32 NightSkipWaitSecond; 
float LocalPlayerAndOtomo_LightRangeScale; 
struct FVector BuildBaseUnitGridDefinition; 
float BuildSimulationVerticalAdjustRate; 
float BuildSimulationVerticalMinLength; 
float BuildSimulationFoundationFloatingAllowance; 
TMap<enum class EPalBuildObjectInstallStrategy, struct FVector> BuildSimulationFoundationCheckCollisionScale; 
float BuildSimulationRoofHeightOffset; 
float BuildSimulationStairHeightOffset; 
float BuildSimulationLeanAngleMax; 
float BuildingProgressInterpolationSpeed; 
int32 PlayerRecord_BuildingObjectMaxNum; 
float BuildingMaxZ; 
int32 BuildObj_HatchedPalCharacterLevel; 
float BuildObj_DamageScarecrowStartRecoveryTime; 
float BaseCampAreaRange; 
float BaseCampPalFindWorkRange; 
float PalArriveToWorkLocationRange; 
float PalArriveToWorkLocationRangeZ; 
float BaseCampNeighborMinimumDistance; 
float PalRotateSpeedToWork; 
float BaseCampFoliageWorkableRange; 
float BaseCampHungerApproachToPlayer; 
float BaseCampHungerUnreachableObjectTimeoutRealSeconds; 
float HungerHUDDisplayRange; 
float WorkAmountBySecForPlayer; 
float BaseCampWorkerEventTriggerInterval; 
float BaseCampWorkerEventTriggerProbability; 
float BaseCampWorkerSanityWarningThreshold; 
float BaseCampWorkerFinishEatingFullStomach; 
float BaseCampWorkerFinishEatingSanity; 
int32 BaseCampWorkerFinishEatCount; 
float BaseCampWorkerRecoverHungryTurnToTargetTimeout; 
float BaseCampWorkerStartSleepHpPercentage; 
float BaseCampWorkerSleepInPlaceRecoverSanityRate; 
float BaseCampWorkerDistancePickableItem; 
TArray<enum class EPalBaseCampItemContainerType> BaseCampBuildingItemContainerPriority; 
float FoliageRespawnFailedExtraRangeOfBaseCamp; 
float BaseCampPalCombatRange_AddCampRange; 
struct FFloatInterval BaseCampPalWalkTime_BeforeSleep; 
float BaseCampTimeFinishBattleModeAfterEmptyEnemy; 
enum class EPalCharacterImportanceType BaseCampWorkerMoveModeChangeThreshold; 
int32 BaseCampWorkerDirectorTickForAssignWorkByCount; 
float BaseCampWorkerLookToTargetWork; 
float ReviveWorkAdditionalRange; 
float WorkAroundRangeDefault; 
TArray<enum class EPalWorkType> IssueNotifyWorkTypes; 
float WorkAmountByManMonth; 
float WorkNotifyDelayTime; 
float WorkFinishDelayCallAddWorkNotifyDelayTime; 
float WorkIgnitionTorchWaitTime; 
TMap<enum class EPalWorkAssignableCheckResult, enum class EPalMonsterControllerBaseCampLogType> WorkAssignFailedLogTypeMap; 
float WorkTransportingSpeedRate; 
TArray<struct FPalDataTableRowName_ItemData> BaseCampNotTransportItemBlackList; 
float WorkTransportingDelayTimeDropItem; 
float BaseCampStopProvideEnergyInterval; 
TMap<enum class EPalBaseCampPassiveEffectWorkHardType, struct FPalBaseCampPassiveEffectWorkHardInfo> BaseCampPassiveEffectWorkHardInfoMap; 
float BaseCampWorkCollectionRestoreStashSeconds; 
TArray<struct FPalWorkTypeSet> WorkTypeAssignPriorityOrder; 
struct FPalWorkAssignDefineDataStaticSetting WorkAssignDefineData_Build; 
  struct FPalWorkAssignDefineDataStaticSetting WorkAssignDefineData_FoliageWork; 
  struct FPalWorkAssignDefineDataStaticSetting WorkAssignDefineData_ReviveCharacterWork; 
  struct FPalWorkAssignDefineDataStaticSetting WorkAssignDefineData_TransportItemInBaseCamp; 
  struct FPalWorkAssignDefineDataStaticSetting WorkAssignDefineData_RepairBuildObjectInBaseCamp; 
  struct FPalWorkAssignDefineDataStaticSetting WorkAssignDefineData_BreedFarm; 
  struct FPalWorkAssignDefineDataStaticSetting WorkAssignDefineData_ExtinguishBurn; 
  int32 WorkSuitabilityMaxRank; 
  TMap<enum class EPalWorkSuitability, struct FPalWorkSuitabilityDefineData> WorkSuitabilityDefineDataMap; 
  struct FPalWorkSuitabilityCollectionDefineData WorkSuitabilityDefineData_Collection; 
  struct FPalWorkSuitabilityDeforestDefineData WorkSuitabilityDefineData_Deforest; 
  struct FPalWorkSuitabilityMiningDefineData WorkSuitabilityDefineData_Mining; 
  int32 DropItemWaitInsertMaxNumPerTick; 
  struct FPalDungeonMarkerPointSpawnParameter DungeonSpawnParameterDefault; 
  float GamePad_NotAimCameraRotateSpeed_DegreePerSecond; 
  float GamePad_AimCameraRotateSpeed_DegreePerSecond; 
  float Mouse_NotAimCameraRotateSpeed; 
  float Mouse_AimCameraRotateSpeed; 
  float YawCameraMaxSpeedRate; 
  float TimeForCameraMaxSpeed; 
  float AimInterpInterval; 
  int32 InvaderSelfDeleteAddTime; 
  float InvadeProbability; 
  int32 InvadeOccurablePlayerLevel; 
  int32 InvadeJudgmentInterval_Minutes; 
  int32 InvadeCollTime_Max_Minutes; 
  int32 InvadeCollTime_Min_Minutes; 
  int32 InvadeReturnTime_Minutes; 
  int32 InvadeStartPoint_BaseCampRadius_Min_cm; 
  int32 InvadeStartPoint_BaseCampRadius_Max_cm; 
  float VisitorNPCProbability; 
  int32 VisitorNPCReturnTime_Minutes; 
  float RidingAimOpacity; 
  float HideUITimeWhenNotConflict; 
  float FirstCapturedUIDisplayTime; 
  float CapturedUIDisplayTime; 
  float FirstActivatedOtomoInfoDisplayTime; 
  float PlayerLevelUpUIDIsplayTime; 
  float PlayerExpGaugeUIDisplayTime; 
  float OtomoExpGaugeUIDisplayTime; 
  float NpcGaugeDisplayDistance; 
  float NpcGaugeDisplayRange_CameraSight; 
  float GuildMemberGaugeDisplayDIstance; 
  float DownPlayerLoupeDisplayDistance; 
  float DownPlayerGaugeDisplayRange_CameraSight; 
  struct FVector2D ReticleOffsetRate; 
  int32 LowHealthEffectParcent; 
  TMap<enum class EPalDamageTextType, int32> DamageTextMargineMap; 
  float DamageTextDisplayLength; 
  struct FVector2D DamageTextMaxOffset; 
  float DamageTextOffsetInterpolationLength; 
  TMap<enum class EPalDamageTextType, float> DamageTextScaleMap; 
  struct FVector2D DamageTextRandomOffset; 
  int32 StrongEnemyMarkLevel; 
  float OtomoInteractUIDisplayDistance; 
  float EnemyMarkUIMinScale; 
  float EnemyMarkScaleInterpolationLength; 
  struct FVector2D NpcHPGaugeGlobalOffset; 
  float DelayGaugeStartTime; 
  float DelayGaugeProgressPerSecond; 
  float InventoryWeaponRangeMaxBorder; 
  float InventoryWeaponStabilityMinBorder; 
  float InventoryWeaponAccuracyMinBorder; 
  float WorldmapUIMaskClearSize; 
  float WorldmapUIFTMergeDistance; 
  int32 WorldmapUIMaxMarker; 
float NPCHPGaugeUpdateSpan; 
float CaptureFailedUIDisplayTime; 
TArray<struct FPalDataTableRowName_ItemData> CaptureSphereSortArray; 
float OpenGameOverUITime; 
float BlockRespawnTime; 
float InventoryWeightAlertRate; 
float InventoryWeightGaugeDIsplayTime; 
float OtomoLevelUpNoticeUIDisplayTime; 
float OtomoMasteredWazaNoticeUIDisplayTime; 
float ProgressGaugeInterpolationSpeed; 
float TeleportFadeInTime; 
float TeleportFadeOutTime; 
float PlayerTeleportTimeoutTime; 
TArray<float> PassiveSkillAppendNumWeights; 
bool bIsEggLauncherExplosion; 
float ThrowPalBattleRadius; 
float ThrowPalWorkRadius; 
float RopeHitPowe; 
float RopePullPower; 
float DefaultMaxInventoryWeight; 
float RaycastLengthForDetectIndoor; 
float MapObjectConnectAnyPlaceRaycastLength; 
float ShootingTargetRayCastDistance; 
TArray<float> CaptureJudgeRateArray; 
int32 CaptureBallBoundCountMax; 
TArray<class FName> ExceptCapturedItemList; 
TMap<enum class EPalCaptureSphereLevelType, int32> CaptureSphereLevelMap; 
TMap<enum class EPalStatusID, float> CaptureRateAddByStatusMap; 
float IgnoreFirstCaptureFailedHPRate; 
float CaptureRateAdd_ByLegHold; 
float LongPressInterval; 
float LongPressInterval_EnemyCampCage; 
float LongPressInterval_GetHatchedPal; 
float CrouchLockAttenuation; 
bool IsEnableCharacterWazaScale; 
bool IsOverrideDamageAdditiveAnimation; 
float BlinkInterval; 
float CrimeStateMaintainDurationBaseDefault; 
int32 TechnologyPointPerLevel; 
int32 bossTechnologyPointPerTowerBoss; 
int32 bossTechnologyPointPerNormalBoss; 
TArray<struct FPalDataTableRowName_RecipeTechnologyData> DefaultUnlockTechnology; 
int32 DefaultTechnologyPoint; 
int32 TechnologyPoint_UnlockFastTravel; 
float DecreaseSanity_DamagedMultiply; 
int32 FullStomachPercent_RecoverySanity; 
float RecoverySanity_FullStomach; 
float DecreaseSanity_Hunger; 
float DecreaseSanity_Starvation; 
bool Spawner_IsCheckLoadedWorldPartition; 
float SpawnerDisableDistanceCM_FromBaseCamp; 
float Spawner_DefaultSpawnRadius_S; 
float Spawner_DefaultSpawnRadius_M; 
float Spawner_DefaultSpawnRadius_L; 
float Spawner_DefaultSpawnRadius_NPC; 
float Spawner_DefaultDespawnDistance_S; 
float Spawner_DefaultDespawnDistance_M; 
float Spawner_DefaultDespawnDistance_L; 
float Spawner_DefaultDespawnDistance_NPC; 
class UDataTable* CharacterHeadMeshDataTable; 
class UDataTable* CharacterBodyMeshDataTable; 
class UDataTable* CharacterHairMeshDataTable; 
class UDataTable* CharacterEquipmentArmorMeshDataTable; 
class UDataTable* CharacterEyeMaterialDataTable; 
float CharacterMakeColorLimit_SV; 
bool IsAutoEquipMasteredWaza; 
bool ActiveUNKO; 
int32 MaxSpawnableDeathPenaltyChest; 
class FName BuildObjectInstallStrategy_SinkAllowCollisionPresetName; 
float MapObjectShakeTimeOnDamaged; 
struct FVector MapObjectShakeOffsetOnDamaged; 
int32 MapObjectOutlineByReticleTargetting; 
int32 MapObjectOutlineByInteractable; 
struct FPalMapObjectRepairInfo MapObjectRepairInfo; 
float FoliageExtentsXY; 
int32 FoliageChunkSeparateScale; 
float MapObjectHPDisplayDistance; 
float MapObjectHPDisplayTime; 
float MapObjectGateLockTime; 
bool bDirectObtainFromTreasureBox; 
float MapObjectEffectTriggerAccumulate_Burn; 
float MapObjectEffect_Burn_DamageHpRate; 
struct FVector MapObjectEffect_Burn_DamageAroundRange; 
float MapObjectEffect_Burn_DamageAroundInterval; 
float MapObjectEffect_Burn_DamageAroundDamageValue; 
float MapObjectEffect_Burn_DamageAroundAccumulateValue; 
float MapObjectEffect_Burn_DamageAroundAccumulateValue_ForCharacter; 
int32 PasswordLockFailedMaxNum; 
float MapObjectItemChestCorruptionRateFromWorkSpeed; 
struct FPalOptimizeParameter RuntimeOptimizeParameter; 
TMap<int32, struct FPalWorldSecurityWantedPoliceSettingData> WorldSecurityWantedPoliceSettingDataMap; 
TMap<int32, struct FPalWorldSecurityWantedPoliceSettingData> WorldSecurityWantedPoliceSettingDataMapForDS; 
int32 StatusPointPerLevel; 
float AddMaxHPPerStatusPoint; 
float AddMaxSPPerStatusPoint; 
float AddPowerPerStatusPoint; 
float AddMaxInventoryWeightPerStatusPoint; 
float AddCaptureLevelPerStatusPoint; 
float AddWorkSpeedPerStatusPoint; 
float AddMaxHPPerHPRank; 
float AddAttackPerAttackRank; 
float AddDefencePerDefenceRank; 
float AddWorkSpeedPerWorkSpeedRank; 
TArray<float> Combi_TalentInheritNum; 
TArray<float> Combi_PassiveInheritNum; 
TArray<float> Combi_PassiveRandomAddNum; 
TArray<struct FPalEggRankInfo> PalEggRankInfoArray; 
TMap<enum class EPalElementType, struct FPalDataTableRowName_MapObjectData> PalEggMapObjectIdMap; 
TMap<int32, float> PalEggHatchingSpeedRateByTemperature; 
class UFont* DebugInfoFont; 
int32 MaxGuildNameLength; 
float JoinGuildRequestInteractLongPushTime; 
float TutorialMinDisplayTime; 
float TutorialDisplayTime; 
TMap<enum class EPalUIRewardDisplayType, float> CommonRewardDisplayTime; 
float DeadBodyDestroySecond; 
float EnemyCampRespawnCoolTime; 
float EnemyCampDespawnDelayTime; 
float PalBoxReviveTime; 
float AfterNPCTalkDelayTime_Interact; 
float MinSprintThreshold; 
float MaxSprintThreshold; 
float MinHPGaugeDisplayTime; 
class UDataTable* SoundSourceDataTable; 
}
```
示例

提高闪光帕鲁捕获概率为 100%
```lua
GameSettingInstance = FindFirstOf("PalGameSetting")
GameSettingInstance.RarePal_AppearanceProbability = 100.0

```
已经生成的帕鲁不会改变 下一次帕鲁触发生成前  概率不会生效
世界设置 APalGameStateInGame
```lua
class UPalCharacterManagerReplicator* CharacterManagerReplicator; 
class UPalBaseCampReplicator* BaseCampReplicator; 
class UPalOptionReplicator* OptionReplicator; 
class UPalStageReplicator* StageReplicator; 
class UPalLocationReplicator* LocationReplicator; 
class APalNetworkTransmitter* DedicatedServerTransmitter; 
class UPalGameSystemInitManagerComponent* GameSystemInitManager; 
TArray<class APalBotBuilderLocationBase*> BotBuilderLocation; 
class FString WorldName; 
class FString WorldSaveDirectoryName; 
bool bIsDedicatedServer; 
int32 MaxPlayerNum; //最大玩家数量
struct FGameDateTime WorldTime; //世界世界
struct FDateTime RealUtcTime; 
float ServerFrameTime; 
class FString ServerSessionId; 
int32 ServerWildMonsterCount; 
int32 ServerOtomoMonsterCount; 
int32 ServerBaseCampMonsterCount; 
int32 ServerNPCCount; 
int32 ServerOtherCharacterCount; 
int32 ImportanceCharacterCount_AllUpdate; 
int32 ImportanceCharacterCount_Nearest; 
int32 ImportanceCharacterCount_Near; 
int32 ImportanceCharacterCount_MidInSight; 
int32 ImportanceCharacterCount_FarInSight; 
int32 ImportanceCharacterCount_MidOutSight; 
int32 ImportanceCharacterCount_FarOutSight; 
int32 ImportanceCharacterCount_Farthest; 
int32 BaseCampCount; 
int32 NavMeshInvokerCount; 
TArray<struct FPalChatMessage> ChatMessages; 
FMulticastInlineDelegateProperty_ OnRecievedChatMessageDelegate; 

```

# 存档
```cpp
/Saved/SaveGames/0/世界id/
                        ->Level.sav //世界存档
                        ->LevelMeta.sav
                        ->Players/
                                 ->UID.sav//个人存档 (索引) 
//UID通常为Steam Id开头的十六进制UUID长度id 
    
```


游戏存档使用了ue5内置的存档管理系统

只不过帕鲁在原存档文件的基础上进行了zlib压缩

有两种选项 (一次压缩  两次压缩)

![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706159624474-f6513b81-1c5b-41c0-8f94-fb626dbb63da.png#averageHue=%23fbfbfb&clientId=u41804305-9356-4&from=paste&height=791&id=u471e6f5e&originHeight=1187&originWidth=1597&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=230522&status=done&style=none&taskId=u6844984b-87c3-4837-83a5-aa57460aaa7&title=&width=1064.6666666666667)
![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706159646521-a7f3c89d-087b-4dda-bce5-02d44b2ce9f3.png#averageHue=%23fbfbfb&clientId=u41804305-9356-4&from=paste&height=805&id=u6e4067e5&originHeight=1208&originWidth=1529&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=187785&status=done&style=none&taskId=u5d28fbc5-d3c9-4ac7-b1eb-0111c84964f&title=&width=1019.3333333333334)

```cpp
TTuple<USaveGame *,enum EPalSaveError> *__fastcall UPalSaveGameManager::PalLoadGameFromSlot(
TTuple<USaveGame *,enum EPalSaveError> *result,
const FString *SlotName,
int UserIndex)
{
    bool v4; // al
    FScriptContainerElement *Data; // rcx
    int CompressedSize; // r12d
    FScriptContainerElement *v7; // r15
    int v8; // edx
    FScriptContainerElement *v9; // rsi
    int v10; // edx
    int v11; // edx
    unsigned int v12; // r14d
    FNameEntryId *v13; // rax
    char *v14; // r9
    int TemporarySize; // r15d
    int ArrayNum; // edx
    FScriptContainerElement *v17; // rbx
    int v18; // ebx
    size_t UncompressedSize; // rbx
    TArray<unsigned char,TSizedDefaultAllocator<32> > ObjectBytes; // [rsp+40h] [rbp-29h] BYREF
    TArray<unsigned char,TSizedDefaultAllocator<32> > ObjectBytesTemporary; // [rsp+50h] [rbp-19h] BYREF
    TArray<unsigned char,TSizedDefaultAllocator<32> > ObjectBytesComped; // [rsp+60h] [rbp-9h] BYREF
    UPalSaveGameManager::FSaveDataHeader Header; // [rsp+70h] [rbp+7h]

    ObjectBytesComped.AllocatorInstance.Data = 0i64;
    *(_QWORD *)&ObjectBytesComped.ArrayNum = 0i64;
    v4 = UGameplayStatics::LoadDataFromSlot(&ObjectBytesComped, SlotName, UserIndex);
    Data = ObjectBytesComped.AllocatorInstance.Data;
    if ( !v4
        || (unsigned __int64)ObjectBytesComped.ArrayNum < 0xC
        || *(_BYTE *)&ObjectBytesComped.AllocatorInstance.Data[8] != 'P'
        || *(_BYTE *)&ObjectBytesComped.AllocatorInstance.Data[9] != 'l'
        || *(_BYTE *)&ObjectBytesComped.AllocatorInstance.Data[10] != 'Z'
        || (unsigned __int8)(*(_BYTE *)&ObjectBytesComped.AllocatorInstance.Data[11] - 48) > 2u )
    {
        result->Key = 0i64;
        result->Value = InProcess;
        goto LABEL_32;
    }
    CompressedSize = ObjectBytesComped.ArrayNum - 12;
    v7 = ObjectBytesComped.AllocatorInstance.Data + 12;
    v8 = HIBYTE(*(_DWORD *)&ObjectBytesComped.AllocatorInstance.Data[8]);
    v9 = 0i64;
    *(_QWORD *)&Header.UncompressedSize = *(_QWORD *)ObjectBytesComped.AllocatorInstance.Data;
    ObjectBytes.AllocatorInstance.Data = 0i64;
    *(_QWORD *)&ObjectBytes.ArrayNum = 0i64;
    v10 = v8 - 48;
    if ( !v10 )
    {
        UncompressedSize = Header.UncompressedSize;
        ObjectBytes.ArrayNum = Header.UncompressedSize;
        if ( Header.UncompressedSize )
        {
            TArray<unsigned char,TSizedDefaultAllocator<32>>::ResizeGrow(
            (TArray<enum FRigVMTemplateArgument::ETypeCategory,TSizedDefaultAllocator<32> > *)&ObjectBytes,
            0);
            v9 = ObjectBytes.AllocatorInstance.Data;
        }
        memcpy_0(v9, v7, UncompressedSize);
        goto LABEL_30;
    }
    v11 = v10 - 1;
    if ( v11 )
    {
        if ( v11 == 1 )
        {
            v12 = Header.UncompressedSize;
            ObjectBytesTemporary.AllocatorInstance.Data = 0i64;
            ObjectBytesTemporary.ArrayMax = 0;
            ObjectBytesTemporary.ArrayNum = Header.TemporarySize;
            if ( Header.TemporarySize )
            {
                TArray<unsigned char,TSizedDefaultAllocator<32>>::ResizeGrow(
                    (TArray<enum FRigVMTemplateArgument::ETypeCategory,TSizedDefaultAllocator<32> > *)&ObjectBytesTemporary,
                    0);
                v9 = ObjectBytesTemporary.AllocatorInstance.Data;
            }
            v13 = FNameEntryId::FromValidEName((FNameEntryId *)&ObjectBytesTemporary, (EName)257);
            v14 = (char *)v7;
            TemporarySize = Header.TemporarySize;
            ObjectBytesTemporary.AllocatorInstance.Data = (FScriptContainerElement *)v13->Value;
            if ( !FCompression::UncompressMemory(
                (FName)ObjectBytesTemporary.AllocatorInstance.Data,
                v9,
                Header.TemporarySize,
                v14,
                CompressedSize,
                COMPRESS_None,
                0) )
                goto LABEL_13;
            ArrayNum = ObjectBytes.ArrayNum;
            ObjectBytes.ArrayNum += v12;
            if ( v12 > ObjectBytes.ArrayMax - ArrayNum )
                TArray<unsigned char,TSizedDefaultAllocator<32>>::ResizeGrow(
                    (TArray<enum FRigVMTemplateArgument::ETypeCategory,TSizedDefaultAllocator<32> > *)&ObjectBytes,
                    ArrayNum);
            v17 = ObjectBytes.AllocatorInstance.Data;
            ObjectBytesTemporary.AllocatorInstance.Data = (FScriptContainerElement *)FNameEntryId::FromValidEName(
                (FNameEntryId *)&ObjectBytesTemporary,
                (EName)257)->Value;
            if ( !FCompression::UncompressMemory(
                (FName)ObjectBytesTemporary.AllocatorInstance.Data,
                v17,
                v12,
                (char *)v9,
                TemporarySize,
                COMPRESS_None,
                0) )
            {
                LABEL_13:
                result->Key = 0i64;
                result->Value = Fail;
                if ( v9 )
                    FMemory::Free(v9);
                goto LABEL_15;
            }
            if ( v9 )
                FMemory::Free(v9);
        }
        goto LABEL_30;
    }
    v18 = Header.UncompressedSize;
    ObjectBytes.ArrayNum = Header.UncompressedSize;
    if ( Header.UncompressedSize )
    {
        TArray<unsigned char,TSizedDefaultAllocator<32>>::ResizeGrow(
            (TArray<enum FRigVMTemplateArgument::ETypeCategory,TSizedDefaultAllocator<32> > *)&ObjectBytes,
            0);
        v9 = ObjectBytes.AllocatorInstance.Data;
    }
    ObjectBytesTemporary.AllocatorInstance.Data = (FScriptContainerElement *)FNameEntryId::FromValidEName(
        (FNameEntryId *)&ObjectBytesTemporary,
        (EName)257)->Value;
    if ( FCompression::UncompressMemory(
        (FName)ObjectBytesTemporary.AllocatorInstance.Data,
        v9,
        v18,
        (char *)v7,
        CompressedSize,
        COMPRESS_None,
        0) )
    {
        LABEL_30:
        result->Key = UGameplayStatics::LoadGameFromMemory(&ObjectBytes);
        result->Value = NotRun;
        goto LABEL_15;
    }
    result->Key = 0i64;
    result->Value = Fail;
    LABEL_15:
    if ( ObjectBytes.AllocatorInstance.Data )
        FMemory::Free(ObjectBytes.AllocatorInstance.Data);
    Data = ObjectBytesComped.AllocatorInstance.Data;
    LABEL_32:
    if ( Data )
        FMemory::Free(Data);
    return result;
}
```

解压后的存档可以使用[uesave-rs](https://github.com/trumank/uesave-rs)解析并生成JSON文件

# 个人存档?
![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706160656505-19598b50-2b1e-41aa-81c2-3d4020cc3275.png#averageHue=%23fefefe&clientId=u41804305-9356-4&from=paste&height=807&id=ube19044a&originHeight=1211&originWidth=1015&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=105547&status=done&style=none&taskId=ue782ea3c-6220-4666-a8a1-0fd12d078cc&title=&width=676.6666666666666)

UID.sav内几乎没有储存玩家信息 只储存了对应的uuid

比如玩家ID 帕鲁盒ID 储存箱ID 等等 实际上这些数据都还是储存在Level.sav里面

现在 我们已经拿到了玩家ID

想要修改玩家数据还得是去看看Level.sav

# Level.sav
试了个多人游戏服务器的备份 


解压后体积巨大 差不多有140M
![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706160572090-ee5f45c0-acb5-4be7-8332-ae04d407ea88.png#averageHue=%23f6f5f4&clientId=u41804305-9356-4&from=paste&height=607&id=Ddh7Q&originHeight=910&originWidth=1338&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=89749&status=done&style=none&taskId=u94a2a476-98e0-4acc-b359-65cc3f5365b&title=&width=892)

最终生成了2GB的JSON文件 编辑它十分的困难

写了个工具遍历打印结构

![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706161001185-a089bcc7-ce24-4d6f-9eaa-3bbd00df60e0.png#averageHue=%23f0ece8&clientId=u41804305-9356-4&from=paste&height=585&id=gMx6U&originHeight=878&originWidth=1441&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=227775&status=done&style=none&taskId=uc17c78f5-de4e-44da-bd6e-b4be9edc1ea&title=&width=960.6666666666666)

玩家对象储存在

root.properties.worldSaveData.Struct.value.Struct.CharacterSaveParameterMap.Map.value[X]

这个数组内

这个结构也包含了NPC角色 (NPC UID为全0)

在此结构中 key.Struct.Struct.PlayerUId.Struct.value.Guid 为玩家id

value.Struct.Struct.RawData.Array.value.Base.Byte.Byte 为玩家数据

玩家数据是二进制序列化数据 储存在RawData中 

这个数据是和虚幻存档是一样的 只不过它没有存档头

![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706161745347-945b2449-632c-451c-b2ec-ab125fb5da2e.png#averageHue=%2334322c&clientId=u41804305-9356-4&from=paste&height=511&id=tledI&originHeight=766&originWidth=1181&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=148677&status=done&style=none&taskId=u71adda88-b55d-44c0-b7a8-70230c75f89&title=&width=787.3333333333334)

稍微魔改了下uesave-rs

[https://github.com/burpheart/uesave-rs/tree/palworld](https://github.com/burpheart/uesave-rs/tree/palworld)

成功解析了玩家数据

![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706161828325-c21f9249-1fe7-4831-866c-8fa315c01c14.png#averageHue=%23fcfbfa&clientId=u41804305-9356-4&from=paste&height=815&id=u6e8e8b1b&originHeight=1222&originWidth=931&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=74422&status=done&style=none&taskId=u0dcc94bb-a406-446c-bb99-35a1fa06fed&title=&width=620.6666666666666)


![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706164986717-c0c45149-0857-4780-92e9-238cd67333b2.png#averageHue=%23998167&clientId=u7aabedb4-4aae-4&from=paste&height=316&id=ud1b03fc6&originHeight=474&originWidth=2060&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=61880&status=done&style=none&taskId=ud8ae4877-0235-4eca-9923-0f51b961df5&title=&width=1373.3333333333333)

之后就可以愉快的修改玩家数据了

整个修改过程

修改玩家数据->序列化玩家数据->替换回原位->序列化世界数据->压缩后替换原存档 ( 替换存档要提前关闭服务器 记得完整备份存档  有可能修改失败 或者是其他未知影响)


![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706162217578-31dfedb7-7868-49b6-a570-139763b21d5d.png#averageHue=%231d2a2b&clientId=u41804305-9356-4&from=paste&height=1067&id=u352995e4&originHeight=1600&originWidth=2560&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=4477381&status=done&style=none&taskId=u0d9eb851-aabc-4cdf-b146-d7176ff2d13&title=&width=1706.6666666666667)
![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706162425941-a05ed2cd-ecc7-466d-8863-c0a4c906aabe.png#averageHue=%2316222c&clientId=u41804305-9356-4&from=paste&height=637&id=ub9f4828f&originHeight=955&originWidth=1477&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=1846280&status=done&style=none&taskId=u31f0dfe1-42cf-4785-b39c-3f87c3ea614&title=&width=984.6666666666666)![图片.png](https://cdn.nlark.com/yuque/0/2024/png/25577536/1706162666223-db5ae191-d856-4912-bcb4-d1d4a6cbcb1e.png#averageHue=%234a3429&clientId=u41804305-9356-4&from=paste&height=1067&id=u12ec2088&originHeight=1600&originWidth=2560&originalType=binary&ratio=1.5&rotation=0&showTitle=false&size=5704653&status=done&style=none&taskId=ucaa65aa4-027e-41c8-b74a-92f6dfb36f7&title=&width=1706.6666666666667)
# 坏档修复

既然可以自由的编辑玩家存档 那么坏档也有一定可能修复

TODO //还没有遇到坏档 手里没有样本

# 附录
| 配置项名 | 描述 |
| --- | --- |
| CharacterMaxLevel | 角色最大等级 |
| GuildCharacterMaxLevel | 公会角色最大等级 |
| OtomoLevelSyncAddMaxLevel | Otomo等级同步增加最大等级 |
| PlayerSprintStaminaType | 玩家冲刺耐力类型 |
| CharacterMaxRank | 角色最大排名 |
| ConsumStamina_PalThrow | 消耗耐力_帕鲁投掷 |
| ReturnOtomoPalCoolTime | 返回Otomo帕鲁冷却时间 |
| OtomoSlotNum | Otomo插槽数量 |
| OtomoWazaCoolDownSpeedRate | Otomo技能冷却速度比率 |
| PlayerHPRateFromRespawn | 从重生开始的玩家HP比率 |
| PlayerStomachRateFromRespawn | 从重生开始的玩家胃口比率 |
| RarePal_AppearanceProbability | 稀有帕鲁出现概率 |
| RarePal_LevelMultiply | 稀有帕鲁等级乘数 |
| BossOrRarePal_TalentMin | Boss或稀有帕鲁最小天赋 |
| CharacterRankUpRequiredNumDefault | 角色升级所需数量默认值 |
| CharacterRankUpRequiredNumMap | 角色升级所需数量映射 |
| NaturalUpdateSaveParameterInterval | 自然更新保存参数间隔 |
| CharacterHUDDisplayRange | 角色HUD显示范围 |
| CharacterFedEatingTime | 角色饱食度吃食时间 |
| CharacterStayingSecondsInSpa | 角色在温泉中停留的秒数 |
| PalBoxPageNum | 帕鲁盒子页数 |
| PalBoxSlotNumInPage | 每页帕鲁盒子插槽数量 |
| PlayerBattleJudge_EnemyDistance | 玩家战斗判断_敌人距离 |
| BodyTemperature_SlipDamage_Percent | 身体温度_滑动伤害百分比 |
| BodyTemperature_StomachDecreaceRate | 身体温度_胃口减少率 |
| SleepFXDefault | 睡眠特效默认值 |
| LiftupCharacterThrownVelocityScalar | 提升角色投掷速度标量 |
| LiftupCharacterClearCollisionDelayTime | 提升角色清除碰撞延迟时间 |
| NickNameMaxLength | 昵称最大长度 |
| IntervalForPalAttackFromBall | 从球开始的帕鲁攻击间隔 |
| CanShootRiderByFullRide | 可以通过全程骑行射击骑手 |
| HitWazaAttackForMapObject | 对地图对象进行HitWaza攻击 |
| KnockBack_MaxHpPercent | 击退_最大HP百分比 |
| KnockBack_Power | 击退_力量 |
| StunTime | 昏迷时间 |
| StepCooldownTime | 步骤冷却时间 |
| Stun_GunDamageRate | 昏迷_枪伤害率 |
| IsEnableAutoReload | 是否启用自动重载 |
| DeadShootImpulseRate | 死亡射击冲量比率 |
| DeadShootImpulseMax | 死亡射击冲量最大值 |
| OtomoDamageRate_Defense | Otomo伤害率_防御 |
| DamageValueMin_MapObject | 伤害值最小_地图对象 |
| DamageRate_WealPoint | 伤害率_弱点 |
| DamageRate_StrongPoint | 伤害率_强点 |
| DamageRate_SleepHit | 伤害率_睡眠命中 |
| FinalDamageRate_Waza | 最终伤害率_技能 |
| FinalDamageRate_Weapon | 最终伤害率_武器 |
| FinalDamageRate_Mine | 最终伤害率_矿石 |
| FinalDamageRate_MapObject_Waza | 最终伤害率_地图对象_技能 |
| FinalDamageRate_MapObject_Weapon | 最终伤害率_地图对象_武器 |
| FinalDamageRate_MapObject_Mine | 最终伤害率_地图对象_矿石 |
| FoliageDefault_Defense | 植被默认_防御 |
| WazaAttackerLevelDamage_Power | 技能攻击者等级伤害_力量 |
| DamageRandomRate_Min | 伤害随机率_最小 |
| DamageRandomRate_Max | 伤害随机率_最大 |
| LevelDamageCorrect | 等级伤害修正 |
| LevelStatusAccumulateCorrect | 等级状态累积修正 |
| DamageElementMatchRate | 伤害元素匹配率 |
| MineAttack_DefensePower | 矿石攻击_防御力 |
| StatusCalculate_LevelMultiply_HP | 状态计算_等级乘数_HP |
| StatusCalculate_TribePlus_HP | 状态计算_部落加_HP |
| StatusCalculate_ConstPlus_HP | 状态计算_常数加_HP |
| StatusCalculate_LevelMultiply_Attack | 状态计算_等级乘数_攻击 |
| StatusCalculate_ConstPlus_Attack | 状态计算_常数加_攻击 |
| StatusCalculate_LevelMultiply_Defense | 状态计算_等级乘数_防御 |
| StatusCalculate_ConstPlus_Defense | 状态计算_常数加_防御 |
| StatusCalculate_TribeMultiply_CraftSpeed | 状态计算_部落乘数_工艺速度 |
| StatusCalculate_GenkaiToppa_PerAdd | 状态计算_极限突破_百分比增加 |
| StatusCalculate_Talent_PerAdd | 状态计算_天赋_百分比增加 |
| BreakedWeaponDamageRate | 破损武器伤害率 |
| BreakedArmorDefenseRate | 破损护甲防御率 |
| ArmorDurabilityDamageDivide | 护甲耐久度伤害分割 |
| ClimbingStamina_Move | 攀爬耐力_移动 |
| ClimbingStamina_Jump | 攀爬耐力_跳跃 |
| RideWazaStaminaRate | 骑行技能耐力比率 |
| bIsEnableJumpPreliminary | 是否启用跳跃预备 |
| JumpInterval | 跳跃间隔 |
| FlyMaxHeight | 飞行最大高度 |
| FlyHover_SP | 飞行悬停_SP |
| FlyHorizon_SP | 飞行水平_SP |
| FlyHorizon_Dash_SP | 飞行水平_冲刺_SP |
| FlyVertical_SP | 飞行垂直_SP |
| AimingSpeedRateInRide | 骑行中的瞄准速度比率 |
| SlidingEndSpeed | 滑行结束速度 |
| JumpSP | 跳跃SP |
| StepSP | 步骤SP |
| MeleeAttackSP | 近战攻击SP |
| SprintSP | 冲刺SP |
| GliderSP | 滑翔SP |
| SwimmingFallWaitTimeSec | 游泳下降等待时间秒 |
| Swimming_SP_Idle | 游泳_SP_空闲 |
| Swimming_SP_Swim | 游泳_SP_游泳 |
| Swimming_SP_DashSwim | 游泳_SP_冲刺游泳 |
| FluidFriction | 流体摩擦 |
| OverWeightSpeedZero_AddPercent | 超重速度零_增加百分比 |
| WalkableFloorAngleForDefault | 默认可行走地板角度 |
| WalkableFloorAngleForRide | 骑行可行走地板角度 |
| IsEnableSpeedCollision | 是否启用速度碰撞 |
| CollisionDamageMinSpeed | 碰撞伤害最小速度 |
| SpeedCollisionDamagePower | 速度碰撞伤害力量 |
| CollisionDamageSpeedMultiplay | 碰撞伤害速度乘数 |
| CollisionDamageWeightThreshold | 碰撞伤害重量阈值 |
| AutoHPRegene_Percent_perSecond | 自动HP再生_百分比_每秒 |
| AutoHPRegene_Percent_perSecond_Sleeping | 自动HP再生_百分比_每秒_睡眠 |
| PosionHPDecrease_Percent_perSecond | 毒药HP减少_百分比_每秒 |
| Starvation_DecreaseHP_Percent_perSecond | 饥饿_减少HP_百分比_每秒 |
| AutoSANRegene_Percent_perSecond_PalStorage | 自动SAN再生_百分比_每秒_帕鲁存储 |
| StomachDecreace_perSecond_Monster | 胃口减少_每秒_怪物 |
| StomachDecreace_perSecond_Player | 胃口减少_每秒_玩家 |
| StomachDecreace_AutoHealing | 胃口减少_自动治疗 |
| StomachDecreace_WorkingRate | 胃口减少_工作率 |
| HungerStart_StomachValue | 饥饿开始_胃口值 |
| FullStomachPalStartEatFood | 饱食度帕鲁开始吃食物 |
| FullStomachCost_ByWazaUse_Base | 饱食度消耗_通过技能使用_基础 |
| FullStomachCost_ByWazaUse_RateMap | 使用技能时的饱腹度消耗 |
| StomachDecreaceRate_GroundRide_Sprint | 地面骑行冲刺时的饱腹度减少率 |
| StomachDecreaceRate_WaterRide | 水上骑行时的饱腹度减少率 |
| StomachDecreaceRate_WaterRide_Sprint | 水上骑行冲刺时的饱腹度减少率 |
| StomachDecreaceRate_FlyRide | 飞行骑行时的饱腹度减少率 |
| StomachDecreaceRate_FlyRide_Sprint | 飞行骑行冲刺时的饱腹度减少率 |
| RemainderOfLife_Second | 生命剩余时间（秒） |
| HpDecreaseRate_Drowning | 溺水时的生命值减少率 |
| PlayerShield_RecoverStartTime | 玩家护盾恢复开始时间 |
| PlayerShield_RecoverPercentPerSecond | 玩家护盾每秒恢复百分比 |
| StaminaRecover_PercentPerSecond | 体力每秒恢复百分比 |
| ResuscitationTime | 复活时间 |
| PlayerDeath_DropOtomoNum | 玩家死亡时掉落的Otomo数量 |
| PlayerDeath_DropOtomoRange | 玩家死亡时掉落的Otomo范围 |
| PlayerDeath_DropOtomoDisappearHours | 玩家死亡后掉落的Otomo消失时间（小时） |
| PlayerDyingDamagePerTime | 玩家濒死时每次受到的伤害 |
| ElementStatus_ResistanceInitialValue | 元素状态的初始抵抗值 |
| ElementStatus_AutoDecreasePerSecond | 元素状态每秒自动减少 |
| ElementStatus_ResetResistanceSecond | 元素状态重置抵抗值的时间（秒） |
| BuildExp | 建造经验 |
| CraftExp | 制作经验 |
| PickupItemOnLevelExp | 拾取物品的等级经验 |
| MapObjectDestroyProceedExp | 地图对象销毁进度经验 |
| MapObjectDistributeExpRange | 地图对象分布经验范围 |
| OtomoExp_LevelDifferenceMap | Otomo经验等级差异图 |
| OtomoExp_HigherPlayerLevel | Otomo经验高于玩家等级 |
| CaptureExpBonus_Tier1_TableReferenceNum | 捕获经验奖励Tier1表引用数 |
| CaptureExpBonus_Tier2_TableReferenceNum | 捕获经验奖励Tier2表引用数 |
| CaptureExpBonus_Tier3_TableReferenceNum | 捕获经验奖励Tier3表引用数 |
| NewGameOtomoPalSet | 新游戏Otomo帕鲁设置 |
| NewGameInventoryItemSet | 新游戏库存物品设置 |
| NewGameLoadoutItemSet | 新游戏装备物品设置 |
| WorldHUDDisplayOffsetDefault | 世界HUD显示默认偏移 |
| WorldHUDDisplayRangeDefault | 世界HUD显示默认范围 |
| WorldHUDDetailDisplayRange | 世界HUD详细显示范围 |
| FarmCropWaterItemIds | 农作物水物品ID |
| FarmCropGrowupSpeedBySec | 农作物每秒生长速度 |
| FarmCropIncreaseRateByWaterFillRate | 农作物由水填充率增加的速度 |
| MaxMoney | 最大金钱 |
| DefaultMoney | 默认金钱 |
| SneakAttackBackJudgeAngle_Degree | 潜行攻击背部判断角度（度） |
| SneakAttack_PalMeleeWaza_AttackRate | 潜行攻击帕鲁近战技能攻击率 |
| AutoAimCameraMoveRate | 自动瞄准摄像机移动速度 |
| AutoAimCharacterMoveRate | 自动瞄准角色移动速度 |
| AutoAimCameraAdsorptionSpeed | 自动瞄准摄像机吸附速度 |
| AutoAimLockOnScreenSpaceRate | 自动瞄准锁定屏幕空间比率 |
| ForceAutoAimTime | 强制自动瞄准时间 |
| SellItemRate | 卖出物品的比率 |
| PalPriceConstantValueA | 帕鲁价格常数值A |
| PalPriceConstantValueB | 帕鲁价格常数值B |
| SellPalRate | 卖出帕鲁的比率 |
| SearchRangeOnThrowedCharacterLanded | 投掷角色着陆时的搜索范围 |
| WorkCompleteReactionRangeFromPlayer | 从玩家开始的工作完成反应范围 |
| WorkerCollectResourceStackMaxNum | 工人收集资源堆叠最大数量 |
| FacialTypeHardWork | 辛勤工作的面部类型 |
| Timeout_WorkerApproachToTarget | 工人接近目标的超时时间 |
| WaitTime_WorkRepairFailedFindPath | 工人修复失败找到路径的等待时间 |
| WorkerWaitingNotifyInterval | 工人等待通知间隔 |
| WarpCheckInterval | 瞬移检查间隔 |
| WarpCheckMoveDistanceThreshold | 瞬移检查移动距离阈值 |
| WarpThreshold | 瞬移阈值 |
| AutoDecreaseHateValue_PercentMaxHP_PerSecond | 每秒自动减少仇恨值_最大生命值百分比 |
| HateDecreaseDamageRate | 仇恨减少伤害率 |
| Hate_ForceUp_HPRate_OtomoActive | 仇恨_强制提升_生命值比率_Otomo活跃 |
| Hate_ForceUp_HPRate_IncidentBattle | 仇恨_强制提升_生命值比率_事件战斗 |
| CombatEndDistance_BattleStartSelfPos_To_SelfPos | 战斗结束距离_战斗开始自身位置_到_自身位置 |
| CombatEndDistance_BattleStartSelfPos_To_TargetPos | 战斗结束距离_战斗开始自身位置_到_目标位置 |
| CombatEndDistance_BattleStartSelfPos_To_TargetPos_AddFirstTargetDistance | 战斗结束距离_战斗开始自身位置_到_目标位置_加上首个目标距离 |
| NavigationAreaDivideExtents | 导航区域划分范围 |
| NavigationUpdateFrequencySettingsFromPlayer | 从玩家开始的导航更新频率设置 |
| AutoSaveSpan | 自动保存间隔 |
| SaveDataName_WorldBaseInfo | 保存数据名_世界基本信息 |
| SaveDataName_World | 保存数据名_世界 |
| SaveDataName_PlayerDirectory | 保存数据名_玩家目录 |
| SaveDataName_LocalData | 保存数据名_本地数据 |
| SaveDataName_WorldOption | 保存数据名_世界选项 |
| MaxWorldDataNumMap | 最大世界数据数量图 |
| PalWorldTime_GameStartHour | 帕鲁世界时间_游戏开始小时 |
| PalWorldMinutes_RealOneMinute | 帕鲁世界分钟_真实一分钟 |
| NightStartHour | 夜晚开始小时 |
| NightEndHour | 夜晚结束小时 |
| PlayerMorningHour | 玩家早晨小时 |
| PlayerSleepStartHour | 玩家睡眠开始小时 |
| NightSkipWaitSecond | 跳过夜晚等待秒数 |
| LocalPlayerAndOtomo_LightRangeScale | 本地玩家和Otomo_光范围比例 |
| BuildBaseUnitGridDefinition | 建造基本单位网格定义 |
| BuildSimulationVerticalAdjustRate | 建造模拟垂直调整率 |
| BuildSimulationFoundationFloatingAllowance | 建造模拟基础浮动容许值 |
| BuildSimulationFoundationCheckCollisionScale | 建造模拟基础检查碰撞规模 |
| BuildSimulationRoofHeightOffset | 建造模拟屋顶高度偏移 |
| BuildSimulationStairHeightOffset | 建造模拟楼梯高度偏移 |
| BuildSimulationLeanAngleMax | 建造模拟倾斜角度最大值 |
| BuildingProgressInterpolationSpeed | 建筑进度插值速度 |
| PlayerRecord_BuildingObjectMaxNum | 玩家记录_建筑物最大数量 |
| BuildingMaxZ | 建筑最大Z值 |
| BuildObj_HatchedPalCharacterLevel | 孵化的帕鲁角色等级 |
| BuildObj_DamageScarecrowStartRecoveryTime | 伤害稻草人开始恢复时间 |
| BaseCampAreaRange | 基地区域范围 |
| BaseCampPalFindWorkRange | 帕鲁在基地寻找工作的范围 |
| PalArriveToWorkLocationRange | 帕鲁到达工作地点的范围 |
| PalArriveToWorkLocationRangeZ | 帕鲁到达工作地点的Z轴范围 |
| BaseCampNeighborMinimumDistance | 基地邻居最小距离 |
| PalRotateSpeedToWork | 帕鲁旋转到工作的速度 |
| BaseCampFoliageWorkableRange | 基地植被可工作范围 |
| BaseCampHungerApproachToPlayer | 基地饥饿接近玩家 |
| BaseCampHungerUnreachableObjectTimeoutRealSeconds | 基地饥饿无法到达物体超时实际秒数 |
| HungerHUDDisplayRange | 饥饿HUD显示范围 |
| WorkAmountBySecForPlayer | 玩家每秒工作量 |
| BaseCampWorkerEventTriggerInterval | 基地工人事件触发间隔 |
| BaseCampWorkerEventTriggerProbability | 基地工人事件触发概率 |
| BaseCampWorkerSanityWarningThreshold | 基地工人理智警告阈值 |
| BaseCampWorkerFinishEatingFullStomach | 基地工人完成吃饭后的饱腹感 |
| BaseCampWorkerFinishEatingSanity | 基地工人完成吃饭后的理智值 |
| BaseCampWorkerFinishEatCount | 基地工人完成吃饭的次数 |
| BaseCampWorkerRecoverHungryTurnToTargetTimeout | 基地工人恢复饥饿转向目标超时 |
| BaseCampWorkerStartSleepHpPercentage | 基地工人开始睡觉的HP百分比 |
| BaseCampWorkerSleepInPlaceRecoverSanityRate | 基地工人在原地睡觉恢复理智率 |
| BaseCampWorkerDistancePickableItem | 基地工人距离可拾取物品的距离 |
| BaseCampBuildingItemContainerPriority | 基地建筑物品容器优先级 |
| FoliageRespawnFailedExtraRangeOfBaseCamp | 基地植被重生失败额外范围 |
| BaseCampPalCombatRange_AddCampRange | 基地帕鲁战斗范围_增加营地范围 |
| BaseCampPalWalkTime_BeforeSleep | 帕鲁在基地行走时间_睡前 |
| BaseCampTimeFinishBattleModeAfterEmptyEnemy | 基地完成战斗模式后的时间_空敌 |
| BaseCampWorkerMoveModeChangeThreshold | 基地工人移动模式改变阈值 |
| BaseCampWorkerDirectorTickForAssignWorkByCount | 基地工人导演为分配工作计数的Tick |
| BaseCampWorkerLookToTargetWork | 基地工人看向目标工作 |
| ReviveWorkAdditionalRange | 复活工作额外范围 |
| WorkAroundRangeDefault | 工作周围范围默认 |
| IssueNotifyWorkTypes | 发布通知工作类型 |
| WorkAmountByManMonth | 按人月计算的工作量 |
| WorkNotifyDelayTime | 工作通知延迟时间 |
| WorkFinishDelayCallAddWorkNotifyDelayTime | 工作完成延迟调用添加工作通知延迟时间 |
| WorkIgnitionTorchWaitTime | 工作点火火炬等待时间 |
| WorkAssignFailedLogTypeMap | 工作分配失败日志类型图 |
| WorkTransportingSpeedRate | 工作运输速度率 |
| BaseCampNotTransportItemBlackList | 基地不运输物品黑名单 |
| WorkTransportingDelayTimeDropItem | 工作运输延迟时间掉落物品 |
| BaseCampStopProvideEnergyInterval | 基地停止提供能量间隔 |
| BaseCampPassiveEffectWorkHardInfoMap | 基地被动效果努力工作信息图 |
| BaseCampWorkCollectionRestoreStashSeconds | 基地工作收集恢复储藏秒数 |
| WorkTypeAssignPriorityOrder | 工作类型分配优先级顺序 |
| WorkAssignDefineData_Build | 工作分配定义数据_建造 |
| WorkAssignDefineData_FoliageWork | 工作分配定义数据_植被工作 |
| WorkAssignDefineData_ReviveCharacterWork | 工作分配定义数据_复活角色工作 |
| WorkAssignDefineData_TransportItemInBaseCamp | 工作分配定义数据_在基地运输物品 |
| WorkAssignDefineData_RepairBuildObjectInBaseCamp | 工作分配定义数据_在基地修理建筑物 |
| WorkAssignDefineData_BreedFarm | 工作分配定义数据_养殖农场 |
| WorkAssignDefineData_ExtinguishBurn | 工作分配定义数据_熄灭燃烧 |
| WorkSuitabilityMaxRank | 工作适应性最大等级 |
| WorkSuitabilityDefineDataMap | 工作适应性定义数据图 |
| WorkSuitabilityDefineData_Collection | 工作适应性定义数据_收集 |
| WorkSuitabilityDefineData_Deforest | 工作适应性定义数据_砍伐 |
| WorkSuitabilityDefineData_Mining | 工作适应性定义数据_采矿 |
| DropItemWaitInsertMaxNumPerTick | 掉落物品等待插入每Tick的最大数量 |
| DungeonSpawnParameterDefault | 地牢生成参数默认 |
| GamePad_NotAimCameraRotateSpeed_DegreePerSecond | 游戏手柄_非瞄准相机旋转速度_每秒度数 |
| GamePad_AimCameraRotateSpeed_DegreePerSecond | 游戏手柄_瞄准相机旋转速度_每秒度数 |
| Mouse_NotAimCameraRotateSpeed | 鼠标_非瞄准相机旋转速度 |
| Mouse_AimCameraRotateSpeed | 鼠标_瞄准相机旋转速度 |
| YawCameraMaxSpeedRate | 偏航相机最大速度率 |
| TimeForCameraMaxSpeed | 相机最大速度的时间 |
| AimInterpInterval | 瞄准插值间隔 |
| InvaderSelfDeleteAddTime | 入侵者自我删除添加时间 |
| InvadeProbability | 入侵概率 |
| InvadeOccurablePlayerLevel | 可发生入侵的玩家等级 |
| InvadeJudgmentInterval_Minutes | 入侵判断间隔_分钟 |
| InvadeCollTime_Max_Minutes | 入侵碰撞时间_最大分钟 |
| InvadeCollTime_Min_Minutes | 入侵碰撞时间_最小分钟 |
| InvadeReturnTime_Minutes | 入侵返回时间_分钟 |
| InvadeStartPoint_BaseCampRadius_Min_cm | 入侵起点_基地半径_最小厘米 |
| InvadeStartPoint_BaseCampRadius_Max_cm | 入侵起点_基地半径_最大厘米 |
| VisitorNPCProbability | 访客NPC概率 |
| VisitorNPCReturnTime_Minutes | 访客NPC返回时间_分钟 |
| RidingAimOpacity | 骑行瞄准不透明度 |
| HideUITimeWhenNotConflict | 非冲突时隐藏UI时间 |
| FirstCapturedUIDisplayTime | 首次捕获UI显示时间 |
| CapturedUIDisplayTime | 捕获UI显示时间 |
| FirstActivatedOtomoInfoDisplayTime | 首次激活Otomo信息显示时间 |
| PlayerLevelUpUIDIsplayTime | 玩家升级UI显示时间 |
| PlayerExpGaugeUIDisplayTime | 玩家经验条UI显示时间 |
| OtomoExpGaugeUIDisplayTime | Otomo经验条UI显示时间 |
| NpcGaugeDisplayDistance | NPC计量器显示距离 |
| NpcGaugeDisplayRange_CameraSight | NPC计量器显示范围_相机视线 |
| GuildMemberGaugeDisplayDIstance | 公会成员计量器显示距离 |
| DownPlayerLoupeDisplayDistance | 下降玩家放大镜显示距离 |
| DownPlayerGaugeDisplayRange_CameraSight | 下降玩家计量器显示范围_相机视线 |
| ReticleOffsetRate | 准星偏移率 |
| LowHealthEffectParcent | 低健康效果百分比 |
| DamageTextMargineMap | 伤害文本边距图 |
| DamageTextDisplayLength | 伤害文本显示长度 |
| DamageTextMaxOffset | 伤害文本最大偏移 |
| DamageTextOffsetInterpolationLength | 伤害文本偏移插值长度 |
| DamageTextScaleMap | 伤害文本缩放图 |
| DamageTextRandomOffset | 伤害文本随机偏移 |
| StrongEnemyMarkLevel | 强敌标记等级 |
| OtomoInteractUIDisplayDistance | Otomo交互UI显示距离 |
| EnemyMarkUIMinScale | 敌人标记UI最小缩放 |
| EnemyMarkScaleInterpolationLength | 敌人标记缩放插值长度 |
| NpcHPGaugeGlobalOffset | NPC血条全局偏移 |
| DelayGaugeStartTime | 延迟计量器开始时间 |
| DelayGaugeProgressPerSecond | 延迟计量器每秒进度 |
| InventoryWeaponRangeMaxBorder | 库存武器范围最大边界 |
| InventoryWeaponStabilityMinBorder | 库存武器稳定性最小边界 |
| InventoryWeaponAccuracyMinBorder | 库存武器精度最小边界 |
| WorldmapUIMaskClearSize | 世界地图UI遮罩清除大小 |
| WorldmapUIFTMergeDistance | 世界地图UI FT合并距离 |
| WorldmapUIMaxMarker | 世界地图UI最大标记 |
| NPCHPGaugeUpdateSpan | NPC血条更新跨度 |
| CaptureFailedUIDisplayTime | 捕获失败UI显示时间 |
| CaptureSphereSortArray | 捕获球排序数组 |
| OpenGameOverUITime | 打开游戏结束UI时间 |
| BlockRespawnTime | 阻止重生时间 |
| InventoryWeightAlertRate | 库存重量警报率 |
| InventoryWeightGaugeDIsplayTime | 库存重量计量器显示时间 |
| OtomoLevelUpNoticeUIDisplayTime | Otomo升级通知UI显示时间 |
| OtomoMasteredWazaNoticeUIDisplayTime | Otomo掌握技术通知UI显示时间 |
| ProgressGaugeInterpolationSpeed | 进度计量器插值速度 |
| TeleportFadeInTime | 传送淡入时间 |
| TeleportFadeOutTime | 传送淡出时间 |
| PlayerTeleportTimeoutTime | 玩家传送超时时间 |
| PassiveSkillAppendNumWeights | 被动技能附加数权重 |
| bIsEggLauncherExplosion | 是否是蛋发射器爆炸 |
| ThrowPalBattleRadius | 投掷帕鲁战斗半径 |
| ThrowPalWorkRadius | 投掷帕鲁工作半径 |
| RopeHitPowe | 绳索命中力量 |
| RopePullPower | 绳索拉力 |
| DefaultMaxInventoryWeight | 默认最大库存重量 |
| RaycastLengthForDetectIndoor | 检测室内的射线长度 |
| MapObjectConnectAnyPlaceRaycastLength | 地图对象连接任何地方的射线长度 |
| ShootingTargetRayCastDistance | 射击目标射线距离 |
| CaptureJudgeRateArray | 捕获判断率数组 |
| CaptureBallBoundCountMax | 捕获球弹跳次数最大值 |
| ExceptCapturedItemList | 捕获排除列表 |
| CaptureSphereLevelMap | 捕获球等级图 |
| CaptureRateAddByStatusMap | 捕获率增加状态图 |
| IgnoreFirstCaptureFailedHPRate | 忽略首次捕获失败的HP率 |
| CaptureRateAdd_ByLegHold | 捕获率增加_腿部固定 |
| LongPressInterval | 长按间隔 |
| LongPressInterval_EnemyCampCage | 长按间隔_敌人营地笼子 |
| LongPressInterval_GetHatchedPal | 长按间隔_获取孵化的帕鲁 |
| CrouchLockAttenuation | 蹲伏锁定衰减 |
| IsEnableCharacterWazaScale | 是否启用角色技术规模 |
| IsOverrideDamageAdditiveAnimation | 是否覆盖伤害添加动画 |
| BlinkInterval | 闪烁间隔 |
| CrimeStateMaintainDurationBaseDefault | 犯罪状态维持持续时间基础默认 |
| TechnologyPointPerLevel | 每级技术点 |
| bossTechnologyPointPerTowerBoss | 塔防boss的boss技术点 |
| bossTechnologyPointPerNormalBoss | 普通boss的boss技术点 |
| DefaultUnlockTechnology | 默认解锁技术 |
| DefaultTechnologyPoint | 默认技术点 |
| TechnologyPoint_UnlockFastTravel | 技术点_解锁快速旅行 |
| DecreaseSanity_DamagedMultiply | 减少理智_受损乘数 |
| FullStomachPercent_RecoverySanity | 全胃百分比_恢复理智 |
| RecoverySanity_FullStomach | 恢复理智_全胃 |
| DecreaseSanity_Hunger | 减少理智_饥饿 |
| DecreaseSanity_Starvation | 减少理智_饥饿 |
| Spawner_IsCheckLoadedWorldPartition | 生成器_检查加载的世界分区 |
| SpawnerDisableDistanceCM_FromBaseCamp | 生成器禁用距离CM_从基地 |
| Spawner_DefaultSpawnRadius_S | 生成器_默认生成半径_S |
| Spawner_DefaultSpawnRadius_M | 生成器_默认生成半径_M |
| Spawner_DefaultSpawnRadius_L | 生成器_默认生成半径_L |
| Spawner_DefaultSpawnRadius_NPC | 生成器_默认生成半径_NPC |
| Spawner_DefaultDespawnDistance_S | 生成器_默认消失距离_S |
| Spawner_DefaultDespawnDistance_M | 生成器_默认消失距离_M |
| Spawner_DefaultDespawnDistance_L | 生成器_默认消失距离_L |
| Spawner_DefaultDespawnDistance_NPC | 生成器_默认消失距离_NPC |
| CharacterHeadMeshDataTable | 角色头部网格数据表 |
| CharacterBodyMeshDataTable | 角色身体网格数据表 |
| CharacterHairMeshDataTable | 角色头发网格数据表 |
| CharacterEquipmentArmorMeshDataTable | 角色装备护甲网格数据表 |
| CharacterEyeMaterialDataTable | 角色眼睛材质数据表 |
| CharacterMakeColorLimit_SV | 角色制作颜色限制_SV |
| IsAutoEquipMasteredWaza | 是否自动装备掌握的技术 |
| ActiveUNKO | 激活UNKO(便便) |
| MaxSpawnableDeathPenaltyChest | 最大可生成的死亡惩罚箱子 |
| BuildObjectInstallStrategy_SinkAllowCollisionPresetName | 构建对象安装策略_允许碰撞预设名称 |
| MapObjectShakeTimeOnDamaged | 地图对象在受损时的震动时间 |
| MapObjectShakeOffsetOnDamaged | 地图对象在受损时的震动偏移 |
| MapObjectOutlineByReticleTargetting | 通过准星目标定位的地图对象轮廓 |
| MapObjectOutlineByInteractable | 通过可交互的地图对象轮廓 |
| MapObjectRepairInfo | 地图对象修复信息 |
| FoliageExtentsXY | 植被范围XY |
| FoliageChunkSeparateScale | 植被块分离比例 |
| MapObjectHPDisplayDistance | 地图对象HP显示距离 |
| MapObjectHPDisplayTime | 地图对象HP显示时间 |
| MapObjectGateLockTime | 地图对象门锁定时间 |
| bDirectObtainFromTreasureBox | 是否直接从宝箱中获取 |
| MapObjectEffectTriggerAccumulate_Burn | 地图对象效果触发累积_燃烧 |
| MapObjectEffect_Burn_DamageHpRate | 地图对象效果_燃烧_伤害Hp率 |
| MapObjectEffect_Burn_DamageAroundRange | 地图对象效果_燃烧_周围伤害范围 |
| MapObjectEffect_Burn_DamageAroundInterval | 地图对象效果_燃烧_周围伤害间隔 |
| MapObjectEffect_Burn_DamageAroundDamageValue | 地图对象效果_燃烧_周围伤害值 |
| MapObjectEffect_Burn_DamageAroundAccumulateValue | 地图对象效果_燃烧_周围伤害累积值 |
| MapObjectEffect_Burn_DamageAroundAccumulateValue_ForCharacter | 地图对象效果_燃烧_角色周围伤害累积值 |
| PasswordLockFailedMaxNum | 密码锁失败最大次数 |
| MapObjectItemChestCorruptionRateFromWorkSpeed | 从工作速度开始的地图对象物品箱子腐败率 |
| RuntimeOptimizeParameter | 运行时优化参数 |
| WorldSecurityWantedPoliceSettingDataMap | 世界安全警察设置数据图 |
| WorldSecurityWantedPoliceSettingDataMapForDS | DS的世界安全警察设置数据图 |
| StatusPointPerLevel | 每级状态点 |
| AddMaxHPPerStatusPoint | 每状态点增加最大HP |
| AddMaxSPPerStatusPoint | 每状态点增加最大SP |
| AddPowerPerStatusPoint | 每状态点增加力量 |
| AddMaxInventoryWeightPerStatusPoint | 每状态点增加最大库存重量 |
| AddCaptureLevelPerStatusPoint | 每状态点增加捕获等级 |
| AddWorkSpeedPerStatusPoint | 每状态点增加工作速度 |
| AddMaxHPPerHPRank | 每HP等级增加最大HP |
| AddAttackPerAttackRank | 每攻击等级增加攻击 |
| AddDefencePerDefenceRank | 每防御等级增加防御 |
| AddWorkSpeedPerWorkSpeedRank | 每工作速度等级增加工作速度 |
| Combi_TalentInheritNum | 组合_才能继承数 |
| Combi_PassiveInheritNum | 组合_被动继承数 |
| Combi_PassiveRandomAddNum | 组合_被动随机添加数 |
| PalEggRankInfoArray | 帕鲁蛋等级信息数组 |
| PalEggMapObjectIdMap | 帕鲁蛋地图对象ID图 |
| PalEggHatchingSpeedRateByTemperature | 根据温度的帕鲁蛋孵化速度率 |
| DebugInfoFont | 调试信息字体 |
| MaxGuildNameLength | 最大公会名称长度 |
| JoinGuildRequestInteractLongPushTime | 加入公会请求互动长按时间 |
| TutorialMinDisplayTime | 教程最小显示时间 |
| TutorialDisplayTime | 教程显示时间 |
| CommonRewardDisplayTime | 常见奖励显示时间 |
| DeadBodyDestroySecond | 死体销毁秒数 |
| EnemyCampRespawnCoolTime | 敌人营地重生冷却时间 |
| EnemyCampDespawnDelayTime | 敌人营地消失延迟时间 |
| PalBoxReviveTime | 帕鲁盒子复活时间 |
| AfterNPCTalkDelayTime_Interact | NPC谈话后延迟时间_互动 |
| MinSprintThreshold | 最小冲刺阈值 |
| MaxSprintThreshold | 最大冲刺阈值 |
| MinHPGaugeDisplayTime | 最小HP计量器显示时间 |
| SoundSourceDataTable | 声音源数据表 |



