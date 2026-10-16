package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 uint16; F1 int32 }

func main() {
    var v1 float32 = 1e10
    var v2 uint16 = 765
    var v3 uint32 = 2147483648
    var v4 int64 = 1
    v5 := []uint32{}
    v6 := make(map[string]bool)
    if true {
        var v7 int = 9223372036854775807
        _ = v7
    }
    if (1 < 2) {
        if true {
            if ((v3 & 42) == (v3 & v3)) {
                switch "foo bar" {
                case "\n\t":
                    v8 := GenS1{F0: (v2 | v2), F1: (v2 / ((255) | 1)), F2: -417}
                    _ = v8
                case "foo bar":
                    fmt.Println(v1, v4)
                }
                _ = len(v6)
                fmt.Println(v4, v2)
            }
            var v9 int16 = -32768
            for fi10 := 0; fi10 < 17; fi10++ {
                switch "hello" {
                case "":
                    var v11 string = "hello"
                    _ = v11
                case "foo bar":
                    v2 += (v2 >> ((0) & 63))
                default:
                    v3 += 42
                }
                switch (("a" + "foo bar") + "a") {
                case "foo bar":
                    v12 := GenS1{F0: 1, F1: (v2 / ((v2) | 1)), F2: 4294967295}
                    _ = v12
                case "\n\t":
                    v4 += v4
                default:
                    var v13 uint16 = (v2 ^ v2)
                    _ = v13
                }
                _ = len(v6)
                _ = fi10
            }
            _ = v9
        }
        var v14 int32 = -1
        _ = v14
    } else {
        switch 4294967295 {
        case 42:
            fmt.Println(v4, v3, v4)
        case 555:
            v5 = append(v5, (v3 / ((v3) | 1)))
        case 32768:
            fmt.Println(v3, v3, v4)
        default:
            v3 = 2
        }
        v2 += v2
        v15 := GenS2{F0: v2, F1: 1}
        _ = v15
    }
    switch 127 {
    case 32768:
        switch 530 {
        case -581:
            var v16 int64 = 42
            _ = v16
        case 0:
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 uint32 = v3
                if (v1 <= 1e10) {
                    _ = len(v5)
                    var v19 string = "hello"
                    _ = v19
                } else {
                    v20 := GenS2{F0: 255, F1: 0}
                    _ = v20
                }
                v2 += v2
                _ = fi17
                _ = v18
            }
        case 32768:
            v4 -= 42
        }
    default:
        v21 := GenS1{F0: v2, F1: (v2 - 0), F2: 690}
        _ = v21
    }
    var v22 uint32 = 2
    v23 := GenS1{F0: v2, F1: (v2 - v2), F2: 1}
    v1 -= (v1 / 1.5)
    var v24 uint32 = (v22 << ((uint(v22)) & 63))
    v25 := GenS2{F0: v2, F1: 460}
    var v26 int64 = -1
    v27 := GenS2{F0: (^(v2)), F1: 56}
    v28 := GenS1{F0: 0, F1: v2, F2: 42}
    var v29 int8 = 56
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
