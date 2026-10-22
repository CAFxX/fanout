package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint16, p1 uint) uint64 {
    var v6 uint16 = p0
    v7 := GenS1{F0: 100, F1: 0, F2: uint64(v6)}
    p0 += 802
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1() int32 {
    var v8 int8 = 126
    for fi9 := 0; fi9 < 20; fi9++ {
        switch "foo bar" {
        case "foo bar":
            var v10 uint32 = 2147483648
            _ = v10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: (v8 + 1), F1: 32767, F2: 18446744073709551615}
            _ = fi11
            _ = v12
        }
        var v13 int = (-(fi9))
        _ = fi9
        _ = v13
    }
    v14 := []int{1, 128, 128, -1}
    _ = v14
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 128
    var v3 uint = 0
    var v4 float32 = -2.25
    v5 := []string{"hello", "a", ""}
    var v15 bool = (65535 == 400)
    v3 += v3
    v16 := GenS2{F0: (v15 || (("a" + ("" + ("a" + ("foo bar" + (("hello" + ("a" + "hello")) + ("" + ("foo bar" + (("a" + ("" + "\n\t")) + ("foo bar" + "hello"))))))))) >= "foo bar")), F1: 1}
    if ((319 <= 32768) && (42 != 65535)) {
        var v17 bool = false
        fmt.Println(v2, v4)
        v17 = (348 >= 42)
        _ = v17
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    switch "foo bar" {
    case "foo bar":
        var v18 uint8 = 235
        _ = v18
    case "\n\t":
        if ((32767 < 1) || (!true)) {
            var v19 int8 = 127
            v2 -= uint8(v3)
            _ = v19
        }
    default:
        if false {
            v4 = v4
            v4 += float32((3.14159 / (float64((v4 - ((v4 + 1e10) * v4))) * 3.14159)))
        } else {
            fmt.Println(v4, v15, v2)
        }
    }
    if (0 < 110) {
        if (1.5 > -2.25) {
            var v20 uint8 = (v2 >> ((65535) & 63))
            v1[42] = ("" + ("" + "foo bar"))
            switch ("foo bar" + "a") {
            case "foo bar":
                v3 += (^(v3))
            default:
                fmt.Println(v2, v4)
            }
            _ = v20
        } else {
            if (v15 || (1000 != -32768)) {
                v21 := GenS2{F0: v15, F1: 393}
                v2 = v2
                _ = v21
            }
            switch -851 {
            case 128:
                var v22 uint64 = 703
                _ = v22
            case 2147483647:
                switch -9223372036854775808 {
                case 1:
                    v2 = (v2 - v2)
                }
            default:
                _ = len(v1)
            }
            switch 2147483647 {
            case 21:
                if (!v15) {
                    fmt.Println(v2)
                }
            case 1:
                fmt.Println(v3, v4, v4)
            }
        }
        v2 = v2
    }
    switch "foo bar" {
    case "a":
        var v23 int8 = 0
        _ = v23
    default:
        if (-265 == -1) {
            for fi24 := 0; fi24 < 14; fi24++ {
                v15 = (!true)
                switch "hello" {
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        var v25 int16 = 42
                        _ = v25
                    case "":
                        delete(v1, 781)
                    case "foo bar":
                        fmt.Println(v15, v3, v2)
                    default:
                        v3 -= 255
                    }
                case "a":
                    switch fi24 {
                    case 42:
                        v4 -= float32(0.1)
                    default:
                        for fi26 := 0; fi26 < 20; fi26++ {
                            var v27 float32 = (((v4 - v4) - ((v4 / 3.14159) + (1e10 * v4))) * (v4 / v4))
                            v2 = v2
                            _ = fi26
                            _ = v27
                        }
                    }
                case "":
                    for fi28 := 0; fi28 < 19; fi28++ {
                        v2 -= v2
                        var v29 uint = (v3 ^ 229)
                        _ = len(v1)
                        _ = fi28
                        _ = v29
                    }
                default:
                    v15 = v15
                }
                _ = fi24
            }
        }
    }
    v2 += v2
    for fi30 := 0; fi30 < 18; fi30++ {
        v3 = (v3 << ((v3) & 63))
        switch ("hello" + "\n\t") {
        case "a":
            if false {
                switch ("\n\t" + "a") {
                case "\n\t":
                    switch "" {
                    case "":
                        var v31 bool = (v3 == (v3 & 829))
                        _ = v31
                    default:
                        if (137 < -108) {
                            fmt.Println(v4, v4)
                        } else {
                            switch ("foo bar" + "a") {
                            case "hello":
                                v4 = v4
                            case "a":
                                var v32 uint16 = 0
                                _ = v32
                            case "\n\t":
                                v33 := v1[106]
                                _ = v33
                            default:
                                var v34 uint8 = (v2 / ((128) | 1))
                                _ = v34
                            }
                            _ = len(v5)
                            v35 := GenS1{F0: -106, F1: -621, F2: 618}
                            _ = v35
                        }
                    }
                case "a":
                    v4 = (v4 + (((1.5 * v4) + (v4 / v4)) - v4))
                case "":
                    var v36 float32 = ((v4 * float32(-2.25)) + 0.0)
                    _ = v36
                default:
                    fmt.Println(v3, fi30)
                }
            }
        case "":
            if (!(1 != 16)) {
                var v37 int = -1
                for fi38 := 0; fi38 < 5; fi38++ {
                    if v15 {
                        v15 = (!(!(!v15)))
                        var v39 uint8 = v2
                        _ = v39
                    } else {
                        v1[42] = ("\n\t" + "hello")
                    }
                    _ = fi38
                }
                v40 := GenS2{F0: (v15 && v15), F1: 766}
                _ = v37
                _ = v40
            } else {
                v41 := GenS3{F0: 0, F1: -751}
                _ = v41
            }
        }
        _ = fi30
    }
    v4 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
