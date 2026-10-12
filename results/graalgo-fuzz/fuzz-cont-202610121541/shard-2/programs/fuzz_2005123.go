package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int8; F2 int64 }
type GenS4 struct { F0 int; F1 bool; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 int8 = 0
    var v2 int32 = -1
    var v3 int64 = 888
    v4 := make(map[int]uint32)
    var v5 uint32 = 2147483647
    var v6 int = 9223372036854775807
    v5 = 2147483647
    fmt.Println(v5, v1)
    delete(v4, 42)
    v3 = (v3 & v3)
    v2 = v2
    fmt.Println(v6)
    fmt.Println(v3, v2, v5)
    if false {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 bool = ((v1 >> ((0) & 63)) != (v1 << ((18446744073709551615) & 63)))
            for fi9 := 0; fi9 < 1; fi9++ {
                delete(v4, fi9)
                _ = fi9
            }
            v10 := GenS1{F0: v6, F1: v2}
            _ = fi7
            _ = v10
            _ = v8
        }
        fmt.Println(v5, v6)
        fmt.Println(v3)
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            if false {
                var v12 int32 = 392
                _ = v12
            } else {
                v13 := GenS1{F0: (v6 + -679), F1: 0}
                fmt.Println(fi11, v1, v5)
                if (-980 > -1) {
                    var v14 int16 = 1
                    _ = v14
                }
                _ = v13
            }
            _ = fi11
        }
        v6 = 4294967295
        for fi15 := 0; fi15 < 1; fi15++ {
            switch "hello" {
            case "a":
                v2 = (v2 >> ((968) & 63))
            default:
                if (true && (("hello" != "\n\t") || true)) {
                    v5 = v5
                    v16 := GenS2{F0: 2147483647}
                    _ = v16
                } else {
                    if false {
                        v17 := GenS2{F0: (v6 % ((v6) | 1))}
                        v5 = v5
                        _ = v17
                    } else {
                        if true {
                            delete(v4, (-(fi15)))
                            v18 := []int{v6, fi15, (fi15 * (^(fi15))), (v6 % (((fi15 * int(v3))) | 1))}
                            _ = v18
                        } else {
                            var v19 int8 = v1
                            for fi20 := 0; fi20 < 16; fi20++ {
                                var v21 int8 = v19
                                _ = fi20
                                _ = v21
                            }
                            v2 -= (v2 % ((2147483647) | 1))
                            _ = v19
                        }
                    }
                    if true {
                        v2 += (v2 * v2)
                        fmt.Println(v1, v2)
                        delete(v4, (v6 % (((fi15 * fi15)) | 1)))
                    }
                }
            }
            _ = fi15
        }
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint16 = 32767
        _ = fi22
        _ = v23
    }
    v24 := v4[(^(v6))]
    v25 := GenS3{F0: 9223372036854775807, F1: 35, F2: (v3 & v3)}
    v24 += 2
    v26 := []int{}
    v2 = v2
    if true {
        v27 := v4[v6]
        _ = v27
    } else {
        v28 := GenS4{F0: v6, F1: true, F2: (v3 & v3)}
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        fmt.Println(v3, v5)
        var v30 float64 = 1e10
        var v31 string = ""
        _ = fi29
        _ = v30
        _ = v31
    }
    _ = len(v4)
    for fi32 := 0; fi32 < 18; fi32++ {
        fmt.Println(v3, v6, v1)
        v33 := GenS5{F0: 2, F1: 32767}
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
