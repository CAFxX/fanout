package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint16) uint32 {
    v5 := GenS1{F0: (-32768 <= 567)}
    if (0 <= 180) {
        switch 592 {
        case -1:
            p0 -= 255
        default:
            for fi6 := 0; fi6 < 2; fi6++ {
                p0 = 0
                _ = fi6
            }
        }
        var v7 float64 = (3.14159 + 0.1)
        _ = v7
    } else {
        p0 = (p0 ^ p0)
    }
    v8 := []int{2147483647, 4294967295, 32767, 934}
    _ = p0
    _ = v5
    _ = v8
    return 253
}
func genFn1() int {
    var v9 uint64 = 2
    var v10 bool = ((!(true && (375 > 502))) || (!(float64(1.5) != (((float64(1.5) + (3.14159 / ((-2.25 / 0.1) - 3.14159))) - 3.14159) + 0.0))))
    for fi11 := 0; fi11 < 11; fi11++ {
        if (1 >= 127) {
            v12 := GenS2{F0: 42, F1: 955}
            var v13 int32 = -748
            _ = v12
            _ = v13
        } else {
            var v14 float32 = 1e10
            _ = v14
        }
        v9 -= (v9 ^ v9)
        _ = fi11
    }
    v15 := GenS2{F0: 255, F1: 0}
    _ = v10
    _ = v15
    _ = v9
    return 127
}

func main() {
    var v1 string = ("\n\t" + "foo bar")
    var v2 uint8 = 211
    var v3 float32 = 0.1
    var v4 uint64 = 4294967295
    switch genFn1() {
    case 703:
        var v16 uint32 = 1
        _ = v16
    case 32768:
        fmt.Println(v2, v3)
    case -1:
        var v17 int32 = -1
        _ = v17
    default:
        v18 := []int{-695}
        _ = v18
    }
    v19 := []int{-9223372036854775808, 1}
    v20 := GenS3{F0: (v4 & 4294967295)}
    if true {
        v21 := GenS2{F0: v2, F1: 527}
        if true {
            var v22 int16 = -1
            if len(v19) > 0 {
                v19[0] = 477
            }
            v3 -= (v3 / ((((float32(-2.25) + (0.1 * (v3 * v3))) - (1.5 * float32(1e10))) + (-2.25 + v3)) + v3))
            _ = v22
        } else {
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 int64 = -9223372036854775808
                v24 = (v24 << ((264) & 63))
                var v25 int64 = (v24 * (v24 ^ (-(v24))))
                _ = fi23
                _ = v24
                _ = v25
            }
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        }
        if (v4 < v4) {
            for fi26 := 0; fi26 < 18; fi26++ {
                _ = len(v19)
                v2 = 2
                _ = fi26
            }
        }
        _ = v21
    } else {
        var v27 string = "foo bar"
        v27 = v27
        _ = v27
    }
    v28 := GenS1{F0: (!(-42 > 42))}
    v1 = ((v1 + ((v1 + "foo bar") + "\n\t")) + ("foo bar" + "foo bar"))
    v3 -= v3
    v29 := GenS3{F0: (v4 >> ((18446744073709551615) & 63))}
    v1 = (v1 + "")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
