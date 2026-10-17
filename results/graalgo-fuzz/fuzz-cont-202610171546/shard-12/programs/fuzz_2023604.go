package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 uint; F1 uint32; F2 int8 }
func genFn0() string {
    var v7 uint16 = 42
    v7 -= (v7 * v7)
    _ = v7
    return "hello"
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]int)
    var v3 string = "a"
    var v4 float32 = (float32(-2.25) * 3.14159)
    var v5 float64 = -2.25
    v6 := make(map[int]bool)
    var v8 uint16 = v1
    v9 := GenS1{F0: 2, F1: 1}
    if false {
        var v10 float32 = v4
        v4 -= (1e10 - float32(float64(0.1)))
        fmt.Println(v3, v4)
        _ = v10
    } else {
        if false {
            if false {
                var v11 bool = true
                if false {
                    v5 += float64(1e10)
                } else {
                    switch 128 {
                    case 128:
                        v12 := GenS2{F0: -1, F1: 0}
                        _ = v12
                    default:
                        for fi13 := 0; fi13 < 12; fi13++ {
                            fmt.Println(fi13)
                            _ = fi13
                        }
                    }
                    v4 -= (v4 + ((1.5 * float32((0.0 - float64((1e10 + v4))))) - v4))
                    var v14 uint32 = 1
                    _ = v14
                }
                if v11 {
                    var v15 int8 = -16
                    if (false || v11) {
                        for fi16 := 0; fi16 < 3; fi16++ {
                            v1 -= v8
                            _ = fi16
                        }
                        switch v3 {
                        case "hello":
                            v17 := GenS1{F0: 230, F1: -2147483648}
                            _ = v17
                        case "a":
                            v1 = 729
                        }
                    }
                    _ = v15
                }
                _ = v11
            }
            v18 := []int{128}
            v8 -= (v8 * 2)
            _ = v18
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := GenS1{F0: 255, F1: 1}
                if (true && (fi19 < 128)) {
                    v8 = v1
                } else {
                    fmt.Println(v1, v5, v8)
                    v22 := GenS2{F0: 127, F1: 87}
                    _ = len(v2)
                    _ = v22
                }
                switch (fi19 - fi20) {
                case 42:
                    fmt.Println(v4)
                default:
                    var v23 int16 = 32767
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            v4 = (3.14159 - v4)
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 float32 = 0.1
                switch v3 {
                case "":
                    _ = len(v2)
                    v6[fi24] = false
                }
                v8 += (v1 | (v8 % ((32767) | 1)))
                v26 := []int{-2147483648, 127, fi24, fi19}
                fmt.Println(v8)
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi19
        }
    }
    v27 := GenS3{F0: false, F1: (0.1 == -2.25)}
    var v28 int8 = 42
    if false {
        v29 := GenS1{F0: 255, F1: 1}
        var v30 uint16 = 32767
        fmt.Println(v5, v28, v5)
        _ = v29
        _ = v30
    }
    if (v1 != (v8 & v8)) {
        switch v3 {
        case "foo bar":
            v31 := []int{}
            _ = v31
        case "":
            _ = len(v2)
        default:
            if true {
                if true {
                    v32 := GenS4{F0: 0, F1: 141}
                    _ = v32
                } else {
                    if (((((float32(v5) / v4) - v4) * v4) > (v4 - v4)) && ((^(v8)) > 42)) {
                        fmt.Println(v8, v4)
                        switch 4294967295 {
                        case 32768:
                            switch 0 {
                            case 543:
                                v1 += v8
                            case 1:
                                v33 := GenS4{F0: 172, F1: 2147483647}
                                _ = v33
                            case 128:
                                switch -9223372036854775808 {
                                case 32767:
                                    if true {
                                        var v34 uint32 = 2
                                        for fi35 := 0; fi35 < 19; fi35++ {
                                            v36 := GenS4{F0: 42, F1: 660}
                                            _ = fi35
                                            _ = v36
                                        }
                                        _ = v34
                                    }
                                }
                            default:
                                v5 = (v5 + (float64((v4 / ((v4 * (1e10 / v4)) - 1.5))) / (((-2.25 * float64(-2.25)) * v5) / v5)))
                            }
                        case 1:
                            v37 := GenS1{F0: 7, F1: -1}
                            _ = v37
                        case 32767:
                            if ((!false) || false) {
                                v5 = -2.25
                                fmt.Println(v4, v3, v8)
                            }
                        default:
                            switch v3 {
                            case "a":
                                var v38 bool = (!(!(true && false)))
                                _ = v38
                            default:
                                v39 := GenS2{F0: 74, F1: 65535}
                                _ = v39
                            }
                        }
                        var v40 string = ("" + (("\n\t" + v3) + (v3 + "\n\t")))
                        _ = v40
                    } else {
                        var v41 uint64 = 108
                        _ = v41
                    }
                    v4 = 1e10
                }
                var v42 int64 = -9223372036854775808
                var v43 int = 257
                _ = v42
                _ = v43
            }
        }
        if (!true) {
            fmt.Println(v3)
            v44 := GenS5{F0: 2, F1: 2, F2: (-(v28))}
            _ = v44
        }
    }
    var v45 uint = 65535
    if (false && ((v45 % ((v45) | 1)) <= 18446744073709551615)) {
        fmt.Println(v28, v4, v45)
        delete(v6, -285)
    } else {
        v2[0] = 4294967295
        _ = len(v2)
    }
    var v46 string = v3
    if true {
        if (!(42 == v45)) {
            var v47 uint = v45
            if ((!("a" != "\n\t")) || false) {
                var v48 uint = (v45 + (^(v45)))
                _ = v48
            }
            _ = v47
        }
    }
    v45 = v45
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
