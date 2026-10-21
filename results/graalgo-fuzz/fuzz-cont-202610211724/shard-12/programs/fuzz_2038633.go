package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 65535
    var v3 uint64 = 0
    var v4 uint32 = v2
    var v5 float64 = -2.25
    v4 += (^(v2))
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := v1[("a" + "a")]
            _ = len(v1)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 uint64 = 547
    var v10 int16 = 1
    var v11 uint8 = 236
    v3 -= (v3 - (^(v3)))
    v12 := GenS1{F0: 1, F1: (v9 >> ((463) & 63))}
    fmt.Println(v2)
    _ = len(v1)
    switch 1 {
    case -1:
        if (42 < 244) {
            var v13 bool = (!true)
            fmt.Println(v13, v10)
            _ = v13
        }
    case 9223372036854775807:
        for fi14 := 0; fi14 < 19; fi14++ {
            if ((v2 * v4) >= v4) {
                delete(v1, ("hello" + ("\n\t" + "a")))
            }
            if true {
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 float64 = v5
                    v9 = (v9 % (((v3 << ((324) & 63))) | 1))
                    _ = fi15
                    _ = v16
                }
                fmt.Println(v5, v2, v3)
                if false {
                    v17 := GenS2{F0: (v4 - (^(v4)))}
                    _ = v17
                }
            }
            _ = fi14
        }
    case 128:
        v10 = v10
    }
    switch "foo bar" {
    case "":
        v9 = (v3 << ((265) & 63))
    case "hello":
        v2 = (v2 + v2)
    case "foo bar":
        var v18 uint32 = v2
        _ = v18
    default:
        v19 := GenS2{F0: v4}
        _ = v19
    }
    var v20 uint16 = 11
    var v21 bool = (-128 == -128)
    v22 := GenS1{F0: 32768, F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
