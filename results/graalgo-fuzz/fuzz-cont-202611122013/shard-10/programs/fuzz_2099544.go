package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int32 }
func genFn0(p0 int16, p1 int16) int64 {
    var v6 int = 42
    var v7 uint = 600
    v8 := []int{-14, v6, int(v7), (v6 & v6)}
    var v9 int = v6
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS1{F0: ((0.1 == 1e10) && ((float32(1e10) + 1e10) <= 1e10))}
            v7 = (^(v7))
            _ = fi11
            _ = v12
        }
        fmt.Println(v6)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 bool, p1 uint64, p2 uint64) string {
    p1 = 18446744073709551615
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS1{F0: ((0.1 == 3.14159) && p0)}
        _ = fi13
        _ = v14
    }
    v15 := []int{2147483647, -2147483648, 32767, 32767}
    switch 42 {
    case 2147483647:
        if (1e10 > -2.25) {
            var v16 float32 = (float32(1e10) * (-2.25 + 3.14159))
            switch ("a" + "foo bar") {
            case "":
                if p0 {
                    var v17 uint8 = 121
                    if false {
                        p2 -= p2
                    }
                    var v18 int = 0
                    _ = v17
                    _ = v18
                }
            default:
                var v19 int = -979
                _ = v19
            }
            _ = v16
        }
    default:
        var v20 int8 = -1
        _ = v20
    }
    if (2 == 0) {
        p1 -= (p1 / ((18446744073709551615) | 1))
        p0 = true
        fmt.Println(p1)
    } else {
        fmt.Println(p1)
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (("a" + "") + "foo bar")
}

func main() {
    v1 := []int{32768, -1, 144}
    v2 := make(map[int]string)
    v3 := []string{((("\n\t" + "a") + "a") + (("a" + "a") + "a")), (("hello" + ("hello" + "\n\t")) + "foo bar")}
    v4 := []int{9223372036854775807, 140, 42, -2147483648}
    v5 := []string{"foo bar"}
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 float64 = (float64((0.0 / -2.25)) / 1.5)
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 bool = (0 != 42)
            v25 := GenS2{F0: v24, F1: -2147483648}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        for fi26 := 0; fi26 < 13; fi26++ {
            if (2 == 2) {
                fmt.Println(fi26, fi26)
                v22 = float64(3.14159)
            }
            if true {
                fmt.Println(fi26, v22, v22)
            }
            _ = fi26
        }
        _ = fi21
        _ = v22
    }
    var v27 uint = 929
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v27 += v27
    v4 = append(v4, 32768)
    if (128 >= 0) {
        for fi28 := 0; fi28 < 10; fi28++ {
            v27 = v27
            _ = fi28
        }
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint = (v27 - (v27 - v27))
        if true {
            switch "foo bar" {
            case "\n\t":
                var v31 int8 = -128
                _ = v31
            case "hello":
                var v32 int8 = -63
                _ = v32
            }
        } else {
            var v33 int64 = -1
            var v34 uint64 = 2
            v30 -= (v27 << ((v30) & 63))
            _ = v33
            _ = v34
        }
        _ = fi29
        _ = v30
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        fmt.Println(fi35, fi35)
        _ = fi35
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        delete(v2, (fi36 - (fi36 << ((42) & 63))))
        _ = fi36
    }
    v37 := GenS1{F0: false}
    var v38 uint16 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v39 uint8 = 102
    var v40 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
