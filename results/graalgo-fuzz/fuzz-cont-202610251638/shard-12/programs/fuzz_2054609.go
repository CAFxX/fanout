package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int; F1 uint; F2 int16 }
type GenS4 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0(p0 uint16) string {
    var v6 int64 = -426
    var v7 uint8 = 128
    v6 = (v6 >> ((747) & 63))
    _ = p0
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    var v1 int16 = 32767
    var v2 int = 0
    var v3 int = 128
    v4 := make(map[uint32]uint32)
    var v5 int8 = -1
    if true {
        if (907 >= 1) {
            switch "hello" {
            case "foo bar":
                for fi8 := 0; fi8 < 18; fi8++ {
                    v9 := GenS1{F0: -1, F1: true}
                    _ = fi8
                    _ = v9
                }
            case "a":
                var v10 string = "foo bar"
                _ = v10
            }
            v11 := GenS1{F0: 0, F1: (!false)}
            _ = v11
        }
    } else {
        v3 -= (v2 / ((v3) | 1))
        if (((0.0 + 3.14159) - (0.0 * (0.0 * 1.5))) == (1.5 + -2.25)) {
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS2{F0: 255, F1: 255}
                _ = fi12
                _ = v13
            }
        } else {
            v14 := GenS2{F0: 32767, F1: 18446744073709551615}
            _ = v14
        }
        var v15 int16 = 1
        _ = v15
    }
    v16 := []int{(^(v2)), -926, v2}
    v17 := v4[331]
    v5 = 0
    v3 -= 128
    v18 := GenS3{F0: (v3 | 127), F1: 491, F2: 778}
    var v19 uint64 = uint64(v17)
    v3 += (v3 % (((v2 >> ((217) & 63))) | 1))
    v20 := GenS4{F0: 42, F1: -1, F2: -9223372036854775808}
    if ((128 >= 2) && (((^(v17)) <= (v17 % ((v17) | 1))) && ((v1 - v1) == (v1 ^ v1)))) {
        switch "\n\t" {
        case "a":
            if false {
                var v21 uint8 = 2
                _ = v21
            }
        case "":
            fmt.Println(v19)
        case "foo bar":
            if (3.14159 != 3.14159) {
                for fi22 := 0; fi22 < 8; fi22++ {
                    v19 -= 1
                    _ = fi22
                }
            }
        default:
            if (false && false) {
                v1 += 32767
            } else {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                fmt.Println(v5, v17)
                var v23 int16 = -664
                _ = v23
            }
        }
        var v24 float32 = -2.25
        v25 := GenS1{F0: 42, F1: ((v5 / (((v5 % ((v5) | 1))) | 1)) != v5)}
        _ = v24
        _ = v25
    }
    delete(v4, (v17 << ((4294967295) & 63)))
    v1 += 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
