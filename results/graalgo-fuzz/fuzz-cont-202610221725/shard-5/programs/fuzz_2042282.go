package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint8; F2 int8 }
type GenS5 struct { F0 int32; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint64 }
func genFn0() uint {
    var v6 int32 = 0
    v7 := GenS1{F0: 0}
    if true {
        v8 := GenS2{F0: 426}
        v9 := GenS2{F0: 1}
        var v10 float32 = (((1e10 - -2.25) / (float32((1.5 + ((0.0 * (0.0 + -2.25)) - 1e10))) + 1e10)) * (1e10 * (-2.25 - (1e10 * 3.14159))))
        _ = v10
        _ = v8
        _ = v9
    } else {
        var v11 float32 = (1e10 + -2.25)
        var v12 float32 = 3.14159
        _ = v11
        _ = v12
    }
    if ((75 < 1) || (89 >= 27)) {
        var v13 int16 = -32768
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 uint32 = 4294967295
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        v16 := []int{-882, 42}
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        fmt.Println(v6)
        _ = v16
    }
    _ = v6
    _ = v7
    return 109
}

func main() {
    var v1 float32 = 0.1
    var v2 float32 = v1
    v3 := make(map[int]bool)
    var v4 float64 = -2.25
    var v5 uint32 = 506
    fmt.Println(v5)
    fmt.Println(v5, v2, v5)
    v1 = ((v2 + (v1 * v1)) + 3.14159)
    v2 += (((float32(((v4 / v4) * v4)) / v1) / ((float32((0.0 - v4)) * ((((v1 / v2) / float32((-2.25 + v4))) - (0.0 + (v1 / -2.25))) - float32(v4))) + (v2 / v2))) * (1e10 - v1))
    var v17 uint16 = 42
    var v18 bool = (true && false)
    if ((v1 >= v1) && (v4 >= (3.14159 / float64((v2 * (v2 - (v2 + 1.5))))))) {
        v19 := []int{42, 901, 42, -224}
        var v20 string = "hello"
        _ = v19
        _ = v20
    } else {
        var v21 float64 = (v4 / v4)
        _ = v21
    }
    if (v17 > (v17 ^ (v17 & v17))) {
        v4 -= ((-2.25 * 1.5) - float64((v1 / v1)))
        for fi22 := 0; fi22 < 2; fi22++ {
            if (1 <= 18) {
                var v23 int8 = 42
                for fi24 := 0; fi24 < 12; fi24++ {
                    switch (fi22 * 128) {
                    case 0:
                        if ((v17 < (v17 | v17)) || (65535 >= v17)) {
                            if v18 {
                                if false {
                                    var v25 int16 = -613
                                    fmt.Println(v23, v23, v25)
                                    _ = v25
                                } else {
                                    delete(v3, fi24)
                                    var v26 uint16 = v17
                                    v2 -= (v2 / ((v2 + v1) / (1e10 + (v2 - ((v1 - 1.5) + ((v1 * v2) - (0.0 / v1)))))))
                                    _ = v26
                                }
                                switch ("" + "\n\t") {
                                case "\n\t":
                                    v17 = (v17 ^ (v17 / ((v17) | 1)))
                                }
                                fmt.Println(v17, v2, v5)
                            } else {
                                var v27 uint64 = 2
                                delete(v3, int(v27))
                                switch "hello" {
                                case "a":
                                    if (!true) {
                                        v23 = v23
                                    }
                                case "foo bar":
                                    v28 := []int{-103, int(v23), fi24}
                                    _ = v28
                                case "":
                                    v29 := []int{(fi24 ^ 445)}
                                    switch fi24 {
                                    case 1:
                                        delete(v3, (fi24 / ((fi22) | 1)))
                                    default:
                                        v29 = append(v29, fi24)
                                    }
                                    _ = v29
                                default:
                                    if (v5 != (v5 + (v5 + uint32(v17)))) {
                                        fmt.Println(v2, v5)
                                        v30 := GenS3{F0: -2147483648, F1: v5, F2: (v5 << ((4294967295) & 63))}
                                        v31 := GenS2{F0: 0}
                                        fmt.Println(v18, v27, v27)
                                        _ = v30
                                        _ = v31
                                    }
                                }
                                _ = v27
                            }
                            fmt.Println(v18, v23, fi24)
                            v32 := []int{int(v5), int(v17), fi24}
                            _ = v32
                        }
                    default:
                        v33 := GenS1{F0: 9223372036854775807}
                        _ = v33
                    }
                    v2 = (v1 / (3.14159 * v1))
                    v34 := GenS1{F0: -9223372036854775808}
                    if v18 {
                        fmt.Println(v2, v2)
                    } else {
                        v35 := []int{int(v5)}
                        v2 = ((v2 / v1) - v2)
                        switch "" {
                        case "hello":
                            v2 = (v2 + v2)
                        case "":
                            var v36 uint32 = (v5 * 771)
                            _ = v36
                        }
                        var v37 int = fi22
                        v23 = (^(v23))
                        _ = v35
                        _ = v37
                    }
                    v3[fi24] = (!v18)
                    _ = fi24
                    _ = v34
                }
                fmt.Println(v23)
                _ = v23
            } else {
                for fi38 := 0; fi38 < 12; fi38++ {
                    if (v18 || (!v18)) {
                        switch (fi22 ^ (^(fi38))) {
                        case 4294967295:
                            v39 := GenS2{F0: 0}
                            _ = v39
                        default:
                            v2 = 0.0
                        }
                        v4 = (((v4 * (v4 / (v4 * v4))) - 3.14159) + -2.25)
                    }
                    delete(v3, -70)
                    v40 := GenS4{F0: v18, F1: uint8(v5), F2: 100}
                    _ = fi38
                    _ = v40
                }
                fmt.Println(v5, v4, v5)
                v41 := v3[fi22]
                _ = v41
            }
            _ = fi22
        }
        switch ("\n\t" + "") {
        case "foo bar":
            switch 244 {
            case -1:
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v1, v17)
                default:
                    var v42 uint64 = 26
                    _ = v42
                }
            default:
                v1 = v2
            }
        default:
            v18 = (-636 <= -1)
        }
    }
    v43 := GenS5{F0: 2147483647, F1: -1, F2: 2}
    v3[531] = v18
    fmt.Println(v18)
    switch ("a" + "\n\t") {
    case "":
        var v44 float64 = (1.5 - float64(float32(v4)))
        _ = v44
    case "hello":
        var v45 int64 = 1
        _ = v45
    case "foo bar":
        for fi46 := 0; fi46 < 20; fi46++ {
            v47 := GenS5{F0: 1, F1: -800, F2: v17}
            fmt.Println(v17, v1, v18)
            _ = fi46
            _ = v47
        }
    default:
        var v48 int8 = -128
        _ = v48
    }
    for fi49 := 0; fi49 < 10; fi49++ {
        v50 := GenS6{F0: v17, F1: 208, F2: 840}
        var v51 uint8 = 255
        for fi52 := 0; fi52 < 5; fi52++ {
            if v18 {
                v4 = v4
                var v53 uint16 = (^(v17))
                v4 += float64(1e10)
                _ = v53
            }
            if (4294967295 <= 18446744073709551615) {
                v2 += (float32(v4) * 3.14159)
                var v54 uint16 = (v17 >> ((uint(v5)) & 63))
                _ = v54
            } else {
                if true {
                    v55 := GenS5{F0: 2147483647, F1: 1, F2: (v17 - 32767)}
                    if true {
                        if ((fi52 < (fi52 ^ (fi52 % ((fi52) | 1)))) || v18) {
                            if false {
                                if (!v18) {
                                    var v56 bool = false
                                    v51 = 2
                                    _ = v56
                                }
                                var v57 bool = true
                                v4 = v4
                                var v58 uint8 = 127
                                v59 := []int{32767}
                                v4 = -2.25
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                            var v60 uint32 = v5
                            _ = v60
                        } else {
                            v51 += 42
                            v61 := v3[-1]
                            _ = v61
                        }
                    }
                    v62 := []int{(fi52 << ((667) & 63)), fi49, -1}
                    _ = v55
                    _ = v62
                } else {
                    v2 += float32(v4)
                    fmt.Println(fi52, v4, fi49)
                    switch fi52 {
                    case 1:
                        var v63 uint16 = v17
                        _ = v63
                    default:
                        fmt.Println(v17, v5, v1)
                    }
                }
                var v64 uint16 = 711
                fmt.Println(v1, v1, v1)
                _ = v64
            }
            _ = fi52
        }
        _ = fi49
        _ = v50
        _ = v51
    }
    v1 += 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
