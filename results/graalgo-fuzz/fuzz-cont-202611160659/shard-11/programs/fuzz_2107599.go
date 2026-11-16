package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int64) uint16 {
    if (65535 > 829) {
        v4 := GenS1{F0: -2147483648}
        var v5 int = 9223372036854775807
        var v6 int64 = -1
        _ = v4
        _ = v5
        _ = v6
    }
    v7 := []int{-2147483648, 1, 81, -194}
    if false {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int8 = 1
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = v7
    return 716
}

func main() {
    var v1 bool = (!true)
    v2 := make(map[uint32]bool)
    v3 := []string{}
    v1 = true
    v1 = (!v1)
    switch -1 {
    case 503:
        var v10 int64 = 9223372036854775807
        _ = v10
    case 330:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v11 string = "\n\t"
    v12 := GenS1{F0: -992}
    v11 = v11
    var v13 uint64 = 1
    var v14 uint32 = 58
    switch v11 {
    case "foo bar":
        _ = len(v2)
    default:
        v15 := GenS1{F0: 369}
        _ = v15
    }
    var v16 float64 = 0.1
    var v17 uint64 = v13
    v18 := v2[(v14 & (v14 / (((v14 ^ 65535)) | 1)))]
    var v19 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
