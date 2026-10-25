package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int8 }
type GenS2 struct { F0 int64; F1 uint16; F2 int }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint64; F1 uint64; F2 int }
type GenS5 struct { F0 uint32; F1 int32 }

func main() {
    v1 := []string{"", "a", ("\n\t" + ("" + "\n\t")), ("hello" + (("foo bar" + "foo bar") + "foo bar"))}
    var v2 int = 2147483647
    var v3 uint = 653
    v4 := make(map[int]bool)
    var v5 uint = (v3 + 18446744073709551615)
    var v6 int16 = 1
    v3 = (v3 & 594)
    v5 += 518
    switch 2147483647 {
    case 9223372036854775807:
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int64 = 9223372036854775807
            for fi9 := 0; fi9 < 1; fi9++ {
                v3 = (v3 ^ (v3 % (((v5 / (((v3 * v5)) | 1))) | 1)))
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    case 819:
        if len(v1) > 0 {
            v1[0] = "a"
        }
    default:
        if false {
            v10 := GenS1{F0: (v6 & 145), F1: 1, F2: 127}
            v11 := GenS1{F0: 42, F1: 4294967295, F2: -107}
            _ = v10
            _ = v11
        } else {
            v12 := GenS2{F0: -349, F1: 255, F2: (v2 % ((v2) | 1))}
            v5 = v3
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS2{F0: -104, F1: 893, F2: fi13}
        v6 += (v6 * v6)
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        _ = fi13
        _ = v14
    }
    if false {
        v6 = v6
        v15 := v4[v2]
        var v16 bool = (!true)
        _ = v15
        _ = v16
    }
    v6 = (v6 / (((v6 + v6)) | 1))
    _ = len(v1)
    var v17 float32 = float32(0.0)
    v4[v2] = false
    v17 = v17
    v3 += (v5 * v5)
    var v18 int16 = v6
    switch "\n\t" {
    case "":
        fmt.Println(v2)
    default:
        for fi19 := 0; fi19 < 14; fi19++ {
            fmt.Println(fi19)
            v20 := GenS1{F0: v18, F1: 2, F2: 42}
            var v21 int16 = 1
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 13; fi23++ {
            switch "a" {
            case "foo bar":
                if false {
                    v3 += v5
                    var v24 bool = true
                    v25 := GenS1{F0: -1, F1: 18446744073709551615, F2: -123}
                    _ = v24
                    _ = v25
                }
            default:
                _ = len(v1)
            }
            v3 += v3
            var v26 uint64 = 0
            switch ("" + (("foo bar" + "a") + "foo bar")) {
            case "a":
                v5 -= (v3 | (^(v5)))
            default:
                v27 := GenS3{F0: 152, F1: false}
                v3 = v5
                v1 = append(v1, (("foo bar" + "") + "a"))
                _ = v27
            }
            var v28 int16 = v6
            switch (v2 + fi23) {
            case -2147483648:
                var v29 uint64 = 0
                _ = v29
            case 780:
                switch "\n\t" {
                case "a":
                    v30 := GenS2{F0: 9223372036854775807, F1: 255, F2: int(v26)}
                    _ = v30
                default:
                    switch (^(fi23)) {
                    case 32767:
                        if ((v5 <= 255) || true) {
                            v17 -= float32(float64((-2.25 * v17)))
                            var v31 uint8 = 127
                            _ = v31
                        }
                    case 128:
                        v32 := v4[(v2 ^ 0)]
                        _ = v32
                    case -9223372036854775808:
                        v26 = v26
                        v33 := v4[fi23]
                        _ = v33
                    }
                }
            case -65:
                v34 := GenS3{F0: 5, F1: (false || (v18 >= 32767))}
                if len(v1) > 0 {
                    v1[0] = ((("hello" + "a") + "\n\t") + ("a" + "hello"))
                }
                _ = v34
            default:
                var v35 int = v2
                _ = v35
            }
            v36 := GenS1{F0: v18, F1: v26, F2: int8(v5)}
            if len(v1) > 0 {
                v1[0] = "a"
            }
            var v37 int64 = -9223372036854775808
            if (fi23 > 9223372036854775807) {
                v38 := v4[(fi23 >> ((748) & 63))]
                v37 -= (v37 % ((v37) | 1))
                if true {
                    fmt.Println(v37)
                    var v39 float32 = v17
                    _ = v39
                }
                _ = v38
            }
            if true {
                v40 := GenS2{F0: (v37 % (((^(v37))) | 1)), F1: uint16(v3), F2: fi22}
                v41 := GenS2{F0: v37, F1: 1, F2: -2147483648}
                _ = v40
                _ = v41
            }
            delete(v4, (fi23 >> ((0) & 63)))
            _ = fi23
            _ = v26
            _ = v28
            _ = v36
            _ = v37
        }
        fmt.Println(v18, v18, v2)
        v42 := GenS4{F0: 350, F1: 42, F2: (^(fi22))}
        _ = fi22
        _ = v42
    }
    v43 := GenS5{F0: 643, F1: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
