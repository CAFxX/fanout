package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint64) int64 {
    var v6 int32 = -24
    var v7 string = "hello"
    fmt.Println(v6, p0, p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 int64) uint64 {
    fmt.Println(p0)
    v8 := []int{1, -9223372036854775808}
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint32 = 0
        _ = fi9
        _ = v10
    }
    _ = len(v8)
    p0 = (p0 % ((1) | 1))
    _ = p0
    _ = v8
    return 42
}

func main() {
    var v1 float32 = 0.1
    var v2 uint8 = 2
    var v3 int = 818
    var v4 bool = true
    v5 := []string{"foo bar"}
    fmt.Println(v1, v4, v4)
    var v11 uint16 = 380
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v4, v2)
        if v4 {
            var v13 uint32 = 829
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v13
        }
        if v4 {
            switch "\n\t" {
            case "a":
                for fi14 := 0; fi14 < 4; fi14++ {
                    v2 += v2
                    v11 += v11
                    v3 -= 4294967295
                    _ = fi14
                }
            case "\n\t":
                for fi15 := 0; fi15 < 6; fi15++ {
                    v16 := GenS1{F0: (v2 | (v2 << ((255) & 63)))}
                    if true {
                        var v17 string = ("foo bar" + "a")
                        var v18 uint64 = 0
                        var v19 uint8 = (v2 - 2)
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        fmt.Println(v4)
                        if len(v5) > 0 {
                            v5[0] = ""
                        }
                    }
                    if (!false) {
                        var v20 int64 = genFn0(324, 18446744073709551615)
                        v3 += (v3 >> ((255) & 63))
                        _ = v20
                    }
                    var v21 uint32 = 42
                    v22 := GenS1{F0: 2}
                    _ = fi15
                    _ = v16
                    _ = v21
                    _ = v22
                }
            default:
                switch "\n\t" {
                case "":
                    for fi23 := 0; fi23 < 8; fi23++ {
                        fmt.Println(v4)
                        if (!(v4 || ((float32(float64((v1 * float32(0.1)))) + v1) < v1))) {
                            var v24 int32 = 957
                            switch "hello" {
                            case "\n\t":
                                switch "a" {
                                case "":
                                    var v25 float32 = ((v1 * v1) - ((((v1 - 0.0) / v1) * ((1.5 - -2.25) - (1e10 * (v1 - v1)))) + (v1 + float32(1.5))))
                                    _ = v25
                                default:
                                    fmt.Println(v24, fi23, fi12)
                                    fmt.Println(fi23, fi12)
                                }
                            case "":
                                v3 += (fi12 * (fi23 / (((fi23 | v3)) | 1)))
                            }
                            _ = v24
                        }
                        v4 = true
                        _ = fi23
                    }
                case "foo bar":
                    for fi26 := 0; fi26 < 10; fi26++ {
                        v1 = float32(1e10)
                        var v27 int16 = 1
                        _ = fi26
                        _ = v27
                    }
                case "hello":
                    if (!v4) {
                        var v28 bool = v4
                        var v29 int = fi12
                        if true {
                            for fi30 := 0; fi30 < 19; fi30++ {
                                if len(v5) > 0 {
                                    v5[0] = ""
                                }
                                _ = len(v5)
                                _ = fi30
                            }
                            for fi31 := 0; fi31 < 15; fi31++ {
                                v32 := GenS1{F0: (v2 >> ((2) & 63))}
                                v3 -= -887
                                fmt.Println(v2, v3)
                                fmt.Println(fi12, fi31)
                                var v33 float32 = 1e10
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            fmt.Println(v4, v1, v4)
                        }
                        _ = v28
                        _ = v29
                    } else {
                        v34 := GenS1{F0: (v2 | (v2 >> ((835) & 63)))}
                        for fi35 := 0; fi35 < 8; fi35++ {
                            if ((100 <= 1) || (42 < -264)) {
                                if len(v5) > 0 {
                                    v5[0] = "foo bar"
                                }
                                _ = len(v5)
                                _ = len(v5)
                            }
                            v36 := GenS1{F0: v2}
                            v11 = v11
                            _ = fi35
                            _ = v36
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v34
                    }
                default:
                    var v37 int16 = -318
                    _ = v37
                }
            }
            v5 = append(v5, ((("" + "\n\t") + "foo bar") + "\n\t"))
        } else {
            var v38 uint = 0
            switch 128 {
            case -404:
                v39 := GenS2{F0: false, F1: 65535}
                _ = v39
            case 4294967295:
                var v40 uint8 = 128
                _ = v40
            case 2147483647:
                v41 := GenS3{F0: (("hello" + "hello") + "a")}
                _ = v41
            default:
                v38 = 206
            }
            fmt.Println(v1)
            _ = v38
        }
        _ = fi12
    }
    for fi42 := 0; fi42 < 15; fi42++ {
        var v43 int8 = 42
        fmt.Println(v1, fi42, v2)
        _ = fi42
        _ = v43
    }
    if v4 {
        if v4 {
            for fi44 := 0; fi44 < 16; fi44++ {
                v45 := GenS4{F0: true, F1: 1}
                _ = fi44
                _ = v45
            }
            switch (v3 * int(v2)) {
            case -287:
                fmt.Println(v1)
            case 9223372036854775807:
                v11 = v11
            }
        } else {
            for fi46 := 0; fi46 < 7; fi46++ {
                if (!(1e10 != v1)) {
                    if (v4 || (v4 && v4)) {
                        v4 = (89 == (v2 << ((374) & 63)))
                    } else {
                        switch "a" {
                        case "":
                            v4 = v4
                        case "foo bar":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        fmt.Println(v4, fi46, v4)
                    }
                    var v47 int64 = 474
                    _ = v47
                } else {
                    v3 -= (v3 % (((fi46 * -875)) | 1))
                    v1 -= (v1 * (3.14159 / ((v1 + v1) - v1)))
                    v4 = ((!(v4 || (!v4))) && false)
                }
                _ = fi46
            }
            for fi48 := 0; fi48 < 12; fi48++ {
                var v49 int = 2147483647
                _ = fi48
                _ = v49
            }
        }
    }
    for fi50 := 0; fi50 < 20; fi50++ {
        _ = len(v5)
        v1 -= ((1e10 - v1) * v1)
        _ = fi50
    }
    v2 -= (v2 ^ v2)
    for fi51 := 0; fi51 < 17; fi51++ {
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
        v52 := GenS3{F0: "foo bar"}
        switch (fi51 << ((1) & 63)) {
        case 32767:
            v53 := GenS5{F0: 4294967295, F1: ((^(v11)) == v11)}
            _ = v53
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi51
        _ = v52
    }
    if false {
        for fi54 := 0; fi54 < 4; fi54++ {
            v3 = 0
            for fi55 := 0; fi55 < 20; fi55++ {
                v3 = (v3 | 4294967295)
                v4 = (!("foo bar" >= "foo bar"))
                _ = fi55
            }
            var v56 int64 = -1
            _ = fi54
            _ = v56
        }
        if v4 {
            var v57 uint64 = 2
            _ = v57
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        if ((!(false && v4)) || (3.14159 == -2.25)) {
            switch (^(v3)) {
            case 0:
                switch "\n\t" {
                case "a":
                    fmt.Println(v4)
                }
            case 1:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 2147483647:
                var v58 bool = v4
                _ = v58
            }
        }
        v59 := GenS6{F0: v2}
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
