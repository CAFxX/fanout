package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint16; F1 bool }

func main() {
    var v1 string = ("foo bar" + ("a" + ("" + ("\n\t" + "hello"))))
    var v2 uint = 1
    v3 := make(map[string]uint32)
    var v4 int8 = 0
    for fi5 := 0; fi5 < 16; fi5++ {
        v4 = v4
        _ = fi5
    }
    v4 = -47
    var v6 string = (v1 + "\n\t")
    if false {
        switch ("\n\t" + v6) {
        case "hello":
            for fi7 := 0; fi7 < 18; fi7++ {
                var v8 int = (fi7 >> (((v2 - (v2 << ((42) & 63)))) & 63))
                _ = fi7
                _ = v8
            }
        default:
            fmt.Println(v6, v6, v4)
        }
        v9 := []int{-952, 4294967295, 9223372036854775807, -685}
        v10 := GenS1{F0: 42, F1: 127, F2: 0}
        _ = v10
        _ = v9
    }
    v2 -= v2
    v6 = "\n\t"
    v11 := []int{531, -2147483648, -178, 32767}
    if false {
        v12 := GenS1{F0: 1, F1: 0, F2: 32767}
        _ = v12
    }
    v13 := GenS2{F0: 2, F1: (true || ((false && false) || true))}
    switch 42 {
    case 32768:
        for fi14 := 0; fi14 < 18; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                v2 = 1
                switch 32768 {
                case 128:
                    v2 = v2
                case -2147483648:
                    switch (fi15 << (((v2 * (v2 - (v2 / ((2) | 1))))) & 63)) {
                    case 363:
                        v6 = "hello"
                    default:
                        v1 = v6
                    }
                case 32767:
                    v16 := GenS2{F0: 65535, F1: (false && (true || (0 < 127)))}
                    _ = v16
                default:
                    v2 -= v2
                }
                if (v2 == v2) {
                    fmt.Println(v6, v4, v6)
                    if len(v11) > 0 {
                        v11[0] = -1
                    }
                }
                _ = fi15
            }
            _ = fi14
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
