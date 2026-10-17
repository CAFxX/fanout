package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int64 }
func genFn0(p0 string, p1 int8, p2 int) bool {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := GenS1{F0: 784, F1: 42, F2: 0}
        switch p0 {
        case "":
            p1 += p1
        case "a":
            var v6 int16 = -659
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    p2 = (-(p2))
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int = p2
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "\n\t":
        v9 := []int{p2}
        _ = v9
    case "foo bar":
        var v10 int32 = 42
        _ = v10
    case "hello":
        var v11 float64 = -2.25
        _ = v11
    default:
        var v12 uint8 = 2
        _ = v12
    }
    v13 := []int{(p2 | -204), (^(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (int64(p1) >= -9223372036854775808)
}

func main() {
    v1 := []string{"\n\t"}
    var v2 uint8 = 194
    var v3 int = -2147483648
    var v14 string = "a"
    var v15 int32 = 757
    v3 = (v3 - 612)
    var v16 int = (^(v3))
    var v17 uint16 = 904
    v18 := GenS1{F0: 1, F1: int64(v2), F2: int64(v17)}
    fmt.Println(v15, v17, v14)
    if false {
        v2 += v2
        var v19 uint32 = 4294967295
        _ = v19
    } else {
        v16 += v16
    }
    v16 = 32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
