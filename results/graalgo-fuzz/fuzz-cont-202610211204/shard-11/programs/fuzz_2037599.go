package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 bool }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 bool; F1 string; F2 string }
type GenS6 struct { F0 uint8; F1 string; F2 uint16 }
func genFn0(p0 int64, p1 int64) uint16 {
    p0 = p1
    p0 -= p1
    v7 := GenS1{F0: 127, F1: 2, F2: (!(((!false) || (!(!(42 == 327)))) && false))}
    var v8 float32 = (0.0 + ((3.14159 * -2.25) - 1.5))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint32, p1 uint8, p2 uint32) uint32 {
    v9 := GenS2{F0: (p1 - 1), F1: 32767, F2: (((("" + "\n\t") + "\n\t") + "a") + "foo bar")}
    var v10 uint8 = (p1 * 2)
    p1 += p1
    fmt.Println(v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (((0.0 / 1e10) - -2.25) + (1.5 + 1e10))
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    var v5 bool = false
    var v6 uint = 65535
    v11 := GenS3{F0: 1}
    v2 += -2.25
    v6 = 494
    for fi12 := 0; fi12 < 15; fi12++ {
        if (1 >= 42) {
            v6 = 42
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            v4[42] = "\n\t"
            _ = fi13
        }
        _ = fi12
    }
    v14 := GenS4{F0: 2147483647, F1: ""}
    if true {
        switch 32767 {
        case -2147483648:
            switch -2147483648 {
            case -9223372036854775808:
                v15 := GenS1{F0: 896, F1: 2, F2: ((!v5) || v5)}
                _ = v15
            case 2147483647:
                for fi16 := 0; fi16 < 1; fi16++ {
                    v2 -= (((1.5 / v2) / ((v2 - float64(((-2.25 / -2.25) * float32(float64((((3.14159 + 3.14159) * 0.1) / 1.0)))))) - v2)) - (0.0 - v2))
                    _ = fi16
                }
            default:
                if (284 != 2) {
                    if v5 {
                        switch "\n\t" {
                        case "a":
                            v5 = (false && ("\n\t" >= ("" + "\n\t")))
                        }
                    }
                    if v5 {
                        v5 = (526 <= genFn0(9223372036854775807, 312))
                        switch 353 {
                        case -1:
                            for fi17 := 0; fi17 < 17; fi17++ {
                                var v18 string = ("" + ((("foo bar" + ("foo bar" + "\n\t")) + "") + "hello"))
                                _ = fi17
                                _ = v18
                            }
                        }
                        var v19 float32 = ((0.1 + -2.25) + float32((((float64((0.1 * 0.1)) - (v2 - 3.14159)) + 0.0) / ((float64((1.5 + 0.1)) - ((0.1 * 0.0) + 3.14159)) * ((v2 + v2) / 1.5)))))
                        _ = v19
                    }
                }
            }
        }
    } else {
        v20 := GenS4{F0: 350, F1: ("\n\t" + "\n\t")}
        v21 := GenS4{F0: -957, F1: "a"}
        v2 += v2
        _ = v20
        _ = v21
    }
    fmt.Println(v5, v2)
    var v22 int = 522
    var v23 bool = v5
    if (!v23) {
        v24 := []int{v22, v22}
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 int = (v22 << ((4294967295) & 63))
            v2 = ((0.1 - (1e10 + v2)) * (((0.1 + v2) / (float64(3.14159) / ((v2 / 1.0) / (v2 * float64((((float32(-2.25) / 0.1) + 0.1) / 1.5)))))) / (v2 + (v2 / 3.14159))))
            for fi27 := 0; fi27 < 6; fi27++ {
                switch fi25 {
                case 42:
                    if true {
                        v28 := GenS4{F0: 719, F1: "\n\t"}
                        var v29 uint32 = 2147483647
                        var v30 bool = ((v29 << (((v6 >> ((v6) & 63))) & 63)) < 65535)
                        v29 += 42
                        v2 = v2
                        switch -137 {
                        case 0:
                            v23 = (!("foo bar" == ""))
                        case 42:
                            var v31 uint16 = 42
                            _ = v31
                        case -509:
                            var v32 float32 = 0.1
                            _ = v32
                        }
                        var v33 float64 = 3.14159
                        var v34 uint64 = 2
                        _ = v28
                        _ = v29
                        _ = v30
                        _ = v33
                        _ = v34
                    } else {
                        var v35 uint64 = 4294967295
                        if true {
                            if v5 {
                                if (0.0 >= 1e10) {
                                    v5 = v5
                                    if v23 {
                                        v5 = (v5 && true)
                                        v35 += v35
                                    } else {
                                        v36 := GenS5{F0: v5, F1: "", F2: ("foo bar" + "")}
                                        _ = v36
                                    }
                                    if len(v24) > 0 {
                                        v24[0] = (v26 & -9223372036854775808)
                                    }
                                }
                                switch (v26 % (((fi25 * fi25)) | 1)) {
                                case 32768:
                                    var v37 int8 = 100
                                    _ = v37
                                case 127:
                                    _ = len(v4)
                                }
                            }
                        } else {
                            fmt.Println(v35)
                        }
                        v2 -= -2.25
                        _ = v35
                    }
                case 128:
                    var v38 int16 = -119
                    _ = v38
                default:
                    if false {
                        switch v22 {
                        case 2147483647:
                            v23 = v23
                        case 128:
                            fmt.Println(fi27, v2)
                        case -1:
                            v2 += 3.14159
                        default:
                            if false {
                                _ = len(v4)
                            } else {
                                var v39 uint8 = 42
                                v40 := v3[(fi25 & (fi27 << ((18446744073709551615) & 63)))]
                                _ = v39
                                _ = v40
                            }
                        }
                        v22 += 2147483647
                        v41 := GenS6{F0: 80, F1: (((((("a" + ("" + "\n\t")) + "hello") + ("" + "a")) + "foo bar") + "foo bar") + "hello"), F2: 8}
                        _ = v41
                    }
                }
                v26 = fi27
                _ = fi27
            }
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    fmt.Println(v2, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
