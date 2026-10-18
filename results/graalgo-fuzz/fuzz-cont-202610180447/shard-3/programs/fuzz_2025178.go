package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int; F1 uint; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 uint; F1 int64 }
func genFn0(p0 int8, p1 uint32, p2 uint64) int64 {
    fmt.Println(p1)
    var v4 uint64 = p2
    if false {
        v5 := GenS1{F0: (p2 >> ((4294967295) & 63)), F1: (("\n\t" + "") + ("" + (("a" + "foo bar") + "foo bar")))}
        p2 = 4294967295
        _ = v5
    } else {
        v6 := GenS2{F0: 1, F1: (p0 << ((1) & 63))}
        if false {
            var v7 uint32 = (^(p1))
            v8 := []int{-721, 2147483647, 127, 127}
            for fi9 := 0; fi9 < 15; fi9++ {
                p1 = (v7 - 393)
                _ = fi9
            }
            _ = v7
            _ = v8
        } else {
            var v10 uint = 1
            _ = v10
        }
        if true {
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := []int{fi11, 0, 0, (fi11 - fi11)}
                var v13 int8 = -72
                var v14 int = fi11
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
        _ = v6
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        fmt.Println(p0)
        switch "foo bar" {
        case "":
            fmt.Println(v4, p2, p1)
        }
        v16 := GenS3{F0: fi15, F1: 4294967295, F2: 2}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -358
}
func genFn1(p0 int16, p1 int64, p2 int64) int16 {
    v17 := []int{}
    if (!(p2 < (p1 ^ 1))) {
        _ = len(v17)
    } else {
        v18 := GenS4{F0: (p2 & p1), F1: 1}
        var v19 uint8 = 255
        var v20 bool = (false && ((!true) || false))
        _ = v18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return (p0 ^ (^(p0)))
}

func main() {
    v1 := []int{-989, 1, 0}
    var v2 int64 = -1
    var v3 uint = 1
    switch 128 {
    case 9223372036854775807:
        if (!false) {
            switch 42 {
            case 2147483647:
                if ((v2 | v2) > genFn0(100, 1, uint64(v3))) {
                    if false {
                        var v21 float32 = 0.1
                        _ = v21
                    } else {
                        fmt.Println(v3, v2)
                        switch "hello" {
                        case "":
                            var v22 uint32 = 65535
                            _ = v22
                        }
                        if false {
                            v23 := GenS3{F0: 32768, F1: v3, F2: 0}
                            var v24 int64 = genFn0(105, 989, 739)
                            _ = v23
                            _ = v24
                        } else {
                            var v25 float64 = 0.0
                            for fi26 := 0; fi26 < 1; fi26++ {
                                for fi27 := 0; fi27 < 17; fi27++ {
                                    v3 += (^(v3))
                                    if (127 == 42) {
                                        fmt.Println(v2, v3)
                                        fmt.Println(v25, fi27)
                                        var v28 int16 = 32767
                                        var v29 int8 = 100
                                        v30 := GenS5{F0: v3}
                                        v2 = (^(v2))
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 uint32 = 42
                                    _ = fi27
                                    _ = v31
                                }
                                v2 = -1
                                switch "\n\t" {
                                case "foo bar":
                                    if (!true) {
                                        v3 -= v3
                                        switch fi26 {
                                        case 128:
                                            v1 = append(v1, fi26)
                                        case 285:
                                            var v32 int8 = 0
                                            _ = v32
                                        }
                                        fmt.Println(v3)
                                    } else {
                                        if (((v3 / ((v3) | 1)) == (^(v3))) || true) {
                                            switch (fi26 & (^(fi26))) {
                                            case 128:
                                                var v33 float64 = (v25 + ((3.14159 + v25) + (3.14159 - v25)))
                                                _ = v33
                                            case 9223372036854775807:
                                                var v34 int = (fi26 - 729)
                                                _ = v34
                                            case -986:
                                                var v35 int16 = 1
                                                _ = v35
                                            }
                                        }
                                        var v36 int64 = v2
                                        for fi37 := 0; fi37 < 18; fi37++ {
                                            v36 = genFn0(int8(fi37), 289, 1)
                                            if true {
                                                switch fi37 {
                                                case 4294967295:
                                                    var v38 int64 = v36
                                                    _ = v38
                                                default:
                                                    fmt.Println(fi26, v2, v25)
                                                }
                                            }
                                            if (((1.5 - v25) - v25) != 1e10) {
                                                v3 = (v3 % ((712) | 1))
                                                v39 := GenS5{F0: v3}
                                                if (!false) {
                                                    v40 := GenS2{F0: fi26, F1: -29}
                                                    if len(v1) > 0 {
                                                        v1[0] = -185
                                                    }
                                                    v36 = 9223372036854775807
                                                    _ = v40
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v41 := GenS4{F0: v36, F1: 42}
                                                    _ = v41
                                                }
                                                v3 = 255
                                                v2 = (v36 & v36)
                                                fmt.Println(v3, fi26, fi26)
                                                _ = v39
                                            } else {
                                                var v42 string = "foo bar"
                                                _ = v42
                                            }
                                            fmt.Println(fi26, v3)
                                            _ = fi37
                                        }
                                        _ = v36
                                    }
                                case "hello":
                                    v43 := GenS6{F0: int64(v3), F1: true}
                                    _ = v43
                                case "a":
                                    if (v25 <= ((v25 - -2.25) + (3.14159 * v25))) {
                                        if false {
                                            v25 += 0.1
                                            v25 = ((v25 * -2.25) / v25)
                                        }
                                    }
                                }
                                _ = fi26
                            }
                            _ = v25
                        }
                    }
                    fmt.Println(v2)
                    v44 := GenS2{F0: 4294967295, F1: 1}
                    _ = v44
                } else {
                    var v45 int = 9223372036854775807
                    v46 := GenS7{F0: true}
                    _ = v45
                    _ = v46
                }
            case 42:
                fmt.Println(v3)
            default:
                var v47 int8 = 127
                _ = v47
            }
            for fi48 := 0; fi48 < 1; fi48++ {
                if (-1 > 2147483647) {
                    v3 = v3
                } else {
                    v2 -= (v2 >> ((v3) & 63))
                    switch (fi48 & fi48) {
                    case -489:
                        if (!false) {
                            for fi49 := 0; fi49 < 8; fi49++ {
                                if ((-128 != -1) || true) {
                                    v2 = (v2 - v2)
                                    var v50 uint16 = 255
                                    _ = v50
                                } else {
                                    fmt.Println(v2, fi48)
                                }
                                v3 -= (v3 * v3)
                                _ = fi49
                            }
                            var v51 uint32 = 2147483647
                            switch fi48 {
                            case 0:
                                fmt.Println(v3, v2, fi48)
                            case 32768:
                                fmt.Println(v3)
                            case 32767:
                                fmt.Println(v51)
                            default:
                                v52 := GenS4{F0: int64(fi48), F1: v51}
                                _ = v52
                            }
                            _ = v51
                        }
                    case -2147483648:
                        for fi53 := 0; fi53 < 6; fi53++ {
                            v2 = (v2 ^ (-(v2)))
                            v3 = (v3 / ((v3) | 1))
                            v54 := GenS8{F0: 169, F1: 65535}
                            _ = fi53
                            _ = v54
                        }
                    case 1:
                        for fi55 := 0; fi55 < 11; fi55++ {
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = (fi55 >> (((^(v3))) & 63))
                                }
                                var v56 int16 = -69
                                v2 -= (v2 / (((v2 + (v2 >> ((1) & 63)))) | 1))
                                _ = v56
                            }
                            v2 = (v2 - v2)
                            _ = fi55
                        }
                    default:
                        for fi57 := 0; fi57 < 10; fi57++ {
                            var v58 bool = (526 != fi48)
                            _ = len(v1)
                            _ = fi57
                            _ = v58
                        }
                    }
                }
                _ = len(v1)
                _ = fi48
            }
        } else {
            v59 := GenS9{F0: (v3 | v3), F1: v2}
            _ = v59
        }
    case 1:
        if (!false) {
            if (v3 > v3) {
                fmt.Println(v3, v3)
                for fi60 := 0; fi60 < 16; fi60++ {
                    v61 := GenS3{F0: (fi60 * 4294967295), F1: v3, F2: 2}
                    v2 += (v2 << ((v3) & 63))
                    _ = fi60
                    _ = v61
                }
            }
            var v62 string = "hello"
            _ = v62
        }
    }
    for fi63 := 0; fi63 < 14; fi63++ {
        v3 = (v3 & v3)
        var v64 uint32 = 65535
        _ = fi63
        _ = v64
    }
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    var v65 int = -9223372036854775808
    v1 = append(v1, v65)
    var v66 int16 = 710
    v2 = (v2 >> (((v3 & 18446744073709551615)) & 63))
    fmt.Println(v3, v3, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v66
}
