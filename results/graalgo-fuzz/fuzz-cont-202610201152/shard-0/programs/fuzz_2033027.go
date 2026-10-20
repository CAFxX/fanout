package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 bool; F1 string; F2 int8 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 int32 }
func genFn0(p0 uint, p1 uint64, p2 int64) uint16 {
    var v4 uint16 = 32767
    v5 := GenS1{F0: 0}
    var v6 uint64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint32) uint8 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 3; fi7++ {
        if false {
            var v8 int16 = 1
            for fi9 := 0; fi9 < 15; fi9++ {
                switch (fi7 ^ (fi9 + (fi9 >> ((42) & 63)))) {
                case 463:
                    var v10 int16 = (v8 * (v8 / (((v8 % (((v8 - (v8 / ((62) | 1)))) | 1))) | 1)))
                    _ = v10
                case 1:
                    v8 -= (v8 % ((v8) | 1))
                    v11 := GenS1{F0: 575}
                    _ = v11
                case 912:
                    var v12 bool = false
                    _ = v12
                default:
                    if (uint64(p0) == 836) {
                        var v13 int64 = 1
                        switch "\n\t" {
                        case "foo bar":
                            switch "hello" {
                            case "hello":
                                switch int(v13) {
                                case -2147483648:
                                    v14 := GenS2{F0: 54}
                                    v13 += v13
                                    v13 = 42
                                    p0 = (^(p0))
                                    v8 += (v8 / (((-(v8))) | 1))
                                    _ = v14
                                case -543:
                                    fmt.Println(p0, fi9)
                                }
                            case "\n\t":
                                var v15 int16 = 42
                                _ = v15
                            case "":
                                var v16 int = (fi9 - (fi7 & (fi7 * (fi7 ^ (fi9 * fi9)))))
                                _ = v16
                            default:
                                if (v8 <= (^(v8))) {
                                    var v17 int8 = 47
                                    if (42 <= 1000) {
                                        switch 32768 {
                                        case 124:
                                            v8 += -959
                                        case 127:
                                            v17 -= v17
                                        default:
                                            if ((1.5 - 0.0) == (3.14159 / -2.25)) {
                                                var v18 uint = 64
                                                v19 := GenS1{F0: 2147483648}
                                                p0 += (p0 << (((v18 ^ (v18 & (v18 << ((1) & 63))))) & 63))
                                                _ = v18
                                                _ = v19
                                            }
                                        }
                                        fmt.Println(v8)
                                        p0 -= (p0 << ((1) & 63))
                                        switch ("hello" + "") {
                                        case "a":
                                            p0 += (^(p0))
                                        case "":
                                            var v20 string = "a"
                                            _ = v20
                                        case "foo bar":
                                            v21 := GenS1{F0: p0}
                                            _ = v21
                                        }
                                        v13 = 42
                                    } else {
                                        v22 := []int{fi7}
                                        _ = v22
                                    }
                                    fmt.Println(v8, fi9, p0)
                                    _ = v17
                                }
                            }
                            v23 := GenS3{F0: 1, F1: (fi9 << ((1) & 63))}
                            var v24 string = ("" + "\n\t")
                            _ = v23
                            _ = v24
                        case "hello":
                            v25 := []int{(fi7 ^ (fi9 << ((42) & 63))), fi9, (fi7 * (fi7 >> ((451) & 63)))}
                            _ = v25
                        }
                        v13 += v13
                        var v26 uint8 = 0
                        var v27 int = fi7
                        v28 := []int{(fi9 + (^(v27))), -227, -9223372036854775808}
                        var v29 int8 = 127
                        switch "hello" {
                        case "hello":
                            var v30 uint16 = 1
                            v13 = int64(v29)
                            _ = v30
                        case "":
                            _ = len(v28)
                            fmt.Println(v29)
                        case "\n\t":
                            switch -6 {
                            case -9223372036854775808:
                                if len(v28) > 0 {
                                    v28[0] = -674
                                }
                            }
                        default:
                            switch (fi7 & -9223372036854775808) {
                            case 32767:
                                v28 = append(v28, (fi7 >> ((18446744073709551615) & 63)))
                            case -411:
                                var v31 int32 = -1
                                _ = v31
                            case 2147483647:
                                var v32 int8 = (v29 - 65)
                                _ = v32
                            }
                        }
                        v29 = v29
                        v29 = (v29 + (v29 - v29))
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                        if true {
                            v8 += (-(v8))
                            v26 += 42
                            v8 = (v8 / (((v8 | v8)) | 1))
                        }
                        var v33 uint16 = 893
                        fmt.Println(v13, v13, p0)
                        var v34 int16 = (v8 - v8)
                        v35 := GenS4{F0: p0, F1: 1}
                        v36 := GenS5{F0: (!true), F1: "\n\t", F2: 1}
                        _ = v13
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v36
                    } else {
                        v37 := GenS6{F0: 65535, F1: 806}
                        _ = v37
                    }
                }
                v38 := []int{(fi9 << ((74) & 63)), (-(fi9))}
                v8 = v8
                _ = fi9
                _ = v38
            }
            _ = v8
        }
        _ = fi7
    }
    v39 := []int{32767, -355, 90}
    _ = p0
    _ = v39
    return 0
}

func main() {
    var v1 float64 = 0.1
    var v2 uint16 = 255
    var v3 bool = ("foo bar" >= "")
    v40 := []int{128, 659}
    fmt.Println(v3)
    var v41 string = "a"
    fmt.Println(v2, v3)
    if len(v40) > 0 {
        v40[0] = 9223372036854775807
    }
    v2 = v2
    if false {
        v2 = 692
        switch 0 {
        case -1:
            v3 = v3
        default:
            v41 = "\n\t"
        }
    } else {
        var v42 uint32 = 42
        _ = v42
    }
    if (("\n\t" + v41) <= v41) {
        fmt.Println(v2)
    } else {
        var v43 int32 = 2147483647
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
        _ = v43
    }
    v3 = (!(v2 == (v2 << ((2) & 63))))
    if (1 != 65535) {
        v3 = v3
        v44 := GenS5{F0: ((v2 & v2) >= 65535), F1: v41, F2: -1}
        v45 := GenS7{F0: 0}
        _ = v44
        _ = v45
    }
    v46 := GenS5{F0: true, F1: (v41 + (v41 + v41)), F2: -49}
    v2 = (v2 * v2)
    fmt.Println(v41, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v46
}
