package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int64; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint32; F1 uint32; F2 int8 }
func genFn0() uint8 {
    v6 := []int{0}
    v7 := GenS1{F0: 68, F1: (false && true), F2: 18446744073709551615}
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int, p1 uint32, p2 int8) int8 {
    v8 := GenS1{F0: 1000, F1: true, F2: 1}
    p0 = (p0 | p0)
    p2 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p2
}

func main() {
    var v1 uint = 4294967295
    v2 := []string{"a", "foo bar"}
    v3 := []int{32768}
    v4 := make(map[string]string)
    v5 := make(map[string]string)
    delete(v5, ("a" + "hello"))
    switch 586 {
    case -1:
        _ = len(v2)
    case 42:
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(v1, v1)
            v10 := GenS2{F0: 172}
            for fi11 := 0; fi11 < 2; fi11++ {
                if false {
                    if true {
                        var v12 float32 = (-2.25 * 1e10)
                        v12 -= v12
                        _ = v12
                    } else {
                        _ = len(v4)
                        v1 = (v1 ^ (^(v1)))
                        fmt.Println(fi9, fi9, fi11)
                    }
                    v1 = v1
                    v1 = (v1 / ((v1) | 1))
                } else {
                    var v13 bool = (!true)
                    var v14 uint32 = 42
                    _ = v13
                    _ = v14
                }
                if false {
                    _ = len(v5)
                } else {
                    v1 = (^(v1))
                }
                v15 := GenS3{F0: fi9, F1: true}
                if ((-110 >= 42) && ((9223372036854775807 < 1) || (65535 >= 2))) {
                    v1 -= (v1 ^ v1)
                    v1 = 42
                } else {
                    delete(v5, "foo bar")
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi11
                _ = v15
            }
            _ = fi9
            _ = v10
        }
    case 2147483647:
        v16 := GenS2{F0: 255}
        _ = v16
    }
    if (42 == 255) {
        v1 = v1
        if (!false) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS4{F0: 1, F1: 267, F2: 888}
            var v19 int16 = 876
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        if (!true) {
            if (128 <= -798) {
                fmt.Println(v1)
                v5["\n\t"] = "a"
                v1 = (v1 | v1)
            } else {
                v1 -= (v1 % (((v1 ^ v1)) | 1))
                fmt.Println(v1)
            }
            if ((0.0 >= 1e10) || false) {
                var v20 bool = true
                _ = v20
            }
            v21 := GenS5{F0: 2, F1: 2147483648}
            _ = v21
        } else {
            v1 = (v1 + v1)
        }
    }
    v1 += 77
    delete(v5, "hello")
    var v22 int = 9223372036854775807
    v22 = 76
    v23 := GenS6{F0: 2, F1: 2147483648, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
