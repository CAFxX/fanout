package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }

func main() {
    var v1 float64 = ((float64(3.14159) - float64(float32(0.1))) + float64(1.5))
    var v2 uint32 = 4294967295
    v3 := make(map[int]uint32)
    fmt.Println(v1, v1)
    v1 += v1
    if (float64((-2.25 / (0.0 + 1.5))) > 1.5) {
        var v4 string = ""
        var v5 float64 = (v1 - 0.0)
        _ = v4
        _ = v5
    }
    switch 32768 {
    case 2147483647:
        if true {
            var v6 int64 = 42
            v7 := v3[-152]
            delete(v3, 32768)
            _ = v6
            _ = v7
        }
    default:
        var v8 bool = (!false)
        _ = v8
    }
    var v9 string = (("\n\t" + "\n\t") + ((("hello" + "a") + ("a" + "a")) + (("foo bar" + "hello") + "\n\t")))
    switch v9 {
    case "":
        var v10 string = (v9 + v9)
        _ = v10
    case "hello":
        v2 = (^(v2))
    default:
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 float64 = v1
            _ = fi11
            _ = v12
        }
    }
    v2 -= (v2 & v2)
    v9 = "hello"
    var v13 uint32 = v2
    if (42 >= 18446744073709551615) {
        var v14 bool = (-675 != -814)
        _ = v14
    }
    var v15 int16 = 1
    fmt.Println(v9, v2)
    var v16 uint8 = 128
    var v17 uint8 = uint8(v13)
    v18 := GenS1{F0: 234, F1: 100}
    switch v9 {
    case "a":
        switch v9 {
        case "foo bar":
            v19 := v3[0]
            _ = v19
        case "hello":
            if false {
                delete(v3, 9223372036854775807)
            } else {
                v15 -= v15
            }
        case "\n\t":
            var v20 bool = true
            _ = v20
        }
    default:
        var v21 float64 = (1.5 + (3.14159 + v1))
        _ = v21
    }
    v16 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
