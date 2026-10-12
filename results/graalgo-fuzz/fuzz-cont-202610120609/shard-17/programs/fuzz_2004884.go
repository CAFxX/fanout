package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int; F1 int32 }

func main() {
    var v1 float64 = (1.5 * ((0.0 / 1.0) + (float64(0.1) * (0.0 / 1e10))))
    var v2 int = 9223372036854775807
    var v3 uint16 = 708
    fmt.Println(v3)
    v4 := GenS1{F0: 42}
    var v5 float64 = 0.1
    switch v2 {
    case 4294967295:
        var v6 int16 = 42
        _ = v6
    case -1:
        if true {
            switch "\n\t" {
            case "foo bar":
                var v7 uint8 = 223
                _ = v7
            case "":
                var v8 float64 = -2.25
                _ = v8
            default:
                v9 := GenS2{F0: 127, F1: (v3 - (v3 & v3)), F2: uint8(v3)}
                _ = v9
            }
            var v10 float32 = 0.1
            _ = v10
        } else {
            for fi11 := 0; fi11 < 17; fi11++ {
                switch (fi11 | 0) {
                case 0:
                    if (false || ((true && (255 == 255)) && (!true))) {
                        var v12 int8 = 42
                        _ = v12
                    }
                case -692:
                    var v13 uint64 = 0
                    _ = v13
                default:
                    v14 := GenS2{F0: 127, F1: (v3 & v3), F2: 0}
                    _ = v14
                }
                if (!false) {
                    var v15 float32 = (1e10 / 3.14159)
                    fmt.Println(fi11, fi11)
                    _ = v15
                }
                v16 := GenS1{F0: -9223372036854775808}
                _ = fi11
                _ = v16
            }
        }
    default:
        var v17 uint32 = 2
        _ = v17
    }
    switch "a" {
    case "hello":
        if (true || (1 != 2)) {
            v18 := []int{0, 42, v2}
            if (!false) {
                v1 += 0.0
            }
            v1 = float64((float32(1.5) * 0.1))
            _ = v18
        } else {
            var v19 string = ("hello" + ((("foo bar" + "a") + ((("a" + "a") + "") + "a")) + "\n\t"))
            _ = v19
        }
    case "\n\t":
        v3 += (v3 << ((905) & 63))
    case "a":
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 int32 = -956
            v22 := []int{-2147483648, 42, (v2 & -173), int(v21)}
            v3 = v3
            _ = fi20
            _ = v21
            _ = v22
        }
    default:
        v5 -= 0.1
    }
    var v23 float64 = (1.5 + (v1 - (v5 / v1)))
    v24 := []int{(v2 % ((128) | 1)), 2147483647}
    var v25 uint8 = 0
    v26 := GenS3{F0: 181, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
