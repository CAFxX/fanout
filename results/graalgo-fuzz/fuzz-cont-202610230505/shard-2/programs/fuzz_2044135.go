package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint8 }
func genFn0(p0 int) int32 {
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(p0, p0)
        fmt.Println(fi7, p0)
        if (true || (2 < 753)) {
            var v8 int32 = 0
            _ = v8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS1{F0: p0, F1: -217}
        var v11 uint = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(p0)
    var v12 float32 = 1e10
    var v13 bool = true
    _ = p0
    _ = v12
    _ = v13
    return 2147483647
}
func genFn1(p0 int8, p1 uint, p2 string) uint32 {
    v14 := []int{213}
    var v15 int = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 2147483648
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"foo bar", "hello"}
    var v3 int8 = 70
    var v4 uint8 = 66
    var v5 uint8 = (v4 - v4)
    var v6 uint = 665
    v16 := GenS1{F0: 637, F1: int16(v5)}
    for fi17 := 0; fi17 < 10; fi17++ {
        switch (("hello" + "") + "hello") {
        case "a":
            if true {
                if false {
                    v18 := GenS2{F0: v3}
                    for fi19 := 0; fi19 < 7; fi19++ {
                        switch ("hello" + "hello") {
                        case "foo bar":
                            v4 += 128
                        }
                        _ = fi19
                    }
                    _ = v18
                } else {
                    var v20 uint8 = (v5 << ((2) & 63))
                    v6 = uint(v5)
                    _ = v20
                }
                var v21 uint64 = 42
                for fi22 := 0; fi22 < 3; fi22++ {
                    if (v5 < 123) {
                        var v23 uint64 = v21
                        v24 := GenS3{F0: 500}
                        v3 += v3
                        _ = v23
                        _ = v24
                    } else {
                        v2 = append(v2, ("\n\t" + ("" + "hello")))
                        var v25 int64 = -1
                        _ = v25
                    }
                    if false {
                        switch "a" {
                        case "hello":
                            var v26 uint = v6
                            _ = v26
                        case "a":
                            v4 = v5
                        case "":
                            switch 42 {
                            case -711:
                                var v27 float32 = 3.14159
                                _ = v27
                            case -609:
                                v3 += (v3 & 127)
                            case 127:
                                v6 += (v6 + 18446744073709551615)
                            }
                        default:
                            if false {
                                v28 := GenS2{F0: (v3 - (v3 % ((v3) | 1)))}
                                var v29 int16 = -1
                                var v30 uint32 = 2147483647
                                _ = v28
                                _ = v29
                                _ = v30
                            } else {
                                if (!((-(fi17)) != (fi17 >> (((v6 + v6)) & 63)))) {
                                    fmt.Println(v21)
                                    v31 := GenS3{F0: 653}
                                    var v32 bool = true
                                    _ = v31
                                    _ = v32
                                }
                                v33 := GenS2{F0: (-(v3))}
                                fmt.Println(v4, v4)
                                _ = v33
                            }
                        }
                        v5 = (v4 - 2)
                        var v34 uint32 = 42
                        _ = v34
                    }
                    _ = fi22
                }
                _ = v21
            } else {
                var v35 int8 = (v3 >> ((v6) & 63))
                _ = v35
            }
        default:
            switch "foo bar" {
            case "hello":
                switch fi17 {
                case 127:
                    v5 = v4
                case 42:
                    var v36 int16 = -434
                    _ = v36
                default:
                    for fi37 := 0; fi37 < 7; fi37++ {
                        v6 = 255
                        delete(v1, 65535)
                        v38 := GenS3{F0: 1}
                        _ = fi37
                        _ = v38
                    }
                }
            case "":
                for fi39 := 0; fi39 < 4; fi39++ {
                    v6 -= v6
                    var v40 uint64 = 2
                    fmt.Println(v5)
                    _ = fi39
                    _ = v40
                }
            case "\n\t":
                switch 32767 {
                case -545:
                    v4 += v4
                default:
                    if (3.14159 == 0.1) {
                        if (!((2 <= 114) && (fi17 > (fi17 * fi17)))) {
                            v5 = (^(v4))
                            _ = len(v2)
                        } else {
                            var v41 float64 = 0.0
                            switch (fi17 & (^(fi17))) {
                            case -9223372036854775808:
                                if true {
                                    var v42 int32 = 2147483647
                                    fmt.Println(v6, fi17, v42)
                                    _ = v42
                                }
                            case -1:
                                switch fi17 {
                                case 42:
                                    v6 = 586
                                }
                            case 0:
                                fmt.Println(fi17, v4, v3)
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            var v43 int64 = -442
                            _ = v41
                            _ = v43
                        }
                        if (65535 == (v6 & (^(v6)))) {
                            v44 := GenS4{F0: 3, F1: 4294967295, F2: (v5 | 88)}
                            _ = v44
                        } else {
                            var v45 uint = (^(v6))
                            var v46 string = "a"
                            _ = v45
                            _ = v46
                        }
                    }
                }
            default:
                v47 := v1[2147483648]
                _ = v47
            }
        }
        v6 = (^(v6))
        _ = fi17
    }
    v4 = v4
    v48 := GenS1{F0: 1, F1: 1000}
    fmt.Println(v6, v5, v3)
    for fi49 := 0; fi49 < 20; fi49++ {
        var v50 int8 = v3
        for fi51 := 0; fi51 < 7; fi51++ {
            v4 += v5
            _ = fi51
        }
        if (((3.14159 - -2.25) * (0.0 / 1.5)) < 1e10) {
            if true {
                fmt.Println(v3, v50, fi49)
            } else {
                var v52 uint = v6
                _ = v52
            }
            if true {
                var v53 uint8 = 80
                v54 := GenS3{F0: 344}
                _ = v53
                _ = v54
            }
            if false {
                var v55 uint16 = 439
                v3 = 100
                _ = v55
            }
        }
        _ = fi49
        _ = v50
    }
    fmt.Println(v6)
    var v56 uint16 = 319
    if (false && true) {
        var v57 float64 = (((3.14159 + 3.14159) - (((1e10 / 1.0) / ((1.5 * 0.0) - (1e10 - 0.0))) * -2.25)) - -2.25)
        _ = v57
    } else {
        var v58 int32 = 42
        if (0 < -1) {
            for fi59 := 0; fi59 < 5; fi59++ {
                v60 := GenS1{F0: fi59, F1: 892}
                for fi61 := 0; fi61 < 4; fi61++ {
                    var v62 uint8 = 225
                    v3 = -128
                    v2 = append(v2, ("\n\t" + ""))
                    switch fi59 {
                    case 32768:
                        v62 = v62
                    case -9223372036854775808:
                        switch (fi61 >> ((uint(v56)) & 63)) {
                        case 387:
                            v56 += (v56 / ((0) | 1))
                        case 4294967295:
                            fmt.Println(v56)
                        }
                    case 42:
                        v2 = append(v2, "")
                    default:
                        var v63 uint = (^(v6))
                        _ = v63
                    }
                    var v64 uint16 = uint16(v6)
                    _ = fi61
                    _ = v62
                    _ = v64
                }
                _ = len(v1)
                _ = fi59
                _ = v60
            }
        }
        _ = v58
    }
    var v65 bool = (0 <= -1)
    var v66 string = ""
    v65 = ((42 != (v3 ^ (v3 & v3))) || v65)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v56
    _ = v6
    _ = v65
    _ = v66
}
