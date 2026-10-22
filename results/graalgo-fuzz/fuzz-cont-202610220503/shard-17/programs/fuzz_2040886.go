package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 string; F1 int8 }
func genFn0(p0 uint, p1 int64, p2 int8) bool {
    p0 = p0
    v6 := []int{-1, 4294967295, 951}
    for fi7 := 0; fi7 < 7; fi7++ {
        switch "foo bar" {
        case "":
            var v8 uint64 = 0
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 20; fi9++ {
                fmt.Println(p0)
                if true {
                    v10 := GenS1{F0: 510, F1: 182, F2: fi9}
                    fmt.Println(p1, p1)
                    v6 = append(v6, (^(fi9)))
                    _ = v10
                }
                var v11 string = ("hello" + "a")
                _ = fi9
                _ = v11
            }
        }
        v6 = append(v6, (fi7 | fi7))
        _ = fi7
    }
    switch "a" {
    case "a":
        p0 += (p0 & 20)
    case "":
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (2147483647 != -720)
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = v1
    var v3 int = 128
    v4 := make(map[string]bool)
    var v5 int64 = -1
    fmt.Println(v1)
    v12 := []int{v3}
    v4[("" + "foo bar")] = (float32(1.5) >= (float32((1.5 + float64(0.0))) + 0.1))
    fmt.Println(v5, v3, v5)
    if false {
        v3 -= v3
    }
    fmt.Println(v2, v3)
    var v13 uint16 = 2
    v13 = (^(v13))
    var v14 string = "foo bar"
    v15 := GenS1{F0: v13, F1: 255, F2: -9223372036854775808}
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 float64 = (3.14159 + ((float64(float32(0.1)) - 1e10) * 0.0))
        switch (^(fi16)) {
        case 1:
            var v18 int16 = 318
            _ = v18
        default:
            for fi19 := 0; fi19 < 10; fi19++ {
                v17 -= (1.5 * v17)
                v5 -= (v5 + -9223372036854775808)
                v3 -= fi16
                _ = fi19
            }
        }
        _ = fi16
        _ = v17
    }
    if false {
        v2 = (v2 - -636)
        var v20 float32 = 1e10
        v20 -= v20
        _ = v20
    } else {
        var v21 int8 = -128
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := GenS1{F0: 21, F1: 2, F2: 1}
            v4[(v14 + v14)] = true
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v24 float64 = -2.25
    v2 -= -1
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 uint = 0
        _ = fi25
        _ = v26
    }
    v2 -= (v2 - 424)
    v27 := GenS2{F0: v14, F1: -117}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
