package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
func genFn0(p0 int32, p1 uint, p2 uint16) int64 {
    v6 := []int{-9223372036854775808, 215}
    var v7 uint = (^(p1))
    v7 -= v7
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]int)
    var v3 float64 = float64(((1.5 / 1.0) / float32(3.14159)))
    var v4 uint32 = 320
    var v5 uint64 = 2
    v3 += (0.1 - 3.14159)
    var v8 string = "foo bar"
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int16 = -1
        var v11 int8 = v1
        for fi12 := 0; fi12 < 18; fi12++ {
            v3 = ((v3 * (-2.25 / v3)) + v3)
            switch v8 {
            case "a":
                v4 += 4294967295
            default:
                v5 += (v5 - v5)
            }
            _ = fi12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    if false {
        if true {
            switch v8 {
            case "a":
                switch v8 {
                case "foo bar":
                    fmt.Println(v3)
                case "a":
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v8 = v8
                        v1 += v1
                        var v14 int16 = -671
                        _ = fi13
                        _ = v14
                    }
                case "hello":
                    v8 = (v8 + v8)
                default:
                    var v15 uint = uint(v5)
                    _ = v15
                }
            case "foo bar":
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := []int{(fi16 | fi16), fi16, 127, 704}
                    _ = fi16
                    _ = v17
                }
            case "":
                v3 = 1e10
            }
            fmt.Println(v4, v1, v1)
        }
        var v18 float32 = 0.0
        _ = v18
    }
    v19 := v2[v4]
    var v20 bool = (true && ((^(v4)) >= v4))
    switch (v19 % ((v19) | 1)) {
    case 0:
        if (!(0 >= 1)) {
            if v20 {
                v21 := GenS1{F0: int(v1), F1: -1, F2: 1}
                _ = v21
            } else {
                var v22 uint16 = 42
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 uint32 = v4
                    if v20 {
                        v25 := GenS1{F0: 127, F1: 1, F2: 1}
                        v3 = v3
                        switch "hello" {
                        case "a":
                            var v26 string = v8
                            _ = v26
                        }
                        _ = v25
                    }
                    v5 -= v5
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
            for fi27 := 0; fi27 < 4; fi27++ {
                v28 := GenS2{F0: (v5 + 4294967295), F1: 921, F2: -310}
                switch ("\n\t" + v8) {
                case "\n\t":
                    var v29 int64 = 1
                    _ = v29
                case "hello":
                    fmt.Println(v20, v1)
                case "":
                    v30 := GenS2{F0: (^(v5)), F1: uint16(v5), F2: 32767}
                    _ = v30
                }
                v31 := []int{1}
                _ = fi27
                _ = v28
                _ = v31
            }
        }
    case 127:
        v32 := GenS3{F0: (v5 / ((v5) | 1)), F1: -128}
        _ = v32
    default:
        v33 := GenS4{F0: 2147483647}
        _ = v33
    }
    var v34 uint64 = v5
    if v20 {
        for fi35 := 0; fi35 < 18; fi35++ {
            v8 = v8
            switch fi35 {
            case 2147483647:
                v20 = v20
            default:
                for fi36 := 0; fi36 < 10; fi36++ {
                    v37 := []int{(fi35 % (((fi35 + 341)) | 1)), (-(fi36))}
                    _ = fi36
                    _ = v37
                }
            }
            v20 = true
            _ = fi35
        }
        v2[v4] = (v19 - int(v5))
    }
    v38 := GenS2{F0: 18446744073709551615, F1: 65535, F2: 1000}
    if ((v20 && v20) || v20) {
        fmt.Println(v3, v8, v1)
    }
    var v39 bool = (99 <= 32767)
    fmt.Println(v4)
    for fi40 := 0; fi40 < 3; fi40++ {
        fmt.Println(v34, v1, v8)
        v34 -= (v5 % ((1) | 1))
        if true {
            for fi41 := 0; fi41 < 18; fi41++ {
                v39 = false
                v8 = ""
                _ = fi41
            }
        } else {
            v42 := GenS4{F0: (fi40 << ((1) & 63))}
            v43 := GenS5{F0: genFn0(1, 18446744073709551615, 584)}
            _ = v42
            _ = v43
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v8
}
