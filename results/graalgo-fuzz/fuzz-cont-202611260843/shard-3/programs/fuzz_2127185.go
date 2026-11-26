package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }

func main() {
    var v1 bool = (!(0 <= 65535))
    var v2 uint64 = 0
    var v3 uint32 = 2147483648
    var v4 uint8 = 184
    var v5 int16 = 32767
    v6 := []int{-9223372036854775808}
    v7 := GenS1{F0: 2147483647}
    v5 += (v5 / ((42) | 1))
    for fi8 := 0; fi8 < 16; fi8++ {
        if v1 {
            fmt.Println(v2, v5)
        }
        _ = fi8
    }
    if (!(-64 == 100)) {
        var v9 uint32 = (^(v3))
        var v10 string = "hello"
        v11 := GenS2{F0: (42 > int(v3))}
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 string = "\n\t"
        fmt.Println(v3, v1)
        v12 = ("a" + v12)
        _ = v12
    }
    v1 = (v5 == (v5 >> ((42) & 63)))
    fmt.Println(v4)
    fmt.Println(v2)
    if true {
        var v13 float32 = 1e10
        _ = v13
    }
    if (v1 && true) {
        fmt.Println(v4)
    } else {
        v4 += v4
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        v5 += (-(v5))
        for fi15 := 0; fi15 < 5; fi15++ {
            if (v5 == v5) {
                v16 := GenS2{F0: v1}
                _ = v16
            } else {
                if len(v6) > 0 {
                    v6[0] = fi14
                }
            }
            var v17 uint8 = v4
            v2 = (v2 * v2)
            _ = fi15
            _ = v17
        }
        switch fi14 {
        case 128:
            v3 += (v3 << ((280) & 63))
        case 32768:
            fmt.Println(v1, v2)
        case -735:
            var v18 uint16 = 396
            _ = v18
        }
        _ = fi14
    }
    fmt.Println(v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
