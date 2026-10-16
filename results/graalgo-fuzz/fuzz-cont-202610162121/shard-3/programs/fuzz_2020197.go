package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int }
type GenS2 struct { F0 int32; F1 uint64; F2 int32 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0() uint64 {
    v7 := []int{0, -1}
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: (fi8 / (((-(fi8))) | 1)), F1: -598, F2: (fi8 << ((4294967295) & 63))}
        var v10 int16 = 0
        var v11 int = fi8
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := GenS2{F0: 681, F1: 4294967295, F2: 787}
    _ = v12
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint8 = 82
    v2 := make(map[string]uint32)
    var v3 uint32 = 42
    v4 := []uint32{2147483648, (v3 % ((2147483648) | 1)), (^(v3))}
    v5 := []uint32{v3}
    var v6 uint64 = 2
    for fi13 := 0; fi13 < 2; fi13++ {
        switch 376 {
        case 2147483647:
            v3 = (v3 - uint32(v1))
        case 32768:
            v3 -= (v3 >> ((0) & 63))
        default:
            if true {
                switch fi13 {
                case 2147483647:
                    v2[("a" + "foo bar")] = v3
                default:
                    var v14 float32 = (0.0 / 1.0)
                    _ = v14
                }
                v15 := GenS3{F0: 28, F1: 4294967295}
                _ = v15
            } else {
                v16 := GenS3{F0: 22, F1: (^(fi13))}
                if (42 > 42) {
                    v6 = 1
                }
                _ = len(v4)
                _ = v16
            }
        }
        _ = fi13
    }
    v17 := GenS4{F0: 1000, F1: (v6 >> ((2) & 63))}
    v6 -= (v6 >> ((255) & 63))
    if len(v5) > 0 {
        v5[0] = uint32(v1)
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := v2["\n\t"]
        if len(v4) > 0 {
            v4[0] = 0
        }
        _ = len(v4)
        _ = fi18
        _ = v19
    }
    _ = len(v4)
    v20 := GenS4{F0: -1, F1: 0}
    v1 += v1
    v3 += v3
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 uint8 = 42
        if (!(("\n\t" + "") != "\n\t")) {
            if ((fi21 % (((fi21 / (((fi21 + fi21)) | 1))) | 1)) >= (fi21 >> ((0) & 63))) {
                var v23 uint16 = 119
                var v24 int64 = 552
                var v25 uint64 = 42
                _ = v23
                _ = v24
                _ = v25
            }
            _ = len(v2)
        }
        _ = fi21
        _ = v22
    }
    v26 := GenS4{F0: -349, F1: (v6 << ((547) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
