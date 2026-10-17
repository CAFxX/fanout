package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint; F2 uint8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int16; F1 int32; F2 uint }
func genFn0() uint8 {
    v5 := GenS1{F0: 533}
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: 255}
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 834}
        _ = fi6
        _ = v9
    }
    _ = v5
    return 225
}
func genFn1(p0 int16) uint64 {
    p0 = (p0 * 1)
    v10 := GenS1{F0: 42}
    var v11 int32 = -525
    v11 = (-(v11))
    _ = p0
    _ = v10
    _ = v11
    return 573
}

func main() {
    v1 := []uint32{2147483647}
    var v2 uint8 = 2
    var v3 int8 = 0
    v4 := []uint32{860, 2147483647, 2, 4294967295}
    v12 := GenS1{F0: 2}
    v2 = (v2 >> ((255) & 63))
    var v13 float64 = 1e10
    var v14 string = ("hello" + "foo bar")
    v15 := GenS1{F0: 65535}
    if true {
        var v16 int8 = v3
        _ = v16
    } else {
        var v17 bool = false
        fmt.Println(v17)
        v18 := GenS1{F0: 18446744073709551615}
        _ = v17
        _ = v18
    }
    var v19 uint16 = uint16(v2)
    v19 -= 704
    v13 += (v13 + v13)
    v13 = v13
    v14 = ("foo bar" + "foo bar")
    var v20 int8 = (v3 | v3)
    var v21 int64 = -1
    if (255 >= 42) {
        if (!(1e10 > 1.5)) {
            v20 -= (v20 * (^(v20)))
            _ = len(v1)
            switch 128 {
            case -2147483648:
                if (1 == 857) {
                    v22 := GenS2{F0: 42, F1: 405, F2: (^(v2))}
                    _ = v22
                }
            case 2147483647:
                var v23 uint32 = 2147483647
                _ = v23
            case 127:
                var v24 int32 = 1
                _ = v24
            }
        } else {
            if false {
                v25 := GenS3{F0: 0, F1: v21}
                var v26 uint64 = 18446744073709551615
                _ = v25
                _ = v26
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 string = "foo bar"
                    _ = fi27
                    _ = v28
                }
                v14 = "foo bar"
            }
            for fi29 := 0; fi29 < 18; fi29++ {
                v14 = ("hello" + "hello")
                _ = fi29
            }
            v13 = v13
        }
    }
    v30 := GenS4{F0: -32768, F1: 2147483647, F2: 18446744073709551615}
    fmt.Println(v14, v20)
    var v31 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
