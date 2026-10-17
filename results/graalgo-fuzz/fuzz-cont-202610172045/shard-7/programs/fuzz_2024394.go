package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint64, p1 int8) string {
    var v4 uint = 4294967295
    if (!(9223372036854775807 < 483)) {
        p0 += p0
    } else {
        v5 := []int{32768}
        _ = v5
    }
    p1 = -128
    if (79 != 221) {
        switch 793 {
        case 9223372036854775807:
            p0 = 460
        default:
            if (true || true) {
                var v6 int32 = 2147483647
                _ = v6
            } else {
                v7 := GenS1{F0: 128, F1: -1, F2: 1}
                v8 := []int{-2147483648}
                var v9 float32 = 1e10
                _ = v7
                _ = v8
                _ = v9
            }
        }
    }
    fmt.Println(v4)
    _ = p0
    _ = p1
    _ = v4
    return "foo bar"
}
func genFn1(p0 int64, p1 bool) uint8 {
    v10 := GenS2{F0: 901}
    p0 = (p0 ^ (p0 & 42))
    var v11 float64 = 1.5
    v12 := []int{-615, 32768}
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(p1)
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]uint32)
    var v3 int64 = -960
    v1 -= 1.5
    var v14 float32 = (1.5 + 3.14159)
    _ = len(v2)
    fmt.Println(v14)
    for fi15 := 0; fi15 < 8; fi15++ {
        v2[350] = 2147483647
        _ = fi15
    }
    delete(v2, 2147483647)
    var v16 int = 965
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS2{F0: 18446744073709551615}
        for fi19 := 0; fi19 < 14; fi19++ {
            switch (fi17 << ((0) & 63)) {
            case -86:
                fmt.Println(fi17)
            }
            _ = fi19
        }
        if (v1 != 0.1) {
            var v20 float32 = 1e10
            for fi21 := 0; fi21 < 19; fi21++ {
                v22 := GenS3{F0: 139, F1: 0}
                v16 = (fi17 << ((0) & 63))
                v1 = (v1 - (v1 - (0.0 * v1)))
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            delete(v2, 42)
            var v23 float64 = ((((v1 + 3.14159) * v1) + (3.14159 / v1)) * (v1 + v1))
            _ = v23
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v16)
    var v24 uint = 65535
    for fi25 := 0; fi25 < 9; fi25++ {
        if (!(-2147483648 > 0)) {
            v1 -= v1
            if (true && false) {
                if true {
                    var v26 int64 = 42
                    for fi27 := 0; fi27 < 14; fi27++ {
                        if ((fi27 > (v16 * fi25)) || ((v16 << ((v24) & 63)) > (fi27 | (v16 ^ (fi25 - 4294967295))))) {
                            fmt.Println(v3, v1)
                            v28 := GenS4{F0: 0}
                            _ = v28
                        } else {
                            switch "\n\t" {
                            case "":
                                var v29 bool = (true && true)
                                _ = v29
                            case "hello":
                                var v30 int64 = v3
                                _ = v30
                            default:
                                v3 -= (v26 + v3)
                            }
                            fmt.Println(v26, v14, v24)
                        }
                        _ = fi27
                    }
                    switch 127 {
                    case 9223372036854775807:
                        v31 := []int{v16, (v16 % ((32768) | 1))}
                        _ = v31
                    }
                    _ = v26
                }
                for fi32 := 0; fi32 < 7; fi32++ {
                    if (548 >= 18446744073709551615) {
                        v33 := GenS1{F0: fi32, F1: 37, F2: -2147483648}
                        if false {
                            v16 = 128
                        } else {
                            var v34 int16 = 0
                            _ = v34
                        }
                        var v35 int = (fi32 + -33)
                        v36 := GenS1{F0: (fi32 | 42), F1: -126, F2: -543}
                        _ = v33
                        _ = v35
                        _ = v36
                    }
                    v37 := []int{(fi32 / ((fi25) | 1)), (v16 * (fi25 + (fi32 >> (((v24 >> (((v24 % ((223) | 1))) & 63))) & 63)))), 41}
                    _ = fi32
                    _ = v37
                }
            } else {
                v38 := GenS3{F0: (v24 - 855), F1: 127}
                switch (v16 | -2147483648) {
                case -126:
                    var v39 uint64 = 1
                    _ = v39
                case 2147483647:
                    v40 := GenS4{F0: 2147483648}
                    _ = v40
                case 32768:
                    fmt.Println(v3, v1)
                }
                _ = v38
            }
        } else {
            fmt.Println(v16)
            delete(v2, 42)
            v24 = v24
        }
        _ = fi25
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        for fi42 := 0; fi42 < 6; fi42++ {
            v43 := GenS5{F0: v24, F1: 127}
            _ = fi42
            _ = v43
        }
        var v44 float32 = v14
        _ = fi41
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v24
    _ = v3
}
