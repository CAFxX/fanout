package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
func genFn0(p0 int64, p1 uint8, p2 uint8) uint64 {
    p1 += p1
    var v7 int64 = p0
    if ((("a" + "foo bar") + "a") <= ("\n\t" + "hello")) {
        var v8 int32 = 1
        if false {
            v9 := []int{}
            _ = v9
        }
        _ = v8
    } else {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    var v1 int = -79
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 2
    var v4 uint16 = 219
    v5 := make(map[uint32]int)
    var v6 uint16 = (v4 + 2)
    var v10 float32 = -2.25
    fmt.Println(v2)
    var v11 uint = 700
    v2 += (v2 >> ((v11) & 63))
    var v12 uint = v11
    switch v1 {
    case -9223372036854775808:
        var v13 string = "foo bar"
        _ = v13
    case -873:
        fmt.Println(v3, v4, v12)
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            switch ("" + "") {
            case "foo bar":
                v16 := []int{(fi14 * (^(v1))), (v1 / ((2147483647) | 1)), (fi14 & fi15), fi14}
                _ = v16
            case "a":
                var v17 float64 = 1e10
                _ = v17
            default:
                v4 = v4
            }
            fmt.Println(fi15, fi14)
            v10 = v10
            _ = fi15
        }
        _ = fi14
    }
    var v18 int32 = -1
    v12 = 674
    v19 := GenS1{F0: "", F1: v1}
    switch (v1 ^ (v1 + (v1 - (-(v1))))) {
    case -645:
        v12 = v11
    default:
        v2 += (^(v2))
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS1{F0: "hello", F1: 127}
        for fi22 := 0; fi22 < 8; fi22++ {
            delete(v5, 65535)
            switch "" {
            case "":
                v23 := v5[158]
                _ = v23
            case "foo bar":
                v1 = -1
            default:
                switch "hello" {
                case "hello":
                    if false {
                        v4 += v4
                    } else {
                        var v24 float32 = float32(3.14159)
                        if false {
                            v12 = (v11 % (((v12 * v11)) | 1))
                            var v25 int32 = (^(v18))
                            _ = len(v5)
                            _ = v25
                        } else {
                            v12 -= (^(v11))
                            v4 += (^(v4))
                        }
                        var v26 uint8 = 0
                        var v27 uint16 = v4
                        _ = v24
                        _ = v26
                        _ = v27
                    }
                case "a":
                    if false {
                        var v28 string = "\n\t"
                        v29 := []int{(^(fi20)), -2147483648, (fi20 + (fi22 / ((fi22) | 1))), (fi22 + fi22)}
                        _ = v28
                        _ = v29
                    }
                }
            }
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    v30 := GenS1{F0: "\n\t", F1: v1}
    var v31 uint8 = 124
    fmt.Println(v4, v4, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
