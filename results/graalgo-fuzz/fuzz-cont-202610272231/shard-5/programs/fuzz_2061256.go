package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int8; F1 int; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 bool = false
    v4 := make(map[uint32]string)
    v5 := []uint32{2, 558}
    v6 := make(map[int]uint32)
    v3 = (-1 == 664)
    var v7 string = "foo bar"
    fmt.Println(v7)
    v8 := GenS1{F0: (v3 || true), F1: 4294967295, F2: 42}
    _ = len(v5)
    v5 = append(v5, 2147483648)
    fmt.Println(v3)
    switch v7 {
    case "foo bar":
        v9 := GenS2{F0: -1, F1: ((2147483647 <= -2147483648) && v3)}
        _ = v9
    case "hello":
        v3 = ((0.0 / -2.25) > 1.5)
    default:
        if false {
            v3 = v3
            var v10 uint64 = 554
            _ = v10
        } else {
            if v3 {
                fmt.Println(v7, v7)
                if v3 {
                    for fi11 := 0; fi11 < 5; fi11++ {
                        var v12 float32 = 3.14159
                        v13 := GenS3{F0: fi11, F1: 408}
                        fmt.Println(fi11)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v3 = (42 >= 65535)
                    if false {
                        var v14 float32 = -2.25
                        _ = v14
                    } else {
                        v15 := GenS4{F0: -37, F1: -2147483648, F2: 4294967295}
                        v7 = ("" + v7)
                        _ = v15
                    }
                }
                v16 := GenS3{F0: 418, F1: 4294967295}
                _ = v16
            } else {
                if v3 {
                    if v3 {
                        fmt.Println(v7, v7)
                        v17 := GenS2{F0: 843, F1: (!(!(110 > 127)))}
                        _ = v17
                    }
                }
            }
            delete(v4, 2147483647)
        }
    }
    if (!v3) {
        for fi18 := 0; fi18 < 18; fi18++ {
            v3 = true
            v19 := GenS2{F0: 1, F1: (395 >= 42)}
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 int8 = 81
                if (!true) {
                    var v22 uint32 = 65535
                    var v23 uint32 = v22
                    v3 = v3
                    fmt.Println(fi20, v23)
                    _ = v22
                    _ = v23
                }
                switch fi20 {
                case -862:
                    v3 = (v3 || (((1 == 59) || v3) || false))
                case -2147483648:
                    switch v7 {
                    case "":
                        _ = len(v6)
                    }
                case -346:
                    v4[707] = (v7 + v7)
                default:
                    var v24 int32 = 0
                    _ = v24
                }
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        v5 = append(v5, 1)
    }
    if (42 < -1) {
        for fi25 := 0; fi25 < 19; fi25++ {
            for fi26 := 0; fi26 < 19; fi26++ {
                if (166 == 42) {
                    var v27 bool = false
                    fmt.Println(v3, v3, v3)
                    _ = v27
                } else {
                    var v28 int = fi25
                    var v29 uint16 = 65535
                    _ = v28
                    _ = v29
                }
                fmt.Println(v3, v3)
                var v30 float64 = 0.1
                _ = fi26
                _ = v30
            }
            var v31 int8 = 100
            _ = fi25
            _ = v31
        }
    } else {
        var v32 bool = v3
        switch 42 {
        case 1:
            var v33 int32 = 1
            _ = v33
        default:
            v34 := GenS4{F0: 0, F1: 9223372036854775807, F2: 2147483648}
            _ = v34
        }
        _ = v32
    }
    var v35 float64 = ((1.5 + float64(float32((1.5 / (3.14159 + ((float64(3.14159) - float64((((1.5 + ((-2.25 + float32(((float64(1.5) * float64(float32(((3.14159 + (3.14159 * -2.25)) / 1.0)))) / 3.14159))) + -2.25)) / ((float32((0.1 - 3.14159)) + 1e10) / (-2.25 + (-2.25 / (3.14159 + -2.25))))) - float32(0.1)))) + (1e10 / float64(3.14159)))))))) * 3.14159)
    fmt.Println(v7, v3, v7)
    v5 = append(v5, 42)
    var v36 int = -9223372036854775808
    v3 = (float32(v35) < 0.1)
    var v37 float64 = v35
    var v38 uint = 42
    var v39 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
