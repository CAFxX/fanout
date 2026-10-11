package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint; F1 string }
func genFn0(p0 int32) string {
    var v5 uint16 = 908
    v6 := GenS1{F0: uint32(v5)}
    _ = p0
    _ = v5
    _ = v6
    return ("a" + "a")
}
func genFn1(p0 uint32, p1 uint32) uint {
    p1 = p1
    var v7 uint64 = 18446744073709551615
    var v8 uint64 = v7
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 737
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -9223372036854775808
    var v3 int8 = 0
    var v4 uint16 = 65535
    v9 := v1[1]
    v3 = (v3 - (^(v3)))
    fmt.Println(v3)
    switch -1 {
    case 32767:
        v10 := GenS2{F0: 2147483647}
        _ = v10
    }
    v11 := v1[2147483647]
    v12 := []int{128, 42, 42}
    v13 := GenS3{F0: 4294967295, F1: 825}
    v14 := GenS4{F0: v2, F1: 42}
    for fi15 := 0; fi15 < 9; fi15++ {
        if (3.14159 > 0.1) {
            var v16 float32 = ((1e10 - 1e10) / float32(float64(((-2.25 * 1e10) * 1e10))))
            _ = v16
        } else {
            if (!false) {
                var v17 int64 = v2
                if false {
                    for fi18 := 0; fi18 < 20; fi18++ {
                        v17 += (-(v17))
                        switch (-(fi18)) {
                        case -474:
                            v17 -= (^(v2))
                        case -665:
                            if len(v12) > 0 {
                                v12[0] = -9223372036854775808
                            }
                        default:
                            v3 = (-(v3))
                        }
                        v19 := GenS2{F0: 2147483647}
                        _ = fi18
                        _ = v19
                    }
                }
                _ = v17
            }
            _ = len(v1)
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            v9 = v11
            _ = fi20
        }
        switch v11 {
        case "foo bar":
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 int16 = 691
                _ = fi21
                _ = v22
            }
        case "hello":
            if (v4 >= v4) {
                switch v11 {
                case "\n\t":
                    for fi23 := 0; fi23 < 5; fi23++ {
                        v2 += 184
                        _ = fi23
                    }
                }
            } else {
                delete(v1, 280)
            }
        case "":
            if false {
                var v24 bool = (!false)
                v25 := v1[0]
                v2 -= 42
                _ = v24
                _ = v25
            } else {
                for fi26 := 0; fi26 < 1; fi26++ {
                    fmt.Println(v4)
                    var v27 uint16 = v4
                    v9 = v11
                    _ = fi26
                    _ = v27
                }
                v28 := GenS5{F0: 546, F1: v11}
                _ = v28
            }
        default:
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
        _ = fi15
    }
    v9 = (v11 + (v9 + v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
