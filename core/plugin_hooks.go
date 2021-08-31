package core

import (
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/plugins"
  "github.com/ethereum/go-ethereum/log"
  "github.com/ethereum/go-ethereum/rlp"
)

func PluginPreProcessBlock(pl *plugins.PluginLoader, block *types.Block) {
  fnList := pl.Lookup("PreProcessBlock", func(item interface{}) bool {
    _, ok := item.(func([]byte))
    return ok
  })
  encoded, _ = rlp.EncodeToBytes(block)
  for _, fni := range fnList {
    if fn, ok := fni.(func([]byte)); ok {
      fn(encoded)
    }
  }
}
func pluginPreProcessBlock(block *types.Block) {
  if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting PreProcessBlock, but default PluginLoader has not been initialized")
    return
  }
  PluginPreProcessBlock(plugins.DefaultPluginLoader, block) // TODO
}
func PluginPreProcessTransaction(pl *plugins.PluginLoader, tx *types.Transaction, block *types.Block, i int) {
  fnList := pl.Lookup("PreProcessTransaction", func(item interface{}) bool {
    _, ok := item.(func([]byte, []byte, int))
    return ok
  })
  txBytes, _ := tx.MarshalBinary()
  blockBytes, _ := rlp.EncodeToBytes(block)
  for _, fni := range fnList {
    if fn, ok := fni.(func([]byte, []byte, int)); ok {
      fn(txBytes, blockBytes, i)
    }
  }
}
func pluginPreProcessTransaction(tx *types.Transaction, block *types.Block, i int) {
  if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting PreProcessTransaction, but default PluginLoader has not been initialized")
    return
  }
  PluginPreProcessTransaction(plugins.DefaultPluginLoader, tx, block, i)
}
func PluginBlockProcessingError(pl *plugins.PluginLoader, tx *types.Transaction, block *types.Block, err error) {
  fnList := pl.Lookup("BlockProcessingError", func(item interface{}) bool {
    _, ok := item.(func(*types.Transaction, *types.Block, error))
    return ok
  })
  for _, fni := range fnList {
    if fn, ok := fni.(func(*types.Transaction, *types.Block, error)); ok {
      fn(tx, block, err)
    }
  }
}
func pluginBlockProcessingError(tx *types.Transaction, block *types.Block, err error) {
  if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting BlockProcessingError, but default PluginLoader has not been initialized")
    return
  }
  PluginBlockProcessingError(plugins.DefaultPluginLoader, tx, block, err)
}
func PluginPostProcessTransaction(pl *plugins.PluginLoader, tx *types.Transaction, block *types.Block, i int, receipt *types.Receipt) {
  fnList := pl.Lookup("PostProcessTransaction", func(item interface{}) bool {
    _, ok := item.(func(*types.Transaction, *types.Block, int, *types.Receipt))
    return ok
  })
  for _, fni := range fnList {
    if fn, ok := fni.(func(*types.Transaction, *types.Block, int, *types.Receipt)); ok {
      fn(tx, block, i, receipt)
    }
  }
}
func pluginPostProcessTransaction(tx *types.Transaction, block *types.Block, i int, receipt *types.Receipt) {
  if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting PostProcessTransaction, but default PluginLoader has not been initialized")
    return
  }
  PluginPostProcessTransaction(plugins.DefaultPluginLoader, tx, block, i, receipt)
}
func PluginPostProcessBlock(pl *plugins.PluginLoader, block *types.Block) {
  fnList := pl.Lookup("PostProcessBlock", func(item interface{}) bool {
    _, ok := item.(func(*types.Block))
    return ok
  })
  for _, fni := range fnList {
    if fn, ok := fni.(func(*types.Block)); ok {
      fn(block)
    }
  }
}
func pluginPostProcessBlock(block *types.Block) {
  if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting PostProcessBlock, but default PluginLoader has not been initialized")
    return
  }
  PluginPostProcessBlock(plugins.DefaultPluginLoader, block)
}


func PluginNewHead(pl *plugins.PluginLoader, block *types.Block, hash common.Hash, logs []*types.Log) {
  fnList := pl.Lookup("NewHead", func(item interface{}) bool {
    _, ok := item.(func(*types.Block, common.Hash, []*types.Log))
    return ok
  })
  for _, fni := range fnList {
    if fn, ok := fni.(func(*types.Block, common.Hash, []*types.Log)); ok {
      fn(block, hash, logs)
    }
  }
}
func pluginNewHead(block *types.Block, hash common.Hash, logs []*types.Log) {
  if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting NewHead, but default PluginLoader has not been initialized")
    return
  }
  PluginNewHead(plugins.DefaultPluginLoader, block, hash, logs)
}

func PluginNewSideBlock(pl *plugins.PluginLoader, block *types.Block, hash common.Hash, logs []*types.Log) {
  fnList := pl.Lookup("NewSideBlock", func(item interface{}) bool {
    _, ok := item.(func(*types.Block, common.Hash, []*types.Log))
    return ok
  })
  for _, fni := range fnList {
    if fn, ok := fni.(func(*types.Block, common.Hash, []*types.Log)); ok {
      fn(block, hash, logs)
    }
  }
}
func pluginNewSideBlock(block *types.Block, hash common.Hash, logs []*types.Log) {
  if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting NewSideBlock, but default PluginLoader has not been initialized")
    return
  }
  PluginNewSideBlock(plugins.DefaultPluginLoader, block, hash, logs)
}

func PluginReorg(pl *plugins.PluginLoader, commonBlock *types.Block, oldChain, newChain types.Blocks) {
  fnList := pl.Lookup("Reorg", func(item interface{}) bool {
    _, ok := item.(func(common *types.Block, oldChain, newChain types.Blocks))
    return ok
  })
  for _, fni := range fnList {
    if fn, ok := fni.(func(common *types.Block, oldChain, newChain types.Blocks)); ok {
      fn(commonBlock, oldChain, newChain)
    }
  }
}
func pluginReorg(commonBlock *types.Block, oldChain, newChain types.Blocks) {
  if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting Reorg, but default PluginLoader has not been initialized")
    return
  }
  PluginReorg(plugins.DefaultPluginLoader, commonBlock, oldChain, newChain)
}
