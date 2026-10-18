package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint, p1 int8, p2 int16) int {
    p0 = p0
    p2 += p2
    var v7 int = 2147483647
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return v7
}

func main() {
    v1 := []uint32{65535}
    var v2 float32 = float32((0.0 - ((0.0 + 3.14159) - -2.25)))
    v3 := make(map[string]string)
    var v4 uint32 = 2147483648
    v5 := make(map[string]bool)
    var v6 int8 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 uint64 = 229
    if ((v2 - ((v2 + v2) + v2)) != 1.5) {
        var v9 string = "hello"
        if true {
            _ = len(v3)
        }
        switch 272 {
        case 32768:
            v10 := GenS1{F0: 2147483647}
            _ = v10
        case 32767:
            switch "\n\t" {
            case "a":
                for fi11 := 0; fi11 < 17; fi11++ {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        var v13 int64 = 0
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
            default:
                var v14 int16 = 1000
                _ = v14
            }
        case 128:
            v9 = (v9 + ((v9 + v9) + v9))
        default:
            if false {
                var v15 int = -63
                v16 := GenS2{F0: false}
                _ = v15
                _ = v16
            }
        }
        _ = v9
    } else {
        v8 += v8
        if (((v2 - (-2.25 / v2)) != ((1.5 - (3.14159 - (float32(float64(0.0)) / (0.1 + float32((-2.25 + 1.5)))))) - ((0.0 / v2) + (v2 / v2)))) || false) {
            switch genFn0(220, (v6 & 127), -15) {
            case 42:
                for fi17 := 0; fi17 < 20; fi17++ {
                    if (!(!((0 > -1) && (-1 > 898)))) {
                        v4 = (v4 / ((40) | 1))
                        v2 = float32(3.14159)
                    } else {
                        switch fi17 {
                        case 0:
                            if ((false && (v8 == v8)) || false) {
                                for fi18 := 0; fi18 < 16; fi18++ {
                                    v2 += v2
                                    v19 := GenS3{F0: -1}
                                    _ = fi18
                                    _ = v19
                                }
                                if (!true) {
                                    v20 := v5[(("a" + (("a" + "a") + "")) + "\n\t")]
                                    for fi21 := 0; fi21 < 5; fi21++ {
                                        fmt.Println(v2, v20)
                                        _ = fi21
                                    }
                                    _ = v20
                                }
                                v22 := GenS4{F0: 32768}
                                _ = v22
                            } else {
                                var v23 int16 = -32768
                                _ = v23
                            }
                        default:
                            _ = len(v3)
                        }
                        var v24 int16 = 747
                        if false {
                            var v25 int32 = 772
                            _ = v25
                        } else {
                            v24 = (v24 << ((4294967295) & 63))
                            v8 = v8
                            v26 := GenS5{F0: -313, F1: 65535}
                            _ = v26
                        }
                        _ = v24
                    }
                    var v27 uint64 = v8
                    _ = fi17
                    _ = v27
                }
            case -9223372036854775808:
                v28 := GenS6{F0: 2}
                _ = v28
            case 812:
                v29 := GenS2{F0: false}
                _ = v29
            default:
                fmt.Println(v8, v8, v2)
            }
        } else {
            v4 += 65535
            if (4294967295 == 255) {
                fmt.Println(v8, v2)
                v2 -= v2
                v30 := GenS7{F0: 0}
                _ = v30
            }
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 uint64 = v8
            for fi33 := 0; fi33 < 10; fi33++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (42 > 42) {
                    v6 = (v6 % (((^(v6))) | 1))
                    v34 := GenS5{F0: fi31, F1: 18446744073709551615}
                    _ = v34
                } else {
                    var v35 uint32 = v4
                    if false {
                        var v36 int32 = -220
                        fmt.Println(v4, fi33, fi33)
                        var v37 int32 = 1
                        var v38 string = "foo bar"
                        var v39 bool = (255 <= 0)
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                    } else {
                        var v40 uint16 = 605
                        var v41 uint8 = 0
                        _ = v40
                        _ = v41
                    }
                    var v42 int = 9223372036854775807
                    delete(v5, "\n\t")
                    _ = v35
                    _ = v42
                }
                _ = fi33
            }
            _ = fi31
            _ = v32
        }
    }
    var v43 uint32 = v4
    if false {
        var v44 int64 = 749
        _ = v44
    }
    for fi45 := 0; fi45 < 6; fi45++ {
        v46 := GenS1{F0: (v43 / ((v43) | 1))}
        var v47 uint8 = 2
        _ = fi45
        _ = v46
        _ = v47
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        var v49 uint = 252
        _ = fi48
        _ = v49
    }
    switch ("" + "a") {
    case "a":
        v50 := v5["a"]
        _ = v50
    }
    for fi51 := 0; fi51 < 20; fi51++ {
        if (!(193 <= 136)) {
            if true {
                var v52 int32 = -333
                fmt.Println(v43, v2, v6)
                v3[""] = ("a" + (("" + "foo bar") + ("" + "a")))
                _ = v52
            }
            var v53 uint8 = uint8(v43)
            _ = v53
        }
        _ = fi51
    }
    var v54 uint = 301
    if ((v6 ^ v6) <= v6) {
        for fi55 := 0; fi55 < 19; fi55++ {
            v56 := GenS2{F0: (true || true)}
            for fi57 := 0; fi57 < 3; fi57++ {
                _ = len(v1)
                v2 = float32(float64(v2))
                v43 = v4
                v2 = v2
                _ = fi57
            }
            _ = fi55
            _ = v56
        }
    }
    v54 -= (v54 % (((v54 % (((^(v54))) | 1))) | 1))
    for fi58 := 0; fi58 < 20; fi58++ {
        v54 = uint(v8)
        if false {
            v6 += v6
            v59 := GenS6{F0: 0}
            _ = v59
        }
        var v60 int8 = (v6 & v6)
        _ = fi58
        _ = v60
    }
    v54 = (v54 * v54)
    v6 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v54
    _ = v6
    _ = v8
}
