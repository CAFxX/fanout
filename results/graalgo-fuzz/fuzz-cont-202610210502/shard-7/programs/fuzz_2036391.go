package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint16; F2 int8 }
type GenS5 struct { F0 int }
func genFn0(p0 uint32) uint16 {
    p0 += p0
    for fi4 := 0; fi4 < 10; fi4++ {
        switch fi4 {
        case 128:
            switch 32768 {
            case 32768:
                var v5 string = "foo bar"
                _ = v5
            case 252:
                var v6 int32 = 2147483647
                _ = v6
            case 473:
                fmt.Println(fi4, p0)
            }
        }
        _ = fi4
    }
    _ = p0
    return 0
}
func genFn1() uint {
    if (1 > 1) {
        v7 := GenS1{F0: false, F1: 32767, F2: (("a" + "\n\t") + "\n\t")}
        if true {
            v8 := []int{}
            var v9 uint8 = 179
            fmt.Println(v9)
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: (true && false), F1: 702, F2: "a"}
        _ = v10
        _ = v7
    }
    if false {
        if (0 <= -32768) {
            v11 := []int{32768, 128, -410}
            _ = v11
        }
    } else {
        v12 := GenS1{F0: (4294967295 == 698), F1: 255, F2: ""}
        var v13 uint16 = 255
        _ = v12
        _ = v13
    }
    return 918
}

