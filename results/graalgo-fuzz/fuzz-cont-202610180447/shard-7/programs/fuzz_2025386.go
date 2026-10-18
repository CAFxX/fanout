package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int64 = 642
    var v2 string = ""
    var v3 uint = 42
    v4 := []int{-137, 505}
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 float32 = (1e10 / 3.14159)
        v7 := GenS1{F0: 0}
        _ = fi5
        _ = v6
        _ = v7
    }
    if ((true || (false && true)) || false) {
        fmt.Println(v2, v1)
        var v8 uint8 = 42
        switch (v2 + v2) {
        case "\n\t":
            var v9 int32 = 420
            _ = v9
        case "a":
            if (true || false) {
                if (!false) {
                    fmt.Println(v3, v2)
                } else {
                    v10 := GenS1{F0: 0}
                    v2 = ("foo bar" + "\n\t")
                    _ = v10
                }
                v3 -= v3
                v1 -= -9223372036854775808
            }
        case "":
            v8 = 1
        default:
            if true {
                v1 = v1
                if false {
                    fmt.Println(v1, v8, v2)
                } else {
                    v2 = "a"
                    switch -668 {
                    case -1:
                        v8 = 42
                    }
                }
                switch (v2 + v2) {
                case "\n\t":
                    v8 += (^(v8))
                case "foo bar":
                    var v11 float32 = 1.5
                    _ = v11
                }
            } else {
                switch v2 {
                case "":
                    fmt.Println(v1, v3, v1)
                default:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v12 := GenS1{F0: 42}
                v13 := GenS1{F0: 571}
                _ = v12
                _ = v13
            }
        }
        _ = v8
    } else {
        var v14 bool = (!false)
        _ = v14
    }
    if true {
        if (0 >= 127) {
            if (802 <= 4294967295) {
                switch 42 {
                case 880:
                    for fi15 := 0; fi15 < 15; fi15++ {
                        v16 := GenS1{F0: (-(fi15))}
                        _ = fi15
                        _ = v16
                    }
                case -435:
                    if len(v4) > 0 {
                        v4[0] = 32768
                    }
                case 32767:
                    v2 = v2
                default:
                    fmt.Println(v3, v2)
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if false {
                fmt.Println(v2, v1)
                v17 := GenS2{F0: -128, F1: 127}
                _ = v17
            }
        }
    }
    var v18 float32 = -2.25
    v4 = append(v4, -9223372036854775808)
    var v19 float32 = (v18 + 1.5)
    v20 := GenS1{F0: 32767}
    if false {
        switch v2 {
        case "":
            v4 = append(v4, 32767)
        case "\n\t":
            v4 = append(v4, 623)
        }
        var v21 uint64 = 42
        _ = v21
    } else {
        v22 := GenS2{F0: 42, F1: -128}
        var v23 int8 = 0
        _ = v22
        _ = v23
    }
    v2 = ("foo bar" + "\n\t")
    v18 += v18
    fmt.Println(v1)
    fmt.Println(v18, v19, v18)
    var v24 int16 = 1
    switch ((v2 + "") + "foo bar") {
    case "":
        v4 = append(v4, int(v1))
    case "\n\t":
        if true {
            var v25 uint = 22
            v26 := GenS3{F0: 4294967295}
            _ = v25
            _ = v26
        } else {
            var v27 string = v2
            v28 := GenS2{F0: -65, F1: -128}
            _ = v27
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
