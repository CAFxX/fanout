package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 uint16, p1 uint32, p2 int8) bool {
    p2 = (p2 + p2)
    v7 := GenS1{F0: 77, F1: (p2 >> ((674) & 63)), F2: 949}
    p1 += uint32(p0)
    switch 32768 {
    case 878:
        v8 := GenS2{F0: 0, F1: 191}
        _ = v8
    case -2147483648:
        if (((-2.25 * 1e10) / (-2.25 - ((0.1 - 3.14159) / (1e10 * float32((1e10 - 0.1)))))) == 0.0) {
            switch 32768 {
            case 0:
                fmt.Println(p2, p2, p2)
            case -454:
                p2 -= 24
            }
        }
    case 382:
        var v9 float32 = 1.5
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    var v3 int16 = 42
    var v4 float32 = float32(3.14159)
    var v5 bool = true
    var v6 uint32 = 0
    if (-2.25 < (3.14159 + float64((float32(0.1) / 1.5)))) {
        v5 = (!true)
    } else {
        v10 := GenS1{F0: 1, F1: 47, F2: 32768}
        _ = v10
    }
    fmt.Println(v6, v4)
    if false {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint32 = (v6 << ((65535) & 63))
            var v13 uint32 = 801
            var v14 float32 = 1.5
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    }
    var v15 bool = (int(v3) != -1)
    v6 = 2147483648
    var v16 int32 = -1
    v6 += (^(v6))
    var v17 int32 = -1
    v18 := GenS3{F0: false, F1: v17}
    var v19 int = -1
    fmt.Println(v17)
    if ((v17 - 1) < (v17 & 0)) {
        fmt.Println(v5)
        switch -1 {
        case -598:
            for fi20 := 0; fi20 < 14; fi20++ {
                v19 = (fi20 ^ (fi20 >> ((18446744073709551615) & 63)))
                var v21 uint16 = 2
                _ = fi20
                _ = v21
            }
        }
    } else {
        v5 = v15
    }
    var v22 uint8 = 0
    var v23 float64 = 1.5
    var v24 string = ("foo bar" + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
