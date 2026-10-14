package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int32 }

func main() {
    var v1 string = "hello"
    var v2 string = (v1 + (v1 + v1))
    var v3 uint = 227
    var v4 uint64 = 424
    var v5 int = -322
    v6 := make(map[string]uint32)
    v4 -= (v4 >> (((v3 & (v3 << (((^(v3))) & 63)))) & 63))
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int8 = int8(v3)
        for fi9 := 0; fi9 < 8; fi9++ {
            switch 349 {
            case 1:
                var v10 int32 = 42
                _ = v10
            case 865:
                var v11 uint64 = v4
                _ = v11
            default:
                fmt.Println(v5)
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v3, v4, v1)
    var v12 int64 = 331
    var v13 float64 = -2.25
    switch v5 {
    case 2147483647:
        if true {
            v12 -= v12
            for fi14 := 0; fi14 < 8; fi14++ {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v1 = v2
                    var v16 uint16 = 42
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
            fmt.Println(v1)
        }
    case -2147483648:
        if false {
            switch (v2 + (v2 + ("\n\t" + v1))) {
            case "hello":
                v17 := GenS1{F0: 42, F1: 969}
                _ = v17
            }
            v18 := []int{}
            fmt.Println(v2, v2, v4)
            _ = v18
        } else {
            switch "hello" {
            case "foo bar":
                var v19 string = v2
                _ = v19
            case "hello":
                switch v2 {
                case "":
                    switch v1 {
                    case "foo bar":
                        v20 := GenS2{F0: 591, F1: 42, F2: 633}
                        _ = v20
                    case "a":
                        var v21 int = (v5 * (v5 & (v5 / ((v5) | 1))))
                        _ = v21
                    case "hello":
                        v22 := GenS3{F0: (v4 & (v4 - 4294967295))}
                        _ = v22
                    default:
                        var v23 float32 = -2.25
                        _ = v23
                    }
                }
            case "a":
                var v24 uint8 = 128
                _ = v24
            }
            v25 := GenS4{F0: 2147483647, F1: 2147483647}
            _ = v25
        }
    case 128:
        v26 := GenS2{F0: 65535, F1: 128, F2: (v3 / ((uint(v4)) | 1))}
        _ = v26
    default:
        var v27 uint16 = 2
        _ = v27
    }
    fmt.Println(v12, v5, v2)
    for fi28 := 0; fi28 < 4; fi28++ {
        v13 = (1.5 / v13)
        v29 := []int{int(v4), v5, (v5 % ((v5) | 1)), (fi28 / ((128) | 1))}
        _ = fi28
        _ = v29
    }
    var v30 int64 = (v12 >> ((v3) & 63))
    for fi31 := 0; fi31 < 2; fi31++ {
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 float64 = -2.25
            _ = len(v6)
            var v34 int32 = -2147483648
            v30 += (v30 ^ (^(v30)))
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    var v35 int64 = (-(v30))
    var v36 float32 = (float32(3.14159) / 3.14159)
    v37 := GenS4{F0: 87, F1: 0}
    if true {
        v2 = "hello"
        v4 -= 469
        var v38 int = int(v4)
        _ = v38
    } else {
        v39 := GenS4{F0: 917, F1: 42}
        _ = v39
    }
    var v40 uint64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
