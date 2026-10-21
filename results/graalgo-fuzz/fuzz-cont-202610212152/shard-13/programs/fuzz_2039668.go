package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int8; F2 uint64 }
type GenS3 struct { F0 int; F1 uint16 }
func genFn0(p0 int) int16 {
    fmt.Println(p0)
    if (!true) {
        if false {
            var v7 uint16 = 1
            var v8 bool = false
            v9 := GenS1{F0: "foo bar"}
            _ = v7
            _ = v8
            _ = v9
        } else {
            p0 += (p0 + p0)
            if (!false) {
                if true {
                    p0 = p0
                } else {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        p0 -= (fi10 & fi10)
                        fmt.Println(fi10, fi10)
                        var v11 int32 = -897
                        _ = fi10
                        _ = v11
                    }
                }
                var v12 uint32 = 753
                if false {
                    switch (((((("" + ("foo bar" + "foo bar")) + "hello") + "") + "hello") + "foo bar") + ((("" + "foo bar") + "hello") + "")) {
                    case "hello":
                        if true {
                            v13 := GenS2{F0: p0, F1: -128, F2: 615}
                            for fi14 := 0; fi14 < 18; fi14++ {
                                var v15 int = 42
                                _ = fi14
                                _ = v15
                            }
                            _ = v13
                        }
                    default:
                        v16 := []int{(p0 >> ((2) & 63)), 0, p0}
                        _ = v16
                    }
                }
                _ = v12
            }
            p0 = (p0 & p0)
        }
    }
    switch (p0 | p0) {
    case 9223372036854775807:
        v17 := GenS3{F0: p0, F1: 0}
        _ = v17
    case 2147483647:
        v18 := GenS3{F0: (p0 & (-(p0))), F1: 255}
        _ = v18
    }
    switch (p0 - (p0 - p0)) {
    case -9223372036854775808:
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := []int{(fi19 * p0), (-(fi19))}
            for fi21 := 0; fi21 < 11; fi21++ {
                if len(v20) > 0 {
                    v20[0] = (-(fi19))
                }
                p0 -= fi21
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
    default:
        for fi22 := 0; fi22 < 4; fi22++ {
            p0 += (fi22 | p0)
            _ = fi22
        }
    }
    p0 = (^(p0))
    _ = p0
    return 250
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = (3.14159 - (float64((0.0 + ((1.5 * ((float32(((float64((3.14159 + float32(0.1))) + ((0.1 * 1e10) * -2.25)) + (0.0 + (1.5 - float64(1e10))))) / 1.5) + ((0.1 / 1.0) / 0.1))) - (0.1 * float32(0.0))))) + float64(((0.1 * float32((3.14159 + (3.14159 / 3.14159)))) / (1e10 * -2.25)))))
    var v3 bool = false
    v4 := []uint32{441, 0, 307}
    var v5 uint32 = 679
    v6 := make(map[string]string)
    if false {
        if v3 {
            v5 = (v5 - (v5 % ((v5) | 1)))
        } else {
            v5 -= 4294967295
        }
    }
    var v23 int = 32767
    v3 = (65535 < 214)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v24 int16 = 1000
    if true {
        fmt.Println(v24, v24, v23)
    }
    switch v23 {
    case 9223372036854775807:
        fmt.Println(v3, v3)
    default:
        v5 += (v5 * 65535)
    }
    var v25 bool = false
    v26 := v6["a"]
    v2 = 1.5
    switch 9223372036854775807 {
    case -215:
        v25 = (0 == 255)
    }
    v5 = (v5 >> ((251) & 63))
    fmt.Println(v5)
    var v27 string = (v26 + ("hello" + "hello"))
    v6[v26] = ("" + v26)
    v26 = v26
    v23 = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
