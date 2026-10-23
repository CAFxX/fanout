package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{418, 9223372036854775807, 128}
    v3 := make(map[string]int)
    var v4 int = 127
    v5 := GenS1{F0: 42}
    for fi6 := 0; fi6 < 8; fi6++ {
        _ = len(v2)
        v7 := GenS2{F0: 4294967295, F1: 1, F2: 1}
        if (false && (42 >= 2147483648)) {
            if (0 < 1) {
                var v8 string = ("a" + (("hello" + "") + ((("hello" + "a") + "") + ("a" + ("" + "hello")))))
                fmt.Println(v8)
                v4 -= -1
                _ = v8
            }
        }
        _ = fi6
        _ = v7
    }
    var v9 float64 = ((1e10 * -2.25) - ((3.14159 - -2.25) / 3.14159))
    v10 := GenS3{F0: 42}
    if (65535 <= 32767) {
        for fi11 := 0; fi11 < 4; fi11++ {
            for fi12 := 0; fi12 < 20; fi12++ {
                if true {
                    _ = len(v3)
                }
                v4 = (fi11 >> ((18446744073709551615) & 63))
                if len(v2) > 0 {
                    v2[0] = fi12
                }
                _ = fi12
            }
            if false {
                v9 = float64(((0.1 + 0.1) * (1e10 - 0.0)))
            } else {
                v4 += -9223372036854775808
                var v13 int8 = 127
                _ = v13
            }
            _ = fi11
        }
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v9 = v9
        v15 := GenS1{F0: 216}
        v16 := GenS3{F0: 0}
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := v3[("a" + "hello")]
        v9 -= (1e10 - 0.1)
        var v19 float64 = float64(3.14159)
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v9 = -2.25
        var v21 int8 = 42
        v22 := GenS3{F0: 90}
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println(v4, v4)
    v9 = v9
    var v23 int64 = 42
    var v24 int16 = 348
    for fi25 := 0; fi25 < 15; fi25++ {
        var v26 int16 = v24
        fmt.Println(v24, v24, fi25)
        _ = fi25
        _ = v26
    }
    var v27 uint32 = 541
    v3["hello"] = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
