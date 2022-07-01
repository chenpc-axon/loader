package loader

import dg "chenpc.com/axon/apis/datagram"

type Loader interface {
	// Key 唯一标识
	Key() string
	// Name 加载器的名称
	Name() string
	// Init 初始化加载器
	Init() error
	// Destroy 销毁加载器
	Destroy()
	// Load 写入 Datagram
	Load(datagrams []dg.Datagram) error
}
