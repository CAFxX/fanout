package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint32; F2 uint8 }
func genFn0(p0 uint) bool {
    var v4 float32 = -2.25
    p0 -= (p0 & p0)
    var v5 bool = (42 <= 2)
    var v6 bool = v5
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return ((0.1 >= -2.25) && (!false))
}
func genFn1(p0 int8, p1 int16) int {
    v7 := []int{236}
    switch "\n\t" {
    case "\n\t":
        var v8 float32 = (1.5 - (float32(1.5) * -2.25))
        _ = v8
    case "":
        var v9 int = -438
        _ = v9
    case "a":
        if ("foo bar" == "hello") {
            v10 := GenS1{F0: -755}
            switch "" {
            case "a":
                for fi11 := 0; fi11 < 17; fi11++ {
                    switch -1 {
                    case 128:
                        var v12 string = "foo bar"
                        _ = v12
                    case -9223372036854775808:
                        var v13 uint32 = 42
                        _ = v13
                    case 487:
                        fmt.Println(p0, fi11)
                    default:
                        var v14 int64 = -9223372036854775808
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 5; fi15++ {
                        _ = len(v7)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        switch "" {
                        case "\n\t":
                            p0 -= 0
                        case "":
                            if false {
                                var v16 uint16 = 255
                                var v17 uint16 = (v16 + v16)
                                _ = v16
                                _ = v17
                            } else {
                                fmt.Println(p1, p0, fi15)
                            }
                        default:
                            if true {
                                p0 += (p0 & p0)
                                var v18 uint16 = 2
                                p1 -= (^(p1))
                                var v19 float32 = float32(float64(((((1.5 - 0.0) + 1e10) + 0.0) / 1e10)))
                                _ = v18
                                _ = v19
                            }
                        }
                        _ = fi15
                    }
                    _ = fi11
                }
            }
            p1 = (-(p1))
            _ = v10
        } else {
            p0 = 127
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint8 = 112
    var v2 uint8 = 194
    var v3 uint64 = 4294967295
    if (!(1 == -128)) {
        var v20 string = ""
        _ = v20
    } else {
        for fi21 := 0; fi21 < 13; fi21++ {
            switch "" {
            case "\n\t":
                v22 := GenS1{F0: -2147483648}
                _ = v22
            case "foo bar":
                for fi23 := 0; fi23 < 5; fi23++ {
                    if false {
                        fmt.Println(v1)
                        v2 = (^(v1))
                        if false {
                            var v24 uint = 874
                            v1 = (v1 << ((771) & 63))
                            _ = v24
                        } else {
                            v2 = 1
                        }
                    }
                    _ = fi23
                }
            case "":
                switch "hello" {
                case "\n\t":
                    v3 += (v3 % (((v3 + (^(v3)))) | 1))
                case "a":
                    fmt.Println(fi21, v1, v1)
                default:
                    fmt.Println(fi21)
                }
            default:
                v1 += (v1 ^ v2)
            }
            v1 -= (^(v2))
            var v25 uint8 = 1
            _ = fi21
            _ = v25
        }
        if false {
            var v26 string = (("foo bar" + ("hello" + "hello")) + "a")
            _ = v26
        }
    }
    v27 := []int{-2147483648}
    var v28 uint32 = 555
    for fi29 := 0; fi29 < 5; fi29++ {
        v27 = append(v27, (fi29 & genFn1(127, -32768)))
        _ = fi29
    }
    if len(v27) > 0 {
        v27[0] = 0
    }
    var v30 uint8 = 255
    var v31 int32 = int32(v3)
    var v32 int16 = 0
    v30 -= (^(v2))
    v33 := GenS2{F0: 782, F1: 1, F2: v30}
    var v34 uint8 = (v2 ^ v1)
    v34 -= uint8(v28)
    v27 = append(v27, 2147483647)
    _ = len(v27)
    v35 := GenS2{F0: -9223372036854775808, F1: 438, F2: 128}
    var v36 float32 = 0.1
    if (42 == v2) {
        for fi37 := 0; fi37 < 17; fi37++ {
            v38 := GenS1{F0: -910}
            _ = fi37
            _ = v38
        }
    }
    switch ("foo bar" + "") {
    case "hello":
        var v39 uint64 = v3
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
