package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS5 struct { F0 int16 }
func genFn0(p0 bool, p1 int32) uint64 {
    if (1000 < 779) {
        var v6 int64 = 1
        if ((p0 && p0) || ("a" < "hello")) {
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := GenS1{F0: 2147483648}
                _ = fi7
                _ = v8
            }
            var v9 float32 = float32(((((float64(((1.5 + 0.1) * 1e10)) / 1e10) * (0.1 - (1.5 / (-2.25 - 1e10)))) + float64(1e10)) / 1.5))
            var v10 int8 = 100
            _ = v10
            _ = v9
        } else {
            v11 := GenS1{F0: 445}
            if true {
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 int16 = 0
                    var v14 uint = 42
                    var v15 string = ("a" + ("hello" + "\n\t"))
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                v16 := GenS1{F0: 2147483647}
                var v17 float32 = 1e10
                _ = v16
                _ = v17
            }
            _ = v11
        }
        _ = v6
    }
    var v18 float64 = 3.14159
    switch ("hello" + "\n\t") {
    case "foo bar":
        v19 := []int{-157, -796, -2147483648}
        _ = v19
    default:
        p1 = 0
    }
    var v20 float32 = 0.1
    switch "foo bar" {
    case "hello":
        fmt.Println(v18, p1, v18)
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v20
    return 956
}
func genFn1(p0 bool, p1 uint32, p2 int8) int16 {
    fmt.Println(p2, p0)
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(p0, p0, p1)
        v22 := GenS1{F0: (p1 * (p1 ^ (p1 / ((2147483648) | 1))))}
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := []int{(-(fi21)), 70}
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := []int{}
    var v2 uint16 = 42
    var v3 float32 = 3.14159
    var v4 uint32 = uint32(v2)
    v5 := make(map[int]uint32)
    v25 := v5[-103]
    fmt.Println(v2, v25, v4)
    v2 -= 65535
    v3 += (-2.25 - 3.14159)
    v26 := GenS2{F0: 0, F1: 42}
    var v27 bool = false
    v28 := GenS2{F0: 42, F1: -2147483648}
    if len(v1) > 0 {
        v1[0] = 32768
    }
    switch -32 {
    case 42:
        var v29 string = ("\n\t" + "\n\t")
        _ = v29
    default:
        var v30 uint16 = (v2 << ((42) & 63))
        _ = v30
    }
    fmt.Println(v25, v25)
    v31 := GenS3{F0: -262, F1: true}
    var v32 int8 = 0
    v1 = append(v1, 553)
    var v33 int32 = 832
    if ((v2 - (v2 * (^(v2)))) <= v2) {
        if ((v32 < 42) || (8 <= 1)) {
            var v34 int32 = (v33 * v33)
            switch "" {
            case "hello":
                v35 := GenS2{F0: 1, F1: -744}
                _ = v35
            case "\n\t":
                v36 := GenS4{F0: 1, F1: 52, F2: 1000}
                _ = v36
            }
            _ = v34
        }
        v37 := GenS1{F0: v25}
        _ = v37
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        v39 := GenS5{F0: 0}
        _ = fi38
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
