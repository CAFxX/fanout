package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint32; F1 bool; F2 int8 }
func genFn0() int {
    if ((((0 > 1) && (158 > 127)) || (1e10 == 1.5)) || false) {
        switch "hello" {
        case "a":
            var v7 float32 = 0.1
            _ = v7
        case "":
            var v8 uint8 = 255
            _ = v8
        case "foo bar":
            v9 := []int{57}
            _ = v9
        default:
            if true {
                var v10 uint32 = 42
                v10 = v10
                var v11 string = (("a" + "") + ("a" + ("" + "a")))
                _ = v10
                _ = v11
            } else {
                var v12 uint16 = 2
                v13 := []int{-2147483648, 631, 32768}
                for fi14 := 0; fi14 < 17; fi14++ {
                    fmt.Println(v12)
                    _ = fi14
                }
                _ = v12
                _ = v13
            }
        }
        var v15 float32 = 1e10
        _ = v15
    }
    v16 := []int{-1, 916, 42, 42}
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 uint16 = 32767
        if false {
            var v19 uint16 = v18
            var v20 uint64 = 620
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: -32768}
            var v22 int = (fi17 | (fi17 - fi17))
            _ = v21
            _ = v22
        }
        _ = fi17
        _ = v18
    }
    _ = v16
    return 4294967295
}
func genFn1(p0 int8, p1 int32) int16 {
    if false {
        v23 := []int{4294967295}
        p1 = p1
        _ = v23
    } else {
        p1 = (p1 - (p1 % ((p1) | 1)))
        p0 -= p0
        v24 := []int{42, 0, -916}
        _ = v24
    }
    var v25 bool = (true && false)
    p1 = p1
    var v26 int = 1
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return 1
}

func main() {
    var v1 int64 = 959
    var v2 uint16 = 876
    v3 := make(map[int]string)
    var v4 int = -930
    v5 := []uint32{4294967295, 980, 2147483648, 42}
    var v6 int64 = v1
    switch "a" {
    case "foo bar":
        _ = len(v3)
    case "a":
        v4 += -914
    case "\n\t":
        v27 := GenS2{F0: 319, F1: (100 > 42), F2: -128}
        _ = v27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        for fi29 := 0; fi29 < 3; fi29++ {
            v6 -= 1
            var v30 uint32 = 4294967295
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    var v31 int = v4
    var v32 float64 = 1e10
    v33 := v3[v31]
    v5 = append(v5, uint32(v2))
    for fi34 := 0; fi34 < 18; fi34++ {
        if (!(0 != 2147483647)) {
            if false {
                fmt.Println(v2, v4, fi34)
            } else {
                v1 = (^(v1))
                fmt.Println(v6, v4)
            }
        }
        _ = fi34
    }
    v1 -= (v1 * v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
