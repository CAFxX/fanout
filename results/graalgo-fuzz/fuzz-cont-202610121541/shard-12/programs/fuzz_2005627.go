package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int32; F2 int16 }
type GenS5 struct { F0 int32; F1 string }
func genFn0(p0 int64, p1 int16, p2 string) uint64 {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint32 = 65535
        for fi9 := 0; fi9 < 10; fi9++ {
            p1 += 1000
            p1 = (p1 ^ p1)
            fmt.Println(v8, p1)
            v8 -= 1
            switch 4294967295 {
            case -1:
                if (274 >= 32767) {
                    p0 = (p0 / ((42) | 1))
                    v10 := []int{(fi7 << ((171) & 63)), (fi7 << ((365) & 63))}
                    var v11 float64 = 3.14159
                    _ = v10
                    _ = v11
                }
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 440
}
func genFn1(p0 bool, p1 int64) int {
    p1 = -9223372036854775808
    v12 := []int{128, 42, 127}
    if (p0 && false) {
        if p0 {
            p0 = (42 >= 557)
        }
        if false {
            _ = len(v12)
            switch "foo bar" {
            case "\n\t":
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 uint8 = 255
                    switch ("a" + "hello") {
                    case "a":
                        var v15 uint16 = uint16(v14)
                        _ = v15
                    case "":
                        v14 = v14
                    case "hello":
                        for fi16 := 0; fi16 < 12; fi16++ {
                            v17 := GenS1{F0: false}
                            var v18 uint32 = 65535
                            if (((((float32((float64(((1.5 / (float32(1.5) / 3.14159)) / (1e10 - float32(((-2.25 - 0.1) / (0.1 - -2.25)))))) + 3.14159)) - 0.0) / (1.5 + float32(1e10))) - (1e10 - (1.5 / ((1e10 * 3.14159) * 3.14159)))) <= 3.14159) || (1.5 >= 1e10)) {
                                if false {
                                    switch fi16 {
                                    case -433:
                                        if (!((p1 * p1) < (p1 + (^(p1))))) {
                                            var v19 uint16 = 0
                                            _ = len(v12)
                                            var v20 uint32 = (v18 * v18)
                                            _ = v19
                                            _ = v20
                                        } else {
                                            v21 := GenS1{F0: ((!true) || (0.0 >= ((1e10 / -2.25) + (1e10 * (3.14159 / 1.0)))))}
                                            v12 = append(v12, fi16)
                                            v22 := GenS1{F0: p0}
                                            switch "" {
                                            case "":
                                                _ = len(v12)
                                            case "foo bar":
                                                if ((v14 <= (v14 >> ((2) & 63))) || false) {
                                                    fmt.Println(v14)
                                                    v23 := GenS1{F0: (!(100 > -128))}
                                                    fmt.Println(v18, p0)
                                                    _ = v23
                                                } else {
                                                    v18 -= v18
                                                    v24 := GenS1{F0: ((fi16 % ((9223372036854775807) | 1)) != fi13)}
                                                    _ = v24
                                                }
                                            case "\n\t":
                                                v25 := GenS1{F0: false}
                                                _ = v25
                                            default:
                                                v12 = append(v12, (-(fi13)))
                                            }
                                            _ = v21
                                            _ = v22
                                        }
                                    default:
                                        var v26 int64 = p1
                                        _ = v26
                                    }
                                    p1 -= p1
                                    fmt.Println(p0, fi16, p1)
                                } else {
                                    v12 = append(v12, fi13)
                                    v27 := GenS2{F0: 721, F1: v18}
                                    _ = v27
                                }
                            } else {
                                v28 := GenS1{F0: (p0 || (!true))}
                                v29 := GenS2{F0: 42, F1: (v18 / ((v18) | 1))}
                                _ = v28
                                _ = v29
                            }
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                    var v30 uint = 65535
                    _ = fi13
                    _ = v14
                    _ = v30
                }
            }
        } else {
            for fi31 := 0; fi31 < 12; fi31++ {
                p1 += 42
                p1 = p1
                p0 = (562 > (fi31 / ((32768) | 1)))
                _ = fi31
            }
        }
        switch -232 {
        case 32768:
            fmt.Println(p1, p1)
        case 440:
            var v32 uint64 = 4294967295
            _ = v32
        case -2147483648:
            var v33 uint8 = 255
            _ = v33
        }
    } else {
        if (((42 != 65535) || true) && (42 >= 2)) {
            fmt.Println(p0, p1)
            p1 += (p1 + (-(p1)))
        } else {
            v34 := GenS2{F0: 186, F1: 2}
            _ = v34
        }
        p1 -= (p1 / ((p1) | 1))
        if true {
            p1 = (p1 ^ 42)
            var v35 int64 = 9223372036854775807
            _ = v35
        }
    }
    switch "" {
    case "foo bar":
        p1 -= (^(p1))
    case "\n\t":
        fmt.Println(p0, p0)
    case "hello":
        for fi36 := 0; fi36 < 20; fi36++ {
            p1 = (p1 << ((2) & 63))
            var v37 string = ""
            _ = fi36
            _ = v37
        }
    }
    _ = p0
    _ = p1
    _ = v12
    return 127
}

func main() {
    v1 := []uint32{1, 501}
    var v2 uint16 = 2
    var v3 int32 = 42
    var v4 uint16 = 147
    var v5 uint = 65535
    v6 := make(map[int]int)
    delete(v6, -925)
    for fi38 := 0; fi38 < 6; fi38++ {
        v3 = (-(v3))
        if ("hello" < "\n\t") {
            v39 := GenS2{F0: 771, F1: 2147483648}
            _ = v39
        } else {
            var v40 int16 = 379
            _ = v40
        }
        _ = fi38
    }
    if ((!true) && true) {
        for fi41 := 0; fi41 < 1; fi41++ {
            if false {
                if true {
                    for fi42 := 0; fi42 < 7; fi42++ {
                        var v43 uint16 = (v2 - v2)
                        switch "a" {
                        case "a":
                            var v44 string = "a"
                            v45 := GenS1{F0: true}
                            _ = v44
                            _ = v45
                        case "\n\t":
                            v43 += 0
                        case "hello":
                            fmt.Println(v43, v5)
                        default:
                            v3 -= v3
                            fmt.Println(fi42, v2, fi41)
                            if false {
                                var v46 float32 = 1.5
                                _ = v46
                            }
                        }
                        _ = fi42
                        _ = v43
                    }
                    var v47 bool = true
                    fmt.Println(v47, v3)
                    _ = v47
                }
            } else {
                v48 := GenS1{F0: (65535 <= uint32(v5))}
                _ = v48
            }
            var v49 bool = true
            fmt.Println(v2)
            _ = fi41
            _ = v49
        }
    }
    v3 = (v3 / (((v3 + 2147483647)) | 1))
    var v50 int64 = 0
    _ = len(v6)
    if false {
        for fi51 := 0; fi51 < 13; fi51++ {
            v52 := GenS3{F0: ("foo bar" + "")}
            _ = fi51
            _ = v52
        }
        var v53 float64 = 0.0
        _ = v53
    }
    fmt.Println(v50, v50, v4)
    if true {
        var v54 int64 = (v50 + v50)
        v1 = append(v1, 1)
        if false {
            fmt.Println(v5, v4, v3)
            var v55 int16 = -1
            switch "foo bar" {
            case "\n\t":
                for fi56 := 0; fi56 < 4; fi56++ {
                    v54 += (v54 - -1)
                    v57 := GenS4{F0: (v4 * (v4 ^ 2)), F1: v3, F2: -32768}
                    v4 -= 2
                    _ = fi56
                    _ = v57
                }
            case "hello":
                if (false && false) {
                    v5 = 65535
                    var v58 int = -1
                    _ = v58
                }
            case "":
                for fi59 := 0; fi59 < 8; fi59++ {
                    v4 = (v2 * v2)
                    v54 += 42
                    _ = fi59
                }
            default:
                v60 := GenS5{F0: -160, F1: "a"}
                _ = v60
            }
            _ = v55
        } else {
            v54 += v50
            delete(v6, 9223372036854775807)
        }
        _ = v54
    }
    var v61 uint8 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v61
}
