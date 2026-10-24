package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 uint64, p2 int16) bool {
    v6 := GenS1{F0: 1, F1: -1}
    v7 := []int{42, -130, 753}
    var v8 uint32 = 2147483647
    if true {
        p0 = p0
    }
    v9 := GenS1{F0: 4294967295, F1: -37}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (true || true)
}

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]bool)
    var v3 uint64 = 1
    v4 := []string{"hello", "foo bar"}
    v5 := []uint32{64, 256, 4294967295}
    v3 = v3
    v10 := GenS1{F0: -807, F1: -545}
    v1 = (v1 & (-(v1)))
    v2[42] = ((float64(0.0) == ((1e10 - ((3.14159 / float64(3.14159)) * 0.1)) - 1.5)) && (!((42 != 2147483647) && false)))
    for fi11 := 0; fi11 < 17; fi11++ {
        v1 = (^(v1))
        v3 += 428
        _ = fi11
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v12 := GenS1{F0: -825, F1: 2147483647}
    var v13 float64 = 1e10
    if false {
        v13 += 0.1
    }
    fmt.Println(v3, v1, v1)
    v14 := GenS2{F0: v3}
    switch "a" {
    case "":
        fmt.Println(v3)
    case "hello":
        fmt.Println(v3, v1, v3)
    case "a":
        v13 = float64(-2.25)
    }
    fmt.Println(v3, v13, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
