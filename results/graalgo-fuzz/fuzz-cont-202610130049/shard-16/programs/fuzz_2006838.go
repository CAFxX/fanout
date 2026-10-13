package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -1
    var v3 uint32 = 2147483647
    var v4 uint16 = 2
    v5 := make(map[int]int)
    v4 = (v4 / ((65535) | 1))
    var v6 uint32 = (v3 / (((^(v3))) | 1))
    fmt.Println(v4, v3)
    for fi7 := 0; fi7 < 10; fi7++ {
        v3 += 1
        if ((!false) && (1 > 42)) {
            var v8 uint = 251
            v3 += 42
            _ = v8
        }
        _ = fi7
    }
    if true {
        var v9 int8 = 42
        if false {
            fmt.Println(v3)
            v10 := GenS1{F0: v6}
            v11 := GenS2{F0: 9223372036854775807, F1: 1}
            _ = v10
            _ = v11
        } else {
            v12 := GenS2{F0: -1, F1: -9223372036854775808}
            _ = v12
        }
        fmt.Println(v9, v3)
        _ = v9
    }
    fmt.Println(v3, v4)
    v13 := []int{623, 127, -2147483648}
    var v14 float32 = -2.25
    if true {
        var v15 uint = 657
        v6 = v3
        fmt.Println(v3)
        _ = v15
    } else {
        switch "" {
        case "foo bar":
            v3 += (v3 ^ v3)
        case "a":
            fmt.Println(v14, v2)
        case "\n\t":
            switch 4294967295 {
            case 294:
                v2 = 142
            case 0:
                fmt.Println(v14)
            case -917:
                var v16 uint64 = 4294967295
                _ = v16
            }
        default:
            switch "" {
            case "hello":
                switch -9223372036854775808 {
                case -1:
                    var v17 uint64 = 959
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 5; fi18++ {
                        switch (-(fi18)) {
                        case 9223372036854775807:
                            v19 := v1[(fi18 % ((fi18) | 1))]
                            _ = v19
                        case 921:
                            v13 = append(v13, fi18)
                        }
                        v14 -= (v14 / v14)
                        fmt.Println(v3)
                        _ = fi18
                    }
                }
            case "foo bar":
                v3 = (v6 & 472)
            }
        }
    }
    var v20 uint8 = 208
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
