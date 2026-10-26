package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
func genFn0(p0 int, p1 int8) int {
    v6 := GenS1{F0: 0, F1: ("" + "foo bar"), F2: -128}
    v7 := []int{p0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    var v3 float64 = (3.14159 / ((1e10 - ((-2.25 + 3.14159) + 1.5)) / 0.1))
    v4 := []uint32{2147483648, 65535, 488, 4294967295}
    v5 := []int{127, 42, -128, 2147483647}
    v3 += v3
    fmt.Println(v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v8 int32 = -833
    v9 := GenS1{F0: 42, F1: "", F2: 100}
    v10 := GenS1{F0: 465, F1: (("foo bar" + ("hello" + ("\n\t" + "a"))) + "hello"), F2: 0}
    for fi11 := 0; fi11 < 20; fi11++ {
        _ = len(v1)
        if true {
            var v12 int16 = -1
            _ = v12
        } else {
            v3 = float64(float32(v3))
        }
        _ = fi11
    }
    var v13 uint = 18446744073709551615
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
