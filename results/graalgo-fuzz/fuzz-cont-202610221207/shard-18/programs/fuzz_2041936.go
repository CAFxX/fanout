package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    var v3 uint8 = 2
    var v4 uint = 42
    var v5 int16 = 242
    v6 := make(map[uint32]bool)
    v7 := []int{455, 128}
    v8 := GenS1{F0: v5}
    v3 -= (^(v3))
    fmt.Println(v5)
    if (v4 != (v4 + 2)) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v9 uint = (v4 >> ((v4) & 63))
        fmt.Println(v9, v9)
        _ = v9
    }
    if false {
        var v10 float32 = (0.0 + 0.0)
        var v11 uint16 = 0
        _ = v10
        _ = v11
    } else {
        var v12 int64 = 885
        var v13 int32 = 42
        fmt.Println(v12, v3)
        _ = v12
        _ = v13
    }
    switch ("\n\t" + "\n\t") {
    case "a":
        switch "foo bar" {
        case "hello":
            if false {
                v14 := GenS2{F0: uint32(v3), F1: 120}
                v6[80] = false
                if (!(true || false)) {
                    var v15 bool = true
                    for fi16 := 0; fi16 < 2; fi16++ {
                        if (1 == 42) {
                            v1[("a" + "foo bar")] = 525
                        }
                        for fi17 := 0; fi17 < 3; fi17++ {
                            v15 = true
                            var v18 uint32 = 42
                            v19 := GenS3{F0: fi17, F1: 1}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        _ = len(v6)
                        _ = fi16
                    }
                    _ = v15
                } else {
                    var v20 bool = true
                    var v21 uint64 = 4294967295
                    _ = v20
                    _ = v21
                }
                _ = v14
            }
        default:
            var v22 int = 32768
            _ = v22
        }
    case "hello":
        if false {
            if (((!false) && (791 > 917)) && false) {
                v4 = (^(v4))
            } else {
                var v23 bool = true
                for fi24 := 0; fi24 < 6; fi24++ {
                    v25 := GenS2{F0: 65535, F1: 42}
                    switch "hello" {
                    case "a":
                        switch fi24 {
                        case -1:
                            var v26 uint = (v4 << ((v4) & 63))
                            _ = v26
                        case 0:
                            _ = len(v6)
                        }
                    case "hello":
                        var v27 bool = (!v23)
                        _ = v27
                    case "\n\t":
                        var v28 uint8 = v3
                        _ = v28
                    default:
                        v3 += (v3 % ((v3) | 1))
                    }
                    _ = len(v7)
                    _ = fi24
                    _ = v25
                }
                fmt.Println(v5, v23, v3)
                _ = v23
            }
            var v29 uint = (v4 ^ 255)
            _ = v29
        }
    case "":
        switch ("hello" + "hello") {
        case "\n\t":
            var v30 int16 = -32768
            _ = v30
        default:
            fmt.Println(v5)
        }
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        switch (("\n\t" + ("" + "hello")) + "hello") {
        case "\n\t":
            var v32 string = "a"
            _ = v32
        default:
            if ((((-2.25 / 3.14159) * -2.25) / 0.1) > (0.0 / (((0.1 + 3.14159) - ((3.14159 / (((0.1 + 1e10) - ((((1.5 * 3.14159) / 1.5) - 1e10) + 0.0)) * 0.1)) - 0.0)) + 1e10))) {
                v1[""] = 2147483647
            } else {
                switch -9223372036854775808 {
                case 42:
                    for fi33 := 0; fi33 < 2; fi33++ {
                        var v34 uint64 = 356
                        _ = fi33
                        _ = v34
                    }
                default:
                    v5 -= v5
                }
            }
        }
        if len(v7) > 0 {
            v7[0] = 42
        }
        var v35 int64 = 431
        _ = fi31
        _ = v35
    }
    v36 := GenS4{F0: v5}
    for fi37 := 0; fi37 < 11; fi37++ {
        fmt.Println(fi37)
        var v38 int16 = v5
        _ = fi37
        _ = v38
    }
    var v39 uint64 = 18446744073709551615
    switch 783 {
    case 4294967295:
        v3 -= (v3 >> ((42) & 63))
    }
    v39 = (v39 % ((uint64(v4)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
