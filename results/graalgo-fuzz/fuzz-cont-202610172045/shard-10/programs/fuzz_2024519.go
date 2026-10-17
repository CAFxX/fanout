package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint16) uint16 {
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 int64 = 9223372036854775807
        v5 = (v5 / ((v5) | 1))
        v5 = -526
        _ = fi4
        _ = v5
    }
    var v6 int8 = -1
    _ = p0
    _ = v6
    return (p0 + p0)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    v3 := []int{4294967295, 657}
    var v7 float32 = 0.1
    if false {
        fmt.Println(v7)
        var v8 uint = 2
        _ = v8
    }
    switch "hello" {
    case "\n\t":
        for fi9 := 0; fi9 < 7; fi9++ {
            switch (fi9 / ((fi9) | 1)) {
            case -908:
                fmt.Println(fi9, fi9)
            default:
                var v10 int64 = 14
                _ = v10
            }
            switch (fi9 - (-(fi9))) {
            case 753:
                switch ((((((("a" + "foo bar") + "foo bar") + "hello") + "hello") + "a") + "") + ("hello" + "a")) {
                case "hello":
                    var v11 uint32 = 666
                    _ = v11
                case "":
                    for fi12 := 0; fi12 < 10; fi12++ {
                        fmt.Println(fi12, v7)
                        switch "" {
                        case "foo bar":
                            v7 = v7
                        case "":
                            var v13 uint32 = 244
                            _ = v13
                        case "hello":
                            switch (fi9 ^ fi12) {
                            case -2147483648:
                                if ((127 > 0) || ((^(fi12)) < (fi9 / (((fi9 | (fi9 * (fi9 + -54)))) | 1)))) {
                                    v7 += float32(1.5)
                                    if false {
                                        v14 := GenS1{F0: "a"}
                                        var v15 uint = 42
                                        v15 += v15
                                        _ = v14
                                        _ = v15
                                    }
                                    var v16 uint32 = 2147483647
                                    v16 = 4294967295
                                    var v17 uint = 632
                                    _ = v16
                                    _ = v17
                                } else {
                                    var v18 uint = 41
                                    v19 := GenS1{F0: "a"}
                                    _ = v18
                                    _ = v19
                                }
                            }
                        default:
                            if len(v3) > 0 {
                                v3[0] = (-(fi12))
                            }
                        }
                        if (-128 < 0) {
                            v20 := GenS2{F0: 175}
                            var v21 float64 = 3.14159
                            var v22 string = ""
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        _ = fi12
                    }
                }
            default:
                for fi23 := 0; fi23 < 1; fi23++ {
                    switch ("" + "hello") {
                    case "hello":
                        switch -9223372036854775808 {
                        case 49:
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                switch -9223372036854775808 {
                                case 4294967295:
                                    switch (fi23 >> ((42) & 63)) {
                                    case -2147483648:
                                        v7 -= float32(1e10)
                                    case 42:
                                        var v24 string = ""
                                        _ = v24
                                    case 32767:
                                        var v25 int32 = 1
                                        _ = v25
                                    }
                                }
                            }
                        }
                    case "\n\t":
                        if (-155 == -9223372036854775808) {
                            v7 = 1e10
                            v26 := GenS2{F0: 127}
                            var v27 int = (fi9 % (((fi23 >> ((0) & 63))) | 1))
                            _ = v26
                            _ = v27
                        } else {
                            fmt.Println(v7, fi9)
                            var v28 int64 = 761
                            var v29 int32 = -1
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = fi23
                }
            }
            var v30 string = "a"
            _ = fi9
            _ = v30
        }
    case "":
        fmt.Println(v7)
    case "hello":
        v31 := v1[128]
        _ = v31
    default:
        v7 += 0.0
    }
    switch 560 {
    case 392:
        var v32 float64 = (1.5 / 1.5)
        _ = v32
    case 282:
        var v33 uint = 677
        _ = v33
    default:
        v34 := GenS2{F0: 128}
        _ = v34
    }
    v7 -= v7
    for fi35 := 0; fi35 < 12; fi35++ {
        fmt.Println(fi35)
        _ = fi35
    }
    fmt.Println(v7)
    v7 -= v7
    var v36 int16 = 42
    var v37 int16 = 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v7
}
