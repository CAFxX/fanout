package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int, p1 uint8, p2 int32) uint16 {
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 uint64 = 647
        if (!true) {
            v6 := []int{(fi4 + p0), (fi4 + int(p2))}
            p1 -= (p1 & 98)
            _ = v6
        } else {
            for fi7 := 0; fi7 < 2; fi7++ {
                v8 := []int{9223372036854775807, (p0 + fi4), (p0 * (p0 + (fi7 ^ fi4))), p0}
                fmt.Println(p1)
                _ = fi7
                _ = v8
            }
            var v9 uint32 = 0
            _ = v9
        }
        var v10 uint64 = v5
        _ = fi4
        _ = v10
        _ = v5
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        switch "foo bar" {
        case "a":
            var v12 int = fi11
            _ = v12
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 uint32 = 951
    var v2 float64 = 0.0
    var v3 int16 = -701
    for fi13 := 0; fi13 < 6; fi13++ {
        v3 -= (v3 - v3)
        _ = fi13
    }
    v1 = (v1 % ((2) | 1))
    var v14 uint8 = 1
    var v15 int32 = -103
    if (!(1 < -340)) {
        v3 += v3
        v15 += v15
        v16 := GenS1{F0: v15, F1: -252, F2: (v14 & 127)}
        _ = v16
    } else {
        var v17 float32 = ((float32(1e10) + float32((v2 * v2))) - float32(0.1))
        if ((v14 ^ v14) <= 42) {
            var v18 int16 = (v3 & (v3 & -481))
            fmt.Println(v18, v14)
            _ = v18
        } else {
            var v19 uint32 = (v1 >> ((775) & 63))
            switch ("hello" + ((("\n\t" + "") + "\n\t") + "a")) {
            case "hello":
                var v20 bool = true
                _ = v20
            }
            _ = v19
        }
        _ = v17
    }
    v21 := []int{-565, -1, -543, -9223372036854775808}
    v2 += (((1.5 + (v2 + (((float64(-2.25) / (1e10 * v2)) - v2) + (v2 + float64((float32(3.14159) / (0.0 + (float32(v2) * float32(3.14159))))))))) * v2) / -2.25)
    v3 = v3
    fmt.Println(v1, v15)
    if ((0.1 >= (float32(v2) + -2.25)) || true) {
        switch 42 {
        case 1:
            var v22 string = ("\n\t" + "hello")
            _ = v22
        default:
            fmt.Println(v15)
        }
        v23 := GenS2{F0: (v15 ^ v15)}
        if len(v21) > 0 {
            v21[0] = -9223372036854775808
        }
        _ = v23
    } else {
        fmt.Println(v14, v15)
        v24 := GenS2{F0: 42}
        _ = v24
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
}
