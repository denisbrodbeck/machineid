// +build zos

package machineid

import (
        "runtime"
        "unsafe"
)

func machineID() (string, error) {
        type sliceHeader struct {
                addr unsafe.Pointer
                len  int
                cap  int
        }

        cvt := uintptr(*(*int32)(unsafe.Pointer(uintptr(16))))
        pccaavt := uintptr(*(*int32)(unsafe.Pointer(uintptr(cvt + 764))))
        pcca := uintptr(*(*int32)(unsafe.Pointer(uintptr(pccaavt))) + 4)

        var b []byte
        hdr := (*sliceHeader)(unsafe.Pointer(&b))
        hdr.addr = unsafe.Pointer(pcca)
        hdr.cap = 12
        hdr.len = 12

        var res [12]byte
        copy(res[0:12], b)
        runtime.CallLeFuncByPtr(runtime.XplinkLibvec+0x6e3<<4,
                []uintptr{uintptr(unsafe.Pointer(&res[0])), uintptr(12)})

        return string(res[:]), nil
}
