package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int; F2 int64 }
type GenS4 struct { F0 bool; F1 uint8; F2 bool }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint64, p2 uint16) int {
    p0 += 0
    v6 := []int{-9223372036854775808, 4294967295, 128, 9223372036854775807}
    var v7 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 uint8) string {
    p0 -= 2
    var v8 int16 = 42
    _ = p0
    _ = p1
    _ = v8
    return "\n\t"
}

func main() {
    var v1 uint16 = 0
    var v2 string = "hello"
    v3 := []string{(((v2 + ((v2 + v2) + v2)) + "hello") + v2), (v2 + (v2 + ("\n\t" + ((((v2 + (v2 + ((v2 + v2) + v2))) + v2) + (v2 + ("" + ""))) + v2)))), v2, (v2 + (v2 + v2))}
    v4 := make(map[string]bool)
    var v5 float64 = ((float64(1.5) + 1.5) - 0.0)
    v1 -= (v1 - 42)
    v5 = v5
    v5 = v5
    for fi9 := 0; fi9 < 8; fi9++ {
        v5 -= (float64((float32(v5) + -2.25)) + 0.1)
        fmt.Println(fi9, v1, fi9)
        for fi10 := 0; fi10 < 1; fi10++ {
            switch ("foo bar" + v2) {
            case "foo bar":
                v1 = (v1 / ((v1) | 1))
            case "\n\t":
                v1 += 501
            default:
                if (1 == 65535) {
                    var v11 uint16 = v1
                    if (!false) {
                        switch -352 {
                        case -261:
                            fmt.Println(v2, v2, v2)
                        case 32767:
                            v2 = "\n\t"
                        case 0:
                            switch genFn0(2147483647, 42, v1) {
                            case -1:
                                if (-624 >= 42) {
                                    if true {
                                        if (1 > v1) {
                                            if len(v3) > 0 {
                                                v3[0] = ((v2 + v2) + v2)
                                            }
                                            fmt.Println(v11, v2, v1)
                                            if (!(v2 < "hello")) {
                                                v11 = v1
                                            }
                                        }
                                        v1 = 255
                                    } else {
                                        if (!false) {
                                            switch "" {
                                            case "":
                                                var v12 int = (fi9 * (fi10 * genFn0(1, 530, v11)))
                                                _ = v12
                                            }
                                            fmt.Println(fi10, v11)
                                            v1 = v1
                                        } else {
                                            v5 -= v5
                                            v13 := GenS1{F0: 1}
                                            switch "" {
                                            case "a":
                                                fmt.Println(fi10, v1)
                                            case "hello":
                                                v14 := GenS2{F0: 903, F1: v11, F2: 996}
                                                v5 = v5
                                                _ = v14
                                            case "foo bar":
                                                if (false && false) {
                                                    v3 = append(v3, "foo bar")
                                                    delete(v4, (v2 + v2))
                                                    if len(v3) > 0 {
                                                        v3[0] = v2
                                                    }
                                                    var v15 int = (^(fi10))
                                                    switch (("" + v2) + v2) {
                                                    case "foo bar":
                                                        var v16 int16 = 1000
                                                        _ = v16
                                                    case "\n\t":
                                                        v4[(((v2 + "a") + (v2 + "a")) + v2)] = (1 < 638)
                                                    case "":
                                                        if true {
                                                            v11 = 1
                                                            fmt.Println(fi10)
                                                            v17 := GenS3{F0: (v11 << ((805) & 63)), F1: 128, F2: -9223372036854775808}
                                                            _ = v17
                                                        }
                                                    default:
                                                        if ((3.14159 + v5) < v5) {
                                                            if (!false) {
                                                                v15 = 32768
                                                                if (-1 == 42) {
                                                                    fmt.Println(fi9)
                                                                    v15 = v15
                                                                }
                                                                if (!true) {
                                                                    fmt.Println(v1)
                                                                    v15 = 2147483647
                                                                    var v18 float32 = (3.14159 * 1.5)
                                                                    _ = v18
                                                                }
                                                            } else {
                                                                var v19 uint16 = (v11 * 2)
                                                                _ = v19
                                                            }
                                                            if ((!((127 < 0) && ((!false) && false))) || true) {
                                                                var v20 float32 = 0.0
                                                                _ = v20
                                                            } else {
                                                                fmt.Println(v2, v1, v5)
                                                                var v21 bool = (860 <= 65535)
                                                                _ = v21
                                                            }
                                                            var v22 int = (v15 * (fi10 >> ((4294967295) & 63)))
                                                            _ = v22
                                                        }
                                                    }
                                                    _ = v15
                                                }
                                            default:
                                                v1 = (v1 ^ v1)
                                            }
                                            _ = v13
                                        }
                                        v2 = ((v2 + "foo bar") + "hello")
                                        v23 := GenS4{F0: true, F1: 2, F2: true}
                                        _ = v23
                                    }
                                    v11 = (v1 | 690)
                                    var v24 uint = 4294967295
                                    if false {
                                        if len(v3) > 0 {
                                            v3[0] = v2
                                        }
                                        switch 1 {
                                        case 128:
                                            v1 += v1
                                        case 0:
                                            var v25 bool = (1.5 != (float32(float64(1e10)) + (1.5 - -2.25)))
                                            _ = v25
                                        case -2147483648:
                                            v26 := GenS5{F0: "hello", F1: 1}
                                            _ = v26
                                        }
                                        if (32767 < -951) {
                                            if len(v3) > 0 {
                                                v3[0] = (v2 + v2)
                                            }
                                        } else {
                                            fmt.Println(fi9)
                                        }
                                    }
                                    v5 += (float64(0.0) + 0.1)
                                    _ = v24
                                }
                            default:
                                v27 := GenS6{F0: 0}
                                _ = v27
                            }
                            var v28 float32 = (3.14159 - float32(((v5 - v5) / float64((0.1 / 1e10)))))
                            _ = v28
                        }
                    }
                    _ = v11
                } else {
                    var v29 uint8 = 228
                    _ = v29
                }
            }
            v30 := GenS4{F0: (!false), F1: 42, F2: (!((244 > 1) && true))}
            var v31 uint8 = 0
            _ = fi10
            _ = v30
            _ = v31
        }
        _ = fi9
    }
    if true {
        if (genFn0(0, 18446744073709551615, v1) < 32767) {
            if (42 > 834) {
                if true {
                    v3 = append(v3, (("" + "hello") + (v2 + v2)))
                }
                v2 = v2
                v5 += 1e10
            }
            v2 = v2
            v32 := GenS7{F0: 935}
            _ = v32
        }
        switch v2 {
        case "hello":
            var v33 uint8 = 42
            _ = v33
        }
    }
    switch v2 {
    case "":
        v4[""] = false
    case "\n\t":
        var v34 uint = 237
        _ = v34
    default:
        v35 := v4[(("hello" + ("\n\t" + "")) + v2)]
        _ = v35
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v5 -= (-2.25 + v5)
        v1 -= (v1 | 380)
        _ = fi36
    }
    if (2147483647 <= 1) {
        fmt.Println(v5, v2, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
