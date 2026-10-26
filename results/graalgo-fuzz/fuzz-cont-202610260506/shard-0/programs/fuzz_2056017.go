package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint8) int8 {
    if false {
        fmt.Println(p0)
    } else {
        p0 += p0
    }
    p0 -= p0
    for fi7 := 0; fi7 < 8; fi7++ {
        if (fi7 < fi7) {
            switch "" {
            case "\n\t":
                for fi8 := 0; fi8 < 19; fi8++ {
                    var v9 uint8 = (p0 << ((0) & 63))
                    var v10 uint8 = 0
                    switch "foo bar" {
                    case "a":
                        var v11 float64 = 0.1
                        _ = v11
                    default:
                        v12 := []int{}
                        _ = v12
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case "":
                if (-206 == -2147483648) {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        v14 := []int{(fi7 & 32767), (^(fi7)), 1}
                        p0 += (^(p0))
                        switch "" {
                        case "":
                            if true {
                                p0 -= 128
                                var v15 uint8 = (p0 | (p0 + p0))
                                fmt.Println(p0)
                                _ = v15
                            } else {
                                fmt.Println(fi7, p0, p0)
                                p0 = (^(p0))
                                var v16 int64 = 660
                                _ = v16
                            }
                        case "\n\t":
                            var v17 uint64 = 2
                            _ = v17
                        case "a":
                            p0 += p0
                        }
                        _ = fi13
                        _ = v14
                    }
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 uint64 = 0
                        v19 = (^(v19))
                        _ = fi18
                        _ = v19
                    }
                    fmt.Println(fi7, fi7)
                }
            default:
                var v20 uint32 = 901
                _ = v20
            }
        } else {
            p0 += p0
            fmt.Println(p0, p0)
        }
        _ = fi7
    }
    p0 += (^(p0))
    switch ("" + "\n\t") {
    case "a":
        for fi21 := 0; fi21 < 9; fi21++ {
            v22 := GenS1{F0: 2}
            v23 := GenS1{F0: 18446744073709551615}
            v24 := GenS1{F0: 18446744073709551615}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    case "foo bar":
        p0 -= (p0 + 80)
    case "\n\t":
        v25 := GenS2{F0: "a"}
        _ = v25
    }
    _ = p0
    return 100
}
func genFn1(p0 int32) int {
    if ((2147483647 > -1) || false) {
        switch "a" {
        case "a":
            fmt.Println(p0)
        case "\n\t":
            if true {
                for fi26 := 0; fi26 < 13; fi26++ {
                    fmt.Println(p0, fi26)
                    _ = fi26
                }
                for fi27 := 0; fi27 < 14; fi27++ {
                    v28 := GenS1{F0: 572}
                    v29 := []int{fi27, fi27, fi27, fi27}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            } else {
                v30 := GenS3{F0: p0, F1: 42}
                _ = v30
            }
        }
    }
    p0 = (p0 | (p0 / ((p0) | 1)))
    p0 = (p0 * 649)
    p0 += (p0 + (p0 - (p0 << ((643) & 63))))
    _ = p0
    return 127
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 324
    v3 := []uint32{42}
    var v4 uint32 = 896
    v5 := []int{-2147483648}
    var v6 uint32 = (v4 * (v4 - v4))
    var v31 int8 = -56
    var v32 string = ("" + "")
    for fi33 := 0; fi33 < 5; fi33++ {
        v6 += v4
        v4 -= v6
        _ = fi33
    }
    var v34 int32 = int32(v6)
    var v35 int32 = 440
    v32 = v32
    fmt.Println(v31)
    var v36 bool = (false && ((^(v4)) > v6))
    var v37 int16 = -32768
    var v38 float32 = float32(0.1)
    fmt.Println(v32, v37)
    var v39 uint32 = 793
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 uint16 = v2
        for fi42 := 0; fi42 < 8; fi42++ {
            var v43 int64 = 9223372036854775807
            _ = fi42
            _ = v43
        }
        var v44 int32 = (^(v35))
        _ = fi40
        _ = v41
        _ = v44
    }
    var v45 int8 = genFn0(22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
