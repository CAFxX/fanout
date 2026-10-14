package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint8 }
func genFn0(p0 uint32, p1 uint32, p2 uint8) int32 {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 string = "hello"
        _ = fi7
        _ = v8
    }
    p1 = 949
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 427
    v3 := make(map[int]bool)
    var v4 int8 = 0
    var v5 uint32 = v2
    var v6 int32 = 2147483647
    v6 -= v6
    var v9 uint = 190
    switch "hello" {
    case "foo bar":
        var v10 uint = (v9 / ((18446744073709551615) | 1))
        _ = v10
    case "\n\t":
        v11 := GenS1{F0: "foo bar", F1: 128}
        _ = v11
    case "hello":
        fmt.Println(v9)
    }
    v9 -= uint(v2)
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 int32 = 1
        _ = fi12
        _ = v13
    }
    var v14 bool = true
    var v15 string = "hello"
    fmt.Println(v2, v2)
    v16 := []int{9223372036854775807, 2147483647}
    switch "" {
    case "hello":
        for fi17 := 0; fi17 < 16; fi17++ {
            if (!true) {
                var v18 int64 = -1
                var v19 uint = v9
                _ = v18
                _ = v19
            } else {
                var v20 int32 = int32(v5)
                _ = v20
            }
            v21 := GenS1{F0: v15, F1: (fi17 >> (((v9 + (v9 & (v9 ^ v9)))) & 63))}
            _ = fi17
            _ = v21
        }
    default:
        if v14 {
            for fi22 := 0; fi22 < 14; fi22++ {
                fmt.Println(v15, v4)
                for fi23 := 0; fi23 < 5; fi23++ {
                    v14 = v14
                    var v24 uint = (v9 % ((v9) | 1))
                    v25 := GenS2{F0: 255, F1: (v2 / ((v5) | 1))}
                    _ = len(v16)
                    var v26 int16 = 832
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                _ = fi22
            }
            var v27 uint = (^(v9))
            for fi28 := 0; fi28 < 20; fi28++ {
                for fi29 := 0; fi29 < 18; fi29++ {
                    if v14 {
                        var v30 uint = v9
                        var v31 uint = (v30 * (v30 >> ((v27) & 63)))
                        fmt.Println(fi29, v2)
                        _ = v30
                        _ = v31
                    }
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = fi29
                }
                if false {
                    if (true || (genFn0(v2, (v2 | 0), 2) > v6)) {
                        for fi32 := 0; fi32 < 16; fi32++ {
                            var v33 int = -2147483648
                            fmt.Println(v6, v5, v27)
                            if v14 {
                                fmt.Println(v4)
                                var v34 int64 = -1
                                v1[v2] = v14
                                _ = v34
                            } else {
                                var v35 int64 = 9223372036854775807
                                _ = v35
                            }
                            _ = fi32
                            _ = v33
                        }
                        v36 := GenS1{F0: v15, F1: (fi28 + fi28)}
                        v16 = append(v16, (-(fi28)))
                        _ = v36
                    } else {
                        for fi37 := 0; fi37 < 6; fi37++ {
                            var v38 int8 = v4
                            _ = fi37
                            _ = v38
                        }
                        var v39 uint8 = 74
                        _ = v39
                    }
                    if true {
                        fmt.Println(v15, fi28, v5)
                    } else {
                        for fi40 := 0; fi40 < 16; fi40++ {
                            var v41 float32 = (3.14159 / 0.1)
                            _ = fi40
                            _ = v41
                        }
                        v42 := GenS3{F0: 42, F1: 255}
                        _ = v42
                    }
                }
                _ = fi28
            }
            _ = v27
        }
    }
    v43 := GenS2{F0: uint8(v9), F1: (v5 + 221)}
    v44 := GenS4{F0: 42, F1: 816, F2: 0}
    var v45 uint16 = 65535
    v5 -= v2
    v9 -= (v9 & v9)
    v15 = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v9
}
