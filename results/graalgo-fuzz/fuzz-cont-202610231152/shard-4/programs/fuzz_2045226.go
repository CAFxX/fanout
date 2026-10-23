package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0(p0 int16, p1 int8) int32 {
    var v5 uint8 = 70
    v6 := GenS1{F0: -712}
    p1 -= (p1 % ((-128) | 1))
    p1 = p1
    v5 = (^(v5))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 883
}
func genFn1(p0 uint8, p1 uint64, p2 int16) string {
    v7 := GenS2{F0: 65535, F1: 2147483648}
    v8 := GenS1{F0: -480}
    var v9 uint16 = 65535
    fmt.Println(p2, p1, p0)
    v9 += v9
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return ("a" + (("foo bar" + "a") + ""))
}

func main() {
    var v1 uint8 = 172
    var v2 uint8 = (v1 - (v1 / ((v1) | 1)))
    var v3 uint32 = 42
    var v4 uint = 1
    v3 -= (v3 / ((v3) | 1))
    var v10 int = -633
    v11 := []int{}
    v12 := GenS2{F0: (v4 / (((v4 ^ (v4 + (v4 | uint(v1))))) | 1)), F1: 1}
    for fi13 := 0; fi13 < 14; fi13++ {
        switch "a" {
        case "a":
            if false {
                v1 += 255
                v14 := GenS1{F0: -1}
                v2 = (v2 * 77)
                _ = v14
            }
        case "\n\t":
            for fi15 := 0; fi15 < 19; fi15++ {
                v10 += (fi13 << (((v4 - (v4 + v4))) & 63))
                _ = fi15
            }
        default:
            if ((true && (false || true)) && false) {
                v3 = (^(v3))
                var v16 int = 127
                for fi17 := 0; fi17 < 11; fi17++ {
                    var v18 int16 = 922
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
        }
        _ = fi13
    }
    var v19 int8 = -92
    var v20 float64 = 0.1
    if len(v11) > 0 {
        v11[0] = 1
    }
    _ = len(v11)
    var v21 bool = true
    if len(v11) > 0 {
        v11[0] = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
