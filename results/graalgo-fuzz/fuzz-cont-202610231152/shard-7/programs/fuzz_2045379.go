package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint32; F2 int }
type GenS4 struct { F0 uint; F1 int }
func genFn0(p0 int8, p1 int16, p2 bool) int16 {
    var v6 int16 = (p1 << ((4294967295) & 63))
    fmt.Println(v6)
    var v7 uint8 = 217
    var v8 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (p1 & (^(p1)))
}
func genFn1(p0 int16, p1 int64, p2 uint) bool {
    v9 := GenS1{F0: 61, F1: 42}
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int = (fi10 >> ((p2) & 63))
        var v12 uint64 = 18446744073709551615
        if (((p1 | 42) == p1) && true) {
            v13 := GenS1{F0: 60, F1: fi10}
            _ = v13
        } else {
            switch ("hello" + "hello") {
            case "":
                v14 := GenS2{F0: genFn0(91, (p0 % ((genFn0(int8(p0), (p0 >> ((p2) & 63)), false)) | 1)), true)}
                _ = v14
            }
            p0 -= p0
            if ((1.5 - (0.1 - (0.0 * 3.14159))) > 3.14159) {
                if false {
                    switch ("" + "\n\t") {
                    case "":
                        if (((false || true) && ((^(p2)) > 877)) && true) {
                            for fi15 := 0; fi15 < 7; fi15++ {
                                p0 += genFn0(-1, p0, (true && (65535 != 255)))
                                _ = fi15
                            }
                        }
                    case "a":
                        var v16 uint = (p2 + (p2 + 255))
                        _ = v16
                    case "\n\t":
                        v12 = 42
                    default:
                        for fi17 := 0; fi17 < 16; fi17++ {
                            switch v11 {
                            case 2147483647:
                                fmt.Println(v12, v11, v12)
                            case 4294967295:
                                switch ("hello" + "hello") {
                                case "":
                                    v18 := GenS2{F0: (p0 / (((p0 + -562)) | 1))}
                                    _ = v18
                                case "hello":
                                    if false {
                                        switch ("foo bar" + "a") {
                                        case "foo bar":
                                            var v19 uint16 = 65535
                                            _ = v19
                                        default:
                                            v12 = v12
                                        }
                                        v20 := []int{0}
                                        _ = v20
                                    }
                                default:
                                    v12 = (v12 * v12)
                                }
                            case -602:
                                var v21 uint32 = 2
                                _ = v21
                            }
                            _ = fi17
                        }
                    }
                    v22 := GenS1{F0: 0, F1: fi10}
                    switch -356 {
                    case 2147483647:
                        v11 = v11
                    case -415:
                        switch 127 {
                        case 32768:
                            fmt.Println(p0, p2)
                        case 9223372036854775807:
                            for fi23 := 0; fi23 < 19; fi23++ {
                                if (117 < 293) {
                                    v24 := []int{}
                                    v12 -= (v12 - (v12 ^ (v12 * 2)))
                                    _ = v24
                                }
                                _ = fi23
                            }
                        default:
                            if (false && (true && true)) {
                                p2 += p2
                                v12 += (v12 & (^(v12)))
                                v25 := GenS3{F0: "", F1: 42, F2: v11}
                                _ = v25
                            }
                        }
                    case 1:
                        v26 := GenS3{F0: "a", F1: 4294967295, F2: fi10}
                        _ = v26
                    }
                    _ = v22
                }
                v27 := []int{v11, fi10, v11}
                v11 -= (v11 << (((p2 % ((368) | 1))) & 63))
                _ = v27
            } else {
                v28 := []int{v11, (fi10 - (v11 % ((fi10) | 1)))}
                switch 798 {
                case 0:
                    v29 := GenS1{F0: 42, F1: fi10}
                    _ = v29
                case 544:
                    for fi30 := 0; fi30 < 12; fi30++ {
                        p0 -= int16(fi30)
                        _ = fi30
                    }
                case 324:
                    for fi31 := 0; fi31 < 12; fi31++ {
                        p2 -= (^(p2))
                        _ = fi31
                    }
                default:
                    if (0.0 == -2.25) {
                        var v32 int8 = 100
                        v12 += 1
                        v32 -= v32
                        _ = v32
                    } else {
                        p1 = (^(p1))
                    }
                }
                _ = v28
            }
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 string = ((((("\n\t" + "foo bar") + "") + "foo bar") + "hello") + "hello")
        fmt.Println(p1)
        _ = fi33
        _ = v34
    }
    p0 = (-(p0))
    v35 := GenS4{F0: 65535, F1: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    _ = v9
    return true
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]bool)
    var v3 int16 = -32768
    var v4 uint32 = 6
    var v5 uint = 42
    var v36 float64 = 0.0
    for fi37 := 0; fi37 < 12; fi37++ {
        v36 = (3.14159 + 1e10)
        var v38 int16 = -505
        if true {
            switch fi37 {
            case 1:
                var v39 bool = false
                _ = v39
            }
            v38 = v38
        }
        _ = fi37
        _ = v38
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        if false {
            var v41 uint32 = v4
            _ = v41
        } else {
            var v42 uint32 = (^(v4))
            v5 = v5
            if (42 >= 42) {
                var v43 int64 = 1
                var v44 int16 = v3
                _ = v43
                _ = v44
            } else {
                var v45 int32 = 0
                var v46 int16 = genFn0(-1, v3, false)
                if false {
                    for fi47 := 0; fi47 < 15; fi47++ {
                        v5 = 65535
                        var v48 int16 = (v3 << ((4294967295) & 63))
                        var v49 int32 = (v45 % ((v45) | 1))
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    v46 = 42
                } else {
                    var v50 bool = false
                    if false {
                        switch ("a" + "") {
                        case "foo bar":
                            v46 += v46
                        default:
                            v1 -= v1
                        }
                    }
                    _ = v50
                }
                _ = v45
                _ = v46
            }
            _ = v42
        }
        v3 = genFn0(v1, genFn0((v1 | (v1 ^ (v1 << (((v5 | v5)) & 63)))), (v3 | genFn0(int8(v5), -851, true)), (false && (true || true))), (false && false))
        v51 := v2[(^(fi40))]
        _ = fi40
        _ = v51
    }
    v4 -= v4
    v36 = (1.5 / (v36 * float64(-2.25)))
    var v52 int = -9223372036854775808
    if false {
        for fi53 := 0; fi53 < 5; fi53++ {
            var v54 uint64 = 984
            switch "hello" {
            case "a":
                for fi55 := 0; fi55 < 3; fi55++ {
                    v56 := v2[1]
                    v5 = (v5 - v5)
                    _ = fi55
                    _ = v56
                }
            default:
                for fi57 := 0; fi57 < 14; fi57++ {
                    v3 += v3
                    var v58 uint16 = 165
                    v58 = (v58 >> ((525) & 63))
                    _ = fi57
                    _ = v58
                }
            }
            _ = fi53
            _ = v54
        }
        v52 += 32768
    }
    if (!false) {
        var v59 uint64 = 2
        var v60 int32 = 2147483647
        fmt.Println(v3, v60, v4)
        _ = v59
        _ = v60
    } else {
        for fi61 := 0; fi61 < 11; fi61++ {
            var v62 int = 127
            _ = fi61
            _ = v62
        }
        _ = len(v2)
        delete(v2, v52)
    }
    fmt.Println(v3)
    for fi63 := 0; fi63 < 7; fi63++ {
        fmt.Println(v52, v5)
        switch fi63 {
        case 4294967295:
            v64 := GenS2{F0: -32768}
            _ = v64
        }
        _ = fi63
    }
    var v65 int16 = genFn0((v1 * 127), 877, true)
    v66 := GenS4{F0: v5, F1: v52}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v52
    _ = v65
    _ = v66
}
