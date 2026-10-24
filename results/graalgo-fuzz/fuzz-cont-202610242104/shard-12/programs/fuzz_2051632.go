package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 string }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int8 }
func genFn0(p0 int, p1 int) int64 {
    fmt.Println(p0, p1)
    var v4 int16 = 995
    _ = p0
    _ = p1
    _ = v4
    return -184
}

func main() {
    var v1 string = (("" + ((("\n\t" + "") + "foo bar") + (("\n\t" + "foo bar") + ("a" + "foo bar")))) + "a")
    var v2 int8 = -1
    var v3 int = 32767
    v5 := []int{83}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, (v3 / ((-771) | 1)))
    switch (v3 + (v3 ^ (v3 * v3))) {
    case 0:
        v6 := GenS1{F0: 1, F1: v2, F2: v1}
        _ = v6
    default:
        v1 = (v1 + v1)
    }
    v1 = ("a" + (v1 + ((v1 + v1) + (v1 + v1))))
    fmt.Println(v1, v2, v1)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch (v3 % ((v3) | 1)) {
    case 1:
        v7 := GenS2{F0: 1, F1: "foo bar"}
        _ = v7
    default:
        var v8 uint16 = 65535
        _ = v8
    }
    v3 -= (v3 ^ (v3 * v3))
    switch (v3 / ((v3) | 1)) {
    case 1:
        if (((1e10 / 1.0) + 0.0) <= 0.1) {
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS3{F0: v2}
                _ = fi9
                _ = v10
            }
            if (!true) {
                for fi11 := 0; fi11 < 14; fi11++ {
                    if (false || (!true)) {
                        v5 = append(v5, v3)
                        v3 -= v3
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        fmt.Println(fi11)
                        if (8 == 255) {
                            v2 = v2
                            fmt.Println(fi11, fi11, v1)
                            for fi12 := 0; fi12 < 15; fi12++ {
                                var v13 string = "a"
                                v1 = (v1 + v1)
                                fmt.Println(v3, fi12)
                                _ = fi12
                                _ = v13
                            }
                        } else {
                            v1 = v1
                            switch v1 {
                            case "\n\t":
                                v3 += 4294967295
                            case "foo bar":
                                for fi14 := 0; fi14 < 20; fi14++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v15 string = v1
                                    if true {
                                        if false {
                                            var v16 int8 = (-(v2))
                                            _ = v16
                                        }
                                        v3 = -493
                                        var v17 int16 = -708
                                        var v18 uint32 = 1
                                        _ = v17
                                        _ = v18
                                    }
                                    _ = fi14
                                    _ = v15
                                }
                            case "":
                                v19 := GenS3{F0: v2}
                                _ = v19
                            default:
                                var v20 string = v1
                                _ = v20
                            }
                            v21 := GenS3{F0: (v2 >> ((2) & 63))}
                            _ = v21
                        }
                        if true {
                            switch v1 {
                            case "hello":
                                v1 = ("" + "hello")
                            case "":
                                if len(v5) > 0 {
                                    v5[0] = (fi11 % ((v3) | 1))
                                }
                            }
                            for fi22 := 0; fi22 < 13; fi22++ {
                                fmt.Println(fi22, v1, v1)
                                switch (v1 + v1) {
                                case "foo bar":
                                    var v23 int16 = 1
                                    _ = v23
                                case "a":
                                    if (-9223372036854775808 != -931) {
                                        var v24 bool = (1 == 1)
                                        fmt.Println(v24)
                                        _ = v24
                                    }
                                default:
                                    if false {
                                        var v25 uint32 = 2147483647
                                        v26 := GenS1{F0: 42, F1: (v2 << ((4294967295) & 63)), F2: v1}
                                        _ = v25
                                        _ = v26
                                    }
                                }
                                v27 := GenS2{F0: 1, F1: ((v1 + "\n\t") + v1)}
                                _ = fi22
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 15; fi28++ {
                                switch v1 {
                                case "\n\t":
                                    v29 := GenS4{F0: (!(-128 <= (-(v2)))), F1: v2}
                                    _ = v29
                                case "":
                                    switch v1 {
                                    case "hello":
                                        fmt.Println(fi28)
                                    case "a":
                                        v2 += v2
                                    default:
                                        v5 = append(v5, fi11)
                                    }
                                default:
                                    v2 -= -128
                                }
                                _ = fi28
                            }
                        }
                    }
                    _ = fi11
                }
                v1 = v1
            }
            var v30 uint32 = 2
            _ = v30
        }
    default:
        v2 += (v2 % (((v2 | v2)) | 1))
    }
    fmt.Println(v2, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
}
