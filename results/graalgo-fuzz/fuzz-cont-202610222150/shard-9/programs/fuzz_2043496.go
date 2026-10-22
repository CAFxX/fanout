package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 bool }
func genFn0(p0 int32, p1 int16) uint16 {
    p0 -= (p0 + p0)
    v6 := []int{int(p1), 1, -813}
    p0 = 53
    var v7 uint8 = 0
    if false {
        v7 = (v7 / ((v7) | 1))
        p1 += p1
        v7 = (v7 / ((239) | 1))
    } else {
        var v8 uint16 = 0
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 973
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"", "foo bar", "hello", "foo bar"}
    var v3 int16 = 1
    var v4 string = ("hello" + ("hello" + "a"))
    var v5 bool = true
    var v9 uint64 = 520
    if v5 {
        v10 := GenS1{F0: 42, F1: 552, F2: (2147483647 < 2147483647)}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: 0, F1: 65535, F2: (!("" > ((v4 + v4) + v4)))}
                var v14 int8 = 127
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := GenS1{F0: v3, F1: 2147483648, F2: (42 > 294)}
            _ = fi11
            _ = v15
        }
        var v16 int16 = (v3 / ((0) | 1))
        _ = v16
    }
    if v5 {
        v5 = (!(!true))
    }
    var v17 float64 = 0.0
    var v18 int64 = 9223372036854775807
    if v5 {
        if v5 {
            var v19 int64 = (^(v18))
            switch 9223372036854775807 {
            case 9223372036854775807:
                if (v3 <= v3) {
                    var v20 int32 = 2147483647
                    v9 = (v9 >> ((255) & 63))
                    _ = v20
                } else {
                    var v21 bool = (false || v5)
                    fmt.Println(v17, v21, v9)
                    v17 = -2.25
                    _ = v21
                }
            case -1:
                fmt.Println(v17, v5, v4)
            }
            _ = v19
        }
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v18 -= -9223372036854775808
        var v23 int64 = (v18 % (((-(v18))) | 1))
        var v24 float32 = 1e10
        _ = fi22
        _ = v23
        _ = v24
    }
    v4 = "foo bar"
    if false {
        v25 := GenS1{F0: int16(v9), F1: 65535, F2: v5}
        for fi26 := 0; fi26 < 8; fi26++ {
            if (!(genFn0(-2147483648, (v3 | int16(v9))) <= uint16(v9))) {
                fmt.Println(fi26)
            }
            var v27 int64 = v18
            for fi28 := 0; fi28 < 4; fi28++ {
                if true {
                    fmt.Println(v5, v5)
                } else {
                    switch "" {
                    case "foo bar":
                        var v29 int32 = 2147483647
                        _ = v29
                    case "hello":
                        v5 = (v5 && (1.5 >= 1e10))
                    default:
                        var v30 float32 = (1e10 / 0.1)
                        _ = v30
                    }
                }
                _ = fi28
            }
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        var v31 uint64 = v9
        _ = len(v2)
        _ = v31
    }
    v32 := v1[v4]
    switch -609 {
    case 127:
        if v5 {
            fmt.Println(v3, v17)
            for fi33 := 0; fi33 < 1; fi33++ {
                fmt.Println(v5, fi33)
                _ = fi33
            }
        }
    default:
        var v34 uint32 = 65535
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
