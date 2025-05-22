package vips

/*
#cgo pkg-config: vips
#include <stdlib.h>
#include <vips/vips.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	"github.com/fyntrix/fyntrix/pkg/logger"
)

var (
	_initOnce sync.Once
	_initErr  error
	_logging  *logger.SubLogger
)

var (
	MajorVersion = int(C.vips_version(0))
	Version      = string(C.GoString(C.vips_version_string()))
)

type Vips struct {
	lk     sync.Mutex
	cfg    *Config
	inited bool
}

func New(cfg *Config) *Vips {
	v := &Vips{
		cfg: cfg,
	}

	_logging = logger.NewSubLogger("_image", v)

	return v
}

func (v *Vips) Init() error {
	v.lk.Lock()
	defer v.lk.Unlock()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if MajorVersion < 8 {
		return errors.New("vips version < 8")
	}

	_initOnce.Do(func() {
		arg := C.CString("vips")
		defer C.free(unsafe.Pointer(arg))

		if C.vips_init(arg) != 0 {
			_initErr = errors.New("vips_init failed")
			return
		}

		if v.cfg.ReportLeaks {
			C.vips_leak_set(toGboolean(v.cfg.ReportLeaks))
			_logging.Debug("enabled report leaks")
		}

		C.vips_concurrency_set(C.int(v.cfg.ConcurrentLevel))
		_logging.Debug("set concurrent level", "level", v.cfg.ConcurrentLevel)
		C.vips_cache_set_max_files(C.int(v.cfg.MaxCacheFiles))
		_logging.Debug("set max cache files", "size", v.cfg.MaxCacheFiles)
		C.vips_cache_set_max_mem(C.size_t(v.cfg.MaxCacheMem))
		_logging.Debug("set cache max memory", "size", v.cfg.MaxCacheMem)
		C.vips_cache_set_max(C.int(v.cfg.MaxCacheSize))
		_logging.Debug("set max cache size", "size", v.cfg.MaxCacheSize)

		if v.cfg.CacheTrace {
			C.vips_cache_set_trace(toGboolean(true))
			_logging.Debug("enabled cache trace")
		}

		_logging.Debug("initialized vips library")
	})

	if _initErr != nil {
		return _initErr
	}

	_logging.Info("initialized vips library")

	v.inited = true
	return nil
}

func (v *Vips) Close() error {
	v.lk.Lock()
	defer v.lk.Unlock()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if v.inited {
		C.vips_shutdown()
		v.inited = false
		_logging.Warn("vips library closed")
	}

	return nil
}

func (v *Vips) Inited() bool {
	v.lk.Lock()
	defer v.lk.Unlock()
	return v.inited
}

func (*Vips) Version() string {
	return Version
}

func (*Vips) String() string {
	return fmt.Sprintf("vips %s", Version)
}

func toGboolean(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func fromGboolean(b C.gboolean) bool {
	return b != 0
}
