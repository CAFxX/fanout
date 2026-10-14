package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
func genFn0(p0 uint64) int {
    v7 := []int{4294967295, 32768}
    _ = len(v7)
    _ = p0
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint = 105
    v2 := []int{-107, -549, 128}
    v3 := make(map[uint32]uint32)
    var v4 int32 = 0
    var v5 int32 = 0
    var v6 float64 = float64(-2.25)
    for fi8 := 0; fi8 < 20; fi8++ {
        switch genFn0(198) {
        case 32767:
            for fi9 := 0; fi9 < 7; fi9++ {
                v1 += (v1 - 982)
                v1 = (v1 * v1)
                _ = fi9
            }
        case -790:
            var v10 uint32 = 810
            _ = v10
        case 128:
            var v11 bool = (v4 < v4)
            _ = v11
        default:
            var v12 int8 = 66
            _ = v12
        }
        switch "\n\t" {
        case "\n\t":
            if true {
                if ("a" != (("a" + ("a" + "a")) + "\n\t")) {
                    v5 -= v4
                    var v13 int16 = 1
                    _ = v13
                }
            }
        case "a":
            if (float32(1.5) != 1.5) {
                if (!false) {
                    v1 += (v1 ^ v1)
                    var v14 uint = (^(v1))
                    var v15 float64 = v6
                    _ = v14
                    _ = v15
                } else {
                    v1 = 65535
                    if false {
                        switch -2147483648 {
                        case -703:
                            v3[65535] = 424
                        case 42:
                            if false {
                                v16 := GenS1{F0: (fi8 + fi8), F1: 255}
                                _ = v16
                            }
                        }
                    }
                }
            }
        }
        v5 += (v5 ^ (^(v5)))
        _ = fi8
    }
    var v17 float32 = -2.25
    v18 := GenS1{F0: 128, F1: v1}
    switch "\n\t" {
    case "hello":
        fmt.Println(v4, v4)
    }
    var v19 float32 = v17
    var v20 string = "foo bar"
    fmt.Println(v1, v5, v6)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v17 += (((v17 + -2.25) * v19) - v17)
    v4 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
