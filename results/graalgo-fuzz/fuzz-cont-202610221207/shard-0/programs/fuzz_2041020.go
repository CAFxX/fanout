package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 float32 = 1e10
    var v2 uint16 = 65535
    var v3 uint8 = 128
    v4 := []uint32{}
    var v5 float32 = (v1 / 1.5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (!((1e10 == ((v5 * (-2.25 / 1.5)) + v1)) && (0 != 18446744073709551615))) {
        switch 359 {
        case 2147483647:
            fmt.Println(v2)
        default:
            var v6 uint32 = 4294967295
            _ = v6
        }
        if (v5 < (v1 - ((v1 * v5) / 1.0))) {
            fmt.Println(v2, v1)
            v7 := GenS1{F0: (^(v2)), F1: 42, F2: v2}
            var v8 int32 = 1
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 1; fi9++ {
                switch (fi9 * 32768) {
                case 32768:
                    fmt.Println(v5, v3, v3)
                case -2147483648:
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v4 = append(v4, 65535)
                        var v11 int16 = 1
                        _ = fi10
                        _ = v11
                    }
                case 776:
                    var v12 float64 = 1e10
                    _ = v12
                }
                var v13 int = (-(fi9))
                _ = fi9
                _ = v13
            }
            v3 -= v3
            switch -2147483648 {
            case 1:
                for fi14 := 0; fi14 < 9; fi14++ {
                    if false {
                        v2 -= (^(v2))
                    }
                    _ = fi14
                }
            case 2147483647:
                fmt.Println(v5)
            case 127:
                fmt.Println(v1, v5, v1)
            }
        }
        switch 579 {
        case 458:
            v2 = v2
        case 42:
            fmt.Println(v1)
        case 32768:
            v1 = (0.1 * -2.25)
        default:
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := GenS1{F0: 32767, F1: 73, F2: (v2 >> ((18446744073709551615) & 63))}
                v1 -= (((float32(0.0) / (v5 * (v5 + v1))) + (v5 * v1)) - (1e10 * v1))
                _ = fi15
                _ = v16
            }
        }
    }
    if ((v2 / (((v2 >> ((255) & 63))) | 1)) == v2) {
        _ = len(v4)
        _ = len(v4)
    }
    v2 = 32767
    for fi17 := 0; fi17 < 8; fi17++ {
        if (1 >= 115) {
            _ = len(v4)
        } else {
            v5 -= (float32(3.14159) - (v5 / float32((1.5 / (1.5 - -2.25)))))
        }
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 uint32 = 725
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 float64 = 1e10
    v2 -= v2
    if ((((((v1 * (v1 / 1.0)) * (float32((((v20 * 3.14159) / v20) - float64(float32(v20)))) - float32(-2.25))) / v1) - v1) - (-2.25 * float32((float64((3.14159 + ((float32(v20) * v5) / v1))) / float64(-2.25))))) <= v5) {
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 float64 = ((((3.14159 / float64(-2.25)) - ((0.1 - (((v20 - (((v20 - v20) - v20) - v20)) + (v20 / (0.1 + (0.1 * v20)))) - (3.14159 * 0.0))) - (v20 * v20))) - 0.1) - v20)
            _ = fi21
            _ = v22
        }
    } else {
        v1 = v5
        v23 := GenS1{F0: (^(v2)), F1: 9, F2: (v2 + v2)}
        var v24 string = ("hello" + ("hello" + ("\n\t" + "")))
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 uint = 255
        var v27 uint = 789
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := GenS2{F0: (1000 == 770)}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
