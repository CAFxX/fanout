package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint8; F2 int32 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int16; F1 int8; F2 int64 }

func main() {
    var v1 bool = false
    var v2 int8 = 0
    v3 := make(map[string]bool)
    v4 := []string{("a" + "")}
    var v5 int32 = -2147483648
    if ((1e10 >= float32(-2.25)) && v1) {
        v6 := GenS1{F0: 255, F1: 42, F2: (float32((((1.5 - (0.1 + (0.0 + (0.1 * 0.0)))) + 3.14159) + 0.0)) == -2.25)}
        _ = v6
    } else {
        var v7 bool = v1
        switch "hello" {
        case "hello":
            v1 = (127 < 97)
        }
        _ = len(v4)
        _ = v7
    }
    if v1 {
        v8 := GenS2{F0: (1 < -32768), F1: 255}
        _ = v8
    } else {
        fmt.Println(v1)
        var v9 int32 = 2147483647
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS3{F0: 1, F1: 2, F2: (v5 * v5)}
        v1 = (32767 < 32767)
        v12 := v3["hello"]
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        v5 += (v5 - v5)
        _ = fi13
    }
    v1 = false
    if true {
        var v14 uint32 = 2147483648
        v15 := GenS1{F0: 42, F1: 102, F2: v1}
        _ = len(v4)
        _ = v14
        _ = v15
    } else {
        switch 2147483647 {
        case -1:
            v2 += (v2 << ((2) & 63))
        case -42:
            var v16 uint8 = 0
            _ = v16
        case 4294967295:
            var v17 int32 = 0
            _ = v17
        }
        var v18 float32 = 0.0
        v5 += v5
        _ = v18
    }
    v5 = (v5 % ((v5) | 1))
    v5 -= v5
    v4 = append(v4, "a")
    v2 = (v2 % ((v2) | 1))
    v19 := GenS1{F0: 658, F1: 128, F2: false}
    var v20 uint16 = 42
    var v21 float32 = 0.1
    if (4294967295 >= 42) {
        if (!v1) {
            v22 := GenS4{F0: ((-1 >= -32768) && ((2 != 290) && false)), F1: "\n\t"}
            fmt.Println(v1, v1)
            _ = v22
        }
        switch (("" + "foo bar") + "foo bar") {
        case "\n\t":
            if true {
                var v23 uint = 456
                if false {
                    var v24 float32 = v21
                    _ = v24
                } else {
                    v25 := GenS4{F0: (0 != 42), F1: "\n\t"}
                    for fi26 := 0; fi26 < 6; fi26++ {
                        switch fi26 {
                        case 478:
                            for fi27 := 0; fi27 < 7; fi27++ {
                                fmt.Println(v1)
                                switch ("" + "") {
                                case "":
                                    fmt.Println(v21, v5, v21)
                                default:
                                    v5 += (-(v5))
                                }
                                _ = fi27
                            }
                        }
                        _ = fi26
                    }
                    _ = v25
                }
                v28 := GenS2{F0: false, F1: (v20 >> (((v23 | 18446744073709551615)) & 63))}
                _ = v23
                _ = v28
            } else {
                v21 = -2.25
            }
        case "a":
            v5 = v5
        case "hello":
            delete(v3, "\n\t")
        }
    }
    v4 = append(v4, "foo bar")
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 11; fi30++ {
            if v1 {
                v31 := GenS2{F0: false, F1: (v20 | v20)}
                v21 += v21
                _ = len(v3)
                _ = v31
            } else {
                if (((v5 >> ((99) & 63)) > v5) || false) {
                    if v1 {
                        fmt.Println(v5, fi29)
                        switch fi30 {
                        case 42:
                            var v32 uint16 = 42
                            _ = v32
                        case 4294967295:
                            var v33 int64 = 9223372036854775807
                            _ = v33
                        default:
                            v34 := GenS4{F0: (v1 || (v1 || (128 != 0))), F1: ("hello" + ("" + ""))}
                            _ = v34
                        }
                    }
                    v21 -= v21
                } else {
                    switch "a" {
                    case "\n\t":
                        v2 = v2
                    case "foo bar":
                        var v35 int = (fi29 + fi29)
                        v36 := GenS4{F0: v1, F1: "hello"}
                        _ = v35
                        _ = v36
                    }
                }
                var v37 uint32 = uint32(v20)
                v38 := GenS4{F0: v1, F1: "foo bar"}
                _ = v37
                _ = v38
            }
            var v39 uint = 65535
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi30
            _ = v39
        }
        if (-9223372036854775808 <= 9223372036854775807) {
            v40 := GenS5{F0: 1000, F1: v2, F2: 9223372036854775807}
            switch -9223372036854775808 {
            case 763:
                if false {
                    fmt.Println(v1, v2)
                    v20 += (^(v20))
                } else {
                    var v41 int64 = 9223372036854775807
                    v41 = v41
                    fmt.Println(v20, v21, v41)
                    _ = v41
                }
            default:
                var v42 uint = 18446744073709551615
                _ = v42
            }
            _ = v40
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
