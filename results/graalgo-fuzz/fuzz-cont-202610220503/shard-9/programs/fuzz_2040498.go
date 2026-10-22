package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0(p0 int32) int32 {
    p0 = 2147483647
    if true {
        v7 := GenS1{F0: 1}
        if (4294967295 != 2147483648) {
            v8 := GenS1{F0: 42}
            if (!false) {
                if (((p0 + (p0 & p0)) < p0) || (-9223372036854775808 >= 0)) {
                    if (!(128 < 255)) {
                        var v9 int32 = 1
                        v10 := []int{-1, 128, -673, 128}
                        if false {
                            v10 = append(v10, 156)
                            if (!true) {
                                v9 = (-(v9))
                            } else {
                                for fi11 := 0; fi11 < 7; fi11++ {
                                    if len(v10) > 0 {
                                        v10[0] = (fi11 + (fi11 & (-(fi11))))
                                    }
                                    _ = fi11
                                }
                            }
                            v12 := GenS2{F0: 18446744073709551615, F1: 255, F2: 819}
                            _ = v12
                        }
                        _ = v10
                        _ = v9
                    } else {
                        var v13 int64 = 0
                        _ = v13
                    }
                } else {
                    var v14 int = 0
                    _ = v14
                }
            }
            fmt.Println(p0)
            _ = v8
        }
        p0 = p0
        _ = v7
    }
    p0 = p0
    v15 := []int{-2147483648, 127}
    for fi16 := 0; fi16 < 19; fi16++ {
        if true {
            if len(v15) > 0 {
                v15[0] = -477
            }
        } else {
            switch (fi16 & fi16) {
            case -2147483648:
                switch fi16 {
                case 247:
                    if true {
                        fmt.Println(p0, fi16)
                        v17 := GenS3{F0: 998, F1: 42}
                        _ = v17
                    } else {
                        var v18 float32 = 1.5
                        v18 = v18
                        p0 = (p0 - (p0 | 42))
                        _ = v18
                    }
                default:
                    v19 := GenS1{F0: -128}
                    _ = v19
                }
            default:
                var v20 uint = 530
                _ = v20
            }
        }
        if (32767 <= 1000) {
            var v21 int = (fi16 / ((fi16) | 1))
            _ = v21
        }
        _ = fi16
    }
    _ = p0
    _ = v15
    return (p0 + p0)
}

func main() {
    var v1 float64 = ((0.0 * ((0.1 / (3.14159 - 1.5)) * 0.0)) + 0.1)
    var v2 float64 = ((v1 - v1) / 1.0)
    v3 := []uint32{42}
    v4 := []uint32{177, 952, 290}
    v5 := []string{("a" + ""), ("hello" + "foo bar")}
    var v6 uint32 = 459
    fmt.Println(v2, v1, v1)
    v6 = 2147483648
    fmt.Println(v6)
    var v22 uint = 144
    if true {
        var v23 int8 = -1
        var v24 uint = v22
        _ = v23
        _ = v24
    } else {
        v2 = (float64(-2.25) / v1)
        var v25 string = ""
        _ = v25
    }
    switch "a" {
    case "a":
        for fi26 := 0; fi26 < 2; fi26++ {
            fmt.Println(fi26)
            _ = fi26
        }
    default:
        v1 = (((v1 * (v2 / 3.14159)) * v1) + v2)
    }
    fmt.Println(v22)
    v2 = ((3.14159 / v2) - 0.1)
    for fi27 := 0; fi27 < 3; fi27++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v28 uint64 = 249
        _ = fi27
        _ = v28
    }
    switch (("\n\t" + "foo bar") + (("\n\t" + "foo bar") + ("" + "a"))) {
    case "hello":
        v2 += (v2 * v2)
    case "\n\t":
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 int32 = 95
            for fi31 := 0; fi31 < 2; fi31++ {
                v6 = v6
                var v32 int16 = 1
                _ = fi31
                _ = v32
            }
            _ = fi29
            _ = v30
        }
    }
    fmt.Println(v22)
    var v33 int32 = 479
    if true {
        v22 = v22
        for fi34 := 0; fi34 < 20; fi34++ {
            v2 = v1
            v33 = 42
            switch "a" {
            case "foo bar":
                var v35 int = fi34
                _ = v35
            case "a":
                if (255 < 48) {
                    v2 = v1
                }
            case "":
                var v36 float64 = 1.5
                _ = v36
            }
            _ = fi34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
