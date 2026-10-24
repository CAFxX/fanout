package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 int }

func main() {
    var v1 float32 = (1e10 + 1e10)
    v2 := []int{}
    v3 := make(map[string]uint32)
    var v4 int = -9223372036854775808
    var v5 string = "\n\t"
    var v6 string = v5
    v1 = v1
    switch v4 {
    case -2147483648:
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(v4)
            fmt.Println(fi7, v5)
            _ = fi7
        }
    case 892:
        if false {
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 int8 = -100
                v4 = (fi8 + 16)
                if false {
                    if (-1 == 1) {
                        if true {
                            var v10 uint16 = 255
                            v6 = (v5 + ("foo bar" + ""))
                            v11 := v3[v5]
                            _ = v10
                            _ = v11
                        }
                        v12 := GenS1{F0: false, F1: 352, F2: 43}
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 14; fi13++ {
                            v2 = append(v2, fi8)
                            _ = fi13
                        }
                        if len(v2) > 0 {
                            v2[0] = -2147483648
                        }
                        v3[v5] = 2147483648
                    }
                    if (!false) {
                        if true {
                            v14 := GenS2{F0: v4}
                            if ((!(!true)) || true) {
                                _ = len(v2)
                            }
                            _ = v14
                        }
                        v9 = v9
                        if (((42 == -9223372036854775808) || (false || (65535 == 2))) && false) {
                            var v15 int16 = -32768
                            var v16 uint64 = 632
                            _ = v15
                            _ = v16
                        } else {
                            if len(v2) > 0 {
                                v2[0] = (v4 * v4)
                            }
                            v1 = v1
                            if ((!(42 != v9)) && false) {
                                var v17 int16 = -659
                                if false {
                                    fmt.Println(v4)
                                }
                                fmt.Println(v9)
                                _ = v17
                            } else {
                                v2 = append(v2, -655)
                                v18 := GenS2{F0: fi8}
                                _ = len(v2)
                                _ = v18
                            }
                        }
                    }
                } else {
                    v19 := GenS1{F0: ((fi8 << ((4294967295) & 63)) != v4), F1: 255, F2: 1}
                    v20 := GenS2{F0: (v4 >> ((319) & 63))}
                    _ = v19
                    _ = v20
                }
                _ = fi8
                _ = v9
            }
            fmt.Println(v1)
        }
    }
    if false {
        var v21 float32 = 1.5
        v2 = append(v2, v4)
        _ = v21
    }
    fmt.Println(v1)
    for fi22 := 0; fi22 < 19; fi22++ {
        switch v5 {
        case "hello":
            v1 = v1
        case "a":
            v5 = v5
        }
        var v23 int8 = 127
        _ = fi22
        _ = v23
    }
    v1 += (v1 * v1)
    delete(v3, ("\n\t" + v5))
    _ = len(v3)
    v4 -= (v4 % (((v4 | v4)) | 1))
    v1 = 0.1
    if (v5 == (v6 + ("\n\t" + "hello"))) {
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 int8 = 0
            for fi26 := 0; fi26 < 5; fi26++ {
                v25 += v25
                delete(v3, ("foo bar" + "hello"))
                _ = fi26
            }
            _ = fi24
            _ = v25
        }
        var v27 uint16 = 0
        _ = v27
    } else {
        var v28 uint = 559
        _ = v28
    }
    var v29 int32 = 42
    v29 -= (v29 - v29)
    switch (v4 ^ v4) {
    case 32768:
        var v30 uint32 = 913
        _ = v30
    default:
        v5 = (v5 + v6)
    }
    switch "hello" {
    case "":
        v29 += (-(v29))
    default:
        v31 := GenS2{F0: (-(v4))}
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
