package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{65535, 2147483648}
    var v2 int8 = 73
    var v3 float32 = (1e10 - 1.5)
    switch ("foo bar" + ("a" + "foo bar")) {
    case "foo bar":
        v2 -= (v2 - 127)
    default:
        fmt.Println(v2)
    }
    if (381 <= 18446744073709551615) {
        v2 = v2
        switch 128 {
        case 9223372036854775807:
            if (745 == 4294967295) {
                switch "hello" {
                case "hello":
                    v3 = float32((1e10 * -2.25))
                }
            } else {
                for fi4 := 0; fi4 < 4; fi4++ {
                    var v5 int64 = -613
                    var v6 int16 = 0
                    var v7 uint32 = 4294967295
                    _ = fi4
                    _ = v5
                    _ = v6
                    _ = v7
                }
            }
        case -401:
            var v8 int16 = 1000
            _ = v8
        }
        switch "\n\t" {
        case "foo bar":
            var v9 string = "hello"
            _ = v9
        default:
            fmt.Println(v2, v3)
        }
    }
    if true {
        _ = len(v1)
        switch 2147483647 {
        case 0:
            var v10 float64 = 0.1
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int16 = 1
        _ = fi11
        _ = v12
    }
    var v13 uint8 = 127
    v2 = (v2 / ((v2) | 1))
    v3 += (v3 * v3)
    v1 = append(v1, 0)
    v13 += v13
    v2 -= 124
    var v14 int = 32767
    var v15 int8 = -14
    v3 -= ((v3 + float32((3.14159 / 0.1))) * v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
