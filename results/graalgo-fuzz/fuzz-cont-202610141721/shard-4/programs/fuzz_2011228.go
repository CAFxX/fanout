package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 uint16, p1 uint) string {
    v7 := []int{9223372036854775807}
    switch (("a" + "") + "\n\t") {
    case "":
        if (!(-67 >= 42)) {
            fmt.Println(p0, p0)
            _ = len(v7)
            var v8 float32 = 1.5
            _ = v8
        } else {
            var v9 float64 = 3.14159
            var v10 float32 = (3.14159 - (((float32(v9) - 3.14159) - ((3.14159 * 1e10) / (0.0 / (-2.25 * (((1e10 * -2.25) * (1.5 - float32((v9 * v9)))) * (1.5 + 3.14159)))))) + (1.5 / 1e10)))
            p0 = p0
            _ = v10
            _ = v9
        }
    }
    if (2147483647 == -274) {
        p0 = p0
        p1 = 0
    } else {
        p0 += p0
    }
    var v11 int = -474
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return ""
}
func genFn1(p0 uint8) bool {
    fmt.Println(p0)
    fmt.Println(p0)
    var v12 int64 = 42
    _ = p0
    _ = v12
    return true
}

func main() {
    var v1 float32 = (0.0 - 0.1)
    var v2 uint16 = 56
    var v3 string = "\n\t"
    var v4 uint64 = 42
    v5 := make(map[int]uint32)
    v6 := make(map[int]bool)
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: 42, F1: (fi13 + -683)}
        var v15 uint64 = (^(v4))
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v3, v1, v4)
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(fi16)
        if (!(32767 > 1)) {
            v1 = 3.14159
            var v17 int8 = 57
            var v18 uint = 575
            _ = v17
            _ = v18
        } else {
            v19 := v5[2147483647]
            fmt.Println(v1)
            _ = v19
        }
        _ = fi16
    }
    var v20 int = 9223372036854775807
    if true {
        for fi21 := 0; fi21 < 3; fi21++ {
            v2 -= 1
            for fi22 := 0; fi22 < 3; fi22++ {
                switch (fi21 % ((v20) | 1)) {
                case -29:
                    v20 -= (-(fi21))
                case 4294967295:
                    var v23 int64 = -9223372036854775808
                    _ = v23
                default:
                    v24 := GenS2{F0: (v4 + 4294967295), F1: 65535, F2: (v2 / ((v2) | 1))}
                    _ = v24
                }
                fmt.Println(v4)
                var v25 int64 = 642
                _ = fi22
                _ = v25
            }
            var v26 uint16 = (v2 & v2)
            _ = fi21
            _ = v26
        }
        v27 := GenS2{F0: v4, F1: 42, F2: 616}
        _ = v27
    } else {
        v28 := GenS2{F0: v4, F1: 18446744073709551615, F2: (v2 ^ 2)}
        v29 := GenS2{F0: v4, F1: 283, F2: (^(v2))}
        _ = v28
        _ = v29
    }
    v30 := GenS3{F0: 736, F1: 0}
    if true {
        v1 = v1
        v3 = v3
        var v31 float64 = 1e10
        _ = v31
    } else {
        for fi32 := 0; fi32 < 7; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                var v34 int16 = -587
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
        var v35 float32 = 0.1
        _ = v35
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        v3 = "a"
        _ = fi36
    }
    if true {
        v37 := []int{v20, v20, (v20 / ((v20) | 1))}
        v4 = (^(v4))
        _ = v37
    }
    var v38 float64 = (-2.25 * 1e10)
    v39 := GenS4{F0: (v3 + v3), F1: 18446744073709551615}
    var v40 int8 = -36
    v4 -= (v4 * (v4 * (v4 ^ (v4 & v4))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
