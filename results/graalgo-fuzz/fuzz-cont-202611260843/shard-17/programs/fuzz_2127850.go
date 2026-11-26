package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0(p0 int64, p1 uint16, p2 uint) bool {
    var v7 int64 = p0
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 bool = true
        v10 := []int{fi8}
        var v11 uint16 = (p1 | (p1 << ((p2) & 63)))
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := []int{2147483647, 128, -9223372036854775808, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return ((p1 <= p1) || true)
}
func genFn1(p0 uint, p1 uint32) bool {
    if false {
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := []int{-702, fi13, (-(fi13))}
            switch ("foo bar" + "\n\t") {
            case "a":
                switch 774 {
                case 2147483647:
                    _ = len(v14)
                case 0:
                    switch "\n\t" {
                    case "a":
                        p0 += (p0 * 42)
                    case "hello":
                        p0 = (p0 & p0)
                    }
                case -9223372036854775808:
                    var v15 uint = (p0 & p0)
                    _ = v15
                default:
                    var v16 uint32 = (p1 << ((935) & 63))
                    _ = v16
                }
            case "hello":
                switch ("foo bar" + ("" + (("hello" + ("\n\t" + ("\n\t" + "\n\t"))) + ("foo bar" + (("hello" + ("foo bar" + ("\n\t" + "\n\t"))) + "hello"))))) {
                case "\n\t":
                    p1 = (p1 >> ((42) & 63))
                case "foo bar":
                    for fi17 := 0; fi17 < 5; fi17++ {
                        if ("a" < ((((("foo bar" + "a") + "\n\t") + "hello") + "foo bar") + "a")) {
                            v14 = append(v14, fi13)
                            switch "foo bar" {
                            case "a":
                                if (false || true) {
                                    p0 += (p0 * p0)
                                } else {
                                    v14 = append(v14, fi13)
                                    _ = len(v14)
                                }
                            default:
                                fmt.Println(fi13, fi13, fi13)
                            }
                        }
                        p1 -= (p1 | p1)
                        v18 := GenS1{F0: (^(p1)), F1: false, F2: (p1 << (((p0 ^ (p0 >> ((p0) & 63)))) & 63))}
                        _ = fi17
                        _ = v18
                    }
                case "a":
                    p1 -= (p1 - p1)
                default:
                    v19 := GenS2{F0: 1, F1: 1, F2: 1}
                    _ = v19
                }
            case "foo bar":
                if true {
                    var v20 int32 = 175
                    _ = v20
                }
            default:
                if (2147483647 <= int32(p1)) {
                    v21 := GenS3{F0: fi13, F1: 731}
                    p1 -= p1
                    _ = v21
                } else {
                    _ = len(v14)
                }
            }
            _ = fi13
            _ = v14
        }
    }
    p1 += (^(p1))
    v22 := []int{9223372036854775807, -2147483648, -1}
    var v23 uint = p0
    p1 = uint32(p0)
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return (127 != -116)
}

func main() {
    v1 := []string{""}
    v2 := []int{127}
    var v3 int64 = -771
    v4 := make(map[int]bool)
    var v5 int = -9223372036854775808
    v6 := make(map[uint32]int)
    v3 += (v3 % (((v3 & v3)) | 1))
    v3 += -1
    v5 = v5
    v24 := GenS4{F0: 255, F1: 0}
    _ = len(v6)
    fmt.Println(v3)
    v5 = v5
    v3 += -1
    var v25 string = "a"
    fmt.Println(v3, v25, v3)
    if true {
        v5 = (v5 >> ((42) & 63))
    } else {
        v26 := v6[2]
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        if false {
            var v28 int64 = (v3 - (v3 / ((0) | 1)))
            v3 -= v3
            var v29 int32 = 42
            _ = v28
            _ = v29
        } else {
            _ = len(v1)
            v3 = -1
        }
        _ = fi27
    }
    switch v25 {
    case "hello":
        switch v25 {
        case "\n\t":
            for fi30 := 0; fi30 < 14; fi30++ {
                fmt.Println(v25, fi30)
                var v31 uint32 = 0
                fmt.Println(v31, fi30)
                _ = fi30
                _ = v31
            }
        default:
            for fi32 := 0; fi32 < 16; fi32++ {
                var v33 int16 = 32767
                _ = fi32
                _ = v33
            }
        }
    default:
        v5 = v5
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        v35 := GenS3{F0: (-(v5)), F1: 809}
        v25 = v25
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi34
        _ = v35
    }
    v5 = (v5 - (v5 ^ 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
