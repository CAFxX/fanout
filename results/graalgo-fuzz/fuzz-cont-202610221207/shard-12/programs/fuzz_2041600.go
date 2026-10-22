package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint8; F2 string }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string) int32 {
    if false {
        v6 := GenS1{F0: 137, F1: 32767}
        switch "" {
        case "":
            var v7 float32 = ((3.14159 - (1.5 + 0.0)) / ((0.0 + (((1.5 - 0.1) + -2.25) / 1.0)) - (1.5 / 3.14159)))
            _ = v7
        case "hello":
            var v8 uint8 = 136
            _ = v8
        case "a":
            var v9 uint32 = 4294967295
            _ = v9
        default:
            for fi10 := 0; fi10 < 9; fi10++ {
                for fi11 := 0; fi11 < 20; fi11++ {
                    v12 := GenS2{F0: -701}
                    _ = fi11
                    _ = v12
                }
                p0 = "foo bar"
                v13 := GenS2{F0: 273}
                _ = fi10
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS1{F0: 42, F1: 255}
            fmt.Println(p0, fi14)
            _ = fi14
            _ = v15
        }
        _ = v6
    } else {
        var v16 int16 = 42
        v16 -= (-(v16))
        _ = v16
    }
    if true {
        v17 := []int{1}
        fmt.Println(p0)
        _ = v17
    } else {
        p0 = (p0 + "hello")
        p0 = p0
        v18 := GenS3{F0: 4294967295, F1: 255, F2: "hello"}
        _ = v18
    }
    _ = p0
    return 2147483647
}
func genFn1(p0 uint32, p1 uint64, p2 uint64) string {
    v19 := []int{-1, -2147483648, 613, 128}
    var v20 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return "a"
}

func main() {
    var v1 uint64 = 2
    var v2 int64 = 1
    var v3 bool = ((255 != 2) && true)
    var v4 uint16 = 2
    v5 := make(map[uint32]int)
    if (v3 && v3) {
        v21 := GenS3{F0: 2, F1: 255, F2: "\n\t"}
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 int = fi22
        for fi24 := 0; fi24 < 19; fi24++ {
            var v25 string = "a"
            var v26 uint64 = v1
            v4 += 693
            var v27 int16 = 0
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        var v28 uint8 = 48
        _ = fi22
        _ = v23
        _ = v28
    }
    v5[0] = 9223372036854775807
    v29 := GenS3{F0: 1, F1: 16, F2: "\n\t"}
    var v30 int32 = 749
    v2 = (v2 | (v2 / ((v2) | 1)))
    v31 := []int{-2147483648, 422, -975, 32768}
    v32 := GenS4{F0: v4}
    var v33 uint64 = 701
    if true {
        v3 = v3
    } else {
        var v34 uint32 = 4294967295
        v34 = 0
        var v35 int16 = -32768
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
