package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int16; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int8; F2 int8 }
type GenS5 struct { F0 uint16; F1 int8; F2 int }

func main() {
    var v1 bool = (20 >= 71)
    v2 := make(map[string]string)
    var v3 bool = v1
    v4 := GenS1{F0: 128}
    switch "\n\t" {
    case "hello":
        v1 = (1 > 686)
    case "":
        v5 := GenS1{F0: 175}
        _ = v5
    }
    v3 = (!(762 >= 1))
    v2["\n\t"] = "foo bar"
    var v6 uint16 = 42
    if v3 {
        v1 = (!false)
        v3 = ((0.0 * (1e10 + 0.0)) >= 3.14159)
    }
    fmt.Println(v1, v3, v3)
    var v7 float64 = -2.25
    switch "\n\t" {
    case "\n\t":
        v8 := GenS2{F0: -1, F1: -1, F2: 738}
        _ = v8
    case "hello":
        if (!(0 < 4294967295)) {
            for fi9 := 0; fi9 < 5; fi9++ {
                if (4294967295 <= 2) {
                    v10 := GenS3{F0: -32768}
                    var v11 string = "hello"
                    _ = v10
                    _ = v11
                }
                v7 = 0.1
                for fi12 := 0; fi12 < 10; fi12++ {
                    v1 = false
                    fmt.Println(v3)
                    _ = fi12
                }
                _ = fi9
            }
        }
    default:
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 uint8 = 0
            v7 += -2.25
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        if true {
            fmt.Println(v1)
            var v16 int = (fi15 - (-(fi15)))
            _ = v16
        } else {
            if ("" == "a") {
                fmt.Println(v7, v7, v3)
                switch 127 {
                case 32767:
                    switch "" {
                    case "\n\t":
                        var v17 uint32 = 0
                        _ = v17
                    default:
                        v1 = (18446744073709551615 >= 18446744073709551615)
                    }
                case -1:
                    v18 := GenS2{F0: 316, F1: -32768, F2: 250}
                    _ = v18
                case 42:
                    v19 := []int{(fi15 >> ((4294967295) & 63)), (fi15 | (fi15 & fi15)), fi15}
                    _ = v19
                default:
                    v3 = (!(!v3))
                }
                var v20 uint64 = 386
                _ = v20
            }
            v21 := GenS4{F0: -9223372036854775808, F1: 118, F2: 1}
            _ = v21
        }
        var v22 string = "hello"
        v23 := GenS2{F0: -9223372036854775808, F1: 178, F2: 1}
        _ = fi15
        _ = v22
        _ = v23
    }
    fmt.Println(v1, v6)
    if true {
        for fi24 := 0; fi24 < 2; fi24++ {
            for fi25 := 0; fi25 < 1; fi25++ {
                if v1 {
                    var v26 uint64 = 18446744073709551615
                    _ = v26
                }
                if false {
                    v1 = (false || v3)
                    v1 = (127 == -85)
                    v6 -= 42
                }
                v27 := v2["foo bar"]
                _ = fi25
                _ = v27
            }
            if v3 {
                for fi28 := 0; fi28 < 10; fi28++ {
                    if true {
                        var v29 uint8 = 42
                        _ = v29
                    } else {
                        var v30 uint32 = 0
                        _ = v30
                    }
                    switch (("a" + ("" + "foo bar")) + "\n\t") {
                    case "\n\t":
                        v31 := GenS2{F0: 9223372036854775807, F1: -1, F2: 741}
                        _ = v31
                    default:
                        v1 = v3
                    }
                    _ = fi28
                }
                v32 := GenS5{F0: (v6 * (^(v6))), F1: 9, F2: fi24}
                _ = v32
            } else {
                v1 = (-1 == -1)
            }
            var v33 uint32 = 37
            _ = fi24
            _ = v33
        }
        v34 := GenS3{F0: 513}
        var v35 int = 796
        _ = v34
        _ = v35
    }
    var v36 int8 = 24
    v37 := []int{-2147483648, -2147483648, -1, 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v6
    _ = v7
}
