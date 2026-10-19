package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int; F1 uint8; F2 int8 }
func genFn0(p0 uint8, p1 bool, p2 bool) int64 {
    v6 := GenS1{F0: 128, F1: 2147483647}
    if (1 <= 1) {
        var v7 int32 = -1
        _ = v7
    }
    switch "" {
    case "hello":
        v8 := GenS1{F0: p0, F1: 150}
        _ = v8
    case "foo bar":
        p2 = ((p0 & p0) <= p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 680
}

func main() {
    var v1 uint64 = 608
    v2 := []string{"", "foo bar", "hello", (("a" + "foo bar") + "hello")}
    v3 := make(map[int]uint32)
    var v4 float64 = 1e10
    v5 := make(map[int]string)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v9 int16 = 42
    v10 := GenS2{F0: 255, F1: 206}
    fmt.Println(v9, v1)
    delete(v3, 0)
    var v11 uint16 = 42
    v11 = 1
    switch -474 {
    case 9223372036854775807:
        if false {
            fmt.Println(v11)
        }
    default:
        v11 -= 2
    }
    v11 += (v11 | (v11 | v11))
    switch 1 {
    case 4294967295:
        fmt.Println(v4)
    }
    var v12 int8 = 42
    v9 = -1
    for fi13 := 0; fi13 < 19; fi13++ {
        switch "hello" {
        case "a":
            var v14 uint = 1
            _ = v14
        default:
            v4 -= 1e10
        }
        var v15 int8 = (v12 * (-(v12)))
        if (42 == 42) {
            switch (fi13 ^ -266) {
            case -9223372036854775808:
                for fi16 := 0; fi16 < 13; fi16++ {
                    switch 2147483647 {
                    case 127:
                        if true {
                            v17 := GenS1{F0: 42, F1: 1}
                            v18 := GenS3{F0: fi16, F1: 128, F2: -128}
                            fmt.Println(v9)
                            v9 += 1000
                            v1 = (v1 + (v1 & (^(v1))))
                            v15 += 1
                            if (!((-(v12)) < v12)) {
                                var v19 int = (fi16 | fi13)
                                fmt.Println(v4, v4)
                                var v20 uint16 = 487
                                var v21 int16 = (v9 * 1000)
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            _ = v17
                            _ = v18
                        } else {
                            v12 += v15
                        }
                    default:
                        var v22 int8 = v15
                        _ = v22
                    }
                    _ = fi16
                }
            case -1:
                v12 -= 74
            }
            v23 := GenS1{F0: 42, F1: 516}
            v24 := v5[fi13]
            _ = v23
            _ = v24
        } else {
            var v25 uint32 = 0
            var v26 uint16 = (v11 & (v11 & 1))
            if false {
                fmt.Println(v11)
                for fi27 := 0; fi27 < 2; fi27++ {
                    if (!(false && (v26 <= 32767))) {
                        v12 = 84
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            switch (fi13 * -1) {
                            case 538:
                                v25 += 2147483647
                            default:
                                v15 = 1
                            }
                        case "a":
                            v2 = append(v2, "a")
                        case "hello":
                            switch 0 {
                            case 9223372036854775807:
                                v15 -= (v15 * 100)
                            default:
                                if (!false) {
                                    v2 = append(v2, ("foo bar" + "\n\t"))
                                    v4 += v4
                                    switch "foo bar" {
                                    case "hello":
                                        var v28 int8 = v15
                                        _ = v28
                                    }
                                    var v29 bool = (1 >= 0)
                                    var v30 uint32 = 0
                                    _ = v29
                                    _ = v30
                                }
                            }
                        }
                    }
                    v31 := GenS3{F0: (fi27 * fi13), F1: 129, F2: (v15 % (((v15 * (v12 / (((v12 / (((v12 + (v12 | (v12 + (-(v15)))))) | 1))) | 1)))) | 1))}
                    v9 = 32767
                    _ = fi27
                    _ = v31
                }
            } else {
                if (!(("hello" + "hello") == "")) {
                    var v32 uint32 = (v25 & (v25 << ((820) & 63)))
                    _ = v32
                }
                v33 := GenS1{F0: 42, F1: 335}
                _ = v33
            }
            _ = v25
            _ = v26
        }
        _ = fi13
        _ = v15
    }
    var v34 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
