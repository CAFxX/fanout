package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int64; F1 int32; F2 uint32 }
type GenS5 struct { F0 bool; F1 int64; F2 uint32 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int64 }
func genFn0(p0 uint32, p1 uint16) uint {
    v7 := GenS1{F0: 2147483647}
    if ((true && false) && (85 == -13)) {
        var v8 uint64 = 156
        if false {
            switch -1 {
            case -464:
                v9 := GenS2{F0: 2147483647, F1: 1}
                _ = v9
            case 32768:
                for fi10 := 0; fi10 < 11; fi10++ {
                    if true {
                        v11 := GenS1{F0: int32(v8)}
                        _ = v11
                    } else {
                        v12 := GenS3{F0: (v8 / ((v8) | 1)), F1: 2}
                        var v13 int32 = 1
                        v14 := []int{(fi10 + (-(fi10))), 4294967295}
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                }
            }
            var v15 float64 = 1.5
            _ = v15
        } else {
            var v16 int = 580
            v17 := []int{v16}
            _ = v16
            _ = v17
        }
        _ = v8
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS4{F0: 9223372036854775807, F1: 1, F2: (p0 | 2147483647)}
        if false {
            v20 := []int{}
            p0 += (p0 + (p0 / (((p0 ^ 65535)) | 1)))
            _ = v20
        }
        v21 := []int{(fi18 * 42), (fi18 / (((fi18 & fi18)) | 1))}
        _ = fi18
        _ = v19
        _ = v21
    }
    v22 := []int{127, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v22
    _ = v7
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    var v3 int64 = -1
    var v4 int16 = 0
    v5 := []int{}
    var v6 uint16 = 42
    switch "\n\t" {
    case "foo bar":
        v6 -= (v6 * (v6 * v6))
    case "\n\t":
        if (false || true) {
            for fi23 := 0; fi23 < 15; fi23++ {
                if false {
                    for fi24 := 0; fi24 < 14; fi24++ {
                        if (43 >= 5) {
                            v4 -= v4
                            var v25 uint16 = v6
                            _ = v25
                        } else {
                            v5 = append(v5, 128)
                            v3 = v3
                            var v26 uint8 = 174
                            _ = v26
                        }
                        delete(v1, ("a" + "\n\t"))
                        _ = fi24
                    }
                }
                v4 -= (v4 + (-(v4)))
                fmt.Println(v4)
                _ = fi23
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            var v27 uint64 = 684
            _ = v27
        }
    case "a":
        fmt.Println(v6, v6)
    default:
        var v28 int64 = (^(v3))
        _ = v28
    }
    switch "a" {
    case "":
        v29 := GenS5{F0: (genFn0(920, v6) <= 65535), F1: v3, F2: 4294967295}
        _ = v29
    case "foo bar":
        if true {
            switch 2147483647 {
            case 127:
                for fi30 := 0; fi30 < 11; fi30++ {
                    if true {
                        v31 := GenS4{F0: 0, F1: 2147483647, F2: 905}
                        _ = v31
                    }
                    _ = fi30
                }
            case 1:
                var v32 int16 = 1
                _ = v32
            default:
                v4 = v4
            }
            _ = len(v5)
            for fi33 := 0; fi33 < 4; fi33++ {
                if false {
                    var v34 int64 = -9223372036854775808
                    var v35 uint8 = 128
                    v5 = append(v5, fi33)
                    _ = v34
                    _ = v35
                } else {
                    switch fi33 {
                    case 2147483647:
                        var v36 int = fi33
                        _ = v36
                    default:
                        if true {
                            if (42 != -2147483648) {
                                var v37 int = (fi33 ^ 2147483647)
                                var v38 bool = ((v3 >> ((42) & 63)) <= v3)
                                v6 += v6
                                _ = v37
                                _ = v38
                            }
                            v1[(((("" + (("a" + "") + "\n\t")) + ((("foo bar" + "a") + "a") + "")) + "") + "")] = 2147483647
                            v39 := GenS6{F0: 255, F1: 153, F2: -226}
                            _ = v39
                        } else {
                            var v40 int32 = 0
                            for fi41 := 0; fi41 < 17; fi41++ {
                                v42 := GenS1{F0: (v40 | v40)}
                                v3 = (v3 * (v3 << ((genFn0(2147483648, v6)) & 63)))
                                _ = fi41
                                _ = v42
                            }
                            _ = v40
                        }
                    }
                    if ((((1e10 >= ((1.5 * float64(((0.0 + 0.0) - 0.1))) / 1e10)) && (false || (!true))) || (true || true)) && ((v6 >= v6) && (float32(0.0) >= ((((3.14159 / 0.1) * float32(1.5)) + float32((0.0 - -2.25))) * 3.14159)))) {
                        v43 := GenS6{F0: 127, F1: v6, F2: v3}
                        v44 := GenS2{F0: 0, F1: 151}
                        for fi45 := 0; fi45 < 4; fi45++ {
                            v3 -= int64(fi33)
                            var v46 float64 = (float64(1e10) / 1.0)
                            _ = fi45
                            _ = v46
                        }
                        _ = v43
                        _ = v44
                    }
                    for fi47 := 0; fi47 < 18; fi47++ {
                        var v48 float32 = float32(((3.14159 + ((0.0 / 1.5) + float64(3.14159))) / 1.0))
                        if true {
                            v5 = append(v5, (^(fi47)))
                        } else {
                            if (!true) {
                                v49 := GenS6{F0: 1, F1: v6, F2: 1}
                                var v50 uint32 = 945
                                v3 = (v3 << ((genFn0(v50, 988)) & 63))
                                _ = v49
                                _ = v50
                            }
                            v6 -= (v6 << ((214) & 63))
                            switch "\n\t" {
                            case "foo bar":
                                if (!false) {
                                    fmt.Println(v6, v4)
                                    var v51 int64 = (v3 * -78)
                                    _ = v51
                                }
                            }
                        }
                        var v52 int = (fi33 % (((-(fi33))) | 1))
                        _ = fi47
                        _ = v48
                        _ = v52
                    }
                }
                for fi53 := 0; fi53 < 16; fi53++ {
                    if len(v5) > 0 {
                        v5[0] = fi33
                    }
                    v4 += int16(fi53)
                    _ = fi53
                }
                _ = fi33
            }
        } else {
            if true {
                v5 = append(v5, 2147483647)
            }
        }
    default:
        var v54 uint64 = 970
        _ = v54
    }
    v6 = 32767
    var v55 int64 = v3
    var v56 bool = false
    switch "hello" {
    case "a":
        if v56 {
            var v57 float64 = 0.1
            _ = v57
        }
    case "":
        v58 := GenS3{F0: 679, F1: 42}
        _ = v58
    }
    delete(v2, "")
    var v59 bool = v56
    var v60 uint64 = 260
    fmt.Println(v3, v6)
    if (!v59) {
        v56 = (100 == -128)
    }
    _ = len(v5)
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v59
    _ = v6
    _ = v60
}
