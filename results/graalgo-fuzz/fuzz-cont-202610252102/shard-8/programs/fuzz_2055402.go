package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int16 }

func main() {
    v1 := []int{0}
    var v2 uint32 = 1
    v3 := make(map[uint32]string)
    v4 := []string{("hello" + ("hello" + "\n\t")), "foo bar", "a", "a"}
    v5 := GenS1{F0: 1, F1: true}
    var v6 uint64 = 1
    var v7 uint32 = (^(v2))
    var v8 float64 = -2.25
    v9 := GenS2{F0: (v2 % ((0) | 1)), F1: v7}
    v10 := GenS2{F0: 889, F1: v7}
    switch ((("\n\t" + "") + ("a" + (("\n\t" + ("" + "a")) + "\n\t"))) + "\n\t") {
    case "hello":
        var v11 int = -9223372036854775808
        _ = v11
    case "":
        v12 := GenS3{F0: v6}
        _ = v12
    case "foo bar":
        v13 := GenS4{F0: v6, F1: -2147483648}
        _ = v13
    default:
        for fi14 := 0; fi14 < 14; fi14++ {
            switch "a" {
            case "\n\t":
                for fi15 := 0; fi15 < 14; fi15++ {
                    fmt.Println(v8, v2, fi15)
                    v2 = v2
                    _ = fi15
                }
            case "foo bar":
                if false {
                    var v16 int8 = 64
                    v17 := GenS5{F0: 2, F1: 9223372036854775807}
                    _ = v16
                    _ = v17
                } else {
                    if (-1 > -32768) {
                        for fi18 := 0; fi18 < 9; fi18++ {
                            v19 := GenS6{F0: 0}
                            v20 := v3[(v7 % ((v7) | 1))]
                            var v21 int = (fi18 * fi14)
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    } else {
                        if (!true) {
                            var v22 uint8 = 2
                            fmt.Println(v22)
                            _ = v22
                        }
                    }
                    v2 = 42
                }
            case "":
                for fi23 := 0; fi23 < 6; fi23++ {
                    fmt.Println(v2, fi14)
                    var v24 uint8 = 73
                    v6 += v6
                    _ = fi23
                    _ = v24
                }
            default:
                if len(v4) > 0 {
                    v4[0] = ""
                }
            }
            var v25 string = "\n\t"
            _ = fi14
            _ = v25
        }
    }
    v6 -= v6
    if false {
        switch "foo bar" {
        case "\n\t":
            var v26 uint16 = 784
            _ = v26
        case "hello":
            v6 = 42
        case "foo bar":
            v8 = 3.14159
        }
        v7 += 932
    }
    switch ("" + "") {
    case "":
        v27 := GenS7{F0: 0, F1: 480}
        _ = v27
    case "foo bar":
        if (!false) {
            if true {
                fmt.Println(v6, v6, v6)
                var v28 uint16 = 305
                _ = v28
            }
            _ = len(v4)
            var v29 int32 = 54
            _ = v29
        } else {
            switch "\n\t" {
            case "hello":
                var v30 bool = false
                _ = v30
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                var v31 string = "hello"
                _ = v31
            }
            v6 += 990
        }
    case "hello":
        var v32 bool = (!false)
        _ = v32
    default:
        fmt.Println(v7)
    }
    v7 = v2
    v33 := GenS8{F0: 255}
    for fi34 := 0; fi34 < 12; fi34++ {
        _ = len(v3)
        v6 = 18446744073709551615
        v8 += v8
        _ = fi34
    }
    v35 := GenS9{F0: 992}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v36 int16 = 0
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    v2 -= (v2 << ((441) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
