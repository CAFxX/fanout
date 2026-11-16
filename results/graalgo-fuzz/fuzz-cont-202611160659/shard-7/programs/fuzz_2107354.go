package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int64; F2 uint }
type GenS4 struct { F0 bool }
func genFn0() string {
    if true {
        switch 2147483647 {
        case -9223372036854775808:
            var v5 uint32 = 713
            _ = v5
        case 128:
            v6 := GenS1{F0: 500}
            _ = v6
        case -2147483648:
            var v7 int8 = 127
            _ = v7
        default:
            var v8 bool = false
            _ = v8
        }
    } else {
        switch "a" {
        case "hello":
            v9 := []int{36, 42}
            _ = v9
        case "\n\t":
            if false {
                var v10 string = ("hello" + (((("" + "hello") + "hello") + "\n\t") + "foo bar"))
                fmt.Println(v10)
                fmt.Println(v10)
                _ = v10
            } else {
                var v11 float32 = ((((3.14159 - (1e10 - (3.14159 / (((-2.25 + -2.25) - 1.5) / 1e10)))) * (0.1 / (-2.25 / 0.1))) - float32(((1e10 / -2.25) - 0.0))) + 1e10)
                _ = v11
            }
        }
    }
    v12 := GenS1{F0: 65535}
    v13 := GenS1{F0: 255}
    var v14 uint8 = 255
    _ = v12
    _ = v13
    _ = v14
    return ((("a" + "foo bar") + "\n\t") + ("a" + "hello"))
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 0
    var v3 int32 = 432
    v4 := make(map[uint32]string)
    var v15 string = "\n\t"
    if ((!true) || (-746 == (^(v3)))) {
        var v16 int8 = 100
        switch (v15 + (v15 + (("\n\t" + (v15 + (v15 + v15))) + v15))) {
        case "a":
            v17 := GenS2{F0: 9223372036854775807}
            _ = v17
        case "":
            v18 := GenS2{F0: -421}
            _ = v18
        case "foo bar":
            if false {
                v16 = (^(v16))
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 int8 = (-(v16))
                    _ = fi19
                    _ = v20
                }
            } else {
                v15 = "a"
                fmt.Println(v15, v2, v2)
            }
        }
        v21 := GenS1{F0: 32767}
        _ = v16
        _ = v21
    }
    v2 = (v2 * v2)
    var v22 uint = 4294967295
    v23 := GenS3{F0: v22, F1: 520, F2: (v22 / (((v22 | 0)) | 1))}
    v24 := GenS1{F0: 324}
    if false {
        var v25 float32 = (0.0 + float32(0.0))
        for fi26 := 0; fi26 < 6; fi26++ {
            v1[v2] = ""
            var v27 uint8 = 21
            _ = fi26
            _ = v27
        }
        v15 = v15
        _ = v25
    } else {
        if true {
            var v28 int32 = -2147483648
            for fi29 := 0; fi29 < 8; fi29++ {
                for fi30 := 0; fi30 < 4; fi30++ {
                    var v31 int16 = -200
                    switch 42 {
                    case 42:
                        v22 = (v22 ^ 4294967295)
                    case -1:
                        v2 = (^(v2))
                    }
                    _ = fi30
                    _ = v31
                }
                v1[-2147483648] = (v15 + v15)
                _ = fi29
            }
            _ = v28
        } else {
            fmt.Println(v2, v22, v22)
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 bool = ((v3 + v3) >= v3)
            _ = fi32
            _ = v33
        }
        v34 := GenS4{F0: true}
        _ = v34
    }
    var v35 int8 = 42
    v36 := []int{128, -709, (v2 + (v2 & (v2 + -2147483648)))}
    fmt.Println(v2, v15)
    v37 := GenS2{F0: int64(v2)}
    delete(v1, (v2 | v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
