package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }

func main() {
    var v1 int32 = 613
    var v2 int64 = 850
    var v3 int8 = 50
    v4 := make(map[int]bool)
    var v5 int16 = -1
    delete(v4, 127)
    fmt.Println(v2, v3, v2)
    v6 := GenS1{F0: int(v5)}
    v7 := GenS1{F0: 32768}
    v3 = int8(v5)
    fmt.Println(v1, v5, v1)
    v1 += -99
    for fi8 := 0; fi8 < 13; fi8++ {
        v2 = -1
        fmt.Println(v5, v2)
        v9 := GenS1{F0: (fi8 + fi8)}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: -539}
    if true {
        fmt.Println(v1, v2, v2)
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            v4[0] = (128 >= 1)
            _ = fi11
        }
        if (false && false) {
            for fi12 := 0; fi12 < 2; fi12++ {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v14 := GenS1{F0: (fi13 ^ fi12)}
                    var v15 int = 9223372036854775807
                    v16 := GenS2{F0: (v3 < (v3 - v3)), F1: 42, F2: 255}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = fi12
            }
            var v17 int16 = v5
            v4[32768] = (!(true && true))
            _ = v17
        }
        switch 128 {
        case 9223372036854775807:
            var v18 float32 = ((1.5 * 0.1) * (0.0 + 0.1))
            _ = v18
        case 127:
            var v19 string = (("a" + "foo bar") + "foo bar")
            _ = v19
        }
    }
    v2 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
