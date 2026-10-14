package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint8 }

func main() {
    var v1 uint64 = 2
    var v2 uint64 = (v1 / (((v1 / (((v1 & (v1 << ((2) & 63)))) | 1))) | 1))
    var v3 uint32 = 2
    var v4 uint64 = v2
    switch "hello" {
    case "":
        if (!(!false)) {
            v5 := []int{2147483647, -93, 620}
            fmt.Println(v1, v4)
            _ = v5
        }
    default:
        v6 := []int{}
        _ = v6
    }
    fmt.Println(v2, v4, v2)
    switch "hello" {
    case "":
        if (42 < 42) {
            var v7 uint64 = v4
            v4 -= v2
            if false {
                switch ("hello" + "foo bar") {
                case "\n\t":
                    v8 := []int{4294967295, 32768, 4294967295, 50}
                    _ = v8
                case "hello":
                    fmt.Println(v3)
                case "a":
                    v9 := []int{127}
                    _ = v9
                }
            }
            _ = v7
        }
    case "\n\t":
        v4 += (v2 | v1)
    }
    v10 := GenS1{F0: -25, F1: 4294967295, F2: 59}
    v11 := GenS1{F0: 34, F1: 255, F2: 42}
    v12 := []int{}
    v2 = 18446744073709551615
    if true {
        for fi13 := 0; fi13 < 10; fi13++ {
            v4 -= 0
            if (!(0.0 <= 0.1)) {
                if true {
                    var v14 bool = (0.1 != 3.14159)
                    v14 = v14
                    _ = v14
                } else {
                    fmt.Println(v4, v1)
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
                var v15 int16 = 95
                v12 = append(v12, (fi13 << ((4294967295) & 63)))
                _ = v15
            }
            var v16 int = (-(fi13))
            _ = fi13
            _ = v16
        }
        fmt.Println(v1)
        if false {
            if len(v12) > 0 {
                v12[0] = 680
            }
            var v17 string = "\n\t"
            _ = v17
        } else {
            v2 = v4
        }
    } else {
        v2 += 42
    }
    fmt.Println(v3)
    switch 4294967295 {
    case 4294967295:
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
