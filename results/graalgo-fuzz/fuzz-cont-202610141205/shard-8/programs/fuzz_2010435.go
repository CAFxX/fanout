package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS7 struct { F0 uint; F1 uint }

func main() {
    v1 := make(map[int]string)
    v2 := []string{(("foo bar" + (("hello" + ("hello" + "a")) + ("foo bar" + "\n\t"))) + ("a" + "")), (((("" + "\n\t") + "\n\t") + "hello") + "\n\t"), "foo bar", ""}
    v3 := make(map[int]bool)
    var v4 uint16 = 2
    var v5 int32 = 247
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: -878}
        switch "hello" {
        case "":
            if ((!true) || false) {
                var v9 float64 = 1.5
                var v10 bool = (v5 >= 224)
                var v11 int64 = -286
                _ = v10
                _ = v11
                _ = v9
            } else {
                v12 := GenS1{F0: int16(v4)}
                for fi13 := 0; fi13 < 11; fi13++ {
                    v4 = v4
                    _ = fi13
                }
                _ = v12
            }
        case "hello":
            if (false || (42 >= 255)) {
                v5 = (v5 ^ -1)
                v4 = 2
            }
        case "a":
            fmt.Println(fi7)
        default:
            var v14 int8 = 1
            _ = v14
        }
        v15 := v3[fi7]
        _ = fi7
        _ = v15
        _ = v8
    }
    switch -401 {
    case 91:
        if ((false || true) || (((((1.5 / (1.5 + 1e10)) - 0.0) / (((0.1 / 3.14159) + 1.5) * 1e10)) < 1.5) && false)) {
            v2 = append(v2, "\n\t")
            switch "" {
            case "foo bar":
                var v16 float64 = 1e10
                _ = v16
            }
            if (!(1000 <= -32768)) {
                var v17 float32 = (3.14159 * 1e10)
                switch "" {
                case "foo bar":
                    if false {
                        v18 := v6["a"]
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 1; fi19++ {
                            var v20 float64 = (1e10 - 1.5)
                            _ = fi19
                            _ = v20
                        }
                        var v21 int8 = -77
                        _ = v21
                    }
                case "\n\t":
                    fmt.Println(v5, v5)
                }
                v4 = (v4 * (v4 / ((v4) | 1)))
                _ = v17
            } else {
                var v22 float32 = 1e10
                _ = v22
            }
        } else {
            var v23 int = 949
            v4 -= 2
            _ = v23
        }
    case 0:
        switch 128 {
        case 127:
            var v24 uint = 4294967295
            _ = v24
        default:
            _ = len(v6)
        }
    default:
        switch "foo bar" {
        case "\n\t":
            switch ("a" + "") {
            case "hello":
                v5 += v5
            case "a":
                var v25 uint = 2
                _ = v25
            default:
                var v26 float32 = 0.1
                _ = v26
            }
        default:
            v27 := GenS2{F0: 1, F1: 4294967295, F2: 42}
            _ = v27
        }
    }
    v4 += v4
    var v28 int8 = 121
    v29 := v6[("foo bar" + "")]
    switch 2147483647 {
    case 42:
        var v30 int16 = 32767
        _ = v30
    }
    if true {
        var v31 int16 = 32767
        _ = v31
    }
    v32 := GenS2{F0: 4294967295, F1: 42, F2: (v4 & 32767)}
    switch "foo bar" {
    case "hello":
        switch -407 {
        case -2147483648:
            var v33 uint8 = 1
            _ = v33
        case 128:
            v5 -= v5
        case 32768:
            v34 := GenS3{F0: (v28 & v28), F1: 42}
            _ = v34
        default:
            var v35 int8 = (v28 << ((0) & 63))
            _ = v35
        }
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v5)
    case "foo bar":
        fmt.Println(v4, v29, v5)
    }
    v4 = v4
    if (!(0 > 42)) {
        if v29 {
            v5 += (v5 - (v5 * 42))
        } else {
            var v36 int64 = -1
            _ = v36
        }
        switch "\n\t" {
        case "\n\t":
            var v37 float64 = (0.0 - float64((1.5 / 1.0)))
            _ = v37
        case "hello":
            if (209 == 69) {
                var v38 bool = false
                v39 := GenS1{F0: 1000}
                var v40 uint8 = 149
                _ = v38
                _ = v39
                _ = v40
            } else {
                for fi41 := 0; fi41 < 17; fi41++ {
                    for fi42 := 0; fi42 < 6; fi42++ {
                        var v43 uint64 = 42
                        var v44 uint32 = 2147483648
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    switch (fi41 / (((fi41 / ((810) | 1))) | 1)) {
                    case 4294967295:
                        var v45 uint = 2
                        _ = v45
                    case 2147483647:
                        v28 -= v28
                    default:
                        var v46 int8 = v28
                        _ = v46
                    }
                    v4 += v4
                    _ = fi41
                }
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
            }
        case "":
            if v29 {
                var v47 int32 = v5
                var v48 float32 = (0.1 + 0.0)
                var v49 int64 = 0
                _ = v47
                _ = v48
                _ = v49
            }
        default:
            var v50 string = "foo bar"
            _ = v50
        }
        for fi51 := 0; fi51 < 11; fi51++ {
            var v52 bool = (!(float64((((1.5 * (3.14159 / (float32((1e10 + float64(-2.25))) + 0.0))) - (3.14159 * ((1e10 * (float32((float64((3.14159 * 3.14159)) - 0.1)) / (((0.0 - (0.0 / ((1.5 - -2.25) + (1.5 + 0.0)))) - -2.25) / -2.25))) * 0.1))) + 3.14159)) < 1e10))
            v28 = 1
            v28 = (^(v28))
            _ = fi51
            _ = v52
        }
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        v54 := GenS4{F0: 2, F1: v4, F2: (v28 % (((v28 - (v28 + v28))) | 1))}
        _ = fi53
        _ = v54
    }
    if (405 == uint(v4)) {
        for fi55 := 0; fi55 < 10; fi55++ {
            fmt.Println(v28, v28)
            for fi56 := 0; fi56 < 6; fi56++ {
                var v57 int32 = 2147483647
                fmt.Println(v29, fi55)
                switch (((("" + "a") + "") + "a") + "hello") {
                case "":
                    var v58 uint8 = 255
                    _ = v58
                case "hello":
                    v1[(fi56 << ((146) & 63))] = "foo bar"
                case "foo bar":
                    fmt.Println(fi55, v29, v4)
                default:
                    v5 = 24
                }
                var v59 int = (fi56 - fi55)
                _ = fi56
                _ = v57
                _ = v59
            }
            switch "a" {
            case "foo bar":
                v5 = v5
            }
            _ = fi55
        }
        if (-2.25 < ((0.1 - 1e10) + 0.0)) {
            var v60 int32 = v5
            _ = v60
        } else {
            switch 1 {
            case 481:
                if (0.0 > float32(0.1)) {
                    v61 := v3[4294967295]
                    if (!v61) {
                        v62 := GenS5{F0: 2147483648, F1: 882, F2: 65535}
                        v2 = append(v2, ("a" + ("foo bar" + (("hello" + ("a" + "a")) + (("a" + "\n\t") + "\n\t")))))
                        _ = v62
                    } else {
                        v63 := GenS6{F0: 174, F1: 65535, F2: 1}
                        switch "" {
                        case "":
                            fmt.Println(v61, v29)
                        case "foo bar":
                            var v64 int = -708
                            _ = v64
                        default:
                            var v65 int = 140
                            _ = v65
                        }
                        v5 = v5
                        _ = v63
                    }
                    _ = v61
                } else {
                    switch "" {
                    case "a":
                        if (1e10 <= (float32(0.0) - (float32(3.14159) + ((1e10 + -2.25) / (3.14159 + 0.1))))) {
                            for fi66 := 0; fi66 < 7; fi66++ {
                                v67 := GenS7{F0: 42, F1: 2}
                                v68 := GenS7{F0: 4294967295, F1: 4294967295}
                                _ = fi66
                                _ = v67
                                _ = v68
                            }
                            if (0.1 < (-2.25 * -2.25)) {
                                var v69 float64 = (float64(3.14159) - ((float64(((0.1 - 0.0) * -2.25)) - float64(0.1)) / 3.14159))
                                fmt.Println(v29, v5)
                                _ = v69
                            }
                        }
                    }
                    var v70 float64 = 0.0
                    if (65535 > 0) {
                        var v71 int16 = -1
                        _ = v71
                    }
                    _ = v70
                }
            default:
                fmt.Println(v28, v4, v29)
            }
            switch 32767 {
            case 0:
                v29 = true
            case -901:
                if v29 {
                    fmt.Println(v28, v28)
                    var v72 uint = 2
                    _ = v72
                }
            case -310:
                v2 = append(v2, "hello")
            }
            v5 = v5
        }
        var v73 uint64 = 206
        _ = v73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
