package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32, p1 uint8, p2 uint16) int8 {
    fmt.Println(p2)
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 bool = (!true)
        p2 = 796
        _ = fi5
        _ = v6
    }
    if false {
        fmt.Println(p1)
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := []int{(^(fi8)), 128}
                fmt.Println(p1)
                _ = fi8
                _ = v9
            }
            var v10 string = "a"
            _ = fi7
            _ = v10
        }
    } else {
        if false {
            switch -1 {
            case -2147483648:
                if true {
                    p0 += (p0 << ((65535) & 63))
                } else {
                    switch ((("" + "hello") + "hello") + ("\n\t" + "")) {
                    case "hello":
                        var v11 int64 = 42
                        _ = v11
                    case "":
                        switch "" {
                        case "":
                            p2 += (p2 - p2)
                        case "hello":
                            var v12 int8 = -107
                            _ = v12
                        case "a":
                            v13 := GenS1{F0: (p2 & (^(p2))), F1: 1}
                            _ = v13
                        }
                    case "\n\t":
                        p1 += p1
                    default:
                        for fi14 := 0; fi14 < 7; fi14++ {
                            var v15 int16 = 1000
                            v16 := GenS2{F0: 42, F1: (p1 ^ (p1 / (((^(p1))) | 1))), F2: 2147483647}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                    p2 -= (p2 * (p2 << ((862) & 63)))
                }
            }
            var v17 uint32 = 905
            _ = v17
        } else {
            var v18 int8 = 100
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 100
}

func main() {
    v1 := []uint32{486}
    v2 := make(map[int]string)
    var v3 uint = 2
    var v4 int64 = 1
    v19 := GenS2{F0: 809, F1: 56, F2: 949}
    _ = len(v2)
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v20 float64 = 1e10
    v20 = ((1e10 * v20) + ((v20 - ((v20 + ((v20 + v20) + v20)) * (v20 - -2.25))) + (0.1 * (float64(1.5) + v20))))
    var v21 float64 = float64(float32(v20))
    for fi22 := 0; fi22 < 3; fi22++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 -= 180
        v23 := GenS2{F0: 565, F1: 127, F2: 42}
        _ = fi22
        _ = v23
    }
    if (2147483647 <= 2147483648) {
        var v24 float64 = float64((0.1 - 1.5))
        var v25 string = "foo bar"
        delete(v2, 42)
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        v27 := GenS3{F0: false}
        v4 = (^(v4))
        _ = fi26
        _ = v27
    }
    if false {
        v28 := GenS1{F0: 255, F1: (v3 - v3)}
        v29 := GenS1{F0: 0, F1: v3}
        _ = v28
        _ = v29
    } else {
        var v30 string = "a"
        _ = v30
    }
    v31 := GenS2{F0: 2147483647, F1: 42, F2: -996}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v4
}
