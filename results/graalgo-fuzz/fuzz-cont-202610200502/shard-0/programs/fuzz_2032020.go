package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint16; F2 bool }
type GenS3 struct { F0 uint32; F1 int64; F2 int16 }

func main() {
    var v1 int = 1
    var v2 int = (v1 & (^(v1)))
    var v3 uint8 = 1
    var v4 int = 32768
    for fi5 := 0; fi5 < 13; fi5++ {
        switch ("" + "hello") {
        case "hello":
            for fi6 := 0; fi6 < 11; fi6++ {
                if ((v3 << ((42) & 63)) > 127) {
                    v7 := []int{}
                    if (1 >= 944) {
                        v4 = v2
                        v2 += (v4 + fi5)
                        v8 := GenS1{F0: 813}
                        _ = v8
                    } else {
                        fmt.Println(v2, v3, fi5)
                    }
                    _ = v7
                } else {
                    v9 := []int{(fi6 + (v2 + v1)), (fi6 | (v1 >> ((1) & 63))), v2, (-(v2))}
                    switch (v1 ^ v2) {
                    case 0:
                        var v10 int = v1
                        _ = v10
                    }
                    var v11 string = "hello"
                    _ = v11
                    _ = v9
                }
                _ = fi6
            }
        default:
            switch "foo bar" {
            case "\n\t":
                if true {
                    fmt.Println(v2, v2, fi5)
                    var v12 int32 = -1
                    switch int(v3) {
                    case 1:
                        v13 := []int{v4}
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 11; fi14++ {
                            v2 -= fi14
                            _ = fi14
                        }
                    }
                    _ = v12
                } else {
                    switch ((("a" + "a") + "") + "hello") {
                    case "\n\t":
                        v15 := GenS1{F0: v4}
                        _ = v15
                    case "":
                        if false {
                            v16 := []int{(v2 >> ((0) & 63)), (v2 % ((fi5) | 1))}
                            var v17 bool = (18446744073709551615 <= 2)
                            _ = v16
                            _ = v17
                        }
                    }
                }
            case "":
                v2 += fi5
            }
        }
        _ = fi5
    }
    v1 -= 9223372036854775807
    v18 := []int{(v2 << ((18446744073709551615) & 63)), (v1 * (v4 | v1)), -2147483648}
    for fi19 := 0; fi19 < 18; fi19++ {
        v2 = (fi19 + -2147483648)
        v1 = fi19
        _ = fi19
    }
    fmt.Println(v2, v1, v4)
    for fi20 := 0; fi20 < 2; fi20++ {
        if (1000 <= 1) {
            v21 := GenS1{F0: v4}
            if true {
                if (1.5 == (3.14159 / (3.14159 / 1.5))) {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        if false {
                            v18 = append(v18, v4)
                            switch v4 {
                            case -722:
                                v2 -= (v1 | v1)
                            case -799:
                                v23 := GenS2{F0: -1, F1: 255, F2: ((130 > 222) && true)}
                                _ = v23
                            case 128:
                                var v24 int8 = 42
                                _ = v24
                            default:
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            }
                            if (false || false) {
                                v1 = v1
                            } else {
                                if len(v18) > 0 {
                                    v18[0] = v2
                                }
                                fmt.Println(fi20, v1, v4)
                            }
                        } else {
                            var v25 uint32 = 42
                            switch "" {
                            case "foo bar":
                                v18 = append(v18, (v1 + (fi22 / (((-(v1))) | 1))))
                            case "hello":
                                v25 = 761
                            case "\n\t":
                                var v26 int = (fi20 << ((2) & 63))
                                fmt.Println(fi22)
                                _ = v26
                            default:
                                fmt.Println(fi20, v4)
                            }
                            _ = v25
                        }
                        fmt.Println(v4, v3)
                        _ = fi22
                    }
                    if (1e10 < 3.14159) {
                        var v27 uint16 = 33
                        _ = len(v18)
                        _ = v27
                    }
                    if true {
                        fmt.Println(fi20)
                        v4 = (v4 & v1)
                    }
                }
            } else {
                for fi28 := 0; fi28 < 2; fi28++ {
                    switch "a" {
                    case "\n\t":
                        v29 := GenS3{F0: 87, F1: 42, F2: 739}
                        switch (v1 - -9223372036854775808) {
                        case 9223372036854775807:
                            var v30 uint32 = 2
                            _ = v30
                        case -2147483648:
                            _ = len(v18)
                        }
                        var v31 int = (fi20 - 479)
                        _ = v29
                        _ = v31
                    case "":
                        if len(v18) > 0 {
                            v18[0] = fi20
                        }
                    default:
                        v3 += (v3 >> ((42) & 63))
                    }
                    v3 -= 2
                    var v32 int64 = 609
                    _ = len(v18)
                    _ = fi28
                    _ = v32
                }
                var v33 float32 = float32(1e10)
                var v34 int16 = 0
                _ = v33
                _ = v34
            }
            if len(v18) > 0 {
                v18[0] = v4
            }
            _ = v21
        }
        v4 = (^(v4))
        v18 = append(v18, v2)
        _ = fi20
    }
    v1 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
