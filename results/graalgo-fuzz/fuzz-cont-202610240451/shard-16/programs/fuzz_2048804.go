package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int16; F2 string }
func genFn0(p0 uint16) uint {
    var v4 bool = false
    v5 := []int{}
    if true {
        var v6 uint32 = 2147483647
        _ = v6
    } else {
        var v7 string = "hello"
        if len(v5) > 0 {
            v5[0] = 656
        }
        _ = v7
    }
    fmt.Println(v4)
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1() int64 {
    switch (("" + ("a" + "")) + "hello") {
    case "\n\t":
        var v8 uint32 = 265
        _ = v8
    default:
        v9 := GenS1{F0: "foo bar"}
        _ = v9
    }
    v10 := []int{128, 0, 32767}
    _ = v10
    return 487
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 13
    v3 := make(map[int]uint32)
    fmt.Println(v2)
    if false {
        switch "\n\t" {
        case "":
            fmt.Println(v2)
        case "\n\t":
            v2 += 255
        case "a":
            v11 := GenS2{F0: -200}
            _ = v11
        default:
            _ = len(v1)
        }
        switch 42 {
        case 1:
            v12 := GenS3{F0: ("\n\t" + ((("" + "a") + "\n\t") + "\n\t")), F1: -787, F2: "foo bar"}
            _ = v12
        case 9223372036854775807:
            if (65535 < 65535) {
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    fmt.Println(v2)
                case "a":
                    if true {
                        if false {
                            for fi13 := 0; fi13 < 19; fi13++ {
                                v14 := GenS3{F0: "hello", F1: -32768, F2: "hello"}
                                var v15 uint = 261
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                        } else {
                            _ = len(v3)
                            var v16 float64 = 1.5
                            var v17 uint8 = (^(v2))
                            _ = v16
                            _ = v17
                        }
                        v18 := GenS1{F0: "\n\t"}
                        var v19 int8 = 125
                        _ = v18
                        _ = v19
                    }
                }
            } else {
                v2 = (v2 | 42)
                var v20 uint16 = 895
                _ = v20
            }
        case -1:
            _ = len(v1)
        }
        v2 += (v2 >> ((1) & 63))
    }
    v2 += v2
    switch "foo bar" {
    case "foo bar":
        var v21 float32 = 1e10
        _ = v21
    default:
        v2 -= (^(v2))
    }
    v22 := []int{32768, 2147483647}
    var v23 string = ""
    switch 4294967295 {
    case 9223372036854775807:
        var v24 uint32 = 4294967295
        _ = v24
    }
    var v25 float32 = -2.25
    if (42 <= 42) {
        var v26 int32 = -374
        _ = v26
    } else {
        _ = len(v1)
        if (!true) {
            v2 += v2
            var v27 int32 = 0
            _ = v27
        }
        var v28 uint64 = 979
        _ = v28
    }
    v25 += 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
}
