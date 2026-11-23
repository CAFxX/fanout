package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 bool; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 int16; F2 uint8 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int; F1 string }

func main() {
    var v1 uint32 = 1
    var v2 int32 = -731
    var v3 float32 = (0.0 * ((3.14159 + (-2.25 / 3.14159)) * (0.1 * 1.5)))
    var v4 bool = true
    var v5 int16 = -555
    if (v4 && v4) {
        v6 := []int{87, 715, 1, 2147483647}
        v6 = append(v6, 9223372036854775807)
        _ = v6
    } else {
        switch "foo bar" {
        case "":
            switch "" {
            case "hello":
                v3 = 0.0
            default:
                v3 += v3
            }
        default:
            if v4 {
                fmt.Println(v5)
                v7 := []int{}
                switch 2147483647 {
                case 2147483647:
                    var v8 uint = 196
                    _ = v8
                case 128:
                    if v4 {
                        for fi9 := 0; fi9 < 18; fi9++ {
                            if ((((42 >= 42) || (851 > uint(v1))) || (v4 || (!(1 > 912)))) && (!v4)) {
                                var v10 uint = 18446744073709551615
                                fmt.Println(v2, v10, v4)
                                _ = v10
                            }
                            _ = fi9
                        }
                    }
                case 32768:
                    v11 := GenS1{F0: ("\n\t" + "a"), F1: 9223372036854775807, F2: -9223372036854775808}
                    _ = v11
                }
                _ = v7
            }
        }
    }
    if v4 {
        v12 := []int{483}
        if (!false) {
            if (!(!false)) {
                _ = len(v12)
                var v13 uint16 = 244
                v14 := GenS1{F0: "", F1: 1, F2: -9223372036854775808}
                _ = v13
                _ = v14
            } else {
                v15 := GenS2{F0: 65535, F1: 32767}
                if (int64(v1) == 1) {
                    if len(v12) > 0 {
                        v12[0] = 323
                    }
                }
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v15
            }
            fmt.Println(v4, v3, v5)
            for fi16 := 0; fi16 < 15; fi16++ {
                fmt.Println(v3, v4)
                _ = fi16
            }
        } else {
            v17 := GenS1{F0: ("hello" + "foo bar"), F1: 42, F2: 0}
            for fi18 := 0; fi18 < 3; fi18++ {
                v12 = append(v12, fi18)
                for fi19 := 0; fi19 < 3; fi19++ {
                    _ = len(v12)
                    v3 += v3
                    _ = fi19
                }
                _ = fi18
            }
            _ = v17
        }
        var v20 uint8 = 128
        _ = v12
        _ = v20
    }
    if (887 >= 255) {
        switch -894 {
        case -851:
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 uint16 = 2
                v1 = v1
                v4 = ((true && ("a" > "hello")) && (203 < 127))
                _ = fi21
                _ = v22
            }
        case -9223372036854775808:
            switch "hello" {
            case "hello":
                switch -91 {
                case 284:
                    for fi23 := 0; fi23 < 8; fi23++ {
                        var v24 uint = 255
                        var v25 int32 = v2
                        for fi26 := 0; fi26 < 6; fi26++ {
                            v25 -= (v2 / (((^(v25))) | 1))
                            if ((true || v4) && (!(false || (v5 <= -1)))) {
                                if (((!v4) || (((v4 && v4) || (!v4)) && (float32((1.5 - 3.14159)) != (3.14159 * v3)))) && ((v3 + (float32(0.1) - -2.25)) <= ((float32(1e10) - 1.5) / v3))) {
                                    v27 := []int{fi26, 127}
                                    v2 = 908
                                    _ = v27
                                } else {
                                    v3 = v3
                                    if (v4 && true) {
                                        v1 = v1
                                        var v28 float64 = 1.5
                                        v25 = v2
                                        _ = v28
                                    }
                                    var v29 int8 = int8(fi23)
                                    _ = v29
                                }
                            } else {
                                if (false && true) {
                                    if true {
                                        v5 = v5
                                        var v30 int = fi23
                                        _ = v30
                                    } else {
                                        var v31 int64 = int64(v1)
                                        _ = v31
                                    }
                                    v32 := []int{(fi23 % ((32767) | 1)), fi23, 42, -9223372036854775808}
                                    v32 = append(v32, (fi23 >> ((0) & 63)))
                                    _ = v32
                                }
                            }
                            v33 := []int{fi23, (^(fi26)), fi26}
                            v34 := GenS3{F0: (42 < 100), F1: (v2 * int32(v1)), F2: 255}
                            _ = fi26
                            _ = v33
                            _ = v34
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case 32768:
                    if v4 {
                        v2 = 596
                        var v35 int8 = 127
                        var v36 uint16 = 738
                        _ = v35
                        _ = v36
                    } else {
                        fmt.Println(v5, v2)
                        fmt.Println(v4, v5, v2)
                    }
                }
            case "foo bar":
                var v37 uint = 2
                _ = v37
            case "a":
                fmt.Println(v3, v2, v4)
            }
        case -458:
            if (18446744073709551615 == 0) {
                if false {
                    v38 := []int{}
                    if len(v38) > 0 {
                        v38[0] = 2147483647
                    }
                    _ = v38
                }
            }
        default:
            v5 -= int16(v1)
        }
        for fi39 := 0; fi39 < 11; fi39++ {
            v40 := GenS4{F0: 320}
            var v41 uint16 = uint16(v1)
            v2 = (v2 << ((42) & 63))
            _ = fi39
            _ = v40
            _ = v41
        }
        var v42 uint64 = 403
        _ = v42
    } else {
        fmt.Println(v3, v4)
    }
    switch 32768 {
    case 127:
        v43 := GenS5{F0: true, F1: v5, F2: 127}
        _ = v43
    case 32768:
        v44 := []int{}
        _ = v44
    case 2147483647:
        var v45 int16 = v5
        _ = v45
    default:
        v4 = v4
    }
    v46 := GenS6{F0: (!true), F1: 255}
    var v47 float32 = v3
    switch 0 {
    case 42:
        fmt.Println(v47)
    case 32768:
        v3 += v3
    case 4294967295:
        var v48 uint = 18446744073709551615
        _ = v48
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        v50 := []int{}
        _ = fi49
        _ = v50
    }
    v3 += ((v3 / -2.25) + v47)
    var v51 uint16 = 42
    var v52 uint64 = 1
    v4 = (!(0.1 > 1e10))
    v51 += (v51 | (v51 >> ((647) & 63)))
    v53 := []int{-9223372036854775808, -525, 80, 32767}
    v4 = v4
    v54 := GenS7{F0: int(v51), F1: ("" + "hello")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v53:", len(v53), v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
