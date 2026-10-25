package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int32; F1 bool; F2 int32 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 bool }
func genFn0(p0 uint8) uint8 {
    v5 := GenS1{F0: -90}
    v6 := GenS2{F0: 2, F1: 2, F2: 0}
    _ = p0
    _ = v5
    _ = v6
    return (p0 & (^(p0)))
}
func genFn1() int {
    var v7 int = 1
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint16 = 0
        var v10 uint = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS3{F0: 42, F1: ((862 != 2) || false), F2: -2147483648}
    if true {
        v12 := []int{0, v7, (-(v7)), 2147483647}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            switch ("hello" + "") {
            case "a":
                if (true || true) {
                    v14 := []int{(v7 % (((fi13 % (((fi13 >> ((375) & 63))) | 1))) | 1)), (v7 % ((fi13) | 1)), (v7 | 662), fi13}
                    _ = v14
                }
            default:
                switch "\n\t" {
                case "foo bar":
                    var v15 uint16 = 65535
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v17 := []int{fi16}
                        _ = fi16
                        _ = v17
                    }
                }
            }
            _ = fi13
        }
    }
    v7 += (-(v7))
    _ = v11
    _ = v7
    return (v7 ^ 2147483647)
}

func main() {
    v1 := []string{"\n\t", (((("foo bar" + (("foo bar" + "\n\t") + "foo bar")) + ("\n\t" + "")) + "\n\t") + ""), "\n\t"}
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 uint8 = 2
    var v18 uint16 = 255
    switch "a" {
    case "foo bar":
        for fi19 := 0; fi19 < 18; fi19++ {
            for fi20 := 0; fi20 < 6; fi20++ {
                v21 := v2[("hello" + "hello")]
                if true {
                    fmt.Println(fi19)
                    var v22 float64 = 0.1
                    switch "a" {
                    case "hello":
                        v23 := GenS3{F0: 995, F1: false, F2: 1}
                        _ = v23
                    case "a":
                        fmt.Println(v4)
                    default:
                        _ = len(v1)
                    }
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    case "":
        v1 = append(v1, "\n\t")
    }
    var v24 float32 = 0.1
    switch 1 {
    case 4294967295:
        v25 := GenS1{F0: -1}
        _ = v25
    case -393:
        v4 = v4
    default:
        v24 += (1.5 - v24)
    }
    switch 199 {
    case -9223372036854775808:
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            switch 9223372036854775807 {
            case 996:
                v24 += v24
            case 2147483647:
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            case 9223372036854775807:
                if false {
                    for fi26 := 0; fi26 < 11; fi26++ {
                        var v27 uint32 = 42
                        var v28 int16 = 42
                        if (false && (((-2.25 / (0.1 / float64(float32(float64((v24 - (0.0 * ((v24 + v24) - float32(float64(float32((1e10 + (1e10 * -2.25))))))))))))) * 1e10) > 0.1)) {
                            if true {
                                v28 -= v28
                            } else {
                                if (113 <= -48) {
                                    fmt.Println(fi26, v27, v24)
                                    _ = len(v1)
                                    fmt.Println(v24, fi26)
                                }
                                fmt.Println(v27)
                            }
                            delete(v3, ("a" + "a"))
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
            default:
                if false {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        v18 = 785
                        for fi30 := 0; fi30 < 8; fi30++ {
                            switch "\n\t" {
                            case "a":
                                if (true || (!false)) {
                                    v31 := GenS4{F0: 692, F1: 0}
                                    _ = v31
                                }
                            case "":
                                var v32 string = "hello"
                                _ = v32
                            case "foo bar":
                                fmt.Println(v24)
                            }
                            _ = fi30
                        }
                        fmt.Println(v18, v4, v24)
                        _ = fi29
                    }
                    v24 = float32(-2.25)
                }
            }
            v24 += (((3.14159 + 1e10) / -2.25) + (v24 / (0.1 - v24)))
        }
    default:
        v1 = append(v1, "a")
    }
    switch "a" {
    case "a":
        var v33 bool = false
        _ = v33
    case "hello":
        v34 := GenS2{F0: (v18 | (v18 ^ v18)), F1: (^(v18)), F2: 935}
        _ = v34
    default:
        if true {
            var v35 uint32 = 2
            _ = v35
        }
    }
    fmt.Println(v4)
    for fi36 := 0; fi36 < 15; fi36++ {
        v37 := GenS5{F0: 5}
        for fi38 := 0; fi38 < 19; fi38++ {
            if len(v1) > 0 {
                v1[0] = (("a" + "") + "\n\t")
            }
            var v39 int64 = -1
            _ = len(v1)
            switch (^(fi36)) {
            case 205:
                fmt.Println(fi36)
            }
            _ = fi38
            _ = v39
        }
        _ = fi36
        _ = v37
    }
    v40 := GenS6{F0: v18, F1: (float64((0.0 + v24)) <= (0.0 - 0.1))}
    var v41 int64 = 9223372036854775807
    v41 = v41
    v24 = (v24 * v24)
    v24 -= (0.1 + v24)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
