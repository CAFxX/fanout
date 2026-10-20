package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 int8; F2 uint32 }
func genFn0() uint {
    var v7 uint64 = 1
    if (255 > 42) {
        v7 = 4294967295
    } else {
        v7 = 2
        for fi8 := 0; fi8 < 20; fi8++ {
            v7 -= (v7 * 685)
            _ = fi8
        }
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        if true {
            v7 -= 883
        } else {
            fmt.Println(fi9)
        }
        v7 = v7
        for fi10 := 0; fi10 < 9; fi10++ {
            if false {
                switch ("a" + "foo bar") {
                case "":
                    if true {
                        v11 := GenS1{F0: (v7 >> ((4294967295) & 63)), F1: "hello", F2: 2}
                        _ = v11
                    }
                case "a":
                    v7 = (v7 - (v7 << ((4294967295) & 63)))
                case "foo bar":
                    v7 = v7
                default:
                    v12 := GenS1{F0: (v7 + (^(v7))), F1: ("\n\t" + (("foo bar" + "\n\t") + "hello")), F2: 0}
                    _ = v12
                }
                v7 = v7
            }
            _ = fi10
        }
        _ = fi9
    }
    var v13 uint = 65535
    var v14 float64 = 0.1
    _ = v13
    _ = v14
    _ = v7
    return (v13 - (^(v13)))
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 179
    var v3 int16 = 939
    v4 := make(map[int]uint32)
    var v5 int = 791
    v6 := []uint32{1, 984}
    if (false || true) {
        var v15 uint8 = 0
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 float64 = 0.0
        v18 := GenS1{F0: 2, F1: "\n\t", F2: 65535}
        fmt.Println(v3)
        _ = fi16
        _ = v17
        _ = v18
    }
    switch ("hello" + "\n\t") {
    case "hello":
        v3 -= v3
    case "a":
        switch "" {
        case "hello":
            var v19 float64 = (-2.25 + 0.0)
            _ = v19
        case "a":
            var v20 uint64 = 0
            _ = v20
        case "foo bar":
            var v21 uint16 = 1
            _ = v21
        }
    default:
        var v22 string = "a"
        _ = v22
    }
    switch "a" {
    case "hello":
        switch ((("hello" + ("foo bar" + "a")) + "\n\t") + "hello") {
        case "hello":
            switch (v5 / (((v5 >> ((255) & 63))) | 1)) {
            case -9223372036854775808:
                v23 := GenS2{F0: 553}
                _ = v23
            default:
                v24 := GenS1{F0: 449, F1: ("\n\t" + ("\n\t" + "a")), F2: 255}
                _ = v24
            }
        }
    default:
        v4[1] = 2147483647
    }
    fmt.Println(v2, v3, v5)
    var v25 bool = (!true)
    v3 += v3
    var v26 uint32 = 2
    v2 += (v2 % (((v2 % (((^(v2))) | 1))) | 1))
    if (false && (305 >= 4294967295)) {
        if (304 != -2147483648) {
            v2 -= v2
            v5 = (^(v5))
        }
    }
    var v27 float64 = 0.0
    if (369 == 1) {
        var v28 int64 = int64(v5)
        _ = v28
    } else {
        if (-833 >= 0) {
            v29 := GenS1{F0: 404, F1: "\n\t", F2: 65535}
            for fi30 := 0; fi30 < 8; fi30++ {
                fmt.Println(v25)
                v31 := GenS3{F0: (fi30 << ((791) & 63)), F1: -128, F2: v26}
                _ = fi30
                _ = v31
            }
            fmt.Println(v5)
            _ = v29
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 uint64 = 446
            _ = fi32
            _ = v33
        }
        switch "" {
        case "":
            var v34 float32 = 0.0
            _ = v34
        default:
            v35 := GenS2{F0: 693}
            _ = v35
        }
    }
    var v36 int64 = 0
    v6 = append(v6, (v26 % ((v26) | 1)))
    v37 := GenS3{F0: (v5 | v5), F1: 127, F2: v26}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
