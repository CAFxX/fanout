package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int16; F2 int64 }
type GenS5 struct { F0 int8; F1 int8; F2 uint }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int32, p1 string) int {
    var v4 float64 = 1.5
    v5 := GenS1{F0: 32767, F1: (p1 + p1), F2: 100}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    var v1 uint64 = 42
    var v2 int16 = 0
    v3 := []int{-980, 9223372036854775807}
    fmt.Println(v2)
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 bool = false
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: 2, F1: "\n\t", F2: 1}
        if (3.14159 >= float64((((float32(1.5) - 0.1) / 0.1) * ((0.0 / (0.0 + 0.1)) - 1e10)))) {
            v3 = append(v3, 127)
            v2 = v2
        }
        switch fi8 {
        case 1:
            var v10 uint16 = 32767
            _ = v10
        default:
            for fi11 := 0; fi11 < 13; fi11++ {
                switch genFn0(2147483647, "foo bar") {
                case 2147483647:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 32768:
                    v2 += (v2 >> ((84) & 63))
                default:
                    var v12 uint64 = 666
                    _ = v12
                }
                _ = fi11
            }
        }
        _ = fi8
        _ = v9
    }
    v2 += (v2 | v2)
    v3 = append(v3, 127)
    switch -103 {
    case -1:
        for fi13 := 0; fi13 < 5; fi13++ {
            _ = len(v3)
            _ = fi13
        }
    default:
        if (true && (65535 >= 688)) {
            if len(v3) > 0 {
                v3[0] = -1
            }
            v1 -= v1
        }
    }
    v14 := GenS1{F0: 0, F1: "hello", F2: 127}
    if true {
        v2 += (v2 - v2)
        var v15 uint16 = 961
        _ = v15
    } else {
        v1 += 0
        v16 := GenS1{F0: 42, F1: "hello", F2: 0}
        _ = v16
    }
    if (false || ("hello" >= ("a" + (((("\n\t" + (("\n\t" + "foo bar") + "a")) + "a") + ("hello" + "hello")) + (("\n\t" + "") + "foo bar"))))) {
        switch 42 {
        case -113:
            switch (("" + "") + ("\n\t" + "")) {
            case "\n\t":
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 uint8 = 42
                    _ = fi17
                    _ = v18
                }
            case "foo bar":
                fmt.Println(v1)
            default:
                v2 -= v2
            }
        case 42:
            v19 := GenS2{F0: 32768, F1: 360}
            _ = v19
        }
        var v20 int16 = v2
        v20 -= -591
        _ = v20
    } else {
        v1 -= (v1 / ((v1) | 1))
        if true {
            var v21 int16 = (-(v2))
            _ = v21
        } else {
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 bool = false
                if (fi22 != (fi22 + (fi22 << ((1) & 63)))) {
                    v2 = v2
                    if v23 {
                        v24 := GenS3{F0: 220}
                        var v25 int64 = 9
                        v25 += v25
                        _ = v24
                        _ = v25
                    } else {
                        var v26 float32 = (((3.14159 + (0.1 / float32((float64(1.5) + (1.5 / float64(1e10)))))) * -2.25) + 0.1)
                        _ = len(v3)
                        v26 += v26
                        _ = v26
                    }
                    v23 = v23
                }
                if ("\n\t" > "\n\t") {
                    v23 = false
                } else {
                    v1 = (^(v1))
                    if false {
                        switch fi22 {
                        case 127:
                            v23 = ((fi22 - (^(fi22))) <= 207)
                        case -460:
                            if ((v1 * v1) == 223) {
                                if (v23 && v23) {
                                    v27 := GenS4{F0: 374, F1: int16(fi22), F2: 288}
                                    _ = v27
                                }
                                if (277 < -921) {
                                    var v28 int32 = 858
                                    if false {
                                        v3 = append(v3, (fi22 % ((fi22) | 1)))
                                    }
                                    fmt.Println(v28)
                                    _ = v28
                                }
                            }
                        default:
                            if len(v3) > 0 {
                                v3[0] = genFn0(1, "foo bar")
                            }
                        }
                    } else {
                        if v23 {
                            var v29 int = (fi22 | (fi22 + 127))
                            var v30 int = (^(v29))
                            _ = v29
                            _ = v30
                        } else {
                            for fi31 := 0; fi31 < 14; fi31++ {
                                var v32 uint16 = 779
                                _ = fi31
                                _ = v32
                            }
                            var v33 uint8 = 0
                            _ = v33
                        }
                        var v34 float32 = float32((-2.25 / 0.1))
                        _ = v34
                    }
                }
                _ = fi22
                _ = v23
            }
            v1 -= 42
            v35 := GenS5{F0: -128, F1: 100, F2: 65535}
            _ = v35
        }
    }
    v3 = append(v3, 114)
    for fi36 := 0; fi36 < 11; fi36++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi36
    }
    v37 := GenS6{F0: 2}
    switch 32767 {
    case -1:
        if false {
            var v38 int = 0
            v2 += v2
            _ = v38
        }
    case 128:
        var v39 bool = (42 <= 1)
        _ = v39
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        v41 := GenS3{F0: 2}
        v2 -= 1000
        if ((v1 << ((65535) & 63)) >= (v1 ^ (v1 & (v1 - v1)))) {
            for fi42 := 0; fi42 < 7; fi42++ {
                v2 = (v2 & v2)
                var v43 int32 = 0
                _ = fi42
                _ = v43
            }
            fmt.Println(v1)
        } else {
            var v44 float64 = (-2.25 - (3.14159 + (0.1 - -2.25)))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v44
        }
        _ = fi40
        _ = v41
    }
    var v45 int8 = 100
    v2 -= (v2 - v2)
    var v46 int64 = 0
    switch "\n\t" {
    case "\n\t":
        for fi47 := 0; fi47 < 16; fi47++ {
            v48 := GenS6{F0: v1}
            var v49 int32 = 1
            _ = fi47
            _ = v48
            _ = v49
        }
    case "a":
        var v50 uint8 = 0
        _ = v50
    case "foo bar":
        v2 += (v2 | v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v45
    _ = v46
}
