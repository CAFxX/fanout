package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 int }
type GenS3 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int8 }
func genFn0(p0 int32) int16 {
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p0, p0)
        v8 := GenS1{F0: -46, F1: 810, F2: 65535}
        var v9 uint = 65535
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int16 = 547
    if (255 > 128) {
        v11 := GenS2{F0: true, F1: 2, F2: 44}
        if (((("foo bar" + "hello") + ("\n\t" + "\n\t")) < "") && (!((3.14159 / 0.1) >= float64((3.14159 + (0.0 / float32(1.5))))))) {
            fmt.Println(v10, p0)
            v10 -= (v10 / ((v10) | 1))
        } else {
            for fi12 := 0; fi12 < 8; fi12++ {
                switch (-(fi12)) {
                case 128:
                    var v13 bool = true
                    _ = v13
                case 42:
                    var v14 uint = 42
                    _ = v14
                case 4294967295:
                    switch "a" {
                    case "hello":
                        v15 := []int{fi12, fi12}
                        _ = v15
                    }
                }
                _ = fi12
            }
            p0 += (^(p0))
        }
        _ = v11
    } else {
        v16 := []int{2147483647}
        v10 -= (v10 ^ v10)
        _ = v16
    }
    var v17 int8 = -128
    _ = p0
    _ = v10
    _ = v17
    return (^(v10))
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{"hello"}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    var v5 string = "foo bar"
    var v6 uint16 = 255
    var v18 int32 = int32(v1)
    v5 = (v5 + v5)
    v2 = append(v2, v5)
    var v19 uint64 = uint64(v6)
    switch 4294967295 {
    case 919:
        v19 = (v19 & (^(v19)))
    case 42:
        v20 := v4[2147483647]
        _ = v20
    default:
        for fi21 := 0; fi21 < 20; fi21++ {
            if false {
                fmt.Println(fi21, v6)
            } else {
                for fi22 := 0; fi22 < 7; fi22++ {
                    v23 := GenS3{F0: (("" <= (v5 + "hello")) || true), F1: uint8(v19), F2: (v19 ^ (v19 % ((4294967295) | 1)))}
                    fmt.Println(fi22, v6)
                    v24 := GenS1{F0: -1, F1: (v18 % ((v18) | 1)), F2: v6}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                var v25 int32 = (v18 % ((int32(v6)) | 1))
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v5)
                default:
                    var v26 int16 = 32767
                    _ = v26
                }
                _ = v25
            }
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 int8 = -46
                var v29 uint64 = v19
                var v30 uint = 65535
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            var v31 string = "a"
            _ = fi21
            _ = v31
        }
    }
    v6 += (^(v6))
    switch v5 {
    case "foo bar":
        v32 := GenS4{F0: 2, F1: (v1 - 466), F2: 42}
        _ = v32
    default:
        v33 := GenS2{F0: ((!false) && true), F1: 18446744073709551615, F2: -9223372036854775808}
        _ = v33
    }
    switch ("hello" + v5) {
    case "a":
        if (-1 < v18) {
            v34 := GenS2{F0: (!((1 <= 18446744073709551615) && true)), F1: 328, F2: 4294967295}
            if true {
                v18 = v18
                for fi35 := 0; fi35 < 17; fi35++ {
                    var v36 int = 1
                    v1 += v1
                    v37 := GenS1{F0: 0, F1: (v18 | (v18 | v18)), F2: v6}
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                v38 := GenS3{F0: true, F1: 0, F2: 697}
                _ = v38
            } else {
                var v39 float64 = float64(((float32(1.5) - -2.25) + 3.14159))
                v1 = v1
                _ = v39
            }
            _ = v34
        } else {
            v6 += 65535
            v40 := GenS1{F0: 482, F1: 42, F2: v6}
            var v41 uint16 = (^(v6))
            _ = v40
            _ = v41
        }
    }
    var v42 int8 = -128
    if (863 > 255) {
        var v43 uint16 = uint16(v19)
        _ = v43
    }
    v44 := GenS5{F0: 2, F1: v42}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v5
    _ = v6
}
