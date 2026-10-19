package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS3 struct { F0 int32; F1 int64; F2 int16 }
type GenS4 struct { F0 int16; F1 uint32; F2 int }
type GenS5 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.1
    v3 := make(map[string]uint32)
    switch -677 {
    case 32768:
        v2 = v2
    case 321:
        for fi4 := 0; fi4 < 14; fi4++ {
            _ = len(v3)
            if (fi4 <= (^(fi4))) {
                for fi5 := 0; fi5 < 9; fi5++ {
                    fmt.Println(v2, v2)
                    v2 += v2
                    switch "foo bar" {
                    case "":
                        switch (("hello" + "foo bar") + "\n\t") {
                        case "":
                            var v6 uint = 42
                            _ = v6
                        case "\n\t":
                            switch "foo bar" {
                            case "\n\t":
                                v2 = ((float64(((0.0 / (1e10 + -2.25)) + 3.14159)) / 1.0) - v2)
                            case "hello":
                                var v7 uint16 = 65535
                                delete(v3, (("a" + "") + "a"))
                                _ = v7
                            case "foo bar":
                                v8 := []int{(fi4 + (^(fi4))), fi4, (fi5 | fi4)}
                                _ = v8
                            }
                            v2 -= -2.25
                            v2 = v2
                        case "hello":
                            v9 := GenS1{F0: "a"}
                            _ = v9
                        default:
                            switch "a" {
                            case "a":
                                v2 += v2
                            case "hello":
                                var v10 uint32 = 480
                                _ = v10
                            default:
                                switch ("\n\t" + "a") {
                                case "":
                                    v11 := GenS2{F0: 127, F1: 42, F2: 255}
                                    _ = v11
                                default:
                                    v2 -= (-2.25 * v2)
                                }
                            }
                            delete(v1, "a")
                        }
                    }
                    _ = fi5
                }
                fmt.Println(v2, v2)
            } else {
                var v12 int = (fi4 % ((127) | 1))
                _ = v12
            }
            v2 = v2
            _ = fi4
        }
    case -9223372036854775808:
        fmt.Println(v2)
    default:
        v13 := GenS3{F0: 42, F1: 899, F2: 0}
        _ = v13
    }
    v2 -= v2
    fmt.Println(v2)
    v14 := GenS4{F0: -313, F1: 2147483647, F2: 1}
    switch "" {
    case "foo bar":
        fmt.Println(v2)
    }
    if (924 >= 32767) {
        switch ((("a" + ((("hello" + "hello") + ("\n\t" + "a")) + "\n\t")) + ("\n\t" + "hello")) + ("\n\t" + "a")) {
        case "\n\t":
            if (!false) {
                fmt.Println(v2)
                for fi15 := 0; fi15 < 16; fi15++ {
                    if (!true) {
                        v16 := GenS5{F0: 1, F1: 0, F2: 178}
                        _ = v16
                    } else {
                        v2 -= v2
                        v17 := GenS1{F0: "\n\t"}
                        _ = v17
                    }
                    var v18 uint16 = 32767
                    fmt.Println(v18, v2)
                    _ = fi15
                    _ = v18
                }
            }
        case "a":
            v19 := []int{420}
            _ = v19
        case "foo bar":
            v2 += (v2 / v2)
        }
    } else {
        v20 := v1["foo bar"]
        _ = v20
    }
    switch "\n\t" {
    case "":
        if (175 == 1) {
            if false {
                v21 := GenS6{F0: 457, F1: 1}
                fmt.Println(v2)
                _ = v21
            }
        } else {
            _ = len(v1)
        }
    case "foo bar":
        v3["\n\t"] = 72
    }
    v2 = float64(-2.25)
    if (42 <= 658) {
        if true {
            v2 += v2
        }
    } else {
        var v22 int = 4294967295
        _ = v22
    }
    v23 := []int{}
    v2 += 1.5
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 bool = true
        v26 := GenS5{F0: -340, F1: 0, F2: 2}
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS7{F0: 389}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v27
    _ = v3
}
