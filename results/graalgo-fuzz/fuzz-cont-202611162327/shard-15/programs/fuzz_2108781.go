package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int = 9223372036854775807
    v3 := make(map[uint32]string)
    var v4 float32 = (-2.25 - (-2.25 * -2.25))
    var v5 int = v2
    var v6 int32 = 0
    var v7 float64 = (float64(float32(1.5)) - 0.0)
    v2 += (v5 >> ((198) & 63))
    var v8 uint = 644
    switch "foo bar" {
    case "a":
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(v8)
            v7 -= v7
            var v10 bool = (v8 != (v8 + 42))
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS1{F0: ("a" + "hello")}
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(v2, v8)
        _ = fi12
    }
    var v13 int = (v5 << ((803) & 63))
    v14 := GenS1{F0: "foo bar"}
    v6 = v6
    var v15 uint8 = 127
    v16 := GenS2{F0: 1000}
    _ = len(v3)
    v17 := GenS1{F0: (("foo bar" + "a") + "\n\t")}
    switch int(v8) {
    case 93:
        var v18 uint = v8
        _ = v18
    default:
        for fi19 := 0; fi19 < 7; fi19++ {
            fmt.Println(v6)
            _ = fi19
        }
    }
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
