package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint64 = 21
    var v2 bool = (!false)
    var v3 int32 = 42
    v4 := []string{}
    var v5 uint16 = 611
    v3 += v3
    var v6 int = 128
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    switch "" {
    case "a":
        var v7 uint8 = 128
        _ = v7
    }
    fmt.Println(v3)
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v1, v3)
        _ = fi8
    }
    var v9 uint64 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 -= (^(v5))
    if (4294967295 == 56) {
        var v10 int32 = v3
        _ = v10
    } else {
        for fi11 := 0; fi11 < 2; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                if len(v4) > 0 {
                    v4[0] = ""
                }
                var v13 int8 = 8
                var v14 uint64 = (v9 >> ((1) & 63))
                _ = fi12
                _ = v13
                _ = v14
            }
            fmt.Println(v3)
            _ = fi11
        }
        switch v6 {
        case 2147483647:
            fmt.Println(v9)
        case -716:
            for fi15 := 0; fi15 < 7; fi15++ {
                v6 = v6
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 string = "foo bar"
                    var v18 bool = (true || (-1 < -261))
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi15
            }
        case -1:
            _ = len(v4)
        }
        var v19 uint64 = v9
        _ = v19
    }
    v20 := GenS1{F0: 0}
    if v2 {
        if ((v5 ^ (v5 % (((v5 ^ (v5 + 1))) | 1))) != 930) {
            v21 := GenS1{F0: 127}
            var v22 int32 = (v3 * 1)
            if v2 {
                v2 = v2
                v6 += -55
            }
            _ = v21
            _ = v22
        }
        fmt.Println(v9, v3)
        fmt.Println(v1, v9, v2)
    } else {
        if len(v4) > 0 {
            v4[0] = "a"
        }
        var v23 int32 = (v3 * v3)
        _ = len(v4)
        _ = v23
    }
    switch 128 {
    case -384:
        v6 = (-(v6))
    default:
        v6 = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
