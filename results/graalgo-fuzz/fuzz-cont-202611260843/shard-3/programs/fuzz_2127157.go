package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint8) uint16 {
    v6 := GenS1{F0: 0, F1: -32768, F2: -876}
    fmt.Println(p0)
    v7 := []int{32768, 32767, -9223372036854775808, -505}
    var v8 int32 = 42
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint64, p1 int8, p2 uint32) string {
    v9 := []int{}
    for fi10 := 0; fi10 < 6; fi10++ {
        _ = len(v9)
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "\n\t"
}

func main() {
    var v1 uint64 = 948
    var v2 int32 = -418
    var v3 int8 = 111
    var v4 int32 = v2
    v5 := make(map[uint32]int)
    v12 := GenS2{F0: v3, F1: v3, F2: 128}
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint8 = 42
        v15 := []int{(^(fi13)), 4294967295}
        if (true && (!false)) {
            for fi16 := 0; fi16 < 19; fi16++ {
                if (!(!true)) {
                    v1 += (v1 ^ (v1 + 2))
                    var v17 uint64 = v1
                    _ = v17
                }
                _ = fi16
            }
            switch "foo bar" {
            case "":
                for fi18 := 0; fi18 < 15; fi18++ {
                    v3 = 1
                    _ = fi18
                }
            case "a":
                v19 := GenS1{F0: v2, F1: -172, F2: 0}
                _ = v19
            default:
                fmt.Println(v1, v2, v1)
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                if true {
                    v2 = (-(v4))
                    _ = len(v5)
                }
            case "\n\t":
                var v20 float32 = -2.25
                _ = v20
            }
            delete(v5, 1)
            for fi21 := 0; fi21 < 7; fi21++ {
                _ = len(v5)
                var v22 float32 = 1e10
                if (1 >= 0) {
                    switch fi13 {
                    case 4294967295:
                        if false {
                            if false {
                                var v23 string = "\n\t"
                                var v24 float32 = (v22 - ((0.0 * v22) - ((-2.25 - ((3.14159 + 1.5) - 1e10)) * ((((v22 - 1.5) - v22) * v22) - v22))))
                                _ = v23
                                _ = v24
                            } else {
                                if ((-(v3)) >= 127) {
                                    if (v14 == (v14 & (v14 & (^(v14))))) {
                                        v2 -= v2
                                        var v25 uint32 = 65535
                                        _ = v25
                                    }
                                } else {
                                    var v26 uint32 = 65535
                                    var v27 float64 = (1.5 / float64(0.1))
                                    _ = v26
                                    _ = v27
                                }
                            }
                        }
                    case 720:
                        v1 -= 559
                    }
                    if (0.0 != (-2.25 + (float64(float32(3.14159)) * float64(((((v22 * float32((1.5 + ((0.1 / 0.1) / 1.0)))) + v22) - 3.14159) * v22))))) {
                        v2 = (^(v2))
                    } else {
                        fmt.Println(v22)
                    }
                    if false {
                        v28 := GenS3{F0: v1, F1: 42, F2: v14}
                        var v29 float32 = v22
                        _ = v28
                        _ = v29
                    } else {
                        if false {
                            v14 = v14
                            switch 42 {
                            case -2147483648:
                                var v30 int64 = -9223372036854775808
                                _ = v30
                            default:
                                if true {
                                    v15 = append(v15, (^(fi13)))
                                    var v31 float32 = v22
                                    _ = v31
                                } else {
                                    var v32 int64 = -653
                                    _ = v32
                                }
                            }
                        }
                        v4 -= v2
                        v15 = append(v15, 128)
                    }
                } else {
                    if ((true && false) || false) {
                        v33 := GenS3{F0: 2, F1: int16(v14), F2: (v14 % ((247) | 1))}
                        v4 -= -1
                        _ = v33
                    }
                    v1 -= 710
                }
                _ = fi21
                _ = v22
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 = v4
    for fi34 := 0; fi34 < 13; fi34++ {
        var v35 uint32 = 42
        v1 -= 1
        _ = fi34
        _ = v35
    }
    if (true || false) {
        for fi36 := 0; fi36 < 6; fi36++ {
            v37 := GenS4{F0: (v2 | 829)}
            var v38 uint16 = 32767
            _ = fi36
            _ = v37
            _ = v38
        }
        var v39 int16 = 32767
        if false {
            v39 = (^(v39))
            v40 := GenS2{F0: 0, F1: (v3 & int8(v1)), F2: 2147483647}
            v39 = 1
            _ = v40
        }
        _ = v39
    }
    fmt.Println(v2, v1)
    for fi41 := 0; fi41 < 12; fi41++ {
        var v42 int64 = 9223372036854775807
        for fi43 := 0; fi43 < 3; fi43++ {
            var v44 int64 = (-(v42))
            v45 := []int{fi41, fi43, fi43}
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = fi41
        _ = v42
    }
    if false {
        switch "\n\t" {
        case "hello":
            switch 1 {
            case -1:
                var v46 uint = 255
                _ = v46
            case 2147483647:
                var v47 float32 = 0.1
                _ = v47
            default:
                v2 += v2
            }
        case "":
            switch -856 {
            case 4294967295:
                v2 -= 1
            }
        }
        var v48 int32 = (v2 * -746)
        _ = v48
    } else {
        v2 += (v4 ^ (v4 % ((0) | 1)))
        v49 := []int{32768, 0}
        for fi50 := 0; fi50 < 10; fi50++ {
            v4 -= int32(fi50)
            _ = fi50
        }
        _ = v49
    }
    v2 += v4
    for fi51 := 0; fi51 < 17; fi51++ {
        v52 := GenS1{F0: v2, F1: -1, F2: fi51}
        if (!(!true)) {
            fmt.Println(v2, v4, v4)
            var v53 string = ("hello" + ("a" + (((("foo bar" + "\n\t") + "") + "a") + "")))
            var v54 string = ((((v53 + (v53 + v53)) + ("" + "a")) + ("" + "")) + v53)
            _ = v53
            _ = v54
        } else {
            var v55 uint32 = 65535
            _ = v55
        }
        v4 += 2147483647
        _ = fi51
        _ = v52
    }
    var v56 uint32 = 4294967295
    v4 = v2
    v56 = (v56 % ((v56) | 1))
    v3 -= (v3 << ((255) & 63))
    v1 -= 763
    if true {
        switch 128 {
        case -9223372036854775808:
            v56 -= (v56 * v56)
        default:
            for fi57 := 0; fi57 < 4; fi57++ {
                v58 := []int{(fi57 >> ((1) & 63)), (fi57 >> ((uint(v1)) & 63))}
                if (2147483647 <= v56) {
                    for fi59 := 0; fi59 < 18; fi59++ {
                        v60 := GenS4{F0: v2}
                        _ = fi59
                        _ = v60
                    }
                } else {
                    var v61 uint64 = 562
                    _ = v61
                }
                _ = fi57
                _ = v58
            }
        }
        fmt.Println(v3)
        v62 := []int{127, -401, 32767, -832}
        _ = v62
    } else {
        v1 = v1
        v2 = (v2 >> ((2) & 63))
        v63 := []int{32768, 127}
        _ = v63
    }
    var v64 uint64 = (v1 | v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v64
}
