package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 string }
func genFn0(p0 int16, p1 int16, p2 int8) bool {
    var v5 uint16 = 0
    var v6 int8 = (p2 & (p2 >> ((1) & 63)))
    v6 = p2
    p1 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 uint, p1 bool) uint32 {
    v7 := []int{-278, 32768}
    switch "foo bar" {
    case "hello":
        v8 := GenS1{F0: 2147483647}
        _ = v8
    case "a":
        if p1 {
            fmt.Println(p1, p0)
        } else {
            var v9 float32 = 3.14159
            var v10 int64 = 0
            _ = v10
            _ = v9
        }
    }
    var v11 string = "hello"
    if len(v7) > 0 {
        v7[0] = 685
    }
    var v12 uint = (p0 | (p0 & (^(p0))))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -267
    var v3 uint16 = 518
    var v4 float64 = ((1.5 - float64(0.0)) - (1e10 - 1.5))
    v3 += v3
    var v13 int16 = (v2 / ((v2) | 1))
    if true {
        v4 = float64((((float32((((v4 - 0.1) / (0.0 / v4)) / v4)) / 0.1) / 1.0) * (float32((v4 / (v4 / ((float64((1.5 * float32(0.1))) * v4) + v4)))) * (0.1 / float32(v4)))))
    } else {
        v14 := GenS2{F0: (v2 * (v2 % (((v13 / ((v13) | 1))) | 1))), F1: false, F2: (("foo bar" + ("hello" + "foo bar")) + "\n\t")}
        _ = v14
    }
    switch "hello" {
    case "a":
        v2 = (v13 % (((v2 + (v13 & int16(v3)))) | 1))
    }
    v4 = 0.1
    if (60 < 2) {
        var v15 int8 = 0
        var v16 int16 = (^(v13))
        _ = v15
        _ = v16
    }
    var v17 float64 = v4
    switch 463 {
    case -2147483648:
        var v18 float32 = (((1.5 * 0.1) / (0.1 * ((0.1 * (-2.25 - 0.0)) - -2.25))) * ((1.5 - 1e10) - 3.14159))
        _ = v18
    case -704:
        v13 -= (^(v2))
    default:
        var v19 float32 = (-2.25 + (((1.5 * (float32(v17) * 0.0)) + (1e10 * (float32(float64(((3.14159 - float32(0.1)) - 1.5))) - 1e10))) - (-2.25 / 1e10)))
        _ = v19
    }
    v2 -= -1
    var v20 uint32 = 42
    fmt.Println(v2)
    for fi21 := 0; fi21 < 1; fi21++ {
        if (0 < 127) {
            if (!(!("hello" <= ("" + "foo bar")))) {
                for fi22 := 0; fi22 < 4; fi22++ {
                    var v23 uint64 = 1
                    _ = len(v1)
                    _ = fi22
                    _ = v23
                }
                switch (fi21 ^ -9223372036854775808) {
                case 0:
                    v24 := v1[(fi21 | 32767)]
                    _ = v24
                default:
                    v13 = (v13 >> ((235) & 63))
                }
                v4 += v17
            } else {
                if true {
                    _ = len(v1)
                    v25 := GenS2{F0: (^(v2)), F1: (42 > 1), F2: ""}
                    _ = v25
                }
                if false {
                    switch 195 {
                    case 1:
                        fmt.Println(v4, v3)
                    }
                    switch "hello" {
                    case "hello":
                        if ((-1 > -9223372036854775808) && false) {
                            v3 = v3
                        }
                    case "\n\t":
                        var v26 int = fi21
                        _ = v26
                    default:
                        v17 = v17
                    }
                    for fi27 := 0; fi27 < 2; fi27++ {
                        var v28 uint64 = 288
                        fmt.Println(v2, v3, v20)
                        switch fi21 {
                        case 42:
                            v20 += genFn1(1, true)
                        }
                        _ = fi27
                        _ = v28
                    }
                }
            }
            if true {
                v4 += v17
            }
        }
        _ = len(v1)
        v29 := v1[fi21]
        _ = fi21
        _ = v29
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        var v31 bool = (true && ((true && false) || false))
        var v32 float64 = 0.0
        var v33 uint8 = uint8(v3)
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
