package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 int32 }

func main() {
    var v1 int16 = -1
    v2 := []string{("hello" + ("hello" + "\n\t")), ""}
    var v3 uint8 = 61
    var v4 uint16 = 733
    var v5 int64 = 1
    var v6 int = 1
    v6 = 940
    switch "a" {
    case "foo bar":
        var v7 uint32 = 916
        _ = v7
    case "hello":
        switch ((("hello" + "foo bar") + "hello") + "foo bar") {
        case "\n\t":
            v6 -= v6
        case "a":
            v1 = (v1 >> ((80) & 63))
        case "":
            for fi8 := 0; fi8 < 1; fi8++ {
                v5 += 460
                _ = fi8
            }
        default:
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint = 798
                for fi11 := 0; fi11 < 15; fi11++ {
                    fmt.Println(v5, v3, fi11)
                    fmt.Println(fi11, v6)
                    var v12 float64 = (0.1 + float64(((-2.25 / 1.0) + (0.0 - ((((-2.25 + 1.5) / 0.1) * 1e10) / 0.1)))))
                    _ = fi11
                    _ = v12
                }
                _ = fi9
                _ = v10
            }
        }
    }
    var v13 bool = false
    v5 = (v5 + (v5 * -9223372036854775808))
    var v14 int16 = (^(v1))
    v2 = append(v2, "\n\t")
    v2 = append(v2, "")
    fmt.Println(v6)
    _ = len(v2)
    v15 := GenS1{F0: -1, F1: 2, F2: 65535}
    v16 := GenS2{F0: v4, F1: 1, F2: -977}
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v6 = v6
    var v17 int64 = (v5 * v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
