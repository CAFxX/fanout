package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint8; F1 uint16; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 int8; F1 int; F2 uint32 }
type GenS9 struct { F0 int32 }
func genFn0() bool {
    v4 := []int{}
    switch "\n\t" {
    case "":
        if false {
            var v5 string = "hello"
            _ = v5
        } else {
            v6 := GenS1{F0: -1, F1: -182, F2: 246}
            _ = v6
        }
    case "a":
        if (float64(float32(1e10)) <= float64(-2.25)) {
            v7 := GenS2{F0: 0, F1: 492}
            var v8 float64 = (float64(-2.25) / 3.14159)
            switch -2147483648 {
            case 127:
                fmt.Println(v8)
            case 512:
                for fi9 := 0; fi9 < 5; fi9++ {
                    v8 += 0.1
                    _ = fi9
                }
            case 2147483647:
                fmt.Println(v8)
            default:
                v10 := GenS3{F0: 237, F1: 239, F2: -674}
                _ = v10
            }
            _ = v7
            _ = v8
        }
    default:
        switch "hello" {
        case "":
            if false {
                if (1.5 <= float64(1e10)) {
                    var v11 float32 = 3.14159
                    var v12 int32 = 0
                    v4 = append(v4, -419)
                    _ = v11
                    _ = v12
                }
                v13 := GenS4{F0: true}
                var v14 float32 = 1.5
                _ = v13
                _ = v14
            } else {
                var v15 int8 = 127
                _ = len(v4)
                _ = v15
            }
        case "foo bar":
            var v16 uint16 = 1
            _ = v16
        case "hello":
            switch ("a" + "a") {
            case "a":
                var v17 uint = 65535
                _ = v17
            case "foo bar":
                v18 := GenS1{F0: 35, F1: 32767, F2: 0}
                _ = v18
            case "\n\t":
                v19 := GenS2{F0: 128, F1: 937}
                _ = v19
            default:
                var v20 uint16 = 42
                _ = v20
            }
        }
    }
    if (467 < -32768) {
        var v21 string = "\n\t"
        _ = v21
    }
    _ = v4
    return false
}

func main() {
    var v1 int = 128
    var v2 int64 = 0
    var v3 int8 = -1
    v2 = v2
    var v22 int32 = 1
    switch "\n\t" {
    case "\n\t":
        v23 := GenS5{F0: 0, F1: 18446744073709551615, F2: 127}
        _ = v23
    }
    if true {
        v24 := []int{-507, (-(v1)), int(v3), v1}
        _ = v24
    }
    v2 += v2
    v22 = v22
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            fmt.Println(v22)
            var v27 float64 = float64(-2.25)
            switch "a" {
            case "foo bar":
                v28 := GenS6{F0: (v22 | v22)}
                _ = v28
            default:
                if true {
                    if (((1e10 / (-2.25 + ((((0.1 + (1e10 * 1.5)) + ((0.0 + float32(v27)) * float32((0.0 * v27)))) - (1e10 + 1.5)) / (0.1 - 1e10)))) <= 0.1) && (0.0 != (0.1 + float64(0.0)))) {
                        v29 := []int{(fi26 * (v1 << ((317) & 63))), (v1 / ((-1) | 1))}
                        v30 := GenS4{F0: true}
                        _ = v29
                        _ = v30
                    }
                    var v31 uint32 = 1
                    _ = v31
                }
            }
            _ = fi26
            _ = v27
        }
        v1 -= (v1 * (v1 * fi25))
        _ = fi25
    }
    if (!((255 != 0) && true)) {
        v32 := GenS2{F0: 148, F1: 312}
        for fi33 := 0; fi33 < 4; fi33++ {
            var v34 float32 = (float32((((float64(float32((3.14159 + 1.5))) / ((float64(0.0) - 0.1) + float64(0.0))) / 0.1) / (1e10 / -2.25))) - 3.14159)
            switch ("a" + "hello") {
            case "":
                switch (fi33 - 32768) {
                case 32768:
                    switch "a" {
                    case "hello":
                        var v35 int8 = -1
                        _ = v35
                    default:
                        v36 := GenS7{F0: 368, F1: 104}
                        _ = v36
                    }
                case 2147483647:
                    if (true || (false && false)) {
                        var v37 float64 = 1e10
                        var v38 int16 = -32768
                        for fi39 := 0; fi39 < 2; fi39++ {
                            fmt.Println(fi33, v22)
                            v40 := GenS8{F0: v3, F1: (v1 + 42), F2: 4294967295}
                            var v41 int8 = (v3 * (-(v3)))
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        _ = v37
                        _ = v38
                    } else {
                        v3 = v3
                        if false {
                            if false {
                                v42 := GenS9{F0: (v22 + (v22 >> ((18446744073709551615) & 63)))}
                                _ = v42
                            } else {
                                fmt.Println(v22)
                                if true {
                                    var v43 uint64 = 0
                                    fmt.Println(v34)
                                    v3 = -1
                                    _ = v43
                                } else {
                                    switch fi33 {
                                    case 283:
                                        for fi44 := 0; fi44 < 7; fi44++ {
                                            v45 := []int{(fi33 % ((9223372036854775807) | 1)), fi33}
                                            _ = fi44
                                            _ = v45
                                        }
                                    case 3:
                                        var v46 float32 = 1e10
                                        _ = v46
                                    default:
                                        var v47 int32 = v22
                                        _ = v47
                                    }
                                }
                                var v48 int = fi33
                                _ = v48
                            }
                            fmt.Println(v34, v3, v2)
                        } else {
                            var v49 uint64 = 18446744073709551615
                            v1 = (fi33 >> ((4294967295) & 63))
                            _ = v49
                        }
                        var v50 float32 = v34
                        _ = v50
                    }
                case 942:
                    v51 := []int{(v1 | v1), 642, v1}
                    _ = v51
                default:
                    var v52 int32 = (v22 | (v22 / ((47) | 1)))
                    _ = v52
                }
            case "foo bar":
                if false {
                    v22 -= (v22 | v22)
                    v34 += (float32(1.5) + v34)
                }
            }
            v53 := []int{v1, (v1 | fi33)}
            _ = fi33
            _ = v34
            _ = v53
        }
        _ = v32
    }
    v54 := []int{}
    var v55 int = 2147483647
    if false {
        v54 = append(v54, 306)
    } else {
        if false {
            v54 = append(v54, (v55 * (v1 + v55)))
            v56 := GenS8{F0: 8, F1: 9223372036854775807, F2: 2}
            _ = v56
        }
        var v57 float32 = (1e10 * (0.1 - 1e10))
        if (true && ("foo bar" == ("a" + "hello"))) {
            for fi58 := 0; fi58 < 1; fi58++ {
                switch ("" + "hello") {
                case "\n\t":
                    _ = len(v54)
                default:
                    v54 = append(v54, (v1 % (((v55 / ((fi58) | 1))) | 1)))
                }
                _ = fi58
            }
        } else {
            var v59 int = (v55 | (v1 + v55))
            var v60 uint = 18446744073709551615
            _ = v59
            _ = v60
        }
        _ = v57
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v54
    _ = v55
}
