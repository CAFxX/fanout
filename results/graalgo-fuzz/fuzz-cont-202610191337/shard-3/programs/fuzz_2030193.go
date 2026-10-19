package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int }
type GenS2 struct { F0 uint64; F1 uint32; F2 int32 }

func main() {
    var v1 uint16 = 2
    var v2 int64 = 478
    var v3 float32 = 0.1
    var v4 int32 = 0
    v5 := []string{"foo bar"}
    v6 := make(map[uint32]bool)
    _ = len(v6)
    for fi7 := 0; fi7 < 11; fi7++ {
        _ = len(v6)
        _ = fi7
    }
    v3 += (-2.25 * (v3 * 1.5))
    v3 = v3
    var v8 bool = false
    v4 -= (v4 >> ((1) & 63))
    v9 := GenS1{F0: 501, F1: v4, F2: 4294967295}
    fmt.Println(v4, v2, v2)
    fmt.Println(v1)
    switch 9223372036854775807 {
    case -562:
        v2 = v2
    case 42:
        v2 = v2
    case 322:
        if false {
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 int = 32768
                _ = fi10
                _ = v11
            }
        }
    }
    delete(v6, 788)
    v8 = v8
    for fi12 := 0; fi12 < 16; fi12++ {
        v2 = (v2 / (((v2 & 318)) | 1))
        _ = fi12
    }
    v13 := GenS2{F0: 1, F1: 0, F2: int32(v1)}
    v2 = (v2 * (^(v2)))
    for fi14 := 0; fi14 < 12; fi14++ {
        switch 9223372036854775807 {
        case 42:
            switch (fi14 - 2147483647) {
            case -556:
                v6[2] = (fi14 > fi14)
            case -9223372036854775808:
                fmt.Println(fi14)
            case 32768:
                for fi15 := 0; fi15 < 8; fi15++ {
                    v16 := v6[393]
                    var v17 uint64 = 2
                    v4 += 0
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            default:
                var v18 int8 = 100
                _ = v18
            }
        case -624:
            v3 = float32((0.1 - -2.25))
        default:
            switch (-(fi14)) {
            case 128:
                fmt.Println(v8, v2)
            case 820:
                v3 = (float32(0.1) / float32(0.1))
            }
        }
        v4 = (-(v4))
        if (2 <= 783) {
            v6[2147483648] = (0 != 4294967295)
            v19 := v6[555]
            fmt.Println(v19)
            _ = v19
        } else {
            var v20 uint32 = 65535
            v20 = v20
            _ = v20
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
