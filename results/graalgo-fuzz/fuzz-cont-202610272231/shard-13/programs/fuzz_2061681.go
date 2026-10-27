package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 string; F2 uint64 }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 uint8 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int16, p1 uint64, p2 uint8) int32 {
    p1 = (^(p1))
    if (!((((!(false || true)) && false) || ((p1 % ((2) | 1)) != (p1 | p1))) || true)) {
        p2 += 2
        v5 := []int{32767, 2147483647}
        _ = v5
    } else {
        if (32767 < 271) {
            switch "a" {
            case "hello":
                if true {
                    if true {
                        v6 := []int{-842, 32767}
                        var v7 uint32 = 0
                        p2 = (^(p2))
                        _ = v6
                        _ = v7
                    }
                    fmt.Println(p2, p1, p2)
                } else {
                    var v8 uint = 65535
                    var v9 int = 42
                    _ = v8
                    _ = v9
                }
            case "\n\t":
                v10 := GenS1{F0: 2147483648, F1: "foo bar"}
                _ = v10
            }
        } else {
            var v11 int = -2147483648
            p1 += 781
            _ = v11
        }
        if true {
            switch (("" + "\n\t") + (((("" + "a") + "\n\t") + "hello") + ((("hello" + "foo bar") + "\n\t") + "a"))) {
            case "hello":
                p0 = 32767
            default:
                var v12 int16 = 787
                _ = v12
            }
            v13 := GenS2{F0: (p0 | (p0 / ((p0) | 1))), F1: ("hello" + ("foo bar" + "\n\t")), F2: (p1 ^ 2)}
            _ = v13
        } else {
            if true {
                var v14 float32 = 0.0
                p1 = p1
                _ = v14
            } else {
                var v15 uint32 = 1
                switch ("foo bar" + "a") {
                case "hello":
                    p2 += 127
                case "":
                    var v16 uint32 = v15
                    _ = v16
                case "foo bar":
                    switch -821 {
                    case 1:
                        v17 := []int{}
                        _ = v17
                    case -346:
                        v18 := GenS1{F0: 65535, F1: "hello"}
                        _ = v18
                    default:
                        p0 = p0
                    }
                }
                _ = v15
            }
            var v19 bool = (523 <= 1)
            _ = v19
        }
        fmt.Println(p2, p2)
    }
    switch ("foo bar" + "\n\t") {
    case "a":
        v20 := GenS2{F0: 42, F1: "\n\t", F2: 2}
        _ = v20
    }
    v21 := GenS2{F0: (p0 ^ (p0 >> ((18446744073709551615) & 63))), F1: "\n\t", F2: (p1 & (p1 % ((p1) | 1)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return -205
}
func genFn1(p0 bool) int {
    p0 = p0
    if p0 {
        fmt.Println(p0)
        v22 := GenS2{F0: -64, F1: "\n\t", F2: 2}
        p0 = p0
        _ = v22
    }
    v23 := GenS3{F0: 1, F1: "hello", F2: 1}
    _ = p0
    _ = v23
    return -621
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]int)
    var v3 uint64 = 18446744073709551615
    var v4 bool = false
    var v24 uint16 = 608
    v4 = v4
    if v4 {
        if (!(127 < 99)) {
            switch "\n\t" {
            case "foo bar":
                v3 = v3
            }
            switch ((("a" + "a") + "a") + "foo bar") {
            case "a":
                if (123 <= -61) {
                    fmt.Println(v4, v3, v4)
                    var v25 string = "hello"
                    for fi26 := 0; fi26 < 1; fi26++ {
                        fmt.Println(fi26, v25, fi26)
                        _ = fi26
                    }
                    _ = v25
                } else {
                    switch -670 {
                    case 981:
                        var v27 bool = v4
                        _ = v27
                    case -299:
                        _ = len(v2)
                    }
                }
            }
            v4 = (!(42 > 15))
        }
        switch -322 {
        case -1:
            fmt.Println(v4, v4)
        case 4294967295:
            v28 := GenS1{F0: 2147483647, F1: ("foo bar" + "\n\t")}
            _ = v28
        default:
            v29 := GenS4{F0: -1, F1: -1, F2: 2}
            _ = v29
        }
    }
    fmt.Println(v4, v3, v24)
    for fi30 := 0; fi30 < 16; fi30++ {
        for fi31 := 0; fi31 < 5; fi31++ {
            v24 = 0
            if false {
                fmt.Println(fi30)
                switch fi30 {
                case 32768:
                    v3 = 4294967295
                case 128:
                    v1 = append(v1, 2147483647)
                default:
                    v32 := GenS1{F0: 2147483647, F1: "foo bar"}
                    _ = v32
                }
            }
            v1 = append(v1, 42)
            _ = fi31
        }
        for fi33 := 0; fi33 < 14; fi33++ {
            v34 := GenS2{F0: -324, F1: "\n\t", F2: 0}
            _ = fi33
            _ = v34
        }
        var v35 uint = 255
        _ = fi30
        _ = v35
    }
    var v36 int16 = 780
    fmt.Println(v4, v4)
    switch "" {
    case "\n\t":
        v4 = (0.1 == 1.5)
    default:
        v37 := GenS5{F0: 9223372036854775807}
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v4
}
