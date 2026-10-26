package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 string; F2 string }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int64; F1 bool; F2 uint32 }
func genFn0() uint16 {
    v5 := GenS1{F0: -2147483648}
    var v6 int64 = 140
    switch "a" {
    case "foo bar":
        v7 := GenS1{F0: 0}
        _ = v7
    default:
        v6 += 42
    }
    _ = v5
    _ = v6
    return 369
}
func genFn1(p0 uint32, p1 uint16) int32 {
    fmt.Println(p1, p0)
    if true {
        p1 -= p1
    } else {
        v8 := []int{-995}
        p1 += 744
        _ = v8
    }
    if false {
        fmt.Println(p1, p0)
        p0 = (^(p0))
    } else {
        v9 := GenS1{F0: 263}
        var v10 int64 = -228
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 954}
    switch ("hello" + "hello") {
    case "":
        var v12 int16 = 32767
        _ = v12
    default:
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int16 = 0
            switch ("foo bar" + "foo bar") {
            case "a":
                if false {
                    if (true || true) {
                        for fi15 := 0; fi15 < 18; fi15++ {
                            p1 = (p1 * (p1 / ((p1) | 1)))
                            _ = fi15
                        }
                    }
                    var v16 float32 = (1e10 / -2.25)
                    _ = v16
                } else {
                    v17 := []int{(fi13 + (fi13 - (fi13 % (((fi13 - fi13)) | 1))))}
                    _ = v17
                }
            case "":
                var v18 string = ("foo bar" + "")
                _ = v18
            case "foo bar":
                var v19 uint32 = 65535
                _ = v19
            }
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v11
    return -2147483648
}

func main() {
    var v1 int8 = 127
    var v2 string = ""
    var v3 float32 = ((1e10 - 0.0) - 1.5)
    var v4 int8 = v1
    v1 += (^(v4))
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 int64 = 226
        switch (v2 + v2) {
        case "hello":
            v22 := []int{(fi20 | (fi20 & 4294967295)), 32768, fi20, (fi20 / ((fi20) | 1))}
            _ = v22
        default:
            var v23 string = "\n\t"
            _ = v23
        }
        v24 := GenS1{F0: 1}
        _ = fi20
        _ = v21
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        if (float32(0.0) != 1.5) {
            var v26 uint64 = 2
            v2 = v2
            v3 += (v3 * (v3 / v3))
            _ = v26
        } else {
            v27 := []int{(fi25 | (fi25 % (((fi25 ^ (fi25 | (fi25 ^ -687)))) | 1)))}
            _ = v27
        }
        _ = fi25
    }
    v28 := GenS1{F0: 609}
    for fi29 := 0; fi29 < 13; fi29++ {
        v1 -= 56
        v30 := GenS2{F0: (fi29 << ((1) & 63)), F1: ("" + ("hello" + v2)), F2: v2}
        v31 := GenS2{F0: (fi29 + (fi29 + (fi29 ^ fi29))), F1: "", F2: "\n\t"}
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 bool = (!true)
    v33 := GenS1{F0: -2147483648}
    v1 -= v1
    for fi34 := 0; fi34 < 14; fi34++ {
        var v35 int16 = -129
        if (127 < 42) {
            switch 127 {
            case -1:
                v35 += v35
            }
        } else {
            if false {
                var v36 int32 = 1
                _ = v36
            } else {
                switch (v2 + v2) {
                case "hello":
                    switch fi34 {
                    case -9223372036854775808:
                        if (v35 < v35) {
                            var v37 uint64 = 42
                            switch (v2 + ((v2 + "hello") + ((("" + v2) + ("\n\t" + ((((((v2 + v2) + "") + "\n\t") + v2) + (("foo bar" + v2) + v2)) + ("a" + ((("a" + v2) + v2) + "hello"))))) + ((("a" + ("" + "hello")) + "foo bar") + "hello")))) {
                            case "\n\t":
                                if (("\n\t" >= v2) || true) {
                                    if ((fi34 + (fi34 & (^(fi34)))) != (fi34 % ((fi34) | 1))) {
                                        v2 = "foo bar"
                                        if ((v3 < float32(0.1)) && (13 > v4)) {
                                            switch -1 {
                                            case 32767:
                                                if (-503 < 42) {
                                                    v35 = (^(v35))
                                                    switch v2 {
                                                    case "":
                                                        for fi38 := 0; fi38 < 8; fi38++ {
                                                            var v39 uint16 = 65535
                                                            switch v2 {
                                                            case "a":
                                                                fmt.Println(v35)
                                                            }
                                                            var v40 uint64 = (v37 << ((4294967295) & 63))
                                                            v41 := GenS1{F0: genFn1(77, v39)}
                                                            _ = fi38
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                    case "foo bar":
                                                        v42 := []int{(fi34 - fi34), fi34}
                                                        _ = v42
                                                    }
                                                    v43 := GenS2{F0: fi34, F1: "\n\t", F2: v2}
                                                    _ = v43
                                                }
                                            case 127:
                                                var v44 int = fi34
                                                _ = v44
                                            }
                                            v1 = (v4 >> ((1) & 63))
                                        } else {
                                            v45 := GenS2{F0: fi34, F1: v2, F2: "\n\t"}
                                            v46 := []int{(^(fi34)), fi34, (fi34 % ((fi34) | 1)), (fi34 % ((fi34) | 1))}
                                            _ = v45
                                            _ = v46
                                        }
                                        v32 = (((0.0 + ((0.1 + (((float64((v3 / (-2.25 - v3))) * 1e10) * 0.0) / -2.25)) - 3.14159)) + ((3.14159 + -2.25) - -2.25)) >= 0.0)
                                    }
                                    v4 = (^(v4))
                                    fmt.Println(v4, v3)
                                }
                            case "a":
                                var v47 bool = ((^(v37)) < v37)
                                _ = v47
                            case "foo bar":
                                v48 := []int{0, fi34, fi34}
                                _ = v48
                            }
                            var v49 uint64 = v37
                            _ = v37
                            _ = v49
                        } else {
                            v50 := []int{4294967295, fi34}
                            for fi51 := 0; fi51 < 7; fi51++ {
                                var v52 int = (fi34 / ((fi34) | 1))
                                _ = len(v50)
                                v32 = (213 <= 4294967295)
                                fmt.Println(v3, fi34)
                                _ = fi51
                                _ = v52
                            }
                            _ = v50
                        }
                    case -864:
                        var v53 uint8 = 83
                        _ = v53
                    }
                }
                fmt.Println(v3)
            }
            fmt.Println(v32, v2, v32)
            var v54 uint16 = 255
            _ = v54
        }
        for fi55 := 0; fi55 < 15; fi55++ {
            var v56 uint32 = 923
            if ((((!v32) || true) || (("\n\t" >= "a") && true)) || v32) {
                v57 := []int{fi34, (^(fi55)), fi55}
                _ = v57
            } else {
                v56 = 69
                var v58 int8 = 0
                v59 := GenS3{F0: 2, F1: (fi55 - (fi55 * 32767))}
                _ = v58
                _ = v59
            }
            _ = fi55
            _ = v56
        }
        _ = fi34
        _ = v35
    }
    v3 -= v3
    v60 := []int{32768, 394, -1, 1}
    if (2147483647 == 1) {
        switch 42 {
        case 128:
            v32 = (32768 < 32768)
        case 9223372036854775807:
            fmt.Println(v1)
        case 2147483647:
            v61 := GenS4{F0: 0, F1: false, F2: 165}
            _ = v61
        default:
            v2 = "hello"
        }
        var v62 int8 = (v4 >> ((584) & 63))
        var v63 bool = v32
        _ = v62
        _ = v63
    }
    for fi64 := 0; fi64 < 5; fi64++ {
        var v65 int8 = (v1 / ((v1) | 1))
        _ = fi64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v60
}
