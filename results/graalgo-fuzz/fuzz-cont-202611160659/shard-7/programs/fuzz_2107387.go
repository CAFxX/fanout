package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int8 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS3 struct { F0 int8; F1 int32; F2 bool }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint32; F1 uint16; F2 uint32 }
func genFn0(p0 int32, p1 uint16, p2 uint32) uint16 {
    if false {
        v5 := GenS1{F0: p2, F1: 942, F2: 100}
        _ = v5
    }
    p2 = p2
    v6 := []int{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    v1 := []uint32{}
    v2 := []string{("foo bar" + ("\n\t" + "foo bar")), "a", (("foo bar" + "") + (("\n\t" + "") + ""))}
    var v3 bool = false
    var v4 int16 = 0
    var v7 int = 687
    v8 := GenS2{F0: genFn0(42, 2, 564), F1: -1, F2: 4294967295}
    fmt.Println(v7, v3, v4)
    v7 = 32768
    v7 = v7
    if (true && (0 == 255)) {
        var v9 uint64 = 0
        switch "foo bar" {
        case "foo bar":
            if (251 < v9) {
                for fi10 := 0; fi10 < 16; fi10++ {
                    if false {
                        for fi11 := 0; fi11 < 19; fi11++ {
                            if (3.14159 > (float32(1e10) / 1e10)) {
                                switch (fi11 % ((fi11) | 1)) {
                                case 586:
                                    v9 += v9
                                default:
                                    switch ("a" + "\n\t") {
                                    case "hello":
                                        switch fi10 {
                                        case 128:
                                            v7 += fi11
                                            v12 := GenS3{F0: 64, F1: -1, F2: ((!((350 != -1) && true)) || (true || (2147483647 == 2147483648)))}
                                            v9 = v9
                                            _ = v12
                                        case 127:
                                            fmt.Println(v4, fi11, v3)
                                        case -749:
                                            if (1 < 42) {
                                                var v13 string = "a"
                                                fmt.Println(v13, v4, v9)
                                                _ = v13
                                            }
                                        default:
                                            v4 = (v4 & v4)
                                        }
                                    case "\n\t":
                                        var v14 uint64 = v9
                                        if v3 {
                                            v7 += (fi10 >> ((18446744073709551615) & 63))
                                        } else {
                                            fmt.Println(v3, v7)
                                        }
                                        _ = v14
                                    case "foo bar":
                                        if v3 {
                                            var v15 uint8 = 0
                                            fmt.Println(fi11, v4)
                                            v7 += v7
                                            _ = v15
                                        }
                                    default:
                                        switch "hello" {
                                        case "foo bar":
                                            v9 = (v9 ^ 331)
                                        case "a":
                                            var v16 int = 128
                                            _ = v16
                                        case "\n\t":
                                            v7 = v7
                                        }
                                    }
                                }
                                v7 -= 42
                                switch "hello" {
                                case "hello":
                                    var v17 uint64 = v9
                                    _ = v17
                                case "foo bar":
                                    v18 := GenS4{F0: (v9 / (((v9 - (^(v9)))) | 1)), F1: 0}
                                    _ = v18
                                case "":
                                    if len(v2) > 0 {
                                        v2[0] = "\n\t"
                                    }
                                }
                                v3 = true
                                var v19 float64 = 1.5
                                _ = v19
                            } else {
                                fmt.Println(fi11, v7, v3)
                                if len(v2) > 0 {
                                    v2[0] = ("" + "")
                                }
                                var v20 bool = v3
                                if (v20 && (!false)) {
                                    fmt.Println(v3, fi10, v20)
                                    v3 = v3
                                } else {
                                    var v21 float64 = 1e10
                                    _ = v21
                                }
                                v20 = (946 > 42)
                                _ = v20
                            }
                            if len(v1) > 0 {
                                v1[0] = 2147483647
                            }
                            var v22 bool = (1 > -968)
                            _ = fi11
                            _ = v22
                        }
                        v23 := GenS5{F0: genFn0(0, 255, 805), F1: 42}
                        v7 = fi10
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 10; fi24++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi24
                    }
                    _ = fi10
                }
                for fi25 := 0; fi25 < 2; fi25++ {
                    if false {
                        for fi26 := 0; fi26 < 14; fi26++ {
                            fmt.Println(v4, v3)
                            fmt.Println(fi25, v7, fi25)
                            _ = fi26
                        }
                        var v27 string = "hello"
                        v28 := GenS4{F0: v9, F1: -1}
                        _ = v27
                        _ = v28
                    }
                    v29 := GenS3{F0: -128, F1: -2147483648, F2: v3}
                    switch v7 {
                    case 128:
                        if (!v3) {
                            _ = len(v1)
                            v7 = v7
                        }
                    default:
                        for fi30 := 0; fi30 < 12; fi30++ {
                            if ((("foo bar" <= "foo bar") || (!(42 == 255))) && v3) {
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                var v31 uint8 = 210
                                var v32 float64 = 0.1
                                _ = v31
                                _ = v32
                            }
                            fmt.Println(fi30)
                            var v33 bool = false
                            _ = fi30
                            _ = v33
                        }
                    }
                    _ = fi25
                    _ = v29
                }
            } else {
                for fi34 := 0; fi34 < 14; fi34++ {
                    for fi35 := 0; fi35 < 17; fi35++ {
                        if true {
                            v1 = append(v1, 2147483648)
                        }
                        var v36 uint8 = 2
                        _ = fi35
                        _ = v36
                    }
                    _ = fi34
                }
            }
        default:
            v7 = v7
        }
        if v3 {
            v9 = (v9 / ((4294967295) | 1))
            var v37 uint32 = 1
            _ = v37
        }
        _ = v9
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (1 <= 255) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v4, v4, v4)
        v4 = (v4 * v4)
    }
    fmt.Println(v3, v4)
    for fi38 := 0; fi38 < 8; fi38++ {
        fmt.Println(v3, fi38)
        _ = fi38
    }
    v7 -= 32768
    v39 := GenS3{F0: 127, F1: -97, F2: (v3 && v3)}
    v40 := GenS5{F0: 566, F1: 983}
    for fi41 := 0; fi41 < 6; fi41++ {
        v4 = (v4 * v4)
        var v42 int = -2147483648
        if (-673 > -323) {
            v2 = append(v2, ("\n\t" + ("a" + ("" + "a"))))
            var v43 int16 = (v4 + v4)
            v4 -= (v43 | 515)
            _ = v43
        }
        _ = fi41
        _ = v42
    }
    v44 := GenS6{F0: 2, F1: 820, F2: 629}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v7
    _ = v8
}
