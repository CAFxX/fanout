package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 string) bool {
    var v5 float64 = 0.1
    if true {
        p0 = p0
        if false {
            var v6 uint = 275
            switch 466 {
            case -9223372036854775808:
                fmt.Println(p0, v5, v5)
            case -869:
                v7 := []int{4294967295}
                _ = v7
            case 4294967295:
                fmt.Println(v5, p0, v6)
            default:
                if (p0 <= p0) {
                    p0 = ""
                    v6 -= v6
                }
            }
            _ = v6
        }
    } else {
        var v8 uint16 = 42
        _ = v8
    }
    _ = p0
    _ = v5
    return ((!(false && (!true))) && false)
}
func genFn1(p0 int, p1 uint, p2 string) uint32 {
    if ((p0 <= 4294967295) || true) {
        var v9 float32 = float32(((-2.25 + (3.14159 - 1.5)) + -2.25))
        v10 := []int{(^(p0)), p0}
        var v11 int8 = 100
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := []int{(p0 & p0), (p0 / ((p0) | 1)), (p0 * (p0 >> ((0) & 63)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 2147483648
}

func main() {
    v1 := []string{"\n\t", "", ("hello" + (("a" + ("foo bar" + "\n\t")) + (("foo bar" + ("hello" + "hello")) + "")))}
    var v2 int16 = 0
    v3 := make(map[int]uint32)
    v4 := []string{}
    fmt.Println(v2)
    _ = len(v3)
    v2 = -19
    for fi13 := 0; fi13 < 8; fi13++ {
        v2 += (^(v2))
        var v14 uint16 = 2
        v14 = (v14 % (((v14 | v14)) | 1))
        _ = fi13
        _ = v14
    }
    switch 9223372036854775807 {
    case 32767:
        _ = len(v3)
    case 9223372036854775807:
        v1 = append(v1, "a")
    }
    v2 = v2
    if true {
        _ = len(v4)
    } else {
        var v15 uint16 = 65535
        _ = v15
    }
    v16 := GenS1{F0: 1}
    var v17 int = 175
    var v18 uint64 = 1
    v17 = (v17 - v17)
    switch (v17 >> ((18446744073709551615) & 63)) {
    case 32768:
        var v19 string = ((((("foo bar" + "") + (("foo bar" + ("hello" + "hello")) + (("\n\t" + "\n\t") + "\n\t"))) + "\n\t") + "") + "\n\t")
        _ = v19
    }
    v18 = (v18 | (v18 & (v18 >> ((65535) & 63))))
    if (0.0 <= 1e10) {
        if true {
            var v20 int8 = 1
            var v21 uint8 = 1
            v17 -= int(v21)
            _ = v20
            _ = v21
        }
        var v22 uint8 = 147
        _ = v22
    }
    v23 := GenS1{F0: 0}
    if false {
        for fi24 := 0; fi24 < 2; fi24++ {
            v18 -= v18
            v2 += -113
            if (1 != 0) {
                var v25 float64 = 1.5
                v26 := v3[fi24]
                v2 -= v2
                _ = v25
                _ = v26
            } else {
                v18 = (v18 / ((v18) | 1))
            }
            _ = fi24
        }
        if true {
            if true {
                for fi27 := 0; fi27 < 12; fi27++ {
                    for fi28 := 0; fi28 < 8; fi28++ {
                        switch v17 {
                        case -821:
                            if (-128 == 127) {
                                fmt.Println(v18, v17, v18)
                            }
                        }
                        v17 = (-(v17))
                        _ = fi28
                    }
                    v2 = (v2 & v2)
                    fmt.Println(v17, v2, v18)
                    _ = fi27
                }
            } else {
                v3[42] = 241
                v18 = (v18 ^ 1)
                var v29 uint = 42
                _ = v29
            }
        } else {
            v30 := GenS2{F0: (-(v2)), F1: 243, F2: 42}
            switch v17 {
            case 32768:
                v31 := GenS3{F0: (v18 & (v18 ^ v18))}
                _ = v31
            case -1:
                v2 -= v2
            }
            var v32 int8 = -128
            _ = v30
            _ = v32
        }
    }
    v33 := v3[v17]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v4
}
