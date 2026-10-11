package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int }
func genFn0(p0 uint16, p1 uint8, p2 uint8) uint8 {
    switch -64 {
    case 0:
        var v6 float32 = -2.25
        _ = v6
    case 273:
        v7 := GenS1{F0: -82, F1: 2}
        _ = v7
    default:
        v8 := []int{}
        _ = v8
    }
    switch 128 {
    case 816:
        v9 := []int{241, 32768}
        _ = v9
    case -2147483648:
        var v10 uint16 = 118
        _ = v10
    case 0:
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint16 = 84
            _ = fi11
            _ = v12
        }
    default:
        switch 0 {
        case 127:
            v13 := []int{127, 42, 178}
            _ = v13
        case -445:
            for fi14 := 0; fi14 < 5; fi14++ {
                if false {
                    var v15 uint32 = 2147483648
                    v16 := GenS2{F0: 4294967295, F1: 0, F2: 0}
                    _ = v15
                    _ = v16
                }
                var v17 bool = false
                _ = fi14
                _ = v17
            }
        }
    }
    p2 -= p1
    if false {
        v18 := []int{128, 239, 128, 128}
        _ = v18
    } else {
        v19 := []int{4294967295, -304}
        if len(v19) > 0 {
            v19[0] = -593
        }
        _ = v19
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1(p0 bool, p1 int16) string {
    v20 := GenS3{F0: (p1 - 32767)}
    v21 := GenS4{F0: 65535, F1: 1}
    p0 = p0
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return ("hello" + "")
}

func main() {
    var v1 string = ("foo bar" + ("" + "a"))
    var v2 float64 = 1.5
    var v3 float64 = (v2 * (((v2 / 0.1) + 0.1) + v2))
    var v4 uint64 = 2
    var v5 int64 = 43
    var v22 uint16 = 354
    for fi23 := 0; fi23 < 17; fi23++ {
        v24 := []int{(fi23 / ((fi23) | 1)), (fi23 - fi23)}
        var v25 float64 = 0.0
        v25 -= 3.14159
        _ = fi23
        _ = v24
        _ = v25
    }
    v4 = (v4 * v4)
    if (!(642 > 2147483647)) {
        switch 1 {
        case 32768:
            v26 := []int{0}
            _ = v26
        case -1:
            var v27 string = v1
            _ = v27
        case -92:
            var v28 uint16 = (^(v22))
            _ = v28
        default:
            for fi29 := 0; fi29 < 10; fi29++ {
                v5 = (v5 | (^(v5)))
                _ = fi29
            }
        }
        v30 := []int{32768, 32768}
        _ = v30
    }
    v31 := []int{-948, 0, 9223372036854775807}
    v5 += v5
    var v32 int = 2147483647
    v5 += (v5 + (v5 << ((65535) & 63)))
    v33 := GenS4{F0: 65535, F1: (v4 + 1)}
    if true {
        if false {
            v4 += (v4 % ((v4) | 1))
            var v34 bool = (v4 >= 42)
            _ = v34
        }
        switch 292 {
        case 42:
            var v35 int64 = (v5 - v5)
            _ = v35
        case 32768:
            _ = len(v31)
        }
    } else {
        v36 := GenS2{F0: 635, F1: v4, F2: (^(v4))}
        _ = v36
    }
    if len(v31) > 0 {
        v31[0] = v32
    }
    if true {
        var v37 float32 = ((0.0 / 1.5) / (0.1 / float32((0.0 + float64(1.5)))))
        if len(v31) > 0 {
            v31[0] = (v32 & v32)
        }
        if (v37 <= v37) {
            if true {
                switch v1 {
                case "\n\t":
                    v38 := GenS5{F0: (v32 * v32)}
                    _ = v38
                case "hello":
                    for fi39 := 0; fi39 < 7; fi39++ {
                        if true {
                            var v40 float64 = (1.5 / (v2 - float64(v37)))
                            _ = v40
                        } else {
                            for fi41 := 0; fi41 < 16; fi41++ {
                                var v42 uint16 = (^(v22))
                                _ = fi41
                                _ = v42
                            }
                            if len(v31) > 0 {
                                v31[0] = -2147483648
                            }
                        }
                        _ = fi39
                    }
                }
                if (1 == 255) {
                    for fi43 := 0; fi43 < 2; fi43++ {
                        fmt.Println(v2, v4, v5)
                        _ = fi43
                    }
                } else {
                    for fi44 := 0; fi44 < 5; fi44++ {
                        switch "hello" {
                        case "a":
                            _ = len(v31)
                        }
                        switch (fi44 / (((^(v32))) | 1)) {
                        case 32:
                            switch v1 {
                            case "\n\t":
                                v37 = float32(((-2.25 + 0.0) * v2))
                            case "foo bar":
                                _ = len(v31)
                            case "hello":
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                            }
                        default:
                            fmt.Println(v32)
                        }
                        _ = fi44
                    }
                    _ = len(v31)
                }
            }
            v4 -= v4
            for fi45 := 0; fi45 < 20; fi45++ {
                var v46 float32 = v37
                v2 = v3
                if true {
                    _ = len(v31)
                }
                _ = fi45
                _ = v46
            }
        }
        _ = v37
    } else {
        v22 -= (v22 % ((72) | 1))
        var v47 uint16 = v22
        _ = v47
    }
    var v48 float64 = (float64(3.14159) / 1e10)
    v4 = (v4 + (v4 & v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v48
    _ = v5
}
