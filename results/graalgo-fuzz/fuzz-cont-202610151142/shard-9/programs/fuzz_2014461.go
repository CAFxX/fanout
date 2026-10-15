package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint16 }

func main() {
    var v1 uint16 = 977
    var v2 uint32 = 0
    var v3 string = "foo bar"
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    var v6 bool = true
    if v6 {
        fmt.Println(v6, v3)
        v6 = v6
        v3 = v3
    }
    fmt.Println(v3, v6, v1)
    var v7 int16 = 32767
    var v8 int8 = 42
    delete(v5, (v2 | v2))
    fmt.Println(v2, v3)
    for fi9 := 0; fi9 < 13; fi9++ {
        v1 = v1
        _ = fi9
    }
    fmt.Println(v7, v3)
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := GenS1{F0: 594, F1: 18446744073709551615, F2: (-(v8))}
            v8 = 70
            v2 -= 849
            if (false || true) {
                if false {
                    v1 = (^(v1))
                    v13 := GenS1{F0: (^(v2)), F1: 2, F2: (v8 << ((325) & 63))}
                    fmt.Println(fi11)
                    _ = v13
                } else {
                    fmt.Println(v6, v2, fi11)
                    if (18446744073709551615 != 1) {
                        switch fi10 {
                        case 32768:
                            v3 = ("a" + (v3 + ""))
                            v8 += (v8 + v8)
                        case 2147483647:
                            fmt.Println(v6, v3)
                        }
                    }
                }
            } else {
                v8 += (v8 * v8)
            }
            var v14 uint32 = v2
            fmt.Println(v6, v14)
            _ = fi11
            _ = v12
            _ = v14
        }
        _ = fi10
    }
    v15 := GenS2{F0: 42, F1: 128, F2: uint16(v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
