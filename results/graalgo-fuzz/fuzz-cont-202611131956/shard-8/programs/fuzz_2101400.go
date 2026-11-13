package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int64 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint8; F2 int8 }
func genFn0() uint8 {
    var v4 float32 = 0.0
    switch 0 {
    case 42:
        var v5 string = "foo bar"
        _ = v5
    }
    _ = v4
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 int8 = 1
    fmt.Println(v3)
    var v6 float32 = 0.0
    fmt.Println(v6)
    if true {
        v7 := GenS1{F0: 2, F1: true, F2: -1}
        _ = v7
    }
    var v8 uint = 262
    switch -977 {
    case 1:
        fmt.Println(v3, v3)
    case -1:
        fmt.Println(v8, v8)
    case 0:
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 string = "foo bar"
        v8 = (v8 / (((v8 % (((^(v8))) | 1))) | 1))
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 42, F1: -660}
    for fi13 := 0; fi13 < 4; fi13++ {
        if false {
            v14 := []int{fi13, 409}
            var v15 uint32 = 136
            _ = v14
            _ = v15
        } else {
            if true {
                var v16 int64 = 1
                for fi17 := 0; fi17 < 13; fi17++ {
                    fmt.Println(fi13, v6)
                    _ = fi17
                }
                _ = v16
            }
            fmt.Println(v3)
        }
        _ = fi13
    }
    var v18 uint32 = 249
    v8 += v8
    v8 = v8
    for fi19 := 0; fi19 < 1; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 int32 = -2147483648
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 18; fi22++ {
            v8 -= 4294967295
            v23 := GenS2{F0: 0, F1: fi22}
            switch "hello" {
            case "foo bar":
                v24 := GenS2{F0: 42, F1: (fi22 - (fi22 >> ((v8) & 63)))}
                _ = v24
            case "hello":
                v25 := GenS3{F0: 42}
                _ = v25
            }
            _ = fi22
            _ = v23
        }
        var v26 uint16 = 42
        _ = fi19
        _ = v26
    }
    v27 := []int{959, int(v8)}
    switch 127 {
    case -1:
        v3 -= int8(v8)
    default:
        fmt.Println(v8, v18, v8)
    }
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    v6 = 1.5
    v28 := GenS4{F0: (v8 ^ v8), F1: 0, F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v6
    _ = v8
}
