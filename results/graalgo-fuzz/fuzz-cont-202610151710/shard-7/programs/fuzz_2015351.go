package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int16; F2 int16 }
type GenS4 struct { F0 string; F1 int64 }
func genFn0() uint64 {
    var v5 int8 = -107
    v6 := GenS1{F0: true, F1: "hello"}
    v7 := []int{32768}
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 int32, p2 int64) int32 {
    if (!(127 < 127)) {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := []int{2147483647, 128, fi8, (fi8 >> ((0) & 63))}
            if len(v9) > 0 {
                v9[0] = (fi8 >> ((92) & 63))
            }
            if false {
                switch ("a" + "hello") {
                case "hello":
                    var v10 float64 = 1.5
                    _ = v10
                default:
                    p1 -= p1
                }
                _ = len(v9)
            } else {
                switch (fi8 + fi8) {
                case -9223372036854775808:
                    switch (-(fi8)) {
                    case -784:
                        if (!false) {
                            var v11 bool = (false && true)
                            if (!v11) {
                                v11 = ((p1 != (p1 | (p1 | (p1 % (((p1 + (p1 ^ p1))) | 1))))) || (!(62 == 42)))
                            }
                            var v12 string = "foo bar"
                            _ = v11
                            _ = v12
                        } else {
                            switch ("hello" + "a") {
                            case "\n\t":
                                fmt.Println(fi8, p0)
                            default:
                                v13 := GenS2{F0: 0, F1: 1}
                                _ = v13
                            }
                            v14 := GenS2{F0: 201, F1: p0}
                            _ = v14
                        }
                    case 2147483647:
                        var v15 string = (("\n\t" + "hello") + ("\n\t" + "foo bar"))
                        _ = v15
                    case 32768:
                        var v16 float64 = 0.0
                        _ = v16
                    }
                }
                v17 := GenS2{F0: 1, F1: p0}
                _ = v17
            }
            _ = fi8
            _ = v9
        }
    } else {
        v18 := []int{42}
        _ = v18
    }
    p0 = (p0 & p0)
    _ = p0
    _ = p1
    _ = p2
    return (^(p1))
}

func main() {
    v1 := []int{}
    v2 := make(map[int]uint32)
    var v3 uint64 = 0
    v4 := []string{("a" + "hello"), "a"}
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 float64 = (-2.25 / float64(((0.0 - 0.0) + (3.14159 * (1e10 / ((3.14159 * 3.14159) / (3.14159 + 1e10)))))))
        _ = fi19
        _ = v20
    }
    if true {
        v3 = genFn0()
        _ = len(v2)
        fmt.Println(v3)
    }
    if ((4294967295 >= 0) && (2147483647 < 2147483647)) {
        var v21 float64 = -2.25
        for fi22 := 0; fi22 < 7; fi22++ {
            fmt.Println(fi22, v3, v3)
            _ = fi22
        }
        _ = v21
    } else {
        v23 := GenS3{F0: v3, F1: 42, F2: 1000}
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        v3 += 0
        v3 = 786
        _ = fi24
    }
    _ = len(v4)
    v3 += (v3 + (^(v3)))
    if (true || ((v3 ^ 18446744073709551615) != genFn0())) {
        fmt.Println(v3)
        v2[-831] = 65535
    } else {
        _ = len(v2)
        v25 := GenS3{F0: 0, F1: 32767, F2: -505}
        v26 := GenS2{F0: v3, F1: 150}
        _ = v25
        _ = v26
    }
    fmt.Println(v3)
    v3 -= v3
    if false {
        v3 -= 1
        for fi27 := 0; fi27 < 3; fi27++ {
            fmt.Println(fi27, v3)
            _ = fi27
        }
        v28 := GenS2{F0: 439, F1: 42}
        _ = v28
    }
    fmt.Println(v3)
    if ("\n\t" != ((("a" + ("" + "\n\t")) + ("" + "\n\t")) + ("hello" + "foo bar"))) {
        v29 := GenS4{F0: "\n\t", F1: 0}
        _ = v29
    }
    fmt.Println(v3)
    v3 -= 42
    v2[-414] = 2147483647
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 int32 = 42
        _ = fi30
        _ = v31
    }
    var v32 int16 = -540
    var v33 uint8 = 96
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
