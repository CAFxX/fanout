package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 float32 = 1.5
    var v2 uint32 = 525
    var v3 int16 = 0
    var v4 int64 = -1
    var v5 int8 = 42
    v6 := GenS1{F0: 246, F1: (-(v3))}
    var v7 uint32 = 953
    fmt.Println(v1, v7)
    if true {
        if (65535 <= 78) {
            var v8 float64 = (-2.25 + 3.14159)
            fmt.Println(v3, v8, v2)
            _ = v8
        } else {
            v9 := GenS1{F0: 70, F1: -32768}
            v1 += 0.0
            _ = v9
        }
        v10 := GenS1{F0: 179, F1: -77}
        _ = v10
    }
    var v11 uint16 = 32767
    v4 = v4
    switch ("" + ("hello" + (("\n\t" + "") + ("foo bar" + (("foo bar" + ("hello" + ("hello" + "hello"))) + "hello"))))) {
    case "\n\t":
        v12 := GenS2{F0: ("" + ("hello" + "\n\t")), F1: 742}
        _ = v12
    case "hello":
        var v13 bool = true
        _ = v13
    default:
        var v14 float32 = (-2.25 - v1)
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        switch fi15 {
        case -9223372036854775808:
            if false {
                var v16 float64 = ((1e10 * (1.5 - 0.0)) - -2.25)
                _ = v16
            } else {
                var v17 string = ("a" + "hello")
                _ = v17
            }
        case 4294967295:
            var v18 float64 = 0.1
            _ = v18
        }
        fmt.Println(v3, v3)
        switch (("\n\t" + "a") + "\n\t") {
        case "a":
            var v19 int64 = 9223372036854775807
            _ = v19
        }
        _ = fi15
    }
    var v20 int64 = v4
    fmt.Println(v20)
    var v21 float32 = v1
    if true {
        v22 := GenS1{F0: 2, F1: (v3 % ((int16(v5)) | 1))}
        _ = v22
    }
    if true {
        if false {
            v20 += v4
            var v23 int16 = (v3 * (v3 - v3))
            v24 := []int{0}
            _ = v23
            _ = v24
        }
        v3 = (v3 % ((v3) | 1))
    }
    var v25 int32 = 2147483647
    var v26 uint32 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
