package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint8; F2 bool }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 float64 = (3.14159 - 1.5)
    var v4 uint32 = 1
    if (!true) {
        v5 := []int{-709, 42}
        if true {
            for fi6 := 0; fi6 < 14; fi6++ {
                var v7 int8 = 55
                for fi8 := 0; fi8 < 8; fi8++ {
                    v9 := GenS1{F0: 0, F1: -402, F2: 42}
                    var v10 int16 = -32768
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        } else {
            var v11 uint32 = 2147483648
            for fi12 := 0; fi12 < 13; fi12++ {
                if len(v5) > 0 {
                    v5[0] = fi12
                }
                _ = fi12
            }
            _ = v11
        }
        _ = v5
    }
    fmt.Println(v3)
    v3 = (((3.14159 / v3) - (float64(0.1) - 1e10)) - (float64(3.14159) / ((1.5 + float64(((float32(v3) / (3.14159 + 0.0)) / -2.25))) * (v3 / v3))))
    var v13 uint = 18446744073709551615
    fmt.Println(v4, v13, v13)
    fmt.Println(v13)
    v14 := GenS1{F0: 32768, F1: 42, F2: 1000}
    var v15 string = "a"
    var v16 string = "foo bar"
    v16 = "hello"
    if (0 <= -396) {
        switch (((("hello" + v15) + v16) + v15) + v16) {
        case "":
            v16 = "\n\t"
        case "a":
            v17 := v2[32767]
            _ = v17
        case "\n\t":
            v18 := []int{9223372036854775807, 127}
            _ = v18
        default:
            for fi19 := 0; fi19 < 6; fi19++ {
                v20 := GenS2{F0: ("" + (v15 + (("hello" + "hello") + (("foo bar" + v16) + v15))))}
                _ = fi19
                _ = v20
            }
        }
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        v4 = v4
        _ = fi21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v13 = v13
        var v23 string = (((v15 + "foo bar") + v15) + "foo bar")
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        switch v16 {
        case "hello":
            if (1e10 >= (float32((-2.25 * v3)) / 1e10)) {
                var v25 string = (("a" + v15) + "foo bar")
                _ = v25
            } else {
                for fi26 := 0; fi26 < 4; fi26++ {
                    delete(v2, (fi26 << ((384) & 63)))
                    _ = fi26
                }
                var v27 string = "hello"
                for fi28 := 0; fi28 < 8; fi28++ {
                    fmt.Println(v13)
                    _ = fi28
                }
                _ = v27
            }
        case "foo bar":
            v13 -= (v13 / (((v13 + v13)) | 1))
        default:
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 int8 = 1
                switch v16 {
                case "\n\t":
                    v31 := []int{fi29, (fi29 - (fi29 - fi24))}
                    _ = v31
                case "a":
                    if (0 <= 2) {
                        var v32 bool = true
                        _ = len(v1)
                        v33 := GenS3{F0: 127, F1: 42, F2: v32}
                        _ = v32
                        _ = v33
                    } else {
                        fmt.Println(v4)
                        fmt.Println(v16, v15, fi24)
                        var v34 uint32 = 52
                        _ = v34
                    }
                default:
                    fmt.Println(fi24, v16)
                }
                _ = fi29
                _ = v30
            }
        }
        _ = fi24
    }
    v35 := v1["a"]
    if true {
        _ = len(v2)
        var v36 int16 = -624
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
