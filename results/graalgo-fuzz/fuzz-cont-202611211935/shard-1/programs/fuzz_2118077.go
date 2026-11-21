package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint16 }

func main() {
    var v1 bool = true
    var v2 uint64 = 2
    var v3 int64 = -1
    var v4 uint32 = 244
    v2 = 2
    switch 1 {
    case 828:
        v5 := GenS1{F0: 1, F1: "foo bar", F2: 699}
        _ = v5
    }
    if v1 {
        v2 = v2
    }
    v1 = (65535 < 1)
    v4 = (^(v4))
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint8 = 2
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        switch (("hello" + "\n\t") + "hello") {
        case "a":
            v3 += -517
        case "":
            var v9 float32 = 0.0
            _ = v9
        case "foo bar":
            var v10 int = 42
            _ = v10
        }
        _ = fi8
    }
    var v11 int64 = 9223372036854775807
    fmt.Println(v2)
    v1 = (false || (255 < 79))
    v2 += 18446744073709551615
    if v1 {
        var v12 string = "hello"
        v11 = (v3 / ((1) | 1))
        _ = v12
    } else {
        var v13 bool = ((v11 & v11) > v3)
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := []int{32767, (-(fi14)), fi14, -1}
        _ = fi14
        _ = v15
    }
    v1 = v1
    v2 = (v2 + 2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
