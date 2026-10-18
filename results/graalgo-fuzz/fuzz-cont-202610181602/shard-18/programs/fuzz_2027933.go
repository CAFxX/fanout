package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 int16; F1 int16; F2 int32 }
type GenS9 struct { F0 string; F1 uint32; F2 uint64 }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 uint64; F1 int64; F2 int64 }
type GenS12 struct { F0 int8 }
func genFn0(p0 int8) int16 {
    if false {
        p0 += p0
        p0 = (^(p0))
    }
    var v5 uint32 = 0
    v6 := []int{int(p0), 905, -2147483648}
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint32 = 65535
        if (178 < 42) {
            switch fi7 {
            case 9223372036854775807:
                for fi9 := 0; fi9 < 1; fi9++ {
                    var v10 uint64 = 2
                    v8 = v5
                    fmt.Println(p0, fi7, p0)
                    _ = fi9
                    _ = v10
                }
            case 42:
                if len(v6) > 0 {
                    v6[0] = (-(fi7))
                }
            case 32768:
                if len(v6) > 0 {
                    v6[0] = fi7
                }
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int64, p1 int64, p2 int64) int16 {
    var v11 uint32 = 2147483647
    var v12 float64 = 0.1
    fmt.Println(p2, p0, p0)
    var v13 int64 = p0
    v14 := GenS1{F0: 1000}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 826
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[string]string)
    var v3 uint32 = 428
    var v4 uint64 = 1
    if false {
        var v15 uint16 = 42
        _ = len(v1)
        v3 = (v3 << ((721) & 63))
        _ = v15
    } else {
        v16 := GenS2{F0: 0, F1: 20, F2: 0}
        var v17 float64 = float64(1.5)
        _ = v16
        _ = v17
    }
    v3 -= 2
    var v18 uint = 42
    var v19 uint = (v18 << (((v18 >> ((18446744073709551615) & 63))) & 63))
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 uint32 = v3
        var v22 uint32 = (v21 | 65535)
        v23 := GenS3{F0: 255, F1: v4}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 int = -9223372036854775808
    var v25 uint64 = v4
    v19 = v18
    _ = len(v2)
    v18 += uint(v4)
    v25 += (v4 >> ((42) & 63))
    v26 := GenS4{F0: 127}
    if (1e10 == 0.1) {
        v27 := GenS1{F0: -378}
        var v28 bool = true
        _ = v27
        _ = v28
    } else {
        var v29 float32 = (1e10 + (((0.1 * (0.0 * 1e10)) - 3.14159) - 0.0))
        v4 += v25
        v30 := GenS1{F0: -32768}
        _ = v29
        _ = v30
    }
    switch (^(v24)) {
    case -980:
        v31 := v2["a"]
        _ = v31
    }
    switch ("" + "") {
    case "":
        v24 += 4294967295
    }
    var v32 uint8 = 128
    switch v24 {
    case -9223372036854775808:
        for fi33 := 0; fi33 < 13; fi33++ {
            switch ("\n\t" + "hello") {
            case "":
                v32 = (^(v32))
            case "\n\t":
                for fi34 := 0; fi34 < 19; fi34++ {
                    _ = len(v2)
                    v35 := GenS5{F0: (v4 - v25), F1: "a"}
                    var v36 bool = (true || true)
                    var v37 uint64 = v25
                    if v36 {
                        fmt.Println(v36, fi33)
                    } else {
                        if (v36 && true) {
                            if ((-381 <= -2147483648) && v36) {
                                var v38 int = v24
                                _ = v38
                            } else {
                                switch (v24 / (((v24 ^ (fi33 / ((2147483647) | 1)))) | 1)) {
                                case 127:
                                    fmt.Println(v19, v19, v24)
                                case 453:
                                    if ((v32 << ((v18) & 63)) != (v32 / ((v32) | 1))) {
                                        var v39 int16 = 1
                                        if v36 {
                                            var v40 uint16 = 42
                                            _ = v40
                                        }
                                        var v41 int64 = 42
                                        var v42 uint = 0
                                        v1 = append(v1, "\n\t")
                                        if false {
                                            switch -432 {
                                            case 9223372036854775807:
                                                v41 = -1
                                            case 0:
                                                v43 := v2["foo bar"]
                                                _ = v43
                                            case 32768:
                                                v42 -= uint(v32)
                                            }
                                            v44 := GenS6{F0: 2147483647, F1: fi34}
                                            if v36 {
                                                var v45 uint16 = uint16(v37)
                                                _ = v45
                                            } else {
                                                var v46 string = "hello"
                                                if v36 {
                                                    delete(v2, v46)
                                                }
                                                v47 := GenS2{F0: 42, F1: 100, F2: fi34}
                                                if (1 != (^(v4))) {
                                                    switch ((v46 + ((v46 + (v46 + "hello")) + "foo bar")) + v46) {
                                                    case "foo bar":
                                                        fmt.Println(v25)
                                                    }
                                                } else {
                                                    fmt.Println(v46, v3, v41)
                                                }
                                                _ = v46
                                                _ = v47
                                            }
                                            v24 = 128
                                            _ = v44
                                        }
                                        _ = v39
                                        _ = v41
                                        _ = v42
                                    } else {
                                        v18 -= v18
                                        switch (fi33 >> ((v18) & 63)) {
                                        case 127:
                                            v48 := GenS7{F0: v32, F1: uint32(v19)}
                                            _ = v48
                                        case 699:
                                            var v49 int64 = 9223372036854775807
                                            _ = v49
                                        case 527:
                                            v19 = v19
                                        default:
                                            if v36 {
                                                fmt.Println(v3)
                                                v50 := GenS5{F0: v4, F1: ("a" + "")}
                                                _ = v50
                                            } else {
                                                v4 -= 0
                                            }
                                        }
                                    }
                                case -2147483648:
                                    var v51 bool = true
                                    _ = v51
                                }
                            }
                            var v52 uint16 = 719
                            fmt.Println(v3, fi33, v24)
                            var v53 int = (fi33 << (((v19 - v19)) & 63))
                            _ = v52
                            _ = v53
                        } else {
                            var v54 uint64 = (v25 | v25)
                            _ = v54
                        }
                        var v55 int16 = 42
                        _ = v55
                    }
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
            var v56 uint32 = 2
            _ = fi33
            _ = v56
        }
    case 9223372036854775807:
        for fi57 := 0; fi57 < 4; fi57++ {
            v58 := GenS2{F0: -1, F1: 1, F2: 32767}
            if (true || true) {
                for fi59 := 0; fi59 < 20; fi59++ {
                    fmt.Println(v24, v18)
                    v60 := GenS8{F0: -83, F1: genFn0(86), F2: -604}
                    v61 := GenS9{F0: ((("\n\t" + "hello") + "a") + (("hello" + ("hello" + "a")) + "\n\t")), F1: (v3 - (v3 / ((uint32(v32)) | 1))), F2: 42}
                    _ = fi59
                    _ = v60
                    _ = v61
                }
                v62 := GenS1{F0: -1}
                _ = v62
            } else {
                if false {
                    if false {
                        var v63 uint16 = 2
                        _ = v63
                    } else {
                        for fi64 := 0; fi64 < 16; fi64++ {
                            if false {
                                v65 := GenS2{F0: -105, F1: 0, F2: fi64}
                                _ = v65
                            }
                            v18 = (v18 % ((199) | 1))
                            var v66 bool = ((v25 * v25) == v25)
                            var v67 int = fi64
                            v4 = (v4 - (v4 / ((2) | 1)))
                            v68 := v2[("\n\t" + ("\n\t" + "\n\t"))]
                            _ = fi64
                            _ = v66
                            _ = v67
                            _ = v68
                        }
                        v69 := GenS10{F0: 32767, F1: (^(v32))}
                        _ = v69
                    }
                    switch (v24 << ((v18) & 63)) {
                    case -668:
                        var v70 uint32 = (v3 - (v3 - (v3 >> ((255) & 63))))
                        _ = v70
                    case -9223372036854775808:
                        switch fi57 {
                        case 128:
                            _ = len(v2)
                        case 32768:
                            _ = len(v2)
                        case 127:
                            var v71 string = "\n\t"
                            _ = v71
                        }
                    case 42:
                        v3 = (v3 / ((v3) | 1))
                    default:
                        v25 = v25
                    }
                    for fi72 := 0; fi72 < 10; fi72++ {
                        v1 = append(v1, ("" + "hello"))
                        v3 += v3
                        _ = fi72
                    }
                } else {
                    for fi73 := 0; fi73 < 18; fi73++ {
                        var v74 int64 = 168
                        _ = fi73
                        _ = v74
                    }
                    v24 -= v24
                }
                v4 = (v25 + v25)
                v75 := GenS9{F0: "\n\t", F1: (v3 / ((4294967295) | 1)), F2: v4}
                _ = v75
            }
            v76 := GenS11{F0: v25, F1: 1, F2: 1}
            _ = fi57
            _ = v58
            _ = v76
        }
    default:
        v77 := GenS12{F0: 42}
        _ = v77
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
}
