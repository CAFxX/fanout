package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int8; F1 int32; F2 int64 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS8 struct { F0 uint64; F1 int; F2 uint }
func genFn0() int8 {
    var v5 int64 = -9223372036854775808
    if (1 <= 255) {
        v6 := GenS1{F0: ""}
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: "foo bar"}
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        v9 := GenS1{F0: "a"}
        _ = v9
    }
    var v10 uint32 = 2147483648
    v10 -= (v10 << ((0) & 63))
    _ = v10
    _ = v5
    return -115
}
func genFn1(p0 uint8, p1 int, p2 uint16) int16 {
    var v11 int64 = 504
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := []int{253, 205, (fi12 % ((p1) | 1)), (fi12 / ((fi12) | 1))}
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = fi12
        _ = v13
    }
    if false {
        v14 := GenS1{F0: "a"}
        v11 = (-(v11))
        _ = v14
    } else {
        var v15 float32 = -2.25
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 0
}

func main() {
    var v1 uint32 = 1
    var v2 int = -301
    var v3 uint16 = 1
    var v4 int32 = 0
    var v16 uint = 65535
    fmt.Println(v3)
    for fi17 := 0; fi17 < 6; fi17++ {
        v16 -= v16
        _ = fi17
    }
    v18 := GenS2{F0: 32767}
    v19 := GenS2{F0: -9}
    var v20 int16 = 42
    v21 := GenS2{F0: v20}
    var v22 float32 = 1.5
    v23 := GenS3{F0: v2, F1: uint32(v16)}
    if true {
        var v24 int8 = -87
        switch "\n\t" {
        case "a":
            v25 := GenS3{F0: 128, F1: v1}
            _ = v25
        case "foo bar":
            for fi26 := 0; fi26 < 19; fi26++ {
                if (2 > 1) {
                    var v27 int = 819
                    _ = v27
                }
                v28 := GenS2{F0: (v20 >> ((42) & 63))}
                if true {
                    for fi29 := 0; fi29 < 10; fi29++ {
                        v22 -= (v22 + 1e10)
                        var v30 int32 = v4
                        var v31 int16 = genFn1(172, (-(fi29)), v3)
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    var v32 float64 = 0.0
                    _ = v32
                }
                _ = fi26
                _ = v28
            }
        case "":
            v33 := GenS1{F0: "a"}
            _ = v33
        default:
            var v34 int8 = int8(v20)
            _ = v34
        }
        if (1.5 > 0.0) {
            var v35 int16 = v20
            v36 := []int{(^(v2)), (v2 << ((498) & 63))}
            var v37 int = v2
            _ = v35
            _ = v36
            _ = v37
        }
        _ = v24
    } else {
        if true {
            v3 = (^(v3))
            var v38 int = (-(v2))
            v39 := GenS4{F0: 0, F1: v4, F2: 9223372036854775807}
            _ = v38
            _ = v39
        } else {
            if false {
                for fi40 := 0; fi40 < 11; fi40++ {
                    if (!false) {
                        fmt.Println(v16, v4)
                    } else {
                        for fi41 := 0; fi41 < 1; fi41++ {
                            var v42 float64 = float64(v22)
                            v43 := GenS3{F0: v2, F1: (v1 + v1)}
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        v44 := []int{}
                        _ = v44
                    }
                    fmt.Println(fi40)
                    _ = fi40
                }
            } else {
                switch "a" {
                case "hello":
                    var v45 int16 = v20
                    _ = v45
                case "foo bar":
                    if (true || true) {
                        if (1 >= v20) {
                            for fi46 := 0; fi46 < 8; fi46++ {
                                v47 := []int{fi46, (v2 * (v2 % ((32768) | 1))), -245}
                                var v48 int = (-(v2))
                                v20 = genFn1(2, v2, 42)
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                            v49 := GenS3{F0: v2, F1: v1}
                            var v50 float64 = 3.14159
                            _ = v49
                            _ = v50
                        } else {
                            switch "" {
                            case "hello":
                                v3 = 2
                            case "\n\t":
                                if false {
                                    v1 -= (v1 - (v1 - (v1 * (v1 % ((v1) | 1)))))
                                    for fi51 := 0; fi51 < 17; fi51++ {
                                        v4 = v4
                                        v52 := GenS2{F0: (v20 + v20)}
                                        _ = fi51
                                        _ = v52
                                    }
                                    if ((!true) && true) {
                                        switch (^(v2)) {
                                        case 42:
                                            if ((genFn1(127, v2, v3) < 199) && (false || false)) {
                                                v53 := []int{}
                                                for fi54 := 0; fi54 < 16; fi54++ {
                                                    switch (("" + "hello") + "a") {
                                                    case "hello":
                                                        v2 += fi54
                                                    case "":
                                                        v55 := GenS5{F0: (v3 ^ v3), F1: (v2 % ((fi54) | 1))}
                                                        _ = v55
                                                    case "foo bar":
                                                        v3 = (v3 % ((v3) | 1))
                                                    }
                                                    v16 += (v16 << (((v16 + (v16 * (v16 % ((4294967295) | 1))))) & 63))
                                                    _ = fi54
                                                }
                                                v22 += float32(-2.25)
                                                _ = v53
                                            } else {
                                                var v56 uint16 = (v3 * v3)
                                                v57 := GenS1{F0: "foo bar"}
                                                for fi58 := 0; fi58 < 4; fi58++ {
                                                    v59 := GenS2{F0: (v20 << ((1) & 63))}
                                                    _ = fi58
                                                    _ = v59
                                                }
                                                _ = v56
                                                _ = v57
                                            }
                                        case 2147483647:
                                            v60 := []int{v2}
                                            _ = v60
                                        default:
                                            var v61 float32 = v22
                                            _ = v61
                                        }
                                    } else {
                                        v62 := []int{(v2 + int(v1)), int(v1)}
                                        _ = v62
                                    }
                                } else {
                                    fmt.Println(v20, v20)
                                    fmt.Println(v22, v1)
                                }
                            case "a":
                                for fi63 := 0; fi63 < 18; fi63++ {
                                    for fi64 := 0; fi64 < 16; fi64++ {
                                        var v65 float64 = (float64(v22) * 3.14159)
                                        var v66 uint16 = v3
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                    _ = fi63
                                }
                            default:
                                v16 -= v16
                            }
                            v67 := GenS6{F0: uint(v3)}
                            v20 -= (v20 >> ((2) & 63))
                            _ = v67
                        }
                    } else {
                        for fi68 := 0; fi68 < 10; fi68++ {
                            if false {
                                v69 := GenS2{F0: v20}
                                _ = v69
                            }
                            _ = fi68
                        }
                    }
                case "":
                    var v70 uint = 18446744073709551615
                    _ = v70
                default:
                    if ((!true) && (42 > 451)) {
                        switch (v2 >> ((v16) & 63)) {
                        case 2147483647:
                            var v71 uint64 = 54
                            _ = v71
                        default:
                            v16 = (v16 & (v16 + (v16 << ((v16) & 63))))
                        }
                    } else {
                        v72 := []int{(v2 + v2), v2, int(v20)}
                        fmt.Println(v16, v4, v16)
                        _ = v72
                    }
                }
                switch "a" {
                case "":
                    v16 = 42
                case "\n\t":
                    for fi73 := 0; fi73 < 10; fi73++ {
                        if false {
                            v74 := []int{-399, v2, v2, v2}
                            for fi75 := 0; fi75 < 7; fi75++ {
                                v3 = 65535
                                switch ("\n\t" + ("" + "")) {
                                case "":
                                    var v76 float32 = (v22 + (v22 + (v22 * 0.1)))
                                    _ = v76
                                }
                                _ = fi75
                            }
                            var v77 uint64 = 42
                            _ = v74
                            _ = v77
                        } else {
                            v78 := []int{fi73}
                            _ = v78
                        }
                        fmt.Println(v4)
                        var v79 float64 = 0.0
                        _ = fi73
                        _ = v79
                    }
                }
                v80 := []int{(v2 ^ v2), (v2 - v2)}
                _ = v80
            }
            v1 += uint32(v3)
        }
        v20 = genFn1(42, v2, v3)
        v22 = 1.5
    }
    v81 := GenS3{F0: 1, F1: v1}
    v82 := GenS4{F0: -128, F1: -2147483648, F2: 0}
    v83 := GenS7{F0: v3, F1: 15, F2: 0}
    fmt.Println(v4, v20, v3)
    v84 := GenS8{F0: 409, F1: 32768, F2: uint(v3)}
    fmt.Println(v2)
    if false {
        var v85 bool = false
        v3 -= v3
        _ = v85
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v81
    _ = v82
    _ = v83
    _ = v84
}
