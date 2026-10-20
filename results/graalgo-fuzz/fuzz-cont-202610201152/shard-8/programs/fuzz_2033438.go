package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 int32; F1 int64; F2 int64 }
func genFn0(p0 int) int64 {
    var v6 uint64 = 42
    var v7 int16 = -155
    v7 -= (v7 | (v7 & 1000))
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int8 = -128
        _ = fi8
        _ = v9
    }
    switch "a" {
    case "foo bar":
        v10 := GenS1{F0: 465}
        _ = v10
    case "hello":
        if (42 <= -1) {
            v11 := GenS2{F0: 1}
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int64 = 9223372036854775807
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    case "":
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := []int{(fi14 | (p0 << ((0) & 63))), fi14}
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                v6 = 18446744073709551615
                if true {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                }
                if ((^(fi16)) <= -2147483648) {
                    var v17 uint = 472
                    _ = v17
                }
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
    default:
        if (!true) {
            switch 32767 {
            case 9223372036854775807:
                for fi18 := 0; fi18 < 2; fi18++ {
                    if (false || (42 >= 56)) {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            v20 := GenS3{F0: 1}
                            if true {
                                p0 = (fi18 + p0)
                                if false {
                                    v21 := []int{}
                                    v22 := GenS4{F0: 0}
                                    _ = v21
                                    _ = v22
                                } else {
                                    v6 = v6
                                }
                            } else {
                                var v23 int32 = 0
                                fmt.Println(v23, v23)
                                _ = v23
                            }
                            _ = fi19
                            _ = v20
                        }
                    }
                    _ = fi18
                }
            case 2147483647:
                var v24 int64 = 51
                _ = v24
            default:
                var v25 float32 = float32(float64(((float32(1.5) - 3.14159) - -2.25)))
                _ = v25
            }
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := []int{(-(fi26)), (p0 | -171), (p0 & (^(p0))), p0}
                _ = fi26
                _ = v27
            }
            v28 := []int{}
            _ = v28
        } else {
            fmt.Println(p0, v6)
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 42
    var v2 float64 = 0.1
    var v3 uint32 = 4294967295
    var v4 int = -192
    v5 := make(map[string]bool)
    switch (v4 | int(v3)) {
    case 4294967295:
        var v29 int = 4294967295
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 int = fi30
            v3 += (v3 | (v3 % (((v3 % ((59) | 1))) | 1)))
            v5[("hello" + "a")] = (-32768 != -1)
            _ = fi31
            _ = v32
        }
        v4 = (v4 >> ((30) & 63))
        _ = fi30
    }
    v3 -= (v3 % (((^(v3))) | 1))
    v1 -= (v1 * 42)
    v33 := []int{v4}
    var v34 string = ""
    if true {
        delete(v5, (v34 + "a"))
        for fi35 := 0; fi35 < 15; fi35++ {
            if (0 < -2147483648) {
                for fi36 := 0; fi36 < 8; fi36++ {
                    delete(v5, v34)
                    v3 = (v3 | v3)
                    switch (fi36 * fi35) {
                    case 41:
                        var v37 uint64 = 42
                        _ = v37
                    case -939:
                        v3 = v3
                    case -9223372036854775808:
                        if (!true) {
                            if true {
                                switch v34 {
                                case "a":
                                    var v38 uint8 = 255
                                    _ = v38
                                case "\n\t":
                                    v2 = ((1.5 * (v2 - float64(1.5))) * v2)
                                    var v39 uint16 = 276
                                    _ = v39
                                case "":
                                    v4 = -9223372036854775808
                                }
                                v40 := v5[v34]
                                v41 := GenS5{F0: 65535, F1: 42}
                                _ = v40
                                _ = v41
                            }
                        } else {
                            if true {
                                var v42 uint8 = 1
                                v2 = v2
                                var v43 float32 = 1.5
                                _ = v42
                                _ = v43
                            }
                            var v44 float32 = (-2.25 + float32((v2 / v2)))
                            var v45 uint8 = 1
                            _ = v44
                            _ = v45
                        }
                    }
                    if true {
                        var v46 float32 = float32(float64(0.0))
                        fmt.Println(v1, v2)
                        var v47 uint16 = 1
                        if false {
                            var v48 int64 = 1
                            var v49 string = (("foo bar" + (v34 + v34)) + v34)
                            switch fi36 {
                            case 9223372036854775807:
                                _ = len(v33)
                            case 42:
                                v2 -= (1e10 + 0.1)
                            default:
                                v50 := GenS4{F0: 16}
                                _ = v50
                            }
                            _ = v48
                            _ = v49
                        } else {
                            if ((0 > 42) || (!true)) {
                                if true {
                                    var v51 int32 = 1
                                    _ = len(v5)
                                    _ = v51
                                }
                            } else {
                                if (false || (42 <= 2)) {
                                    v52 := GenS1{F0: v47}
                                    var v53 int = fi35
                                    var v54 bool = true
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                                v55 := GenS1{F0: 1}
                                _ = v55
                            }
                            if (0 <= uint8(v47)) {
                                switch (v34 + "") {
                                case "hello":
                                    fmt.Println(fi35, fi36, v3)
                                case "a":
                                    var v56 int8 = 127
                                    _ = v56
                                }
                                _ = len(v5)
                                v4 -= 32767
                                v57 := GenS1{F0: v47}
                                var v58 uint32 = (v3 << ((2) & 63))
                                if true {
                                    var v59 float32 = v46
                                    v60 := GenS6{F0: 2147483647, F1: -1, F2: 9223372036854775807}
                                    _ = v59
                                    _ = v60
                                }
                                var v61 int = fi35
                                fmt.Println(v58, v2)
                                _ = v57
                                _ = v58
                                _ = v61
                            }
                        }
                        _ = v46
                        _ = v47
                    }
                    v34 = ("hello" + v34)
                    _ = fi36
                }
                fmt.Println(v2)
                v5[""] = true
            } else {
                v1 += 2
            }
            v2 = float64(3.14159)
            for fi62 := 0; fi62 < 10; fi62++ {
                if true {
                    var v63 bool = (v34 <= ((v34 + ("\n\t" + ("hello" + v34))) + "\n\t"))
                    v34 = (v34 + (v34 + (("hello" + v34) + "foo bar")))
                    _ = v63
                }
                v3 = v3
                v64 := v5[v34]
                _ = fi62
                _ = v64
            }
            _ = fi35
        }
    }
    var v65 float64 = v2
    if true {
        v34 = (("" + ("a" + "\n\t")) + v34)
        v5[(v34 + v34)] = (-109 > 100)
    } else {
        for fi66 := 0; fi66 < 10; fi66++ {
            var v67 int64 = 1
            _ = fi66
            _ = v67
        }
        delete(v5, "a")
    }
    for fi68 := 0; fi68 < 17; fi68++ {
        v2 = v65
        var v69 uint64 = v1
        var v70 float32 = 0.0
        _ = fi68
        _ = v69
        _ = v70
    }
    fmt.Println(v4, v65, v65)
    for fi71 := 0; fi71 < 12; fi71++ {
        v4 -= (v4 & fi71)
        for fi72 := 0; fi72 < 10; fi72++ {
            if len(v33) > 0 {
                v33[0] = (v4 % (((fi72 << ((4294967295) & 63))) | 1))
            }
            v73 := v5["\n\t"]
            var v74 uint8 = 0
            _ = fi72
            _ = v73
            _ = v74
        }
        _ = fi71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v65
}
