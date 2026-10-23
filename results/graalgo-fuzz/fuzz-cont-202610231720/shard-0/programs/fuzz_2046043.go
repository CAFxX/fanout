package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0() bool {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 824, F1: 178, F2: -128}
        var v9 uint32 = 735
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int = 32768
    fmt.Println(v10)
    var v11 float32 = 1.5
    v10 += v10
    _ = v10
    _ = v11
    return (("" <= (("a" + "hello") + "foo bar")) && (255 <= 2))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 223
    v3 := make(map[int]bool)
    var v4 string = ((("\n\t" + "a") + "foo bar") + "a")
    var v5 int = 0
    v6 := []int{(v5 & v5)}
    if len(v6) > 0 {
        v6[0] = (v5 - (v5 * v5))
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v1[1] = (!true)
    for fi12 := 0; fi12 < 13; fi12++ {
        switch v5 {
        case 127:
            for fi13 := 0; fi13 < 13; fi13++ {
                v14 := GenS2{F0: 2, F1: (v2 << ((18446744073709551615) & 63))}
                var v15 int32 = -2147483648
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            var v16 uint8 = 2
            _ = v16
        }
        var v17 uint = 2
        _ = fi12
        _ = v17
    }
    var v18 uint64 = 42
    var v19 int64 = v2
    v20 := GenS2{F0: 229, F1: 834}
    var v21 int = v5
    var v22 uint = 0
    v1[468] = (false || ((v21 << ((v22) & 63)) <= 564))
    var v23 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
