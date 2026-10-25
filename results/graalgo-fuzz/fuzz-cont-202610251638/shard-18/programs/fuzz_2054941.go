package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int; F1 int64; F2 int }
func genFn0(p0 uint32) uint64 {
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(fi5, p0)
        for fi6 := 0; fi6 < 16; fi6++ {
            p0 = (p0 << ((0) & 63))
            v7 := GenS1{F0: 255}
            p0 = (p0 >> ((42) & 63))
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    switch "a" {
    case "\n\t":
        p0 -= (p0 & 643)
    case "a":
        var v8 float64 = 3.14159
        _ = v8
    }
    _ = p0
    return 0
}
func genFn1(p0 int16, p1 int) uint16 {
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(p1)
        var v10 int64 = 797
        _ = fi9
        _ = v10
    }
    var v11 int = p1
    if false {
        v12 := GenS2{F0: 326}
        _ = v12
    }
    v13 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v13
    return 65535
}

func main() {
    var v1 int = 32767
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]int)
    var v4 float32 = ((1.5 + (0.0 - 1e10)) - float32(0.1))
    v14 := GenS3{F0: (359 < 127)}
    v4 += v4
    v3[v1] = v1
    var v15 uint = 531
    if ("" > (("foo bar" + ("foo bar" + (("foo bar" + "\n\t") + "foo bar"))) + ("foo bar" + ("a" + "foo bar")))) {
        if false {
            v2 += v2
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 bool = false
                v18 := GenS4{F0: genFn1(-32768, fi16), F1: -114}
                v3[v1] = (fi16 ^ 9223372036854775807)
                _ = fi16
                _ = v17
                _ = v18
            }
            if (!(true && true)) {
                v4 = v4
            } else {
                v19 := GenS1{F0: 2}
                var v20 uint = (^(v15))
                v1 += v1
                _ = v19
                _ = v20
            }
        } else {
            for fi21 := 0; fi21 < 5; fi21++ {
                fmt.Println(v15, v4, v2)
                switch -534 {
                case 2147483647:
                    v22 := GenS3{F0: false}
                    _ = v22
                case 1:
                    var v23 float64 = (1e10 - float64(((v4 / v4) - v4)))
                    _ = v23
                }
                v24 := GenS3{F0: true}
                _ = fi21
                _ = v24
            }
            if false {
                switch (v1 + (v1 / (((v1 - v1)) | 1))) {
                case -2147483648:
                    fmt.Println(v2)
                default:
                    v25 := GenS5{F0: v2}
                    _ = v25
                }
                switch "foo bar" {
                case "foo bar":
                    if false {
                        _ = len(v3)
                        v15 = (^(v15))
                    }
                case "hello":
                    if false {
                        var v26 string = "\n\t"
                        if (-600 != -47) {
                            switch (v26 + v26) {
                            case "foo bar":
                                v4 = v4
                            case "\n\t":
                                v1 = v1
                            case "hello":
                                var v27 int8 = 42
                                _ = v27
                            }
                            switch (-(v1)) {
                            case 127:
                                v1 += -1
                            case 32768:
                                v15 = v15
                            case 0:
                                var v28 uint8 = 128
                                _ = v28
                            }
                            v26 = ("foo bar" + ("\n\t" + (v26 + ("a" + v26))))
                        }
                        _ = v26
                    }
                case "\n\t":
                    for fi29 := 0; fi29 < 12; fi29++ {
                        v30 := []int{v1, fi29, -9223372036854775808}
                        for fi31 := 0; fi31 < 9; fi31++ {
                            fmt.Println(fi31, v2, fi29)
                            _ = fi31
                        }
                        v15 -= v15
                        _ = fi29
                        _ = v30
                    }
                }
                if (!false) {
                    _ = len(v3)
                } else {
                    if (false || true) {
                        var v32 uint8 = 0
                        var v33 int16 = -32768
                        _ = v32
                        _ = v33
                    }
                    var v34 int8 = -104
                    v35 := GenS5{F0: 18446744073709551615}
                    _ = v34
                    _ = v35
                }
            } else {
                var v36 string = ""
                _ = v36
            }
        }
        if true {
            if false {
                v37 := GenS4{F0: genFn1(32767, v1), F1: 97}
                _ = v37
            }
            fmt.Println(v15, v1)
            v15 += (v15 % ((uint(v2)) | 1))
        }
    } else {
        if ((3.14159 > 1.5) || (true || (-2147483648 < -2147483648))) {
            delete(v3, (v1 | (v1 ^ (v1 - v1))))
            switch ("foo bar" + "foo bar") {
            case "hello":
                fmt.Println(v15, v1, v2)
            case "":
                for fi38 := 0; fi38 < 17; fi38++ {
                    var v39 string = ("\n\t" + "")
                    for fi40 := 0; fi40 < 14; fi40++ {
                        var v41 uint64 = (v2 % (((v2 / (((v2 & uint64(v15))) | 1))) | 1))
                        v39 = v39
                        _ = fi40
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                }
            }
        } else {
            var v42 uint32 = uint32(v2)
            _ = v42
        }
        var v43 int32 = 2147483647
        _ = v43
    }
    v4 -= (0.0 * (-2.25 + (v4 * 3.14159)))
    v1 = 1
    var v44 int32 = 42
    for fi45 := 0; fi45 < 8; fi45++ {
        var v46 int8 = 30
        if (true && (3.14159 == 1.5)) {
            var v47 int32 = v44
            fmt.Println(v44)
            var v48 float32 = float32(-2.25)
            _ = v47
            _ = v48
        } else {
            for fi49 := 0; fi49 < 17; fi49++ {
                fmt.Println(v46, v4)
                _ = fi49
            }
            v50 := []int{fi45, (fi45 - (fi45 & fi45))}
            _ = v50
        }
        _ = fi45
        _ = v46
    }
    if (int64(v44) != -9223372036854775808) {
        var v51 int32 = (v44 | v44)
        if false {
            var v52 string = "a"
            v53 := GenS4{F0: 2, F1: -1}
            var v54 float64 = float64((v4 / ((1.5 + v4) - (v4 + v4))))
            _ = v52
            _ = v53
            _ = v54
        } else {
            if (false || ((false || (float32((0.0 - 1e10)) > 0.0)) || false)) {
                var v55 int32 = v51
                if false {
                    for fi56 := 0; fi56 < 17; fi56++ {
                        v57 := GenS6{F0: 127}
                        v15 -= 4294967295
                        v58 := []int{(v1 | (^(v1))), fi56}
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                } else {
                    if false {
                        fmt.Println(v55, v55, v44)
                        var v59 uint16 = 2
                        v60 := []int{-986, (-(v1)), int(v51)}
                        _ = v59
                        _ = v60
                    }
                    v4 = v4
                    fmt.Println(v15, v51, v15)
                }
                _ = v55
            }
            var v61 uint64 = genFn0(4294967295)
            _ = v61
        }
        v62 := []int{(v1 + v1), v1, v1}
        _ = v51
        _ = v62
    }
    if false {
        var v63 bool = true
        for fi64 := 0; fi64 < 20; fi64++ {
            v1 = (v1 * v1)
            v2 = 562
            switch v1 {
            case 127:
                var v65 int16 = 32767
                _ = v65
            case 4294967295:
                v66 := GenS4{F0: uint16(v2), F1: -92}
                _ = v66
            default:
                var v67 string = "hello"
                _ = v67
            }
            _ = fi64
        }
        v2 += (v2 % (((v2 - (v2 + (v2 / (((v2 / ((2) | 1))) | 1))))) | 1))
        _ = v63
    } else {
        var v68 bool = false
        _ = v68
    }
    var v69 bool = (genFn1(1, 352) >= 2)
    if (!((v4 * float32(((1.5 + (1e10 - 0.1)) - 0.1))) < v4)) {
        v70 := v3[128]
        _ = v70
    } else {
        fmt.Println(v4, v4, v69)
        var v71 uint16 = 2
        v15 += uint(v71)
        _ = v71
    }
    var v72 float64 = (float64(v4) * ((0.1 / 0.1) * 0.1))
    if (!v69) {
        v15 += (v15 - v15)
        var v73 uint = v15
        for fi74 := 0; fi74 < 2; fi74++ {
            switch fi74 {
            case 0:
                v73 = 42
            case -293:
                if false {
                    _ = len(v3)
                    fmt.Println(v69, v44)
                    var v75 float32 = v4
                    _ = v75
                }
            }
            if true {
                var v76 string = "hello"
                v77 := GenS2{F0: 1}
                v69 = (!(((v72 - 3.14159) < (0.0 / v72)) || false))
                _ = v76
                _ = v77
            } else {
                fmt.Println(v69)
                var v78 int8 = 42
                var v79 bool = (!(942 != (fi74 - fi74)))
                _ = v78
                _ = v79
            }
            fmt.Println(v4, v69, v2)
            _ = fi74
        }
        _ = v73
    } else {
        v44 = v44
        v80 := GenS7{F0: v1, F1: 0, F2: (v1 / ((int(v2)) | 1))}
        _ = v80
    }
    fmt.Println(v4, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v69:", v69)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v69
    _ = v72
}
