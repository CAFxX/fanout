package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int16; F1 int64; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint32, p1 int16) string {
    v5 := GenS1{F0: -1, F1: 618}
    var v6 int = 4294967295
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{}
    v3 := []uint32{0, 42}
    var v4 int8 = 100
    for fi7 := 0; fi7 < 15; fi7++ {
        if (2 != 2) {
            v1 = v1
            var v8 uint32 = 4294967295
            v2 = append(v2, fi7)
            _ = v8
        } else {
            switch v1 {
            case "hello":
                if (0 != 464) {
                    v1 = "a"
                    for fi9 := 0; fi9 < 7; fi9++ {
                        var v10 float32 = (float32(1e10) - (0.1 + (-2.25 * 1.5)))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi9
                        _ = v10
                    }
                    if ((((((1e10 * -2.25) / ((3.14159 - ((((0.1 - 3.14159) / 3.14159) - 1e10) * (0.0 + (float64((-2.25 - float32((0.1 + ((-2.25 * (0.0 - 1.5)) * 0.0))))) * (1e10 - -2.25))))) + float64(0.0))) / (3.14159 - -2.25)) - (0.0 + 3.14159)) >= 0.0) && true) {
                        v4 = 21
                        switch "hello" {
                        case "a":
                            var v11 string = ("hello" + ((((v1 + "a") + ((((("a" + ((v1 + (((v1 + "hello") + v1) + v1)) + ("" + v1))) + v1) + "foo bar") + "\n\t") + v1)) + (("foo bar" + ((v1 + (v1 + (v1 + ("\n\t" + v1)))) + "")) + "")) + v1))
                            _ = v11
                        case "foo bar":
                            fmt.Println(v4, v4)
                        case "hello":
                            for fi12 := 0; fi12 < 14; fi12++ {
                                var v13 uint64 = 18446744073709551615
                                _ = fi12
                                _ = v13
                            }
                        }
                        v1 = ""
                    }
                } else {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        var v15 int = (fi14 + fi7)
                        var v16 uint16 = 410
                        v15 = fi14
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    if true {
                        v17 := GenS2{F0: 936, F1: 9223372036854775807, F2: ("hello" + v1)}
                        _ = v17
                    }
                    v4 = 1
                }
            default:
                fmt.Println(v1)
            }
            switch v1 {
            case "hello":
                var v18 int16 = -32768
                _ = v18
            case "foo bar":
                for fi19 := 0; fi19 < 5; fi19++ {
                    v1 = (("" + "foo bar") + (("foo bar" + ("a" + v1)) + (v1 + (v1 + ("hello" + ("a" + v1))))))
                    var v20 uint8 = 30
                    _ = fi19
                    _ = v20
                }
            }
        }
        switch "foo bar" {
        case "":
            fmt.Println(v1, v4)
        default:
            _ = len(v2)
        }
        _ = fi7
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v21 int16 = -32768
    fmt.Println(v4)
    v21 -= (v21 / (((v21 - (v21 - v21))) | 1))
    var v22 string = (v1 + (v1 + v1))
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS3{F0: 2}
        for fi25 := 0; fi25 < 6; fi25++ {
            if false {
                if ((v22 + (v22 + (v1 + v1))) <= ((((("hello" + "") + v22) + (v1 + "a")) + (v1 + "")) + (v22 + (v22 + (v22 + ""))))) {
                    fmt.Println(v1, v4)
                    var v26 uint32 = 65535
                    v27 := GenS1{F0: 9223372036854775807, F1: 77}
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(v4, fi23, v4)
                    var v28 float64 = 3.14159
                    _ = v28
                }
            }
            if len(v2) > 0 {
                v2[0] = fi23
            }
            v22 = "a"
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v4)
    var v29 uint32 = 0
    v4 -= int8(v21)
    switch v1 {
    case "foo bar":
        _ = len(v3)
    case "hello":
        v30 := GenS2{F0: v21, F1: -1, F2: v1}
        _ = v30
    default:
        _ = len(v2)
    }
    if len(v2) > 0 {
        v2[0] = -862
    }
    if (float64((1.5 * (3.14159 - 1.5))) != 0.0) {
        v29 += 465
        v2 = append(v2, -1)
    } else {
        if (!true) {
            if len(v3) > 0 {
                v3[0] = 137
            }
        }
        _ = len(v2)
    }
    switch 128 {
    case 313:
        v29 -= v29
    case 42:
        v31 := GenS4{F0: v4}
        _ = v31
    case 128:
        v22 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
}
