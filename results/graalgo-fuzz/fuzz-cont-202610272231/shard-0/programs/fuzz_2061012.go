package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 string }
type GenS3 struct { F0 int; F1 uint8; F2 int }
func genFn0(p0 uint8) int32 {
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int8 = 42
            fmt.Println(p0)
            _ = fi7
            _ = v8
        }
        p0 += p0
        _ = fi6
    }
    if ((1 == 391) && ((("" + "") + "hello") <= "foo bar")) {
        var v9 uint64 = 4294967295
        fmt.Println(v9, p0)
        p0 -= (p0 - p0)
        _ = v9
    }
    _ = p0
    return -2147483648
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]int)
    v3 := []int{128, 9223372036854775807}
    v4 := make(map[int]int)
    var v5 uint32 = 0
    if (361 > 255) {
        var v10 uint8 = 65
        var v11 int = 2147483647
        if true {
            for fi12 := 0; fi12 < 18; fi12++ {
                if len(v3) > 0 {
                    v3[0] = (fi12 * fi12)
                }
                _ = fi12
            }
            v11 += 54
            var v13 int = (-(v11))
            _ = v13
        } else {
            if true {
                for fi14 := 0; fi14 < 17; fi14++ {
                    v15 := GenS1{F0: -939}
                    v1 = ""
                    var v16 uint32 = (v5 / (((^(v5))) | 1))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                var v17 int64 = -65
                var v18 uint64 = 934
                _ = v17
                _ = v18
            }
            v10 = v10
            switch v1 {
            case "a":
                var v19 int = v11
                _ = v19
            case "hello":
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 float64 = (-2.25 - 1e10)
                    _ = fi20
                    _ = v21
                }
            default:
                var v22 float32 = 0.1
                _ = v22
            }
        }
        _ = v10
        _ = v11
    }
    var v23 int = -153
    if len(v3) > 0 {
        v3[0] = (v23 + -2147483648)
    }
    v5 = (^(v5))
    switch v23 {
    case 42:
        v23 = 42
    }
    var v24 int32 = 1
    var v25 float32 = 1.5
    var v26 int64 = 0
    var v27 int16 = 0
    var v28 int16 = v27
    switch v1 {
    case "":
        var v29 uint16 = 65535
        _ = v29
    case "a":
        var v30 int32 = genFn0(127)
        _ = v30
    case "\n\t":
        for fi31 := 0; fi31 < 8; fi31++ {
            fmt.Println(v23, v24, v23)
            v32 := GenS2{F0: 185, F1: 1, F2: "\n\t"}
            v33 := GenS3{F0: fi31, F1: 42, F2: (v23 % ((fi31) | 1))}
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
