package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 uint32 = 589
    var v4 uint64 = 4294967295
    v5 := make(map[int]uint32)
    var v6 bool = (65535 == 2)
    if v6 {
        if (1 != 127) {
            v7 := []int{397, -725}
            if v6 {
                fmt.Println(v3, v4)
                var v8 uint8 = 128
                v4 = (v4 << ((952) & 63))
                _ = v8
            } else {
                v9 := GenS1{F0: (^(v4))}
                var v10 string = "foo bar"
                _ = v10
                _ = v9
            }
            v4 = v4
            _ = v7
        }
    } else {
        _ = len(v5)
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        switch ("" + ("hello" + "\n\t")) {
        case "\n\t":
            v5[(fi11 << ((456) & 63))] = v3
        case "foo bar":
            var v12 int64 = int64(v4)
            _ = v12
        }
        v2[fi11] = (^(v3))
        _ = fi11
    }
    v6 = ((1e10 <= 3.14159) || true)
    v4 = 2
    v13 := GenS1{F0: v4}
    for fi14 := 0; fi14 < 11; fi14++ {
        v6 = (!true)
        fmt.Println(v3, v6, v4)
        _ = fi14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int = 0
        _ = fi15
        _ = v16
    }
    v17 := []int{}
    var v18 uint16 = 723
    if ("foo bar" != ("" + "")) {
        var v19 uint = 220
        _ = v19
    } else {
        for fi20 := 0; fi20 < 4; fi20++ {
            fmt.Println(v4, v4, v3)
            var v21 int16 = -32768
            _ = fi20
            _ = v21
        }
        v3 = (v3 << ((65535) & 63))
        switch -2147483648 {
        case 158:
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        default:
            _ = len(v1)
        }
    }
    v3 = (v3 * 438)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
