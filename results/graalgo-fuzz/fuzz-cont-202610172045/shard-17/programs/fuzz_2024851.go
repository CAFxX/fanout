package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint8; F1 uint16 }
func genFn0(p0 bool) int32 {
    fmt.Println(p0)
    v4 := []int{42, 32}
    _ = p0
    _ = v4
    return 0
}
func genFn1(p0 uint64, p1 int, p2 int16) uint64 {
    p2 = 1000
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(fi5)
        switch "" {
        case "":
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 int64 = -244
                _ = fi6
                _ = v7
            }
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := []int{}
    var v2 uint32 = 364
    var v3 int = 127
    v2 += (^(v2))
    var v8 uint8 = 255
    fmt.Println(v3)
    var v9 string = "\n\t"
    var v10 uint64 = 24
    for fi11 := 0; fi11 < 16; fi11++ {
        v3 -= fi11
        if len(v1) > 0 {
            v1[0] = (-(v3))
        }
        var v12 int32 = -2147483648
        _ = fi11
        _ = v12
    }
    v2 -= v2
    v13 := GenS1{F0: 2, F1: (v2 << ((65535) & 63))}
    var v14 int = v3
    var v15 uint16 = uint16(v10)
    v2 = v2
    fmt.Println(v10, v9, v2)
    var v16 uint8 = (v8 * (v8 ^ 128))
    switch (v14 & (v3 ^ v14)) {
    case -573:
        v17 := GenS2{F0: (v8 & v16), F1: v2, F2: 353}
        _ = v17
    case 32768:
        if false {
            if (-1 < 0) {
                switch v9 {
                case "":
                    var v18 float64 = 1.5
                    _ = v18
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi19 := 0; fi19 < 16; fi19++ {
                    fmt.Println(v9, v3)
                    v20 := GenS3{F0: -1, F1: 0}
                    fmt.Println(v8, v3)
                    _ = fi19
                    _ = v20
                }
            }
            if (-1 < -1) {
                var v21 float64 = 1.5
                _ = v21
            } else {
                for fi22 := 0; fi22 < 11; fi22++ {
                    _ = len(v1)
                    _ = fi22
                }
                var v23 uint16 = (v15 & (v15 ^ v15))
                var v24 uint8 = (^(v8))
                _ = v23
                _ = v24
            }
        }
    case 277:
        if (true || true) {
            switch -561 {
            case 128:
                var v25 uint32 = 2
                _ = v25
            case -1:
                v26 := GenS1{F0: uint(v2), F1: (v2 >> ((2) & 63))}
                _ = v26
            }
            v8 = (v8 << ((243) & 63))
            var v27 uint8 = v16
            _ = v27
        }
    default:
        v1 = append(v1, v14)
    }
    if false {
        v28 := GenS3{F0: 0, F1: -9223372036854775808}
        switch -1 {
        case -2147483648:
            for fi29 := 0; fi29 < 9; fi29++ {
                switch 128 {
                case 32768:
                    v3 = 0
                case 128:
                    if (!((("hello" + v9) + v9) > v9)) {
                        fmt.Println(v9)
                        var v30 float32 = 3.14159
                        _ = v30
                    }
                default:
                    var v31 int64 = 42
                    _ = v31
                }
                v3 -= (fi29 - (^(fi29)))
                _ = fi29
            }
        case 207:
            v16 = v8
        }
        v32 := GenS4{F0: 42, F1: (v15 | (v15 >> ((255) & 63)))}
        _ = v28
        _ = v32
    }
    var v33 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v8
    _ = v9
}
