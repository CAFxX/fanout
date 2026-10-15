package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
type GenS2 struct { F0 int; F1 int; F2 uint16 }
func genFn0(p0 int64, p1 uint8) int32 {
    for fi5 := 0; fi5 < 2; fi5++ {
        p1 = 127
        fmt.Println(p0, p1, p0)
        _ = fi5
    }
    var v6 int16 = 42
    v7 := []int{}
    switch "\n\t" {
    case "":
        if (-117 >= -1) {
            fmt.Println(p0, p1, p1)
            _ = len(v7)
            if (2 > uint(p1)) {
                p1 -= (^(p1))
                p0 += 9223372036854775807
                v6 = (v6 << ((258) & 63))
            } else {
                var v8 uint = 255
                fmt.Println(v6)
                _ = v8
            }
        } else {
            v6 -= 732
            v6 -= (v6 * -1)
            if false {
                v9 := GenS1{F0: 65535, F1: 484, F2: -128}
                v10 := GenS1{F0: 2, F1: -1, F2: 100}
                _ = v10
                _ = v9
            }
        }
    }
    if (1 <= p1) {
        p1 = p1
        _ = len(v7)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    var v3 string = "hello"
    var v4 string = v3
    var v11 uint = 8
    if true {
        if false {
            var v12 uint = v11
            var v13 int32 = 42
            _ = v12
            _ = v13
        }
        var v14 bool = false
        v11 = 1
        _ = v14
    } else {
        var v15 string = "hello"
        _ = v15
    }
    v1 = 0
    v11 = v11
    v16 := GenS2{F0: 1, F1: 128, F2: 65535}
    if false {
        fmt.Println(v4, v3, v11)
        fmt.Println(v11, v3)
    } else {
        for fi17 := 0; fi17 < 18; fi17++ {
            v4 = (v4 + (v4 + v3))
            v4 = "foo bar"
            _ = fi17
        }
    }
    v11 -= v11
    v4 = v3
    switch -597 {
    case 9223372036854775807:
        switch (v4 + v3) {
        case "\n\t":
            if (false || (genFn0(9223372036854775807, (^(v1))) <= int32(v11))) {
                var v18 uint16 = 32767
                _ = v18
            } else {
                var v19 int32 = 1
                var v20 int = -792
                _ = v19
                _ = v20
            }
        default:
            var v21 int32 = 0
            _ = v21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
