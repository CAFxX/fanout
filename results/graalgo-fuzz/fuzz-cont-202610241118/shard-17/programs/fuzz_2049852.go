package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint32 }

func main() {
    var v1 uint = 1
    v2 := []string{"", ("foo bar" + "foo bar"), "\n\t", (("\n\t" + "a") + "")}
    var v3 string = "foo bar"
    v4 := []uint32{2, 160, 234}
    var v5 uint16 = 771
    v5 -= (^(v5))
    v5 = v5
    switch v3 {
    case "hello":
        if len(v2) > 0 {
            v2[0] = v3
        }
    case "\n\t":
        if ((1e10 / 1.5) != (0.0 - (1e10 * 3.14159))) {
            if (4294967295 > 4294967295) {
                var v6 uint16 = (v5 + v5)
                if (v3 > (v3 + ((v3 + "\n\t") + v3))) {
                    v2 = append(v2, v3)
                    v7 := GenS1{F0: 2}
                    _ = v7
                } else {
                    v5 = (v5 >> ((v1) & 63))
                    for fi8 := 0; fi8 < 6; fi8++ {
                        v1 = v1
                        var v9 int = (fi8 / ((fi8) | 1))
                        _ = fi8
                        _ = v9
                    }
                }
                _ = v6
            } else {
                var v10 bool = (!false)
                _ = v10
            }
        }
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v5 -= (v5 - v5)
    }
    if len(v4) > 0 {
        v4[0] = 2
    }
    v11 := GenS2{F0: 42, F1: 0, F2: 65535}
    v3 = ""
    for fi12 := 0; fi12 < 17; fi12++ {
        v5 -= 32767
        _ = fi12
    }
    var v13 int64 = 1
    v13 += v13
    for fi14 := 0; fi14 < 4; fi14++ {
        v1 = 4294967295
        var v15 float64 = 3.14159
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
