package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint; F1 int; F2 int64 }
func genFn0(p0 uint16) uint {
    fmt.Println(p0)
    if true {
        var v6 bool = false
        _ = v6
    }
    if false {
        var v7 int32 = -2147483648
        _ = v7
    }
    v8 := []int{-9223372036854775808, 665, 369}
    var v9 uint = 18446744073709551615
    _ = p0
    _ = v8
    _ = v9
    return (v9 / ((18446744073709551615) | 1))
}

func main() {
    var v1 uint8 = 152
    var v2 int32 = 42
    var v3 float64 = float64(3.14159)
    var v4 uint16 = 422
    var v5 string = "a"
    var v10 float64 = v3
    var v11 uint16 = v4
    var v12 int64 = 1
    var v13 uint64 = 538
    v2 += v2
    if true {
        v1 = v1
        switch 42 {
        case 812:
            switch 84 {
            case -1:
                v14 := GenS1{F0: 32767, F1: (v5 + ("hello" + ((v5 + v5) + v5)))}
                _ = v14
            default:
                var v15 int16 = 1
                _ = v15
            }
        case 32768:
            v16 := []int{1}
            _ = v16
        case 127:
            var v17 bool = (!(242 > 255))
            _ = v17
        }
    }
    fmt.Println(v5)
    switch v5 {
    case "":
        if (2 == 1) {
            var v18 int8 = 90
            var v19 int8 = (v18 % (((v18 + v18)) | 1))
            _ = v18
            _ = v19
        }
    case "hello":
        switch 386 {
        case -404:
            switch -49 {
            case 9223372036854775807:
                v20 := []int{}
                _ = v20
            case 0:
                v4 = v11
            }
        case 374:
            var v21 uint = 295
            _ = v21
        }
    case "foo bar":
        v22 := []int{127, 679, 124}
        _ = v22
    }
    v23 := GenS2{F0: 259, F1: 128, F2: (-(v12))}
    if ((!false) || false) {
        var v24 int16 = -1
        _ = v24
    }
    var v25 int8 = 100
    v26 := []int{493, -9223372036854775808}
    var v27 uint32 = 42
    v3 += v3
    var v28 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
