package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint32, p2 uint16) uint16 {
    v4 := []int{-1}
    for fi5 := 0; fi5 < 7; fi5++ {
        if ((false || ((3.14159 / (1e10 + -2.25)) == (-2.25 + 3.14159))) && (float64(1e10) <= 3.14159)) {
            fmt.Println(p1, p1)
            var v6 int8 = -70
            fmt.Println(fi5)
            _ = v6
        } else {
            if len(v4) > 0 {
                v4[0] = (fi5 * (fi5 << ((65535) & 63)))
            }
            v7 := GenS1{F0: 2147483647}
            var v8 float64 = (0.0 * ((0.1 + -2.25) + -2.25))
            _ = v7
            _ = v8
        }
        fmt.Println(p1, p2)
        _ = fi5
    }
    p2 += (^(p2))
    p1 = (p1 - 500)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}
func genFn1() bool {
    var v9 uint = 701
    v9 = 42
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 string = "foo bar"
        _ = fi10
        _ = v11
    }
    v12 := []int{32767, 4294967295, -2147483648}
    switch 32768 {
    case -2147483648:
        v9 = (v9 | (v9 % ((v9) | 1)))
    case -425:
        v13 := GenS1{F0: 46}
        _ = v13
    case 207:
        fmt.Println(v9)
    }
    _ = v12
    _ = v9
    return false
}

func main() {
    v1 := []uint32{922, 4294967295, 699}
    v2 := make(map[int]string)
    var v3 int32 = 0
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 uint16 = 32767
        var v16 float32 = ((0.1 - 1e10) + -2.25)
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v3 = v3
        var v18 uint8 = 127
        if (0 == (v18 | (^(v18)))) {
            v3 = v3
        } else {
            var v19 int64 = 42
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    v20 := GenS1{F0: (v3 << ((279) & 63))}
    var v21 uint8 = 128
    var v22 int64 = 9223372036854775807
    for fi23 := 0; fi23 < 20; fi23++ {
        fmt.Println(fi23, v22, v3)
        var v24 int8 = -1
        var v25 int8 = (v24 | v24)
        _ = fi23
        _ = v24
        _ = v25
    }
    v22 = v22
    fmt.Println(v3)
    fmt.Println(v3, v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
