package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int8; F1 string; F2 int8 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int; F1 uint64 }

func main() {
    var v1 int8 = 1
    var v2 int64 = 878
    v3 := make(map[string]int)
    v4 := []string{""}
    v5 := []int{2147483647, -9223372036854775808}
    var v6 uint8 = 0
    for fi7 := 0; fi7 < 12; fi7++ {
        switch ("hello" + ((("a" + "hello") + "") + "foo bar")) {
        case "foo bar":
            var v8 int8 = (v1 ^ v1)
            _ = v8
        case "":
            if (656 <= 4294967295) {
                var v9 float32 = -2.25
                v10 := GenS1{F0: 1, F1: v2}
                v11 := GenS2{F0: (v6 << ((255) & 63)), F1: 62}
                _ = v10
                _ = v11
                _ = v9
            }
        case "hello":
            v12 := GenS3{F0: v2, F1: -2147483648}
            _ = v12
        default:
            if ((!(18446744073709551615 == 195)) && true) {
                switch (fi7 / (((^(fi7))) | 1)) {
                case -942:
                    v13 := GenS2{F0: (^(v6)), F1: (fi7 ^ (fi7 & fi7))}
                    _ = v13
                case -1:
                    if ((true && (127 == 1)) || true) {
                        v1 -= (v1 * (v1 / ((v1) | 1)))
                        v14 := GenS4{F0: v1, F1: "\n\t", F2: (^(v1))}
                        var v15 uint32 = 0
                        _ = v14
                        _ = v15
                    }
                default:
                    v2 = v2
                }
                switch 0 {
                case 4294967295:
                    var v16 int32 = -84
                    _ = v16
                default:
                    if (true && (970 < 0)) {
                        v1 = v1
                        v1 -= v1
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
            } else {
                fmt.Println(v2, v6, fi7)
                _ = len(v4)
                v6 -= v6
            }
        }
        var v17 bool = (!(v6 > v6))
        _ = fi7
        _ = v17
    }
    v18 := GenS5{F0: false, F1: 591}
    switch "\n\t" {
    case "hello":
        v19 := v3["\n\t"]
        _ = v19
    case "foo bar":
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := GenS4{F0: v1, F1: (((("" + "hello") + "\n\t") + "a") + "\n\t"), F2: (^(v1))}
            fmt.Println(v6)
            _ = fi20
            _ = v21
        }
    default:
        var v22 uint16 = 691
        _ = v22
    }
    if ((v6 ^ v6) > v6) {
        var v23 bool = false
        if (42 == 2147483647) {
            _ = len(v3)
        }
        _ = v23
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v3)
    }
    v2 = 0
    v6 -= (v6 * (v6 >> ((65535) & 63)))
    switch "hello" {
    case "a":
        v24 := GenS6{F0: 0, F1: 61}
        _ = v24
    case "hello":
        var v25 int32 = 0
        _ = v25
    case "":
        for fi26 := 0; fi26 < 20; fi26++ {
            v27 := GenS2{F0: (^(v6)), F1: fi26}
            _ = fi26
            _ = v27
        }
    }
    if (378 > 280) {
        var v28 int8 = (v1 & v1)
        for fi29 := 0; fi29 < 3; fi29++ {
            v30 := v3[(("" + "hello") + "hello")]
            v28 = 127
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    delete(v3, "foo bar")
    var v31 uint32 = 2147483647
    fmt.Println(v2)
    var v32 float32 = 3.14159
    var v33 float64 = 3.14159
    var v34 float32 = v32
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
