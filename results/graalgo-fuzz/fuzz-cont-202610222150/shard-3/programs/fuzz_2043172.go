package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int32 }
type GenS2 struct { F0 int32; F1 int; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint32) uint32 {
    switch ("a" + "") {
    case "hello":
        p0 = 569
    default:
        v6 := GenS1{F0: 255, F1: "a", F2: -2147483648}
        _ = v6
    }
    switch ((("hello" + "\n\t") + "\n\t") + ((("" + "hello") + "") + "a")) {
    case "foo bar":
        p0 = 401
    case "":
        fmt.Println(p0)
    }
    _ = p0
    return (^(p0))
}
func genFn1(p0 uint32) uint16 {
    v7 := []int{2147483647, -2147483648}
    for fi8 := 0; fi8 < 5; fi8++ {
        switch "a" {
        case "":
            switch fi8 {
            case 687:
                switch "\n\t" {
                case "foo bar":
                    var v9 int64 = 0
                    _ = v9
                case "a":
                    for fi10 := 0; fi10 < 11; fi10++ {
                        var v11 int32 = 42
                        var v12 float32 = 3.14159
                        v11 -= v11
                        var v13 uint64 = uint64(p0)
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
            case -605:
                if len(v7) > 0 {
                    v7[0] = fi8
                }
            case 9223372036854775807:
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 int32 = -2147483648
                    _ = fi14
                    _ = v15
                }
            default:
                var v16 int64 = 1
                _ = v16
            }
        default:
            fmt.Println(fi8)
        }
        _ = fi8
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        p0 -= (p0 >> ((2) & 63))
        _ = fi17
    }
    if (float32(-2.25) >= (3.14159 * -2.25)) {
        for fi18 := 0; fi18 < 16; fi18++ {
            fmt.Println(fi18, p0)
            _ = fi18
        }
    } else {
        p0 = genFn0(885)
        if (1 == 128) {
            p0 = p0
        }
    }
    switch -9223372036854775808 {
    case 32768:
        var v19 float32 = float32(0.0)
        _ = v19
    case 9223372036854775807:
        fmt.Println(p0)
    case 128:
        v7 = append(v7, 4294967295)
    default:
        var v20 uint16 = 73
        _ = v20
    }
    _ = p0
    _ = v7
    return 255
}

func main() {
    var v1 uint = 303
    v2 := make(map[int]int)
    var v3 uint64 = 42
    v4 := []uint32{42, 645, 65535, 65535}
    var v5 uint = (v1 / (((v1 | (v1 << (((v1 & v1)) & 63)))) | 1))
    v3 = v3
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v1, v5, v1)
    case "":
        fmt.Println(v5)
    case "\n\t":
        v2[2147483647] = 79
    default:
        for fi21 := 0; fi21 < 20; fi21++ {
            _ = len(v4)
            _ = fi21
        }
    }
    v5 += v1
    switch -381 {
    case -132:
        var v22 string = ("foo bar" + "\n\t")
        _ = v22
    case -9223372036854775808:
        var v23 int16 = -1
        _ = v23
    default:
        switch -1 {
        case 4294967295:
            var v24 int16 = -935
            _ = v24
        case 917:
            var v25 int16 = -32768
            _ = v25
        case 42:
            v26 := v2[4294967295]
            _ = v26
        }
    }
    var v27 float32 = -2.25
    for fi28 := 0; fi28 < 2; fi28++ {
        if false {
            var v29 int32 = -924
            _ = v29
        } else {
            v30 := GenS2{F0: 488, F1: -848, F2: 27}
            for fi31 := 0; fi31 < 17; fi31++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v32 := GenS1{F0: 42, F1: ("" + ("a" + "a")), F2: -673}
                _ = fi31
                _ = v32
            }
            _ = v30
        }
        switch "hello" {
        case "\n\t":
            var v33 uint32 = 301
            _ = v33
        default:
            if true {
                var v34 string = "hello"
                _ = v34
            }
        }
        v35 := GenS3{F0: 244, F1: -1, F2: 2147483647}
        _ = fi28
        _ = v35
    }
    if true {
        for fi36 := 0; fi36 < 6; fi36++ {
            switch (fi36 + fi36) {
            case 128:
                var v37 int16 = 349
                _ = v37
            default:
                var v38 uint8 = 244
                _ = v38
            }
            v39 := GenS3{F0: 0, F1: 0, F2: 4294967295}
            _ = fi36
            _ = v39
        }
    }
    var v40 int32 = -2147483648
    if (true && (true && (-9223372036854775808 <= 2147483647))) {
        var v41 string = ""
        for fi42 := 0; fi42 < 17; fi42++ {
            switch (fi42 / ((32768) | 1)) {
            case 4294967295:
                v43 := GenS2{F0: (^(v40)), F1: -369, F2: 127}
                _ = v43
            }
            _ = fi42
        }
        for fi44 := 0; fi44 < 16; fi44++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch ("foo bar" + v41) {
            case "hello":
                v45 := GenS2{F0: (^(v40)), F1: (fi44 / ((fi44) | 1)), F2: 102}
                _ = v45
            case "":
                v41 = v41
            case "foo bar":
                switch fi44 {
                case 42:
                    v27 = ((((float32(float64((float32((3.14159 / 1e10)) / (float32(0.1) - float32(-2.25))))) + v27) * v27) - ((v27 * 0.0) + v27)) / float32((0.1 / 1.5)))
                case -750:
                    switch fi44 {
                    case 0:
                        for fi46 := 0; fi46 < 7; fi46++ {
                            var v47 uint32 = 1
                            var v48 uint8 = 212
                            delete(v2, 9223372036854775807)
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                    }
                case 9223372036854775807:
                    var v49 uint8 = 0
                    _ = v49
                default:
                    _ = len(v4)
                }
            default:
                for fi50 := 0; fi50 < 10; fi50++ {
                    fmt.Println(v1, v40)
                    var v51 float32 = v27
                    v1 -= v1
                    _ = fi50
                    _ = v51
                }
            }
            _ = fi44
        }
        _ = v41
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (-800 == -32768) {
            for fi52 := 0; fi52 < 15; fi52++ {
                v53 := GenS2{F0: v40, F1: 32767, F2: 42}
                _ = fi52
                _ = v53
            }
            var v54 uint16 = 456
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 13; fi55++ {
        if (true || (!(125 > 100))) {
            if (0 >= 949) {
                switch (fi55 & (fi55 & fi55)) {
                case 884:
                    fmt.Println(v3)
                case -347:
                    fmt.Println(v3, v1)
                case 128:
                    var v56 uint64 = (v3 | v3)
                    _ = v56
                }
            }
            for fi57 := 0; fi57 < 10; fi57++ {
                v27 -= float32(1.5)
                switch ("" + "") {
                case "foo bar":
                    v1 -= 895
                case "":
                    fmt.Println(v3)
                case "\n\t":
                    v5 -= (^(v1))
                default:
                    _ = len(v2)
                }
                delete(v2, (fi55 % ((fi57) | 1)))
                _ = fi57
            }
        } else {
            v4 = append(v4, genFn0(915))
            var v58 uint8 = 156
            switch (fi55 | (fi55 + (fi55 << (((v5 >> ((560) & 63))) & 63)))) {
            case -9223372036854775808:
                var v59 int8 = -41
                _ = v59
            case 4294967295:
                if ((true && true) || true) {
                    var v60 uint8 = (v58 * (v58 + v58))
                    v4 = append(v4, 4294967295)
                    _ = len(v2)
                    _ = v60
                }
            default:
                for fi61 := 0; fi61 < 18; fi61++ {
                    var v62 int = (fi55 >> (((v5 / ((v5) | 1))) & 63))
                    v5 -= 471
                    _ = fi61
                    _ = v62
                }
            }
            _ = v58
        }
        v63 := GenS4{F0: genFn0(989), F1: 761}
        v40 = v40
        _ = fi55
        _ = v63
    }
    v5 -= v1
    var v64 string = (("" + "hello") + ("a" + "foo bar"))
    if len(v4) > 0 {
        v4[0] = genFn0(4294967295)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v64
}
