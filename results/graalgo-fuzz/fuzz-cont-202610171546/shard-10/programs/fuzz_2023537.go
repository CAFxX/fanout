package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint; F2 uint8 }

func main() {
    var v1 int64 = 245
    var v2 float32 = 0.1
    v3 := []uint32{337, 0}
    var v4 int16 = 1
    var v5 int32 = 2147483647
    v6 := make(map[string]uint32)
    v2 += (((v2 * (v2 * 0.0)) / (1.5 + v2)) / v2)
    v7 := GenS1{F0: v4}
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: (^(v4))}
        switch 127 {
        case 178:
            switch ("hello" + "") {
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            }
        case -787:
            v6[""] = 2147483647
        case 120:
            var v10 int = (fi8 - (fi8 & fi8))
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    if true {
        v11 := GenS2{F0: "hello", F1: 18446744073709551615, F2: 127}
        v12 := GenS2{F0: "hello", F1: 419, F2: 55}
        _ = v11
        _ = v12
    }
    var v13 string = ""
    var v14 int64 = v1
    var v15 float32 = 3.14159
    delete(v6, v13)
    fmt.Println(v1, v5)
    _ = len(v3)
    var v16 string = "foo bar"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v17 int16 = -32768
    switch 9223372036854775807 {
    case 4294967295:
        v15 = v15
    case 9223372036854775807:
        var v18 uint32 = 0
        _ = v18
    default:
        var v19 int8 = 0
        _ = v19
    }
    var v20 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
