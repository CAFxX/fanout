package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint }
func genFn0() int {
    v7 := GenS1{F0: 255}
    var v8 float32 = 0.0
    var v9 uint32 = 697
    if (0.1 != (((0.1 + (((-2.25 * (0.0 * (0.1 * 3.14159))) + (((-2.25 / 3.14159) / (1e10 * 0.1)) / 1e10)) / 1e10)) + 3.14159) / 1.0)) {
        v9 -= (v9 + v9)
    }
    _ = v7
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 int8) string {
    var v10 int32 = 514
    switch "\n\t" {
    case "hello":
        var v11 bool = ((!false) || ((true && ((!(false || (((0.0 * 1e10) / 1.5) > 1.5))) || (255 >= 687))) || (true && false)))
        _ = v11
    case "\n\t":
        p0 -= (p0 >> ((65535) & 63))
    default:
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int16 = -734
            for fi14 := 0; fi14 < 10; fi14++ {
                switch fi12 {
                case -89:
                    p0 = 1
                default:
                    var v15 uint = 2
                    _ = v15
                }
                if true {
                    v16 := []int{108}
                    p0 = p0
                    switch (fi12 / ((genFn0()) | 1)) {
                    case 1:
                        var v17 float64 = (0.1 * -2.25)
                        _ = v17
                    }
                    _ = v16
                } else {
                    v18 := []int{}
                    _ = v18
                }
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
    }
    switch 0 {
    case 32767:
        v10 -= v10
    case -1:
        v19 := []int{}
        _ = v19
    case 604:
        v20 := []int{377, 4294967295}
        _ = v20
    }
    fmt.Println(v10, p0)
    _ = p0
    _ = v10
    return (("a" + "foo bar") + "foo bar")
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"foo bar", "a", "a", "a"}
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    v5 := []uint32{353, 2, 2147483648}
    v6 := []uint32{2147483647}
    var v21 uint16 = 255
    for fi22 := 0; fi22 < 12; fi22++ {
        _ = len(v2)
        v23 := v3[fi22]
        _ = fi22
        _ = v23
    }
    if false {
        var v24 string = ("" + "\n\t")
        fmt.Println(v21, v24)
        _ = v24
    } else {
        var v25 float32 = 1.5
        _ = v25
    }
    var v26 int64 = 1
    v21 = 2
    v27 := GenS2{F0: 2, F1: 127}
    var v28 int16 = 1
    v29 := GenS3{F0: 4294967295}
    var v30 float64 = (((0.0 / 3.14159) * 0.0) - (0.0 / (1e10 * 0.1)))
    if (!(!false)) {
        var v31 int = genFn0()
        fmt.Println(v31)
        _ = v31
    } else {
        v5 = append(v5, 2147483647)
    }
    var v32 float64 = ((v30 - v30) + v30)
    fmt.Println(v30, v30)
    v33 := GenS2{F0: 0, F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
