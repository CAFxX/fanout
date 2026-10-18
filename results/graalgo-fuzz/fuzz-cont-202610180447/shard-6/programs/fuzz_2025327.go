package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 int64) string {
    fmt.Println(p0)
    switch "hello" {
    case "hello":
        var v7 int16 = 32767
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{fi9, (fi9 >> ((694) & 63)), fi8, fi9}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 float64 = -2.25
    _ = p0
    _ = v11
    return ("hello" + "hello")
}
func genFn1(p0 int16) int16 {
    var v12 bool = true
    p0 = p0
    switch -843 {
    case -1:
        fmt.Println(p0)
    }
    var v13 uint16 = 0
    _ = p0
    _ = v12
    _ = v13
    return (p0 * p0)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 712
    var v3 int32 = -2147483648
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]int)
    var v6 uint16 = 2
    var v14 uint16 = v6
    v6 -= v6
    v6 -= v14
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS1{F0: (v2 ^ v2)}
        _ = fi15
        _ = v16
    }
    var v17 int = 127
    switch (-(v17)) {
    case 314:
        for fi18 := 0; fi18 < 18; fi18++ {
            if ((((float32(float64(0.1)) - (0.0 - (0.0 * (((1.5 - 1.5) / 1.0) / float32(0.1))))) != float32((float64(((1e10 + 1e10) * float32(float64(1e10)))) - 0.0))) || (!((!true) || false))) || (468 <= -17)) {
                fmt.Println(fi18, fi18, v3)
                v19 := GenS1{F0: v2}
                _ = v19
            }
            _ = fi18
        }
    case 1:
        v14 = (v14 - (v6 >> ((42) & 63)))
    }
    var v20 int64 = 42
    var v21 string = "a"
    var v22 uint16 = (v6 - v6)
    switch v21 {
    case "foo bar":
        v23 := GenS2{F0: ((v21 + "\n\t") + "foo bar"), F1: 492}
        _ = v23
    case "hello":
        var v24 uint8 = 0
        _ = v24
    case "\n\t":
        var v25 int8 = 1
        _ = v25
    default:
        fmt.Println(v3, v17)
    }
    var v26 uint32 = 711
    for fi27 := 0; fi27 < 1; fi27++ {
        v20 = (v20 ^ v20)
        if (!false) {
            v28 := GenS1{F0: genFn1(genFn1((v2 << ((970) & 63))))}
            _ = v28
        }
        _ = fi27
    }
    v20 += int64(v14)
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
