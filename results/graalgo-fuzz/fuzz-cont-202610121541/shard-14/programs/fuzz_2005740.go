package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 int64 }

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    var v4 int64 = -269
    var v5 int32 = -439
    if (-1 < v5) {
        v6 := []int{2147483647, 42, 410}
        _ = v6
    }
    v4 += 902
    var v7 uint16 = 8
    v7 = 42
    v4 -= 9223372036854775807
    if (42 == 942) {
        v4 += (v4 + (v4 & (v4 - (v4 >> ((uint(v7)) & 63)))))
        var v8 bool = (1 != 873)
        var v9 uint32 = 2147483648
        _ = v8
        _ = v9
    }
    switch ("" + v1) {
    case "foo bar":
        v10 := GenS1{F0: 65535}
        _ = v10
    case "":
        v11 := []int{748, -9223372036854775808}
        _ = v11
    case "\n\t":
        v12 := GenS2{F0: int32(v4)}
        _ = v12
    default:
        var v13 uint16 = v7
        _ = v13
    }
    var v14 int32 = -590
    var v15 uint = 18446744073709551615
    fmt.Println(v5, v15)
    v4 += (^(v4))
    var v16 int64 = -1
    v17 := GenS3{F0: -701, F1: 18446744073709551615}
    var v18 int16 = -821
    var v19 uint64 = 233
    if (true && false) {
        if ((0.1 >= 1.5) && false) {
            v19 += (v19 / ((v19) | 1))
            fmt.Println(v15)
        } else {
            if (((("a" + (((v1 + v1) + "") + v1)) + "hello") == "") || (!(!false))) {
                v3[("foo bar" + (v1 + v1))] = (2147483648 != 1)
                var v20 uint8 = 2
                _ = v20
            } else {
                v21 := GenS4{F0: 42}
                _ = v21
            }
            for fi22 := 0; fi22 < 8; fi22++ {
                v15 = (v15 << ((v15) & 63))
                _ = fi22
            }
        }
        var v23 uint8 = 127
        var v24 uint64 = (^(v19))
        _ = v23
        _ = v24
    } else {
        var v25 bool = (1 <= 34)
        for fi26 := 0; fi26 < 2; fi26++ {
            if v25 {
                switch "hello" {
                case "":
                    v19 = uint64(v7)
                case "hello":
                    v14 = (^(v14))
                case "a":
                    for fi27 := 0; fi27 < 12; fi27++ {
                        switch "hello" {
                        case "":
                            switch (v1 + v1) {
                            case "\n\t":
                                v2[42] = -1
                                v1 = v1
                                switch fi27 {
                                case -951:
                                    var v28 int16 = v18
                                    _ = v28
                                case 127:
                                    switch v1 {
                                    case "\n\t":
                                        v15 = (v15 - (^(v15)))
                                    case "":
                                        v29 := GenS3{F0: 2147483647, F1: (v19 | 42)}
                                        v18 = -1
                                        var v30 float64 = (-2.25 - ((3.14159 + 0.0) + -2.25))
                                        if (false && false) {
                                            var v31 string = "a"
                                            switch fi27 {
                                            case -1:
                                                if ((-(v4)) > 42) {
                                                    fmt.Println(v15, v5, v1)
                                                    if v25 {
                                                        var v32 int8 = 1
                                                        _ = v32
                                                    }
                                                }
                                            case 0:
                                                fmt.Println(fi26, v1, v5)
                                            default:
                                                var v33 int8 = -17
                                                _ = v33
                                            }
                                            _ = v31
                                        } else {
                                            _ = len(v2)
                                            v34 := GenS3{F0: -927, F1: 987}
                                            v35 := []int{32767}
                                            switch (fi26 + 4294967295) {
                                            case -1:
                                                v36 := GenS2{F0: 0}
                                                _ = v36
                                            default:
                                                var v37 uint32 = 4294967295
                                                _ = v37
                                            }
                                            var v38 string = "a"
                                            _ = v34
                                            _ = v35
                                            _ = v38
                                        }
                                        _ = v29
                                        _ = v30
                                    case "a":
                                        if (0 >= 0) {
                                            v39 := []int{(fi27 + fi26), 9223372036854775807, fi26}
                                            switch fi26 {
                                            case -42:
                                                var v40 uint64 = 1
                                                _ = v40
                                            case 42:
                                                v15 = (^(v15))
                                            case 0:
                                                v41 := GenS2{F0: v5}
                                                _ = v41
                                            default:
                                                if len(v39) > 0 {
                                                    v39[0] = fi27
                                                }
                                            }
                                            _ = v39
                                        }
                                    }
                                }
                            case "foo bar":
                                var v42 float64 = -2.25
                                _ = v42
                            case "hello":
                                v43 := GenS2{F0: -2147483648}
                                _ = v43
                            }
                        case "hello":
                            var v44 uint32 = 65535
                            _ = v44
                        case "\n\t":
                            if ((v4 > v16) && v25) {
                                v45 := v2[2147483648]
                                var v46 float64 = (0.0 * 1.5)
                                v46 = v46
                                _ = v45
                                _ = v46
                            }
                        default:
                            v1 = v1
                            var v47 int8 = 0
                            _ = v47
                        }
                        fmt.Println(v15, v25)
                        delete(v2, 2)
                        _ = fi27
                    }
                default:
                    for fi48 := 0; fi48 < 4; fi48++ {
                        v49 := GenS5{F0: ""}
                        v50 := GenS5{F0: v1}
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                }
                var v51 int64 = (v4 | -659)
                _ = len(v3)
                _ = v51
            } else {
                var v52 uint32 = 2147483648
                switch (fi26 % (((fi26 ^ fi26)) | 1)) {
                case -729:
                    var v53 uint64 = v19
                    _ = v53
                case 4294967295:
                    for fi54 := 0; fi54 < 14; fi54++ {
                        if (!v25) {
                            var v55 int32 = (-(v14))
                            if (v5 != (v55 + (v5 / (((v5 * v55)) | 1)))) {
                                v4 = 0
                                v5 = 2147483647
                            }
                            _ = v55
                        } else {
                            v56 := GenS6{F0: (v19 / ((v19) | 1)), F1: v16}
                            if (!(v25 && (((0.1 - (3.14159 / (((0.0 + 1e10) - 0.1) - (0.1 * (((1.5 - (0.0 * (0.1 * (0.0 + 3.14159)))) * 1e10) + (0.1 * float64(float32((1e10 * ((1.5 * -2.25) * (((3.14159 * (1.5 + 1e10)) - float64(3.14159)) + 1.5))))))))))) / 0.1) >= 3.14159))) {
                                v3[v1] = (((((v18 << ((255) & 63)) >= v18) && v25) && false) && v25)
                            }
                            _ = v56
                        }
                        _ = fi54
                    }
                case 1:
                    var v57 int = 683
                    _ = v57
                }
                switch v1 {
                case "":
                    if true {
                        for fi58 := 0; fi58 < 6; fi58++ {
                            var v59 string = v1
                            v18 = v18
                            _ = fi58
                            _ = v59
                        }
                        v1 = ((v1 + v1) + v1)
                    }
                case "hello":
                    for fi60 := 0; fi60 < 18; fi60++ {
                        v61 := GenS1{F0: 246}
                        switch (-(fi26)) {
                        case 1:
                            if v25 {
                                v62 := v2[2147483648]
                                v63 := GenS3{F0: 32767, F1: v19}
                                _ = v62
                                _ = v63
                            }
                        case 4294967295:
                            v64 := v3[("hello" + v1)]
                            _ = v64
                        }
                        var v65 int32 = 42
                        _ = fi60
                        _ = v61
                        _ = v65
                    }
                case "\n\t":
                    var v66 uint8 = 0
                    _ = v66
                default:
                    v14 += (v14 << ((18446744073709551615) & 63))
                }
                _ = v52
            }
            _ = fi26
        }
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
