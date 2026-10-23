package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 uint8, p1 int64, p2 uint32) int64 {
    v7 := GenS1{F0: 0, F1: 42}
    var v8 uint64 = 4294967295
    var v9 float32 = (float32(-2.25) * 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p1 / ((9223372036854775807) | 1))
}
func genFn1(p0 int32, p1 uint16, p2 int32) bool {
    var v10 int32 = 1
    if true {
        v11 := GenS2{F0: (^(p1))}
        p1 = 255
        _ = v11
    }
    var v12 uint64 = 2
    v13 := GenS2{F0: (p1 * (p1 | (p1 ^ p1)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v12
    _ = v13
    return (true && ("" <= "hello"))
}

func main() {
    v1 := []string{((((("hello" + "\n\t") + "a") + "\n\t") + ("hello" + "")) + ("foo bar" + "hello"))}
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    var v5 int8 = 10
    var v6 bool = true
    var v14 int16 = int16(v5)
    switch "\n\t" {
    case "hello":
        v15 := GenS3{F0: 4294967295, F1: 18446744073709551615}
        _ = v15
    case "":
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 uint32 = 2147483648
            _ = fi16
            _ = v17
        }
    default:
        switch "hello" {
        case "\n\t":
            v6 = (((0.0 / 1e10) > 0.0) && true)
        case "a":
            var v18 int32 = -1
            _ = v18
        case "foo bar":
            v19 := GenS1{F0: 310, F1: 88}
            _ = v19
        }
    }
    var v20 string = ""
    if v6 {
        fmt.Println(v6)
        if (!true) {
            v14 -= (v14 % ((v14) | 1))
            var v21 uint8 = 127
            _ = v21
        }
        v22 := GenS2{F0: 213}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 7; fi23++ {
            v5 += v5
            if v6 {
                for fi24 := 0; fi24 < 11; fi24++ {
                    fmt.Println(v14, v5, v6)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi24
                }
                for fi25 := 0; fi25 < 9; fi25++ {
                    v5 += -128
                    v26 := GenS2{F0: 266}
                    _ = fi25
                    _ = v26
                }
            } else {
                v20 = v20
                var v27 bool = (1 > 171)
                _ = v27
            }
            v4["\n\t"] = (422 != 2147483648)
            _ = fi23
        }
        _ = len(v2)
    }
    v14 = (v14 % ((81) | 1))
    var v28 uint64 = 228
    var v29 uint16 = 255
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := v4[v20]
        _ = fi30
        _ = v31
    }
    delete(v2, 42)
    var v32 uint16 = 900
    var v33 int8 = (-(v5))
    if (v14 <= (v14 | 994)) {
        var v34 string = v20
        _ = v34
    } else {
        v29 = (v29 / ((v32) | 1))
        v32 = (^(v29))
        var v35 uint = 653
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
