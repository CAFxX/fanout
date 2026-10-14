package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int8 }
func genFn0(p0 string) int16 {
    p0 = p0
    p0 = (p0 + "foo bar")
    _ = p0
    return -345
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    v3 := []string{"hello", "a", "foo bar", ("\n\t" + ((("foo bar" + (("foo bar" + ("a" + "\n\t")) + ((("a" + "\n\t") + "") + (((("hello" + ("a" + "")) + "a") + (("hello" + "a") + (("\n\t" + ("foo bar" + "")) + ""))) + ("hello" + "hello"))))) + "\n\t") + ("hello" + "foo bar")))}
    var v4 uint64 = 2
    var v5 uint32 = 2147483647
    var v6 float64 = ((float64(float32((0.0 + 1.5))) - (1e10 / 1e10)) + (1e10 - (1.5 + 3.14159)))
    var v7 uint16 = 1
    v4 = (v4 >> ((4294967295) & 63))
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint = 18446744073709551615
            _ = fi9
            _ = v10
        }
        switch fi8 {
        case 1:
            switch (fi8 & fi8) {
            case 1:
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
            case 32768:
                switch (fi8 - fi8) {
                case 127:
                    fmt.Println(v5, v5, v6)
                case 803:
                    if (v4 == (^(v4))) {
                        for fi11 := 0; fi11 < 10; fi11++ {
                            v12 := GenS1{F0: 490}
                            v7 = v7
                            v7 += (v7 ^ 423)
                            v13 := GenS1{F0: 2147483647}
                            v5 -= 2147483647
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    } else {
                        switch "a" {
                        case "hello":
                            v6 += 0.1
                        case "foo bar":
                            var v14 uint8 = 128
                            _ = v14
                        case "":
                            v4 = v4
                        default:
                            for fi15 := 0; fi15 < 11; fi15++ {
                                var v16 string = "foo bar"
                                _ = fi15
                                _ = v16
                            }
                        }
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v18 := GenS1{F0: -2147483648}
                            delete(v2, 971)
                            v4 = v4
                            v7 -= 0
                            switch fi17 {
                            case 42:
                                v19 := GenS2{F0: -217, F1: int(v7), F2: uint8(v7)}
                                _ = v19
                            case 726:
                                fmt.Println(v7)
                            }
                            v4 -= v4
                            v1[fi8] = true
                            v6 += v6
                            _ = fi17
                            _ = v18
                        }
                    }
                default:
                    switch 127 {
                    case 4294967295:
                        for fi20 := 0; fi20 < 18; fi20++ {
                            var v21 int8 = -128
                            _ = fi20
                            _ = v21
                        }
                    case 335:
                        var v22 int64 = 42
                        _ = v22
                    default:
                        for fi23 := 0; fi23 < 16; fi23++ {
                            fmt.Println(v5, v7)
                            v24 := GenS3{F0: uint8(v5), F1: v7, F2: 42}
                            _ = fi23
                            _ = v24
                        }
                    }
                }
            case 733:
                var v25 int32 = int32(v4)
                _ = v25
            }
        case -1:
            v7 = v7
        case 32767:
            v6 += (float64(((0.0 / ((1.5 * 1.5) + ((float32(0.1) + (0.1 + ((float32(((0.1 * (v6 * v6)) / v6)) * 0.0) + (((((1.5 / (1e10 - ((1e10 + (float32((v6 - -2.25)) + 0.0)) * float32((0.0 * 0.0))))) * (float32(float64(3.14159)) * (1.5 / 1.0))) * float32(v6)) + (0.0 * -2.25)) - ((1.5 / (-2.25 + (0.0 * ((float32(v6) / 1.5) * -2.25)))) - float32(float64((-2.25 + float32(-2.25))))))))) / ((0.1 / -2.25) + ((float32(1e10) / 1.0) * 0.1))))) + (float32(v6) - (0.0 + (-2.25 - 0.0))))) * ((float64((1.5 * 1.5)) * v6) / v6))
        }
        _ = fi8
    }
    v26 := GenS1{F0: int32(v7)}
    var v27 int32 = 42
    v6 = 0.0
    switch ((("foo bar" + "hello") + "foo bar") + (((("a" + "foo bar") + "hello") + (("a" + "\n\t") + ("" + "a"))) + "a")) {
    case "a":
        for fi28 := 0; fi28 < 1; fi28++ {
            v27 = (v27 % (((-(v27))) | 1))
            if false {
                for fi29 := 0; fi29 < 4; fi29++ {
                    if false {
                        v30 := GenS3{F0: 167, F1: 295, F2: 100}
                        switch 4294967295 {
                        case 32768:
                            _ = len(v1)
                        default:
                            v7 -= v7
                        }
                        _ = v30
                    }
                    _ = fi29
                }
                if false {
                    if ((0.1 / v6) == v6) {
                        fmt.Println(v27, v4, v5)
                    }
                } else {
                    fmt.Println(v4)
                }
            } else {
                for fi31 := 0; fi31 < 19; fi31++ {
                    v6 += ((v6 / 0.1) + v6)
                    if (false || ((v7 / (((v7 ^ (v7 & 65535))) | 1)) == (v7 << ((2) & 63)))) {
                        var v32 uint16 = uint16(v5)
                        var v33 uint16 = (^(v7))
                        var v34 int16 = -1
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                }
                if false {
                    switch int(v7) {
                    case -2147483648:
                        v27 += 42
                    case 221:
                        var v35 string = "\n\t"
                        _ = v35
                    default:
                        var v36 int16 = genFn0("a")
                        _ = v36
                    }
                    if true {
                        var v37 int16 = -344
                        fmt.Println(fi28, v4, v37)
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 17; fi38++ {
                            v5 -= v5
                            v39 := v2[960]
                            _ = fi38
                            _ = v39
                        }
                        v4 = (v4 >> ((65535) & 63))
                        if (65535 >= 0) {
                            var v40 int16 = -32768
                            switch ("\n\t" + ("" + ("" + ""))) {
                            case "a":
                                v27 = int32(v7)
                            case "hello":
                                v6 -= v6
                            case "":
                                v3 = append(v3, ("hello" + "foo bar"))
                            }
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                            _ = v40
                        }
                    }
                    switch "\n\t" {
                    case "a":
                        if true {
                            for fi41 := 0; fi41 < 2; fi41++ {
                                fmt.Println(v5, fi41, v27)
                                _ = fi41
                            }
                            fmt.Println(v7, v4, fi28)
                        }
                    default:
                        switch 127 {
                        case -975:
                            for fi42 := 0; fi42 < 3; fi42++ {
                                fmt.Println(fi42)
                                var v43 bool = (!(v7 > 2))
                                switch "\n\t" {
                                case "":
                                    v44 := GenS1{F0: (v27 % ((v27) | 1))}
                                    _ = v44
                                case "a":
                                    var v45 float32 = 3.14159
                                    var v46 string = ""
                                    _ = v45
                                    _ = v46
                                default:
                                    fmt.Println(v7, v27)
                                }
                                _ = fi42
                                _ = v43
                            }
                        case -1:
                            if (-9223372036854775808 <= 42) {
                                var v47 int32 = v27
                                _ = v47
                            }
                        default:
                            var v48 int64 = -1
                            _ = v48
                        }
                    }
                }
                fmt.Println(v7, v7)
            }
            _ = fi28
        }
    default:
        v49 := GenS1{F0: v27}
        _ = v49
    }
    v50 := GenS3{F0: 255, F1: uint16(v4), F2: -1}
    v51 := v1[456]
    var v52 uint32 = v5
    var v53 int32 = -1
    v54 := GenS2{F0: -1, F1: -9223372036854775808, F2: 233}
    if ((0 != v5) || (-1 > genFn0(("" + "a")))) {
        _ = len(v3)
    } else {
        var v55 string = (("foo bar" + "") + "a")
        fmt.Println(v27, v27, v6)
        var v56 int = 1
        _ = v55
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v7
}
