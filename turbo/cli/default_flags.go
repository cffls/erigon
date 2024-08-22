package cli

import (
	"github.com/urfave/cli/v2"

	"github.com/ledgerwatch/erigon/cmd/utils"
	"github.com/ledgerwatch/erigon/turbo/logging"
)

// DefaultFlags contains all flags that are used and supported by Erigon binary.
var DefaultFlags = []cli.Flag{
	&utils.DataDirFlag,
	&utils.EthashDatasetDirFlag,
	&utils.SnapshotFlag,
	&utils.ExternalConsensusFlag,
	&utils.TxPoolDisableFlag,
	&utils.TxPoolLocalsFlag,
	&utils.TxPoolNoLocalsFlag,
	&utils.TxPoolPriceLimitFlag,
	&utils.TxPoolPriceBumpFlag,
	&utils.TxPoolAccountSlotsFlag,
	&utils.TxPoolGlobalSlotsFlag,
	&utils.TxPoolGlobalBaseFeeSlotsFlag,
	&utils.TxPoolAccountQueueFlag,
	&utils.TxPoolGlobalQueueFlag,
	&utils.TxPoolLifetimeFlag,
	&utils.TxPoolTraceSendersFlag,
	&utils.TxPoolCommitEveryFlag,
	&PruneFlag,
	&PruneHistoryFlag,
	&PruneReceiptFlag,
	&PruneTxIndexFlag,
	&PruneCallTracesFlag,
	&PruneHistoryBeforeFlag,
	&PruneReceiptBeforeFlag,
	&PruneTxIndexBeforeFlag,
	&PruneCallTracesBeforeFlag,
	&BatchSizeFlag,
	&BodyCacheLimitFlag,
	&DatabaseVerbosityFlag,
	&PrivateApiAddr,
	&PrivateApiRateLimit,
	&EtlBufferSizeFlag,
	&TLSFlag,
	&TLSCertFlag,
	&TLSKeyFlag,
	&TLSCACertFlag,
	&StateStreamDisableFlag,
	&SyncLoopThrottleFlag,
	&BadBlockFlag,

	&utils.HTTPEnabledFlag,
	&utils.GraphQLEnabledFlag,
	&utils.HTTPListenAddrFlag,
	&utils.HTTPPortFlag,
	&utils.AuthRpcAddr,
	&utils.AuthRpcPort,
	&utils.JWTSecretPath,
	&utils.HttpCompressionFlag,
	&utils.HTTPCORSDomainFlag,
	&utils.HTTPVirtualHostsFlag,
	&utils.AuthRpcVirtualHostsFlag,
	&utils.HTTPApiFlag,
	&utils.WSEnabledFlag,
	&utils.WsCompressionFlag,
	&utils.HTTPTraceFlag,
	&utils.StateCacheFlag,
	&utils.RpcBatchConcurrencyFlag,
	&utils.RpcStreamingDisableFlag,
	&utils.DBReadConcurrencyFlag,
	&utils.RpcAccessListFlag,
	&utils.RpcTraceCompatFlag,
	&utils.RpcGasCapFlag,
	&utils.RpcBatchLimit,
	&utils.RpcReturnDataLimit,
	&utils.TxpoolApiAddrFlag,
	&utils.TraceMaxtracesFlag,
	&HTTPReadTimeoutFlag,
	&HTTPWriteTimeoutFlag,
	&HTTPIdleTimeoutFlag,
	&AuthRpcReadTimeoutFlag,
	&AuthRpcWriteTimeoutFlag,
	&AuthRpcIdleTimeoutFlag,
	&EvmCallTimeoutFlag,

	&utils.SnapKeepBlocksFlag,
	&utils.SnapStopFlag,
	&utils.DbPageSizeFlag,
	&utils.DbSizeLimitFlag,
	&utils.TorrentPortFlag,
	&utils.TorrentMaxPeersFlag,
	&utils.TorrentConnsPerFileFlag,
	&utils.TorrentDownloadSlotsFlag,
	&utils.TorrentStaticPeersFlag,
	&utils.TorrentUploadRateFlag,
	&utils.TorrentDownloadRateFlag,
	&utils.TorrentVerbosityFlag,
	&utils.ListenPortFlag,
	&utils.P2pProtocolVersionFlag,
	&utils.P2pProtocolAllowedPorts,
	&utils.NATFlag,
	&utils.NoDiscoverFlag,
	&utils.DiscoveryV5Flag,
	&utils.NetrestrictFlag,
	&utils.NodeKeyFileFlag,
	&utils.NodeKeyHexFlag,
	&utils.DNSDiscoveryFlag,
	&utils.BootnodesFlag,
	&utils.StaticPeersFlag,
	&utils.TrustedPeersFlag,
	&utils.MaxPeersFlag,
	&utils.ChainFlag,
	&utils.DeveloperPeriodFlag,
	&utils.VMEnableDebugFlag,
	&utils.NetworkIdFlag,
	&utils.FakePoWFlag,
	&utils.GpoBlocksFlag,
	&utils.GpoPercentileFlag,
	&utils.InsecureUnlockAllowedFlag,
	&utils.MetricsEnabledFlag,
	&utils.MetricsHTTPFlag,
	&utils.MetricsPortFlag,
	&utils.HistoryV3Flag,
	&utils.TransactionV3Flag,
	&utils.IdentityFlag,
	&utils.CliqueSnapshotCheckpointIntervalFlag,
	&utils.CliqueSnapshotInmemorySnapshotsFlag,
	&utils.CliqueSnapshotInmemorySignaturesFlag,
	&utils.CliqueDataDirFlag,
	&utils.MiningEnabledFlag,
	&utils.ProposingDisableFlag,
	&utils.MinerNotifyFlag,
	&utils.MinerGasLimitFlag,
	&utils.MinerEtherbaseFlag,
	&utils.MinerExtraDataFlag,
	&utils.MinerNoVerfiyFlag,
	&utils.MinerSigningKeyFileFlag,
	&utils.SentryAddrFlag,
	&utils.SentryLogPeerInfoFlag,
	&utils.SentryDropUselessPeers,
	&utils.DownloaderAddrFlag,
	&utils.DisableIPV4,
	&utils.DisableIPV6,
	&utils.NoDownloaderFlag,
	&utils.DownloaderVerifyFlag,
	&HealthCheckFlag,
	&utils.HeimdallURLFlag,
	&utils.WithoutHeimdallFlag,
	&utils.HeimdallgRPCAddressFlag,
	&utils.EthStatsURLFlag,
	&utils.OverrideShanghaiTime,

	&utils.ConfigFlag,
	&logging.LogConsoleVerbosityFlag,
	&logging.LogDirVerbosityFlag,
	&logging.LogDirPathFlag,
	&logging.LogConsoleJsonFlag,
	&logging.LogJsonFlag,
	&logging.LogDirJsonFlag,

	&utils.LightClientDiscoveryAddrFlag,
	&utils.LightClientDiscoveryPortFlag,
	&utils.LightClientDiscoveryTCPPortFlag,
	&utils.SentinelAddrFlag,
	&utils.SentinelPortFlag,
	&utils.YieldSizeFlag,

	&utils.L2ChainIdFlag,
	&utils.L2RpcUrlFlag,
	&utils.L2DataStreamerUrlFlag,
	&utils.L2DataStreamerTimeout,
	&utils.L1SyncStartBlock,
	&utils.L1SyncStopBatch,
	&utils.L1ChainIdFlag,
	&utils.L1RpcUrlFlag,
	&utils.L1CacheEnabledFlag,
	&utils.L1CachePortFlag,
	&utils.AddressSequencerFlag,
	&utils.AddressAdminFlag,
	&utils.AddressRollupFlag,
	&utils.AddressZkevmFlag,
	&utils.AddressGerManagerFlag,
	&utils.L1RollupIdFlag,
	&utils.L1BlockRangeFlag,
	&utils.L1QueryDelayFlag,
	&utils.L1HighestBlockTypeFlag,
	&utils.L1MaticContractAddressFlag,
	&utils.L1FirstBlockFlag,
	&utils.L1ContractAddressCheckFlag,
	&utils.RpcRateLimitsFlag,
	&utils.RpcGetBatchWitnessConcurrencyLimitFlag,
	&utils.DatastreamVersionFlag,
	&utils.RebuildTreeAfterFlag,
	&utils.IncrementTreeAlways,
	&utils.SmtRegenerateInMemory,
	&utils.SequencerBlockSealTime,
	&utils.SequencerBatchSealTime,
	&utils.SequencerBatchVerificationTimeout,
	&utils.SequencerHaltOnBatchNumber,
	&utils.SequencerResequence,
	&utils.SequencerResequenceStrict,
	&utils.SequencerResequenceReuseL1InfoIndex,
	&utils.ExecutorUrls,
	&utils.ExecutorStrictMode,
	&utils.ExecutorRequestTimeout,
	&utils.DatastreamNewBlockTimeout,
	&utils.WitnessMemdbSize,
	&utils.ExecutorMaxConcurrentRequests,
	&utils.Limbo,
	&utils.AllowFreeTransactions,
	&utils.AllowPreEIP155Transactions,
	&utils.EffectiveGasPriceForEthTransfer,
	&utils.EffectiveGasPriceForErc20Transfer,
	&utils.EffectiveGasPriceForContractInvocation,
	&utils.EffectiveGasPriceForContractDeployment,
	&utils.DefaultGasPrice,
	&utils.MaxGasPrice,
	&utils.GasPriceFactor,
	&utils.DataStreamHost,
	&utils.DataStreamPort,
	&utils.DataStreamWriteTimeout,
	&utils.DataStreamInactivityTimeout,
	&utils.DataStreamInactivityCheckInterval,
	&utils.WitnessFullFlag,
	&utils.SyncLimit,
	&utils.SupportGasless,
	&utils.ExecutorPayloadOutput,
	&utils.DebugTimers,
	&utils.DebugNoSync,
	&utils.DebugLimit,
	&utils.DebugStep,
	&utils.DebugStepAfter,
	&utils.PoolManagerUrl,
	&utils.TxPoolRejectSmartContractDeployments,
	&utils.DisableVirtualCounters,
	&utils.DAUrl,
	&utils.VirtualCountersSmtReduction,
}
