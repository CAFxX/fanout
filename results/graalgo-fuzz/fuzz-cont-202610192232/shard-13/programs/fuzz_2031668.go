package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint8; F1 int }
func genFn0(p0 int16, p1 uint) uint32 {
    p0 = p0
    for fi4 := 0; fi4 < 19; fi4++ {
        p1 += (p1 ^ (p1 >> ((p1) & 63)))
        p1 += p1
        _ = fi4
    }
    if false {
        v5 := GenS1{F0: 0, F1: 849, F2: -629}
        for fi6 := 0; fi6 < 4; fi6++ {
            if true {
                var v7 int64 = 42
                p1 = p1
                _ = v7
            } else {
                p0 -= (p0 % ((p0) | 1))
            }
            v8 := []int{(-(fi6)), (fi6 & (fi6 * fi6)), 1}
            for fi9 := 0; fi9 < 9; fi9++ {
                p1 += (p1 % (((^(p1))) | 1))
                _ = fi9
            }
            _ = fi6
            _ = v8
        }
        var v10 uint16 = 1000
        _ = v10
        _ = v5
    } else {
        v11 := GenS1{F0: 0, F1: 128, F2: 32767}
        var v12 int = -731
        _ = v11
        _ = v12
    }
    fmt.Println(p0)
    p0 += (p0 ^ -85)
    _ = p0
    _ = p1
    return 928
}
func genFn1(p0 int16, p1 uint, p2 int32) uint8 {
    p0 = p0
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(fi13, p0, fi13)
        var v14 int64 = 9223372036854775807
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint8 = 5
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS1{F0: 1, F1: (fi17 + (fi17 >> ((p1) & 63))), F2: 32768}
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := []int{691, 2147483647, 2147483647, 32767}
    var v2 float32 = (-2.25 / 1.0)
    var v3 uint64 = 420
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS1{F0: (v3 - (v3 & (^(v3)))), F1: fi19, F2: fi19}
        v21 := GenS1{F0: 0, F1: (fi19 ^ -252), F2: fi19}
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: v3, F1: 9223372036854775807, F2: 32768}
    if true {
        var v23 bool = true
        fmt.Println(v2, v3, v23)
        if len(v1) > 0 {
            v1[0] = -769
        }
        _ = v23
    }
    if ((v3 - v3) >= (v3 / ((v3) | 1))) {
        v3 = (v3 + (v3 % ((4294967295) | 1)))
        fmt.Println(v3)
        if (128 < 42) {
            for fi24 := 0; fi24 < 12; fi24++ {
                if (32767 <= 0) {
                    var v25 int = (^(fi24))
                    _ = v25
                }
                if (!true) {
                    if false {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            if false {
                                switch "foo bar" {
                                case "\n\t":
                                    v3 -= 1
                                case "":
                                    switch fi26 {
                                    case 29:
                                        switch 42 {
                                        case 42:
                                            var v27 uint = 65535
                                            _ = v27
                                        case -863:
                                            if true {
                                                v2 = v2
                                                v2 -= v2
                                            }
                                        }
                                    case -292:
                                        v28 := GenS1{F0: 18446744073709551615, F1: fi26, F2: fi26}
                                        _ = v28
                                    case 4294967295:
                                        if true {
                                            if false {
                                                v2 -= (v2 * v2)
                                                if false {
                                                    var v29 float32 = 3.14159
                                                    fmt.Println(v2)
                                                    v1 = append(v1, (fi26 * -9223372036854775808))
                                                    _ = v29
                                                } else {
                                                    _ = len(v1)
                                                    v3 -= (v3 ^ v3)
                                                    var v30 int64 = int64(fi24)
                                                    _ = v30
                                                }
                                                v31 := GenS2{F0: 65, F1: 0}
                                                _ = v31
                                            } else {
                                                v2 = 1e10
                                                v32 := GenS3{F0: -123, F1: 42}
                                                v33 := GenS3{F0: 1, F1: 42}
                                                _ = v32
                                                _ = v33
                                            }
                                            var v34 uint = 2
                                            if len(v1) > 0 {
                                                v1[0] = (^(fi24))
                                            }
                                            _ = v34
                                        } else {
                                            v2 -= v2
                                            var v35 uint8 = 2
                                            if len(v1) > 0 {
                                                v1[0] = fi24
                                            }
                                            _ = v35
                                        }
                                    }
                                }
                                v36 := GenS4{F0: 0, F1: fi26}
                                if true {
                                    var v37 string = ""
                                    switch "a" {
                                    case "foo bar":
                                        switch fi26 {
                                        case 9223372036854775807:
                                            var v38 float64 = (-2.25 - (0.1 - float64(v2)))
                                            _ = v38
                                        case 666:
                                            fmt.Println(fi26, fi24, v2)
                                        }
                                    case "":
                                        v37 = "hello"
                                    case "\n\t":
                                        fmt.Println(v3, v37)
                                    default:
                                        v3 = (v3 % ((v3) | 1))
                                    }
                                    _ = v37
                                }
                                _ = v36
                            } else {
                                if (false || (v2 < (v2 - float32((0.1 * (1e10 - float64(float32((1.5 / (-2.25 / 1.0)))))))))) {
                                    v2 += 0.0
                                } else {
                                    v1 = append(v1, fi26)
                                    switch fi24 {
                                    case 32768:
                                        if len(v1) > 0 {
                                            v1[0] = fi24
                                        }
                                    case 728:
                                        var v39 int64 = 0
                                        _ = v39
                                    case -769:
                                        v3 += (v3 * v3)
                                    default:
                                        var v40 int16 = 42
                                        _ = v40
                                    }
                                }
                                v2 = -2.25
                                v3 -= (v3 / (((v3 % (((^(v3))) | 1))) | 1))
                            }
                            v3 = (v3 * v3)
                            _ = fi26
                        }
                    } else {
                        v1 = append(v1, (fi24 % ((736) | 1)))
                        fmt.Println(v2, v2, fi24)
                    }
                    var v41 bool = (((false || (2147483648 != 772)) || (!((9223372036854775807 >= -849) || ((((("foo bar" + "\n\t") + ("a" + "")) + "") != "") && false)))) && (42 < -32))
                    _ = v41
                } else {
                    fmt.Println(v3)
                    if true {
                        var v42 int16 = 1000
                        for fi43 := 0; fi43 < 14; fi43++ {
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                var v44 uint32 = 2147483647
                                _ = v44
                            }
                            v42 = (v42 + (v42 * -1))
                            var v45 int = fi24
                            _ = fi43
                            _ = v45
                        }
                        v46 := GenS1{F0: (v3 >> ((884) & 63)), F1: -381, F2: -2147483648}
                        _ = v42
                        _ = v46
                    }
                    fmt.Println(fi24)
                }
                _ = fi24
            }
        } else {
            for fi47 := 0; fi47 < 3; fi47++ {
                for fi48 := 0; fi48 < 2; fi48++ {
                    switch (fi48 + -2147483648) {
                    case -558:
                        v2 = v2
                    case 373:
                        var v49 uint16 = 255
                        _ = v49
                    }
                    v1 = append(v1, 46)
                    _ = fi48
                }
                _ = fi47
            }
            var v50 float64 = (0.0 * ((-2.25 - ((3.14159 / 1.0) / (-2.25 + 3.14159))) / 1.5))
            _ = v50
        }
    }
    v3 = (v3 << ((42) & 63))
    var v51 int16 = 667
    for fi52 := 0; fi52 < 20; fi52++ {
        var v53 uint = 255
        for fi54 := 0; fi54 < 20; fi54++ {
            v51 -= v51
            v2 -= 3.14159
            var v55 uint = (v53 % ((v53) | 1))
            _ = fi54
            _ = v55
        }
        _ = fi52
        _ = v53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v51
}
