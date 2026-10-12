package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -1
    var v3 int8 = -27
    v4 := []int{32768}
    var v5 uint8 = 2
    var v6 int32 = v2
    v6 = v2
    for fi7 := 0; fi7 < 12; fi7++ {
        if (!((false && false) || ((v5 * (v5 ^ (v5 + v5))) <= (v5 ^ 127)))) {
            var v8 int = (-(fi7))
            _ = v8
        } else {
            fmt.Println(v6, fi7, fi7)
            v4 = append(v4, (^(fi7)))
            var v9 int = fi7
            _ = v9
        }
        if true {
            var v10 int64 = -1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v11 uint32 = 0
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            switch (-(fi7)) {
            case 1:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 128:
                v2 = (^(v6))
            default:
                var v13 string = "foo bar"
                _ = v13
            }
            if false {
                v2 += v2
                if true {
                    v14 := GenS1{F0: 675, F1: 1, F2: v6}
                    _ = v14
                }
            } else {
                if false {
                    v6 = 166
                } else {
                    v3 -= v3
                }
            }
            _ = fi12
        }
        _ = fi7
    }
    var v15 uint32 = 2
    var v16 int8 = (-(v3))
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS1{F0: 1, F1: 4294967295, F2: v2}
        fmt.Println(v5, v6, v3)
        var v19 int16 = int16(v2)
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
