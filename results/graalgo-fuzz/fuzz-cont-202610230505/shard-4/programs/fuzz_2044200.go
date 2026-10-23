package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint8; F1 string; F2 int64 }
type GenS4 struct { F0 int8; F1 bool }
func genFn0(p0 bool, p1 int64, p2 uint8) int8 {
    for fi7 := 0; fi7 < 13; fi7++ {
        switch 127 {
        case 9223372036854775807:
            v8 := GenS1{F0: 1, F1: false, F2: 255}
            _ = v8
        case 32767:
            p2 = (p2 % ((217) | 1))
        }
        switch "\n\t" {
        case "hello":
            switch fi7 {
            case 2147483647:
                v9 := []int{}
                _ = v9
            default:
                v10 := GenS2{F0: -659, F1: (-(p1))}
                _ = v10
            }
        case "\n\t":
            v11 := GenS1{F0: 853, F1: ((fi7 & -1) > -676), F2: p2}
            _ = v11
        case "foo bar":
            switch "\n\t" {
            case "foo bar":
                v12 := GenS3{F0: (^(p2)), F1: (("hello" + (("\n\t" + "\n\t") + "")) + "\n\t"), F2: (p1 << ((2) & 63))}
                _ = v12
            case "\n\t":
                var v13 int16 = 42
                _ = v13
            }
        }
        var v14 uint16 = 2
        _ = fi7
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int32 = 1
        _ = fi15
        _ = v16
    }
    if p0 {
        fmt.Println(p0, p2, p0)
        if p0 {
            p0 = p0
        } else {
            v17 := []int{9223372036854775807}
            v18 := GenS3{F0: p2, F1: ("hello" + "\n\t"), F2: 42}
            _ = v17
            _ = v18
        }
        if p0 {
            for fi19 := 0; fi19 < 10; fi19++ {
                p0 = (460 == 42)
                if (0 == 9223372036854775807) {
                    var v20 uint8 = p2
                    v21 := GenS1{F0: 42, F1: ((722 <= 32767) && (!(false && p0))), F2: 210}
                    p0 = false
                    _ = v20
                    _ = v21
                }
                var v22 int8 = -1
                _ = fi19
                _ = v22
            }
        } else {
            fmt.Println(p0, p0)
            for fi23 := 0; fi23 < 19; fi23++ {
                p0 = (1e10 <= float64(1e10))
                _ = fi23
            }
            fmt.Println(p0)
        }
    } else {
        fmt.Println(p0)
    }
    v24 := GenS3{F0: p2, F1: "", F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 1
}

func main() {
    var v1 uint8 = 0
    var v2 int64 = 166
    var v3 int64 = 1
    v4 := []uint32{958, 2147483647, 2}
    v5 := make(map[uint32]int)
    var v6 int16 = 10
    _ = len(v5)
    var v25 uint16 = 255
    fmt.Println(v6, v3, v3)
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 int8 = -128
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 uint16 = v25
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        v6 += (v6 << ((255) & 63))
        fmt.Println(v25, v2, v25)
        _ = fi30
    }
    if ((-2.25 < 1e10) && false) {
        var v31 uint = 1
        v2 = 42
        _ = v31
    } else {
        v32 := GenS4{F0: genFn0((!(!false)), -674, 197), F1: (148 <= 0)}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        for fi34 := 0; fi34 < 7; fi34++ {
            var v35 uint = 65535
            switch fi34 {
            case 128:
                var v36 int8 = -119
                _ = v36
            case 9223372036854775807:
                v35 -= uint(v1)
            default:
                v3 = 9223372036854775807
            }
            _ = fi34
            _ = v35
        }
        fmt.Println(fi33, v6, v2)
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
