package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int64; F1 uint16; F2 int8 }

func main() {
    v1 := []int{127, 32767}
    var v2 uint64 = 1
    v3 := []uint32{0}
    if (-699 > 623) {
        var v4 float64 = -2.25
        v4 = v4
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        _ = v4
    } else {
        if (!(((("" + ("" + "hello")) + "\n\t") + "a") < ("" + (((("" + "hello") + ("foo bar" + "foo bar")) + (((("\n\t" + "hello") + "a") + "foo bar") + "foo bar")) + (("\n\t" + "foo bar") + ("hello" + "a")))))) {
            var v5 bool = true
            v2 += (v2 & 985)
            fmt.Println(v5, v5)
            _ = v5
        } else {
            v6 := GenS1{F0: 2, F1: 703, F2: 322}
            _ = v6
        }
    }
    var v7 int8 = -1
    var v8 float32 = (1.5 / 0.1)
    v2 = (v2 >> ((2) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 := GenS2{F0: 561, F1: 4294967295}
    if false {
        switch 32768 {
        case 9223372036854775807:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 500:
            v10 := GenS3{F0: 1, F1: 255, F2: v7}
            _ = v10
        case -823:
            var v11 uint = 18446744073709551615
            _ = v11
        }
        var v12 uint32 = 2147483648
        _ = v12
    }
    var v13 uint8 = 255
    if (!(v2 >= v2)) {
        v14 := GenS2{F0: 2, F1: 2}
        _ = v14
    } else {
        fmt.Println(v8, v13)
        var v15 int64 = -9223372036854775808
        if len(v3) > 0 {
            v3[0] = 65535
        }
        _ = v15
    }
    if len(v1) > 0 {
        v1[0] = 32767
    }
    var v16 uint8 = v13
    fmt.Println(v8)
    v1 = append(v1, 4294967295)
    var v17 float32 = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
