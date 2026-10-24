package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int16 }

func main() {
    v1 := []int{-9223372036854775808, 0, 0}
    v2 := make(map[int]uint32)
    var v3 bool = true
    if v3 {
        for fi4 := 0; fi4 < 13; fi4++ {
            var v5 uint16 = 1
            if (793 < 1) {
                var v6 int32 = 295
                _ = v6
            }
            _ = fi4
            _ = v5
        }
        for fi7 := 0; fi7 < 18; fi7++ {
            v1 = append(v1, (fi7 & 128))
            for fi8 := 0; fi8 < 9; fi8++ {
                v3 = (!(((0 <= 255) && (1.5 <= (1.5 / 1.0))) && (!((0.1 <= 0.0) || (!(65535 > 2))))))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi8
            }
            _ = fi7
        }
        fmt.Println(v3)
    }
    var v9 uint32 = 217
    if (895 >= 65535) {
        if v3 {
            var v10 int8 = 42
            _ = v10
        }
    } else {
        delete(v2, 4294967295)
        v11 := GenS1{F0: 9223372036854775807, F1: -873}
        switch 4294967295 {
        case -945:
            fmt.Println(v9, v3)
        case 127:
            var v12 bool = v3
            _ = v12
        case 42:
            if v3 {
                var v13 string = "foo bar"
                v3 = (!(!((float32(-2.25) - 0.0) >= 0.1)))
                _ = v13
            } else {
                v14 := GenS2{F0: 0}
                _ = v14
            }
        }
        _ = v11
    }
    var v15 bool = (v3 && v3)
    if len(v1) > 0 {
        v1[0] = -328
    }
    if ((!(1 >= 32)) && (!false)) {
        v16 := v2[127]
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 int = (fi17 - fi17)
        v9 -= (v9 << ((0) & 63))
        _ = fi17
        _ = v18
    }
    fmt.Println(v3, v9)
    if false {
        if v15 {
            v19 := GenS3{F0: (0.1 >= ((-2.25 + float64(0.0)) / -2.25))}
            v3 = (!true)
            switch "hello" {
            case "\n\t":
                v3 = (9223372036854775807 != -1)
            case "hello":
                _ = len(v1)
            case "a":
                v20 := GenS3{F0: false}
                _ = v20
            default:
                var v21 float32 = 1.5
                _ = v21
            }
            _ = v19
        } else {
            switch -435 {
            case 1:
                v22 := GenS4{F0: (949 >= 4294967295), F1: -1}
                _ = v22
            }
            v2[969] = (v9 / ((v9) | 1))
        }
        v15 = false
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
