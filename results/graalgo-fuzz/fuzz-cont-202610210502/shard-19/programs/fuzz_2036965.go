package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int32; F1 uint32; F2 int }

func main() {
    var v1 uint16 = 42
    var v2 uint8 = 127
    v3 := []string{}
    var v4 int64 = -153
    v5 := make(map[uint32]bool)
    var v6 int8 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (127 <= 42) {
        if (uint64(v1) < 1) {
            _ = len(v3)
            for fi7 := 0; fi7 < 19; fi7++ {
                if ((!(((0.0 + ((3.14159 * 0.1) - 1e10)) * -2.25) == 0.0)) || (!true)) {
                    var v8 int = int(v1)
                    _ = v8
                }
                var v9 int32 = 42
                var v10 uint = 18446744073709551615
                _ = fi7
                _ = v10
                _ = v9
            }
        } else {
            v2 -= (v2 ^ 107)
            v4 += (v4 + v4)
            for fi11 := 0; fi11 < 13; fi11++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi11
            }
        }
        v12 := GenS1{F0: 0, F1: 986}
        v2 -= v2
        _ = v12
    }
    v4 += v4
    v2 -= (v2 % ((81) | 1))
    var v13 string = "\n\t"
    if false {
        v14 := GenS1{F0: (v6 - (v6 ^ (v6 % ((0) | 1)))), F1: 0}
        switch 174 {
        case 2147483647:
            var v15 int = -553
            _ = v15
        default:
            if true {
                fmt.Println(v2)
                v16 := GenS2{F0: 2147483647, F1: 2147483648, F2: 9223372036854775807}
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            if (394 == 0) {
                for fi18 := 0; fi18 < 3; fi18++ {
                    var v19 string = v13
                    _ = fi18
                    _ = v19
                }
                var v20 float64 = (-2.25 - 1.5)
                _ = v20
            } else {
                var v21 bool = false
                var v22 int64 = (v4 << ((769) & 63))
                _ = v21
                _ = v22
            }
            var v23 string = (v13 + v13)
            fmt.Println(v2)
            _ = fi17
            _ = v23
        }
        _ = v14
    }
    var v24 uint64 = 0
    if true {
        var v25 uint8 = 42
        var v26 uint32 = 939
        v6 += v6
        _ = v25
        _ = v26
    }
    v5[1] = (v1 <= v1)
    var v27 uint = 42
    v28 := GenS1{F0: (v6 | v6), F1: v27}
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
