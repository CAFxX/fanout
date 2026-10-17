package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint32; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 int64; F2 int16 }
type GenS6 struct { F0 uint32; F1 string }

func main() {
    var v1 int64 = 967
    v2 := []string{}
    v3 := make(map[uint32]int)
    var v4 uint = 1
    v5 := []string{"\n\t"}
    var v6 string = ((("a" + "foo bar") + ("" + "")) + ("" + "\n\t"))
    if false {
        v4 += 255
        fmt.Println(v4)
        v4 = v4
    } else {
        v6 = v6
    }
    var v7 float64 = -2.25
    switch 32767 {
    case 50:
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: v6}
            if true {
                delete(v3, 380)
            } else {
                v10 := GenS2{F0: -41, F1: 42, F2: 4294967295}
                if true {
                    var v11 uint = 953
                    _ = v11
                }
                if (!false) {
                    var v12 int16 = -87
                    var v13 int8 = 42
                    if (42 <= 321) {
                        v14 := GenS3{F0: 240, F1: 236, F2: v6}
                        for fi15 := 0; fi15 < 2; fi15++ {
                            var v16 uint64 = 18446744073709551615
                            if true {
                                v17 := GenS2{F0: v13, F1: 2, F2: v16}
                                _ = v17
                            } else {
                                var v18 int64 = -9223372036854775808
                                switch (v6 + "a") {
                                case "":
                                    if (false || true) {
                                        v1 = (v1 - 116)
                                    } else {
                                        fmt.Println(v16, v12)
                                        var v19 uint8 = 128
                                        var v20 int32 = 2147483647
                                        _ = v19
                                        _ = v20
                                    }
                                }
                                _ = v18
                            }
                            _ = fi15
                            _ = v16
                        }
                        if true {
                            delete(v3, 1)
                        } else {
                            fmt.Println(v7, v7)
                            var v21 uint = v4
                            _ = v21
                        }
                        _ = v14
                    } else {
                        switch -12 {
                        case 32767:
                            if (("\n\t" + (("hello" + v6) + v6)) >= v6) {
                                v22 := GenS4{F0: 47}
                                _ = v22
                            } else {
                                var v23 int16 = v12
                                var v24 int = 127
                                var v25 string = "hello"
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        case 190:
                            fmt.Println(v12, v1, v13)
                        case 4294967295:
                            switch fi8 {
                            case 1:
                                fmt.Println(v4, v7, v4)
                            default:
                                var v26 bool = (124 >= 42)
                                _ = v26
                            }
                        }
                    }
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
            var v27 uint = (v4 * 0)
            _ = fi8
            _ = v27
            _ = v9
        }
    case 2147483647:
        v28 := GenS5{F0: 502, F1: (v1 ^ 922), F2: 1000}
        _ = v28
    }
    v4 = v4
    _ = len(v3)
    v29 := GenS6{F0: 175, F1: v6}
    v4 -= (v4 << ((18446744073709551615) & 63))
    v7 = v7
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
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
}
