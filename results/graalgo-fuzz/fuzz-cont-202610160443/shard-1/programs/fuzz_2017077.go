package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint; F2 int16 }
type GenS4 struct { F0 int16; F1 string }
func genFn0(p0 uint64, p1 uint) uint8 {
    if true {
        fmt.Println(p0, p0)
        v4 := []int{128}
        _ = v4
    }
    v5 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = v5
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    var v3 uint64 = 0
    v2 = (-128 != -56)
    switch "hello" {
    case "foo bar":
        var v6 float32 = 0.1
        _ = v6
    }
    v7 := []int{}
    var v8 string = (("foo bar" + "hello") + "\n\t")
    var v9 uint64 = (^(v3))
    switch -2147483648 {
    case -1:
        v2 = false
    case -2147483648:
        if len(v7) > 0 {
            v7[0] = -39
        }
    default:
        v3 = 641
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        if v2 {
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 float64 = -2.25
                _ = fi11
                _ = v12
            }
            v7 = append(v7, (fi10 >> ((2) & 63)))
            v13 := GenS2{F0: -1}
            _ = v13
        }
        _ = fi10
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        switch "" {
        case "\n\t":
            fmt.Println(v8, v2)
        case "hello":
            var v15 int = (fi14 + fi14)
            _ = v15
        }
        _ = fi14
    }
    v16 := GenS3{F0: -473, F1: 2, F2: 42}
    v2 = (1.5 < 1e10)
    if (112 >= 2) {
        v17 := GenS4{F0: 0, F1: v8}
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
