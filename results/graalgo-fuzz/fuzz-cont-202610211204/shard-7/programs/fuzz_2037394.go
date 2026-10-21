package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 32767
    var v3 int8 = 42
    var v4 float64 = 1e10
    var v5 int64 = 1
    v6 := make(map[uint32]string)
    if true {
        fmt.Println(v2)
    } else {
        fmt.Println(v5, v5, v4)
        var v7 uint = 289
        _ = v7
    }
    var v8 bool = (false || true)
    v9 := []int{32767}
    v8 = v8
    var v10 string = "a"
    var v11 int8 = (-(v3))
    v3 = (-(v3))
    v12 := GenS1{F0: 2147483648, F1: v10}
    v3 += (v11 % (((v11 & v11)) | 1))
    v4 = 0.0
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    switch (v10 + "") {
    case "foo bar":
        fmt.Println(v8, v5)
    case "a":
        v10 = (v10 + v10)
    case "":
        v11 = 13
    default:
        switch (v10 + ("" + "")) {
        case "":
            if (v10 <= "\n\t") {
                switch -9223372036854775808 {
                case 4294967295:
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                case 42:
                    if (2 > 0) {
                        for fi13 := 0; fi13 < 11; fi13++ {
                            if (fi13 > (fi13 + 4294967295)) {
                                for fi14 := 0; fi14 < 14; fi14++ {
                                    v10 = v10
                                    if (!true) {
                                        var v15 uint32 = 65535
                                        v10 = (v10 + v10)
                                        _ = v15
                                    } else {
                                        if (v8 || v8) {
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                        }
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        var v16 uint = 1
                                        var v17 uint64 = 1
                                        _ = v16
                                        _ = v17
                                    }
                                    _ = fi14
                                }
                                switch (fi13 ^ 127) {
                                case 331:
                                    fmt.Println(v10, v3, v8)
                                default:
                                    for fi18 := 0; fi18 < 4; fi18++ {
                                        fmt.Println(v3)
                                        var v19 string = v10
                                        _ = fi18
                                        _ = v19
                                    }
                                }
                                v20 := GenS1{F0: 697, F1: ""}
                                _ = v20
                            } else {
                                v2 -= 806
                            }
                            var v21 uint64 = 1
                            _ = fi13
                            _ = v21
                        }
                    }
                default:
                    var v22 int16 = int16(v2)
                    _ = v22
                }
                fmt.Println(v10, v3)
            } else {
                var v23 int64 = 9223372036854775807
                v5 = v5
                _ = v23
            }
        case "foo bar":
            v24 := GenS1{F0: 420, F1: "hello"}
            _ = v24
        case "hello":
            _ = len(v9)
        default:
            var v25 int = -2147483648
            _ = v25
        }
    }
    fmt.Println(v4, v2)
    v26 := GenS1{F0: 1, F1: v10}
    var v27 int8 = v3
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 float64 = 0.0
        _ = fi28
        _ = v29
    }
    v10 = v10
    v30 := GenS2{F0: v10}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
