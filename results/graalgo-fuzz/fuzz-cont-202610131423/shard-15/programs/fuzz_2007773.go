package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0(p0 int32, p1 bool, p2 uint8) int64 {
    fmt.Println(p0, p0, p1)
    p2 = (p2 | p2)
    switch 127 {
    case -505:
        switch 2147483647 {
        case 127:
            for fi4 := 0; fi4 < 6; fi4++ {
                p0 = (^(p0))
                _ = fi4
            }
        case -9223372036854775808:
            var v5 uint8 = p2
            _ = v5
        default:
            var v6 float64 = (-2.25 / 1e10)
            _ = v6
        }
    case 127:
        var v7 uint8 = 1
        _ = v7
    case -611:
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := []int{fi8}
            var v10 bool = p1
            var v11 string = "hello"
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    default:
        fmt.Println(p2, p1)
    }
    v12 := []int{127, 2147483647, 32768, 32767}
    if true {
        fmt.Println(p2, p1, p2)
        v13 := GenS1{F0: (p0 % ((p0) | 1))}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 42
}
func genFn1(p0 int16, p1 int8, p2 uint32) uint16 {
    var v14 int8 = (p1 + (p1 << ((1) & 63)))
    for fi15 := 0; fi15 < 8; fi15++ {
        if ("\n\t" == (("" + "foo bar") + "a")) {
            p2 -= (p2 * 65535)
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 42
}

func main() {
    var v1 uint = 0
    var v2 uint32 = 2147483648
    var v3 uint32 = v2
    for fi16 := 0; fi16 < 9; fi16++ {
        if false {
            v3 += 1
            fmt.Println(v1, fi16)
            var v17 uint = v1
            _ = v17
        }
        switch "foo bar" {
        case "hello":
            switch (fi16 & (fi16 << ((42) & 63))) {
            case 42:
                switch "" {
                case "a":
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 float32 = 0.0
                        _ = fi18
                        _ = v19
                    }
                case "hello":
                    if false {
                        v2 = (^(v3))
                        for fi20 := 0; fi20 < 11; fi20++ {
                            switch "" {
                            case "foo bar":
                                switch ("" + "hello") {
                                case "":
                                    var v21 uint8 = 128
                                    _ = v21
                                default:
                                    fmt.Println(fi16, fi16, fi16)
                                }
                                v22 := GenS1{F0: 1}
                                _ = v22
                            case "hello":
                                fmt.Println(v3, fi20)
                                switch "" {
                                case "foo bar":
                                    switch fi16 {
                                    case -929:
                                        v23 := []int{0}
                                        _ = v23
                                    case -298:
                                        fmt.Println(v3, fi20, v1)
                                    case -1:
                                        v24 := []int{(fi16 | 42)}
                                        _ = v24
                                    default:
                                        switch (fi20 - fi20) {
                                        case 2147483647:
                                            if (genFn1(-1, 42, (v2 % (((v2 | 92)) | 1))) < 2) {
                                                if (127 != uint8(v3)) {
                                                    var v25 int8 = 125
                                                    if ((fi16 == 731) && (1e10 < ((0.1 / ((float64(1.5) + 0.0) * -2.25)) + 1.5))) {
                                                        v25 = 127
                                                        v26 := []int{fi16, 776, 127, -1}
                                                        v26 = append(v26, (-(fi20)))
                                                        var v27 uint8 = 128
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                    _ = v25
                                                }
                                                v1 = v1
                                                var v28 int16 = 1000
                                                _ = v28
                                            }
                                        case 128:
                                            v3 += v3
                                        case 281:
                                            v29 := []int{(-(fi16)), (-(fi16)), (-(fi16))}
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                            _ = v29
                                        }
                                    }
                                default:
                                    v30 := []int{fi20}
                                    _ = v30
                                }
                            case "\n\t":
                                fmt.Println(fi20, fi20, v1)
                            default:
                                var v31 uint = 255
                                _ = v31
                            }
                            v32 := []int{}
                            v1 = (v1 ^ (v1 + v1))
                            _ = fi20
                            _ = v32
                        }
                        v33 := GenS1{F0: 566}
                        _ = v33
                    }
                case "":
                    for fi34 := 0; fi34 < 12; fi34++ {
                        var v35 float32 = (-2.25 * ((1.5 + (((0.1 + (3.14159 * (0.0 - float32(0.1)))) - (0.1 - 0.1)) - 0.1)) - 1.5))
                        v36 := GenS2{F0: -1, F1: (fi34 + -1), F2: 127}
                        v37 := []int{fi34, fi16}
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                default:
                    v1 = (^(v1))
                }
            default:
                v38 := []int{0}
                _ = v38
            }
        default:
            fmt.Println(v3)
        }
        _ = fi16
    }
    var v39 int = 1
    v40 := []int{v39, v39, v39, (v39 ^ (v39 << ((v1) & 63)))}
    v39 = v39
    if false {
        var v41 float32 = (-2.25 + (float32(((1.5 - 1e10) - -2.25)) * 1e10))
        var v42 bool = false
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        if len(v40) > 0 {
            v40[0] = 32767
        }
        v3 = (v3 & (v3 >> (((^(v1))) & 63)))
        _ = fi43
    }
    for fi44 := 0; fi44 < 1; fi44++ {
        v45 := GenS1{F0: -2147483648}
        v1 -= (v1 | (v1 - (^(v1))))
        var v46 bool = (((881 <= -597) || true) || (1 == 127))
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println(v3, v2)
    v47 := GenS3{F0: 128, F1: -32768}
    v3 -= (v2 % (((v3 ^ uint32(v1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v47
}
