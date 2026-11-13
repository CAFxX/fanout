package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int; F2 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := []uint32{42, 994}
    var v4 uint64 = 1
    fmt.Println(v4)
    fmt.Println(v4)
    v5 := GenS1{F0: 42, F1: 42, F2: 89}
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint8 = 91
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: 579}
    switch 0 {
    case -2147483648:
        v4 = (v4 - v4)
    case -616:
        v4 = (v4 + v4)
    case -948:
        var v9 uint = 18446744073709551615
        _ = v9
    }
    var v10 int8 = -23
    v10 = v10
    v10 = (v10 << ((104) & 63))
    if false {
        if ((v10 / (((v10 - (v10 - 100))) | 1)) < v10) {
            v4 = 2
        } else {
            v4 -= (v4 % ((v4) | 1))
            v10 += 1
        }
        var v11 int16 = 0
        _ = v11
    } else {
        if (!false) {
            var v12 int = 685
            if (!(0 <= 1)) {
                if (65535 <= 65535) {
                    switch v12 {
                    case 773:
                        var v13 int = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 18446744073709551615, F1: v10, F2: 1}
                    _ = v14
                }
                v12 += 50
            }
            var v15 float32 = ((0.0 + float32(1.5)) * 1e10)
            _ = v12
            _ = v15
        } else {
            var v16 bool = false
            var v17 bool = false
            switch "\n\t" {
            case "foo bar":
                v17 = (42 == 2147483648)
            case "\n\t":
                if (4294967295 != 489) {
                    var v18 uint = 689
                    v18 -= v18
                    var v19 int8 = 0
                    _ = v18
                    _ = v19
                }
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v16
            _ = v17
        }
        v20 := GenS3{F0: -1}
        for fi21 := 0; fi21 < 14; fi21++ {
            v22 := GenS4{F0: 100, F1: fi21, F2: 59}
            var v23 float32 = (1e10 - 1e10)
            v10 -= (v10 >> ((255) & 63))
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = v20
    }
    var v24 bool = (!("a" < (((("hello" + "foo bar") + "") + (("hello" + "\n\t") + "")) + ("" + "a"))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
