package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int8) uint64 {
    v6 := []int{}
    var v7 float64 = 1.5
    var v8 int64 = 693
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 215
}
func genFn1() uint32 {
    var v9 bool = true
    switch ("hello" + ((("hello" + "hello") + (("a" + ("hello" + "")) + ((("foo bar" + ("foo bar" + "a")) + "foo bar") + "hello"))) + ("" + ""))) {
    case "":
        v10 := GenS1{F0: 755, F1: "foo bar"}
        _ = v10
    default:
        for fi11 := 0; fi11 < 6; fi11++ {
            v9 = v9
            if v9 {
                v9 = (506 != 4294967295)
                var v12 uint16 = 2
                _ = v12
            }
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 int32 = 114
                _ = fi13
                _ = v14
            }
            _ = fi11
        }
    }
    v9 = (65535 > 1)
    _ = v9
    return 929
}

func main() {
    var v1 int8 = -20
    v2 := make(map[uint32]int)
    var v3 int = -831
    v4 := []uint32{65535, 42}
    v5 := []int{(^(v3)), 1, (v3 / ((127) | 1))}
    for fi15 := 0; fi15 < 12; fi15++ {
        _ = len(v2)
        fmt.Println(v1)
        for fi16 := 0; fi16 < 5; fi16++ {
            fmt.Println(fi16, fi15, v3)
            _ = fi16
        }
        _ = fi15
    }
    v3 += (v3 % ((v3) | 1))
    switch "\n\t" {
    case "\n\t":
        var v17 float64 = ((0.0 * (1.5 * 3.14159)) * 0.0)
        _ = v17
    case "":
        for fi18 := 0; fi18 < 3; fi18++ {
            if ((v1 % (((v1 & v1)) | 1)) < (v1 << ((881) & 63))) {
                if false {
                    var v19 uint = 899
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 int = 454
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
            } else {
                if (false && true) {
                    switch 32768 {
                    case -2147483648:
                        var v22 int8 = v1
                        _ = v22
                    default:
                        var v23 uint64 = genFn0(v1)
                        _ = v23
                    }
                    v1 = (v1 - -128)
                    if true {
                        v3 = 127
                        v24 := GenS2{F0: 65535}
                        _ = v24
                    }
                } else {
                    for fi25 := 0; fi25 < 10; fi25++ {
                        var v26 int16 = 42
                        var v27 uint16 = 65535
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    v1 -= (v1 & v1)
                }
                fmt.Println(v3, v1)
            }
            var v28 uint16 = 255
            _ = fi18
            _ = v28
        }
    case "foo bar":
        v1 -= v1
    default:
        if true {
            var v29 int = 1
            _ = v29
        } else {
            v30 := GenS3{F0: -29}
            fmt.Println(v3, v3)
            _ = v30
        }
    }
    if true {
        if false {
            var v31 uint64 = 2
            _ = v31
        }
    }
    fmt.Println(v1, v3)
    _ = len(v4)
    if (42 > 112) {
        v1 -= v1
        v3 = (v3 ^ -172)
        v3 += v3
    }
    var v32 uint32 = 0
    v32 += genFn1()
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
