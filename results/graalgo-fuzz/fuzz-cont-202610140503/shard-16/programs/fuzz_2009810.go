package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int32 }
func genFn0(p0 int16, p1 int8) string {
    var v5 int16 = (-(p0))
    v5 = (p0 / (((^(p0))) | 1))
    if false {
        var v6 uint32 = 9
        _ = v6
    }
    v7 := GenS1{F0: 183}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    _ = v8
    return ("hello" + ("a" + "hello"))
}

func main() {
    var v1 string = (("a" + ("" + (("hello" + "foo bar") + "foo bar"))) + "hello")
    var v2 float32 = 0.0
    v3 := make(map[int]bool)
    var v4 int16 = 32767
    v9 := GenS1{F0: 0}
    v10 := []int{127, 84, 127}
    v11 := GenS2{F0: 2147483648, F1: 42, F2: 952}
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS1{F0: -1}
        fmt.Println(fi12, v1)
        _ = fi12
        _ = v13
    }
    v1 = (v1 + v1)
    var v14 uint = 18446744073709551615
    v15 := GenS1{F0: -601}
    var v16 uint64 = 445
    var v17 bool = true
    var v18 int16 = v4
    fmt.Println(v18, v14, v1)
    switch "\n\t" {
    case "foo bar":
        for fi19 := 0; fi19 < 19; fi19++ {
            v2 -= ((1.5 * (float32(3.14159) + float32(float64((v2 - v2))))) * float32(-2.25))
            _ = fi19
        }
    case "\n\t":
        v1 = v1
    default:
        fmt.Println(v4)
    }
    fmt.Println(v14, v17, v17)
    v1 = (v1 + "foo bar")
    if len(v10) > 0 {
        v10[0] = -9223372036854775808
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
