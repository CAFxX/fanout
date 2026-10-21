package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 uint64 }

func main() {
    var v1 uint64 = 569
    v2 := make(map[string]string)
    var v3 uint32 = 193
    v4 := []string{"a", ("" + "foo bar"), ("\n\t" + "")}
    v5 := make(map[uint32]string)
    switch (("" + "\n\t") + "\n\t") {
    case "hello":
        v1 = v1
    case "":
        var v6 float64 = 0.1
        _ = v6
    case "a":
        v3 = 65535
    default:
        v1 -= (v1 >> ((42) & 63))
    }
    v1 += (v1 ^ v1)
    fmt.Println(v3, v1)
    var v7 uint = 4294967295
    var v8 float64 = float64(-2.25)
    v3 += v3
    var v9 uint64 = v1
    for fi10 := 0; fi10 < 8; fi10++ {
        v3 = (v3 / ((v3) | 1))
        v7 = v7
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: 0}
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        switch 155 {
        case 0:
            v14 := GenS1{F0: 9223372036854775807}
            _ = v14
        case 870:
            var v15 uint16 = 198
            _ = v15
        case -9223372036854775808:
            if (42 == 617) {
                if (!(9223372036854775807 < 1)) {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        v1 += v9
                        v17 := GenS1{F0: 0}
                        if true {
                            v4 = append(v4, "a")
                            switch "\n\t" {
                            case "hello":
                                var v18 uint = (v7 - 937)
                                _ = v18
                            case "a":
                                v9 += 2
                                v7 = 65535
                                v19 := GenS1{F0: -368}
                                _ = v19
                            case "foo bar":
                                v8 = 1e10
                            default:
                                v20 := v2["\n\t"]
                                _ = v20
                            }
                            if (!((!((v3 / ((v3) | 1)) != (^(v3)))) && true)) {
                                v7 -= v7
                                if false {
                                    switch 9223372036854775807 {
                                    case 9223372036854775807:
                                        var v21 int16 = -377
                                        _ = v21
                                    case 127:
                                        v22 := GenS1{F0: -1}
                                        _ = v22
                                    }
                                    switch "" {
                                    case "a":
                                        switch "foo bar" {
                                        case "":
                                            v8 -= float64(0.1)
                                        default:
                                            fmt.Println(v3)
                                        }
                                    case "":
                                        fmt.Println(v9, v7)
                                    case "\n\t":
                                        switch ("" + ("\n\t" + "hello")) {
                                        case "a":
                                            v8 = ((v8 / v8) / (v8 - v8))
                                        case "\n\t":
                                            fmt.Println(fi16)
                                            v23 := GenS1{F0: -1}
                                            v3 = (v3 | uint32(v1))
                                            v1 += (^(v9))
                                            v24 := GenS1{F0: 736}
                                            _ = v23
                                            _ = v24
                                        default:
                                            var v25 int8 = 0
                                            _ = v25
                                        }
                                    }
                                    v9 += 0
                                }
                                var v26 int16 = -1
                                _ = v26
                            } else {
                                switch fi16 {
                                case 42:
                                    switch ("\n\t" + ("a" + ("hello" + "a"))) {
                                    case "":
                                        v27 := GenS1{F0: 376}
                                        _ = v27
                                    }
                                case 127:
                                    _ = len(v5)
                                default:
                                    v28 := GenS1{F0: 1}
                                    _ = v28
                                }
                                var v29 int = 386
                                v30 := GenS2{F0: uint16(v3), F1: v9}
                                v31 := GenS1{F0: 42}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        }
                        var v32 bool = true
                        if true {
                            v9 = 42
                            var v33 bool = true
                            _ = v33
                        } else {
                            v4 = append(v4, "\n\t")
                            v34 := v2[((("" + ("a" + ("" + "hello"))) + "") + (("hello" + (("\n\t" + "hello") + ("" + ("\n\t" + "a")))) + ("foo bar" + "")))]
                            var v35 uint = v7
                            _ = v34
                            _ = v35
                        }
                        v36 := v5[v3]
                        v37 := GenS3{F0: (v1 << ((18446744073709551615) & 63))}
                        var v38 uint = v7
                        var v39 int = fi16
                        _ = len(v4)
                        _ = fi16
                        _ = v17
                        _ = v32
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 2; fi40++ {
                        v41 := GenS1{F0: -587}
                        if false {
                            v9 = v9
                            var v42 string = ("a" + ("foo bar" + "a"))
                            var v43 uint16 = 65535
                            _ = v42
                            _ = v43
                        } else {
                            if (false && false) {
                                v8 -= float64((1.5 / (float32((v8 + -2.25)) - (((0.0 + (0.1 / 1.0)) - (1.5 - 0.1)) * (-2.25 / (0.1 / 0.1))))))
                                v9 = 884
                                fmt.Println(v3)
                            }
                            if false {
                                switch ("\n\t" + (("hello" + "foo bar") + "foo bar")) {
                                case "":
                                    v9 += v9
                                case "\n\t":
                                    v1 -= (v9 + (v1 ^ uint64(v3)))
                                default:
                                    v3 -= (^(v3))
                                }
                                v44 := GenS3{F0: v9}
                                fmt.Println(v3)
                                _ = v44
                            } else {
                                var v45 bool = false
                                _ = v45
                            }
                        }
                        fmt.Println(v3, v9)
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(fi40, v9, v1)
                        case "hello":
                            fmt.Println(v8)
                        case "foo bar":
                            var v46 uint8 = 127
                            _ = v46
                        }
                        _ = fi40
                        _ = v41
                    }
                    for fi47 := 0; fi47 < 15; fi47++ {
                        v48 := GenS3{F0: v1}
                        var v49 uint8 = 127
                        var v50 string = "hello"
                        v51 := GenS4{F0: -93}
                        fmt.Println(v3, v49)
                        switch 1 {
                        case -311:
                            var v52 string = (("a" + v50) + v50)
                            v52 = (v50 + ((("hello" + v50) + ((v52 + (((v52 + "") + v52) + (v52 + (v52 + "")))) + v50)) + v52))
                            _ = v52
                        case -1:
                            var v53 int32 = -1
                            _ = v53
                        case -401:
                            var v54 float32 = (-2.25 + (0.0 / (1.5 + 3.14159)))
                            _ = v54
                        }
                        _ = fi47
                        _ = v48
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                } else {
                    for fi55 := 0; fi55 < 20; fi55++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch "" {
                        case "\n\t":
                            if (196 != 78) {
                                v4 = append(v4, "a")
                                _ = len(v5)
                                v1 = 183
                                if true {
                                    switch (fi13 / ((fi55) | 1)) {
                                    case 2147483647:
                                        var v56 uint16 = 690
                                        _ = v56
                                    default:
                                        if true {
                                            var v57 int32 = 42
                                            if (false && false) {
                                                if (922 < 42) {
                                                    var v58 uint8 = 36
                                                    _ = v58
                                                } else {
                                                    v59 := GenS2{F0: 275, F1: (^(v1))}
                                                    _ = v59
                                                }
                                                v1 = 0
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v57
                                        }
                                    }
                                    v9 -= v1
                                    switch "" {
                                    case "hello":
                                        v8 = v8
                                    case "foo bar":
                                        if len(v4) > 0 {
                                            v4[0] = ("\n\t" + "")
                                        }
                                    case "a":
                                        if (!(!true)) {
                                            v7 += (v7 >> (((v7 | (v7 * v7))) & 63))
                                            fmt.Println(fi13, v9)
                                            var v60 string = ""
                                            if ((0 <= -1) || true) {
                                                v7 -= (v7 / ((42) | 1))
                                                v60 = v60
                                            }
                                            var v61 bool = false
                                            _ = v60
                                            _ = v61
                                        }
                                    }
                                    v62 := GenS4{F0: 1}
                                    _ = v62
                                }
                                var v63 uint16 = 995
                                _ = v63
                            }
                        case "":
                            var v64 uint32 = 0
                            _ = v64
                        }
                        _ = fi55
                    }
                    v1 -= v9
                    var v65 uint = 838
                    _ = v65
                }
            } else {
                var v66 float64 = (v8 / v8)
                v5[v3] = "a"
                v66 = (float64(-2.25) * 3.14159)
                _ = v66
            }
        default:
            if (((1e10 / 1e10) * float32(((v8 + v8) - ((((0.0 - v8) / -2.25) - ((((v8 * 0.1) * v8) / (float64(1e10) + (v8 * (0.1 - 0.1)))) + v8)) - (v8 + float64((1.5 / float32(v8)))))))) != (0.0 * 0.0)) {
                v9 = v1
                fmt.Println(fi13, fi13, v7)
                var v67 uint = (v7 ^ (v7 | (^(v7))))
                _ = v67
            } else {
                _ = len(v2)
            }
        }
        _ = fi13
    }
    if (9223372036854775807 > -632) {
        v8 = v8
        v3 -= v3
    } else {
        var v68 int16 = 1
        var v69 uint32 = (v3 >> (((v7 >> ((v7) & 63))) & 63))
        fmt.Println(v3, v9)
        _ = v68
        _ = v69
    }
    v2[(("a" + (("hello" + (("" + "") + (("\n\t" + "") + ("a" + "a")))) + "a")) + "foo bar")] = "hello"
    var v70 uint16 = uint16(v1)
    v8 -= v8
    v71 := GenS5{F0: 42, F1: 4294967295}
    fmt.Println(v1, v70)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v70
    _ = v71
    _ = v8
    _ = v9
}
