package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    var v1 int = 287
    var v2 int32 = 2147483647
    var v3 int32 = (v2 % ((596) | 1))
    v4 := make(map[int]bool)
    v5 := []int{}
    var v6 string = ""
    v2 = (v3 >> ((1) & 63))
    v5 = append(v5, (v1 ^ (v1 >> ((206) & 63))))
    _ = len(v4)
    if true {
        var v7 int64 = -290
        var v8 int16 = -279
        _ = v7
        _ = v8
    }
    v5 = append(v5, int(v3))
    if (true || true) {
        v4[32767] = (!(-2.25 <= 3.14159))
        v1 += -208
    }
    switch "foo bar" {
    case "\n\t":
        v9 := GenS1{F0: 100, F1: 127}
        _ = v9
    case "foo bar":
        switch (v1 ^ (v1 ^ (v1 - (v1 | (-(v1)))))) {
        case 32768:
            v2 = -1
        case 0:
            switch 42 {
            case 128:
                for fi10 := 0; fi10 < 16; fi10++ {
                    if (!(!false)) {
                        var v11 string = (v6 + v6)
                        for fi12 := 0; fi12 < 9; fi12++ {
                            if (!true) {
                                v1 = fi10
                            } else {
                                var v13 uint16 = 32767
                                switch ("" + ((v11 + "hello") + v11)) {
                                case "foo bar":
                                    v11 = ""
                                case "\n\t":
                                    if true {
                                        switch ("" + v6) {
                                        case "hello":
                                            v3 += (v2 << ((485) & 63))
                                            v14 := GenS1{F0: -1, F1: 255}
                                            _ = v14
                                        case "foo bar":
                                            v15 := v4[748]
                                            _ = v15
                                        }
                                        var v16 int64 = 42
                                        _ = v16
                                    }
                                case "a":
                                    if ((3.14159 * float64(0.1)) <= -2.25) {
                                        v2 = (v2 / ((v2) | 1))
                                        v4[32768] = ((v11 + "hello") > ("hello" + v6))
                                    }
                                default:
                                    var v17 uint32 = 2
                                    _ = v17
                                }
                                fmt.Println(v1, v2, v6)
                                _ = v13
                            }
                            switch fi10 {
                            case 65:
                                v5 = append(v5, fi10)
                            }
                            fmt.Println(v3, v11, fi10)
                            v6 = ("" + v6)
                            v11 = "a"
                            _ = fi12
                        }
                        _ = v11
                    } else {
                        var v18 float32 = ((3.14159 / 1.0) - 1e10)
                        switch (v1 - 4294967295) {
                        case 766:
                            v19 := GenS1{F0: 42, F1: 128}
                            _ = v19
                        case 42:
                            var v20 float64 = (0.1 - 0.0)
                            _ = v20
                        case 1:
                            var v21 int = 32768
                            _ = v21
                        }
                        _ = v18
                    }
                    _ = fi10
                }
            default:
                v22 := GenS1{F0: 1, F1: 2}
                _ = v22
            }
        case -351:
            v6 = ""
        }
    default:
        var v23 int32 = v2
        _ = v23
    }
    var v24 uint = 298
    v5 = append(v5, (-(v1)))
    if false {
        var v25 string = "\n\t"
        var v26 bool = (-2.25 != float32(0.0))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v25
        _ = v26
    } else {
        v3 = v2
    }
    if (-1 < 0) {
        fmt.Println(v1)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v24, v3)
    }
    v27 := GenS2{F0: 78, F1: true}
    v24 = v24
    switch v1 {
    case -2147483648:
        var v28 int64 = 9223372036854775807
        _ = v28
    case -279:
        if len(v5) > 0 {
            v5[0] = v1
        }
    case -9223372036854775808:
        var v29 int16 = 211
        _ = v29
    }
    v4[-213] = true
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
