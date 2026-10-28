package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string, p1 bool) uint {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 bool = (1 < 114)
        p1 = v8
        _ = fi7
        _ = v8
    }
    v9 := []int{809, 32767, 2147483647}
    fmt.Println(p0, p1)
    if ((2 > 128) && ((p1 || (176 != 1)) && p1)) {
        p0 = ("foo bar" + (p0 + (p0 + p0)))
    }
    _ = p0
    _ = p1
    _ = v9
    return 972
}
func genFn1(p0 uint16, p1 uint64, p2 uint32) int64 {
    fmt.Println(p0, p0, p1)
    switch 127 {
    case 128:
        v10 := GenS1{F0: 42}
        _ = v10
    case -9223372036854775808:
        v11 := GenS1{F0: 242}
        _ = v11
    case 4294967295:
        if false {
            if (-967 > 141) {
                var v12 bool = (true || (!(true && (false || ((^(p1)) == p1)))))
                var v13 float32 = (0.1 + (-2.25 - (3.14159 + -2.25)))
                _ = v12
                _ = v13
            } else {
                v14 := []int{-800, 1, 1}
                switch 612 {
                case -945:
                    p2 -= 2147483647
                case 0:
                    p2 = p2
                case 748:
                    fmt.Println(p2, p1, p0)
                }
                var v15 float64 = 3.14159
                _ = v14
                _ = v15
            }
        } else {
            fmt.Println(p2, p2, p2)
            v16 := []int{283, 42}
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -2147483648
    v3 := make(map[int]int)
    var v4 uint16 = 65535
    var v5 uint64 = 1
    var v6 string = "a"
    v17 := []int{2147483647, 127}
    var v18 bool = (2147483647 == 0)
    v4 = v4
    delete(v1, 597)
    if v18 {
        if (0.0 > 1e10) {
            var v19 uint = 4294967295
            _ = v19
        }
        v20 := GenS1{F0: 248}
        v4 -= 2
        _ = v20
    } else {
        _ = len(v3)
    }
    v5 = v5
    v21 := v1[-2147483648]
    v2 = v2
    for fi22 := 0; fi22 < 4; fi22++ {
        fmt.Println(v6, v5, v4)
        fmt.Println(v6, v2, v21)
        _ = fi22
    }
    var v23 uint16 = (v4 / ((uint16(v5)) | 1))
    fmt.Println(v18)
    fmt.Println(v6, v2)
    v24 := GenS2{F0: v4}
    v2 = (v2 ^ (v2 | 42))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
