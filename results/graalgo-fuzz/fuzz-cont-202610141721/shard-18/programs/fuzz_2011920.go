package main

import (
    "fmt"
)


func main() {
    var v1 uint = 0
    var v2 string = "hello"
    v3 := []string{v2}
    var v4 int8 = 42
    var v5 uint = 255
    _ = len(v3)
    fmt.Println(v5)
    v4 += (v4 ^ v4)
    v2 = v2
    var v6 uint16 = 42
    switch v2 {
    case "":
        if len(v3) > 0 {
            v3[0] = v2
        }
    }
    v4 += v4
    if len(v3) > 0 {
        v3[0] = v2
    }
    v2 = ""
    if false {
        for fi7 := 0; fi7 < 9; fi7++ {
            v4 = (v4 >> (((v5 - 191)) & 63))
            _ = fi7
        }
    } else {
        v4 += -108
        fmt.Println(v4)
        fmt.Println(v5, v5)
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v6 = 65535
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v6 = (v6 >> ((v1) & 63))
        _ = fi8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v2 = "hello"
        _ = fi9
    }
    fmt.Println(v1, v1, v4)
    var v10 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
