package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint8; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 uint16; F1 bool }
type GenS11 struct { F0 uint }
func genFn0(p0 uint) uint {
    v7 := []int{127, 0, 42}
    if len(v7) > 0 {
        v7[0] = 128
    }
    if true {
        p0 = p0
        _ = len(v7)
    }
    p0 = p0
    p0 -= p0
    _ = p0
    _ = v7
    return 255
}

func main() {
    var v1 int = -2147483648
    var v2 string = "\n\t"
    v3 := make(map[int]string)
    var v4 float64 = -2.25
    v5 := make(map[uint32]uint32)
    v6 := []int{(v1 / (((v1 / ((303) | 1))) | 1))}
    v8 := GenS1{F0: 1, F1: 2147483647, F2: 42}
    var v9 int = 952
    v10 := GenS1{F0: 2, F1: 661, F2: genFn0(0)}
    if (3.14159 > ((1e10 * (0.0 - 0.1)) + 3.14159)) {
        _ = len(v3)
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            switch v2 {
            case "a":
                v9 = (v9 / (((v1 >> ((42) & 63))) | 1))
            case "":
                if false {
                    if (!true) {
                        var v12 uint32 = 428
                        fmt.Println(v2, v9)
                        for fi13 := 0; fi13 < 16; fi13++ {
                            if len(v6) > 0 {
                                v6[0] = 127
                            }
                            var v14 float64 = 0.0
                            _ = fi13
                            _ = v14
                        }
                        _ = v12
                    } else {
                        _ = len(v3)
                        switch v1 {
                        case 917:
                            var v15 int16 = 32767
                            _ = v15
                        case 9223372036854775807:
                            if true {
                                v16 := GenS1{F0: 403, F1: 942, F2: 255}
                                for fi17 := 0; fi17 < 5; fi17++ {
                                    var v18 uint8 = 69
                                    _ = fi17
                                    _ = v18
                                }
                                _ = v16
                            }
                        case 4294967295:
                            var v19 bool = true
                            _ = v19
                        }
                        v20 := GenS2{F0: 1}
                        _ = v20
                    }
                }
            }
            _ = fi11
        }
        v21 := GenS3{F0: (v1 & v9), F1: 1, F2: ("foo bar" + v2)}
        for fi22 := 0; fi22 < 12; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                v24 := GenS4{F0: 18446744073709551615}
                v25 := GenS4{F0: 4294967295}
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
        _ = v21
    }
    fmt.Println(v4, v2)
    v26 := GenS5{F0: 100}
    for fi27 := 0; fi27 < 3; fi27++ {
        for fi28 := 0; fi28 < 3; fi28++ {
            switch v2 {
            case "hello":
                v4 -= v4
            case "foo bar":
                switch v2 {
                case "foo bar":
                    v1 = fi28
                default:
                    if true {
                        var v29 float64 = v4
                        v29 -= (0.0 / (1e10 / (v4 + v29)))
                        v30 := GenS2{F0: -243}
                        var v31 bool = false
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        if (2 > 0) {
                            if (true || false) {
                                v32 := GenS4{F0: 446}
                                _ = v32
                            }
                        }
                        v6 = append(v6, (fi28 | -248))
                        var v33 int64 = 9223372036854775807
                        _ = v33
                    }
                }
            default:
                if false {
                    fmt.Println(fi27, v9)
                    v34 := GenS1{F0: 0, F1: 2, F2: 0}
                    _ = v34
                } else {
                    if false {
                        delete(v5, 4294967295)
                        if false {
                            fmt.Println(v2, v9)
                            if (4294967295 != 4294967295) {
                                v2 = (v2 + v2)
                                var v35 int64 = 1
                                _ = v35
                            }
                            v1 -= 894
                        } else {
                            v9 = -2147483648
                            if (4294967295 > 65535) {
                                v36 := GenS5{F0: -1}
                                _ = v36
                            }
                            v37 := GenS4{F0: 0}
                            _ = v37
                        }
                        v9 = (fi28 << ((65535) & 63))
                        if (!(42 != 1)) {
                            switch (v9 >> ((2) & 63)) {
                            case -9223372036854775808:
                                v38 := GenS6{F0: false}
                                _ = v38
                            }
                            if false {
                                var v39 uint8 = 255
                                switch ("" + "\n\t") {
                                case "hello":
                                    v4 = v4
                                }
                                v3[v1] = ("a" + (("hello" + "hello") + v2))
                                v1 -= 32767
                                v39 += v39
                                _ = v39
                            }
                        }
                    }
                    if (true && (((genFn0(1) < 4294967295) || true) && (65535 != 2))) {
                        var v40 float32 = (0.0 - float32(3.14159))
                        var v41 int32 = 2147483647
                        if (0 != 4294967295) {
                            switch fi28 {
                            case 982:
                                v42 := GenS3{F0: (v9 & fi27), F1: 75, F2: "hello"}
                                _ = v42
                            default:
                                v40 = 3.14159
                            }
                        } else {
                            v9 = int(v41)
                        }
                        _ = v40
                        _ = v41
                    }
                    var v43 bool = true
                    _ = v43
                }
            }
            if (!true) {
                v44 := GenS4{F0: 4294967295}
                var v45 int16 = 910
                v6 = append(v6, (v9 << ((115) & 63)))
                _ = v44
                _ = v45
            }
            v46 := GenS7{F0: 18446744073709551615}
            if true {
                v1 -= (fi28 % ((-1) | 1))
                v4 -= v4
                v1 = fi27
                var v47 float64 = v4
                _ = v47
            } else {
                v4 = 1e10
                if true {
                    v48 := GenS8{F0: "", F1: -804}
                    _ = len(v6)
                    _ = len(v3)
                    if true {
                        switch v1 {
                        case 1:
                            v2 = "foo bar"
                        case 128:
                            _ = len(v3)
                        }
                        v6 = append(v6, (-(v9)))
                    }
                    var v49 int8 = -128
                    switch (fi28 & v1) {
                    case 1:
                        v49 += v49
                    case 2147483647:
                        if false {
                            fmt.Println(fi28, v49)
                            _ = len(v5)
                            fmt.Println(v2)
                            delete(v5, 2147483647)
                            fmt.Println(v4, fi28, v1)
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    case -1:
                        v50 := GenS9{F0: v49, F1: false}
                        _ = v50
                    }
                    _ = v48
                    _ = v49
                } else {
                    v2 = (v2 + "")
                }
                if false {
                    if false {
                        v51 := GenS3{F0: (fi27 >> ((255) & 63)), F1: 42, F2: v2}
                        fmt.Println(v9)
                        v4 = v4
                        _ = v51
                    } else {
                        var v52 uint = 4294967295
                        var v53 float64 = v4
                        if (v4 > 3.14159) {
                            v4 -= 0.0
                        } else {
                            v54 := GenS10{F0: 42, F1: (272 != 32767)}
                            _ = v54
                        }
                        var v55 uint16 = 1
                        _ = v52
                        _ = v53
                        _ = v55
                    }
                    var v56 float64 = ((((v4 / ((v4 + (v4 / float64(1e10))) / 1.5)) + v4) + (((((3.14159 / (((((float64(-2.25) - ((v4 + -2.25) / v4)) * v4) - v4) / ((v4 + v4) - (1e10 / (1e10 - v4)))) - (v4 * ((v4 - (v4 - v4)) * 0.1)))) / ((v4 + (v4 * (v4 / 1.5))) * v4)) - float64((0.0 * float32(v4)))) + v4) + v4)) + v4)
                    _ = v56
                }
            }
            _ = fi28
            _ = v46
        }
        _ = fi27
    }
    _ = len(v6)
    if (!(!(-1 >= 1))) {
        switch "foo bar" {
        case "\n\t":
            var v57 bool = true
            _ = v57
        case "hello":
            fmt.Println(v4, v9)
        case "foo bar":
            var v58 uint32 = 21
            _ = v58
        }
        if false {
            switch v2 {
            case "":
                for fi59 := 0; fi59 < 14; fi59++ {
                    v4 = float64(0.1)
                    var v60 uint = 65535
                    _ = fi59
                    _ = v60
                }
            case "foo bar":
                var v61 uint32 = 0
                _ = v61
            case "hello":
                v6 = append(v6, v9)
            }
        } else {
            v62 := GenS10{F0: 0, F1: false}
            v1 = v9
            v63 := GenS11{F0: 18446744073709551615}
            _ = v62
            _ = v63
        }
    }
    v5[2147483647] = 258
    v9 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
