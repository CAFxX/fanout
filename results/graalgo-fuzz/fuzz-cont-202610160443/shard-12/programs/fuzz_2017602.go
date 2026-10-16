package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0(p0 int32, p1 uint8, p2 uint16) uint8 {
    var v6 uint = 0
    v7 := GenS1{F0: ((p2 / ((p2) | 1)) == 869), F1: p0}
    var v8 int = 32768
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (p1 - p1)
}

func main() {
    var v1 int64 = -441
    v2 := []int{42, 42}
    var v3 uint32 = 1
    var v4 int64 = 263
    var v5 int8 = 127
    switch "foo bar" {
    case "a":
        v1 = -9223372036854775808
    case "\n\t":
        if (42 > -2147483648) {
            fmt.Println(v5, v3, v3)
        }
    case "":
        v3 += (v3 % ((v3) | 1))
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 bool = true
            _ = fi9
            _ = v10
        }
    }
    if len(v2) > 0 {
        v2[0] = -639
    }
    v4 = 1
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(v1, v4)
        if len(v2) > 0 {
            v2[0] = fi11
        }
        v1 = (^(v1))
        _ = fi11
    }
    v2 = append(v2, 32767)
    fmt.Println(v4, v4)
    v3 += 4294967295
    var v12 string = "a"
    var v13 uint8 = 1
    var v14 uint = 1
    v13 = uint8(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
