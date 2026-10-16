package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int; F1 int; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    var v5 uint16 = (v1 ^ 1)
    if (18 < 1) {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 float32 = 1.5
            v7 = v7
            fmt.Println(fi6, fi6, v5)
            _ = fi6
            _ = v7
        }
        switch "\n\t" {
        case "foo bar":
            if false {
                var v8 uint64 = 408
                v9 := GenS1{F0: -1, F1: 0}
                for fi10 := 0; fi10 < 14; fi10++ {
                    var v11 float64 = (float64(-2.25) - float64(0.0))
                    v3[2147483647] = (!(32767 == 1))
                    if (0 >= 187) {
                        v1 = 65535
                    } else {
                        v11 = (0.0 / v11)
                    }
                    _ = fi10
                    _ = v11
                }
                _ = v8
                _ = v9
            } else {
                var v12 int = 127
                var v13 uint16 = (v1 & (v5 ^ (v1 * (v5 ^ 65535))))
                _ = v12
                _ = v13
            }
        default:
            v14 := GenS2{F0: int8(v5), F1: 1}
            _ = v14
        }
        fmt.Println(v1, v5)
    } else {
        _ = len(v2)
    }
    var v15 uint32 = 1
    v16 := GenS3{F0: 42, F1: 69, F2: 1}
    v5 -= v1
    switch ("hello" + "\n\t") {
    case "\n\t":
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS4{F0: ("foo bar" + ("foo bar" + ("foo bar" + "\n\t")))}
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v1 = (v5 >> ((2) & 63))
        var v20 uint32 = v15
        v5 = (^(v1))
        _ = fi19
        _ = v20
    }
    v2[v15] = (404 < -540)
    if (v15 >= 619) {
        var v21 int64 = -751
        var v22 uint16 = v5
        var v23 string = ("foo bar" + "\n\t")
        _ = v21
        _ = v22
        _ = v23
    } else {
        v1 = (^(v1))
        delete(v4, "")
        fmt.Println(v5, v1, v1)
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        v1 += 32767
    default:
        v15 -= 4294967295
    }
    v15 -= v15
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 15; fi25++ {
            _ = len(v4)
            v1 += v5
            fmt.Println(v15)
            v5 = (v5 & (v1 - v5))
            v15 = (v15 - v15)
            _ = fi25
        }
        _ = fi24
    }
    var v26 int16 = -1
    for fi27 := 0; fi27 < 18; fi27++ {
        for fi28 := 0; fi28 < 10; fi28++ {
            v29 := GenS5{F0: (fi28 % ((int(v1)) | 1))}
            v30 := v2[v15]
            switch fi27 {
            case 0:
                var v31 int8 = 100
                _ = v31
            default:
                v5 += (^(v1))
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        v15 -= 2147483647
        _ = fi27
    }
    fmt.Println(v15)
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 float64 = (-2.25 * 0.1)
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        fmt.Println(v26, v1)
        _ = fi34
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := []int{fi35, (^(fi35)), fi35}
        _ = fi35
        _ = v36
    }
    v26 += (^(v26))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
