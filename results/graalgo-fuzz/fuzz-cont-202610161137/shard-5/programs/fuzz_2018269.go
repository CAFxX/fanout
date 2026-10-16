package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint32; F2 string }
type GenS3 struct { F0 int64; F1 string; F2 int16 }
func genFn0(p0 int32) uint8 {
    p0 -= p0
    fmt.Println(p0)
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := []int{(fi7 / ((32768) | 1)), -2147483648, fi7, fi7}
        var v9 int = 2147483647
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-2147483648, 42, 154}
    var v11 int64 = 398
    _ = p0
    _ = v10
    _ = v11
    return 42
}
func genFn1(p0 uint16, p1 string) uint64 {
    switch 2147483647 {
    case 2147483647:
        p0 -= (p0 ^ 2)
    case 0:
        switch p1 {
        case "\n\t":
            v12 := GenS1{F0: 1}
            _ = v12
        }
    case -1:
        var v13 uint64 = uint64(p0)
        _ = v13
    default:
        p0 += (p0 * (p0 * p0))
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int64 = 42
        v16 := GenS2{F0: ("" + p1), F1: 2147483647, F2: (("\n\t" + ("foo bar" + "")) + p1)}
        if (uint(p0) < 834) {
            v17 := []int{-555}
            _ = v17
        } else {
            switch "" {
            case "\n\t":
                v18 := []int{}
                _ = v18
            case "":
                fmt.Println(v15, p0)
            case "a":
                v19 := []int{(fi14 % ((fi14) | 1)), (fi14 | fi14), 399, 32767}
                _ = v19
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v20 string = p1
    p1 = ("foo bar" + "hello")
    _ = p0
    _ = p1
    _ = v20
    return 42
}

func main() {
    var v1 uint8 = 1
    var v2 int = 32767
    var v3 int64 = -9223372036854775808
    var v4 uint8 = v1
    v5 := []uint32{2147483648, 2147483647}
    var v6 int16 = 686
    if ((1.5 <= 0.1) || false) {
        for fi21 := 0; fi21 < 6; fi21++ {
            fmt.Println(v6)
            var v22 uint = 2
            var v23 uint = 42
            _ = fi21
            _ = v22
            _ = v23
        }
        fmt.Println(v6)
        v6 -= v6
    } else {
        if false {
            var v24 float32 = 1.5
            var v25 uint16 = 318
            _ = v24
            _ = v25
        }
        switch "\n\t" {
        case "a":
            var v26 uint64 = 337
            _ = v26
        case "hello":
            var v27 float32 = -2.25
            _ = v27
        default:
            fmt.Println(v6, v3)
        }
    }
    var v28 uint = 322
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 17; fi30++ {
            v31 := GenS3{F0: 9223372036854775807, F1: "foo bar", F2: (v6 >> (((^(v28))) & 63))}
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    if (float32((float64(0.1) / 1.0)) >= -2.25) {
        fmt.Println(v6)
        v32 := GenS1{F0: 2}
        _ = v32
    }
    v3 = -968
    v3 = v3
    var v33 int64 = v3
    if (0.0 == (-2.25 / 1.5)) {
        var v34 int64 = v3
        if true {
            for fi35 := 0; fi35 < 10; fi35++ {
                fmt.Println(v6)
                v36 := GenS1{F0: 0}
                _ = fi35
                _ = v36
            }
            var v37 string = ((("hello" + "") + "foo bar") + "\n\t")
            var v38 uint = 42
            _ = v37
            _ = v38
        }
        var v39 uint32 = 2147483648
        _ = v34
        _ = v39
    } else {
        if false {
            v2 -= v2
            switch 30 {
            case 32767:
                var v40 float32 = 0.1
                _ = v40
            case 42:
                fmt.Println(v1, v4, v6)
            }
        }
    }
    if false {
        v5 = append(v5, 1)
        var v41 uint = v28
        v41 += 90
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
