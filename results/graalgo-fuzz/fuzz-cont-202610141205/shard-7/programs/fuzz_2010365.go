package main

import (
    "fmt"
)


func main() {
    var v1 float32 = 0.0
    var v2 float32 = 0.1
    v3 := []string{("" + ("a" + "hello"))}
    var v4 float64 = 3.14159
    var v5 uint8 = 2
    v5 -= (v5 % (((^(v5))) | 1))
    v1 -= (float32(float64((3.14159 * 1e10))) - ((0.1 - (v1 + 1e10)) - v2))
    var v6 int64 = -1
    v2 = v1
    var v7 uint = 65535
    switch "a" {
    case "a":
        v5 = 42
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        if ((!true) || (true && false)) {
            v2 = v1
            for fi9 := 0; fi9 < 15; fi9++ {
                _ = len(v3)
                var v10 int16 = 0
                _ = fi9
                _ = v10
            }
        } else {
            var v11 uint = v7
            var v12 float64 = v4
            fmt.Println(v7)
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    var v13 uint = (v7 >> (((v7 + 65535)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
