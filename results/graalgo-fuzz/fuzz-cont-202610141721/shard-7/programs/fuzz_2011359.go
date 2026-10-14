package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint; F1 int; F2 int16 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 bool; F1 string; F2 int32 }
func genFn0(p0 uint64) int64 {
    var v6 int64 = -842
    v6 = v6
    v7 := []int{}
    _ = p0
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 uint) int8 {
    fmt.Println(p0)
    var v8 int = 1
    v9 := []int{-1, -228, (v8 << (((p0 % (((^(p0))) | 1))) & 63))}
    _ = p0
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]int)
    var v3 int8 = 1
    v4 := make(map[uint32]string)
    v5 := []string{"\n\t", "hello", ("a" + "\n\t")}
    switch ("" + ("a" + "foo bar")) {
    case "a":
        var v10 uint8 = 2
        _ = v10
    case "foo bar":
        v1 = genFn0(879)
    case "\n\t":
        for fi11 := 0; fi11 < 7; fi11++ {
            if (3.14159 == float64(-2.25)) {
                fmt.Println(v3)
            }
            if true {
                var v12 uint64 = 42
                v13 := GenS1{F0: (v12 % ((v12) | 1)), F1: v12, F2: (-(v3))}
                _ = v12
                _ = v13
            }
            if true {
                v14 := GenS1{F0: 1, F1: 0, F2: -13}
                fmt.Println(v3, v3)
                if (genFn1(65535) == genFn1(0)) {
                    v15 := GenS1{F0: 36, F1: 2, F2: (v3 ^ 49)}
                    _ = v15
                } else {
                    v16 := GenS1{F0: 1, F1: 18446744073709551615, F2: (v3 + 100)}
                    for fi17 := 0; fi17 < 16; fi17++ {
                        var v18 uint16 = 65535
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v3 += v3
                        _ = fi19
                    }
                    _ = v16
                }
                _ = v14
            }
            _ = fi11
        }
    default:
        if (32767 != 0) {
            for fi20 := 0; fi20 < 7; fi20++ {
                if (2147483647 >= 2147483647) {
                    fmt.Println(v1, v3, v3)
                    v1 = genFn0(662)
                    v1 = (v1 << ((931) & 63))
                } else {
                    for fi21 := 0; fi21 < 4; fi21++ {
                        var v22 float32 = 0.1
                        _ = fi21
                        _ = v22
                    }
                    v23 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: v3}
                    _ = v23
                }
                var v24 int32 = 0
                fmt.Println(fi20)
                _ = fi20
                _ = v24
            }
            for fi25 := 0; fi25 < 19; fi25++ {
                if (-32768 > int16(fi25)) {
                    v5 = append(v5, "hello")
                    switch fi25 {
                    case 42:
                        _ = len(v4)
                    case 32768:
                        var v26 int32 = 42
                        _ = v26
                    case 88:
                        v2[(fi25 + (fi25 >> ((4294967295) & 63)))] = 32768
                    default:
                        var v27 int64 = -9223372036854775808
                        _ = v27
                    }
                }
                _ = len(v2)
                var v28 uint64 = 344
                _ = fi25
                _ = v28
            }
            var v29 int32 = 1
            _ = v29
        } else {
            var v30 string = "\n\t"
            var v31 int64 = v1
            _ = v30
            _ = v31
        }
    }
    var v32 uint8 = 127
    for fi33 := 0; fi33 < 6; fi33++ {
        var v34 int64 = (v1 % ((int64(v3)) | 1))
        var v35 int16 = 1
        for fi36 := 0; fi36 < 7; fi36++ {
            v37 := GenS1{F0: 1000, F1: 18446744073709551615, F2: int8(v34)}
            if false {
                var v38 int = fi33
                _ = v38
            } else {
                fmt.Println(v34, fi33)
                v39 := GenS2{F0: 42, F1: -184}
                switch "\n\t" {
                case "":
                    var v40 uint = 4294967295
                    _ = v40
                case "a":
                    v32 = v32
                default:
                    fmt.Println(v1, v32)
                }
                _ = v39
            }
            var v41 int16 = int16(v32)
            _ = fi36
            _ = v37
            _ = v41
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    if false {
        var v42 int64 = (v1 << ((822) & 63))
        fmt.Println(v1, v32, v42)
        _ = v42
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        if ((!true) || true) {
            v44 := GenS3{F0: 2147483647, F1: 42}
            if (v3 > (v3 | (v3 & genFn1(65535)))) {
                v1 = (v1 * -9223372036854775808)
                var v45 int32 = -245
                _ = v45
            } else {
                for fi46 := 0; fi46 < 20; fi46++ {
                    v3 -= (v3 & 0)
                    if (("a" + "") == "\n\t") {
                        var v47 int32 = 446
                        v32 += 42
                        var v48 uint = 4294967295
                        _ = v47
                        _ = v48
                    } else {
                        switch (fi43 & fi43) {
                        case -9223372036854775808:
                            var v49 int = fi43
                            switch "" {
                            case "a":
                                v50 := GenS1{F0: 18446744073709551615, F1: 210, F2: -128}
                                _ = v50
                            case "foo bar":
                                switch fi46 {
                                case 42:
                                    fmt.Println(fi43)
                                default:
                                    v32 -= (^(v32))
                                }
                            case "hello":
                                v51 := GenS4{F0: 1, F1: v49, F2: -909}
                                var v52 uint8 = (^(v32))
                                _ = v51
                                _ = v52
                            default:
                                v53 := GenS2{F0: 381, F1: 42}
                                _ = v53
                            }
                            _ = v49
                        }
                        if true {
                            delete(v2, fi46)
                            v3 = (v3 & v3)
                        }
                        delete(v2, -1)
                        if (v1 == (v1 % ((v1) | 1))) {
                            v54 := GenS2{F0: 110, F1: 0}
                            v32 = v32
                            _ = v54
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi46
                }
                for fi55 := 0; fi55 < 6; fi55++ {
                    if false {
                        var v56 int8 = v3
                        _ = v56
                    }
                    v57 := GenS2{F0: uint64(v32), F1: 2147483647}
                    _ = fi55
                    _ = v57
                }
            }
            _ = v44
        }
        v3 = (v3 ^ -128)
        fmt.Println(v1, v1, v32)
        _ = fi43
    }
    for fi58 := 0; fi58 < 5; fi58++ {
        v59 := v2[fi58]
        v3 += (v3 & -13)
        _ = fi58
        _ = v59
    }
    if false {
        v60 := v2[127]
        if false {
            var v61 uint64 = 704
            var v62 int = v60
            _ = v61
            _ = v62
        } else {
            var v63 int32 = -2147483648
            switch (v60 % (((^(v60))) | 1)) {
            case -777:
                v64 := GenS5{F0: false, F1: 2147483648}
                _ = v64
            case 4294967295:
                var v65 float64 = (float64(3.14159) / (1e10 + 0.1))
                _ = v65
            }
            v66 := GenS3{F0: 42, F1: 255}
            _ = v63
            _ = v66
        }
        _ = v60
    } else {
        v3 -= (v3 & v3)
    }
    var v67 int64 = v1
    if (42 > 895) {
        var v68 int = -320
        _ = v68
    } else {
        var v69 string = "hello"
        for fi70 := 0; fi70 < 10; fi70++ {
            for fi71 := 0; fi71 < 18; fi71++ {
                var v72 uint32 = 431
                _ = fi71
                _ = v72
            }
            _ = fi70
        }
        _ = v69
    }
    v3 = v3
    var v73 uint = 1
    var v74 float32 = 3.14159
    _ = len(v5)
    if (true || (0.0 < (0.1 / float64(3.14159)))) {
        for fi75 := 0; fi75 < 8; fi75++ {
            for fi76 := 0; fi76 < 1; fi76++ {
                switch (fi76 ^ (fi76 | (fi75 ^ (-(fi75))))) {
                case 2147483647:
                    v1 = v1
                case -2147483648:
                    v3 = 127
                }
                switch (("foo bar" + "\n\t") + "a") {
                case "a":
                    var v77 float64 = 0.1
                    _ = v77
                }
                fmt.Println(v73, fi75)
                _ = fi76
            }
            var v78 string = "a"
            _ = fi75
            _ = v78
        }
        v3 = (v3 << ((v73) & 63))
        v79 := GenS6{F0: false, F1: "foo bar", F2: int32(v1)}
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v67:", v67)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v67
    _ = v73
    _ = v74
}
