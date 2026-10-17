package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int8 }
func genFn0(p0 string) int8 {
    p0 = p0
    p0 = p0
    _ = p0
    return -1
}
func genFn1(p0 int32, p1 string, p2 uint) string {
    if false {
        var v6 uint16 = 32767
        _ = v6
    }
    var v7 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ((p1 + p1) + "a")
}

func main() {
    v1 := []int{-863, -1, 323, 4294967295}
    var v2 uint32 = 2
    var v3 uint16 = 65535
    v4 := make(map[string]string)
    v5 := []string{"hello", ("" + "a"), ("\n\t" + "hello"), (("hello" + ("\n\t" + "foo bar")) + "foo bar")}
    switch -712 {
    case 42:
        for fi8 := 0; fi8 < 1; fi8++ {
            fmt.Println(fi8, v2, v2)
            var v9 float64 = 0.0
            _ = fi8
            _ = v9
        }
    case 32768:
        v2 -= (^(v2))
    default:
        _ = len(v1)
    }
    _ = len(v4)
    var v10 int16 = -146
    v11 := GenS1{F0: 42}
    fmt.Println(v10, v3, v3)
    var v12 float32 = -2.25
    var v13 float64 = 0.0
    for fi14 := 0; fi14 < 9; fi14++ {
        _ = len(v4)
        _ = fi14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 int8 = 1
        if true {
            v12 += 1e10
        }
        var v17 bool = (false || false)
        _ = fi15
        _ = v16
        _ = v17
    }
    if (-1 != (v10 | v10)) {
        v13 += float64(float32(((v13 * 0.1) * v13)))
    } else {
        delete(v4, "\n\t")
        v18 := GenS2{F0: 0, F1: 100}
        _ = v18
    }
    v19 := GenS1{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
