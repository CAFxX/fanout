package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 bool; F1 bool; F2 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 uint32 }
type GenS4 struct { F0 int64; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint }
func genFn0(p0 int32) uint64 {
    v6 := GenS1{F0: ("foo bar" + "a"), F1: (127 <= 247)}
    p0 = -27
    p0 += (p0 * (p0 + p0))
    var v7 int32 = 1
    _ = p0
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint32, p1 uint32, p2 bool) int {
    var v8 int8 = -128
    var v9 uint64 = 1
    if true {
        fmt.Println(v9)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{178}
    var v2 uint16 = 1
    v3 := make(map[string]bool)
    var v4 float32 = float32(-2.25)
    var v5 float64 = 0.1
    v10 := GenS1{F0: "a", F1: true}
    fmt.Println(v2)
    if true {
        v11 := GenS1{F0: "hello", F1: (849 == v2)}
        var v12 float32 = (v4 * (v4 - (float32(v5) - 0.1)))
        v13 := GenS1{F0: "", F1: false}
        _ = v11
        _ = v12
        _ = v13
    }
    v5 += (-2.25 - ((v5 * (v5 * float64(((((float32(3.14159) + float32(1.5)) * v4) * (1.5 / v4)) * 3.14159)))) * (v5 / v5)))
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(v5)
        if ((!(!true)) || (-1 == 85)) {
            v15 := GenS2{F0: (!false), F1: false, F2: 163}
            _ = v15
        } else {
            if ((-1 <= -246) || false) {
                v4 += (float32(-2.25) + (v4 * (v4 / 3.14159)))
            } else {
                for fi16 := 0; fi16 < 1; fi16++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi16
                }
            }
            v2 = v2
            if true {
                var v17 bool = (127 == 42)
                var v18 string = "\n\t"
                v19 := GenS3{F0: 42, F1: 452, F2: 42}
                _ = v17
                _ = v18
                _ = v19
            } else {
                v5 += v5
                switch ("foo bar" + "hello") {
                case "a":
                    for fi20 := 0; fi20 < 17; fi20++ {
                        var v21 uint8 = 85
                        _ = fi20
                        _ = v21
                    }
                case "foo bar":
                    v2 -= (v2 / ((v2) | 1))
                case "\n\t":
                    if true {
                        if true {
                            var v22 uint8 = 127
                            _ = v22
                        }
                        v5 = ((v5 + 1e10) / ((1e10 - (-2.25 / 1e10)) / 1.5))
                        fmt.Println(fi14, v5)
                    }
                default:
                    var v23 int64 = 0
                    _ = v23
                }
                delete(v3, "a")
            }
        }
        v5 = (-2.25 * (-2.25 * 1.5))
        _ = fi14
    }
    v24 := GenS4{F0: 0, F1: (v4 <= (-2.25 - v4)), F2: v2}
    fmt.Println(v2)
    for fi25 := 0; fi25 < 9; fi25++ {
        v5 = 3.14159
        _ = fi25
    }
    v26 := GenS3{F0: 1, F1: 42, F2: 1}
    fmt.Println(v4, v4, v5)
    v4 += v4
    v3["foo bar"] = (42 == 9223372036854775807)
    switch 2147483647 {
    case 280:
        fmt.Println(v5, v2, v5)
    case 128:
        for fi27 := 0; fi27 < 14; fi27++ {
            switch "" {
            case "":
                if (350 != 42) {
                    if false {
                        for fi28 := 0; fi28 < 16; fi28++ {
                            var v29 int64 = 0
                            v29 += v29
                            v30 := GenS5{F0: 2147483647, F1: 65535}
                            v31 := GenS6{F0: -74}
                            v29 = (^(v29))
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        v4 -= (((v4 + -2.25) * v4) * ((v4 / 3.14159) + v4))
                        var v32 uint = 42
                        _ = v32
                    } else {
                        var v33 string = ("foo bar" + "hello")
                        v34 := v3[""]
                        v4 = (v4 / ((float32((float64(v4) / 1.5)) + -2.25) * ((v4 / float32(1e10)) / 1e10)))
                        _ = v33
                        _ = v34
                    }
                }
            case "foo bar":
                v35 := GenS7{F0: 927}
                _ = v35
            case "a":
                v5 = float64(v4)
            }
            _ = fi27
        }
    }
    v2 = 669
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
