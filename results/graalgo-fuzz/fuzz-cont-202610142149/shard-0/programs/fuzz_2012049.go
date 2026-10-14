package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int16 }
func genFn0(p0 uint8, p1 string, p2 uint64) int64 {
    p2 = (^(p2))
    p0 -= p0
    p2 = (p2 / ((p2) | 1))
    switch 127 {
    case 2147483647:
        for fi4 := 0; fi4 < 1; fi4++ {
            fmt.Println(p1)
            v5 := []int{(fi4 - fi4), (fi4 >> ((402) & 63))}
            fmt.Println(fi4, p2)
            _ = fi4
            _ = v5
        }
    }
    var v6 uint32 = 222
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -109
}

func main() {
    var v1 int16 = 1
    var v2 uint32 = 4294967295
    v3 := []int{560}
    if false {
        var v7 uint8 = 42
        _ = v7
    } else {
        var v8 uint16 = 171
        v9 := GenS1{F0: -841, F1: 2147483647, F2: 2}
        switch 690 {
        case -127:
            v1 = (v1 << ((65535) & 63))
        case 1:
            if len(v3) > 0 {
                v3[0] = 32768
            }
        case 42:
            if (23 == 100) {
                switch 128 {
                case 32767:
                    var v10 uint8 = 127
                    _ = v10
                case 9223372036854775807:
                    _ = len(v3)
                default:
                    var v11 int8 = -128
                    _ = v11
                }
                switch (("foo bar" + ("a" + "foo bar")) + (("a" + "a") + "hello")) {
                case "foo bar":
                    var v12 int16 = (v1 | (v1 - (v1 % ((v1) | 1))))
                    _ = v12
                case "":
                    for fi13 := 0; fi13 < 9; fi13++ {
                        _ = len(v3)
                        switch "hello" {
                        case "":
                            if true {
                                var v14 int64 = -542
                                v1 = v1
                                var v15 uint16 = (v8 / ((0) | 1))
                                _ = v14
                                _ = v15
                            }
                        case "\n\t":
                            for fi16 := 0; fi16 < 10; fi16++ {
                                v1 -= (v1 % (((v1 >> ((65535) & 63))) | 1))
                                _ = fi16
                            }
                        case "hello":
                            fmt.Println(v8, v2, v8)
                        default:
                            v17 := GenS1{F0: -2147483648, F1: (-(fi13)), F2: (v2 + v2)}
                            _ = v17
                        }
                        _ = fi13
                    }
                }
            }
        }
        _ = v8
        _ = v9
    }
    var v18 int64 = 42
    var v19 float64 = 3.14159
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            switch fi21 {
            case 127:
                fmt.Println(v18, v2)
                var v22 uint8 = 1
                v23 := GenS2{F0: v2}
                switch "\n\t" {
                case "\n\t":
                    var v24 uint16 = 42
                    _ = v24
                default:
                    v25 := GenS1{F0: -2, F1: (fi21 + fi21), F2: v2}
                    _ = v25
                }
                _ = v22
                _ = v23
            case 32768:
                if (42 > 1) {
                    fmt.Println(v1)
                    v19 = v19
                    var v26 uint16 = 0
                    _ = len(v3)
                    v27 := GenS1{F0: 757, F1: fi21, F2: 388}
                    _ = v26
                    _ = v27
                } else {
                    v2 -= 252
                }
                var v28 uint64 = 2
                _ = v28
            case 682:
                v18 -= (v18 + (v18 ^ (^(v18))))
            }
            _ = fi21
        }
        _ = fi20
    }
    var v29 float64 = ((((v19 - ((0.0 / (float64(-2.25) + (((3.14159 - v19) / v19) - (((v19 - v19) - v19) - v19)))) * -2.25)) * (v19 * v19)) / float64(((-2.25 / float32(v19)) + -2.25))) / v19)
    v3 = append(v3, 32767)
    if true {
        fmt.Println(v19, v29, v29)
        var v30 int64 = v18
        if (v29 < ((v29 + v19) * v29)) {
            switch ("hello" + "foo bar") {
            case "":
                v2 = 289
            }
            var v31 int32 = 2147483647
            _ = v31
        } else {
            v32 := GenS2{F0: 1}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if ((127 == 1) || (((v29 / v29) * v29) <= v29)) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v33 int = 536
                _ = v33
            }
            _ = v32
        }
        _ = v30
    }
    fmt.Println(v19)
    v18 = v18
    switch 128 {
    case 4294967295:
        var v34 int = 475
        _ = v34
    }
    v2 += 630
    var v35 uint = 255
    for fi36 := 0; fi36 < 3; fi36++ {
        var v37 uint64 = 2
        _ = fi36
        _ = v37
    }
    v3 = append(v3, -9223372036854775808)
    v38 := GenS3{F0: -2147483648, F1: v1}
    v3 = append(v3, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v38
}
