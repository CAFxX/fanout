package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int }
type GenS2 struct { F0 uint; F1 int; F2 int }
func genFn0(p0 bool, p1 int64, p2 bool) uint32 {
    v6 := GenS1{F0: 1, F1: (p1 + -1), F2: -1}
    v7 := GenS1{F0: 128, F1: 42, F2: -452}
    for fi8 := 0; fi8 < 7; fi8++ {
        p0 = (1e10 >= float64(3.14159))
        var v9 int64 = (p1 | -9223372036854775808)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 255, F1: -9223372036854775808, F2: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    return 557
}
func genFn1(p0 bool) uint {
    var v11 int64 = 711
    v12 := []int{-9223372036854775808, 127, -347}
    var v13 int = 4294967295
    switch ("hello" + "hello") {
    case "hello":
        fmt.Println(p0)
    default:
        p0 = true
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    v1 := []string{"foo bar", "a", "hello"}
    var v2 int32 = 1
    v3 := []uint32{144, 42, 2147483647, 0}
    v4 := []string{"foo bar", "\n\t", ""}
    v5 := []int{4294967295}
    if true {
        if (false || false) {
            if true {
                switch 1 {
                case 1:
                    v2 = (v2 % ((-236) | 1))
                default:
                    var v14 uint8 = 255
                    _ = v14
                }
            }
            var v15 float64 = 1.5
            fmt.Println(v2)
            _ = v15
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                switch (fi16 | (^(fi16))) {
                case 0:
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                default:
                    v2 = v2
                }
                var v17 string = (("a" + ("hello" + (("\n\t" + "") + "foo bar"))) + "hello")
                fmt.Println(fi16, fi16)
                _ = fi16
                _ = v17
            }
        }
        v18 := GenS1{F0: 127, F1: 1, F2: 128}
        _ = v18
    }
    if (((187 == 27) && false) && (42 != 0)) {
        var v19 int16 = 1000
        if false {
            v19 += v19
            if false {
                if len(v5) > 0 {
                    v5[0] = 127
                }
                _ = len(v3)
            } else {
                var v20 uint = 0
                if true {
                    v20 = (v20 ^ 384)
                    if (true && (((1 <= 1) || false) || (uint16(v20) == 2))) {
                        v2 -= 42
                    } else {
                        fmt.Println(v19)
                    }
                    var v21 int = -75
                    _ = v21
                }
                _ = v20
            }
        } else {
            if len(v4) > 0 {
                v4[0] = ("a" + "\n\t")
            }
            v19 -= (v19 ^ v19)
        }
        _ = v19
    }
    if (-840 < -1) {
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := GenS1{F0: 111, F1: -1, F2: (fi22 >> ((42) & 63))}
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := GenS2{F0: 18446744073709551615, F1: 0, F2: 998}
                var v26 float64 = -2.25
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi22
            _ = v23
        }
    } else {
        v2 = (v2 >> ((1) & 63))
        if false {
            var v27 float32 = (1.5 - 1.5)
            var v28 int = 331
            _ = v27
            _ = v28
        }
        fmt.Println(v2)
    }
    var v29 float64 = (3.14159 - 3.14159)
    fmt.Println(v29, v2)
    fmt.Println(v2, v2)
    v1 = append(v1, "\n\t")
    fmt.Println(v29)
    var v30 int = -711
    if (!(0 != 19)) {
        fmt.Println(v29, v30)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
