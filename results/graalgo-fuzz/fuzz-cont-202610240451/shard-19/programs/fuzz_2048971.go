package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
func genFn0(p0 int64, p1 string) int32 {
    var v4 uint16 = 255
    if (true && false) {
        v5 := GenS1{F0: (v4 + 0), F1: 31, F2: 213}
        _ = v5
    } else {
        v6 := GenS1{F0: v4, F1: v4, F2: 0}
        v7 := GenS1{F0: (v4 + v4), F1: 65535, F2: 121}
        switch "foo bar" {
        case "a":
            v8 := GenS2{F0: p0}
            _ = v8
        case "foo bar":
            v4 -= 2
        case "":
            v9 := GenS2{F0: p0}
            _ = v9
        }
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return 2147483647
}

func main() {
    var v1 uint64 = 960
    var v2 uint32 = 2147483647
    var v3 float64 = (((float64(1.5) - (0.0 * 0.1)) - 1e10) / 1.0)
    v10 := GenS3{F0: genFn0(-1, "\n\t"), F1: 1000}
    var v11 int32 = -1
    var v12 uint16 = uint16(v1)
    v2 = 2
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(v12, fi13, v3)
            _ = fi14
        }
        _ = fi13
    }
    if true {
        v15 := []int{}
        v12 += 179
        _ = v15
    }
    switch -9223372036854775808 {
    case 127:
        switch "hello" {
        case "\n\t":
            if false {
                var v16 bool = ((!(0 == (v2 % (((v2 - v2)) | 1)))) && true)
                v17 := []int{127, 127, -675}
                _ = v16
                _ = v17
            } else {
                switch -9223372036854775808 {
                case 964:
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v19 := []int{fi18}
                        var v20 uint16 = (v12 - (v12 ^ v12))
                        var v21 int16 = 42
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                case 903:
                    switch (("\n\t" + "\n\t") + "\n\t") {
                    case "\n\t":
                        if ((false && false) && false) {
                            v11 += 2147483647
                        } else {
                            if true {
                                v2 = 2147483648
                                switch 9223372036854775807 {
                                case -632:
                                    if true {
                                        v22 := GenS4{F0: 1}
                                        fmt.Println(v11, v12, v12)
                                        _ = v22
                                    }
                                case -9223372036854775808:
                                    for fi23 := 0; fi23 < 3; fi23++ {
                                        v3 = ((0.0 / v3) - float64((-2.25 * 3.14159)))
                                        _ = fi23
                                    }
                                }
                                if true {
                                    v12 = (^(v12))
                                    if true {
                                        fmt.Println(v2, v1, v12)
                                        v2 = (^(v2))
                                    } else {
                                        var v24 uint = 0
                                        switch -327 {
                                        case -768:
                                            v25 := GenS5{F0: 32767}
                                            _ = v25
                                        case 159:
                                            if (((v11 >> (((v24 / (((v24 % ((v24) | 1))) | 1))) & 63)) > (v11 * v11)) || (42 <= 77)) {
                                                v26 := []int{-678, 9223372036854775807, 32768}
                                                _ = v26
                                            }
                                        case 128:
                                            var v27 int = 127
                                            _ = v27
                                        default:
                                            v28 := []int{32768, -746}
                                            _ = v28
                                        }
                                        _ = v24
                                    }
                                    v1 = (^(v1))
                                }
                            }
                            v29 := GenS1{F0: v12, F1: 32767, F2: 126}
                            _ = v29
                        }
                    case "hello":
                        fmt.Println(v3, v2, v1)
                    }
                case 4294967295:
                    v2 = 42
                }
                var v30 float64 = (3.14159 * v3)
                _ = v30
            }
        case "hello":
            v31 := []int{127}
            _ = v31
        case "":
            v32 := GenS6{F0: 9223372036854775807}
            _ = v32
        }
    case 128:
        var v33 float64 = v3
        _ = v33
    }
    var v34 int8 = 1
    var v35 int32 = (v11 | (v11 ^ (v11 & v11)))
    v36 := []int{-1}
    fmt.Println(v11)
    switch 32768 {
    case -875:
        if (0 == 1000) {
            fmt.Println(v11, v12)
            for fi37 := 0; fi37 < 5; fi37++ {
                v34 += 42
                var v38 bool = ((fi37 <= 899) && true)
                fmt.Println(v12, fi37, v3)
                _ = fi37
                _ = v38
            }
            var v39 bool = (381 != 9223372036854775807)
            _ = v39
        }
    case 1:
        var v40 uint = 379
        _ = v40
    }
    v3 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
