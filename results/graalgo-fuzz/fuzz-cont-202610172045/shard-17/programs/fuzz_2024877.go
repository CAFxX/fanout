package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 float64 = (-2.25 - 3.14159)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    v4 := GenS1{F0: 2147483647}
    v5 := GenS1{F0: 2147483647}
    if true {
        v6 := []int{-1, 9223372036854775807, -9223372036854775808, -60}
        var v7 int16 = -1
        _ = v6
        _ = v7
    }
    delete(v3, -2147483648)
    v8 := []int{1}
    fmt.Println(v1)
    if (18446744073709551615 >= 42) {
        var v9 int16 = 690
        v10 := GenS2{F0: -1}
        _ = v10
        _ = v9
    } else {
        var v11 int16 = 42
        v12 := GenS2{F0: 42}
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        switch fi13 {
        case 128:
            var v14 uint = 1
            _ = v14
        case 4294967295:
            var v15 int64 = 0
            _ = v15
        case -2147483648:
            switch "" {
            case "foo bar":
                v1 = v1
            case "\n\t":
                var v16 int8 = 56
                _ = v16
            case "hello":
                for fi17 := 0; fi17 < 10; fi17++ {
                    v1 += 1.5
                    _ = fi17
                }
            }
        default:
            fmt.Println(fi13, v1)
        }
        if false {
            switch fi13 {
            case -374:
                v8 = append(v8, fi13)
            case 9223372036854775807:
                v1 = v1
            default:
                v1 = v1
            }
        } else {
            v1 -= float64((1e10 / 1.5))
        }
        var v18 float64 = v1
        _ = fi13
        _ = v18
    }
    var v19 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
