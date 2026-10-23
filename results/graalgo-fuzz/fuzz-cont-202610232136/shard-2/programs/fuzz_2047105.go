package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := []uint32{835}
    var v2 string = ""
    var v3 int = 32768
    v4 := make(map[uint32]int)
    v5 := make(map[int]int)
    v2 = v2
    if (100 < 0) {
        v2 = "a"
        v6 := GenS1{F0: (v3 * (v3 ^ v3)), F1: 750, F2: 42}
        _ = v6
    } else {
        fmt.Println(v2, v2)
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := v5[(fi7 >> ((1) & 63))]
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            var v10 uint8 = 2
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v2 = "\n\t"
    fmt.Println(v2)
    v1 = append(v1, 2147483647)
    v11 := GenS1{F0: (v3 / (((v3 % ((-588) | 1))) | 1)), F1: 42, F2: 2}
    v12 := GenS1{F0: v3, F1: 1, F2: 1}
    v2 = v2
    v2 = v2
    v3 = v3
    v2 = v2
    for fi13 := 0; fi13 < 2; fi13++ {
        if false {
            v3 += fi13
            switch 2147483647 {
            case 1:
                v5[v3] = (fi13 * (-(fi13)))
            case 2147483647:
                v3 -= (fi13 - (fi13 + fi13))
            case 343:
                var v14 int16 = -1
                _ = v14
            }
            var v15 float64 = (0.1 / (3.14159 + (float64(((float32(3.14159) * (1e10 + (float32(1e10) * -2.25))) / 1e10)) / -2.25)))
            _ = v15
        } else {
            v3 += 32767
            v16 := GenS1{F0: (v3 << ((919) & 63)), F1: 659, F2: 42}
            _ = v16
        }
        v17 := GenS2{F0: 32767}
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS2{F0: 2}
            v3 -= (v3 * 127)
            _ = fi18
            _ = v19
        }
        _ = fi13
        _ = v17
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        fmt.Println(v3, v3, v3)
        v3 -= v3
        _ = fi20
    }
    v3 += 128
    v21 := GenS1{F0: 4294967295, F1: 42, F2: 2}
    if true {
        var v22 string = v2
        var v23 string = v2
        if (65535 != 42) {
            if len(v1) > 0 {
                v1[0] = 820
            }
            var v24 uint8 = 255
            v25 := GenS1{F0: v3, F1: 18446744073709551615, F2: 73}
            _ = v24
            _ = v25
        } else {
            v23 = "hello"
        }
        _ = v22
        _ = v23
    } else {
        fmt.Println(v2, v3)
        v26 := GenS2{F0: 65535}
        _ = v26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 17; fi28++ {
            v2 = v2
            if (!(862 <= 2)) {
                fmt.Println(v2, v2, v3)
                switch (-(v3)) {
                case -1:
                    v2 = ("\n\t" + "foo bar")
                default:
                    var v29 float64 = 0.1
                    _ = v29
                }
                v3 += fi27
            } else {
                v3 += fi28
                v3 = 2147483647
            }
            delete(v4, 4294967295)
            _ = fi28
        }
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
