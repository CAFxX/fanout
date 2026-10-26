package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint, p2 bool) uint16 {
    p2 = (2 == 65535)
    v4 := GenS1{F0: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 255
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (float32(0.0) * (float32(0.1) - 3.14159))
    var v3 int16 = 887
    v5 := []int{}
    if (4294967295 < 430) {
        if false {
            v2 = -2.25
        } else {
            fmt.Println(v2)
        }
        var v6 uint32 = 823
        _ = v6
    } else {
        v3 = 84
        switch "foo bar" {
        case "hello":
            if len(v5) > 0 {
                v5[0] = 32767
            }
        }
        if false {
            var v7 float64 = 0.0
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v3, v7, v2)
            _ = v7
        } else {
            var v8 uint64 = 1
            _ = v8
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        fmt.Println(v2)
        var v9 int64 = 1
        switch "foo bar" {
        case "":
            var v10 uint = 0
            _ = v10
        }
        _ = v9
    } else {
        v2 = v2
    }
    var v11 int = -531
    var v12 uint32 = 2147483648
    v3 = v3
    v2 = v2
    fmt.Println(v12, v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
}
