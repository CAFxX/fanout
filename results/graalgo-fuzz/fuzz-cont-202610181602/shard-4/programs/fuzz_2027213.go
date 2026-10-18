package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint32 }

func main() {
    var v1 uint16 = 42
    var v2 uint = 1
    var v3 int16 = -32768
    fmt.Println(v2, v2, v3)
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        v1 -= (v1 << (((^(v2))) & 63))
    default:
        v2 += v2
    }
    v4 := GenS1{F0: 0}
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 float64 = 0.0
        fmt.Println(fi5)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int64 = int64(v2)
        var v9 uint64 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 string = "\n\t"
            v2 = 65535
            _ = fi11
            _ = v12
        }
        v13 := []int{(fi10 * fi10), fi10}
        v2 -= 65535
        _ = fi10
        _ = v13
    }
    v14 := GenS2{F0: -665, F1: 2147483647}
    var v15 int32 = 42
    var v16 int16 = v3
    if (1 != 2) {
        var v17 uint8 = 2
        if (!((("foo bar" + ("hello" + "")) + ("a" + "")) == ((("foo bar" + "a") + "") + "\n\t"))) {
            for fi18 := 0; fi18 < 4; fi18++ {
                if true {
                    var v19 int8 = 1
                    var v20 int64 = -1
                    _ = v19
                    _ = v20
                }
                v15 = -2147483648
                _ = fi18
            }
            v3 += (v16 + (v3 | v3))
            var v21 uint16 = v1
            _ = v21
        } else {
            v17 += v17
            v22 := []int{-1, int(v16), -2147483648, 2147483647}
            _ = v22
        }
        var v23 int8 = 68
        _ = v17
        _ = v23
    }
    var v24 int32 = v15
    v15 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
