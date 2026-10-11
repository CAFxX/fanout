package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int16; F1 int32 }
func genFn0() int {
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = v5
    return 128
}
func genFn1(p0 int) int64 {
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int32 = -2147483648
        _ = fi6
        _ = v7
    }
    var v8 string = ("hello" + ("a" + ""))
    if (1000 != -505) {
        switch "a" {
        case "\n\t":
            v9 := []int{(p0 << ((227) & 63)), genFn0()}
            _ = v9
        case "foo bar":
            p0 += (p0 >> ((2) & 63))
        }
        if false {
            switch p0 {
            case 32768:
                fmt.Println(p0)
            case -9223372036854775808:
                v10 := GenS1{F0: 190}
                _ = v10
            default:
                var v11 uint32 = 4294967295
                _ = v11
            }
            v8 = v8
            v12 := GenS1{F0: 42}
            _ = v12
        } else {
            v13 := []int{p0}
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := []int{p0, (p0 >> ((18446744073709551615) & 63))}
            v16 := GenS1{F0: -519}
            v15 = append(v15, p0)
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    var v17 bool = ((true && false) || ((true || (0 < 4294967295)) || (!(!true))))
    _ = p0
    _ = v17
    _ = v8
    return 0
}

func main() {
    var v1 uint = 42
    var v2 uint = (v1 ^ v1)
    v3 := make(map[string]int)
    v4 := []uint32{65535, 2147483647}
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v1, v2, fi18)
            v2 -= v2
            _ = fi19
        }
        _ = fi18
    }
    switch 128 {
    case 320:
        v20 := GenS1{F0: 1}
        _ = v20
    case 4294967295:
        v2 -= v2
    case 152:
        if true {
            v21 := GenS2{F0: 529, F1: 127}
            var v22 uint64 = 0
            _ = v21
            _ = v22
        } else {
            fmt.Println(v2, v2)
        }
    default:
        var v23 bool = (!(false || ((float32(1.5) > (3.14159 - 0.0)) && (42 >= 4294967295))))
        _ = v23
    }
    var v24 int = 42
    v2 = 65535
    v1 = 18446744073709551615
    var v25 string = "\n\t"
    switch v25 {
    case "":
        v1 = (v1 & v2)
    case "\n\t":
        v1 = (v1 - 940)
    default:
        switch ((v25 + ("" + v25)) + "a") {
        case "a":
            var v26 int64 = -9223372036854775808
            _ = v26
        case "hello":
            _ = len(v3)
        default:
            v27 := v3[((("a" + v25) + "foo bar") + v25)]
            _ = v27
        }
    }
    v28 := GenS2{F0: int16(v24), F1: 0}
    switch ("foo bar" + "foo bar") {
    case "hello":
        v25 = ((((v25 + "a") + v25) + v25) + (v25 + v25))
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if ((-32768 < -17) && (2147483647 > 0)) {
        v25 = v25
        v29 := GenS2{F0: 1000, F1: 127}
        _ = v29
    } else {
        switch (-(v24)) {
        case -814:
            for fi30 := 0; fi30 < 14; fi30++ {
                for fi31 := 0; fi31 < 18; fi31++ {
                    var v32 int64 = 42
                    switch genFn0() {
                    case 1:
                        if false {
                            v32 += v32
                            if false {
                                var v33 uint = 65535
                                _ = v33
                            }
                        }
                    default:
                        switch fi30 {
                        case 127:
                            v1 = (v2 / ((0) | 1))
                        default:
                            var v34 uint16 = 2
                            _ = v34
                        }
                    }
                    v35 := GenS1{F0: 42}
                    var v36 int = (v24 & v24)
                    _ = fi31
                    _ = v32
                    _ = v35
                    _ = v36
                }
                var v37 uint16 = 32767
                v38 := GenS3{F0: uint8(v37), F1: 2}
                _ = fi30
                _ = v37
                _ = v38
            }
        case 904:
            for fi39 := 0; fi39 < 14; fi39++ {
                delete(v3, v25)
                var v40 int64 = -529
                for fi41 := 0; fi41 < 13; fi41++ {
                    var v42 int16 = 341
                    fmt.Println(v24)
                    v43 := GenS1{F0: (v42 / (((v42 >> ((v2) & 63))) | 1))}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = fi39
                _ = v40
            }
        case 0:
            switch (v24 | (v24 % (((v24 | v24)) | 1))) {
            case -961:
                switch ("a" + v25) {
                case "foo bar":
                    v1 = 18446744073709551615
                case "":
                    if false {
                        for fi44 := 0; fi44 < 10; fi44++ {
                            var v45 int8 = 100
                            v3["foo bar"] = (v24 >> (((^(v2))) & 63))
                            _ = fi44
                            _ = v45
                        }
                        v24 -= v24
                        var v46 int64 = 1
                        _ = v46
                    } else {
                        var v47 float32 = (float32(-2.25) / 1.5)
                        _ = v47
                    }
                case "\n\t":
                    for fi48 := 0; fi48 < 11; fi48++ {
                        switch fi48 {
                        case -9223372036854775808:
                            var v49 uint8 = 2
                            _ = v49
                        default:
                            fmt.Println(v1, v1, v24)
                        }
                        var v50 uint64 = 4294967295
                        v2 += (v1 - (v2 ^ (v2 + (v1 & 42))))
                        _ = fi48
                        _ = v50
                    }
                default:
                    if (!(42 == 42)) {
                        if false {
                            v1 = v2
                            var v51 float32 = 3.14159
                            _ = v51
                        }
                    } else {
                        for fi52 := 0; fi52 < 5; fi52++ {
                            var v53 int32 = 0
                            _ = fi52
                            _ = v53
                        }
                        for fi54 := 0; fi54 < 1; fi54++ {
                            var v55 uint64 = 18446744073709551615
                            _ = len(v4)
                            fmt.Println(fi54)
                            _ = fi54
                            _ = v55
                        }
                        fmt.Println(v24, v1)
                    }
                }
            case 4294967295:
                v56 := GenS4{F0: -32768, F1: 42}
                _ = v56
            default:
                var v57 uint = 0
                _ = v57
            }
        default:
            for fi58 := 0; fi58 < 6; fi58++ {
                v59 := GenS2{F0: 1, F1: 100}
                switch "foo bar" {
                case "":
                    v24 += fi58
                case "\n\t":
                    var v60 float64 = ((float64(0.0) - 3.14159) - 3.14159)
                    _ = v60
                }
                var v61 string = "\n\t"
                _ = fi58
                _ = v59
                _ = v61
            }
        }
        v2 = 994
        switch "a" {
        case "\n\t":
            var v62 uint32 = 65535
            _ = v62
        case "foo bar":
            fmt.Println(v25, v25, v1)
        case "hello":
            var v63 float64 = ((0.0 - (1e10 * 3.14159)) - (1.5 - ((float64(3.14159) / (3.14159 + 0.0)) + 0.0)))
            _ = v63
        default:
            v64 := GenS2{F0: 1, F1: 42}
            _ = v64
        }
    }
    fmt.Println(v1)
    var v65 float64 = 1e10
    if ((^(v1)) > v2) {
        var v66 bool = (255 >= (^(v2)))
        var v67 string = v25
        var v68 bool = (v67 > "\n\t")
        _ = v66
        _ = v67
        _ = v68
    }
    fmt.Println(v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v65
}
