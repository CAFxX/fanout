package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int; F2 bool }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int; F1 int; F2 int }
func genFn0(p0 uint8, p1 uint) uint {
    v7 := GenS1{F0: p0, F1: 100, F2: -784}
    var v8 string = (((((("a" + "hello") + ((("a" + "\n\t") + "a") + ("a" + "\n\t"))) + "") + "hello") + ("\n\t" + "foo bar")) + "a")
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p1 / ((604) | 1))
}
func genFn1(p0 int64, p1 uint) int {
    p0 += -444
    if true {
        switch "hello" {
        case "hello":
            if ((!false) && true) {
                switch "\n\t" {
                case "":
                    switch 32767 {
                    case 127:
                        for fi9 := 0; fi9 < 7; fi9++ {
                            for fi10 := 0; fi10 < 15; fi10++ {
                                v11 := GenS2{F0: 1}
                                p1 += (p1 & (p1 & p1))
                                _ = fi10
                                _ = v11
                            }
                            p1 = (p1 >> (((p1 - 65535)) & 63))
                            _ = fi9
                        }
                    default:
                        fmt.Println(p0)
                    }
                case "hello":
                    if (((1.5 / 1.0) / -2.25) == float32(float64(0.1))) {
                        var v12 uint16 = 42
                        _ = v12
                    } else {
                        v13 := []int{32768, -1}
                        _ = v13
                    }
                case "\n\t":
                    var v14 uint64 = 958
                    _ = v14
                }
                p1 -= (p1 - (p1 >> ((genFn0(1, p1)) & 63)))
                var v15 uint32 = 421
                _ = v15
            } else {
                p0 -= 9223372036854775807
            }
        case "":
            if true {
                p0 = (p0 - p0)
                if true {
                    p0 = 0
                    p1 = (p1 + (^(p1)))
                }
            } else {
                fmt.Println(p0, p0)
            }
        case "\n\t":
            v16 := GenS3{F0: 255, F1: 4294967295}
            _ = v16
        }
    }
    var v17 uint64 = 799
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 float64 = (((((1e10 + (1.5 * -2.25)) / -2.25) / 1.5) / -2.25) + ((1.5 * 0.1) / 3.14159))
        fmt.Println(v19, p0, v17)
        if (!(false && false)) {
            if ("a" > "\n\t") {
                fmt.Println(v17, fi18)
                v20 := GenS3{F0: 127, F1: 2147483647}
                if false {
                    switch "" {
                    case "\n\t":
                        for fi21 := 0; fi21 < 2; fi21++ {
                            fmt.Println(p0, v17)
                            _ = fi21
                        }
                    case "a":
                        if (!true) {
                            p0 -= (p0 / ((9223372036854775807) | 1))
                        } else {
                            for fi22 := 0; fi22 < 11; fi22++ {
                                p0 = p0
                                _ = fi22
                            }
                            v23 := []int{}
                            _ = v23
                        }
                    case "foo bar":
                        var v24 int16 = -933
                        _ = v24
                    }
                } else {
                    if (!(255 >= 128)) {
                        p1 += (p1 << ((738) & 63))
                        if ((65535 >= genFn0(0, p1)) || ((!(p1 != p1)) || false)) {
                            for fi25 := 0; fi25 < 13; fi25++ {
                                v17 += v17
                                fmt.Println(fi18)
                                var v26 int16 = 32767
                                var v27 uint8 = 42
                                p0 = (p0 % (((p0 / ((p0) | 1))) | 1))
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            var v28 bool = (p0 <= p0)
                            _ = v28
                        } else {
                            switch "a" {
                            case "foo bar":
                                v19 = ((v19 * float64(0.1)) + (v19 * v19))
                            }
                            var v29 float32 = 0.0
                            fmt.Println(p0, v19, v17)
                            _ = v29
                        }
                    }
                }
                _ = v20
            }
        }
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v17
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 65535
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    var v5 float64 = 1e10
    var v6 int32 = 829
    var v30 int16 = 1000
    if (-128 == 42) {
        if false {
            var v31 bool = (0 >= 597)
            v6 += -1
            fmt.Println(v2, v30)
            _ = v31
        }
    } else {
        v32 := GenS3{F0: 127, F1: 2147483647}
        v6 = -512
        v33 := GenS4{F0: 200, F1: genFn1(-781, 42), F2: (float64(1.5) != ((v5 / 1.0) + (v5 / float64((3.14159 * 1.5)))))}
        _ = v32
        _ = v33
    }
    _ = len(v1)
    v30 -= (v30 & -823)
    v30 += (v30 << ((18446744073709551615) & 63))
    v30 = v30
    delete(v3, 4294967295)
    v30 = (v30 - 1000)
    fmt.Println(v6, v2)
    if (!(128 >= uint8(v2))) {
        var v34 int8 = -63
        _ = v34
    }
    if (!true) {
        v35 := GenS5{F0: -2147483648, F1: v2}
        var v36 uint64 = 42
        var v37 uint64 = (v36 | (v36 % ((435) | 1)))
        _ = v35
        _ = v36
        _ = v37
    } else {
        v38 := GenS6{F0: 32767, F1: 32768, F2: 0}
        var v39 int16 = v30
        _ = v38
        _ = v39
    }
    v30 = (^(v30))
    fmt.Println(v30, v2, v30)
    v40 := GenS3{F0: 42, F1: 76}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
