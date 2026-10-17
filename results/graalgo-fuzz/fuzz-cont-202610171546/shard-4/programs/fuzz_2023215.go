package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -1
    v3 := make(map[uint32]uint32)
    v4 := []int{-165, 127, 548, 127}
    v2 = v2
    v2 -= (-(v2))
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint = 581
        _ = fi5
        _ = v6
    }
    fmt.Println(v2)
    v7 := GenS1{F0: 2}
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            if (!(true && false)) {
                v4 = append(v4, (fi8 >> ((2) & 63)))
                if false {
                    switch (fi9 & (-(fi9))) {
                    case -656:
                        v2 = -791
                    }
                    fmt.Println(v2, fi8, fi9)
                    var v10 uint32 = 1
                    if false {
                        var v11 float64 = (float64(-2.25) + ((-2.25 - float64(float32(-2.25))) / (0.0 - 1.5)))
                        v12 := GenS1{F0: 589}
                        _ = v11
                        _ = v12
                    } else {
                        v2 += 1
                    }
                    fmt.Println(v10, fi9, v2)
                    var v13 uint = 65535
                    _ = v10
                    _ = v13
                }
            } else {
                switch fi8 {
                case -9223372036854775808:
                    v14 := GenS1{F0: 4294967295}
                    _ = v14
                default:
                    v2 = (v2 ^ -1)
                }
            }
            fmt.Println(fi9)
            _ = fi9
        }
        _ = fi8
    }
    switch 32768 {
    case 32767:
        v4 = append(v4, 973)
    }
    v2 += (v2 - v2)
    v15 := GenS1{F0: 1}
    switch ("\n\t" + ((((("foo bar" + "hello") + "hello") + "") + "foo bar") + "")) {
    case "foo bar":
        fmt.Println(v2)
    }
    var v16 uint16 = 255
    v17 := GenS2{F0: 1, F1: 208, F2: -378}
    fmt.Println(v2, v2)
    v16 += v16
    var v18 uint32 = 1
    for fi19 := 0; fi19 < 10; fi19++ {
        v2 -= v2
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
