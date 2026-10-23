package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 42
    var v3 int16 = 32767
    v4 := make(map[int]int)
    var v5 string = ("" + "foo bar")
    v6 := []int{127, 2147483647}
    var v7 uint64 = 0
    if true {
        fmt.Println(v1, v3, v5)
        v3 = -56
    } else {
        fmt.Println(v2, v7, v2)
        if false {
            if (2 > (^(v7))) {
                v8 := GenS1{F0: (false && true), F1: 18446744073709551615}
                _ = v8
            }
            var v9 uint16 = 65535
            v10 := GenS2{F0: (v5 + "foo bar")}
            _ = v10
            _ = v9
        }
    }
    if true {
        v11 := GenS3{F0: (true && false)}
        var v12 uint8 = v2
        v7 -= 1
        _ = v11
        _ = v12
    } else {
        if false {
            var v13 int32 = v1
            for fi14 := 0; fi14 < 11; fi14++ {
                v5 = v5
                _ = fi14
            }
            var v15 int64 = 42
            _ = v13
            _ = v15
        }
        v3 += -1
        fmt.Println(v2, v7)
    }
    if false {
        if true {
            switch -9223372036854775808 {
            case -2147483648:
                v5 = v5
            default:
                v1 = v1
            }
            delete(v4, -512)
        } else {
            var v16 int8 = 79
            _ = len(v4)
            v1 -= (v1 / ((v1) | 1))
            _ = v16
        }
        v2 = 127
    } else {
        v17 := GenS3{F0: true}
        v3 = 0
        _ = v17
    }
    v3 += (v3 + v3)
    if (false && false) {
        fmt.Println(v1)
        var v18 int8 = -65
        v19 := GenS2{F0: "foo bar"}
        _ = v18
        _ = v19
    }
    v5 = v5
    v3 += v3
    v1 -= (v1 + -787)
    var v20 uint16 = 937
    fmt.Println(v3)
    var v21 uint8 = v2
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
