package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 string }
func genFn0(p0 uint) int32 {
    if (4294967295 <= 32768) {
        if false {
            for fi4 := 0; fi4 < 2; fi4++ {
                v5 := GenS1{F0: 65535, F1: 9223372036854775807, F2: (((("foo bar" + ("\n\t" + "foo bar")) + ("a" + "foo bar")) <= ("" + "")) || false)}
                _ = fi4
                _ = v5
            }
            v6 := GenS1{F0: 40, F1: 1, F2: (!false)}
            for fi7 := 0; fi7 < 7; fi7++ {
                for fi8 := 0; fi8 < 15; fi8++ {
                    if (65535 <= p0) {
                        var v9 bool = (true || (p0 < (p0 / (((p0 >> (((p0 << (((p0 | 84)) & 63))) & 63))) | 1))))
                        var v10 int = 435
                        _ = v10
                        _ = v9
                    } else {
                        var v11 int16 = 0
                        _ = v11
                    }
                    switch 9223372036854775807 {
                    case 4294967295:
                        fmt.Println(fi8)
                    case 32767:
                        v12 := GenS2{F0: -1, F1: 1}
                        _ = v12
                    case -662:
                        p0 = p0
                    }
                    _ = fi8
                }
                var v13 uint16 = 363
                v13 = v13
                _ = fi7
                _ = v13
            }
            _ = v6
        }
        fmt.Println(p0)
    } else {
        var v14 uint64 = 904
        fmt.Println(v14)
        var v15 uint16 = 65535
        _ = v14
        _ = v15
    }
    p0 += (p0 * p0)
    p0 = p0
    _ = p0
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 2147483647
    var v3 uint16 = 42
    v16 := GenS2{F0: 127, F1: 976}
    var v17 int32 = v2
    v3 = (v3 + v3)
    for fi18 := 0; fi18 < 16; fi18++ {
        v3 += 255
        v2 = (v17 % (((-(v17))) | 1))
        _ = fi18
    }
    _ = len(v1)
    var v19 bool = (((65535 != 795) && (false && (1 >= 42))) || (!(("a" + "\n\t") < ("" + (("\n\t" + ("a" + "a")) + "foo bar")))))
    for fi20 := 0; fi20 < 18; fi20++ {
        _ = len(v1)
        if v19 {
            if (0 >= 2) {
                var v21 bool = (true || v19)
                _ = v21
            }
            fmt.Println(v2, v3, v3)
        }
        _ = fi20
    }
    fmt.Println(v2)
    v17 = v2
    v22 := []int{4294967295, 42, -2147483648, 1}
    if false {
        if (775 < 1) {
            v17 += 174
            if v19 {
                if v19 {
                    for fi23 := 0; fi23 < 1; fi23++ {
                        if v19 {
                            for fi24 := 0; fi24 < 17; fi24++ {
                                var v25 int16 = -32768
                                _ = fi24
                                _ = v25
                            }
                        } else {
                            v2 = v17
                            v19 = (793 >= 32767)
                            v26 := GenS3{F0: 42, F1: 42}
                            _ = v26
                        }
                        v19 = (v19 && v19)
                        v22 = append(v22, fi23)
                        _ = fi23
                    }
                    delete(v1, "\n\t")
                } else {
                    v3 += v3
                    v17 += v2
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        var v27 float64 = ((0.0 / float64(((1e10 + 3.14159) * 3.14159))) * (1.5 - 1.5))
                        _ = v27
                    }
                }
                var v28 int32 = genFn0(4294967295)
                v29 := GenS4{F0: "\n\t"}
                _ = v28
                _ = v29
            } else {
                var v30 uint64 = 42
                var v31 uint32 = 65535
                _ = v30
                _ = v31
            }
            v3 += (v3 >> ((255) & 63))
        }
        v32 := GenS4{F0: ""}
        _ = v32
    } else {
        for fi33 := 0; fi33 < 14; fi33++ {
            fmt.Println(v2, v19)
            _ = fi33
        }
    }
    var v34 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v34
}
