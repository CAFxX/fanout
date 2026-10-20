package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint; F1 string; F2 int16 }
func genFn0(p0 uint32, p1 string) int64 {
    v6 := GenS1{F0: -1, F1: -1}
    var v7 float32 = -2.25
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 1
    var v4 uint32 = 112
    v5 := make(map[string]bool)
    var v8 uint32 = (v4 ^ v4)
    fmt.Println(v4, v1, v8)
    for fi9 := 0; fi9 < 6; fi9++ {
        if ("" <= "foo bar") {
            var v10 uint = uint(v3)
            fmt.Println(fi9, v4)
            _ = v10
        } else {
            v11 := GenS2{F0: 915, F1: "a", F2: 32767}
            v12 := GenS1{F0: -101, F1: -42}
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    v1 = (v1 + (v1 & (-(v1))))
    v3 = (v3 + 65)
    for fi13 := 0; fi13 < 18; fi13++ {
        v3 += v3
        if true {
            _ = len(v2)
            v4 -= (v8 / ((uint32(v3)) | 1))
        } else {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 bool = false
                v4 -= uint32(v3)
                _ = fi14
                _ = v15
            }
            var v16 uint64 = uint64(v4)
            _ = v16
        }
        v1 -= (v1 / ((-9223372036854775808) | 1))
        _ = fi13
    }
    var v17 int32 = int32(v3)
    fmt.Println(v1, v3, v17)
    for fi18 := 0; fi18 < 19; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            v17 = (v17 & (v17 & -1))
            _ = fi19
        }
        delete(v2, 65535)
        var v20 uint = 2
        _ = fi18
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
