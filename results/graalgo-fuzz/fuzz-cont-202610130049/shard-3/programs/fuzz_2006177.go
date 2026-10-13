package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int; F2 string }

func main() {
    var v1 uint32 = 2147483648
    var v2 bool = (128 <= 42)
    var v3 int32 = -2147483648
    var v4 uint = 18446744073709551615
    var v5 uint16 = 1
    if (0 > (v5 >> ((uint(v1)) & 63))) {
        var v6 float64 = float64((1e10 - float32(1.5)))
        if (!v2) {
            v5 = 251
            if ((-454 != -1) && (1 < -63)) {
                if (496 <= 0) {
                    switch 21 {
                    case -2:
                        v7 := []int{}
                        _ = v7
                    case 32768:
                        var v8 uint8 = 2
                        _ = v8
                    default:
                        var v9 uint = (^(v4))
                        _ = v9
                    }
                }
            }
            fmt.Println(v3, v2)
        }
        _ = v6
    }
    var v10 int = 948
    v10 = v10
    if (v2 && true) {
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(v1, v5, fi11)
            v1 -= v1
            _ = fi11
        }
        if (v2 || true) {
            v12 := GenS1{F0: -32768, F1: v2}
            _ = v12
        }
    } else {
        var v13 int64 = 455
        fmt.Println(v2, v2, v10)
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := []int{32768, 697}
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    v1 -= (v1 % ((434) | 1))
    var v16 int8 = 127
    v17 := []int{(v10 + (-(v10))), v10}
    var v18 uint8 = 42
    v2 = true
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        if (!((v3 == v3) || ((^(v3)) == v3))) {
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 int64 = 1
                _ = fi20
                _ = v21
            }
        } else {
            v2 = (v4 <= (v4 << ((255) & 63)))
            v10 -= (^(v10))
            for fi22 := 0; fi22 < 5; fi22++ {
                v5 += v5
                v10 = (v10 & (fi22 | (fi19 & 149)))
                v23 := GenS1{F0: 1, F1: (!(455 == 4294967295))}
                _ = fi22
                _ = v23
            }
        }
        _ = fi19
    }
    v24 := GenS2{F0: v4}
    var v25 int64 = -9223372036854775808
    var v26 int = (v10 << ((4294967295) & 63))
    v1 = v1
    for fi27 := 0; fi27 < 16; fi27++ {
        if (3.14159 > (-2.25 + -2.25)) {
            v26 = (-(v26))
            if (v2 || (!v2)) {
                v17 = append(v17, 2147483647)
                var v28 int16 = 32767
                v29 := GenS3{F0: (v3 >> ((2) & 63)), F1: (v26 << ((514) & 63)), F2: "hello"}
                _ = v28
                _ = v29
            } else {
                var v30 int32 = 2147483647
                v25 = (v25 << (((^(v4))) & 63))
                _ = v30
            }
            v2 = false
        } else {
            var v31 float32 = float32(3.14159)
            _ = v31
        }
        for fi32 := 0; fi32 < 3; fi32++ {
            var v33 int8 = v16
            v25 = (v25 + (v25 ^ (v25 | int64(v4))))
            switch 0 {
            case 0:
                v10 = 9223372036854775807
            case 127:
                if (int32(v4) >= -2147483648) {
                    v3 += v3
                    switch "hello" {
                    case "foo bar":
                        v5 += (v5 / ((v5) | 1))
                    }
                    v34 := GenS2{F0: uint(v18)}
                    _ = v34
                }
            case 32768:
                v4 = 255
            default:
                v17 = append(v17, v26)
            }
            _ = fi32
            _ = v33
        }
        for fi35 := 0; fi35 < 8; fi35++ {
            fmt.Println(v4)
            _ = fi35
        }
        _ = fi27
    }
    var v36 int16 = 831
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
