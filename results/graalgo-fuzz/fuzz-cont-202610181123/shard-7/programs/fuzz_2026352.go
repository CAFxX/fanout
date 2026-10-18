package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 160
    var v3 int64 = v2
    var v4 uint8 = 2
    v5 := make(map[uint32]string)
    var v6 uint32 = 203
    fmt.Println(v4, v3, v4)
    var v7 float64 = float64(0.1)
    v8 := GenS1{F0: (v3 - (v2 % (((v2 ^ (v2 & (^(v3))))) | 1)))}
    switch -137 {
    case -348:
        v3 = (v2 * (v2 + 42))
    case 42:
        v9 := GenS1{F0: (v3 - (v2 % (((v2 / (((v2 * (v2 * (v3 - 713)))) | 1))) | 1)))}
        _ = v9
    case 128:
        v10 := []int{-954}
        _ = v10
    }
    v11 := GenS2{F0: (^(v4)), F1: (v2 / (((^(v2))) | 1))}
    switch -606 {
    case 42:
        v3 += v2
    case -146:
        var v12 string = "foo bar"
        _ = v12
    case -2147483648:
        var v13 bool = true
        _ = v13
    default:
        switch ("a" + ("a" + "a")) {
        case "\n\t":
            v1["foo bar"] = 65535
        default:
            v14 := GenS2{F0: v4, F1: v3}
            _ = v14
        }
    }
    v4 -= 128
    var v15 string = "a"
    fmt.Println(v2, v7)
    fmt.Println(v2, v15)
    var v16 uint = 934
    if false {
        var v17 uint64 = 2
        v17 = (v17 * (v17 * (v17 << ((uint(v17)) & 63))))
        var v18 float64 = v7
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 5; fi19++ {
            switch 42 {
            case 128:
                v16 -= (v16 % ((v16) | 1))
            case -9223372036854775808:
                switch (v15 + (v15 + (((("foo bar" + (v15 + v15)) + "hello") + v15) + v15))) {
                case "foo bar":
                    var v20 bool = (!(!(v15 >= v15)))
                    _ = v20
                }
            }
            _ = fi19
        }
    }
    var v21 uint8 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
