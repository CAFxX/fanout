package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint16; F1 uint16; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint8 = 58
    v2 := []string{"", "", "", ("a" + "\n\t")}
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 float64 = -2.25
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: (false || true), F1: (v1 >> ((86) & 63)), F2: 427}
    v5 += float64((-2.25 + (0.1 - float32(v5))))
    v8 := GenS1{F0: (!true), F1: (v1 - 0), F2: 42}
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(v5, v1)
        _ = fi9
    }
    if (v1 <= (v1 << ((65535) & 63))) {
        switch 32767 {
        case 32767:
            v10 := GenS2{F0: ("foo bar" + ("foo bar" + ("a" + "hello"))), F1: 255}
            _ = v10
        case 128:
            v1 += 91
        }
    } else {
        fmt.Println(v5)
        if false {
            v11 := GenS3{F0: 698, F1: true}
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 float32 = 1e10
                fmt.Println(v5, v13)
                _ = fi12
                _ = v13
            }
            v14 := GenS4{F0: 689, F1: 707, F2: "foo bar"}
            _ = v11
            _ = v14
        } else {
            if ((18446744073709551615 != 65535) || ((true || (-2147483648 < -2147483648)) || (65535 <= 1))) {
                v1 = (v1 | (v1 & 2))
                v2 = append(v2, "")
                for fi15 := 0; fi15 < 5; fi15++ {
                    v1 -= 2
                    _ = fi15
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = ("" + "")
                }
                fmt.Println(v5, v5)
            }
            var v16 int64 = 279
            _ = v16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            switch 128 {
            case 4294967295:
                v5 = (((float64((0.1 + -2.25)) / float64(-2.25)) / v5) - -2.25)
            case 42:
                for fi18 := 0; fi18 < 14; fi18++ {
                    v19 := v3[1]
                    v19 = -2147483648
                    _ = fi18
                    _ = v19
                }
            default:
                for fi20 := 0; fi20 < 8; fi20++ {
                    fmt.Println(v5, v1, fi20)
                    var v21 int32 = -428
                    switch "" {
                    case "hello":
                        var v22 int64 = int64(v1)
                        _ = v22
                    }
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi17
        }
    }
    v2 = append(v2, ("\n\t" + ""))
    for fi23 := 0; fi23 < 8; fi23++ {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        v24 := GenS5{F0: 0}
        v25 := GenS6{F0: 1}
        _ = fi23
        _ = v24
        _ = v25
    }
    v1 += v1
    v5 -= 3.14159
    if (true && false) {
        v1 = v1
        v26 := v4[4294967295]
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
