package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int8; F2 string }
func genFn0(p0 int64, p1 uint8) uint32 {
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint32, p1 int64, p2 int64) int {
    var v6 float32 = (-2.25 + 1e10)
    v7 := GenS1{F0: 9223372036854775807, F1: 65535}
    v8 := GenS1{F0: (p2 * p2), F1: 2}
    var v9 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 int8 = -31
    v2 := make(map[string]bool)
    var v3 float32 = -2.25
    var v4 uint64 = 606
    var v5 string = "foo bar"
    v3 = float32((1.5 * 1e10))
    v3 = (((v3 / 1.5) + v3) - float32(float64(v3)))
    var v10 uint = 206
    if false {
        v4 = 169
        if true {
            v11 := GenS2{F0: 128, F1: 42, F2: (false || true)}
            var v12 uint8 = 0
            _ = v11
            _ = v12
        } else {
            var v13 bool = false
            _ = v13
        }
    } else {
        var v14 int32 = 42
        var v15 int32 = (-(v14))
        _ = v14
        _ = v15
    }
    v16 := []int{2147483647}
    for fi17 := 0; fi17 < 6; fi17++ {
        v4 -= 263
        var v18 float32 = float32(1.5)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 int32 = -1
            v21 -= v21
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    switch 42 {
    case 9223372036854775807:
        fmt.Println(v4, v5, v10)
    case 127:
        _ = len(v2)
    case -614:
        for fi22 := 0; fi22 < 2; fi22++ {
            if len(v16) > 0 {
                v16[0] = (fi22 << (((^(v10))) & 63))
            }
            if ((1.5 / v3) >= (3.14159 * (0.1 - float32(1.5)))) {
                for fi23 := 0; fi23 < 2; fi23++ {
                    switch -1 {
                    case -522:
                        v4 += 535
                        switch (((v5 + (("foo bar" + (("foo bar" + (((((("foo bar" + ((("a" + v5) + (v5 + ("hello" + (("hello" + (((((v5 + (v5 + v5)) + v5) + "a") + "") + "hello")) + "hello")))) + "foo bar")) + ("" + v5)) + ("foo bar" + v5)) + "hello") + "\n\t") + (v5 + ("\n\t" + ("" + (v5 + (v5 + (v5 + v5)))))))) + v5)) + (v5 + v5))) + (v5 + "a")) + v5) {
                        case "\n\t":
                            fmt.Println(v3)
                        case "foo bar":
                            v24 := GenS1{F0: 42, F1: 952}
                            _ = v24
                        case "hello":
                            var v25 int16 = -42
                            _ = len(v2)
                            fmt.Println(v5, v4)
                            _ = v25
                        }
                        v26 := GenS3{F0: (fi23 - (fi22 ^ (fi23 - (fi22 << ((v10) & 63)))))}
                        _ = v26
                    case 128:
                        _ = len(v16)
                    default:
                        var v27 uint8 = 242
                        _ = v27
                    }
                    v28 := GenS1{F0: 1, F1: 1}
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = fi23
                    _ = v28
                }
                var v29 int16 = 0
                _ = v29
            } else {
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 bool = true
                    var v32 uint8 = 128
                    if (0 >= 1) {
                        fmt.Println(v31, v10)
                        fmt.Println(fi30, v4, v31)
                    }
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                var v33 string = "\n\t"
                for fi34 := 0; fi34 < 10; fi34++ {
                    if len(v16) > 0 {
                        v16[0] = genFn1(702, -193, int64(fi34))
                    }
                    v35 := GenS4{F0: v10}
                    var v36 bool = false
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                _ = v33
            }
            v37 := GenS5{F0: (int64(v1) != 42), F1: (v1 % ((v1) | 1)), F2: v5}
            _ = fi22
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
