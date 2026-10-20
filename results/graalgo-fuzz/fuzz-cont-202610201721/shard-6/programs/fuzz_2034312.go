package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int64; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint32 }

func main() {
    var v1 int16 = 1
    v2 := make(map[int]string)
    v3 := []string{"", ("a" + "a"), "\n\t", ("hello" + "foo bar")}
    var v4 float64 = (3.14159 * 3.14159)
    var v5 uint16 = 349
    v6 := []string{}
    var v7 uint16 = v5
    for fi8 := 0; fi8 < 19; fi8++ {
        delete(v2, fi8)
        _ = fi8
    }
    v9 := GenS1{F0: 0}
    if ((!((3.14159 + v4) != v4)) || (!(!false))) {
        v10 := GenS1{F0: 1}
        _ = v10
    } else {
        if true {
            v4 -= v4
            var v11 float32 = 0.1
            if false {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v7 -= v7
                    v13 := GenS2{F0: 0, F1: fi12}
                    _ = fi12
                    _ = v13
                }
                if ("foo bar" > (("a" + (("\n\t" + "foo bar") + "hello")) + "\n\t")) {
                    if false {
                        fmt.Println(v4, v11, v11)
                    } else {
                        var v14 string = ("" + "hello")
                        switch -526 {
                        case 989:
                            fmt.Println(v1)
                        case 9223372036854775807:
                            v2[2147483647] = v14
                        }
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 8; fi15++ {
                        var v16 int64 = 42
                        _ = fi15
                        _ = v16
                    }
                } else {
                    if true {
                        for fi17 := 0; fi17 < 13; fi17++ {
                            var v18 int64 = -9
                            _ = fi17
                            _ = v18
                        }
                        v19 := GenS2{F0: 65535, F1: 127}
                        _ = v19
                    }
                    var v20 string = ((("\n\t" + "") + "hello") + ("" + "a"))
                    switch 1 {
                    case -986:
                        v21 := GenS3{F0: 0, F1: (v1 - (v1 + (v1 * v1))), F2: 2147483647}
                        _ = v21
                    case -193:
                        v22 := GenS1{F0: 255}
                        _ = v22
                    case 0:
                        v11 = 1.5
                    }
                    _ = v20
                }
                v11 = -2.25
            }
            _ = v11
        } else {
            v4 = ((((3.14159 + (v4 * 0.0)) * 0.0) / v4) - (float64(1.5) + v4))
            v6 = append(v6, "\n\t")
            var v23 int8 = 127
            _ = v23
        }
        _ = len(v6)
    }
    v2[325] = "\n\t"
    v5 += 65535
    var v24 float32 = ((0.1 - -2.25) / float32(((((3.14159 / v4) - v4) / 1.0) / float64(3.14159))))
    _ = len(v2)
    fmt.Println(v4, v24)
    switch 685 {
    case 42:
        v1 = (-(v1))
    default:
        fmt.Println(v5)
    }
    if (false || ("a" != "hello")) {
        v25 := GenS4{F0: 843, F1: true}
        var v26 uint8 = 42
        v27 := GenS5{F0: 2147483648, F1: v5, F2: 207}
        _ = v25
        _ = v26
        _ = v27
    } else {
        var v28 string = "hello"
        fmt.Println(v4)
        var v29 bool = (v5 > (v7 + 848))
        _ = v28
        _ = v29
    }
    var v30 uint16 = v5
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 uint = 2
        _ = fi31
        _ = v32
    }
    v7 = (v7 / ((v5) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
