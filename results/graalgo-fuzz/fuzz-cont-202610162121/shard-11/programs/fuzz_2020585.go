package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 string; F1 uint8 }
func genFn0() bool {
    var v4 bool = (100 <= 0)
    v5 := []int{32767, -306, 32767, 600}
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 int64, p1 uint32) string {
    v6 := []int{1, 128, 32768}
    if true {
        var v7 uint32 = (p1 * 138)
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return "foo bar"
}

func main() {
    v1 := []string{("foo bar" + ("foo bar" + "a"))}
    v2 := make(map[uint32]uint32)
    var v3 int16 = 1
    var v8 uint32 = 2
    fmt.Println(v8, v3)
    var v9 int64 = 60
    if false {
        if false {
            v10 := GenS1{F0: (42 != (v8 & 415))}
            v3 = v3
            _ = v10
        } else {
            v9 -= int64(v8)
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        }
        var v11 bool = true
        v12 := GenS1{F0: (false && v11)}
        _ = v11
        _ = v12
    } else {
        var v13 uint = 18446744073709551615
        if (42 <= 32767) {
            for fi14 := 0; fi14 < 5; fi14++ {
                switch fi14 {
                case 735:
                    var v15 bool = false
                    _ = v15
                }
                var v16 bool = (-2.25 <= ((0.0 + 0.0) + -2.25))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi14
                _ = v16
            }
            for fi17 := 0; fi17 < 20; fi17++ {
                v3 = (v3 ^ (v3 << (((v13 >> ((v13) & 63))) & 63)))
                _ = fi17
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                v3 = -1
                var v19 float64 = (float64((float32(0.1) / 1.0)) - ((float64(0.0) * ((1.5 - 0.1) * 1.5)) / ((3.14159 / -2.25) + 3.14159)))
                _ = fi18
                _ = v19
            }
        } else {
            var v20 int64 = v9
            var v21 int64 = int64(v3)
            switch -114 {
            case 9223372036854775807:
                v8 = (v8 ^ (v8 / ((v8) | 1)))
            }
            _ = v20
            _ = v21
        }
        switch 32768 {
        case -1:
            v22 := GenS2{F0: v3}
            _ = v22
        case -617:
            var v23 float32 = (0.1 - ((1.5 / 3.14159) / (1e10 * 0.1)))
            _ = v23
        }
        _ = v13
    }
    if true {
        v2[(v8 | v8)] = 1
        v24 := GenS3{F0: v9}
        _ = v24
    }
    fmt.Println(v9, v3, v9)
    v9 = v9
    for fi25 := 0; fi25 < 13; fi25++ {
        switch "" {
        case "hello":
            switch fi25 {
            case -1:
                fmt.Println(v3)
            case 1:
                v26 := GenS4{F0: ((true || (v9 <= -9223372036854775808)) && ((127 >= 71) || (v3 != (-(v3))))), F1: -403}
                _ = v26
            }
        default:
            v27 := GenS5{F0: ("hello" + ((("" + "\n\t") + (("a" + "a") + "foo bar")) + "\n\t")), F1: 2}
            _ = v27
        }
        if true {
            if false {
                v9 = 9223372036854775807
                v3 -= (v3 / (((v3 * v3)) | 1))
                var v28 uint16 = 32767
                _ = v28
            }
            for fi29 := 0; fi29 < 15; fi29++ {
                fmt.Println(v3, v9)
                _ = fi29
            }
        }
        var v30 uint8 = 127
        _ = fi25
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
