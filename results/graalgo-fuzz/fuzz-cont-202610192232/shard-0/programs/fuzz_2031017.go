package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int; F1 bool; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int8; F2 uint32 }
type GenS5 struct { F0 int32; F1 int16; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64; F1 uint16; F2 bool }
func genFn0(p0 uint32, p1 int8) int32 {
    var v6 float64 = 0.0
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int = fi7
        v6 = (v6 + (((v6 / 1.0) / (float64(((float32((0.1 / (0.0 * v6))) / 3.14159) + -2.25)) + (v6 - v6))) - (3.14159 + 1.5)))
        var v9 string = ""
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"a"}
    var v3 float64 = 0.0
    var v4 int64 = -9223372036854775808
    var v5 int = 0
    fmt.Println(v4, v5)
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 uint64 = 18446744073709551615
        switch (^(v5)) {
        case 128:
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 int8 = -38
                var v14 int = (fi10 % ((fi10) | 1))
                var v15 int16 = 32767
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        case 42:
            var v16 int32 = -106
            _ = v16
        default:
            switch -2147483648 {
            case 128:
                var v17 float64 = v3
                _ = v17
            default:
                fmt.Println(v5, v3)
            }
        }
        _ = fi10
        _ = v11
    }
    var v18 string = "hello"
    for fi19 := 0; fi19 < 6; fi19++ {
        if false {
            if ((v4 - -9223372036854775808) != (v4 % ((v4) | 1))) {
                v18 = "foo bar"
                var v20 string = v18
                _ = v20
            }
            for fi21 := 0; fi21 < 15; fi21++ {
                v22 := GenS1{F0: 666, F1: 1}
                if (fi21 < -1) {
                    fmt.Println(v4)
                    var v23 uint8 = 110
                    if true {
                        var v24 string = ("foo bar" + v18)
                        v5 = 42
                        _ = v24
                    }
                    fmt.Println(v23, v5)
                    _ = v23
                }
                if (42 >= 0) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    v25 := GenS1{F0: 42, F1: 1}
                    v26 := GenS1{F0: 613, F1: 255}
                    _ = v25
                    _ = v26
                }
                _ = fi21
                _ = v22
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                _ = len(v2)
                switch "\n\t" {
                case "foo bar":
                    v5 -= (v5 + (v5 / (((^(fi27))) | 1)))
                default:
                    var v28 int = fi27
                    _ = v28
                }
                _ = fi27
            }
        }
        var v29 float64 = -2.25
        v4 = v4
        _ = fi19
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        v31 := GenS2{F0: fi30, F1: (!((1.5 < (float32(0.0) * 1e10)) || false)), F2: 0}
        var v32 uint8 = 42
        for fi33 := 0; fi33 < 6; fi33++ {
            v18 = ("\n\t" + v18)
            switch fi33 {
            case -769:
                var v34 uint64 = 2
                _ = v34
            case -821:
                fmt.Println(v32, v32, v5)
            }
            _ = fi33
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    if false {
        for fi35 := 0; fi35 < 5; fi35++ {
            for fi36 := 0; fi36 < 12; fi36++ {
                v37 := GenS3{F0: true}
                switch v18 {
                case "a":
                    if true {
                        if (42 != 42) {
                            var v38 float64 = (3.14159 - 1.5)
                            _ = v38
                        } else {
                            v39 := GenS3{F0: (!true)}
                            v3 = v3
                            v5 += (fi35 * (fi35 % ((-1) | 1)))
                            v40 := GenS2{F0: fi35, F1: false, F2: 42}
                            switch v5 {
                            case 127:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                v41 := GenS4{F0: -188, F1: 127, F2: 851}
                                _ = v41
                            }
                            _ = v39
                            _ = v40
                        }
                    } else {
                        v42 := GenS5{F0: 42, F1: 1000, F2: true}
                        fmt.Println(v18, v4)
                        if true {
                            if (3.14159 < (((0.0 - 3.14159) - 1.5) - (-2.25 - 1e10))) {
                                if false {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            } else {
                                _ = len(v1)
                                delete(v1, v18)
                                v43 := GenS4{F0: 0, F1: 1, F2: 713}
                                _ = v43
                            }
                            fmt.Println(v4, v5, v18)
                            var v44 uint8 = 42
                            _ = v44
                        } else {
                            if (!false) {
                                _ = len(v1)
                                var v45 int64 = v4
                                _ = v45
                            } else {
                                v18 = "a"
                            }
                        }
                        _ = v42
                    }
                default:
                    var v46 int = (v5 + fi36)
                    _ = v46
                }
                v2 = append(v2, (((v18 + v18) + v18) + (((("foo bar" + ("\n\t" + ((("\n\t" + v18) + "a") + v18))) + v18) + v18) + v18)))
                switch (v18 + (v18 + "\n\t")) {
                case "\n\t":
                    v47 := GenS6{F0: 0}
                    _ = v47
                default:
                    _ = len(v1)
                }
                switch fi35 {
                case 2147483647:
                    fmt.Println(fi35, fi35)
                case 127:
                    v48 := GenS5{F0: 1, F1: 42, F2: true}
                    _ = v48
                }
                _ = fi36
                _ = v37
            }
            var v49 uint32 = 354
            _ = fi35
            _ = v49
        }
        v18 = v18
        if (!((18446744073709551615 != 1) && false)) {
            v18 = ("hello" + "hello")
            for fi50 := 0; fi50 < 1; fi50++ {
                var v51 int8 = 0
                _ = fi50
                _ = v51
            }
            var v52 int = 128
            _ = v52
        } else {
            var v53 uint = 1
            _ = v53
        }
    } else {
        switch (v5 << ((0) & 63)) {
        case 127:
            fmt.Println(v5, v3)
        case -988:
            var v54 int = (v5 % ((128) | 1))
            _ = v54
        }
        var v55 bool = false
        _ = v55
    }
    v3 += v3
    var v56 int16 = 1000
    if (true || (!(3.14159 > 1.5))) {
        fmt.Println(v18)
        for fi57 := 0; fi57 < 14; fi57++ {
            fmt.Println(v3)
            var v58 uint32 = 2147483648
            _ = fi57
            _ = v58
        }
    } else {
        switch v5 {
        case 9223372036854775807:
            v18 = ("foo bar" + "foo bar")
        default:
            v3 += float64((3.14159 + (-2.25 * (1e10 / float32(1e10)))))
        }
    }
    v56 = v56
    for fi59 := 0; fi59 < 16; fi59++ {
        for fi60 := 0; fi60 < 11; fi60++ {
            if ((0 > 4294967295) && (false && true)) {
                var v61 int32 = 1
                var v62 int64 = 1
                _ = v61
                _ = v62
            } else {
                if (42 > 0) {
                    if (true && true) {
                        if ((((1e10 / (0.1 - -2.25)) * float32(v3)) > 1e10) || false) {
                            var v63 bool = true
                            var v64 int8 = 100
                            var v65 string = "hello"
                            _ = len(v1)
                            _ = v63
                            _ = v64
                            _ = v65
                        } else {
                            v4 -= v4
                            _ = len(v2)
                            if len(v2) > 0 {
                                v2[0] = v18
                            }
                        }
                        switch ("foo bar" + (v18 + v18)) {
                        case "hello":
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = v18
                                }
                                v18 = ((v18 + ((("a" + v18) + "hello") + (v18 + v18))) + "foo bar")
                            }
                        }
                        var v66 int8 = -128
                        if true {
                            fmt.Println(v66, v4)
                            v4 = v4
                            v2 = append(v2, "foo bar")
                        } else {
                            var v67 int8 = 23
                            var v68 uint = 975
                            fmt.Println(v56, v5)
                            fmt.Println(fi59)
                            v3 = v3
                            _ = v67
                            _ = v68
                        }
                        fmt.Println(v4, v5)
                        var v69 int = v5
                        fmt.Println(v3, v69)
                        v4 = (-(v4))
                        v70 := v1[(v18 + (v18 + (v18 + (("hello" + ("hello" + v18)) + (("a" + v18) + ("a" + ("foo bar" + "foo bar")))))))]
                        if (1 < (v4 + v4)) {
                            v4 = (v4 % ((v4) | 1))
                            v18 = (v18 + v18)
                        }
                        _ = v66
                        _ = v69
                        _ = v70
                    } else {
                        var v71 float32 = 0.0
                        _ = v71
                    }
                    var v72 bool = false
                    fmt.Println(v4, v5, v3)
                    _ = v72
                } else {
                    v3 -= float64(-2.25)
                    switch (fi60 ^ (fi59 ^ 32767)) {
                    case -208:
                        var v73 uint16 = 65535
                        _ = v73
                    case 232:
                        v4 = v4
                    case 127:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                if (v4 < 369) {
                    v74 := GenS6{F0: 18446744073709551615}
                    _ = v74
                }
                v5 = (fi59 | 32768)
            }
            if len(v2) > 0 {
                v2[0] = (("foo bar" + "foo bar") + v18)
            }
            v4 += -1
            var v75 int16 = v56
            _ = fi60
            _ = v75
        }
        if ((65535 <= 455) && (!false)) {
            v4 -= v4
        }
        _ = fi59
    }
    v76 := GenS7{F0: 2, F1: 0, F2: ((!(false && (false && false))) && true)}
    v18 = v18
    for fi77 := 0; fi77 < 8; fi77++ {
        var v78 bool = false
        _ = fi77
        _ = v78
    }
    v18 = ""
    v56 = (v56 * v56)
    var v79 bool = ((true && true) || (2147483647 <= 2147483647))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v76
    _ = v79
}
