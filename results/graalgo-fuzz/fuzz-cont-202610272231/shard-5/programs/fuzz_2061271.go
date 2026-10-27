package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 uint32, p1 int8) bool {
    v5 := []int{}
    var v6 uint32 = (^(p0))
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return true
}
func genFn1() int64 {
    var v7 uint = 1
    var v8 int = -294
    _ = v7
    _ = v8
    return -26
}

func main() {
    var v1 int8 = -128
    v2 := []string{"foo bar", ("hello" + ("foo bar" + "foo bar")), "\n\t"}
    v3 := make(map[int]string)
    var v4 int16 = -121
    _ = len(v2)
    var v9 int32 = 42
    for fi10 := 0; fi10 < 15; fi10++ {
        if (379 == 18446744073709551615) {
            v11 := GenS1{F0: -844, F1: 18446744073709551615}
            var v12 float32 = -2.25
            v1 = v1
            _ = v11
            _ = v12
        } else {
            var v13 float32 = 1e10
            var v14 int = (fi10 >> ((954) & 63))
            var v15 float32 = ((v13 * (0.0 / 1.0)) - v13)
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi10
    }
    v1 = 127
    var v16 int = 127
    fmt.Println(v9)
    fmt.Println(v4)
    switch (v16 / ((128) | 1)) {
    case 9223372036854775807:
        for fi17 := 0; fi17 < 20; fi17++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v18 int32 = (v9 ^ v9)
            if (false && false) {
                fmt.Println(v9, v1, v9)
            } else {
                switch (^(fi17)) {
                case 1:
                    v19 := GenS2{F0: v16, F1: 0, F2: 2}
                    _ = v19
                case 4294967295:
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v16 = 42
                        v21 := GenS3{F0: ""}
                        v18 -= (v9 + 0)
                        _ = fi20
                        _ = v21
                    }
                case 895:
                    if len(v2) > 0 {
                        v2[0] = ("a" + "")
                    }
                default:
                    for fi22 := 0; fi22 < 18; fi22++ {
                        switch "a" {
                        case "foo bar":
                            switch 32767 {
                            case 379:
                                v23 := GenS1{F0: 42, F1: 18446744073709551615}
                                _ = v23
                            case -320:
                                var v24 uint32 = 175
                                _ = v24
                            case 32768:
                                delete(v3, (fi22 | (-(v16))))
                            default:
                                v1 -= 100
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        v16 = fi22
                        _ = fi22
                    }
                }
                var v25 int = v16
                _ = v25
            }
            _ = fi17
            _ = v18
        }
    case -1:
        v26 := GenS2{F0: -1, F1: 76, F2: 2147483648}
        _ = v26
    }
    fmt.Println(v9)
    fmt.Println(v4, v9)
    for fi27 := 0; fi27 < 7; fi27++ {
        _ = len(v2)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
