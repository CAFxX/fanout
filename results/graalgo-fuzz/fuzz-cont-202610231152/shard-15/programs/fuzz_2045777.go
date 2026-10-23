package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 bool; F2 int32 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int16, p1 uint8, p2 uint) bool {
    v6 := GenS1{F0: 65535}
    var v7 string = ("a" + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (1e10 > float32(float64(0.1)))
}
func genFn1(p0 uint16) int32 {
    p0 = (p0 & 0)
    v8 := GenS1{F0: 0}
    var v9 uint64 = 18446744073709551615
    _ = p0
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{(("hello" + "a") + "a"), ("hello" + "")}
    var v3 int32 = 1
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    var v10 float32 = (-2.25 - -2.25)
    if len(v2) > 0 {
        v2[0] = ("foo bar" + ((("foo bar" + ("\n\t" + "hello")) + "hello") + "\n\t"))
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if (false && (true && (42 == 342))) {
            v12 := GenS1{F0: 2147483647}
            for fi13 := 0; fi13 < 8; fi13++ {
                if false {
                    switch -200 {
                    case 868:
                        v14 := GenS2{F0: "", F1: 4294967295}
                        _ = v14
                    case 0:
                        if ("hello" <= "") {
                            var v15 uint32 = 65535
                            _ = len(v4)
                            v1 = -9223372036854775808
                            switch "foo bar" {
                            case "\n\t":
                                v2 = append(v2, ("" + "foo bar"))
                            case "":
                                v4[(v15 % ((v15) | 1))] = (("a" + ((("a" + "hello") + "\n\t") + "a")) + "")
                            }
                            delete(v5, (fi13 >> ((0) & 63)))
                            _ = v15
                        }
                    case 42:
                        switch "foo bar" {
                        case "a":
                            _ = len(v4)
                        case "foo bar":
                            var v16 string = "\n\t"
                            _ = v16
                        default:
                            switch "hello" {
                            case "\n\t":
                                v3 = (^(v3))
                            default:
                                v4[2147483647] = "a"
                            }
                        }
                    default:
                        if false {
                            var v17 int8 = 0
                            var v18 int8 = v17
                            v17 = (v17 ^ (^(v17)))
                            _ = v17
                            _ = v18
                        }
                    }
                } else {
                    fmt.Println(fi13, v3, v3)
                    if (!false) {
                        if false {
                            v10 = v10
                            delete(v4, 2)
                        }
                    }
                }
                v1 = v1
                switch (-(fi13)) {
                case 152:
                    var v19 uint32 = 863
                    _ = v19
                case -909:
                    fmt.Println(fi13)
                }
                _ = fi13
            }
            _ = v12
        } else {
            v20 := GenS3{F0: 255, F1: 0, F2: 648}
            _ = v20
        }
        v3 += v3
        _ = fi11
    }
    v21 := GenS4{F0: 42}
    if false {
        switch "foo bar" {
        case "foo bar":
            v22 := GenS5{F0: 18446744073709551615, F1: ((v1 >= v1) && true), F2: (v3 << ((0) & 63))}
            _ = v22
        case "\n\t":
            for fi23 := 0; fi23 < 10; fi23++ {
                v1 = v1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v24 := GenS6{F0: (-(v1))}
                _ = fi23
                _ = v24
            }
        case "a":
            v1 -= (^(v1))
        default:
            v25 := GenS3{F0: 0, F1: 42, F2: -247}
            _ = v25
        }
        if (93 <= 42) {
            var v26 float64 = 1e10
            _ = v26
        }
    }
    switch ("a" + "a") {
    case "\n\t":
        v3 = 2147483647
    case "":
        v1 = (-(v1))
    case "hello":
        var v27 uint16 = 454
        _ = v27
    }
    v28 := GenS4{F0: 270}
    var v29 uint32 = 1
    var v30 uint64 = 0
    v1 = (v1 | v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
