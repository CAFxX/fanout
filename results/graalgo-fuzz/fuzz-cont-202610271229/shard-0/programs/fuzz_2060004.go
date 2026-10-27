package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 int; F2 bool }
type GenS6 struct { F0 int64; F1 uint64; F2 int8 }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]bool)
    v3 := []uint32{0, 2147483648, 17}
    var v4 bool = true
    var v5 int8 = 42
    var v6 uint8 = 58
    var v7 int8 = -74
    v3 = append(v3, 65535)
    v6 = (v6 | 128)
    switch 4294967295 {
    case 212:
        v8 := GenS1{F0: "hello", F1: "\n\t", F2: 2}
        _ = v8
    case -35:
        switch -561 {
        case 32767:
            if (((0 < v6) && ((!(v4 && v4)) && false)) || v4) {
                var v9 int8 = 127
                switch "\n\t" {
                case "foo bar":
                    for fi10 := 0; fi10 < 16; fi10++ {
                        for fi11 := 0; fi11 < 5; fi11++ {
                            fmt.Println(v5)
                            v3 = append(v3, 2147483648)
                            _ = fi11
                        }
                        v12 := GenS1{F0: "\n\t", F1: ("hello" + "hello"), F2: uint(v6)}
                        var v13 int8 = (v7 / (((v9 - (v5 << ((255) & 63)))) | 1))
                        _ = fi10
                        _ = v12
                        _ = v13
                    }
                }
                if (1 != 1000) {
                    var v14 uint8 = (v6 | (^(v6)))
                    _ = v14
                }
                _ = v9
            } else {
                if false {
                    fmt.Println(v6)
                }
                v15 := GenS2{F0: 811}
                var v16 int8 = v5
                _ = v15
                _ = v16
            }
        default:
            fmt.Println(v7, v5)
        }
    case -902:
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 float32 = 1e10
            v18 += (v18 - (v18 - float32(v1)))
            _ = fi17
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 int16 = 32767
                var v22 uint64 = 4294967295
                v21 = 1
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
    }
    delete(v2, 1)
    switch ("" + "") {
    case "":
        switch 1 {
        case 128:
            if true {
                switch 32768 {
                case 1:
                    var v23 int64 = 534
                    _ = v23
                case -2147483648:
                    var v24 int8 = (v5 & 100)
                    _ = v24
                case 128:
                    for fi25 := 0; fi25 < 10; fi25++ {
                        for fi26 := 0; fi26 < 17; fi26++ {
                            var v27 float32 = (1.5 + 0.1)
                            v28 := GenS1{F0: ((("" + "\n\t") + ("a" + "\n\t")) + ""), F1: "", F2: 4294967295}
                            var v29 int64 = 994
                            fmt.Println(v27, v4, v4)
                            _ = fi26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        var v30 string = ("\n\t" + "foo bar")
                        _ = fi25
                        _ = v30
                    }
                default:
                    v7 = v5
                }
            }
        }
    case "a":
        for fi31 := 0; fi31 < 12; fi31++ {
            if v4 {
                var v32 uint8 = (^(v6))
                if v4 {
                    var v33 int32 = -2147483648
                    _ = len(v2)
                    _ = v33
                } else {
                    for fi34 := 0; fi34 < 11; fi34++ {
                        fmt.Println(v32)
                        v3 = append(v3, 1)
                        _ = fi34
                    }
                }
                var v35 int8 = v7
                _ = v32
                _ = v35
            }
            v6 = v6
            _ = fi31
        }
    }
    if (0 >= 255) {
        fmt.Println(v1, v5)
        for fi36 := 0; fi36 < 4; fi36++ {
            if (true || (!false)) {
                v1 = (v1 + 0.0)
                v7 += v7
            } else {
                v37 := GenS3{F0: 377}
                var v38 float64 = 1e10
                _ = v37
                _ = v38
            }
            for fi39 := 0; fi39 < 13; fi39++ {
                switch ("" + ("hello" + "hello")) {
                case "\n\t":
                    v4 = false
                case "foo bar":
                    var v40 int32 = 461
                    _ = v40
                }
                v41 := GenS1{F0: "hello", F1: ((("\n\t" + "\n\t") + "a") + ("\n\t" + "a")), F2: 255}
                _ = fi39
                _ = v41
            }
            _ = fi36
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v42 := GenS4{F0: (1 < 18446744073709551615)}
    _ = len(v3)
    var v43 float32 = 0.1
    v4 = false
    v44 := GenS5{F0: 127, F1: 32768, F2: true}
    v3 = append(v3, 2147483647)
    for fi45 := 0; fi45 < 14; fi45++ {
        var v46 float32 = 1e10
        _ = fi45
        _ = v46
    }
    if true {
        switch 2147483647 {
        case 4294967295:
            for fi47 := 0; fi47 < 4; fi47++ {
                if (!false) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    if v4 {
                        for fi48 := 0; fi48 < 18; fi48++ {
                            var v49 uint16 = 1
                            if true {
                                v5 = v7
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v50 float32 = ((((3.14159 - float32(v1)) - 1e10) / v43) - v43)
                                _ = v50
                            } else {
                                if (!(2 < 42)) {
                                    fmt.Println(fi47, v43, v7)
                                    switch ("" + "") {
                                    case "a":
                                        v51 := GenS5{F0: (fi48 / ((fi47) | 1)), F1: (fi48 >> ((255) & 63)), F2: v4}
                                        _ = v51
                                    case "":
                                        v49 += (v49 >> ((42) & 63))
                                    case "hello":
                                        v52 := v2[286]
                                        _ = v52
                                    default:
                                        v5 = 127
                                    }
                                    v7 += 127
                                }
                                var v53 int16 = 893
                                var v54 uint64 = 88
                                _ = v53
                                _ = v54
                            }
                            _ = fi48
                            _ = v49
                        }
                        for fi55 := 0; fi55 < 4; fi55++ {
                            v5 += v7
                            _ = fi55
                        }
                    }
                }
                _ = fi47
            }
        case 688:
            if false {
                for fi56 := 0; fi56 < 17; fi56++ {
                    for fi57 := 0; fi57 < 16; fi57++ {
                        fmt.Println(v43, v4, fi56)
                        _ = fi57
                    }
                    _ = fi56
                }
                var v58 uint64 = 18446744073709551615
                var v59 int64 = 42
                _ = v58
                _ = v59
            } else {
                v6 += 0
                v60 := GenS6{F0: 42, F1: 785, F2: -128}
                _ = v60
            }
        }
        _ = len(v2)
    } else {
        if v4 {
            v61 := GenS3{F0: 1}
            for fi62 := 0; fi62 < 8; fi62++ {
                var v63 int = (fi62 / ((fi62) | 1))
                _ = fi62
                _ = v63
            }
            fmt.Println(v43)
            _ = v61
        }
        var v64 uint32 = 943
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
