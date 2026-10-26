package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint8; F1 int8 }

func main() {
    var v1 bool = (112 == 2)
    var v2 uint64 = 0
    var v3 int = 1
    v4 := make(map[string]int)
    switch "" {
    case "a":
        if v1 {
            for fi5 := 0; fi5 < 7; fi5++ {
                if (fi5 > 629) {
                    v2 += v2
                    switch ("" + "\n\t") {
                    case "":
                        if false {
                            fmt.Println(v3, fi5)
                        } else {
                            v6 := GenS1{F0: v1, F1: ("hello" + "\n\t"), F2: 2}
                            _ = v6
                        }
                    }
                    for fi7 := 0; fi7 < 15; fi7++ {
                        fmt.Println(v1, fi5, v2)
                        var v8 uint8 = 128
                        if ((fi5 == fi7) && false) {
                            var v9 uint16 = 697
                            var v10 uint64 = v2
                            _ = v10
                            _ = v9
                        } else {
                            if v1 {
                                if ((float64((1e10 + -2.25)) != 0.0) && v1) {
                                    var v11 int64 = -9223372036854775808
                                    _ = v11
                                } else {
                                    var v12 int64 = -1
                                    var v13 uint8 = (^(v8))
                                    var v14 string = "a"
                                    _ = v12
                                    _ = v13
                                    _ = v14
                                }
                                v3 = (-(v3))
                                v3 = (fi7 * 977)
                            }
                        }
                        _ = fi7
                        _ = v8
                    }
                } else {
                    if (v1 || (v1 && (65535 >= 255))) {
                        v2 -= (v2 | (v2 - v2))
                    }
                }
                v2 = (^(v2))
                v15 := GenS1{F0: v1, F1: "a", F2: 0}
                _ = fi5
                _ = v15
            }
            v2 = v2
        }
    }
    var v16 int = -2147483648
    v16 = v3
    switch (("foo bar" + (("hello" + "foo bar") + ("foo bar" + ("a" + "hello")))) + "foo bar") {
    case "foo bar":
        v2 = v2
    case "\n\t":
        switch "" {
        case "hello":
            var v17 bool = true
            _ = v17
        case "a":
            for fi18 := 0; fi18 < 16; fi18++ {
                fmt.Println(v1)
                _ = fi18
            }
        }
    }
    v1 = v1
    _ = len(v4)
    var v19 int16 = -322
    v4[("foo bar" + "\n\t")] = v3
    fmt.Println(v2, v3, v3)
    for fi20 := 0; fi20 < 9; fi20++ {
        if true {
            v21 := GenS2{F0: true, F1: 1}
            v1 = (!(42 >= 42))
            _ = len(v4)
            _ = v21
        }
        _ = fi20
    }
    var v22 float64 = 3.14159
    v23 := GenS3{F0: v19, F1: -941}
    if (!("foo bar" < "hello")) {
        v22 = float64(-2.25)
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 bool = v1
            fmt.Println(v3)
            _ = fi24
            _ = v25
        }
    }
    switch "a" {
    case "hello":
        fmt.Println(v22, v16, v16)
    case "foo bar":
        for fi26 := 0; fi26 < 13; fi26++ {
            v27 := GenS4{F0: 1, F1: 0}
            var v28 int32 = 2147483647
            _ = fi26
            _ = v27
            _ = v28
        }
    case "":
        v16 = 1
    }
    fmt.Println(v22, v22, v1)
    switch v3 {
    case 866:
        v2 = v2
    case 4294967295:
        for fi29 := 0; fi29 < 19; fi29++ {
            v30 := []int{(-(fi29)), (v16 << ((255) & 63)), (fi29 & (v16 | (-(fi29)))), (v16 >> ((265) & 63))}
            var v31 uint64 = v2
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
