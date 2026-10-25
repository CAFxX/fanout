package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int = 32767
    var v2 int16 = 1000
    var v3 uint32 = 2
    var v4 int = 0
    if false {
        if (0 > 42) {
            v5 := []int{9223372036854775807, v4, v4, (v1 << ((724) & 63))}
            v6 := GenS1{F0: v2}
            switch 127 {
            case 32768:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case -874:
                var v7 int16 = (v2 + v2)
                _ = v7
            case 2147483647:
                var v8 uint8 = 0
                _ = v8
            }
            _ = v5
            _ = v6
        }
        if true {
            var v9 uint16 = 42
            v9 += v9
            _ = v9
        } else {
            var v10 int32 = 108
            _ = v10
        }
    } else {
        v11 := []int{42, 42}
        v3 = (v3 % ((v3) | 1))
        _ = v11
    }
    fmt.Println(v3, v4)
    var v12 int16 = (v2 % ((v2) | 1))
    fmt.Println(v3)
    var v13 int8 = 0
    v14 := GenS1{F0: (v12 - v2)}
    var v15 bool = (1 < -9223372036854775808)
    switch v4 {
    case -2147483648:
        var v16 uint64 = 2
        _ = v16
    default:
        if (!true) {
            v17 := GenS1{F0: (v2 * (v2 | (v12 / (((v12 | 42)) | 1))))}
            v18 := GenS1{F0: v2}
            var v19 bool = v15
            _ = v17
            _ = v18
            _ = v19
        }
    }
    var v20 uint32 = v3
    v1 -= v4
    fmt.Println(v1)
    v13 -= v13
    v12 = v2
    switch "" {
    case "foo bar":
        var v21 uint32 = v3
        _ = v21
    case "":
        fmt.Println(v4, v12)
    case "hello":
        v22 := GenS1{F0: int16(v20)}
        _ = v22
    }
    var v23 uint = 744
    if v15 {
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 int16 = v2
            _ = fi24
            _ = v25
        }
    } else {
        var v26 uint16 = 547
        _ = v26
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
}
