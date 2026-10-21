package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int64; F1 int64 }
func genFn0(p0 int32, p1 int32, p2 uint16) string {
    p2 += 515
    var v6 int = 343
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "foo bar"
}
func genFn1(p0 int32, p1 uint8, p2 int) uint {
    fmt.Println(p2, p1, p1)
    var v7 int8 = 1
    v8 := []int{32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 float64 = 0.1
    var v2 uint32 = 2
    v3 := []int{42}
    var v4 uint32 = (v2 - (^(v2)))
    var v5 int16 = 42
    fmt.Println(v2)
    switch -9223372036854775808 {
    case 804:
        fmt.Println(v1, v4, v1)
    case -9223372036854775808:
        v5 = (v5 | (v5 >> ((4294967295) & 63)))
    }
    switch -2147483648 {
    case 0:
        v9 := GenS1{F0: (v5 % ((v5) | 1)), F1: 841}
        _ = v9
    case 128:
        var v10 uint16 = 404
        _ = v10
    }
    v4 = 2147483648
    if len(v3) > 0 {
        v3[0] = 1
    }
    if false {
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS1{F0: v5, F1: 42}
            fmt.Println(fi11)
            _ = fi11
            _ = v12
        }
    }
    var v13 bool = false
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 float64 = 0.1
        if false {
            v16 := GenS2{F0: 0, F1: 880}
            for fi17 := 0; fi17 < 16; fi17++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v2 -= 65535
                var v18 uint = 662
                var v19 string = ("a" + "a")
                fmt.Println(v13, v13, fi17)
                _ = len(v3)
                v20 := GenS1{F0: 0, F1: genFn1(42, 1, fi17)}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = v16
        } else {
            switch fi14 {
            case 32768:
                var v21 uint8 = 2
                _ = v21
            }
            v2 = v4
        }
        v22 := GenS2{F0: 242, F1: 18446744073709551615}
        _ = fi14
        _ = v15
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        switch -1 {
        case 32767:
            fmt.Println(v5, v2, fi23)
        case 128:
            var v24 float32 = -2.25
            _ = v24
        case -1:
            v25 := GenS2{F0: -1, F1: 1}
            _ = v25
        default:
            if true {
                var v26 string = "\n\t"
                _ = v26
            } else {
                var v27 uint8 = uint8(v4)
                if len(v3) > 0 {
                    v3[0] = (fi23 - int(v2))
                }
                _ = v27
            }
        }
        if (127 != 42) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi23
    }
    var v28 float64 = 3.14159
    switch 32768 {
    case 32767:
        var v29 float64 = (float64(0.1) * (1e10 + (1e10 * 0.1)))
        _ = v29
    case -978:
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 uint64 = 2
            v31 += v31
            _ = fi30
            _ = v31
        }
    case 423:
        var v32 uint16 = 255
        _ = v32
    default:
        switch "" {
        case "hello":
            fmt.Println(v4)
        case "":
            var v33 float64 = (0.1 - ((((float64(((1.5 / 1.0) / (3.14159 / 0.1))) + v1) + v1) - v1) + (0.1 * (3.14159 * v1))))
            _ = v33
        case "a":
            if true {
                if (true && (v13 || (false || v13))) {
                    switch ("" + (("a" + "foo bar") + "")) {
                    case "hello":
                        for fi34 := 0; fi34 < 5; fi34++ {
                            for fi35 := 0; fi35 < 11; fi35++ {
                                v36 := GenS2{F0: -2147483648, F1: 375}
                                _ = fi35
                                _ = v36
                            }
                            _ = fi34
                        }
                    case "foo bar":
                        v2 = (v4 >> ((733) & 63))
                    }
                    var v37 uint8 = 42
                    if len(v3) > 0 {
                        v3[0] = -1
                    }
                    _ = v37
                }
                if false {
                    v38 := GenS3{F0: -1, F1: 0}
                    v5 = v5
                    _ = v38
                }
                if (0 != 1) {
                    v1 += 0.1
                    v2 = v2
                    if v13 {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v39 int = 32768
                        var v40 uint16 = uint16(v4)
                        _ = v39
                        _ = v40
                    } else {
                        if (45 >= 18446744073709551615) {
                            for fi41 := 0; fi41 < 6; fi41++ {
                                fmt.Println(fi41, v5)
                                v42 := GenS3{F0: 9223372036854775807, F1: 42}
                                if (-1 > 127) {
                                    v43 := GenS3{F0: 367, F1: -1}
                                    v44 := GenS2{F0: 42, F1: 1}
                                    v13 = (160 >= 184)
                                    _ = v43
                                    _ = v44
                                } else {
                                    var v45 string = "\n\t"
                                    _ = v45
                                }
                                _ = fi41
                                _ = v42
                            }
                        } else {
                            var v46 uint8 = 133
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v47 := GenS2{F0: 42, F1: 18446744073709551615}
                            _ = v46
                            _ = v47
                        }
                    }
                } else {
                    v1 += v28
                    v5 -= v5
                }
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v48 int8 = -128
    fmt.Println(v13, v48, v1)
    for fi49 := 0; fi49 < 1; fi49++ {
        var v50 float64 = v1
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v48
    _ = v5
}
