package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0() int16 {
    var v6 uint = 255
    if (127 > 146) {
        v7 := GenS1{F0: "a", F1: 4294967295, F2: 726}
        _ = v7
    }
    _ = v6
    return 32767
}

func main() {
    v1 := []uint32{42, 4294967295, 0, 0}
    var v2 float32 = float32(((3.14159 + float64((float32(1.5) * ((0.1 - 3.14159) * 0.0)))) - ((1.5 - (((3.14159 + 3.14159) * 3.14159) + (((((3.14159 + (3.14159 + 1.5)) * 0.1) * 1.5) + float64(-2.25)) - -2.25))) + (0.1 + (0.1 + 3.14159)))))
    var v3 int64 = 0
    var v4 uint16 = 53
    var v5 int32 = 252
    switch "a" {
    case "foo bar":
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                _ = len(v1)
                v1 = append(v1, uint32(v4))
                v4 = (^(v4))
                _ = fi9
            }
            fmt.Println(v3, v2, v5)
            _ = fi8
        }
    case "hello":
        v3 += (v3 | 42)
    default:
        for fi10 := 0; fi10 < 5; fi10++ {
            if (127 == 1) {
                v11 := GenS2{F0: 0, F1: 0}
                var v12 int = fi10
                v3 += v3
                _ = v11
                _ = v12
            }
            _ = fi10
        }
    }
    var v13 uint8 = 1
    fmt.Println(v4, v13)
    switch 127 {
    case -942:
        for fi14 := 0; fi14 < 6; fi14++ {
            switch (fi14 * fi14) {
            case -2147483648:
                for fi15 := 0; fi15 < 10; fi15++ {
                    if (1.5 < float64((float32(-2.25) / v2))) {
                        var v16 string = (("a" + "") + ("a" + "a"))
                        v17 := GenS2{F0: 873, F1: 53}
                        fmt.Println(fi15, v4)
                        var v18 int = (fi14 & fi14)
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    var v19 string = ("foo bar" + ("a" + "foo bar"))
                    _ = fi15
                    _ = v19
                }
            }
            _ = len(v1)
            fmt.Println(v5, v3)
            _ = fi14
        }
    case 32768:
        fmt.Println(v5, v4, v5)
    case -714:
        v2 -= 0.0
    default:
        fmt.Println(v5)
    }
    fmt.Println(v5)
    v20 := GenS3{F0: false, F1: (v4 << ((218) & 63))}
    var v21 int16 = -1
    for fi22 := 0; fi22 < 16; fi22++ {
        fmt.Println(v3, v13, fi22)
        var v23 int64 = v3
        var v24 uint8 = (v13 | (v13 - 42))
        _ = fi22
        _ = v23
        _ = v24
    }
    switch ("hello" + ("hello" + ("hello" + "\n\t"))) {
    case "foo bar":
        if (!true) {
            var v25 float64 = (float64(v2) / 1.5)
            _ = v25
        } else {
            var v26 uint32 = 42
            _ = v26
        }
    }
    fmt.Println(v21, v3, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
