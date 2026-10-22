package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }

func main() {
    var v1 uint64 = 42
    v2 := []uint32{uint32(v1), 4294967295, 334, 646}
    var v3 int64 = int64(v1)
    fmt.Println(v3, v1)
    var v4 string = ("\n\t" + "hello")
    fmt.Println(v3)
    v3 += (v3 % ((v3) | 1))
    var v5 int16 = -945
    fmt.Println(v4, v1, v5)
    v5 += int16(v3)
    if false {
        _ = len(v2)
        switch v4 {
        case "\n\t":
            if (0 > 371) {
                for fi6 := 0; fi6 < 19; fi6++ {
                    v7 := GenS1{F0: 4294967295, F1: (v1 | v1)}
                    _ = fi6
                    _ = v7
                }
                v1 = v1
                var v8 int32 = 595
                _ = v8
            } else {
                var v9 int = 0
                v1 += 479
                for fi10 := 0; fi10 < 14; fi10++ {
                    v1 -= 1
                    fmt.Println(v3, v9)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi10
                }
                _ = v9
            }
        case "a":
            v11 := GenS2{F0: -1, F1: 42, F2: 139}
            _ = v11
        case "foo bar":
            switch 1 {
            case 127:
                switch 305 {
                case -417:
                    var v12 uint64 = (v1 / ((v1) | 1))
                    _ = v12
                default:
                    v1 += (^(v1))
                }
            case 55:
                v13 := GenS1{F0: 659, F1: (^(v1))}
                _ = v13
            case 32768:
                v5 = (-(v5))
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
