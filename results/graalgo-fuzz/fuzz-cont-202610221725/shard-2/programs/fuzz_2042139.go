package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 bool; F1 int64; F2 int32 }
type GenS4 struct { F0 int; F1 uint; F2 bool }
type GenS5 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 int8 }
func genFn0() uint {
    var v7 bool = false
    switch "a" {
    case "foo bar":
        var v8 uint = 255
        _ = v8
    case "":
        for fi9 := 0; fi9 < 12; fi9++ {
            if true {
                var v10 int = fi9
                v11 := []int{(^(v10))}
                _ = len(v11)
                _ = v10
                _ = v11
            } else {
                switch "" {
                case "hello":
                    v12 := []int{fi9}
                    _ = v12
                case "a":
                    fmt.Println(v7)
                case "":
                    if ((0.1 + 3.14159) <= float32(0.0)) {
                        if (127 != 253) {
                            if v7 {
                                v7 = v7
                                v7 = (42 == 153)
                                v13 := GenS1{F0: ((42 < -84) && true), F1: 200, F2: 65535}
                                _ = v13
                            } else {
                                for fi14 := 0; fi14 < 19; fi14++ {
                                    v7 = true
                                    _ = fi14
                                }
                                var v15 uint8 = 255
                                var v16 bool = ((3.14159 >= (((float64((-2.25 / 0.1)) * 0.0) + (((-2.25 * -2.25) - 1.5) / 1e10)) * 0.0)) && v7)
                                _ = v15
                                _ = v16
                            }
                        }
                        for fi17 := 0; fi17 < 15; fi17++ {
                            v18 := []int{(fi9 ^ fi17), -1, fi17}
                            fmt.Println(fi9, fi17)
                            _ = fi17
                            _ = v18
                        }
                        v7 = (0 != 18446744073709551615)
                    }
                }
                if ((1 <= 42) && v7) {
                    v7 = false
                    for fi19 := 0; fi19 < 16; fi19++ {
                        if v7 {
                            v7 = (1 > 4294967295)
                            var v20 int8 = 127
                            _ = v20
                        }
                        switch (fi9 * fi9) {
                        case 32767:
                            v21 := GenS1{F0: (false && true), F1: 389, F2: 895}
                            _ = v21
                        }
                        if v7 {
                            switch "a" {
                            case "\n\t":
                                v22 := GenS2{F0: 1, F1: -954}
                                _ = v22
                            case "a":
                                v23 := GenS3{F0: false, F1: -9223372036854775808, F2: -2147483648}
                                _ = v23
                            case "foo bar":
                                var v24 uint8 = 193
                                _ = v24
                            default:
                                if v7 {
                                    var v25 uint64 = 42
                                    _ = v25
                                }
                                v7 = true
                                fmt.Println(fi9, v7, fi9)
                            }
                        }
                        v26 := []int{fi9, 953}
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        switch 9223372036854775807 {
                        case 128:
                            _ = len(v26)
                        case 32767:
                            var v27 int16 = -32768
                            _ = v27
                        case 2147483647:
                            fmt.Println(fi9, fi9, v7)
                        default:
                            fmt.Println(fi9)
                        }
                        if ((18446744073709551615 != 42) || ((!true) || (false || (fi19 < 1)))) {
                            fmt.Println(fi9, fi19, fi19)
                            fmt.Println(v7, fi9, fi9)
                            var v28 int64 = -1
                            _ = v28
                        } else {
                            v7 = (0 >= 127)
                            if (!true) {
                                v26 = append(v26, 128)
                                v7 = (!false)
                                v29 := GenS1{F0: (!v7), F1: 446, F2: 525}
                                _ = v29
                            }
                            v7 = v7
                        }
                        _ = fi19
                        _ = v26
                    }
                } else {
                    v7 = (!(0.1 != 1.5))
                    v7 = v7
                }
                var v30 float64 = -2.25
                _ = v30
            }
            _ = fi9
        }
    }
    if v7 {
        v31 := GenS4{F0: 382, F1: 42, F2: v7}
        var v32 uint = 4294967295
        _ = v31
        _ = v32
    } else {
        v33 := GenS5{F0: 737, F1: 42, F2: 42}
        v34 := GenS6{F0: 42}
        _ = v33
        _ = v34
    }
    _ = v7
    return 40
}

func main() {
    v1 := []int{-75, 127}
    var v2 float32 = float32(-2.25)
    v3 := []uint32{4294967295, 1, 42, 670}
    var v4 float32 = v2
    var v5 uint8 = 182
    var v6 bool = true
    fmt.Println(v2, v6, v4)
    fmt.Println(v6, v4)
    var v35 uint = 1
    if v6 {
        for fi36 := 0; fi36 < 18; fi36++ {
            v4 = 1.5
            var v37 int = (fi36 - (fi36 * (^(fi36))))
            var v38 string = "a"
            _ = fi36
            _ = v37
            _ = v38
        }
        if (1 == 748) {
            var v39 bool = (42 <= 9223372036854775807)
            v40 := GenS7{F0: 2, F1: 100}
            _ = v39
            _ = v40
        }
        fmt.Println(v5, v6)
    }
    _ = len(v3)
    fmt.Println(v2, v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        for fi42 := 0; fi42 < 8; fi42++ {
            v35 += 1
            v3 = append(v3, 2147483647)
            _ = fi42
        }
        _ = fi41
    }
    switch ((("hello" + "hello") + "a") + "foo bar") {
    case "":
        v6 = v6
    default:
        v6 = false
    }
    if v6 {
        v2 -= 1.5
        v4 -= ((v4 - v2) + float32((-2.25 - (0.0 / 3.14159))))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
