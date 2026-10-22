package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int32; F1 uint64; F2 uint32 }

func main() {
    var v1 int16 = 1000
    v2 := []string{}
    var v3 int32 = 42
    var v4 uint8 = 127
    v5 := make(map[uint32]int)
    var v6 uint = 255
    v1 -= v1
    if false {
        v1 = -32768
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        _ = len(v2)
        v8 := GenS1{F0: 32767, F1: false}
        _ = fi7
        _ = v8
    }
    if true {
        v9 := v5[0]
        v9 = 2147483647
        v3 += v3
        _ = v9
    } else {
        if (!(false && (42 >= 32767))) {
            v10 := GenS2{F0: (v3 | v3), F1: 594, F2: 211}
            var v11 uint16 = 895
            v4 = 127
            _ = v10
            _ = v11
        } else {
            switch "" {
            case "":
                v5[0] = -17
            case "a":
                v1 = (v1 & v1)
            default:
                v1 -= (v1 << ((0) & 63))
            }
        }
        fmt.Println(v4, v6, v3)
        switch -99 {
        case 0:
            var v12 int32 = 42
            _ = v12
        default:
            v6 -= 255
        }
    }
    v6 = (^(v6))
    v6 -= 2
    switch -2147483648 {
    case -563:
        for fi13 := 0; fi13 < 1; fi13++ {
            if true {
                var v14 float32 = 3.14159
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v14 = v14
                _ = v14
            } else {
                v4 = (v4 << ((0) & 63))
            }
            v4 = 249
            _ = fi13
        }
    case 4294967295:
        fmt.Println(v4)
    }
    var v15 int8 = -1
    v16 := GenS2{F0: 0, F1: 1, F2: 65535}
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 uint = 65535
        if (0.1 == (-2.25 * 1e10)) {
            var v19 int64 = 9223372036854775807
            _ = v19
        }
        fmt.Println(fi17, v18, v18)
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
