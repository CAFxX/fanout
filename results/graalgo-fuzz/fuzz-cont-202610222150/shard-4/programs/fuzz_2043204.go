package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int64; F2 int32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    var v4 bool = (v1 && (1 != 42))
    v5 := []string{"a", "a", "foo bar", "\n\t"}
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint64 = 18446744073709551615
        _ = fi6
        _ = v7
    }
    if (-2.25 <= (-2.25 + 0.0)) {
        if (0 < -170) {
            if true {
                v4 = (((v1 || v4) && (false || (0 == 985))) || ((1.5 - (((float32(((((1.5 + (-2.25 / 1.0)) * 0.1) / 0.1) / (-2.25 - (((1e10 + 1e10) - 0.0) * 1e10)))) - 1.5) + -2.25) - float32(0.1))) != 1.5))
                _ = len(v3)
            } else {
                v8 := GenS1{F0: 30, F1: 42, F2: 65535}
                v9 := GenS1{F0: 127, F1: -158, F2: 255}
                _ = v8
                _ = v9
            }
            switch 484 {
            case 127:
                fmt.Println(v1)
            case 42:
                var v10 int16 = 32767
                _ = v10
            case -803:
                for fi11 := 0; fi11 < 20; fi11++ {
                    v4 = v4
                    _ = fi11
                }
            }
        }
        if v1 {
            for fi12 := 0; fi12 < 8; fi12++ {
                switch fi12 {
                case -52:
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v4 = (!((true && (126 < -128)) && v4))
                        fmt.Println(v4)
                        _ = fi13
                    }
                }
                v14 := v3[335]
                if (false && v14) {
                    var v15 uint16 = 65535
                    if true {
                        v16 := GenS2{F0: 176, F1: -6, F2: -208}
                        for fi17 := 0; fi17 < 7; fi17++ {
                            v15 -= 255
                            switch 9223372036854775807 {
                            case -409:
                                var v18 uint64 = 18446744073709551615
                                _ = v18
                            }
                            _ = len(v3)
                            var v19 int = fi12
                            switch ("" + "a") {
                            case "hello":
                                fmt.Println(v4, v19)
                            default:
                                v19 = v19
                            }
                            _ = fi17
                            _ = v19
                        }
                        v20 := GenS2{F0: 5, F1: -90, F2: 572}
                        _ = v16
                        _ = v20
                    }
                    var v21 int64 = 740
                    _ = v15
                    _ = v21
                } else {
                    v1 = true
                }
                _ = fi12
                _ = v14
            }
        }
        v22 := GenS2{F0: 0, F1: -9223372036854775808, F2: 1}
        _ = v22
    }
    if v4 {
        fmt.Println(v1)
        var v23 string = "hello"
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v23
    }
    if (!(4 == 0)) {
        v24 := GenS1{F0: 1, F1: -440, F2: 0}
        if (0.1 != (float32(3.14159) + 1.5)) {
            var v25 bool = v4
            _ = v25
        }
        _ = v24
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        v1 = (-128 < 100)
        _ = fi26
    }
    if true {
        var v27 float32 = 3.14159
        _ = v27
    } else {
        switch 9223372036854775807 {
        case 2147483647:
            v28 := GenS1{F0: 100, F1: 944, F2: 2}
            _ = v28
        case -2147483648:
            v4 = v1
        }
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 int = 4294967295
            _ = fi29
            _ = v30
        }
    }
    var v31 uint64 = 2
    var v32 int64 = 532
    v1 = (!false)
    delete(v3, 2147483648)
    v1 = v1
    v33 := GenS3{F0: (2147483647 <= 837)}
    v4 = (!true)
    v34 := GenS3{F0: v4}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
