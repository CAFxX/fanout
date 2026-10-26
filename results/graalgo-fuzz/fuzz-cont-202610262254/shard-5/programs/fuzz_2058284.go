package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint; F1 bool }
func genFn0(p0 int, p1 uint32, p2 uint16) bool {
    var v6 uint8 = 0
    var v7 uint16 = (p2 & p2)
    if true {
        var v8 bool = false
        var v9 uint = 0
        _ = v8
        _ = v9
    }
    fmt.Println(p2)
    v7 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (1.5 <= float64((((3.14159 * 1e10) - float32(0.1)) / (-2.25 * -2.25))))
}
func genFn1(p0 int64) int32 {
    p0 -= p0
    if true {
        v10 := []int{128, 127}
        _ = v10
    }
    _ = p0
    return 923
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 uint64 = 2
    var v4 int32 = 0
    var v5 bool = (!(-227 != int16(v3)))
    v11 := GenS1{F0: (v4 << ((4294967295) & 63)), F1: 42, F2: 42}
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint16 = 42
        _ = fi12
        _ = v13
    }
    var v14 int8 = 100
    var v15 uint64 = v3
    switch 1 {
    case 42:
        switch 32767 {
        case 32768:
            v16 := GenS2{F0: 4294967295, F1: true}
            _ = v16
        case 128:
            fmt.Println(v14, v15)
        }
    case 127:
        v3 = v15
    default:
        v17 := []int{129, 299, -171, 32768}
        _ = v17
    }
    var v18 uint = 18446744073709551615
    v19 := []int{}
    if ((0.1 + (1.5 * 1.5)) != (((1e10 / 1.0) + -2.25) * 0.0)) {
        v14 -= v14
        for fi20 := 0; fi20 < 6; fi20++ {
            v3 -= (v15 / (((v3 - (v15 | 18446744073709551615))) | 1))
            for fi21 := 0; fi21 < 10; fi21++ {
                if v5 {
                    v4 = v4
                }
                fmt.Println(fi21)
                var v22 int8 = (^(v14))
                _ = fi21
                _ = v22
            }
            v23 := GenS1{F0: (v4 | v4), F1: -890, F2: (v15 >> ((v18) & 63))}
            _ = fi20
            _ = v23
        }
        var v24 float32 = ((0.1 + -2.25) / ((0.0 - 1.5) - ((0.0 * ((1.5 / 1.5) * (float32(((-2.25 + 3.14159) + ((0.0 / 1e10) - float64(1.5)))) / 3.14159))) - 0.0)))
        _ = v24
    }
    fmt.Println(v3, v18, v3)
    for fi25 := 0; fi25 < 18; fi25++ {
        if ((!(fi25 >= fi25)) && v5) {
            var v26 string = ((("hello" + "\n\t") + (("a" + "hello") + "a")) + "\n\t")
            v27 := GenS1{F0: v4, F1: 711, F2: 2}
            _ = v26
            _ = v27
        } else {
            var v28 int16 = 494
            for fi29 := 0; fi29 < 7; fi29++ {
                fmt.Println(v15)
                switch fi29 {
                case 128:
                    switch "a" {
                    case "":
                        if ("" <= "hello") {
                            v15 = v3
                            var v30 int64 = 605
                            v31 := GenS2{F0: (^(v18)), F1: true}
                            _ = v30
                            _ = v31
                        }
                    case "foo bar":
                        v2["a"] = false
                    }
                case -1:
                    var v32 uint32 = 65535
                    _ = v32
                default:
                    v4 += (v4 << ((253) & 63))
                }
                var v33 uint16 = 32767
                _ = fi29
                _ = v33
            }
            var v34 uint32 = 1
            _ = v28
            _ = v34
        }
        fmt.Println(v3, v5, v5)
        _ = len(v2)
        _ = fi25
    }
    var v35 bool = v5
    if len(v19) > 0 {
        v19[0] = 127
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
