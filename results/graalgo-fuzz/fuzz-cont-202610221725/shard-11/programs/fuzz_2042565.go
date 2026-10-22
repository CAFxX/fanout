package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS5 struct { F0 bool; F1 int; F2 bool }
type GenS6 struct { F0 int; F1 bool; F2 int64 }
type GenS7 struct { F0 uint16; F1 int64; F2 int16 }
type GenS8 struct { F0 int64; F1 int8; F2 uint64 }
func genFn0() uint64 {
    v5 := []int{42}
    var v6 int8 = 97
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int) uint8 {
    for fi7 := 0; fi7 < 8; fi7++ {
        p0 = (fi7 | p0)
        p0 = (p0 >> ((10) & 63))
        v8 := GenS1{F0: -1, F1: 0}
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    _ = p0
    _ = v9
    return 239
}

func main() {
    var v1 int16 = -1
    var v2 uint8 = 1
    var v3 float32 = (3.14159 - 0.1)
    v4 := make(map[string]bool)
    for fi10 := 0; fi10 < 6; fi10++ {
        _ = len(v4)
        if (65535 >= 42) {
            v2 -= (v2 | v2)
            v3 = 1e10
        } else {
            if true {
                v11 := GenS1{F0: 1, F1: 42}
                v3 -= v3
                _ = v11
            } else {
                for fi12 := 0; fi12 < 4; fi12++ {
                    if (true && true) {
                        var v13 float32 = v3
                        if false {
                            switch (^(fi10)) {
                            case 2147483647:
                                if false {
                                    if true {
                                        v3 = (float32((0.0 * float64(v3))) - (-2.25 + 1e10))
                                    }
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(v2, fi10)
                                    case "hello":
                                        if ((!(float64(float32(1e10)) <= (3.14159 + 1.5))) && false) {
                                            fmt.Println(fi10)
                                        } else {
                                            v14 := GenS1{F0: 0, F1: 42}
                                            _ = v14
                                        }
                                    case "":
                                        fmt.Println(fi10)
                                    default:
                                        v3 = ((((v3 * (v13 - 1.5)) / v3) * v3) + (0.1 * v3))
                                    }
                                }
                            default:
                                switch -258 {
                                case 128:
                                    var v15 uint = 4294967295
                                    _ = v15
                                case 0:
                                    var v16 uint16 = 42
                                    _ = v16
                                case 4294967295:
                                    switch "a" {
                                    case "a":
                                        v17 := GenS1{F0: -51, F1: 1}
                                        _ = v17
                                    default:
                                        v13 = v3
                                        fmt.Println(fi12)
                                        if false {
                                            v2 = (v2 % ((v2) | 1))
                                            if (((-2147483648 == -396) || false) || (-1 > int16(fi10))) {
                                                v3 = ((float32((((((((-2.25 / 1.0) - 0.1) * ((-2.25 - float64(v3)) / (1e10 * 1.5))) / float64(v3)) / 3.14159) + 0.0) * 1.5)) / (-2.25 * (v3 - v13))) - v3)
                                                fmt.Println(v13, v1, v3)
                                                v1 += v1
                                            } else {
                                                v18 := GenS2{F0: (v2 & v2)}
                                                switch "foo bar" {
                                                case "hello":
                                                    delete(v4, "")
                                                    fmt.Println(fi10)
                                                    var v19 uint = 65535
                                                    v3 = 3.14159
                                                    _ = v19
                                                default:
                                                    if (42 >= 592) {
                                                        v20 := GenS3{F0: 127, F1: 18446744073709551615}
                                                        v2 = (v2 ^ (v2 + genFn1(fi10)))
                                                        v21 := GenS4{F0: 1, F1: 2, F2: (v1 & (^(v1)))}
                                                        _ = v20
                                                        _ = v21
                                                    }
                                                }
                                                _ = v18
                                            }
                                            v3 = v13
                                            v1 = (v1 >> ((42) & 63))
                                            v1 -= (v1 << ((65535) & 63))
                                        }
                                    }
                                default:
                                    v3 -= (v3 / (v3 - 0.0))
                                }
                            }
                            fmt.Println(v2, fi12, v3)
                        }
                        _ = v13
                    }
                    v22 := v4[""]
                    v23 := GenS5{F0: ("" != "foo bar"), F1: 128, F2: false}
                    var v24 float64 = 1.5
                    if (v1 >= 425) {
                        var v25 int = fi12
                        _ = v25
                    }
                    _ = fi12
                    _ = v22
                    _ = v23
                    _ = v24
                }
                v26 := GenS1{F0: 0, F1: 22}
                _ = v26
            }
            var v27 float64 = 0.0
            _ = v27
        }
        v28 := []int{fi10, (-(fi10)), 128}
        _ = fi10
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        v2 += v2
        v30 := GenS6{F0: (fi29 << ((0) & 63)), F1: (float32(-2.25) == v3), F2: -9223372036854775808}
        _ = fi29
        _ = v30
    }
    var v31 int = 9223372036854775807
    for fi32 := 0; fi32 < 15; fi32++ {
        fmt.Println(v1)
        v33 := GenS4{F0: 2147483647, F1: 299, F2: 1000}
        _ = fi32
        _ = v33
    }
    v34 := GenS7{F0: 482, F1: 42, F2: v1}
    for fi35 := 0; fi35 < 7; fi35++ {
        v31 -= (-(fi35))
        _ = fi35
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        v37 := GenS8{F0: 42, F1: 100, F2: 42}
        if (2 != 182) {
            v2 = 2
            var v38 float32 = 0.0
            fmt.Println(v31, v1, fi36)
            _ = v38
        } else {
            if true {
                var v39 uint = 249
                v3 += v3
                fmt.Println(v1, v31)
                _ = v39
            }
            var v40 int = v31
            _ = v40
        }
        _ = fi36
        _ = v37
    }
    v31 = (v31 * v31)
    v3 -= v3
    v41 := GenS6{F0: -1, F1: true, F2: -9223372036854775808}
    switch "a" {
    case "hello":
        fmt.Println(v2)
    case "foo bar":
        var v42 int16 = v1
        _ = v42
    case "a":
        fmt.Println(v1)
    default:
        delete(v4, ("\n\t" + ""))
    }
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v41
}
