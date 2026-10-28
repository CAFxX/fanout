package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 int32, p2 int16) int8 {
    if (((p2 ^ p2) > -32768) || (("foo bar" + "foo bar") != "")) {
        var v7 string = ""
        _ = v7
    } else {
        p1 += p1
        v8 := []int{4294967295, 9223372036854775807}
        fmt.Println(p2, p0)
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        p2 = p2
        _ = fi9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := []int{fi10}
        if true {
            var v12 float64 = float64(((0.0 + ((0.1 + 1e10) + 0.0)) / 3.14159))
            v12 += (float64(0.1) / 1.0)
            v13 := GenS1{F0: -1, F1: -1}
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() string {
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint32 = 669
        _ = fi14
        _ = v15
    }
    v16 := []int{-9223372036854775808, 1, -2147483648}
    _ = len(v16)
    _ = v16
    return ""
}

func main() {
    var v1 int8 = -34
    var v2 uint16 = 320
    var v3 int16 = 1000
    v4 := []uint32{249, 0, 828}
    var v5 uint64 = 2
    var v6 int16 = 880
    for fi17 := 0; fi17 < 14; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            v6 = 876
            _ = fi18
        }
        v5 += (v5 * v5)
        _ = fi17
    }
    v3 += (v3 / (((v3 >> ((550) & 63))) | 1))
    var v19 int32 = 548
    var v20 int = -9223372036854775808
    var v21 uint64 = v5
    fmt.Println(v1, v5, v3)
    if true {
        fmt.Println(v6, v6, v6)
        _ = len(v4)
        if ((false || false) || (!((!true) || (((!((2 < 0) && ((!true) && (65535 <= 1)))) && ((1e10 > ((0.1 * 0.0) - 3.14159)) || ((v1 < 1) && (int(v2) != -2147483648)))) || (false && (42 >= v3)))))) {
            v20 += 32768
            if false {
                for fi22 := 0; fi22 < 4; fi22++ {
                    for fi23 := 0; fi23 < 15; fi23++ {
                        fmt.Println(v1, v19)
                        v5 = (v5 + (v5 / ((1) | 1)))
                        switch 56 {
                        case 127:
                            var v24 float32 = (float32(-2.25) + 0.1)
                            _ = v24
                        case -776:
                            v20 += (fi22 % ((v20) | 1))
                        default:
                            if true {
                                v4 = append(v4, 156)
                                v4 = append(v4, 339)
                                switch ("hello" + "hello") {
                                case "foo bar":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    var v25 float32 = 1.5
                                    _ = v25
                                }
                            }
                        }
                        var v26 float32 = 1.5
                        _ = fi23
                        _ = v26
                    }
                    var v27 int = 127
                    _ = fi22
                    _ = v27
                }
            } else {
                for fi28 := 0; fi28 < 12; fi28++ {
                    var v29 int16 = v3
                    v20 += 9223372036854775807
                    _ = fi28
                    _ = v29
                }
                v30 := GenS2{F0: (v19 * (v19 ^ 2147483647))}
                v20 = (v20 + (v20 - 32767))
                _ = v30
            }
        } else {
            for fi31 := 0; fi31 < 6; fi31++ {
                for fi32 := 0; fi32 < 18; fi32++ {
                    v5 = (v21 << ((255) & 63))
                    _ = fi32
                }
                var v33 int = v20
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi31
                _ = v33
            }
        }
    }
    v3 += (-(v6))
    if true {
        v20 -= (v20 - 9223372036854775807)
        v3 = (v3 / (((v6 % (((v6 ^ v3)) | 1))) | 1))
        var v34 int16 = (v6 << ((570) & 63))
        _ = v34
    } else {
        switch "\n\t" {
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "\n\t":
            var v35 uint16 = v2
            _ = v35
        case "":
            var v36 uint8 = 2
            _ = v36
        default:
            var v37 int16 = 1000
            _ = v37
        }
    }
    v19 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
