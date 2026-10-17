package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint8; F1 string; F2 uint32 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 int64 = -1
    var v5 int = 127
    if true {
        if true {
            fmt.Println(v5, v5)
        } else {
            var v6 uint16 = 42
            for fi7 := 0; fi7 < 10; fi7++ {
                v5 = 119
                var v8 string = ((("hello" + "foo bar") + (("" + "hello") + ("\n\t" + "\n\t"))) + ("foo bar" + "hello"))
                delete(v1, v5)
                _ = fi7
                _ = v8
            }
            _ = v6
        }
    }
    var v9 uint = 42
    if true {
        switch v5 {
        case -602:
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 int16 = -868
                _ = fi10
                _ = v11
            }
        }
        _ = len(v1)
        var v12 uint32 = 1
        _ = v12
    } else {
        v13 := []int{v5, v5, v5, 156}
        _ = v13
    }
    fmt.Println(v4, v4, v9)
    v4 = (-(v4))
    if true {
        switch "foo bar" {
        case "a":
            v5 += v5
        case "hello":
            var v14 bool = false
            _ = v14
        default:
            var v15 uint64 = 743
            _ = v15
        }
    }
    var v16 uint8 = 127
    var v17 uint16 = uint16(v9)
    v18 := []int{(v5 * v5), (^(v5)), -876}
    if false {
        fmt.Println(v4)
        for fi19 := 0; fi19 < 11; fi19++ {
            switch v5 {
            case 1:
                v20 := GenS1{F0: 632, F1: v4, F2: 0}
                _ = v20
            default:
                v17 -= uint16(v16)
            }
            _ = fi19
        }
    } else {
        if (!false) {
            v5 = -2147483648
            for fi21 := 0; fi21 < 8; fi21++ {
                v4 = (v4 << ((v9) & 63))
                var v22 uint64 = 353
                fmt.Println(v5, v9, v4)
                _ = fi21
                _ = v22
            }
        } else {
            v23 := GenS1{F0: 616, F1: v4, F2: -1}
            var v24 uint = (^(v9))
            _ = v23
            _ = v24
        }
    }
    if (!(2147483647 <= 65535)) {
        v25 := GenS2{F0: false, F1: (v16 | (v16 / ((v16) | 1)))}
        v9 += uint(v17)
        _ = v25
    } else {
        v26 := GenS2{F0: ((^(v5)) == 1), F1: 255}
        _ = v26
    }
    delete(v2, -1)
    var v27 uint8 = v16
    var v28 uint32 = 796
    for fi29 := 0; fi29 < 6; fi29++ {
        if (-9223372036854775808 != 1) {
            if (v28 < v28) {
                var v30 uint64 = 2
                fmt.Println(v28, v27)
                _ = v30
            } else {
                v18 = append(v18, -2147483648)
            }
            v5 += v5
            fmt.Println(v4, v4)
        }
        if (!false) {
            var v31 string = "foo bar"
            switch (v5 + int(v16)) {
            case -9223372036854775808:
                if (uint16(v16) != v17) {
                    v32 := GenS3{F0: 127, F1: v31, F2: (v28 << ((v9) & 63))}
                    for fi33 := 0; fi33 < 10; fi33++ {
                        fmt.Println(v4)
                        var v34 uint32 = (v28 & (v28 + 213))
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                } else {
                    fmt.Println(v27, v5, v4)
                    v35 := v1[(v5 | (fi29 >> (((v9 & v9)) & 63)))]
                    if false {
                        for fi36 := 0; fi36 < 6; fi36++ {
                            if (1.5 != float32(((3.14159 - 3.14159) + (0.0 / (0.0 + 1e10))))) {
                                if true {
                                    if true {
                                        v37 := GenS4{F0: 127, F1: int64(v27)}
                                        v3[v31] = (2147483647 <= 42)
                                        _ = v37
                                    } else {
                                        _ = len(v18)
                                    }
                                    var v38 int = fi29
                                    _ = v38
                                }
                                fmt.Println(fi29, v4)
                                if false {
                                    v16 -= (^(v16))
                                    if ((v27 * v27) != v16) {
                                        switch (fi36 * fi29) {
                                        case -9223372036854775808:
                                            switch fi29 {
                                            case 32768:
                                                v18 = append(v18, (v5 << ((v9) & 63)))
                                            case 1:
                                                delete(v3, v35)
                                            }
                                        case 9223372036854775807:
                                            delete(v3, ("foo bar" + v31))
                                        }
                                        v31 = v35
                                    } else {
                                        v27 += 1
                                        if ((v9 & v9) <= (v9 + (v9 & v9))) {
                                            fmt.Println(v28)
                                            _ = len(v1)
                                            var v39 bool = (!true)
                                            _ = v39
                                        } else {
                                            v31 = (v35 + (v31 + v31))
                                            var v40 string = v35
                                            _ = v40
                                        }
                                        var v41 uint8 = (v16 / ((uint8(v28)) | 1))
                                        _ = v41
                                    }
                                    v42 := GenS5{F0: (0.0 >= 1.5)}
                                    var v43 int64 = (v4 | (^(v4)))
                                    _ = v42
                                    _ = v43
                                } else {
                                    v35 = v31
                                    fmt.Println(fi29, v31)
                                    var v44 string = "foo bar"
                                    _ = v44
                                }
                            }
                            _ = fi36
                        }
                        _ = len(v1)
                        if true {
                            fmt.Println(v9, fi29, v31)
                            v45 := GenS3{F0: v27, F1: ((("\n\t" + v31) + v35) + v35), F2: (v28 % ((v28) | 1))}
                            if false {
                                v18 = append(v18, v5)
                                v46 := GenS3{F0: 1, F1: v35, F2: v28}
                                fmt.Println(v35, v4, v27)
                                _ = v46
                            } else {
                                v28 += (v28 + 1)
                                if false {
                                    fmt.Println(fi29)
                                    for fi47 := 0; fi47 < 11; fi47++ {
                                        v3[(v35 + "hello")] = (((v28 * v28) < v28) && false)
                                        if ((false && true) && false) {
                                            v48 := v2[(fi47 + 130)]
                                            v31 = ((v35 + v31) + ("foo bar" + v35))
                                            delete(v3, (((v31 + "a") + v31) + (v35 + (("a" + v31) + v31))))
                                            _ = v48
                                        } else {
                                            switch (v5 + (fi47 >> ((1) & 63))) {
                                            case 821:
                                                if true {
                                                    v17 = 403
                                                } else {
                                                    var v49 int = (v5 & (-(fi29)))
                                                    v50 := v3[v31]
                                                    var v51 uint64 = 569
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                }
                                            default:
                                                v27 = v16
                                            }
                                            fmt.Println(v4, v17, v28)
                                        }
                                        _ = fi47
                                    }
                                    for fi52 := 0; fi52 < 17; fi52++ {
                                        var v53 float32 = -2.25
                                        _ = fi52
                                        _ = v53
                                    }
                                }
                            }
                            _ = v45
                        }
                    }
                    _ = v35
                }
            case 128:
                _ = len(v18)
            default:
                for fi54 := 0; fi54 < 4; fi54++ {
                    var v55 uint32 = (v28 % ((v28) | 1))
                    v18 = append(v18, (^(v5)))
                    _ = fi54
                    _ = v55
                }
            }
            _ = v31
        }
        _ = fi29
    }
    if (v9 > (v9 >> (((v9 % ((4294967295) | 1))) & 63))) {
        v16 = (v16 + 42)
    } else {
        v27 = 128
        var v56 uint8 = v27
        if false {
            var v57 float64 = 0.0
            v56 += 238
            v17 = (v17 - (v17 + (v17 % ((255) | 1))))
            _ = v57
        } else {
            for fi58 := 0; fi58 < 11; fi58++ {
                switch (-(fi58)) {
                case -9223372036854775808:
                    v59 := GenS5{F0: (int(v28) >= 969)}
                    _ = v59
                default:
                    fmt.Println(v27, v27)
                }
                _ = fi58
            }
            v60 := GenS6{F0: (v9 | (v9 >> ((v9) & 63)))}
            _ = v60
        }
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
