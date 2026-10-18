package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 string; F1 int64; F2 bool }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int32) uint8 {
    v6 := GenS1{F0: 565, F1: "foo bar", F2: 206}
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{fi7, fi7, (fi7 - 32767), (^(fi7))}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 255
}

func main() {
    var v1 int8 = 42
    var v2 uint = 173
    var v3 float32 = (0.1 - -2.25)
    v4 := []uint32{157}
    var v5 uint32 = 65535
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(fi9, fi9)
        if (!true) {
            switch "\n\t" {
            case "":
                if len(v4) > 0 {
                    v4[0] = 1
                }
            case "foo bar":
                for fi10 := 0; fi10 < 13; fi10++ {
                    if false {
                        v11 := GenS1{F0: 841, F1: "", F2: 2}
                        var v12 int32 = -444
                        v2 -= (v2 | (v2 & v2))
                        _ = v11
                        _ = v12
                    }
                    v2 -= (v2 - v2)
                    if false {
                        if true {
                            var v13 uint16 = 42
                            _ = v13
                        } else {
                            switch (fi9 / (((fi10 * (fi9 * fi9))) | 1)) {
                            case 32767:
                                v2 = (v2 >> ((18446744073709551615) & 63))
                            case -9223372036854775808:
                                v3 = (float32(0.1) - 0.1)
                                v14 := GenS2{F0: 2, F1: 18446744073709551615}
                                _ = v14
                            case -989:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            if true {
                                var v15 float32 = v3
                                v15 += (v3 * v15)
                                switch "foo bar" {
                                case "hello":
                                    switch 32768 {
                                    case 940:
                                        v16 := GenS3{F0: "\n\t", F1: -954, F2: false}
                                        _ = v16
                                    case 9223372036854775807:
                                        v5 = (v5 << ((v2) & 63))
                                    }
                                }
                                _ = v15
                            }
                            var v17 uint = (v2 + (v2 / (((v2 << (((v2 | v2)) & 63))) | 1)))
                            var v18 uint = (^(v17))
                            _ = v17
                            _ = v18
                        }
                        switch "\n\t" {
                        case "\n\t":
                            var v19 bool = false
                            v20 := GenS4{F0: 32767, F1: v1}
                            var v21 uint16 = uint16(v5)
                            _ = v19
                            _ = v20
                            _ = v21
                        case "":
                            var v22 float64 = 3.14159
                            _ = v22
                        }
                        var v23 string = "\n\t"
                        _ = v23
                    } else {
                        var v24 int = 32768
                        if true {
                            var v25 string = ""
                            _ = v25
                        }
                        _ = v24
                    }
                    if false {
                        v2 -= (v2 << ((v2) & 63))
                        v5 += (v5 >> (((v2 | v2)) & 63))
                    }
                    v3 = v3
                    _ = fi10
                }
            case "\n\t":
                for fi26 := 0; fi26 < 17; fi26++ {
                    if false {
                        if (714 == 1) {
                            var v27 int64 = 434
                            if false {
                                var v28 int64 = v27
                                switch ("\n\t" + "") {
                                case "\n\t":
                                    var v29 float64 = (3.14159 / 3.14159)
                                    _ = v29
                                case "":
                                    var v30 int64 = v28
                                    _ = v30
                                case "hello":
                                    if true {
                                        if (true || true) {
                                            fmt.Println(v27, v3)
                                        }
                                    } else {
                                        v2 -= 1
                                        v2 -= 0
                                        var v31 int64 = v27
                                        _ = v31
                                    }
                                default:
                                    v32 := GenS4{F0: 1000, F1: (v1 / ((v1) | 1))}
                                    _ = v32
                                }
                                _ = v28
                            }
                            if true {
                                switch fi9 {
                                case 450:
                                    fmt.Println(fi26, v27)
                                case -1:
                                    v3 = 0.0
                                case 127:
                                    v4 = append(v4, (^(v5)))
                                default:
                                    var v33 int16 = 1
                                    _ = v33
                                }
                                if len(v4) > 0 {
                                    v4[0] = 42
                                }
                                if true {
                                    if len(v4) > 0 {
                                        v4[0] = v5
                                    }
                                }
                                switch "" {
                                case "hello":
                                    if len(v4) > 0 {
                                        v4[0] = v5
                                    }
                                case "":
                                    v27 += (v27 ^ v27)
                                case "foo bar":
                                    v1 -= v1
                                }
                                var v34 int8 = -1
                                _ = v34
                            }
                            if true {
                                if ((((false || (32767 != 1)) || (fi26 != -5)) || false) || ((!(("hello" + "") <= (("foo bar" + "hello") + "foo bar"))) && false)) {
                                    if len(v4) > 0 {
                                        v4[0] = (v5 >> (((v2 + 324)) & 63))
                                    }
                                    fmt.Println(v27)
                                    var v35 uint64 = 0
                                    _ = v35
                                }
                                var v36 string = ("hello" + "foo bar")
                                fmt.Println(v27)
                                v1 = (v1 - 0)
                                switch "\n\t" {
                                case "a":
                                    v1 -= -87
                                case "":
                                    var v37 uint32 = v5
                                    _ = v37
                                case "\n\t":
                                    fmt.Println(v1, fi26)
                                default:
                                    v5 += (v5 & v5)
                                }
                                _ = v36
                            }
                            _ = v27
                        } else {
                            switch "\n\t" {
                            case "hello":
                                v38 := GenS1{F0: 2, F1: "a", F2: 255}
                                _ = v38
                            case "\n\t":
                                var v39 int64 = 1
                                _ = len(v4)
                                var v40 uint64 = 812
                                _ = v39
                                _ = v40
                            default:
                                v41 := GenS3{F0: (("\n\t" + "") + ("foo bar" + "foo bar")), F1: 42, F2: false}
                                _ = v41
                            }
                            if true {
                                var v42 float32 = 0.1
                                _ = v42
                            } else {
                                if (!true) {
                                    var v43 int = fi9
                                    fmt.Println(v3)
                                    _ = v43
                                } else {
                                    fmt.Println(fi26)
                                }
                                var v44 uint64 = 580
                                _ = v44
                            }
                            var v45 float32 = float32(1.5)
                            _ = v45
                        }
                        v2 = 0
                        v4 = append(v4, (v5 >> ((v2) & 63)))
                    }
                    v5 -= 4294967295
                    v3 -= (0.1 / v3)
                    _ = fi26
                }
            }
            for fi46 := 0; fi46 < 13; fi46++ {
                switch "a" {
                case "":
                    v2 = 4294967295
                case "foo bar":
                    var v47 uint8 = 128
                    _ = v47
                }
                if ((2 <= 1) && false) {
                    var v48 uint = v2
                    v5 += 2147483647
                    v4 = append(v4, (v5 % (((v5 - v5)) | 1)))
                    _ = v48
                }
                _ = fi46
            }
        }
        if len(v4) > 0 {
            v4[0] = 2
        }
        _ = fi9
    }
    v49 := GenS5{F0: 255}
    var v50 int64 = -804
    var v51 uint16 = 0
    v4 = append(v4, 42)
    var v52 uint64 = 0
    var v53 int32 = -1
    var v54 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
