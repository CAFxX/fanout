package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int16 }
type GenS2 struct { F0 int8; F1 string; F2 int32 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint32 }
func genFn0() uint32 {
    var v7 float64 = 0.0
    if ((true && false) && true) {
        switch "foo bar" {
        case "a":
            if false {
                if (-426 >= -489) {
                    var v8 uint16 = 549
                    v9 := GenS1{F0: -2147483648, F1: (("foo bar" + "a") + ""), F2: -32768}
                    _ = v8
                    _ = v9
                }
                fmt.Println(v7)
            } else {
                if (1e10 <= float32(v7)) {
                    v7 = 1e10
                    v7 = v7
                } else {
                    v7 -= 0.0
                    if true {
                        v7 -= 3.14159
                        var v10 int32 = 0
                        _ = v10
                    }
                    v11 := GenS2{F0: -65, F1: ("hello" + ""), F2: -2147483648}
                    _ = v11
                }
            }
        }
        var v12 int = 346
        v12 = v12
        _ = v12
    }
    _ = v7
    return 0
}
func genFn1(p0 bool, p1 uint) uint32 {
    switch "\n\t" {
    case "foo bar":
        if p0 {
            for fi13 := 0; fi13 < 7; fi13++ {
                for fi14 := 0; fi14 < 8; fi14++ {
                    if (1 < 55) {
                        var v15 uint32 = 732
                        p0 = (234 != 128)
                        p1 = (p1 - 18446744073709551615)
                        _ = v15
                    } else {
                        var v16 float32 = (1e10 + 3.14159)
                        v17 := GenS1{F0: 42, F1: "hello", F2: -78}
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS3{F0: ((fi14 - fi14) != fi14), F1: 296, F2: 124}
                    _ = fi14
                    _ = v18
                }
                var v19 int32 = 0
                var v20 uint16 = 255
                _ = fi13
                _ = v19
                _ = v20
            }
            switch ("foo bar" + "a") {
            case "foo bar":
                v21 := GenS2{F0: 127, F1: "", F2: 42}
                _ = v21
            case "hello":
                var v22 uint8 = 255
                _ = v22
            default:
                var v23 uint32 = genFn0()
                _ = v23
            }
            v24 := []int{-128, -459, 9223372036854775807, -1}
            _ = v24
        }
    }
    switch 128 {
    case 1:
        v25 := GenS2{F0: -128, F1: ("a" + "\n\t"), F2: 0}
        _ = v25
    case 626:
        v26 := GenS4{F0: 74}
        _ = v26
    }
    if (-2.25 != 0.0) {
        p0 = (9223372036854775807 >= 128)
        switch 127 {
        case 0:
            switch 4294967295 {
            case 32768:
                var v27 uint = p1
                _ = v27
            case 1:
                v28 := []int{4294967295, 887, 9223372036854775807, 4294967295}
                _ = v28
            case -1:
                if true {
                    if p0 {
                        for fi29 := 0; fi29 < 4; fi29++ {
                            switch fi29 {
                            case 32768:
                                p1 = (p1 & (p1 << ((255) & 63)))
                            case 587:
                                switch fi29 {
                                case 758:
                                    p1 = (p1 % (((p1 ^ 61)) | 1))
                                }
                            }
                            var v30 int8 = -128
                            _ = fi29
                            _ = v30
                        }
                    }
                    for fi31 := 0; fi31 < 4; fi31++ {
                        v32 := []int{(fi31 >> ((2) & 63)), fi31, fi31, fi31}
                        v32 = append(v32, (fi31 % ((-1) | 1)))
                        _ = fi31
                        _ = v32
                    }
                } else {
                    p0 = p0
                    v33 := []int{133, 42, 42, 678}
                    _ = v33
                }
            default:
                var v34 uint16 = 0
                _ = v34
            }
        case 32767:
            var v35 uint16 = 1
            _ = v35
        }
        v36 := GenS1{F0: 1, F1: "\n\t", F2: 1000}
        _ = v36
    }
    v37 := []int{0}
    switch 2147483647 {
    case 42:
        if len(v37) > 0 {
            v37[0] = 4294967295
        }
    case 9223372036854775807:
        if false {
            p1 = p1
            fmt.Println(p1)
        }
    }
    _ = p0
    _ = p1
    _ = v37
    return 514
}

func main() {
    var v1 int16 = 42
    v2 := make(map[int]int)
    var v3 int8 = 0
    var v4 float32 = ((-2.25 + 0.0) * (float32((1.5 - (1e10 * 3.14159))) * 1e10))
    var v5 float32 = v4
    v6 := []string{}
    v38 := GenS3{F0: true, F1: 0, F2: 32767}
    v4 = (0.0 / 0.1)
    switch 1 {
    case 128:
        v5 = (1.5 / v4)
    }
    var v39 uint32 = 637
    v40 := GenS4{F0: 2}
    v41 := GenS5{F0: 255, F1: v39}
    fmt.Println(v5, v1)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
