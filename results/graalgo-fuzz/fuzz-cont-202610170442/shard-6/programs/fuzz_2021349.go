package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint32, p1 uint32) bool {
    var v5 bool = ((4294967295 >= 908) || true)
    p1 -= (^(p1))
    p1 = (^(p0))
    p1 = 1
    _ = p0
    _ = p1
    _ = v5
    return v5
}
func genFn1(p0 int32, p1 int16, p2 int16) uint8 {
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    p2 = p1
    fmt.Println(p1)
    switch 0 {
    case -1:
        switch "hello" {
        case "":
            fmt.Println(p0)
        default:
            switch "a" {
            case "\n\t":
                switch 128 {
                case 0:
                    var v7 bool = ((18446744073709551615 < 0) && false)
                    _ = v7
                }
            case "hello":
                v8 := []int{-645}
                _ = v8
            }
        }
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if true {
            var v10 int = -655
            fmt.Println(fi9, p1)
            switch (("\n\t" + "\n\t") + "\n\t") {
            case "foo bar":
                v11 := []int{fi9, 9223372036854775807, (v10 % (((fi9 ^ (v10 + 4294967295))) | 1))}
                _ = v11
            case "":
                v12 := GenS1{F0: true, F1: p0}
                _ = v12
            }
            _ = v10
        } else {
            v13 := []int{-9223372036854775808, 2147483647, fi9, (fi9 >> ((253) & 63))}
            if (!false) {
                _ = len(v13)
            }
            _ = v13
        }
        v14 := GenS2{F0: -128, F1: 683}
        _ = fi9
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 int32 = -1
    var v2 int32 = v1
    v3 := make(map[int]int)
    var v4 float32 = (-2.25 - 0.1)
    v15 := GenS3{F0: "foo bar"}
    v4 = 1e10
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := []int{}
        _ = fi16
        _ = v17
    }
    fmt.Println(v1, v4, v1)
    v18 := GenS2{F0: 28, F1: 2147483647}
    v4 = ((float32(3.14159) + float32(1.5)) * v4)
    for fi19 := 0; fi19 < 7; fi19++ {
        switch (fi19 - fi19) {
        case 2147483647:
            fmt.Println(v4)
        case 128:
            v20 := []int{}
            _ = v20
        case 32767:
            v4 = 3.14159
        }
        var v21 int16 = -83
        _ = fi19
        _ = v21
    }
    fmt.Println(v4, v4, v1)
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := v3[fi22]
        _ = len(v3)
        _ = fi22
        _ = v23
    }
    v4 = float32(3.14159)
    if (!false) {
        v24 := []int{}
        _ = v24
    } else {
        var v25 float64 = (((float64(v4) / 1.0) - 0.1) - (0.0 / 1.0))
        if (2 > 0) {
            var v26 float64 = v25
            _ = v26
        }
        _ = v25
    }
    v1 = (v2 / ((641) | 1))
    v27 := []int{}
    v28 := v3[2147483647]
    for fi29 := 0; fi29 < 6; fi29++ {
        v30 := GenS4{F0: 42, F1: 1}
        _ = fi29
        _ = v30
    }
    if false {
        v4 -= ((0.0 / v4) * float32(float64(((v4 / 1.0) / (float32(1e10) / v4)))))
        if (182 == 172) {
            if (!false) {
                for fi31 := 0; fi31 < 1; fi31++ {
                    if (false || (67 >= 273)) {
                        fmt.Println(v28)
                        fmt.Println(v4, fi31)
                    }
                    if false {
                        fmt.Println(v2, v2)
                        v1 -= 0
                    } else {
                        var v32 uint32 = 2147483648
                        fmt.Println(v32, v32)
                        _ = v32
                    }
                    var v33 bool = (!(true || false))
                    _ = fi31
                    _ = v33
                }
            } else {
                var v34 string = "\n\t"
                for fi35 := 0; fi35 < 15; fi35++ {
                    var v36 string = "\n\t"
                    for fi37 := 0; fi37 < 18; fi37++ {
                        var v38 int = fi37
                        v3[fi35] = (fi37 % ((v28) | 1))
                        v2 = -234
                        _ = fi37
                        _ = v38
                    }
                    _ = fi35
                    _ = v36
                }
                _ = v34
            }
        } else {
            _ = len(v27)
            fmt.Println(v4, v4, v4)
            v1 = v2
        }
    } else {
        if len(v27) > 0 {
            v27[0] = (^(v28))
        }
        v28 += v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
