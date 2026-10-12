package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int }

func main() {
    var v1 int8 = 42
    v2 := make(map[int]int)
    var v3 float32 = 0.0
    var v4 int32 = 2147483647
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int16 = -32768
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 string = ""
        _ = fi7
        _ = v8
    }
    v3 += (float32(1e10) / v3)
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 float64 = -2.25
        _ = fi9
        _ = v10
    }
    if (-645 > 838) {
        v3 += float32(1e10)
        v4 -= (-(v4))
    }
    if (true || (!(1e10 >= 1e10))) {
        v11 := GenS1{F0: -359}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := v2[fi12]
            v3 = (1e10 / ((3.14159 * float32((3.14159 - 0.0))) + v3))
            _ = fi12
            _ = v13
        }
        v1 = (v1 + (v1 % ((v1) | 1)))
        if (!true) {
            for fi14 := 0; fi14 < 17; fi14++ {
                if (false && false) {
                    var v15 string = "foo bar"
                    fmt.Println(v1, v4, v1)
                    v4 -= (v4 << ((2) & 63))
                    _ = v15
                }
                fmt.Println(v1, v4, fi14)
                _ = fi14
            }
        }
    }
    v4 = v4
    for fi16 := 0; fi16 < 6; fi16++ {
        v1 -= (v1 >> ((4294967295) & 63))
        fmt.Println(v4, v3, v4)
        _ = fi16
    }
    if ((!((-2.25 < 1e10) || (100 > (-(v1))))) || ((42 < 1) || (!(0 > (v4 & (v4 % ((-1) | 1))))))) {
        v3 = (v3 * (3.14159 + v3))
        _ = len(v2)
        if (int64(v1) == 9223372036854775807) {
            var v17 int16 = 32767
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            if false {
                v4 = (v4 | v4)
            } else {
                var v19 int8 = v1
                for fi20 := 0; fi20 < 20; fi20++ {
                    v3 = (-2.25 + v3)
                    _ = fi20
                }
                _ = v19
            }
            switch "hello" {
            case "foo bar":
                var v21 int8 = v1
                _ = v21
            case "\n\t":
                switch "hello" {
                case "\n\t":
                    v22 := GenS2{F0: -1, F1: 0, F2: (fi18 + (fi18 & fi18))}
                    _ = v22
                default:
                    fmt.Println(v4, v4)
                }
            case "":
                v23 := []int{(^(fi18))}
                _ = v23
            }
            delete(v2, 0)
            _ = fi18
        }
        var v24 uint32 = 42
        _ = v24
    }
    var v25 int16 = 42
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 uint8 = 255
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
