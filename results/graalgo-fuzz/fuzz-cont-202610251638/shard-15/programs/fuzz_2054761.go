package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int64; F1 int; F2 bool }
func genFn0(p0 uint8, p1 uint8, p2 uint32) uint8 {
    fmt.Println(p2, p2, p2)
    switch 983 {
    case 252:
        v4 := GenS1{F0: 42, F1: -605, F2: p2}
        _ = v4
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 * (^(p1)))
}

func main() {
    v1 := []int{0, 2147483647}
    var v2 uint64 = 1
    v3 := make(map[string]string)
    switch ("" + "\n\t") {
    case "foo bar":
        var v5 uint64 = (v2 & v2)
        _ = v5
    case "hello":
        var v6 uint32 = 2147483647
        _ = v6
    case "a":
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 int16 = 1000
                if (119 != 185) {
                    v9 += v9
                    if (-55 != 100) {
                        v10 := v3["hello"]
                        v1 = append(v1, fi8)
                        _ = v10
                    } else {
                        if true {
                            var v11 int32 = -1
                            v3[("foo bar" + "a")] = "\n\t"
                            v2 -= v2
                            _ = v11
                        }
                    }
                    v12 := GenS2{F0: 65535, F1: -847}
                    _ = v12
                }
                if (!(1 < -128)) {
                    v2 += 2
                    v9 = -349
                } else {
                    v2 -= (v2 / (((v2 - (v2 + (v2 ^ v2)))) | 1))
                    v13 := v3["\n\t"]
                    _ = v13
                }
                _ = fi8
                _ = v9
            }
            if (!(603 > -9223372036854775808)) {
                v2 += 2
                _ = len(v1)
                fmt.Println(v2)
            } else {
                var v14 int64 = -9223372036854775808
                _ = v14
            }
            _ = fi7
        }
    default:
        v2 = (v2 ^ v2)
    }
    fmt.Println(v2)
    if true {
        if (42 == -1) {
            _ = len(v3)
        } else {
            var v15 uint = 18446744073709551615
            var v16 bool = true
            _ = v15
            _ = v16
        }
        v2 = v2
    } else {
        v17 := v3["foo bar"]
        _ = v17
    }
    v18 := GenS3{F0: 9223372036854775807, F1: -53, F2: (!(!(false || true)))}
    _ = len(v3)
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int64 = -1
        if true {
            v1 = append(v1, (^(fi19)))
        } else {
            v20 += v20
        }
        switch ("hello" + "foo bar") {
        case "\n\t":
            if false {
                var v21 int = (fi19 % (((^(fi19))) | 1))
                v21 = v21
                if (!((1.5 - 0.0) > ((1.5 + ((float64(1e10) * 0.1) * float64(0.1))) + (-2.25 - -2.25)))) {
                    v2 = v2
                    if false {
                        v22 := GenS1{F0: 149, F1: fi19, F2: 2147483648}
                        if len(v1) > 0 {
                            v1[0] = (^(fi19))
                        }
                        v23 := GenS1{F0: -1, F1: (v21 % ((0) | 1)), F2: 1}
                        _ = v22
                        _ = v23
                    }
                }
                _ = v21
            }
        }
        _ = fi19
        _ = v20
    }
    if ((334 >= 1) || (1 <= 414)) {
        v2 = 764
    } else {
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 int64 = 603
            _ = fi24
            _ = v25
        }
    }
    var v26 int16 = 42
    for fi27 := 0; fi27 < 12; fi27++ {
        if ((!(42 == 860)) || true) {
            var v28 int = fi27
            _ = v28
        }
        var v29 bool = (!(true && false))
        _ = fi27
        _ = v29
    }
    var v30 uint64 = (v2 & (v2 | (v2 ^ (v2 * (v2 % ((2) | 1))))))
    switch "\n\t" {
    case "\n\t":
        v26 = (v26 % ((v26) | 1))
    case "a":
        var v31 int64 = -522
        _ = v31
    default:
        var v32 int16 = v26
        _ = v32
    }
    v30 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v30
}
