package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
func genFn0(p0 int64, p1 int64) uint64 {
    p1 = (p0 * 894)
    v7 := GenS1{F0: 1, F1: 100, F2: "hello"}
    p0 = p0
    switch 0 {
    case 2147483647:
        var v8 float64 = float64(1.5)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    v1 := []int{128, -790}
    v2 := make(map[int]string)
    var v3 bool = true
    var v4 int = -777
    v5 := make(map[uint32]bool)
    v6 := make(map[int]bool)
    switch "" {
    case "":
        if false {
            var v9 uint8 = 128
            _ = v9
        } else {
            if ("a" > "a") {
                var v10 int = 127
                _ = v10
            }
            v11 := GenS1{F0: 42, F1: -1, F2: ("\n\t" + "foo bar")}
            _ = v11
        }
    case "a":
        var v12 float64 = ((1.5 * 3.14159) - 0.1)
        _ = v12
    default:
        var v13 int64 = 9223372036854775807
        _ = v13
    }
    v3 = (0.0 != 3.14159)
    switch "foo bar" {
    case "\n\t":
        switch (("hello" + "") + ("hello" + "\n\t")) {
        case "":
            switch 128 {
            case 32768:
                v4 = (v4 / (((v4 << ((2) & 63))) | 1))
            case 32767:
                switch "\n\t" {
                case "":
                    var v14 float32 = (1e10 / 0.1)
                    _ = v14
                case "hello":
                    delete(v2, v4)
                case "foo bar":
                    var v15 int16 = -32768
                    _ = v15
                }
            case 127:
                v4 += (v4 % ((v4) | 1))
            }
        default:
            v1 = append(v1, (-(v4)))
        }
    }
    var v16 uint16 = 0
    if (9223372036854775807 <= 0) {
        var v17 uint64 = 18446744073709551615
        switch (v4 ^ v4) {
        case 127:
            _ = len(v1)
        case 32767:
            if (-251 > -1) {
                var v18 uint = 0
                _ = v18
            }
        case 4294967295:
            v17 += genFn0(-32, -67)
        }
        _ = v17
    } else {
        var v19 int = 0
        _ = v19
    }
    switch "hello" {
    case "hello":
        for fi20 := 0; fi20 < 19; fi20++ {
            if (!(2 > 2)) {
                v2[(fi20 / ((fi20) | 1))] = ""
                _ = len(v1)
                for fi21 := 0; fi21 < 3; fi21++ {
                    if len(v1) > 0 {
                        v1[0] = -1
                    }
                    fmt.Println(fi21, v4)
                    v4 -= (fi20 & (fi21 << ((0) & 63)))
                    if false {
                        if v3 {
                            v22 := v6[fi20]
                            var v23 uint32 = 461
                            _ = v22
                            _ = v23
                        }
                    } else {
                        var v24 int32 = 42
                        _ = v24
                    }
                    _ = fi21
                }
            }
            var v25 int64 = 447
            fmt.Println(v25)
            _ = fi20
            _ = v25
        }
    case "a":
        fmt.Println(v4, v16, v16)
    case "foo bar":
        var v26 int64 = -9223372036854775808
        _ = v26
    }
    fmt.Println(v16, v3)
    fmt.Println(v3, v16)
    for fi27 := 0; fi27 < 11; fi27++ {
        switch "a" {
        case "":
            switch (("\n\t" + "hello") + ("hello" + "hello")) {
            case "foo bar":
                delete(v5, 88)
            case "a":
                switch -582 {
                case -9223372036854775808:
                    v28 := GenS1{F0: 42, F1: 42, F2: ("foo bar" + ("hello" + "foo bar"))}
                    _ = v28
                case -1:
                    v1 = append(v1, (-(v4)))
                }
            }
        case "a":
            v16 -= (v16 * v16)
        case "\n\t":
            switch "" {
            case "foo bar":
                v29 := GenS2{F0: 42, F1: (^(v16)), F2: "hello"}
                _ = v29
            default:
                fmt.Println(fi27, fi27)
            }
        }
        _ = fi27
    }
    fmt.Println(v4)
    v16 = v16
    var v30 uint16 = 1
    v30 -= (v30 << ((873) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
