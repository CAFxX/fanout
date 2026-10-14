package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool }

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]bool)
    var v3 uint32 = (v1 / ((630) | 1))
    var v4 uint32 = (v1 & v1)
    var v5 int32 = 81
    var v6 string = "\n\t"
    v6 = v6
    switch v6 {
    case "hello":
        var v7 uint16 = 479
        _ = v7
    case "":
        v8 := GenS1{F0: 42, F1: 32767, F2: 255}
        _ = v8
    case "foo bar":
        v9 := v2[42]
        _ = v9
    }
    if (false || (true && (!((1.5 + (0.1 / 3.14159)) == (0.1 * 0.0))))) {
        v10 := GenS1{F0: 950, F1: 1, F2: 255}
        var v11 float32 = -2.25
        _ = v10
        _ = v11
    } else {
        v3 = 928
        v12 := GenS2{F0: -346, F1: 917}
        switch v6 {
        case "":
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS2{F0: v5, F1: v1}
                _ = fi13
                _ = v14
            }
        default:
            var v15 bool = (0 != uint16(v4))
            _ = v15
        }
        _ = v12
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            v4 -= v1
            var v18 uint16 = 42
            v5 = 42
            _ = fi17
            _ = v18
        }
        var v19 int16 = 32767
        v19 += v19
        _ = fi16
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        if (!(0 == 18446744073709551615)) {
            for fi21 := 0; fi21 < 13; fi21++ {
                v3 -= 418
                _ = fi21
            }
        } else {
            for fi22 := 0; fi22 < 2; fi22++ {
                fmt.Println(v6, v6, v4)
                v4 -= (v1 * v1)
                v2[v3] = (!false)
                _ = fi22
            }
            v23 := GenS3{F0: 128, F1: 1}
            var v24 int32 = 1
            _ = v23
            _ = v24
        }
        v25 := GenS2{F0: (v5 + (^(v5))), F1: (v4 | 67)}
        _ = fi20
        _ = v25
    }
    if (!true) {
        v3 = (^(v4))
        if true {
            var v26 int = -1
            v26 = v26
            if true {
                v27 := []int{(^(v26))}
                _ = v27
            } else {
                v28 := []int{v26}
                switch (("a" + v6) + "\n\t") {
                case "a":
                    for fi29 := 0; fi29 < 6; fi29++ {
                        var v30 bool = false
                        fmt.Println(fi29, v3, v5)
                        _ = fi29
                        _ = v30
                    }
                case "foo bar":
                    var v31 int8 = 42
                    _ = v31
                case "":
                    v32 := GenS4{F0: (0 < 1)}
                    _ = v32
                default:
                    _ = len(v2)
                }
                _ = v28
            }
            _ = v26
        } else {
            if false {
                v6 = v6
            } else {
                var v33 float32 = 3.14159
                var v34 uint8 = 183
                _ = v33
                _ = v34
            }
        }
    } else {
        fmt.Println(v4, v6)
        var v35 uint16 = 804
        _ = v35
    }
    v36 := GenS3{F0: 11, F1: 199}
    v37 := GenS5{F0: 118}
    var v38 int = -9223372036854775808
    var v39 float32 = 0.0
    v40 := GenS6{F0: true}
    var v41 float64 = (float64(1e10) - 0.0)
    fmt.Println(v39, v3, v38)
    var v42 int32 = -430
    var v43 uint32 = (v4 % ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
