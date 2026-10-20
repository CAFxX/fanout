package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint; F1 int32; F2 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 uint8 }
func genFn0(p0 bool, p1 bool, p2 bool) int8 {
    var v5 uint64 = 1
    v6 := GenS1{F0: 537, F1: 0, F2: -858}
    v7 := GenS2{F0: 0, F1: false}
    var v8 float64 = (3.14159 / 1e10)
    switch 128 {
    case -705:
        if p2 {
            p0 = (!(0 < 194))
        }
    default:
        if (-2.25 == 1e10) {
            if p0 {
                v9 := []int{}
                for fi10 := 0; fi10 < 6; fi10++ {
                    for fi11 := 0; fi11 < 13; fi11++ {
                        p0 = p2
                        _ = len(v9)
                        v12 := GenS1{F0: 354, F1: -1, F2: -1}
                        _ = fi11
                        _ = v12
                    }
                    p2 = p1
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v9 = append(v9, (fi10 / (((fi13 / ((fi10) | 1))) | 1)))
                        p0 = p1
                        p0 = ((float64(1.5) - 3.14159) <= v8)
                        _ = fi13
                    }
                    _ = fi10
                }
                _ = v9
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 100
}

func main() {
    v1 := []uint32{539, 2}
    v2 := make(map[string]string)
    var v3 bool = (!false)
    var v4 uint16 = 0
    for fi14 := 0; fi14 < 8; fi14++ {
        _ = len(v1)
        _ = fi14
    }
    var v15 int8 = 0
    for fi16 := 0; fi16 < 18; fi16++ {
        if len(v1) > 0 {
            v1[0] = uint32(v4)
        }
        v4 = (v4 ^ v4)
        _ = fi16
    }
    if ((v4 > v4) || v3) {
        v3 = v3
        var v17 int = 2147483647
        _ = v17
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (true || (v3 && (genFn0((0 > 0), true, false) > genFn0(v3, (!v3), (v3 && false)))))
    var v18 int8 = v15
    if v3 {
        if v3 {
            var v19 float32 = float32(1.5)
            v4 = (^(v4))
            _ = v19
        } else {
            v20 := GenS1{F0: 42, F1: genFn0(v3, true, false), F2: 766}
            _ = v20
        }
    } else {
        for fi21 := 0; fi21 < 19; fi21++ {
            fmt.Println(v3)
            var v22 uint16 = 276
            _ = fi21
            _ = v22
        }
        var v23 int32 = 2147483647
        v24 := GenS2{F0: 127, F1: v3}
        _ = v23
        _ = v24
    }
    fmt.Println(v4, v18, v15)
    var v25 bool = (!v3)
    if v25 {
        v26 := GenS3{F0: uint(v4), F1: 0, F2: uint32(v4)}
        var v27 int16 = 1
        switch 32768 {
        case -618:
            if v3 {
                fmt.Println(v4, v18, v18)
                if v3 {
                    v4 += v4
                    v15 -= (v18 + v18)
                }
            } else {
                var v28 uint64 = 119
                if true {
                    fmt.Println(v27, v3)
                    switch (("hello" + "a") + "hello") {
                    case "hello":
                        var v29 float64 = 0.0
                        _ = v29
                    default:
                        var v30 int16 = v27
                        _ = v30
                    }
                    if true {
                        v25 = v25
                        v31 := GenS4{F0: 945}
                        var v32 int16 = v27
                        _ = v31
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 20; fi33++ {
                            var v34 int64 = 660
                            _ = fi33
                            _ = v34
                        }
                        v18 -= 58
                    }
                }
                for fi35 := 0; fi35 < 14; fi35++ {
                    if (((((0.1 * -2.25) / ((((0.1 + 0.1) - (3.14159 - 0.1)) + (0.1 - float64(1.5))) * ((1.5 / -2.25) - 0.0))) * 1.5) == 3.14159) && v25) {
                        var v36 uint8 = 131
                        _ = v36
                    }
                    _ = fi35
                }
                _ = v28
            }
        case 128:
            v37 := v2["foo bar"]
            _ = v37
        case 42:
            if (!v25) {
                v18 -= (v18 / (((v15 * (v15 % ((genFn0(((((("hello" + "\n\t") + "") + "hello") < ("hello" + ("foo bar" + "hello"))) && v3), v3, ((0.1 + float32((float64(0.1) - 0.1))) >= ((((((0.1 * -2.25) - (float32(0.0) * ((float32(0.1) - (-2.25 * 0.0)) - ((-2.25 - (-2.25 + (0.1 / 0.1))) + 1.5)))) + (0.1 * 0.0)) / (1.5 / (0.0 + -2.25))) / 1e10) / 1.0)))) | 1)))) | 1))
            } else {
                if (22 > 65535) {
                    v15 -= genFn0(v3, v25, true)
                    var v38 int8 = (^(v15))
                    var v39 float32 = 3.14159
                    _ = v38
                    _ = v39
                }
                var v40 bool = (!v25)
                fmt.Println(v27)
                _ = v40
            }
        }
        _ = v26
        _ = v27
    }
    fmt.Println(v3, v25, v3)
    for fi41 := 0; fi41 < 17; fi41++ {
        v25 = ("a" < ("foo bar" + "\n\t"))
        _ = fi41
    }
    v42 := GenS5{F0: (v18 * v18), F1: 255}
    for fi43 := 0; fi43 < 11; fi43++ {
        var v44 uint16 = v4
        v44 = v44
        _ = fi43
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v42
}
