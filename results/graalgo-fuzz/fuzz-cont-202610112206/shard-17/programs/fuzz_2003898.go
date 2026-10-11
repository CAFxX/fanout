package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 bool; F1 uint64; F2 uint8 }
func genFn0(p0 uint8) uint32 {
    p0 += 44
    v6 := GenS1{F0: 572, F1: ((("a" + (("hello" + "a") + "a")) + "foo bar") + "foo bar")}
    var v7 int64 = 551
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint8 = 1
    var v2 int64 = 163
    var v3 uint32 = 485
    v4 := make(map[int]bool)
    var v5 uint = 18446744073709551615
    var v8 string = ("a" + "hello")
    if (0 <= (v3 << (((v5 << ((v5) & 63))) & 63))) {
        for fi9 := 0; fi9 < 14; fi9++ {
            fmt.Println(v3)
            var v10 uint = (^(v5))
            var v11 float64 = (-2.25 - (3.14159 - 0.0))
            _ = fi9
            _ = v10
            _ = v11
        }
        v3 = genFn0(uint8(v3))
        for fi12 := 0; fi12 < 2; fi12++ {
            v8 = v8
            v3 = v3
            v8 = (((("\n\t" + v8) + (v8 + v8)) + v8) + v8)
            _ = fi12
        }
    } else {
        v1 = v1
    }
    fmt.Println(v3, v1, v2)
    v2 = v2
    v13 := GenS2{F0: true, F1: 788, F2: v1}
    var v14 bool = true
    v15 := []int{}
    fmt.Println(v8, v5)
    v2 += v2
    var v16 string = v8
    switch -767 {
    case -658:
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    case 926:
        v5 = (v5 - v5)
    default:
        v17 := v4[9223372036854775807]
        _ = v17
    }
    v15 = append(v15, 500)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
