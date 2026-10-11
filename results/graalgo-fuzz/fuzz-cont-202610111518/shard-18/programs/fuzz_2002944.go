package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int; F1 uint32; F2 uint }
type GenS4 struct { F0 uint32; F1 uint32; F2 int32 }
func genFn0(p0 int) int8 {
    var v6 int64 = -9223372036854775808
    v7 := GenS1{F0: 424}
    if ("foo bar" > "hello") {
        var v8 uint64 = 0
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int8 = 100
    var v2 int32 = -1
    v3 := []uint32{}
    var v4 int16 = 1
    var v5 uint8 = 42
    v4 = (v4 >> ((2) & 63))
    var v9 uint16 = 255
    if (v4 == (v4 >> ((18446744073709551615) & 63))) {
        var v10 uint32 = 42
        if len(v3) > 0 {
            v3[0] = (v10 + (^(v10)))
        }
        _ = v10
    } else {
        var v11 int = 2147483647
        for fi12 := 0; fi12 < 8; fi12++ {
            v3 = append(v3, 2147483648)
            v13 := GenS2{F0: 18446744073709551615, F1: (!(false || true))}
            v14 := GenS3{F0: (fi12 + 4294967295), F1: 42, F2: 658}
            _ = fi12
            _ = v13
            _ = v14
        }
        v11 = v11
        _ = v11
    }
    v3 = append(v3, 2147483648)
    v5 = v5
    var v15 int32 = v2
    v16 := GenS4{F0: 4294967295, F1: 4294967295, F2: v2}
    var v17 int8 = (v1 | -128)
    if false {
        switch 4294967295 {
        case -721:
            v18 := GenS2{F0: 18446744073709551615, F1: true}
            _ = v18
        case -493:
            var v19 int32 = v15
            _ = v19
        case 32768:
            v17 += (v17 ^ v17)
        }
        v5 = v5
        if (true && (!(1 < 940))) {
            if false {
                var v20 uint32 = 145
                _ = v20
            }
            v3 = append(v3, 2147483648)
        } else {
            if false {
                v5 += (v5 / ((208) | 1))
                v5 -= (v5 - v5)
            } else {
                v1 -= genFn0(869)
                switch "hello" {
                case "foo bar":
                    v17 += genFn0(2147483647)
                case "a":
                    v17 -= int8(v4)
                }
            }
            var v21 int8 = v1
            var v22 bool = (false || false)
            _ = v21
            _ = v22
        }
    } else {
        fmt.Println(v17, v1, v1)
        var v23 bool = true
        _ = len(v3)
        _ = v23
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = v15
    var v24 int32 = (^(v2))
    v24 = v24
    if true {
        switch -66 {
        case 254:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            v25 := GenS4{F0: 589, F1: 4294967295, F2: v24}
            _ = v25
        }
    }
    var v26 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
