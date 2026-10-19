package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 uint8; F1 uint64; F2 int }
func genFn0(p0 int32, p1 int32) uint64 {
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint16 = 763
        var v9 int32 = (p0 & p0)
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 += -1
    _ = p0
    _ = p1
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 int, p1 uint64, p2 bool) int {
    if p2 {
        v10 := GenS1{F0: 2147483647, F1: 908, F2: 2}
        var v11 int16 = -32768
        v12 := []int{p0, p0, (p0 ^ (p0 - (p0 + 42))), 0}
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint16 = 1
        v14 += v14
        var v15 int16 = -1
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 bool = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -693
}

func main() {
    var v1 uint = 0
    var v2 int8 = 127
    var v3 uint32 = 2147483647
    var v4 uint = 255
    var v5 int = 128
    switch "a" {
    case "\n\t":
        v17 := []int{}
        _ = v17
    case "a":
        switch (v5 + v5) {
        case 9223372036854775807:
            for fi18 := 0; fi18 < 14; fi18++ {
                v5 = 2147483647
                _ = fi18
            }
        }
    default:
        for fi19 := 0; fi19 < 6; fi19++ {
            v5 = (^(v5))
            v20 := GenS1{F0: -2147483648, F1: v1, F2: v3}
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(v5, v5, v1)
    v21 := GenS2{F0: -1, F1: -1}
    v22 := GenS3{F0: (v2 % (((-(v2))) | 1)), F1: (42 <= 507)}
    v3 = v3
    v23 := GenS4{F0: v3, F1: -1}
    if true {
        v24 := GenS5{F0: 42}
        var v25 uint16 = 65535
        _ = v24
        _ = v25
    } else {
        v5 += (-(v5))
    }
    switch "a" {
    case "a":
        v26 := GenS2{F0: 841, F1: -9223372036854775808}
        _ = v26
    case "foo bar":
        for fi27 := 0; fi27 < 14; fi27++ {
            var v28 uint = v4
            for fi29 := 0; fi29 < 14; fi29++ {
                switch fi27 {
                case 127:
                    v30 := []int{fi29, (fi29 ^ v5)}
                    _ = v30
                case 128:
                    fmt.Println(v3)
                }
                v1 += (v28 ^ v28)
                _ = fi29
            }
            _ = fi27
            _ = v28
        }
    case "":
        v3 -= (v3 % (((v3 << ((v4) & 63))) | 1))
    default:
        for fi31 := 0; fi31 < 8; fi31++ {
            if true {
                switch ("hello" + "hello") {
                case "":
                    if false {
                        if true {
                            for fi32 := 0; fi32 < 11; fi32++ {
                                if false {
                                    v3 = v3
                                } else {
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v4, fi31)
                                    case "foo bar":
                                        var v33 float32 = 1e10
                                        _ = v33
                                    default:
                                        v3 = 609
                                    }
                                }
                                v34 := []int{2147483647, -2147483648, (fi31 & 128), fi32}
                                v1 -= v1
                                _ = fi32
                                _ = v34
                            }
                        } else {
                            fmt.Println(fi31)
                            var v35 uint = 0
                            _ = v35
                        }
                        if (32767 < 687) {
                            v36 := GenS6{F0: v2, F1: (("" + "") + "a")}
                            _ = v36
                        }
                        fmt.Println(v1, v1)
                    } else {
                        var v37 uint32 = v3
                        _ = v37
                    }
                }
                switch (fi31 % (((v5 >> (((v4 + (v4 ^ 401))) & 63))) | 1)) {
                case 127:
                    v38 := GenS4{F0: 2, F1: -451}
                    _ = v38
                case 32768:
                    v39 := GenS3{F0: (^(v2)), F1: false}
                    _ = v39
                case -2147483648:
                    fmt.Println(v4)
                default:
                    fmt.Println(v3, fi31)
                }
            }
            _ = fi31
        }
    }
    v40 := []int{(v5 << ((v4) & 63))}
    v41 := GenS3{F0: -71, F1: false}
    switch 4294967295 {
    case -1:
        v2 -= v2
    }
    v42 := GenS7{F0: 42, F1: 42, F2: v5}
    var v43 uint = (v1 - (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
