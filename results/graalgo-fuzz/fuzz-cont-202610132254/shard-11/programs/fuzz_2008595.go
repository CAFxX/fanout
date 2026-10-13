package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint32; F2 bool }
func genFn0(p0 int, p1 int16, p2 string) string {
    v6 := GenS1{F0: 1, F1: -1, F2: 255}
    p1 += 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "hello"
}
func genFn1(p0 uint16, p1 uint) int16 {
    switch "\n\t" {
    case "foo bar":
        v7 := GenS1{F0: 9223372036854775807, F1: 1, F2: (^(p1))}
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int = fi8
            var v10 int32 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
    case "a":
        var v11 uint16 = 381
        _ = v11
    }
    if false {
        p0 = 728
        for fi12 := 0; fi12 < 5; fi12++ {
            switch "foo bar" {
            case "a":
                for fi13 := 0; fi13 < 8; fi13++ {
                    v14 := GenS1{F0: 0, F1: -115, F2: 746}
                    _ = fi13
                    _ = v14
                }
            default:
                var v15 int64 = -455
                _ = v15
            }
            p0 = (p0 & p0)
            for fi16 := 0; fi16 < 1; fi16++ {
                if false {
                    v17 := GenS1{F0: -9223372036854775808, F1: -18, F2: (p1 << ((p1) & 63))}
                    var v18 int = -1
                    _ = v17
                    _ = v18
                }
                p1 = (p1 * p1)
                _ = fi16
            }
            _ = fi12
        }
    }
    var v19 float64 = 3.14159
    v20 := GenS2{F0: "hello"}
    if false {
        v21 := GenS3{F0: 1, F1: 2, F2: ((1e10 == -2.25) || false)}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return -533
}

func main() {
    v1 := []int{42, 9223372036854775807}
    var v2 bool = (-2.25 == 0.0)
    v3 := make(map[string]string)
    var v4 uint8 = 98
    var v5 int32 = -1
    switch 141 {
    case -725:
        if len(v1) > 0 {
            v1[0] = 939
        }
    default:
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 int32 = v5
            _ = fi22
            _ = v23
        }
    }
    if false {
        for fi24 := 0; fi24 < 3; fi24++ {
            v1 = append(v1, (fi24 - (fi24 | (fi24 + -9223372036854775808))))
            v25 := GenS1{F0: -1, F1: (v5 * v5), F2: 18446744073709551615}
            _ = fi24
            _ = v25
        }
    } else {
        delete(v3, "a")
        if v2 {
            if (v2 && false) {
                if len(v1) > 0 {
                    v1[0] = 423
                }
                v26 := GenS2{F0: "foo bar"}
                v5 -= 0
                _ = v26
            }
            v2 = v2
            var v27 float32 = 0.0
            _ = v27
        }
    }
    var v28 int = 4294967295
    var v29 bool = v2
    if (!(42 >= 2)) {
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 bool = (v2 || (!((!true) || v29)))
            for fi32 := 0; fi32 < 16; fi32++ {
                if ((v29 || v31) || (0 == 42)) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = (4294967295 == 425)
                    switch v28 {
                    case 1:
                        v33 := GenS2{F0: "a"}
                        _ = v33
                    case 128:
                        v31 = v31
                    default:
                        v31 = v31
                    }
                }
                fmt.Println(v31, v28)
                _ = fi32
            }
            _ = fi30
            _ = v31
        }
        v4 = 23
        v34 := GenS1{F0: -1, F1: v5, F2: 18446744073709551615}
        _ = v34
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        v2 = (-843 <= 0)
        switch "hello" {
        case "\n\t":
            var v36 uint8 = v4
            _ = v36
        case "a":
            var v37 float64 = 0.1
            _ = v37
        case "":
            if false {
                var v38 uint64 = 2
                _ = v38
            }
        }
        v2 = (v4 != v4)
        _ = fi35
    }
    v28 += (v28 + (v28 / ((v28) | 1)))
    v2 = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
