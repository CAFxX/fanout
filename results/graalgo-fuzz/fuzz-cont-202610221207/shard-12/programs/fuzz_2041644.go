package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32) int16 {
    switch 2147483647 {
    case -9223372036854775808:
        for fi5 := 0; fi5 < 10; fi5++ {
            if true {
                for fi6 := 0; fi6 < 10; fi6++ {
                    var v7 uint8 = 0
                    var v8 int32 = (p0 | p0)
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                v9 := GenS1{F0: p0}
                var v10 string = ""
                _ = v10
                _ = v9
            }
            _ = fi5
        }
    case -1:
        p0 = (p0 + p0)
    case 49:
        for fi11 := 0; fi11 < 3; fi11++ {
            p0 = p0
            p0 -= (p0 * -2147483648)
            v12 := []int{2147483647, fi11, fi11, fi11}
            _ = fi11
            _ = v12
        }
    }
    var v13 float32 = 1.5
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS2{F0: 42}
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = v13
    return 0
}

func main() {
    var v1 int32 = 42
    var v2 int64 = 42
    var v3 uint = 42
    v4 := make(map[uint32]bool)
    var v16 float64 = (3.14159 + 1.5)
    var v17 uint32 = 4294967295
    for fi18 := 0; fi18 < 1; fi18++ {
        if true {
            switch "foo bar" {
            case "foo bar":
                v19 := GenS1{F0: v1}
                _ = v19
            case "a":
                for fi20 := 0; fi20 < 11; fi20++ {
                    switch (fi18 - (fi18 >> ((v3) & 63))) {
                    case 95:
                        var v21 float64 = v16
                        _ = v21
                    case 32768:
                        var v22 int8 = 100
                        _ = v22
                    }
                    v23 := v4[42]
                    var v24 float32 = -2.25
                    _ = fi20
                    _ = v23
                    _ = v24
                }
            case "\n\t":
                for fi25 := 0; fi25 < 8; fi25++ {
                    v26 := GenS2{F0: 1}
                    _ = fi25
                    _ = v26
                }
            }
        }
        _ = fi18
    }
    v27 := []int{128, 128, -9223372036854775808, 42}
    v16 -= ((float64(0.1) * -2.25) * v16)
    if ((v16 <= float64((float32((v16 + v16)) + -2.25))) && false) {
        var v28 int32 = (v1 & (v1 / (((v1 >> (((v3 * (v3 - 2))) & 63))) | 1)))
        fmt.Println(v28, v16)
        _ = v28
    } else {
        v29 := v4[0]
        var v30 uint8 = 128
        _ = v29
        _ = v30
    }
    switch "" {
    case "":
        var v31 uint = v3
        _ = v31
    default:
        v16 += v16
    }
    var v32 uint = (v3 | (v3 / ((v3) | 1)))
    if false {
        fmt.Println(v16)
        var v33 uint32 = v17
        v27 = append(v27, 2147483647)
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
}
