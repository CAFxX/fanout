package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 bool; F2 string }
type GenS5 struct { F0 uint64; F1 int; F2 uint16 }
type GenS6 struct { F0 int; F1 uint16; F2 uint }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint16) uint16 {
    p0 = (^(p0))
    v4 := []int{128, -295, -2147483648}
    v5 := GenS1{F0: 426, F1: 4294967295, F2: 65535}
    _ = p0
    _ = v4
    _ = v5
    return (p0 << ((0) & 63))
}
func genFn1(p0 int) int8 {
    p0 = p0
    p0 = p0
    if (!(719 <= 651)) {
        for fi6 := 0; fi6 < 11; fi6++ {
            switch "" {
            case "":
                p0 -= fi6
            case "\n\t":
                v7 := GenS2{F0: 2147483647, F1: genFn0(1), F2: 1}
                _ = v7
            case "a":
                var v8 uint16 = 197
                _ = v8
            default:
                fmt.Println(fi6)
            }
            if ((1e10 / (float64(1.5) - 0.0)) > 1.5) {
                v9 := []int{}
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 float32 = -2.25
                    _ = fi10
                    _ = v11
                }
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = v9
            } else {
                if (669 < 2) {
                    if (1e10 != 0.0) {
                        var v12 int64 = 0
                        _ = v12
                    } else {
                        v13 := []int{fi6, (p0 % (((fi6 | p0)) | 1)), (fi6 % ((9223372036854775807) | 1))}
                        v14 := GenS1{F0: 259, F1: 1, F2: 965}
                        if true {
                            var v15 int16 = -1
                            _ = v15
                        } else {
                            v16 := GenS3{F0: 4294967295}
                            _ = v16
                        }
                        _ = v13
                        _ = v14
                    }
                } else {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        fmt.Println(fi17)
                        var v18 float64 = -2.25
                        _ = fi17
                        _ = v18
                    }
                    p0 = p0
                }
                switch "a" {
                case "":
                    switch "hello" {
                    case "foo bar":
                        p0 = (fi6 / (((^(fi6))) | 1))
                    }
                case "foo bar":
                    v19 := GenS4{F0: 4294967295, F1: false, F2: "hello"}
                    _ = v19
                default:
                    p0 = -581
                }
                switch "\n\t" {
                case "hello":
                    var v20 uint16 = 2
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 6; fi21++ {
                        p0 -= 0
                        _ = fi21
                    }
                }
            }
            p0 = (fi6 - (p0 + (^(p0))))
            _ = fi6
        }
        p0 = p0
        var v22 string = ("hello" + ("hello" + "foo bar"))
        _ = v22
    } else {
        fmt.Println(p0)
    }
    _ = p0
    return -128
}

func main() {
    var v1 int32 = -750
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v3[9223372036854775807] = 9223372036854775807
    switch 152 {
    case 1:
        v23 := GenS5{F0: 42, F1: -1, F2: 32767}
        _ = v23
    default:
        v1 = (v1 * 42)
    }
    fmt.Println(v1)
    var v24 float32 = 3.14159
    fmt.Println(v1, v1)
    var v25 uint = 908
    switch "" {
    case "foo bar":
        if true {
            var v26 int64 = 42
            v27 := GenS6{F0: 519, F1: 731, F2: (v25 & v25)}
            switch 127 {
            case 0:
                switch 32767 {
                case -9223372036854775808:
                    v26 = v26
                default:
                    var v28 uint8 = 42
                    _ = v28
                }
            case -527:
                var v29 string = (("foo bar" + "") + "a")
                _ = v29
            default:
                var v30 uint16 = genFn0(65535)
                _ = v30
            }
            _ = v26
            _ = v27
        } else {
            var v31 bool = (-677 >= 9223372036854775807)
            _ = v31
        }
    case "":
        v1 = (v1 ^ (v1 % ((1) | 1)))
    case "\n\t":
        var v32 float32 = 0.0
        _ = v32
    }
    switch -348 {
    case 1:
        var v33 int32 = (-(v1))
        _ = v33
    case 128:
        delete(v2, "a")
    case 42:
        if true {
            delete(v2, ("hello" + "\n\t"))
        } else {
            v34 := GenS7{F0: 0, F1: 42}
            _ = v34
        }
    default:
        v35 := []int{-753, 788, 2147483647, -390}
        _ = v35
    }
    if false {
        v36 := GenS7{F0: -104, F1: 127}
        for fi37 := 0; fi37 < 16; fi37++ {
            if (!("hello" != "hello")) {
                for fi38 := 0; fi38 < 13; fi38++ {
                    v39 := GenS2{F0: 1, F1: 1, F2: v25}
                    _ = fi38
                    _ = v39
                }
                fmt.Println(v24)
                fmt.Println(fi37, v1, v24)
            }
            fmt.Println(v24)
            _ = fi37
        }
        _ = v36
    } else {
        _ = len(v3)
    }
    v40 := GenS8{F0: -128, F1: genFn0(65535)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v40
}
