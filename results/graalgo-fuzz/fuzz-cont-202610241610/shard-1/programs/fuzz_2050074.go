package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int8; F1 int32; F2 int8 }
func genFn0(p0 uint, p1 uint8, p2 int32) int64 {
    switch "hello" {
    case "foo bar":
        fmt.Println(p2)
    case "a":
        var v6 int16 = 1
        _ = v6
    default:
        v7 := []int{}
        _ = v7
    }
    p0 += (p0 ^ (p0 * (p0 + p0)))
    var v8 float64 = float64((1.5 / -2.25))
    var v9 int8 = -102
    for fi10 := 0; fi10 < 11; fi10++ {
        if false {
            var v11 int32 = p2
            _ = v11
        }
        switch "" {
        case "":
            if true {
                var v12 float32 = 3.14159
                for fi13 := 0; fi13 < 1; fi13++ {
                    v8 = v8
                    var v14 uint16 = 1
                    _ = fi13
                    _ = v14
                }
                p0 = (p0 & p0)
                _ = v12
            } else {
                v15 := []int{-596, (fi10 | fi10)}
                if (fi10 != -9223372036854775808) {
                    switch ("foo bar" + ("foo bar" + "a")) {
                    case "":
                        fmt.Println(fi10)
                    default:
                        var v16 uint64 = 4294967295
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 17; fi17++ {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        if false {
                            if len(v15) > 0 {
                                v15[0] = -2147483648
                            }
                        } else {
                            v18 := GenS1{F0: 2147483648}
                            _ = v18
                        }
                        if (((v9 - (v9 >> ((p0) & 63))) == (v9 * (-(v9)))) || true) {
                            _ = len(v15)
                        } else {
                            v8 += 1.5
                            v9 = v9
                            p0 = (^(p0))
                        }
                        _ = fi17
                    }
                } else {
                    if (v9 <= v9) {
                        v19 := GenS1{F0: 2147483648}
                        _ = v19
                    } else {
                        p1 = (^(p1))
                        if (!true) {
                            var v20 bool = true
                            var v21 int = (^(fi10))
                            _ = v20
                            _ = v21
                        }
                    }
                    var v22 float32 = 0.0
                    _ = v22
                }
                _ = v15
            }
        case "hello":
            var v23 int8 = v9
            _ = v23
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -245
}

func main() {
    var v1 float64 = (-2.25 * ((float64(3.14159) * ((0.0 / 1.0) - float64((1e10 * 0.1)))) - 1.5))
    var v2 bool = false
    var v3 uint = 65535
    var v4 int64 = -9223372036854775808
    var v5 int32 = 1
    if v2 {
        fmt.Println(v1, v2)
    } else {
        v24 := []int{-75, 9223372036854775807}
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        fmt.Println(v4)
        fmt.Println(fi25, v3)
        _ = fi25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 float32 = 1e10
        _ = fi26
        _ = v27
    }
    if (!((5 <= 255) && (-2.25 > (3.14159 + -2.25)))) {
        for fi28 := 0; fi28 < 19; fi28++ {
            if v2 {
                for fi29 := 0; fi29 < 10; fi29++ {
                    var v30 float32 = (((0.1 * 0.0) * (3.14159 * (0.1 / float32(v1)))) + float32(v1))
                    switch fi28 {
                    case 803:
                        if v2 {
                            v1 = (1.5 + v1)
                            var v31 int = fi28
                            _ = v31
                        }
                    case 128:
                        v32 := GenS2{F0: (!(18446744073709551615 < 741)), F1: (true || (((false && (42 < 58)) && (!v2)) && v2)), F2: 255}
                        _ = v32
                    case -10:
                        v1 = v1
                        v33 := GenS2{F0: v2, F1: (-919 < 32767), F2: 42}
                        _ = v33
                    }
                    _ = fi29
                    _ = v30
                }
                var v34 int8 = 25
                _ = v34
            } else {
                v35 := GenS3{F0: 0, F1: 0, F2: 2}
                if false {
                    v3 = 0
                    v1 = float64((-2.25 / 3.14159))
                    v2 = v2
                }
                _ = v35
            }
            switch (fi28 / ((fi28) | 1)) {
            case 127:
                fmt.Println(v2, v2, v4)
            case 371:
                v36 := GenS1{F0: 4294967295}
                _ = v36
            }
            _ = fi28
        }
    } else {
        v3 = v3
        v37 := GenS4{F0: (v5 | (v5 / ((580) | 1))), F1: v3}
        _ = v37
    }
    var v38 uint = 1
    v5 -= v5
    switch (("hello" + "\n\t") + ("foo bar" + "")) {
    case "\n\t":
        if true {
            switch 0 {
            case -404:
                v4 = 9223372036854775807
            case -9223372036854775808:
                v2 = ((float32(v1) / 1e10) > -2.25)
            }
        }
    default:
        v39 := GenS4{F0: v5, F1: v38}
        _ = v39
    }
    v40 := []int{}
    for fi41 := 0; fi41 < 11; fi41++ {
        _ = len(v40)
        _ = fi41
    }
    var v42 int = -2147483648
    for fi43 := 0; fi43 < 8; fi43++ {
        switch fi43 {
        case -230:
            if false {
                var v44 uint64 = 821
                v42 += v42
                _ = v44
            }
        case -567:
            if v2 {
                v42 += (v42 ^ (fi43 * fi43))
            }
        default:
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
        }
        if false {
            if (100 <= -1) {
                if (4294967295 < 4294967295) {
                    for fi45 := 0; fi45 < 14; fi45++ {
                        var v46 uint32 = 1
                        v47 := GenS1{F0: (v46 * v46)}
                        v40 = append(v40, (v42 << ((255) & 63)))
                        if true {
                            var v48 uint64 = 240
                            var v49 float32 = (1e10 - 1e10)
                            v50 := GenS2{F0: v2, F1: (2 <= uint8(v48)), F2: 566}
                            _ = v48
                            _ = v49
                            _ = v50
                        } else {
                            var v51 string = "a"
                            v51 = v51
                            var v52 string = (("hello" + v51) + "a")
                            _ = v51
                            _ = v52
                        }
                        if false {
                            if (!((-2.25 + 3.14159) >= ((0.1 / 1.5) - (1.5 / (float32(v1) * 1e10))))) {
                                v42 += (fi45 >> (((v3 + v38)) & 63))
                                if len(v40) > 0 {
                                    v40[0] = (^(fi43))
                                }
                            }
                            v4 -= genFn0((v38 & 0), 255, v5)
                        } else {
                            var v53 uint = v38
                            var v54 uint64 = 2
                            _ = v53
                            _ = v54
                        }
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    fmt.Println(v2, fi43, v38)
                }
            }
            v5 = (v5 ^ (v5 + v5))
        }
        _ = fi43
    }
    if ((!((1 >= 1000) || v2)) && v2) {
        for fi55 := 0; fi55 < 5; fi55++ {
            for fi56 := 0; fi56 < 20; fi56++ {
                v57 := GenS2{F0: (int32(v42) == (v5 + v5)), F1: v2, F2: 42}
                v58 := GenS3{F0: 255, F1: 42, F2: 42}
                switch ("" + (("\n\t" + "\n\t") + "")) {
                case "foo bar":
                    switch "\n\t" {
                    case "hello":
                        var v59 string = ("" + (("hello" + ("" + "a")) + "hello"))
                        _ = v59
                    case "foo bar":
                        var v60 uint32 = uint32(v3)
                        _ = v60
                    }
                case "a":
                    if false {
                        v1 += ((v1 / v1) / v1)
                        if ((v5 - 0) >= (v5 | int32(fi56))) {
                            v61 := GenS2{F0: false, F1: ((v2 && v2) || v2), F2: 65535}
                            _ = v61
                        }
                    }
                    v62 := GenS5{F0: 70, F1: v5, F2: -117}
                    _ = v62
                }
                if len(v40) > 0 {
                    v40[0] = fi55
                }
                fmt.Println(fi55)
                _ = fi56
                _ = v57
                _ = v58
            }
            _ = fi55
        }
        var v63 uint16 = 858
        _ = v63
    }
    var v64 int = int(v38)
    v2 = (951 <= 596)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v64
}
