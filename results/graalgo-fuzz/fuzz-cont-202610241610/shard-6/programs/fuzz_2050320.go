package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint8; F1 int8; F2 int16 }
type GenS3 struct { F0 int16; F1 uint16 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]int)
    var v3 uint8 = 40
    v4 := make(map[string]bool)
    var v5 float32 = float32(3.14159)
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            _ = len(v2)
            v7 := GenS1{F0: 1, F1: 42}
            _ = fi6
            _ = v7
        }
        var v8 int16 = -1
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := []int{fi9}
        _ = fi9
        _ = v10
    }
    if (!(270 <= 65535)) {
        v1 = v1
        v11 := GenS1{F0: 4294967295, F1: 42}
        _ = v11
    } else {
        v1 = 471
    }
    var v12 int8 = 100
    v1 += (^(v1))
    switch 1 {
    case 2147483647:
        var v13 uint64 = (v1 >> ((18446744073709551615) & 63))
        _ = v13
    case 906:
        v3 = 127
    case 1:
        v14 := v2[""]
        _ = v14
    }
    var v15 int32 = -1
    var v16 int = 2147483647
    if (v5 < v5) {
        for fi17 := 0; fi17 < 13; fi17++ {
            fmt.Println(v12)
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 int64 = -9223372036854775808
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
        var v20 int64 = 9223372036854775807
        v16 = (v16 / ((v16) | 1))
        _ = v20
    } else {
        fmt.Println(v16, v1)
        fmt.Println(v3, v5)
    }
    var v21 int64 = 0
    var v22 int8 = 0
    var v23 int64 = (v21 << ((42) & 63))
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := GenS2{F0: (v3 << ((306) & 63)), F1: -31, F2: 0}
        if (87 <= -327) {
            v15 += (v15 ^ (v15 & (v15 - (v15 << ((494) & 63)))))
            for fi26 := 0; fi26 < 4; fi26++ {
                fmt.Println(v1, fi24)
                switch ("" + "foo bar") {
                case "foo bar":
                    var v27 uint = 594
                    _ = v27
                default:
                    var v28 int8 = 0
                    _ = v28
                }
                _ = fi26
            }
        } else {
            if ((1000 < -353) || (!(((v1 << ((uint(v1)) & 63)) != (v1 + 2)) || false))) {
                var v29 int = (v16 & v16)
                _ = v29
            }
            v3 = v3
            v16 += 0
        }
        _ = fi24
        _ = v25
    }
    switch v16 {
    case 4294967295:
        var v30 int = 902
        _ = v30
    case 32767:
        v22 = 127
    default:
        v31 := GenS3{F0: 640, F1: 305}
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
