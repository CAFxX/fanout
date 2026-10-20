package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    var v1 float64 = ((1e10 / 0.1) - (1.5 * ((float64(1e10) / ((0.1 / float64(1e10)) * -2.25)) + 0.1)))
    v2 := make(map[string]string)
    var v3 uint64 = 2
    var v4 uint32 = 65535
    v5 := make(map[uint32]uint32)
    var v6 int = -9223372036854775808
    switch v6 {
    case 32767:
        for fi7 := 0; fi7 < 13; fi7++ {
            switch fi7 {
            case 32767:
                v6 += (v6 - (v6 ^ v6))
            case 128:
                for fi8 := 0; fi8 < 7; fi8++ {
                    fmt.Println(v3, fi8)
                    _ = fi8
                }
            case 32768:
                fmt.Println(v6, fi7)
            }
            _ = fi7
        }
    case -949:
        var v9 string = ""
        _ = v9
    case 42:
        var v10 bool = (((0.0 + 0.1) + (v1 - (float64((-2.25 + ((float32(v1) / float32((((float64(3.14159) - (v1 * v1)) * (float64(0.0) + (1.5 * (float64((0.1 * ((1.5 + ((float32(1e10) + -2.25) + 1e10)) + 1e10))) / 1.5)))) - 0.1))) - 1.5))) / v1))) < v1)
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        fmt.Println(v3)
        v4 -= v4
        _ = fi11
    }
    v1 = (float64(float32(float64(1.5))) - (v1 + 3.14159))
    fmt.Println(v4, v1, v3)
    v12 := GenS1{F0: -1, F1: 42}
    fmt.Println(v3)
    fmt.Println(v1, v4)
    fmt.Println(v4, v3, v6)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint16 = 42
        if ((true || (!("foo bar" <= "foo bar"))) && false) {
            v4 += (v4 / ((v4) | 1))
            v15 := []int{int(v3)}
            fmt.Println(v14)
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
