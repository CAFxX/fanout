package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []uint32{2147483647, 907, 659, 4294967295}
    v2 := make(map[int]uint32)
    var v3 int8 = -128
    v4 := make(map[int]int)
    v5 := make(map[uint32]bool)
    if false {
        var v6 int8 = v3
        v3 += v6
        _ = v6
    }
    v3 = 42
    for fi7 := 0; fi7 < 1; fi7++ {
        if len(v1) > 0 {
            v1[0] = 100
        }
        _ = fi7
    }
    if false {
        switch "hello" {
        case "foo bar":
            if false {
                v3 -= v3
                if (!false) {
                    for fi8 := 0; fi8 < 5; fi8++ {
                        var v9 int64 = 0
                        _ = fi8
                        _ = v9
                    }
                } else {
                    v1 = append(v1, 875)
                }
                v10 := GenS1{F0: 0}
                _ = v10
            }
        default:
            v3 = v3
        }
        var v11 int32 = -8
        _ = len(v1)
        _ = v11
    }
    v3 += (v3 & (v3 - v3))
    var v12 int32 = -1
    v13 := GenS1{F0: 1000}
    v14 := v4[1]
    _ = len(v1)
    v12 += (v12 >> ((1) & 63))
    v15 := GenS1{F0: -104}
    for fi16 := 0; fi16 < 4; fi16++ {
        v14 = v14
        if ((v12 << ((2) & 63)) <= -921) {
            var v17 uint8 = 0
            fmt.Println(v17, v3, v14)
            _ = v17
        } else {
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int = 4294967295
                v19 -= fi16
                var v20 uint16 = 32767
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        _ = fi16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
