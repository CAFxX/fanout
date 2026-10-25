package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 bool; F1 int16 }

func main() {
    var v1 int16 = 1
    var v2 uint8 = 1
    v3 := make(map[int]string)
    if true {
        var v4 float64 = 1.5
        v4 -= v4
        switch ("\n\t" + ((("hello" + ("hello" + ("\n\t" + ""))) + ((("hello" + ("a" + "hello")) + "\n\t") + "foo bar")) + "hello")) {
        case "a":
            fmt.Println(v4, v4)
        case "hello":
            v2 = v2
        }
        _ = v4
    }
    v1 += v1
    v1 = v1
    if ((((3.14159 / (((3.14159 - 0.1) * -2.25) - (float32(1e10) - (0.0 / 1e10)))) / 1e10) == float32(1.5)) || (false && true)) {
        switch ((("hello" + "a") + "hello") + ("hello" + "a")) {
        case "a":
            switch 42 {
            case 4294967295:
                var v5 uint32 = 65535
                _ = v5
            default:
                fmt.Println(v2)
            }
        }
        var v6 int16 = v1
        v7 := GenS1{F0: false, F1: 42}
        _ = v6
        _ = v7
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if true {
            v1 = (v1 / ((v1) | 1))
        } else {
            for fi8 := 0; fi8 < 15; fi8++ {
                v1 = (^(v1))
                switch (fi8 << ((4294967295) & 63)) {
                case 0:
                    v9 := v3[fi8]
                    _ = v9
                default:
                    var v10 bool = (true || false)
                    _ = v10
                }
                v1 += (v1 ^ (v1 ^ v1))
                _ = fi8
            }
            var v11 int64 = -924
            _ = v11
        }
    case "hello":
        var v12 int = -9223372036854775808
        _ = v12
    default:
        fmt.Println(v1, v2)
    }
    v13 := GenS1{F0: (!(true && (201 >= 0))), F1: 2147483647}
    v14 := GenS2{F0: 127, F1: "\n\t"}
    v15 := []int{633, 9223372036854775807, -9223372036854775808, -688}
    if false {
        fmt.Println(v1, v1)
    }
    v16 := GenS3{F0: false, F1: v1}
    var v17 uint8 = 48
    for fi18 := 0; fi18 < 5; fi18++ {
        fmt.Println(v17, fi18, v17)
        fmt.Println(v17, v2, v2)
        _ = fi18
    }
    fmt.Println(v17, v2, v2)
    var v19 float32 = float32(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
}
