package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint8; F1 uint16; F2 string }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 bool; F1 bool; F2 uint16 }
func genFn0(p0 uint16) uint32 {
    p0 += p0
    for fi6 := 0; fi6 < 9; fi6++ {
        if (0.1 >= float32(float64(1e10))) {
            p0 = p0
            var v7 int32 = -1
            for fi8 := 0; fi8 < 1; fi8++ {
                if (!false) {
                    v9 := []int{(fi8 | 2147483647), -1, fi8}
                    v10 := GenS1{F0: 422, F1: 94}
                    var v11 uint16 = 342
                    _ = v10
                    _ = v11
                    _ = v9
                }
                _ = fi8
            }
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(p0)
    var v12 uint16 = 61
    switch 633 {
    case 362:
        if (42 == 1) {
            if (0 >= 2147483647) {
                v13 := []int{4294967295, 2147483647, 128, 9223372036854775807}
                _ = v13
            } else {
                if true {
                    v14 := []int{478, 388}
                    if false {
                        _ = len(v14)
                    }
                    p0 += (^(v12))
                    _ = v14
                }
                v12 -= 32767
            }
        } else {
            var v15 int16 = 32767
            _ = v15
        }
    case -9223372036854775808:
        v16 := []int{-541, -2147483648, -70, 127}
        _ = v16
    case -1:
        switch ((("foo bar" + "a") + ((("\n\t" + "hello") + "a") + "")) + "a") {
        case "a":
            p0 = 32767
        case "foo bar":
            v17 := GenS1{F0: 218, F1: 128}
            _ = v17
        default:
            var v18 float64 = 1.5
            _ = v18
        }
    }
    _ = p0
    _ = v12
    return 2147483647
}
func genFn1(p0 uint64, p1 uint64, p2 uint16) uint8 {
    fmt.Println(p0, p0)
    p2 = 1
    v19 := GenS1{F0: -1, F1: 107}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 1
}

func main() {
    var v1 uint64 = 2
    var v2 uint = 1
    var v3 int16 = 1
    v4 := make(map[int]int)
    var v5 uint16 = 146
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 int16 = v3
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v2 = (v2 | (^(v2)))
    if (0 <= 1) {
        if (!(!(0.0 >= (1e10 * (1.5 - (1.5 - (1.5 - 3.14159))))))) {
            v5 = 975
            v23 := GenS2{F0: v2, F1: true}
            _ = v23
        }
        fmt.Println(v1)
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 int8 = 100
            switch 256 {
            case 9223372036854775807:
                if (!false) {
                    if false {
                        v26 := v4[fi24]
                        v27 := GenS3{F0: 127, F1: v5, F2: ("foo bar" + "a")}
                        _ = v26
                        _ = v27
                    }
                    if (true && false) {
                        var v28 int32 = -662
                        _ = v28
                    }
                    if false {
                        v29 := []int{int(v2), -635, (fi24 >> (((v2 + 83)) & 63)), fi24}
                        var v30 int32 = -749
                        _ = v29
                        _ = v30
                    }
                }
            case -1:
                for fi31 := 0; fi31 < 20; fi31++ {
                    if (178 > 42) {
                        var v32 float32 = 0.1
                        switch "foo bar" {
                        case "foo bar":
                            var v33 uint8 = 101
                            _ = v33
                        default:
                            v25 = v25
                        }
                        if ((v5 | v5) > 255) {
                            v32 -= v32
                            v34 := []int{-2147483648, fi24, 921, fi24}
                            _ = v34
                        }
                        _ = v32
                    }
                    var v35 int32 = 2147483647
                    v2 = v2
                    _ = fi31
                    _ = v35
                }
            default:
                v36 := GenS1{F0: 530, F1: genFn1(39, v1, (v5 ^ v5))}
                _ = v36
            }
            _ = fi24
            _ = v25
        }
    } else {
        v37 := GenS4{F0: 0, F1: -128}
        var v38 string = ("foo bar" + ("" + ("foo bar" + "a")))
        v3 = (v3 | v3)
        _ = v37
        _ = v38
    }
    v39 := []int{}
    v39 = append(v39, 560)
    var v40 string = "hello"
    for fi41 := 0; fi41 < 11; fi41++ {
        var v42 float32 = (1e10 - (0.0 / 0.1))
        if true {
            if len(v39) >= 2 {
                v39 = v39[1:len(v39)]
            }
            _ = len(v39)
            fmt.Println(v5, v3, v42)
        }
        _ = fi41
        _ = v42
    }
    v5 = (v5 * v5)
    v43 := GenS5{F0: true, F1: (genFn0((v5 ^ (v5 & 0))) != uint32(v2)), F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v5
}
