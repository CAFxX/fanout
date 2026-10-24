package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint8 }

func main() {
    var v1 int32 = 0
    v2 := []int{32768, -38, -193, -418}
    var v3 uint16 = 42
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: 18446744073709551615}
        if len(v2) > 0 {
            v2[0] = fi4
        }
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            switch ("" + "a") {
            case "\n\t":
                v2 = append(v2, 32768)
            case "a":
                v3 = (v3 + 2)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (0 >= 18446744073709551615) {
                    switch "hello" {
                    case "hello":
                        v1 -= (v1 * (v1 | 265))
                    case "a":
                        var v8 float64 = ((1e10 - (1e10 - 0.1)) / ((3.14159 / ((-2.25 - 3.14159) / 0.1)) + (1e10 * 0.0)))
                        if (18446744073709551615 >= 65535) {
                            v3 -= 1
                            var v9 uint64 = 0
                            if (411 < fi7) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v2 = append(v2, -2147483648)
                                var v10 int8 = 4
                                _ = v10
                            }
                            _ = v9
                        }
                        _ = v8
                    case "\n\t":
                        if (fi7 <= fi6) {
                            var v11 int16 = 42
                            v1 = (v1 >> ((4294967295) & 63))
                            v1 += -1
                            switch (fi7 * (fi7 ^ -9223372036854775808)) {
                            case 4294967295:
                                v3 = 32767
                            }
                            _ = v11
                        } else {
                            var v12 float32 = (((-2.25 + 0.0) - (float32((1e10 - (-2.25 * (1.5 + 0.0)))) * 3.14159)) + 3.14159)
                            v1 = v1
                            if (!(2 >= 2)) {
                                if false {
                                    var v13 uint32 = 178
                                    _ = v13
                                } else {
                                    v3 = (v3 % ((v3) | 1))
                                    v1 -= (v1 << ((759) & 63))
                                    if true {
                                        fmt.Println(fi7)
                                        var v14 int16 = 1
                                        _ = v14
                                    }
                                }
                                v3 -= v3
                                var v15 string = ("" + "a")
                                _ = v15
                            }
                            v1 = (v1 + (-(v1)))
                            v2 = append(v2, (fi6 * 32767))
                            _ = v12
                        }
                    default:
                        v1 = v1
                    }
                    switch fi6 {
                    case -139:
                        var v16 uint = 282
                        _ = v16
                    case 0:
                        fmt.Println(v3, fi7)
                    case 9223372036854775807:
                        _ = len(v2)
                    default:
                        if (true || (42 <= 562)) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v3 = v3
                            var v17 int16 = 1
                            v3 = (v3 % ((v3) | 1))
                            _ = v17
                        }
                    }
                    switch fi7 {
                    case -9223372036854775808:
                        fmt.Println(v3, v3, fi6)
                    default:
                        if (18446744073709551615 < 0) {
                            v1 = (^(v1))
                        }
                    }
                    if true {
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            var v18 uint32 = 65535
                            _ = v18
                        }
                        var v19 int8 = 1
                        var v20 uint8 = 2
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(fi7)
                        v21 := GenS1{F0: 4294967295}
                        _ = v21
                    }
                }
            default:
                v3 = (v3 & (v3 >> ((4294967295) & 63)))
            }
            _ = fi7
        }
        fmt.Println(fi6, v1)
        _ = fi6
    }
    v2 = append(v2, 240)
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int = fi22
        _ = fi22
        _ = v23
    }
    v3 += 65535
    v1 -= 2147483647
    v2 = append(v2, -110)
    v24 := GenS1{F0: 2}
    switch 127 {
    case 850:
        fmt.Println(v1, v1)
    case 2147483647:
        v25 := GenS1{F0: 1}
        _ = v25
    }
    v26 := GenS1{F0: 42}
    v27 := GenS2{F0: 0, F1: 42, F2: 127}
    var v28 string = ""
    v1 -= v1
    var v29 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
