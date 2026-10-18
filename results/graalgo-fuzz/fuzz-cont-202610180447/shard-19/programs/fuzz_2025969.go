package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64; F1 int32; F2 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int }
func genFn0(p0 uint64, p1 uint8) uint64 {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float64 = (0.1 / 1.0)
        var v9 int8 = 100
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint8 = p1
            fmt.Println(v9)
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        v12 := GenS1{F0: 255}
        _ = v12
    } else {
        switch -9223372036854775808 {
        case -555:
            for fi13 := 0; fi13 < 16; fi13++ {
                p1 += p1
                _ = fi13
            }
        case 42:
            v14 := GenS1{F0: 65535}
            _ = v14
        case 1:
            var v15 float32 = 1e10
            _ = v15
        }
        var v16 int16 = 88
        _ = v16
    }
    v17 := GenS2{F0: 2}
    p0 = 204
    _ = p0
    _ = p1
    _ = v17
    return 4294967295
}
func genFn1() uint8 {
    var v18 uint32 = 2147483648
    v18 -= (v18 % ((v18) | 1))
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS3{F0: 240, F1: 18446744073709551615}
        v21 := []int{fi19, -2147483648, fi19, fi19}
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := GenS3{F0: 0, F1: 675}
            _ = fi22
            _ = v23
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    _ = v18
    return 140
}

func main() {
    var v1 uint8 = 151
    var v2 float64 = 1.5
    var v3 float64 = -2.25
    v4 := make(map[string]string)
    var v5 bool = (64 <= 540)
    v6 := []uint32{}
    if ("a" <= (("\n\t" + "") + ("foo bar" + ("\n\t" + ("foo bar" + "hello"))))) {
        v24 := GenS4{F0: "hello"}
        for fi25 := 0; fi25 < 16; fi25++ {
            if v5 {
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 uint8 = (v1 << ((255) & 63))
                    if (false || v5) {
                        if (!(v1 != genFn1())) {
                            v28 := GenS1{F0: 901}
                            _ = v28
                        }
                    } else {
                        var v29 int32 = 0
                        v30 := GenS2{F0: 1}
                        _ = v29
                        _ = v30
                    }
                    _ = fi26
                    _ = v27
                }
            }
            if (42 > 255) {
                v5 = v5
                var v31 int8 = -103
                v32 := GenS5{F0: 858, F1: 4294967295, F2: 1}
                _ = v31
                _ = v32
            } else {
                for fi33 := 0; fi33 < 17; fi33++ {
                    v34 := GenS6{F0: 687, F1: genFn1()}
                    if false {
                        v2 += v3
                    } else {
                        _ = len(v4)
                        v1 = v1
                    }
                    _ = fi33
                    _ = v34
                }
                fmt.Println(v5)
                for fi35 := 0; fi35 < 17; fi35++ {
                    _ = len(v4)
                    var v36 int = -1
                    if false {
                        switch fi25 {
                        case -1:
                            switch "\n\t" {
                            case "hello":
                                v37 := GenS7{F0: 1}
                                _ = v37
                            case "\n\t":
                                v38 := GenS2{F0: 878}
                                _ = v38
                            default:
                                if true {
                                    v39 := GenS3{F0: v1, F1: 839}
                                    _ = v39
                                } else {
                                    v3 += (3.14159 / 0.1)
                                }
                            }
                        default:
                            v36 += v36
                        }
                        v40 := GenS4{F0: ("\n\t" + "")}
                        _ = v40
                    }
                    var v41 int16 = 1000
                    _ = fi35
                    _ = v36
                    _ = v41
                }
            }
            v42 := GenS8{F0: 0, F1: 0, F2: -128}
            _ = fi25
            _ = v42
        }
        v43 := GenS9{F0: 18446744073709551615}
        _ = v24
        _ = v43
    } else {
        var v44 uint64 = 454
        v45 := GenS1{F0: 42}
        _ = v44
        _ = v45
    }
    var v46 int8 = -128
    v2 = (1.5 * float64(1e10))
    v47 := GenS5{F0: 211, F1: 2, F2: 817}
    switch 32767 {
    case 9223372036854775807:
        var v48 uint8 = v1
        _ = v48
    }
    var v49 int = 2147483647
    switch v49 {
    case -321:
        var v50 uint32 = 2147483648
        _ = v50
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    for fi51 := 0; fi51 < 5; fi51++ {
        fmt.Println(fi51, v2, v49)
        v52 := GenS10{F0: 32768}
        var v53 string = "foo bar"
        _ = fi51
        _ = v52
        _ = v53
    }
    fmt.Println(v3, v1)
    var v54 uint64 = 4294967295
    for fi55 := 0; fi55 < 2; fi55++ {
        v4[("hello" + ((("hello" + "foo bar") + ("hello" + "foo bar")) + ""))] = "\n\t"
        _ = fi55
    }
    if true {
        v5 = (v5 && v5)
    }
    if v5 {
        for fi56 := 0; fi56 < 1; fi56++ {
            delete(v4, ("" + "foo bar"))
            _ = fi56
        }
        fmt.Println(v3, v1)
    }
    if v5 {
        v5 = ((v46 < (v46 & (v46 % ((v46) | 1)))) && true)
        switch "\n\t" {
        case "\n\t":
            switch (((("hello" + "foo bar") + "a") + "") + "hello") {
            case "hello":
                fmt.Println(v54, v54)
            case "":
                v3 -= (0.0 + v3)
            default:
                for fi57 := 0; fi57 < 3; fi57++ {
                    v54 += 1
                    delete(v4, ("a" + "\n\t"))
                    v2 += ((v2 + (float64(1.5) / v3)) - v3)
                    _ = fi57
                }
            }
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    _ = len(v6)
    if (("a" != "hello") && v5) {
        v2 = ((v2 - ((v2 - (1.5 - v2)) * v2)) + 3.14159)
        var v58 int32 = 1
        _ = v58
    } else {
        if len(v6) > 0 {
            v6[0] = 0
        }
        fmt.Println(v49, v46, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v49
    _ = v5
    _ = v54
    _ = v6
}
