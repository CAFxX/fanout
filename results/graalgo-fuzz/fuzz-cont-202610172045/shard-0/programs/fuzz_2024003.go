package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 string; F1 string; F2 bool }

func main() {
    var v1 uint16 = 498
    var v2 int32 = -627
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    var v5 uint8 = 42
    v6 := []uint32{2, 0, 2147483647}
    var v7 uint16 = v1
    v8 := GenS1{F0: (^(v5)), F1: 1}
    var v9 uint32 = 4294967295
    if false {
        v10 := GenS1{F0: (v5 / ((v5) | 1)), F1: int16(v9)}
        v2 = (v2 ^ -2147483648)
        v9 -= v9
        _ = v10
    } else {
        for fi11 := 0; fi11 < 9; fi11++ {
            if (1 < 42) {
                var v12 uint32 = (v9 / (((^(v9))) | 1))
                v9 = v12
                _ = v12
            }
            var v13 float32 = (1e10 * 0.1)
            switch "\n\t" {
            case "hello":
                _ = len(v6)
            case "":
                var v14 string = (("foo bar" + "foo bar") + "foo bar")
                _ = v14
            default:
                var v15 bool = false
                _ = v15
            }
            _ = fi11
            _ = v13
        }
        v16 := GenS1{F0: uint8(v1), F1: 0}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v16
    }
    v7 = (v1 / ((255) | 1))
    v17 := GenS1{F0: v5, F1: -32768}
    for fi18 := 0; fi18 < 6; fi18++ {
        v7 -= v1
        v19 := GenS1{F0: 0, F1: 497}
        _ = fi18
        _ = v19
    }
    var v20 uint32 = (v9 + v9)
    v21 := GenS2{F0: ("\n\t" + ("\n\t" + "foo bar")), F1: "\n\t", F2: (!(!true))}
    fmt.Println(v7, v2, v1)
    delete(v3, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
