package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0(p0 uint8, p1 uint8) uint {
    var v7 int8 = 0
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint32 = 2
        for fi10 := 0; fi10 < 9; fi10++ {
            v7 += v7
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float64 = 1.5
    var v3 uint32 = (v1 / ((v1) | 1))
    var v4 uint = 4294967295
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]uint32)
    v11 := []int{128, 2147483647}
    var v12 bool = (!(-128 == 1))
    v12 = true
    v5[811] = 1
    if (("hello" + "foo bar") >= ("foo bar" + (("" + "foo bar") + ""))) {
        v13 := GenS1{F0: -1, F1: 9223372036854775807}
        if (v12 && (2147483647 != -2147483648)) {
            v1 -= v1
            var v14 int8 = -33
            _ = v14
        }
        v1 -= v1
        _ = v13
    }
    if ((!(v12 || (!(255 == 42)))) || (2 >= 65535)) {
        v15 := GenS1{F0: -1, F1: 0}
        v4 -= genFn0(199, 158)
        _ = v15
    } else {
        v1 -= (v1 - 46)
        var v16 int32 = -2147483648
        _ = v16
    }
    v1 -= v1
    for fi17 := 0; fi17 < 11; fi17++ {
        if (!v12) {
            var v18 uint32 = v3
            _ = v18
        } else {
            v19 := GenS1{F0: -1, F1: -9223372036854775808}
            _ = v19
        }
        v2 -= 3.14159
        _ = fi17
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 uint16 = 65535
        for fi22 := 0; fi22 < 20; fi22++ {
            if (true || ((fi20 + fi20) > (fi22 & fi20))) {
                if true {
                    var v23 float32 = 3.14159
                    if len(v11) > 0 {
                        v11[0] = fi20
                    }
                    _ = v23
                }
                var v24 float64 = (v2 / (float64(0.1) - v2))
                v25 := GenS1{F0: 42, F1: 9223372036854775807}
                var v26 int32 = 294
                _ = len(v11)
                _ = v24
                _ = v25
                _ = v26
            } else {
                var v27 float64 = v2
                _ = v27
            }
            if ((2147483647 > -1) && (((fi22 * fi20) < (-(fi20))) && (-1 == 42))) {
                v1 -= 216
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        v28 := GenS1{F0: 1, F1: 42}
                        _ = v28
                    case "foo bar":
                        var v29 uint8 = 0
                        _ = v29
                    case "hello":
                        v21 = 920
                    }
                    var v30 float64 = v2
                    var v31 float64 = 3.14159
                    if v12 {
                        var v32 bool = (2 < 42)
                        _ = v32
                    } else {
                        switch (fi20 % (((^(fi20))) | 1)) {
                        case -9223372036854775808:
                            var v33 int32 = 2147483647
                            _ = v33
                        default:
                            var v34 uint64 = 315
                            _ = v34
                        }
                        var v35 uint64 = 1
                        var v36 int32 = 1
                        _ = v35
                        _ = v36
                    }
                    _ = v30
                    _ = v31
                } else {
                    var v37 float32 = ((0.1 / (0.1 + (3.14159 / 1.0))) + (1.5 - -2.25))
                    fmt.Println(fi22, fi22, fi22)
                    if v12 {
                        if (!false) {
                            var v38 uint = (v4 ^ (v4 + (v4 / (((v4 | v4)) | 1))))
                            if v12 {
                                var v39 float32 = ((v37 + v37) - 3.14159)
                                v6[v1] = v1
                                _ = v39
                            } else {
                                var v40 uint64 = 881
                                _ = v40
                            }
                            switch "hello" {
                            case "a":
                                _ = len(v5)
                            case "foo bar":
                                v37 += (((((0.0 / 3.14159) / (v37 - 0.0)) / 1.5) * v37) - -2.25)
                            case "":
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    var v41 uint8 = 218
                                    _ = v41
                                case "a":
                                    var v42 int32 = 1
                                    _ = v42
                                case "foo bar":
                                    if false {
                                        var v43 int = (-(fi22))
                                        _ = v43
                                    }
                                }
                            }
                            _ = v38
                        }
                        if ((0 == -89) || (255 < 42)) {
                            var v44 uint8 = 127
                            _ = v44
                        }
                        var v45 int = (fi22 * 509)
                        _ = v45
                    } else {
                        var v46 float32 = 0.1
                        _ = v46
                    }
                    v1 -= v3
                    if v12 {
                        v47 := GenS1{F0: -9223372036854775808, F1: -376}
                        _ = v47
                    }
                    _ = v37
                }
                var v48 uint = 42
                _ = v48
            }
            v49 := GenS1{F0: -672, F1: 0}
            _ = fi22
            _ = v49
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
