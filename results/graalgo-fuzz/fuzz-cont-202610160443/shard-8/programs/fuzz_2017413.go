package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int16 = 686
    var v2 string = "foo bar"
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 uint = 255
    v6 := v4[32767]
    switch -1 {
    case 42:
        v7 := []int{(v6 - v6), v6}
        _ = v7
    case 381:
        if false {
            var v8 int32 = 243
            v8 -= 2147483647
            _ = v8
        }
    }
    var v9 int32 = -609
    switch v6 {
    case 0:
        v10 := v3[32767]
        _ = v10
    case 9223372036854775807:
        var v11 bool = false
        _ = v11
    case 32768:
        v2 = v2
    default:
        v6 = 294
    }
    fmt.Println(v5, v5)
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 int64 = 47
            v9 = v9
            v4[-571] = 32767
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: v1, F1: 128}
        _ = fi12
        _ = v15
    }
    v16 := GenS2{F0: v1, F1: -1}
    switch -2147483648 {
    case -131:
        var v17 string = v2
        _ = v17
    case 32767:
        if (((v2 + v2) >= (((v2 + ("hello" + v2)) + v2) + v2)) && false) {
            v18 := GenS1{F0: 463, F1: 0}
            fmt.Println(v2, v5, v2)
            _ = v18
        } else {
            v2 = "foo bar"
            v19 := GenS3{F0: 255}
            v3[v6] = "hello"
            _ = v19
        }
    case 2147483647:
        var v20 int32 = (v9 & v9)
        _ = v20
    }
    v6 = (-(v6))
    for fi21 := 0; fi21 < 4; fi21++ {
        if false {
            fmt.Println(v9, v1)
        }
        v22 := []int{(fi21 << (((v5 % (((^(v5))) | 1))) & 63))}
        _ = fi21
        _ = v22
    }
    var v23 int32 = (v9 * -151)
    v6 += (v6 >> ((912) & 63))
    var v24 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
