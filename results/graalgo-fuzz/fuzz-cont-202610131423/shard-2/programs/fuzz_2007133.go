package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 string }
func genFn0(p0 int16) int32 {
    var v5 string = ("foo bar" + "")
    v6 := GenS1{F0: (false && (!(-719 >= 4294967295))), F1: 682, F2: ("a" + "")}
    var v7 int64 = -1
    v7 = (v7 - 0)
    fmt.Println(v7, p0)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint) uint32 {
    v8 := []int{981, 32768, 32767}
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 127
    var v3 float32 = -2.25
    var v4 float32 = v3
    v9 := GenS1{F0: true, F1: 48, F2: "hello"}
    var v10 uint8 = 0
    switch "foo bar" {
    case "":
        var v11 float64 = 1e10
        _ = v11
    case "hello":
        fmt.Println(v10, v2, v4)
    default:
        if (-1 < 42) {
            for fi12 := 0; fi12 < 11; fi12++ {
                switch fi12 {
                case 42:
                    var v13 int32 = 0
                    _ = v13
                default:
                    var v14 uint64 = 1
                    _ = v14
                }
                switch (fi12 + (fi12 - fi12)) {
                case 974:
                    v15 := []int{-9223372036854775808, fi12, (fi12 / (((fi12 >> ((18446744073709551615) & 63))) | 1))}
                    _ = v15
                case 2147483647:
                    if false {
                        v3 = (0.0 / (v3 + float32(0.1)))
                        fmt.Println(fi12, v3)
                    }
                default:
                    v16 := []int{2147483647}
                    _ = v16
                }
                _ = fi12
            }
        } else {
            if false {
                fmt.Println(v2, v4)
            }
            if true {
                v17 := GenS2{F0: 42, F1: false}
                if (1 == 4294967295) {
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        if (1 < 1) {
                            v18 := GenS1{F0: false, F1: 42, F2: "a"}
                            v19 := v1[42]
                            _ = v18
                            _ = v19
                        } else {
                            var v20 float32 = (v4 * v4)
                            v21 := GenS1{F0: (true || true), F1: 42, F2: "foo bar"}
                            if (3.14159 != float64((3.14159 / (v3 / v20)))) {
                                v10 -= 191
                                v22 := GenS3{F0: v2, F1: 42}
                                _ = v22
                            }
                            _ = v20
                            _ = v21
                        }
                    default:
                        _ = len(v1)
                    }
                    switch ("a" + ("a" + "hello")) {
                    case "\n\t":
                        v4 -= -2.25
                    }
                    if ("\n\t" >= "a") {
                        for fi23 := 0; fi23 < 13; fi23++ {
                            fmt.Println(fi23, v10, v10)
                            for fi24 := 0; fi24 < 6; fi24++ {
                                fmt.Println(fi24)
                                v25 := []int{42, fi24, (fi24 | (fi24 % ((fi23) | 1))), fi23}
                                switch "hello" {
                                case "":
                                    if (!true) {
                                        v1[991] = true
                                        if false {
                                            v26 := GenS1{F0: true, F1: 18446744073709551615, F2: ("hello" + "\n\t")}
                                            var v27 uint64 = 2
                                            if true {
                                                switch (fi23 & (fi24 + (fi24 & (^(fi24))))) {
                                                case 9223372036854775807:
                                                    var v28 int64 = 42
                                                    _ = v28
                                                case -681:
                                                    _ = len(v25)
                                                case 32768:
                                                    switch 2147483647 {
                                                    case -365:
                                                        v27 -= 845
                                                    }
                                                }
                                                var v29 int16 = -439
                                                var v30 uint32 = 2147483647
                                                v31 := GenS3{F0: v2, F1: v27}
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            } else {
                                                switch "hello" {
                                                case "hello":
                                                    var v32 float64 = (((1.5 / 1e10) / 0.1) + float64(v4))
                                                    v2 += v2
                                                    _ = v32
                                                case "":
                                                    var v33 uint64 = (v27 - uint64(v10))
                                                    _ = v33
                                                default:
                                                    var v34 uint64 = (v27 - (v27 ^ 194))
                                                    v35 := GenS3{F0: (v2 << ((767) & 63)), F1: 4294967295}
                                                    var v36 string = "a"
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                                switch ("" + ("\n\t" + "")) {
                                                case "a":
                                                    var v37 uint = 42
                                                    _ = v37
                                                }
                                            }
                                            _ = v26
                                            _ = v27
                                        } else {
                                            fmt.Println(v4, fi24)
                                            if (0 == 118) {
                                                var v38 int64 = 42
                                                v4 += 0.1
                                                _ = v38
                                            }
                                            fmt.Println(v4)
                                        }
                                        if (true || true) {
                                            v39 := v1[68]
                                            _ = v39
                                        }
                                    }
                                case "\n\t":
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                }
                                _ = fi24
                                _ = v25
                            }
                            _ = fi23
                        }
                        v10 = 1
                    } else {
                        v40 := []int{-319, -95, 32767}
                        fmt.Println(v4)
                        if ((v2 ^ 100) >= 127) {
                            var v41 uint64 = 161
                            v4 -= v4
                            var v42 float64 = (1.5 * (1e10 / 0.1))
                            _ = v41
                            _ = v42
                        } else {
                            v43 := GenS4{F0: (^(v10))}
                            var v44 uint32 = 42
                            _ = v43
                            _ = v44
                        }
                        _ = v40
                    }
                }
                _ = v17
            } else {
                v45 := GenS2{F0: 829, F1: false}
                if (1 == -1) {
                    v2 = v2
                } else {
                    fmt.Println(v10, v10)
                    for fi46 := 0; fi46 < 13; fi46++ {
                        v4 += v3
                        _ = fi46
                    }
                    v10 += v10
                }
                _ = v45
            }
        }
    }
    var v47 uint32 = 42
    v48 := GenS5{F0: 0, F1: ("a" + ((("hello" + ("a" + "hello")) + ("" + (("\n\t" + "foo bar") + "a"))) + "foo bar"))}
    v49 := GenS4{F0: 0}
    var v50 int32 = 0
    for fi51 := 0; fi51 < 19; fi51++ {
        v52 := []int{433}
        v1[2147483647] = ((v47 <= genFn1(0)) && (!false))
        for fi53 := 0; fi53 < 11; fi53++ {
            if true {
                var v54 uint16 = 665
                _ = v54
            }
            var v55 int32 = (v50 * v50)
            v47 -= genFn1(18446744073709551615)
            _ = fi53
            _ = v55
        }
        _ = fi51
        _ = v52
    }
    switch 42 {
    case -9223372036854775808:
        var v56 int = -18
        _ = v56
    default:
        v57 := []int{1, 9223372036854775807, 4294967295}
        _ = v57
    }
    for fi58 := 0; fi58 < 8; fi58++ {
        var v59 int32 = (-(v50))
        _ = fi58
        _ = v59
    }
    v3 = (v3 + (v3 / (v3 * v3)))
    v60 := []int{1, -358, 9223372036854775807}
    fmt.Println(v10)
    var v61 uint32 = (v47 >> ((65535) & 63))
    switch "\n\t" {
    case "foo bar":
        if false {
            v10 = v10
        } else {
            v60 = append(v60, 128)
        }
    case "hello":
        v3 = 1e10
    case "a":
        for fi62 := 0; fi62 < 19; fi62++ {
            if false {
                var v63 float64 = (float64(0.1) / 3.14159)
                v47 -= (v61 + 1)
                _ = v63
            } else {
                fmt.Println(v10)
                v10 = (^(v10))
            }
            fmt.Println(v50)
            _ = fi62
        }
    }
    for fi64 := 0; fi64 < 13; fi64++ {
        fmt.Println(v3, v2)
        v50 = 42
        if false {
            if (255 == 2) {
                v65 := GenS5{F0: -422, F1: "foo bar"}
                v66 := GenS3{F0: v2, F1: 18446744073709551615}
                v67 := GenS3{F0: (v2 << ((1) & 63)), F1: 606}
                _ = v65
                _ = v66
                _ = v67
            }
        }
        _ = fi64
    }
    v10 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v60:", len(v60), v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v60
    _ = v61
    _ = v9
}
