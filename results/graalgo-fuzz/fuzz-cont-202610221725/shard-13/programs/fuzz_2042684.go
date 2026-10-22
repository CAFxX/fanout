package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint; F2 string }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int8; F1 int8; F2 int16 }
type GenS6 struct { F0 uint8 }
func genFn0() int8 {
    var v5 bool = true
    if (true || v5) {
        switch "\n\t" {
        case "foo bar":
            var v6 float32 = float32(1.5)
            _ = v6
        case "":
            var v7 string = ("foo bar" + "hello")
            _ = v7
        case "\n\t":
            v8 := []int{-944, -375, -849, 691}
            _ = v8
        }
    } else {
        v9 := []int{-2147483648, 32767, 127, 32768}
        _ = v9
    }
    _ = v5
    return 83
}
func genFn1(p0 int8) uint {
    var v10 bool = false
    var v11 int64 = 768
    v11 -= v11
    switch ("" + "") {
    case "foo bar":
        v12 := GenS1{F0: "a", F1: 2147483647, F2: -457}
        _ = v12
    case "hello":
        v13 := []int{30, -9223372036854775808, -338}
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    return 483
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = ((float64(0.1) + (-2.25 / 1.5)) + 0.1)
    v3 := make(map[string]int)
    var v4 float64 = 3.14159
    v14 := GenS1{F0: ("hello" + (("\n\t" + "a") + "a")), F1: 127, F2: 2147483647}
    var v15 int = 641
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 string = "\n\t"
        v18 := GenS2{F0: 42}
        _ = fi16
        _ = v17
        _ = v18
    }
    v15 = (v15 << ((1) & 63))
    v19 := v1[""]
    fmt.Println(v4)
    var v20 int16 = -32768
    if (0 == 1) {
        v15 -= v15
        var v21 float32 = 0.1
        if (2 >= 103) {
            switch ("foo bar" + (("a" + "foo bar") + ("" + ("\n\t" + "a")))) {
            case "a":
                for fi22 := 0; fi22 < 3; fi22++ {
                    var v23 int32 = 42
                    fmt.Println(v2, v19, v20)
                    v4 += v4
                    _ = fi22
                    _ = v23
                }
            }
        }
        _ = v21
    } else {
        v19 = (true && (866 <= 65535))
        v24 := v1["a"]
        var v25 float64 = 0.0
        _ = v24
        _ = v25
    }
    v19 = (!((v15 * v15) < v15))
    v26 := GenS3{F0: (v20 % ((32767) | 1)), F1: 65535, F2: "hello"}
    fmt.Println(v4)
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := GenS4{F0: 2147483647, F1: "foo bar"}
        switch "\n\t" {
        case "foo bar":
            v4 = (v4 + ((float64(0.1) * v2) + ((v4 - (0.0 / v4)) / (0.0 + v4))))
        }
        var v29 int32 = 2147483647
        _ = fi27
        _ = v28
        _ = v29
    }
    v20 -= -1
    if (0.0 <= -2.25) {
        var v30 bool = (128 < 42)
        switch v15 {
        case 1:
            fmt.Println(v4, v4)
        case -1:
            v1["a"] = v19
        case 9223372036854775807:
            switch "foo bar" {
            case "\n\t":
                if v19 {
                    if true {
                        _ = len(v1)
                        var v31 uint8 = 127
                        v32 := []int{127, v15, 8, (v15 % ((-220) | 1))}
                        _ = v31
                        _ = v32
                    }
                    var v33 uint8 = 239
                    _ = v33
                } else {
                    switch "\n\t" {
                    case "a":
                        v19 = v30
                    case "hello":
                        v34 := GenS2{F0: 713}
                        _ = v34
                    default:
                        var v35 uint = 956
                        _ = v35
                    }
                }
            case "":
                for fi36 := 0; fi36 < 17; fi36++ {
                    if true {
                        for fi37 := 0; fi37 < 20; fi37++ {
                            if v30 {
                                var v38 uint64 = 4294967295
                                fmt.Println(fi37)
                                v15 = (^(fi37))
                                _ = v38
                            }
                            v39 := GenS5{F0: 1, F1: 1, F2: (^(v20))}
                            var v40 uint16 = 0
                            var v41 int = 76
                            if v30 {
                                switch (v15 - 42) {
                                case -2147483648:
                                    _ = len(v3)
                                case -1:
                                    v42 := []int{66}
                                    _ = v42
                                default:
                                    fmt.Println(v15, v2)
                                    var v43 uint64 = 1
                                    if v19 {
                                        v44 := GenS2{F0: uint(v43)}
                                        v45 := []int{748, fi36, (v15 | (v15 << ((833) & 63))), (v15 & fi36)}
                                        _ = v44
                                        _ = v45
                                    } else {
                                        v19 = ((float64(0.1) != v4) || true)
                                        delete(v3, "a")
                                        v46 := GenS4{F0: 801, F1: ("\n\t" + "foo bar")}
                                        _ = v46
                                    }
                                    _ = v43
                                }
                                var v47 float64 = v2
                                _ = v47
                            } else {
                                var v48 float64 = float64(((0.0 + 1.5) + 3.14159))
                                v49 := []int{v41, fi36}
                                _ = v48
                                _ = v49
                            }
                            _ = fi37
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        switch (^(v15)) {
                        case 2147483647:
                            v3[("" + ("" + "hello"))] = (v15 / (((fi36 ^ fi36)) | 1))
                        }
                    }
                    if v30 {
                        switch (-(fi36)) {
                        case 128:
                            if (((162 < 108) || (false && v30)) || ((v20 | (v20 << ((255) & 63))) >= (v20 >> ((4294967295) & 63)))) {
                                var v50 int32 = 2147483647
                                _ = v50
                            } else {
                                for fi51 := 0; fi51 < 6; fi51++ {
                                    v19 = ((!(!(255 < 169))) || (1.5 < 3.14159))
                                    v52 := v3[(("\n\t" + "foo bar") + "foo bar")]
                                    _ = fi51
                                    _ = v52
                                }
                                for fi53 := 0; fi53 < 13; fi53++ {
                                    if true {
                                        var v54 uint = 493
                                        v55 := GenS6{F0: 128}
                                        v3["foo bar"] = fi53
                                        _ = v54
                                        _ = v55
                                    }
                                    v2 += v4
                                    _ = fi53
                                }
                            }
                        case 332:
                            var v56 bool = (!v30)
                            _ = v56
                        case 551:
                            var v57 float32 = 0.1
                            _ = v57
                        }
                    }
                    _ = fi36
                }
            }
        default:
            v58 := GenS5{F0: genFn0(), F1: 100, F2: v20}
            _ = v58
        }
        delete(v1, "a")
        _ = v30
    }
    v19 = ((!(1 == 2)) && (2 >= 32767))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
}
