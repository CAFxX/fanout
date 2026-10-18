package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }

func main() {
    var v1 string = ""
    var v2 uint8 = 127
    var v3 uint16 = 255
    v4 := make(map[string]uint32)
    v2 += v2
    var v5 uint = 1
    fmt.Println(v3, v2)
    var v6 uint16 = 2
    if (1 != -2147483648) {
        _ = len(v4)
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 int32 = 2147483647
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        fmt.Println(v5, v2, v2)
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        fmt.Println(v6)
        _ = fi10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        switch -1 {
        case 1:
            if true {
                fmt.Println(v5, v5, v6)
                v5 -= (v5 >> (((v5 - uint(v2))) & 63))
            }
        default:
            v12 := GenS1{F0: (v1 + v1), F1: (true && (281 > -1))}
            _ = v12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        fmt.Println(fi13)
        _ = fi13
    }
    switch ((v1 + (v1 + v1)) + v1) {
    case "foo bar":
        fmt.Println(v1, v5, v1)
    case "\n\t":
        v14 := v4[(v1 + (v1 + (v1 + v1)))]
        _ = v14
    default:
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 5; fi16++ {
                fmt.Println(v1, v1, v5)
                v17 := []int{fi16, (fi16 % ((fi16) | 1))}
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    }
    switch "\n\t" {
    case "":
        var v18 int64 = 1
        _ = v18
    case "a":
        var v19 int8 = 127
        _ = v19
    case "\n\t":
        v1 = ("foo bar" + ((v1 + v1) + "hello"))
    default:
        if (false || true) {
            fmt.Println(v6)
            v5 += v5
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 bool = (!true)
                var v22 int8 = 127
                _ = fi20
                _ = v21
                _ = v22
            }
        }
    }
    v4[v1] = 2
    var v23 uint16 = v3
    for fi24 := 0; fi24 < 3; fi24++ {
        fmt.Println(v1, v1, v2)
        var v25 uint16 = v6
        v25 += v6
        _ = fi24
        _ = v25
    }
    var v26 int32 = 1
    _ = len(v4)
    v23 += v23
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
