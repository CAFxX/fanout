package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
type GenS4 struct { F0 int32; F1 int; F2 int32 }
func genFn0() int16 {
    if true {
        var v7 int = -2147483648
        _ = v7
    } else {
        for fi8 := 0; fi8 < 15; fi8++ {
            if false {
                var v9 int32 = 0
                var v10 uint32 = 42
                _ = v10
                _ = v9
            } else {
                v11 := []int{}
                _ = len(v11)
                _ = v11
            }
            switch "" {
            case "foo bar":
                var v12 float64 = 0.0
                _ = v12
            case "\n\t":
                var v13 int64 = 42
                _ = v13
            case "a":
                switch ("\n\t" + "") {
                case "a":
                    var v14 uint32 = 2
                    _ = v14
                default:
                    var v15 int16 = 1
                    _ = v15
                }
            default:
                var v16 int16 = 1
                _ = v16
            }
            _ = fi8
        }
        var v17 bool = false
        v18 := []int{128, -638}
        _ = v17
        _ = v18
    }
    var v19 uint8 = 202
    v20 := GenS1{F0: -1}
    _ = v19
    _ = v20
    return -549
}
func genFn1(p0 string, p1 uint8) uint32 {
    p1 = p1
    var v21 string = "\n\t"
    if false {
        var v22 int8 = 42
        p1 -= 127
        _ = v22
    } else {
        var v23 uint16 = 65535
        var v24 float64 = -2.25
        v25 := GenS2{F0: 9223372036854775807}
        _ = v23
        _ = v24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v21
    return 189
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -647
    var v3 string = "\n\t"
    v4 := make(map[string]uint32)
    var v5 int = 1
    var v6 int = v5
    var v26 uint8 = 2
    v6 = 0
    var v27 int64 = v2
    var v28 float32 = ((-2.25 / 3.14159) - 3.14159)
    switch v3 {
    case "":
        var v29 uint = 608
        _ = v29
    default:
        fmt.Println(v6)
    }
    v30 := GenS3{F0: 2, F1: (v27 - -9223372036854775808), F2: 42}
    fmt.Println(v2)
    v31 := GenS4{F0: 133, F1: v6, F2: -2147483648}
    var v32 uint8 = (v26 - v26)
    if false {
        v3 = v3
        for fi33 := 0; fi33 < 11; fi33++ {
            v6 = -364
            var v34 uint = 0
            _ = fi33
            _ = v34
        }
    } else {
        var v35 uint = 18446744073709551615
        _ = v35
    }
    var v36 uint8 = v26
    switch v3 {
    case "":
        v37 := v4[v3]
        _ = v37
    case "a":
        var v38 float64 = ((3.14159 + 1.5) * ((((1.5 * (float64(v28) * 0.0)) - -2.25) - (0.0 / ((float64(v28) + 1e10) * 1e10))) - 0.1))
        _ = v38
    }
    v39 := []int{(^(v6)), v6, v6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
