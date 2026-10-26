package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int16 }
func genFn0() bool {
    v4 := []int{}
    if (2 > 2) {
        var v5 int32 = 23
        v5 += 42
        var v6 int = 535
        _ = v5
        _ = v6
    }
    var v7 int16 = 300
    if true {
        switch -580 {
        case 4294967295:
            if false {
                var v8 uint16 = 411
                if len(v4) > 0 {
                    v4[0] = -2147483648
                }
                _ = v8
            }
        case -1:
            var v9 uint8 = 127
            _ = v9
        case -611:
            v10 := GenS1{F0: 65535}
            _ = v10
        default:
            fmt.Println(v7)
        }
        if false {
            var v11 int = -9223372036854775808
            _ = v11
        }
        if ((float64((0.1 + 3.14159)) - (1.5 * (0.0 - (float64(float32(3.14159)) / (-2.25 + 0.0))))) != 1e10) {
            if (0 != -1) {
                v12 := GenS1{F0: 42}
                fmt.Println(v7)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v12
            }
        }
    }
    _ = v4
    _ = v7
    return false
}

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    v13 := v2[0]
    if (-95 <= -1) {
        var v14 uint16 = 0
        var v15 uint = 65535
        _ = v14
        _ = v15
    } else {
        v16 := []int{4294967295, -615}
        v13 = 2147483648
        _ = v16
    }
    var v17 int16 = 0
    var v18 uint32 = v13
    fmt.Println(v13)
    if (1 == int64(v1)) {
        var v19 uint16 = 0
        v18 = 65535
        if (v19 == (v19 << ((65535) & 63))) {
            for fi20 := 0; fi20 < 10; fi20++ {
                v21 := v2[v18]
                _ = fi20
                _ = v21
            }
            var v22 uint64 = 2
            var v23 uint16 = 42
            _ = v22
            _ = v23
        } else {
            v24 := GenS2{F0: (("hello" > ((("foo bar" + ("foo bar" + "")) + ("foo bar" + "foo bar")) + (("foo bar" + "foo bar") + "hello"))) || (false || (-2.25 <= 1.5))), F1: (-(v17))}
            _ = v24
        }
        _ = v19
    }
    v13 += 0
    if (true || false) {
        v25 := GenS2{F0: true, F1: v17}
        for fi26 := 0; fi26 < 18; fi26++ {
            _ = len(v3)
            var v27 int16 = (v17 ^ v17)
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
