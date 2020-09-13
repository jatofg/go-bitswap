package decision

import intdec "github.com/ipfs/go-bitswap/internal/decision"

// Expose Receipt externally
type Receipt = intdec.Receipt

// Expose ScoreLedger externally
type ScoreLedger = intdec.ScoreLedger

// Expose ScorePeerFunc externally
type ScorePeerFunc = intdec.ScorePeerFunc

// Expose want-list cache externally
var EnableWantlistCaching = intdec.EnableWantlistCaching
var GetWantlistCache = intdec.GetWantlistCache
var GetAndResetWantlistCache = intdec.GetAndResetWantlistCache
var ResetWantlistCache = intdec.ResetWantlistCache