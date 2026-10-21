package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
func genFn0(p0 uint16, p1 int32, p2 string) int8 {
    switch 4294967295 {
    case 403:
        var v5 int16 = 0
        _ = v5
    }
    if ((0.1 * -2.25) >= 1.5) {
        v6 := GenS1{F0: 65535, F1: 2, F2: true}
        _ = v6
    } else {
        v7 := GenS2{F0: -204}
        _ = v7
    }
    fmt.Println(p2)
    p0 += p0
    v8 := []int{-2147483648, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -128
}
func genFn1(p0 bool, p1 uint8, p2 string) int64 {
    switch 128 {
    case 464:
        v9 := []int{-1}
        _ = v9
    default:
        v10 := GenS3{F0: p1}
        _ = v10
    }
    v11 := []int{42, 9223372036854775807, 32768, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}

func main() {
    var v1 string = (("" + "hello") + ("" + ("hello" + "")))
    var v2 uint8 = 127
    var v3 uint = 18446744073709551615
    var v4 int16 = -351
    var v12 uint16 = 32767
    if true {
        var v13 int8 = 29
        v14 := GenS3{F0: (v2 & v2)}
        v4 -= -80
        _ = v13
        _ = v14
    }
    v15 := []int{32767, 128, 127, -531}
    v4 += (v4 / (((v4 / (((v4 << ((42) & 63))) | 1))) | 1))
    var v16 uint16 = v12
    if ((false && false) && (!true)) {
        _ = len(v15)
        switch 0 {
        case -2147483648:
            v3 += v3
        default:
            v16 -= (v12 + (v16 + (v12 / ((0) | 1))))
        }
        var v17 uint64 = 853
        _ = v17
    } else {
        var v18 uint64 = 2
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 2; fi19++ {
            _ = len(v15)
            _ = fi19
        }
        v20 := GenS4{F0: 128}
        _ = v20
    }
    _ = len(v15)
    v21 := GenS5{F0: (1e10 == float32(float64((((0.0 - 1.5) * 0.1) - float32(0.0)))))}
    switch v1 {
    case "foo bar":
        if (false || (!(!(140 == (v2 % ((0) | 1)))))) {
            v4 -= (v4 ^ (v4 - v4))
            v12 += (v16 << ((v3) & 63))
            var v22 float32 = (0.0 - (1.5 + (-2.25 + (0.1 / 1e10))))
            _ = v22
        } else {
            if len(v15) > 0 {
                v15[0] = 0
            }
            v3 -= (v3 - v3)
        }
    case "a":
        for fi23 := 0; fi23 < 10; fi23++ {
            fmt.Println(v12, v4)
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            if (42 == 42) {
                var v24 uint16 = (^(v12))
                _ = v24
            }
            _ = fi23
        }
    case "\n\t":
        v25 := GenS1{F0: 0, F1: 4294967295, F2: (!(true || (!((v1 + v1) > ((v1 + ((v1 + ("foo bar" + (v1 + (v1 + ("\n\t" + v1))))) + (v1 + "a"))) + ("" + v1))))))}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
