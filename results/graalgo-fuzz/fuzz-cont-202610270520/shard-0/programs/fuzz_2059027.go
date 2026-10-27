package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 bool }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint16; F1 bool; F2 int64 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 string; F1 int32 }
func genFn0(p0 string, p1 bool, p2 int8) uint64 {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 string = ""
        p2 -= p2
        p1 = (!(!false))
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: -32768, F2: p1}
    var v9 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 852
}
func genFn1(p0 uint8, p1 uint) uint8 {
    var v10 int16 = 787
    v11 := []int{271, 0, -49}
    var v12 uint16 = 65535
    switch 32767 {
    case 2147483647:
        var v13 int64 = -9223372036854775808
        _ = v13
    case 32767:
        var v14 uint8 = 255
        _ = v14
    case -2147483648:
        v12 += (v12 % (((v12 % ((32767) | 1))) | 1))
    }
    var v15 float32 = 0.1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    return (p0 ^ 67)
}

func main() {
    var v1 int8 = 15
    var v2 int16 = 42
    v3 := []uint32{2, 1, 42}
    v4 := []uint32{4294967295, 65535, 583, 273}
    var v5 uint8 = 254
    if (-265 <= 310) {
        var v16 int8 = (-(v1))
        v5 = genFn1(v5, 4294967295)
        _ = v16
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            v2 = v2
            if (((((fi17 + (fi17 - (fi17 * (fi17 + fi17)))) == (-(fi17))) || (!false)) || (true || (false && false))) || (false && true)) {
                var v18 int32 = -1
                switch fi17 {
                case 32768:
                    for fi19 := 0; fi19 < 4; fi19++ {
                        fmt.Println(fi19, fi17, fi19)
                        _ = fi19
                    }
                case -649:
                    if len(v4) > 0 {
                        v4[0] = 65535
                    }
                case -9223372036854775808:
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                }
                fmt.Println(fi17, v2, fi17)
                _ = v18
            } else {
                fmt.Println(v5, fi17)
            }
            switch fi17 {
            case 42:
                var v20 int64 = 42
                _ = v20
            case -1:
                v1 = (^(v1))
            case 32767:
                switch "" {
                case "a":
                    fmt.Println(v1, v5, fi17)
                case "hello":
                    v21 := GenS1{F0: 2147483647, F1: (v2 - v2), F2: true}
                    _ = v21
                default:
                    v1 = (v1 << ((0) & 63))
                }
            }
            _ = fi17
        }
        var v22 string = ("hello" + "")
        _ = v22
    }
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(fi23)
            v1 -= (v1 | v1)
            for fi24 := 0; fi24 < 9; fi24++ {
                fmt.Println(fi24, fi23)
                if (208 < 0) {
                    if ((!((v1 * v1) > (^(v1)))) && false) {
                        v25 := GenS1{F0: -885, F1: 42, F2: (42 > 42)}
                        if (false || (true && false)) {
                            var v26 uint8 = 255
                            _ = v26
                        } else {
                            v2 = (v2 % (((^(v2))) | 1))
                            switch "" {
                            case "":
                                if true {
                                    v2 += -1
                                    var v27 string = (("a" + "hello") + ("\n\t" + ("hello" + ("hello" + "a"))))
                                    var v28 string = (((v27 + (v27 + v27)) + v27) + v27)
                                    _ = v27
                                    _ = v28
                                } else {
                                    if false {
                                        if len(v4) > 0 {
                                            v4[0] = 42
                                        }
                                        if ((((3.14159 / 1.0) > -2.25) || (false && (false || true))) && (!(!true))) {
                                            switch "\n\t" {
                                            case "":
                                                if len(v4) > 0 {
                                                    v4[0] = 2147483648
                                                }
                                            case "\n\t":
                                                v29 := GenS2{F0: (^(fi24)), F1: 9223372036854775807}
                                                _ = v29
                                            }
                                        }
                                        var v30 uint = 738
                                        _ = v30
                                    } else {
                                        fmt.Println(fi24, v2, v5)
                                        var v31 uint64 = 0
                                        if true {
                                            v31 = v31
                                            fmt.Println(v2)
                                            v1 += (v1 * 42)
                                        }
                                        _ = v31
                                    }
                                    v3 = append(v3, 959)
                                    if true {
                                        if (!(0.0 == 0.0)) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            v5 = genFn1((^(v5)), 30)
                                        }
                                        switch fi24 {
                                        case 32767:
                                            if true {
                                                v32 := GenS3{F0: 0, F1: (false || true), F2: int64(v5)}
                                                var v33 int64 = -864
                                                v2 = (v2 >> ((4294967295) & 63))
                                                _ = v32
                                                _ = v33
                                            }
                                        case 32768:
                                            var v34 int8 = v1
                                            _ = v34
                                        }
                                        var v35 int16 = v2
                                        _ = v35
                                    }
                                }
                            default:
                                fmt.Println(v1, fi24)
                                if false {
                                    v36 := GenS3{F0: 255, F1: true, F2: -27}
                                    var v37 uint = 220
                                    _ = len(v3)
                                    _ = v36
                                    _ = v37
                                } else {
                                    v38 := GenS4{F0: -290, F1: v1}
                                    if (0 < 278) {
                                        var v39 float32 = 3.14159
                                        v5 = 78
                                        if (true && (!true)) {
                                            var v40 uint64 = 333
                                            var v41 uint16 = 833
                                            _ = v40
                                            _ = v41
                                        } else {
                                            v42 := GenS3{F0: 173, F1: false, F2: -711}
                                            var v43 int64 = 279
                                            var v44 uint16 = 1
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        }
                                        v2 = (v2 << ((444) & 63))
                                        _ = v39
                                    }
                                    _ = v38
                                }
                                v45 := GenS4{F0: -2147483648, F1: (v1 | v1)}
                                switch (fi23 * (fi23 << ((561) & 63))) {
                                case 9223372036854775807:
                                    v2 += v2
                                case 2147483647:
                                    var v46 string = "\n\t"
                                    _ = v46
                                default:
                                    var v47 int8 = 65
                                    _ = v47
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v45
                            }
                        }
                        if false {
                            v2 += (v2 % (((v2 ^ v2)) | 1))
                            if ((-55 <= 1000) || ((!true) || (!("hello" <= "a")))) {
                                var v48 uint64 = 0
                                var v49 int32 = -2147483648
                                v2 += 338
                                _ = v48
                                _ = v49
                            } else {
                                v2 += (-(v2))
                                var v50 string = ""
                                var v51 float32 = ((1.5 / -2.25) + 1.5)
                                _ = v50
                                _ = v51
                            }
                        }
                        _ = v25
                    }
                }
                var v52 uint16 = 448
                _ = fi24
                _ = v52
            }
            _ = fi23
        }
    }
    var v53 int16 = (v2 - v2)
    fmt.Println(v5, v53)
    v1 -= (v1 | -105)
    if len(v3) > 0 {
        v3[0] = 564
    }
    fmt.Println(v5, v2, v53)
    v54 := GenS5{F0: "foo bar", F1: 42}
    fmt.Println(v1)
    for fi55 := 0; fi55 < 4; fi55++ {
        var v56 int16 = v2
        _ = fi55
        _ = v56
    }
    switch "" {
    case "\n\t":
        var v57 uint64 = 2
        _ = v57
    default:
        var v58 uint64 = 440
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
}
