package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 bool, p1 uint32) string {
    p0 = p0
    for fi5 := 0; fi5 < 10; fi5++ {
        switch "\n\t" {
        case "hello":
            v6 := []int{127, 2147483647}
            _ = v6
        case "a":
            if ((float32(-2.25) - float32(-2.25)) >= ((0.0 / 1.0) + 1e10)) {
                var v7 uint64 = 249
                _ = v7
            } else {
                p0 = (float32((1.5 + -2.25)) <= 0.0)
            }
        }
        p0 = p0
        _ = fi5
    }
    var v8 uint64 = 424
    _ = p0
    _ = p1
    _ = v8
    return ("a" + "foo bar")
}

func main() {
    v1 := []string{"foo bar"}
    var v2 float32 = 0.0
    var v3 int8 = -128
    var v4 int8 = v3
    v2 += 0.0
    fmt.Println(v3)
    var v9 uint32 = 2147483647
    switch "\n\t" {
    case "foo bar":
        var v10 uint64 = 176
        _ = v10
    case "hello":
        v2 -= (0.0 + (v2 * (v2 + v2)))
    case "\n\t":
        if (!(false || true)) {
            v9 += v9
            for fi11 := 0; fi11 < 7; fi11++ {
                if false {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    v2 = v2
                    for fi12 := 0; fi12 < 1; fi12++ {
                        switch ("" + "\n\t") {
                        case "\n\t":
                            v13 := GenS1{F0: false}
                            _ = v13
                        }
                        if ("foo bar" == ("foo bar" + "")) {
                            _ = len(v1)
                            var v14 string = "hello"
                            fmt.Println(v3)
                            _ = v14
                        } else {
                            v3 += 55
                            var v15 int = fi11
                            _ = v15
                        }
                        _ = fi12
                    }
                } else {
                    v16 := GenS2{F0: -32768, F1: (v3 | -128)}
                    for fi17 := 0; fi17 < 6; fi17++ {
                        var v18 uint64 = 4294967295
                        var v19 int64 = 0
                        var v20 int = 127
                        if true {
                            if (true || (4294967295 > 462)) {
                                v4 = v3
                                _ = len(v1)
                                if (v9 >= v9) {
                                    v9 += (v9 & v9)
                                } else {
                                    fmt.Println(v18, v9)
                                }
                            } else {
                                fmt.Println(v3, v19)
                                v2 -= (((v2 - 0.1) + 0.0) / ((float32(0.0) + 3.14159) / ((float32((0.1 + 1.5)) + 0.1) * (((-2.25 * ((((float32(0.1) - v2) * -2.25) / v2) - float32((1e10 * -2.25)))) * v2) / (float32(1.5) - (v2 * v2))))))
                            }
                        }
                        fmt.Println(v9)
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    var v21 uint = 65535
                    _ = v16
                    _ = v21
                }
                v3 = v4
                fmt.Println(v9, v9)
                _ = fi11
            }
            v22 := GenS3{F0: (0.1 > 1e10), F1: -44}
            _ = v22
        } else {
            switch (("" + "") + "foo bar") {
            case "foo bar":
                var v23 int64 = 0
                _ = v23
            default:
                v9 += 1
            }
            switch 128 {
            case 2147483647:
                var v24 int = -890
                _ = v24
            case 1:
                var v25 float32 = v2
                _ = v25
            case -520:
                v4 = (v4 / ((1) | 1))
            }
            var v26 uint = 249
            _ = v26
        }
    }
    var v27 uint64 = 4294967295
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := GenS3{F0: true, F1: (v4 % (((v4 << ((48) & 63))) | 1))}
        v30 := GenS4{F0: -1}
        _ = fi28
        _ = v29
        _ = v30
    }
    v9 = (^(v9))
    fmt.Println(v9, v2)
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
