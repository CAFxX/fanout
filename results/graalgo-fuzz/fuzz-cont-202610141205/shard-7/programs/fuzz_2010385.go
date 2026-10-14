package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint64; F1 bool; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int8, p1 int64, p2 bool) uint8 {
    var v7 uint16 = 32767
    var v8 float64 = 3.14159
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 float64 = v8
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 string = "a"
            v13 := []int{2147483647}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    v14 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := []int{127}
    v2 := make(map[string]int)
    var v3 uint16 = 781
    v4 := []string{("foo bar" + ("\n\t" + "\n\t")), "", "foo bar"}
    var v5 int = -9223372036854775808
    v6 := []int{}
    var v15 uint16 = (v3 % (((v3 + v3)) | 1))
    switch (^(v5)) {
    case 787:
        v16 := GenS1{F0: 0, F1: 42}
        _ = v16
    case -866:
        fmt.Println(v15)
    case 32768:
        var v17 uint16 = (v15 | v3)
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        _ = len(v1)
        var v19 int16 = 1
        var v20 string = "hello"
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v5)
    v15 += 32767
    v5 = (v5 ^ 0)
    if false {
        v21 := v2[("" + "")]
        var v22 bool = false
        _ = v21
        _ = v22
    }
    var v23 int8 = 100
    switch (v5 | 128) {
    case 4294967295:
        for fi24 := 0; fi24 < 11; fi24++ {
            if false {
                if ((280 != 32767) && ((0.0 / (float64(3.14159) * (((((-2.25 / (((3.14159 + (1e10 * (0.0 + -2.25))) + 3.14159) + 0.1)) - (((-2.25 + 0.0) / 1.0) + 0.0)) + (0.1 * (-2.25 + 0.1))) / 0.1) - float64(((float32((1.5 - 1.5)) / 1.0) / ((((1e10 / (3.14159 - float32(((-2.25 * float64(float32(3.14159))) + 1.5)))) * (3.14159 * ((3.14159 * 3.14159) / 1.0))) * -2.25) / (-2.25 - float32((((0.1 + (-2.25 + 0.0)) / (-2.25 + 1e10)) - float64(3.14159)))))))))) <= 1e10)) {
                    v25 := GenS1{F0: 600, F1: -149}
                    _ = v25
                } else {
                    v5 += (v5 | (v5 << ((18446744073709551615) & 63)))
                    for fi26 := 0; fi26 < 17; fi26++ {
                        if true {
                            var v27 int16 = 42
                            if true {
                                switch int(v27) {
                                case 4294967295:
                                    v3 -= v15
                                default:
                                    fmt.Println(v3, v27)
                                }
                                v1 = append(v1, (fi26 % (((fi26 % ((-226) | 1))) | 1)))
                            }
                            v28 := GenS1{F0: 292, F1: 929}
                            _ = v27
                            _ = v28
                        } else {
                            if true {
                                if false {
                                    v29 := GenS2{F0: 397, F1: true, F2: fi24}
                                    switch "hello" {
                                    case "\n\t":
                                        if false {
                                            _ = len(v2)
                                            fmt.Println(fi26)
                                        } else {
                                            fmt.Println(v23, v15, fi24)
                                            v30 := GenS3{F0: 2}
                                            _ = v30
                                        }
                                    default:
                                        if (false || (genFn0((v23 ^ 42), 9223372036854775807, true) < 42)) {
                                            fmt.Println(v15, v23)
                                            switch ("a" + "hello") {
                                            case "hello":
                                                var v31 string = ("a" + ("\n\t" + "a"))
                                                _ = v31
                                            case "\n\t":
                                                v32 := GenS2{F0: 2, F1: false, F2: (^(v5))}
                                                v33 := v2["foo bar"]
                                                switch "a" {
                                                case "":
                                                    if true {
                                                        v33 = 32768
                                                    }
                                                case "a":
                                                    v23 = (^(v23))
                                                default:
                                                    var v34 uint = 42
                                                    _ = v34
                                                }
                                                _ = v32
                                                _ = v33
                                            case "a":
                                                var v35 int16 = -1
                                                _ = v35
                                            }
                                            switch "" {
                                            case "hello":
                                                v36 := GenS3{F0: 661}
                                                _ = v36
                                            case "foo bar":
                                                v15 = (v3 & (v3 / (((v15 >> ((907) & 63))) | 1)))
                                            case "\n\t":
                                                v2["a"] = (fi26 / (((v5 % ((823) | 1))) | 1))
                                                v2[("" + ("" + ""))] = 128
                                            default:
                                                if true {
                                                    v5 = 112
                                                    var v37 int = fi24
                                                    _ = v37
                                                }
                                            }
                                        } else {
                                            var v38 int32 = 1
                                            v23 += int8(v15)
                                            var v39 uint64 = 492
                                            switch (("" + "") + "") {
                                            case "a":
                                                var v40 int32 = (v38 / ((0) | 1))
                                                _ = v40
                                            case "foo bar":
                                                if false {
                                                    v3 -= 65535
                                                } else {
                                                    switch (fi24 ^ (^(fi26))) {
                                                    case 128:
                                                        var v41 uint = 255
                                                        _ = v41
                                                    }
                                                    var v42 int64 = -1
                                                    _ = v42
                                                }
                                            case "\n\t":
                                                fmt.Println(fi24, fi26, v23)
                                            }
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                    v43 := GenS4{F0: "", F1: (127 > 56)}
                                    _ = v29
                                    _ = v43
                                }
                            }
                            if (("\n\t" + ((("foo bar" + "\n\t") + "a") + "a")) != ("foo bar" + "foo bar")) {
                                var v44 uint16 = 255
                                v3 -= v44
                                _ = v44
                            }
                        }
                        _ = fi26
                    }
                }
                v23 += v23
            }
            _ = fi24
        }
    case -9223372036854775808:
        var v45 int8 = -45
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
