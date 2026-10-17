package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 bool }

func main() {
    var v1 int64 = 797
    var v2 int8 = -128
    v3 := []int{}
    var v4 int64 = (-(v1))
    v5 := make(map[int]bool)
    var v6 string = ("\n\t" + "hello")
    var v7 float64 = 0.0
    v8 := GenS1{F0: 0, F1: false}
    v4 = (v4 - -9223372036854775808)
    v9 := GenS2{F0: 0}
    if (!(-1 == int16(v2))) {
        v1 -= v4
    }
    switch v6 {
    case "":
        v1 -= -1
    case "foo bar":
        var v10 float64 = -2.25
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 7; fi11++ {
            if false {
                if (((v6 + v6) + (v6 + "a")) >= ((v6 + v6) + v6)) {
                    if (100 == v2) {
                        v7 -= (1.5 / 1e10)
                    } else {
                        v12 := GenS2{F0: 5}
                        _ = v12
                    }
                    var v13 uint = 1
                    v6 = (v6 + v6)
                    _ = v13
                }
                v3 = append(v3, (-(fi11)))
                v7 += 0.1
            } else {
                var v14 int32 = 553
                for fi15 := 0; fi15 < 12; fi15++ {
                    fmt.Println(v7)
                    _ = len(v3)
                    _ = fi15
                }
                _ = v14
            }
            v1 = (v4 & 387)
            v3 = append(v3, -2147483648)
            _ = fi11
        }
    default:
        if false {
            var v16 float32 = 3.14159
            fmt.Println(v4, v1)
            _ = v16
        }
    }
    var v17 uint8 = 127
    var v18 string = "a"
    fmt.Println(v4, v4)
    v4 = v1
    v19 := GenS3{F0: 2, F1: true}
    fmt.Println(v7, v2, v4)
    v5[-2147483648] = (false || false)
    var v20 uint = 18446744073709551615
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v21 uint64 = 18446744073709551615
    switch 111 {
    case 483:
        v18 = ((v18 + "a") + v6)
    case 0:
        var v22 uint64 = v21
        _ = v22
    default:
        v23 := GenS3{F0: 259, F1: ((v1 << ((v20) & 63)) > int64(v2))}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
