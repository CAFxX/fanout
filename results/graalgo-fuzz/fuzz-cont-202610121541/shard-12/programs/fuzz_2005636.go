package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint32; F1 uint8; F2 int8 }
func genFn0(p0 bool, p1 int) uint {
    var v6 string = (("" + ("\n\t" + "hello")) + ("a" + "foo bar"))
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 string = v6
        var v9 int16 = -1
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v6, p1, p1)
    _ = p0
    _ = p1
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 42
    v3 := []string{"foo bar"}
    var v4 uint8 = v2
    v5 := make(map[string]int)
    v10 := GenS1{F0: (!(!true)), F1: 1}
    if false {
        v2 = v4
        var v11 int8 = -5
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        switch fi12 {
        case 0:
            var v13 uint16 = 65535
            _ = v13
        case 169:
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 bool = false
                fmt.Println(fi12)
                switch "foo bar" {
                case "a":
                    if v15 {
                        if true {
                            var v16 int64 = 9223372036854775807
                            _ = v16
                        } else {
                            v15 = v15
                            switch "" {
                            case "foo bar":
                                var v17 bool = v15
                                _ = v17
                            case "\n\t":
                                if v15 {
                                    var v18 uint8 = v2
                                    var v19 int8 = 12
                                    fmt.Println(fi12, v18, fi14)
                                    _ = v18
                                    _ = v19
                                }
                            }
                            var v20 uint = 120
                            _ = v20
                        }
                    } else {
                        var v21 float32 = (1.5 * 0.0)
                        switch "\n\t" {
                        case "a":
                            var v22 uint32 = 216
                            _ = v22
                        case "":
                            switch "" {
                            case "":
                                _ = len(v3)
                            default:
                                v2 += (v4 * (v4 ^ v2))
                            }
                        }
                        v4 = v4
                        _ = v21
                    }
                case "\n\t":
                    _ = len(v3)
                case "":
                    v4 = v4
                default:
                    v2 = v2
                }
                _ = fi14
                _ = v15
            }
        case 32768:
            v4 -= (^(v2))
        }
        v23 := GenS2{F0: 32767, F1: 9223372036854775807, F2: 1}
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 float64 = ((3.14159 - ((1.5 * -2.25) + -2.25)) - 3.14159)
            _ = fi24
            _ = v25
        }
        _ = fi12
        _ = v23
    }
    v26 := GenS1{F0: (!false), F1: 18446744073709551615}
    fmt.Println(v2, v2)
    delete(v1, "\n\t")
    if (genFn0(true, 546) >= genFn0(false, 463)) {
        delete(v5, "hello")
        var v27 int64 = 0
        _ = v27
    } else {
        v28 := GenS1{F0: false, F1: 18446744073709551615}
        if (!true) {
            if true {
                var v29 bool = (0 >= -1)
                var v30 float64 = 1.5
                _ = v29
                _ = v30
            }
            v31 := GenS3{F0: -128, F1: 549}
            _ = v31
        }
        _ = v28
    }
    v4 = (v4 * v4)
    if true {
        if ("\n\t" != "") {
            var v32 uint32 = 4294967295
            var v33 int16 = -1
            fmt.Println(v2, v33)
            _ = v32
            _ = v33
        } else {
            _ = len(v1)
        }
    } else {
        var v34 uint16 = 429
        fmt.Println(v34)
        fmt.Println(v4)
        _ = v34
    }
    v35 := GenS2{F0: 40, F1: 1, F2: 127}
    v2 += v2
    switch 127 {
    case 127:
        v36 := GenS4{F0: 42, F1: v2, F2: 100}
        _ = v36
    case 128:
        fmt.Println(v2, v2)
    case 32768:
        v2 = v4
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        var v38 bool = (!(!(921 != -1)))
        for fi39 := 0; fi39 < 17; fi39++ {
            if true {
                var v40 uint64 = 2
                _ = v40
            } else {
                v38 = true
            }
            _ = fi39
        }
        fmt.Println(v4, fi37)
        _ = fi37
        _ = v38
    }
    var v41 int64 = -207
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v4
    _ = v41
    _ = v5
}
