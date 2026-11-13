package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 42
    var v3 int32 = 1
    var v4 int32 = 92
    for fi5 := 0; fi5 < 2; fi5++ {
        v2 += (v2 + (^(v2)))
        _ = fi5
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint = 669
        var v8 uint8 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -953 {
    case 128:
        switch "foo bar" {
        case "\n\t":
            var v9 uint32 = 2147483647
            _ = v9
        case "":
            var v10 uint8 = 0
            _ = v10
        case "foo bar":
            v2 = 65535
        }
    case -675:
        var v11 bool = true
        _ = v11
    }
    v12 := []int{-1, 2147483647, -140, -353}
    if (18446744073709551615 != 4294967295) {
        fmt.Println(v4, v3)
        var v13 int32 = (^(v3))
        _ = v13
    } else {
        fmt.Println(v2, v4, v2)
        v4 -= -2147483648
    }
    var v14 uint16 = 65535
    var v15 float64 = -2.25
    v4 -= 0
    switch "foo bar" {
    case "a":
        fmt.Println(v4, v2)
    case "foo bar":
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 bool = (-128 > -1)
            v3 -= (v3 | v4)
            _ = fi16
            _ = v17
        }
    case "":
        fmt.Println(v3, v2)
    default:
        v2 -= 32767
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
