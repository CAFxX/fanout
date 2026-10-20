package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint32; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 int64; F1 uint16 }
func genFn0(p0 int8, p1 int64, p2 int64) int32 {
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(p2, p2, fi5)
        _ = fi5
    }
    var v6 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}

func main() {
    v1 := []int{32768, -767, 0, -406}
    var v2 uint8 = 255
    v3 := []int{863, 42, 9223372036854775807, 0}
    var v4 int32 = 958
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float64 = (-2.25 - 0.0)
        if true {
            v9 := GenS1{F0: ((0.0 != 1.5) && false), F1: 255, F2: 537}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v10 uint16 = 32767
    switch 739 {
    case 127:
        var v11 bool = false
        _ = v11
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2, v10, v2)
    v10 += (v10 & (v10 | v10))
    if (873 <= 0) {
        v2 += 2
    } else {
        var v12 uint16 = 255
        var v13 int16 = 1
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        switch fi14 {
        case 704:
            if ((v10 < 65535) && true) {
                fmt.Println(v4)
                switch "hello" {
                case "a":
                    var v15 int = (-(fi14))
                    _ = v15
                }
            } else {
                if len(v1) > 0 {
                    v1[0] = fi14
                }
                v16 := GenS2{F0: v4, F1: 4294967295, F2: (v4 == (v4 % ((v4) | 1)))}
                _ = v16
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v17 := GenS3{F0: fi14}
        _ = fi14
        _ = v17
    }
    var v18 float32 = (((3.14159 * 3.14159) / float32(-2.25)) - ((float32((-2.25 * -2.25)) + 0.0) - (-2.25 + -2.25)))
    v19 := GenS2{F0: v4, F1: uint32(v2), F2: true}
    for fi20 := 0; fi20 < 15; fi20++ {
        for fi21 := 0; fi21 < 19; fi21++ {
            switch fi20 {
            case 482:
                if (-9223372036854775808 <= 9223372036854775807) {
                    _ = len(v3)
                    v18 = 1e10
                    if (false || true) {
                        var v22 bool = true
                        _ = v22
                    } else {
                        if false {
                            if (!(!("hello" > "\n\t"))) {
                                v23 := GenS4{F0: (v2 != (v2 >> ((18446744073709551615) & 63)))}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v23
                            }
                        } else {
                            v10 = (v10 | uint16(v2))
                            if false {
                                var v24 int = 32767
                                if ((v10 % ((150) | 1)) > v10) {
                                    switch "a" {
                                    case "":
                                        var v25 uint = 940
                                        _ = v25
                                    }
                                    v26 := GenS5{F0: -1, F1: false}
                                    _ = v26
                                } else {
                                    fmt.Println(fi20, fi20)
                                    if true {
                                        v27 := GenS3{F0: -960}
                                        _ = v27
                                    }
                                }
                                _ = v24
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            fmt.Println(v4, v10, fi21)
                            var v28 bool = false
                            v2 -= 2
                            _ = v28
                        }
                        v18 = (float32(3.14159) + float32(1e10))
                    }
                    v2 += v2
                    v29 := GenS6{F0: v10}
                    v18 += v18
                    switch (fi20 + 127) {
                    case 32767:
                        fmt.Println(v10, v10, v10)
                    case 480:
                        var v30 int16 = 1000
                        _ = v30
                    case 0:
                        v31 := GenS6{F0: v10}
                        _ = v31
                    default:
                        var v32 uint = 18446744073709551615
                        _ = v32
                    }
                    _ = v29
                } else {
                    fmt.Println(fi21)
                    switch "hello" {
                    case "hello":
                        v33 := GenS7{F0: 910, F1: 711}
                        _ = v33
                    default:
                        fmt.Println(v10, v2, v4)
                    }
                    v1 = append(v1, fi21)
                    var v34 string = ("a" + "hello")
                    v34 = ("\n\t" + v34)
                    _ = v34
                }
            case 2147483647:
                v35 := GenS5{F0: 1, F1: false}
                _ = v35
            }
            if true {
                var v36 int = fi20
                _ = v36
            } else {
                if true {
                    switch "hello" {
                    case "foo bar":
                        v18 += v18
                    case "hello":
                        var v37 int32 = v4
                        _ = v37
                    }
                }
            }
            _ = fi21
        }
        _ = len(v3)
        v3 = append(v3, (fi20 ^ fi20))
        _ = fi20
    }
    v38 := GenS8{F0: -9223372036854775808, F1: (v10 & (v10 / ((v10) | 1)))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
