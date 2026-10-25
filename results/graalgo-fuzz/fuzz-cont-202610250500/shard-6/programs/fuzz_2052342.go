package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 bool }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := []int{-9223372036854775808}
    var v2 uint64 = 4294967295
    var v3 string = "\n\t"
    var v4 float32 = -2.25
    var v5 uint8 = 42
    var v6 uint16 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 := GenS1{F0: 127, F1: v3, F2: false}
    if false {
        v8 := GenS2{F0: -9223372036854775808, F1: 42}
        switch (((("hello" + ("hello" + ((("" + v3) + "a") + v3))) + v3) + v3) + v3) {
        case "hello":
            v3 = v3
        default:
            v5 = (v5 << ((0) & 63))
        }
        _ = v8
    }
    if ((!false) || (false && (1000 <= -32768))) {
        v3 = "hello"
    }
    var v9 int64 = 537
    v1 = append(v1, -1)
    var v10 int64 = v9
    if true {
        if len(v1) > 0 {
            v1[0] = 127
        }
        if false {
            v11 := GenS2{F0: -563, F1: 839}
            switch "hello" {
            case "foo bar":
                if true {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        var v13 float32 = v4
                        _ = fi12
                        _ = v13
                    }
                }
            case "\n\t":
                v10 += v10
            default:
                if len(v1) > 0 {
                    v1[0] = -905
                }
            }
            _ = v11
        }
        fmt.Println(v5, v10)
    } else {
        if false {
            fmt.Println(v6, v4, v4)
        }
        v3 = "a"
    }
    v2 = (v2 ^ (v2 << ((435) & 63)))
    var v14 float64 = 0.1
    switch -773 {
    case 9223372036854775807:
        v3 = v3
    case 0:
        v15 := GenS3{F0: v2}
        _ = v15
    default:
        switch int(v5) {
        case -2147483648:
            var v16 uint64 = (v2 % ((2) | 1))
            _ = v16
        default:
            switch "a" {
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    }
    var v17 int64 = v10
    if (false && (65535 > 42)) {
        for fi18 := 0; fi18 < 3; fi18++ {
            for fi19 := 0; fi19 < 9; fi19++ {
                fmt.Println(v3)
                var v20 int32 = -1
                _ = len(v1)
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v1 = append(v1, 9223372036854775807)
    if true {
        if true {
            fmt.Println(v5, v9)
            var v21 float32 = v4
            _ = v21
        }
        var v22 uint8 = (v5 * v5)
        _ = v22
    }
    v9 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
