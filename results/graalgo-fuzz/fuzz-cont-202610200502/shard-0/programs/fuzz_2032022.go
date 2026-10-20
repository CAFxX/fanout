package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 string }

func main() {
    var v1 uint8 = 0
    var v2 int16 = -353
    var v3 int16 = (v2 ^ -32768)
    var v4 int = 1
    var v5 uint16 = 255
    var v6 uint32 = 54
    v7 := []int{}
    switch (v4 ^ (^(v4))) {
    case -9223372036854775808:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case -301:
        fmt.Println(v5, v2)
    }
    switch v4 {
    case -1:
        var v8 bool = false
        _ = v8
    case 9223372036854775807:
        v5 = v5
    case 32768:
        if true {
            v9 := GenS1{F0: (v5 % ((v5) | 1)), F1: 9223372036854775807, F2: "hello"}
            _ = v9
        }
    default:
        if ((!true) && (v1 == (v1 << ((0) & 63)))) {
            fmt.Println(v2)
            v7 = append(v7, (-(v4)))
        }
    }
    v3 = v3
    v7 = append(v7, (^(v4)))
    var v10 float32 = 0.1
    switch int(v6) {
    case -607:
        if len(v7) > 0 {
            v7[0] = (v4 ^ v4)
        }
    case 2147483647:
        switch "foo bar" {
        case "a":
            v11 := GenS1{F0: 32767, F1: 863, F2: ("" + (("a" + ("\n\t" + "foo bar")) + (("a" + ((("\n\t" + "hello") + "a") + "a")) + "foo bar")))}
            _ = v11
        case "foo bar":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case "\n\t":
            v12 := GenS1{F0: 1, F1: -9223372036854775808, F2: ("hello" + "a")}
            _ = v12
        }
    }
    v4 -= 9223372036854775807
    if (127 >= (^(v1))) {
        var v13 int32 = -1
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v4 += (v4 - 32768)
        _ = fi14
    }
    _ = len(v7)
    fmt.Println(v4, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
