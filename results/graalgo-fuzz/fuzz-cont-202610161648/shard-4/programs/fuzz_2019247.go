package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint8 }

func main() {
    v1 := []string{}
    var v2 int8 = -98
    var v3 float64 = -2.25
    v4 := []uint32{2147483647}
    v5 := make(map[int]bool)
    var v6 int32 = -1
    v7 := GenS1{F0: 32768}
    var v8 int32 = (v6 / ((v6) | 1))
    var v9 float64 = ((v3 * v3) / ((1.5 * (float64(0.0) - v3)) / v3))
    var v10 int32 = (v8 >> ((387) & 63))
    v11 := GenS1{F0: 127}
    if true {
        fmt.Println(v9, v2, v3)
        var v12 uint16 = 0
        var v13 float64 = float64(-2.25)
        _ = v12
        _ = v13
    }
    var v14 int = 1
    switch ((("a" + "foo bar") + (("hello" + ((("hello" + ("" + "a")) + ("a" + "")) + "\n\t")) + "hello")) + "\n\t") {
    case "\n\t":
        fmt.Println(v10)
    case "a":
        switch -557 {
        case 127:
            var v15 int8 = 100
            _ = v15
        }
    }
    var v16 string = ((((((((("foo bar" + "hello") + ("a" + "hello")) + "") + "") + "") + "hello") + "foo bar") + "hello") + ("\n\t" + ""))
    if false {
        v2 = v2
        v16 = "\n\t"
        if (true && (!false)) {
            v17 := GenS1{F0: (v14 & 976)}
            var v18 string = (v16 + v16)
            _ = v17
            _ = v18
        }
    }
    switch "a" {
    case "\n\t":
        v19 := v5[(v14 * v14)]
        _ = v19
    case "foo bar":
        if false {
            v20 := GenS2{F0: 1}
            if (v3 == v9) {
                for fi21 := 0; fi21 < 17; fi21++ {
                    v9 = v9
                    for fi22 := 0; fi22 < 11; fi22++ {
                        v23 := GenS3{F0: 0, F1: 279}
                        _ = fi22
                        _ = v23
                    }
                    v6 -= -2147483648
                    _ = fi21
                }
                if (533 >= 1000) {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v4 = append(v4, 2147483647)
                        v25 := GenS4{F0: 18446744073709551615, F1: 201}
                        _ = fi24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 1; fi26++ {
                        var v27 int32 = -131
                        _ = fi26
                        _ = v27
                    }
                    v28 := v5[v14]
                    _ = v28
                }
                if (int32(v2) > v8) {
                    fmt.Println(v16, v9)
                    v14 = -433
                }
            }
            _ = v20
        } else {
            v6 = 0
            var v29 uint64 = 2
            _ = v29
        }
    case "":
        v14 = 9223372036854775807
    }
    switch ((("" + v16) + (((v16 + v16) + v16) + v16)) + v16) {
    case "hello":
        if false {
            v30 := v5[32767]
            v14 = 48
            v6 -= 2147483647
            _ = v30
        }
    case "foo bar":
        v6 -= -2147483648
    case "\n\t":
        v10 = v6
    default:
        switch v14 {
        case 1:
            fmt.Println(v9, v2, v9)
        }
    }
    switch "\n\t" {
    case "\n\t":
        _ = len(v5)
    case "foo bar":
        fmt.Println(v2, v9, v10)
    case "":
        v10 = v6
    }
    var v31 string = "hello"
    v31 = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
