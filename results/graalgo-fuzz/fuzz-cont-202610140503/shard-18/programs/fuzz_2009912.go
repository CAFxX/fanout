package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 string }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int16; F1 int8; F2 string }
type GenS4 struct { F0 int16; F1 int }
func genFn0(p0 string, p1 uint32, p2 uint16) int {
    fmt.Println(p2, p0, p2)
    if ((p0 != ("\n\t" + (("\n\t" + "") + p0))) || false) {
        p0 = p0
        if (true && (0 <= 2)) {
            if (18446744073709551615 <= 2) {
                p2 = (p2 << ((uint(p2)) & 63))
                p2 += (p2 & (p2 ^ (p2 << ((0) & 63))))
            } else {
                var v4 uint64 = 0
                var v5 int32 = -2147483648
                _ = v4
                _ = v5
            }
            v6 := GenS1{F0: 128, F1: (p2 >> ((25) & 63)), F2: p0}
            var v7 uint32 = 59
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: 1, F1: (p2 << ((0) & 63)), F2: (p0 + p0)}
            _ = v8
        }
        switch 1 {
        case 128:
            if (p2 > 1) {
                p0 = p0
                var v9 int32 = 919
                _ = v9
            }
        case -1:
            if ((1e10 != 0.1) || true) {
                var v10 bool = (83 != 751)
                _ = v10
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 uint16) uint32 {
    var v11 bool = false
    var v12 uint32 = 2147483647
    _ = p0
    _ = v11
    _ = v12
    return (v12 + (v12 * (v12 | 2147483647)))
}

func main() {
    var v1 int16 = -1
    v2 := []string{"foo bar", ((((("hello" + ("hello" + ("foo bar" + ("a" + "\n\t")))) + (("\n\t" + "\n\t") + ("foo bar" + "foo bar"))) + "foo bar") + "a") + "foo bar")}
    var v3 uint8 = 2
    if (!false) {
        var v13 int8 = 127
        _ = v13
    }
    var v14 uint = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v15 bool = (!false)
    if true {
        var v16 int16 = (v1 ^ v1)
        _ = v16
    } else {
        if v15 {
            if (128 < 1) {
                v17 := GenS2{F0: v15, F1: 65535}
                _ = v17
            } else {
                for fi18 := 0; fi18 < 15; fi18++ {
                    if (9223372036854775807 != -9223372036854775808) {
                        var v19 string = "a"
                        v14 = (v14 / (((^(v14))) | 1))
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v15 = false
                        var v21 float32 = 1e10
                        var v22 bool = false
                        v3 += (v3 << (((v14 & 42)) & 63))
                        var v23 int8 = 39
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi18
                }
            }
            var v24 uint64 = 42
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = v24
        }
        fmt.Println(v3, v15, v3)
        v1 -= (v1 - int16(v14))
    }
    v25 := GenS3{F0: v1, F1: 0, F2: ((("hello" + (("foo bar" + "hello") + "")) + ("\n\t" + (("foo bar" + "hello") + "foo bar"))) + "")}
    v14 -= (v14 << ((v14) & 63))
    if false {
        switch "" {
        case "a":
            if true {
                switch 42 {
                case -1:
                    var v26 uint16 = 65535
                    _ = v26
                case -9223372036854775808:
                    v27 := GenS4{F0: 623, F1: 42}
                    _ = v27
                default:
                    var v28 int8 = -79
                    _ = v28
                }
            }
        }
    }
    if false {
        var v29 int64 = -777
        if (v14 != 255) {
            v2 = append(v2, "\n\t")
            var v30 int32 = 2147483647
            _ = v30
        }
        v31 := GenS3{F0: v1, F1: 100, F2: ("\n\t" + ("a" + "\n\t"))}
        _ = v29
        _ = v31
    }
    var v32 uint = (v14 << ((v14) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v32
}
