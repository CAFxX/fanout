package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0(p0 int16) string {
    var v4 int16 = 1000
    fmt.Println(v4, v4)
    fmt.Println(v4)
    fmt.Println(p0)
    _ = p0
    _ = v4
    return "\n\t"
}
func genFn1(p0 string) int16 {
    p0 = ("hello" + "")
    v5 := []int{42, 883, 42, -680}
    if false {
        p0 = ((p0 + p0) + p0)
        v5 = append(v5, 128)
    } else {
        if (true || false) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            p0 = ""
            for fi6 := 0; fi6 < 8; fi6++ {
                var v7 uint8 = 1
                for fi8 := 0; fi8 < 12; fi8++ {
                    var v9 uint8 = v7
                    if true {
                        var v10 uint = 65535
                        _ = v10
                    } else {
                        fmt.Println(p0, p0, v7)
                        var v11 uint8 = 163
                        _ = v11
                    }
                    v12 := GenS1{F0: uint64(v9), F1: -480, F2: 0}
                    _ = fi8
                    _ = v12
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        } else {
            var v13 uint16 = 0
            _ = v13
        }
        var v14 string = p0
        fmt.Println(v14, v14)
        _ = v14
    }
    var v15 float32 = ((1.5 + 0.1) * float32(1.5))
    _ = p0
    _ = v15
    _ = v5
    return -1
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v16 bool = (-114 != 100)
    fmt.Println(v1)
    v16 = (!false)
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS2{F0: (v1 % ((v1) | 1)), F1: 267}
        v16 = (127 > 127)
        v1 = v1
        _ = fi17
        _ = v18
    }
    fmt.Println(v16, v16)
    fmt.Println(v16, v16)
    v1 = (v1 % ((v1) | 1))
    fmt.Println(v1, v1)
    var v19 uint16 = 65535
    fmt.Println(v19, v19, v1)
    if (-32768 < -32768) {
        var v20 int16 = -32768
        v1 += v1
        var v21 uint = 1
        _ = v20
        _ = v21
    }
    v1 = v1
    var v22 uint32 = 861
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := GenS1{F0: uint64(v19), F1: v1, F2: (v19 - 0)}
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: v1, F1: v22}
    v3["a"] = 127
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v25
    _ = v3
}
