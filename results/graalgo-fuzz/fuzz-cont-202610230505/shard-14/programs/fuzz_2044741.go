package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint16 = 962
    var v2 uint8 = 42
    var v3 int64 = 47
    var v4 bool = (-512 >= -192)
    var v5 int64 = (v3 >> ((681) & 63))
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            if false {
                if v4 {
                    v3 -= (v5 % (((v3 ^ (v3 | v3))) | 1))
                } else {
                    var v9 uint32 = 42
                    fmt.Println(v5, fi7, v1)
                    v4 = (18 > 20)
                    _ = v9
                }
                v10 := v6[fi8]
                _ = v10
            } else {
                v1 += v1
                v6[(^(fi7))] = ""
                switch fi8 {
                case 4294967295:
                    v11 := GenS1{F0: v2}
                    _ = v11
                case 820:
                    v4 = v4
                case -780:
                    v12 := []int{int(v1), (fi7 / ((fi8) | 1))}
                    _ = v12
                }
            }
            _ = fi8
        }
        _ = fi7
    }
    v5 = -675
    if v4 {
        v3 -= 1
        v2 -= v2
    }
    var v13 int32 = 736
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 int64 = 1
        switch "hello" {
        case "\n\t":
            fmt.Println(fi14, v5, v13)
        case "a":
            for fi16 := 0; fi16 < 16; fi16++ {
                v15 += v3
                fmt.Println(fi14, v5, v13)
                _ = fi16
            }
        case "":
            v17 := []int{42, fi14, fi14}
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    if v4 {
        var v18 float64 = -2.25
        v19 := GenS1{F0: v2}
        _ = v18
        _ = v19
    }
    v4 = v4
    if v4 {
        for fi20 := 0; fi20 < 19; fi20++ {
            fmt.Println(fi20, v4)
            _ = fi20
        }
        switch 127 {
        case -546:
            v4 = (0 == -2147483648)
        }
    } else {
        for fi21 := 0; fi21 < 16; fi21++ {
            v4 = v4
            _ = fi21
        }
    }
    var v22 float32 = (((((0.0 / 1.0) / 3.14159) - (-2.25 - 0.0)) * 0.0) + 1e10)
    var v23 int32 = int32(v5)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
