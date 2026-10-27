package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint64 }
func genFn0(p0 int8) uint {
    if false {
        v4 := []int{}
        p0 += (p0 - 1)
        _ = v4
    } else {
        v5 := []int{9223372036854775807}
        if (0 >= 757) {
            p0 = (p0 - -1)
            var v6 uint16 = 65535
            if false {
                v7 := GenS1{F0: v6, F1: 32767}
                v6 = 258
                p0 += (p0 ^ (-(p0)))
                _ = v7
            } else {
                v5 = append(v5, 204)
                var v8 uint32 = 2147483647
                _ = v8
            }
            _ = v6
        }
        var v9 int32 = 668
        _ = v5
        _ = v9
    }
    if (!(((65535 > 2147483647) || (0.0 < 3.14159)) && (152 > 0))) {
        v10 := []int{32767, -9223372036854775808, 32767}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = v10
    }
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 bool = false
    v2 := []uint32{2, 2147483647, 0}
    var v3 int = -548
    v1 = (18446744073709551615 <= 671)
    fmt.Println(v1, v1)
    if v1 {
        v11 := GenS1{F0: 1, F1: -1}
        v3 += v3
        _ = v11
    } else {
        _ = len(v2)
        var v12 int8 = 1
        _ = v12
    }
    v3 = 2147483647
    switch (("foo bar" + "") + "\n\t") {
    case "\n\t":
        fmt.Println(v3)
    case "":
        v3 = v3
    case "a":
        var v13 bool = (9223372036854775807 < 9223372036854775807)
        _ = v13
    default:
        fmt.Println(v1)
    }
    v1 = ("hello" >= "a")
    var v14 uint8 = 255
    v2 = append(v2, 2147483647)
    v14 += (v14 & v14)
    var v15 int = -2147483648
    v1 = false
    v2 = append(v2, 42)
    v1 = v1
    for fi16 := 0; fi16 < 8; fi16++ {
        v15 += v3
        _ = fi16
    }
    if v1 {
        var v17 int8 = 127
        if false {
            if false {
                v3 -= v15
            } else {
                var v18 string = "\n\t"
                var v19 int16 = -97
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v1)
                case "a":
                    if (644 > 1) {
                        var v20 int = 765
                        if false {
                            var v21 uint8 = 1
                            v20 -= (v15 - (v20 / ((32768) | 1)))
                            _ = v21
                        } else {
                            var v22 uint64 = 2
                            _ = v22
                        }
                        _ = v20
                    }
                default:
                    if false {
                        if false {
                            v14 -= v14
                            for fi23 := 0; fi23 < 17; fi23++ {
                                v24 := GenS2{F0: 4294967295, F1: v14, F2: 4294967295}
                                var v25 int64 = -369
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS2{F0: 42, F1: (v14 + (v14 << ((863) & 63))), F2: 967}
                            _ = v26
                        }
                    }
                }
                _ = v18
                _ = v19
            }
        } else {
            v14 = (v14 / (((v14 >> ((791) & 63))) | 1))
        }
        v27 := GenS1{F0: 164, F1: -265}
        _ = v17
        _ = v27
    } else {
        fmt.Println(v15, v15, v3)
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 bool = ("\n\t" <= "hello")
        _ = fi28
        _ = v29
    }
    v3 = (v15 - (v3 & v15))
    var v30 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
}
