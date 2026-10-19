package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint64, p1 uint8, p2 int) uint32 {
    v7 := GenS1{F0: p2}
    var v8 bool = (false && (0 == -18))
    fmt.Println(v8, p2, p0)
    if (0.0 != (((0.1 / ((3.14159 + (1.5 / 1.0)) / (float32((3.14159 / float64(3.14159))) * 0.1))) - float32(3.14159)) + (1.5 + (1e10 * (0.1 - (-2.25 * 1e10)))))) {
        if (!v8) {
            fmt.Println(p2)
            if true {
                switch 392 {
                case 2147483647:
                    if (!(("foo bar" <= "") || true)) {
                        if (!v8) {
                            v8 = false
                            if v8 {
                                fmt.Println(p0, p1, p0)
                            }
                        }
                        v9 := GenS1{F0: p2}
                        _ = v9
                    }
                case -1:
                    p2 = (p2 - (p2 * (^(p2))))
                case -2147483648:
                    for fi10 := 0; fi10 < 6; fi10++ {
                        v8 = (987 >= 1)
                        var v11 uint64 = p0
                        _ = fi10
                        _ = v11
                    }
                }
            } else {
                if true {
                    fmt.Println(p1, p0, v8)
                    var v12 uint16 = 917
                    _ = v12
                }
                fmt.Println(v8)
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 string, p1 string) int32 {
    var v13 uint16 = 65535
    var v14 int64 = 739
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 32767
    v3 := make(map[string]int)
    var v4 uint = 2
    var v5 float32 = 1e10
    v6 := []uint32{}
    if len(v6) > 0 {
        v6[0] = genFn0(1, 255, 505)
    }
    var v15 uint16 = 2
    v4 = (v4 ^ (v4 % ((2) | 1)))
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v4, v5, v4)
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 uint32 = 42
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    var v19 int8 = 1
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    v19 += 100
    if true {
        switch 127 {
        case 32768:
            v20 := v1[-1]
            _ = v20
        default:
            switch 42 {
            case -1:
                var v21 int64 = 0
                _ = v21
            }
        }
        v22 := GenS1{F0: -2147483648}
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
