package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32; F1 uint16; F2 int }
func genFn0() uint {
    v4 := []int{42, 2147483647, 688}
    switch 840 {
    case 127:
        v5 := GenS1{F0: 127}
        _ = v5
    }
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = 32767
    var v3 int = 32768
    var v6 float64 = ((0.1 / -2.25) * (1e10 + 0.1))
    if (!(0 == 255)) {
        v7 := GenS2{F0: 1, F1: genFn0(), F2: 18446744073709551615}
        for fi8 := 0; fi8 < 1; fi8++ {
            if ((!(-2.25 < (float32(v6) / float32((float64(((((3.14159 * (-2.25 / (float32(v6) + (1.5 * 0.1)))) - 0.1) / float32((float64(1.5) + 0.0))) - 1e10)) / v6))))) && false) {
                if true {
                    if (!false) {
                        if (682 == 65535) {
                            var v9 float64 = float64(3.14159)
                            v9 = v6
                            _ = v9
                        }
                        if (2147483648 >= 4294967295) {
                            fmt.Println(v6, v3)
                        } else {
                            var v10 uint16 = 255
                            var v11 uint64 = 42
                            _ = v10
                            _ = v11
                        }
                        v1 -= 0
                    } else {
                        v1 = (^(v1))
                        v3 = v3
                        var v12 string = "a"
                        _ = v12
                    }
                    if true {
                        v13 := []int{v3, v3, v3, (v3 - fi8)}
                        _ = v13
                    }
                    var v14 uint32 = 439
                    _ = v14
                }
                var v15 int = 128
                _ = v15
            } else {
                for fi16 := 0; fi16 < 5; fi16++ {
                    v17 := []int{fi8, fi16, (fi16 - (v3 + (-(fi16))))}
                    var v18 uint = 65535
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            var v19 int32 = -1
            _ = fi8
            _ = v19
        }
        _ = v7
    } else {
        v6 = float64(3.14159)
    }
    if ((127 != 56) || true) {
        switch ("foo bar" + ("\n\t" + "a")) {
        case "foo bar":
            fmt.Println(v6)
        default:
            if ((false || true) || true) {
                switch ("foo bar" + "\n\t") {
                case "hello":
                    v1 += (v1 / ((v1) | 1))
                case "a":
                    v20 := GenS3{F0: 0, F1: true, F2: (1 < 255)}
                    _ = v20
                case "\n\t":
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v22 := []int{}
                        fmt.Println(v6, v6, v2)
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        _ = fi21
                        _ = v22
                    }
                }
                var v23 uint = genFn0()
                _ = v23
            }
        }
    }
    v24 := []int{v3, (v3 * (^(v3))), (v3 << ((2) & 63))}
    v1 = 42
    for fi25 := 0; fi25 < 18; fi25++ {
        for fi26 := 0; fi26 < 13; fi26++ {
            if ((42 < 1) || ("hello" < "hello")) {
                v1 -= v1
                v27 := GenS1{F0: -114}
                v3 -= -2147483648
                if true {
                    v1 -= 42
                    _ = len(v24)
                }
                _ = v27
            } else {
                v3 -= int(v2)
                var v28 uint64 = 4294967295
                if false {
                    v2 = (v2 ^ v2)
                    _ = len(v24)
                    if (1 <= 65535) {
                        if true {
                            v29 := GenS4{F0: "", F1: 65535}
                            _ = v29
                        } else {
                            if true {
                                v28 -= 214
                                var v30 float64 = v6
                                v1 = v1
                                _ = v30
                            }
                            var v31 uint32 = 2147483647
                            if (true && true) {
                                if ((v2 - (v2 + (^(v2)))) != 42) {
                                    switch (fi26 - (fi25 & -9223372036854775808)) {
                                    case 32767:
                                        v6 -= v6
                                    case 128:
                                        fmt.Println(v6, fi25)
                                    default:
                                        v32 := GenS1{F0: -95}
                                        _ = v32
                                    }
                                    v33 := GenS5{F0: 492}
                                    _ = v33
                                }
                            }
                            _ = v31
                        }
                        if len(v24) > 0 {
                            v24[0] = fi26
                        }
                    }
                    if false {
                        var v34 uint8 = 1
                        _ = v34
                    } else {
                        if (0 <= 0) {
                            v24 = append(v24, fi26)
                            v24 = append(v24, v3)
                            v2 -= v2
                            v35 := GenS6{F0: 18446744073709551615}
                            _ = v35
                        } else {
                            fmt.Println(v2, fi25)
                            var v36 float32 = 1.5
                            _ = v36
                        }
                    }
                }
                _ = v28
            }
            switch (^(v3)) {
            case 4294967295:
                switch ("" + "\n\t") {
                case "foo bar":
                    var v37 uint = 1
                    _ = v37
                default:
                    switch v3 {
                    case 42:
                        switch (v3 % ((fi25) | 1)) {
                        case -2147483648:
                            _ = len(v24)
                        case -342:
                            v38 := GenS3{F0: 4294967295, F1: false, F2: false}
                            _ = v38
                        default:
                            fmt.Println(v2, fi26, fi26)
                        }
                    case 9223372036854775807:
                        var v39 int = (fi25 - 1)
                        _ = v39
                    case 227:
                        var v40 int8 = 1
                        _ = v40
                    default:
                        v41 := GenS7{F0: -1}
                        if len(v24) > 0 {
                            v24[0] = fi26
                        }
                        _ = v41
                    }
                }
            case 42:
                var v42 uint = 0
                _ = v42
            case 946:
                var v43 string = "foo bar"
                _ = v43
            default:
                v44 := GenS1{F0: -128}
                _ = v44
            }
            var v45 uint16 = 2
            var v46 int64 = v1
            switch (fi25 - (v3 | -302)) {
            case 128:
                v47 := GenS7{F0: 42}
                _ = v47
            }
            _ = fi26
            _ = v45
            _ = v46
        }
        var v48 float32 = (0.0 + float32(v6))
        for fi49 := 0; fi49 < 15; fi49++ {
            var v50 uint64 = 169
            v1 -= v1
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            var v51 bool = false
            v1 += (v1 * (v1 % (((^(v1))) | 1)))
            _ = len(v24)
            v1 = (v1 * 127)
            v48 = (v48 * (3.14159 * (1e10 - (0.1 - (v48 + float32((v6 - 1.5)))))))
            v52 := GenS3{F0: 0, F1: true, F2: (v51 || v51)}
            _ = fi49
            _ = v50
            _ = v51
            _ = v52
        }
        _ = fi25
        _ = v48
    }
    v53 := GenS6{F0: 144}
    var v54 int8 = -1
    var v55 bool = true
    switch v3 {
    case -9223372036854775808:
        if len(v24) > 0 {
            v24[0] = (v3 >> ((18446744073709551615) & 63))
        }
    case 42:
        var v56 int8 = (^(v54))
        _ = v56
    default:
        v55 = false
    }
    _ = len(v24)
    _ = len(v24)
    var v57 bool = v55
    v24 = append(v24, v3)
    var v58 string = ""
    switch (("\n\t" + v58) + v58) {
    case "\n\t":
        v59 := GenS2{F0: 127, F1: 569, F2: 4294967295}
        _ = v59
    case "hello":
        for fi60 := 0; fi60 < 5; fi60++ {
            fmt.Println(v1, v3)
            for fi61 := 0; fi61 < 12; fi61++ {
                if (v2 == v2) {
                    var v62 uint16 = 42
                    v63 := GenS2{F0: 42, F1: genFn0(), F2: 4294967295}
                    if v55 {
                        var v64 int = fi60
                        v64 += (fi61 << ((255) & 63))
                        _ = v64
                    }
                    _ = v62
                    _ = v63
                }
                _ = fi61
            }
            v65 := GenS7{F0: 1}
            _ = fi60
            _ = v65
        }
    case "":
        if (!(-1 >= 90)) {
            var v66 uint16 = 2
            var v67 int64 = (^(v1))
            _ = v66
            _ = v67
        } else {
            var v68 int8 = (v54 + 0)
            v58 = v58
            _ = v68
        }
    }
    v69 := GenS8{F0: -2147483648, F1: 973, F2: (v3 ^ (v3 % ((-319) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v53
    _ = v54
    _ = v55
    _ = v57
    _ = v58
    _ = v6
    _ = v69
}
