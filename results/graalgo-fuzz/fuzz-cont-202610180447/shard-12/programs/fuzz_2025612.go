package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 bool }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int = 0
    v2 := make(map[string]uint32)
    v3 := []int{2147483647, 442, v1, (v1 * v1)}
    v4 := []uint32{}
    v5 := []int{}
    v6 := make(map[int]string)
    v7 := GenS1{F0: 952, F1: 0}
    fmt.Println(v1)
    if false {
        switch "a" {
        case "hello":
            v1 = v1
        case "":
            v8 := v6[v1]
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 18; fi9++ {
                v1 += 0
                _ = fi9
            }
        }
        var v10 uint32 = 423
        v10 += (v10 & 65535)
        _ = v10
    } else {
        if false {
            for fi11 := 0; fi11 < 7; fi11++ {
                v12 := v6[fi11]
                switch (v12 + ((v12 + v12) + v12)) {
                case "":
                    if false {
                        if false {
                            var v13 uint64 = 0
                            _ = v13
                        } else {
                            switch fi11 {
                            case 581:
                                _ = len(v6)
                            case -166:
                                v14 := GenS2{F0: 2, F1: 255, F2: ((54 < 21) && (false || false))}
                                _ = v14
                            case 485:
                                if true {
                                    v12 = v12
                                    v12 = v12
                                    for fi15 := 0; fi15 < 19; fi15++ {
                                        v12 = v12
                                        _ = fi15
                                    }
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = 127
                                    }
                                    var v16 uint64 = 42
                                    var v17 uint16 = 322
                                    _ = v16
                                    _ = v17
                                }
                            default:
                                switch v1 {
                                case 2147483647:
                                    _ = len(v6)
                                case -780:
                                    var v18 float64 = (0.0 - (float64(3.14159) + 1e10))
                                    _ = v18
                                case 128:
                                    var v19 uint8 = 255
                                    _ = v19
                                }
                            }
                        }
                        for fi20 := 0; fi20 < 16; fi20++ {
                            v12 = (v12 + v12)
                            _ = fi20
                        }
                        if false {
                            v12 = ("a" + (v12 + ""))
                            var v21 float64 = 3.14159
                            _ = v21
                        }
                    } else {
                        switch fi11 {
                        case 286:
                            v1 += 127
                        default:
                            var v22 int = fi11
                            _ = v22
                        }
                    }
                case "hello":
                    _ = len(v2)
                }
                var v23 float32 = 0.1
                _ = fi11
                _ = v12
                _ = v23
            }
        } else {
            v1 = v1
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            _ = len(v3)
            _ = fi24
        }
        for fi25 := 0; fi25 < 19; fi25++ {
            v1 -= v1
            if (21 == 0) {
                switch "a" {
                case "foo bar":
                    v1 -= (^(fi25))
                case "\n\t":
                    if false {
                        if ("" <= ("hello" + "a")) {
                            switch v1 {
                            case 2147483647:
                                v1 -= 0
                            default:
                                var v26 int64 = 0
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 17; fi27++ {
                                var v28 bool = true
                                v28 = (fi27 != -1)
                                if v28 {
                                    var v29 float64 = 1e10
                                    var v30 int8 = -128
                                    _ = v29
                                    _ = v30
                                } else {
                                    _ = len(v6)
                                    fmt.Println(fi25)
                                    v31 := GenS2{F0: 4294967295, F1: 766, F2: (255 != 65535)}
                                    _ = v31
                                }
                                _ = fi27
                                _ = v28
                            }
                            delete(v2, "\n\t")
                        }
                    } else {
                        v32 := v2["\n\t"]
                        var v33 int = fi25
                        _ = v32
                        _ = v33
                    }
                default:
                    fmt.Println(fi25)
                }
            } else {
                if true {
                    _ = len(v2)
                    var v34 int64 = 1
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 12; fi35++ {
                        var v36 int8 = 127
                        v36 = v36
                        _ = fi35
                        _ = v36
                    }
                    var v37 int = 127
                    _ = v37
                }
            }
            _ = fi25
        }
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        v39 := GenS1{F0: 32767, F1: 42}
        v1 = 127
        switch "hello" {
        case "foo bar":
            if (42 < 1) {
                var v40 int64 = 192
                _ = v40
            } else {
                v41 := GenS3{F0: 128}
                v1 = (fi38 ^ v1)
                var v42 uint64 = 13
                _ = v41
                _ = v42
            }
        case "hello":
            switch "foo bar" {
            case "a":
                if false {
                    for fi43 := 0; fi43 < 20; fi43++ {
                        _ = len(v6)
                        _ = fi43
                    }
                    v44 := GenS1{F0: 1, F1: 838}
                    _ = v44
                }
            default:
                v1 -= (fi38 % ((fi38) | 1))
            }
        case "\n\t":
            v45 := GenS3{F0: 1}
            _ = v45
        default:
            for fi46 := 0; fi46 < 9; fi46++ {
                fmt.Println(v1, fi46)
                _ = fi46
            }
        }
        _ = fi38
        _ = v39
    }
    v1 -= (v1 * -9223372036854775808)
    for fi47 := 0; fi47 < 17; fi47++ {
        var v48 int16 = 42
        var v49 int16 = (v48 * v48)
        fmt.Println(fi47, fi47, fi47)
        _ = fi47
        _ = v48
        _ = v49
    }
    v2[("foo bar" + "hello")] = 65535
    var v50 uint = 1
    if ((^(v50)) <= v50) {
        if false {
            v50 -= v50
        }
        if (9223372036854775807 == 19) {
            v51 := GenS3{F0: 0}
            _ = v51
        } else {
            fmt.Println(v50)
            if true {
                var v52 int = (v1 - 0)
                _ = v52
            }
        }
        v53 := GenS2{F0: 2147483647, F1: 336, F2: (2 <= 307)}
        _ = v53
    } else {
        v2[((("a" + ("a" + "\n\t")) + "foo bar") + "hello")] = 0
    }
    var v54 uint = v50
    var v55 uint8 = 203
    if (!(42 >= 255)) {
        v56 := GenS2{F0: 2, F1: 597, F2: (-32768 == -771)}
        _ = v56
    } else {
        for fi57 := 0; fi57 < 9; fi57++ {
            v55 -= (v55 / (((v55 << ((v54) & 63))) | 1))
            _ = fi57
        }
        switch "\n\t" {
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = (v1 / (((v1 * (v1 * (v1 >> ((100) & 63))))) | 1))
            }
        case "":
            var v58 int = v1
            _ = v58
        }
    }
    v50 += v50
    for fi59 := 0; fi59 < 20; fi59++ {
        v60 := GenS2{F0: 0, F1: 65535, F2: false}
        for fi61 := 0; fi61 < 17; fi61++ {
            v50 += 363
            _ = fi61
        }
        _ = fi59
        _ = v60
    }
    v62 := GenS1{F0: 2, F1: 284}
    v55 = v55
    v1 = v1
    if true {
        v63 := GenS2{F0: 144, F1: 919, F2: (!(!true))}
        v64 := v2[(("foo bar" + "\n\t") + "hello")]
        _ = v63
        _ = v64
    } else {
        var v65 uint8 = v55
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v54
    _ = v55
    _ = v6
    _ = v62
    _ = v7
}
