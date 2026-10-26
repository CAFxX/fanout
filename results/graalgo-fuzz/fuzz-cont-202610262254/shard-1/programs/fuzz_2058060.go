package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0() bool {
    if (false && (659 == 32767)) {
        var v6 float64 = 0.1
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                fmt.Println(fi7)
                _ = fi8
            }
            var v9 bool = (1 == 1)
            var v10 uint64 = 4294967295
            _ = fi7
            _ = v10
            _ = v9
        }
        _ = v6
    } else {
        v11 := []int{-907, 1, 4294967295, 2147483647}
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint = 255
            v14 := GenS1{F0: 32767}
            var v15 float64 = 0.1
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = v11
    }
    var v16 uint16 = 42
    var v17 int16 = 538
    v17 += v17
    fmt.Println(v16, v17)
    _ = v16
    _ = v17
    return (127 <= -124)
}

func main() {
    v1 := []uint32{968, 4294967295, 4294967295}
    var v2 uint = 352
    var v3 float64 = (3.14159 * (1e10 + 1.5))
    var v4 uint16 = 32767
    v5 := make(map[uint32]bool)
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := GenS1{F0: 1}
            if false {
                var v21 uint64 = 319
                v21 += (^(v21))
                _ = v21
            } else {
                v3 -= (v3 * float64(1e10))
                var v22 string = "a"
                v1 = append(v1, 2)
                _ = v22
            }
            v4 -= (v4 >> (((v2 ^ (v2 << ((v2) & 63)))) & 63))
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    if true {
        if (137 >= 220) {
            if (-2.25 > (((float64(1e10) / 3.14159) - float64((1e10 * (1e10 - float32(1e10))))) - ((1.5 * (3.14159 / 3.14159)) / -2.25))) {
                v3 -= v3
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 uint16 = (v4 + v4)
                    var v25 float32 = float32(-2.25)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                v26 := GenS1{F0: 1000}
                _ = v26
            }
        } else {
            v2 += v2
            if (1 <= 2147483647) {
                var v27 uint16 = (v4 << ((v2) & 63))
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                v28 := GenS1{F0: 42}
                _ = v27
                _ = v28
            }
            if (1 <= 42) {
                fmt.Println(v3, v2)
            }
        }
        if (!("a" >= ("foo bar" + "\n\t"))) {
            fmt.Println(v2, v3)
            _ = len(v5)
            switch ("" + "foo bar") {
            case "\n\t":
                if (true || false) {
                    v29 := GenS1{F0: 32767}
                    _ = v29
                } else {
                    if (!false) {
                        if false {
                            if (151 != 86) {
                                v4 = 612
                            }
                            switch 0 {
                            case 313:
                                var v30 uint16 = 2
                                _ = v30
                            case 127:
                                if ((!(42 <= -747)) && true) {
                                    for fi31 := 0; fi31 < 12; fi31++ {
                                        fmt.Println(fi31, v3)
                                        v32 := GenS2{F0: -834, F1: uint32(v2)}
                                        fmt.Println(fi31)
                                        _ = fi31
                                        _ = v32
                                    }
                                    fmt.Println(v2, v2)
                                    v33 := GenS2{F0: 1, F1: 65535}
                                    _ = v33
                                } else {
                                    for fi34 := 0; fi34 < 12; fi34++ {
                                        if (!(!true)) {
                                            for fi35 := 0; fi35 < 8; fi35++ {
                                                fmt.Println(fi35)
                                                _ = fi35
                                            }
                                            fmt.Println(v3, fi34)
                                            v36 := GenS2{F0: 616, F1: 1}
                                            _ = v36
                                        } else {
                                            var v37 uint32 = 156
                                            _ = v37
                                        }
                                        _ = fi34
                                    }
                                }
                            }
                            _ = len(v1)
                        }
                    }
                    _ = len(v1)
                    var v38 int = 32767
                    _ = v38
                }
            default:
                switch "foo bar" {
                case "hello":
                    for fi39 := 0; fi39 < 17; fi39++ {
                        var v40 bool = (0 >= -9223372036854775808)
                        switch "foo bar" {
                        case "\n\t":
                            var v41 uint16 = (v4 << ((v2) & 63))
                            _ = v41
                        case "a":
                            if (42 <= 636) {
                                fmt.Println(v2, v2)
                                var v42 string = (("" + "a") + (("\n\t" + "hello") + ("foo bar" + "foo bar")))
                                _ = v42
                            }
                        default:
                            v40 = v40
                        }
                        for fi43 := 0; fi43 < 18; fi43++ {
                            v40 = false
                            var v44 int = -526
                            fmt.Println(v3, v44, v2)
                            _ = fi43
                            _ = v44
                        }
                        _ = fi39
                        _ = v40
                    }
                default:
                    for fi45 := 0; fi45 < 15; fi45++ {
                        v46 := v5[0]
                        _ = fi45
                        _ = v46
                    }
                }
            }
        } else {
            v3 += ((1.5 / -2.25) + v3)
        }
    }
    v2 -= v2
    var v47 uint16 = 65535
    v2 -= (v2 + (v2 & (v2 - v2)))
    v47 += (v47 >> (((^(v2))) & 63))
    fmt.Println(v47)
    if false {
        fmt.Println(v3, v47, v4)
        v48 := GenS3{F0: false, F1: false}
        for fi49 := 0; fi49 < 6; fi49++ {
            switch (-(fi49)) {
            case 42:
                var v50 int8 = 127
                _ = v50
            case 4294967295:
                if (!false) {
                    if false {
                        for fi51 := 0; fi51 < 3; fi51++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v47)
                            v52 := GenS2{F0: 32767, F1: 42}
                            _ = fi51
                            _ = v52
                        }
                    } else {
                        v4 = (^(v4))
                        for fi53 := 0; fi53 < 1; fi53++ {
                            switch fi53 {
                            case 9223372036854775807:
                                if (-451 != -484) {
                                    var v54 float32 = 1e10
                                    v3 -= (v3 + float64(((v54 + v54) / 1.0)))
                                    delete(v5, 1)
                                    _ = v54
                                } else {
                                    v4 = (^(v4))
                                    if true {
                                        if true {
                                            var v55 int32 = 2147483647
                                            fmt.Println(v3)
                                            v2 += v2
                                            fmt.Println(v55)
                                            v55 = (^(v55))
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v55
                                        } else {
                                            switch (fi49 % ((42) | 1)) {
                                            case 32768:
                                                v56 := GenS4{F0: "foo bar", F1: 2147483647}
                                                _ = v56
                                            }
                                            if (!(2147483648 > 482)) {
                                                var v57 int32 = -2147483648
                                                if len(v1) > 0 {
                                                    v1[0] = 65535
                                                }
                                                _ = v57
                                            }
                                        }
                                    }
                                    var v58 bool = (4294967295 <= v2)
                                    _ = v58
                                }
                            }
                            var v59 uint32 = 2
                            _ = fi53
                            _ = v59
                        }
                        var v60 float32 = 1.5
                        _ = v60
                    }
                    v61 := GenS1{F0: 695}
                    _ = v61
                }
            default:
                var v62 bool = false
                _ = v62
            }
            _ = fi49
        }
        _ = v48
    }
    v4 -= 65535
    if false {
        fmt.Println(v2, v2)
        v2 += (^(v2))
        fmt.Println(v47)
    }
    if false {
        if len(v1) > 0 {
            v1[0] = 0
        }
        var v63 bool = false
        v47 = 0
        _ = v63
    } else {
        v4 -= v4
        var v64 int32 = 2147483647
        _ = v64
    }
    v3 = v3
    var v65 uint32 = uint32(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v65
}
