package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint, p1 uint32, p2 string) bool {
    p0 = (p0 >> ((p0) & 63))
    v5 := []int{2147483647, 4294967295, -9223372036854775808}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return false
}

func main() {
    var v1 uint16 = 1
    var v2 uint32 = 556
    v3 := make(map[uint32]bool)
    var v4 bool = (((!(1.5 > float32(float64(1.5)))) || false) || true)
    var v6 uint8 = 45
    switch int(v1) {
    case 32768:
        switch 128 {
        case 128:
            var v7 uint8 = 1
            _ = v7
        default:
            _ = len(v3)
        }
    default:
        var v8 int16 = 490
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            switch fi9 {
            case 128:
                v1 = (v1 & 186)
            case -9223372036854775808:
                fmt.Println(v6)
            default:
                if true {
                    var v11 uint8 = (v6 << ((4294967295) & 63))
                    var v12 uint8 = 128
                    _ = v11
                    _ = v12
                }
            }
            if false {
                v4 = v4
            } else {
                v13 := GenS1{F0: -1, F1: (v1 + v1), F2: "hello"}
                _ = v13
            }
            _ = fi10
        }
        v14 := []int{128, fi9}
        _ = fi9
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v2 = 0
        var v16 uint8 = (^(v6))
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS2{F0: 645}
            if v4 {
                var v19 uint64 = 4294967295
                var v20 int16 = 222
                _ = v19
                _ = v20
            }
            v16 = (^(v16))
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    var v21 uint8 = 2
    var v22 string = "hello"
    if v4 {
        var v23 int = 1
        _ = v23
    } else {
        var v24 string = v22
        var v25 float32 = 0.0
        if v4 {
            fmt.Println(v25)
        }
        _ = v24
        _ = v25
    }
    fmt.Println(v6)
    v4 = v4
    var v26 uint16 = (v1 ^ (v1 << ((18446744073709551615) & 63)))
    v27 := GenS3{F0: 42}
    fmt.Println(v1, v1, v1)
    switch -2147483648 {
    case 248:
        fmt.Println(v21)
    case 128:
        for fi28 := 0; fi28 < 7; fi28++ {
            v6 -= (v21 << ((42) & 63))
            for fi29 := 0; fi29 < 19; fi29++ {
                if v4 {
                    var v30 uint16 = (v26 - v1)
                    var v31 int8 = 127
                    _ = v30
                    _ = v31
                } else {
                    var v32 int64 = -1
                    var v33 int8 = 127
                    fmt.Println(v22, fi28)
                    _ = v32
                    _ = v33
                }
                _ = fi29
            }
            v34 := v3[(v2 - v2)]
            _ = fi28
            _ = v34
        }
    default:
        var v35 uint32 = 541
        _ = v35
    }
    v22 = (("" + v22) + (v22 + ((v22 + v22) + "hello")))
    if (0 == (v2 >> ((2) & 63))) {
        _ = len(v3)
        var v36 uint = 4294967295
        var v37 string = "foo bar"
        _ = v36
        _ = v37
    }
    v4 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v6
}
