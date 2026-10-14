package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32, p1 int64) int {
    p1 = -300
    p0 += p0
    v7 := GenS1{F0: 409, F1: 1}
    v8 := []int{32768, 371, 1}
    if true {
        for fi9 := 0; fi9 < 3; fi9++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi9
        }
        v10 := GenS1{F0: 939, F1: 2}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    v1 := []string{("" + ("foo bar" + "\n\t"))}
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 1
    var v4 uint = 442
    var v5 int32 = 42
    v6 := []string{"hello", "\n\t", ((((("foo bar" + (("a" + "foo bar") + "foo bar")) + ("a" + "hello")) + "") + "a") + "hello"), ("hello" + "")}
    fmt.Println(v5)
    var v11 string = ("a" + ("\n\t" + "a"))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ("" + "hello") {
    case "foo bar":
        var v12 uint = (v4 ^ 510)
        _ = v12
    case "":
        fmt.Println(v3)
    case "a":
        v13 := GenS2{F0: v4, F1: v4}
        _ = v13
    default:
        for fi14 := 0; fi14 < 15; fi14++ {
            v11 = ""
            var v15 float64 = ((((0.1 / ((3.14159 * (1e10 + (3.14159 - (1.5 + ((0.1 - -2.25) - 1.5))))) - float64((1.5 * 1e10)))) + 1.5) - float64(0.1)) + 0.1)
            _ = fi14
            _ = v15
        }
    }
    var v16 int64 = -180
    var v17 int64 = v16
    var v18 uint = (v4 - v4)
    _ = len(v6)
    v16 -= (v17 << ((255) & 63))
    var v19 int = -202
    switch (v19 - v19) {
    case 688:
        v5 += 1
    case 128:
        switch v19 {
        case 32768:
            v5 = v5
        case -148:
            if (false && true) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if false {
                    _ = len(v6)
                } else {
                    v18 -= v4
                    v17 -= -1
                    for fi20 := 0; fi20 < 8; fi20++ {
                        if (1.5 != 0.1) {
                            var v21 uint8 = 0
                            switch fi20 {
                            case 128:
                                v21 = (v21 >> ((v18) & 63))
                            case -944:
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    if false {
                                        var v22 uint16 = 2
                                        _ = v22
                                    } else {
                                        var v23 int16 = -32768
                                        v24 := GenS2{F0: v4, F1: uint(v3)}
                                        if (v18 < v4) {
                                            v4 = 18446744073709551615
                                            var v25 int8 = 1
                                            _ = v25
                                        }
                                        _ = v23
                                        _ = v24
                                    }
                                default:
                                    var v26 int32 = v5
                                    _ = v26
                                }
                            case 0:
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    switch v19 {
                                    case -1:
                                        var v28 uint16 = 1
                                        _ = v28
                                    case 1:
                                        switch "a" {
                                        case "hello":
                                            v6 = append(v6, ("foo bar" + ""))
                                        case "foo bar":
                                            var v29 int32 = -197
                                            _ = v29
                                        default:
                                            if false {
                                                v30 := GenS3{F0: (v21 / ((143) | 1))}
                                                fmt.Println(fi27, fi20, v19)
                                                var v31 uint64 = 690
                                                _ = v30
                                                _ = v31
                                            } else {
                                                if false {
                                                    if (42 != 42) {
                                                        v1 = append(v1, (((("a" + v11) + v11) + "foo bar") + ""))
                                                    }
                                                } else {
                                                    v17 -= -380
                                                }
                                            }
                                        }
                                    }
                                    var v32 string = "hello"
                                    var v33 uint8 = v21
                                    _ = fi27
                                    _ = v32
                                    _ = v33
                                }
                            default:
                                v34 := GenS1{F0: (fi20 >> (((v18 % ((v18) | 1))) & 63)), F1: 307}
                                _ = v34
                            }
                            _ = v21
                        } else {
                            v3 = (^(v3))
                        }
                        for fi35 := 0; fi35 < 20; fi35++ {
                            fmt.Println(v16)
                            v4 = uint(v3)
                            var v36 float64 = ((((0.1 * 1e10) - float64(((3.14159 / 1.0) + 0.0))) * -2.25) * 0.1)
                            _ = fi35
                            _ = v36
                        }
                        _ = fi20
                    }
                }
            }
        case 42:
            v2[(v3 / ((v3) | 1))] = 4294967295
        default:
            v16 = (v17 % ((-9223372036854775808) | 1))
        }
    }
    v18 = 65535
    v2[2147483648] = 2147483648
    var v37 int = (v19 ^ (-(v19)))
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v38 := GenS3{F0: 2}
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
