package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int; F2 int }
type GenS3 struct { F0 int64 }

func main() {
    var v1 float32 = (1e10 / (-2.25 * 0.1))
    var v2 int = 32768
    var v3 int32 = -2147483648
    var v4 int = v2
    switch "foo bar" {
    case "\n\t":
        if true {
            if true {
                var v5 uint32 = 1
                var v6 float64 = float64(v1)
                _ = v5
                _ = v6
            } else {
                if true {
                    var v7 uint16 = 1
                    var v8 float64 = 1e10
                    _ = v7
                    _ = v8
                }
                var v9 int = v2
                var v10 int = (v2 * (v9 & (v2 % (((v9 - v9)) | 1))))
                _ = v10
                _ = v9
            }
            v4 += (v2 & v4)
            fmt.Println(v4, v4)
        } else {
            if true {
                var v11 uint16 = 255
                v12 := GenS1{F0: 255}
                _ = v11
                _ = v12
            } else {
                var v13 bool = false
                v14 := GenS1{F0: 4294967295}
                var v15 int = (v2 + (v2 | (-(v4))))
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 uint32 = 1
            _ = v16
        }
    case "hello":
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 18; fi18++ {
                if (!true) {
                    v19 := []int{v4}
                    _ = v19
                }
                fmt.Println(fi18, v4)
                _ = fi18
            }
            v4 += (^(v4))
            _ = fi17
        }
    case "a":
        var v20 uint = 90
        _ = v20
    default:
        v21 := []int{(v4 & (v4 & 879)), (v2 | v2), -1}
        _ = v21
    }
    v22 := GenS2{F0: 1, F1: v4, F2: (v2 << ((65535) & 63))}
    v23 := GenS1{F0: 65535}
    var v24 int = (v4 + (v4 + v4))
    var v25 int8 = 95
    switch "" {
    case "hello":
        v24 = (v4 ^ 0)
    case "\n\t":
        var v26 string = "hello"
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        v28 := GenS3{F0: 1}
        _ = fi27
        _ = v28
    }
    var v29 int64 = 125
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
