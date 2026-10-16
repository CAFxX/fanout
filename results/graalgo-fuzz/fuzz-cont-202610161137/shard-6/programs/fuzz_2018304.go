package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 int64) int {
    fmt.Println(p0)
    if false {
        switch 32767 {
        case -9223372036854775808:
            p0 -= (p0 & p0)
        default:
            v4 := []int{}
            _ = v4
        }
        p0 += p0
        v5 := []int{4294967295, 175, -1, 0}
        _ = v5
    } else {
        fmt.Println(p0)
    }
    _ = p0
    return 9223372036854775807
}
func genFn1() uint {
    var v6 uint64 = 42
    if (-1 == -765) {
        v6 -= (v6 & (v6 & 42))
        if false {
            v6 -= (^(v6))
            v7 := []int{0, 127, -2147483648, 128}
            v8 := GenS1{F0: 2, F1: 2, F2: 32767}
            _ = v7
            _ = v8
        } else {
            if false {
                fmt.Println(v6)
            }
            var v9 uint = 540
            _ = v9
        }
    } else {
        if (-9223372036854775808 >= 4294967295) {
            v10 := []int{4294967295, 127}
            var v11 uint64 = v6
            v12 := GenS1{F0: 65535, F1: 30, F2: -1}
            _ = v10
            _ = v11
            _ = v12
        } else {
            if ((true && (!((!(!(65535 != 32767))) && ((42 >= -1) || false)))) && false) {
                var v13 bool = (42 < 623)
                fmt.Println(v13, v6)
                _ = v13
            }
            fmt.Println(v6)
        }
    }
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 int8 = 1
    var v2 uint16 = 111
    var v3 uint16 = (v2 << ((18) & 63))
    switch "hello" {
    case "":
        switch 249 {
        case 467:
            fmt.Println(v1, v3, v1)
        case 42:
            v14 := []int{}
            _ = v14
        }
    case "foo bar":
        fmt.Println(v2, v3)
    case "a":
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS1{F0: 2, F1: 1, F2: 32767}
            _ = fi15
            _ = v16
        }
    }
    fmt.Println(v1, v3)
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v2)
        _ = fi17
    }
    var v18 uint = 255
    var v19 int64 = 288
    v20 := []int{128, -351, 9223372036854775807}
    v21 := GenS1{F0: v2, F1: (v2 >> ((42) & 63)), F2: 1}
    v22 := GenS2{F0: 392}
    if false {
        fmt.Println(v2)
        v23 := GenS3{F0: -1, F1: v1}
        _ = v23
    }
    v18 += v18
    v24 := GenS3{F0: -1, F1: (v1 + v1)}
    if len(v20) > 0 {
        v20[0] = 32767
    }
    v3 -= 32767
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
}
