package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint = 65535
    var v2 uint32 = uint32(v1)
    var v3 int = -869
    var v4 uint8 = 128
    var v5 int16 = 42
    var v6 int = 4294967295
    var v7 uint32 = (v2 & (^(v2)))
    v8 := GenS1{F0: 0, F1: (true || false)}
    v9 := GenS2{F0: 42}
    if false {
        v7 = 42
    }
    switch v3 {
    case 127:
        v10 := []int{}
        _ = v10
    case 2147483647:
        v2 = 154
    default:
        var v11 float64 = 1.5
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        if true {
            v13 := GenS1{F0: -684, F1: false}
            fmt.Println(v4, v1)
            var v14 int8 = 1
            _ = v13
            _ = v14
        } else {
            v1 = (v1 >> ((18446744073709551615) & 63))
            v3 -= (v3 - 42)
            switch "foo bar" {
            case "hello":
                var v15 float32 = -2.25
                _ = v15
            case "a":
                for fi16 := 0; fi16 < 13; fi16++ {
                    switch (fi16 / (((v3 + -654)) | 1)) {
                    case 241:
                        var v17 uint32 = v7
                        _ = v17
                    }
                    v2 += (^(v2))
                    var v18 uint = 1
                    var v19 uint = (v1 - 65535)
                    switch "hello" {
                    case "":
                        var v20 int16 = 1
                        _ = v20
                    case "a":
                        if true {
                            v21 := GenS2{F0: (v5 | v5)}
                            _ = v21
                        } else {
                            v5 = v5
                            var v22 float64 = float64((0.0 * -2.25))
                            if true {
                                fmt.Println(fi16, v5, v2)
                            }
                            _ = v22
                        }
                    default:
                        var v23 bool = (!(((^(v2)) != (v7 - 446)) && false))
                        _ = v23
                    }
                    _ = fi16
                    _ = v18
                    _ = v19
                }
            case "\n\t":
                v7 = (v2 ^ (^(v7)))
            default:
                switch (("\n\t" + "hello") + ("hello" + ("a" + "hello"))) {
                case "foo bar":
                    for fi24 := 0; fi24 < 6; fi24++ {
                        v2 = (^(v7))
                        _ = fi24
                    }
                case "":
                    fmt.Println(fi12, v1, v3)
                default:
                    if (-1 == 833) {
                        fmt.Println(v1)
                    } else {
                        for fi25 := 0; fi25 < 17; fi25++ {
                            var v26 uint = 18446744073709551615
                            _ = fi25
                            _ = v26
                        }
                        var v27 uint16 = 32767
                        var v28 int = 9223372036854775807
                        _ = v27
                        _ = v28
                    }
                }
            }
        }
        v1 = v1
        _ = fi12
    }
    v29 := []int{(v3 >> ((v1) & 63)), v3, int(v1), v3}
    switch "foo bar" {
    case "hello":
        fmt.Println(v6)
    case "":
        v6 -= (v3 | int(v1))
    case "\n\t":
        v1 += (v1 ^ 4294967295)
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 int = (^(v3))
        for fi32 := 0; fi32 < 12; fi32++ {
            if len(v29) > 0 {
                v29[0] = (v31 / ((v31) | 1))
            }
            fmt.Println(v6)
            var v33 int = (fi30 >> ((42) & 63))
            _ = fi32
            _ = v33
        }
        fmt.Println(v5, v1, v5)
        _ = fi30
        _ = v31
    }
    fmt.Println(v7, v1)
    if false {
        if len(v29) > 0 {
            v29[0] = v3
        }
        v2 -= uint32(v4)
        var v34 uint32 = v2
        _ = v34
    } else {
        v1 = (v1 / ((v1) | 1))
        for fi35 := 0; fi35 < 15; fi35++ {
            if len(v29) > 0 {
                v29[0] = v3
            }
            v36 := GenS2{F0: v5}
            _ = fi35
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
