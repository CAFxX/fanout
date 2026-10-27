package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint16 }
func genFn0() string {
    var v4 uint16 = 32767
    var v5 uint16 = v4
    v6 := []int{4294967295, -170}
    if (((-32768 > -663) && ((!true) && (!false))) && true) {
        if (-1 >= 1) {
            v6 = append(v6, -2147483648)
            var v7 bool = (!false)
            var v8 float32 = (-2.25 / ((((((((1e10 + 0.0) * 1e10) * 3.14159) + ((1.5 + -2.25) + float32(0.1))) / 1e10) - -2.25) + 1e10) / 3.14159))
            _ = v7
            _ = v8
        } else {
            v5 = (v5 & v5)
            for fi9 := 0; fi9 < 16; fi9++ {
                v4 += v5
                v10 := GenS1{F0: 710, F1: -32768}
                _ = fi9
                _ = v10
            }
            v11 := GenS2{F0: (^(v5))}
            _ = v11
        }
        if true {
            switch 32768 {
            case 127:
                var v12 int16 = -1
                _ = v12
            case 2147483647:
                if (32768 > 0) {
                    v4 -= 0
                    v5 = (v4 & 0)
                }
            default:
                for fi13 := 0; fi13 < 8; fi13++ {
                    v14 := GenS1{F0: fi13, F1: 992}
                    _ = fi13
                    _ = v14
                }
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            var v15 uint16 = v5
            if true {
                v4 += 32767
                for fi16 := 0; fi16 < 6; fi16++ {
                    if false {
                        fmt.Println(v15, fi16)
                        var v17 float64 = 0.0
                        _ = v17
                    }
                    var v18 uint64 = 182
                    var v19 uint8 = 203
                    _ = fi16
                    _ = v18
                    _ = v19
                }
            } else {
                var v20 int64 = -671
                _ = v20
            }
            _ = v15
        }
        switch 0 {
        case 128:
            var v21 int32 = -1
            _ = v21
        case 4294967295:
            fmt.Println(v5, v5)
        case 419:
            var v22 uint = 956
            _ = v22
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return ""
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    _ = len(v3)
    v1 = v1
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 int32 = -1
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        var v26 float32 = 1.5
        v1 = v1
        var v27 int64 = 518
        _ = fi25
        _ = v26
        _ = v27
    }
    v1 = v1
    fmt.Println(v1)
    var v28 int = 128
    var v29 float32 = (((0.1 / 1.0) + 0.1) + float32(float64((3.14159 + 0.1))))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
