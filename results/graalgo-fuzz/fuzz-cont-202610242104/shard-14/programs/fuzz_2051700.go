package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int8; F2 uint32 }
type GenS5 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 int8 = -128
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    if (!(1 != 42)) {
        v5 := GenS1{F0: -2147483648, F1: 2}
        _ = v5
    }
    var v6 int64 = 418
    switch -339 {
    case 290:
        switch "" {
        case "foo bar":
            v6 = (^(v6))
        case "a":
            fmt.Println(v6, v6)
        }
    case 128:
        v7 := GenS1{F0: 1, F1: 1}
        _ = v7
    case 0:
        for fi8 := 0; fi8 < 14; fi8++ {
            v1 = 23
            v6 = (v6 ^ -9223372036854775808)
            v9 := GenS2{F0: 65535, F1: 0, F2: v1}
            _ = fi8
            _ = v9
        }
    default:
        switch ("" + "a") {
        case "foo bar":
            if (false || (0 < 639)) {
                var v10 uint64 = 18446744073709551615
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 uint8 = 0
                    v10 = v10
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                if false {
                    _ = len(v3)
                } else {
                    v6 += 42
                    v1 = (^(v1))
                }
                var v13 bool = ((-1 >= 42) || (3.14159 < 0.0))
                v13 = v13
                _ = v13
            }
        case "":
            var v14 string = "hello"
            _ = v14
        }
    }
    v15 := GenS1{F0: 0, F1: 203}
    if true {
        var v16 uint = 42
        var v17 uint64 = 511
        _ = v16
        _ = v17
    }
    v6 = (v6 % (((-(v6))) | 1))
    v6 += (v6 + (-(v6)))
    for fi18 := 0; fi18 < 7; fi18++ {
        if ((("\n\t" + ("" + "foo bar")) != "\n\t") || false) {
            v6 = (-(v6))
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 float64 = (1e10 - 1e10)
                if (!(1 != -32768)) {
                    fmt.Println(v6, fi19)
                    v21 := GenS2{F0: 18446744073709551615, F1: 128, F2: v1}
                    v22 := GenS3{F0: 128}
                    _ = v21
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
            v23 := []int{}
            _ = v23
        } else {
            var v24 int16 = -1
            var v25 int16 = (v24 & (v24 / (((v24 << ((4294967295) & 63))) | 1)))
            switch "\n\t" {
            case "\n\t":
                v1 -= (v1 % ((v1) | 1))
            case "":
                var v26 uint16 = 255
                _ = v26
            case "hello":
                if true {
                    v27 := []int{(fi18 ^ (fi18 % (((-(fi18))) | 1))), fi18, (fi18 * (fi18 / ((-1) | 1)))}
                    v25 = v25
                    if false {
                        _ = len(v27)
                    }
                    _ = v27
                } else {
                    var v28 uint64 = 4294967295
                    _ = v28
                }
            default:
                var v29 string = "foo bar"
                _ = v29
            }
            _ = v24
            _ = v25
        }
        if (255 >= 0) {
            fmt.Println(v1, v6, v6)
        } else {
            fmt.Println(v1, v1, v1)
            if (!false) {
                var v30 string = "a"
                for fi31 := 0; fi31 < 2; fi31++ {
                    var v32 int32 = 0
                    v6 -= -1
                    _ = fi31
                    _ = v32
                }
                if (-2.25 < -2.25) {
                    var v33 bool = true
                    v34 := GenS1{F0: 1, F1: 255}
                    switch (((v30 + "foo bar") + v30) + "a") {
                    case "foo bar":
                        if (!v33) {
                            v35 := v4[(fi18 & (fi18 * fi18))]
                            _ = v35
                        }
                    default:
                        v36 := GenS4{F0: v1, F1: (v1 - v1), F2: 647}
                        _ = v36
                    }
                    _ = v33
                    _ = v34
                } else {
                    var v37 uint16 = 42
                    _ = v37
                }
                _ = v30
            } else {
                fmt.Println(fi18, fi18)
                switch (fi18 & 128) {
                case 128:
                    v38 := GenS5{F0: 255, F1: 2147483648}
                    _ = v38
                }
                switch fi18 {
                case -2147483648:
                    fmt.Println(v6, v6)
                default:
                    v6 = 510
                }
            }
            var v39 uint = 65535
            _ = v39
        }
        _ = fi18
    }
    v40 := []int{}
    var v41 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v6
}
