package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 float32 = ((1.5 - 0.1) / ((3.14159 + (1e10 + (-2.25 + 1.5))) * ((3.14159 - (1e10 + (0.0 / (-2.25 / float32(float64(1.5)))))) / -2.25)))
    v2 := []uint32{1, 1, 2, 2147483647}
    var v3 int32 = 0
    var v4 int = 684
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi5 := 0; fi5 < 9; fi5++ {
        v1 = v1
        switch "a" {
        case "foo bar":
            _ = len(v2)
        default:
            fmt.Println(v3)
        }
        _ = fi5
    }
    var v6 uint16 = 724
    v7 := GenS1{F0: 9223372036854775807}
    switch 496 {
    case -850:
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint8 = 1
            switch ("\n\t" + (("hello" + ("foo bar" + "foo bar")) + "hello")) {
            case "a":
                var v10 int = v4
                _ = v10
            }
            _ = fi8
            _ = v9
        }
    case 248:
        var v11 uint64 = 418
        _ = v11
    default:
        v12 := GenS1{F0: -9223372036854775808}
        _ = v12
    }
    var v13 string = ""
    fmt.Println(v3)
    if ((!false) && (!(v4 == (v4 & 127)))) {
        if false {
            var v14 int = v4
            _ = v14
        }
    } else {
        v4 += 697
        v3 -= v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
