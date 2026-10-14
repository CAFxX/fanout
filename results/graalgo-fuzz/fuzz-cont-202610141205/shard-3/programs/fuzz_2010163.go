package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint32; F1 bool }
func genFn0(p0 uint, p1 uint8, p2 int) bool {
    var v5 float32 = (1.5 - -2.25)
    var v6 bool = true
    var v7 uint64 = 42
    p1 += 254
    v8 := GenS1{F0: (42 <= -128), F1: 32767, F2: (v7 ^ v7)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return ((!v6) || ((p0 * p0) >= p0))
}
func genFn1(p0 uint16, p1 int8) uint64 {
    if false {
        var v9 int64 = -9223372036854775808
        switch (("hello" + "") + "a") {
        case "\n\t":
            v10 := GenS2{F0: 1, F1: (p0 << ((732) & 63)), F2: (p0 % ((65535) | 1))}
            _ = v10
        case "foo bar":
            var v11 float32 = float32(1.5)
            _ = v11
        case "":
            v9 = -1
        default:
            switch int(v9) {
            case 42:
                if (!(618 <= 1)) {
                    v12 := GenS1{F0: true, F1: -1, F2: 4294967295}
                    _ = v12
                } else {
                    var v13 string = "hello"
                    v14 := GenS3{F0: 2147483647, F1: v13}
                    var v15 int32 = -431
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case 9223372036854775807:
                v16 := []int{-2147483648, -9223372036854775808, 4294967295, 0}
                _ = v16
            case 1:
                v17 := GenS1{F0: (9223372036854775807 > -9223372036854775808), F1: 1, F2: 1}
                _ = v17
            default:
                fmt.Println(p1, p1, p0)
            }
        }
        var v18 float32 = (0.1 / 3.14159)
        _ = v18
        _ = v9
    }
    var v19 uint32 = 2147483647
    p0 -= 255
    var v20 uint32 = v19
    v19 = 80
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 4294967295
}

func main() {
    var v1 int16 = 1000
    v2 := []int{127, 127}
    var v3 uint8 = 1
    var v4 int = 385
    var v21 string = ("\n\t" + "foo bar")
    v21 = "a"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v21 = v21
    v1 = v1
    switch v4 {
    case 2147483647:
        if (true || false) {
            if false {
                v2 = append(v2, (v4 / (((v4 * -726)) | 1)))
            }
            if true {
                v2 = append(v2, v4)
            } else {
                if false {
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v21 = "a"
                        if ((int(v3) <= (fi22 / ((9223372036854775807) | 1))) && (!false)) {
                            if (!(18446744073709551615 <= 0)) {
                                switch v21 {
                                case "foo bar":
                                    v23 := GenS1{F0: false, F1: (v1 & -32768), F2: 4294967295}
                                    _ = v23
                                case "hello":
                                    if len(v2) > 0 {
                                        v2[0] = (v4 / (((fi22 / ((127) | 1))) | 1))
                                    }
                                case "":
                                    fmt.Println(fi22, fi22, fi22)
                                }
                            } else {
                                v3 -= (v3 | (v3 * (^(v3))))
                                v21 = v21
                            }
                            v1 += (v1 + 0)
                        } else {
                            v3 -= v3
                        }
                        _ = fi22
                    }
                    var v24 bool = false
                    _ = v24
                }
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                var v26 uint64 = 737
                _ = fi25
                _ = v26
            }
        }
    default:
        v3 -= v3
    }
    var v27 int8 = 42
    v27 += v27
    var v28 int = v4
    v29 := GenS4{F0: 65535, F1: (14 <= 193)}
    v2 = append(v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
