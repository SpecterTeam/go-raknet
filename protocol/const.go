/**
 *     SpecterGO  Copyright (C) 2018  SpecterTeam
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package protocol

const (
	IDConnectedPing                                    = 0x00
	IDUnconnectedPing                                  = 0x01
	IDUnconnectedPingOpenConnections                   = 0x02
	IDConnectedPong                                    = 0x03
	IDDetectLostConnections                            = 0x04
	IDOpenConnectionRequest1                           = 0x05
	IDOpenConnectionReply1                             = 0x06
	IDOpenConnectionRequest2                           = 0x07
	IDOpenConnectionReply2                             = 0x08
	IDConnectionRequest                                = 0x09
	IDRemoteSystemRequiresPublicKey                    = 0x0a
	IDOurSystemRequiresSecurity                        = 0x0b
	IDPublicKeyMismatch                                = 0x0c
	IDOutOfBandInternal                                = 0x0d
	IDSndReceiptAcked                                  = 0x0e
	IDSndReceiptLoss                                   = 0x0f
	IDConnectionRequestAccepted                        = 0x10
	IDConnectionAttemptFailed                          = 0x11
	IDAlreadyConnected                                 = 0x12
	IDNewIncomingConnection                            = 0x13
	IDNoFreeIncomingConnections                        = 0x14
	IDDisconnectionNotification                        = 0x15
	IDConnectionLost                                   = 0x16
	IDConnectionBanned                                 = 0x17
	IDInvalIDPassword                                  = 0x18
	IDIncompatibleProtocolVersion                      = 0x19
	IDIpRecentlyConnected                              = 0x1a
	IDTimestamp                                        = 0x1b
	IDUnconnectedPong                                  = 0x1c
	IDAdvertiseSystem                                  = 0x1d
	IDDownloadProgress                                 = 0x1e
	IDRemoteDisconnectionNotification                  = 0x1f
	IDRemoteConnectionLost                             = 0x20
	IDRemoteNewIncomingConnection                      = 0x21
	IDFileListTransferHeader                           = 0x22
	IDFileListTransferFile                             = 0x23
	IDFileListReferencePushAck                         = 0x24
	IDDdtDownloadRequest                               = 0x25
	IDTransportString                                  = 0x26
	IDReplicaManagerConstruction                       = 0x27
	IDReplicaManagerScopeChange                        = 0x28
	IDReplicaManagerSerialize                          = 0x29
	IDReplicaManagerDownloadStarted                    = 0x2a
	IDReplicaManagerDownloadComplete                   = 0x2b
	IDRakvoiceOpenChannelRequest                       = 0x2c
	IDRakvoiceOpenChannelReply                         = 0x2d
	IDRakvoiceCloseChannel                             = 0x2e
	IDRakvoiceData                                     = 0x2f
	IDAutopatcherGetChangelistSinceDate                = 0x30
	IDAutopatcherCreationList                          = 0x31
	IDAutopatcherDeletionList                          = 0x32
	IDAutopatcherGetPatch                              = 0x33
	IDAutopatcherPatchList                             = 0x34
	IDAutopatcherRepositoryFatalError                  = 0x35
	IDAutopatcherCannotDownloadOriginalUnmodifiedFiles = 0x36
	IDAutopatcherFinishedInternal                      = 0x37
	IDAutopatcherFinished                              = 0x38
	IDAutopatcherRestartApplication                    = 0x39
	IDNatPunchthroughRequest                           = 0x3a
	IDNatConnectAtTime                                 = 0x3b
	IDNatGetMostRecentPort                             = 0x3c
	IDNatClientReady                                   = 0x3d
	IDNatTargetNotConnected                            = 0x3e
	IDNatTargetUnresponsive                            = 0x3f
	IDNatConnectionToTargetLost                        = 0x40
	IDNatAlreadyInProgress                             = 0x41
	IDNatPunchthroughFailed                            = 0x42
	IDNatPunchthroughSucceeded                         = 0x43
	IDReadyEventSet                                    = 0x44
	IDReadyEventUnset                                  = 0x45
	IDReadyEventAllSet                                 = 0x46
	IDReadyEventQuery                                  = 0x47
	IDLobbyGeneral                                     = 0x48
	IDRpcRemoteError                                   = 0x49
	IDRpcPlugin                                        = 0x4a
	IDFileListReferencePush                            = 0x4b
	IDReadyEventForceAllSet                            = 0x4c
	IDRoomsExecuteFunc                                 = 0x4d
	IDRoomsLogonStatus                                 = 0x4e
	IDRoomsHandleChange                                = 0x4f
	IDLobby2SendMessage                                = 0x50
	IDLobby2ServerError                                = 0x51
	IDFcm2NewHost                                      = 0x52
	IDFcm2RequestFcmguID                               = 0x53
	IDFcm2RespondConnectionCount                       = 0x54
	IDFcm2InformFcmguID                                = 0x55
	IDFcm2UpdateMinTotalConnectionCount                = 0x56
	IDFcm2VerifiedJoinStart                            = 0x57
	IDFcm2VerifiedJoinCapable                          = 0x58
	IDFcm2VerifiedJoinFailed                           = 0x59
	IDFcm2VerifiedJoinAccepted                         = 0x5a
	IDFcm2VerifiedJoinRejected                         = 0x5b
	IDUdpProxyGeneral                                  = 0x5c
	IDSqlite3Exec                                      = 0x5d
	IDSqlite3UnknownDb                                 = 0x5e
	IDSqlliteLogger                                    = 0x5f
	IDNatTypeDetectionRequest                          = 0x60
	IDNatTypeDetectionResult                           = 0x61
	IDRouter2Internal                                  = 0x62
	IDRouter2ForwardingNoPath                          = 0x63
	IDRouter2ForwardingEstablished                     = 0x64
	IDRouter2Rerouted                                  = 0x65
	IDTeamBalancerInternal                             = 0x66
	IDTeamBalancerRequestedTeamFull                    = 0x67
	IDTeamBalancerRequestedTeamLocked                  = 0x68
	IDTeamBalancerTeamRequestedCancelled               = 0x69
	IDTeamBalancerTeamAssigned                         = 0x6a
	IDLightspeedIntegration                            = 0x6b
	IDXboxLobby                                        = 0x6c
	IDTwoWayAuthenticationIncomingChallengeSuccess     = 0x6d
	IDTwoWayAuthenticationOutgoingChallengeSuccess     = 0x6e
	IDTwoWayAuthenticationIncomingChallengeFailure     = 0x6f
	IDTwoWayAuthenticationOutgoingChallengeFailure     = 0x70
	IDTwoWayAuthenticationOutgoingChallengeTimeout     = 0x71
	IDTwoWayAuthenticationNegotiation                  = 0x72
	IDCloudPostRequest                                 = 0x73
	IDCloudReleaseRequest                              = 0x74
	IDCloudGetRequest                                  = 0x75
	IDCloudGetResponse                                 = 0x76
	IDCloudUnsubscribeRequest                          = 0x77
	IDCloudServerToServerCommand                       = 0x78
	IDCloudSubscriptionNotification                    = 0x79
	IDLibVoice                                         = 0x7a
	IDRelayPlugin                                      = 0x7b
	IDNatRequestBoundAddresses                         = 0x7c
	IDNatRespondBoundAddresses                         = 0x7d
	IDFcm2UpdateUserContext                            = 0x7e
	IDReserved3                                        = 0x7f
	IDReserved4                                        = 0x80
	IDReserved5                                        = 0x81
	IDReserved6                                        = 0x82
	IDReserved7                                        = 0x83
	IDReserved8                                        = 0x84
	IDReserved9                                        = 0x85
	IDUserPacketEnum                                   = 0x86
	IDCustom0                                          = 0x80
	IDCustom1                                          = 0x81
	IDCustom2                                          = 0x82
	IDCustom3                                          = 0x83
	IDCustom4                                          = 0x84
	IDCustom5                                          = 0x85
	IDCustom6                                          = 0x86
	IDCustom7                                          = 0x87
	IDCustom8                                          = 0x88
	IDCustom9                                          = 0x89
	IDCustomA                                          = 0x8a
	IDCustomB                                          = 0x8b
	IDCustomC                                          = 0x8c
	IDCustomD                                          = 0x8d
	IDCustomE                                          = 0x8e
	IDCustomF                                          = 0x8f

	IDACK  = 0xc0
	IDNACK = 0xa0
)
