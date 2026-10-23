package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint8; F1 bool; F2 string }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int32; F1 int32 }
func genFn0(p0 bool) int32 {
    p0 = p0
    v7 := GenS1{F0: 695}
    _ = p0
    _ = v7
    return -1
}

func main() {
    var v1 uint8 = 203
    var v2 bool = true
    var v3 float32 = 1.5
    var v4 uint8 = v1
    v5 := make(map[uint32]string)
    var v6 int64 = 42
    v8 := GenS1{F0: (v6 >> ((0) & 63))}
    if (v2 && (-128 < -1)) {
        v2 = true
    } else {
        v9 := GenS1{F0: v6}
        var v10 string = ("foo bar" + "")
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: -1}
        v13 := GenS2{F0: (v6 % ((-753) | 1)), F1: (v6 + v6)}
        _ = fi11
        _ = v12
        _ = v13
    }
    switch -2147483648 {
    case 9223372036854775807:
        v14 := GenS3{F0: -59, F1: "\n\t"}
        _ = v14
    default:
        v6 -= int64(v4)
    }
    if ((((0.0 + 0.1) > (float64((v3 - (((0.1 - 3.14159) + v3) + ((v3 + 1e10) + (v3 - 1e10))))) - (-2.25 - ((float64(-2.25) * 0.0) / 1.5)))) && (65535 <= 0)) || (!true)) {
        v15 := GenS4{F0: (v4 + v1), F1: ((521 > 673) && v2), F2: "\n\t"}
        v2 = v2
        _ = v15
    } else {
        v16 := GenS5{F0: (-(v6)), F1: 0}
        _ = v16
    }
    var v17 string = "foo bar"
    for fi18 := 0; fi18 < 15; fi18++ {
        v4 = (v1 & 255)
        switch (v17 + v17) {
        case "\n\t":
            if false {
                var v19 int64 = (v6 / ((v6) | 1))
                v1 = 255
                for fi20 := 0; fi20 < 3; fi20++ {
                    fmt.Println(v3)
                    v21 := GenS6{F0: 1, F1: 264}
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
        default:
            _ = len(v5)
        }
        _ = fi18
    }
    var v22 int = 128
    v2 = (v2 || v2)
    v2 = ((((v17 + v17) + ("" + (v17 + "foo bar"))) + v17) >= v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
