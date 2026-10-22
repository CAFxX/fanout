package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint64; F1 int8; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0(p0 int) int8 {
    switch "foo bar" {
    case "\n\t":
        for fi4 := 0; fi4 < 3; fi4++ {
            fmt.Println(fi4)
            v5 := GenS1{F0: 931, F1: 794}
            _ = fi4
            _ = v5
        }
    case "foo bar":
        v6 := []int{p0, p0, 128, 128}
        _ = v6
    case "hello":
        fmt.Println(p0)
    default:
        v7 := []int{(p0 * p0), (p0 << ((42) & 63)), (p0 * 1)}
        _ = v7
    }
    fmt.Println(p0)
    var v8 string = "foo bar"
    p0 = (p0 + p0)
    _ = p0
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    v9 := []int{-9223372036854775808, 193, -9223372036854775808, 42}
    var v10 uint16 = 32767
    v10 -= 1
    var v11 int64 = 42
    switch -57 {
    case -2147483648:
        if (false || (1e10 >= 1.5)) {
            v12 := GenS1{F0: -1, F1: 612}
            v11 += (v11 | 1)
            if (false && (42 < -1)) {
                switch 128 {
                case 9223372036854775807:
                    var v13 uint64 = 2
                    _ = v13
                default:
                    v14 := GenS2{F0: 0, F1: -92, F2: -1}
                    _ = v14
                }
                for fi15 := 0; fi15 < 7; fi15++ {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        var v17 int = fi16
                        v18 := v1[("\n\t" + (((("\n\t" + (("foo bar" + "a") + ("foo bar" + (("a" + "foo bar") + "")))) + (("a" + ("" + "a")) + "\n\t")) + "\n\t") + ""))]
                        var v19 int16 = 36
                        var v20 int32 = 1
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    _ = fi15
                }
                var v21 float32 = (0.1 - -2.25)
                _ = v21
            }
            _ = v12
        } else {
            for fi22 := 0; fi22 < 14; fi22++ {
                if (3.14159 != ((float32(1.5) * -2.25) * 3.14159)) {
                    if len(v9) > 0 {
                        v9[0] = 0
                    }
                    var v23 bool = (65535 < 776)
                    fmt.Println(v11, v10, v11)
                    _ = v23
                }
                switch (fi22 ^ 32768) {
                case -22:
                    v11 -= 42
                default:
                    var v24 int = -1
                    _ = v24
                }
                _ = fi22
            }
            v9 = append(v9, -1)
        }
    case 578:
        v11 = (v11 / (((^(v11))) | 1))
    case 9223372036854775807:
        switch -2147483648 {
        case 9223372036854775807:
            for fi25 := 0; fi25 < 14; fi25++ {
                v26 := GenS1{F0: -1, F1: -153}
                v11 += (^(v11))
                _ = fi25
                _ = v26
            }
        case -2147483648:
            v11 += (^(v11))
        }
    }
    var v27 int = -785
    var v28 int32 = -1
    v11 = v11
    if (!((!(!((-2.25 / 1.0) < float32(-2.25)))) && (4294967295 != 4294967295))) {
        v29 := GenS3{F0: 3}
        _ = v29
    } else {
        switch (("" + ("" + "a")) + ("" + "hello")) {
        case "":
            var v30 float64 = (3.14159 - 3.14159)
            _ = v30
        case "hello":
            var v31 int8 = 100
            _ = v31
        case "\n\t":
            v11 += 42
        default:
            v32 := GenS1{F0: -592, F1: 1}
            _ = v32
        }
    }
    if false {
        if (127 != 127) {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            fmt.Println(v10, v27)
            fmt.Println(v11, v27, v10)
        } else {
            v11 += (-(v11))
        }
    }
    fmt.Println(v11, v28, v27)
    v33 := GenS3{F0: -128}
    v34 := GenS4{F0: (false && (!true)), F1: 0}
    if (true || ((v28 + v28) < (v28 * v28))) {
        fmt.Println(v10, v28)
        v35 := GenS1{F0: v28, F1: 42}
        switch "hello" {
        case "hello":
            var v36 bool = (false && true)
            _ = v36
        case "foo bar":
            if (((v27 * v27) >= v27) || ((!(false && (116 >= 0))) || ((v10 * v10) != v10))) {
                for fi37 := 0; fi37 < 7; fi37++ {
                    v38 := GenS2{F0: 466, F1: 127, F2: (v28 / ((664) | 1))}
                    var v39 float32 = 1.5
                    var v40 uint = 255
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
            }
        case "\n\t":
            for fi41 := 0; fi41 < 5; fi41++ {
                _ = len(v2)
                if (105 == 109) {
                    v27 = v27
                }
                _ = fi41
            }
        default:
            if len(v9) > 0 {
                v9[0] = 42
            }
        }
        _ = v35
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v9
}
