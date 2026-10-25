package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    var v4 string = "foo bar"
    var v5 int32 = 2147483647
    v6 := []int{}
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 65535, F1: 65535}
        v9 := GenS1{F0: 649, F1: 0}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v5, v1)
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int = (fi10 ^ fi10)
            _ = fi11
            _ = v12
        }
        v1 += 4294967295
        _ = fi10
    }
    var v13 int = 548
    v5 = v5
    var v14 uint16 = 32767
    switch int(v14) {
    case 898:
        v1 = v1
    }
    var v15 uint64 = (^(v1))
    var v16 uint8 = 1
    if (-678 != 0) {
        v17 := GenS2{F0: (v5 ^ (v5 - v5)), F1: 18446744073709551615}
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 float64 = 0.1
            if false {
                var v20 uint32 = 42
                for fi21 := 0; fi21 < 2; fi21++ {
                    v15 += v1
                    var v22 uint16 = (v14 | 42)
                    _ = fi21
                    _ = v22
                }
                _ = len(v2)
                _ = v20
            }
            var v23 int8 = -128
            _ = fi18
            _ = v19
            _ = v23
        }
        _ = v17
    }
    fmt.Println(v4, v13, v15)
    _ = len(v2)
    v24 := v3[65535]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
