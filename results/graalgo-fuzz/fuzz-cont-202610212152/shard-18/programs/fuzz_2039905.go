package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0(p0 int16) int32 {
    p0 = p0
    v5 := GenS1{F0: (p0 << ((2) & 63))}
    p0 = -412
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := []int{fi6}
            v9 := GenS1{F0: -87}
            v10 := GenS1{F0: p0}
            fmt.Println(p0, p0, p0)
            v11 := GenS2{F0: "\n\t", F1: "\n\t"}
            _ = fi7
            _ = v10
            _ = v11
            _ = v8
            _ = v9
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            p0 = p0
            _ = fi12
        }
        switch fi6 {
        case 9223372036854775807:
            p0 -= 822
        default:
            var v13 uint32 = 681
            _ = v13
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return 1
}
func genFn1() uint8 {
    var v14 float32 = ((((1e10 - 3.14159) - 3.14159) + 3.14159) - (float32(0.1) - (float32((3.14159 + 1e10)) + -2.25)))
    fmt.Println(v14)
    _ = v14
    return 128
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 838
    var v3 bool = true
    var v4 bool = v3
    v4 = v3
    v4 = (6 <= 0)
    var v15 bool = v3
    if (0 == 98) {
        switch "a" {
        case "foo bar":
            v4 = true
        case "hello":
            if v4 {
                var v16 int16 = 1
                fmt.Println(v3, v16)
                for fi17 := 0; fi17 < 4; fi17++ {
                    v16 += (v16 >> ((657) & 63))
                    var v18 uint16 = 670
                    v19 := []int{(fi17 / (((fi17 + fi17)) | 1)), fi17, 962, fi17}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            } else {
                switch ((("hello" + ("hello" + "a")) + (("" + "") + "\n\t")) + ((("a" + ("foo bar" + (("" + "hello") + "hello"))) + "a") + "\n\t")) {
                case "a":
                    delete(v1, 42)
                }
                if (18446744073709551615 > 564) {
                    var v20 string = "hello"
                    _ = v20
                }
            }
        case "a":
            var v21 uint8 = 255
            _ = v21
        default:
            switch -9223372036854775808 {
            case -9223372036854775808:
                var v22 int = -9223372036854775808
                _ = v22
            }
        }
        var v23 float32 = (3.14159 - (3.14159 - (0.0 / float32((float64(((float32(float64(0.1)) - -2.25) - 1e10)) - 0.0)))))
        switch 9223372036854775807 {
        case -9223372036854775808:
            if true {
                switch -436 {
                case 802:
                    v24 := GenS3{F0: -128, F1: 44}
                    _ = v24
                case 0:
                    var v25 string = ("hello" + ("a" + ("foo bar" + (("a" + ("" + ("\n\t" + ((("foo bar" + "") + ("" + "")) + "")))) + ""))))
                    _ = v25
                case 32768:
                    v23 = v23
                default:
                    v26 := v1[32767]
                    _ = v26
                }
                var v27 float64 = (1.5 - -2.25)
                _ = v27
            } else {
                for fi28 := 0; fi28 < 2; fi28++ {
                    v15 = (v3 && true)
                    switch fi28 {
                    case 2147483647:
                        if v4 {
                            var v29 int16 = -231
                            v30 := GenS2{F0: (((("\n\t" + "hello") + "a") + "foo bar") + (("\n\t" + "a") + ("foo bar" + "\n\t"))), F1: ((("\n\t" + ("hello" + ("foo bar" + "hello"))) + ("foo bar" + ((("a" + "hello") + "") + "foo bar"))) + "a")}
                            _ = v29
                            _ = v30
                        }
                    case 32768:
                        if false {
                            v3 = (((v2 >= v2) && (-1 != 1000)) || true)
                            v31 := GenS3{F0: -1, F1: 2}
                            _ = v31
                        } else {
                            v32 := GenS1{F0: 0}
                            if false {
                                v33 := []int{}
                                v23 = (float32(float64(v23)) - (((v23 * v23) + (v23 / (v23 - v23))) - v23))
                                _ = v33
                            }
                            switch 128 {
                            case -792:
                                for fi34 := 0; fi34 < 12; fi34++ {
                                    fmt.Println(fi28)
                                    var v35 int64 = (^(v2))
                                    _ = fi34
                                    _ = v35
                                }
                            case 0:
                                for fi36 := 0; fi36 < 19; fi36++ {
                                    v37 := []int{(fi28 | fi36), (fi28 * 127)}
                                    _ = fi36
                                    _ = v37
                                }
                            default:
                                var v38 string = "hello"
                                _ = v38
                            }
                            _ = v32
                        }
                    case 32767:
                        switch 2147483647 {
                        case 9223372036854775807:
                            var v39 int = -209
                            _ = v39
                        }
                    default:
                        var v40 bool = (!v3)
                        _ = v40
                    }
                    v2 -= v2
                    _ = fi28
                }
                var v41 int8 = 0
                var v42 int64 = v2
                _ = v41
                _ = v42
            }
        case 555:
            fmt.Println(v2, v3)
        case -34:
            switch "" {
            case "foo bar":
                v43 := GenS3{F0: 42, F1: 10}
                _ = v43
            case "hello":
                var v44 int64 = v2
                _ = v44
            case "\n\t":
                v45 := []int{127, 32768, 2147483647}
                _ = v45
            }
        default:
            var v46 uint32 = 2147483647
            _ = v46
        }
        _ = v23
    }
    for fi47 := 0; fi47 < 3; fi47++ {
        var v48 int64 = v2
        for fi49 := 0; fi49 < 3; fi49++ {
            var v50 float64 = (0.0 + 1e10)
            v3 = (false && ("hello" == "foo bar"))
            v4 = v3
            _ = fi49
            _ = v50
        }
        var v51 float32 = 1.5
        _ = fi47
        _ = v48
        _ = v51
    }
    if (v3 && (1000 < 1)) {
        v52 := []int{42}
        _ = v52
    } else {
        for fi53 := 0; fi53 < 7; fi53++ {
            v54 := GenS2{F0: "a", F1: ("foo bar" + "hello")}
            if v4 {
                v4 = (!v3)
            } else {
                v15 = true
            }
            for fi55 := 0; fi55 < 3; fi55++ {
                switch "\n\t" {
                case "a":
                    fmt.Println(v3, fi55, v3)
                default:
                    v4 = v15
                }
                delete(v1, fi55)
                v3 = (458 != 2)
                v3 = v3
                v15 = v4
                _ = fi55
            }
            _ = fi53
            _ = v54
        }
        fmt.Println(v15)
    }
    v15 = (!v3)
    fmt.Println(v15)
    var v56 float64 = (1e10 + 3.14159)
    v57 := v1[4294967295]
    if false {
        v58 := GenS3{F0: 100, F1: 0}
        delete(v1, 9223372036854775807)
        _ = v58
    } else {
        if (v15 || (false && v3)) {
            v3 = (4294967295 <= 0)
            for fi59 := 0; fi59 < 12; fi59++ {
                var v60 float64 = (v56 - 0.1)
                var v61 bool = v3
                v62 := v1[57]
                _ = fi59
                _ = v60
                _ = v61
                _ = v62
            }
        }
        fmt.Println(v15)
        var v63 uint32 = 2
        _ = v63
    }
    var v64 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v64
}
