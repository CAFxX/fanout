package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int8; F1 uint32; F2 uint32 }
func genFn0(p0 uint64, p1 uint8, p2 uint64) uint {
    if (14 != 65535) {
        v4 := GenS1{F0: (p1 | (^(p1)))}
        v5 := []int{127, -9223372036854775808, 42, 2147483647}
        fmt.Println(p0, p0)
        _ = v4
        _ = v5
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint = 255
    v2 := []int{}
    v3 := []string{"", "foo bar", "foo bar", ("foo bar" + "hello")}
    if (!(!(false || true))) {
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: 255}
            var v8 uint = (^(v1))
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 int8 = 0
            var v11 uint8 = 2
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 uint32 = 513
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(fi14)
        _ = fi14
    }
    v1 -= 1
    v15 := GenS1{F0: 0}
    fmt.Println(v1)
    var v16 int32 = 1
    for fi17 := 0; fi17 < 10; fi17++ {
        if false {
            var v18 uint = 1
            _ = v18
        } else {
            v16 = 0
            _ = len(v3)
        }
        fmt.Println(v1)
        if true {
            v19 := GenS1{F0: 0}
            fmt.Println(v16)
            if (!(42 != uint64(v1))) {
                v16 = v16
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 float64 = (float64(1.5) * 1e10)
                    if (!false) {
                        v21 -= (v21 + (v21 - float64((0.1 * 1e10))))
                        v1 = (v1 % ((v1) | 1))
                    }
                    _ = fi20
                    _ = v21
                }
            } else {
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 int16 = 1
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            v1 -= (^(v1))
                        default:
                            v16 += v16
                        }
                    }
                    var v24 int = (fi17 - fi17)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v16 -= v16
            }
            _ = v19
        } else {
            var v25 uint32 = 1
            if true {
                if true {
                    v26 := GenS2{F0: 100}
                    for fi27 := 0; fi27 < 16; fi27++ {
                        var v28 int64 = -1
                        if true {
                            if len(v2) > 0 {
                                v2[0] = (fi17 * fi27)
                            }
                            var v29 int64 = (v28 ^ int64(v1))
                            _ = v29
                        }
                        if ("hello" != "a") {
                            var v30 uint16 = 65535
                            _ = v30
                        }
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                } else {
                    v25 += 468
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v2 = append(v2, 32768)
                        fmt.Println(v25, fi17)
                        var v32 int8 = 127
                        _ = fi31
                        _ = v32
                    }
                }
                if true {
                    switch (fi17 >> ((v1) & 63)) {
                    case -2147483648:
                        v1 = genFn0(703, 1, 547)
                    }
                    switch fi17 {
                    case 0:
                        var v33 uint16 = 65535
                        _ = v33
                    default:
                        fmt.Println(v1)
                    }
                    v34 := GenS3{F0: "a"}
                    _ = v34
                } else {
                    var v35 string = ""
                    _ = v35
                }
            }
            for fi36 := 0; fi36 < 18; fi36++ {
                switch fi36 {
                case 42:
                    v16 = (^(v16))
                case -9223372036854775808:
                    switch (^(fi17)) {
                    case 548:
                        switch fi17 {
                        case 128:
                            v37 := GenS1{F0: 127}
                            _ = v37
                        case -208:
                            _ = len(v3)
                        case 4294967295:
                            if (false || (1 != 32767)) {
                                v38 := GenS1{F0: 42}
                                var v39 uint64 = 540
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v38
                                _ = v39
                            }
                        default:
                            v1 = genFn0(2, 128, 2)
                        }
                    case 50:
                        v16 += v16
                    case 127:
                        if len(v3) > 0 {
                            v3[0] = (("foo bar" + "") + "")
                        }
                    default:
                        fmt.Println(fi36, fi17, v1)
                    }
                case 438:
                    fmt.Println(v25, v16, v16)
                default:
                    switch fi36 {
                    case 42:
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v16, fi36, fi17)
                        case "a":
                            v1 = genFn0(836, 255, 171)
                        case "hello":
                            var v40 uint16 = 42
                            _ = v40
                        default:
                            if false {
                                switch (("\n\t" + "a") + ("foo bar" + ("\n\t" + ""))) {
                                case "\n\t":
                                    switch "a" {
                                    case "\n\t":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        var v41 float64 = -2.25
                                        _ = v41
                                    }
                                case "a":
                                    var v42 uint = (v1 - v1)
                                    _ = v42
                                default:
                                    var v43 int16 = 32767
                                    _ = v43
                                }
                                v25 = (v25 / ((v25) | 1))
                            } else {
                                var v44 bool = ((((false || (false && true)) && false) || (v25 == (v25 | (^(v25))))) && false)
                                if false {
                                    if false {
                                        var v45 float64 = float64(3.14159)
                                        _ = v45
                                    } else {
                                        v1 = 1
                                    }
                                }
                                var v46 int32 = int32(v25)
                                _ = v44
                                _ = v46
                            }
                        }
                    case -2147483648:
                        if (1.5 >= 0.0) {
                            v1 = 2
                            var v47 int8 = 100
                            v47 -= (v47 * (v47 >> ((v1) & 63)))
                            _ = v47
                        }
                    default:
                        fmt.Println(v16, fi17)
                    }
                }
                if len(v3) > 0 {
                    v3[0] = ("hello" + "hello")
                }
                v1 = v1
                _ = fi36
            }
            _ = v25
        }
        _ = fi17
    }
    var v48 bool = ((float64(float32(1.5)) * ((0.0 - 1e10) - 3.14159)) >= 0.1)
    switch 127 {
    case -9223372036854775808:
        _ = len(v3)
    }
    switch "foo bar" {
    case "a":
        if (2 < 551) {
            v16 += 2147483647
        }
    case "\n\t":
        v16 += (-(v16))
    }
    v49 := GenS4{F0: v48, F1: 2}
    v48 = (0.1 == -2.25)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 42 {
    case 1:
        v1 += v1
    case 32768:
        v50 := GenS3{F0: "a"}
        _ = v50
    }
    switch 9223372036854775807 {
    case 42:
        if (-5 >= 1) {
            v48 = v48
        } else {
            switch (("a" + "\n\t") + "") {
            case "hello":
                for fi51 := 0; fi51 < 2; fi51++ {
                    v16 += v16
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi51
                }
            case "":
                var v52 uint = 1
                _ = v52
            default:
                _ = len(v2)
            }
            var v53 string = "a"
            if v48 {
                for fi54 := 0; fi54 < 6; fi54++ {
                    for fi55 := 0; fi55 < 6; fi55++ {
                        if true {
                            var v56 string = (("a" + "hello") + "foo bar")
                            if (-2.25 > float32(0.0)) {
                                fmt.Println(v16)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                v57 := GenS5{F0: 466, F1: 1000}
                                v16 = (v16 % (((v16 / ((v16) | 1))) | 1))
                                _ = v57
                            }
                            v58 := GenS5{F0: 1, F1: 42}
                            _ = v56
                            _ = v58
                        } else {
                            fmt.Println(v48)
                            v3 = append(v3, v53)
                        }
                        fmt.Println(v48, v53, fi54)
                        var v59 int16 = 0
                        v60 := GenS6{F0: 38, F1: 42, F2: 2}
                        var v61 uint64 = 18446744073709551615
                        v1 += v1
                        _ = fi55
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                    _ = fi54
                }
            }
            _ = v53
        }
    case 4294967295:
        var v62 int64 = -9223372036854775808
        _ = v62
    default:
        v63 := GenS4{F0: ((32767 <= 2147483647) || v48), F1: -9223372036854775808}
        _ = v63
    }
    v1 = genFn0(346, 1, 257)
    var v64 uint16 = uint16(v1)
    v65 := GenS1{F0: 122}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v48:", v48)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v64
    _ = v65
}
