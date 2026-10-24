package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int64; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 uint8; F2 uint }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 bool; F1 int32; F2 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32; F1 int8 }
type GenS10 struct { F0 int; F1 string }
type GenS11 struct { F0 string }
func genFn0(p0 uint16, p1 int64) uint {
    fmt.Println(p1, p1)
    var v5 bool = (41 <= 4294967295)
    if (!(42 >= 565)) {
        var v6 uint32 = 513
        if (-961 > -9223372036854775808) {
            var v7 float32 = -2.25
            if true {
                var v8 uint = 760
                _ = v8
            }
            v9 := GenS1{F0: 18446744073709551615, F1: 453}
            _ = v7
            _ = v9
        }
        _ = v6
    }
    var v10 uint16 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return 844
}
func genFn1() int16 {
    if (2147483647 < 2) {
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := []int{fi11}
            _ = fi11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint = 18446744073709551615
            v14 = (v14 % (((^(v14))) | 1))
            _ = fi13
            _ = v14
        }
        var v15 uint = 255
        v15 = (v15 - genFn0(716, 1))
        _ = v15
    }
    if (!(-2.25 != -2.25)) {
        switch "a" {
        case "foo bar":
            for fi16 := 0; fi16 < 13; fi16++ {
                fmt.Println(fi16)
                var v17 uint32 = 98
                _ = fi16
                _ = v17
            }
        case "":
            var v18 bool = (!(0 != -335))
            _ = v18
        }
        var v19 float64 = 0.0
        _ = v19
    }
    if (42 <= 782) {
        var v20 uint16 = 255
        v20 = 739
        v21 := []int{-543, 32767}
        _ = v20
        _ = v21
    } else {
        var v22 float32 = 0.0
        _ = v22
    }
    return 1000
}

