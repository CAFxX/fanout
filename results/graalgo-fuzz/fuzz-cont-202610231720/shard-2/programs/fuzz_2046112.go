package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }

func main() {
    var v1 float64 = (float64((1.5 - (0.0 * float32(3.14159)))) * (float64(((3.14159 / 0.1) / 1.0)) + 1.5))
    var v2 uint8 = 44
    var v3 int32 = 0
    var v4 bool = true
    if v4 {
        var v5 int8 = -14
        _ = v5
    }
    v6 := GenS1{F0: 2, F1: 0, F2: 100}
    v7 := GenS1{F0: 176, F1: 32767, F2: 0}
    var v8 uint = 803
    switch "hello" {
    case "":
        switch "foo bar" {
        case "":
            switch -9223372036854775808 {
            case 281:
                if v4 {
                    switch (((("foo bar" + "foo bar") + "foo bar") + "\n\t") + "") {
                    case "foo bar":
                        var v9 int16 = 42
                        _ = v9
                    case "a":
                        v1 = v1
                    case "\n\t":
                        fmt.Println(v2, v8, v1)
                    default:
                        fmt.Println(v1, v3, v1)
                    }
                    v10 := []int{-9223372036854775808}
                    var v11 uint16 = 32767
                    _ = v10
                    _ = v11
                }
            }
        case "hello":
            if (!(1 <= 456)) {
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 float64 = (((v1 + v1) / (3.14159 - v1)) / 1e10)
                    var v14 string = ((("a" + ("a" + ("foo bar" + "foo bar"))) + "a") + "foo bar")
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            } else {
                var v15 int32 = -1
                v3 = (v15 + (-(v3)))
                _ = v15
            }
        }
    case "a":
        if true {
            fmt.Println(v4, v2)
            v8 += 18446744073709551615
        }
    case "foo bar":
        v2 = (^(v2))
    }
    var v16 string = ("foo bar" + "")
    var v17 int16 = 0
    v16 = "foo bar"
    v18 := GenS1{F0: 724, F1: int(v2), F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
