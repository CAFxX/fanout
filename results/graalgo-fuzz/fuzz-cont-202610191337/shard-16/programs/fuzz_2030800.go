package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint8 = 42
    var v2 uint = 720
    var v3 uint16 = 2
    v4 := GenS1{F0: -83}
    v5 := GenS2{F0: v1}
    v6 := GenS2{F0: 252}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 float64 = 0.0
        v9 := []int{fi7, (fi7 % ((fi7) | 1)), (-(fi7)), (^(fi7))}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch ((("a" + (("foo bar" + "") + "hello")) + "foo bar") + "a") {
    case "\n\t":
        v1 = (^(v1))
    case "hello":
        v10 := GenS1{F0: 127}
        _ = v10
    }
    if true {
        v1 = v1
        v11 := GenS3{F0: -9223372036854775808}
        fmt.Println(v2, v3, v2)
        _ = v11
    }
    if (1000 <= 1) {
        v1 = (v1 * v1)
        var v12 bool = (!true)
        switch 2147483647 {
        case 0:
            fmt.Println(v2, v3)
        case 1:
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := []int{}
                var v15 int64 = -9223372036854775808
                switch "a" {
                case "a":
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(fi13, v15, v15)
                    case "":
                        fmt.Println(v15)
                    default:
                        v15 = v15
                    }
                case "hello":
                    var v16 float64 = (0.1 - 3.14159)
                    _ = v16
                }
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            v2 += (^(v2))
        }
        _ = v12
    }
    fmt.Println(v2, v2, v1)
    if (!((-2.25 - 1.5) > float64(0.0))) {
        v17 := []int{-853, 32767, 32767}
        if len(v17) > 0 {
            v17[0] = 874
        }
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
