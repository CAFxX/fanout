package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int; F1 int8; F2 bool }
type GenS3 struct { F0 int16; F1 int8; F2 int }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 string }
func genFn0(p0 string) uint {
    var v6 int8 = -79
    p0 = (((p0 + ("hello" + "hello")) + (((((p0 + "foo bar") + p0) + ((p0 + p0) + p0)) + (p0 + "foo bar")) + p0)) + "foo bar")
    var v7 float32 = (float32(3.14159) + (-2.25 / 1e10))
    for fi8 := 0; fi8 < 2; fi8++ {
        if ((2 != 0) && (1.5 < v7)) {
            v6 = v6
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS1{F0: true, F1: 947}
                if true {
                    var v11 float32 = 0.1
                    _ = v11
                }
                v12 := GenS1{F0: false, F1: 65535}
                _ = fi9
                _ = v10
                _ = v12
            }
        }
        _ = fi8
    }
    _ = p0
    _ = v6
    _ = v7
    return 430
}
func genFn1(p0 uint32, p1 uint16) uint16 {
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 float32 = 3.14159
        if true {
            if true {
                v14 += 1.5
                v15 := GenS2{F0: (-(fi13)), F1: -89, F2: false}
                v16 := GenS1{F0: true, F1: 1}
                _ = v15
                _ = v16
            } else {
                var v17 uint = 65535
                _ = v17
            }
        }
        var v18 uint = genFn0(("hello" + ("hello" + ((("hello" + ("\n\t" + "a")) + "hello") + ("hello" + "foo bar")))))
        _ = fi13
        _ = v14
        _ = v18
    }
    v19 := GenS1{F0: false, F1: 907}
    v20 := []int{32768, 42}
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{32768}
    v3 := []string{}
    v4 := make(map[int]bool)
    v5 := []uint32{}
    var v21 float32 = -2.25
    v21 = float32(0.1)
    v21 = v21
    v22 := GenS1{F0: ((("a" + "foo bar") + "a") <= ("a" + "\n\t")), F1: 18446744073709551615}
    if false {
        v21 = (v21 - (((0.0 * v21) - (0.0 / v21)) + 1e10))
        var v23 uint16 = genFn1(450, 255)
        switch 94 {
        case 42:
            delete(v4, -1)
        case 128:
            for fi24 := 0; fi24 < 18; fi24++ {
                for fi25 := 0; fi25 < 20; fi25++ {
                    if (true || (float32(1e10) > -2.25)) {
                        v26 := GenS3{F0: -32768, F1: 1, F2: 0}
                        var v27 uint16 = genFn1(uint32(v23), v23)
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS4{F0: -1, F1: 42}
                    delete(v1, (("foo bar" + "\n\t") + "foo bar"))
                    _ = fi25
                    _ = v28
                }
                _ = fi24
            }
        default:
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 uint16 = (v23 ^ (v23 - genFn1(318, v23)))
                var v31 uint16 = v23
                var v32 float32 = ((1.5 / -2.25) - float32((1.5 - 0.0)))
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        }
        _ = v23
    } else {
        var v33 float32 = 1.5
        var v34 int32 = 637
        for fi35 := 0; fi35 < 7; fi35++ {
            if len(v2) > 0 {
                v2[0] = (fi35 - (fi35 & fi35))
            }
            fmt.Println(v21)
            fmt.Println(v21, fi35, v34)
            _ = fi35
        }
        _ = v33
        _ = v34
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        v21 = v21
        var v37 uint16 = 65535
        _ = fi36
        _ = v37
    }
    v21 += v21
    v38 := v4[0]
    fmt.Println(v38, v21)
    v39 := GenS5{F0: "hello"}
    v21 -= ((float32(float64(v21)) / v21) - float32(0.0))
    v40 := GenS2{F0: 9223372036854775807, F1: -94, F2: v38}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
