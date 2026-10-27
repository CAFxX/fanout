package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 int8 }
func genFn0(p0 uint8, p1 uint) int {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    v9 := []int{-1, 42, 704}
    if true {
        v10 := GenS1{F0: 18446744073709551615, F1: (^(p0)), F2: 1}
        var v11 int64 = 1
        _ = v10
        _ = v11
    } else {
        var v12 int32 = 596
        fmt.Println(v12, p1)
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        p1 = p1
        fmt.Println(p0, p1, p0)
        switch (fi13 << ((p1) & 63)) {
        case -451:
            if (!(!(-44 < 1))) {
                v14 := GenS1{F0: 18446744073709551615, F1: 128, F2: 596}
                v15 := GenS1{F0: 905, F1: p0, F2: 0}
                if true {
                    if (int32(p0) <= -2147483648) {
                        fmt.Println(p1, p1)
                    } else {
                        for fi16 := 0; fi16 < 12; fi16++ {
                            fmt.Println(p1)
                            v17 := GenS1{F0: 2, F1: p0, F2: -1}
                            var v18 int = fi16
                            fmt.Println(v18, v18, fi16)
                            if (!true) {
                                switch (v18 % ((v18) | 1)) {
                                case -2147483648:
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                case 4294967295:
                                    var v19 int32 = 220
                                    _ = v19
                                case -9223372036854775808:
                                    if (1.5 < 3.14159) {
                                        v20 := GenS1{F0: 4294967295, F1: p0, F2: -719}
                                        if true {
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            v21 := GenS1{F0: 1, F1: p0, F2: 0}
                                            _ = v21
                                        } else {
                                            if (!true) {
                                                var v22 float32 = 0.1
                                                _ = v22
                                            }
                                            switch fi16 {
                                            case -2147483648:
                                                v9 = append(v9, 32767)
                                            }
                                            _ = len(v9)
                                        }
                                        _ = v20
                                    }
                                }
                                if (0 == 1) {
                                    v23 := GenS1{F0: 0, F1: (p0 / ((p0) | 1)), F2: 247}
                                    _ = v23
                                } else {
                                    var v24 uint = (p1 >> ((602) & 63))
                                    v9 = append(v9, (fi16 << (((v24 & 2)) & 63)))
                                    p0 -= (p0 + (p0 << (((p1 | (p1 * 1))) & 63)))
                                    _ = v24
                                }
                                var v25 int64 = -9223372036854775808
                                _ = v25
                            } else {
                                var v26 uint = 42
                                switch v18 {
                                case -550:
                                    var v27 int32 = -602
                                    _ = v27
                                case 128:
                                    fmt.Println(p1, fi16, fi13)
                                default:
                                    fmt.Println(p0, p1, v18)
                                }
                                _ = v26
                            }
                            if (!false) {
                                var v28 int16 = 1
                                p0 += 0
                                _ = v28
                            }
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                    var v29 uint16 = 864
                    _ = v29
                }
                _ = v14
                _ = v15
            }
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v9
    return -489
}
func genFn1() string {
    if false {
        v30 := []int{9223372036854775807, 1, 32767}
        var v31 int8 = 0
        _ = v30
        _ = v31
    } else {
        var v32 uint64 = 4294967295
        _ = v32
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        fmt.Println(fi33)
        var v34 uint = 373
        v34 -= 47
        _ = fi33
        _ = v34
    }
    return (("" + ("hello" + ("" + ("\n\t" + "hello")))) + ((((("\n\t" + "") + "foo bar") + "foo bar") + "a") + "a"))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    v3 := []string{"", ("foo bar" + "hello"), ""}
    v4 := []uint32{2147483647, 2147483648}
    var v5 int16 = -1
    var v6 bool = true
    v2[0] = 32768
    for fi35 := 0; fi35 < 18; fi35++ {
        var v36 uint = 4294967295
        v6 = (false || false)
        _ = fi35
        _ = v36
    }
    v6 = v6
    v37 := GenS2{F0: 1, F1: 42, F2: -49}
    if (v6 && (false || v6)) {
        var v38 float64 = 1.5
        _ = v38
    }
    v39 := GenS1{F0: 467, F1: 229, F2: v5}
    v6 = v6
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
