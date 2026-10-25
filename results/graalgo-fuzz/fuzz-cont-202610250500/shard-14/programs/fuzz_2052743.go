package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int }
func genFn0(p0 uint8, p1 uint8, p2 string) bool {
    var v4 int = 32768
    p2 = "\n\t"
    v5 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 int8 = 58
    var v2 int8 = v1
    v3 := make(map[string]int)
    v6 := GenS1{F0: 2147483647}
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := v3[(("foo bar" + "a") + ("a" + ("\n\t" + (("foo bar" + "foo bar") + ""))))]
        var v9 int32 = -2147483648
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: -2147483648}
    v1 = v1
    v11 := GenS2{F0: 0, F1: 1}
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            v1 -= (-(v2))
            _ = fi13
        }
        fmt.Println(fi12, fi12)
        _ = fi12
    }
    v14 := []int{}
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(fi15, v2, fi15)
        fmt.Println(v1)
        _ = fi15
    }
    v16 := GenS1{F0: 42}
    switch 5 {
    case -772:
        var v17 uint = 459
        _ = v17
    default:
        var v18 float32 = -2.25
        _ = v18
    }
    var v19 string = ("\n\t" + "a")
    v1 = (v1 ^ (^(v2)))
    fmt.Println(v2, v1, v19)
    var v20 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
}
