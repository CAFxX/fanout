package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int32; F1 int16; F2 uint8 }
func genFn0(p0 string, p1 int, p2 uint) int32 {
    switch p0 {
    case "\n\t":
        if (42 <= (p2 + p2)) {
            p0 = ((p0 + ("foo bar" + p0)) + "hello")
        }
    case "a":
        v6 := GenS1{F0: 2}
        _ = v6
    case "":
        var v7 int8 = 100
        _ = v7
    default:
        v8 := GenS1{F0: 255}
        _ = v8
    }
    if (!((1 <= 1) && true)) {
        if (!true) {
            p0 = (p0 + (("" + "hello") + (p0 + p0)))
        }
        var v9 float32 = 1.5
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 822
}

func main() {
    var v1 uint = 972
    var v2 string = "foo bar"
    var v3 string = ""
    var v4 uint64 = 757
    var v5 bool = (true && false)
    var v10 int64 = -1
    v5 = (v5 || ((true && v5) && v5))
    v11 := GenS1{F0: 739}
    var v12 uint = (^(v1))
    var v13 int = -951
    v14 := GenS2{F0: 0, F1: -2147483648}
    v2 = v2
    var v15 uint = (v1 & (v12 ^ (v1 ^ 13)))
    v4 = 2
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            v18 := GenS2{F0: 311, F1: 1}
            var v19 float64 = float64(1e10)
            v5 = (!(!(((v5 || v5) || (!(42 >= 0))) && false)))
            _ = fi17
            _ = v18
            _ = v19
        }
        if (0 <= 2) {
            v20 := []int{636, fi16, fi16, v13}
            _ = v20
        } else {
            if (!v5) {
                v13 -= (fi16 - v13)
            } else {
                fmt.Println(v5, v10)
            }
        }
        switch -1 {
        case 42:
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 int64 = v10
                var v23 uint16 = 2
                var v24 int64 = v10
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
        case 127:
            v4 = (v4 + (^(v4)))
        }
        _ = fi16
    }
    if ((v5 && v5) && v5) {
        if (((v4 & v4) < v4) && (9223372036854775807 < v10)) {
            var v25 int64 = v10
            _ = v25
        }
        v26 := GenS3{F0: 0, F1: 42, F2: 243}
        _ = v26
    } else {
        v27 := []int{}
        fmt.Println(v13, v12, v10)
        var v28 uint16 = 135
        _ = v27
        _ = v28
    }
    v4 = (v4 ^ 148)
    v10 -= v10
    v29 := GenS3{F0: 2147483647, F1: int16(v10), F2: 14}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
