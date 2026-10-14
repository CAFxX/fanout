package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint; F2 int64 }
type GenS5 struct { F0 uint64; F1 string; F2 uint64 }
type GenS6 struct { F0 uint; F1 int16; F2 uint }
func genFn0(p0 bool, p1 uint8) int32 {
    p1 = 1
    var v5 int8 = 100
    var v6 uint8 = p1
    var v7 bool = (!p0)
    p1 -= (v6 * (v6 + 127))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[string]string)
    var v2 int = -2147483648
    v3 := make(map[uint32]bool)
    v4 := []uint32{}
    for fi8 := 0; fi8 < 14; fi8++ {
        delete(v1, ("" + ""))
        var v9 string = "foo bar"
        _ = fi8
        _ = v9
    }
    var v10 uint64 = 0
    if false {
        var v11 uint8 = 2
        switch v2 {
        case -2147483648:
            var v12 int64 = -1
            _ = v12
        case 140:
            v10 = (^(v10))
        }
        var v13 uint64 = v10
        _ = v11
        _ = v13
    } else {
        fmt.Println(v10, v2)
    }
    var v14 int = (v2 * (v2 | -531))
    v2 -= 0
    v14 -= v2
    switch "hello" {
    case "a":
        v10 -= (v10 + (v10 << ((42) & 63)))
    case "foo bar":
        var v15 float64 = (1.5 + 1e10)
        _ = v15
    case "":
        v2 = v14
    default:
        _ = len(v4)
    }
    var v16 uint64 = 234
    var v17 float64 = (float64(3.14159) + (0.1 - 3.14159))
    v14 += 9223372036854775807
    for fi18 := 0; fi18 < 20; fi18++ {
        v2 -= v2
        _ = fi18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        _ = len(v3)
        v2 = 128
        _ = fi19
    }
    v20 := GenS1{F0: v10}
    for fi21 := 0; fi21 < 3; fi21++ {
        switch (^(fi21)) {
        case 1:
            v1["foo bar"] = "hello"
        case 631:
            v16 += v10
        }
        v17 = v17
        v17 = ((v17 / 0.1) - (((1.5 / (v17 - ((3.14159 + -2.25) * v17))) - 0.1) + ((float64(1e10) + 3.14159) - ((0.0 - 0.1) + (v17 + v17)))))
        _ = fi21
    }
    v10 = (v10 - 337)
    if (!false) {
        fmt.Println(v10, v10, v10)
        for fi22 := 0; fi22 < 6; fi22++ {
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 bool = ("" == "")
                if (654 <= -744) {
                    v25 := GenS1{F0: 79}
                    var v26 uint64 = 0
                    if (0 >= 1) {
                        var v27 uint8 = 1
                        _ = v27
                    } else {
                        switch v2 {
                        case 665:
                            if (21 > 9223372036854775807) {
                                if (114 <= 2147483647) {
                                    var v28 int32 = genFn0((!(fi23 != fi22)), uint8(v16))
                                    _ = v28
                                } else {
                                    v29 := GenS2{F0: 2}
                                    switch ("a" + ("foo bar" + "hello")) {
                                    case "":
                                        v4 = append(v4, 2147483647)
                                    case "a":
                                        if (v10 == 392) {
                                            if (("" + "") <= "hello") {
                                                if true {
                                                    delete(v1, "")
                                                    if v24 {
                                                        var v30 uint16 = 32767
                                                        v17 -= v17
                                                        _ = v30
                                                    }
                                                }
                                                v10 -= v16
                                                var v31 int32 = 2147483647
                                                v10 = 42
                                                _ = v31
                                            }
                                            var v32 string = "a"
                                            fmt.Println(v16, v32)
                                            _ = v32
                                        } else {
                                            if v24 {
                                                var v33 uint32 = 65535
                                                if v24 {
                                                    var v34 float32 = (((1e10 / 3.14159) - -2.25) + (0.0 * (0.1 / 1.0)))
                                                    fmt.Println(v34)
                                                    if (true || v24) {
                                                        if false {
                                                            v14 += 2147483647
                                                            if false {
                                                                fmt.Println(v14)
                                                                var v35 int16 = 1000
                                                                var v36 uint = uint(v33)
                                                                v34 = v34
                                                                _ = v35
                                                                _ = v36
                                                            } else {
                                                                if v24 {
                                                                    v37 := GenS1{F0: 484}
                                                                    var v38 int16 = 0
                                                                    v39 := GenS3{F0: false}
                                                                    _ = v37
                                                                    _ = v38
                                                                    _ = v39
                                                                } else {
                                                                    switch (fi23 % (((fi23 & v14)) | 1)) {
                                                                    case -285:
                                                                        v34 = v34
                                                                    default:
                                                                        var v40 float64 = 0.0
                                                                        _ = v40
                                                                    }
                                                                }
                                                                var v41 string = ("foo bar" + "\n\t")
                                                                _ = v41
                                                            }
                                                            v2 = (fi22 % (((-(v14))) | 1))
                                                        }
                                                    }
                                                    _ = v34
                                                }
                                                _ = v33
                                            } else {
                                                v42 := GenS4{F0: ("a" + "hello"), F1: 827, F2: 1}
                                                _ = v42
                                            }
                                            v26 += v16
                                        }
                                    default:
                                        var v43 int64 = 0
                                        _ = v43
                                    }
                                    _ = v29
                                }
                                v14 = v14
                                if ("\n\t" != (("\n\t" + ("" + ((("a" + ("foo bar" + ("" + (("hello" + "\n\t") + "")))) + ((("" + "") + "\n\t") + "foo bar")) + ("\n\t" + "a")))) + "\n\t")) {
                                    v10 = v10
                                }
                            } else {
                                if false {
                                    fmt.Println(fi23)
                                }
                            }
                            fmt.Println(fi23)
                            _ = len(v4)
                        case 390:
                            switch fi23 {
                            case 0:
                                v24 = false
                            case -2147483648:
                                if (34 <= -5) {
                                    v44 := GenS2{F0: uint8(v16)}
                                    var v45 uint = 2
                                    if ((v17 != (1e10 + (v17 * -2.25))) && v24) {
                                        v46 := GenS5{F0: (^(v16)), F1: ("\n\t" + "a"), F2: v26}
                                        if (81 != (fi23 + (v2 & (v2 ^ (v14 * -1))))) {
                                            v47 := GenS2{F0: 156}
                                            _ = v47
                                        }
                                        var v48 string = ("a" + "\n\t")
                                        _ = v46
                                        _ = v48
                                    }
                                    _ = v44
                                    _ = v45
                                } else {
                                    v2 += -224
                                    var v49 int = (^(v2))
                                    _ = v49
                                }
                            }
                        case 9223372036854775807:
                            if v24 {
                                var v50 uint64 = v16
                                var v51 int = (v14 - v2)
                                _ = v50
                                _ = v51
                            }
                        default:
                            v2 = (v14 >> ((18446744073709551615) & 63))
                        }
                        v17 -= (-2.25 - (0.0 / v17))
                        var v52 string = (((("" + "hello") + (("foo bar" + ("" + "\n\t")) + "")) + "a") + ("foo bar" + ("a" + (("" + "") + ""))))
                        _ = v52
                    }
                    _ = v25
                    _ = v26
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        for fi53 := 0; fi53 < 1; fi53++ {
            if (953 == -32768) {
                v16 = 42
            }
            _ = fi53
        }
    } else {
        for fi54 := 0; fi54 < 19; fi54++ {
            for fi55 := 0; fi55 < 1; fi55++ {
                v56 := GenS6{F0: 1, F1: 1, F2: 2}
                switch "foo bar" {
                case "":
                    switch (fi54 | (v14 ^ v2)) {
                    case 128:
                        v4 = append(v4, 2147483647)
                    case 515:
                        var v57 int64 = -538
                        _ = v57
                    case 127:
                        var v58 uint = 2
                        _ = v58
                    }
                case "\n\t":
                    var v59 uint = 2
                    _ = v59
                default:
                    v60 := v3[958]
                    _ = v60
                }
                _ = fi55
                _ = v56
            }
            _ = fi54
        }
        fmt.Println(v10)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
