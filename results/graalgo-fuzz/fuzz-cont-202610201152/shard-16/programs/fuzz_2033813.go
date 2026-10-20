package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint }
type GenS3 struct { F0 int }
func genFn0(p0 bool) uint16 {
    v4 := []int{}
    var v5 float32 = (0.1 - 0.1)
    v6 := GenS1{F0: 65535, F1: 1}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int8 = 42
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v7 bool = ("foo bar" == "\n\t")
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: 942, F1: 4294967295}
        _ = fi8
        _ = v9
    }
    if v7 {
        if true {
            v7 = v7
        }
        v7 = (65535 >= 2147483647)
        _ = len(v2)
    } else {
        v7 = (!(42 < 9223372036854775807))
        fmt.Println(v7, v1)
    }
    v10 := GenS1{F0: 42, F1: 2147483647}
    fmt.Println(v7)
    v7 = ("hello" > ("hello" + ("foo bar" + ("a" + "\n\t"))))
    var v11 int8 = (-(v1))
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint8 = 231
        var v14 float32 = (((3.14159 + -2.25) - 3.14159) * (-2.25 - 0.1))
        switch "hello" {
        case "hello":
            if ((("" + "foo bar") + "a") != ("" + ("foo bar" + (("foo bar" + "\n\t") + ("hello" + "\n\t"))))) {
                var v15 int = fi12
                _ = v15
            } else {
                var v16 int32 = 0
                for fi17 := 0; fi17 < 2; fi17++ {
                    fmt.Println(v7)
                    v14 = float32(-2.25)
                    var v18 uint16 = 32767
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
        case "\n\t":
            for fi19 := 0; fi19 < 12; fi19++ {
                if v7 {
                    v14 = ((float32(float64(v14)) / v14) + v14)
                    delete(v2, (fi12 % ((fi19) | 1)))
                    if true {
                        v20 := GenS1{F0: uint16(v13), F1: 65535}
                        v1 += (v11 - (-(v11)))
                        v21 := GenS2{F0: 787, F1: 42, F2: 0}
                        _ = v20
                        _ = v21
                    } else {
                        if (42 <= 42) {
                            v13 -= 0
                        }
                        v13 -= v13
                    }
                }
                v11 -= (v1 ^ v1)
                v22 := GenS1{F0: 261, F1: 2147483647}
                _ = fi19
                _ = v22
            }
        default:
            if ((0.1 - 0.1) < 0.1) {
                for fi23 := 0; fi23 < 2; fi23++ {
                    fmt.Println(v11, v13)
                    v24 := GenS1{F0: genFn0(v7), F1: 1}
                    _ = fi23
                    _ = v24
                }
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v25 float32 = 0.1
    if v7 {
        fmt.Println(v7)
    }
    v26 := []int{}
    var v27 int64 = 9223372036854775807
    if false {
        fmt.Println(v25)
        switch (("" + "foo bar") + "foo bar") {
        case "":
            if ((1 <= -1) && v7) {
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 uint64 = 1
                    _ = fi28
                    _ = v29
                }
            } else {
                v7 = (-706 >= 32768)
                var v30 float64 = (float64(3.14159) / 1.5)
                _ = v30
            }
        case "\n\t":
            var v31 float64 = 3.14159
            _ = v31
        }
    } else {
        v25 += -2.25
        if (genFn0(v7) < genFn0(((false && v7) || true))) {
            switch "foo bar" {
            case "\n\t":
                v32 := GenS3{F0: 2147483647}
                _ = v32
            default:
                var v33 int = 0
                _ = v33
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v27, v27, v11)
            case "hello":
                v1 = (v11 & -27)
            }
        }
    }
    fmt.Println(v27, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v7
}
