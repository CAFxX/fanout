package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint16 }
func genFn0(p0 int32, p1 bool, p2 int64) int {
    fmt.Println(p1, p2, p0)
    p1 = ((42 >= 0) || (-27 > 1))
    _ = p0
    _ = p1
    _ = p2
    return -89
}
func genFn1(p0 uint64) int32 {
    v4 := []int{-483}
    p0 += (p0 >> ((42) & 63))
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: 65535, F1: 128, F2: 2}
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int = (fi7 >> ((18446744073709551615) & 63))
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            if len(v4) > 0 {
                v4[0] = 32767
            }
            var v10 int16 = -32768
            _ = fi9
            _ = v10
        }
        _ = fi5
        _ = v6
    }
    var v11 uint16 = 32767
    _ = p0
    _ = v11
    _ = v4
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = ((1.5 * (0.1 + (1.5 - 1.5))) / -2.25)
    var v3 uint8 = 2
    v12 := v1[1]
    var v13 uint16 = 416
    v2 -= v2
    var v14 bool = ((42 != 42) && v12)
    switch 794 {
    case -9223372036854775808:
        for fi15 := 0; fi15 < 13; fi15++ {
            for fi16 := 0; fi16 < 19; fi16++ {
                switch genFn0(genFn1(42), true, -1) {
                case 543:
                    switch (fi15 | (fi15 + (fi15 >> ((4294967295) & 63)))) {
                    case 1:
                        var v17 int32 = 1
                        _ = v17
                    case 0:
                        v13 = (^(v13))
                    case 9223372036854775807:
                        v18 := []int{fi15}
                        _ = v18
                    }
                }
                _ = fi16
            }
            _ = fi15
        }
    case 0:
        for fi19 := 0; fi19 < 17; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                if v14 {
                    v13 = 65535
                    if (((-1 > 1) || (-2.25 >= (0.0 - 1e10))) || ((!false) || (952 == 439))) {
                        if false {
                            v1[0] = true
                            v3 = 1
                        }
                        v21 := GenS1{F0: 266, F1: 0, F2: 42}
                        var v22 int32 = -851
                        if (-27 < 127) {
                            var v23 string = ("a" + ("foo bar" + (("a" + ("foo bar" + ("\n\t" + "\n\t"))) + ("" + "\n\t"))))
                            v22 += (v22 >> ((18446744073709551615) & 63))
                            _ = v23
                        } else {
                            fmt.Println(v2, fi20)
                            v22 += v22
                            if (!true) {
                                var v24 float64 = 0.0
                                v25 := GenS1{F0: (v13 - v13), F1: 225, F2: 4294967295}
                                _ = v24
                                _ = v25
                            }
                        }
                        fmt.Println(fi19)
                        _ = v21
                        _ = v22
                    }
                } else {
                    switch 42 {
                    case -2147483648:
                        fmt.Println(v12, v12)
                    case 0:
                        v3 += 127
                        switch "" {
                        case "a":
                            _ = len(v1)
                        case "foo bar":
                            v26 := v1[872]
                            _ = v26
                        case "\n\t":
                            var v27 bool = true
                            _ = v27
                        }
                        if (2 == 188) {
                            v28 := GenS1{F0: v13, F1: v3, F2: uint32(v3)}
                            if v14 {
                                v29 := []int{(fi20 + (fi19 << ((uint(v13)) & 63))), 1, 2147483647}
                                fmt.Println(v2, v12)
                                var v30 float64 = 3.14159
                                _ = v29
                                _ = v30
                            } else {
                                v12 = (-2147483648 < 524)
                                var v31 uint8 = 2
                                _ = v31
                            }
                            _ = v28
                        }
                        var v32 uint32 = 65535
                        _ = v32
                    case -9223372036854775808:
                        var v33 bool = true
                        _ = v33
                    }
                }
                fmt.Println(v2)
                _ = fi20
            }
            v34 := v1[2]
            _ = fi19
            _ = v34
        }
    default:
        fmt.Println(v2)
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        if true {
            v14 = v14
            v14 = ((!v14) && ((!v14) || (v14 || v12)))
            v2 = 1e10
        }
        if ((!false) && (v12 && (!(!false)))) {
            v14 = (!v14)
        } else {
            for fi36 := 0; fi36 < 17; fi36++ {
                if (!(v3 <= (v3 ^ (^(v3))))) {
                    v37 := GenS2{F0: 179, F1: 32767, F2: 2}
                    _ = v37
                }
                var v38 int = (fi36 + genFn0(42, true, 272))
                v39 := GenS3{F0: 2147483647}
                _ = fi36
                _ = v38
                _ = v39
            }
        }
        _ = fi35
    }
    v14 = (false || v14)
    v13 = (v13 ^ (v13 << ((195) & 63)))
    var v40 uint16 = v13
    fmt.Println(v12, v13, v13)
    switch 1 {
    case 128:
        fmt.Println(v14, v13)
    case -410:
        v13 -= (v13 ^ 255)
    }
    var v41 uint64 = 4294967295
    fmt.Println(v14)
    if true {
        if true {
            v1[654] = (-64 >= -292)
            v42 := GenS1{F0: (^(v40)), F1: (^(v3)), F2: 145}
            fmt.Println(v12)
            _ = v42
        }
        v3 = v3
        switch ("\n\t" + "") {
        case "\n\t":
            var v43 int = -9223372036854775808
            _ = v43
        case "a":
            for fi44 := 0; fi44 < 1; fi44++ {
                v45 := v1[4294967295]
                _ = fi44
                _ = v45
            }
        default:
            if true {
                var v46 string = "foo bar"
                for fi47 := 0; fi47 < 11; fi47++ {
                    var v48 uint16 = (v13 >> ((uint(v3)) & 63))
                    v12 = false
                    _ = fi47
                    _ = v48
                }
                var v49 float32 = ((v2 / (float32((1e10 / (-2.25 - (0.1 - 1e10)))) + float32(0.1))) + ((v2 * v2) * (v2 / v2)))
                _ = v46
                _ = v49
            }
        }
    }
    v50 := GenS4{F0: 1, F1: v13}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v50
}
