package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 bool }
type GenS3 struct { F0 uint16; F1 int; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int }
func genFn0(p0 int64) string {
    if false {
        var v6 uint16 = 32767
        var v7 float64 = 0.1
        _ = v6
        _ = v7
    }
    var v8 int32 = -1
    _ = p0
    _ = v8
    return "hello"
}
func genFn1(p0 int64) int32 {
    p0 = (p0 | (p0 + (p0 + 370)))
    if true {
        if false {
            v9 := []int{}
            _ = v9
        }
    } else {
        var v10 int32 = 2147483647
        p0 -= (p0 * p0)
        p0 = (p0 - p0)
        _ = v10
    }
    v11 := GenS1{F0: 255}
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint8 = 0
        _ = fi12
        _ = v13
    }
    if (122 <= 1) {
        p0 = p0
        v14 := []int{32768, 42, 32767, 32767}
        p0 -= p0
        _ = v14
    }
    _ = p0
    _ = v11
    return -1
}

func main() {
    var v1 int64 = 1
    v2 := []string{"a"}
    var v3 bool = false
    var v4 uint64 = 0
    var v5 int8 = -1
    v4 = (v4 * v4)
    if v3 {
        var v15 float64 = (1e10 / float64(0.1))
        if ((v4 + 309) < 437) {
            v1 = (^(v1))
            v16 := GenS1{F0: 0}
            if true {
                switch 1 {
                case -9223372036854775808:
                    _ = len(v2)
                case 2147483647:
                    var v17 uint64 = (v4 / ((4294967295) | 1))
                    _ = v17
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                switch -9223372036854775808 {
                case -1:
                    if false {
                        if (!v3) {
                            _ = len(v2)
                            v5 += (v5 >> ((2) & 63))
                            v18 := GenS2{F0: 890, F1: 93, F2: (!(v3 || ((v15 - float64(0.1)) == (v15 * 3.14159))))}
                            _ = v18
                        } else {
                            if v3 {
                                var v19 uint64 = 18446744073709551615
                                fmt.Println(v1)
                                _ = v19
                            } else {
                                v4 = (v4 / ((48) | 1))
                                if true {
                                    v1 = (v1 * v1)
                                    v20 := GenS2{F0: 1, F1: 0, F2: false}
                                    _ = v20
                                } else {
                                    if v3 {
                                        var v21 uint64 = 2
                                        _ = v21
                                    }
                                    v2 = append(v2, "foo bar")
                                }
                                for fi22 := 0; fi22 < 14; fi22++ {
                                    var v23 uint64 = (^(v4))
                                    var v24 string = "foo bar"
                                    v25 := GenS3{F0: 280, F1: 987, F2: 127}
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            }
                        }
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v3, v15)
                        case "hello":
                            fmt.Println(v15)
                        case "a":
                            v1 = v1
                        }
                        for fi26 := 0; fi26 < 1; fi26++ {
                            v4 += (v4 >> ((166) & 63))
                            v4 = (v4 - v4)
                            _ = fi26
                        }
                    }
                default:
                    v27 := GenS2{F0: 2147483647, F1: 128, F2: false}
                    _ = v27
                }
            }
            _ = v16
        }
        v4 += v4
        _ = v15
    } else {
        for fi28 := 0; fi28 < 2; fi28++ {
            if false {
                for fi29 := 0; fi29 < 6; fi29++ {
                    switch fi29 {
                    case 926:
                        var v30 int8 = (v5 ^ (v5 / ((127) | 1)))
                        _ = v30
                    default:
                        var v31 int16 = 1
                        _ = v31
                    }
                    v3 = (true || v3)
                    var v32 uint8 = 42
                    var v33 uint32 = 65535
                    v34 := GenS4{F0: fi29}
                    _ = fi29
                    _ = v32
                    _ = v33
                    _ = v34
                }
            }
            if false {
                v3 = v3
                if false {
                    v2 = append(v2, ("foo bar" + "\n\t"))
                    if (false || v3) {
                        fmt.Println(v3, v4, fi28)
                        for fi35 := 0; fi35 < 4; fi35++ {
                            v5 = v5
                            _ = fi35
                        }
                    }
                    for fi36 := 0; fi36 < 7; fi36++ {
                        var v37 string = ("a" + ("hello" + "foo bar"))
                        var v38 int = int(v5)
                        v3 = (!true)
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                }
                if v3 {
                    v2 = append(v2, "hello")
                    v2 = append(v2, "a")
                    v1 = (v1 & v1)
                } else {
                    var v39 uint8 = 239
                    _ = v39
                }
            } else {
                if true {
                    var v40 int8 = -128
                    v4 = v4
                    v41 := GenS5{F0: true, F1: false}
                    _ = v40
                    _ = v41
                }
            }
            v42 := GenS6{F0: 1}
            _ = fi28
            _ = v42
        }
        v43 := GenS2{F0: 2, F1: 204, F2: (false || v3)}
        if ((v3 || (((0.1 * -2.25) / ((1e10 + (-2.25 * (float64(0.1) - float64((0.0 - -2.25))))) / 0.1)) >= ((float64((-2.25 - ((1.5 - 0.1) / (float32(3.14159) * 3.14159)))) / 1.0) / (float64(3.14159) / float64(0.1))))) || true) {
            if (v3 || true) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            if (-2147483648 != 32768) {
                var v44 string = ("foo bar" + ("hello" + "\n\t"))
                if (!v3) {
                    _ = len(v2)
                }
                _ = v44
            }
            v3 = v3
        } else {
            var v45 int = 656
            fmt.Println(v4)
            _ = v45
        }
        _ = v43
    }
    v46 := GenS7{F0: 4294967295}
    fmt.Println(v3, v4, v5)
    v2 = append(v2, "hello")
    var v47 uint64 = (v4 - v4)
    for fi48 := 0; fi48 < 11; fi48++ {
        v1 -= v1
        v4 -= v47
        var v49 int64 = int64(v4)
        _ = fi48
        _ = v49
    }
    v1 += (v1 * -656)
    var v50 int64 = v1
    switch -9223372036854775808 {
    case -863:
        var v51 int8 = v5
        _ = v51
    case 9223372036854775807:
        v50 = v50
    case 32767:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v5 = -19
    v2 = append(v2, ("foo bar" + (("\n\t" + "hello") + "\n\t")))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v50
}
