package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int }
func genFn0(p0 int64, p1 uint, p2 int) uint64 {
    p0 -= (^(p0))
    var v5 uint8 = 0
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}
func genFn1(p0 int16) uint32 {
    var v6 int = -9223372036854775808
    v6 -= (v6 % ((4294967295) | 1))
    v7 := []int{v6, -1, v6, -2147483648}
    _ = p0
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 int16 = 329
    var v2 uint32 = 0
    v3 := make(map[uint32]bool)
    var v4 uint = 822
    for fi8 := 0; fi8 < 16; fi8++ {
        v4 = (^(v4))
        _ = fi8
    }
    v2 = v2
    for fi9 := 0; fi9 < 9; fi9++ {
        delete(v3, (v2 << ((18446744073709551615) & 63)))
        var v10 float64 = 1.5
        _ = fi9
        _ = v10
    }
    switch (("" + "") + "\n\t") {
    case "\n\t":
        var v11 uint16 = 255
        _ = v11
    case "hello":
        if true {
            if (18446744073709551615 != 2) {
                var v12 float64 = (0.1 / 3.14159)
                _ = v12
            }
            switch "a" {
            case "a":
                switch int(v4) {
                case 1:
                    if true {
                        _ = len(v3)
                    }
                default:
                    for fi13 := 0; fi13 < 13; fi13++ {
                        for fi14 := 0; fi14 < 8; fi14++ {
                            var v15 float32 = float32(1e10)
                            if true {
                                var v16 bool = ((-(v1)) >= (^(v1)))
                                if ((0.0 > 1.5) && false) {
                                    v17 := GenS1{F0: 823, F1: 42, F2: 2}
                                    _ = v17
                                }
                                switch fi13 {
                                case 42:
                                    fmt.Println(fi14, v4)
                                default:
                                    v18 := GenS2{F0: true, F1: uint8(v2)}
                                    _ = v18
                                }
                                delete(v3, genFn1(-533))
                                v19 := GenS3{F0: 42}
                                _ = v16
                                _ = v19
                            }
                            v20 := GenS4{F0: "hello"}
                            _ = fi14
                            _ = v15
                            _ = v20
                        }
                        _ = fi13
                    }
                }
            case "hello":
                fmt.Println(v4, v2)
            }
            v21 := GenS4{F0: ("foo bar" + "\n\t")}
            _ = v21
        }
    default:
        v1 = (v1 / ((-172) | 1))
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 uint64 = uint64(v2)
        v24 := GenS4{F0: ("a" + "a")}
        fmt.Println(v2, v23)
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 int = (fi25 / (((fi25 / (((fi25 / (((fi25 % (((fi25 | fi25)) | 1))) | 1))) | 1))) | 1))
        var v27 float64 = 0.1
        v27 -= v27
        _ = fi25
        _ = v26
        _ = v27
    }
    v1 = (^(v1))
    fmt.Println(v2, v1, v2)
    var v28 float32 = 1e10
    for fi29 := 0; fi29 < 15; fi29++ {
        _ = len(v3)
        _ = fi29
    }
    fmt.Println(v28)
    if true {
        var v30 uint = v4
        _ = v30
    } else {
        fmt.Println(v28, v2)
    }
    if (true && ((false && false) || true)) {
        switch "hello" {
        case "a":
            for fi31 := 0; fi31 < 13; fi31++ {
                v28 = (-2.25 + float32(((-2.25 + (3.14159 - (3.14159 * 0.1))) * (0.1 - float64(v28)))))
                _ = fi31
            }
        case "":
            v28 = ((v28 * v28) - (v28 * 0.1))
        case "hello":
            if false {
                v32 := []int{-376, -9223372036854775808}
                if true {
                    v3[v2] = false
                }
                _ = v32
            } else {
                for fi33 := 0; fi33 < 14; fi33++ {
                    v34 := GenS5{F0: -159, F1: v4}
                    _ = fi33
                    _ = v34
                }
                switch 128 {
                case 42:
                    v2 = (v2 | (^(v2)))
                }
                v35 := GenS6{F0: 32768}
                _ = v35
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
