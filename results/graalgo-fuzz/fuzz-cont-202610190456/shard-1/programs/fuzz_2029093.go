package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
func genFn0(p0 int64) int {
    var v7 uint = 18446744073709551615
    fmt.Println(v7, v7)
    v8 := []int{}
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []int{-537, 2147483647}
    var v2 int16 = 159
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    var v5 string = "\n\t"
    var v6 bool = (!(!false))
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 100, F1: 42, F2: 0}
        fmt.Println(v5)
        if false {
            v2 -= -32768
            for fi11 := 0; fi11 < 19; fi11++ {
                v5 = ((v5 + ("a" + "foo bar")) + v5)
                v2 = v2
                switch fi9 {
                case 127:
                    v12 := v3[2147483647]
                    _ = v12
                default:
                    var v13 int32 = 591
                    _ = v13
                }
                _ = fi11
            }
        } else {
            v14 := GenS1{F0: 0, F1: 1, F2: 65535}
            fmt.Println(v6, v6, v6)
            if (18446744073709551615 == 0) {
                var v15 int8 = 42
                _ = v15
            }
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    var v16 uint = 2
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 uint8 = 42
            fmt.Println(v6, v2)
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 bool = true
        _ = fi20
        _ = v21
    }
    v6 = true
    v5 = ("foo bar" + "a")
    v22 := GenS1{F0: -88, F1: 255, F2: 849}
    var v23 uint = (v16 % (((v16 * (v16 * (v16 % (((v16 << ((0) & 63))) | 1))))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
