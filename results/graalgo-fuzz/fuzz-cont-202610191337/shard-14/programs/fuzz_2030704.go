package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int; F2 int8 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint32, p2 uint64) uint64 {
    var v5 int64 = 9223372036854775807
    v6 := GenS1{F0: 1, F1: 42}
    v7 := []int{42, -683, -751, -672}
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (p2 * uint64(p1))
}
func genFn1(p0 string, p1 int) uint16 {
    v8 := GenS1{F0: 128, F1: -128}
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int64 = 479
        fmt.Println(fi9, fi9, p0)
        _ = fi9
        _ = v10
    }
    v11 := []int{p1}
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return 843
}

func main() {
    v1 := []string{(("hello" + "\n\t") + ""), "foo bar", ("hello" + ((((("\n\t" + "\n\t") + "") + "foo bar") + ("foo bar" + ("foo bar" + "\n\t"))) + ("a" + ("\n\t" + ("" + (("\n\t" + "a") + ("hello" + "foo bar"))))))), "hello"}
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    v4 := make(map[uint32]bool)
    var v12 uint32 = 4294967295
    switch 128 {
    case -9223372036854775808:
        v12 = (v12 % ((v12) | 1))
    case 1:
        v13 := GenS2{F0: "hello", F1: 2}
        _ = v13
    }
    var v14 string = ("\n\t" + "a")
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: 1, F1: 1}
        if false {
            switch v14 {
            case "":
                if (!(42 >= 0)) {
                    switch v14 {
                    case "foo bar":
                        switch (fi15 + fi15) {
                        case 2147483647:
                            fmt.Println(fi15, v12)
                        case -273:
                            var v17 float32 = 3.14159
                            _ = v17
                        case -594:
                            var v18 string = v14
                            _ = v18
                        }
                    case "\n\t":
                        var v19 int32 = 0
                        _ = v19
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v12 = (v12 / ((2147483648) | 1))
                }
            }
            if (!false) {
                v14 = v14
                var v20 int = fi15
                _ = v20
            } else {
                v21 := GenS3{F0: -1}
                _ = v21
            }
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
        } else {
            v12 -= (v12 << ((480) & 63))
            fmt.Println(v12)
        }
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 int8 = -91
            if ((false && (2 < 255)) && (true && false)) {
                v12 = v12
                switch v14 {
                case "a":
                    if false {
                        switch fi15 {
                        case -1:
                            v14 = (v14 + v14)
                            v24 := GenS4{F0: (("hello" + "a") + "\n\t"), F1: fi15, F2: 47}
                            _ = v24
                        default:
                            switch (fi15 * (-(fi15))) {
                            case 4294967295:
                                v25 := GenS4{F0: v14, F1: fi22, F2: int8(fi22)}
                                _ = v25
                            case 42:
                                var v26 int8 = -100
                                _ = v26
                            case -539:
                                fmt.Println(v23)
                            }
                        }
                        if (2 != 255) {
                            if false {
                                if false {
                                    v27 := GenS3{F0: 781}
                                    var v28 uint = 0
                                    _ = v27
                                    _ = v28
                                } else {
                                    if (65535 == 4294967295) {
                                        if true {
                                            switch int(v23) {
                                            case 128:
                                                var v29 float32 = 1.5
                                                _ = v29
                                            case -9223372036854775808:
                                                var v30 bool = true
                                                _ = v30
                                            default:
                                                v14 = (v14 + ((v14 + "foo bar") + v14))
                                            }
                                            delete(v4, v12)
                                        }
                                        v12 = v12
                                        var v31 int64 = int64(fi22)
                                        v14 = ("a" + "foo bar")
                                        switch -2147483648 {
                                        case 9223372036854775807:
                                            if (false || (!false)) {
                                                delete(v2, 127)
                                            } else {
                                                _ = len(v1)
                                                var v32 int = (fi15 | int(v23))
                                                var v33 int16 = -587
                                                _ = v32
                                                _ = v33
                                            }
                                        case 0:
                                            switch v14 {
                                            case "a":
                                                v23 += 46
                                            default:
                                                v4[v12] = false
                                            }
                                        }
                                        _ = len(v2)
                                        _ = v31
                                    } else {
                                        v34 := GenS2{F0: ("" + "foo bar"), F1: 787}
                                        var v35 int64 = 0
                                        v23 += (v23 / ((v23) | 1))
                                        _ = v34
                                        _ = v35
                                    }
                                    v14 = (v14 + v14)
                                }
                                var v36 int16 = -1
                                _ = v36
                            } else {
                                fmt.Println(fi22)
                                v12 += v12
                                switch v14 {
                                case "a":
                                    var v37 int32 = 2147483647
                                    _ = v37
                                case "":
                                    v38 := v4[4294967295]
                                    _ = v38
                                }
                            }
                            v39 := GenS2{F0: v14, F1: v12}
                            if len(v3) > 0 {
                                v3[0] = (v12 >> ((255) & 63))
                            }
                            _ = v39
                        } else {
                            var v40 uint32 = (v12 >> ((4294967295) & 63))
                            _ = v40
                        }
                        v23 = v23
                    } else {
                        v12 = 24
                        var v41 float64 = 3.14159
                        var v42 uint32 = v12
                        _ = v41
                        _ = v42
                    }
                case "":
                    v14 = v14
                case "foo bar":
                    v14 = v14
                }
            } else {
                v43 := GenS5{F0: -755}
                if (1e10 == (1e10 - 1e10)) {
                    if true {
                        fmt.Println(v12)
                    } else {
                        switch (fi22 / (((fi15 * fi22)) | 1)) {
                        case -9223372036854775808:
                            v14 = (("" + (v14 + v14)) + "a")
                            var v44 uint8 = 2
                            _ = len(v4)
                            _ = v44
                        case 1:
                            var v45 uint8 = 155
                            fmt.Println(v23)
                            var v46 int16 = 1
                            _ = v45
                            _ = v46
                        }
                    }
                    v47 := GenS4{F0: v14, F1: fi15, F2: v23}
                    if true {
                        var v48 bool = true
                        var v49 uint16 = 443
                        _ = v48
                        _ = v49
                    }
                    _ = v47
                } else {
                    fmt.Println(v12, v23)
                    v12 += v12
                }
                if ((v12 ^ (^(v12))) != v12) {
                    v2[(^(fi15))] = 243
                    var v50 int64 = 552
                    v23 += v23
                    _ = v50
                }
                v12 -= v12
                _ = v43
            }
            var v51 float64 = 0.1
            _ = fi22
            _ = v23
            _ = v51
        }
        _ = fi15
        _ = v16
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v14)
    var v52 uint = 721
    var v53 int8 = 127
    for fi54 := 0; fi54 < 8; fi54++ {
        if true {
            for fi55 := 0; fi55 < 18; fi55++ {
                var v56 int32 = -702
                _ = fi55
                _ = v56
            }
            v12 = (v12 / ((v12) | 1))
        }
        _ = fi54
    }
    switch "\n\t" {
    case "":
        v53 = -1
    }
    var v57 uint64 = 134
    v14 = ("" + v14)
    v57 = (v57 % ((v57) | 1))
    if true {
        for fi58 := 0; fi58 < 6; fi58++ {
            if (true || (false || true)) {
                v52 = (v52 / ((991) | 1))
                v57 -= (v57 / (((^(v57))) | 1))
            } else {
                v59 := GenS3{F0: 175}
                v53 += (v53 >> (((v52 / ((v52) | 1))) & 63))
                var v60 int16 = 1000
                _ = v59
                _ = v60
            }
            var v61 bool = false
            _ = fi58
            _ = v61
        }
        fmt.Println(v52, v12, v53)
        delete(v4, (v12 + v12))
    }
    if (1 == 1) {
        fmt.Println(v52, v57)
        switch v14 {
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            if false {
                fmt.Println(v52)
            }
        case "a":
            var v62 float32 = 0.0
            _ = v62
        default:
            v57 -= v57
        }
    } else {
        _ = len(v2)
        switch 449 {
        case 788:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 4294967295:
            var v63 uint = v52
            _ = v63
        case -2147483648:
            switch 356 {
            case 9223372036854775807:
                v57 = (v57 % ((0) | 1))
            default:
                v64 := GenS2{F0: v14, F1: v12}
                _ = v64
            }
        default:
            var v65 float32 = 3.14159
            _ = v65
        }
        var v66 uint16 = 42
        _ = v66
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v57
}
