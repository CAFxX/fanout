package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int8; F1 int32; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 int32 }
func genFn0(p0 int) uint16 {
    var v6 int8 = -128
    v7 := []int{p0, (-(p0)), p0}
    v6 = 82
    switch "" {
    case "hello":
        fmt.Println(p0)
    case "\n\t":
        switch p0 {
        case 128:
            for fi8 := 0; fi8 < 2; fi8++ {
                if (((fi8 / ((fi8) | 1)) == (p0 ^ 127)) || false) {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        fmt.Println(fi8)
                        var v10 string = "\n\t"
                        _ = fi9
                        _ = v10
                    }
                } else {
                    v6 += v6
                    fmt.Println(fi8, v6, fi8)
                }
                for fi11 := 0; fi11 < 15; fi11++ {
                    v12 := GenS1{F0: (v6 - 100), F1: 816, F2: 1}
                    _ = fi11
                    _ = v12
                }
                _ = fi8
            }
        case 4294967295:
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 uint32 = 0
                v15 := GenS2{F0: 2147483648, F1: -2147483648}
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            p0 += (p0 & 1)
        }
    default:
        switch 42 {
        case -1:
            if (311 == 18446744073709551615) {
                fmt.Println(v6, p0)
            }
        case 2147483647:
            v6 = v6
        case 32767:
            v6 += (v6 % ((127) | 1))
        }
    }
    var v16 uint16 = 65535
    _ = p0
    _ = v16
    _ = v6
    _ = v7
    return (v16 | v16)
}
func genFn1() int {
    switch -9223372036854775808 {
    case 32767:
        if (65535 > 0) {
            if (true && (1 <= 0)) {
                var v17 float64 = ((1e10 / 1e10) + -2.25)
                _ = v17
            }
            v18 := []int{277, -1}
            _ = v18
        }
    case 797:
        switch "foo bar" {
        case "":
            switch 80 {
            case 913:
                var v19 uint32 = 0
                _ = v19
            case -2147483648:
                v20 := GenS1{F0: 100, F1: 2, F2: 1}
                _ = v20
            case 32768:
                v21 := []int{}
                _ = v21
            default:
                var v22 int64 = 24
                _ = v22
            }
        case "foo bar":
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 float64 = 1.5
                v25 := GenS2{F0: 155, F1: -868}
                _ = fi23
                _ = v24
                _ = v25
            }
        case "\n\t":
            var v26 uint8 = 255
            _ = v26
        }
    case 42:
        switch "foo bar" {
        case "":
            var v27 uint64 = 18446744073709551615
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        for fi29 := 0; fi29 < 13; fi29++ {
            if ((false || false) && false) {
                var v30 int8 = 1
                _ = v30
            }
            _ = fi29
        }
        var v31 bool = (((true || true) || true) || false)
        _ = fi28
        _ = v31
    }
    if (1 == 248) {
        for fi32 := 0; fi32 < 5; fi32++ {
            v33 := []int{(fi32 / ((-2147483648) | 1)), fi32, 42, (fi32 % ((127) | 1))}
            _ = fi32
            _ = v33
        }
        if false {
            v34 := GenS1{F0: 0, F1: 0, F2: -323}
            switch "foo bar" {
            case "":
                v35 := []int{}
                _ = v35
            case "foo bar":
                var v36 float64 = 1.5
                _ = v36
            default:
                var v37 int64 = -1
                _ = v37
            }
            _ = v34
        } else {
            var v38 int8 = -128
            _ = v38
        }
    }
    switch 0 {
    case -2147483648:
        var v39 float64 = 1e10
        _ = v39
    case 9223372036854775807:
        for fi40 := 0; fi40 < 14; fi40++ {
            var v41 int8 = -128
            v42 := GenS3{F0: 0, F1: 1, F2: -1}
            _ = fi40
            _ = v41
            _ = v42
        }
    case 127:
        var v43 uint64 = 4294967295
        _ = v43
    default:
        var v44 int16 = 123
        _ = v44
    }
    return 32768
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 18446744073709551615
    var v3 int64 = 1
    v4 := make(map[int]int)
    var v5 int = 0
    var v45 uint8 = 1
    v2 = (v2 / ((v2) | 1))
    switch "foo bar" {
    case "\n\t":
        switch ((("" + "\n\t") + "") + ("foo bar" + (("foo bar" + "foo bar") + ((("" + ("hello" + "hello")) + (("hello" + "\n\t") + ("\n\t" + "hello"))) + "\n\t")))) {
        case "":
            fmt.Println(v5)
        case "foo bar":
            v2 = (^(v2))
        default:
            v45 = 48
        }
    }
    v3 = (v3 - v3)
    var v46 int64 = (v3 | v3)
    switch ("hello" + ("\n\t" + "foo bar")) {
    case "\n\t":
        if (true && ((!((v45 << ((255) & 63)) < v45)) || true)) {
            fmt.Println(v5, v2)
        }
    }
    fmt.Println(v2, v2)
    v5 -= (v5 | -814)
    var v47 string = ((("hello" + ((("" + ("a" + "hello")) + ("\n\t" + "")) + "a")) + "") + "foo bar")
    if false {
        for fi48 := 0; fi48 < 11; fi48++ {
            v49 := GenS4{F0: 42}
            v5 += (v5 | (-(v5)))
            _ = fi48
            _ = v49
        }
        var v50 int8 = 24
        v51 := GenS5{F0: (v3 * (v46 - (v46 * v46))), F1: 1}
        _ = v50
        _ = v51
    }
    switch genFn1() {
    case 127:
        var v52 float32 = 3.14159
        _ = v52
    default:
        var v53 float64 = (1.5 + 0.1)
        _ = v53
    }
    v5 = int(v45)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
