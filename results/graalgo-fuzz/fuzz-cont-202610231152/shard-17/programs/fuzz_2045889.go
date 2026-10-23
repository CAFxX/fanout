package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 string; F1 string }
func genFn0(p0 uint8, p1 int64) uint32 {
    v4 := []int{128, -1, 153}
    if (false && true) {
        fmt.Println(p0)
    }
    fmt.Println(p0, p0)
    var v5 float32 = (0.1 - (0.1 + 0.1))
    for fi6 := 0; fi6 < 18; fi6++ {
        switch fi6 {
        case -826:
            p0 += p0
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483648
}
func genFn1(p0 int64, p1 int16, p2 string) int16 {
    var v7 bool = false
    v8 := GenS1{F0: 1, F1: 42, F2: 1}
    var v9 bool = v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 string = ("\n\t" + (("a" + "hello") + "\n\t"))
    v2 := []string{v1, "foo bar", (("\n\t" + "") + ("foo bar" + v1)), v1}
    var v3 uint16 = 255
    v10 := GenS2{F0: (v1 + (v1 + (v1 + v1))), F1: v1}
    var v11 int = -2147483648
    v3 = 2
    v3 += v3
    var v12 uint32 = genFn0(127, -971)
    v3 += (v3 | v3)
    var v13 float64 = (1.5 / 3.14159)
    switch 0 {
    case 128:
        if false {
            if true {
                var v14 int8 = -119
                fmt.Println(v1, v14)
                _ = v14
            }
        }
    case -35:
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS2{F0: v1, F1: ("\n\t" + ("" + v1))}
            _ = fi15
            _ = v16
        }
    case -71:
        fmt.Println(v11, v13, v1)
    default:
        v13 -= 3.14159
    }
    v2 = append(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
