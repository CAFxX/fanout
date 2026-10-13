package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint8 }
func genFn0(p0 string, p1 bool, p2 uint16) uint {
    p2 = 2
    var v5 int32 = 42
    p2 += 297
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 int8 = 1
    var v4 int = 787
    fmt.Println(v4, v3)
    v4 = (v4 & -423)
    v6 := []int{(v4 << ((2) & 63)), v4, (v4 - v4), v4}
    fmt.Println(v3, v3)
    fmt.Println(v3, v3)
    if (!(-2.25 < (0.0 + 1e10))) {
        if false {
            var v7 uint64 = 4294967295
            v3 = (v3 % (((v3 | -48)) | 1))
            _ = v7
        } else {
            v8 := GenS1{F0: ((!true) || true), F1: 255, F2: false}
            _ = v8
        }
    } else {
        v4 += (v4 % ((v4) | 1))
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS2{F0: 0, F1: -32768, F2: (false || (128 != 42))}
            if (!false) {
                var v11 int = (v4 + -482)
                if true {
                    _ = len(v6)
                    _ = len(v6)
                    for fi12 := 0; fi12 < 2; fi12++ {
                        v13 := v1[256]
                        var v14 uint = 2
                        fmt.Println(v13)
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                } else {
                    var v15 float64 = 1.5
                    if ("" < ("a" + "hello")) {
                        v2[("a" + "\n\t")] = (true || true)
                        v16 := GenS3{F0: 65535}
                        if true {
                            v4 = v4
                        } else {
                            v17 := GenS1{F0: false, F1: 18446744073709551615, F2: ((v4 / ((fi9) | 1)) > v4)}
                            v15 = 0.1
                            var v18 float32 = float32(v15)
                            _ = v17
                            _ = v18
                        }
                        _ = v16
                    } else {
                        v19 := GenS1{F0: (!true), F1: 139, F2: false}
                        _ = v19
                    }
                    _ = v15
                }
                _ = v11
            }
            var v20 uint64 = 191
            _ = fi9
            _ = v10
            _ = v20
        }
    }
    fmt.Println(v3)
    v1[645] = 2
    fmt.Println(v3, v4)
    fmt.Println(v3)
    switch (v4 + (v4 | 271)) {
    case 307:
        v4 -= (v4 * v4)
    case 42:
        for fi21 := 0; fi21 < 14; fi21++ {
            switch ("" + "foo bar") {
            case "hello":
                v3 += (v3 >> ((genFn0((("a" + "a") + ("\n\t" + "hello")), false, 1)) & 63))
            case "\n\t":
                if ((762 <= 32767) && (false || (-799 >= 42))) {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        delete(v2, "foo bar")
                        var v23 uint64 = 4294967295
                        v24 := GenS4{F0: (v3 * v3), F1: 1}
                        v3 = -108
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            }
            _ = fi21
        }
    case -9223372036854775808:
        v25 := v2[(("foo bar" + "hello") + "foo bar")]
        _ = v25
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
