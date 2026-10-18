package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int16 }
func genFn0(p0 string) int64 {
    fmt.Println(p0)
    var v4 uint16 = 65535
    if true {
        v4 = 2
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 int = (fi5 ^ fi5)
            if ((1.5 + 0.1) != 1e10) {
                var v7 uint = 18446744073709551615
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 uint8 = 65
                    _ = fi8
                    _ = v9
                }
                switch (p0 + "\n\t") {
                case "foo bar":
                    fmt.Println(v7, v6)
                default:
                    for fi10 := 0; fi10 < 8; fi10++ {
                        v11 := GenS1{F0: 65535}
                        var v12 uint8 = uint8(v7)
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
                _ = v7
            } else {
                for fi13 := 0; fi13 < 4; fi13++ {
                    v6 = (v6 << ((65535) & 63))
                    _ = fi13
                }
            }
            _ = fi5
            _ = v6
        }
    }
    _ = p0
    _ = v4
    return int64(v4)
}
func genFn1(p0 uint) bool {
    if (false && (2 != p0)) {
        var v14 uint64 = 0
        for fi15 := 0; fi15 < 4; fi15++ {
            v14 += 946
            var v16 uint8 = 255
            _ = fi15
            _ = v16
        }
        v14 = (^(v14))
        _ = v14
    } else {
        p0 = 2
        p0 = (p0 | (p0 - p0))
    }
    p0 -= (p0 * 4294967295)
    _ = p0
    return (-9223372036854775808 > -66)
}

func main() {
    var v1 uint16 = 1
    var v2 string = "\n\t"
    v3 := make(map[int]uint32)
    var v17 uint64 = 1
    var v18 int64 = -401
    var v19 int32 = 1
    fmt.Println(v2, v18, v2)
    _ = len(v3)
    v17 -= 977
    switch 32767 {
    case 2147483647:
        if false {
            var v20 int16 = 1
            var v21 uint64 = (v17 % ((42) | 1))
            _ = v20
            _ = v21
        }
    }
    v22 := GenS2{F0: v1, F1: v1, F2: 729}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
}
