package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint32 = 171
    var v2 float64 = (0.0 * -2.25)
    v3 := []int{222, 32768}
    v4 := make(map[int]int)
    var v5 uint8 = 225
    var v6 int = 9223372036854775807
    v6 -= v6
    switch "" {
    case "a":
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint = 4294967295
            fmt.Println(v8)
            v2 = ((float64(-2.25) + v2) / 1.0)
            _ = fi7
            _ = v8
        }
    case "foo bar":
        if true {
            if (false || false) {
                v5 = v5
            }
            v9 := GenS1{F0: 1000, F1: 1}
            var v10 uint16 = 2
            _ = v10
            _ = v9
        }
    case "\n\t":
        v2 += v2
    }
    if (!true) {
        v1 += (v1 << ((65535) & 63))
        var v11 uint64 = 0
        _ = v11
    } else {
        switch -342 {
        case 4294967295:
            if len(v3) > 0 {
                v3[0] = (-(v6))
            }
        default:
            var v12 int = (v6 % ((int(v1)) | 1))
            _ = v12
        }
        v13 := v4[v6]
        _ = v13
    }
    var v14 int16 = -1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = ((v2 * (v2 + v2)) / 1.0)
    var v15 float64 = v2
    var v16 uint64 = 1
    var v17 int32 = -2147483648
    fmt.Println(v1)
    for fi18 := 0; fi18 < 9; fi18++ {
        if false {
            var v19 uint16 = 65535
            var v20 bool = (("a" + "\n\t") <= (("a" + "foo bar") + "a"))
            delete(v4, fi18)
            _ = v19
            _ = v20
        } else {
            v15 += (v15 * 3.14159)
            v6 += (^(v6))
        }
        fmt.Println(v2)
        var v21 uint = 0
        _ = fi18
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
