package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 bool; F2 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 uint64; F2 string }
type GenS8 struct { F0 int16; F1 string }
func genFn0(p0 int8) int8 {
    var v4 uint16 = 1
    v4 -= (v4 | (v4 - 1))
    _ = p0
    _ = v4
    return p0
}

func main() {
    v1 := []int{32767, 967}
    var v2 int32 = -5
    var v3 uint16 = 65535
    var v5 float32 = 3.14159
    var v6 uint64 = 552
    var v7 uint = 65535
    v8 := GenS1{F0: (v7 | 950), F1: (v2 * -317), F2: 556}
    if (v5 != (v5 * (float32(float64(-2.25)) / float32(3.14159)))) {
        v5 += v5
        v2 -= v2
    } else {
        var v9 uint = v7
        if false {
            if false {
                v10 := GenS2{F0: 2147483647, F1: 33}
                v9 = 313
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 uint16 = v3
                    for fi13 := 0; fi13 < 5; fi13++ {
                        var v14 float32 = float32((0.1 * 3.14159))
                        var v15 uint16 = 222
                        var v16 uint8 = 1
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    var v17 uint = (v7 - (^(v7)))
                    _ = fi11
                    _ = v12
                    _ = v17
                }
                _ = v10
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi18 := 0; fi18 < 19; fi18++ {
                v19 := GenS3{F0: v9}
                v3 += (^(v3))
                _ = fi18
                _ = v19
            }
        }
        _ = v9
    }
    if len(v1) > 0 {
        v1[0] = -713
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    v7 -= 2
    for fi20 := 0; fi20 < 17; fi20++ {
        if true {
            if false {
                v7 -= (^(v7))
                v3 -= (v3 + (v3 | (^(v3))))
            }
        } else {
            v1 = append(v1, fi20)
            _ = len(v1)
        }
        v7 = 18446744073709551615
        _ = len(v1)
        _ = fi20
    }
    switch "foo bar" {
    case "foo bar":
        v2 -= v2
    case "hello":
        v21 := GenS3{F0: v7}
        _ = v21
    }
    v5 += v5
    v22 := GenS3{F0: (v7 | (^(v7)))}
    switch 267 {
    case 127:
        for fi23 := 0; fi23 < 16; fi23++ {
            switch fi23 {
            case 9223372036854775807:
                var v24 int = fi23
                _ = v24
            case 1:
                var v25 int16 = 0
                _ = v25
            case -1:
                for fi26 := 0; fi26 < 17; fi26++ {
                    var v27 uint8 = 255
                    if false {
                        var v28 uint16 = v3
                        if false {
                            if ((42 <= 1) && ((true || (!(-9223372036854775808 < 0))) || (1 < 4294967295))) {
                                v1 = append(v1, fi26)
                                v3 += (v3 & 753)
                                var v29 string = ("a" + ("hello" + ""))
                                var v30 uint16 = 0
                                switch 32768 {
                                case 1:
                                    switch v29 {
                                    case "\n\t":
                                        var v31 int8 = 68
                                        _ = v31
                                    }
                                    v32 := GenS4{F0: 256, F1: false, F2: "\n\t"}
                                    _ = v32
                                case -251:
                                    v30 = 507
                                }
                                v33 := GenS5{F0: 944}
                                var v34 uint32 = 458
                                v35 := GenS2{F0: fi26, F1: (v27 - v27)}
                                v36 := GenS2{F0: (fi23 >> (((v7 | v7)) & 63)), F1: (v27 ^ v27)}
                                if (4294967295 < 127) {
                                    var v37 string = v29
                                    var v38 int64 = 1
                                    v34 -= v34
                                    _ = v37
                                    _ = v38
                                } else {
                                    var v39 uint64 = (v6 * (v6 % ((v6) | 1)))
                                    _ = v39
                                }
                                _ = v29
                                _ = v30
                                _ = v33
                                _ = v34
                                _ = v35
                                _ = v36
                            } else {
                                v40 := GenS6{F0: 32}
                                _ = v40
                            }
                        } else {
                            if (422 > 567) {
                                v41 := GenS1{F0: v7, F1: v2, F2: 1}
                                v42 := GenS7{F0: 547, F1: (^(v6)), F2: (("\n\t" + "\n\t") + "")}
                                _ = v41
                                _ = v42
                            } else {
                                v3 -= 32767
                            }
                            switch (^(fi26)) {
                            case -624:
                                if true {
                                    if false {
                                        var v43 uint16 = v3
                                        _ = v43
                                    }
                                    if (false || (!(0 < 1))) {
                                        v44 := GenS8{F0: -32768, F1: "foo bar"}
                                        _ = v44
                                    }
                                    var v45 int64 = 42
                                    _ = v45
                                }
                            }
                            switch ("hello" + "") {
                            case "foo bar":
                                var v46 bool = ("a" != "foo bar")
                                _ = v46
                            case "\n\t":
                                v3 = (v3 << (((v7 >> ((65535) & 63))) & 63))
                                switch 2147483647 {
                                case -9223372036854775808:
                                    v6 = v6
                                }
                            default:
                                switch (^(fi26)) {
                                case 32767:
                                    fmt.Println(v3)
                                case -728:
                                    v6 += uint64(v27)
                                }
                            }
                        }
                        v6 = v6
                        var v47 int16 = 32767
                        v48 := GenS1{F0: 1, F1: -492, F2: 1}
                        if true {
                            v49 := GenS8{F0: 42, F1: ""}
                            var v50 float64 = 1e10
                            var v51 int64 = 1
                            v5 -= (0.1 / ((v5 + v5) * (0.1 - v5)))
                            _ = v49
                            _ = v50
                            _ = v51
                        } else {
                            switch 2147483647 {
                            case -859:
                                v28 = v3
                            default:
                                var v52 string = "foo bar"
                                _ = v52
                            }
                        }
                        v5 = -2.25
                        var v53 int16 = v47
                        v28 = uint16(v6)
                        _ = v28
                        _ = v47
                        _ = v48
                        _ = v53
                    }
                    v6 = v6
                    _ = fi26
                    _ = v27
                }
            default:
                fmt.Println(fi23)
            }
            _ = fi23
        }
    case 42:
        v5 = float32(1.5)
    case -1:
        _ = len(v1)
    default:
        v54 := GenS7{F0: 65535, F1: 975, F2: "a"}
        _ = v54
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
