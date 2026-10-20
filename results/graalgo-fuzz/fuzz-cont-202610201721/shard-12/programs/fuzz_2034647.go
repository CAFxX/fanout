package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 float64 = ((((3.14159 / 0.1) + (1e10 + (-2.25 / (0.0 + 3.14159)))) + 3.14159) / 1.5)
    v4 := []int{128, 42, -9223372036854775808, -1}
    var v5 uint8 = 0
    switch 32767 {
    case 127:
        if false {
            var v6 uint32 = 42
            _ = v6
        } else {
            v5 = v5
            if (639 > 42) {
                var v7 float32 = -2.25
                v1[127] = 42
                for fi8 := 0; fi8 < 5; fi8++ {
                    v9 := GenS1{F0: 65535, F1: (v5 | v5)}
                    _ = fi8
                    _ = v9
                }
                _ = v7
            }
        }
    case 42:
        if (0.0 <= v3) {
            var v10 uint32 = 2
            _ = v10
        } else {
            var v11 float64 = 1.5
            v5 = v5
            _ = v11
        }
    case -2147483648:
        v5 = v5
    default:
        fmt.Println(v3)
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        if (!false) {
            v13 := GenS1{F0: 2, F1: v5}
            _ = v13
        }
        _ = fi12
    }
    var v14 uint32 = 412
    var v15 int32 = -309
    var v16 int = -142
    v1[(v16 << ((1) & 63))] = (v14 + v14)
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := GenS2{F0: 451}
        if (!(((fi17 < (-(fi17))) || true) || false)) {
            v14 = 65535
            if (!("a" > "a")) {
                var v19 uint = 633
                v20 := GenS1{F0: 32767, F1: 0}
                for fi21 := 0; fi21 < 13; fi21++ {
                    if true {
                        switch (fi21 * (^(fi17))) {
                        case 9223372036854775807:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v22 := v2[fi21]
                            _ = v22
                        case 32768:
                            var v23 int = fi17
                            _ = v23
                        case 42:
                            var v24 int32 = 42
                            _ = v24
                        default:
                            if true {
                                var v25 bool = (42 != fi21)
                                fmt.Println(fi21, v5, v3)
                                _ = v25
                            } else {
                                v3 += ((v3 / ((-2.25 * 1e10) * 3.14159)) + float64(((3.14159 + 1.5) / 1.5)))
                            }
                        }
                        switch v16 {
                        case 47:
                            var v26 uint32 = v14
                            _ = v26
                        case 718:
                            var v27 int32 = (v15 >> ((uint(v5)) & 63))
                            _ = v27
                        case -2147483648:
                            v19 -= v19
                        }
                    }
                    _ = fi21
                }
                _ = v19
                _ = v20
            } else {
                switch v16 {
                case 1:
                    fmt.Println(v16, v3, fi17)
                case 9223372036854775807:
                    fmt.Println(v16, fi17, fi17)
                case 32768:
                    v2[(v16 % ((-612) | 1))] = false
                }
                delete(v2, v16)
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            v14 = v14
            fmt.Println(v3, v3, v3)
            v29 := GenS2{F0: 250}
            _ = fi28
            _ = v29
        }
        _ = fi17
        _ = v18
    }
    v3 = v3
    var v30 int16 = 42
    var v31 int32 = int32(v14)
    var v32 int = (v16 >> ((18446744073709551615) & 63))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
