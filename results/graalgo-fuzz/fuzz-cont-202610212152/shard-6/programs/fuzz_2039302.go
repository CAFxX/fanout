package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int16; F1 int8 }
func genFn0(p0 bool, p1 uint8, p2 uint16) int16 {
    if (1 == -1) {
        p1 = p1
        var v5 uint16 = 255
        var v6 bool = false
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: -1, F1: 42}
    p1 -= (p1 & 1)
    p1 = (p1 ^ (p1 ^ p1))
    var v8 uint8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int16 = -582
    var v2 uint32 = 134
    var v3 int8 = -1
    var v4 int32 = -401
    var v9 uint8 = 1
    v4 += 0
    fmt.Println(v9, v3, v1)
    v10 := []int{4294967295}
    if false {
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS2{F0: v3, F1: 0}
            v13 := GenS2{F0: v3, F1: 28}
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 uint8 = 127
        for fi15 := 0; fi15 < 6; fi15++ {
            switch ("foo bar" + ("" + "foo bar")) {
            case "a":
                if (375 > 65535) {
                    if len(v10) > 0 {
                        v10[0] = -2147483648
                    }
                    var v16 uint32 = v2
                    _ = v16
                } else {
                    fmt.Println(v14, v1, v14)
                }
            case "hello":
                if false {
                    v17 := GenS1{F0: (fi15 % ((fi15) | 1)), F1: (fi15 & (fi15 + (fi15 / (((fi15 / (((fi15 % ((285) | 1))) | 1))) | 1))))}
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 float32 = 1e10
                        _ = len(v10)
                        _ = fi18
                        _ = v19
                    }
                    _ = v17
                }
            default:
                v20 := GenS3{F0: v2, F1: (v1 % ((genFn0(false, (v14 ^ 0), 42)) | 1))}
                _ = v20
            }
            _ = fi15
        }
        _ = v14
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 1; fi22++ {
            if (!(!(true || true))) {
                v1 = (v1 | (v1 + (v1 * 1)))
                v1 = genFn0(false, v9, 65535)
                if false {
                    fmt.Println(v4, fi22, fi22)
                    v23 := GenS2{F0: (v3 % ((31) | 1)), F1: 2}
                    v9 += 2
                    var v24 bool = false
                    _ = v23
                    _ = v24
                } else {
                    var v25 float32 = 1e10
                    v3 = -1
                    _ = v25
                }
                v26 := GenS3{F0: v2, F1: v1}
                v27 := GenS4{F0: v1, F1: (v3 & (v3 % (((v3 % (((v3 >> ((18446744073709551615) & 63))) | 1))) | 1)))}
                _ = v26
                _ = v27
            }
            var v28 uint64 = 4294967295
            _ = fi22
            _ = v28
        }
        _ = fi21
    }
    v9 -= v9
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v1, v3, v9)
    case "a":
        if (!(false || false)) {
            switch "a" {
            case "\n\t":
                var v29 string = "a"
                _ = v29
            }
        } else {
            if len(v10) > 0 {
                v10[0] = 9223372036854775807
            }
        }
    default:
        if true {
            v30 := GenS4{F0: (v1 ^ genFn0(false, 245, 73)), F1: (v3 + -128)}
            _ = v30
        } else {
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 string = "foo bar"
                _ = fi31
                _ = v32
            }
            _ = len(v10)
            var v33 int32 = v4
            _ = v33
        }
    }
    v4 = (v4 ^ -437)
    v34 := GenS4{F0: v1, F1: 42}
    for fi35 := 0; fi35 < 6; fi35++ {
        v3 = (v3 % (((^(v3))) | 1))
        v36 := GenS1{F0: 9223372036854775807, F1: 4294967295}
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v9
}
