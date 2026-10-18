package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS5 struct { F0 int32; F1 uint8 }
func genFn0(p0 int8) bool {
    v6 := GenS1{F0: 1}
    var v7 uint32 = 913
    v8 := GenS2{F0: 18446744073709551615}
    v9 := []int{338}
    v7 = (v7 + v7)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1e10
    v3 := []string{"hello", "\n\t", "hello", "\n\t"}
    v4 := []uint32{151, 2147483648, 1, 380}
    v5 := []string{"foo bar", "\n\t", "\n\t"}
    v2 -= v2
    v2 += (v2 * v2)
    if (true || (false && (!("hello" < (("a" + ("hello" + ("hello" + "\n\t"))) + ("hello" + "")))))) {
        var v10 float64 = -2.25
        v2 += 0.0
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(fi11)
        _ = fi11
    }
    if false {
        for fi12 := 0; fi12 < 13; fi12++ {
            v2 = v2
            switch "foo bar" {
            case "foo bar":
                v13 := GenS3{F0: 1}
                _ = v13
            case "":
                v2 -= 3.14159
            }
            _ = fi12
        }
    } else {
        var v14 int64 = -9223372036854775808
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS4{F0: 0, F1: -322, F2: 42}
        var v17 uint32 = 1
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v2)
    fmt.Println(v2)
    delete(v1, 42)
    fmt.Println(v2)
    v2 = v2
    v2 = ((1e10 / (v2 * (float32((1e10 / 3.14159)) - (1e10 * v2)))) - v2)
    var v18 int = 0
    v2 = 1.5
    v19 := GenS1{F0: 18446744073709551615}
    var v20 uint64 = 4294967295
    if (0.1 == 1.5) {
        v21 := GenS5{F0: 42, F1: 127}
        var v22 int8 = 82
        _ = v21
        _ = v22
    }
    v2 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
