package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint; F2 uint8 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint16 }
func genFn0(p0 int, p1 uint32, p2 uint8) uint32 {
    v6 := []int{(^(p0))}
    _ = len(v6)
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 += fi7
        p1 = (p1 - (p1 - (p1 * (p1 + uint32(p2)))))
        _ = fi7
    }
    if true {
        var v8 int16 = 42
        p2 = p2
        _ = v8
    }
    p2 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}
func genFn1(p0 int32) uint8 {
    p0 = p0
    v9 := []int{740, 42}
    p0 -= (p0 + p0)
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int = (fi10 & (fi10 | fi10))
        if false {
            v12 := GenS1{F0: 1, F1: 2}
            var v13 float32 = 0.0
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: (v11 >> ((255) & 63))}
            _ = v14
        }
        p0 = -139
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v9
    return 255
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 127
    var v3 float64 = 3.14159
    var v4 uint8 = 2
    v5 := []int{-2147483648, 127, -838, -2147483648}
    switch "hello" {
    case "foo bar":
        _ = len(v1)
    case "a":
        v2 = 100
    default:
        if len(v5) > 0 {
            v5[0] = 968
        }
    }
    if false {
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 bool = true
            v17 := GenS3{F0: (v4 * (v4 ^ (^(v4))))}
            var v18 int64 = 42
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        switch "" {
        case "a":
            fmt.Println(v4, v3, v4)
        case "":
            if len(v5) > 0 {
                v5[0] = 42
            }
        case "foo bar":
            var v19 uint = 2
            _ = v19
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 bool = false
            v22 := v1[(fi20 + fi20)]
            switch "" {
            case "hello":
                for fi23 := 0; fi23 < 11; fi23++ {
                    switch (("foo bar" + ("" + ((("" + (("\n\t" + (((("hello" + "a") + ("foo bar" + ("foo bar" + "\n\t"))) + "foo bar") + "hello")) + "\n\t")) + "hello") + "\n\t"))) + "foo bar") {
                    case "\n\t":
                        var v24 uint = 163
                        _ = v24
                    default:
                        if v21 {
                            v4 -= v4
                            switch "" {
                            case "\n\t":
                                v21 = (false || false)
                            default:
                                var v25 int64 = -542
                                _ = v25
                            }
                        }
                    }
                    var v26 float64 = (v3 - v3)
                    _ = fi23
                    _ = v26
                }
            case "a":
                v22 = false
            }
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v27 := GenS1{F0: 32767, F1: 479}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (42 < 0) {
        var v28 int8 = 100
        if (32767 <= -839) {
            var v29 uint = 0
            _ = v29
        }
        if (v3 > -2.25) {
            v4 += (v4 << ((4294967295) & 63))
            for fi30 := 0; fi30 < 7; fi30++ {
                if (!(!false)) {
                    v4 += 128
                    _ = len(v1)
                    var v31 float32 = 0.1
                    _ = v31
                }
                _ = fi30
            }
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v28 = (v2 >> ((65535) & 63))
            }
        } else {
            v32 := GenS4{F0: 700, F1: 255, F2: (v4 >> ((876) & 63))}
            v33 := GenS2{F0: -1}
            _ = v32
            _ = v33
        }
        _ = v28
    }
    v34 := GenS5{F0: (-(v2)), F1: 65535, F2: 65535}
    var v35 uint32 = 4294967295
    var v36 bool = ((716 != 2) || (true && ((genFn1(-1) == v4) || false)))
    fmt.Println(v36, v4, v3)
    if false {
        switch -31 {
        case -2147483648:
            var v37 int = 2147483647
            _ = v37
        default:
            v36 = false
        }
        v1[4294967295] = v36
        if (((!false) && v36) && v36) {
            fmt.Println(v4, v35)
            if v36 {
                var v38 float64 = v3
                if v36 {
                    for fi39 := 0; fi39 < 12; fi39++ {
                        v1[int(v2)] = ((v36 && (!((1e10 - -2.25) < -2.25))) && (v2 >= -1))
                        for fi40 := 0; fi40 < 19; fi40++ {
                            switch (("foo bar" + "") + ("hello" + "")) {
                            case "hello":
                                var v41 uint16 = 458
                                _ = v41
                            default:
                                var v42 bool = false
                                _ = v42
                            }
                            var v43 int64 = -1
                            v5 = append(v5, 42)
                            _ = fi40
                            _ = v43
                        }
                        _ = fi39
                    }
                    var v44 int8 = 100
                    v45 := v1[4294967295]
                    _ = v44
                    _ = v45
                }
                for fi46 := 0; fi46 < 9; fi46++ {
                    var v47 int16 = -813
                    var v48 float32 = -2.25
                    for fi49 := 0; fi49 < 14; fi49++ {
                        v50 := GenS2{F0: (fi49 >> ((4294967295) & 63))}
                        _ = fi49
                        _ = v50
                    }
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                _ = v38
            }
        } else {
            v5 = append(v5, -238)
            for fi51 := 0; fi51 < 6; fi51++ {
                for fi52 := 0; fi52 < 17; fi52++ {
                    var v53 uint8 = v4
                    var v54 uint16 = 2
                    v54 -= (v54 % ((v54) | 1))
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                _ = fi51
            }
        }
    } else {
        var v55 uint64 = 0
        v56 := v1[-2147483648]
        _ = v55
        _ = v56
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 785)
    var v57 uint64 = 267
    if (v4 <= (^(v4))) {
        v35 = uint32(v4)
        var v58 uint = 255
        if v36 {
            if (!(!true)) {
                if true {
                    if (v58 != v58) {
                        v35 -= (v35 >> ((42) & 63))
                        v35 = v35
                        v59 := GenS4{F0: (v58 << ((v58) & 63)), F1: (v58 ^ (v58 - (v58 - (v58 << (((^(v58))) & 63))))), F2: 103}
                        _ = v59
                    } else {
                        v60 := GenS2{F0: 4294967295}
                        _ = v60
                    }
                    switch "foo bar" {
                    case "a":
                        switch "\n\t" {
                        case "hello":
                            v57 += (^(v57))
                        default:
                            for fi61 := 0; fi61 < 1; fi61++ {
                                v62 := GenS5{F0: v2, F1: (v35 % (((v35 / ((4294967295) | 1))) | 1)), F2: uint16(v58)}
                                _ = len(v5)
                                _ = fi61
                                _ = v62
                            }
                        }
                    case "\n\t":
                        v58 = v58
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                v35 += 4294967295
                delete(v1, 695)
            }
        } else {
            var v63 int8 = v2
            for fi64 := 0; fi64 < 2; fi64++ {
                _ = len(v1)
                _ = fi64
            }
            _ = v63
        }
        _ = v58
    } else {
        var v65 uint64 = 0
        if len(v5) > 0 {
            v5[0] = 620
        }
        _ = v65
    }
    var v66 uint = 300
    var v67 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v57
    _ = v66
    _ = v67
}
