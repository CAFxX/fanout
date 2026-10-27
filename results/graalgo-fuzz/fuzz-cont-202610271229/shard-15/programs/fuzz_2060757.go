package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint16; F2 int16 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS5 struct { F0 bool; F1 bool; F2 int16 }

func main() {
    v1 := []string{("hello" + "\n\t"), ("" + "a"), "a"}
    var v2 int64 = 645
    var v3 string = (("hello" + "") + (("\n\t" + "\n\t") + ""))
    var v4 bool = (false || true)
    var v5 int8 = 119
    fmt.Println(v4, v3, v5)
    var v6 int8 = (-(v5))
    fmt.Println(v5, v2, v3)
    v2 -= (^(v2))
    _ = len(v1)
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: (v5 * -117)}
        var v9 uint = 102
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = (float64(0.0) * -2.25)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 uint64 = 2
    _ = len(v1)
    if v4 {
        v12 := GenS2{F0: 2, F1: 255, F2: 1000}
        var v13 int = 32768
        _ = v12
        _ = v13
    } else {
        if len(v1) > 0 {
            v1[0] = ("foo bar" + ((v3 + (("a" + ("foo bar" + (v3 + (v3 + (v3 + v3))))) + v3)) + "a"))
        }
        if v4 {
            v4 = (!v4)
            var v14 int32 = -806
            v2 -= 9223372036854775807
            _ = v14
        } else {
            v15 := GenS2{F0: 311, F1: 598, F2: 1000}
            var v16 uint = 255
            switch (((v3 + v3) + "") + "\n\t") {
            case "\n\t":
                for fi17 := 0; fi17 < 17; fi17++ {
                    fmt.Println(v11, v4)
                    _ = fi17
                }
            case "a":
                v2 += -500
            default:
                if v4 {
                    v6 = (v5 + (^(v5)))
                    _ = len(v1)
                    v18 := GenS3{F0: 0, F1: -270}
                    _ = v18
                } else {
                    v19 := GenS3{F0: -775, F1: -1}
                    v20 := GenS4{F0: -1, F1: uint8(v16), F2: 606}
                    var v21 int32 = -1
                    _ = v19
                    _ = v20
                    _ = v21
                }
            }
            _ = v15
            _ = v16
        }
        if v4 {
            switch v3 {
            case "foo bar":
                if ((((0.0 * (-2.25 / 3.14159)) + (-2.25 / float32(float64(float32(3.14159))))) / 1e10) < -2.25) {
                    if v4 {
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v10 += v10
                            var v23 int8 = (v5 / ((127) | 1))
                            _ = fi22
                            _ = v23
                        }
                        v2 = (v2 % ((v2) | 1))
                    } else {
                        var v24 float64 = (v10 + v10)
                        var v25 uint32 = 2
                        _ = v24
                        _ = v25
                    }
                    v26 := GenS5{F0: v4, F1: ((true && (2147483647 != -1)) && true), F2: 620}
                    if ((v6 / ((v5) | 1)) > v6) {
                        var v27 float64 = (v10 * v10)
                        var v28 uint64 = v11
                        _ = v27
                        _ = v28
                    }
                    _ = v26
                } else {
                    if (!false) {
                        var v29 uint64 = 137
                        v30 := GenS3{F0: 79, F1: -1}
                        _ = v29
                        _ = v30
                    }
                }
            }
            v3 = "\n\t"
        } else {
            for fi31 := 0; fi31 < 18; fi31++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v11 = v11
                v4 = (245 <= 42)
                _ = fi31
            }
        }
    }
    v4 = (false || ((v11 ^ 2) != (v11 << ((742) & 63))))
    switch ((("a" + v3) + v3) + "\n\t") {
    case "foo bar":
        _ = len(v1)
    case "":
        fmt.Println(v11)
    case "\n\t":
        v32 := GenS5{F0: true, F1: true, F2: 1000}
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
