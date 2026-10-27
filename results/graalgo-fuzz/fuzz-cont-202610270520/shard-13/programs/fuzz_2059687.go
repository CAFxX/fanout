package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 bool; F2 uint }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int32; F2 int64 }
func genFn0() uint64 {
    var v7 int32 = 871
    var v8 float32 = 0.0
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS1{F0: -32768}
        var v11 int = fi9
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int32 = 765
    v2 := make(map[int]bool)
    var v3 float64 = (((0.1 / 0.1) / 1.5) - -2.25)
    var v4 float64 = 3.14159
    v5 := []uint32{916, 42, 42}
    var v6 int16 = -427
    fmt.Println(v6, v3)
    switch -345 {
    case 0:
        v12 := GenS2{F0: 42}
        _ = v12
    case -102:
        for fi13 := 0; fi13 < 19; fi13++ {
            if false {
                var v14 int16 = 1000
                v1 -= v1
                _ = len(v2)
                _ = v14
            }
            v5 = append(v5, 65535)
            for fi15 := 0; fi15 < 18; fi15++ {
                v16 := GenS2{F0: v1}
                v17 := GenS3{F0: ("a" + "foo bar")}
                switch "hello" {
                case "":
                    _ = len(v5)
                case "a":
                    switch "" {
                    case "hello":
                        v18 := GenS3{F0: "hello"}
                        switch "hello" {
                        case "foo bar":
                            v19 := GenS4{F0: v1, F1: (!true), F2: 0}
                            _ = v19
                        case "a":
                            var v20 uint = 2
                            _ = v20
                        default:
                            fmt.Println(v4, fi13)
                        }
                        _ = v18
                    case "\n\t":
                        v21 := GenS5{F0: "a", F1: -2147483648}
                        if false {
                            fmt.Println(fi15, fi13, v3)
                            fmt.Println(fi13)
                        }
                        _ = v21
                    }
                default:
                    v1 = (-(v1))
                    switch "foo bar" {
                    case "a":
                        v22 := GenS4{F0: 0, F1: (0 != 1), F2: 65535}
                        _ = v22
                    default:
                        v23 := GenS6{F0: 42}
                        _ = v23
                    }
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi13
        }
    case 32767:
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := v2[(fi24 ^ (fi24 | -233))]
            switch int(v6) {
            case -2147483648:
                var v26 uint16 = 32767
                _ = v26
            }
            var v27 bool = true
            _ = fi24
            _ = v25
            _ = v27
        }
    }
    var v28 int16 = (-(v6))
    var v29 float32 = 0.1
    if (!false) {
        fmt.Println(v1, v28)
        var v30 uint16 = 981
        var v31 uint64 = 18446744073709551615
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 6; fi32++ {
            switch (fi32 - (fi32 * (fi32 * (^(fi32))))) {
            case 0:
                var v33 bool = true
                _ = v33
            case 377:
                delete(v2, fi32)
            case 1:
                v34 := GenS6{F0: 65535}
                _ = v34
            }
            _ = fi32
        }
    }
    var v35 uint16 = 712
    v29 = 0.1
    v36 := GenS7{F0: -128, F1: -2147483648, F2: 667}
    v35 -= v35
    v35 -= v35
    var v37 float64 = v4
    v5 = append(v5, 42)
    if (496 >= 18446744073709551615) {
        switch 128 {
        case 768:
            v6 -= (v28 << ((18446744073709551615) & 63))
        case 545:
            delete(v2, -671)
        }
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        if true {
            var v39 uint32 = 53
            _ = v39
        } else {
            v29 = (v29 * v29)
            var v40 bool = (!(v29 < v29))
            _ = v40
        }
        for fi41 := 0; fi41 < 13; fi41++ {
            v37 += 1.5
            if false {
                _ = len(v5)
            }
            if (!((-(v1)) < v1)) {
                switch (fi38 ^ (fi41 & fi41)) {
                case 32767:
                    v4 -= v4
                case 32768:
                    if (!(!false)) {
                        switch "\n\t" {
                        case "foo bar":
                            v35 = v35
                        case "\n\t":
                            v29 += float32(v3)
                        }
                    }
                case 9223372036854775807:
                    _ = len(v5)
                default:
                    _ = len(v2)
                }
            }
            _ = fi41
        }
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
