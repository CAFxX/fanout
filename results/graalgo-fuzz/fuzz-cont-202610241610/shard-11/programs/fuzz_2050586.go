package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint; F1 int; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 string; F2 int64 }
type GenS6 struct { F0 uint64; F1 int64 }
func genFn0(p0 int64, p1 string) uint64 {
    p0 = p0
    v5 := GenS1{F0: 110, F1: 0}
    if (-128 > 100) {
        if true {
            if false {
                p1 = p1
                if false {
                    var v6 uint = 65535
                    switch p1 {
                    case "foo bar":
                        if (!true) {
                            var v7 uint64 = 1
                            for fi8 := 0; fi8 < 9; fi8++ {
                                switch fi8 {
                                case 128:
                                    for fi9 := 0; fi9 < 6; fi9++ {
                                        v10 := GenS2{F0: (v6 >> ((v6) & 63)), F1: fi8, F2: 42}
                                        v6 += 65535
                                        switch (p1 + ("\n\t" + p1)) {
                                        case "hello":
                                            v11 := []int{fi8, (^(fi9)), (^(fi8))}
                                            _ = v11
                                        }
                                        var v12 string = p1
                                        v13 := GenS3{F0: v7}
                                        _ = fi9
                                        _ = v10
                                        _ = v12
                                        _ = v13
                                    }
                                case 417:
                                    var v14 int32 = 637
                                    _ = v14
                                default:
                                    if false {
                                        for fi15 := 0; fi15 < 11; fi15++ {
                                            v16 := []int{(fi15 >> ((v6) & 63)), fi15}
                                            _ = fi15
                                            _ = v16
                                        }
                                    } else {
                                        if false {
                                            v7 += (^(v7))
                                            var v17 int64 = p0
                                            p1 = p1
                                            _ = v17
                                        }
                                    }
                                }
                                _ = fi8
                            }
                            _ = v7
                        } else {
                            var v18 uint = v6
                            v19 := GenS1{F0: 0, F1: 98}
                            _ = v18
                            _ = v19
                        }
                    case "hello":
                        if (0 != 0) {
                            var v20 int8 = 100
                            switch 32767 {
                            case 726:
                                fmt.Println(v20, p1, p1)
                            default:
                                switch (p1 + p1) {
                                case "a":
                                    v21 := []int{406}
                                    _ = v21
                                case "hello":
                                    if true {
                                        v22 := []int{-9223372036854775808, -2147483648, -9223372036854775808}
                                        _ = v22
                                    }
                                }
                            }
                            _ = v20
                        } else {
                            var v23 uint8 = 136
                            var v24 uint64 = 447
                            _ = v23
                            _ = v24
                        }
                    case "":
                        v6 += (v6 + v6)
                    }
                    if true {
                        v6 -= v6
                    }
                    _ = v6
                }
                var v25 uint = 824
                _ = v25
            } else {
                fmt.Println(p0, p0)
                for fi26 := 0; fi26 < 13; fi26++ {
                    v27 := GenS4{F0: -74}
                    _ = fi26
                    _ = v27
                }
            }
        }
        var v28 int16 = -1
        _ = v28
    } else {
        for fi29 := 0; fi29 < 13; fi29++ {
            for fi30 := 0; fi30 < 19; fi30++ {
                v31 := GenS3{F0: 0}
                p0 += p0
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v5
    return 606
}
func genFn1(p0 uint8) uint32 {
    v32 := GenS5{F0: "hello", F1: "foo bar", F2: 9223372036854775807}
    v33 := GenS4{F0: -32768}
    switch "a" {
    case "":
        p0 = (p0 | (p0 / ((p0) | 1)))
    case "\n\t":
        var v34 int16 = 419
        _ = v34
    }
    fmt.Println(p0)
    _ = p0
    _ = v32
    _ = v33
    return 334
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -1
    var v3 uint64 = 0
    var v4 float32 = (3.14159 * 1e10)
    var v35 uint32 = 2147483648
    switch 9223372036854775807 {
    case 0:
        v36 := GenS2{F0: 840, F1: -282, F2: 0}
        _ = v36
    }
    var v37 uint8 = 42
    var v38 uint8 = (v37 ^ uint8(v35))
    fmt.Println(v35)
    for fi39 := 0; fi39 < 11; fi39++ {
        v3 = v3
        _ = fi39
    }
    v4 += ((v4 / v4) / (float32((((1.5 / 0.1) / 0.1) - 3.14159)) + -2.25))
    v40 := GenS6{F0: (v3 * v3), F1: 9223372036854775807}
    var v41 uint64 = genFn0(-1, (((("a" + "foo bar") + "\n\t") + ("" + "")) + ("\n\t" + "foo bar")))
    switch -9223372036854775808 {
    case 42:
        v2 = v2
    default:
        v1[49] = ("hello" + "a")
    }
    var v42 uint16 = 421
    v41 = v41
    if false {
        var v43 int = 305
        fmt.Println(v3, v38, v4)
        switch 9223372036854775807 {
        case -874:
            var v44 bool = false
            _ = v44
        case -9223372036854775808:
            var v45 uint64 = v41
            _ = v45
        default:
            var v46 uint = 2
            _ = v46
        }
        _ = v43
    }
    delete(v1, -2147483648)
    v47 := []int{int(v2), 32768}
    fmt.Println(v3, v2)
    var v48 int = 197
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v47
    _ = v48
}
