package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint64 }
type GenS2 struct { F0 bool }
func genFn0(p0 int16, p1 int16, p2 int16) bool {
    var v5 uint32 = 269
    v6 := GenS1{F0: 402, F1: -1, F2: 433}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ((p1 & 32767) < (^(p1)))
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{79}
    var v3 float32 = ((-2.25 + 1.5) + float32(1e10))
    v4 := []string{"\n\t", "\n\t"}
    fmt.Println(v1)
    if false {
        var v7 uint16 = 0
        var v8 int32 = v1
        if (!false) {
            var v9 uint8 = 3
            v4 = append(v4, "foo bar")
            _ = v9
        } else {
            v1 -= (v8 % ((v1) | 1))
            if false {
                if len(v2) > 0 {
                    v2[0] = 996
                }
            }
            if (!(((42 == 65535) && true) && (v8 >= -152))) {
                fmt.Println(v8, v3, v3)
                v4 = append(v4, "foo bar")
                for fi10 := 0; fi10 < 4; fi10++ {
                    v7 -= v7
                    v7 += (v7 & v7)
                    if len(v4) > 0 {
                        v4[0] = ((("a" + "\n\t") + ("\n\t" + "a")) + "")
                    }
                    _ = fi10
                }
            }
        }
        _ = v7
        _ = v8
    } else {
        var v11 float64 = float64(v3)
        _ = v11
    }
    _ = len(v2)
    var v12 string = "a"
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS1{F0: 572, F1: (fi13 + (fi13 & (fi13 | fi13))), F2: 2}
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 65535
    v12 = (("" + v12) + "foo bar")
    v12 = v12
    fmt.Println(v1, v12)
    var v16 uint16 = (v15 + 42)
    switch 4294967295 {
    case 2147483647:
        v17 := GenS2{F0: (((true || ((0.1 - (3.14159 * -2.25)) != 1.5)) || false) || (18446744073709551615 == 42))}
        _ = v17
    case -2147483648:
        fmt.Println(v15, v3)
    default:
        if len(v2) > 0 {
            v2[0] = 65535
        }
    }
    v18 := GenS2{F0: false}
    for fi19 := 0; fi19 < 13; fi19++ {
        v3 = v3
        _ = fi19
    }
    fmt.Println(v15, v3, v1)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
