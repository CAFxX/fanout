package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 int64 }
func genFn0(p0 uint) int32 {
    switch "\n\t" {
    case "":
        for fi7 := 0; fi7 < 11; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(fi8, fi7)
                _ = fi8
            }
            var v9 int64 = -9223372036854775808
            v10 := []int{}
            _ = fi7
            _ = v10
            _ = v9
        }
    default:
        var v11 int32 = 2147483647
        _ = v11
    }
    var v12 int8 = -108
    _ = p0
    _ = v12
    return 1
}
func genFn1(p0 uint32, p1 int16) int32 {
    var v13 uint16 = 128
    if (true && (-547 <= 1)) {
        fmt.Println(p0, p1, p1)
    } else {
        if false {
            v14 := []int{42, 32768, int(v13), 211}
            p1 -= p1
            v14 = append(v14, 9223372036854775807)
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v13
    return genFn0(uint(v13))
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 string = "a"
    var v3 float64 = ((float64(3.14159) - -2.25) + 1e10)
    var v4 uint64 = 2
    v5 := make(map[string]bool)
    v6 := []string{"foo bar", "hello", ""}
    var v15 int16 = 32767
    for fi16 := 0; fi16 < 12; fi16++ {
        if (v2 > v2) {
            if (4294967295 == (fi16 / ((fi16) | 1))) {
                switch v2 {
                case "":
                    v3 = float64(-2.25)
                default:
                    if (176 < 235) {
                        fmt.Println(fi16, v4, v2)
                        for fi17 := 0; fi17 < 18; fi17++ {
                            v18 := GenS1{F0: 786, F1: -2147483648, F2: (("" + "a") + "\n\t")}
                            var v19 int32 = -81
                            v6 = append(v6, v2)
                            _ = len(v5)
                            v3 = 1.5
                            fmt.Println(v1, v15)
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        fmt.Println(v1, v4)
                    }
                }
                v2 = v2
                switch int(v15) {
                case 0:
                    for fi20 := 0; fi20 < 18; fi20++ {
                        var v21 uint32 = 2
                        _ = fi20
                        _ = v21
                    }
                }
            }
            var v22 int64 = (v1 % (((^(v1))) | 1))
            var v23 int16 = v15
            _ = v22
            _ = v23
        }
        var v24 uint = 4294967295
        _ = fi16
        _ = v24
    }
    var v25 bool = false
    switch v2 {
    case "foo bar":
        if v25 {
            v4 = 1
            v2 = "hello"
        } else {
            fmt.Println(v25, v1)
            v6 = append(v6, "foo bar")
        }
    case "\n\t":
        v26 := GenS2{F0: genFn0(2), F1: -754}
        _ = v26
    }
    v4 = (v4 % ((42) | 1))
    var v27 uint32 = 2147483647
    fmt.Println(v27, v15, v4)
    for fi28 := 0; fi28 < 10; fi28++ {
        v29 := GenS3{F0: -128}
        v30 := GenS2{F0: 800, F1: -1}
        for fi31 := 0; fi31 < 11; fi31++ {
            if ((v1 ^ (v1 >> ((2) & 63))) < v1) {
                v4 -= 1
                v2 = "hello"
                if v25 {
                    fmt.Println(fi31, fi31)
                    fmt.Println(fi28, v2, v2)
                    v4 = (^(v4))
                    _ = len(v5)
                }
            } else {
                var v32 int32 = -192
                var v33 uint64 = (v4 ^ v4)
                _ = v32
                _ = v33
            }
            switch (^(fi28)) {
            case -841:
                switch fi31 {
                case -1:
                    fmt.Println(v2, v4, v3)
                }
            case -2147483648:
                fmt.Println(v2)
                var v34 uint8 = 52
                _ = v34
            }
            _ = fi31
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    v6 = append(v6, v2)
    v1 = (^(v1))
    delete(v5, "foo bar")
    v6 = append(v6, v2)
    if (v4 > 0) {
        if true {
            var v35 bool = (!false)
            var v36 int32 = -1
            _ = v35
            _ = v36
        }
        v3 += v3
        for fi37 := 0; fi37 < 7; fi37++ {
            var v38 float32 = 0.1
            v25 = ((^(v15)) < (v15 % (((v15 << ((42) & 63))) | 1)))
            _ = fi37
            _ = v38
        }
    } else {
        v15 = v15
    }
    var v39 int16 = (^(v15))
    v6 = append(v6, v2)
    for fi40 := 0; fi40 < 9; fi40++ {
        fmt.Println(v27, v4, v39)
        var v41 int32 = 42
        _ = fi40
        _ = v41
    }
    fmt.Println(v39)
    v42 := GenS4{F0: "", F1: (v1 & (-(v1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
