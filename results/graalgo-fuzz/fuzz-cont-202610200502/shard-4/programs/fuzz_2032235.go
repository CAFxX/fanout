package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 string) string {
    p0 = p0
    p0 = "foo bar"
    v7 := GenS1{F0: 183}
    _ = p0
    _ = v7
    return (p0 + (("a" + p0) + "a"))
}

func main() {
    var v1 uint8 = 225
    v2 := []string{("" + (((((("" + "") + "foo bar") + "\n\t") + "a") + ("foo bar" + "foo bar")) + "\n\t")), (((("hello" + ("hello" + "\n\t")) + "") + (((((("hello" + "") + "") + "\n\t") + ("foo bar" + "\n\t")) + "\n\t") + ("foo bar" + "\n\t"))) + "a")}
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2
    v5 := make(map[int]bool)
    var v6 int = 446
    v4 = (v4 & v4)
    if ((v4 - v4) >= uint32(v1)) {
        var v8 int32 = 1
        fmt.Println(v8, v4)
        _ = v8
    }
    var v9 uint = 2
    _ = len(v3)
    var v10 float32 = float32(3.14159)
    var v11 uint8 = 2
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint64 = uint64(v4)
        fmt.Println(v6, v10)
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 float64 = (float64((float32(float64((3.14159 + 1.5))) + 0.1)) / (-2.25 / 1e10))
        delete(v5, 1)
        if false {
            fmt.Println(v15, v15, v9)
            v16 := GenS2{F0: -32768, F1: 0, F2: (v4 & (v4 ^ 414))}
            v11 += v1
            _ = v16
        } else {
            var v17 int32 = -1
            v18 := GenS3{F0: 935, F1: 260}
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    var v19 float64 = 0.0
    v20 := v3[uint32(v9)]
    if v20 {
        v20 = (222 >= 465)
        var v21 uint32 = 267
        _ = v21
    }
    if (false || (v20 || v20)) {
        var v22 uint = v9
        _ = v22
    } else {
        var v23 float64 = (1.5 * (0.0 + v19))
        var v24 float64 = -2.25
        if v20 {
            _ = len(v2)
            fmt.Println(v10)
            v25 := GenS3{F0: v6, F1: 671}
            _ = v25
        }
        _ = v23
        _ = v24
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        v1 = (v1 & (^(v1)))
        _ = fi26
    }
    v20 = ((-130 == -1) && ("" != ((("" + "a") + "a") + "a")))
    if (696 == 42) {
        var v27 int16 = 0
        v6 -= (v6 << (((v9 - (v9 | (v9 / (((v9 - uint(v1))) | 1))))) & 63))
        _ = v27
    } else {
        for fi28 := 0; fi28 < 20; fi28++ {
            for fi29 := 0; fi29 < 19; fi29++ {
                v4 -= v4
                var v30 int32 = 701
                _ = fi29
                _ = v30
            }
            v6 -= v6
            if ((v19 - ((v19 + v19) / (v19 - -2.25))) != (v19 / v19)) {
                if ((0 != 0) || (float32((0.0 - v19)) < ((1e10 + v10) - 0.0))) {
                    for fi31 := 0; fi31 < 3; fi31++ {
                        switch v6 {
                        case 4294967295:
                            v10 -= v10
                        case 2147483647:
                            var v32 bool = (v4 >= (v4 << (((v9 + (^(v9)))) & 63)))
                            _ = v32
                        case -487:
                            v33 := GenS3{F0: 733, F1: 42}
                            _ = v33
                        }
                        v34 := GenS1{F0: -787}
                        _ = fi31
                        _ = v34
                    }
                }
            } else {
                fmt.Println(v9)
                v35 := GenS1{F0: 0}
                _ = v35
            }
            _ = fi28
        }
        var v36 int8 = 124
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
