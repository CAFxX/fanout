package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint16; F1 int8; F2 int }
func genFn0(p0 int16) string {
    v7 := []int{2147483647, 2147483647}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    p0 += (p0 & -1)
    _ = p0
    _ = v7
    return ("" + "foo bar")
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 3.14159
    var v3 bool = false
    var v4 int8 = 42
    var v5 int8 = v4
    var v6 int64 = 0
    v3 = ((-548 <= 2147483647) || false)
    v2 -= 0.0
    var v8 string = "hello"
    v4 -= (v4 % (((v4 & v5)) | 1))
    v9 := GenS1{F0: 2}
    if (!(127 >= 45)) {
        v1[847] = -937
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int16 = 1
        v12 := GenS1{F0: 0}
        v1[65535] = fi10
        _ = fi10
        _ = v11
        _ = v12
    }
    v4 += (v4 >> ((18446744073709551615) & 63))
    v1[65535] = 9223372036854775807
    var v13 int64 = v6
    if v3 {
        if ("" != (v8 + v8)) {
            v4 -= (v4 | (v5 - v4))
            fmt.Println(v5, v6, v6)
        } else {
            v14 := []int{128, -1}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = v14
        }
        var v15 int = 870
        _ = v15
    } else {
        if ((0 < 128) && ((-(v4)) != v5)) {
            fmt.Println(v6, v5, v13)
            if v3 {
                fmt.Println(v13, v5, v2)
                v16 := v1[366]
                v4 += (v4 << ((522) & 63))
                _ = v16
            }
            fmt.Println(v8)
        } else {
            v17 := GenS2{F0: 0, F1: v6}
            v3 = (32767 <= 1)
            _ = v17
        }
        v8 = v8
        var v18 int8 = (v4 | v4)
        _ = v18
    }
    switch 2147483647 {
    case -102:
        if (2 > 2) {
            switch v8 {
            case "a":
                v19 := GenS3{F0: 32767, F1: v4, F2: int(v4)}
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := v1[0]
                _ = fi20
                _ = v21
            }
        }
    }
    var v22 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
