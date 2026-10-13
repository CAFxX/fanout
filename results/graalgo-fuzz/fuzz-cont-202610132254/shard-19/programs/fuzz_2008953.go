package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 int32; F2 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS8 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int64, p1 uint8, p2 string) uint16 {
    var v7 uint16 = 2
    var v8 float32 = float32(1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]int)
    var v3 uint64 = 42
    var v4 float32 = float32((0.1 + (float64(0.0) - 3.14159)))
    var v5 int64 = -9223372036854775808
    v6 := make(map[int]int)
    v3 -= v3
    fmt.Println(v3)
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            fmt.Println(v4, fi9)
            var v11 uint32 = (v1 / ((65535) | 1))
            var v12 uint16 = 42
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 uint32 = 2
        fmt.Println(fi9)
        _ = fi9
        _ = v13
    }
    v14 := GenS1{F0: 68, F1: 228, F2: 37}
    var v15 bool = (2 == 128)
    var v16 float32 = v4
    var v17 uint64 = uint64(v1)
    v18 := GenS1{F0: 127, F1: 333, F2: 42}
    v15 = v15
    v17 += 2
    fmt.Println(v3)
    var v19 int8 = -128
    v17 = v3
    var v20 int16 = 1
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS1{F0: 1, F1: (-(fi21)), F2: 0}
            v24 := GenS1{F0: 128, F1: (^(fi22)), F2: 42}
            var v25 string = ""
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        v26 := GenS1{F0: 0, F1: (fi21 << ((2) & 63)), F2: -387}
        var v27 float32 = v16
        _ = fi21
        _ = v26
        _ = v27
    }
    var v28 uint16 = 2
    if (42 >= 255) {
        v29 := GenS1{F0: 120, F1: -657, F2: 0}
        for fi30 := 0; fi30 < 13; fi30++ {
            if v15 {
                var v31 int16 = (v20 | -1)
                v15 = (2 <= 0)
                if ((fi30 >> ((2) & 63)) < fi30) {
                    switch "foo bar" {
                    case "foo bar":
                        if (!true) {
                            fmt.Println(v3, v16, v15)
                            switch "foo bar" {
                            case "":
                                var v32 int16 = v20
                                _ = v32
                            }
                            v1 = v1
                        }
                    case "\n\t":
                        fmt.Println(v5)
                    case "":
                        fmt.Println(v16)
                    }
                }
                _ = v31
            } else {
                v2["hello"] = (fi30 / ((fi30) | 1))
                v33 := GenS2{F0: v28, F1: (v20 % ((v20) | 1)), F2: v19}
                delete(v6, (fi30 << ((uint(v3)) & 63)))
                _ = v33
            }
            v34 := GenS3{F0: 18446744073709551615}
            _ = fi30
            _ = v34
        }
        _ = v29
    } else {
        for fi35 := 0; fi35 < 3; fi35++ {
            _ = len(v2)
            _ = fi35
        }
        v36 := GenS1{F0: 128, F1: 1, F2: 0}
        if (((!(!(!(!false)))) && ((v1 * 4294967295) == v1)) || v15) {
            var v37 float64 = -2.25
            var v38 uint8 = 127
            if v15 {
                switch 127 {
                case 42:
                    v39 := []int{}
                    _ = v39
                case 32767:
                    switch ((("\n\t" + "foo bar") + "hello") + "hello") {
                    case "":
                        var v40 int32 = int32(v28)
                        _ = v40
                    case "\n\t":
                        var v41 uint8 = 1
                        _ = v41
                    case "foo bar":
                        v42 := []int{-817, -651}
                        _ = v42
                    }
                default:
                    _ = len(v6)
                }
                if v15 {
                    v20 += 1000
                    v43 := GenS4{F0: v5}
                    v44 := GenS5{F0: v3, F1: -1, F2: "hello"}
                    _ = v43
                    _ = v44
                } else {
                    for fi45 := 0; fi45 < 20; fi45++ {
                        for fi46 := 0; fi46 < 19; fi46++ {
                            var v47 uint8 = (v38 & v38)
                            if v15 {
                                var v48 bool = (!false)
                                var v49 bool = v15
                                var v50 bool = (!(v15 || (((v16 + v4) == v4) || (4294967295 > 18446744073709551615))))
                                _ = v48
                                _ = v49
                                _ = v50
                            } else {
                                switch "hello" {
                                case "a":
                                    var v51 uint32 = v1
                                    _ = v51
                                }
                                var v52 int32 = -1
                                _ = v52
                            }
                            _ = fi46
                            _ = v47
                        }
                        _ = fi45
                    }
                    v3 -= v17
                }
            } else {
                v19 += -4
                delete(v6, int(v1))
                v53 := GenS6{F0: uint16(v3)}
                _ = v53
            }
            _ = v37
            _ = v38
        } else {
            if (!v15) {
                switch "" {
                case "foo bar":
                    if ((int32(v28) > int32(v20)) || (false && (true && (2147483647 >= 42)))) {
                        switch "a" {
                        case "\n\t":
                            for fi54 := 0; fi54 < 1; fi54++ {
                                if false {
                                    switch (fi54 + fi54) {
                                    case 32768:
                                        fmt.Println(v16, v16, v4)
                                    case -1:
                                        var v55 int32 = -150
                                        _ = v55
                                    default:
                                        var v56 uint64 = uint64(v1)
                                        _ = v56
                                    }
                                    _ = len(v2)
                                }
                                if (!false) {
                                    v57 := GenS7{F0: 255, F1: (v28 - 1), F2: v19}
                                    _ = v57
                                } else {
                                    var v58 uint16 = 929
                                    _ = v58
                                }
                                fmt.Println(v20, fi54)
                                _ = fi54
                            }
                        default:
                            var v59 uint64 = (v3 | v3)
                            _ = v59
                        }
                        var v60 uint32 = (v1 >> ((255) & 63))
                        var v61 uint16 = 1
                        _ = v60
                        _ = v61
                    } else {
                        for fi62 := 0; fi62 < 18; fi62++ {
                            var v63 uint16 = (v28 + genFn0(v5, 2, ("foo bar" + "foo bar")))
                            v64 := []int{(^(fi62)), fi62, (fi62 | (fi62 << ((18446744073709551615) & 63))), fi62}
                            v65 := GenS8{F0: 2, F1: (v28 * 255)}
                            _ = fi62
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        v17 += (v17 % ((v17) | 1))
                    }
                case "\n\t":
                    var v66 string = "\n\t"
                    _ = v66
                case "a":
                    var v67 int32 = -95
                    _ = v67
                default:
                    v3 += v3
                }
                switch "hello" {
                case "foo bar":
                    v3 = (v17 / (((v17 | 801)) | 1))
                case "a":
                    if (true || (127 != 115)) {
                        if v15 {
                            v6[2147483647] = 42
                            for fi68 := 0; fi68 < 7; fi68++ {
                                delete(v6, 4294967295)
                                _ = fi68
                            }
                        }
                    } else {
                        var v69 float32 = 1.5
                        var v70 uint64 = v17
                        v19 = -63
                        _ = v69
                        _ = v70
                    }
                case "\n\t":
                    var v71 uint8 = 2
                    _ = v71
                default:
                    v28 -= 758
                }
                fmt.Println(v3, v16, v1)
            }
            for fi72 := 0; fi72 < 10; fi72++ {
                if (!(v20 >= 1000)) {
                    if ((v20 == v20) || (!v15)) {
                        v73 := GenS1{F0: 128, F1: (fi72 % (((fi72 >> ((632) & 63))) | 1)), F2: 2147483647}
                        v74 := []int{fi72, 112, fi72, int(v20)}
                        _ = v73
                        _ = v74
                    } else {
                        v75 := GenS5{F0: v17, F1: 1, F2: ""}
                        _ = v75
                    }
                    v16 = (v4 + 3.14159)
                    v17 -= v17
                }
                for fi76 := 0; fi76 < 1; fi76++ {
                    v1 += v1
                    v77 := GenS1{F0: 100, F1: int(v5), F2: 562}
                    v6[fi72] = fi76
                    _ = fi76
                    _ = v77
                }
                v78 := v6[fi72]
                _ = fi72
                _ = v78
            }
        }
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
