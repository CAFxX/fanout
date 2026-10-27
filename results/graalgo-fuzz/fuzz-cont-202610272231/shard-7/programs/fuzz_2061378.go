package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint; F1 string; F2 uint32 }
func genFn0(p0 int) int64 {
    v5 := GenS1{F0: 1, F1: "foo bar", F2: 294}
    var v6 bool = ((-1 >= 0) && false)
    switch p0 {
    case 4294967295:
        switch (("\n\t" + "hello") + ((("\n\t" + "") + "foo bar") + "\n\t")) {
        case "a":
            if false {
                v7 := []int{p0, (p0 ^ 4294967295), p0}
                fmt.Println(v6, v6)
                var v8 bool = ((1e10 >= 1e10) && (365 == 42))
                _ = v7
                _ = v8
            }
        }
    case 0:
        v9 := []int{32767, p0, (-(p0)), (p0 & p0)}
        _ = v9
    case 1:
        v10 := GenS1{F0: 42, F1: "\n\t", F2: 335}
        _ = v10
    default:
        var v11 string = "a"
        _ = v11
    }
    v6 = v6
    _ = p0
    _ = v5
    _ = v6
    return 986
}
func genFn1(p0 bool, p1 uint16, p2 int64) int16 {
    switch ("hello" + (("a" + ("foo bar" + "\n\t")) + ("\n\t" + ("hello" + ("" + "foo bar"))))) {
    case "":
        fmt.Println(p0)
    case "hello":
        v12 := GenS1{F0: genFn0(42), F1: "", F2: p1}
        _ = v12
    case "a":
        v13 := GenS1{F0: 9223372036854775807, F1: ("foo bar" + "foo bar"), F2: p1}
        _ = v13
    default:
        if false {
            v14 := GenS1{F0: genFn0(-2147483648), F1: "", F2: (p1 << ((819) & 63))}
            var v15 uint64 = 940
            _ = v14
            _ = v15
        } else {
            p2 = (p2 * (p2 / (((p2 >> ((0) & 63))) | 1)))
        }
    }
    p1 = (p1 - 255)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    var v4 uint8 = 56
    v4 += 42
    for fi16 := 0; fi16 < 14; fi16++ {
        v4 = v4
        switch "" {
        case "hello":
            v4 = v4
        case "a":
            for fi17 := 0; fi17 < 18; fi17++ {
                if (255 == 65535) {
                    v4 = (^(v4))
                } else {
                    if false {
                        if (-1 != 1) {
                            var v18 int64 = 1
                            var v19 uint32 = 538
                            var v20 uint = 1
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    }
                    if (("\n\t" != "hello") && false) {
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(fi16, fi17)
                        case "a":
                            var v21 uint16 = 32767
                            _ = v21
                        case "\n\t":
                            fmt.Println(fi16, fi16)
                        }
                    } else {
                        if true {
                            v22 := GenS1{F0: 42, F1: "", F2: 0}
                            v4 += (^(v4))
                            v23 := GenS2{F0: 877, F1: 638, F2: 736}
                            _ = v22
                            _ = v23
                        } else {
                            v4 = v4
                        }
                        var v24 int8 = 17
                        _ = v24
                    }
                    v25 := GenS3{F0: 18446744073709551615, F1: ("hello" + ""), F2: 248}
                    _ = v25
                }
                v26 := []int{}
                _ = fi17
                _ = v26
            }
        case "foo bar":
            v4 -= (v4 & 28)
        }
        var v27 uint8 = (v4 / ((v4) | 1))
        _ = fi16
        _ = v27
    }
    v28 := []int{458, -1}
    var v29 int8 = -128
    _ = len(v28)
    for fi30 := 0; fi30 < 10; fi30++ {
        if false {
            fmt.Println(fi30, v4, v29)
            v4 -= (v4 | (^(v4)))
            fmt.Println(v4, v29, v4)
        } else {
            delete(v1, 2147483647)
            fmt.Println(v4)
        }
        _ = fi30
    }
    fmt.Println(v29, v29)
    v4 += 77
    fmt.Println(v4, v4)
    var v31 float64 = ((3.14159 + 0.1) * 0.0)
    v29 = (v29 >> ((65535) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
}
