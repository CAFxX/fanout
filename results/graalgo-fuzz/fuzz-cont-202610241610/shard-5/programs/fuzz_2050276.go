package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 string; F2 uint8 }

func main() {
    var v1 uint32 = 65535
    var v2 int64 = -9223372036854775808
    var v3 float32 = 1e10
    var v4 int32 = 42
    v5 := []int{-2147483648, 9223372036854775807, 860}
    if (2 < 820) {
        v6 := GenS1{F0: (v1 % (((v1 * (^(v1)))) | 1)), F1: 2, F2: 255}
        fmt.Println(v2)
        if (!(-1 == 0)) {
            if false {
                var v7 bool = false
                _ = v7
            } else {
                switch 0 {
                case -311:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                switch "" {
                case "\n\t":
                    v8 := GenS2{F0: 42, F1: -734, F2: -1}
                    _ = v8
                case "":
                    v5 = append(v5, 127)
                }
                v3 = 0.0
            }
            var v9 uint8 = 108
            var v10 int32 = v4
            _ = v10
            _ = v9
        } else {
            switch 1 {
            case 127:
                var v11 int16 = 0
                _ = v11
            default:
                fmt.Println(v1, v3, v3)
            }
        }
        _ = v6
    } else {
        var v12 uint16 = 255
        if false {
            var v13 uint16 = v12
            if false {
                v13 += v12
            }
            _ = v13
        }
        _ = v12
    }
    fmt.Println(v3, v2, v2)
    _ = len(v5)
    switch ("" + (("\n\t" + ("foo bar" + "a")) + "\n\t")) {
    case "a":
        switch ("hello" + ("" + "hello")) {
        case "foo bar":
            v14 := GenS2{F0: -647, F1: 1, F2: 97}
            _ = v14
        }
    }
    if true {
        for fi15 := 0; fi15 < 13; fi15++ {
            v2 += (^(v2))
            fmt.Println(v3, v3)
            _ = fi15
        }
        var v16 float64 = -2.25
        var v17 int = 0
        _ = v16
        _ = v17
    }
    if len(v5) > 0 {
        v5[0] = 32768
    }
    switch ("\n\t" + "hello") {
    case "a":
        v4 = (v4 << ((909) & 63))
    default:
        switch "" {
        case "foo bar":
            v18 := GenS2{F0: 1000, F1: -143, F2: 100}
            _ = v18
        case "a":
            var v19 float32 = 3.14159
            _ = v19
        default:
            v20 := GenS3{F0: v2}
            _ = v20
        }
    }
    var v21 uint8 = 42
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 = (v2 % ((360) | 1))
    var v22 uint32 = v1
    if (!true) {
        for fi23 := 0; fi23 < 13; fi23++ {
            switch fi23 {
            case -2147483648:
                var v24 uint16 = uint16(v22)
                _ = v24
            case -754:
                v5 = append(v5, (fi23 + (fi23 | -94)))
            case 282:
                v25 := GenS4{F0: 18446744073709551615, F1: "\n\t", F2: (v21 & (v21 * 0))}
                _ = v25
            }
            _ = fi23
        }
        var v26 int32 = v4
        v3 = (v3 / (float32(1e10) / v3))
        _ = v26
    } else {
        _ = len(v5)
        var v27 uint64 = 42
        _ = v27
    }
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 int64 = -1
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
