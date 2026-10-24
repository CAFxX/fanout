package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 bool; F2 string }
type GenS7 struct { F0 string; F1 int16; F2 uint64 }

func main() {
    var v1 bool = true
    var v2 float32 = 0.1
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: -1, F1: 42, F2: 539}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := v3[(fi6 ^ 871)]
        switch "hello" {
        case "\n\t":
            if v1 {
                v1 = ("hello" < "hello")
                if true {
                    fmt.Println(v7, v7, fi6)
                    v3[fi6] = (^(v7))
                    delete(v3, fi6)
                } else {
                    v8 := GenS2{F0: 0, F1: 0, F2: 127}
                    _ = v8
                }
            }
        case "foo bar":
            v3[-2147483648] = v7
        default:
            v9 := GenS2{F0: 47, F1: 0, F2: 100}
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    switch "" {
    case "a":
        if false {
            v1 = (42 == -1)
        }
    case "":
        if (!((!false) || (v1 && (0 < 480)))) {
            v1 = (255 >= 255)
        } else {
            if (!true) {
                v10 := GenS3{F0: 42, F1: 825}
                v2 = v2
                _ = v10
            }
            v2 = 0.1
            if false {
                if (!v1) {
                    if (!(9223372036854775807 < 1)) {
                        for fi11 := 0; fi11 < 18; fi11++ {
                            v1 = true
                            if true {
                                if v1 {
                                    if (float64(v2) > 0.1) {
                                        if (1 <= -2147483648) {
                                            v3[(fi11 >> ((42) & 63))] = 65535
                                        }
                                        v2 = (v2 * (0.0 / v2))
                                        v12 := GenS4{F0: 0, F1: -107}
                                        _ = v12
                                    }
                                }
                                var v13 int8 = -128
                                _ = v13
                            }
                            v1 = v1
                            _ = fi11
                        }
                        v14 := GenS2{F0: 13, F1: 18446744073709551615, F2: 100}
                        switch "\n\t" {
                        case "":
                            for fi15 := 0; fi15 < 7; fi15++ {
                                v16 := GenS5{F0: 1000}
                                var v17 bool = (!((-1 >= 42) || true))
                                v2 += v2
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        default:
                            if false {
                                var v18 int64 = 42
                                _ = v18
                            }
                        }
                        _ = v14
                    } else {
                        v19 := GenS4{F0: 65535, F1: -300}
                        v1 = (42 != 2)
                        _ = v19
                    }
                    v1 = true
                } else {
                    switch ("" + "foo bar") {
                    case "":
                        var v20 uint16 = 255
                        _ = v20
                    case "\n\t":
                        if (2 > 4294967295) {
                            v21 := []int{-9223372036854775808, -1}
                            var v22 int64 = 25
                            _ = v21
                            _ = v22
                        } else {
                            v2 = -2.25
                        }
                    case "hello":
                        var v23 uint32 = 0
                        _ = v23
                    }
                    if ((true || ((("a" + "") + "\n\t") <= ("hello" + ""))) && v1) {
                        switch -207 {
                        case 171:
                            if (v1 && v1) {
                                for fi24 := 0; fi24 < 1; fi24++ {
                                    v25 := GenS3{F0: 2147483647, F1: 0}
                                    var v26 uint64 = 42
                                    fmt.Println(fi24, v2, v1)
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                                fmt.Println(v1, v1)
                                for fi27 := 0; fi27 < 7; fi27++ {
                                    v28 := GenS1{F0: 42, F1: 236, F2: 65535}
                                    if (v1 || (-656 < -909)) {
                                        var v29 uint = 703
                                        v30 := GenS3{F0: 4294967295, F1: 32767}
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = fi27
                                    _ = v28
                                }
                            } else {
                                switch ("foo bar" + "a") {
                                case "a":
                                    var v31 int8 = -89
                                    _ = v31
                                case "\n\t":
                                    v32 := []int{459, -9223372036854775808}
                                    _ = v32
                                }
                                var v33 int16 = 42
                                _ = v33
                            }
                        }
                    } else {
                        var v34 int64 = 1
                        _ = v34
                    }
                }
            } else {
                v1 = (v1 || v1)
                var v35 uint16 = 32767
                _ = v35
            }
        }
    case "\n\t":
        v2 += (v2 + (((((v2 / (1e10 / v2)) / float32(-2.25)) * (v2 * 3.14159)) * float32(0.1)) + ((v2 + 0.1) * (v2 * -2.25))))
    default:
        v3[-422] = 867
    }
    v1 = (v1 || (false || false))
    if true {
        v1 = v1
    }
    switch "foo bar" {
    case "":
        v36 := GenS5{F0: -1}
        _ = v36
    }
    if v1 {
        v37 := v3[-9223372036854775808]
        fmt.Println(v1, v37, v2)
        var v38 uint32 = (v37 << ((42) & 63))
        _ = v37
        _ = v38
    }
    v1 = true
    switch -611 {
    case 2147483647:
        if (370 > 1) {
            fmt.Println(v1)
            v39 := GenS2{F0: 42, F1: 65535, F2: -128}
            _ = v39
        } else {
            var v40 uint64 = 1
            v40 = v40
            fmt.Println(v40, v40)
            _ = v40
        }
    case 0:
        if (802 <= 128) {
            if false {
                for fi41 := 0; fi41 < 2; fi41++ {
                    fmt.Println(fi41)
                    v1 = v1
                    v42 := []int{-9223372036854775808}
                    _ = fi41
                    _ = v42
                }
                v1 = false
                switch -243 {
                case -1:
                    v43 := []int{}
                    _ = v43
                case 32767:
                    v44 := GenS6{F0: 42, F1: true, F2: "foo bar"}
                    _ = v44
                case 128:
                    v2 = (0.1 - v2)
                default:
                    fmt.Println(v2)
                }
            }
        } else {
            if ((((v2 / v2) / v2) <= v2) || (!true)) {
                v2 += (v2 * (v2 - v2))
                var v45 int32 = -694
                v45 = v45
                _ = v45
            } else {
                switch ("hello" + "") {
                case "a":
                    v46 := GenS2{F0: 42, F1: 2, F2: 1}
                    _ = v46
                case "hello":
                    fmt.Println(v1, v1)
                default:
                    switch "a" {
                    case "":
                        v47 := []int{32768, 32767, 42, 127}
                        _ = v47
                    default:
                        if v1 {
                            var v48 uint8 = 128
                            var v49 int32 = 42
                            _ = v48
                            _ = v49
                        } else {
                            if (1 > 571) {
                                v2 = v2
                                fmt.Println(v1, v2)
                            } else {
                                v2 = float32(3.14159)
                                if v1 {
                                    v50 := GenS2{F0: -1, F1: 562, F2: 127}
                                    if v1 {
                                        delete(v3, 127)
                                        for fi51 := 0; fi51 < 16; fi51++ {
                                            var v52 float64 = 1e10
                                            v2 += v2
                                            _ = fi51
                                            _ = v52
                                        }
                                        delete(v3, -1)
                                    }
                                    v2 = 0.0
                                    _ = v50
                                } else {
                                    v1 = false
                                }
                            }
                            var v53 string = ("a" + "foo bar")
                            _ = v53
                        }
                    }
                }
            }
        }
    case 61:
        fmt.Println(v2, v1)
    }
    for fi54 := 0; fi54 < 6; fi54++ {
        var v55 float64 = float64(((0.0 / (3.14159 + ((0.0 - v2) * (0.1 / 1e10)))) / (v2 / (v2 - v2))))
        v56 := GenS7{F0: ("" + ("foo bar" + ("a" + ""))), F1: 32767, F2: 0}
        if v1 {
            var v57 int8 = 105
            _ = v57
        } else {
            var v58 uint = 649
            v59 := []int{fi54}
            _ = v58
            _ = v59
        }
        _ = fi54
        _ = v55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}
