package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int64, p1 string) uint {
    v5 := []int{-688}
    p0 -= p0
    p0 = p0
    _ = p0
    _ = p1
    _ = v5
    return 196
}
func genFn1(p0 int8, p1 string, p2 int16) int32 {
    fmt.Println(p0, p0, p0)
    p0 -= (p0 << ((0) & 63))
    switch int(p2) {
    case -83:
        v6 := GenS1{F0: 674, F1: 2}
        _ = v6
    default:
        var v7 float32 = (float32((-2.25 - 0.1)) - (3.14159 / (3.14159 + float32(1.5))))
        _ = v7
    }
    fmt.Println(p1, p1, p2)
    v8 := GenS1{F0: 109, F1: 851}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2147483647
}

func main() {
    var v1 int32 = 0
    v2 := make(map[string]int)
    var v3 int = 0
    var v4 int = v3
    v1 = 2147483647
    switch "foo bar" {
    case "a":
        if true {
            var v9 int32 = genFn1(0, ("foo bar" + ("a" + "\n\t")), 32767)
            switch (v3 | (-(v4))) {
            case 0:
                fmt.Println(v3)
            default:
                var v10 string = (("hello" + ("a" + "a")) + "\n\t")
                _ = v10
            }
            _ = v9
        }
    case "hello":
        v11 := []int{}
        _ = v11
    default:
        var v12 int16 = -75
        _ = v12
    }
    var v13 string = "\n\t"
    v14 := []int{(v3 | (v3 >> ((211) & 63))), v4}
    var v15 float64 = 3.14159
    v15 += (1.5 * v15)
    v1 = (v1 + genFn1(0, v13, -1))
    var v16 int8 = 100
    if (v1 == -2147483648) {
        v13 = "foo bar"
        if (float64(1e10) > (((float64(1.5) - float64(1e10)) / 1.0) + 0.1)) {
            for fi17 := 0; fi17 < 11; fi17++ {
                v3 += v3
                _ = fi17
            }
        } else {
            var v18 int = v3
            _ = v18
        }
    }
    v19 := GenS1{F0: 384, F1: 0}
    fmt.Println(v13, v15, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
