package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS3 struct { F0 bool }
func genFn0(p0 int64) uint8 {
    v4 := []int{-2147483648, 32767}
    var v5 string = "a"
    v5 = "\n\t"
    if true {
        _ = len(v4)
    }
    _ = p0
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int32, p1 uint64, p2 int) int16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(p2, p1, p2)
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int16 = -32768
            _ = fi7
            _ = v8
        }
        switch "" {
        case "a":
            var v9 uint64 = 0
            _ = v9
        default:
            v10 := GenS1{F0: "\n\t"}
            _ = v10
        }
        _ = fi6
    }
    v11 := []int{-1, p2, -1, (p2 | (p2 ^ (p2 >> ((86) & 63))))}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    var v1 int16 = 32767
    var v2 uint8 = 127
    v3 := []uint32{4294967295, 802}
    v1 += genFn1(597, 124, -2147483648)
    v1 = v1
    for fi12 := 0; fi12 < 19; fi12++ {
        switch fi12 {
        case -9223372036854775808:
            var v13 float32 = 0.1
            _ = v13
        default:
            fmt.Println(v2, v2, v2)
        }
        v14 := GenS1{F0: "\n\t"}
        if (float64(((float32(((0.0 - 1e10) * 0.0)) / 1e10) * (-2.25 + float32(((-2.25 - -2.25) - (1.5 - ((float64(1e10) / 1.5) * 3.14159))))))) != 1.5) {
            v2 = v2
            var v15 bool = true
            _ = v15
        } else {
            var v16 uint32 = 65535
            v1 -= 0
            for fi17 := 0; fi17 < 20; fi17++ {
                v2 -= 1
                _ = fi17
            }
            _ = v16
        }
        _ = fi12
        _ = v14
    }
    fmt.Println(v2, v2)
    v18 := GenS1{F0: ((("\n\t" + (("foo bar" + (("hello" + "") + "foo bar")) + ("foo bar" + "\n\t"))) + "foo bar") + "\n\t")}
    if ((42 != 0) || false) {
        v2 -= v2
        v2 = (v2 >> ((42) & 63))
    } else {
        var v19 uint32 = 65535
        v19 += v19
        _ = v19
    }
    v2 -= v2
    if (!true) {
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := GenS2{F0: 2, F1: 61, F2: (v2 | 127)}
            v22 := GenS3{F0: true}
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 uint16 = 42
        _ = v23
    }
    var v24 int64 = 1
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 bool = false
        _ = fi25
        _ = v26
    }
    var v27 uint16 = 662
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v27
    _ = v3
}
