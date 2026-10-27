package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 uint16 }
func genFn0(p0 int64) int32 {
    if true {
        p0 = p0
        switch "\n\t" {
        case "a":
            v4 := []int{}
            _ = v4
        case "foo bar":
            p0 -= -9223372036854775808
        default:
            p0 = -1
        }
    }
    var v5 bool = true
    var v6 int8 = 22
    _ = p0
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{2147483648, 409, 129, 364}
    v2 := make(map[uint32]int)
    var v3 int16 = 0
    if ((!true) || true) {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 float32 = ((-2.25 + -2.25) - -2.25)
            var v9 uint8 = 128
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 uint8 = v9
                var v12 int64 = -1
                v3 += (v3 << ((255) & 63))
                var v13 int8 = -1
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    if false {
        for fi14 := 0; fi14 < 18; fi14++ {
            for fi15 := 0; fi15 < 16; fi15++ {
                v3 = v3
                v3 += 1
                _ = fi15
            }
            _ = fi14
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            v3 = -32768
            var v17 uint64 = 4294967295
            _ = fi16
            _ = v17
        }
        var v18 uint8 = 210
        _ = v18
    } else {
        v19 := GenS1{F0: 1, F1: (true && true), F2: 1}
        _ = v19
    }
    if (1.5 > -2.25) {
        fmt.Println(v3)
    }
    v20 := GenS1{F0: -1, F1: true, F2: 2147483648}
    v21 := GenS2{F0: 763, F1: 588, F2: 0}
    for fi22 := 0; fi22 < 3; fi22++ {
        fmt.Println(v3, v3)
        v3 = v3
        fmt.Println(fi22, v3)
        _ = fi22
    }
    v3 += (v3 + (v3 << ((404) & 63)))
    v23 := GenS3{F0: 1, F1: 414, F2: 646}
    if len(v1) > 0 {
        v1[0] = 2
    }
    var v24 uint64 = 535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
}
