package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
func genFn0(p0 uint) int64 {
    switch 9223372036854775807 {
    case 127:
        v6 := []int{}
        _ = v6
    default:
        p0 -= (p0 << ((p0) & 63))
    }
    var v7 float64 = 3.14159
    v7 = (v7 - (v7 - ((v7 * v7) * v7)))
    for fi8 := 0; fi8 < 14; fi8++ {
        v7 = (((1.5 / (1e10 / (v7 / ((1e10 + 1.5) - v7)))) * v7) - 0.0)
        _ = fi8
    }
    fmt.Println(v7)
    _ = p0
    _ = v7
    return -1
}
func genFn1(p0 string) uint8 {
    p0 = p0
    p0 = "a"
    _ = p0
    return 42
}

func main() {
    var v1 float32 = 0.1
    var v2 uint16 = 2
    var v3 float64 = 3.14159
    var v4 uint8 = 2
    var v5 float64 = (float64(v1) / -2.25)
    var v9 string = ("hello" + "hello")
    v1 += (3.14159 + (float32((((v3 + v5) / v3) + ((float64(v1) * v3) / v3))) * ((v1 + v1) + 0.1)))
    var v10 bool = false
    switch 2147483647 {
    case -2147483648:
        switch "\n\t" {
        case "hello":
            var v11 string = "a"
            _ = v11
        case "a":
            v2 = (v2 | 255)
        case "\n\t":
            v9 = v9
        default:
            switch v9 {
            case "foo bar":
                fmt.Println(v2, v2)
            default:
                var v12 int16 = 42
                _ = v12
            }
        }
    default:
        v10 = false
    }
    v13 := []int{716}
    if len(v13) > 0 {
        v13[0] = 4294967295
    }
    v14 := GenS1{F0: v9, F1: 16}
    if ((!(18446744073709551615 == 65535)) || (!((-161 != 0) && (!v10)))) {
        v10 = (42 == 2)
    } else {
        v1 = ((v1 + 1e10) / v1)
        var v15 int = 4294967295
        fmt.Println(v10)
        _ = v15
    }
    v5 += v3
    var v16 string = v9
    for fi17 := 0; fi17 < 17; fi17++ {
        if len(v13) > 0 {
            v13[0] = fi17
        }
        _ = fi17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 uint32 = uint32(v2)
        for fi20 := 0; fi20 < 20; fi20++ {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            v9 = (v16 + (v9 + (v9 + "hello")))
            _ = fi21
        }
        _ = fi18
        _ = v19
    }
    var v22 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
