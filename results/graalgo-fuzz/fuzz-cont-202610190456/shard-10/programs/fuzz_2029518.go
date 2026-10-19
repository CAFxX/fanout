package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint16; F1 string }
func genFn0(p0 uint16, p1 uint8) uint16 {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int32 = -859
        fmt.Println(p0, p1)
        v7 = v7
        _ = fi6
        _ = v7
    }
    if (0 == p0) {
        var v8 int8 = -1
        var v9 uint = 65535
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 7
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    var v3 uint = 148
    var v4 int64 = 1
    var v5 float64 = 1e10
    switch 127 {
    case -448:
        switch 2147483647 {
        case -724:
            if ((true && (13 <= 780)) || true) {
                switch "foo bar" {
                case "":
                    v10 := GenS1{F0: 112, F1: (v4 | (v4 % ((561) | 1)))}
                    _ = v10
                case "hello":
                    fmt.Println(v4, v3, v5)
                case "foo bar":
                    var v11 float64 = 1e10
                    _ = v11
                default:
                    var v12 bool = true
                    _ = v12
                }
                if (2 <= 883) {
                    v13 := []int{-1, 32768}
                    _ = v13
                } else {
                    v14 := []int{42, -948, -718, 1}
                    if (!(2 == 42)) {
                        v3 = (v3 & (v3 % ((v3) | 1)))
                        var v15 uint8 = 128
                        _ = v15
                    }
                    _ = len(v2)
                    _ = v14
                }
                switch 42 {
                case 573:
                    v16 := GenS2{F0: 65535, F1: (("\n\t" + "\n\t") + (("\n\t" + "foo bar") + "\n\t"))}
                    _ = v16
                default:
                    switch "a" {
                    case "foo bar":
                        _ = len(v1)
                    case "":
                        var v17 int32 = -1
                        _ = v17
                    case "hello":
                        v3 = 1
                    default:
                        v3 = v3
                    }
                }
            }
        case -2147483648:
            v5 = (v5 * 0.1)
        }
    }
    switch ("" + "") {
    case "a":
        var v18 int = 707
        _ = v18
    }
    var v19 float64 = v5
    v20 := []int{1, -384, 2147483647}
    fmt.Println(v19, v5)
    fmt.Println(v19, v4)
    var v21 uint32 = 2
    if false {
        for fi22 := 0; fi22 < 19; fi22++ {
            v21 = uint32(v3)
            _ = fi22
        }
        v20 = append(v20, -1)
    }
    var v23 float64 = v19
    v21 = v21
    var v24 string = ""
    v23 += ((((v23 - 0.1) * 3.14159) / 1e10) / (0.1 / v23))
    v25 := GenS2{F0: 255, F1: (("a" + v24) + "foo bar")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
