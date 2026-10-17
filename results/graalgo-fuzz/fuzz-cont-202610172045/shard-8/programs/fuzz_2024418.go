package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []uint32{991, 4294967295}
    var v2 float64 = (0.0 / ((1.5 + 1e10) - ((3.14159 - (1e10 / 3.14159)) / -2.25)))
    v3 := make(map[uint32]string)
    var v4 float64 = v2
    v2 = (v2 + (1.5 * (0.1 + v4)))
    v1 = append(v1, 4294967295)
    for fi5 := 0; fi5 < 12; fi5++ {
        if true {
            if true {
                _ = len(v3)
                var v6 int8 = 100
                var v7 int16 = 0
                _ = v6
                _ = v7
            }
            v1 = append(v1, 2)
        } else {
            fmt.Println(v2)
        }
        if (658 >= 42) {
            var v8 int64 = 9223372036854775807
            _ = v8
        }
        _ = fi5
    }
    v2 = (v4 * 1e10)
    switch "hello" {
    case "":
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(v2, v2)
            v2 = v4
            _ = fi9
        }
    case "foo bar":
        switch -9223372036854775808 {
        case -2147483648:
            var v10 int = 42
            _ = v10
        case 32767:
            v4 = (v4 + float64(0.1))
        }
    }
    fmt.Println(v4, v2)
    v1 = append(v1, 65535)
    fmt.Println(v2, v2)
    switch "hello" {
    case "foo bar":
        v2 -= 1e10
    case "a":
        v4 = v2
    case "hello":
        var v11 float64 = v4
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        switch fi12 {
        case 1:
            fmt.Println(v2)
        }
        v4 = float64(((0.0 * (float32(v2) / 1.0)) - float32(((0.1 * v4) / v2))))
        _ = fi12
    }
    v2 += -2.25
    v1 = append(v1, 62)
    v13 := GenS1{F0: 4294967295}
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if ((!true) && ((42 >= -9223372036854775808) || (1 < 4294967295))) {
        if (2 >= 2) {
            if (v2 > v4) {
                v4 = (v4 / float64(1e10))
            } else {
                if (251 > 2) {
                    var v14 uint32 = 2147483648
                    var v15 bool = (2 > 1)
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 15; fi16++ {
                    switch 0 {
                    case 42:
                        fmt.Println(fi16)
                    case 4294967295:
                        for fi17 := 0; fi17 < 6; fi17++ {
                            v18 := v3[434]
                            _ = fi17
                            _ = v18
                        }
                    default:
                        var v19 string = "a"
                        _ = v19
                    }
                    _ = fi16
                }
            }
        } else {
            var v20 uint8 = 128
            _ = v20
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            v4 -= 1.5
            if (1000 > 0) {
                delete(v3, 42)
                v22 := GenS1{F0: 4294967295}
                fmt.Println(v2)
                _ = v22
            } else {
                var v23 int32 = 42
                for fi24 := 0; fi24 < 18; fi24++ {
                    var v25 uint64 = 42
                    if false {
                        var v26 int8 = -27
                        _ = v26
                    } else {
                        fmt.Println(v23, fi21, v25)
                        var v27 float32 = 1.5
                        if true {
                            if true {
                                var v28 float64 = 3.14159
                                v29 := GenS1{F0: (^(v25))}
                                fmt.Println(v25, v23, v28)
                                _ = v28
                                _ = v29
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
                fmt.Println(v4, v4)
                _ = v23
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi21
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            v2 += v4
            _ = fi30
        }
    }
    v2 = v4
    var v31 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
