package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8, p1 int16, p2 uint64) int32 {
    var v7 bool = true
    if true {
        for fi8 := 0; fi8 < 2; fi8++ {
            switch "" {
            case "":
                p1 = p1
            default:
                v9 := GenS1{F0: -845}
                _ = v9
            }
            var v10 float32 = (1.5 - (((1.5 + -2.25) * (3.14159 - -2.25)) / 1.0))
            _ = fi8
            _ = v10
        }
    } else {
        v7 = (!(-1 > -9223372036854775808))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint32 = 0
    var v2 float32 = 1.5
    var v3 uint64 = 18446744073709551615
    var v4 int64 = 1
    var v5 uint16 = 42
    var v6 int = 127
    if (668 != 1) {
        var v11 int16 = 1
        v12 := []int{v6, (v6 | (v6 << ((622) & 63))), (v6 | 1)}
        _ = v11
        _ = v12
    }
    var v13 bool = (-2.25 != ((v2 / v2) / v2))
    v13 = true
    switch "hello" {
    case "\n\t":
        for fi14 := 0; fi14 < 4; fi14++ {
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 uint = 255
                var v17 float32 = 0.1
                v18 := []int{fi14, fi14, -499}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            fmt.Println(v2)
            _ = fi14
        }
    default:
        v19 := []int{2147483647, 127, v6}
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 3; fi21++ {
            fmt.Println(fi21, v1)
            _ = fi21
        }
        if false {
            if (false || v13) {
                v4 += v4
                v4 += (^(v4))
                var v22 bool = v13
                _ = v22
            } else {
                var v23 string = ("\n\t" + "")
                _ = v23
            }
        }
        _ = fi20
    }
    v13 = ((-22 < 23) || (v1 >= v1))
    v24 := []int{v6}
    var v25 uint16 = (v5 << ((739) & 63))
    switch "hello" {
    case "":
        switch v6 {
        case -2147483648:
            v3 -= (v3 << ((255) & 63))
        case 0:
            fmt.Println(v5, v6)
        case 4294967295:
            if (v2 != float32(((-2.25 / ((1e10 / ((1.5 / float64(v2)) / ((((1e10 + (3.14159 + 1e10)) / 1.0) * 1e10) + (1e10 / float64(v2))))) / 0.1)) - -2.25))) {
                v26 := GenS2{F0: 2}
                v4 -= v4
                _ = v26
            }
        default:
            var v27 uint = 255
            _ = v27
        }
    case "foo bar":
        if v13 {
            var v28 bool = v13
            fmt.Println(v5, v28, v25)
            v29 := GenS1{F0: 42}
            _ = v28
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
