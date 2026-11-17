package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 string; F1 bool; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{2147483647, 698, 2147483647, 65535}
    var v3 float32 = -2.25
    for fi4 := 0; fi4 < 12; fi4++ {
        fmt.Println(v3, v3)
        v5 := GenS1{F0: 2147483647, F1: -319}
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 65535, F1: 9223372036854775807}
    switch "\n\t" {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = 0
        }
    case "":
        v7 := GenS2{F0: ("foo bar" + (("hello" + ("" + "")) + "\n\t")), F1: (!false), F2: "hello"}
        _ = v7
    default:
        v8 := GenS1{F0: 147, F1: -1}
        _ = v8
    }
    if (!true) {
        switch 1 {
        case -1:
            if true {
                if ((0 >= -24) || (546 < 42)) {
                    var v9 uint = 723
                    v10 := GenS3{F0: 849}
                    _ = v10
                    _ = v9
                }
                for fi11 := 0; fi11 < 16; fi11++ {
                    v3 = (1.5 / float32(3.14159))
                    var v12 string = "hello"
                    fmt.Println(v3)
                    _ = fi11
                    _ = v12
                }
            }
        case 92:
            if (false && (true || false)) {
                v13 := GenS3{F0: 65535}
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 int8 = 0
                    var v16 uint8 = 1
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v13
            } else {
                v3 += ((float32(((0.1 * float64(v3)) - ((0.1 / (-2.25 / 1.5)) / ((0.1 * float64(v3)) * -2.25)))) * (float32(-2.25) * float32((((1e10 / 3.14159) * (float64((v3 - v3)) + float64(0.1))) + ((0.1 * float64(0.1)) * (1.5 - 3.14159)))))) - v3)
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 string = "a"
                    if (-2147483648 >= 1) {
                        fmt.Println(v18, v18)
                        fmt.Println(v18, v18)
                    } else {
                        if true {
                            var v19 int = fi17
                            if (18446744073709551615 != 0) {
                                var v20 uint = 1
                                _ = len(v2)
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 14; fi21++ {
                                fmt.Println(v19, v3)
                                var v22 float64 = (float64(1.5) - 0.0)
                                var v23 int32 = 42
                                switch "" {
                                case "a":
                                    v24 := GenS4{F0: (v23 == 42)}
                                    _ = v24
                                case "":
                                    v25 := GenS1{F0: 2, F1: 511}
                                    _ = v25
                                }
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            _ = v19
                        } else {
                            if true {
                                var v26 int8 = 127
                                v3 = (v3 - (v3 - v3))
                                for fi27 := 0; fi27 < 11; fi27++ {
                                    if (!(true || false)) {
                                        if (!(!true)) {
                                            var v28 int32 = -2147483648
                                            _ = v28
                                        }
                                        v3 += v3
                                        var v29 uint8 = 255
                                        _ = v29
                                    }
                                    var v30 int64 = -9223372036854775808
                                    if len(v2) > 0 {
                                        v2[0] = 57
                                    }
                                    var v31 uint16 = 2
                                    _ = fi27
                                    _ = v30
                                    _ = v31
                                }
                                _ = v26
                            }
                        }
                    }
                    _ = fi17
                    _ = v18
                }
            }
        default:
            v3 += float32((1.5 - (0.1 / -2.25)))
        }
    } else {
        delete(v1, "\n\t")
    }
    if ("hello" <= "") {
        switch -2147483648 {
        case 381:
            v32 := GenS3{F0: 65535}
            _ = v32
        case -1:
            v33 := GenS2{F0: "foo bar", F1: (42 > 1), F2: "foo bar"}
            _ = v33
        default:
            var v34 uint16 = 42
            _ = v34
        }
        for fi35 := 0; fi35 < 9; fi35++ {
            switch "" {
            case "\n\t":
                var v36 bool = (697 > 24)
                _ = v36
            case "foo bar":
                _ = len(v1)
            case "hello":
                for fi37 := 0; fi37 < 14; fi37++ {
                    if false {
                        if len(v2) > 0 {
                            v2[0] = 2
                        }
                    }
                    _ = fi37
                }
            }
            _ = fi35
        }
        v3 += (((v3 + (float32(((1.5 - (1e10 + (((-2.25 - (1.5 - 3.14159)) * float64(0.1)) + 1e10))) + 1.5)) + (v3 / (v3 + (v3 * v3))))) - v3) / v3)
    } else {
        v38 := v1["\n\t"]
        for fi39 := 0; fi39 < 12; fi39++ {
            if len(v2) > 0 {
                v2[0] = 763
            }
            _ = fi39
        }
        v3 += (1.5 / (v3 * -2.25))
        _ = v38
    }
    v3 -= v3
    v3 = 3.14159
    switch "a" {
    case "foo bar":
        v3 = 0.0
    default:
        v3 = ((v3 + -2.25) / -2.25)
    }
    v2 = append(v2, 2)
    v40 := GenS2{F0: "foo bar", F1: ((true || true) || (1 == 5)), F2: ("a" + "hello")}
    for fi41 := 0; fi41 < 15; fi41++ {
        switch "\n\t" {
        case "a":
            fmt.Println(fi41, fi41)
        case "hello":
            if true {
                var v42 uint8 = 2
                v2 = append(v2, 2)
                var v43 int64 = -9223372036854775808
                _ = v42
                _ = v43
            } else {
                for fi44 := 0; fi44 < 5; fi44++ {
                    if true {
                        var v45 string = "a"
                        v46 := GenS1{F0: 89, F1: 9223372036854775807}
                        var v47 uint16 = 1
                        fmt.Println(fi44)
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    _ = fi44
                }
                switch ("hello" + "hello") {
                case "foo bar":
                    fmt.Println(fi41)
                case "a":
                    switch "\n\t" {
                    case "foo bar":
                        var v48 int64 = -1
                        _ = v48
                    case "hello":
                        fmt.Println(fi41, v3)
                    }
                default:
                    var v49 uint = 0
                    _ = v49
                }
            }
        case "\n\t":
            var v50 int16 = 42
            _ = v50
        default:
            fmt.Println(fi41, fi41)
        }
        v51 := v1[""]
        var v52 uint = 994
        _ = fi41
        _ = v51
        _ = v52
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    for fi53 := 0; fi53 < 4; fi53++ {
        for fi54 := 0; fi54 < 20; fi54++ {
            v3 += v3
            switch "" {
            case "hello":
                v3 = float32((3.14159 / float64(float32(float64(1.5)))))
            }
            _ = fi54
        }
        for fi55 := 0; fi55 < 16; fi55++ {
            v3 += v3
            _ = fi55
        }
        v3 -= 1.5
        _ = fi53
    }
    v56 := GenS5{F0: "hello", F1: -2147483648}
    for fi57 := 0; fi57 < 5; fi57++ {
        switch 893 {
        case 2147483647:
            var v58 float64 = 0.0
            _ = v58
        }
        var v59 int16 = 1000
        if (127 <= 42) {
            for fi60 := 0; fi60 < 7; fi60++ {
                _ = len(v1)
                _ = fi60
            }
        } else {
            if len(v2) > 0 {
                v2[0] = 65535
            }
            if (v59 <= (v59 >> ((410) & 63))) {
                v59 = v59
                if false {
                    delete(v1, "hello")
                }
            } else {
                fmt.Println(v59)
            }
            var v61 int8 = 100
            _ = v61
        }
        _ = fi57
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v56
    _ = v6
}
