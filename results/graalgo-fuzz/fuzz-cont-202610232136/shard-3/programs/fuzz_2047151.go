package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint64; F1 int; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 int32) uint16 {
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint8 = 2
        if false {
            switch -301 {
            case 127:
                var v8 float32 = (0.0 / 1e10)
                _ = v8
            case 42:
                fmt.Println(v7)
            case 1:
                for fi9 := 0; fi9 < 19; fi9++ {
                    v7 = (v7 ^ (v7 << ((330) & 63)))
                    v10 := GenS1{F0: -464, F1: 2}
                    _ = fi9
                    _ = v10
                }
            }
        } else {
            p0 = p0
        }
        _ = fi6
        _ = v7
    }
    if false {
        var v11 uint = 0
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := []int{(fi12 - fi12), fi12, (fi12 / (((fi12 % ((-279) | 1))) | 1))}
            p0 += p0
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    var v14 int8 = -1
    var v15 uint16 = 1
    _ = p0
    _ = v14
    _ = v15
    return v15
}

func main() {
    var v1 uint8 = 208
    var v2 uint32 = 65535
    var v3 float64 = -2.25
    var v4 int = -580
    v5 := []int{v4, v4, 436}
    fmt.Println(v1, v4, v2)
    v16 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615}
    v17 := GenS1{F0: -899, F1: 926}
    var v18 int8 = 1
    var v19 int32 = 566
    for fi20 := 0; fi20 < 18; fi20++ {
        switch 9223372036854775807 {
        case 32767:
            switch 42 {
            case 0:
                var v21 int16 = 1
                _ = v21
            case 32768:
                switch -9223372036854775808 {
                case -792:
                    v4 = 4294967295
                case 9223372036854775807:
                    v18 = (v18 / (((v18 - v18)) | 1))
                }
            default:
                v18 = 0
            }
        case 393:
            var v22 uint8 = (v1 << ((65535) & 63))
            _ = v22
        case -767:
            v18 -= (v18 >> ((740) & 63))
        }
        var v23 int16 = 614
        v24 := GenS2{F0: 4294967295, F1: fi20, F2: genFn0(v19)}
        _ = fi20
        _ = v23
        _ = v24
    }
    var v25 int32 = (v19 * (v19 - -717))
    var v26 string = ("foo bar" + "a")
    v18 += (^(v18))
    var v27 uint32 = v2
    var v28 int32 = v25
    v2 -= (v2 & (v2 ^ (v2 % ((v27) | 1))))
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 uint32 = v27
        v27 = (^(v30))
        v31 := GenS3{F0: 42, F1: 2}
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int = v4
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
