package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint64 }

func main() {
    var v1 float32 = 0.1
    var v2 uint64 = 18446744073709551615
    var v3 uint64 = (v2 | (v2 << ((65535) & 63)))
    v4 := make(map[uint32]bool)
    var v5 bool = ((false || (32767 >= 2147483647)) && (406 <= 193))
    v5 = true
    v1 = float32(3.14159)
    if (!(1 == 175)) {
        for fi6 := 0; fi6 < 3; fi6++ {
            v5 = v5
            _ = fi6
        }
        v7 := []int{}
        _ = v7
    }
    fmt.Println(v5, v1)
    v8 := []int{816, 32767, -9223372036854775808, 985}
    for fi9 := 0; fi9 < 17; fi9++ {
        v2 -= 18446744073709551615
        v10 := GenS1{F0: 312}
        v4[76] = v5
        _ = fi9
        _ = v10
    }
    var v11 int16 = 32767
    if ((v5 && v5) || (v5 && true)) {
        v3 += 41
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint32 = 1
            v14 := v4[65535]
            var v15 int8 = -25
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        if v5 {
            v4[856] = (926 > 515)
        }
    }
    var v16 int64 = 1
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    switch ("foo bar" + "a") {
    case "hello":
        v2 = (v2 + 643)
    case "foo bar":
        var v17 float64 = 0.0
        _ = v17
    case "":
        v1 = (v1 - ((3.14159 / float32(((3.14159 - ((0.0 - ((float64((v1 * v1)) * float64(-2.25)) * 0.1)) + 0.1)) / 3.14159))) / v1))
    }
    v18 := GenS2{F0: 185, F1: 42, F2: v2}
    var v19 int = -956
    var v20 uint8 = 127
    v5 = v5
    for fi21 := 0; fi21 < 14; fi21++ {
        v20 -= 236
        _ = fi21
    }
    fmt.Println(v16, v1)
    for fi22 := 0; fi22 < 4; fi22++ {
        v1 += (v1 + v1)
        for fi23 := 0; fi23 < 10; fi23++ {
            v4[2147483647] = (v5 && v5)
            _ = fi23
        }
        v2 = v2
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
