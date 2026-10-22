package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 uint }
type GenS3 struct { F0 string }

func main() {
    var v1 string = "a"
    var v2 bool = true
    v3 := make(map[int]bool)
    var v4 uint8 = 255
    var v5 float64 = 1.5
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 float64 = v5
        for fi8 := 0; fi8 < 19; fi8++ {
            v4 += v4
            fmt.Println(v4, fi8, v4)
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        switch (fi9 & fi9) {
        case 9223372036854775807:
            var v10 int8 = 42
            _ = v10
        case 399:
            v5 += v5
        }
        if (1 >= 65535) {
            if (!(v1 <= "")) {
                v2 = v2
                v5 = v5
            } else {
                v5 -= v5
                _ = len(v3)
            }
            v11 := GenS1{F0: 850}
            v5 = float64(0.1)
            _ = v11
        }
        fmt.Println(v1)
        _ = fi9
    }
    v12 := GenS1{F0: 18446744073709551615}
    v1 = ("hello" + v1)
    v4 -= (^(v4))
    fmt.Println(v5, v2, v2)
    var v13 string = "\n\t"
    fmt.Println(v2)
    switch v1 {
    case "foo bar":
        v5 = -2.25
    case "hello":
        v5 = 0.1
    }
    if false {
        v3[0] = (!(0 != 2))
        switch "foo bar" {
        case "a":
            switch 42 {
            case 32768:
                switch 240 {
                case 4294967295:
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v2 = (!true)
                        v3[(fi14 | fi14)] = v2
                        fmt.Println(v1, v13, fi14)
                        _ = fi14
                    }
                }
            default:
                fmt.Println(v2, v2, v5)
            }
        }
    }
    var v15 uint64 = 18446744073709551615
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(v15, v15)
        v17 := GenS2{F0: (v15 & (v15 % ((4294967295) | 1))), F1: 4294967295, F2: 4294967295}
        v4 -= v4
        _ = fi16
        _ = v17
    }
    v18 := GenS3{F0: ("a" + v13)}
    switch -9223372036854775808 {
    case -2147483648:
        v3[-98] = (v2 || (887 == uint32(v4)))
    case 32767:
        v1 = v13
    case -901:
        fmt.Println(v2, v15)
    default:
        v19 := []int{0, 128, 42}
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
