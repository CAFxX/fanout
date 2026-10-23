package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int32) int16 {
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint64 = 1
        _ = fi5
        _ = v6
    }
    v7 := []int{42}
    v8 := GenS1{F0: 285}
    for fi9 := 0; fi9 < 6; fi9++ {
        switch 32768 {
        case -62:
            var v10 int32 = 42
            _ = v10
        case -9223372036854775808:
            if (true || true) {
                var v11 uint64 = 0
                _ = v11
            } else {
                p0 -= (p0 + (-(p0)))
                p0 = p0
            }
        default:
            for fi12 := 0; fi12 < 10; fi12++ {
                p0 = -2147483648
                _ = len(v7)
                if len(v7) > 0 {
                    v7[0] = (fi12 - fi12)
                }
                p0 = -54
                if len(v7) > 0 {
                    v7[0] = fi12
                }
                fmt.Println(p0, fi9)
                _ = fi12
            }
        }
        _ = fi9
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v15 := GenS1{F0: -711}
            v7 = append(v7, -2147483648)
            _ = fi14
            _ = v15
        }
        if false {
            fmt.Println(p0)
            if false {
                var v16 uint32 = 2147483647
                v17 := GenS1{F0: -734}
                p0 = p0
                _ = v16
                _ = v17
            }
        }
        if (true || true) {
            fmt.Println(p0)
            var v18 float64 = (float64(((3.14159 + 1e10) / (-2.25 - 3.14159))) - float64(0.1))
            _ = v18
        }
        _ = fi13
    }
    _ = p0
    _ = v7
    _ = v8
    return 1000
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 float64 = float64((3.14159 * -2.25))
    var v4 int8 = 100
    v19 := GenS1{F0: 32767}
    var v20 uint8 = 255
    var v21 string = ("foo bar" + "hello")
    v3 += 3.14159
    v4 = v4
    v20 = v20
    v21 = (v21 + "a")
    if true {
        var v22 uint = 2
        for fi23 := 0; fi23 < 1; fi23++ {
            v3 -= v3
            var v24 int8 = (v4 % ((100) | 1))
            _ = fi23
            _ = v24
        }
        if false {
            var v25 string = (v21 + ("hello" + ("" + ("hello" + ""))))
            v26 := []int{4294967295, -970}
            _ = v25
            _ = v26
        }
        _ = v22
    }
    fmt.Println(v3)
    var v27 bool = (-1 <= 438)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
}
