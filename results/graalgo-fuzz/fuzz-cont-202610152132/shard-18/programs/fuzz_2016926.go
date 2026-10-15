package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int32; F2 int16 }
type GenS4 struct { F0 int; F1 int64 }
func genFn0(p0 int, p1 int8) uint64 {
    v6 := []int{p0, 42, p0}
    v7 := GenS1{F0: int(p1), F1: ("a" + "hello"), F2: false}
    v8 := GenS1{F0: p0, F1: ("\n\t" + "a"), F2: ((-1 > -143) && ((^(p0)) >= p0))}
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 bool = (1 <= 0)
            var v12 float32 = float32(0.1)
            var v13 int16 = 42
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        v14 := GenS2{F0: p1}
        _ = fi9
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int8 = 100
    var v2 int = 127
    var v3 uint64 = 42
    var v4 int8 = (v1 * v1)
    v5 := []uint32{}
    v15 := GenS2{F0: (^(v4))}
    switch "hello" {
    case "hello":
        v1 = (v4 << ((0) & 63))
    case "\n\t":
        v4 -= v1
    case "":
        if (v1 < 71) {
            if true {
                for fi16 := 0; fi16 < 1; fi16++ {
                    if len(v5) > 0 {
                        v5[0] = 2
                    }
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v18 := GenS3{F0: (true || true), F1: 2147483647, F2: -32768}
                        v19 := GenS2{F0: v1}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = fi16
                }
                if false {
                    fmt.Println(v3)
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            } else {
                _ = len(v5)
                v20 := GenS4{F0: v2, F1: 1}
                var v21 int = 128
                _ = v20
                _ = v21
            }
            switch "" {
            case "\n\t":
                v2 = 1
            default:
                v4 += (v1 - int8(v2))
            }
        }
    }
    fmt.Println(v1, v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch v2 {
    case 261:
        var v22 uint = 0
        _ = v22
    case 42:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        v5 = append(v5, 2)
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        fmt.Println(fi23)
        _ = len(v5)
        var v24 float64 = -2.25
        _ = fi23
        _ = v24
    }
    v3 -= (^(v3))
    v25 := GenS3{F0: ((2147483647 <= 4294967295) && (false || false)), F1: 0, F2: 42}
    var v26 uint32 = 2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
