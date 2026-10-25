package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int; F2 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 bool; F1 bool; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32 }
func genFn0(p0 string, p1 int16) uint32 {
    var v4 int = 903
    v5 := []int{}
    for fi6 := 0; fi6 < 3; fi6++ {
        if (127 != v4) {
            var v7 uint64 = 2
            if len(v5) > 0 {
                v5[0] = (v4 + (v4 & fi6))
            }
            _ = v7
        } else {
            v8 := GenS1{F0: true, F1: 384}
            p0 = p0
            _ = v8
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        if false {
            p0 = ("a" + ("\n\t" + (p0 + (p0 + p0))))
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1() uint64 {
    var v10 float32 = (1e10 - -2.25)
    v11 := GenS2{F0: 899, F1: -9223372036854775808, F2: 42}
    _ = v10
    _ = v11
    return 825
}

func main() {
    v1 := []string{"hello", "hello"}
    var v2 float64 = (0.1 / 1.0)
    v3 := make(map[int]bool)
    var v12 uint8 = 1
    _ = len(v3)
    if false {
        if (!true) {
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS2{F0: 2147483648, F1: (fi13 / (((^(fi13))) | 1)), F2: 2}
                for fi15 := 0; fi15 < 18; fi15++ {
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    _ = fi15
                }
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 int = fi13
                    v18 := GenS3{F0: 0, F1: 32767, F2: 6}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi13
                _ = v14
            }
            var v19 uint64 = uint64(v12)
            _ = v19
        }
    }
    if (true || true) {
        v12 -= v12
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 uint64 = 451
            v2 = 1.5
            switch "\n\t" {
            case "a":
                switch "\n\t" {
                case "":
                    for fi22 := 0; fi22 < 20; fi22++ {
                        switch int(v12) {
                        case -1:
                            v2 += v2
                        case -243:
                            if true {
                                v23 := GenS3{F0: fi22, F1: 2, F2: 760}
                                v12 = (v12 % (((v12 + (v12 % (((v12 ^ 42)) | 1)))) | 1))
                                v12 = (v12 >> ((0) & 63))
                                var v24 uint8 = (v12 - (v12 / (((v12 >> ((4294967295) & 63))) | 1)))
                                _ = v23
                                _ = v24
                            }
                        default:
                            var v25 int16 = 0
                            _ = v25
                        }
                        if true {
                            v1 = append(v1, "\n\t")
                            var v26 int = (fi22 << ((uint(v21)) & 63))
                            _ = v26
                        } else {
                            v12 = (v12 * (^(v12)))
                        }
                        v27 := GenS4{F0: (false && false), F1: true, F2: 623}
                        _ = fi22
                        _ = v27
                    }
                }
            case "":
                for fi28 := 0; fi28 < 19; fi28++ {
                    v29 := GenS5{F0: -762}
                    v12 -= v12
                    fmt.Println(fi20, fi20)
                    _ = fi28
                    _ = v29
                }
            case "foo bar":
                v30 := GenS6{F0: 255}
                _ = v30
            default:
                fmt.Println(v21, fi20, fi20)
            }
            _ = fi20
            _ = v21
        }
        var v31 bool = (false && (!(2147483647 <= -1)))
        _ = v31
    } else {
        if (!(!false)) {
            for fi32 := 0; fi32 < 15; fi32++ {
                fmt.Println(v12)
                _ = len(v3)
                for fi33 := 0; fi33 < 14; fi33++ {
                    if (-9223372036854775808 != 9223372036854775807) {
                        v34 := GenS7{F0: 989}
                        var v35 bool = false
                        if v35 {
                            fmt.Println(v2)
                            var v36 string = "hello"
                            _ = v36
                        }
                        _ = v34
                        _ = v35
                    } else {
                        v2 -= (1e10 + v2)
                    }
                    _ = fi33
                }
                _ = fi32
            }
            var v37 int32 = -2147483648
            _ = v37
        }
        var v38 float32 = ((3.14159 - 0.1) * 1.5)
        _ = v38
    }
    fmt.Println(v12)
    var v39 int8 = -128
    v2 = v2
    v3[-515] = false
    _ = len(v1)
    var v40 uint = 42
    var v41 uint32 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
}
