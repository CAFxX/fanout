package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int; F2 bool }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint16; F1 bool; F2 int16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := []string{(("" + "a") + "foo bar"), "hello"}
    v4 := []string{("hello" + ""), ""}
    var v5 string = ""
    var v6 int64 = 892
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint16 = 65535
        _ = fi7
        _ = v8
    }
    if (!(false && (-1 == 42))) {
        if (884 <= 42) {
            fmt.Println(v6, v5)
            _ = len(v3)
            if false {
                var v9 uint32 = 2
                v5 = ""
                _ = v9
            } else {
                var v10 uint16 = 42
                v11 := GenS1{F0: true}
                _ = v10
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 10; fi12++ {
                if len(v4) > 0 {
                    v4[0] = (v5 + v5)
                }
                _ = fi12
            }
            v13 := v1["foo bar"]
            _ = v13
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (1e10 > 1e10) {
            v6 += (v6 + (v6 + v6))
        } else {
            v14 := GenS1{F0: true}
            v5 = (v5 + (((v5 + (v5 + (((v5 + v5) + "hello") + v5))) + v5) + v5))
            switch 1 {
            case -219:
                fmt.Println(v6, v6)
            case 2147483647:
                var v15 int64 = (-(v6))
                _ = v15
            }
            _ = v14
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v2)
    for fi16 := 0; fi16 < 9; fi16++ {
        if true {
            for fi17 := 0; fi17 < 18; fi17++ {
                fmt.Println(v5, v5, v5)
                var v18 int = (fi16 * fi16)
                _ = fi17
                _ = v18
            }
        } else {
            if false {
                for fi19 := 0; fi19 < 4; fi19++ {
                    v3 = append(v3, (v5 + "\n\t"))
                    _ = fi19
                }
            }
            switch "foo bar" {
            case "foo bar":
                v20 := GenS2{F0: 0, F1: (fi16 | fi16), F2: (!false)}
                _ = v20
            case "hello":
                for fi21 := 0; fi21 < 16; fi21++ {
                    fmt.Println(v6)
                    switch (^(fi21)) {
                    case 956:
                        fmt.Println(v5)
                    case 42:
                        if true {
                            switch 954 {
                            case 32768:
                                v5 = (((v5 + ("" + v5)) + ("a" + "hello")) + v5)
                            case 4294967295:
                                var v22 int64 = v6
                                _ = v22
                            case -2147483648:
                                var v23 bool = true
                                _ = v23
                            }
                        }
                    case 0:
                        switch "foo bar" {
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                            v6 = v6
                            switch (fi16 << ((18446744073709551615) & 63)) {
                            case 850:
                                v24 := v1[((v5 + v5) + v5)]
                                _ = v24
                            case 32767:
                                fmt.Println(v6, fi16)
                            default:
                                if (77 > 154) {
                                    fmt.Println(fi16, v5)
                                    v25 := GenS1{F0: false}
                                    _ = v25
                                } else {
                                    var v26 uint64 = 18446744073709551615
                                    v6 = v6
                                    _ = v26
                                }
                                v5 = ((v5 + v5) + (((v5 + "foo bar") + ((v5 + (("\n\t" + ((v5 + "a") + v5)) + (v5 + ("\n\t" + "hello")))) + (v5 + v5))) + ""))
                                v27 := GenS3{F0: 229, F1: 443}
                                _ = v27
                            }
                            var v28 string = "hello"
                            _ = v28
                        case "hello":
                            v5 = v5
                            v5 = "hello"
                        default:
                            fmt.Println(fi21)
                        }
                    }
                    v29 := v2[("" + "")]
                    _ = fi21
                    _ = v29
                }
            default:
                v30 := GenS2{F0: 32, F1: 0, F2: (fi16 <= 0)}
                _ = v30
            }
            for fi31 := 0; fi31 < 19; fi31++ {
                v32 := GenS1{F0: (!(-2.25 != (3.14159 / -2.25)))}
                v6 += v6
                _ = fi31
                _ = v32
            }
        }
        _ = fi16
    }
    if (133 < 42) {
        switch 341 {
        case -1:
            if (((0.0 - ((0.1 * 1.5) * float32(((0.0 + 1e10) + float64((0.0 + (((0.0 - 0.1) * -2.25) / ((float32(1.5) * 0.1) - 0.0)))))))) / 1.0) >= (-2.25 * 3.14159)) {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
            } else {
                v33 := GenS3{F0: 1, F1: 0}
                _ = v33
            }
        }
    }
    switch 42 {
    case 128:
        var v34 string = (v5 + "hello")
        _ = v34
    case 42:
        v6 += v6
    case -826:
        v5 = v5
    default:
        switch v5 {
        case "foo bar":
            if (1 == -183) {
                for fi35 := 0; fi35 < 12; fi35++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v36 := GenS3{F0: 2, F1: 65535}
                    var v37 uint64 = 1
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            } else {
                delete(v2, v5)
            }
        }
    }
    fmt.Println(v5, v5)
    if (0.0 >= -2.25) {
        for fi38 := 0; fi38 < 2; fi38++ {
            var v39 bool = true
            var v40 uint32 = 2
            _ = fi38
            _ = v39
            _ = v40
        }
        v41 := GenS3{F0: 127, F1: 0}
        v5 = v5
        _ = v41
    } else {
        fmt.Println(v5)
        v6 -= (-(v6))
    }
    fmt.Println(v6, v6)
    if (0 >= 4294967295) {
        for fi42 := 0; fi42 < 8; fi42++ {
            v43 := GenS3{F0: 113, F1: 662}
            _ = fi42
            _ = v43
        }
        if false {
            v6 = (v6 ^ (v6 * (v6 << ((917) & 63))))
            for fi44 := 0; fi44 < 17; fi44++ {
                var v45 int16 = 1
                var v46 int64 = v6
                if (-2.25 < 0.0) {
                    if ((1 > 2147483648) && true) {
                        if (2 < 2) {
                            v45 = (v45 - 533)
                        } else {
                            if true {
                                for fi47 := 0; fi47 < 20; fi47++ {
                                    var v48 int16 = 367
                                    v49 := GenS3{F0: 2, F1: 2}
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            } else {
                                var v50 int64 = -243
                                _ = v50
                            }
                            v45 = (v45 + (v45 - (v45 & -1)))
                        }
                        if true {
                            delete(v1, v5)
                            switch (-(fi44)) {
                            case -1:
                                v5 = (v5 + v5)
                            case -617:
                                var v51 bool = false
                                _ = v51
                            }
                            if false {
                                var v52 uint = 2
                                v53 := GenS3{F0: 239, F1: 2147483648}
                                _ = v52
                                _ = v53
                            }
                        }
                    } else {
                        for fi54 := 0; fi54 < 19; fi54++ {
                            delete(v1, ("a" + ""))
                            v1[(v5 + v5)] = fi44
                            _ = fi54
                        }
                        var v55 int32 = -2147483648
                        _ = v55
                    }
                    var v56 int = fi44
                    if (true || false) {
                        v57 := GenS2{F0: -65, F1: -9223372036854775808, F2: (!true)}
                        _ = v57
                    }
                    _ = v56
                }
                _ = fi44
                _ = v45
                _ = v46
            }
            v5 = ((v5 + ("a" + "\n\t")) + (("hello" + ("a" + "")) + "foo bar"))
        } else {
            fmt.Println(v5, v5)
            if (!(true && false)) {
                v5 = v5
            } else {
                v6 = v6
                v58 := GenS4{F0: 1, F1: true, F2: -577}
                _ = v58
            }
        }
        var v59 int32 = -1
        _ = v59
    } else {
        v2[v5] = true
    }
    var v60 uint16 = 761
    var v61 float32 = ((1e10 - 0.0) - 0.1)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    delete(v1, v5)
    v61 -= (v61 - (float32(0.1) + (float32(((3.14159 + float64(((v61 * v61) / v61))) / 0.1)) / (v61 / (((0.0 / (0.0 + 0.1)) - v61) / (float32(1e10) + ((v61 * float32(-2.25)) - 0.1)))))))
    v5 = (v5 + "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
