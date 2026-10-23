package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 string; F2 uint16 }
type GenS5 struct { F0 int; F1 uint64; F2 string }

func main() {
    v1 := []string{((((("" + "\n\t") + "") + "foo bar") + "a") + "hello"), "a", "hello", "foo bar"}
    var v2 uint = 18446744073709551615
    v3 := []int{}
    var v4 string = "a"
    var v5 int32 = 2147483647
    v6 := GenS1{F0: 2147483647}
    var v7 float32 = 1.5
    if false {
        switch 267 {
        case 9223372036854775807:
            v8 := GenS2{F0: 528}
            _ = v8
        default:
            var v9 uint = 179
            _ = v9
        }
        switch ((("a" + v4) + "hello") + "\n\t") {
        case "\n\t":
            if (-1 > 9223372036854775807) {
                v4 = v4
            }
        case "foo bar":
            v7 -= -2.25
        case "a":
            v4 = v4
        }
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v4 = (v4 + v4)
        switch 4294967295 {
        case -9223372036854775808:
            var v11 float32 = v7
            _ = v11
        case -230:
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 uint = v2
                switch (fi12 * (fi10 ^ fi12)) {
                case -1:
                    fmt.Println(fi12, v13, v2)
                default:
                    v13 -= (v13 ^ (v2 - 190))
                }
                var v14 bool = true
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = fi10
    }
    switch -2147483648 {
    case 402:
        switch "" {
        case "foo bar":
            v15 := GenS1{F0: 783}
            _ = v15
        case "\n\t":
            if (0 > v5) {
                var v16 int = 32768
                var v17 float32 = -2.25
                switch ("a" + (("foo bar" + "foo bar") + ("a" + "hello"))) {
                case "a":
                    var v18 uint16 = 603
                    _ = v18
                }
                _ = v16
                _ = v17
            } else {
                switch 128 {
                case -163:
                    for fi19 := 0; fi19 < 5; fi19++ {
                        fmt.Println(v2, v4)
                        switch ("foo bar" + v4) {
                        case "hello":
                            switch v4 {
                            case "":
                                var v20 string = ("\n\t" + v4)
                                _ = v20
                            default:
                                _ = len(v3)
                            }
                        case "":
                            if false {
                                v7 -= float32((1.5 / 1.5))
                                v21 := GenS1{F0: 32768}
                                var v22 bool = (!false)
                                _ = v21
                                _ = v22
                            }
                        case "foo bar":
                            v23 := GenS3{F0: false}
                            _ = v23
                        default:
                            v4 = v4
                        }
                        _ = fi19
                    }
                case 42:
                    for fi24 := 0; fi24 < 10; fi24++ {
                        var v25 int = fi24
                        var v26 uint = v2
                        v1 = append(v1, v4)
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
            }
        }
    case -85:
        var v27 float32 = (v7 - v7)
        _ = v27
    default:
        var v28 uint = (v2 / (((v2 & v2)) | 1))
        _ = v28
    }
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    v4 = v4
    if len(v1) > 0 {
        v1[0] = (("a" + "") + "hello")
    }
    var v29 int64 = -150
    switch -363 {
    case 4294967295:
        var v30 float32 = -2.25
        _ = v30
    case 808:
        var v31 uint8 = 42
        _ = v31
    case 127:
        _ = len(v1)
    }
    v29 += (v29 >> ((255) & 63))
    v32 := GenS2{F0: (v29 >> (((v2 >> ((4294967295) & 63))) & 63))}
    for fi33 := 0; fi33 < 14; fi33++ {
        v34 := GenS3{F0: (!false)}
        if ((false || ((v29 << (((v2 << (((v2 * 386)) & 63))) & 63)) != v29)) || false) {
            v29 += (v29 >> (((v2 * 439)) & 63))
            var v35 uint32 = 2147483648
            var v36 uint16 = 1
            _ = v35
            _ = v36
        } else {
            v37 := GenS4{F0: 2147483648, F1: v4, F2: 0}
            var v38 bool = false
            _ = v37
            _ = v38
        }
        _ = fi33
        _ = v34
    }
    if false {
        v5 = (^(v5))
        v2 = (v2 % (((^(v2))) | 1))
    }
    v39 := GenS5{F0: 127, F1: 700, F2: "foo bar"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
