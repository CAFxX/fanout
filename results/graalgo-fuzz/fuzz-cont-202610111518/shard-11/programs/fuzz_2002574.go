package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 uint16, p1 uint8, p2 uint16) string {
    var v4 uint32 = 708
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "a"
}
func genFn1() uint8 {
    var v5 uint32 = 2147483647
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            switch "a" {
            case "\n\t":
                v8 := []int{fi6, (^(fi6))}
                _ = v8
            case "foo bar":
                switch 42 {
                case -656:
                    if true {
                        v5 -= 4294967295
                        switch "\n\t" {
                        case "\n\t":
                            v9 := []int{}
                            _ = v9
                        }
                        v10 := []int{466, fi7}
                        _ = v10
                    }
                }
            case "hello":
                v11 := GenS1{F0: 32767}
                _ = v11
            }
            fmt.Println(fi7)
            v5 -= 2147483647
            _ = fi7
        }
        _ = fi6
    }
    if (74 == 626) {
        if (!false) {
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 float64 = (0.1 * (0.1 + (((float64(1e10) - (3.14159 * (0.1 * 1e10))) * 3.14159) * ((((1.5 + 1e10) * (((0.0 + float64(0.0)) + (0.0 - 1e10)) * 0.1)) - 0.1) / float64(((((((1e10 + (1.5 + float32((((float64(((1.5 / 0.1) - float32(((float64(1.5) * (((float64((float32(((0.0 * (0.1 / 0.1)) - (3.14159 - -2.25))) * 0.1)) / (0.1 / -2.25)) - 1.5) + 0.1)) - (3.14159 / (1.5 / (-2.25 + 0.1))))))) * 3.14159) * 1.5) / (float64(-2.25) - 1.5))))) / float32(0.1)) * 1.5) / 1e10) * 1.5) * 1e10))))))
                _ = fi12
                _ = v13
            }
            fmt.Println(v5)
            var v14 uint16 = 1
            _ = v14
        } else {
            var v15 bool = true
            var v16 uint16 = 227
            _ = v15
            _ = v16
        }
    }
    _ = v5
    return 189
}

func main() {
    var v1 int = 771
    var v2 int8 = -58
    var v3 int16 = 42
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 bool = true
        _ = fi17
        _ = v18
    }
    v2 -= v2
    switch (v1 | v1) {
    case -2147483648:
        fmt.Println(v2, v1, v1)
    case 173:
        v19 := GenS1{F0: v3}
        _ = v19
    }
    fmt.Println(v2, v1, v1)
    switch "\n\t" {
    case "":
        fmt.Println(v3, v2, v3)
    case "a":
        v2 -= (v2 & (-(v2)))
    default:
        fmt.Println(v2, v1)
    }
    var v20 uint = 18446744073709551615
    v21 := GenS1{F0: v3}
    v22 := GenS2{F0: 2, F1: v20}
    fmt.Println(v20)
    var v23 string = "\n\t"
    fmt.Println(v2, v3)
    v24 := []int{9223372036854775807, v1, v1, (v1 % ((v1) | 1))}
    fmt.Println(v20, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