func main() {
    var v1 string = "a"
    var v2 uint = 42
    var v3 bool = false
    v4 := []string{}
    v4 = append(v4, v1)
    for fi23 := 0; fi23 < 10; fi23++ {
        switch (fi23 >> (((^(v2))) & 63)) {
        case 881:
            if (v3 || v3) {
                v24 := GenS1{F0: 578, F1: 834}
                if v3 {
                    fmt.Println(v1, fi23)
                    v25 := GenS2{F0: -599, F1: -258, F2: 0}
                    if (-2147483648 == -859) {
                        var v26 int64 = -1
                        if ((v3 || (v3 || v3)) || (42 > 18446744073709551615)) {
                            v27 := GenS3{F0: 42}
                            for fi28 := 0; fi28 < 3; fi28++ {
                                v29 := GenS4{F0: (^(fi28))}
                                var v30 float64 = ((((0.0 / (((3.14159 - ((float64(1e10) + 0.0) * (float64(((3.14159 + float32(0.1)) / (3.14159 / 1.5))) * (0.1 / 0.1)))) - (0.0 + float64((((float32(((((1e10 * ((3.14159 * float64(1e10)) - 1e10)) + 1.5) - ((float64(3.14159) / 1.0) - 0.1)) / float64(float32(0.1)))) * 3.14159) * (1.5 - float32(((float64(0.0) + 1.5) / 1.5)))) + float32(0.0))))) + (-2.25 + 3.14159))) * 0.1) / 1.5) * 0.1)
                                v31 := GenS2{F0: v26, F1: 42, F2: 1}
                                _ = fi28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            if (0.1 == 1e10) {
                                v32 := GenS1{F0: 4294967295, F1: 285}
                                v3 = (fi23 >= fi23)
                                v1 = v1
                                _ = v32
                            } else {
                                var v33 uint8 = 107
                                var v34 bool = v3
                                _ = v33
                                _ = v34
                            }
                            _ = v27
                        } else {
                            var v35 int32 = -2147483648
                            v4 = append(v4, (v1 + v1))
                            v36 := GenS5{F0: (v2 + v2), F1: 2, F2: 1}
                            _ = v35
                            _ = v36
                        }
                        _ = v26
                    }
                    _ = v25
                }
                switch "hello" {
                case "hello":
                    v4 = append(v4, v1)
                default:
                    var v37 int32 = -2147483648
                    _ = v37
                }
                _ = v24
            }
        }
        var v38 uint = v2
        _ = fi23
        _ = v38
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        if len(v4) > 0 {
            v4[0] = v1
        }
        _ = fi39
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, v1)
    if (2147483647 != 9223372036854775807) {
        v40 := GenS6{F0: 30, F1: 0}
        _ = v40
    } else {
        v2 = v2
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v41 int = 580
        _ = v41
    }
    var v42 uint8 = 149
    if v3 {
        if (42 == 0) {
            if (42 > 0) {
                if len(v4) > 0 {
                    v4[0] = (v1 + v1)
                }
                _ = len(v4)
                v4 = append(v4, v1)
            } else {
                var v43 uint = 2
                v1 = ((("\n\t" + ((("hello" + v1) + v1) + v1)) + (v1 + (v1 + ""))) + v1)
                _ = v43
            }
            if v3 {
                for fi44 := 0; fi44 < 9; fi44++ {
                    switch v1 {
                    case "foo bar":
                        v2 -= 42
                    case "hello":
                        v1 = (v1 + v1)
                    }
                    _ = fi44
                }
                fmt.Println(v1, v1)
            } else {
                if v3 {
                    if v3 {
                        switch 693 {
                        case 32767:
                            fmt.Println(v1)
                        case 62:
                            var v45 uint = v2
                            _ = v45
                        default:
                            if (!v3) {
                                for fi46 := 0; fi46 < 4; fi46++ {
                                    for fi47 := 0; fi47 < 3; fi47++ {
                                        if true {
                                            v48 := GenS7{F0: (65535 != 903), F1: 2147483647, F2: (18446744073709551615 < 261)}
                                            v49 := GenS8{F0: 32767}
                                            if len(v4) > 0 {
                                                v4[0] = (v1 + ("foo bar" + v1))
                                            }
                                            _ = v48
                                            _ = v49
                                        }
                                        _ = fi47
                                    }
                                    _ = fi46
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                fmt.Println(v1, v42, v1)
                            }
                        }
                    }
                    v50 := GenS6{F0: (^(v42)), F1: 0}
                    if false {
                        if ((false && false) || false) {
                            if v3 {
                                for fi51 := 0; fi51 < 12; fi51++ {
                                    var v52 int16 = 1
                                    _ = fi51
                                    _ = v52
                                }
                            }
                        } else {
                            v53 := GenS9{F0: 2147483647, F1: 42}
                            if v3 {
                                v3 = false
                            }
                            v54 := GenS2{F0: 0, F1: -1, F2: 42}
                            _ = v53
                            _ = v54
                        }
                        _ = len(v4)
                        for fi55 := 0; fi55 < 4; fi55++ {
                            var v56 uint64 = 106
                            _ = fi55
                            _ = v56
                        }
                    }
                    _ = v50
                }
                var v57 string = ("" + v1)
                v42 += v42
                _ = v57
            }
            if (v3 && true) {
                var v58 bool = (1.5 <= float32(1e10))
                v59 := GenS10{F0: -9223372036854775808, F1: v1}
                _ = v58
                _ = v59
            }
        } else {
            if ((!v3) || ((!(!(!(v42 >= (v42 / (((v42 >> (((v2 % (((v2 & v2)) | 1))) & 63))) | 1)))))) && (v3 || v3))) {
                fmt.Println(v42, v2, v3)
            } else {
                v3 = (!(v3 || (((!v3) || false) && (42 == 42))))
            }
        }
        v60 := GenS11{F0: "hello"}
        if (1 > 2) {
            var v61 string = v1
            _ = v61
        } else {
            if (!(238 < 250)) {
                fmt.Println(v1, v1)
                var v62 uint16 = 255
                _ = v62
            }
            var v63 uint64 = 2
            v64 := GenS2{F0: -599, F1: -1, F2: uint32(v2)}
            _ = v63
            _ = v64
        }
        _ = v60
    } else {
        v65 := GenS5{F0: (v2 + v2), F1: (v42 - (v42 >> ((255) & 63))), F2: (v2 & (v2 << (((v2 + 567)) & 63)))}
        var v66 int64 = 1
        fmt.Println(v2, v66, v3)
        _ = v65
        _ = v66
    }
    fmt.Println(v1, v42, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
