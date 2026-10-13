package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint32, p1 uint, p2 int32) int {
    p2 = p2
    v4 := GenS1{F0: "\n\t", F1: 278, F2: 0}
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint32 = p0
        _ = fi5
        _ = v6
    }
    if ((0 <= 25) || (!true)) {
        if (1 >= 42) {
            var v7 float32 = 1e10
            _ = v7
        }
        var v8 uint32 = p0
        _ = v8
    } else {
        p0 = (p0 >> ((p1) & 63))
        switch -204 {
        case 4294967295:
            var v9 int8 = -1
            _ = v9
        }
        var v10 int64 = int64(p0)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -657
}
func genFn1(p0 int16, p1 uint8, p2 int64) string {
    v11 := []int{32768, 128, 251}
    switch -9223372036854775808 {
    case -439:
        fmt.Println(p0, p1, p0)
    case 32767:
        p0 = p0
    default:
        v12 := GenS1{F0: ("\n\t" + "a"), F1: 0, F2: 340}
        _ = v12
    }
    v13 := GenS1{F0: ("" + ("a" + "a")), F1: 32767, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v13
    return ("foo bar" + "")
}

func main() {
    var v1 uint8 = 76
    var v2 uint32 = 756
    var v3 uint = 1
    for fi14 := 0; fi14 < 8; fi14++ {
        switch (fi14 % (((fi14 % ((42) | 1))) | 1)) {
        case -247:
            v15 := GenS1{F0: "hello", F1: 2, F2: v3}
            _ = v15
        default:
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 uint64 = uint64(v1)
                _ = fi16
                _ = v17
            }
        }
        if (!(2 <= 212)) {
            var v18 int = fi14
            fmt.Println(v2, v3, fi14)
            _ = v18
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := []int{}
        for fi21 := 0; fi21 < 15; fi21++ {
            fmt.Println(v1)
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    v22 := GenS1{F0: "hello", F1: 65535, F2: v3}
    var v23 float32 = ((1e10 - (0.0 * (1e10 - 1e10))) / (0.1 * (1.5 - -2.25)))
    var v24 uint16 = 358
    var v25 bool = (!true)
    var v26 int8 = -1
    var v27 int16 = 585
    var v28 int8 = (v26 >> ((v3) & 63))
    for fi29 := 0; fi29 < 5; fi29++ {
        v23 = 3.14159
        v24 += v24
        var v30 float64 = float64(float32(float64((1.5 / 3.14159))))
        _ = fi29
        _ = v30
    }
    v31 := GenS2{F0: v2, F1: 4294967295}
    fmt.Println(v24, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
}
