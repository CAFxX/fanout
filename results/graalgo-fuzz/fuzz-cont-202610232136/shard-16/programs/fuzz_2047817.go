package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0() uint32 {
    var v6 float64 = 0.0
    switch 128 {
    case 128:
        v6 -= (3.14159 * ((1.5 * 0.1) + ((v6 + float64(0.1)) / (0.0 / (v6 - v6)))))
    }
    if true {
        var v7 bool = (42 < 1)
        var v8 bool = v7
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: -1, F1: 1}
        var v10 bool = false
        _ = v10
        _ = v9
    }
    _ = v6
    return 251
}
func genFn1(p0 uint) string {
    p0 = 65535
    fmt.Println(p0)
    v11 := GenS1{F0: -9223372036854775808, F1: 0}
    v12 := GenS1{F0: 673, F1: 137}
    _ = p0
    _ = v11
    _ = v12
    return "a"
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]uint32)
    var v3 int = 128
    var v4 int8 = int8(v3)
    v5 := []string{"foo bar", "\n\t"}
    v1 = true
    var v13 uint8 = 19
    v14 := GenS1{F0: (v3 << ((65535) & 63)), F1: 42}
    var v15 int16 = 1000
    if (v1 || true) {
        var v16 string = "a"
        if false {
            v17 := GenS1{F0: 32768, F1: 65535}
            _ = v17
        } else {
            fmt.Println(v16)
        }
        _ = v16
    }
    var v18 int64 = 0
    if (18446744073709551615 < 2) {
        var v19 int8 = v4
        _ = v19
    } else {
        v3 -= 262
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 int = v3
            v2[2] = genFn0()
            for fi22 := 0; fi22 < 18; fi22++ {
                v23 := GenS2{F0: (v21 >= (-(v3))), F1: 47}
                switch (fi22 & -843) {
                case 9223372036854775807:
                    v24 := GenS3{F0: 255}
                    _ = v24
                case 3:
                    if (v1 || (!v1)) {
                        var v25 int16 = (v15 / ((-1) | 1))
                        var v26 float32 = 0.1
                        fmt.Println(v21)
                        _ = v25
                        _ = v26
                    }
                case 1:
                    switch "a" {
                    case "foo bar":
                        v15 = (v15 / ((int16(v18)) | 1))
                    case "":
                        if (v13 < v13) {
                            if (((^(v18)) > v18) || true) {
                                var v27 uint16 = 0
                                if (!(!true)) {
                                    if len(v5) > 0 {
                                        v5[0] = "hello"
                                    }
                                    if (!(true || (v1 && (v1 && (0.0 >= 0.0))))) {
                                        _ = len(v2)
                                        v18 = v18
                                    } else {
                                        v3 += (^(fi20))
                                    }
                                    v27 -= v27
                                } else {
                                    v5 = append(v5, "\n\t")
                                    v28 := v2[2147483647]
                                    _ = v28
                                }
                                _ = v27
                            } else {
                                _ = len(v2)
                            }
                        } else {
                            var v29 int8 = -1
                            v30 := GenS1{F0: v3, F1: 42}
                            _ = v29
                            _ = v30
                        }
                    }
                default:
                    fmt.Println(v21, v21)
                }
                var v31 float32 = 0.0
                _ = fi22
                _ = v23
                _ = v31
            }
            _ = fi20
            _ = v21
        }
        v15 += 537
    }
    if (v18 >= v18) {
        if (v1 && true) {
            for fi32 := 0; fi32 < 11; fi32++ {
                delete(v2, 958)
                v4 = v4
                _ = fi32
            }
            for fi33 := 0; fi33 < 17; fi33++ {
                v15 += (v15 & v15)
                var v34 bool = false
                v13 = (v13 * (v13 * v13))
                _ = fi33
                _ = v34
            }
        }
        for fi35 := 0; fi35 < 16; fi35++ {
            for fi36 := 0; fi36 < 8; fi36++ {
                if (42 > 583) {
                    v18 = v18
                    switch fi36 {
                    case 1:
                        switch "\n\t" {
                        case "":
                            if (395 != -1) {
                                if len(v5) > 0 {
                                    v5[0] = "foo bar"
                                }
                                var v37 uint8 = 118
                                _ = v37
                            } else {
                                var v38 uint64 = 42
                                fmt.Println(v38, fi36, v15)
                                _ = v38
                            }
                        case "foo bar":
                            fmt.Println(v18, v18, v1)
                        }
                    case 128:
                        v18 = v18
                    default:
                        if (int8(fi35) >= (^(v4))) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if v1 {
                                fmt.Println(fi35, v1, fi36)
                            }
                            var v39 float32 = 3.14159
                            _ = v39
                        }
                    }
                }
                if false {
                    var v40 float64 = 0.1
                    _ = v40
                } else {
                    v13 -= v13
                    v41 := GenS3{F0: v13}
                    _ = v41
                }
                v42 := GenS4{F0: 255, F1: (!v1)}
                var v43 uint = 42
                _ = fi36
                _ = v42
                _ = v43
            }
            v4 = v4
            _ = fi35
        }
    }
    _ = len(v5)
    v44 := GenS1{F0: v3, F1: 942}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
