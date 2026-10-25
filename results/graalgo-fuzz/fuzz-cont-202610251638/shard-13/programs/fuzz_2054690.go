package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 bool) uint8 {
    var v5 bool = (-2147483648 >= 1)
    v6 := []int{2147483647, 1}
    switch "foo bar" {
    case "foo bar":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        var v7 float64 = 0.0
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 0
    v3 := make(map[int]string)
    v4 := []uint32{}
    v3[42] = ("" + (("" + ("foo bar" + "hello")) + "\n\t"))
    v2 -= (v2 % (((v2 << ((255) & 63))) | 1))
    if (0 < 255) {
        fmt.Println(v2)
    } else {
        if (true && true) {
            if false {
                switch "foo bar" {
                case "a":
                    v4 = append(v4, 2)
                case "hello":
                    switch "" {
                    case "a":
                        v2 = v2
                    default:
                        if len(v4) > 0 {
                            v4[0] = 2
                        }
                    }
                case "foo bar":
                    var v8 uint8 = 2
                    _ = v8
                }
                fmt.Println(v2)
            }
            v9 := GenS1{F0: 533}
            fmt.Println(v2)
            _ = v9
        }
        _ = len(v4)
        v1[-113] = "hello"
    }
    v10 := GenS1{F0: 65535}
    v11 := GenS2{F0: 2}
    var v12 string = ""
    switch "foo bar" {
    case "\n\t":
        v2 += (v2 - v2)
    case "foo bar":
        for fi13 := 0; fi13 < 18; fi13++ {
            delete(v1, 1)
            _ = fi13
        }
    default:
        v14 := GenS3{F0: 1}
        _ = v14
    }
    var v15 int = -9223372036854775808
    var v16 uint = 18446744073709551615
    v16 += (v16 / ((576) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
