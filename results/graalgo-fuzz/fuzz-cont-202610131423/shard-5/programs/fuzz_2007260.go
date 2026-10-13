package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 string; F1 uint; F2 uint32 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8 }
func genFn0(p0 uint64) uint64 {
    var v7 string = ("\n\t" + ("a" + (("hello" + "foo bar") + "hello")))
    v7 = "a"
    _ = p0
    _ = v7
    return 4294967295
}
func genFn1(p0 uint16, p1 int16) uint16 {
    v8 := []int{2147483647}
    p1 -= -1
    p1 = (p1 % ((p1) | 1))
    v9 := GenS1{F0: 2}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    var v3 float64 = ((1e10 * 1.5) + 3.14159)
    v4 := []int{32767, 127, 32768, -893}
    var v5 uint8 = 0
    var v6 int16 = 1000
    v5 -= v5
    v5 += (v5 * v5)
    for fi10 := 0; fi10 < 15; fi10++ {
        v6 = (v6 >> ((1) & 63))
        _ = fi10
    }
    fmt.Println(v6, v6)
    v5 = v5
    var v11 uint16 = 44
    var v12 int64 = -682
    v3 = v3
    if false {
        switch -910 {
        case 1:
            _ = len(v4)
        case -475:
            var v13 uint8 = v5
            _ = v13
        default:
            var v14 uint8 = 0
            _ = v14
        }
    } else {
        var v15 uint8 = v5
        _ = v15
    }
    delete(v2, 65535)
    if (v12 >= (v12 / ((1) | 1))) {
        v16 := GenS2{F0: "a", F1: 9223372036854775807}
        _ = v16
    } else {
        v11 = v11
        if (0 == genFn0(18446744073709551615)) {
            v6 += (v6 / ((v6) | 1))
            var v17 int32 = -2147483648
            var v18 uint8 = (v5 | v5)
            _ = v17
            _ = v18
        } else {
            fmt.Println(v5, v3, v12)
        }
        switch "" {
        case "foo bar":
            var v19 uint8 = (v5 | v5)
            _ = v19
        case "hello":
            fmt.Println(v5)
        case "a":
            if (float32(((v3 + (1.5 + ((v3 + 0.1) / ((1e10 / v3) * -2.25)))) + v3)) >= (((3.14159 - -2.25) - (0.1 - 1e10)) + 1e10)) {
                v20 := GenS1{F0: 324}
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 int = fi21
                    var v23 int32 = 0
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = v20
            }
        default:
            switch "a" {
            case "hello":
                var v24 uint = 1
                _ = v24
            }
        }
    }
    fmt.Println(v11, v11)
    v25 := GenS1{F0: 884}
    var v26 uint = 255
    v27 := v2[42]
    var v28 int8 = 120
    for fi29 := 0; fi29 < 20; fi29++ {
        if true {
            v2[v27] = v27
            v5 += v5
        } else {
            fmt.Println(v5, v11, v11)
        }
        if false {
            for fi30 := 0; fi30 < 6; fi30++ {
                v31 := GenS3{F0: 32767, F1: (^(v26))}
                if true {
                    if false {
                        if (!(fi30 != (fi30 << (((v26 / ((653) | 1))) & 63)))) {
                            switch (fi29 - 4294967295) {
                            case 802:
                                if (793 != (^(v27))) {
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v28)
                                    case "\n\t":
                                        switch "" {
                                        case "hello":
                                            v11 = v11
                                        case "":
                                            if (!(v26 <= 0)) {
                                                var v32 string = "\n\t"
                                                switch (fi29 ^ fi29) {
                                                case 35:
                                                    var v33 int64 = (v12 - v12)
                                                    _ = v33
                                                case 1:
                                                    v26 += (v26 >> ((v26) & 63))
                                                }
                                                fmt.Println(fi30, v28, v3)
                                                _ = v32
                                            } else {
                                                if true {
                                                    v4 = append(v4, (fi30 % (((-(fi29))) | 1)))
                                                    var v34 uint64 = 1
                                                    v35 := GenS2{F0: "", F1: v12}
                                                    v36 := GenS3{F0: v6, F1: v26}
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                            }
                                        }
                                    case "":
                                        var v37 int = fi30
                                        _ = v37
                                    }
                                } else {
                                    if false {
                                        var v38 float32 = (1.5 + (1.5 / (-2.25 / (0.1 / 1.0))))
                                        v39 := GenS1{F0: 414}
                                        v40 := GenS2{F0: (("\n\t" + ("\n\t" + "\n\t")) + ("foo bar" + "\n\t")), F1: v12}
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    } else {
                                        var v41 int = -451
                                        delete(v1, v41)
                                        v42 := GenS4{F0: 973, F1: (v28 << ((uint(v5)) & 63))}
                                        _ = v41
                                        _ = v42
                                    }
                                }
                            case 1:
                                var v43 int64 = (v12 & 12)
                                _ = v43
                            case 900:
                                switch "\n\t" {
                                case "foo bar":
                                    v44 := GenS5{F0: ("hello" + ""), F1: (v26 - (^(v26))), F2: (v27 * (v27 * (v27 / (((^(v27))) | 1))))}
                                    _ = v44
                                }
                            default:
                                if false {
                                    var v45 float64 = v3
                                    v46 := GenS3{F0: 1000, F1: 65535}
                                    if false {
                                        v47 := GenS6{F0: (("" + "foo bar") + "foo bar"), F1: (^(v26))}
                                        if false {
                                            var v48 float64 = 0.1
                                            v49 := GenS7{F0: fi30}
                                            _ = v48
                                            _ = v49
                                        }
                                        v50 := GenS2{F0: "foo bar", F1: 42}
                                        _ = v47
                                        _ = v50
                                    }
                                    _ = v45
                                    _ = v46
                                }
                            }
                        } else {
                            var v51 uint64 = 42
                            var v52 int64 = v12
                            _ = v51
                            _ = v52
                        }
                        v53 := GenS8{F0: v5}
                        _ = v53
                    } else {
                        var v54 int32 = 316
                        v55 := GenS2{F0: ("hello" + ("hello" + "foo bar")), F1: (v12 << ((4294967295) & 63))}
                        _ = v54
                        _ = v55
                    }
                    var v56 uint = (v26 << ((v26) & 63))
                    _ = len(v2)
                    _ = v56
                }
                v12 = 9223372036854775807
                _ = fi30
                _ = v31
            }
            v26 -= 18446744073709551615
        }
        _ = fi29
    }
    v11 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
