package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 bool, p1 uint64) uint32 {
    v4 := GenS1{F0: 42}
    switch "\n\t" {
    case "hello":
        var v5 float64 = 0.0
        _ = v5
    case "":
        var v6 int8 = -128
        _ = v6
    default:
        p1 = p1
    }
    v7 := []int{42, 375, 2147483647}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 842
}

func main() {
    v1 := []string{}
    var v2 int32 = -210
    var v3 int32 = v2
    var v8 float64 = 1.5
    v3 -= v3
    if (4294967295 == 1) {
        var v9 uint = 42
        v2 -= v2
        _ = v9
    } else {
        fmt.Println(v2, v8)
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 string = "a"
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            fmt.Println(fi12, v8, fi12)
            v2 += -81
            var v13 float32 = (0.1 / -2.25)
            _ = fi12
            _ = v13
        }
    }
    if (!(!false)) {
        var v14 int16 = 0
        fmt.Println(v3, v8)
        _ = v14
    }
    var v15 uint32 = 567
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 uint16 = 42
    _ = len(v1)
    v17 := GenS1{F0: 127}
    var v18 string = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
}
