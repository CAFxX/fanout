package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int64, p1 int8, p2 uint32) uint64 {
    v6 := []int{-223, 4294967295}
    if (42 == 128) {
        var v7 float64 = (0.1 * ((3.14159 + ((1e10 / -2.25) - (3.14159 / (3.14159 / 1e10)))) / 1.0))
        if false {
            var v8 uint = 78
            _ = v8
        } else {
            var v9 uint8 = 2
            if ((p2 / ((42) | 1)) >= (p2 % ((2147483648) | 1))) {
                p1 += p1
            }
            _ = v9
        }
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int16 = -1
            if (true && (2147483647 < fi10)) {
                var v12 uint8 = 2
                _ = v12
            } else {
                var v13 float64 = (((float64(0.0) * v7) * (v7 - (-2.25 / (-2.25 / 1e10)))) * (((1.5 / (v7 * float64(1e10))) - v7) * 1.5))
                _ = v13
            }
            var v14 uint64 = 310
            _ = fi10
            _ = v11
            _ = v14
        }
        _ = v7
    } else {
        var v15 int16 = 626
        _ = v15
    }
    p1 -= int8(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 899
}

func main() {
    var v1 int = 341
    v2 := make(map[string]uint32)
    var v3 float64 = ((0.1 - ((float64(3.14159) / 1e10) + 3.14159)) / float64(-2.25))
    v4 := make(map[uint32]string)
    var v5 int8 = -128
    v1 += (v1 / ((127) | 1))
    fmt.Println(v5, v3, v3)
    v3 += ((float64(1.5) + 1e10) / v3)
    v16 := []int{v1, (v1 + v1), -1, v1}
    v1 += (v1 % ((v1) | 1))
    if false {
        v5 -= v5
        var v17 int = (^(v1))
        v1 -= (-(v17))
        _ = v17
    } else {
        var v18 int64 = 42
        _ = v18
    }
    switch "" {
    case "a":
        if len(v16) > 0 {
            v16[0] = v1
        }
    case "":
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    default:
        v1 += v1
    }
    v1 = (^(v1))
    v19 := GenS1{F0: 850}
    if (!true) {
        for fi20 := 0; fi20 < 2; fi20++ {
            if (false && (true && true)) {
                if false {
                    if false {
                        v3 += -2.25
                        _ = len(v16)
                    }
                    v1 -= (fi20 << ((2) & 63))
                    v21 := GenS1{F0: 4294967295}
                    _ = v21
                }
            }
            if false {
                v4[65535] = "foo bar"
                fmt.Println(v5, fi20)
            }
            fmt.Println(v3, v3, fi20)
            _ = fi20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
