package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 uint64, p1 string, p2 int8) int32 {
    v5 := []int{}
    v6 := GenS1{F0: 2147483647, F1: 32767, F2: 32767}
    if false {
        v7 := GenS2{F0: 42, F1: 213}
        _ = v7
    } else {
        fmt.Println(p2, p2)
        for fi8 := 0; fi8 < 12; fi8++ {
            switch p1 {
            case "foo bar":
                if true {
                    var v9 int64 = -1
                    var v10 bool = true
                    p1 = (p1 + p1)
                    _ = v10
                    _ = v9
                }
            case "\n\t":
                p1 = ("" + (p1 + p1))
            }
            p2 = 100
            p2 += 1
            _ = fi8
        }
        for fi11 := 0; fi11 < 7; fi11++ {
            p0 = (^(p0))
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -285
}

func main() {
    var v1 uint8 = 128
    var v2 int = 32767
    var v3 int8 = 127
    var v4 int32 = 2147483647
    var v12 float64 = 1.5
    if (953 >= 866) {
        v1 = v1
        v4 = (v4 - (v4 * v4))
        fmt.Println(v4, v12, v2)
    }
    fmt.Println(v2, v3)
    fmt.Println(v2, v1)
    v13 := GenS3{F0: 128, F1: v1}
    var v14 int = (v2 - v2)
    var v15 uint8 = (v1 + 1)
    switch v2 {
    case 4294967295:
        fmt.Println(v15, v4)
    default:
        if ("foo bar" >= ("\n\t" + "\n\t")) {
            var v16 int16 = 32767
            var v17 string = "hello"
            _ = v16
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
