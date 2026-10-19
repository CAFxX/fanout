package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 string }
type GenS3 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS4 struct { F0 int32; F1 int64; F2 bool }
type GenS5 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint8, p1 int, p2 int32) uint8 {
    if (-1 > 103) {
        v6 := GenS1{F0: 42, F1: 2}
        v7 := GenS1{F0: 0, F1: 1}
        var v8 uint16 = 63
        _ = v6
        _ = v7
        _ = v8
    }
    if (!(!true)) {
        v9 := []int{(p1 + int(p0)), -9223372036854775808}
        p1 = (p1 % ((1) | 1))
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 42}
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(p0, fi11)
            _ = fi12
        }
        fmt.Println(p2, p1)
        _ = fi11
    }
    var v13 uint32 = 841
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    return p0
}
func genFn1(p0 int, p1 int64) int32 {
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            if false {
                var v16 int = (p0 & 42)
                v17 := GenS1{F0: 1, F1: 374}
                v18 := GenS1{F0: 2147483648, F1: 2}
                _ = v16
                _ = v17
                _ = v18
            } else {
                v19 := []int{p0, 2147483647}
                _ = v19
            }
            _ = fi15
        }
        if true {
            v20 := []int{9223372036854775807, p0, (p0 * p0), (p0 | (fi14 / ((4294967295) | 1)))}
            v21 := GenS2{F0: fi14, F1: 272, F2: "\n\t"}
            p1 = p1
            _ = v20
            _ = v21
        }
        p0 = p0
        _ = fi14
    }
    var v22 bool = (0 < p0)
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(p0)
        _ = fi23
    }
    var v24 int64 = (p1 / ((42) | 1))
    _ = p0
    _ = p1
    _ = v22
    _ = v24
    return -1
}

func main() {
    var v1 float32 = (-2.25 * (0.1 + -2.25))
    var v2 uint = 255
    var v3 int16 = 42
    var v4 uint = 378
    v5 := make(map[string]uint32)
    v1 = v1
    v3 += (v3 & (v3 + (v3 ^ 209)))
    v25 := []int{-873, 0}
    var v26 uint16 = 453
    var v27 bool = (!(32767 > 42))
    var v28 int = -1
    var v29 bool = false
    var v30 uint = (v2 ^ 2)
    switch "\n\t" {
    case "foo bar":
        v31 := v5["\n\t"]
        _ = v31
    default:
        v26 = v26
    }
    _ = len(v5)
    if v27 {
        var v32 uint16 = 255
        var v33 uint8 = 18
        switch "" {
        case "hello":
            v29 = (0.1 >= v1)
        }
        _ = v32
        _ = v33
    } else {
        if (3.14159 <= (((-2.25 + float64(3.14159)) - (1.5 + (float64(1e10) / 1.0))) / 0.1)) {
            var v34 bool = true
            for fi35 := 0; fi35 < 5; fi35++ {
                var v36 float64 = 1.5
                _ = fi35
                _ = v36
            }
            _ = v34
        } else {
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            var v37 string = "foo bar"
            v26 -= v26
            _ = v37
        }
        if (127 >= 100) {
            if ((v26 > 1) || (((3.14159 + ((float64(((v1 + v1) - v1)) - 0.0) - 1e10)) / 1.5) != (1e10 - (-2.25 + 1.5)))) {
                v38 := GenS1{F0: 2, F1: 42}
                switch ("" + "hello") {
                case "\n\t":
                    v26 = (v26 >> ((v4) & 63))
                case "a":
                    v3 = (-(v3))
                default:
                    for fi39 := 0; fi39 < 12; fi39++ {
                        if (!v27) {
                            if v29 {
                                if false {
                                    if v29 {
                                        for fi40 := 0; fi40 < 19; fi40++ {
                                            fmt.Println(v3, v3, v26)
                                            var v41 uint32 = 65535
                                            _ = fi40
                                            _ = v41
                                        }
                                        var v42 int16 = (v3 | v3)
                                        v3 = v3
                                        _ = v42
                                    }
                                }
                                _ = len(v25)
                            } else {
                                v43 := GenS1{F0: 42, F1: 32}
                                v44 := GenS3{F0: v26, F1: (127 == 128), F2: 326}
                                var v45 bool = v27
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                        }
                        _ = fi39
                    }
                }
                if v27 {
                    v29 = (0 > -2147483648)
                }
                _ = v38
            }
            for fi46 := 0; fi46 < 13; fi46++ {
                var v47 uint16 = 1
                _ = fi46
                _ = v47
            }
            delete(v5, "a")
        } else {
            for fi48 := 0; fi48 < 9; fi48++ {
                if (!v29) {
                    switch (fi48 | 0) {
                    case 32767:
                        fmt.Println(v4)
                    default:
                        if ("a" >= "\n\t") {
                            var v49 int8 = 42
                            var v50 int = 1
                            _ = v49
                            _ = v50
                        }
                    }
                } else {
                    var v51 uint64 = 4294967295
                    _ = v51
                }
                _ = fi48
            }
            v26 = v26
            var v52 uint = v4
            _ = v52
        }
        v2 = 255
    }
    if v27 {
        v4 = 4294967295
        var v53 int16 = (-(v3))
        if len(v25) > 0 {
            v25[0] = (v28 * (v28 ^ 0))
        }
        _ = v53
    }
    v54 := GenS4{F0: -2147483648, F1: 0, F2: (1.5 >= 0.1)}
    for fi55 := 0; fi55 < 8; fi55++ {
        if (v27 || ((v26 % ((1) | 1)) > (v26 >> (((v4 / (((v4 ^ v30)) | 1))) & 63)))) {
            switch v28 {
            case -151:
                v26 = v26
            case 4294967295:
                v56 := GenS2{F0: (v28 % ((v28) | 1)), F1: 255, F2: ("hello" + ("a" + "hello"))}
                _ = v56
            case -895:
                for fi57 := 0; fi57 < 9; fi57++ {
                    v28 -= 128
                    _ = fi57
                }
            default:
                for fi58 := 0; fi58 < 20; fi58++ {
                    v59 := GenS5{F0: 4294967295, F1: 42}
                    var v60 uint8 = 44
                    _ = fi58
                    _ = v59
                    _ = v60
                }
            }
        }
        _ = fi55
    }
    v1 = float32((-2.25 / 3.14159))
    var v61 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v54
    _ = v61
}
