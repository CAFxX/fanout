package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 string; F1 bool; F2 int }
type GenS3 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS4 struct { F0 string; F1 string; F2 bool }
type GenS5 struct { F0 int16 }
func genFn0() int8 {
    if true {
        var v7 float64 = (0.0 + ((1.5 / 1.5) * 0.1))
        var v8 uint32 = 188
        for fi9 := 0; fi9 < 8; fi9++ {
            v8 -= v8
            v10 := GenS1{F0: "hello", F1: 42}
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    } else {
        v11 := GenS2{F0: ("a" + "foo bar"), F1: false, F2: 127}
        _ = v11
    }
    if true {
        if (!(42 <= 18446744073709551615)) {
            var v12 int16 = 32767
            _ = v12
        } else {
            var v13 int16 = 2
            if true {
                switch "\n\t" {
                case "foo bar":
                    switch 1 {
                    case 9223372036854775807:
                        v14 := []int{}
                        _ = v14
                    case 42:
                        switch 119 {
                        case 9223372036854775807:
                            var v15 uint64 = 1
                            _ = v15
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                v16 := []int{-874, 327, -891, 0}
                                _ = v16
                            default:
                                var v17 bool = (!true)
                                _ = v17
                            }
                        }
                    default:
                        var v18 uint8 = 127
                        _ = v18
                    }
                case "\n\t":
                    v13 += 32767
                default:
                    v13 = v13
                }
            } else {
                v13 -= (^(v13))
            }
            v19 := GenS1{F0: "hello", F1: (^(v13))}
            _ = v13
            _ = v19
        }
    }
    return 61
}
func genFn1(p0 int64, p1 uint32, p2 int16) uint32 {
    if ((0 == -2147483648) && (true || false)) {
        p1 += (p1 - (p1 | (p1 >> ((18446744073709551615) & 63))))
    }
    v20 := GenS2{F0: (("\n\t" + (("a" + ("a" + "a")) + "")) + (("foo bar" + ("foo bar" + "\n\t")) + "")), F1: (2147483647 < 42), F2: 2147483647}
    v21 := []int{-2147483648, 128}
    if (p2 == (p2 << ((65535) & 63))) {
        var v22 int8 = 1
        _ = v22
    } else {
        p1 += (^(p1))
        p2 = p2
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return (p1 ^ p1)
}

func main() {
    var v1 float64 = (0.0 / (1.5 + float64(float32(0.0))))
    var v2 uint8 = 200
    var v3 int16 = -1
    v4 := make(map[string]bool)
    var v5 int16 = v3
    var v6 float64 = v1
    v23 := GenS3{F0: 4294967295, F1: 999, F2: (v5 ^ (v3 << ((255) & 63)))}
    v1 = v1
    v6 = 1e10
    if (!(genFn0() == 1)) {
        if ((4294967295 < 1) || (210 == 127)) {
            v4["foo bar"] = true
            if true {
                v6 = v1
            } else {
                switch "foo bar" {
                case "\n\t":
                    v24 := v4[""]
                    _ = v24
                case "":
                    if false {
                        for fi25 := 0; fi25 < 20; fi25++ {
                            if false {
                                v3 = v3
                                if (!false) {
                                    v26 := []int{}
                                    _ = v26
                                }
                            }
                            var v27 uint8 = (v2 - (v2 >> ((314) & 63)))
                            for fi28 := 0; fi28 < 17; fi28++ {
                                v29 := v4[""]
                                _ = fi28
                                _ = v29
                            }
                            _ = fi25
                            _ = v27
                        }
                        var v30 string = "foo bar"
                        var v31 int = -2147483648
                        _ = v30
                        _ = v31
                    } else {
                        if (9223372036854775807 <= -9223372036854775808) {
                            for fi32 := 0; fi32 < 2; fi32++ {
                                v3 = (v5 >> ((0) & 63))
                                v33 := GenS4{F0: "\n\t", F1: "hello", F2: ((v6 * (3.14159 / v1)) <= v1)}
                                _ = fi32
                                _ = v33
                            }
                            for fi34 := 0; fi34 < 18; fi34++ {
                                var v35 uint = 18446744073709551615
                                _ = fi34
                                _ = v35
                            }
                            v36 := []int{1, 9223372036854775807, 2147483647}
                            _ = v36
                        } else {
                            var v37 uint32 = genFn1(-816, 2147483647, v5)
                            _ = v37
                        }
                        v38 := GenS2{F0: "\n\t", F1: true, F2: 0}
                        _ = v38
                    }
                }
                delete(v4, "\n\t")
                v1 = (v1 + (v1 - 0.1))
            }
        } else {
            var v39 int = -489
            var v40 int = (v39 << ((8) & 63))
            v1 -= v6
            _ = v39
            _ = v40
        }
        for fi41 := 0; fi41 < 2; fi41++ {
            var v42 uint8 = v2
            var v43 string = "\n\t"
            _ = fi41
            _ = v42
            _ = v43
        }
        v2 = 1
    } else {
        var v44 uint16 = 1
        var v45 float32 = (3.14159 / 1.5)
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        _ = len(v4)
        v5 -= v3
        fmt.Println(v1, v6)
        _ = fi46
    }
    v47 := []int{127, 203}
    for fi48 := 0; fi48 < 20; fi48++ {
        _ = len(v47)
        v2 += 127
        v2 += 197
        _ = fi48
    }
    switch 1 {
    case 1:
        switch "\n\t" {
        case "a":
            if false {
                v49 := GenS3{F0: 483, F1: 1, F2: v3}
                var v50 uint8 = v2
                var v51 uint16 = 2
                _ = v49
                _ = v50
                _ = v51
            }
        case "hello":
            switch "\n\t" {
            case "\n\t":
                var v52 uint8 = 248
                _ = v52
            case "foo bar":
                var v53 uint16 = 65535
                _ = v53
            case "":
                if (-1 < 0) {
                    switch 2147483647 {
                    case 4294967295:
                        var v54 uint = 0
                        _ = v54
                    case 32768:
                        v1 = 0.0
                    }
                } else {
                    v6 -= v1
                }
            default:
                if (((42 < 0) || (((float64((0.1 - 0.1)) * (v6 / ((v6 * v1) - v1))) + v1) != v6)) && false) {
                    for fi55 := 0; fi55 < 2; fi55++ {
                        for fi56 := 0; fi56 < 7; fi56++ {
                            if (!false) {
                                v57 := GenS1{F0: "hello", F1: v3}
                                _ = v57
                            }
                            if false {
                                if false {
                                    var v58 uint8 = 26
                                    _ = v58
                                } else {
                                    if true {
                                        v6 -= v6
                                        fmt.Println(v5, v1, fi56)
                                        if len(v47) > 0 {
                                            v47[0] = -927
                                        }
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        if ((-1 != 0) || true) {
                                            fmt.Println(v5)
                                            v3 += v5
                                            var v59 int32 = -2147483648
                                            _ = v59
                                        } else {
                                            var v60 int = 127
                                            var v61 uint16 = 1
                                            _ = v60
                                            _ = v61
                                        }
                                    case "\n\t":
                                        v5 += v5
                                    case "":
                                        if false {
                                            switch (("\n\t" + "foo bar") + "\n\t") {
                                            case "hello":
                                                v3 = 311
                                            }
                                        } else {
                                            if (("" + "\n\t") <= ("hello" + "")) {
                                                delete(v4, "foo bar")
                                                fmt.Println(fi55)
                                            }
                                        }
                                    default:
                                        if false {
                                            v62 := GenS5{F0: (v5 ^ (^(v5)))}
                                            _ = v62
                                        }
                                        v1 -= 1.5
                                    }
                                }
                            }
                            _ = fi56
                        }
                        _ = fi55
                    }
                }
            }
        case "":
            fmt.Println(v2, v6)
        }
    case -605:
        var v63 int32 = 0
        _ = v63
    default:
        if (18446744073709551615 != 18446744073709551615) {
            v5 -= v5
            var v64 int64 = -881
            _ = v64
        }
    }
    var v65 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v65
}
