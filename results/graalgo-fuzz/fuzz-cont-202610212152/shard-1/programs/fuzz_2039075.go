package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }
type GenS2 struct { F0 uint8; F1 int; F2 bool }

func main() {
    var v1 uint32 = 2147483648
    var v2 int = 32768
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 193
    var v5 bool = (!(65535 < 65535))
    v6 := []int{v2, (v2 + (v2 | -453)), 677}
    v7 := v3[(v1 | (^(v1)))]
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: -1, F1: ("foo bar" + "foo bar"), F2: 629}
        _ = len(v3)
        _ = fi8
        _ = v9
    }
    fmt.Println(v4, v5)
    fmt.Println(v2)
    v10 := GenS2{F0: 255, F1: (v2 ^ v2), F2: (!false)}
    if len(v6) > 0 {
        v6[0] = -313
    }
    v11 := GenS1{F0: 32767, F1: "a", F2: 2147483647}
    fmt.Println(v4, v7)
    if (v5 && v5) {
        v1 = 65535
        switch v2 {
        case 127:
            _ = len(v6)
        case 32768:
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 float32 = -2.25
                for fi14 := 0; fi14 < 2; fi14++ {
                    if len(v6) > 0 {
                        v6[0] = fi12
                    }
                    _ = fi14
                }
                for fi15 := 0; fi15 < 16; fi15++ {
                    v16 := GenS1{F0: 42, F1: ("hello" + ("hello" + "foo bar")), F2: -281}
                    _ = fi15
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
        default:
            var v17 uint = 374
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
