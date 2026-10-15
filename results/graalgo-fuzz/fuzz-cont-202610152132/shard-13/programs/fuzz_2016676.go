package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint8; F1 bool; F2 string }
type GenS3 struct { F0 int64; F1 int8; F2 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int }
func genFn0() int {
    var v6 float32 = 0.0
    var v7 int8 = -1
    for fi8 := 0; fi8 < 10; fi8++ {
        v6 -= 3.14159
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: 42, F1: false}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint32, p1 uint, p2 int32) int32 {
    fmt.Println(p1)
    if (!false) {
        if true {
            v11 := []int{411, genFn0()}
            var v12 int8 = -1
            v13 := GenS1{F0: 1, F1: true}
            _ = v11
            _ = v12
            _ = v13
        }
    }
    v14 := GenS2{F0: 0, F1: false, F2: "\n\t"}
    v15 := []int{156, 2147483647, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 uint = 255
    v4 := []string{"a", "hello"}
    v5 := []uint32{}
    v3 = (v3 / ((18446744073709551615) | 1))
    var v16 int32 = 579
    switch 1 {
    case -529:
        fmt.Println(v16, v3)
    case 0:
        if (true && true) {
            switch -2147483648 {
            case 128:
                switch "\n\t" {
                case "hello":
                    var v17 uint8 = 255
                    _ = v17
                case "foo bar":
                    if false {
                        if true {
                            fmt.Println(v3)
                            if true {
                                var v18 int64 = 1
                                for fi19 := 0; fi19 < 13; fi19++ {
                                    for fi20 := 0; fi20 < 6; fi20++ {
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(fi19)
                                        case "\n\t":
                                            v21 := GenS2{F0: 42, F1: (("hello" >= "\n\t") && (!(!((fi19 - fi19) < 0)))), F2: "hello"}
                                            var v22 uint16 = 255
                                            var v23 uint16 = v22
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        default:
                                            v16 -= (v16 - (v16 * v16))
                                        }
                                        _ = fi20
                                    }
                                    _ = fi19
                                }
                                v3 += v3
                                _ = v18
                            } else {
                                var v24 int = 127
                                _ = v24
                            }
                            var v25 float32 = 3.14159
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 7; fi26++ {
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                                v27 := GenS3{F0: 1, F1: 1, F2: (v16 - -2147483648)}
                                _ = fi26
                                _ = v27
                            }
                            v28 := GenS4{F0: 57}
                            _ = v28
                        }
                        var v29 int16 = -32768
                        _ = v29
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            for fi30 := 0; fi30 < 3; fi30++ {
                                v1[2] = (float32(-2.25) == 0.1)
                                switch "\n\t" {
                                case "\n\t":
                                    delete(v2, (("foo bar" + "") + ("foo bar" + (("\n\t" + (("\n\t" + "") + "foo bar")) + ""))))
                                default:
                                    v16 = v16
                                }
                                _ = fi30
                            }
                        default:
                            switch 2147483647 {
                            case 32767:
                                v3 = (^(v3))
                            default:
                                fmt.Println(v16, v16)
                            }
                        }
                        if false {
                            v3 = 2
                            var v31 int8 = 0
                            v4 = append(v4, "a")
                            _ = v31
                        } else {
                            if true {
                                v16 = genFn1(2147483647, v3, v16)
                                v16 = genFn1(2147483647, v3, 2147483647)
                                var v32 uint16 = 900
                                _ = v32
                            }
                        }
                    }
                default:
                    v33 := v2[""]
                    _ = v33
                }
            default:
                if len(v5) > 0 {
                    v5[0] = 1
                }
            }
            v34 := GenS2{F0: 127, F1: true, F2: "hello"}
            _ = v34
        }
    case 1:
        v3 += v3
    default:
        for fi35 := 0; fi35 < 18; fi35++ {
            v36 := GenS5{F0: 1, F1: 549}
            _ = fi35
            _ = v36
        }
    }
    if true {
        for fi37 := 0; fi37 < 7; fi37++ {
            v4 = append(v4, ("hello" + ""))
            if ("" < ((("" + "a") + "a") + "\n\t")) {
                var v38 int16 = 32767
                _ = v38
            }
            _ = fi37
        }
        v16 = (v16 + genFn1(1, (v3 % (((v3 | 18446744073709551615)) | 1)), -2147483648))
        if (4294967295 == 606) {
            var v39 int = -9223372036854775808
            if len(v4) > 0 {
                v4[0] = ""
            }
            v40 := v2["hello"]
            _ = v39
            _ = v40
        } else {
            if len(v4) > 0 {
                v4[0] = ""
            }
            var v41 int64 = 9223372036854775807
            _ = v41
        }
    }
    var v42 int16 = 0
    for fi43 := 0; fi43 < 12; fi43++ {
        if ((v16 & (v16 % ((v16) | 1))) >= (v16 + -2147483648)) {
            var v44 uint16 = 255
            for fi45 := 0; fi45 < 19; fi45++ {
                if (false && true) {
                    if false {
                        var v46 int32 = (v16 * -197)
                        _ = v46
                    } else {
                        v44 = (v44 | (v44 / (((^(v44))) | 1)))
                    }
                    v47 := GenS3{F0: -9223372036854775808, F1: 100, F2: (v16 * 943)}
                    var v48 int32 = 279
                    _ = v47
                    _ = v48
                }
                var v49 uint64 = 42
                switch "a" {
                case "a":
                    fmt.Println(v3, v42, fi45)
                case "":
                    var v50 int8 = -1
                    _ = v50
                case "\n\t":
                    var v51 int16 = (v42 << ((v3) & 63))
                    _ = v51
                }
                _ = fi45
                _ = v49
            }
            v44 = 287
            _ = v44
        } else {
            switch (fi43 / (((fi43 & (fi43 + (fi43 >> (((v3 << (((v3 - v3)) & 63))) & 63))))) | 1)) {
            case -899:
                if ((0 > (fi43 * 42)) || ((v42 | -441) <= (^(v42)))) {
                    var v52 uint32 = 0
                    for fi53 := 0; fi53 < 14; fi53++ {
                        v42 = (v42 >> ((v3) & 63))
                        _ = fi53
                    }
                    _ = v52
                } else {
                    var v54 int32 = v16
                    _ = v54
                }
            }
            v55 := GenS2{F0: 127, F1: (((fi43 != genFn0()) && (42 >= -9223372036854775808)) || (!(!true))), F2: ("\n\t" + ("a" + "hello"))}
            if (1 != 2) {
                v42 = (-(v42))
            }
            _ = v55
        }
        switch (fi43 * genFn0()) {
        case -851:
            for fi56 := 0; fi56 < 6; fi56++ {
                var v57 string = "foo bar"
                var v58 uint = v3
                var v59 uint64 = 42
                _ = fi56
                _ = v57
                _ = v58
                _ = v59
            }
        }
        v3 += (v3 >> (((v3 / (((v3 >> ((0) & 63))) | 1))) & 63))
        _ = fi43
    }
    fmt.Println(v42, v16)
    for fi60 := 0; fi60 < 8; fi60++ {
        var v61 float64 = (-2.25 / (1e10 - (0.0 + 3.14159)))
        v61 -= 0.0
        v16 -= (v16 + (v16 >> ((18446744073709551615) & 63)))
        _ = fi60
        _ = v61
    }
    switch ("hello" + "\n\t") {
    case "\n\t":
        var v62 uint = (^(v3))
        _ = v62
    case "foo bar":
        var v63 int32 = (v16 & -1)
        _ = v63
    case "hello":
        v16 = (v16 % (((v16 | (v16 & (v16 ^ (^(v16)))))) | 1))
    default:
        switch -617 {
        case 9223372036854775807:
            if len(v4) > 0 {
                v4[0] = ""
            }
        case 128:
            for fi64 := 0; fi64 < 9; fi64++ {
                v65 := GenS6{F0: (fi64 % ((-537) | 1))}
                var v66 uint16 = 42
                for fi67 := 0; fi67 < 4; fi67++ {
                    v16 += genFn1(390, 180, v16)
                    var v68 int16 = 42
                    _ = fi67
                    _ = v68
                }
                _ = fi64
                _ = v65
                _ = v66
            }
        case 127:
            fmt.Println(v42, v16)
        default:
            v16 += v16
        }
    }
    if true {
        var v69 int64 = 0
        v70 := v2["\n\t"]
        _ = v69
        _ = v70
    } else {
        v42 += (v42 % ((-1) | 1))
        var v71 uint64 = 18446744073709551615
        v42 = (v42 ^ (v42 >> (((v3 << (((v3 | 18446744073709551615)) & 63))) & 63)))
        _ = v71
    }
    var v72 uint64 = 676
    v73 := GenS3{F0: 1, F1: -1, F2: (v16 ^ (v16 - 42))}
    fmt.Println(v3, v16, v16)
    var v74 int8 = 1
    var v75 float32 = 0.1
    switch int(v74) {
    case -136:
        v75 = (-2.25 + v75)
    case -326:
        v74 += 33
    }
    var v76 uint8 = 42
    fmt.Println(v16, v76)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
