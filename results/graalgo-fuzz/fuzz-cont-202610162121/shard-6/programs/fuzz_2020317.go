package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int32 }
func genFn0(p0 string, p1 int32, p2 int32) int32 {
    p1 = p1
    var v5 bool = false
    v5 = v5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]int)
    var v3 int = 32768
    v4 := []int{(v3 % ((v3) | 1)), (v3 >> ((4294967295) & 63)), v3, v3}
    v6 := GenS1{F0: 2, F1: false}
    for fi7 := 0; fi7 < 1; fi7++ {
        if false {
            v2["\n\t"] = v3
            var v8 uint16 = 65535
            fmt.Println(v8, v1, v8)
            _ = v8
        }
        switch (^(fi7)) {
        case 128:
            v9 := GenS1{F0: 193, F1: true}
            _ = v9
        case -576:
            var v10 int16 = 1
            _ = v10
        case 32767:
            v3 -= fi7
        default:
            switch "" {
            case "":
                v1 = 1
            }
        }
        _ = fi7
    }
    var v11 uint8 = 255
    v12 := GenS1{F0: (v11 * v11), F1: (!false)}
    v1 += v1
    v11 = (v11 & 2)
    v1 -= (v1 ^ (v1 / (((v1 % ((-128) | 1))) | 1)))
    fmt.Println(v3, v3, v1)
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 float32 = 1.5
        var v15 bool = false
        if (42 <= 1) {
            var v16 uint = 18446744073709551615
            _ = v16
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    if false {
        var v17 bool = true
        v18 := GenS2{F0: 255}
        for fi19 := 0; fi19 < 5; fi19++ {
            fmt.Println(v11, v1, v1)
            _ = fi19
        }
        _ = v17
        _ = v18
    } else {
        v4 = append(v4, 9223372036854775807)
        v3 += 32768
    }
    switch "" {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "":
        var v20 int32 = 108
        _ = v20
    default:
        if false {
            v21 := GenS3{F0: -335, F1: 2, F2: 4294967295}
            switch (v3 % ((v3) | 1)) {
            case 32768:
                var v22 int64 = -377
                _ = v22
            }
            switch "a" {
            case "":
                for fi23 := 0; fi23 < 6; fi23++ {
                    v24 := GenS3{F0: 336, F1: 141, F2: 306}
                    _ = fi23
                    _ = v24
                }
            case "hello":
                var v25 int16 = -78
                _ = v25
            }
            _ = v21
        }
    }
    var v26 uint16 = 2
    v27 := GenS4{F0: 255, F1: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v6
}
