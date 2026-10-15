package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []string{"", ""}
    var v2 string = "hello"
    var v3 int8 = -128
    var v4 uint64 = 898
    v3 = (v3 + (-(v3)))
    if ((false || (true && (!(2 >= 2)))) || true) {
        var v5 int8 = v3
        _ = len(v1)
        var v6 float32 = float32(1e10)
        _ = v5
        _ = v6
    }
    v4 = 0
    switch 372 {
    case 128:
        v3 -= (v3 ^ (^(v3)))
    case -9223372036854775808:
        switch 42 {
        case 2147483647:
            v4 = v4
        case 128:
            for fi7 := 0; fi7 < 7; fi7++ {
                _ = len(v1)
                _ = fi7
            }
        case 127:
            if true {
                if (170 >= 92) {
                    fmt.Println(v2)
                }
                v8 := GenS1{F0: 128, F1: v4}
                _ = v8
            } else {
                var v9 int64 = -1
                var v10 int = 9223372036854775807
                if false {
                    if true {
                        v2 = ""
                        var v11 bool = (!(65535 != 32767))
                        v3 += v3
                        _ = v11
                    } else {
                        switch ("" + "foo bar") {
                        case "a":
                            if (v2 >= "\n\t") {
                                var v12 uint8 = 128
                                var v13 uint = 1
                                switch v10 {
                                case 4294967295:
                                    v3 -= 127
                                case 889:
                                    fmt.Println(v13)
                                }
                                _ = v12
                                _ = v13
                            } else {
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    v14 := GenS1{F0: 0, F1: (v4 | (v4 ^ 2))}
                                    var v15 int64 = 615
                                    v16 := GenS1{F0: 39, F1: (v4 - v4)}
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                            }
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v9 = (v9 + v9)
                    }
                    v17 := GenS1{F0: 0, F1: v4}
                    _ = v17
                }
                _ = v10
                _ = v9
            }
        default:
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 int32 = 0
                var v20 int8 = -1
                for fi21 := 0; fi21 < 9; fi21++ {
                    v22 := GenS1{F0: 112, F1: v4}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch fi21 {
                    case 349:
                        fmt.Println(v4, v2)
                    case 42:
                        v23 := GenS2{F0: 255}
                        _ = v23
                    case -365:
                        var v24 string = ""
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    case 1:
        v3 += (v3 | 41)
    }
    var v25 uint = 0
    var v26 int32 = 0
    switch ((v2 + (("\n\t" + v2) + v2)) + v2) {
    case "hello":
        fmt.Println(v4, v2, v3)
    case "a":
        v4 = v4
    default:
        var v27 int32 = 829
        _ = v27
    }
    v28 := GenS1{F0: 2, F1: v4}
    v3 = (v3 & 100)
    var v29 float32 = (float32((-2.25 / 0.1)) + (0.1 - float32(0.1)))
    switch (v2 + v2) {
    case "":
        v30 := GenS2{F0: uint(v4)}
        _ = v30
    case "hello":
        v31 := GenS1{F0: 37, F1: v4}
        _ = v31
    case "foo bar":
        if (189 >= 249) {
            _ = len(v1)
            v25 = v25
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v1) > 0 {
        v1[0] = v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
