package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int16; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 uint8; F2 int16 }

func main() {
    var v1 float32 = (float32(1e10) / 1e10)
    var v2 int8 = -128
    v3 := []int{323, 42, 569, 32768}
    var v4 bool = ((!false) || (true && false))
    switch 535 {
    case 9223372036854775807:
        switch "" {
        case "":
            v5 := GenS1{F0: 879, F1: 0}
            _ = v5
        case "a":
            for fi6 := 0; fi6 < 17; fi6++ {
                v7 := GenS2{F0: 0}
                v4 = v4
                v1 += v1
                _ = fi6
                _ = v7
            }
        case "\n\t":
            for fi8 := 0; fi8 < 2; fi8++ {
                v4 = (3.14159 > (v1 * float32(1e10)))
                for fi9 := 0; fi9 < 3; fi9++ {
                    v10 := GenS1{F0: -1, F1: 142}
                    if true {
                        var v11 uint64 = 4294967295
                        _ = v11
                    }
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
        default:
            if len(v3) > 0 {
                v3[0] = 298
            }
        }
    case 32768:
        switch ((("a" + "hello") + ("hello" + ("" + ((("a" + "hello") + "") + "")))) + ("\n\t" + "a")) {
        case "":
            v1 += (v1 + (float32((0.1 * 1e10)) - (v1 * (-2.25 + v1))))
        case "\n\t":
            if (v1 == v1) {
                if (!(-618 < 0)) {
                    v1 = float32(1.5)
                }
                v1 += float32(-2.25)
            }
        }
    case -599:
        var v12 uint8 = 128
        _ = v12
    default:
        v1 = v1
    }
    _ = len(v3)
    if v4 {
        var v13 uint = 1
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 18; fi15++ {
            if (v1 == v1) {
                var v16 uint16 = 89
                var v17 float64 = 0.0
                _ = v16
                _ = v17
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v18 := GenS1{F0: 42, F1: int64(fi14)}
                v1 -= v1
                _ = v18
            }
            v2 = v2
            v19 := GenS3{F0: 2147483648, F1: 42, F2: (0.1 <= ((1.5 * -2.25) + 1.5))}
            var v20 float32 = (float32(1.5) - float32(float64((0.0 * (v1 * 0.1)))))
            switch ("a" + "a") {
            case "\n\t":
                if v4 {
                    v2 -= v2
                    var v21 uint16 = 682
                    v22 := GenS2{F0: 751}
                    v23 := GenS2{F0: 2}
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    v2 -= 44
                }
            }
            v24 := GenS4{F0: v4}
            _ = fi15
            _ = v19
            _ = v20
            _ = v24
        }
        _ = fi14
    }
    if v4 {
        var v25 uint = 18446744073709551615
        v26 := GenS5{F0: -289, F1: 102, F2: -768}
        _ = v25
        _ = v26
    } else {
        var v27 int64 = -1
        _ = v27
    }
    fmt.Println(v1)
    var v28 int = 32768
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
