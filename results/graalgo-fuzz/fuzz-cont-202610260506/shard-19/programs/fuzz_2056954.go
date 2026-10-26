package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool) string {
    v5 := GenS1{F0: 136}
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int8 = -1
        var v8 int64 = -1
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    return (("" + "\n\t") + (("\n\t" + ("\n\t" + ("a" + "\n\t"))) + ("a" + ("foo bar" + ""))))
}
func genFn1(p0 int8, p1 int16) int16 {
    fmt.Println(p0, p1)
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(p0, fi9)
        _ = fi9
    }
    p1 -= 32767
    p1 += -32768
    for fi10 := 0; fi10 < 3; fi10++ {
        p1 = p1
        _ = fi10
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]bool)
    var v3 int8 = 94
    var v4 int16 = 0
    v11 := []int{-2147483648}
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS1{F0: 18446744073709551615}
        fmt.Println(v4)
        if ((true || false) && (genFn1(v3, v4) > v4)) {
            var v14 float64 = 0.0
            v15 := GenS2{F0: v1}
            var v16 int32 = 42
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v4, v1)
    switch 1 {
    case -225:
        var v17 uint32 = 886
        _ = v17
    case 32768:
        var v18 int8 = v3
        _ = v18
    default:
        if (65535 <= 0) {
            v1 = (v1 * v1)
        }
    }
    var v19 bool = (!(!(float32(1e10) < -2.25)))
    v4 = -59
    if (0 <= (v3 * v3)) {
        v1 += (v1 << ((2) & 63))
        fmt.Println(v4)
        for fi20 := 0; fi20 < 1; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS2{F0: (v1 - 2)}
                delete(v2, 2147483648)
                v23 := GenS1{F0: 685}
                v19 = (1e10 == 1.5)
                v4 = v4
                _ = fi21
                _ = v22
                _ = v23
            }
            v24 := v2[883]
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := GenS2{F0: v1}
                _ = fi25
                _ = v26
            }
            _ = fi20
            _ = v24
        }
    }
    v1 = 65535
    var v27 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
