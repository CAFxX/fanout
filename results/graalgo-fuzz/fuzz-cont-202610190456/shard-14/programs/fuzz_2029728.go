package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16; F1 int8; F2 int8 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint8 }
func genFn0(p0 uint, p1 int8, p2 uint16) string {
    p1 -= 42
    var v5 uint16 = (p2 * 65535)
    if false {
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := []int{}
            var v8 int32 = 2147483647
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ("\n\t" + "foo bar")
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]int)
    var v3 int = 32767
    v4 := make(map[uint32]uint32)
    var v9 float32 = 0.0
    v9 += 1e10
    fmt.Println(v3, v1)
    v10 := GenS1{F0: true, F1: ""}
    for fi11 := 0; fi11 < 18; fi11++ {
        v3 = (v3 & -436)
        switch v1 {
        case "hello":
            var v12 int8 = -128
            _ = v12
        default:
            var v13 uint16 = 0
            _ = v13
        }
        v9 -= float32(0.1)
        _ = fi11
    }
    v3 -= -377
    v9 = ((((((float32((float64(v9) / 3.14159)) + ((v9 - v9) + 1.5)) + v9) + v9) + v9) - v9) / -2.25)
    v9 -= v9
    v14 := []int{325, -2147483648}
    switch 42 {
    case -310:
        v15 := GenS1{F0: (v3 <= 0), F1: v1}
        _ = v15
    case -2147483648:
        var v16 uint16 = 2
        _ = v16
    case 806:
        switch ("foo bar" + v1) {
        case "\n\t":
            v9 = float32((1e10 / 1e10))
        case "foo bar":
            v17 := GenS2{F0: -32768, F1: 42, F2: 79}
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 20; fi18++ {
            v3 = fi18
            switch "foo bar" {
            case "\n\t":
                switch fi18 {
                case -666:
                    v9 = v9
                case 32767:
                    v1 = v1
                case 0:
                    v19 := GenS2{F0: 623, F1: int8(v3), F2: -93}
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 7; fi20++ {
                        v21 := GenS3{F0: 42, F1: -128}
                        v22 := GenS2{F0: 42, F1: 1, F2: 98}
                        var v23 string = v1
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                }
            }
            var v24 int32 = 0
            _ = fi18
            _ = v24
        }
    }
    switch ("foo bar" + ("" + "a")) {
    case "a":
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 int32 = 42
            v27 := GenS4{F0: 255, F1: 48, F2: 127}
            if ((!(569 != 463)) && true) {
                if false {
                    var v28 string = v1
                    switch v1 {
                    case "hello":
                        if false {
                            if len(v14) > 0 {
                                v14[0] = (fi25 | (v3 | fi25))
                            }
                        }
                    }
                    _ = v28
                }
                fmt.Println(v3, v9, v3)
                v29 := GenS1{F0: false, F1: v1}
                _ = v29
            } else {
                var v30 int64 = 9223372036854775807
                v31 := GenS2{F0: 32767, F1: -82, F2: 127}
                v32 := v2[(v1 + v1)]
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi25
            _ = v26
            _ = v27
        }
    default:
        var v33 int64 = 0
        _ = v33
    }
    v9 += (v9 * v9)
    for fi34 := 0; fi34 < 12; fi34++ {
        for fi35 := 0; fi35 < 2; fi35++ {
            v9 = 1.5
            _ = fi35
        }
        v1 = v1
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi34
    }
    v3 = 4294967295
    for fi36 := 0; fi36 < 6; fi36++ {
        v3 = fi36
        v3 = fi36
        v14 = append(v14, v3)
        _ = fi36
    }
    if len(v14) > 0 {
        v14[0] = v3
    }
    var v37 float32 = v9
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v9
}
