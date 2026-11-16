package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 string }

func main() {
    var v1 float32 = 1.5
    var v2 bool = true
    var v3 int32 = 2147483647
    fmt.Println(v2, v2, v3)
    if (("hello" < "hello") || v2) {
        for fi4 := 0; fi4 < 6; fi4++ {
            for fi5 := 0; fi5 < 8; fi5++ {
                var v6 string = ""
                _ = fi5
                _ = v6
            }
            v7 := []int{}
            if len(v7) > 0 {
                v7[0] = (fi4 >> ((822) & 63))
            }
            _ = fi4
            _ = v7
        }
        v8 := GenS1{F0: (11 != 127), F1: 1, F2: 0}
        var v9 uint16 = 56
        _ = v8
        _ = v9
    }
    switch 4294967295 {
    case 32767:
        if (("foo bar" + "hello") < "\n\t") {
            v1 = (float32(3.14159) - v1)
            if (0.1 > ((0.0 - (((3.14159 + 0.0) + 0.1) * ((-2.25 / 1.5) * ((1.5 * 0.0) * float64((float32((3.14159 / 3.14159)) + 1e10)))))) / 1.5)) {
                v1 = (1e10 - 3.14159)
                if ((v3 + v3) >= (v3 ^ (v3 | (v3 + (^(v3)))))) {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        var v11 int8 = -128
                        _ = fi10
                        _ = v11
                    }
                    v2 = (false && (int16(v3) < 1000))
                    var v12 uint64 = 906
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v14 := GenS2{F0: 128, F1: 1}
                        var v15 uint8 = 127
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
            } else {
                if (!(669 < -877)) {
                    var v16 uint64 = 600
                    _ = v16
                }
                for fi17 := 0; fi17 < 17; fi17++ {
                    v1 += 3.14159
                    _ = fi17
                }
                var v18 uint16 = 1
                _ = v18
            }
            v1 = (v1 * v1)
        }
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        v3 -= -2147483648
        v3 -= (v3 % ((-2147483648) | 1))
        var v20 int8 = 1
        _ = fi19
        _ = v20
    }
    v3 += -2147483648
    v1 -= (v1 / v1)
    v21 := []int{786, 587, -9223372036854775808, 977}
    for fi22 := 0; fi22 < 4; fi22++ {
        if (151 > 1) {
            switch "foo bar" {
            case "foo bar":
                var v23 float64 = ((1e10 - 0.0) - ((0.0 * 0.0) + 1.5))
                _ = v23
            }
            v24 := GenS1{F0: ((!v2) && (207 < 688)), F1: 0, F2: 1}
            switch -2147483648 {
            case -323:
                v25 := GenS1{F0: false, F1: 255, F2: 0}
                _ = v25
            default:
                fmt.Println(v2, fi22)
            }
            _ = v24
        } else {
            v26 := GenS3{F0: -1, F1: 1, F2: "foo bar"}
            _ = v26
        }
        v1 -= (0.0 + (v1 + (float32(0.1) - -2.25)))
        _ = fi22
    }
    v2 = v2
    switch 1 {
    case 584:
        switch "foo bar" {
        case "\n\t":
            v2 = v2
        case "":
            v1 = 0.0
        default:
            fmt.Println(v1)
        }
    case 127:
        v27 := GenS1{F0: v2, F1: 255, F2: 128}
        _ = v27
    case 32767:
        v1 = 3.14159
    default:
        if v2 {
            v28 := GenS1{F0: (true && ((^(v3)) <= (-(v3)))), F1: 127, F2: 42}
            _ = v28
        } else {
            var v29 float64 = 0.0
            _ = v29
        }
    }
    v2 = true
    v2 = (!v2)
    var v30 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
}
