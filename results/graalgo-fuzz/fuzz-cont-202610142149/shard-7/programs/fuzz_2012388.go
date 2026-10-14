package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0(p0 uint64, p1 uint8, p2 uint) string {
    v7 := GenS1{F0: -810}
    var v8 int = 42
    fmt.Println(v8, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 int32 = 0
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    var v4 uint = 18446744073709551615
    var v5 uint = 42
    var v6 uint16 = 42
    var v9 uint8 = 255
    v10 := v2["hello"]
    if (!(v6 < (v6 / ((v6) | 1)))) {
        for fi11 := 0; fi11 < 19; fi11++ {
            v10 = (763 != v5)
            var v12 uint8 = (v9 ^ 42)
            v13 := GenS2{F0: -187, F1: -36}
            _ = fi11
            _ = v12
            _ = v13
        }
        v5 += (v4 & 2)
    } else {
        if (32767 >= -2147483648) {
            v4 = (v4 >> ((847) & 63))
            var v14 uint64 = 1
            _ = v14
        } else {
            v15 := GenS3{F0: v6, F1: -9223372036854775808, F2: -112}
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 bool = (("" + (("\n\t" + ("\n\t" + "")) + "a")) == "")
                v18 := GenS3{F0: (^(v6)), F1: -91, F2: 100}
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        }
    }
    var v19 uint64 = 530
    fmt.Println(v5)
    var v20 float32 = (1e10 - 0.0)
    var v21 int16 = -32768
    switch "a" {
    case "\n\t":
        v20 = v20
    default:
        switch "" {
        case "\n\t":
            v19 = v19
        default:
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS1{F0: (v21 * int16(v9))}
                v24 := []int{(fi22 | -640), (fi22 >> ((v5) & 63))}
                var v25 uint16 = (v6 / ((668) | 1))
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
        }
    }
    v4 -= (v5 >> ((255) & 63))
    v10 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
