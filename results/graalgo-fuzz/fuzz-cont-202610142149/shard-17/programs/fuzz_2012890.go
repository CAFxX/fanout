package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 float32 = -2.25
    var v2 bool = false
    v3 := make(map[int]bool)
    v4 := []string{(("foo bar" + ((("foo bar" + ("a" + "hello")) + "\n\t") + ("\n\t" + "a"))) + (("\n\t" + "\n\t") + "\n\t")), (("\n\t" + "foo bar") + "foo bar"), (((((("foo bar" + "foo bar") + ("" + "hello")) + "foo bar") + "hello") + "a") + "\n\t"), (("" + "hello") + "foo bar")}
    var v5 uint8 = 0
    v6 := make(map[uint32]int)
    v5 = 127
    v1 = v1
    v2 = ((v1 <= float32(0.1)) && v2)
    _ = len(v4)
    if ("a" <= "") {
        fmt.Println(v1, v1, v5)
    } else {
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                v5 += 128
                v1 = v1
                _ = fi8
            }
            _ = fi7
        }
    }
    v9 := GenS1{F0: ((32767 == 0) && (v5 > (^(v5))))}
    switch -97 {
    case -1:
        v10 := GenS1{F0: (!(!(307 == 18446744073709551615)))}
        _ = v10
    case 32767:
        switch "\n\t" {
        case "":
            for fi11 := 0; fi11 < 6; fi11++ {
                for fi12 := 0; fi12 < 6; fi12++ {
                    if (0 == 1) {
                        var v13 uint8 = (v5 ^ v5)
                        _ = v13
                    }
                    _ = fi12
                }
                v14 := GenS2{F0: 257, F1: 18446744073709551615}
                v1 = v1
                _ = fi11
                _ = v14
            }
        default:
            fmt.Println(v2, v1)
        }
    case 285:
        v1 = v1
    default:
        if ((true || v2) && v2) {
            if v2 {
                var v15 int64 = 1
                _ = v15
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            var v16 uint16 = 2
            fmt.Println(v16)
            _ = v16
        } else {
            fmt.Println(v1)
            v6[939] = -1
            if len(v4) > 0 {
                v4[0] = ("" + "a")
            }
        }
    }
    var v17 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
