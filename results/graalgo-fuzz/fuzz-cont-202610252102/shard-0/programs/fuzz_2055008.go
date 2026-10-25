package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint; F2 bool }
func genFn0(p0 uint32, p1 int32, p2 uint32) string {
    var v4 uint64 = 4294967295
    v5 := GenS1{F0: 540}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "\n\t"
}

func main() {
    var v1 int8 = 1
    var v2 uint8 = 32
    v3 := []string{}
    for fi6 := 0; fi6 < 15; fi6++ {
        v1 += v1
        _ = fi6
    }
    v7 := GenS1{F0: 640}
    fmt.Println(v1, v2)
    if (true && false) {
        v2 = (v2 << ((484) & 63))
        v1 -= v1
    }
    if (42 > 865) {
        var v8 uint16 = 138
        v9 := GenS1{F0: -283}
        _ = v8
        _ = v9
    }
    var v10 uint32 = 257
    if (false || (false || (false && (42 == -32768)))) {
        v1 = -128
    }
    switch -1 {
    case -2147483648:
        var v11 int64 = 42
        _ = v11
    default:
        switch ("foo bar" + "foo bar") {
        case "a":
            v10 = v10
        case "\n\t":
            v12 := GenS1{F0: -1}
            _ = v12
        default:
            if (false && (true || true)) {
                v10 += v10
                v10 += (^(v10))
            } else {
                var v13 uint32 = (v10 / (((v10 ^ v10)) | 1))
                v14 := GenS1{F0: 4294967295}
                _ = v13
                _ = v14
            }
        }
    }
    fmt.Println(v1, v1)
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            _ = len(v3)
            v17 := GenS2{F0: 1}
            v1 = v1
            v18 := GenS2{F0: (^(v2))}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v19 := GenS3{F0: 1, F1: 2, F2: (v2 <= (^(v2)))}
            v1 -= (v1 | (v1 | v1))
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi15
    }
    switch ("a" + "hello") {
    case "":
        var v20 float32 = 0.1
        _ = v20
    case "\n\t":
        v3 = append(v3, "a")
    }
    var v21 uint = 304
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v7
}
