package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 uint8; F1 int8; F2 bool }
type GenS3 struct { F0 int64; F1 uint8; F2 bool }
func genFn0(p0 uint64, p1 uint8, p2 string) int8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := []int{}
            if true {
                v9 := GenS1{F0: 2, F1: true, F2: 1000}
                p1 += uint8(p0)
                _ = v9
            } else {
                v10 := GenS2{F0: p1, F1: 127, F2: ("foo bar" >= ("hello" + p2))}
                var v11 int64 = 42
                switch fi6 {
                case 860:
                    var v12 int64 = v11
                    _ = v12
                case 878:
                    p0 = (^(p0))
                case -587:
                    v13 := GenS1{F0: 18446744073709551615, F1: (-971 != 991), F2: -32768}
                    _ = v13
                }
                _ = v10
                _ = v11
            }
            p2 = p2
            _ = fi7
            _ = v8
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            fmt.Println(fi14, fi14)
            var v15 string = p2
            v16 := []int{}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi6
    }
    if ((p2 + ("a" + "")) < p2) {
        for fi17 := 0; fi17 < 3; fi17++ {
            if false {
                v18 := GenS1{F0: 824, F1: (181 < -1), F2: 0}
                v19 := GenS3{F0: 596, F1: (p1 + 128), F2: ((false || false) && false)}
                _ = v18
                _ = v19
            } else {
                var v20 int = 4294967295
                v20 = (v20 ^ v20)
                _ = v20
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                fmt.Println(p1)
                var v22 float32 = -2.25
                switch (p2 + (p2 + (("" + p2) + ""))) {
                case "hello":
                    if (-32768 == 1) {
                        p0 += p0
                    } else {
                        var v23 string = "\n\t"
                        v24 := []int{fi21, -9223372036854775808, 4294967295}
                        _ = v23
                        _ = v24
                    }
                    switch p2 {
                    case "\n\t":
                        if false {
                            v25 := GenS2{F0: p1, F1: -128, F2: true}
                            _ = v25
                        } else {
                            var v26 float64 = 3.14159
                            fmt.Println(p1)
                            switch p2 {
                            case "foo bar":
                                var v27 int = (-(fi21))
                                _ = v27
                            default:
                                var v28 string = (((p2 + p2) + p2) + "a")
                                _ = v28
                            }
                            _ = v26
                        }
                    default:
                        if false {
                            p2 = p2
                        } else {
                            var v29 bool = (0.0 > float32((((-2.25 - 1.5) - (1e10 + 3.14159)) - ((1.5 * 0.0) / 1.0))))
                            _ = v29
                        }
                    }
                case "foo bar":
                    fmt.Println(p0, fi17)
                case "\n\t":
                    v30 := []int{-78, fi21, fi17, 32768}
                    _ = v30
                default:
                    v31 := []int{(fi21 & (fi21 / ((127) | 1)))}
                    _ = v31
                }
                v22 = ((0.0 * 0.0) + float32((0.1 * 1.5)))
                _ = fi21
                _ = v22
            }
            _ = fi17
        }
    } else {
        p1 += 42
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        for fi33 := 0; fi33 < 2; fi33++ {
            v34 := GenS1{F0: 700, F1: (((!false) && true) && (0.0 >= (float32(-2.25) - 3.14159))), F2: 0}
            var v35 uint = 42
            switch fi32 {
            case 652:
                var v36 int16 = 32767
                _ = v36
            case 0:
                v37 := GenS3{F0: 9223372036854775807, F1: p1, F2: false}
                _ = v37
            }
            var v38 int = fi33
            _ = fi33
            _ = v34
            _ = v35
            _ = v38
        }
        fmt.Println(p1, p2, fi32)
        _ = fi32
    }
    v39 := []int{}
    if len(v39) > 0 {
        v39[0] = -1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return -128
}
func genFn1(p0 int64, p1 int16) int {
    p1 -= 677
    v40 := []int{1, 501}
    _ = p0
    _ = p1
    _ = v40
    return int(p1)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    v3 := make(map[uint32]string)
    var v4 int8 = 1
    var v5 uint32 = 65535
    if ((!(((float64(float32(1.5)) - -2.25) * ((((0.0 / 1.0) + (float64(((1.5 / (0.1 + 1.5)) - (3.14159 / 1e10))) - (0.0 + float64(0.0)))) - 0.0) - 0.0)) != (0.1 + (3.14159 + (((3.14159 + 0.1) / (((3.14159 / (float64(float32(float64((0.1 - 3.14159)))) * float64(0.1))) + float64(1e10)) - (1e10 / float64(1e10)))) / (((-2.25 / ((1e10 + 1e10) - 3.14159)) * 0.1) / 0.1)))))) && false) {
        for fi41 := 0; fi41 < 17; fi41++ {
            if len(v2) > 0 {
                v2[0] = (("hello" + "a") + (("hello" + (("" + "a") + ("" + "\n\t"))) + ("foo bar" + "\n\t")))
            }
            _ = fi41
        }
        var v42 int32 = -2147483648
        fmt.Println(v5, v42)
        _ = v42
    }
    var v43 int16 = -32768
    for fi44 := 0; fi44 < 15; fi44++ {
        v43 = v43
        if (3.14159 > 3.14159) {
            v45 := GenS3{F0: 0, F1: uint8(v5), F2: (18446744073709551615 == 4294967295)}
            if true {
                var v46 uint32 = 461
                switch ((("" + (("" + "hello") + ((("" + ("hello" + "a")) + "foo bar") + "hello"))) + ("foo bar" + "")) + ("a" + "hello")) {
                case "":
                    var v47 float32 = -2.25
                    _ = v47
                default:
                    var v48 int = (fi44 + fi44)
                    _ = v48
                }
                _ = v46
            } else {
                var v49 int8 = 127
                _ = v49
            }
            _ = v45
        } else {
            fmt.Println(v4, v43, fi44)
            v5 = (v5 - v5)
        }
        _ = fi44
    }
    for fi50 := 0; fi50 < 3; fi50++ {
        v4 = (v4 >> ((255) & 63))
        if false {
            var v51 int8 = v4
            for fi52 := 0; fi52 < 18; fi52++ {
                switch genFn1(int64(fi52), -1) {
                case 32767:
                    v43 -= (v43 | v43)
                case 2147483647:
                    var v53 int = (fi50 - fi50)
                    _ = v53
                case 0:
                    fmt.Println(v43, fi52)
                default:
                    fmt.Println(fi52)
                }
                _ = fi52
            }
            var v54 int32 = 42
            _ = v51
            _ = v54
        }
        _ = fi50
    }
    v2 = append(v2, "hello")
    v55 := GenS3{F0: 1, F1: 128, F2: (1.5 != ((0.0 + 1e10) * float64(-2.25)))}
    for fi56 := 0; fi56 < 16; fi56++ {
        v4 += genFn0(628, 37, "hello")
        var v57 int64 = 941
        _ = fi56
        _ = v57
    }
    var v58 uint16 = 32767
    fmt.Println(v4, v4)
    v58 = 1
    var v59 int = 128
    var v60 uint16 = v58
    if true {
        if false {
            v60 = 882
        }
    }
    var v61 float32 = 0.0
    v61 += (0.0 / v61)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v55
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
