package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint8; F2 int32 }

func main() {
    var v1 int8 = -108
    var v2 int64 = -1
    var v3 int8 = (v1 >> ((1) & 63))
    var v4 int64 = (v2 / ((42) | 1))
    v5 := []int{32767, 4294967295}
    switch 524 {
    case 128:
        if true {
            v5 = append(v5, 1)
        }
    case -2147483648:
        if len(v5) > 0 {
            v5[0] = 1
        }
    default:
        if (!false) {
            if ((39 > 0) || (!(519 >= 42))) {
                v2 += v2
            }
        } else {
            for fi6 := 0; fi6 < 5; fi6++ {
                if ((v4 & (v4 | 1)) >= v4) {
                    v7 := GenS1{F0: 1, F1: 65535}
                    fmt.Println(v4, v2)
                    _ = v7
                }
                for fi8 := 0; fi8 < 17; fi8++ {
                    v9 := GenS2{F0: 4294967295, F1: 821}
                    v2 = v4
                    _ = fi8
                    _ = v9
                }
                _ = fi6
            }
        }
    }
    v3 = (v1 + (v1 + 26))
    switch int(v3) {
    case 1:
        switch "foo bar" {
        case "hello":
            var v10 uint16 = 42
            _ = v10
        }
    case 127:
        fmt.Println(v2)
    case -9223372036854775808:
        v11 := GenS2{F0: 18446744073709551615, F1: v2}
        _ = v11
    default:
        v3 = 0
    }
    v1 -= (v3 >> ((0) & 63))
    v12 := GenS2{F0: 808, F1: -9223372036854775808}
    if (1 == 42) {
        var v13 float64 = 0.0
        v13 -= v13
        v3 = v1
        _ = v13
    }
    _ = len(v5)
    var v14 int16 = -758
    v15 := GenS2{F0: 42, F1: v4}
    v16 := GenS2{F0: 2, F1: 0}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v17 := GenS3{F0: 799, F1: 977}
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 int32 = -1
        _ = fi18
        _ = v19
    }
    v20 := GenS3{F0: 2147483647, F1: 65535}
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(v1)
        v22 := GenS4{F0: 386, F1: 2, F2: 1}
        _ = fi21
        _ = v22
    }
    fmt.Println(v1, v14)
    if false {
        for fi23 := 0; fi23 < 1; fi23++ {
            v14 = (^(v14))
            if (false || true) {
                for fi24 := 0; fi24 < 14; fi24++ {
                    fmt.Println(v2)
                    var v25 int = fi24
                    var v26 uint16 = 0
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            _ = fi23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
