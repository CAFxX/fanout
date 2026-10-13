package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int64 }
type GenS2 struct { F0 int8; F1 string; F2 bool }
func genFn0(p0 int16) uint32 {
    var v7 uint = 382
    fmt.Println(v7, p0)
    var v8 int = 640
    var v9 int8 = -1
    v10 := GenS1{F0: true, F1: false, F2: -9223372036854775808}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 340
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{}
    var v3 int16 = 882
    var v4 int16 = v3
    var v5 int8 = 42
    var v6 uint8 = 2
    var v11 uint16 = 65535
    var v12 uint8 = uint8(v11)
    v13 := GenS1{F0: false, F1: true, F2: int64(v5)}
    var v14 string = "foo bar"
    fmt.Println(v6)
    switch v14 {
    case "hello":
        if false {
            switch ("a" + (v14 + v14)) {
            case "a":
                _ = len(v1)
            case "hello":
                var v15 int = 32768
                _ = v15
            }
        }
    case "a":
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int16 = (-(v4))
            var v18 uint = 949
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    if (((((v14 + ("" + ((v14 + v14) + (v14 + "a")))) + "hello") + v14) < "hello") && false) {
        v11 += v11
        switch v14 {
        case "foo bar":
            fmt.Println(v12, v4, v4)
        case "\n\t":
            fmt.Println(v11, v5, v14)
        default:
            var v19 int16 = (v4 / (((v3 + -271)) | 1))
            _ = v19
        }
        v20 := v1[42]
        _ = v20
    } else {
        if true {
            switch 0 {
            case -9223372036854775808:
                v5 = -81
            case 128:
                v2 = append(v2, "hello")
            default:
                v1[-387] = genFn0(int16(v11))
            }
            var v21 bool = (0.0 != 1.5)
            _ = v21
        }
        v22 := GenS2{F0: v5, F1: v14, F2: (!(-922 < 0))}
        v12 -= (v12 ^ 2)
        _ = v22
    }
    v1[9223372036854775807] = 65535
    v4 -= (v3 << ((65535) & 63))
    var v23 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
