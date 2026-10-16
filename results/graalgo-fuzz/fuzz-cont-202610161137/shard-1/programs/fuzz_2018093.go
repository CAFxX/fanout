package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 float64 = ((3.14159 * 3.14159) / ((((1.5 / 3.14159) * float64(0.1)) * 0.0) - 1.5))
    var v2 int8 = -1
    var v3 int32 = 2147483647
    v4 := []string{}
    v3 -= v3
    var v5 uint8 = 0
    var v6 uint16 = 616
    fmt.Println(v1, v6, v3)
    switch 9223372036854775807 {
    case -307:
        v6 += (v6 * v6)
    default:
        fmt.Println(v5, v2)
    }
    if true {
        fmt.Println(v6, v2, v1)
        var v7 string = "a"
        var v8 float32 = 0.1
        _ = v7
        _ = v8
    }
    var v9 int64 = 0
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(fi10)
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int = fi11
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    if ((0.0 * float32((v1 - 0.0))) >= 3.14159) {
        for fi13 := 0; fi13 < 8; fi13++ {
            switch "" {
            case "\n\t":
                v6 = 0
            case "a":
                var v14 int8 = 42
                _ = v14
            case "":
                v6 -= v6
            default:
                v3 -= (^(v3))
            }
            _ = fi13
        }
        _ = len(v4)
        v15 := GenS1{F0: (v9 & v9)}
        _ = v15
    }
    var v16 int = -9223372036854775808
    for fi17 := 0; fi17 < 20; fi17++ {
        switch "a" {
        case "\n\t":
            for fi18 := 0; fi18 < 4; fi18++ {
                v2 = (v2 >> ((2) & 63))
                switch ("hello" + "foo bar") {
                case "\n\t":
                    v4 = append(v4, "foo bar")
                case "foo bar":
                    v19 := GenS2{F0: v5}
                    _ = v19
                case "hello":
                    var v20 uint32 = 2147483647
                    _ = v20
                }
                var v21 string = "a"
                _ = fi18
                _ = v21
            }
        case "a":
            var v22 uint32 = 0
            _ = v22
        case "foo bar":
            var v23 uint32 = 918
            _ = v23
        }
        _ = fi17
    }
    var v24 uint32 = 2
    v9 -= (v9 ^ int64(v2))
    var v25 int64 = -363
    var v26 uint = 42
    var v27 int32 = v3
    fmt.Println(v27, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
