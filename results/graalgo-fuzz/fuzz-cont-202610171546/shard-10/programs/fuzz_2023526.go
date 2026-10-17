package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int64 = -954
    var v2 int = 0
    var v3 int16 = -1
    if (!false) {
        v2 += v2
        v3 = 32767
    }
    v4 := []int{v2, (v2 ^ v2), 32768, -289}
    v5 := GenS1{F0: false, F1: 129, F2: (v2 | v2)}
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (0 < 0) {
            switch "\n\t" {
            case "a":
                switch (v2 << ((18446744073709551615) & 63)) {
                case 9223372036854775807:
                    var v6 int32 = 1
                    _ = v6
                }
            case "":
                v1 += (v1 & -9223372036854775808)
            case "\n\t":
                var v7 uint16 = 129
                _ = v7
            default:
                v2 = -400
            }
            var v8 float64 = 0.1
            var v9 uint = 1
            _ = v8
            _ = v9
        }
    } else {
        var v10 uint8 = 0
        if false {
            var v11 uint64 = 811
            _ = v11
        } else {
            var v12 uint = 987
            _ = v12
        }
        fmt.Println(v2, v1)
        _ = v10
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint64 = 216
        v15 := GenS1{F0: (1 <= (v14 * (^(v14)))), F1: 1, F2: v2}
        if (((v1 - -374) <= (^(v1))) && (56 > 1)) {
            for fi16 := 0; fi16 < 12; fi16++ {
                v3 += (v3 >> ((739) & 63))
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v17 := GenS2{F0: 461}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 float32 = -2.25
                var v20 int64 = v1
                var v21 int = (fi13 - fi18)
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v22 int32 = 1
                    var v23 int16 = v3
                    v24 := GenS1{F0: false, F1: 128, F2: 42}
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            v14 = (^(v14))
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v25 uint32 = 1
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 string = "hello"
        v2 = (fi26 ^ v2)
        fmt.Println(v1)
        _ = fi26
        _ = v27
    }
    v1 -= v1
    v3 = v3
    v3 = (v3 + -32768)
    switch v2 {
    case 4294967295:
        var v28 bool = true
        _ = v28
    case 127:
        switch (-(v2)) {
        case -2147483648:
            v29 := GenS1{F0: (v2 == (v2 % (((v2 / ((-9223372036854775808) | 1))) | 1))), F1: 77, F2: 2147483647}
            _ = v29
        case 745:
            for fi30 := 0; fi30 < 4; fi30++ {
                switch "hello" {
                case "\n\t":
                    if (100 == 100) {
                        fmt.Println(v1, v2)
                    }
                case "":
                    v31 := GenS1{F0: (1 >= v1), F1: 1, F2: (fi30 - 32768)}
                    _ = v31
                default:
                    v4 = append(v4, (v2 * v2))
                }
                v2 += v2
                _ = fi30
            }
        case 127:
            v25 -= 0
        }
    case -2147483648:
        for fi32 := 0; fi32 < 6; fi32++ {
            v33 := GenS1{F0: (42 <= 235), F1: 128, F2: (v2 - (^(fi32)))}
            var v34 int16 = v3
            fmt.Println(v25, v25, v2)
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
