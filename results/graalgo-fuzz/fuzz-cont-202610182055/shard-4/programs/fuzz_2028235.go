package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint8; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 bool; F2 int32 }
type GenS5 struct { F0 string; F1 int64 }
func genFn0(p0 bool, p1 uint64) int32 {
    v4 := GenS1{F0: (p1 | (^(p1)))}
    p1 = (p1 ^ p1)
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return -1
}

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    if true {
        var v5 string = "a"
        v1 += v1
        v6 := []int{128, 0}
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: 1}
        _ = v7
    }
    var v8 float32 = 0.0
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v8)
        v1 -= (v1 / ((v1) | 1))
        switch "hello" {
        case "a":
            v8 = (v8 - v8)
        case "hello":
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := GenS1{F0: 1}
                if false {
                    v8 = v8
                    v1 = (-(v1))
                    fmt.Println(fi9, fi10)
                    fmt.Println(fi10)
                }
                v1 = (v1 - -1)
                _ = fi10
                _ = v11
            }
        default:
            v12 := GenS1{F0: 42}
            _ = v12
        }
        _ = fi9
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v1 = -128
        if (!true) {
            var v14 float64 = 0.0
            for fi15 := 0; fi15 < 1; fi15++ {
                _ = len(v2)
                if (764 == 809) {
                    fmt.Println(fi13)
                    var v16 float32 = (v8 - v8)
                    _ = v16
                } else {
                    v17 := GenS1{F0: 1}
                    _ = v17
                }
                _ = fi15
            }
            _ = v14
        } else {
            for fi18 := 0; fi18 < 12; fi18++ {
                if (2 < 175) {
                    switch 32767 {
                    case 42:
                        var v19 int32 = 0
                        _ = v19
                    default:
                        delete(v3, 32768)
                    }
                    var v20 uint8 = 217
                    _ = v20
                }
                fmt.Println(v8)
                v8 = 3.14159
                _ = fi18
            }
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            if false {
                var v22 int = -9223372036854775808
                var v23 uint = 0
                _ = v22
                _ = v23
            }
            v1 += int8(fi13)
            switch fi21 {
            case 42:
                v24 := v2[2147483647]
                _ = v24
            case 595:
                v25 := GenS1{F0: 1}
                if true {
                    fmt.Println(v8, fi21)
                    delete(v2, 0)
                    v26 := []int{}
                    _ = v26
                } else {
                    v27 := []int{fi21, 4294967295, 0, (-(fi13))}
                    _ = v27
                }
                switch fi13 {
                case -18:
                    v28 := GenS2{F0: ((-1 > -957) || false), F1: 2, F2: ("a" + ("" + ("hello" + "foo bar")))}
                    _ = v28
                default:
                    if false {
                        _ = len(v3)
                    }
                }
                var v29 bool = (2 <= 497)
                _ = v25
                _ = v29
            default:
                fmt.Println(v1, v8, fi21)
            }
            _ = fi21
        }
        _ = fi13
    }
    v1 += (^(v1))
    v1 = v1
    switch "hello" {
    case "hello":
        v1 = (v1 & (v1 << ((834) & 63)))
    case "\n\t":
        v30 := []int{0, -566}
        _ = v30
    case "foo bar":
        var v31 float64 = (3.14159 - (1e10 * ((-2.25 * 1.5) * 0.1)))
        _ = v31
    default:
        var v32 int16 = 974
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        v34 := GenS3{F0: false}
        var v35 float32 = (((0.0 * 1e10) + v8) * (v8 - (0.0 * v8)))
        var v36 uint = 1
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    if false {
        for fi37 := 0; fi37 < 5; fi37++ {
            v38 := GenS4{F0: 9223372036854775807, F1: false, F2: 366}
            _ = fi37
            _ = v38
        }
    } else {
        var v39 float32 = v8
        v3[-1] = "\n\t"
        v40 := GenS5{F0: "\n\t", F1: 42}
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 3; fi41++ {
        v3[fi41] = "a"
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
