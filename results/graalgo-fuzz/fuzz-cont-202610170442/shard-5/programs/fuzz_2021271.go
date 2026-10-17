package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int32; F2 int32 }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 uint8; F1 uint8; F2 uint64 }
func genFn0(p0 uint64, p1 uint16, p2 int16) uint32 {
    p2 -= (p2 ^ p2)
    v5 := GenS1{F0: 2, F1: 255, F2: 0}
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(p2, fi6)
        v7 := []int{fi6, (fi6 & fi6), fi6}
        v8 := GenS1{F0: 255, F1: 128, F2: -2147483648}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{}
    var v10 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v9
    return 1
}
func genFn1(p0 uint16, p1 string, p2 int64) uint8 {
    var v11 uint32 = 2147483647
    v12 := GenS2{F0: 1}
    switch p1 {
    case "hello":
        if true {
            var v13 uint64 = 2
            _ = v13
        } else {
            p0 = p0
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := GenS3{F0: p1, F1: 236, F2: (p0 - p0)}
                _ = fi14
                _ = v15
            }
            var v16 int = 9223372036854775807
            _ = v16
        }
    case "a":
        var v17 uint32 = v11
        _ = v17
    case "foo bar":
        v18 := []int{-840}
        _ = v18
    }
    fmt.Println(p2, p0)
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 128
}

