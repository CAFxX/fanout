package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int }
func genFn0() int64 {
    v4 := []int{2147483647, 128}
    var v5 uint16 = 2
    var v6 int64 = 893
    v6 += v6
    v5 += (v5 & v5)
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 float32 = (0.0 - -2.25)
    var v2 float32 = 0.1
    var v3 uint16 = 42
    var v7 uint32 = 0
    v8 := []int{128, -9223372036854775808, -1}
    var v9 int32 = 0
    v1 = (-2.25 * (1.5 / v2))
    switch 0 {
    case 127:
        var v10 int64 = genFn0()
        _ = v10
    case -739:
        v7 -= v7
    case 341:
        if true {
            if (((-2.25 * (-2.25 * 0.1)) + ((1e10 - ((3.14159 * 0.1) - 1.5)) + 0.0)) < 1.5) {
                v7 -= (v7 | v7)
            }
            _ = len(v8)
        }
    default:
        var v11 int64 = genFn0()
        _ = v11
    }
    var v12 uint16 = (v3 - (v3 % ((v3) | 1)))
    if len(v8) > 0 {
        v8[0] = -102
    }
    v13 := GenS1{F0: 113, F1: (931 != 42), F2: 42}
    if ((2 != 18446744073709551615) && (false && (255 == (v3 ^ v12)))) {
        v3 -= v3
    }
    fmt.Println(v3, v9)
    var v14 int = 42
    var v15 int64 = 31
    for fi16 := 0; fi16 < 12; fi16++ {
        if (!false) {
            switch ("" + ("\n\t" + "")) {
            case "foo bar":
                fmt.Println(v7, v2)
            case "\n\t":
                v17 := GenS1{F0: -1, F1: false, F2: 684}
                _ = v17
            case "hello":
                if false {
                    v18 := GenS1{F0: 1, F1: false, F2: 354}
                    switch (^(v14)) {
                    case -3:
                        var v19 int = (fi16 * (fi16 | (fi16 ^ int(v9))))
                        _ = v19
                    case 127:
                        fmt.Println(v3, v14, v1)
                    case 32767:
                        switch -497 {
                        case 384:
                            v15 += (v15 ^ (v15 & v15))
                        case -2147483648:
                            v2 = (v2 / (0.1 / (v2 / 0.1)))
                        }
                    default:
                        if ((^(v15)) <= 42) {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            v12 = v3
                            v8 = append(v8, (^(fi16)))
                        }
                    }
                    _ = v18
                }
            }
            v8 = append(v8, (fi16 >> ((572) & 63)))
        }
        v20 := GenS2{F0: v12, F1: v3, F2: fi16}
        v1 += ((v1 - v2) / 1.5)
        _ = fi16
        _ = v20
    }
    var v21 uint16 = (v3 | (v3 >> ((1) & 63)))
    if (!(4294967295 > 4294967295)) {
        if len(v8) > 0 {
            v8[0] = v14
        }
        v7 = (v7 - 0)
    } else {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 float32 = (v2 - float32(((1.5 / 1.0) * float64(float32(-2.25)))))
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            if false {
                var v25 bool = true
                for fi26 := 0; fi26 < 6; fi26++ {
                    if (!true) {
                        var v27 int = v14
                        _ = v27
                    }
                    _ = fi26
                }
                _ = v25
            }
            if false {
                if false {
                    v3 += v12
                    v3 = (v12 + 860)
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                v28 := GenS1{F0: -1, F1: (uint64(v12) > 2), F2: 1}
                _ = v28
            }
            _ = fi24
        }
    }
    var v29 uint = 789
    var v30 uint32 = 2147483648
    v21 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v7
    _ = v8
    _ = v9
}
