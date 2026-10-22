package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 bool }
type GenS3 struct { F0 int8; F1 uint64; F2 uint16 }
func genFn0(p0 uint64, p1 int64, p2 int8) bool {
    v5 := GenS1{F0: 18446744073709551615}
    var v6 string = ""
    var v7 int16 = -1
    p0 = 444
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (true || false)
}
func genFn1(p0 uint32, p1 bool, p2 int16) bool {
    fmt.Println(p0)
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint8 = 128
    v2 := []uint32{}
    var v3 uint8 = (v1 << ((18446744073709551615) & 63))
    v4 := make(map[uint32]bool)
    fmt.Println(v3, v3)
    v8 := GenS2{F0: 2147483647, F1: 621, F2: false}
    fmt.Println(v3)
    fmt.Println(v3)
    var v9 float64 = (1.5 + (0.0 - -2.25))
    if (42 == -571) {
        var v10 uint8 = v1
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS3{F0: 0, F1: 18446744073709551615, F2: 65535}
            _ = fi11
            _ = v12
        }
        if true {
            if false {
                var v13 int = -458
                var v14 float32 = -2.25
                fmt.Println(v1, v14)
                _ = v13
                _ = v14
            } else {
                var v15 float32 = (float32((-2.25 + float64(3.14159))) / (0.0 + -2.25))
                for fi16 := 0; fi16 < 3; fi16++ {
                    fmt.Println(v15)
                    _ = fi16
                }
                var v17 float32 = v15
                _ = v15
                _ = v17
            }
            fmt.Println(v10, v10, v1)
        } else {
            if (float64(0.1) != ((v9 + 0.1) / v9)) {
                fmt.Println(v3)
            } else {
                v10 -= 2
                if len(v2) > 0 {
                    v2[0] = 0
                }
                v18 := v4[645]
                _ = v18
            }
        }
        _ = v10
    }
    v3 = 42
    v3 -= 255
    v1 = (v3 - (v1 >> ((18446744073709551615) & 63)))
    var v19 uint = 0
    var v20 uint64 = 442
    switch -807 {
    case 0:
        var v21 int8 = 1
        _ = v21
    case -735:
        v3 += v1
    case 2147483647:
        fmt.Println(v19, v3, v19)
    }
    fmt.Println(v19, v3)
    fmt.Println(v3, v1)
    var v22 int8 = 80
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
