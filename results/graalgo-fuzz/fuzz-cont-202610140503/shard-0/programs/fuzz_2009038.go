package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 string; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int64; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool; F1 int64; F2 uint8 }
func genFn0() int8 {
    var v5 int64 = 1
    var v6 uint64 = 18446744073709551615
    v6 -= (v6 >> ((4294967295) & 63))
    v5 -= v5
    v7 := GenS1{F0: -2147483648, F1: 261}
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -671
    v3 := []int{(^(v2))}
    v4 := make(map[string]string)
    var v8 int16 = 0
    v8 += (v8 + 42)
    v9 := GenS1{F0: 1, F1: 0}
    var v10 bool = (0 == 42)
    if (((float64(float32((1e10 + 1.5))) + float64(3.14159)) / 3.14159) != (3.14159 * 0.0)) {
        v4["\n\t"] = "a"
        v2 -= (v2 * (v2 ^ -247))
    } else {
        var v11 uint32 = 42
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        if (0 < 82) {
            var v13 string = "\n\t"
            if ((-9223372036854775808 == 0) && v10) {
                v2 = (fi12 & fi12)
            }
            for fi14 := 0; fi14 < 18; fi14++ {
                if (855 <= 42) {
                    var v15 int8 = genFn0()
                    _ = v15
                } else {
                    var v16 bool = (557 >= 2)
                    var v17 int8 = -108
                    var v18 uint8 = 230
                    if v16 {
                        v16 = v16
                        if len(v3) > 0 {
                            v3[0] = int(v8)
                        }
                        fmt.Println(v8)
                    } else {
                        fmt.Println(v18)
                        switch "hello" {
                        case "\n\t":
                            var v19 int8 = (^(v17))
                            _ = v19
                        case "":
                            fmt.Println(v13)
                        default:
                            if (v16 && (-1 == 9223372036854775807)) {
                                v10 = (71 > 2147483648)
                                var v20 uint8 = 128
                                _ = v20
                            } else {
                                v21 := GenS2{F0: fi12}
                                fmt.Println(v18, v8, v8)
                                switch v13 {
                                case "hello":
                                    var v22 float32 = (3.14159 / (((3.14159 * float32((-2.25 / 1e10))) - 0.0) - float32((0.0 - ((1.5 * 0.1) / (1.5 / 1.0))))))
                                    _ = v22
                                case "":
                                    var v23 float64 = 3.14159
                                    _ = v23
                                default:
                                    v24 := GenS3{F0: v8, F1: "", F2: 255}
                                    _ = v24
                                }
                                _ = v21
                            }
                        }
                        delete(v1, fi14)
                        v4[v13] = "foo bar"
                    }
                    _ = v16
                    _ = v17
                    _ = v18
                }
                switch 128 {
                case 9223372036854775807:
                    fmt.Println(v8, v13, v13)
                case 200:
                    fmt.Println(v2)
                case -541:
                    switch "\n\t" {
                    case "\n\t":
                        v25 := v1[42]
                        _ = v25
                    case "":
                        var v26 uint = 0
                        _ = v26
                    case "a":
                        if (-2.25 != float64((((0.1 / float32(1e10)) / 1e10) - 1e10))) {
                            if true {
                                v3 = append(v3, fi14)
                                delete(v1, 0)
                                var v27 int64 = 811
                                _ = v27
                            } else {
                                v8 += (^(v8))
                                if (0 != 65535) {
                                    var v28 uint64 = 827
                                    if (-484 != 0) {
                                        fmt.Println(fi14, v13)
                                    } else {
                                        var v29 bool = false
                                        v30 := GenS4{F0: 101}
                                        var v31 int16 = v8
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                    _ = v28
                                }
                                var v32 string = (v13 + ("foo bar" + v13))
                                _ = v32
                            }
                            fmt.Println(v13, v2, v8)
                            var v33 uint64 = 435
                            _ = v33
                        }
                    default:
                        _ = len(v3)
                    }
                default:
                    fmt.Println(v2, v13)
                }
                _ = fi14
            }
            _ = v13
        } else {
            if v10 {
                for fi34 := 0; fi34 < 18; fi34++ {
                    var v35 int32 = -1
                    _ = len(v4)
                    _ = fi34
                    _ = v35
                }
            }
            _ = len(v1)
        }
        v2 -= fi12
        for fi36 := 0; fi36 < 9; fi36++ {
            if true {
                v37 := v1[fi36]
                var v38 float32 = float32(1e10)
                var v39 float64 = 0.0
                switch v37 {
                case "hello":
                    if (!v10) {
                        v40 := GenS5{F0: 42, F1: 9223372036854775807, F2: v37}
                        v41 := GenS3{F0: (v8 + (v8 + v8)), F1: v37, F2: 845}
                        _ = v40
                        _ = v41
                    }
                case "":
                    v39 = (v39 + v39)
                default:
                    v10 = (316 != 18446744073709551615)
                }
                _ = v37
                _ = v38
                _ = v39
            }
            _ = fi36
        }
        _ = fi12
    }
    switch ("" + ("\n\t" + "a")) {
    case "\n\t":
        v42 := GenS2{F0: (v2 % ((v2) | 1))}
        _ = v42
    case "hello":
        v43 := GenS6{F0: 555}
        _ = v43
    case "":
        fmt.Println(v2)
    }
    v10 = (v10 || false)
    if (v2 != 32767) {
        switch -2147483648 {
        case -2147483648:
            for fi44 := 0; fi44 < 17; fi44++ {
                if ((("hello" + "a") == "") || ((("\n\t" < "hello") || (65535 == 4294967295)) || (!(-149 <= 0)))) {
                    _ = len(v3)
                    switch v2 {
                    case 171:
                        var v45 int32 = -102
                        _ = v45
                    case -566:
                        if v10 {
                            var v46 float64 = 1e10
                            _ = v46
                        } else {
                            v2 += (^(fi44))
                        }
                    case 1:
                        if (!false) {
                            v47 := GenS3{F0: v8, F1: "a", F2: 758}
                            _ = v47
                        } else {
                            var v48 bool = (!false)
                            v8 += v8
                            _ = v48
                        }
                    }
                    switch v2 {
                    case 127:
                        if (0 == 1) {
                            delete(v1, (^(v2)))
                            _ = len(v1)
                        }
                    default:
                        v49 := v1[fi44]
                        _ = v49
                    }
                } else {
                    if (!(2 >= 0)) {
                        var v50 string = ""
                        fmt.Println(fi44, v2)
                        _ = v50
                    } else {
                        v8 = (v8 * (v8 % (((v8 ^ v8)) | 1)))
                        if v10 {
                            v8 = v8
                            if (2 > 42) {
                                var v51 int8 = -128
                                _ = v51
                            }
                        }
                        fmt.Println(v2, fi44, fi44)
                    }
                    var v52 uint32 = 1
                    if (uint64(v52) < 872) {
                        fmt.Println(v10, v52, v10)
                        switch (((((("" + "foo bar") + ("foo bar" + "foo bar")) + ("" + "")) + "hello") + "\n\t") + "") {
                        case "hello":
                            v53 := GenS1{F0: v2, F1: 2147483647}
                            _ = v53
                        case "a":
                            var v54 int16 = (^(v8))
                            _ = v54
                        case "\n\t":
                            v52 = v52
                        default:
                            _ = len(v1)
                        }
                        var v55 uint32 = 459
                        _ = v55
                    }
                    _ = v52
                }
                v56 := GenS7{F0: 100}
                _ = fi44
                _ = v56
            }
        case -598:
            if false {
                fmt.Println(v2, v10, v10)
                var v57 string = ("foo bar" + "")
                var v58 int16 = (-(v8))
                _ = v57
                _ = v58
            } else {
                for fi59 := 0; fi59 < 17; fi59++ {
                    if (-1 == -9223372036854775808) {
                        v3 = append(v3, 0)
                        var v60 uint16 = 255
                        v8 = v8
                        _ = v60
                    }
                    if (239 < 2147483647) {
                        for fi61 := 0; fi61 < 5; fi61++ {
                            v62 := GenS8{F0: true, F1: -9223372036854775808, F2: 2}
                            _ = fi61
                            _ = v62
                        }
                    }
                    _ = fi59
                }
            }
        case 127:
            fmt.Println(v2)
        }
        var v63 int16 = v8
        _ = v63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
