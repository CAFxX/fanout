package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint8; F2 string }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 uint, p2 bool) int32 {
    var v6 uint16 = 86
    var v7 string = "a"
    fmt.Println(v7)
    if true {
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(v7, v6, fi8)
            _ = fi8
        }
    } else {
        v6 -= 488
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 uint8) int16 {
    p0 += p0
    var v9 int = 32768
    var v10 float32 = ((-2.25 / 1e10) - (0.0 + 0.1))
    _ = p0
    _ = v10
    _ = v9
    return -1
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]bool)
    var v3 int = -183
    v4 := []string{"foo bar", "foo bar", ("hello" + ("a" + "hello")), "a"}
    var v5 float32 = ((3.14159 - (float32(0.0) * -2.25)) * (1.5 - 3.14159))
    v11 := GenS1{F0: 4294967295, F1: "", F2: 863}
    fmt.Println(v5, v5, v1)
    v12 := GenS1{F0: 0, F1: "foo bar", F2: genFn0(v1, 255, (1 < -1))}
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            if ((false && (32767 == 1000)) && false) {
                var v15 uint = 18446744073709551615
                fmt.Println(fi13, v1, v5)
                _ = v15
            } else {
                var v16 bool = (108 >= (v1 & (v1 * -1)))
                var v17 int64 = -1
                fmt.Println(v17, fi14, fi14)
                _ = v16
                _ = v17
            }
            if (0 <= -9223372036854775808) {
                v18 := GenS2{F0: 1, F1: 1, F2: "\n\t"}
                switch v3 {
                case 1:
                    fmt.Println(v3, v3, v5)
                case 128:
                    switch (-(v3)) {
                    case 128:
                        fmt.Println(fi13)
                    case -2147483648:
                        var v19 int8 = (v1 << ((0) & 63))
                        _ = v19
                    default:
                        var v20 int16 = 1
                        _ = v20
                    }
                default:
                    var v21 bool = true
                    _ = v21
                }
                var v22 float32 = (v5 / (float32(3.14159) * (v5 / -2.25)))
                switch ("a" + ("foo bar" + "a")) {
                case "foo bar":
                    v22 = ((0.0 / v22) - (v5 + 0.0))
                case "\n\t":
                    v23 := GenS2{F0: 4294967295, F1: 0, F2: "a"}
                    _ = v23
                case "a":
                    v1 = v1
                default:
                    v4 = append(v4, "\n\t")
                }
                delete(v2, 0)
                _ = v18
                _ = v22
            }
            _ = fi14
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 int32 = -1
            if true {
                v3 = (fi24 ^ (fi24 * v3))
                var v26 string = "\n\t"
                fmt.Println(v1, v5, fi13)
                _ = v26
            }
            _ = fi24
            _ = v25
        }
        _ = fi13
    }
    fmt.Println(v5, v1)
    if (457 < 65535) {
        v27 := GenS2{F0: 18446744073709551615, F1: 42, F2: ""}
        _ = v27
    } else {
        fmt.Println(v3)
        var v28 int = v3
        _ = v28
    }
    if true {
        v29 := GenS3{F0: 65535, F1: 681}
        v3 -= (-(v3))
        _ = v29
    } else {
        v30 := GenS4{F0: 42}
        var v31 uint16 = 783
        _ = v30
        _ = v31
    }
    var v32 uint16 = 87
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v33 int = v3
    v34 := v2[(v33 * 127)]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
