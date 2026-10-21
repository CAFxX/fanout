package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 bool; F2 uint8 }
type GenS3 struct { F0 int; F1 int8 }

func main() {
    v1 := []uint32{2147483648, 0}
    v2 := []uint32{2147483647}
    var v3 int = 128
    v4 := make(map[string]int)
    var v5 string = (("hello" + (("hello" + "hello") + (("a" + "") + "hello"))) + ("a" + "foo bar"))
    var v6 uint = 1
    _ = len(v4)
    v3 = (-(v3))
    v5 = "foo bar"
    switch "\n\t" {
    case "hello":
        switch v3 {
        case -938:
            var v7 uint32 = 835
            _ = v7
        case -916:
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v8 uint16 = 1
                if (true || true) {
                    v9 := GenS1{F0: -1}
                    var v10 uint8 = 30
                    _ = v10
                    _ = v9
                } else {
                    if ((v6 % ((v6) | 1)) < 1) {
                        if false {
                            v2 = append(v2, 65535)
                        } else {
                            v11 := GenS1{F0: 1}
                            v12 := GenS1{F0: 0}
                            _ = v11
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 4; fi13++ {
                            var v14 int64 = 183
                            _ = fi13
                            _ = v14
                        }
                        fmt.Println(v8, v3, v3)
                    } else {
                        v15 := GenS2{F0: -1, F1: ((924 <= 2147483647) || (!(!true))), F2: 128}
                        fmt.Println(v3, v5, v6)
                        _ = v15
                    }
                    var v16 string = ""
                    _ = v16
                }
                _ = v8
            }
        case 0:
            fmt.Println(v6, v3)
        default:
            var v17 int32 = -2147483648
            _ = v17
        }
    case "a":
        var v18 int = -514
        _ = v18
    default:
        v19 := GenS3{F0: (^(v3)), F1: 32}
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 float64 = 0.0
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS1{F0: 64}
            v24 := GenS1{F0: -102}
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi20
        _ = v21
    }
    v5 = v5
    v4[(("hello" + "hello") + v5)] = 107
    for fi25 := 0; fi25 < 7; fi25++ {
        for fi26 := 0; fi26 < 15; fi26++ {
            v27 := GenS1{F0: -128}
            var v28 int64 = -586
            _ = fi26
            _ = v27
            _ = v28
        }
        switch "\n\t" {
        case "a":
            fmt.Println(v6, v5)
        }
        _ = fi25
    }
    var v29 uint = 1
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 uint8 = 10
        _ = fi30
        _ = v31
    }
    var v32 int = (v3 | v3)
    if len(v1) > 0 {
        v1[0] = 65535
    }
    if true {
        var v33 int16 = 1000
        if true {
            var v34 int64 = 9223372036854775807
            var v35 uint32 = 1
            _ = v34
            _ = v35
        }
        _ = v33
    }
    var v36 float32 = (0.1 + (float32(-2.25) - float32(float64(3.14159))))
    v37 := GenS3{F0: 369, F1: -31}
    if (0 < 1) {
        v32 += (v32 - v3)
    }
    v3 += (^(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
