package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint16 = 659
    var v2 int32 = -94
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    v2 -= v2
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int8 = 86
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: 42}
    v4[("hello" + "")] = false
    v2 += (v2 << ((259) & 63))
    v8 := GenS1{F0: -1, F1: 9223372036854775807}
    delete(v3, ("a" + ((("foo bar" + "\n\t") + "hello") + "")))
    v1 = v1
    v9 := v4[""]
    if v9 {
        if ((!true) || (638 > 0)) {
            v10 := []int{32768, -980}
            _ = v10
        } else {
            fmt.Println(v1, v9, v9)
        }
        switch 0 {
        case -488:
            var v11 float64 = 0.0
            _ = v11
        case 4294967295:
            for fi12 := 0; fi12 < 11; fi12++ {
                v13 := GenS2{F0: v2}
                if v9 {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        switch (fi14 ^ 1) {
                        case -2147483648:
                            if v9 {
                                var v15 int64 = 42
                                _ = v15
                            }
                        case 32767:
                            v3["foo bar"] = 2
                        }
                        _ = fi14
                    }
                    fmt.Println(v9)
                }
                _ = fi12
                _ = v13
            }
        case -277:
            var v16 int32 = v2
            _ = v16
        default:
            v9 = (0 >= 868)
        }
        if v9 {
            v4["hello"] = v9
        } else {
            var v17 float64 = (((-2.25 / float64(float32((float64(1.5) - 0.1)))) - (float64(((1.5 - -2.25) / 3.14159)) / 1e10)) * 3.14159)
            v1 = (v1 + 456)
            _ = v17
        }
    } else {
        v1 -= 65535
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
