package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 string }

func main() {
    var v1 int16 = -638
    v2 := []int{}
    var v3 string = "foo bar"
    var v4 float32 = 0.0
    var v5 string = v3
    v4 += (3.14159 * float32(3.14159))
    if true {
        v2 = append(v2, 42)
        var v6 uint64 = 0
        fmt.Println(v6)
        _ = v6
    } else {
        v1 = v1
    }
    var v7 uint = 616
    _ = len(v2)
    for fi8 := 0; fi8 < 2; fi8++ {
        v5 = ("a" + "foo bar")
        _ = fi8
    }
    if true {
        var v9 int8 = 41
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int64 = 0
            v12 := GenS1{F0: (42 < 4294967295), F1: "hello"}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = len(v2)
        _ = v9
    }
    var v13 float64 = (((float64(1.5) + (1.5 + float64(v4))) / 0.1) - 1.5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v14 := GenS2{F0: v7}
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            switch v5 {
            case "hello":
                fmt.Println(v1, fi15, v3)
                v4 = v4
            case "":
                var v17 float32 = (v4 / v4)
                _ = v17
            case "a":
                v3 = v5
            default:
                var v18 string = (v3 + "\n\t")
                _ = v18
            }
            var v19 uint32 = 65535
            _ = fi16
            _ = v19
        }
        _ = fi15
    }
    switch 826 {
    case -116:
        for fi20 := 0; fi20 < 5; fi20++ {
            if (895 == 2147483648) {
                var v21 uint32 = 2147483647
                _ = v21
            }
            v2 = append(v2, (fi20 % (((fi20 | (fi20 % (((fi20 / ((4294967295) | 1))) | 1)))) | 1)))
            var v22 uint = (v7 ^ 18446744073709551615)
            _ = fi20
            _ = v22
        }
    default:
        v4 += v4
    }
    v23 := GenS3{F0: 4294967295, F1: (v5 + v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
