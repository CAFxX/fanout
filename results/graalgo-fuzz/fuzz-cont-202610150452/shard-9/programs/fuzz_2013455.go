package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int; F2 uint }
func genFn0(p0 bool, p1 bool) int8 {
    var v6 uint = 595
    v7 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -128
}
func genFn1(p0 string) uint8 {
    v8 := GenS1{F0: 2147483647}
    p0 = p0
    _ = p0
    _ = v8
    return 255
}

func main() {
    var v1 float64 = (0.1 / (3.14159 / 0.1))
    var v2 uint8 = 123
    var v3 uint = 18446744073709551615
    var v4 string = ("foo bar" + "foo bar")
    v5 := make(map[int]int)
    var v9 uint16 = 113
    v3 = v3
    switch "" {
    case "a":
        var v10 float64 = v1
        _ = v10
    case "foo bar":
        var v11 uint64 = 2
        _ = v11
    default:
        switch 9223372036854775807 {
        case 32768:
            _ = len(v5)
        }
    }
    var v12 int32 = 619
    var v13 uint64 = 934
    v14 := GenS1{F0: (v12 / ((580) | 1))}
    v12 = v12
    v15 := []int{-1, -1, 983}
    var v16 float32 = 3.14159
    var v17 int8 = 0
    var v18 uint = v3
    fmt.Println(v2, v16, v2)
    if (297 >= 42) {
        fmt.Println(v9)
        var v19 int8 = v17
        _ = v19
    } else {
        if false {
            if true {
                v17 = v17
                for fi20 := 0; fi20 < 12; fi20++ {
                    v5[(-(fi20))] = -2147483648
                    _ = fi20
                }
            } else {
                if false {
                    v9 = v9
                    v2 = v2
                    v12 -= 2147483647
                } else {
                    if (v12 == 1) {
                        v9 = (^(v9))
                    } else {
                        _ = len(v5)
                        var v21 uint16 = v9
                        _ = v21
                    }
                    switch "" {
                    case "foo bar":
                        for fi22 := 0; fi22 < 4; fi22++ {
                            var v23 uint64 = (^(v13))
                            _ = fi22
                            _ = v23
                        }
                    case "hello":
                        var v24 uint64 = (v13 % (((^(v13))) | 1))
                        _ = v24
                    case "a":
                        fmt.Println(v16, v9)
                    default:
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    }
                    v25 := GenS2{F0: v2, F1: 0, F2: 2}
                    _ = v25
                }
                v4 = "a"
            }
            var v26 uint32 = 335
            v27 := GenS1{F0: (v12 & v12)}
            _ = v26
            _ = v27
        }
        fmt.Println(v4, v12, v12)
    }
    v16 += (v16 / (-2.25 - 1e10))
    if ((^(v9)) <= (v9 / (((v9 >> ((v18) & 63))) | 1))) {
        switch -544 {
        case 32767:
            fmt.Println(v16, v16, v13)
        }
    }
    v3 = (v18 % ((uint(v2)) | 1))
    v28 := GenS2{F0: (^(v2)), F1: 9223372036854775807, F2: (v3 + (^(v18)))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
