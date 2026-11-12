package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 int64; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int, p1 string, p2 uint) int64 {
    fmt.Println(p0, p0, p0)
    var v5 uint8 = 89
    p1 = (p1 + (("a" + p1) + ("\n\t" + p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 18446744073709551615
    var v3 bool = true
    var v4 int = 100
    var v6 uint16 = 0
    var v7 uint16 = (^(v6))
    fmt.Println(v2, v2)
    v6 = v7
    v6 -= 2
    v6 += uint16(v2)
    if false {
        var v8 uint16 = 42
        var v9 int64 = 9223372036854775807
        var v10 uint8 = 2
        _ = v10
        _ = v8
        _ = v9
    } else {
        var v11 string = ("hello" + (("foo bar" + ("\n\t" + "")) + "a"))
        var v12 int8 = 0
        _ = v11
        _ = v12
    }
    if (("a" + "hello") >= "") {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 int16 = 42
            if v3 {
                if (-509 >= 42) {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        var v16 int16 = (-(v14))
                        _ = fi15
                        _ = v16
                    }
                } else {
                    v17 := GenS1{F0: 168}
                    _ = len(v1)
                    v18 := []int{-1, fi13, fi13}
                    _ = v17
                    _ = v18
                }
            } else {
                v4 += (fi13 % ((fi13) | 1))
                var v19 uint16 = 2
                _ = v19
            }
            v1[109] = 42
            _ = fi13
            _ = v14
        }
        if false {
            v20 := GenS1{F0: 150}
            if (!false) {
                var v21 float64 = 0.0
                if true {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v23 := GenS2{F0: v3}
                        for fi24 := 0; fi24 < 14; fi24++ {
                            var v25 string = ""
                            v26 := GenS3{F0: genFn0(fi24, "", (^(v2)))}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        fmt.Println(v2)
                        _ = fi22
                        _ = v23
                    }
                } else {
                    delete(v1, 2147483648)
                    fmt.Println(v21)
                }
                if true {
                    v2 = v2
                } else {
                    var v27 uint16 = v7
                    v3 = false
                    _ = v27
                }
                _ = v21
            } else {
                v7 += 890
            }
            _ = v20
        }
    } else {
        fmt.Println(v2, v7, v7)
        if (v4 <= (^(v4))) {
            v6 -= v7
            switch "\n\t" {
            case "hello":
                v4 = 32767
            case "a":
                v3 = false
            default:
                var v28 bool = v3
                _ = v28
            }
        }
        v3 = v3
    }
    v6 -= (v6 & v7)
    if v3 {
        v29 := GenS2{F0: true}
        if v3 {
            var v30 int64 = 42
            v31 := GenS1{F0: 42}
            _ = v30
            _ = v31
        }
        _ = v29
    } else {
        if v3 {
            v32 := GenS4{F0: 42}
            delete(v1, 4294967295)
            _ = v32
        }
        v7 = (v7 / ((737) | 1))
        for fi33 := 0; fi33 < 8; fi33++ {
            for fi34 := 0; fi34 < 7; fi34++ {
                fmt.Println(v3, v7, fi33)
                var v35 int16 = -1
                v36 := GenS5{F0: false, F1: -679, F2: uint32(v2)}
                _ = fi34
                _ = v35
                _ = v36
            }
            var v37 int8 = -72
            v38 := GenS3{F0: -9223372036854775808}
            _ = fi33
            _ = v37
            _ = v38
        }
    }
    v2 += v2
    switch "a" {
    case "":
        var v39 float32 = (((1.5 * -2.25) - 0.0) - (float32(0.1) + 1e10))
        _ = v39
    case "\n\t":
        for fi40 := 0; fi40 < 1; fi40++ {
            switch "" {
            case "a":
                v41 := GenS5{F0: v3, F1: 9223372036854775807, F2: 303}
                _ = v41
            case "hello":
                for fi42 := 0; fi42 < 20; fi42++ {
                    var v43 int32 = -2147483648
                    _ = fi42
                    _ = v43
                }
            default:
                v3 = (2 > 18446744073709551615)
            }
            _ = fi40
        }
    case "a":
        for fi44 := 0; fi44 < 17; fi44++ {
            if true {
                var v45 float32 = float32(3.14159)
                _ = v45
            }
            fmt.Println(v6)
            _ = fi44
        }
    default:
        for fi46 := 0; fi46 < 18; fi46++ {
            v2 = v2
            _ = fi46
        }
    }
    if v3 {
        var v47 bool = v3
        delete(v1, 2147483648)
        _ = v47
    } else {
        v2 -= v2
        if true {
            switch (v4 | (v4 - (v4 % (((^(v4))) | 1)))) {
            case -9223372036854775808:
                v6 = v7
            case -619:
                v48 := v1[4294967295]
                _ = v48
            case 42:
                v49 := GenS2{F0: true}
                _ = v49
            }
        }
    }
    v50 := []int{-1, (v4 ^ v4), v4}
    switch ("a" + "a") {
    case "\n\t":
        switch -493 {
        case 127:
            v50 = append(v50, (v4 >> (((v2 % ((18446744073709551615) | 1))) & 63)))
        default:
            if v3 {
                for fi51 := 0; fi51 < 17; fi51++ {
                    if (4294967295 != (v2 ^ v2)) {
                        v52 := GenS6{F0: v7, F1: 255}
                        if (4294967295 >= 4294967295) {
                            if (0 <= 42) {
                                _ = len(v1)
                                fmt.Println(v2, v2, v4)
                            }
                        } else {
                            switch "foo bar" {
                            case "":
                                v53 := GenS2{F0: true}
                                _ = v53
                            }
                            v3 = (655 == 42)
                        }
                        var v54 int = fi51
                        _ = v52
                        _ = v54
                    } else {
                        v2 -= 32
                        if len(v50) > 0 {
                            v50[0] = (v4 & fi51)
                        }
                    }
                    v4 -= 2147483647
                    _ = fi51
                }
            } else {
                switch "foo bar" {
                case "hello":
                    v55 := GenS3{F0: -956}
                    _ = v55
                case "":
                    var v56 string = "a"
                    _ = v56
                case "\n\t":
                    var v57 int64 = 235
                    _ = v57
                }
                if len(v50) >= 2 {
                    v50 = v50[1:len(v50)]
                }
                var v58 uint = 740
                _ = v58
            }
        }
    case "hello":
        v7 += v7
    case "":
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v6
    _ = v7
}
