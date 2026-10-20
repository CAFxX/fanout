package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint32; F2 int16 }
type GenS4 struct { F0 bool; F1 uint; F2 uint }
type GenS5 struct { F0 string; F1 int }
func genFn0(p0 string, p1 int) int32 {
    v4 := GenS1{F0: 32, F1: (2 >= 42)}
    v5 := []int{2147483647, 42, p1, (p1 - (p1 / (((p1 | (p1 - 42))) | 1)))}
    fmt.Println(p0)
    v6 := GenS1{F0: 127, F1: false}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -566
}
func genFn1(p0 uint8, p1 uint64) uint64 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 52, F1: (2147483647 > 713)}
        switch 142 {
        case 128:
            p0 += (p0 & p0)
        case -1:
            v9 := GenS2{F0: fi7, F1: 782}
            _ = v9
        case 528:
            p1 += p1
        default:
            var v10 bool = (false && false)
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint16 = 959
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := []int{(-(fi13)), fi13}
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    return uint64(p0)
}

func main() {
    v1 := []int{-339, -958, 1, 0}
    var v2 int16 = 1000
    v3 := make(map[uint32]string)
    v2 = (v2 / ((v2) | 1))
    v2 = (v2 / (((v2 << ((1) & 63))) | 1))
    var v15 float32 = (0.1 - 3.14159)
    if (!true) {
        fmt.Println(v2, v15)
        if true {
            switch "foo bar" {
            case "":
                v16 := GenS1{F0: 243, F1: (!(true && false))}
                _ = v16
            case "a":
                v17 := GenS2{F0: -9223372036854775808, F1: 955}
                _ = v17
            default:
                fmt.Println(v2, v2)
            }
            var v18 float64 = float64((1.5 / v15))
            _ = v18
        } else {
            v19 := GenS3{F0: 1, F1: 2147483647, F2: 200}
            _ = v19
        }
        switch "hello" {
        case "":
            var v20 uint = 4294967295
            _ = v20
        case "hello":
            fmt.Println(v15, v2)
        case "foo bar":
            if (!("\n\t" > (("\n\t" + ("a" + "")) + "a"))) {
                v15 = v15
                fmt.Println(v2, v2)
            } else {
                switch "a" {
                case "hello":
                    v2 += (v2 + (v2 / ((-32768) | 1)))
                case "":
                    for fi21 := 0; fi21 < 8; fi21++ {
                        _ = len(v3)
                        var v22 float32 = (0.0 / v15)
                        _ = fi21
                        _ = v22
                    }
                default:
                    v15 = -2.25
                }
                v3[0] = "a"
                v23 := v3[65535]
                _ = v23
            }
        default:
            var v24 uint16 = 42
            _ = v24
        }
    }
    if true {
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 int16 = v2
            _ = fi25
            _ = v26
        }
        v27 := GenS4{F0: true, F1: 144, F2: 1}
        for fi28 := 0; fi28 < 11; fi28++ {
            v29 := v3[107]
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    v15 = (v15 + float32(float64((v15 - ((float32(float64(v15)) + v15) * (1e10 - v15))))))
    switch "a" {
    case "hello":
        v30 := GenS1{F0: 1, F1: true}
        _ = v30
    case "a":
        switch "hello" {
        case "a":
            v2 = (v2 | (v2 + 42))
        case "\n\t":
            v31 := GenS2{F0: 2147483647, F1: 930}
            _ = v31
        }
    case "foo bar":
        v3[4294967295] = ("a" + ((("a" + "") + ("\n\t" + (("\n\t" + "") + "\n\t"))) + "\n\t"))
    }
    if (((!("" == "\n\t")) || (0 <= 235)) || false) {
        v32 := GenS5{F0: "foo bar", F1: 0}
        for fi33 := 0; fi33 < 5; fi33++ {
            if (710 < 42) {
                var v34 int = (fi33 % (((fi33 + (^(fi33)))) | 1))
                switch "a" {
                case "hello":
                    if true {
                        var v35 uint = 255
                        _ = v35
                    } else {
                        v2 = (v2 / ((v2) | 1))
                        v36 := GenS2{F0: v34, F1: 4294967295}
                        _ = v36
                    }
                case "foo bar":
                    var v37 uint8 = 42
                    _ = v37
                default:
                    v15 = v15
                }
                var v38 int = v34
                _ = v34
                _ = v38
            }
            _ = fi33
        }
        _ = v32
    } else {
        v2 += (v2 % (((v2 ^ 794)) | 1))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
}
