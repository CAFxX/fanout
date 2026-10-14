package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint }
type GenS4 struct { F0 int8; F1 bool }
func genFn0(p0 uint, p1 int64) int64 {
    if true {
        var v5 uint16 = 255
        var v6 int32 = 42
        if false {
            for fi7 := 0; fi7 < 16; fi7++ {
                fmt.Println(v6, fi7, v6)
                var v8 int16 = -32768
                v9 := GenS1{F0: false}
                _ = fi7
                _ = v8
                _ = v9
            }
            if true {
                p0 = (p0 / ((p0) | 1))
            } else {
                v10 := []int{32767, 9223372036854775807}
                _ = v10
            }
        }
        _ = v5
        _ = v6
    } else {
        p1 += 84
    }
    if false {
        fmt.Println(p0)
        v11 := GenS2{F0: 746, F1: 2}
        _ = v11
    } else {
        p1 = (p1 + 621)
        if (291 != 770) {
            if (!(false && (42 >= 2147483647))) {
                v12 := GenS2{F0: 76, F1: 255}
                _ = v12
            } else {
                switch 0 {
                case 9223372036854775807:
                    switch 32768 {
                    case 528:
                        v13 := GenS1{F0: (42 >= 42)}
                        _ = v13
                    case -2147483648:
                        v14 := []int{-579, 32768, 724}
                        _ = v14
                    case 149:
                        var v15 int64 = 1
                        _ = v15
                    }
                case -505:
                    p1 = (p1 + (p1 | -1))
                case 2147483647:
                    p1 = (p1 << ((p0) & 63))
                default:
                    v16 := GenS2{F0: 42, F1: 41}
                    _ = v16
                }
                if false {
                    p1 -= -1
                }
            }
        }
    }
    p1 += (p1 ^ (^(p1)))
    var v17 int64 = -1
    v18 := []int{-9223372036854775808, 420}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return (v17 - 307)
}
func genFn1() uint {
    var v19 float32 = float32(float64(0.0))
    if (111 == 28) {
        v20 := []int{32768, 722, 490}
        fmt.Println(v19)
        _ = v20
    } else {
        v21 := []int{4294967295}
        v19 = (((v19 + 0.0) * float32((1.5 + (1.5 + 1.5)))) / (v19 - float32(0.1)))
        _ = v21
    }
    _ = v19
    return 2
}

func main() {
    v1 := []string{((("" + "hello") + "\n\t") + "a"), ("a" + "hello"), "hello"}
    var v2 uint8 = 255
    var v3 uint = 0
    v4 := make(map[uint32]uint32)
    fmt.Println(v2)
    if (42 != 1) {
        if true {
            switch "a" {
            case "\n\t":
                var v22 uint8 = v2
                _ = v22
            }
        } else {
            v2 -= v2
            var v23 int = -2147483648
            switch (-(v23)) {
            case 2147483647:
                var v24 float64 = (0.0 / -2.25)
                _ = v24
            case 32767:
                var v25 int16 = 227
                _ = v25
            case -511:
                if true {
                    v3 = (^(v3))
                    fmt.Println(v23, v3, v3)
                }
            }
            _ = v23
        }
        var v26 uint64 = 18446744073709551615
        _ = v26
    } else {
        fmt.Println(v2, v3)
        for fi27 := 0; fi27 < 19; fi27++ {
            v3 -= v3
            _ = fi27
        }
    }
    var v28 uint16 = 0
    var v29 uint16 = 0
    var v30 uint16 = 0
    v31 := GenS3{F0: 4294967295, F1: 2, F2: genFn1()}
    v30 -= (v30 | (v29 - v30))
    v32 := GenS4{F0: -24, F1: false}
    v33 := GenS3{F0: 4294967295, F1: (v3 & (v3 % (((^(v3))) | 1))), F2: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
