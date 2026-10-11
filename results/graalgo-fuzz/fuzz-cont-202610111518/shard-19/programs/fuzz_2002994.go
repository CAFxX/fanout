package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int; F1 int8; F2 uint64 }
func genFn0() uint8 {
    var v6 string = ("" + "")
    for fi7 := 0; fi7 < 20; fi7++ {
        v6 = ((v6 + "hello") + ("hello" + v6))
        if (1000 >= -1) {
            v6 = v6
        } else {
            if true {
                switch fi7 {
                case 4294967295:
                    if true {
                        v6 = "hello"
                    }
                case -386:
                    var v8 int = 2147483647
                    _ = v8
                default:
                    switch (fi7 / ((1) | 1)) {
                    case -9223372036854775808:
                        v9 := GenS1{F0: false, F1: -937}
                        _ = v9
                    default:
                        v10 := []int{fi7, (-(fi7))}
                        _ = v10
                    }
                }
                v11 := GenS1{F0: ((1e10 * 1e10) >= (((-2.25 + (0.1 - -2.25)) + 1e10) * 1e10)), F1: 916}
                _ = v11
            } else {
                v12 := GenS1{F0: (1.5 == 1.5), F1: -9223372036854775808}
                _ = v12
            }
            v6 = "foo bar"
        }
        _ = fi7
    }
    v6 = v6
    var v13 uint64 = 4294967295
    _ = v13
    _ = v6
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"foo bar", "a"}
    v3 := make(map[uint32]bool)
    v4 := []string{"hello", "", "foo bar", ("" + "a")}
    v5 := make(map[int]int)
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := v1[97]
        v15 = ((0.0 + (1.5 / -2.25)) < 1e10)
        fmt.Println(fi14, fi14)
        _ = fi14
        _ = v15
    }
    var v16 int = 42
    var v17 int64 = -1
    v16 += (-(v16))
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 int8 = 42
        for fi20 := 0; fi20 < 7; fi20++ {
            v19 = v19
            _ = fi20
        }
        if true {
            var v21 uint16 = 255
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    v22 := GenS1{F0: (!true), F1: v17}
    for fi23 := 0; fi23 < 17; fi23++ {
        if true {
            _ = len(v4)
            fmt.Println(fi23)
        }
        v24 := GenS1{F0: false, F1: v17}
        var v25 uint64 = 2
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 bool = true
    var v27 float64 = (-2.25 * (1e10 * 1e10))
    v27 = (v27 / v27)
    v16 -= v16
    v27 += v27
    switch ("" + "a") {
    case "foo bar":
        var v28 float32 = 3.14159
        _ = v28
    case "\n\t":
        v26 = v26
    case "hello":
        v29 := GenS2{F0: v16, F1: -51, F2: 63}
        _ = v29
    }
    var v30 uint32 = 4294967295
    switch "" {
    case "foo bar":
        v27 += ((((v27 - v27) / -2.25) + v27) * ((v27 + -2.25) - 1.5))
    case "\n\t":
        var v31 uint = 1
        _ = v31
    default:
        var v32 uint64 = 734
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
