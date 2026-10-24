package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint8; F1 string; F2 string }
type GenS6 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint8, p1 uint16, p2 int32) uint16 {
    if (0.1 >= -2.25) {
        var v4 float32 = float32(1e10)
        for fi5 := 0; fi5 < 6; fi5++ {
            if false {
                p0 = (p0 >> ((65535) & 63))
                v6 := GenS1{F0: true, F1: fi5}
                p2 -= (p2 << ((231) & 63))
                _ = v6
            }
            v7 := GenS1{F0: (127 >= (^(p0))), F1: fi5}
            var v8 float32 = float32(0.0)
            _ = fi5
            _ = v7
            _ = v8
        }
        var v9 uint32 = 4294967295
        _ = v4
        _ = v9
    }
    p0 = p0
    p2 += (p2 << ((255) & 63))
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint16 = 724
    var v2 bool = false
    var v3 uint16 = v1
    v3 -= 2
    v10 := GenS2{F0: false, F1: 1, F2: (v3 << ((18446744073709551615) & 63))}
    var v11 int = -9223372036854775808
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(v11, v11)
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(fi12, v11, v11)
            _ = fi13
        }
        _ = fi12
    }
    if (!false) {
        v2 = false
        if (((((3.14159 + -2.25) * 3.14159) + 0.1) <= 3.14159) || (((true && true) && (!v2)) && (1 == 127))) {
            v14 := []int{v11, (-(v11)), v11}
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := []int{}
                if len(v17) > 0 {
                    v17[0] = v11
                }
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    } else {
        v18 := GenS3{F0: 2, F1: 515}
        var v19 int16 = -32768
        for fi20 := 0; fi20 < 19; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                v11 -= (^(fi21))
                if (!(fi20 >= (^(fi21)))) {
                    v22 := GenS4{F0: 1, F1: 99}
                    var v23 int16 = -1
                    fmt.Println(fi21)
                    _ = v22
                    _ = v23
                }
                switch fi21 {
                case -2147483648:
                    v24 := []int{1, (fi21 % (((fi21 | (v11 & fi20))) | 1))}
                    _ = v24
                case 32767:
                    v25 := GenS5{F0: 127, F1: "\n\t", F2: ""}
                    _ = v25
                default:
                    var v26 int32 = 1
                    _ = v26
                }
                _ = fi21
            }
            fmt.Println(fi20, v19)
            _ = fi20
        }
        _ = v18
        _ = v19
    }
    v27 := GenS6{F0: 953, F1: v3}
    v3 = 269
    var v28 float32 = (0.1 * -2.25)
    fmt.Println(v3, v28)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
