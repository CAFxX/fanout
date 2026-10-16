package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 bool }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 uint8; F1 int; F2 bool }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 uint8; F2 uint8 }
func genFn0(p0 uint16, p1 uint64) uint16 {
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: 1, F1: 42, F2: "hello"}
        _ = fi5
        _ = v6
    }
    v7 := []int{-856, 4294967295, -1, 32767}
    _ = p0
    _ = p1
    _ = v7
    return (p0 & (p0 / (((p0 & 255)) | 1)))
}
func genFn1(p0 uint32, p1 uint64) uint {
    if (!(false || (42 != 255))) {
        v8 := []int{379, 958, -356}
        _ = v8
    } else {
        v9 := GenS2{F0: (^(p0))}
        v10 := GenS3{F0: p1, F1: -1, F2: (p1 | 1)}
        _ = v10
        _ = v9
    }
    v11 := []int{}
    switch "" {
    case "":
        p0 += (p0 * p0)
    case "\n\t":
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    case "foo bar":
        for fi12 := 0; fi12 < 16; fi12++ {
            p0 -= p0
            var v13 uint64 = p1
            var v14 bool = true
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    p0 = 42
    _ = p0
    _ = p1
    _ = v11
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 1
    v3 := make(map[int]bool)
    var v4 uint32 = 64
    v15 := GenS1{F0: -134, F1: 4294967295, F2: "\n\t"}
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(fi16)
        _ = fi16
    }
    v17 := v3[-2]
    v18 := GenS4{F0: 23, F1: 0, F2: true}
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 bool = v17
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v1["\n\t"] = 2
        if ((^(fi21)) >= fi21) {
            v4 = v4
        }
        for fi22 := 0; fi22 < 14; fi22++ {
            if ((42 >= 18446744073709551615) || (true || true)) {
                delete(v1, (("\n\t" + "\n\t") + "foo bar"))
                v23 := GenS1{F0: 409, F1: 221, F2: ((("" + (("a" + "\n\t") + "\n\t")) + (("\n\t" + (("" + "\n\t") + "a")) + "foo bar")) + (("hello" + "") + ("\n\t" + "\n\t")))}
                v17 = (v4 < v4)
                _ = v23
            } else {
                var v24 int64 = 188
                if v17 {
                    delete(v1, ("hello" + "hello"))
                    v24 = v24
                    if (-2.25 != (3.14159 + 1.5)) {
                        switch ("foo bar" + "a") {
                        case "":
                            v17 = (v17 || (v17 || v17))
                            switch "\n\t" {
                            case "a":
                                var v25 bool = v17
                                _ = v25
                            case "foo bar":
                                v26 := GenS5{F0: 0, F1: fi21}
                                _ = v26
                            }
                        default:
                            switch "hello" {
                            case "":
                                if false {
                                    if (v24 > 1) {
                                        var v27 uint8 = 84
                                        v17 = (641 > 2)
                                        var v28 float32 = (0.1 * 1.5)
                                        var v29 int16 = -859
                                        v30 := v3[546]
                                        v31 := GenS3{F0: 1, F1: v24, F2: 4294967295}
                                        v2 -= v2
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                }
                            default:
                                v32 := GenS2{F0: v4}
                                _ = v32
                            }
                        }
                        v3[(^(fi22))] = ("foo bar" < "a")
                        switch 2147483647 {
                        case -1:
                            v33 := []int{fi22, 32768, (fi21 >> ((42) & 63)), fi21}
                            _ = v33
                        }
                    } else {
                        v4 = 65535
                    }
                } else {
                    v24 -= v24
                    switch ((("hello" + "\n\t") + "\n\t") + "a") {
                    case "":
                        v34 := GenS2{F0: (v4 ^ v4)}
                        _ = v34
                    default:
                        switch "a" {
                        case "":
                            var v35 bool = (32767 <= 42)
                            _ = v35
                        }
                    }
                    v2 = genFn0(uint16(v4), 4294967295)
                }
                if false {
                    if ((v17 && v17) && (v17 && (!true))) {
                        v36 := []int{(-(fi21)), (fi22 / (((-(fi21))) | 1)), fi22, fi22}
                        delete(v1, "a")
                        _ = v36
                    }
                    v24 = (v24 >> ((425) & 63))
                    switch "foo bar" {
                    case "\n\t":
                        switch "\n\t" {
                        case "foo bar":
                            v17 = (("hello" + "hello") <= ("hello" + ""))
                        case "":
                            v17 = (true || v17)
                        default:
                            switch fi22 {
                            case 127:
                                var v37 int16 = 32767
                                _ = v37
                            default:
                                var v38 int32 = 2147483647
                                _ = v38
                            }
                        }
                    default:
                        v1["\n\t"] = v4
                        v39 := GenS6{F0: 0, F1: (fi22 ^ -137), F2: (!false)}
                        v2 += genFn0(genFn0(v2, 2), 42)
                        _ = v39
                    }
                    v40 := GenS4{F0: (v4 & 816), F1: 4294967295, F2: ((3.14159 != (((1e10 * 1.5) + ((3.14159 * (1.5 + 3.14159)) - 1e10)) + 0.1)) || false)}
                    var v41 uint8 = 42
                    _ = v40
                    _ = v41
                }
                v42 := []int{}
                fmt.Println(v24, v17, v4)
                if len(v42) > 0 {
                    v42[0] = (fi21 ^ fi21)
                }
                v2 = (v2 >> ((255) & 63))
                _ = v24
                _ = v42
            }
            _ = fi22
        }
        _ = fi21
    }
    var v43 float64 = 0.0
    var v44 uint32 = (v4 & (v4 >> ((18446744073709551615) & 63)))
    v17 = (v17 || (v17 && (-2147483648 <= 42)))
    if v17 {
        v45 := GenS4{F0: v4, F1: 18446744073709551615, F2: false}
        _ = v45
    } else {
        var v46 uint = 255
        v46 = (^(v46))
        _ = v46
    }
    switch ("a" + "\n\t") {
    case "":
        if ((!v17) && (true && v17)) {
            v43 = v43
            var v47 uint = genFn1((v44 - v4), 614)
            var v48 float64 = (v43 - v43)
            _ = v47
            _ = v48
        } else {
            var v49 int8 = 0
            _ = v49
        }
    case "hello":
        var v50 int64 = 9223372036854775807
        _ = v50
    default:
        v44 = (v4 << ((430) & 63))
    }
    if (128 >= uint8(v2)) {
        v51 := []int{127, -9223372036854775808, 127, -342}
        switch 9223372036854775807 {
        case 0:
            v17 = (v17 && v17)
        case -1:
            if len(v51) >= 2 {
                v51 = v51[1:len(v51)]
            }
        case 341:
            var v52 int = -1
            _ = v52
        default:
            var v53 float32 = 0.1
            _ = v53
        }
        v54 := GenS7{F0: 32768, F1: 1}
        _ = v51
        _ = v54
    } else {
        if (2 <= 18446744073709551615) {
            var v55 int32 = 0
            _ = v55
        }
        var v56 bool = false
        _ = v56
    }
    v4 = v4
    var v57 uint = 255
    var v58 uint8 = 235
    if true {
        if v17 {
            v59 := GenS8{F0: 771}
            switch "hello" {
            case "":
                var v60 float64 = float64((1.5 - 1.5))
                _ = v60
            }
            v61 := []int{-821, 128, -9223372036854775808, 2147483647}
            _ = v59
            _ = v61
        }
        v62 := []int{}
        for fi63 := 0; fi63 < 18; fi63++ {
            if len(v62) > 0 {
                v62[0] = (fi63 % (((fi63 % (((fi63 & (fi63 >> (((v57 >> (((v57 % (((v57 >> (((v57 << ((42) & 63))) & 63))) | 1))) & 63))) & 63)))) | 1))) | 1))
            }
            fmt.Println(fi63, v4, fi63)
            _ = fi63
        }
        _ = v62
    } else {
        for fi64 := 0; fi64 < 7; fi64++ {
            for fi65 := 0; fi65 < 4; fi65++ {
                fmt.Println(v17)
                fmt.Println(v44)
                _ = fi65
            }
            for fi66 := 0; fi66 < 2; fi66++ {
                if (!((fi66 % ((fi64) | 1)) < fi64)) {
                    if v17 {
                        var v67 float32 = (3.14159 / 1.5)
                        v3[-2147483648] = (v17 && ((v58 + 0) != 127))
                        _ = v67
                    } else {
                        if (1e10 == (-2.25 + (v43 / v43))) {
                            v68 := GenS9{F0: 1000, F1: uint8(v4), F2: v58}
                            if ("a" > "") {
                                if v17 {
                                    var v69 float32 = 1.5
                                    _ = v69
                                }
                                v70 := GenS2{F0: v44}
                                var v71 int16 = 1000
                                _ = v70
                                _ = v71
                            }
                            _ = v68
                        } else {
                            var v72 int8 = 0
                            var v73 int8 = v72
                            _ = v72
                            _ = v73
                        }
                    }
                    if (1 <= -724) {
                        v2 += v2
                        fmt.Println(fi66)
                        var v74 uint32 = (v4 % (((v4 / (((v44 * (v44 + 65535))) | 1))) | 1))
                        _ = v74
                    } else {
                        fmt.Println(fi64, v44)
                        var v75 bool = false
                        _ = v75
                    }
                } else {
                    if (((1.5 / 1e10) / 1.0) > (1.5 / 0.1)) {
                        v76 := GenS9{F0: 42, F1: (v58 ^ (v58 - v58)), F2: uint8(v2)}
                        _ = v76
                    } else {
                        var v77 uint = 4294967295
                        _ = v77
                    }
                    v43 = v43
                }
                v57 += 269
                _ = fi66
            }
            _ = fi64
        }
        v78 := GenS4{F0: (v44 * uint32(v57)), F1: 0, F2: (v17 || ((!false) || (2147483647 >= 111)))}
        _ = v78
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v57
    _ = v58
}
