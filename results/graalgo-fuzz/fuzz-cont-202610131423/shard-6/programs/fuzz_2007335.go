package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 string }
type GenS2 struct { F0 uint32; F1 string }
func genFn0(p0 string) uint {
    switch 127 {
    case 0:
        var v5 int64 = -9223372036854775808
        _ = v5
    default:
        v6 := []int{32768, 4294967295, 0}
        _ = v6
    }
    var v7 uint = 663
    switch -1 {
    case 765:
        var v8 uint32 = uint32(v7)
        _ = v8
    case 290:
        var v9 int = 32767
        _ = v9
    }
    var v10 uint32 = 2147483648
    v7 -= (v7 | v7)
    _ = p0
    _ = v10
    _ = v7
    return 922
}
func genFn1(p0 uint64, p1 uint, p2 uint64) int {
    if (true || false) {
        p0 = p0
        var v11 uint16 = 2
        _ = v11
    } else {
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 int64 = -212
            if true {
                fmt.Println(p2, fi12, fi12)
                v14 := []int{fi12}
                v13 = -9223372036854775808
                _ = v14
            } else {
                var v15 uint8 = 255
                v16 := GenS1{F0: -128, F1: -1, F2: "foo bar"}
                _ = v15
                _ = v16
            }
            switch (fi12 * 42) {
            case 32767:
                v17 := []int{(fi12 >> ((4294967295) & 63)), fi12, 176}
                _ = v17
            case -1:
                p2 -= p2
            default:
                var v18 uint8 = 115
                _ = v18
            }
            _ = fi12
            _ = v13
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := []int{fi19, fi19}
            for fi21 := 0; fi21 < 14; fi21++ {
                v22 := GenS1{F0: 127, F1: -818, F2: "hello"}
                var v23 int16 = -1
                var v24 float64 = 0.1
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            var v25 uint = (^(p1))
            _ = fi19
            _ = v20
            _ = v25
        }
        if true {
            switch int(p1) {
            case -9223372036854775808:
                fmt.Println(p2, p0)
            default:
                v26 := GenS2{F0: 66, F1: "a"}
                _ = v26
            }
            if true {
                var v27 float64 = 0.0
                if false {
                    p2 += 767
                }
                _ = v27
            }
            var v28 bool = false
            _ = v28
        }
    }
    p2 = (p0 << ((genFn0((("" + ("foo bar" + "foo bar")) + ("hello" + ("a" + ""))))) & 63))
    var v29 float32 = 1.5
    for fi30 := 0; fi30 < 8; fi30++ {
        v31 := GenS1{F0: -70, F1: 31, F2: ("hello" + "hello")}
        p2 = p2
        _ = fi30
        _ = v31
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return -516
}

func main() {
    var v1 int64 = -1
    var v2 int64 = (-(v1))
    v3 := make(map[int]bool)
    v4 := []uint32{2147483648}
    v1 -= v1
    if true {
        if false {
            v1 = v1
        } else {
            var v32 int8 = -128
            _ = v32
        }
        fmt.Println(v1, v1)
        var v33 int8 = -1
        _ = v33
    } else {
        v2 += (v2 & 42)
        v2 = (v2 | (v1 - (v1 * v1)))
        v34 := v3[9223372036854775807]
        _ = v34
    }
    var v35 int32 = 42
    switch 32767 {
    case 648:
        fmt.Println(v35, v1, v1)
    }
    v1 = (v1 - -9223372036854775808)
    v35 -= v35
    fmt.Println(v2, v1, v35)
    v35 = v35
    v1 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
