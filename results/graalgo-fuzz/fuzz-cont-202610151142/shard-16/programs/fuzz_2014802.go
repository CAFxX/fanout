package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint }

func main() {
    var v1 uint16 = 32767
    var v2 int8 = 1
    var v3 bool = ((^(v1)) <= 42)
    var v4 int32 = -2147483648
    for fi5 := 0; fi5 < 7; fi5++ {
        v3 = v3
        v6 := GenS1{F0: "foo bar"}
        if (((((0.1 / (-2.25 - 1e10)) + float64(float32(float64((((0.0 / 1e10) / -2.25) * (float32(-2.25) * 1.5)))))) * 0.0) + 3.14159) > 0.1) {
            switch "a" {
            case "a":
                v7 := GenS1{F0: "hello"}
                _ = v7
            case "foo bar":
                var v8 int16 = 570
                _ = v8
            case "hello":
                fmt.Println(v3, v1)
            default:
                var v9 float32 = 3.14159
                _ = v9
            }
            var v10 float32 = 0.1
            _ = v10
        } else {
            var v11 int64 = -1
            _ = v11
        }
        _ = fi5
        _ = v6
    }
    v12 := []int{32768, 0, -178, -944}
    if v3 {
        fmt.Println(v2, v1)
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint8 = 169
            if v3 {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                var v15 uint16 = 644
                _ = v15
            } else {
                v3 = true
                _ = len(v12)
            }
            if (!v3) {
                if v3 {
                    var v16 uint64 = 326
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = v16
                }
            } else {
                var v17 uint64 = 0
                var v18 uint = 55
                var v19 uint64 = v17
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi13
            _ = v14
        }
        var v20 int16 = 1
        _ = v20
    } else {
        if v3 {
            v21 := GenS1{F0: "foo bar"}
            var v22 string = "hello"
            v1 -= v1
            _ = v21
            _ = v22
        } else {
            var v23 uint8 = 255
            var v24 int16 = -340
            _ = v23
            _ = v24
        }
        var v25 uint8 = 127
        _ = v25
    }
    v2 -= v2
    if v3 {
        if len(v12) > 0 {
            v12[0] = 2147483647
        }
        if v3 {
            fmt.Println(v1, v4, v1)
            var v26 int = 42
            _ = v26
        } else {
            v27 := GenS1{F0: "a"}
            switch "" {
            case "hello":
                var v28 uint16 = v1
                _ = v28
            case "":
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            default:
                if v3 {
                    var v29 int8 = v2
                    var v30 int8 = (v2 / ((v2) | 1))
                    _ = v29
                    _ = v30
                } else {
                    v3 = (((1.5 * 3.14159) * 0.0) == (((-2.25 + -2.25) / -2.25) + (1.5 / 1.0)))
                }
            }
            _ = v27
        }
    }
    if true {
        switch 9223372036854775807 {
        case 0:
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
    }
    v31 := GenS1{F0: "a"}
    var v32 int = -2147483648
    v33 := GenS2{F0: 32767}
    v32 = v32
    if (985 > 4294967295) {
        if v3 {
            v34 := GenS2{F0: 1}
            _ = v34
        }
        switch (v32 + v32) {
        case 2147483647:
            var v35 int8 = (v2 | (v2 % ((v2) | 1)))
            _ = v35
        }
    }
    v36 := GenS3{F0: 2147483648, F1: 2}
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v4
}