func main() {
    var v1 uint16 = 112
    v2 := make(map[uint32]bool)
    var v3 uint64 = 4294967295
    switch 2147483647 {
    case 538:
        var v14 uint64 = (v3 - v3)
        _ = v14
    case 402:
        v15 := GenS1{F0: (2 != 2), F1: v1, F2: ("" + "hello")}
        _ = v15
    default:
        switch 32767 {
        case -546:
            switch 9223372036854775807 {
            case -986:
                fmt.Println(v1, v1)
            default:
                var v16 int16 = -743
                _ = v16
            }
        case -2147483648:
            var v17 uint64 = (^(v3))
            _ = v17
        case 32767:
            v18 := GenS2{F0: 1, F1: 1, F2: (v1 + genFn0(55))}
            _ = v18
        }
    }
    var v19 int = 32767
    v20 := GenS2{F0: 42, F1: 1, F2: 65535}
    fmt.Println(v3)
    var v21 uint = genFn1()
    v21 = (v21 / ((v21) | 1))
    v3 = uint64(v21)
    _ = len(v2)
    var v22 uint = (v21 / ((uint(v3)) | 1))
    v23 := GenS1{F0: ((genFn0(2147483648) > (v1 >> ((v21) & 63))) || true), F1: v1, F2: "\n\t"}
    v1 = 1
    var v24 uint32 = 2147483648
    for fi25 := 0; fi25 < 5; fi25++ {
        fmt.Println(v22, v1, v22)
        var v26 uint = v22
        if false {
            v3 = v3
            v27 := GenS3{F0: (("foo bar" + "hello") + ("foo bar" + ((("a" + "") + "foo bar") + ("hello" + "a"))))}
            _ = v27
        } else {
            v21 = v22
            var v28 int16 = int16(v19)
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    if (!true) {
        if (!true) {
            switch (v19 | v19) {
            case 0:
                for fi29 := 0; fi29 < 18; fi29++ {
                    if ((fi29 / (((v19 * (fi29 + fi29))) | 1)) < 387) {
                        switch "foo bar" {
                        case "\n\t":
                            var v30 uint8 = 255
                            _ = v30
                        case "":
                            v31 := []int{(fi29 + 2147483647), int(v24), (v19 >> ((v22) & 63))}
                            _ = v31
                        case "foo bar":
                            if ((v1 % ((v1) | 1)) <= v1) {
                                if true {
                                    v22 = v22
                                    var v32 uint = v21
                                    if (-1 == 1) {
                                        switch "" {
                                        case "hello":
                                            v3 += (^(v3))
                                        default:
                                            var v33 float32 = 0.0
                                            _ = v33
                                        }
                                    } else {
                                        fmt.Println(v21)
                                    }
                                    _ = v32
                                }
                            }
                        }
                        for fi34 := 0; fi34 < 13; fi34++ {
                            v35 := []int{v19}
                            var v36 bool = (!(true && (!false)))
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    } else {
                        var v37 uint8 = 157
                        var v38 float64 = 3.14159
                        var v39 int32 = 42
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 4; fi40++ {
                        var v41 float32 = 3.14159
                        v1 = (^(v1))
                        var v42 float64 = 3.14159
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    _ = fi29
                }
            case -453:
                for fi43 := 0; fi43 < 10; fi43++ {
                    for fi44 := 0; fi44 < 12; fi44++ {
                        var v45 uint = v22
                        v46 := []int{fi44, (fi43 | -9223372036854775808), fi44, int(v1)}
                        fmt.Println(v24)
                        if false {
                            fmt.Println(v19)
                            v3 = v3
                            fmt.Println(v3, v21)
                            v47 := v2[(v24 * v24)]
                            _ = v47
                        }
                        fmt.Println(v21, fi44)
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                    switch v19 {
                    case 4294967295:
                        v3 += 104
                    default:
                        if false {
                            var v48 int = fi43
                            _ = v48
                        } else {
                            if false {
                                var v49 string = "foo bar"
                                v50 := GenS4{F0: 1, F1: (v1 & 32767), F2: -31}
                                _ = v49
                                _ = v50
                            } else {
                                for fi51 := 0; fi51 < 4; fi51++ {
                                    var v52 int64 = 0
                                    _ = fi51
                                    _ = v52
                                }
                                var v53 int = v19
                                _ = v53
                            }
                        }
                    }
                    var v54 int8 = 41
                    _ = fi43
                    _ = v54
                }
            default:
                delete(v2, v24)
            }
            if true {
                v22 -= (v21 - (^(v22)))
            } else {
                var v55 uint32 = v24
                _ = v55
            }
            for fi56 := 0; fi56 < 15; fi56++ {
                if true {
                    switch (v19 * v19) {
                    case -557:
                        fmt.Println(v24)
                    }
                    var v57 float32 = -2.25
                    v19 = v19
                    _ = v57
                } else {
                    var v58 int8 = 127
                    _ = v58
                }
                if false {
                    switch (-(fi56)) {
                    case 942:
                        fmt.Println(v3, v22, v22)
                    case 9223372036854775807:
                        v2[(^(v24))] = (!(true || false))
                    case -9223372036854775808:
                        v3 = 4294967295
                    default:
                        var v59 float32 = float32(1e10)
                        _ = v59
                    }
                    v2[42] = (!true)
                    var v60 int64 = 0
                    _ = v60
                } else {
                    v61 := []int{32767, v19}
                    _ = v61
                }
                _ = fi56
            }
        }
        if (false && (803 != -32768)) {
            if (false && ((!true) || false)) {
                if false {
                    if false {
                        v1 = 255
                        for fi62 := 0; fi62 < 2; fi62++ {
                            v1 -= 32767
                            v63 := GenS5{F0: v19}
                            v1 += (^(v1))
                            _ = fi62
                            _ = v63
                        }
                        v2[v24] = (!(false && (!false)))
                    } else {
                        var v64 int16 = -506
                        v3 -= v3
                        _ = v64
                    }
                } else {
                    for fi65 := 0; fi65 < 6; fi65++ {
                        fmt.Println(v22, v22, v19)
                        _ = fi65
                    }
                }
            } else {
                var v66 uint = v22
                for fi67 := 0; fi67 < 13; fi67++ {
                    v68 := []int{fi67, v19, 4294967295, int(v1)}
                    v69 := GenS5{F0: v19}
                    _ = fi67
                    _ = v68
                    _ = v69
                }
                _ = v66
            }
        }
        v1 += (^(v1))
    } else {
        fmt.Println(v1, v24, v21)
        var v70 int16 = -32768
        switch ("a" + "hello") {
        case "hello":
            switch "\n\t" {
            case "hello":
                v71 := []int{v19, v19, (v19 - (v19 % ((v19) | 1))), v19}
                _ = v71
            case "a":
                v72 := GenS1{F0: true, F1: (^(v1)), F2: "hello"}
                _ = v72
            default:
                var v73 uint16 = v1
                _ = v73
            }
        default:
            var v74 int8 = -61
            _ = v74
        }
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
