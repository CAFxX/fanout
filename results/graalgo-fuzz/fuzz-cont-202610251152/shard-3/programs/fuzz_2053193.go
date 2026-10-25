package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 string; F2 int8 }
type GenS5 struct { F0 int }
func genFn0(p0 uint) uint8 {
    fmt.Println(p0)
    if (32767 < 1000) {
        if true {
            p0 = p0
            var v5 float32 = (-2.25 / 1.0)
            switch -2147483648 {
            case 4294967295:
                var v6 int = 1
                _ = v6
            }
            _ = v5
        }
        var v7 int64 = -1
        for fi8 := 0; fi8 < 7; fi8++ {
            switch "" {
            case "hello":
                var v9 int16 = -888
                _ = v9
            case "a":
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 int64 = v7
                    if true {
                        if false {
                            if ((!true) && true) {
                                v7 += (-(v7))
                            } else {
                                p0 -= (p0 % ((p0) | 1))
                                if false {
                                    v12 := GenS1{F0: "\n\t", F1: fi10}
                                    v7 += v11
                                    _ = v12
                                } else {
                                    var v13 uint8 = 255
                                    v14 := []int{(fi8 - fi8), 924, fi10, fi8}
                                    v15 := GenS1{F0: "hello", F1: fi10}
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                }
                            }
                        }
                        v16 := GenS2{F0: v7, F1: 18446744073709551615, F2: 4294967295}
                        v17 := []int{fi8}
                        var v18 int16 = 42
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        v7 = (v7 + v11)
                        v19 := []int{(fi8 + fi8), (fi10 + (fi10 + 32768)), (fi10 << (((p0 ^ (p0 + 244))) & 63)), fi10}
                        _ = v19
                    }
                    fmt.Println(v7, fi10)
                    _ = fi10
                    _ = v11
                }
            case "\n\t":
                v7 = -1
            }
            v20 := GenS3{F0: 4294967295}
            _ = fi8
            _ = v20
        }
        _ = v7
    } else {
        if false {
            var v21 int = 4294967295
            _ = v21
        }
        v22 := []int{-766, 9223372036854775807, 781}
        var v23 int = -855
        _ = v22
        _ = v23
    }
    var v24 uint = (p0 + p0)
    _ = p0
    _ = v24
    return 82
}
func genFn1(p0 int16) int8 {
    p0 = -1
    if true {
        var v25 uint32 = 2147483647
        fmt.Println(p0, p0)
        _ = v25
    } else {
        v26 := GenS2{F0: -9223372036854775808, F1: 2, F2: 65535}
        _ = v26
    }
    _ = p0
    return 92
}

func main() {
    var v1 uint32 = 636
    var v2 bool = false
    v3 := []int{-107, -717}
    var v4 float64 = float64(-2.25)
    _ = len(v3)
    switch ("" + "\n\t") {
    case "\n\t":
        var v27 uint32 = (^(v1))
        _ = v27
    }
    v1 = (v1 ^ (v1 + 2147483648))
    v2 = (192 > genFn0(192))
    v1 -= (v1 + v1)
    var v28 int8 = -128
    switch 32767 {
    case 2147483647:
        v3 = append(v3, 1)
    case -1:
        v2 = ((v2 || (false || (v1 < (v1 * v1)))) || true)
    }
    v29 := GenS3{F0: 917}
    v28 = 42
    v28 -= (v28 ^ -1)
    v4 += 1e10
    v30 := GenS4{F0: 42, F1: ("foo bar" + "hello"), F2: v28}
    var v31 int8 = v28
    switch "foo bar" {
    case "":
        for fi32 := 0; fi32 < 18; fi32++ {
            for fi33 := 0; fi33 < 18; fi33++ {
                if v2 {
                    fmt.Println(v4, v31, fi32)
                    var v34 float64 = 1e10
                    v35 := GenS5{F0: fi32}
                    _ = v34
                    _ = v35
                } else {
                    if (609 < 264) {
                        var v36 int16 = 1000
                        _ = v36
                    } else {
                        if (-9223372036854775808 > -9223372036854775808) {
                            fmt.Println(fi33, v1)
                            v4 = v4
                        }
                    }
                    v37 := GenS3{F0: 1}
                    _ = v37
                }
                fmt.Println(v1, v4, fi32)
                _ = fi33
            }
            _ = fi32
        }
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        if (42 == 257) {
            for fi38 := 0; fi38 < 18; fi38++ {
                var v39 uint16 = 65535
                v31 -= (v31 << ((1) & 63))
                _ = fi38
                _ = v39
            }
        } else {
            for fi40 := 0; fi40 < 8; fi40++ {
                var v41 bool = (!false)
                var v42 int = fi40
                _ = fi40
                _ = v41
                _ = v42
            }
            v1 = (v1 >> ((0) & 63))
        }
    default:
        for fi43 := 0; fi43 < 13; fi43++ {
            v31 = (^(v28))
            for fi44 := 0; fi44 < 2; fi44++ {
                v3 = append(v3, fi43)
                v3 = append(v3, (fi44 >> ((42) & 63)))
                var v45 string = "a"
                _ = fi44
                _ = v45
            }
            _ = fi43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
