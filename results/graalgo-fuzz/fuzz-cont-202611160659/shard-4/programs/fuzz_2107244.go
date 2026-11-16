package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16 }
func genFn0(p0 bool) uint64 {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 uint8 = 42
        if p0 {
            v6 := GenS1{F0: (42 != 18446744073709551615), F1: false}
            v7 := GenS2{F0: (v5 ^ (v5 >> ((856) & 63))), F1: -9223372036854775808}
            var v8 uint32 = 0
            _ = v6
            _ = v7
            _ = v8
        } else {
            var v9 uint64 = uint64(v5)
            p0 = p0
            v10 := []int{fi4, fi4, 1, fi4}
            _ = v10
            _ = v9
        }
        v11 := []int{-1}
        _ = fi4
        _ = v11
        _ = v5
    }
    if true {
        switch "\n\t" {
        case "foo bar":
            var v12 int32 = 412
            _ = v12
        case "hello":
            switch "foo bar" {
            case "foo bar":
                if p0 {
                    fmt.Println(p0)
                    switch ("hello" + "a") {
                    case "hello":
                        p0 = p0
                    case "a":
                        for fi13 := 0; fi13 < 8; fi13++ {
                            var v14 uint64 = 42
                            for fi15 := 0; fi15 < 3; fi15++ {
                                v16 := []int{(fi15 >> ((700) & 63))}
                                p0 = (42 != 0)
                                _ = fi15
                                _ = v16
                            }
                            fmt.Println(v14)
                            _ = fi13
                            _ = v14
                        }
                    case "foo bar":
                        switch 9223372036854775807 {
                        case 32768:
                            fmt.Println(p0)
                        default:
                            for fi17 := 0; fi17 < 4; fi17++ {
                                v18 := []int{(fi17 / ((fi17) | 1)), 372, fi17}
                                p0 = (72 != 0)
                                switch ("" + "hello") {
                                case "a":
                                    if len(v18) > 0 {
                                        v18[0] = fi17
                                    }
                                case "hello":
                                    for fi19 := 0; fi19 < 9; fi19++ {
                                        fmt.Println(p0, p0, fi17)
                                        var v20 int32 = 42
                                        fmt.Println(v20, p0, fi17)
                                        _ = fi19
                                        _ = v20
                                    }
                                }
                                _ = fi17
                                _ = v18
                            }
                        }
                    default:
                        p0 = p0
                    }
                }
            case "":
                if (65535 < 2) {
                    var v21 int64 = -9223372036854775808
                    var v22 uint = 18446744073709551615
                    v23 := []int{42, 360, -1}
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    if p0 {
                        var v24 bool = (true || p0)
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 18; fi25++ {
                            v26 := GenS3{F0: fi25}
                            _ = fi25
                            _ = v26
                        }
                    }
                }
            case "\n\t":
                var v27 int64 = 1
                _ = v27
            }
        case "":
            p0 = (!p0)
        default:
            v28 := GenS2{F0: 1, F1: 1}
            _ = v28
        }
        v29 := []int{}
        v30 := GenS4{F0: 127}
        _ = v29
        _ = v30
    } else {
        var v31 uint64 = 2
        v32 := GenS3{F0: 0}
        _ = v31
        _ = v32
    }
    p0 = p0
    var v33 int32 = -2147483648
    p0 = true
    _ = p0
    _ = v33
    return 0
}
func genFn1(p0 uint64, p1 uint32) int {
    if false {
        if true {
            for fi34 := 0; fi34 < 4; fi34++ {
                var v35 int8 = 0
                _ = fi34
                _ = v35
            }
            v36 := GenS3{F0: 9223372036854775807}
            _ = v36
        }
    }
    if (42 > 255) {
        if false {
            v37 := GenS5{F0: "", F1: p0}
            _ = v37
        }
        p0 -= genFn0((208 > 65535))
    }
    if true {
        fmt.Println(p1, p0)
        var v38 uint64 = p0
        _ = v38
    } else {
        var v39 uint32 = 2147483648
        var v40 uint = 307
        _ = v39
        _ = v40
    }
    var v41 int64 = -1
    p1 -= (p1 % ((p1) | 1))
    _ = p0
    _ = p1
    _ = v41
    return -2147483648
}

