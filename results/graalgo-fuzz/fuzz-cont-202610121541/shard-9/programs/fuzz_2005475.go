package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 string; F2 uint32 }
func genFn0(p0 uint, p1 uint) int {
    fmt.Println(p0, p0)
    var v5 uint16 = 255
    _ = p0
    _ = p1
    _ = v5
    return 598
}
func genFn1(p0 uint) uint8 {
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: (p0 << (((p0 / (((p0 | 863)) | 1))) & 63)), F1: v6, F2: 874}
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 bool = false
            _ = fi9
            _ = v10
        }
        v11 := []int{genFn0((p0 / ((199) | 1)), p0), genFn0(p0, p0), fi7}
        _ = fi7
        _ = v11
        _ = v8
    }
    var v12 float32 = ((float32(0.1) * ((float32(0.1) / ((-2.25 + 3.14159) / (0.0 - (1e10 + 1e10)))) - 0.1)) / (-2.25 / 1.5))
    switch 32768 {
    case 9223372036854775807:
        v12 = ((float32((3.14159 + float64(v12))) + (1e10 / 3.14159)) / (((-2.25 + ((1e10 * v12) * ((((v12 - v12) + 0.0) * (v12 - v12)) - float32((1.5 * (3.14159 - 3.14159)))))) + v12) * v12))
    default:
        for fi13 := 0; fi13 < 10; fi13++ {
            v6 -= uint32(p0)
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := []int{(-(fi13)), fi14, (fi13 + 834)}
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    }
    var v16 uint16 = 42
    _ = p0
    _ = v12
    _ = v16
    _ = v6
    return 56
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    var v4 int8 = 100
    var v17 uint = 1
    v4 -= 1
    v17 += 2
    if true {
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 uint8 = 128
            for fi20 := 0; fi20 < 5; fi20++ {
                switch "\n\t" {
                case "foo bar":
                    v19 += (v19 | v19)
                case "a":
                    if false {
                        switch (("" + ("" + "a")) + ("foo bar" + "foo bar")) {
                        case "a":
                            var v21 float64 = float64(-2.25)
                            _ = v21
                        case "":
                            v22 := GenS1{F0: (v17 & (v17 >> (((^(v17))) & 63))), F1: 42, F2: -1}
                            if true {
                                if (42 < 243) {
                                    v23 := []int{int(v19)}
                                    _ = v23
                                } else {
                                    v24 := []int{fi20, fi20, 42}
                                    _ = v24
                                }
                            } else {
                                v25 := v2[529]
                                _ = v25
                            }
                            _ = v22
                        }
                    } else {
                        v26 := []int{genFn0((v17 ^ 4294967295), v17)}
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        _ = v26
                    }
                case "":
                    _ = len(v2)
                default:
                    if (fi18 == genFn0(360, (v17 + v17))) {
                        if (0 >= int64(v17)) {
                            v19 -= uint8(v17)
                            v4 += v4
                            var v27 uint64 = 4294967295
                            _ = v27
                        }
                        var v28 uint8 = v19
                        v17 += 255
                        _ = v28
                    } else {
                        switch "" {
                        case "a":
                            fmt.Println(v4)
                        case "\n\t":
                            var v29 int8 = v4
                            _ = v29
                        }
                        var v30 int = (fi20 + -9223372036854775808)
                        _ = v30
                    }
                }
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
        v3[-9223372036854775808] = 9223372036854775807
        v4 -= 0
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        if false {
            var v32 int32 = -114
            v4 -= (v4 & (v4 % (((-(v4))) | 1)))
            for fi33 := 0; fi33 < 13; fi33++ {
                switch "\n\t" {
                case "hello":
                    v17 += (v17 / ((18446744073709551615) | 1))
                default:
                    var v34 float32 = (1.5 * 1e10)
                    _ = v34
                }
                v35 := GenS2{F0: 0}
                _ = fi33
                _ = v35
            }
            _ = v32
        } else {
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 float32 = (-2.25 * 1.5)
                v17 = 694
                _ = fi36
                _ = v37
            }
            for fi38 := 0; fi38 < 10; fi38++ {
                var v39 int64 = -62
                if true {
                    v40 := GenS1{F0: v17, F1: 478, F2: -9223372036854775808}
                    _ = v40
                }
                _ = fi38
                _ = v39
            }
        }
        _ = fi31
    }
    v41 := GenS1{F0: 456, F1: 2147483648, F2: -1}
    v42 := GenS2{F0: 42}
    var v43 uint64 = 2
    delete(v1, "")
    if false {
        if (2 == 42) {
            if true {
                v43 -= v43
            }
            v44 := GenS3{F0: -450, F1: ("foo bar" + "a"), F2: 2147483648}
            _ = v44
        }
    } else {
        v17 -= v17
        for fi45 := 0; fi45 < 9; fi45++ {
            v43 = v43
            var v46 string = ("a" + ("\n\t" + "hello"))
            _ = fi45
            _ = v46
        }
        fmt.Println(v17, v4, v4)
    }
    v47 := []int{}
    var v48 uint = v17
    var v49 int32 = -2147483648
    v48 = v48
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v47
    _ = v48
    _ = v49
}
