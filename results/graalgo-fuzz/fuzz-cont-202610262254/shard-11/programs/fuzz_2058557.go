package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 bool, p1 uint16, p2 int32) int8 {
    var v6 string = ("hello" + ("a" + ("foo bar" + ((("" + "foo bar") + "foo bar") + "foo bar"))))
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 100
}
func genFn1(p0 int8, p1 uint8) bool {
    v8 := GenS1{F0: 94}
    if (32767 < 613) {
        var v9 float64 = 0.0
        if false {
            p0 = (^(p0))
            p0 = 127
        }
        _ = v9
    } else {
        v10 := GenS1{F0: 2147483647}
        v11 := GenS1{F0: 42}
        _ = v10
        _ = v11
    }
    var v12 int = 42
    var v13 int = (v12 >> ((101) & 63))
    p0 -= genFn0(((v12 * (v12 | -88)) == (v12 + -293)), 42, 1)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v8
    return true
}

func main() {
    var v1 uint32 = 2147483648
    var v2 float64 = (1.5 / 1.5)
    var v3 uint64 = 4294967295
    v4 := make(map[string]uint32)
    v5 := make(map[int]string)
    if true {
        v1 = 1
        v14 := GenS2{F0: 110, F1: (v3 % ((18446744073709551615) | 1))}
        v1 = v1
        _ = v14
    } else {
        var v15 string = ""
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 float64 = (v2 / (v2 - 3.14159))
            var v18 uint16 = 1
            v19 := []int{42}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        if true {
            v20 := []int{-2147483648, -9223372036854775808, -2147483648}
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 uint8 = 255
                _ = fi21
                _ = v22
            }
            var v23 uint16 = 677
            _ = v20
            _ = v23
        }
        _ = v15
    }
    v24 := []int{4294967295}
    switch int(v3) {
    case 4294967295:
        for fi25 := 0; fi25 < 15; fi25++ {
            for fi26 := 0; fi26 < 4; fi26++ {
                v3 = (v3 - (^(v3)))
                v3 = (v3 & v3)
                fmt.Println(fi25, v3, fi26)
                fmt.Println(fi26, fi25, fi25)
                var v27 int64 = -1
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    }
    if (4 == -405) {
        if (!true) {
            _ = len(v24)
        }
        fmt.Println(v2, v2, v1)
        var v28 int8 = 127
        _ = v28
    }
    v2 += (float64((1e10 * 0.0)) - 0.1)
    fmt.Println(v3, v3)
    fmt.Println(v3)
    var v29 int8 = genFn0(true, 32767, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
