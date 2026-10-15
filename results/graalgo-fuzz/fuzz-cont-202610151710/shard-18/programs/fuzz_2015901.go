package main

import (
    "fmt"
)

func genFn0(p0 uint32, p1 int8) int16 {
    v4 := []int{-1, 0, -2147483648}
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v5 uint32 = 2147483647
        switch -801 {
        case 42:
            var v6 int = -471
            _ = v6
        case -154:
            fmt.Println(v5)
        default:
            var v7 uint32 = p0
            _ = v7
        }
        _ = v5
    }
    switch "hello" {
    case "foo bar":
        var v8 int32 = 1
        _ = v8
    case "\n\t":
        var v9 uint = 65535
        _ = v9
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(fi10, p1)
            _ = fi10
        }
    }
    v4 = append(v4, 0)
    _ = p0
    _ = p1
    _ = v4
    return -599
}

func main() {
    var v1 uint32 = 661
    var v2 uint32 = (v1 ^ v1)
    v3 := make(map[string]uint32)
    if true {
        v1 = v1
        v2 += (v1 | 517)
    } else {
        fmt.Println(v2, v2)
    }
    v1 -= 2
    v1 -= 2147483647
    var v11 int16 = -1
    var v12 uint8 = 42
    var v13 int16 = (v11 / ((v11) | 1))
    v2 = (v1 * v1)
    fmt.Println(v11, v1, v2)
    var v14 uint16 = 65535
    switch 0 {
    case 9223372036854775807:
        var v15 uint16 = v14
        _ = v15
    default:
        switch ("\n\t" + "hello") {
        case "\n\t":
            v16 := []int{32768, -363, 9223372036854775807}
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 int16 = (v13 ^ genFn0((v1 >> ((783) & 63)), 0))
            v20 := []int{-876}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
