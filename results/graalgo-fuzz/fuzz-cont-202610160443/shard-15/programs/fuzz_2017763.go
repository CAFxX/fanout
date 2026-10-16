package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS4 struct { F0 string; F1 int64; F2 int }
type GenS5 struct { F0 int64 }
func genFn0() string {
    var v7 int8 = 0
    v8 := GenS1{F0: (v7 + v7), F1: 2, F2: 1}
    _ = v7
    _ = v8
    return ("hello" + "\n\t")
}
func genFn1(p0 uint16) int16 {
    p0 = p0
    if (!(-678 <= -102)) {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 bool = true
            _ = fi9
            _ = v10
        }
        v11 := []int{127, -2147483648}
        _ = v11
    }
    _ = p0
    return 780
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"a", ("\n\t" + ""), ("hello" + (("" + "a") + ""))}
    v3 := []string{(("hello" + ("hello" + "\n\t")) + "\n\t")}
    var v4 uint64 = 18446744073709551615
    var v5 bool = true
    var v6 bool = true
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 string = ("a" + ("a" + ("\n\t" + "foo bar")))
        var v14 uint16 = 1
        v4 = (v4 >> ((372) & 63))
        _ = fi12
        _ = v13
        _ = v14
    }
    v6 = ((32768 < 9223372036854775807) || v6)
    fmt.Println(v4, v6)
    fmt.Println(v5, v6)
    if false {
        var v15 int8 = -1
        _ = v15
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v16 uint16 = 160
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS2{F0: 394, F1: 1}
            if v6 {
                v1[2147483647] = 42
                v5 = v5
            }
            switch (^(fi18)) {
            case -2147483648:
                switch "foo bar" {
                case "foo bar":
                    var v20 int = fi17
                    _ = v20
                case "":
                    if true {
                        v21 := GenS2{F0: 42, F1: 255}
                        var v22 int32 = 196
                        if len(v3) > 0 {
                            v3[0] = (("\n\t" + "a") + ("foo bar" + "foo bar"))
                        }
                        _ = v21
                        _ = v22
                    }
                case "hello":
                    if v6 {
                        var v23 int32 = -1
                        _ = v23
                    } else {
                        var v24 string = ("\n\t" + ("hello" + ""))
                        _ = v24
                    }
                default:
                    fmt.Println(fi17)
                }
            case -9223372036854775808:
                var v25 int16 = 1
                _ = v25
            case 9223372036854775807:
                fmt.Println(fi17, v6, fi17)
            default:
                if (!("hello" == "")) {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    var v26 uint8 = 128
                    var v27 uint64 = 18446744073709551615
                    _ = v26
                    _ = v27
                }
            }
            fmt.Println(fi17)
            v28 := GenS3{F0: (v4 >> ((37) & 63)), F1: -9223372036854775808, F2: 65535}
            _ = fi18
            _ = v19
            _ = v28
        }
        delete(v1, 2)
        _ = fi17
    }
    v5 = (v6 || true)
    var v29 float64 = (1.5 + -2.25)
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 19; fi31++ {
            v6 = (0 > 1)
            var v32 int8 = -62
            v2 = append(v2, "foo bar")
            v29 = 0.0
            _ = fi31
            _ = v32
        }
        v33 := GenS4{F0: "", F1: -9223372036854775808, F2: (fi30 + fi30)}
        v4 = v4
        _ = fi30
        _ = v33
    }
    if v5 {
        var v34 int16 = -1
        if v5 {
            switch 32768 {
            case 1:
                for fi35 := 0; fi35 < 20; fi35++ {
                    fmt.Println(fi35, fi35)
                    v36 := GenS4{F0: "", F1: 9223372036854775807, F2: fi35}
                    v6 = (v5 || v6)
                    _ = fi35
                    _ = v36
                }
            case -2147483648:
                v37 := GenS2{F0: 1, F1: 306}
                _ = v37
            default:
                switch "a" {
                case "\n\t":
                    var v38 int8 = -1
                    _ = v38
                case "foo bar":
                    var v39 uint = 2
                    _ = v39
                case "hello":
                    v40 := GenS4{F0: ("hello" + ("" + "\n\t")), F1: 0, F2: 0}
                    _ = v40
                }
            }
            if (1 < 42) {
                var v41 string = "foo bar"
                var v42 uint16 = 1
                _ = v41
                _ = v42
            } else {
                v4 -= 158
            }
            fmt.Println(v5, v29)
        }
        _ = v34
    }
    for fi43 := 0; fi43 < 16; fi43++ {
        v5 = v6
        for fi44 := 0; fi44 < 20; fi44++ {
            switch (fi43 << ((42) & 63)) {
            case -514:
                v4 = (v4 ^ v4)
            case -9223372036854775808:
                switch "hello" {
                case "hello":
                    v45 := GenS4{F0: "hello", F1: 9223372036854775807, F2: (fi44 + fi44)}
                    _ = v45
                case "":
                    var v46 uint64 = 1
                    _ = v46
                default:
                    v47 := GenS2{F0: 65535, F1: 2}
                    _ = v47
                }
            case -2147483648:
                v5 = v5
            default:
                v29 = v29
            }
            _ = len(v2)
            _ = fi44
        }
        for fi48 := 0; fi48 < 3; fi48++ {
            var v49 float64 = ((float64((-2.25 / (float32(float64((-2.25 / ((1e10 / -2.25) * (float32(v29) * 1.5))))) - 3.14159))) * 0.1) - (0.0 + 0.0))
            _ = fi48
            _ = v49
        }
        _ = fi43
    }
    v50 := GenS4{F0: "\n\t", F1: 1, F2: 128}
    v29 = (float64(-2.25) - v29)
    for fi51 := 0; fi51 < 12; fi51++ {
        v6 = v5
        if false {
            for fi52 := 0; fi52 < 14; fi52++ {
                v29 += v29
                var v53 bool = (255 == 2)
                var v54 uint64 = v4
                if (float64((0.1 + 0.1)) <= v29) {
                    fmt.Println(v5, v6, v54)
                    v4 = (v54 | v54)
                } else {
                    var v55 uint64 = v4
                    v56 := GenS3{F0: (v54 << ((2) & 63)), F1: 9223372036854775807, F2: 690}
                    if (false || v53) {
                        v54 -= (^(v4))
                        v57 := v1[uint32(v4)]
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v58 int8 = -53
                            _ = v58
                        }
                        fmt.Println(v6, v6, v55)
                        _ = v57
                    } else {
                        v59 := GenS3{F0: (v4 - v55), F1: 0, F2: 745}
                        _ = len(v3)
                        var v60 float64 = float64((1e10 / (-2.25 + 1e10)))
                        var v61 string = "a"
                        switch fi52 {
                        case 128:
                            fmt.Println(v60, v61)
                            v62 := GenS5{F0: -279}
                            v63 := GenS1{F0: -66, F1: 2, F2: (^(v55))}
                            _ = v62
                            _ = v63
                        }
                        if ((v55 - v54) > (v54 - (^(v4)))) {
                            var v64 string = v61
                            _ = v64
                        } else {
                            if (!v5) {
                                var v65 uint16 = 0
                                var v66 int = (fi52 & 9223372036854775807)
                                v61 = (v61 + v61)
                                _ = v65
                                _ = v66
                            }
                        }
                        var v67 bool = v53
                        _ = v59
                        _ = v60
                        _ = v61
                        _ = v67
                    }
                    _ = v55
                    _ = v56
                }
                _ = fi52
                _ = v53
                _ = v54
            }
            var v68 bool = false
            for fi69 := 0; fi69 < 8; fi69++ {
                if v5 {
                    fmt.Println(v4, fi69)
                    v70 := GenS1{F0: 0, F1: 65535, F2: v4}
                    var v71 uint = 0
                    _ = v70
                    _ = v71
                } else {
                    var v72 int16 = 42
                    if (18446744073709551615 >= 942) {
                        delete(v1, 527)
                        v68 = (127 <= 42)
                        var v73 uint16 = 32767
                        _ = v73
                    } else {
                        v74 := GenS1{F0: 6, F1: 42, F2: 0}
                        switch (-(fi69)) {
                        case 32767:
                            switch "\n\t" {
                            case "":
                                v68 = (-1 >= 42)
                            case "\n\t":
                                fmt.Println(v72)
                            default:
                                v75 := GenS2{F0: 4294967295, F1: 811}
                                _ = v75
                            }
                        case 4294967295:
                            v3 = append(v3, "hello")
                        case 127:
                            v6 = v6
                        }
                        v4 += v4
                        _ = v74
                    }
                    v76 := GenS4{F0: "foo bar", F1: 9223372036854775807, F2: -133}
                    _ = v72
                    _ = v76
                }
                fmt.Println(v6, v29, v4)
                _ = fi69
            }
            _ = v68
        }
        v77 := GenS1{F0: -84, F1: 1, F2: (v4 << ((2) & 63))}
        _ = fi51
        _ = v77
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
