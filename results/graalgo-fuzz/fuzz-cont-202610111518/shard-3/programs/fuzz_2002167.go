package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int }
type GenS3 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS4 struct { F0 int; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int16) uint {
    p0 += p0
    if false {
        var v6 string = "a"
        _ = v6
    }
    _ = p0
    return 848
}
func genFn1(p0 bool, p1 uint32) int8 {
    v7 := GenS1{F0: -32768, F1: 1}
    var v8 bool = (-2.25 < 1.5)
    for fi9 := 0; fi9 < 7; fi9++ {
        if p0 {
            v8 = false
            var v10 string = "\n\t"
            fmt.Println(v10, v8, p0)
            _ = v10
        } else {
            v8 = ("a" >= "foo bar")
            if v8 {
                v8 = p0
            } else {
                v11 := GenS2{F0: 27, F1: 235, F2: 2147483647}
                v12 := []int{}
                _ = len(v12)
                _ = v11
                _ = v12
            }
        }
        _ = fi9
    }
    var v13 int64 = -616
    p0 = (p0 || (9223372036854775807 <= 791))
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int = 128
    v2 := []uint32{1, 1, 42}
    v3 := make(map[string]string)
    var v4 uint16 = 32767
    v5 := []uint32{}
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            v1 = 32767
            v16 := GenS2{F0: 255, F1: 2, F2: fi14}
            v4 -= v4
            var v17 int64 = 9223372036854775807
            var v18 int = (v1 & -1)
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        if (-663 < 1) {
            if (1 != -7) {
                switch "hello" {
                case "\n\t":
                    v19 := GenS3{F0: (v4 ^ v4), F1: false, F2: 216}
                    _ = v19
                }
                v4 -= (v4 + 65535)
            }
            if (1 != 365) {
                v20 := GenS1{F0: 1000, F1: 24}
                _ = v20
            } else {
                for fi21 := 0; fi21 < 18; fi21++ {
                    v5 = append(v5, uint32(v4))
                    _ = fi21
                }
            }
        }
        _ = fi14
    }
    v1 += v1
    v22 := GenS4{F0: -1, F1: 4294967295, F2: -9223372036854775808}
    switch v1 {
    case 128:
        for fi23 := 0; fi23 < 1; fi23++ {
            if false {
                var v24 int32 = 0
                _ = v24
            } else {
                var v25 int8 = -128
                _ = v25
            }
            if false {
                v1 = (^(v1))
                fmt.Println(v1, fi23, fi23)
                if ((3.14159 == float32((1.5 * ((-2.25 - 0.1) / 0.1)))) || false) {
                    v2 = append(v2, 2147483648)
                } else {
                    for fi26 := 0; fi26 < 8; fi26++ {
                        v3["hello"] = "hello"
                        var v27 int32 = -449
                        _ = fi26
                        _ = v27
                    }
                    v4 -= (^(v4))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            } else {
                v28 := GenS3{F0: v4, F1: (!(int16(fi23) < 32767)), F2: 2}
                _ = v28
            }
            for fi29 := 0; fi29 < 8; fi29++ {
                if true {
                    v1 = 128
                } else {
                    v30 := v3["hello"]
                    switch ("a" + (("foo bar" + (v30 + v30)) + v30)) {
                    case "":
                        fmt.Println(v4, fi23)
                    case "foo bar":
                        fmt.Println(v1, fi29, v1)
                    default:
                        v30 = v30
                    }
                    v31 := GenS4{F0: (fi23 << ((677) & 63)), F1: 1, F2: -1}
                    _ = v30
                    _ = v31
                }
                var v32 int64 = 873
                v1 -= (v1 % ((-2147483648) | 1))
                _ = fi29
                _ = v32
            }
            _ = fi23
        }
    case 563:
        v4 -= v4
    }
    v33 := GenS5{F0: 42}
    fmt.Println(v4, v4)
    v1 = v1
    var v34 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
