package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    v1 := []uint32{42, 42}
    var v2 string = "foo bar"
    var v3 int = 127
    var v4 bool = (9223372036854775807 <= 9223372036854775807)
    _ = len(v1)
    if (v4 && false) {
        var v5 uint = 189
        _ = v5
    }
    var v6 uint16 = 320
    if (!v4) {
        var v7 string = ((v2 + "foo bar") + v2)
        v8 := GenS1{F0: 0}
        switch "" {
        case "":
            for fi9 := 0; fi9 < 1; fi9++ {
                fmt.Println(v4, v2)
                _ = fi9
            }
        case "\n\t":
            for fi10 := 0; fi10 < 13; fi10++ {
                switch ((v2 + (v7 + v7)) + v7) {
                case "a":
                    switch v2 {
                    case "":
                        _ = len(v1)
                    case "hello":
                        v6 = (v6 / ((v6) | 1))
                    case "foo bar":
                        fmt.Println(v6, v2)
                    }
                case "":
                    var v11 int32 = 372
                    _ = v11
                default:
                    var v12 string = v7
                    _ = v12
                }
                _ = fi10
            }
        case "hello":
            var v13 bool = v4
            _ = v13
        default:
            if (0 != 419) {
                var v14 int16 = 1
                fmt.Println(v3)
                _ = v14
            }
        }
        _ = v7
        _ = v8
    }
    v15 := GenS2{F0: 2147483648, F1: 2147483647}
    var v16 uint32 = 0
    v16 -= v16
    v3 = (v3 - v3)
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint = 0
        _ = fi17
        _ = v18
    }
    fmt.Println(v4, v4)
    v19 := GenS2{F0: v16, F1: (^(v3))}
    var v20 int64 = 42
    if v4 {
        if (673 <= 1) {
            v3 = -711
        } else {
            var v21 float64 = 3.14159
            var v22 int8 = 1
            _ = v21
            _ = v22
        }
    } else {
        var v23 int = -429
        _ = v23
    }
    var v24 bool = v4
    var v25 uint32 = 2147483647
    fmt.Println(v3, v4)
    for fi26 := 0; fi26 < 6; fi26++ {
        switch v2 {
        case "\n\t":
            var v27 int8 = -128
            _ = v27
        case "foo bar":
            for fi28 := 0; fi28 < 5; fi28++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi28
            }
        case "a":
            var v29 uint = 0
            _ = v29
        }
        v2 = ("hello" + v2)
        var v30 uint = 4294967295
        _ = fi26
        _ = v30
    }
    v25 -= 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
}
