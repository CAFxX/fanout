package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 int = 1
    var v2 uint16 = 255
    v3 := []int{v1}
    var v4 uint = 684
    v5 := make(map[uint32]string)
    var v6 int32 = 0
    for fi7 := 0; fi7 < 10; fi7++ {
        if true {
            v8 := GenS1{F0: 255, F1: 32767}
            v1 = fi7
            _ = v8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint16 = 32767
        _ = fi9
        _ = v10
    }
    if (850 != 906) {
        v4 = (^(v4))
        if false {
            var v11 int16 = 1000
            v2 += v2
            _ = v11
        }
    } else {
        v12 := GenS1{F0: 128, F1: 32767}
        v6 = v6
        v2 = (v2 % ((65535) | 1))
        _ = v12
    }
    var v13 uint16 = 32767
    v13 -= uint16(v4)
    v1 -= (^(v1))
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS1{F0: 127, F1: v13}
        _ = fi14
        _ = v15
    }
    v13 = v2
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            switch 127 {
            case 1:
                if len(v3) > 0 {
                    v3[0] = fi16
                }
            default:
                v18 := GenS1{F0: 2, F1: v2}
                _ = v18
            }
            if (0.0 > (((0.0 + (-2.25 / 1.0)) / (1.5 + 0.0)) * 0.0)) {
                v4 = 2
                var v19 float64 = 0.1
                fmt.Println(v1, v6, v19)
                _ = v19
            } else {
                var v20 int32 = 0
                _ = v20
            }
            var v21 float64 = ((0.0 * float64(3.14159)) - 0.1)
            _ = fi17
            _ = v21
        }
        _ = fi16
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
