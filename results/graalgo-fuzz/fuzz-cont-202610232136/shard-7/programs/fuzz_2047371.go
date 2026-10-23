package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 string; F2 uint }
func genFn0(p0 int32, p1 uint32) uint16 {
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(p0)
        v7 := GenS1{F0: 42, F1: (p0 >> ((2) & 63))}
        v8 := []int{fi6, (^(fi6))}
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 = p0
    var v9 int32 = p0
    _ = p0
    _ = p1
    _ = v9
    return 2
}
func genFn1(p0 uint) string {
    v10 := GenS2{F0: 997}
    v11 := GenS3{F0: 32767, F1: (((("foo bar" + ("foo bar" + "\n\t")) + "hello") + "\n\t") + "hello"), F2: p0}
    var v12 uint64 = 42
    v13 := []int{9223372036854775807, 513, -417, 4294967295}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return ("hello" + ("" + "\n\t"))
}

func main() {
    v1 := []string{"\n\t", ("\n\t" + "a"), "hello"}
    var v2 bool = false
    var v3 uint64 = 0
    v4 := make(map[string]string)
    v5 := make(map[uint32]int)
    var v14 float64 = ((((((((3.14159 - float64(((1e10 - 0.0) - (((-2.25 - -2.25) / 3.14159) * 1e10)))) / 3.14159) + -2.25) + 1e10) + 0.1) + 3.14159) * (0.0 - -2.25)) - 1.5)
    v1 = append(v1, "foo bar")
    var v15 int16 = 0
    var v16 float64 = float64(0.1)
    for fi17 := 0; fi17 < 8; fi17++ {
        v14 = (v16 / (((0.0 / (v16 + float64(((1e10 - (((0.0 / -2.25) + (1e10 / 1.0)) - float32(v16))) * (-2.25 - float32(v16)))))) / 1e10) - ((0.1 / v14) - v14)))
        _ = fi17
    }
    var v18 float64 = v16
    v1 = append(v1, "foo bar")
    v15 -= v15
    v5[2147483648] = -248
    if true {
        var v19 int16 = (v15 >> ((4294967295) & 63))
        _ = v19
    } else {
        fmt.Println(v14, v14)
    }
    v5[65535] = 343
    v14 -= float64(3.14159)
    if (42 == -1) {
        v16 = 0.1
        _ = len(v1)
        for fi20 := 0; fi20 < 6; fi20++ {
            for fi21 := 0; fi21 < 19; fi21++ {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                var v22 uint = 42
                switch fi21 {
                case 9223372036854775807:
                    fmt.Println(v18)
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v23 uint64 = 4294967295
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = len(v5)
            _ = fi20
        }
    } else {
        var v24 uint8 = 128
        _ = v24
    }
    v15 = 32767
    var v25 uint8 = 1
    v16 += v18
    switch "a" {
    case "a":
        if v2 {
            fmt.Println(v25, v16)
        }
    case "":
        var v26 uint16 = genFn0(609, 65535)
        _ = v26
    case "hello":
        if true {
            switch "" {
            case "\n\t":
                var v27 uint = 18446744073709551615
                _ = v27
            case "foo bar":
                var v28 int16 = v15
                _ = v28
            }
            if v2 {
                if v2 {
                    var v29 uint32 = 732
                    v30 := GenS1{F0: 1, F1: -1}
                    v1 = append(v1, (("\n\t" + "") + "\n\t"))
                    _ = v29
                    _ = v30
                } else {
                    v15 += (v15 - (-(v15)))
                }
                if true {
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    var v31 uint = 65535
                    _ = v31
                } else {
                    v1 = append(v1, "\n\t")
                }
                fmt.Println(v3, v25)
            } else {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
            }
            for fi32 := 0; fi32 < 13; fi32++ {
                var v33 string = ("" + ((((("hello" + "foo bar") + "foo bar") + "hello") + "") + ""))
                _ = fi32
                _ = v33
            }
        } else {
            v1 = append(v1, "a")
        }
    default:
        var v34 float64 = (3.14159 / 0.1)
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
