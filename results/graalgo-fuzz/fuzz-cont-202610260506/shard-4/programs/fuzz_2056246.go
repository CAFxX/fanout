package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint16; F1 string }
func genFn0() uint16 {
    v4 := GenS1{F0: 65535, F1: -714, F2: 865}
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 string = "a"
        fmt.Println(v6)
        v7 := []int{(fi5 + fi5), (^(fi5)), fi5, fi5}
        _ = fi5
        _ = v6
        _ = v7
    }
    if ((775 == 73) && true) {
        var v8 bool = false
        _ = v8
    }
    var v9 float32 = float32((((((1.5 * 0.1) - (((float64((3.14159 - 0.1)) / 1.0) * (((((0.1 - 0.1) / 1e10) * (3.14159 * 3.14159)) * ((1e10 / 3.14159) * 3.14159)) - (float64(1e10) / -2.25))) / ((0.0 * 1.5) + float64(-2.25)))) * float64(0.1)) / 3.14159) * ((1e10 / 3.14159) + float64((0.0 + -2.25)))))
    var v10 uint32 = 765
    _ = v10
    _ = v4
    _ = v9
    return 153
}

func main() {
    var v1 float32 = 0.0
    var v2 int16 = 42
    v3 := make(map[uint32]uint32)
    var v11 float32 = (v1 - 1.5)
    v2 -= (^(v2))
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int32 = -733
        v2 += v2
        _ = fi12
        _ = v13
    }
    v2 = (v2 >> ((700) & 63))
    var v14 float64 = (3.14159 * 1.5)
    var v15 float64 = 0.1
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 uint32 = 65535
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v19 int = fi17
                _ = v19
            case 684:
                switch (fi16 & -89) {
                case 32768:
                    fmt.Println(v11, v2, v14)
                case 1:
                    var v20 uint64 = 70
                    _ = v20
                case 0:
                    var v21 uint8 = 127
                    _ = v21
                default:
                    fmt.Println(fi16)
                }
            default:
                if (v15 > v14) {
                    fmt.Println(v14, v2)
                }
            }
            var v22 bool = false
            _ = fi17
            _ = v18
            _ = v22
        }
        var v23 uint64 = 660
        _ = fi16
        _ = v23
    }
    fmt.Println(v1)
    switch 4294967295 {
    case 698:
        fmt.Println(v2, v14)
    case 1:
        v24 := GenS2{F0: 278, F1: 9223372036854775807, F2: 0}
        _ = v24
    case 42:
        fmt.Println(v1, v1)
    default:
        fmt.Println(v1, v2, v1)
    }
    var v25 bool = true
    for fi26 := 0; fi26 < 3; fi26++ {
        fmt.Println(fi26)
        _ = fi26
    }
    var v27 uint8 = 255
    if (!(1 <= 300)) {
        fmt.Println(v2, v1, v14)
    } else {
        v28 := GenS3{F0: 930, F1: "\n\t"}
        v3[988] = 30
        if v25 {
            var v29 uint16 = genFn0()
            _ = len(v3)
            v30 := v3[0]
            _ = v29
            _ = v30
        } else {
            var v31 int = 0
            v15 += (0.1 - -2.25)
            _ = v31
        }
        _ = v28
    }
    if (195 < v27) {
        v32 := GenS3{F0: 546, F1: "\n\t"}
        var v33 float32 = v11
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v27
    _ = v3
}
