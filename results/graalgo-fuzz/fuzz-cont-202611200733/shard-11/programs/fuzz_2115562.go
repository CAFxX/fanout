package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int64; F1 int; F2 uint64 }
func genFn0(p0 uint64, p1 string, p2 int64) uint16 {
    var v4 bool = (false || true)
    fmt.Println(p0, p1)
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v5 := []int{705, 45, 201}
        _ = v5
    case "a":
        v6 := GenS1{F0: 122, F1: 127}
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int8 = 1
            v9 := []int{fi7, fi7, fi7, -9223372036854775808}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}
func genFn1(p0 int64) int8 {
    if (false && ((!(65535 >= 2)) || (42 >= -200))) {
        p0 = 42
    } else {
        var v10 float64 = float64(0.0)
        _ = v10
    }
    if (p0 == 42) {
        switch "foo bar" {
        case "\n\t":
            p0 -= p0
        case "foo bar":
            var v11 uint32 = 744
            _ = v11
        default:
            fmt.Println(p0)
        }
    }
    _ = p0
    return -107
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (false || true)
    var v3 uint8 = 42
    fmt.Println(v3, v2)
    fmt.Println(v2)
    var v12 int = 42
    v13 := GenS1{F0: -9223372036854775808, F1: (v12 % (((-(v12))) | 1))}
    var v14 uint8 = 127
    var v15 float32 = (float32(((3.14159 / 1.0) + float64(((float32((float64(1e10) * (float64((3.14159 - 0.0)) / 3.14159))) / 0.1) / 0.1)))) + 1e10)
    v16 := []int{v12, v12, 0}
    for fi17 := 0; fi17 < 17; fi17++ {
        if v2 {
            fmt.Println(fi17, v3)
            switch "a" {
            case "a":
                v12 = (fi17 / ((fi17) | 1))
            case "\n\t":
                v18 := GenS2{F0: 0, F1: 1, F2: 1}
                _ = v18
            }
        }
        v12 -= v12
        if (!(((v15 + 1e10) + v15) != 0.1)) {
            v2 = (!v2)
        }
        _ = fi17
    }
    var v19 int16 = 42
    if true {
        fmt.Println(v12, v15)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
