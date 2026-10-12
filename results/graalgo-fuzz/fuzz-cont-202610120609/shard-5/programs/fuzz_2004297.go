package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int8 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 float64 = v1
    v5 := GenS1{F0: 128}
    if false {
        v1 = v1
    } else {
        v6 := GenS1{F0: -9223372036854775808}
        v4 = 0.0
        if (((-2.25 == (1.5 / 1e10)) || false) || false) {
            for fi7 := 0; fi7 < 7; fi7++ {
                if false {
                    v4 = float64(0.1)
                }
                v8 := GenS2{F0: 0, F1: 366, F2: -622}
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 uint64 = 1
                var v11 uint64 = (^(v10))
                _ = fi9
                _ = v10
                _ = v11
            }
            v12 := GenS2{F0: 805, F1: 65535, F2: 526}
            _ = v12
        } else {
            switch "" {
            case "a":
                for fi13 := 0; fi13 < 10; fi13++ {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        if true {
                            v3[("a" + "\n\t")] = 1
                            var v15 uint64 = 2
                            fmt.Println(fi13, v1)
                            _ = v15
                        }
                        v16 := []int{(fi13 ^ (fi13 & (fi14 % ((fi13) | 1)))), (fi14 + (-(fi14))), -1}
                        var v17 int8 = 127
                        _ = fi14
                        _ = v16
                        _ = v17
                    }
                    if true {
                        switch ("foo bar" + ("hello" + "foo bar")) {
                        case "":
                            v18 := GenS2{F0: 2, F1: 0, F2: -446}
                            _ = v18
                        case "hello":
                            v19 := []int{32767, (fi13 >> ((333) & 63)), fi13}
                            _ = v19
                        case "a":
                            var v20 int64 = 56
                            _ = v20
                        }
                        v21 := []int{}
                        _ = v21
                    }
                    v1 += 3.14159
                    _ = fi13
                }
            case "foo bar":
                var v22 string = ("foo bar" + "")
                _ = v22
            default:
                v23 := v2[954]
                _ = v23
            }
            v24 := []int{32768, 32767}
            _ = v24
        }
        _ = v6
    }
    if (!true) {
        if true {
            fmt.Println(v4, v1)
            for fi25 := 0; fi25 < 17; fi25++ {
                for fi26 := 0; fi26 < 16; fi26++ {
                    fmt.Println(fi26, fi26)
                    _ = fi26
                }
                _ = fi25
            }
        }
        v4 = (1e10 + float64((0.0 / 1.0)))
    }
    v1 += ((v4 - 0.0) + ((float64(3.14159) - v4) + (((v1 + v4) / (((0.1 - float64(float32(v1))) / v4) + (v1 - (v4 / v1)))) / 0.1)))
    var v27 uint = 742
    var v28 uint16 = 786
    v29 := GenS3{F0: 961, F1: 0}
    for fi30 := 0; fi30 < 5; fi30++ {
        v4 = v1
        _ = fi30
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        v28 += (v28 - (^(v28)))
        _ = fi31
    }
    var v32 int = -763
    switch (v32 - v32) {
    case 128:
        v33 := GenS4{F0: 32767, F1: 255, F2: 127}
        _ = v33
    case -957:
        fmt.Println(v27, v27, v27)
    case 32768:
        for fi34 := 0; fi34 < 15; fi34++ {
            fmt.Println(fi34)
            _ = fi34
        }
    }
    v28 += (v28 << ((v27) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
