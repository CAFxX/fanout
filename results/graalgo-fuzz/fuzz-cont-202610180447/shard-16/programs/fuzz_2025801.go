package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8; F1 int32; F2 uint8 }

func main() {
    var v1 uint8 = 153
    var v2 bool = true
    v3 := []string{}
    var v4 uint = 18446744073709551615
    _ = len(v3)
    var v5 uint32 = 4294967295
    fmt.Println(v2)
    v5 -= (v5 / (((v5 | (^(v5)))) | 1))
    if v2 {
        if v2 {
            v2 = v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if v2 {
                v3 = append(v3, "")
            } else {
                fmt.Println(v1, v2)
                _ = len(v3)
            }
        } else {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            switch -747 {
            case 423:
                for fi6 := 0; fi6 < 15; fi6++ {
                    var v7 uint32 = v5
                    var v8 uint16 = 32767
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            case -34:
                v1 = 255
            }
        }
        v4 -= 1
        if (347 == 150) {
            v9 := GenS1{F0: -2147483648, F1: (v4 / ((v4) | 1)), F2: 884}
            fmt.Println(v2)
            v10 := GenS1{F0: -1, F1: v4, F2: 255}
            _ = v10
            _ = v9
        }
    } else {
        v5 -= v5
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    }
    if true {
        var v11 int8 = -103
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 uint8 = (v1 << ((v4) & 63))
            v13 -= (v13 - v1)
            var v14 float64 = 0.0
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    } else {
        if (!v2) {
            v3 = append(v3, (("hello" + "\n\t") + "\n\t"))
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 uint = 1
                _ = fi15
                _ = v16
            }
        }
        v17 := GenS1{F0: 964, F1: v4, F2: 0}
        switch 9223372036854775807 {
        case 0:
            v5 -= (v5 - (^(v5)))
        case 42:
            v18 := GenS1{F0: -2147483648, F1: (v4 ^ uint(v1)), F2: 32767}
            _ = v18
        }
        _ = v17
    }
    v19 := GenS1{F0: 2147483647, F1: 42, F2: uint16(v1)}
    v20 := GenS1{F0: 938, F1: v4, F2: 608}
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS2{F0: -50, F1: -2147483648, F2: (v1 % (((v1 | uint8(v5))) | 1))}
        _ = fi21
        _ = v22
    }
    switch 2147483647 {
    case 32767:
        for fi23 := 0; fi23 < 15; fi23++ {
            _ = len(v3)
            var v24 uint16 = 1
            _ = fi23
            _ = v24
        }
    }
    v25 := GenS2{F0: 100, F1: 105, F2: (v1 + (v1 ^ (v1 ^ 166)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
