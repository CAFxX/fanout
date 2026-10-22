package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    v4 := GenS1{F0: 65535}
    v1 += v1
    fmt.Println(v1)
    v5 := GenS2{F0: 2}
    var v6 uint8 = 255
    v6 = (v6 >> ((18446744073709551615) & 63))
    if false {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint64 = 42
            _ = fi7
            _ = v8
        }
        switch -9223372036854775808 {
        case 4294967295:
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 int32 = -2147483648
                if true {
                    var v11 int16 = 42
                    fmt.Println(fi9)
                    if false {
                        var v12 float64 = 0.0
                        _ = v12
                    }
                    _ = v11
                } else {
                    v6 += v6
                    fmt.Println(v1, v10)
                    var v13 bool = false
                    _ = v13
                }
                v6 = 42
                _ = fi9
                _ = v10
            }
        case 422:
            v6 += (^(v6))
        default:
            if true {
                if false {
                    if false {
                        v1 = v1
                    }
                    var v14 uint16 = 0
                    _ = v14
                }
            } else {
                v15 := []int{}
                if (668 <= 1) {
                    var v16 uint = 255
                    var v17 uint16 = 42
                    v18 := GenS3{F0: v1}
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v15
            }
        }
    } else {
        switch 127 {
        case 4294967295:
            for fi19 := 0; fi19 < 14; fi19++ {
                v20 := GenS3{F0: v1}
                _ = fi19
                _ = v20
            }
        case 127:
            v6 += (v6 & (v6 + v6))
        case 9223372036854775807:
            var v21 uint8 = (^(v6))
            _ = v21
        }
        v6 -= v6
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 string = ((("foo bar" + (("foo bar" + ("" + "\n\t")) + "hello")) + "hello") + ("\n\t" + "\n\t"))
            if (!false) {
                fmt.Println(v1, fi22)
                var v24 float32 = 0.0
                _ = v24
            } else {
                if ((1 <= 2) || false) {
                    var v25 int64 = -1
                    _ = v25
                }
                switch v23 {
                case "foo bar":
                    fmt.Println(v1)
                case "\n\t":
                    if false {
                        fmt.Println(fi22)
                        v23 = v23
                        fmt.Println(v1, v6, fi22)
                    }
                case "hello":
                    if false {
                        if ((v1 << ((65535) & 63)) < v1) {
                            var v26 int16 = -242
                            _ = v26
                        } else {
                            v6 += 1
                        }
                        v23 = (v23 + v23)
                        v27 := []int{}
                        _ = v27
                    }
                default:
                    v28 := GenS2{F0: (v6 + v6)}
                    _ = v28
                }
            }
            v6 = v6
            _ = fi22
            _ = v23
        }
    }
    fmt.Println(v1, v1)
    v2[65535] = 414
    v29 := []int{717, -2147483648, 32768, 102}
    if true {
        if ((^(v1)) < 2147483647) {
            for fi30 := 0; fi30 < 14; fi30++ {
                v1 -= v1
                switch "foo bar" {
                case "hello":
                    var v31 int = 13
                    _ = v31
                case "\n\t":
                    if ((1000 > 107) || false) {
                        v6 -= (^(v6))
                        if false {
                            v1 = -486
                            if false {
                                if (737 < 255) {
                                    fmt.Println(v1, v6, fi30)
                                } else {
                                    v32 := GenS4{F0: "foo bar", F1: v6, F2: 1}
                                    var v33 uint64 = 2
                                    _ = v32
                                    _ = v33
                                }
                            }
                        }
                        switch int(v6) {
                        case 32768:
                            v1 += (v1 * (v1 | v1))
                        default:
                            for fi34 := 0; fi34 < 17; fi34++ {
                                v35 := GenS2{F0: (v6 - 252)}
                                _ = fi34
                                _ = v35
                            }
                        }
                    } else {
                        var v36 uint8 = 16
                        fmt.Println(v1)
                        _ = v36
                    }
                }
                _ = fi30
            }
            var v37 float32 = (0.0 - 0.1)
            _ = v37
        }
    } else {
        if true {
            var v38 int16 = 42
            _ = v38
        } else {
            var v39 string = "foo bar"
            _ = v39
        }
        var v40 uint = 461
        if false {
            v1 += (v1 >> (((v40 >> (((^(v40))) & 63))) & 63))
            v41 := GenS3{F0: (v1 ^ v1)}
            _ = v41
        } else {
            v42 := GenS4{F0: ("a" + "\n\t"), F1: 2, F2: 2}
            v29 = append(v29, 4294967295)
            v40 += v40
            _ = v42
        }
        _ = v40
    }
    fmt.Println(v6, v6)
    if true {
        if (!(422 == 0)) {
            v6 = v6
        } else {
            for fi43 := 0; fi43 < 14; fi43++ {
                for fi44 := 0; fi44 < 5; fi44++ {
                    var v45 uint64 = 2
                    _ = fi44
                    _ = v45
                }
                _ = fi43
            }
            var v46 float64 = 0.1
            var v47 uint8 = (v6 << ((18446744073709551615) & 63))
            _ = v46
            _ = v47
        }
        delete(v3, "\n\t")
    } else {
        for fi48 := 0; fi48 < 14; fi48++ {
            switch "a" {
            case "foo bar":
                v29 = append(v29, (fi48 & fi48))
            default:
                v49 := GenS5{F0: 712}
                _ = v49
            }
            v6 = 255
            if false {
                if true {
                    if true {
                        v6 = 127
                        switch ("hello" + "\n\t") {
                        case "hello":
                            for fi50 := 0; fi50 < 5; fi50++ {
                                fmt.Println(v1, v6)
                                _ = fi50
                            }
                        }
                    }
                    if ((false && true) || true) {
                        if ((v1 == v1) || true) {
                            v6 += (v6 / (((v6 << ((0) & 63))) | 1))
                        } else {
                            fmt.Println(fi48, fi48)
                        }
                        var v51 uint32 = 65535
                        v52 := GenS2{F0: (v6 | 255)}
                        _ = v51
                        _ = v52
                    } else {
                        var v53 float32 = float32((3.14159 / float64(float32(3.14159))))
                        _ = v53
                    }
                }
                if (v1 <= v1) {
                    var v54 bool = true
                    var v55 bool = v54
                    for fi56 := 0; fi56 < 12; fi56++ {
                        var v57 uint8 = 127
                        _ = fi56
                        _ = v57
                    }
                    _ = v54
                    _ = v55
                }
            } else {
                v1 += v1
            }
            _ = fi48
        }
        for fi58 := 0; fi58 < 4; fi58++ {
            v6 += v6
            var v59 uint32 = 2147483647
            _ = fi58
            _ = v59
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
