package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int64 }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    var v1 int16 = 42
    var v2 float32 = (1e10 + (0.0 - 0.1))
    var v3 string = "a"
    v4 := []int{2147483647, 32768, -332}
    var v5 uint32 = 4294967295
    if len(v4) > 0 {
        v4[0] = 41
    }
    var v6 uint = 0
    switch 42 {
    case -9223372036854775808:
        if (1e10 == (0.1 - (1.5 + 1.5))) {
            fmt.Println(v5)
            v2 = -2.25
            v4 = append(v4, 128)
        } else {
            v7 := GenS1{F0: 0, F1: 669, F2: -752}
            _ = v7
        }
    }
    switch 128 {
    case -2147483648:
        _ = len(v4)
    case -182:
        fmt.Println(v5, v2)
    case 127:
        v4 = append(v4, 2147483647)
    }
    var v8 float64 = 0.0
    v9 := GenS1{F0: 42, F1: 65535, F2: 9223372036854775807}
    if true {
        v5 -= v5
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v6 = uint(v5)
    }
    switch "foo bar" {
    case "\n\t":
        var v10 uint16 = 435
        _ = v10
    case "foo bar":
        var v11 uint = 0
        _ = v11
    case "a":
        if true {
            var v12 uint = (v6 + v6)
            v2 = v2
            v13 := GenS1{F0: 265, F1: 0, F2: -475}
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 int32 = 953
                fmt.Println(v2, v5, v1)
                _ = fi14
                _ = v15
            }
        }
    default:
        v1 += (^(v1))
    }
    v3 = v3
    v6 = 4294967295
    if (32767 < 255) {
        _ = len(v4)
        if false {
            if len(v4) > 0 {
                v4[0] = 1
            }
            var v16 int64 = 9223372036854775807
            _ = v16
        }
    }
    v17 := GenS2{F0: 127, F1: -1}
    fmt.Println(v6, v3, v3)
    for fi18 := 0; fi18 < 7; fi18++ {
        v5 -= v5
        var v19 float32 = (0.0 / v2)
        var v20 int16 = v1
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int16 = v1
    switch -2147483648 {
    case -2147483648:
        var v22 uint16 = 65535
        _ = v22
    case -650:
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 string = (v3 + v3)
            _ = fi23
            _ = v24
        }
    case 2147483647:
        var v25 uint64 = 2
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
