package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int16; F1 uint16; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint16; F2 int8 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 string }
func genFn0() int8 {
    v7 := GenS1{F0: 32767, F1: 100}
    var v8 int16 = 0
    if (float32(0.0) == 3.14159) {
        v9 := GenS1{F0: 42, F1: 102}
        _ = v9
    } else {
        v10 := []int{-24}
        fmt.Println(v8)
        _ = v10
    }
    v11 := GenS2{F0: v8, F1: 2, F2: -9223372036854775808}
    v8 = v8
    _ = v11
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 bool) uint64 {
    if (!p0) {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint = 255
            _ = fi12
            _ = v13
        }
        fmt.Println(p0)
        v14 := GenS2{F0: 32767, F1: 32767, F2: 0}
        _ = v14
    } else {
        v15 := GenS3{F0: ("a" + "\n\t")}
        var v16 int64 = -9223372036854775808
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            p0 = (((0.1 < 0.1) && (!(-1 != 9223372036854775807))) && ((0.1 + 3.14159) >= ((-2.25 + 0.0) + (((1.5 + 0.0) + 1.5) / (float32(float64(-2.25)) * 0.1)))))
            _ = fi18
        }
        fmt.Println(fi17)
        _ = fi17
    }
    if (65535 < 1) {
        var v19 float64 = -2.25
        var v20 uint16 = 526
        v21 := []int{}
        _ = v19
        _ = v20
        _ = v21
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 255
    v3 := []uint32{0, 42, 342, 2}
    var v4 int64 = 0
    v5 := []string{""}
    var v6 float32 = 1e10
    var v22 int16 = 1
    var v23 int = 944
    v24 := v1[(^(v23))]
    v24 += 128
    v6 += 1e10
    if true {
        v24 = v24
        v24 = (v23 % ((-2147483648) | 1))
        for fi25 := 0; fi25 < 19; fi25++ {
            v2 = (^(v2))
            v1[(v24 & (v24 ^ v23))] = 2147483647
            for fi26 := 0; fi26 < 8; fi26++ {
                if (42 < 876) {
                    if true {
                        v27 := GenS4{F0: false, F1: (^(v2)), F2: -53}
                        v28 := GenS5{F0: (v22 * v22), F1: -9223372036854775808}
                        v6 = ((float32((0.1 * 1e10)) / 1.0) - v6)
                        _ = v27
                        _ = v28
                    } else {
                        switch ("foo bar" + "") {
                        case "a":
                            switch v24 {
                            case 128:
                                var v29 int32 = -1
                                _ = v29
                            default:
                                v30 := GenS1{F0: fi26, F1: 16}
                                if ((4294967295 <= 2147483647) || false) {
                                    v6 = v6
                                    v6 = v6
                                    var v31 int64 = (v4 - 42)
                                    _ = v31
                                } else {
                                    if false {
                                        v32 := GenS2{F0: (v22 | -1), F1: v2, F2: v4}
                                        var v33 int = fi25
                                        fmt.Println(v33)
                                        _ = v32
                                        _ = v33
                                    }
                                    fmt.Println(v22, v22, v6)
                                    var v34 float32 = ((v6 - ((v6 * ((((-2.25 / ((v6 * v6) / v6)) * ((-2.25 - v6) / v6)) - v6) - float32(float64((v6 * (1.5 - 0.1)))))) + ((v6 - (v6 * (float32(1e10) - v6))) - v6))) * float32(float64((((float32((0.0 - (0.0 - (3.14159 * (((3.14159 - ((1e10 - 0.1) - (3.14159 - 1e10))) + ((-2.25 / 0.1) * 0.1)) - (-2.25 / 1e10)))))) + 1e10) - 3.14159) * v6))))
                                    _ = v34
                                }
                                _ = v30
                            }
                        case "foo bar":
                            var v35 uint64 = 540
                            _ = v35
                        case "hello":
                            switch 961 {
                            case 138:
                                fmt.Println(v2, v2, v6)
                            case 226:
                                if (0 < -2147483648) {
                                    v36 := GenS3{F0: "foo bar"}
                                    fmt.Println(v23, v23, v2)
                                    _ = v36
                                } else {
                                    v23 = (fi25 >> ((18446744073709551615) & 63))
                                }
                            case -9223372036854775808:
                                v37 := GenS2{F0: (v22 / ((-32768) | 1)), F1: v2, F2: (-(v4))}
                                v6 -= 3.14159
                                var v38 int32 = 906
                                v2 = v2
                                var v39 int16 = (v22 - (v22 / ((-614) | 1)))
                                v39 = (v22 % (((v39 << ((0) & 63))) | 1))
                                var v40 float64 = 0.0
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v41 uint16 = 425
                                fmt.Println(v22)
                                v42 := GenS3{F0: ""}
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                                _ = v41
                                _ = v42
                            default:
                                if (!true) {
                                    switch ("a" + "\n\t") {
                                    case "\n\t":
                                        if (743 >= 516) {
                                            var v43 int32 = -727
                                            if (3.14159 > ((1.5 / (0.0 * v6)) / v6)) {
                                                switch (v23 + (v24 ^ fi26)) {
                                                case 128:
                                                    v6 -= v6
                                                case 1:
                                                    if ((!true) || true) {
                                                        if true {
                                                            v22 += 42
                                                            var v44 uint32 = 312
                                                            _ = v44
                                                        }
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v45 int16 = -316
                                                            switch ("a" + "a") {
                                                            case "hello":
                                                                v1[4294967295] = (fi25 >> ((1) & 63))
                                                            case "a":
                                                                v43 = 1
                                                            }
                                                            _ = v45
                                                        default:
                                                            switch ("a" + ("hello" + ("foo bar" + "\n\t"))) {
                                                            case "foo bar":
                                                                v22 += (v22 / ((-673) | 1))
                                                            case "a":
                                                                v46 := GenS1{F0: v23, F1: -128}
                                                                _ = v46
                                                            case "\n\t":
                                                                v43 = v43
                                                            }
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    } else {
                                                        fmt.Println(v2, v23, fi25)
                                                        v47 := GenS6{F0: ""}
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        fmt.Println(fi26, fi26)
                                                        var v48 string = "a"
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                default:
                                                    var v49 float64 = float64(v6)
                                                    _ = v49
                                                }
                                            }
                                            _ = v43
                                        }
                                        v50 := GenS5{F0: (v22 ^ (v22 << ((1) & 63))), F1: (fi26 * 2147483647)}
                                        delete(v1, v23)
                                        _ = v50
                                    case "foo bar":
                                        var v51 int16 = (v22 + (v22 & v22))
                                        _ = v51
                                    }
                                    if (4294967295 != 2147483648) {
                                        var v52 int32 = 0
                                        var v53 int8 = 127
                                        _ = v52
                                        _ = v53
                                    }
                                }
                            }
                        }
                    }
                    var v54 uint16 = (v2 - 42)
                    var v55 int32 = -2147483648
                    _ = v54
                    _ = v55
                } else {
                    v6 = (1e10 * ((v6 - 0.0) + 1e10))
                }
                v56 := v1[fi26]
                _ = fi26
                _ = v56
            }
            _ = fi25
        }
    } else {
        v2 = (^(v2))
    }
    v57 := GenS2{F0: (v22 - (^(v22))), F1: (v2 * v2), F2: 42}
    v22 -= (v22 | v22)
    v6 += (float32(-2.25) - float32(3.14159))
    v24 = (v24 >> ((42) & 63))
    fmt.Println(v4, v6)
    v2 -= (v2 - (v2 + v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
