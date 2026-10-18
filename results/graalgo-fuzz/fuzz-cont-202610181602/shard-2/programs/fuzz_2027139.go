package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int8; F2 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 671
    var v3 uint8 = 203
    var v4 uint8 = v3
    var v5 uint32 = 679
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: 150}
        _ = fi6
        _ = v7
    }
    if (!(v2 < v2)) {
        switch 1 {
        case -953:
            for fi8 := 0; fi8 < 9; fi8++ {
                if (!(!true)) {
                    if (9223372036854775807 > 445) {
                        v9 := []int{(fi8 * 9223372036854775807)}
                        for fi10 := 0; fi10 < 19; fi10++ {
                            v11 := GenS2{F0: 344, F1: 1}
                            _ = fi10
                            _ = v11
                        }
                        v3 = 83
                        _ = v9
                    }
                    switch "\n\t" {
                    case "foo bar":
                        var v12 int8 = -128
                        _ = v12
                    }
                } else {
                    v13 := GenS2{F0: 42, F1: -88}
                    if (false || (!(!((fi8 >> ((18446744073709551615) & 63)) != -9223372036854775808)))) {
                        fmt.Println(fi8)
                    } else {
                        v14 := GenS3{F0: -465}
                        if false {
                            v15 := GenS4{F0: (3.14159 != float32(float64(3.14159))), F1: -1, F2: 10}
                            var v16 int64 = 99
                            _ = v15
                            _ = v16
                        }
                        _ = v14
                    }
                    _ = v13
                }
                _ = fi8
            }
        case -2147483648:
            v3 -= 128
        default:
            v17 := []int{496, 2147483647, 441}
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 11; fi18++ {
            if (true || (!(false || ((!false) || (!(v2 < (^(v2)))))))) {
                var v19 uint16 = 0
                _ = v19
            }
            switch (-(fi18)) {
            case 249:
                switch -2147483648 {
                case -720:
                    v3 -= (v4 << ((255) & 63))
                case -1:
                    var v20 int64 = int64(v5)
                    _ = v20
                }
            case -222:
                var v21 float32 = 0.1
                _ = v21
            }
            _ = fi18
        }
    }
    v3 = (v3 | v3)
    fmt.Println(v3, v3, v4)
    v2 = (v2 >> ((1) & 63))
    for fi22 := 0; fi22 < 5; fi22++ {
        v23 := GenS4{F0: (0 != v2), F1: -128, F2: -128}
        v4 = 1
        _ = fi22
        _ = v23
    }
    var v24 float32 = float32(0.0)
    v3 = (v4 | 239)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
