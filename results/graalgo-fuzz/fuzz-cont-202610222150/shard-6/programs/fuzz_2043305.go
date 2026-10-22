package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint32; F2 int16 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS5 struct { F0 bool; F1 int; F2 int16 }
type GenS6 struct { F0 int16; F1 uint }

func main() {
    var v1 uint8 = 121
    var v2 string = "a"
    var v3 int16 = 1000
    v4 := make(map[string]bool)
    v5 := []int{}
    if true {
        if true {
            v2 = v2
            var v6 uint16 = 682
            _ = v6
        } else {
            var v7 string = v2
            v8 := GenS1{F0: 0}
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(v2)
            _ = fi9
        }
    } else {
        for fi10 := 0; fi10 < 12; fi10++ {
            if false {
                var v11 int64 = 42
                fmt.Println(v11, v1, fi10)
                var v12 float64 = 0.1
                _ = v11
                _ = v12
            } else {
                var v13 int = (fi10 % (((fi10 % ((fi10) | 1))) | 1))
                _ = v13
            }
            if (42 > 1) {
                v14 := GenS1{F0: 1}
                fmt.Println(v2)
                _ = v14
            }
            _ = fi10
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS2{F0: (!true), F1: 42, F2: (v3 - 234)}
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int16 = int16(fi17)
            _ = fi17
            _ = v18
        }
    }
    v19 := GenS2{F0: ((!false) || false), F1: 529, F2: -908}
    if false {
        var v20 uint = 4294967295
        fmt.Println(v20, v2)
        for fi21 := 0; fi21 < 11; fi21++ {
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 string = ""
                _ = fi22
                _ = v23
            }
            v24 := GenS3{F0: "foo bar", F1: 2}
            v25 := GenS4{F0: 2, F1: 559, F2: 42}
            _ = fi21
            _ = v24
            _ = v25
        }
        _ = v20
    } else {
        if len(v5) > 0 {
            v5[0] = 0
        }
        v26 := GenS2{F0: false, F1: 2147483647, F2: (^(v3))}
        _ = v26
    }
    v3 -= v3
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v27 uint64 = 850
    if (v3 != v3) {
        if (!((0.1 + -2.25) <= 0.1)) {
            v2 = v2
        } else {
            var v28 uint8 = v1
            fmt.Println(v28)
            var v29 float64 = (((1e10 / 1.5) - 1.5) - 3.14159)
            _ = v28
            _ = v29
        }
        v5 = append(v5, -2147483648)
        var v30 int16 = (v3 & v3)
        _ = v30
    } else {
        if false {
            var v31 bool = (9223372036854775807 >= 42)
            v31 = ((^(v27)) < 42)
            for fi32 := 0; fi32 < 6; fi32++ {
                var v33 bool = (fi32 != 2147483647)
                fmt.Println(v31, v2, v1)
                _ = fi32
                _ = v33
            }
            _ = v31
        }
        var v34 uint16 = 225
        v3 += (v3 / ((v3) | 1))
        _ = v34
    }
    if false {
        v2 = v2
        v3 = (v3 ^ (v3 << ((uint(v27)) & 63)))
    }
    v1 += (^(v1))
    if (false || true) {
        v35 := GenS4{F0: uint32(v1), F1: v27, F2: -539}
        fmt.Println(v2, v1, v27)
        _ = v35
    } else {
        for fi36 := 0; fi36 < 4; fi36++ {
            fmt.Println(v27, v1)
            var v37 int32 = -2147483648
            var v38 int16 = v3
            _ = fi36
            _ = v37
            _ = v38
        }
        if (42 == 182) {
            v27 -= (^(v27))
            v2 = v2
            if (v2 == v2) {
                var v39 float32 = 0.1
                for fi40 := 0; fi40 < 15; fi40++ {
                    var v41 bool = (!true)
                    _ = fi40
                    _ = v41
                }
                _ = v39
            } else {
                for fi42 := 0; fi42 < 17; fi42++ {
                    if (!false) {
                        v43 := GenS3{F0: v2, F1: 53}
                        _ = v43
                    }
                    v44 := GenS1{F0: 65535}
                    _ = fi42
                    _ = v44
                }
                v2 = v2
                var v45 float64 = -2.25
                _ = v45
            }
        } else {
            v2 = "hello"
        }
        fmt.Println(v27, v2)
    }
    for fi46 := 0; fi46 < 9; fi46++ {
        fmt.Println(v1, fi46, fi46)
        _ = len(v5)
        if true {
            if true {
                v4[((((v2 + ((v2 + v2) + v2)) + "\n\t") + v2) + v2)] = true
            } else {
                fmt.Println(v1)
                for fi47 := 0; fi47 < 13; fi47++ {
                    var v48 float64 = ((0.1 + (1.5 + (0.1 - 0.0))) + 0.0)
                    v1 -= v1
                    _ = fi47
                    _ = v48
                }
                switch (v2 + "foo bar") {
                case "a":
                    for fi49 := 0; fi49 < 19; fi49++ {
                        var v50 int64 = 1
                        switch fi46 {
                        case 9223372036854775807:
                            v27 -= 619
                        }
                        v50 += -280
                        _ = fi49
                        _ = v50
                    }
                case "":
                    v1 -= (v1 >> ((2) & 63))
                }
            }
            switch fi46 {
            case 32767:
                var v51 string = ""
                _ = v51
            case 232:
                v1 = (v1 | (v1 << ((65535) & 63)))
            case 4294967295:
                v3 = (v3 & (v3 >> ((255) & 63)))
            default:
                switch fi46 {
                case 4294967295:
                    v52 := GenS3{F0: "", F1: 762}
                    _ = v52
                case 180:
                    if (1.5 < 0.0) {
                        for fi53 := 0; fi53 < 7; fi53++ {
                            switch 127 {
                            case -2147483648:
                                fmt.Println(fi53)
                            case 4294967295:
                                v2 = "hello"
                            case 795:
                                switch (fi53 % ((fi46) | 1)) {
                                case 42:
                                    if true {
                                        v3 = v3
                                        v27 -= uint64(v1)
                                        v54 := GenS5{F0: false, F1: fi46, F2: v3}
                                        _ = v54
                                    } else {
                                        v55 := GenS6{F0: 319, F1: 618}
                                        _ = v55
                                    }
                                case 1:
                                    v2 = v2
                                case 461:
                                    _ = len(v4)
                                default:
                                    v56 := GenS2{F0: (!(!(1 > (-(fi53))))), F1: 0, F2: v3}
                                    _ = v56
                                }
                            default:
                                v3 += (-(v3))
                            }
                            _ = fi53
                        }
                    } else {
                        var v57 uint = 645
                        _ = v57
                    }
                case 127:
                    delete(v4, v2)
                }
            }
            var v58 int8 = -113
            _ = v58
        }
        _ = fi46
    }
    if false {
        v27 = (^(v27))
        v59 := v4["hello"]
        _ = v59
    }
    fmt.Println(v2)
    for fi60 := 0; fi60 < 2; fi60++ {
        v61 := GenS4{F0: 699, F1: (v27 >> ((42) & 63)), F2: 2147483647}
        fmt.Println(v1, v1, v27)
        v5 = append(v5, (fi60 % ((fi60) | 1)))
        _ = fi60
        _ = v61
    }
    var v62 uint16 = 0
    v63 := v4["foo bar"]
    v3 = (v3 % (((v3 / (((v3 / ((1000) | 1))) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
}
