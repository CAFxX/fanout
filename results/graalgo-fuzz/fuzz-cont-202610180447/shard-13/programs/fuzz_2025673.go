package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 bool; F2 bool }
func genFn0(p0 uint32, p1 uint8, p2 uint32) int64 {
    v7 := GenS1{F0: 0, F1: 127}
    var v8 uint16 = 2
    if false {
        v9 := GenS1{F0: 32767, F1: 127}
        var v10 int8 = -50
        _ = v10
        _ = v9
    } else {
        var v11 int32 = 2147483647
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 uint = 604
            var v14 uint16 = 42
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int16 = 998
        _ = v11
        _ = v15
    }
    var v16 int = 9223372036854775807
    v16 += 127
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 float32 = ((0.0 / 3.14159) - ((-2.25 - float32(0.1)) * ((-2.25 - float32(-2.25)) - 0.1)))
    var v2 uint = 65535
    var v3 string = "foo bar"
    var v4 string = v3
    var v5 int64 = -1
    var v6 int32 = 42
    if false {
        v2 -= (v2 / ((v2) | 1))
        var v17 int32 = v6
        _ = v17
    } else {
        var v18 int8 = 108
        switch ("\n\t" + (v3 + v4)) {
        case "":
            v19 := GenS2{F0: "a"}
            _ = v19
        case "a":
            if false {
                var v20 uint16 = 1
                v21 := []int{}
                var v22 float64 = 1.5
                _ = v20
                _ = v21
                _ = v22
            }
        }
        _ = v18
    }
    v5 = v5
    switch "foo bar" {
    case "a":
        switch 0 {
        case 128:
            v23 := GenS3{F0: -1}
            _ = v23
        case 673:
            for fi24 := 0; fi24 < 18; fi24++ {
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 uint32 = 1
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
        case 127:
            v1 = v1
        default:
            v27 := []int{623, 127}
            _ = v27
        }
    case "":
        v4 = ("" + "\n\t")
    case "\n\t":
        var v28 uint8 = 42
        _ = v28
    default:
        fmt.Println(v6, v3, v5)
    }
    v29 := []int{128, -503, -554, 32767}
    switch "" {
    case "a":
        var v30 uint32 = 238
        _ = v30
    case "\n\t":
        fmt.Println(v6, v2, v6)
    case "hello":
        if ((false && (2147483648 < 1)) || false) {
            v29 = append(v29, -657)
            var v31 int16 = 1000
            for fi32 := 0; fi32 < 9; fi32++ {
                v29 = append(v29, fi32)
                _ = fi32
            }
            _ = v31
        }
    }
    fmt.Println(v2, v2, v4)
    var v33 uint = v2
    v2 -= 2
    for fi34 := 0; fi34 < 14; fi34++ {
        if false {
            v2 = v33
        }
        if true {
            _ = len(v29)
            var v35 int8 = -107
            v36 := GenS4{F0: (true || (v35 < 7))}
            _ = v35
            _ = v36
        }
        _ = fi34
    }
    v37 := GenS5{F0: 128, F1: false, F2: (2 != 0)}
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    switch -9223372036854775808 {
    case 32768:
        _ = len(v29)
    default:
        v1 = v1
    }
    v1 -= (v1 / (((float32((1.5 / -2.25)) * v1) / v1) - 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
