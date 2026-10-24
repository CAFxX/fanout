package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int32; F2 int16 }
type GenS8 struct { F0 int64; F1 bool; F2 int16 }
type GenS9 struct { F0 string }
func genFn0(p0 bool, p1 uint64) string {
    v5 := GenS1{F0: 4294967295, F1: 128, F2: 40}
    var v6 float64 = 0.1
    v7 := GenS1{F0: 886, F1: 255, F2: -91}
    fmt.Println(v6, p0, v6)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1() int32 {
    if true {
        if (0.0 != 3.14159) {
            switch "hello" {
            case "":
                v8 := []int{32767, 32768, -183, 4294967295}
                _ = v8
            case "hello":
                var v9 int32 = 509
                _ = v9
            case "\n\t":
                switch ("foo bar" + "") {
                case "foo bar":
                    for fi10 := 0; fi10 < 14; fi10++ {
                        var v11 int8 = 101
                        v12 := GenS2{F0: 811}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case "":
                    var v13 bool = ((false || true) && (2 > 127))
                    _ = v13
                case "\n\t":
                    var v14 bool = true
                    _ = v14
                }
            }
        }
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        v16 := []int{(-(fi15)), fi15, fi15}
        _ = fi15
        _ = v16
    }
    var v17 int8 = -48
    var v18 uint64 = 2
    _ = v17
    _ = v18
    return 1
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{(v1 / ((2147483648) | 1)), 42, (v1 & v1), 126}
    v3 := []string{((("foo bar" + "a") + "foo bar") + "foo bar"), "hello", ("a" + "\n\t")}
    v4 := []string{}
    var v19 float64 = 0.0
    var v20 int8 = 1
    v1 -= v1
    v20 -= v20
    if false {
        v21 := GenS2{F0: 471}
        var v22 int = 42
        var v23 uint8 = 255
        _ = v21
        _ = v22
        _ = v23
    }
    if false {
        v24 := GenS2{F0: genFn1()}
        _ = len(v2)
        v25 := GenS3{F0: (v1 | v1), F1: v1}
        _ = v24
        _ = v25
    }
    if (1 <= v20) {
        switch 646 {
        case -333:
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 uint64 = 1
                v28 := GenS4{F0: 65535, F1: fi26}
                _ = fi26
                _ = v27
                _ = v28
            }
        case 19:
            switch -2147483648 {
            case 42:
                if (true || (!true)) {
                    v29 := GenS3{F0: 2147483648, F1: 0}
                    _ = v29
                }
            default:
                v20 += (v20 - v20)
            }
        default:
            switch -378 {
            case -1:
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 float32 = float32(v19)
                    _ = fi30
                    _ = v31
                }
            case 128:
                switch -1 {
                case 761:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 1:
                    fmt.Println(v1, v20)
                }
            }
        }
        _ = len(v2)
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    } else {
        switch -538 {
        case 127:
            if false {
                v19 = float64(((0.1 / 1.0) / (float32(float64(-2.25)) * 1e10)))
            }
        case -1:
            v32 := GenS5{F0: -9223372036854775808, F1: 1}
            _ = v32
        default:
            v3 = append(v3, "")
        }
    }
    if true {
        if true {
            fmt.Println(v1, v19)
            var v33 int = 768
            _ = v33
        } else {
            var v34 int32 = 1
            if (true && (true && (!(true && ((v19 + v19) != -2.25))))) {
                if true {
                    fmt.Println(v1, v1, v1)
                    v19 = (0.0 * 3.14159)
                    v34 += v34
                } else {
                    if ((v1 >> ((2) & 63)) != v1) {
                        var v35 float64 = float64(1e10)
                        v20 = 1
                        _ = v35
                    }
                }
            } else {
                fmt.Println(v1)
                for fi36 := 0; fi36 < 5; fi36++ {
                    v37 := GenS5{F0: 1, F1: (-(fi36))}
                    _ = fi36
                    _ = v37
                }
                v38 := GenS6{F0: 18446744073709551615}
                _ = v38
            }
            _ = v34
        }
        _ = len(v3)
        v39 := GenS7{F0: (^(v20)), F1: -1, F2: 42}
        _ = v39
    }
    var v40 float64 = 0.1
    v40 = (((v19 + 0.1) / 3.14159) - ((3.14159 - ((v19 - 0.1) * float64((-2.25 / 1.5)))) + v40))
    var v41 uint32 = 2147483647
    v19 += 0.1
    if true {
        if false {
            if (0 != 1) {
                fmt.Println(v19)
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
            } else {
                if false {
                    for fi42 := 0; fi42 < 2; fi42++ {
                        var v43 uint16 = 255
                        _ = fi42
                        _ = v43
                    }
                    v44 := GenS8{F0: 9223372036854775807, F1: true, F2: -737}
                    _ = v44
                }
                var v45 uint32 = 42
                _ = v45
            }
        }
        var v46 int8 = 1
        _ = v46
    } else {
        switch (("\n\t" + ((("foo bar" + "\n\t") + ("foo bar" + "")) + "foo bar")) + "") {
        case "a":
            for fi47 := 0; fi47 < 8; fi47++ {
                var v48 bool = false
                if (false || (v48 || ((("foo bar" > ("hello" + "")) && v48) || v48))) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi47
                _ = v48
            }
        case "foo bar":
            switch 42 {
            case 128:
                v40 += (float64(0.0) - (-2.25 - (float64(0.1) - (v19 + 1e10))))
            case 127:
                var v49 uint8 = 0
                _ = v49
            case 1:
                v50 := GenS2{F0: 2147483647}
                _ = v50
            }
        case "":
            v51 := GenS3{F0: v41, F1: 799}
            _ = v51
        default:
            if len(v2) > 0 {
                v2[0] = (v41 << ((65535) & 63))
            }
        }
        v3 = append(v3, "foo bar")
    }
    switch "" {
    case "\n\t":
        if (false && false) {
            var v52 uint64 = 18446744073709551615
            v19 = v19
            _ = v52
        } else {
            v40 += ((v40 + float64(0.1)) + (float64(0.1) / ((v40 * v40) + (-2.25 * v40))))
            fmt.Println(v20, v40)
            var v53 uint = 0
            _ = v53
        }
    case "foo bar":
        v1 = (v41 / (((v41 ^ v41)) | 1))
    default:
        for fi54 := 0; fi54 < 8; fi54++ {
            v41 = v1
            var v55 uint8 = 42
            switch "" {
            case "\n\t":
                v56 := GenS4{F0: 65535, F1: (fi54 + fi54)}
                _ = v56
            case "a":
                fmt.Println(fi54, v1, v20)
            case "":
                var v57 string = "a"
                _ = v57
            }
            _ = fi54
            _ = v55
        }
    }
    if (genFn1() == 2147483647) {
        v58 := GenS4{F0: 764, F1: -103}
        for fi59 := 0; fi59 < 6; fi59++ {
            fmt.Println(v40, v1, v19)
            v41 += 65535
            _ = fi59
        }
        _ = len(v3)
        _ = v58
    }
    var v60 int16 = 0
    if true {
        v20 = -109
    } else {
        switch ("\n\t" + "foo bar") {
        case "hello":
            v3 = append(v3, ("\n\t" + ("foo bar" + "hello")))
        case "":
            var v61 int32 = -14
            _ = v61
        case "foo bar":
            var v62 uint = 0
            _ = v62
        default:
            if (false && true) {
                for fi63 := 0; fi63 < 18; fi63++ {
                    for fi64 := 0; fi64 < 11; fi64++ {
                        v65 := GenS5{F0: 630, F1: (fi64 % (((fi64 - fi63)) | 1))}
                        var v66 int16 = 32767
                        switch fi64 {
                        case -658:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            switch (fi64 ^ fi63) {
                            case -1:
                                if true {
                                    v67 := GenS9{F0: "a"}
                                    _ = v67
                                }
                            default:
                                if (65535 != 183) {
                                    fmt.Println(v66)
                                }
                            }
                        }
                        v68 := GenS3{F0: 65535, F1: 0}
                        _ = fi64
                        _ = v65
                        _ = v66
                        _ = v68
                    }
                    var v69 int = (fi63 >> ((0) & 63))
                    var v70 int = 60
                    _ = fi63
                    _ = v69
                    _ = v70
                }
                if false {
                    var v71 uint32 = v41
                    _ = v71
                }
                fmt.Println(v60, v1)
            }
        }
        for fi72 := 0; fi72 < 17; fi72++ {
            fmt.Println(v1)
            if true {
                for fi73 := 0; fi73 < 11; fi73++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi73
                }
            }
            for fi74 := 0; fi74 < 20; fi74++ {
                switch "a" {
                case "a":
                    var v75 string = "foo bar"
                    _ = v75
                case "":
                    fmt.Println(fi72, v41)
                case "hello":
                    v3 = append(v3, (("" + "a") + "\n\t"))
                default:
                    switch "a" {
                    case "a":
                        fmt.Println(v41, v40)
                    default:
                        fmt.Println(v60, v19)
                    }
                }
                _ = fi74
            }
            _ = fi72
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v60
}
