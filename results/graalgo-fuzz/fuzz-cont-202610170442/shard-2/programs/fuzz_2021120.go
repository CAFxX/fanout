package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 bool; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -557
    var v3 int8 = 42
    v4 := make(map[string]int)
    var v5 bool = (1 != -565)
    var v6 float64 = (1.5 + float64((0.1 - 1.5)))
    v7 := GenS1{F0: 21, F1: v2, F2: v3}
    var v8 int16 = 123
    v9 := GenS1{F0: v3, F1: v2, F2: v3}
    switch -1 {
    case -9223372036854775808:
        fmt.Println(v8, v3, v2)
    default:
        v5 = v5
    }
    v2 -= (v2 & v2)
    var v10 uint16 = 2
    switch ("hello" + "") {
    case "\n\t":
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 float64 = v6
            _ = fi11
            _ = v12
        }
    case "hello":
        v2 += int64(v8)
    case "":
        var v13 float32 = 1e10
        _ = v13
    }
    if v5 {
        v6 = v6
    }
    v5 = ((1 > 729) || v5)
    fmt.Println(v8)
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(v10, v6)
        _ = fi14
    }
    if true {
        var v15 int16 = 0
        _ = v15
    } else {
        var v16 int8 = v3
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        switch ("foo bar" + "hello") {
        case "":
            v18 := []int{fi17, 491}
            _ = v18
        }
        _ = fi17
    }
    switch "foo bar" {
    case "foo bar":
        if (!v5) {
            v19 := GenS2{F0: 0}
            switch -106 {
            case -1:
                v8 += (v8 << ((65535) & 63))
            case 127:
                v20 := GenS1{F0: v3, F1: (v2 / ((v2) | 1)), F2: (v3 ^ v3)}
                _ = v20
            default:
                v21 := GenS3{F0: v8, F1: v5, F2: 179}
                _ = v21
            }
            _ = v19
        }
    default:
        v22 := GenS3{F0: 0, F1: (-32768 >= v8), F2: 2}
        _ = v22
    }
    var v23 float64 = 3.14159
    for fi24 := 0; fi24 < 3; fi24++ {
        v2 = (v2 + 331)
        fmt.Println(fi24)
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
