package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
func genFn0(p0 int32, p1 uint32) int32 {
    v5 := GenS1{F0: 127, F1: (p0 - (p0 >> ((1) & 63)))}
    for fi6 := 0; fi6 < 1; fi6++ {
        p1 += p1
        var v7 float32 = -2.25
        _ = fi6
        _ = v7
    }
    var v8 float64 = ((((float64((0.0 * 3.14159)) * 0.0) - -2.25) / -2.25) - 3.14159)
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    _ = v9
    return (^(p0))
}

func main() {
    var v1 bool = (18446744073709551615 <= 136)
    var v2 int16 = -32768
    v3 := make(map[int]string)
    var v4 int8 = 127
    v1 = v1
    v1 = (2 > 112)
    v1 = v1
    v3[4294967295] = "foo bar"
    switch ("hello" + "a") {
    case "hello":
        switch -941 {
        case -9223372036854775808:
            v10 := GenS1{F0: (^(v4)), F1: -2147483648}
            _ = v10
        }
    default:
        v11 := []int{1, -2147483648}
        _ = v11
    }
    v1 = v1
    v4 += (v4 >> ((255) & 63))
    if (-123 > 396) {
        _ = len(v3)
        delete(v3, 4294967295)
        v12 := []int{127}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int16 = -667
            v2 = 42
            var v15 uint8 = 199
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := []int{42}
        _ = v16
    }
    v17 := []int{32768}
    var v18 uint64 = 4294967295
    v1 = v1
    v19 := GenS2{F0: 42, F1: 253, F2: 2}
    switch -2147483648 {
    case 0:
        var v20 int8 = v4
        _ = v20
    case -30:
        var v21 int8 = 1
        _ = v21
    }
    _ = len(v3)
    switch (("a" + "a") + "a") {
    case "\n\t":
        if (("" != ("\n\t" + "\n\t")) && ((1.5 * ((1.5 * (1e10 / (float32(-2.25) * 1.5))) - (0.1 / 1.0))) > ((float32((3.14159 - (3.14159 + (3.14159 + float64(1.5))))) - (0.1 / 0.1)) * 0.1))) {
            var v22 uint32 = 530
            _ = v22
        }
    case "a":
        v4 -= (v4 >> ((65535) & 63))
    case "":
        v23 := GenS1{F0: (v4 + (-(v4))), F1: -2147483648}
        _ = v23
    default:
        var v24 uint32 = 2147483647
        _ = v24
    }
    v4 = (v4 & v4)
    var v25 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
