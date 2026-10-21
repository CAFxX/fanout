package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint8) uint {
    v4 := GenS1{F0: -1, F1: 0, F2: 0}
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint = 255
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{(^(fi7)), (fi7 + 42)}
        if (0 <= 1) {
            if ("a" >= "\n\t") {
                p0 = p0
                var v9 uint32 = 4294967295
                fmt.Println(fi7)
                _ = v9
            }
            v10 := GenS1{F0: fi7, F1: 18446744073709551615, F2: 9223372036854775807}
            var v11 float32 = (0.0 / 1.5)
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 bool = false
            _ = fi13
            _ = v14
        }
        v15 := []int{fi12}
        p0 -= p0
        _ = fi12
        _ = v15
    }
    v16 := GenS2{F0: p0, F1: 397, F2: -1}
    _ = p0
    _ = v16
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 128
    var v2 uint32 = 142
    var v3 int64 = -9223372036854775808
    v3 = (v3 >> ((2) & 63))
    v17 := GenS3{F0: true}
    v1 += (v1 >> ((genFn0(v1)) & 63))
    var v18 string = ""
    v19 := []int{128, -355}
    var v20 float64 = (1e10 / 1.0)
    switch v18 {
    case "a":
        for fi21 := 0; fi21 < 14; fi21++ {
            v2 = v2
            _ = fi21
        }
    case "hello":
        v20 += 3.14159
    }
    if false {
        var v22 uint = 65535
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS1{F0: (fi23 | fi23), F1: 502, F2: 9223372036854775807}
            switch fi23 {
            case 42:
                v20 = v20
            default:
                v18 = v18
            }
            var v25 int16 = 0
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v22
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 uint = genFn0((v1 + v1))
            switch v18 {
            case "foo bar":
                v1 += (v1 + (v1 << ((v28) & 63)))
            case "a":
                v18 = ""
            default:
                switch fi27 {
                case 947:
                    v20 += 0.0
                }
            }
            var v29 uint32 = v2
            fmt.Println(fi27)
            fmt.Println(v18, v3, v28)
            v1 = 68
            v30 := GenS3{F0: (1 != 1)}
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        v31 := GenS2{F0: (v1 & v1), F1: 255, F2: 0}
        _ = fi26
        _ = v31
    }
    v20 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
