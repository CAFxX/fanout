package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0(p0 uint) uint8 {
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 += p0
        v7 := GenS1{F0: 42, F1: (p0 * (p0 + (p0 % (((p0 << ((p0) & 63))) | 1))))}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return 128
}

func main() {
    var v1 float32 = (float32(0.1) - 1.5)
    var v2 uint64 = 0
    var v3 int8 = 127
    v4 := []string{"", "hello", "a"}
    var v5 int8 = v3
    for fi8 := 0; fi8 < 10; fi8++ {
        if ((!(-1 <= 1)) || false) {
            v9 := GenS2{F0: (fi8 & (fi8 - -554)), F1: 1}
            _ = v9
        }
        _ = fi8
    }
    var v10 float64 = 1e10
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = (v2 << ((89) & 63))
    switch "hello" {
    case "":
        switch 127 {
        case -376:
            fmt.Println(v1, v2)
        case 9223372036854775807:
            switch 4294967295 {
            case 0:
                if len(v4) > 0 {
                    v4[0] = ("\n\t" + "\n\t")
                }
            case -402:
                var v11 bool = false
                _ = v11
            case -741:
                for fi12 := 0; fi12 < 15; fi12++ {
                    fmt.Println(v10, v10, v1)
                    _ = fi12
                }
            }
        }
    case "hello":
        if true {
            fmt.Println(v2, v3)
            v5 = v3
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                v1 += 0.0
                for fi14 := 0; fi14 < 13; fi14++ {
                    v3 -= v3
                    _ = fi14
                }
                _ = fi13
            }
            fmt.Println(v2, v2, v5)
            switch (("foo bar" + "foo bar") + "hello") {
            case "\n\t":
                if true {
                    _ = len(v4)
                    switch -189 {
                    case 2147483647:
                        for fi15 := 0; fi15 < 8; fi15++ {
                            _ = len(v4)
                            _ = fi15
                        }
                    default:
                        v3 = 29
                    }
                }
            default:
                if false {
                    if true {
                        v1 = (float32((v10 * v10)) / v1)
                    }
                    v4 = append(v4, ("hello" + "foo bar"))
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            }
        }
    default:
        var v16 int8 = (v5 / (((v3 << ((4294967295) & 63))) | 1))
        _ = v16
    }
    var v17 uint16 = 65535
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v18 float32 = v1
    var v19 int = 868
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
