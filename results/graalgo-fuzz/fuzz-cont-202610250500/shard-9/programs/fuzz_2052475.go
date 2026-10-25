package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint64; F1 string; F2 int8 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint8, p1 bool, p2 int64) bool {
    v6 := GenS1{F0: 836}
    switch 767 {
    case 32767:
        if p1 {
            if false {
                v7 := GenS2{F0: 565, F1: 1}
                var v8 bool = ((p2 >> ((2) & 63)) >= (p2 | p2))
                p2 -= p2
                _ = v7
                _ = v8
            }
            v9 := GenS1{F0: 1}
            _ = v9
        } else {
            fmt.Println(p2, p0, p0)
        }
    case 42:
        p2 = 795
    default:
        p2 = p2
    }
    var v10 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return p1
}

func main() {
    var v1 int32 = 42
    var v2 uint32 = 1
    v3 := make(map[string]uint32)
    var v4 bool = true
    var v5 int64 = -845
    var v11 int = 4294967295
    v11 = (v11 / (((v11 << ((2) & 63))) | 1))
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS3{F0: uint64(v2), F1: "foo bar", F2: 100}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 int32 = v1
        v16 := GenS2{F0: 2, F1: -38}
        var v17 float64 = 0.0
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    fmt.Println(v4, v4, v11)
    fmt.Println(v1, v11, v11)
    if ((-2.25 + 1e10) != float32(3.14159)) {
        v18 := v3[""]
        _ = v18
    } else {
        v19 := []int{}
        v5 -= 1
        _ = v19
    }
    v20 := []int{v11, (-(v11)), 4294967295, 9223372036854775807}
    for fi21 := 0; fi21 < 5; fi21++ {
        v22 := GenS4{F0: 42}
        for fi23 := 0; fi23 < 3; fi23++ {
            fmt.Println(v1, fi21)
            var v24 float32 = 0.0
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    v25 := GenS3{F0: 0, F1: ("" + "foo bar"), F2: 100}
    for fi26 := 0; fi26 < 20; fi26++ {
        v11 += fi26
        switch ((("" + "\n\t") + ("foo bar" + "foo bar")) + "hello") {
        case "hello":
            v11 += (v11 | -103)
        default:
            fmt.Println(v1, v1)
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
