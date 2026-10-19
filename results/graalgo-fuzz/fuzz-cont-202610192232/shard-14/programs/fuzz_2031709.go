package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8) int {
    v5 := GenS1{F0: 1, F1: 4294967295}
    var v6 int8 = p0
    _ = p0
    _ = v5
    _ = v6
    return -29
}
func genFn1(p0 string) int {
    if true {
        fmt.Println(p0)
        for fi7 := 0; fi7 < 5; fi7++ {
            fmt.Println(fi7, p0)
            switch p0 {
            case "a":
                p0 = p0
            case "hello":
                p0 = "a"
            case "foo bar":
                var v8 uint32 = 605
                _ = v8
            }
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := GenS2{F0: 2}
                v11 := []int{(fi7 >> ((255) & 63)), (fi7 | fi9), fi7, fi7}
                p0 = "hello"
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi7
        }
        fmt.Println(p0)
    }
    p0 = p0
    _ = p0
    return genFn0(0)
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    v4 := []int{4294967295, 1, 9223372036854775807}
    var v12 float64 = 1e10
    if true {
        v13 := GenS2{F0: 255}
        var v14 bool = (!false)
        _ = v13
        _ = v14
    } else {
        v1 = v1
    }
    var v15 uint = 4294967295
    v12 += v12
    v15 = v15
    v16 := GenS1{F0: (v15 + (^(v15))), F1: v1}
    var v17 float32 = (1.5 * (0.1 + -2.25))
    for fi18 := 0; fi18 < 1; fi18++ {
        fmt.Println(v12)
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(v17, fi18, fi19)
            _ = fi19
        }
        _ = fi18
    }
    var v20 string = ""
    var v21 float64 = (v12 - float64(v17))
    var v22 uint16 = 364
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            switch "" {
            case "hello":
                fmt.Println(fi24, v20, v17)
            default:
                fmt.Println(fi23, fi24)
            }
            v1 = 42
            v22 -= 42
            _ = fi24
        }
        if false {
            for fi25 := 0; fi25 < 6; fi25++ {
                v26 := GenS1{F0: v15, F1: (v1 ^ 42)}
                v12 += (v21 + (((3.14159 * ((v21 / ((v12 / v12) - float64(v17))) + (v12 * v21))) / ((v21 * float64(v17)) * v12)) + v12))
                switch ("foo bar" + ("" + v20)) {
                case "hello":
                    if (v20 != v20) {
                        v1 -= v1
                        var v27 int32 = -996
                        _ = v27
                    }
                case "\n\t":
                    var v28 uint8 = 127
                    _ = v28
                case "a":
                    v29 := GenS2{F0: (v22 - 255)}
                    _ = v29
                }
                _ = fi25
                _ = v26
            }
            v15 += 1
        } else {
            var v30 bool = (v20 > "a")
            v31 := v2[v1]
            _ = v30
            _ = v31
        }
        _ = fi23
    }
    var v32 uint64 = 18446744073709551615
    var v33 float32 = (v17 * (float32(v12) * (v17 * (v17 / v17))))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
