package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int16, p1 bool, p2 int16) int16 {
    p2 = p0
    p0 += p0
    if ((!p1) && (-32768 < (^(p2)))) {
        switch "a" {
        case "hello":
            for fi7 := 0; fi7 < 1; fi7++ {
                v8 := GenS1{F0: 2147483647, F1: fi7}
                _ = fi7
                _ = v8
            }
        case "a":
            var v9 uint8 = 234
            _ = v9
        case "foo bar":
            var v10 int64 = -814
            _ = v10
        }
        var v11 float32 = -2.25
        _ = v11
    } else {
        fmt.Println(p1, p1)
    }
    var v12 bool = p1
    p2 = (-(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p2
}
func genFn1(p0 uint32, p1 int32) uint64 {
    switch 291 {
    case 128:
        if ((!false) || (!true)) {
            v13 := GenS2{F0: p0}
            _ = v13
        }
    case -543:
        var v14 uint16 = 255
        _ = v14
    case 42:
        var v15 uint16 = 65535
        _ = v15
    }
    var v16 string = ("" + "")
    switch 575 {
    case 605:
        v17 := GenS2{F0: 1}
        _ = v17
    default:
        var v18 int = 32767
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v16
    return 4294967295
}

func main() {
    var v1 float64 = (1.5 * 0.0)
    var v2 int16 = -146
    v3 := []uint32{890}
    var v4 float32 = 3.14159
    var v5 int32 = 1
    var v6 float32 = v4
    for fi19 := 0; fi19 < 16; fi19++ {
        if (894 < -9223372036854775808) {
            if false {
                if false {
                    var v20 int32 = (v5 % ((v5) | 1))
                    var v21 int32 = -2147483648
                    switch "" {
                    case "\n\t":
                        switch fi19 {
                        case 796:
                            var v22 int64 = -977
                            _ = v22
                        case 1:
                            if true {
                                var v23 uint8 = 0
                                _ = v23
                            }
                        }
                    }
                    _ = v20
                    _ = v21
                } else {
                    v2 = genFn0((^(v2)), true, genFn0((-(v2)), true, -32768))
                }
                var v24 int32 = v5
                _ = v24
            } else {
                switch ("\n\t" + ("foo bar" + ("a" + "foo bar"))) {
                case "hello":
                    var v25 int8 = -1
                    _ = v25
                case "\n\t":
                    if true {
                        if false {
                            v6 -= (v6 - 1e10)
                            if (!true) {
                                var v26 int32 = 1
                                _ = v26
                            }
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 2
                            }
                            for fi27 := 0; fi27 < 3; fi27++ {
                                var v28 int32 = v5
                                if (false || (!(false || true))) {
                                    if true {
                                        fmt.Println(v2, v5, fi19)
                                    } else {
                                        fmt.Println(v1, v2)
                                        var v29 int16 = int16(fi19)
                                        _ = v29
                                    }
                                } else {
                                    v30 := GenS1{F0: (v28 / (((v5 + (v28 | (v5 * (v28 | (v5 % (((v28 * (v28 * 42))) | 1))))))) | 1)), F1: fi19}
                                    fmt.Println(v4)
                                    v5 = (v28 << ((255) & 63))
                                    _ = v30
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi27
                                _ = v28
                            }
                        }
                        fmt.Println(v4, v2, v6)
                        v2 -= (-(v2))
                    } else {
                        _ = len(v3)
                        if false {
                            v3 = append(v3, 224)
                            v2 = genFn0((v2 - 1000), false, (v2 + v2))
                        } else {
                            var v31 uint64 = 18446744073709551615
                            _ = v31
                        }
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                        case "":
                            var v32 int32 = 0
                            _ = v32
                        case "\n\t":
                            switch (^(fi19)) {
                            case -2147483648:
                                v2 = -112
                            case 1:
                                switch "" {
                                case "\n\t":
                                    v33 := GenS3{F0: 73}
                                    _ = v33
                                case "hello":
                                    v5 = v5
                                }
                            case -438:
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    v2 = (v2 % (((v2 >> ((4294967295) & 63))) | 1))
                                    _ = fi34
                                }
                            default:
                                for fi35 := 0; fi35 < 11; fi35++ {
                                    var v36 int = fi19
                                    _ = fi35
                                    _ = v36
                                }
                            }
                        }
                    }
                default:
                    if (255 < 202) {
                        if false {
                            v4 = -2.25
                        }
                        var v37 float64 = v1
                        _ = v37
                    }
                }
                switch ("foo bar" + "") {
                case "foo bar":
                    v1 -= v1
                case "hello":
                    v38 := GenS4{F0: v5}
                    _ = v38
                case "":
                    if (false && false) {
                        var v39 uint16 = 255
                        v40 := GenS4{F0: (v5 * v5)}
                        if (42 > 2) {
                            if (false || false) {
                                v3 = append(v3, 42)
                            }
                        } else {
                            var v41 int = (fi19 + (fi19 & fi19))
                            v39 += v39
                            _ = v41
                        }
                        _ = v39
                        _ = v40
                    }
                default:
                    v1 -= 0.0
                }
            }
            if (!(!(!((v4 + v4) < v6)))) {
                v42 := GenS4{F0: v5}
                _ = v42
            } else {
                switch "foo bar" {
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    var v43 float64 = (v1 * -2.25)
                    _ = v43
                }
            }
            v44 := GenS3{F0: v5}
            _ = v44
        }
        _ = len(v3)
        if (0 < -858) {
            for fi45 := 0; fi45 < 18; fi45++ {
                if false {
                    var v46 uint8 = 124
                    fmt.Println(v46, v46)
                    _ = v46
                }
                var v47 int32 = -646
                _ = fi45
                _ = v47
            }
            v48 := GenS4{F0: (v5 % (((v5 | (v5 % (((^(v5))) | 1)))) | 1))}
            if len(v3) > 0 {
                v3[0] = 0
            }
            _ = v48
        }
        _ = fi19
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 += (-2.25 - v4)
    v3 = append(v3, 27)
    v3 = append(v3, 2147483647)
    v49 := GenS3{F0: v5}
    if (2 < 4294967295) {
        v6 -= 1.5
        fmt.Println(v4, v6, v5)
        var v50 int16 = (v2 ^ v2)
        _ = v50
    } else {
        v3 = append(v3, 1)
    }
    for fi51 := 0; fi51 < 1; fi51++ {
        var v52 float64 = (v1 / -2.25)
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
