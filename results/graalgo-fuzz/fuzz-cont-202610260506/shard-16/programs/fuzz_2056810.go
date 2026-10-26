package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 string }
func genFn0(p0 int32, p1 int, p2 int8) uint8 {
    v7 := []int{42, p1}
    p2 -= p2
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 float32 = float32(3.14159)
            _ = fi9
            _ = v10
        }
        p1 += fi8
        var v11 uint32 = 2
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 11
}
func genFn1(p0 int32, p1 uint) bool {
    var v12 int32 = 0
    fmt.Println(v12, p1)
    p0 = p0
    _ = p0
    _ = p1
    _ = v12
    return false
}

func main() {
    var v1 int = 0
    var v2 uint64 = 1
    var v3 float64 = (float64((1.5 / (0.1 - (3.14159 + 1.5)))) * 3.14159)
    v4 := make(map[uint32]string)
    var v5 int64 = 42
    var v6 bool = (true || true)
    v13 := GenS1{F0: "", F1: 128, F2: 127}
    if false {
        v6 = true
    }
    switch "a" {
    case "a":
        var v14 uint8 = 42
        _ = v14
    case "hello":
        fmt.Println(v3, v2, v3)
    case "foo bar":
        v6 = false
    }
    v15 := GenS2{F0: (v1 * v1), F1: 18446744073709551615, F2: ""}
    switch ("hello" + "foo bar") {
    case "":
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := []int{(v1 ^ v1), (v1 / ((fi16) | 1))}
            v6 = v6
            _ = fi16
            _ = v17
        }
    case "foo bar":
        var v18 uint16 = 255
        _ = v18
    case "a":
        v2 = v2
    }
    fmt.Println(v5, v1, v6)
    var v19 int32 = 42
    v1 = (v1 % ((-1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
