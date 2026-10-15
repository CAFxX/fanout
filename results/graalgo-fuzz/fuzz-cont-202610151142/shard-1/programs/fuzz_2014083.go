package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 int = -2147483648
    v2 := []uint32{2147483647}
    v3 := make(map[uint32]bool)
    var v4 uint8 = 102
    var v5 uint32 = 65535
    fmt.Println(v5, v1)
    fmt.Println(v5, v1, v1)
    var v6 int64 = 0
    for fi7 := 0; fi7 < 20; fi7++ {
        switch ("a" + "hello") {
        case "foo bar":
            switch "a" {
            case "":
                v4 = 0
            default:
                var v8 int = 4294967295
                _ = v8
            }
        case "\n\t":
            var v9 uint16 = 1
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 uint8 = 42
            _ = len(v2)
            v12 := GenS1{F0: 629, F1: 32767}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    fmt.Println(v1)
    v4 = v4
    var v13 int8 = -1
    var v14 uint32 = 0
    var v15 int64 = -184
    var v16 int = 32767
    var v17 float64 = (1.5 - ((((((3.14159 * (-2.25 - float64((-2.25 - 0.1)))) / -2.25) + float64(3.14159)) + (1e10 - 0.0)) - -2.25) - 3.14159))
    var v18 int16 = 744
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 int32 = 10
        _ = fi19
        _ = v20
    }
    var v21 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
