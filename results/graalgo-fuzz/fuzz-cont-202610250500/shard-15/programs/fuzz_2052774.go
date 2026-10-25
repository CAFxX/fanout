package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int16 }
func genFn0(p0 int32, p1 int32) uint64 {
    var v6 int16 = -32768
    v7 := []int{0, 32768, 9223372036854775807, 4294967295}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -1
    var v3 int64 = -1
    var v4 uint64 = 18446744073709551615
    v5 := []string{"\n\t", "\n\t", "a", ("\n\t" + (("\n\t" + "foo bar") + ("" + ("a" + ""))))}
    v8 := v1[572]
    switch (v8 + v8) {
    case "\n\t":
        var v9 bool = true
        _ = v9
    }
    var v10 int32 = -1
    v11 := GenS1{F0: 42, F1: -99, F2: -1}
    if (true && false) {
        if (!(42 != 1)) {
            var v12 uint = 2
            for fi13 := 0; fi13 < 9; fi13++ {
                if true {
                    v4 -= 45
                }
                _ = fi13
            }
            var v14 float64 = 1.5
            _ = v12
            _ = v14
        } else {
            fmt.Println(v2, v10)
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            if (((float64(0.0) - 1.5) + float64(0.1)) > (3.14159 / 1.5)) {
                var v16 int = fi15
                v17 := GenS2{F0: 0, F1: 0}
                for fi18 := 0; fi18 < 12; fi18++ {
                    v8 = v8
                    switch fi18 {
                    case 32767:
                        _ = len(v5)
                    case 4294967295:
                        var v19 uint32 = 65535
                        _ = v19
                    default:
                        var v20 uint64 = (v4 << ((18446744073709551615) & 63))
                        _ = v20
                    }
                    fmt.Println(v3, v3)
                    _ = fi18
                }
                _ = v16
                _ = v17
            }
            if false {
                if true {
                    var v21 float64 = float64(1e10)
                    for fi22 := 0; fi22 < 11; fi22++ {
                        v1[32] = ("foo bar" + v8)
                        fmt.Println(fi15, v8)
                        fmt.Println(v21, v3)
                        _ = fi22
                    }
                    _ = v21
                } else {
                    v2 = v2
                    var v23 int16 = 827
                    for fi24 := 0; fi24 < 11; fi24++ {
                        v25 := GenS3{F0: v4, F1: (true && (127 <= 2))}
                        _ = fi24
                        _ = v25
                    }
                    _ = v23
                }
                var v26 float64 = -2.25
                _ = v26
            } else {
                v5 = append(v5, v8)
                v2 -= v2
                var v27 int16 = v2
                _ = v27
            }
            v8 = (v8 + v8)
            _ = fi15
        }
        if ((v8 + v8) != ((v8 + "hello") + v8)) {
            v28 := GenS3{F0: (^(v4)), F1: false}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 int = 32768
        v30 = -976
        _ = len(v1)
        _ = fi29
        _ = v30
    }
    var v31 int8 = 127
    v8 = (((v8 + v8) + v8) + (v8 + (v8 + (v8 + "a"))))
    delete(v1, 82)
    v10 = v10
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    var v32 int8 = v31
    for fi33 := 0; fi33 < 19; fi33++ {
        if true {
            v34 := GenS2{F0: (v3 | v3), F1: (fi33 ^ (fi33 >> ((18446744073709551615) & 63)))}
            v3 += (v3 >> ((2) & 63))
            fmt.Println(v31, v31)
            _ = v34
        }
        if true {
            if false {
                _ = len(v1)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v32 -= v31
            } else {
                switch (fi33 | fi33) {
                case -176:
                    var v35 int32 = v10
                    _ = v35
                case 586:
                    for fi36 := 0; fi36 < 17; fi36++ {
                        var v37 bool = (!false)
                        _ = fi36
                        _ = v37
                    }
                default:
                    fmt.Println(v2, v4)
                }
                v3 -= (v3 & v3)
            }
            if ((1 != 479) || ((v10 ^ (^(v10))) != (v10 + v10))) {
                fmt.Println(v31, v3)
            }
        } else {
            v2 -= 32767
            v4 -= (v4 >> ((42) & 63))
            v4 -= 4294967295
        }
        _ = fi33
    }
    if false {
        fmt.Println(v8, v2, v8)
        if false {
            fmt.Println(v10, v3)
            var v38 float64 = ((3.14159 / (((0.0 + (3.14159 / 1.0)) - ((0.1 / 3.14159) / 1.5)) / (((1.5 + ((-2.25 * 0.0) / 1.5)) / 1.5) / 1e10))) * float64((0.0 + (3.14159 * ((-2.25 * 0.1) - 0.0)))))
            v2 -= v2
            _ = v38
        } else {
            var v39 int64 = (^(v3))
            if (!(!(32767 != 1))) {
                switch v8 {
                case "a":
                    v32 = v31
                default:
                    fmt.Println(v8, v3, v31)
                }
                for fi40 := 0; fi40 < 8; fi40++ {
                    v41 := GenS4{F0: 1}
                    v42 := GenS3{F0: 866, F1: false}
                    var v43 int16 = -1
                    _ = fi40
                    _ = v41
                    _ = v42
                    _ = v43
                }
            } else {
                v44 := v1[4294967295]
                if true {
                    switch -112 {
                    case -1:
                        _ = len(v5)
                    case 32767:
                        fmt.Println(v44)
                    default:
                        for fi45 := 0; fi45 < 2; fi45++ {
                            for fi46 := 0; fi46 < 17; fi46++ {
                                var v47 int = (fi46 / (((^(fi46))) | 1))
                                _ = fi46
                                _ = v47
                            }
                            if (v2 > 42) {
                                v2 = 1
                            }
                            v48 := v1[803]
                            _ = fi45
                            _ = v48
                        }
                    }
                } else {
                    v4 = 0
                    fmt.Println(v2, v31, v4)
                }
                v5 = append(v5, "a")
                _ = v44
            }
            delete(v1, 762)
            _ = v39
        }
    }
    v49 := GenS5{F0: (^(v4)), F1: 631, F2: v2}
    fmt.Println(v10)
    fmt.Println(v2, v32)
    for fi50 := 0; fi50 < 12; fi50++ {
        v3 = v3
        var v51 uint64 = v4
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi50
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v49
    _ = v5
    _ = v8
}
