package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int; F2 int }
func genFn0(p0 uint8, p1 uint64) uint16 {
    fmt.Println(p0, p0)
    p0 += p0
    v6 := []int{}
    if false {
        var v7 int32 = 1
        _ = v7
    } else {
        var v8 int8 = 111
        var v9 int = 0
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 714
}
func genFn1(p0 int) bool {
    switch (p0 | 2147483647) {
    case -934:
        v10 := GenS1{F0: 2}
        _ = v10
    default:
        p0 += (p0 | p0)
    }
    fmt.Println(p0)
    p0 = (^(p0))
    p0 = (p0 ^ (p0 * (p0 + -9223372036854775808)))
    switch p0 {
    case 128:
        p0 = (p0 / (((p0 ^ -9223372036854775808)) | 1))
    case 2147483647:
        if (!(((!(255 < 65535)) && false) && true)) {
            var v11 bool = (false && true)
            var v12 float32 = ((float32((3.14159 + (1.5 - 0.0))) / 0.1) * 0.0)
            _ = v11
            _ = v12
        } else {
            fmt.Println(p0)
            switch ("hello" + "") {
            case "":
                var v13 float64 = -2.25
                _ = v13
            case "hello":
                v14 := []int{p0, 42}
                _ = v14
            }
            v15 := []int{p0}
            _ = v15
        }
    default:
        switch p0 {
        case -328:
            switch 32767 {
            case 42:
                fmt.Println(p0)
            }
        }
    }
    _ = p0
    return (100 < 0)
}

func main() {
    var v1 uint8 = 0
    var v2 int32 = -1
    v3 := []int{int(v2)}
    var v4 uint16 = 1
    var v5 float64 = 3.14159
    if (656 > 374) {
        if false {
            var v16 uint = 65535
            _ = v16
        } else {
            var v17 int = -305
            v5 += -2.25
            _ = v17
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            if len(v3) > 0 {
                v3[0] = (fi18 ^ fi18)
            }
            fmt.Println(v5, v5)
            _ = fi18
        }
        fmt.Println(v5, v1)
    }
    var v19 float64 = v5
    fmt.Println(v1, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch ("" + "foo bar") {
    case "hello":
        v4 = (v4 * genFn0((v1 * v1), 42))
    case "":
        var v20 int32 = 2147483647
        _ = v20
    case "foo bar":
        v21 := GenS1{F0: 4294967295}
        _ = v21
    }
    switch -2147483648 {
    case 127:
        for fi22 := 0; fi22 < 16; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 float64 = v19
                var v25 uint8 = 167
                v26 := GenS1{F0: 4294967295}
                fmt.Println(fi22, v2, v24)
                switch "foo bar" {
                case "hello":
                    var v27 uint16 = v4
                    _ = v27
                case "foo bar":
                    v28 := GenS2{F0: (v2 & 630), F1: int(v1), F2: (fi22 >> ((18446744073709551615) & 63))}
                    _ = v28
                default:
                    var v29 int32 = (v2 * (-(v2)))
                    _ = v29
                }
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi22
        }
    default:
        if (false && (378 < 98)) {
            fmt.Println(v2)
            v2 = v2
        }
    }
    fmt.Println(v2, v1, v5)
    var v30 int8 = 31
    var v31 int32 = 42
    v32 := GenS1{F0: 65535}
    switch "foo bar" {
    case "\n\t":
        _ = len(v3)
    case "":
        for fi33 := 0; fi33 < 8; fi33++ {
            v5 -= (v19 + 1.5)
            fmt.Println(v2)
            _ = fi33
        }
    case "hello":
        v31 -= int32(v30)
    default:
        var v34 int = 32768
        _ = v34
    }
    v19 += 1e10
    v3 = append(v3, -892)
    for fi35 := 0; fi35 < 14; fi35++ {
        v36 := GenS1{F0: 4294967295}
        _ = fi35
        _ = v36
    }
    v4 -= (v4 % (((v4 >> ((65535) & 63))) | 1))
    v2 -= (^(v31))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
