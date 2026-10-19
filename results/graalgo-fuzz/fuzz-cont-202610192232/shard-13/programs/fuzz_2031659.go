package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 uint16, p1 uint32) bool {
    switch "" {
    case "hello":
        v5 := []int{}
        _ = v5
    case "":
        p1 = p1
    case "foo bar":
        v6 := GenS1{F0: 42}
        _ = v6
    }
    var v7 float64 = (0.1 / 1e10)
    v8 := GenS1{F0: -341}
    for fi9 := 0; fi9 < 5; fi9++ {
        p1 -= (p1 / ((p1) | 1))
        var v10 float64 = (((v7 / v7) + v7) + (v7 * (0.0 + v7)))
        _ = fi9
        _ = v10
    }
    if (!true) {
        v7 = 3.14159
        v11 := GenS2{F0: 2}
        var v12 bool = (!true)
        _ = v11
        _ = v12
    } else {
        if true {
            p1 = p1
        } else {
            fmt.Println(p0)
        }
        fmt.Println(v7, p0)
        v13 := GenS1{F0: -309}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p1 == 1)
}

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]string)
    var v3 int32 = -1
    var v4 uint32 = 2147483647
    var v14 uint = 4294967295
    v14 += v14
    var v15 float32 = (3.14159 * -2.25)
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := []int{(fi16 << ((168) & 63)), -994, (fi16 + 1), -9223372036854775808}
        switch 1 {
        case 127:
            v18 := GenS3{F0: true, F1: 2147483647}
            _ = v18
        case -9223372036854775808:
            v19 := v2[1]
            _ = v19
        case 2147483647:
            if false {
                switch (fi16 * (^(fi16))) {
                case -1:
                    v3 += v3
                default:
                    var v20 float32 = v15
                    _ = v20
                }
                v3 -= v3
                if false {
                    var v21 uint = v14
                    v22 := GenS1{F0: -865}
                    _ = v21
                    _ = v22
                }
            } else {
                switch 0 {
                case 128:
                    var v23 float32 = float32(((3.14159 - 0.0) / 0.1))
                    _ = v23
                }
                _ = len(v2)
            }
        default:
            switch 42 {
            case 4294967295:
                var v24 uint32 = (v4 >> ((v14) & 63))
                _ = v24
            case 128:
                v15 += float32(3.14159)
            default:
                v2[2147483648] = ("a" + "hello")
            }
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v4)
    var v25 int16 = 42
    var v26 uint32 = v4
    for fi27 := 0; fi27 < 5; fi27++ {
        if true {
            v14 = (v14 ^ (v14 / ((v14) | 1)))
            v25 = 42
        } else {
            v14 = 65535
        }
        var v28 uint32 = (v4 % ((2147483648) | 1))
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
