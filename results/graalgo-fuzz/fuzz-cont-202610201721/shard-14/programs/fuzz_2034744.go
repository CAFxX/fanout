package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint8; F1 int; F2 int }
func genFn0() uint32 {
    v4 := []int{745, 0}
    if true {
        var v5 int = 127
        v5 += 42
        _ = v5
    } else {
        if len(v4) > 0 {
            v4[0] = -2147483648
        }
    }
    var v6 int = 1
    _ = v4
    _ = v6
    return 1
}

func main() {
    v1 := []string{(("foo bar" + "") + "foo bar")}
    var v2 uint64 = 301
    var v3 uint = 42
    switch -155 {
    case -9223372036854775808:
        var v7 int8 = -48
        _ = v7
    case -2147483648:
        fmt.Println(v3)
    }
    _ = len(v1)
    if true {
        if true {
            switch "foo bar" {
            case "a":
                for fi8 := 0; fi8 < 4; fi8++ {
                    for fi9 := 0; fi9 < 18; fi9++ {
                        switch (("" + "a") + "a") {
                        case "hello":
                            if (0 <= 2147483648) {
                                var v10 uint = 18446744073709551615
                                v1 = append(v1, "")
                                var v11 uint8 = 2
                                _ = v10
                                _ = v11
                            } else {
                                var v12 float32 = -2.25
                                v3 = v3
                                var v13 uint8 = 127
                                _ = v12
                                _ = v13
                            }
                        case "foo bar":
                            v3 = (^(v3))
                        }
                        v14 := GenS1{F0: -6, F1: "hello"}
                        v3 -= (v3 << ((v3) & 63))
                        _ = fi9
                        _ = v14
                    }
                    _ = fi8
                }
            case "hello":
                var v15 uint32 = 42
                _ = v15
            case "\n\t":
                var v16 uint64 = v2
                _ = v16
            default:
                v2 = (v2 + uint64(v3))
            }
        } else {
            if true {
                fmt.Println(v2)
                v2 += 42
            } else {
                var v17 uint = 0
                if (true && (127 > 42)) {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        _ = len(v1)
                        _ = fi18
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v19 int64 = 150
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        switch "" {
                        case "\n\t":
                            v17 = (v3 / ((v17) | 1))
                        case "a":
                            v2 = (v2 / ((uint64(v17)) | 1))
                        case "hello":
                            v21 := GenS2{F0: 255, F1: 32767, F2: -9223372036854775808}
                            _ = v21
                        default:
                            var v22 string = "\n\t"
                            _ = v22
                        }
                        _ = fi20
                    }
                }
                _ = v17
            }
            var v23 uint8 = 216
            _ = v23
        }
        v24 := GenS1{F0: -128, F1: "\n\t"}
        _ = v24
    } else {
        var v25 int = -380
        var v26 uint64 = 983
        _ = v25
        _ = v26
    }
    var v27 uint32 = 0
    v28 := GenS1{F0: -121, F1: "hello"}
    if (!(223 <= 1)) {
        var v29 float32 = (1e10 * (((float32((0.0 - 0.0)) / ((3.14159 - 0.0) - (1e10 - 0.1))) / 1.0) * (-2.25 / 1e10)))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v29
    }
    if (32768 == 534) {
        var v30 int = -348
        _ = v30
    }
    var v31 float32 = 3.14159
    var v32 int16 = 423
    v31 = float32(1.5)
    for fi33 := 0; fi33 < 5; fi33++ {
        v31 = (0.1 - v31)
        v34 := GenS1{F0: -123, F1: "hello"}
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 int8 = 127
        var v37 uint64 = (^(v2))
        _ = fi35
        _ = v36
        _ = v37
    }
    if len(v1) > 0 {
        v1[0] = ("hello" + ("hello" + ""))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
}
