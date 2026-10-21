package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (float64(0.1) / 1.0)
    var v3 uint16 = 0
    v4 := []string{}
    if (0 >= 1) {
        v3 -= (v3 - 42)
        v3 -= 2
        v3 += 32767
    }
    var v5 uint64 = 18446744073709551615
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 uint = 2
    fmt.Println(v6, v3)
    v7 := GenS1{F0: (v6 == v6), F1: 65535}
    var v8 int64 = 9223372036854775807
    var v9 int32 = -919
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1[2] = false
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS2{F0: 42}
        if true {
            if ((v8 + (v8 << (((^(v6))) & 63))) >= v8) {
                fmt.Println(fi10)
                _ = len(v4)
                v5 = (v5 >> ((65535) & 63))
            }
        } else {
            var v12 int16 = 927
            switch 9223372036854775807 {
            case -1:
                var v13 int64 = (v8 ^ v8)
                _ = v13
            case -9223372036854775808:
                for fi14 := 0; fi14 < 6; fi14++ {
                    v9 = (v9 >> ((v6) & 63))
                    fmt.Println(fi14)
                    _ = fi14
                }
            case -577:
                fmt.Println(v8, v9)
            }
            _ = v12
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 bool = true
            if true {
                if false {
                    v3 = uint16(v5)
                    if v16 {
                        var v17 int32 = int32(v3)
                        _ = v17
                    }
                } else {
                    var v18 int = (^(fi15))
                    fmt.Println(v3, v3)
                    _ = v18
                }
                var v19 uint32 = 826
                _ = v19
            } else {
                if len(v4) > 0 {
                    v4[0] = ""
                }
                if (v5 != v5) {
                    var v20 uint = (v6 >> ((v6) & 63))
                    _ = len(v1)
                    fmt.Println(fi10)
                    _ = v20
                } else {
                    var v21 float64 = -2.25
                    v8 -= (v8 / ((v8) | 1))
                    v16 = false
                    _ = v21
                }
            }
            fmt.Println(v3)
            var v22 string = ("a" + "hello")
            _ = fi15
            _ = v16
            _ = v22
        }
        _ = fi10
        _ = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
