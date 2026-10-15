package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 bool }
func genFn0(p0 int32, p1 string, p2 int64) uint {
    var v5 uint = 42
    switch 711 {
    case 9223372036854775807:
        var v6 float64 = 1.5
        _ = v6
    case 2147483647:
        p2 = 998
    case 1:
        fmt.Println(p1, p2)
    }
    v7 := GenS1{F0: 32768, F1: ((255 > 0) && (1000 <= -1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return v5
}
func genFn1(p0 int8, p1 int, p2 int) string {
    if ((p0 >> ((1) & 63)) == 0) {
        v8 := GenS1{F0: 40, F1: true}
        _ = v8
    }
    fmt.Println(p0)
    switch (^(p2)) {
    case 0:
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: p1, F1: (true || ((p0 > 118) && (828 > 652)))}
            for fi11 := 0; fi11 < 6; fi11++ {
                v12 := GenS1{F0: fi9, F1: (true || (1 != 1))}
                p1 = 9223372036854775807
                _ = fi11
                _ = v12
            }
            v13 := []int{(-(p1))}
            _ = fi9
            _ = v10
            _ = v13
        }
    case -579:
        for fi14 := 0; fi14 < 14; fi14++ {
            switch p1 {
            case -9223372036854775808:
                if (4294967295 <= 2147483647) {
                    v15 := GenS1{F0: p1, F1: ((-1 > 1) && true)}
                    var v16 uint16 = 32767
                    for fi17 := 0; fi17 < 14; fi17++ {
                        var v18 int8 = p0
                        var v19 uint8 = 0
                        v19 += (^(v19))
                        v16 += (^(v16))
                        v18 -= 127
                        v20 := GenS1{F0: (fi17 >> ((2) & 63)), F1: false}
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    _ = v15
                    _ = v16
                } else {
                    var v21 uint64 = 114
                    p0 = (p0 - 31)
                    if false {
                        v22 := GenS2{F0: 2}
                        v21 -= v21
                        for fi23 := 0; fi23 < 12; fi23++ {
                            switch p2 {
                            case -401:
                                v24 := GenS2{F0: 255}
                                _ = v24
                            default:
                                var v25 bool = true
                                _ = v25
                            }
                            v21 += 386
                            fmt.Println(fi14, fi14, fi14)
                            v26 := GenS3{F0: 495, F1: (v21 & v21)}
                            _ = fi23
                            _ = v26
                        }
                        _ = v22
                    }
                    _ = v21
                }
            default:
                v27 := []int{p2, -428}
                _ = v27
            }
            v28 := []int{}
            var v29 uint32 = 96
            _ = fi14
            _ = v28
            _ = v29
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    v1 := []string{}
    var v2 uint16 = 0
    var v3 uint8 = 255
    var v4 int64 = 42
    if ((!false) || (0.1 == 1.5)) {
        v30 := GenS2{F0: (v3 | v3)}
        _ = v30
    } else {
        v31 := GenS3{F0: 1, F1: uint64(v3)}
        fmt.Println(v2)
        for fi32 := 0; fi32 < 11; fi32++ {
            for fi33 := 0; fi33 < 11; fi33++ {
                v3 = v3
                _ = fi33
            }
            for fi34 := 0; fi34 < 20; fi34++ {
                v4 -= (v4 & v4)
                v35 := GenS3{F0: 1000, F1: 745}
                var v36 float64 = 0.0
                _ = fi34
                _ = v35
                _ = v36
            }
            switch "hello" {
            case "hello":
                var v37 uint64 = 2
                _ = v37
            case "a":
                v3 += v3
            default:
                switch 32768 {
                case 32767:
                    fmt.Println(v3, v4, v2)
                case -68:
                    if true {
                        v1 = append(v1, ((("" + ("hello" + "a")) + "\n\t") + "a"))
                        if false {
                            if false {
                                fmt.Println(fi32)
                                _ = len(v1)
                            } else {
                                v1 = append(v1, "\n\t")
                                v38 := GenS3{F0: 1000, F1: 661}
                                _ = v38
                            }
                        }
                        var v39 int = fi32
                        _ = v39
                    } else {
                        v40 := GenS1{F0: 2147483647, F1: true}
                        _ = v40
                    }
                case 881:
                    for fi41 := 0; fi41 < 4; fi41++ {
                        v42 := GenS2{F0: v3}
                        if (!((v3 & v3) <= (v3 & v3))) {
                            switch ("" + "foo bar") {
                            case "hello":
                                v43 := GenS4{F0: (v4 / (((v4 - 0)) | 1)), F1: fi41}
                                _ = v43
                            default:
                                if (0 < uint(v3)) {
                                    if true {
                                        v3 = (v3 / (((v3 & (v3 | v3))) | 1))
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(v3, v4)
                                }
                            }
                            v2 -= (v2 + uint16(v3))
                            _ = len(v1)
                        } else {
                            var v44 uint8 = (v3 % (((v3 << ((255) & 63))) | 1))
                            _ = v44
                        }
                        v45 := GenS5{F0: true}
                        v4 += v4
                        var v46 bool = (473 != 4294967295)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v47 := GenS5{F0: (7 <= v3)}
                        if (1000 < -581) {
                            fmt.Println(v2)
                            var v48 int8 = -128
                            v49 := GenS2{F0: v3}
                            _ = v48
                            _ = v49
                        } else {
                            var v50 float32 = 1.5
                            _ = v50
                        }
                        _ = fi41
                        _ = v42
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                default:
                    v1 = append(v1, "foo bar")
                }
            }
            _ = fi32
        }
        _ = v31
    }
    v51 := GenS4{F0: v4, F1: 2147483647}
    var v52 uint32 = 394
    fmt.Println(v4, v3, v52)
    if true {
        v52 = (v52 & (v52 ^ (v52 ^ 65535)))
    }
    fmt.Println(v2, v52)
    v53 := GenS3{F0: 1, F1: 42}
    var v54 float32 = (1.5 / ((((3.14159 * -2.25) - 3.14159) - 0.0) / (1e10 + (1e10 * -2.25))))
    var v55 int16 = -359
    for fi56 := 0; fi56 < 19; fi56++ {
        switch ("foo bar" + "a") {
        case "\n\t":
            v3 -= (v3 / ((v3) | 1))
        case "hello":
            switch fi56 {
            case -2147483648:
                v54 += ((v54 + (v54 * float32((0.1 / 1.5)))) - v54)
            default:
                _ = len(v1)
            }
        }
        for fi57 := 0; fi57 < 20; fi57++ {
            var v58 uint16 = 65535
            fmt.Println(v55, v55)
            var v59 float32 = (v54 + 3.14159)
            var v60 int = int(v3)
            var v61 int32 = 393
            _ = fi57
            _ = v58
            _ = v59
            _ = v60
            _ = v61
        }
        switch int(v52) {
        case -780:
            if (!true) {
                v62 := GenS2{F0: v3}
                _ = v62
            }
        }
        _ = fi56
    }
    v54 = (3.14159 * float32((1.5 * (0.1 * (-2.25 / (3.14159 * float64(-2.25)))))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
