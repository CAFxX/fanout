package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 string; F1 int32; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := []uint32{2147483648, 1, 2147483648}
    var v2 uint64 = 1
    v3 := make(map[uint32]uint32)
    switch ("hello" + ("\n\t" + "hello")) {
    case "\n\t":
        var v4 int = 32768
        _ = v4
    case "foo bar":
        v2 = (v2 >> ((2) & 63))
    case "a":
        _ = len(v3)
    default:
        for fi5 := 0; fi5 < 5; fi5++ {
            v2 -= v2
            _ = fi5
        }
    }
    var v6 uint16 = 42
    v7 := GenS1{F0: -9223372036854775808, F1: -128, F2: -1}
    if false {
        v8 := GenS2{F0: 1, F1: 73}
        var v9 int64 = -457
        delete(v3, 1)
        _ = v8
        _ = v9
    }
    v10 := GenS3{F0: ("foo bar" + ("a" + "a")), F1: -2147483648, F2: 335}
    if false {
        if true {
            var v11 int32 = 42
            fmt.Println(v6)
            var v12 bool = ((!false) && (32767 == 1))
            _ = v11
            _ = v12
        }
        v13 := GenS1{F0: 344, F1: 86, F2: -96}
        _ = v13
    } else {
        var v14 int16 = 1
        v15 := GenS4{F0: (((v14 << ((18446744073709551615) & 63)) < v14) || (true || true))}
        _ = v14
        _ = v15
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    var v16 uint32 = 65535
    var v17 int16 = 1000
    fmt.Println(v16)
    v3[(v16 / ((v16) | 1))] = (v16 << ((18446744073709551615) & 63))
    var v18 uint = 148
    var v19 uint8 = 2
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS2{F0: uint32(v18), F1: (v18 * (^(v18)))}
        v1 = append(v1, v16)
        v22 := GenS5{F0: (v6 ^ uint16(v2))}
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        if (100 == 127) {
            switch "\n\t" {
            case "\n\t":
                var v24 int16 = -1
                _ = v24
            case "":
                if (!true) {
                    v25 := GenS2{F0: (v16 << ((v18) & 63)), F1: (v18 << ((v18) & 63))}
                    v26 := GenS1{F0: 0, F1: 0, F2: -114}
                    _ = v25
                    _ = v26
                }
            case "foo bar":
                v19 = (^(v19))
            default:
                v6 -= v6
            }
            var v27 int8 = -113
            _ = v27
        } else {
            v28 := GenS5{F0: (v6 ^ 825)}
            _ = v28
        }
        switch ("a" + "\n\t") {
        case "foo bar":
            var v29 int16 = (v17 - 490)
            _ = v29
        case "\n\t":
            fmt.Println(v17, v16, fi23)
        case "":
            _ = len(v1)
        }
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