func main() {
    var v1 string = "foo bar"
    var v2 int32 = 2147483647
    v3 := make(map[string]string)
    var v4 int16 = -292
    var v19 float64 = 3.14159
    if (false || ((1 == 0) && (0 < -71))) {
        v20 := []int{4294967295, -714, -118}
        v4 = v4
        if false {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 int8 = 127
                _ = fi21
                _ = v22
            }
            switch -2147483648 {
            case 9223372036854775807:
                if (1 <= 42) {
                    for fi23 := 0; fi23 < 13; fi23++ {
                        var v24 int64 = 1
                        switch (fi23 << ((2) & 63)) {
                        case -2147483648:
                            v2 += 42
                        case 1:
                            fmt.Println(v1)
                        default:
                            switch v1 {
                            case "\n\t":
                                v25 := GenS2{F0: 255}
                                _ = v25
                            default:
                                var v26 uint64 = 18446744073709551615
                                _ = v26
                            }
                        }
                        _ = fi23
                        _ = v24
                    }
                    var v27 float64 = 3.14159
                    v4 -= v4
                    _ = v27
                }
            case 116:
                var v28 uint16 = 0
                _ = v28
            default:
                v4 = (v4 * v4)
            }
            for fi29 := 0; fi29 < 20; fi29++ {
                if ((1e10 * 1e10) >= (((float32((1e10 - (v19 - v19))) - 1.5) + (float32(v19) * float32(0.1))) * 0.0)) {
                    for fi30 := 0; fi30 < 13; fi30++ {
                        var v31 uint8 = 127
                        _ = fi30
                        _ = v31
                    }
                    var v32 uint16 = 65535
                    fmt.Println(v32, v4)
                    _ = v32
                }
                var v33 uint = 18446744073709551615
                var v34 int16 = (v4 % ((-1) | 1))
                _ = fi29
                _ = v33
                _ = v34
            }
        } else {
            v35 := GenS4{F0: 2147483647}
            if len(v20) > 0 {
                v20[0] = -633
            }
            fmt.Println(v1, v1, v2)
            _ = v35
        }
        _ = v20
    }
    v1 = v1
    v1 = v1
    if (222 > 255) {
        for fi36 := 0; fi36 < 14; fi36++ {
            for fi37 := 0; fi37 < 4; fi37++ {
                v4 += (v4 - v4)
                switch ((v1 + v1) + (v1 + ((v1 + (v1 + v1)) + "foo bar"))) {
                case "\n\t":
                    var v38 uint16 = 32767
                    _ = v38
                case "foo bar":
                    v1 = v1
                default:
                    delete(v3, ((v1 + ((v1 + "") + (("a" + ("a" + (("a" + "\n\t") + v1))) + v1))) + v1))
                }
                _ = fi37
            }
            fmt.Println(fi36)
            var v39 int = fi36
            _ = fi36
            _ = v39
        }
        v19 -= ((v19 * ((v19 * 1e10) - v19)) * ((1.5 * float64(-2.25)) + v19))
    }
    for fi40 := 0; fi40 < 1; fi40++ {
        fmt.Println(fi40)
        var v41 uint16 = 42
        for fi42 := 0; fi42 < 7; fi42++ {
            if true {
                if (4294967295 < 4294967295) {
                    v43 := GenS5{F0: -3}
                    v41 += 1
                    _ = v43
                } else {
                    if (!((!(!true)) && ((1.5 / 1e10) < (0.1 / float32((float64(1e10) + v19)))))) {
                        delete(v3, "a")
                        if (false && false) {
                            v44 := v3[""]
                            v41 = 32767
                            var v45 uint32 = 874
                            _ = v44
                            _ = v45
                        }
                        var v46 uint16 = (v41 % ((2) | 1))
                        _ = v46
                    }
                }
                v47 := GenS6{F0: v2}
                v41 = 2
                _ = v47
            } else {
                v48 := []int{4294967295, (fi40 ^ (fi40 & fi42)), (fi40 << ((65535) & 63))}
                switch fi40 {
                case -9223372036854775808:
                    if len(v48) > 0 {
                        v48[0] = (fi42 << ((65535) & 63))
                    }
                case 2147483647:
                    if true {
                        var v49 uint8 = 158
                        if len(v48) > 0 {
                            v48[0] = (fi40 - fi40)
                        }
                        _ = v49
                    }
                case 1:
                    switch "hello" {
                    case "foo bar":
                        delete(v3, ("foo bar" + ("hello" + v1)))
                    case "":
                        v1 = "\n\t"
                    }
                }
                _ = v48
            }
            _ = fi42
        }
        _ = fi40
        _ = v41
    }
    if (-1 <= 0) {
        if (!(v4 != (-(v4)))) {
            v3[""] = ("hello" + "foo bar")
        }
    }
    var v50 string = ((((v1 + v1) + ((v1 + "a") + "a")) + v1) + "foo bar")
    for fi51 := 0; fi51 < 18; fi51++ {
        v3[(v50 + "")] = v50
        v4 = v4
        _ = fi51
    }
    v1 = ("foo bar" + ((((v1 + v50) + (((("hello" + (((v50 + v1) + ("foo bar" + (("hello" + v50) + (((v50 + ((v50 + ("\n\t" + ((v50 + (("hello" + (v1 + ((v50 + v1) + "\n\t"))) + v50)) + v1))) + (v50 + "hello"))) + v50) + v1)))) + (("a" + v1) + v50))) + v50) + v50) + v1)) + (((v1 + v50) + (v1 + (v50 + v50))) + (v50 + v50))) + v1))
    switch 0 {
    case 42:
        for fi52 := 0; fi52 < 16; fi52++ {
            if (!true) {
                switch fi52 {
                case 32767:
                    if true {
                        v53 := GenS7{F0: 4294967295, F1: (v2 >> ((278) & 63)), F2: int32(v4)}
                        var v54 float64 = (((3.14159 + float64((((3.14159 * 1e10) / float32((v19 + 1e10))) - 1.5))) * float64(float32(0.1))) * (v19 * ((v19 * v19) + (v19 * v19))))
                        v2 = (v2 + (v2 / ((v2) | 1)))
                        _ = v53
                        _ = v54
                    }
                case 127:
                    fmt.Println(v1, v19, v1)
                case 128:
                    var v55 string = ((("a" + ("" + "\n\t")) + v50) + v50)
                    _ = v55
                default:
                    fmt.Println(v1, v19, v19)
                }
            }
            var v56 float32 = 0.0
            _ = fi52
            _ = v56
        }
    default:
        fmt.Println(v2)
    }
    v4 = -939
    v57 := GenS8{F0: 0, F1: 0}
    v50 = "a"
    switch -9223372036854775808 {
    case 630:
        v58 := GenS9{F0: 861, F1: 692}
        _ = v58
    case 42:
        for fi59 := 0; fi59 < 17; fi59++ {
            var v60 uint8 = genFn1(930, v50, 168)
            v61 := GenS10{F0: 995, F1: (v2 >> ((867) & 63))}
            if true {
                if (2147483647 >= fi59) {
                    var v62 int = (fi59 + (fi59 >> ((42) & 63)))
                    v62 += v62
                    _ = v62
                }
            }
            _ = fi59
            _ = v60
            _ = v61
        }
    case -9223372036854775808:
        switch "foo bar" {
        case "\n\t":
            var v63 string = v1
            _ = v63
        case "hello":
            v64 := GenS11{F0: 42, F1: genFn1(1, v50, 315), F2: 376}
            _ = v64
        }
    }
    var v65 float32 = (((float32(v19) - 1e10) + -2.25) * (float32(v19) / ((3.14159 * 0.0) + 1e10)))
    v65 = (v65 - v65)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v57
    _ = v65
}
