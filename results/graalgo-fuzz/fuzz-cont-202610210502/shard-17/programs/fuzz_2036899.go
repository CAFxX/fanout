package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0() uint16 {
    var v6 int32 = -460
    var v7 int64 = 263
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int, p1 bool) int8 {
    v8 := []int{}
    for fi9 := 0; fi9 < 6; fi9++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
    }
    p0 = p0
    _ = len(v8)
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    var v1 uint8 = 128
    var v2 uint16 = 42
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    fmt.Println(v1, v2)
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(v1, fi10, v2)
        v11 := GenS1{F0: (v2 | (^(v2)))}
        var v12 uint64 = 645
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS1{F0: v2}
            _ = fi13
            _ = v14
        }
    }
    if false {
        v15 := GenS1{F0: v2}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 8; fi16++ {
            fmt.Println(v1, v1, v1)
            _ = fi16
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 string = "a"
        _ = fi17
        _ = v18
    }
    switch -17 {
    case 9223372036854775807:
        fmt.Println(v2, v2)
    case 42:
        v1 = (v1 % ((v1) | 1))
    default:
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 string = ("a" + ("a" + "\n\t"))
            _ = len(v4)
            _ = fi19
            _ = v20
        }
    }
    v2 -= 2
    if (1 == 0) {
        var v21 float32 = 0.1
        _ = v21
    } else {
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := []int{(fi22 << ((42) & 63)), -2147483648, 32767, fi22}
            for fi24 := 0; fi24 < 7; fi24++ {
                switch "hello" {
                case "hello":
                    v25 := GenS2{F0: 2147483648, F1: -32768}
                    _ = v25
                case "\n\t":
                    switch "\n\t" {
                    case "a":
                        v26 := v5[("" + (("foo bar" + "") + ""))]
                        _ = v26
                    case "\n\t":
                        if false {
                            v1 = 87
                        } else {
                            fmt.Println(fi24)
                            v2 += (v2 ^ (v2 & genFn0()))
                            var v27 int16 = 1
                            if (3.14159 >= (0.1 + 0.1)) {
                                _ = len(v3)
                                switch 4294967295 {
                                case -2147483648:
                                    v28 := GenS1{F0: (v2 | v2)}
                                    v27 -= 0
                                    _ = v28
                                case 2147483647:
                                    v2 = (^(v2))
                                default:
                                    v2 = (v2 << ((0) & 63))
                                }
                            } else {
                                var v29 uint16 = (v2 & (v2 - (v2 + 1)))
                                v1 = uint8(v2)
                                _ = v29
                            }
                            _ = v27
                        }
                    case "foo bar":
                        var v30 int = (^(fi24))
                        _ = v30
                    }
                case "foo bar":
                    v1 += (v1 >> ((617) & 63))
                }
                var v31 int = (fi22 / ((2147483647) | 1))
                _ = fi24
                _ = v31
            }
            _ = fi22
            _ = v23
        }
        v32 := GenS1{F0: 255}
        v2 -= (v2 % ((genFn0()) | 1))
        _ = v32
    }
    if ((false || (2 < 198)) && (!((10 <= 9) && true))) {
        var v33 uint16 = (v2 ^ 65535)
        var v34 int32 = 519
        _ = v33
        _ = v34
    }
    v2 -= (v2 % ((65535) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
