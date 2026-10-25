package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }

func main() {
    var v1 float32 = -2.25
    var v2 int32 = -1
    var v3 uint8 = 128
    var v4 string = ((("a" + ("a" + ("" + "foo bar"))) + ("foo bar" + "\n\t")) + "a")
    var v5 uint = 1
    if true {
        for fi6 := 0; fi6 < 7; fi6++ {
            switch ("" + "a") {
            case "\n\t":
                v7 := GenS1{F0: 2, F1: fi6}
                _ = v7
            case "":
                fmt.Println(v2)
            default:
                var v8 uint8 = v3
                _ = v8
            }
            v9 := []int{fi6, -288}
            v4 = v4
            _ = fi6
            _ = v9
        }
    }
    var v10 uint = (v5 + 42)
    v2 -= -683
    var v11 float32 = (((v1 + v1) + v1) - 3.14159)
    switch ((("hello" + v4) + (v4 + ((v4 + v4) + v4))) + "foo bar") {
    case "foo bar":
        v11 -= v1
    case "hello":
        if true {
            v5 = (v10 / ((v5) | 1))
        }
    case "a":
        v4 = v4
    default:
        var v12 int64 = 0
        _ = v12
    }
    var v13 int64 = 1
    v13 -= (v13 * -715)
    var v14 float32 = float32(float64((1.5 + (v1 + v1))))
    v3 -= (v3 & 1)
    var v15 int16 = 0
    if true {
        v16 := GenS1{F0: uint64(v5), F1: 0}
        v17 := GenS1{F0: 42, F1: 32768}
        _ = v16
        _ = v17
    }
    v1 = float32((3.14159 + (float64(1e10) - (1e10 - (3.14159 / (1e10 - 0.1))))))
    if (767 == 806) {
        var v18 uint16 = 65535
        v19 := GenS1{F0: 42, F1: int(v18)}
        var v20 bool = (!false)
        _ = v18
        _ = v19
        _ = v20
    } else {
        v21 := []int{}
        v14 -= (v1 / (v1 + v11))
        var v22 bool = false
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
