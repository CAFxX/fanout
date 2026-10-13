package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int64; F1 int; F2 bool }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0(p0 int32, p1 int16, p2 string) int16 {
    if (-2.25 < (1e10 * (0.0 - -2.25))) {
        v7 := GenS1{F0: p1, F1: 42}
        for fi8 := 0; fi8 < 17; fi8++ {
            p2 = p2
            if true {
                if true {
                    v9 := GenS1{F0: (p1 % ((32767) | 1)), F1: fi8}
                    p0 = 42
                    _ = v9
                } else {
                    var v10 bool = true
                    _ = v10
                }
                p0 -= (^(p0))
            }
            _ = fi8
        }
        _ = v7
    } else {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint8 = 2
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 float32 = 3.14159
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    var v15 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 351
}

func main() {
    var v1 float32 = ((0.1 / (1e10 + 3.14159)) * 0.0)
    v2 := make(map[string]int)
    var v3 uint = 2
    v4 := make(map[string]int)
    v5 := make(map[uint32]string)
    var v6 float64 = 0.0
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 uint = (v3 + v3)
        var v18 float32 = v1
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS2{F0: -1, F1: fi19, F2: false}
        v21 := GenS1{F0: -446, F1: 127}
        if false {
            switch fi19 {
            case 32768:
                fmt.Println(v6)
            case -303:
                switch 127 {
                case -29:
                    if (!((fi19 ^ (fi19 >> ((v3) & 63))) >= (fi19 * (fi19 % ((fi19) | 1))))) {
                        delete(v4, "")
                        var v22 int32 = -593
                        _ = v22
                    }
                case -9223372036854775808:
                    v23 := GenS1{F0: -32768, F1: -606}
                    _ = v23
                case 42:
                    var v24 int = (fi19 % (((fi19 % (((fi19 << (((v3 << (((v3 >> ((18446744073709551615) & 63))) & 63))) & 63))) | 1))) | 1))
                    _ = v24
                default:
                    var v25 uint64 = 18446744073709551615
                    _ = v25
                }
            }
            if (!((-1 >= 32767) || true)) {
                v6 = v6
                v26 := GenS2{F0: 1, F1: 668, F2: true}
                _ = v26
            } else {
                delete(v4, "a")
            }
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    v3 -= (^(v3))
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 uint8 = 0
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        if false {
            fmt.Println(fi29, v3)
            v30 := GenS3{F0: 4294967295, F1: (!false)}
            v6 = float64((v1 - 1e10))
            _ = v30
        }
        if false {
            delete(v5, 0)
            v3 = (v3 | v3)
        } else {
            v3 -= v3
        }
        _ = fi29
    }
    var v31 string = ("foo bar" + "foo bar")
    v3 += (v3 + 441)
    v1 = float32(((v6 * float64(v1)) - -2.25))
    var v32 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
