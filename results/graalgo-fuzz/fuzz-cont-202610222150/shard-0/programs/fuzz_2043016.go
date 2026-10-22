package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int }
type GenS2 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 float32 = -2.25
    var v2 float64 = 0.0
    v3 := []int{119, -362, 9223372036854775807, 2147483647}
    if true {
        v2 += ((v2 / v2) * (v2 - v2))
    }
    var v4 int = 1
    v1 += (((0.0 - (v1 + (3.14159 + (-2.25 - (0.0 / v1))))) * v1) * v1)
    v1 = 3.14159
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint16 = 65535
            var v8 uint16 = 32767
            if false {
                v8 -= 862
            } else {
                v7 += v8
                v9 := GenS1{F0: "foo bar", F1: 42, F2: v4}
                _ = v9
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        fmt.Println(fi5, v2, fi5)
        _ = fi5
    }
    if false {
        if (!(846 == 65535)) {
            for fi10 := 0; fi10 < 19; fi10++ {
                v3 = append(v3, (fi10 & v4))
                var v11 uint16 = 165
                _ = fi10
                _ = v11
            }
            v4 = (v4 * v4)
        }
    } else {
        if true {
            v3 = append(v3, v4)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v12 int8 = 42
            _ = v12
        } else {
            var v13 int = v4
            if false {
                v1 += (v1 + (v1 / v1))
                v2 = (v2 / ((float64((0.0 - float32(v2))) + (3.14159 - v2)) / 1e10))
                var v14 uint16 = 0
                _ = v14
            }
            _ = v13
        }
    }
    var v15 uint32 = 42
    _ = len(v3)
    v15 -= v15
    v16 := GenS2{F0: 142, F1: (v15 << ((2) & 63))}
    var v17 uint32 = (v15 << ((65535) & 63))
    v17 = v15
    var v18 uint32 = v15
    switch -2147483648 {
    case 247:
        v17 += (^(v15))
    case 9223372036854775807:
        if ((v4 & v4) < (v4 ^ v4)) {
            v2 = float64(((3.14159 * (1.5 * 1e10)) + v1))
            var v19 uint = 717
            _ = v19
        }
    }
    var v20 int16 = 0
    for fi21 := 0; fi21 < 4; fi21++ {
        v1 = (-2.25 * v1)
        for fi22 := 0; fi22 < 19; fi22++ {
            if len(v3) > 0 {
                v3[0] = v4
            }
            fmt.Println(v4, v1, fi22)
            fmt.Println(v18, v17, v15)
            _ = fi22
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
