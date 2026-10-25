package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 uint64; F1 bool; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint64; F1 int8; F2 int32 }
type GenS6 struct { F0 uint64; F1 bool }
func genFn0() uint {
    var v4 int = 9223372036854775807
    v4 += (v4 % (((v4 << ((65535) & 63))) | 1))
    _ = v4
    return 65535
}
func genFn1(p0 uint16) uint64 {
    p0 -= p0
    var v5 int32 = -2147483648
    var v6 int8 = 1
    v7 := GenS1{F0: 42, F1: ((("\n\t" + "") + "") + "a"), F2: 1}
    v6 -= -1
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 164
    v2 := []uint32{42}
    var v3 string = "\n\t"
    _ = len(v2)
    v2 = append(v2, 892)
    if true {
        for fi8 := 0; fi8 < 13; fi8++ {
            v1 = (v1 % ((65535) | 1))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi8
        }
        var v9 int8 = 127
        _ = v9
    } else {
        var v10 int64 = -246
        _ = v10
    }
    switch ((v3 + v3) + (v3 + (v3 + (v3 + "")))) {
    case "foo bar":
        v11 := GenS2{F0: 18446744073709551615, F1: (true || true), F2: 2}
        _ = v11
    case "":
        if true {
            var v12 float32 = (float32((1e10 - (0.0 * 1e10))) + 1.5)
            v13 := GenS2{F0: 1, F1: (!(!true)), F2: 447}
            _ = v12
            _ = v13
        } else {
            switch v3 {
            case "foo bar":
                switch 42 {
                case 0:
                    var v14 int = 910
                    _ = v14
                case 32767:
                    v1 = v1
                }
            }
            var v15 int = 2147483647
            fmt.Println(v3, v1, v1)
            _ = v15
        }
    }
    var v16 float32 = (-2.25 / 1e10)
    var v17 uint8 = 42
    fmt.Println(v17, v3)
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int64 = -1
        var v20 uint64 = 18446744073709551615
        if ((^(fi18)) < fi18) {
            if (v17 <= v17) {
                _ = len(v2)
                if (false && ((((3.14159 * 3.14159) - 1e10) / float64(v16)) != (((1.5 + 1e10) * 0.1) + (0.1 * (1e10 / 3.14159))))) {
                    var v21 int16 = -32768
                    _ = v21
                } else {
                    v3 = v3
                    switch v3 {
                    case "\n\t":
                        v22 := GenS2{F0: genFn1(v1), F1: true, F2: genFn0()}
                        _ = v22
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v17 -= v17
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (492 < v19) {
                    v2 = append(v2, 0)
                }
            }
            if (true && false) {
                if (1e10 <= v16) {
                    var v23 uint16 = (^(v1))
                    _ = v23
                }
                switch "hello" {
                case "hello":
                    v2 = append(v2, 502)
                }
                v24 := GenS3{F0: genFn1(32767)}
                _ = v24
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi25 := 0; fi25 < 13; fi25++ {
                    if (1 != 42) {
                        var v26 int64 = v19
                        if (int(v17) == fi18) {
                            switch "a" {
                            case "hello":
                                v17 = uint8(v20)
                            }
                            v17 = uint8(v1)
                            var v27 int64 = v19
                            _ = v27
                        } else {
                            fmt.Println(v26, v19, fi18)
                            _ = len(v2)
                            v17 = (^(v17))
                        }
                        _ = v26
                    } else {
                        v20 = (v20 | v20)
                    }
                    v1 += 459
                    _ = fi25
                }
            }
            for fi28 := 0; fi28 < 9; fi28++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(fi18)
                switch -9223372036854775808 {
                case -580:
                    if true {
                        fmt.Println(v16, v17, v16)
                        v20 += (v20 % (((^(v20))) | 1))
                        var v29 uint32 = 373
                        _ = v29
                    }
                case 32767:
                    v30 := GenS3{F0: (v20 + (v20 * (v20 * v20)))}
                    _ = v30
                default:
                    if false {
                        fmt.Println(v3, fi18, v17)
                        var v31 uint64 = genFn1(v1)
                        _ = v31
                    } else {
                        v16 += float32(1e10)
                        v19 = -1
                        if true {
                            if (1 <= uint32(v17)) {
                                if true {
                                    v17 -= (v17 | 255)
                                    fmt.Println(fi18)
                                    v16 += (((v16 - -2.25) - float32(-2.25)) - v16)
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v32 uint64 = (v20 - 4294967295)
                                    _ = v32
                                }
                                var v33 string = v3
                                _ = v33
                            }
                        } else {
                            var v34 int32 = 2147483647
                            _ = v34
                        }
                    }
                }
                v16 = v16
                _ = fi28
            }
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    v35 := GenS4{F0: 42, F1: 359}
    if false {
        v36 := GenS3{F0: 18446744073709551615}
        v37 := GenS2{F0: 42, F1: (-354 >= 42), F2: 983}
        var v38 int = -2147483648
        _ = v36
        _ = v37
        _ = v38
    } else {
        if len(v2) > 0 {
            v2[0] = 65535
        }
    }
    var v39 int32 = 808
    fmt.Println(v17)
    var v40 uint8 = v17
    var v41 uint8 = (v17 | (v40 & (v40 | v17)))
    var v42 float64 = (float64((1e10 + (v16 - float32((3.14159 * (-2.25 * -2.25)))))) / float64(3.14159))
    v39 = (v39 & (v39 % (((-(v39))) | 1)))
    if false {
        if (-32768 == 0) {
            v43 := GenS2{F0: genFn1(v1), F1: false, F2: 18446744073709551615}
            v2 = append(v2, 656)
            v44 := GenS5{F0: 0, F1: 113, F2: v39}
            _ = v43
            _ = v44
        } else {
            for fi45 := 0; fi45 < 5; fi45++ {
                var v46 int32 = (^(v39))
                v47 := GenS6{F0: 1, F1: (fi45 == fi45)}
                v3 = v3
                _ = fi45
                _ = v46
                _ = v47
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
