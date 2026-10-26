package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint32 }
func genFn0() uint8 {
    var v6 float32 = (1.5 + -2.25)
    var v7 int8 = 0
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: ((("foo bar" + "\n\t") + "foo bar") + "foo bar"), F1: -32768}
            if (v6 <= v6) {
                var v11 uint = 65535
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        var v12 uint32 = 4294967295
        v6 += 3.14159
        _ = fi8
        _ = v12
    }
    _ = v6
    _ = v7
    return 93
}

func main() {
    v1 := []int{42, 2147483647, 109, -298}
    v2 := []int{1, 9223372036854775807, 1, 128}
    var v3 int16 = 0
    var v4 int16 = (^(v3))
    v5 := []uint32{2147483648}
    v3 = v3
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := GenS1{F0: (("" + "hello") + "\n\t"), F1: v3}
        var v15 int = fi13
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 float64 = 0.1
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            if (v16 <= v16) {
                v3 = -1
            } else {
                var v19 uint = 0
                _ = v19
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v20 := GenS2{F0: (fi17 ^ 614), F1: 127}
            _ = fi18
            _ = v20
        }
        if false {
            for fi21 := 0; fi21 < 18; fi21++ {
                if true {
                    v22 := GenS2{F0: (fi17 % ((9223372036854775807) | 1)), F1: genFn0()}
                    if false {
                        fmt.Println(v16)
                        switch fi17 {
                        case 2147483647:
                            fmt.Println(v16)
                        case 0:
                            fmt.Println(v4, fi17)
                        default:
                            var v23 string = ("a" + "foo bar")
                            _ = v23
                        }
                    }
                    if true {
                        var v24 int64 = 31
                        var v25 int16 = 1000
                        fmt.Println(v24, v25, v3)
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                }
                fmt.Println(v3, v16, v3)
                v3 = v3
                _ = fi21
            }
            fmt.Println(fi17, fi17)
            fmt.Println(fi17)
        }
        v26 := GenS1{F0: ("" + "hello"), F1: (v4 >> ((1) & 63))}
        _ = fi17
        _ = v26
    }
    if false {
        v4 = (-(v3))
        v3 = (-(v3))
    }
    var v27 int16 = (-(v3))
    if false {
        fmt.Println(v16, v3)
        if false {
            v28 := GenS3{F0: 0, F1: 856, F2: 0}
            if true {
                switch "foo bar" {
                case "a":
                    for fi29 := 0; fi29 < 8; fi29++ {
                        for fi30 := 0; fi30 < 1; fi30++ {
                            var v31 float32 = 1e10
                            v1 = append(v1, fi29)
                            _ = fi30
                            _ = v31
                        }
                        _ = fi29
                    }
                case "foo bar":
                    if (255 != 829) {
                        fmt.Println(v3)
                        if (2147483647 < 669) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v4 = (v4 % ((v3) | 1))
                            v16 += 0.1
                        } else {
                            v16 -= v16
                            fmt.Println(v3, v16)
                        }
                    }
                case "\n\t":
                    if (((65535 > 32767) || false) && true) {
                        v3 -= -32768
                        if true {
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v16, v16)
                            case "":
                                v32 := GenS1{F0: "a", F1: 807}
                                _ = v32
                            case "foo bar":
                                var v33 int8 = 0
                                _ = v33
                            default:
                                v16 = v16
                            }
                        }
                    }
                }
                if true {
                    if true {
                        if (!(42 != 583)) {
                            var v34 int32 = 0
                            var v35 uint = 539
                            _ = v34
                            _ = v35
                        }
                        var v36 float64 = float64(1e10)
                        switch "\n\t" {
                        case "foo bar":
                            var v37 uint16 = 65535
                            _ = v37
                        case "a":
                            v27 += v27
                        case "\n\t":
                            var v38 uint64 = 182
                            _ = v38
                        }
                        _ = v36
                    } else {
                        v16 += 1.5
                        _ = len(v1)
                        for fi39 := 0; fi39 < 7; fi39++ {
                            if true {
                                v16 = ((float64(-2.25) - (((v16 - 1e10) / v16) * v16)) - (v16 - (v16 + v16)))
                            }
                            for fi40 := 0; fi40 < 16; fi40++ {
                                v4 -= -850
                                _ = fi40
                            }
                            v27 = (v4 >> ((0) & 63))
                            _ = fi39
                        }
                    }
                    v27 += (-(v4))
                } else {
                    if ((32767 <= 9223372036854775807) || (76 != 827)) {
                        v1 = append(v1, 42)
                    } else {
                        var v41 bool = true
                        if (32767 <= 0) {
                            _ = len(v2)
                            v4 += v27
                            var v42 string = "a"
                            _ = v42
                        }
                        _ = v41
                    }
                    v43 := GenS4{F0: 2147483648}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v43
                }
            }
            _ = v28
        }
    }
    v44 := GenS2{F0: 42, F1: 100}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = (v27 << ((42) & 63))
    v27 += v3
    switch "foo bar" {
    case "foo bar":
        v16 -= 1e10
    default:
        v4 += (^(v27))
    }
    v4 = v27
    var v45 int32 = 330
    var v46 bool = true
    v47 := GenS1{F0: "foo bar", F1: v27}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
