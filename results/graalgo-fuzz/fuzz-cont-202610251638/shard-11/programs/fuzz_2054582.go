package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64; F1 int8; F2 int16 }
func genFn0(p0 uint64, p1 int8) int {
    if true {
        v5 := GenS1{F0: 65535}
        fmt.Println(p1, p0)
        v6 := GenS1{F0: 117}
        _ = v5
        _ = v6
    }
    switch 573 {
    case 677:
        switch "\n\t" {
        case "foo bar":
            var v7 uint16 = 65535
            _ = v7
        case "a":
            var v8 uint16 = 607
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 4; fi9++ {
                switch ("a" + "hello") {
                case "":
                    p1 += p1
                }
                for fi10 := 0; fi10 < 8; fi10++ {
                    p0 = (p0 >> ((4294967295) & 63))
                    _ = fi10
                }
                _ = fi9
            }
        default:
            p1 -= (^(p1))
        }
    default:
        p1 = 1
    }
    fmt.Println(p0)
    var v11 uint32 = 2147483647
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v11
    return 4294967295
}

func main() {
    v1 := []string{"", "\n\t"}
    var v2 uint = 18446744073709551615
    var v3 uint8 = 255
    v4 := make(map[string]string)
    v12 := v4["\n\t"]
    fmt.Println(v3)
    v12 = v12
    var v13 int = -2147483648
    v14 := GenS2{F0: -9223372036854775808, F1: v3, F2: 2}
    switch genFn0(42, 100) {
    case 0:
        if true {
            var v15 int16 = 99
            _ = v15
        } else {
            var v16 int8 = -4
            v17 := v4["\n\t"]
            _ = v16
            _ = v17
        }
    case 1:
        v1 = append(v1, (v12 + "a"))
    case 4294967295:
        var v18 uint8 = 42
        _ = v18
    default:
        if (!(v12 <= ((((v12 + (((v12 + v12) + "") + v12)) + v12) + v12) + ("a" + v12)))) {
            var v19 int16 = 42
            _ = v19
        }
    }
    var v20 int32 = 1
    if true {
        v12 = v12
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        if false {
            if (391 > 42) {
                if ((v13 - (fi21 << ((2) & 63))) <= (-(v13))) {
                    fmt.Println(v3, fi21, v13)
                    if (0 != 18446744073709551615) {
                        for fi22 := 0; fi22 < 17; fi22++ {
                            var v23 int = v13
                            _ = fi22
                            _ = v23
                        }
                    } else {
                        v24 := GenS2{F0: -9223372036854775808, F1: (v3 | (v3 % ((v3) | 1))), F2: 130}
                        _ = v24
                    }
                    var v25 float64 = ((3.14159 + (3.14159 * 3.14159)) + 1.5)
                    _ = v25
                }
                v20 += v20
            }
            _ = len(v4)
        } else {
            v12 = "a"
            switch v13 {
            case -10:
                v26 := v4[(("\n\t" + v12) + v12)]
                _ = v26
            case -1:
                for fi27 := 0; fi27 < 7; fi27++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v28 := GenS3{F0: 778, F1: 100, F2: -1}
                    var v29 int16 = -1
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
        }
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