func main() {
    var v1 uint = 65535
    var v2 uint16 = 42
    v3 := []int{491, 2147483647}
    v42 := GenS6{F0: 127}
    if len(v3) > 0 {
        v3[0] = 797
    }
    var v43 int16 = int16(v2)
    v43 += v43
    v44 := GenS7{F0: 9223372036854775807}
    var v45 bool = ((^(v2)) != (v2 & (^(v2))))
    v43 = (v43 * (v43 + -81))
    v46 := GenS5{F0: "\n\t", F1: 868}
    switch 9223372036854775807 {
    case 603:
        if ((1.5 + (float32(-2.25) - (float32(1.5) * 1.5))) >= (((((3.14159 * (float32(float64(1.5)) - 1e10)) / -2.25) * 3.14159) + (3.14159 / 0.1)) + 0.1)) {
            if v45 {
                for fi47 := 0; fi47 < 3; fi47++ {
                    var v48 bool = ((v2 >> (((v1 / (((v1 | 0)) | 1))) & 63)) != 133)
                    _ = fi47
                    _ = v48
                }
                if (v45 && ((9223372036854775807 >= 9223372036854775807) && true)) {
                    v49 := GenS8{F0: -1}
                    if v45 {
                        if (v43 < -768) {
                            v45 = (0 >= 147)
                            for fi50 := 0; fi50 < 2; fi50++ {
                                switch "foo bar" {
                                case "\n\t":
                                    v51 := GenS9{F0: 42}
                                    _ = v51
                                case "a":
                                    if v45 {
                                        var v52 int8 = int8(v1)
                                        fmt.Println(fi50)
                                        _ = v52
                                    } else {
                                        var v53 int16 = (v43 | (v43 / (((-(v43))) | 1)))
                                        for fi54 := 0; fi54 < 14; fi54++ {
                                            fmt.Println(v1, v2)
                                            fmt.Println(fi50, v43)
                                            var v55 float32 = 0.0
                                            _ = fi54
                                            _ = v55
                                        }
                                        var v56 int32 = 1
                                        _ = v53
                                        _ = v56
                                    }
                                case "hello":
                                    v45 = (v45 || true)
                                }
                                v57 := GenS2{F0: 2, F1: -9223372036854775808}
                                _ = fi50
                                _ = v57
                            }
                            switch "" {
                            case "":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                v3 = append(v3, -571)
                            }
                        }
                        var v58 uint8 = uint8(v2)
                        _ = v58
                    }
                    v3 = append(v3, 9223372036854775807)
                    _ = v49
                } else {
                    for fi59 := 0; fi59 < 10; fi59++ {
                        if true {
                            fmt.Println(v45)
                            if (18446744073709551615 != 0) {
                                var v60 uint16 = (v2 | v2)
                                _ = v60
                            } else {
                                for fi61 := 0; fi61 < 10; fi61++ {
                                    var v62 int16 = -32768
                                    _ = fi61
                                    _ = v62
                                }
                            }
                        } else {
                            v2 += (v2 << (((v1 - v1)) & 63))
                        }
                        _ = fi59
                    }
                    fmt.Println(v2)
                }
                if (!(!(((808 != 2) && v45) && (42 > 42)))) {
                    var v63 float64 = 3.14159
                    v45 = false
                    _ = v63
                } else {
                    v3 = append(v3, 0)
                    if (1 > 218) {
                        for fi64 := 0; fi64 < 13; fi64++ {
                            var v65 float32 = 3.14159
                            _ = fi64
                            _ = v65
                        }
                        var v66 uint = v1
                        _ = v66
                    } else {
                        if len(v3) > 0 {
                            v3[0] = -512
                        }
                        var v67 float32 = 1.5
                        v68 := GenS4{F0: 127}
                        _ = v67
                        _ = v68
                    }
                }
            } else {
                v3 = append(v3, 0)
            }
        }
    case -885:
        if ((-98 >= 127) || ((float32(0.0) != 3.14159) || (true || false))) {
            v69 := GenS1{F0: (!v45), F1: v45}
            var v70 int = -169
            _ = v69
            _ = v70
        } else {
            var v71 uint32 = 0
            var v72 int64 = -1
            if (1 >= -2147483648) {
                v1 = 4294967295
                var v73 bool = (true && false)
                _ = v73
            } else {
                v45 = v45
            }
            _ = v71
            _ = v72
        }
    case -9223372036854775808:
        fmt.Println(v1)
    default:
        v74 := GenS2{F0: 0, F1: 9223372036854775807}
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
