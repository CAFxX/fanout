package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 uint; F2 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 int32; F2 string }
type GenS7 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 int32; F1 int; F2 int32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int8; F1 uint32; F2 string }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 bool; F1 int16; F2 string }
type GenS14 struct { F0 int32; F1 int16; F2 int8 }
func genFn0() int16 {
    v4 := GenS1{F0: -44}
    v5 := GenS1{F0: -111}
    if (false || (4294967295 > 2)) {
        v6 := GenS2{F0: 1, F1: 18446744073709551615, F2: 196}
        _ = v6
    }
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 string) int16 {
    v7 := []int{-2147483648, 247, -33, 25}
    var v8 int64 = -637
    if false {
        var v9 float32 = -2.25
        fmt.Println(v9, v9)
        fmt.Println(v8)
        _ = v9
    }
    _ = p0
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 bool = (1e10 > float32((1e10 - ((1.5 * ((0.1 / ((1e10 * 1.5) + 0.0)) + -2.25)) * (1.5 + 1e10)))))
    var v2 string = ("\n\t" + "\n\t")
    v3 := []string{(v2 + v2)}
    var v10 uint64 = 1
    for fi11 := 0; fi11 < 15; fi11++ {
        v3 = append(v3, "")
        for fi12 := 0; fi12 < 2; fi12++ {
            v10 -= (v10 << ((18446744073709551615) & 63))
            v1 = (!v1)
            var v13 bool = false
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v3 = append(v3, v2)
        var v15 string = v2
        _ = fi14
        _ = v15
    }
    v3 = append(v3, (v2 + v2))
    switch ((((v2 + v2) + (v2 + v2)) + "hello") + (((v2 + "\n\t") + "\n\t") + v2)) {
    case "":
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS3{F0: v1}
            if (((0 >= 32767) || (!true)) || v1) {
                v18 := GenS1{F0: -1}
                var v19 float32 = 0.1
                v2 = v2
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 14; fi20++ {
                    switch v2 {
                    case "a":
                        fmt.Println(fi20, fi16, v10)
                    }
                    var v21 int64 = -381
                    switch 380 {
                    case -117:
                        if ((!v1) || (25 <= 255)) {
                            v2 = v2
                        } else {
                            v2 = (v2 + (v2 + (v2 + "a")))
                            v22 := GenS4{F0: 255, F1: uint(v10), F2: fi16}
                            switch 2147483647 {
                            case -419:
                                fmt.Println(fi16, v21, v2)
                            case 948:
                                switch fi20 {
                                case 1:
                                    fmt.Println(v10, fi16)
                                default:
                                    fmt.Println(v10, v10, v10)
                                }
                            }
                            _ = v22
                        }
                    case 127:
                        if v1 {
                            if false {
                                v23 := GenS5{F0: v21}
                                v2 = "a"
                                _ = v23
                            }
                            if (!(-2147483648 < 1)) {
                                var v24 int64 = v21
                                v24 = v21
                                _ = v24
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = (v2 + "\n\t")
                                }
                            }
                        }
                    default:
                        fmt.Println(v1, v1, v1)
                    }
                    _ = fi20
                    _ = v21
                }
                v1 = (!false)
                switch (fi16 - (fi16 % ((fi16) | 1))) {
                case -600:
                    fmt.Println(v2, v2, fi16)
                case 1:
                    if v1 {
                        v25 := GenS3{F0: false}
                        _ = v25
                    } else {
                        if (9223372036854775807 == -9223372036854775808) {
                            for fi26 := 0; fi26 < 11; fi26++ {
                                v27 := GenS1{F0: 127}
                                v28 := GenS4{F0: 64, F1: 42, F2: (fi26 / ((2147483647) | 1))}
                                if v1 {
                                    var v29 float32 = (((1.5 * 1.5) + 0.0) * 0.1)
                                    fmt.Println(v29, v1, v29)
                                    _ = v29
                                } else {
                                    switch ("" + v2) {
                                    case "foo bar":
                                        if ((0.0 / -2.25) > ((3.14159 - 3.14159) - -2.25)) {
                                            var v30 uint8 = 91
                                            _ = v30
                                        } else {
                                            v31 := GenS6{F0: 2, F1: 2147483647, F2: v2}
                                            _ = len(v3)
                                            _ = v31
                                        }
                                        v32 := GenS5{F0: -9223372036854775808}
                                        _ = v32
                                    case "\n\t":
                                        v33 := GenS2{F0: fi16, F1: 18446744073709551615, F2: 0}
                                        _ = v33
                                    case "a":
                                        v34 := GenS7{F0: 0, F1: 972, F2: 14}
                                        _ = v34
                                    default:
                                        var v35 bool = (v1 && (v2 != ("" + (v2 + v2))))
                                        _ = v35
                                    }
                                    v2 = v2
                                    switch "\n\t" {
                                    case "hello":
                                        v36 := GenS8{F0: (v10 - (^(v10))), F1: 2147483647}
                                        _ = v36
                                    }
                                }
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            v37 := GenS9{F0: 2147483647, F1: 1, F2: -1}
                            v38 := GenS3{F0: true}
                            _ = v37
                            _ = v38
                        } else {
                            fmt.Println(fi16, fi16)
                            v3 = append(v3, (v2 + ("foo bar" + v2)))
                        }
                        var v39 uint8 = 200
                        v40 := GenS5{F0: -9223372036854775808}
                        _ = v39
                        _ = v40
                    }
                }
            }
            _ = fi16
            _ = v17
        }
    case "hello":
        if false {
            v10 = v10
            if false {
                var v41 float64 = 0.0
                v42 := GenS10{F0: (v41 >= 1e10)}
                v3 = append(v3, v2)
                _ = v41
                _ = v42
            }
            v1 = false
        } else {
            v43 := GenS11{F0: 127, F1: 2, F2: ""}
            _ = v43
        }
    case "\n\t":
        for fi44 := 0; fi44 < 7; fi44++ {
            if (v1 && (42 != -9223372036854775808)) {
                var v45 uint16 = 2
                _ = v45
            } else {
                v46 := GenS7{F0: 76, F1: v10, F2: 128}
                _ = v46
            }
            var v47 float64 = 1.5
            _ = fi44
            _ = v47
        }
    default:
        v1 = true
    }
    v2 = v2
    if true {
        for fi48 := 0; fi48 < 12; fi48++ {
            var v49 float64 = 1.5
            if false {
                var v50 bool = v1
                var v51 string = v2
                var v52 float32 = float32(((0.0 / v49) * float64(1e10)))
                _ = v50
                _ = v51
                _ = v52
            } else {
                if v1 {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v1 = true
                    var v53 int = 42
                    _ = v53
                }
            }
            v54 := GenS12{F0: 42}
            _ = fi48
            _ = v49
            _ = v54
        }
    }
    v3 = append(v3, v2)
    v1 = (v2 > "hello")
    var v55 int8 = 100
    switch v2 {
    case "a":
        var v56 uint32 = 2
        _ = v56
    case "\n\t":
        if (true && (!(!v1))) {
            fmt.Println(v1, v1)
            v57 := GenS13{F0: true, F1: 1, F2: v2}
            for fi58 := 0; fi58 < 9; fi58++ {
                if (v1 || v1) {
                    for fi59 := 0; fi59 < 19; fi59++ {
                        var v60 uint8 = 2
                        v1 = ((761 >= 838) || (!v1))
                        v61 := GenS10{F0: (42 <= -128)}
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                }
                v62 := GenS14{F0: -2147483648, F1: int16(fi58), F2: v55}
                v55 += 0
                _ = fi58
                _ = v62
            }
            _ = v57
        } else {
            v63 := GenS14{F0: 880, F1: 42, F2: v55}
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            fmt.Println(v2, v10)
            _ = v63
        }
    case "foo bar":
        if (!v1) {
            v64 := GenS2{F0: -2147483648, F1: 374, F2: -9223372036854775808}
            _ = len(v3)
            _ = v64
        } else {
            for fi65 := 0; fi65 < 19; fi65++ {
                v55 = 100
                _ = fi65
            }
        }
    default:
        fmt.Println(v55, v1)
    }
    for fi66 := 0; fi66 < 5; fi66++ {
        fmt.Println(v55, v55)
        fmt.Println(v55, v55, v1)
        fmt.Println(fi66)
        _ = fi66
    }
    var v67 bool = false
    var v68 int32 = 42
    var v69 int8 = (v55 | v55)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v55:", v55)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v55
    _ = v67
    _ = v68
    _ = v69
}
