package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
type GenS2 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint32, p2 uint64) int8 {
    v4 := []int{32767, int(p1), 1}
    v5 := GenS1{F0: 32767, F1: (p2 / ((2) | 1)), F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -41
}

func main() {
    v1 := []uint32{551, 661, 0, 0}
    var v2 float64 = 1.5
    var v3 int32 = 1
    var v6 uint32 = 0
    if (!true) {
        if false {
            v2 += v2
        } else {
            for fi7 := 0; fi7 < 7; fi7++ {
                if false {
                    for fi8 := 0; fi8 < 8; fi8++ {
                        v9 := GenS1{F0: 0, F1: 941, F2: (fi8 + fi7)}
                        var v10 uint8 = 18
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    switch -9223372036854775808 {
                    case 127:
                        var v11 float32 = 3.14159
                        _ = v11
                    case 571:
                        switch "hello" {
                        case "\n\t":
                            var v12 uint16 = 851
                            _ = v12
                        case "a":
                            switch fi7 {
                            case 127:
                                if (((((v2 / -2.25) * (((v2 + (v2 * -2.25)) - 0.0) / float64(-2.25))) * v2) + (((((1e10 * ((1e10 - v2) * ((v2 / 1e10) / v2))) - 1.5) + (v2 / v2)) * (v2 - float64(1e10))) * (v2 - v2))) <= 3.14159) {
                                    var v13 uint16 = 855
                                    v14 := GenS1{F0: 32767, F1: 0, F2: fi7}
                                    var v15 uint16 = v13
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                }
                            default:
                                for fi16 := 0; fi16 < 10; fi16++ {
                                    var v17 uint16 = 784
                                    _ = fi16
                                    _ = v17
                                }
                            }
                        case "hello":
                            v18 := GenS1{F0: 121, F1: 884, F2: (fi7 | fi7)}
                            _ = v18
                        }
                    case -416:
                        v6 += (^(v6))
                    }
                    var v19 bool = (-1 >= 0)
                    _ = v19
                }
                _ = fi7
            }
        }
    } else {
        v20 := GenS2{F0: (v3 / ((-2147483648) | 1)), F1: 4294967295, F2: 42}
        for fi21 := 0; fi21 < 15; fi21++ {
            if (106 >= -933) {
                var v22 uint16 = 13
                _ = v22
            } else {
                var v23 float32 = (1.5 - ((((-2.25 + (((0.0 * (((1.5 / 3.14159) * 0.0) * (0.1 - -2.25))) - 3.14159) * 1.5)) - 1e10) + -2.25) + 1.5))
                for fi24 := 0; fi24 < 5; fi24++ {
                    v3 += v3
                    _ = fi24
                }
                _ = v23
            }
            switch (fi21 + 142) {
            case 179:
                var v25 float64 = 1e10
                _ = v25
            default:
                v2 = 0.0
            }
            v2 -= (v2 - (0.1 - (((3.14159 / v2) * (v2 - 1e10)) + 0.0)))
            _ = fi21
        }
        _ = v20
    }
    switch 2147483647 {
    case -709:
        v2 = float64((float32((v2 / 3.14159)) - 0.1))
    default:
        if (42 <= -32768) {
            if (0 >= -44) {
                var v26 string = "\n\t"
                for fi27 := 0; fi27 < 17; fi27++ {
                    v3 -= (v3 | 0)
                    var v28 bool = (0 != 816)
                    if true {
                        for fi29 := 0; fi29 < 6; fi29++ {
                            var v30 int = 598
                            v2 += v2
                            _ = fi29
                            _ = v30
                        }
                        var v31 int64 = 0
                        _ = v31
                    } else {
                        for fi32 := 0; fi32 < 18; fi32++ {
                            var v33 bool = v28
                            var v34 int = (fi32 * -446)
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    }
                    _ = fi27
                    _ = v28
                }
                var v35 uint32 = v6
                _ = v26
                _ = v35
            } else {
                v1 = append(v1, v6)
            }
            for fi36 := 0; fi36 < 15; fi36++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v6 = (v6 >> ((18446744073709551615) & 63))
                _ = fi36
            }
            var v37 string = "foo bar"
            _ = v37
        } else {
            if false {
                var v38 bool = false
                _ = v38
            } else {
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v39 uint = 827
                    _ = v39
                } else {
                    fmt.Println(v6)
                }
            }
            fmt.Println(v3, v3)
        }
    }
    for fi40 := 0; fi40 < 4; fi40++ {
        for fi41 := 0; fi41 < 3; fi41++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v2, fi40)
            if false {
                fmt.Println(v2)
            } else {
                if true {
                    _ = len(v1)
                } else {
                    v6 = 2147483648
                    if true {
                        var v42 uint64 = 18446744073709551615
                        var v43 uint32 = v6
                        v43 = v6
                        _ = v42
                        _ = v43
                    } else {
                        fmt.Println(v2)
                        v1 = append(v1, v6)
                        v2 -= 0.0
                    }
                    if false {
                        if false {
                            v2 = v2
                            v44 := GenS3{F0: -128}
                            _ = v44
                        }
                    }
                }
                if true {
                    var v45 int32 = 42
                    v45 = -50
                    _ = v45
                } else {
                    v6 = v6
                }
            }
            _ = fi41
        }
        v46 := GenS2{F0: (v3 / (((v3 - -2147483648)) | 1)), F1: 42, F2: 0}
        _ = fi40
        _ = v46
    }
    if true {
        var v47 string = "\n\t"
        if (false || (2147483647 < (v6 - (v6 / ((v6) | 1))))) {
            for fi48 := 0; fi48 < 7; fi48++ {
                var v49 float32 = (0.1 / ((-2.25 / 0.1) / 1.0))
                var v50 float64 = (1.5 - (3.14159 - v2))
                _ = fi48
                _ = v49
                _ = v50
            }
            v51 := GenS4{F0: (v47 + "hello")}
            _ = v51
        } else {
            var v52 float64 = -2.25
            _ = v52
        }
        v3 -= v3
        _ = v47
    } else {
        var v53 int32 = 0
        v3 += v53
        fmt.Println(v6)
        _ = v53
    }
    switch 9223372036854775807 {
    case -600:
        v54 := GenS5{F0: (v6 << ((2) & 63)), F1: v6, F2: 9223372036854775807}
        _ = v54
    default:
        v2 = 0.0
    }
    var v55 bool = true
    v56 := GenS6{F0: 255}
    var v57 bool = (v55 || v55)
    var v58 uint16 = 42
    fmt.Println(v3, v58, v58)
    v55 = v55
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
