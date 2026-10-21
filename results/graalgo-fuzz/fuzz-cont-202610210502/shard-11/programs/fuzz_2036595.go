package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint16; F2 uint }
func genFn0(p0 uint32) int64 {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int32 = 511
        for fi9 := 0; fi9 < 6; fi9++ {
            if (1000 >= 0) {
                var v10 int32 = v8
                p0 -= (p0 % (((p0 >> ((65535) & 63))) | 1))
                _ = v10
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    v11 := []int{729}
    fmt.Println(p0)
    p0 += 42
    v12 := GenS1{F0: (false || (!true)), F1: 255}
    _ = p0
    _ = v11
    _ = v12
    return 818
}

func main() {
    var v1 uint = 255
    v2 := []string{("\n\t" + "\n\t"), ("a" + ("" + "foo bar")), "a"}
    var v3 int = 32767
    var v4 float64 = 3.14159
    var v5 bool = false
    var v6 string = "\n\t"
    fmt.Println(v3, v5, v6)
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 string = v6
        var v15 float32 = ((-2.25 - (1.5 / 1.0)) + 1e10)
        if (65535 <= 65535) {
            if v5 {
                var v16 int16 = -32768
                _ = len(v2)
                for fi17 := 0; fi17 < 7; fi17++ {
                    if (("\n\t" + v6) != v14) {
                        if (!false) {
                            var v18 float32 = v15
                            var v19 uint = 267
                            _ = v18
                            _ = v19
                        }
                        var v20 float64 = float64((0.1 - (0.1 * v15)))
                        if false {
                            var v21 bool = (!(127 > 0))
                            var v22 int32 = 1
                            _ = v21
                            _ = v22
                        } else {
                            fmt.Println(v15)
                        }
                        _ = v20
                    } else {
                        var v23 bool = true
                        switch 32768 {
                        case 128:
                            v2 = append(v2, "foo bar")
                        default:
                            v24 := GenS1{F0: v23, F1: 0}
                            _ = v24
                        }
                        _ = v23
                    }
                    switch v14 {
                    case "\n\t":
                        v25 := GenS1{F0: true, F1: 202}
                        _ = v25
                    case "":
                        switch 483 {
                        case -1:
                            v26 := GenS2{F0: v1}
                            _ = v26
                        default:
                            v1 -= v1
                        }
                    case "a":
                        v16 -= v16
                    default:
                        if ((37 >= 189) || (int(v16) == (-(v3)))) {
                            var v27 int = 127
                            v28 := GenS1{F0: false, F1: 2}
                            v14 = (v14 + "a")
                            _ = v27
                            _ = v28
                        } else {
                            if len(v2) > 0 {
                                v2[0] = ("hello" + ("" + (((v14 + "hello") + v14) + v6)))
                            }
                        }
                    }
                    fmt.Println(v5, v1, v16)
                    _ = fi17
                }
                _ = v16
            }
            var v29 string = v14
            var v30 bool = v5
            _ = v29
            _ = v30
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v31 int32 = 279
    var v32 uint32 = 65535
    _ = len(v2)
    var v33 bool = (v5 && (-1 != 54))
    v4 = v4
    v31 = v31
    var v34 uint = 1
    if v5 {
        switch "a" {
        case "foo bar":
            if v33 {
                _ = len(v2)
                for fi35 := 0; fi35 < 5; fi35++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi36 := 0; fi36 < 10; fi36++ {
                        var v37 float32 = -2.25
                        _ = fi36
                        _ = v37
                    }
                    var v38 bool = v33
                    _ = fi35
                    _ = v38
                }
            }
        }
        v39 := GenS3{F0: 1, F1: v31, F2: 18446744073709551615}
        _ = v39
    } else {
        fmt.Println(v32)
        fmt.Println(v33)
    }
    switch (v3 & (^(v3))) {
    case 0:
        switch (-(v3)) {
        case -295:
            v2 = append(v2, v6)
        }
    default:
        if ((v1 * v34) != 0) {
            var v40 bool = v5
            v32 -= (v32 - (v32 / ((v32) | 1)))
            _ = v40
        } else {
            if v5 {
                var v41 int32 = (-(v31))
                if len(v2) > 0 {
                    v2[0] = v6
                }
                v32 -= (^(v32))
                _ = v41
            }
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v31 -= (-(v31))
    var v42 uint = (v1 / ((v1) | 1))
    for fi43 := 0; fi43 < 4; fi43++ {
        for fi44 := 0; fi44 < 3; fi44++ {
            v45 := GenS4{F0: v32}
            fmt.Println(v33)
            v46 := GenS5{F0: 602, F1: 2, F2: (v34 & v1)}
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = fi43
    }
    var v47 float32 = (1e10 + (1e10 + 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v47
    _ = v5
    _ = v6
}
